// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package configuration

import (
	"fmt"
	"github.com/haproxytech/client-native/v5/misc"
	"github.com/haproxytech/client-native/v5/models"
	"log"
)

// ServiceGrowthTypeLinear indicates linear growth type in ScalingParams.
const ServiceGrowthTypeLinear = "linear"

// ServiceGrowthTypeExponential indicates exponential growth type in ScalingParams.
const ServiceGrowthTypeExponential = "exponential"

// ServiceServer contains information for one server in the service.
type ServiceServer struct {
	Address string
	Port    int
	Weight  *int64 // Optional weight field
	Backup  string
}

type serviceNode struct {
	address  string
	name     string
	port     int64
	disabled bool
	modified bool
	weight   *int64 // Optional weight field
	backup   string
}

// Service represents the mapping from a discovery service into a configuration backend.
type Service struct {
	client        Configuration
	name          string
	nodes         []*serviceNode
	usedNames     map[string]struct{}
	transactionID string
	scaling       ScalingParams
}

type ServiceI interface {
	NewService(name string, scaling ScalingParams) (*Service, error)
	DeleteService(name string)
}

// ScalingParams defines parameter for dynamic server scaling of the Service backend.
type ScalingParams struct {
	SlotsGrowthType string
	BaseSlots       int
	SlotsIncrement  int
}

// NewService creates and returns a new Service instance.
// name indicates the name of the service and only one Service instance with the given name can be created.
func (c *client) NewService(name string, scaling ScalingParams) (*Service, error) {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if _, ok := c.services[name]; ok {
		return nil, fmt.Errorf("service with name %s already exists", name)
	}
	service := &Service{
		client:    c,
		name:      name,
		nodes:     make([]*serviceNode, 0),
		usedNames: make(map[string]struct{}),
		scaling:   scaling,
	}
	c.services[name] = service
	return service, nil
}

// DeleteService removes the Service instance specified by name from the client.
func (c *client) DeleteService(name string) {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	delete(c.services, name)
}

// Delete removes the service from the client with all the associated configuration resources.
func (s *Service) Delete() error {
	err := s.client.DeleteBackend(s.name, s.transactionID, 0)
	if err != nil {
		return err
	}
	s.client.DeleteService(s.name)
	return nil
}

// Init initiates the client by reading the configuration associated with it or created the initial configuration if it does not exist.
func (s *Service) Init(transactionID string, from string) (bool, error) {
	s.SetTransactionID(transactionID)
	newBackend, err := s.createBackend(from)
	if err != nil {
		return false, err
	}
	if newBackend {
		return true, s.createNewNodes(s.scaling.BaseSlots)
	}
	return s.loadNodes()
}

// SetTransactionID updates the transaction ID to be used for modifications on the configuration associated with the service.
func (s *Service) SetTransactionID(transactionID string) {
	s.transactionID = transactionID
}

// UpdateScalingParams updates parameters used for dynamic server scaling of the Service backend
func (s *Service) UpdateScalingParams(scaling ScalingParams) error {
	s.scaling = scaling
	if s.serverCount() < s.scaling.BaseSlots {
		return s.createNewNodes(s.scaling.BaseSlots - s.serverCount())
	}
	return nil
}

// Update updates the backend associated with the server based on the list of servers provided.
func (s *Service) Update(servers []ServiceServer) (bool, error) {
	log.Printf("Updating service: %s", s.name)
	reload := false
	r, err := s.expandNodes(len(servers))
	if err != nil {
		log.Printf("Error expanding nodes for service %s: %v", s.name, err)
		return false, err
	}
	reload = reload || r

	s.markRemovedNodes(servers)
	for _, server := range servers {
		log.Printf("Handling server update - Address: %s, Port: %d, Backup: %s", server.Address, server.Port, server.Backup)
		if err = s.handleNode(server); err != nil {
			log.Printf("Error handling node for service %s: %v", s.name, err)
			return false, err
		}
	}
	s.reorderNodes(len(servers))
	r, err = s.updateConfig()
	if err != nil {
		log.Printf("Error updating config for service %s: %v", s.name, err)
		return false, err
	}
	reload = reload || r
	r, err = s.removeExcessNodes(len(servers))
	if err != nil {
		log.Printf("Error removing excess nodes for service %s: %v", s.name, err)
		return false, err
	}
	reload = reload || r

	return reload, nil
}

// GetServers returns the list of servers as they are currently configured in the services backend
func (s *Service) GetServers() (models.Servers, error) {
	_, servers, err := s.client.GetServers("backend", s.name, s.transactionID)
	return servers, err
}

func (s *Service) expandNodes(nodeCount int) (bool, error) {
	currentNodeCount := s.serverCount()
	if nodeCount < currentNodeCount {
		return false, nil
	}
	newNodeCount := s.calculateNodeCount(nodeCount)
	if err := s.createNewNodes(newNodeCount - currentNodeCount); err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) serverCount() int {
	return len(s.nodes)
}

func (s *Service) calculateNodeCount(nodeCount int) int {
	if s.scaling.SlotsGrowthType == ServiceGrowthTypeLinear {
		return s.calculateNextLinearCount(nodeCount)
	}
	currentNodeCount := s.serverCount()
	for currentNodeCount < nodeCount {
		currentNodeCount *= 2
	}
	return currentNodeCount
}

func (s *Service) calculateNextLinearCount(nodeCount int) int {
	return nodeCount + s.scaling.SlotsIncrement - nodeCount%s.scaling.SlotsIncrement
}

func (s *Service) markRemovedNodes(servers []ServiceServer) {
	for _, node := range s.nodes {
		if node.disabled {
			continue
		}
		if s.nodeRemoved(node, servers) {
			node.modified = true
			node.disabled = true
			node.address = "127.0.0.1"
			node.port = 80
			node.weight = misc.Int64P(128)
			node.backup = ""
		}
	}
}

func (s *Service) handleNode(server ServiceServer) error {
	// Existing logic to check if server exists...
	log.Printf("Node handled - Backup: %s", server.Backup)
	if s.serverExists(server) {
		return nil
	}
	return s.setServer(server)
}

func (s *Service) createNewNodes(nodeCount int) error {
	for i := 0; i < nodeCount; i++ {
		if err := s.addNode(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) removeExcessNodes(newNodes int) (bool, error) {
	if newNodes < s.serverCount() {
		if s.serverCount() == s.scaling.BaseSlots {
			return false, nil
		}
		if newNodes < s.scaling.BaseSlots {
			return true, s.removeNodesAfterIndex(s.scaling.BaseSlots)
		}
	}
	lastIndex, reduce := s.getLastNodeIndex(newNodes)
	if !reduce {
		return false, nil
	}
	return true, s.removeNodesAfterIndex(lastIndex)
}

func (s *Service) getLastNodeIndex(nodeCount int) (int, bool) {
	if s.scaling.SlotsGrowthType == ServiceGrowthTypeLinear {
		if nodeCount+s.scaling.SlotsIncrement > s.serverCount() {
			return 0, false
		}
		return s.calculateNextLinearCount(nodeCount), true
	}
	if nodeCount*2 > s.serverCount() {
		return 0, false
	}

	currentNodeCount := s.serverCount()
	for {
		if currentNodeCount/2 < s.scaling.BaseSlots {
			break
		}
		if currentNodeCount/2 <= nodeCount {
			break
		}
		currentNodeCount /= 2
	}
	return currentNodeCount, true
}

func (s *Service) removeNodesAfterIndex(lastIndex int) error {
	for i := lastIndex; i < len(s.nodes); i++ {
		err := s.client.DeleteServer(s.nodes[i].name, "backend", s.name, s.transactionID, 0)
		if err != nil {
			return err
		}
	}
	s.nodes = s.nodes[:lastIndex]
	return nil
}

func (s *Service) createBackend(from string) (bool, error) {
	_, _, err := s.client.GetBackend(s.name, s.transactionID)
	if err != nil {
		err := s.client.CreateBackend(&models.Backend{
			From: from,
			Name: s.name,
		}, s.transactionID, 0)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// loadNodes loads the existing nodes for the service.
func (s *Service) loadNodes() (bool, error) {
	log.Printf("Loading nodes for service %s", s.name)
	_, servers, err := s.client.GetServers("backend", s.name, s.transactionID)
	if err != nil {
		log.Printf("Error loading nodes for service %s: %v", s.name, err)
		return false, err
	}
	for _, server := range servers {
		sNode := &serviceNode{
			name:     server.Name,
			address:  server.Address,
			port:     *server.Port,
			weight:   server.Weight,
			modified: false,
			backup:   server.Backup,
		}
		if server.Maintenance == "enabled" {
			sNode.disabled = true
		}
		s.nodes = append(s.nodes, sNode)
	}
	if s.serverCount() < s.scaling.BaseSlots {
		return true, s.createNewNodes(s.scaling.BaseSlots - s.serverCount())
	}
	return false, nil
}

// updateConfig updates the configuration for the service.
func (s *Service) updateConfig() (bool, error) {
	log.Printf("Updating configuration for service %s", s.name)
	reload := false
	for _, node := range s.nodes {
		log.Printf("Updating node - Name: %s, Backup: %s", node.name, node.backup)
		if node.modified {
			weight := int64(128) // Default weight
			if node.weight != nil {
				weight = *node.weight // Use the node's weight if set
			}
			backup := ""                  // Default disable
			if node.backup == "enabled" { // If backup value is set enable as a backup server
				backup = "enabled"
			}
			server := &models.Server{
				Name:    node.name,
				Address: node.address,
				Port:    misc.Ptr(node.port),
				ServerParams: models.ServerParams{
					Weight: misc.Int64P(int(weight)),
					Check:  "enabled",
					Backup: backup,
				},
			}
			if node.disabled {
				server.Maintenance = "enabled"
			}
			err := s.client.EditServer(node.name, "backend", s.name, server, s.transactionID, 0)
			if err != nil {
				log.Printf("Failed to update configuration for node %s in service %s: %v", node.name, s.name, err)
				return false, err
			}
			node.modified = false
			reload = true
		}
	}
	log.Printf("Configuration update completed for service %s. Reload required: %t", s.name, reload)
	return reload, nil
}

func (s *Service) nodeRemoved(node *serviceNode, servers []ServiceServer) bool {
	for _, server := range servers {
		if s.nodesMatch(node, server) {
			return false
		}
	}
	return true
}

func (s *Service) nodesMatch(sNode *serviceNode, servers ServiceServer) bool {
	return !sNode.disabled && sNode.address == servers.Address && sNode.port == int64(servers.Port) && sNode.weight == servers.Weight && sNode.backup == servers.Backup
}

func (s *Service) serverExists(server ServiceServer) bool {
	for _, sNode := range s.nodes {
		if s.nodesMatch(sNode, server) {
			return true
		}
	}
	return false
}

func (s *Service) setServer(server ServiceServer) error {
	for _, sNode := range s.nodes {
		if sNode.disabled {
			sNode.modified = true
			sNode.disabled = false
			sNode.address = server.Address
			sNode.port = int64(server.Port)
			sNode.weight = server.Weight
			sNode.backup = server.Backup
			break
		}
	}
	return nil
}

// addNode adds a new node to the service.
func (s *Service) addNode() error {
	name := s.getNodeName()
	log.Printf("Adding node to service %s. Node name: %s", s.name, name)
	// Create a temporary ServiceServer instance to use getNodeWeight method
	newServer := ServiceServer{}
	weight := newServer.getNodeWeight()
	server := &models.Server{
		Name:    name,
		Address: "127.0.0.1",
		Port:    misc.Int64P(80),
		ServerParams: models.ServerParams{
			Weight:      misc.Int64P(int(weight)),
			Maintenance: "enabled",
			Backup:      "",
		},
	}
	err := s.client.CreateServer("backend", s.name, server, s.transactionID, 0)
	if err != nil {
		log.Printf("Error adding node %s to service %s: %v", name, s.name, err)
		return err
	}
	s.nodes = append(s.nodes, &serviceNode{
		name:     name,
		address:  "127.0.0.1",
		port:     80,
		weight:   &weight,
		modified: false,
		disabled: true,
		backup:   "",
	})
	return nil
}

func (s *Service) getNodeName() string {
	name := fmt.Sprintf("SRV_%s", misc.RandomString(5))
	for _, ok := s.usedNames[name]; ok; {
		name = fmt.Sprintf("SRV_%s", misc.RandomString(5))
	}
	s.usedNames[name] = struct{}{}
	return name
}

func (s *ServiceServer) getNodeWeight() int64 {
	defaultWeight := int64(128) // Default weight

	if s != nil && s.Weight != nil {
		return *s.Weight // Use provided weight
	}
	return defaultWeight // Use default weight if not provided
}

func (s *Service) reorderNodes(count int) {
	for i := 0; i < count; i++ {
		if s.nodes[i].disabled {
			s.swapDisabledNode(i)
		}
	}
}

func (s *Service) swapDisabledNode(index int) {
	for i := len(s.nodes) - 1; i > index; i-- {
		if !s.nodes[i].disabled {
			s.nodes[i].disabled = true
			s.nodes[i].modified = true
			s.nodes[index].address = s.nodes[i].address
			s.nodes[index].port = s.nodes[i].port
			s.nodes[index].weight = s.nodes[i].weight
			s.nodes[index].backup = s.nodes[i].backup
			s.nodes[index].disabled = false
			s.nodes[index].modified = true
			s.nodes[i].address = "127.0.0.1"
			s.nodes[i].port = 80
			s.nodes[i].weight = misc.Int64P(128)
			s.nodes[i].backup = ""
			break
		}
	}
}
