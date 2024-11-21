// Code generated by go-swagger; DO NOT EDIT.

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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServerSwitchingRule Server Switching Rule
//
// HAProxy server switching rule configuration (corresponds to use-server directive)
// Example: {"cond":"if","cond_test":"{ req_ssl_sni -i www.example.com }","index":0,"target_server":"www"}
//
// swagger:model server_switching_rule
type ServerSwitchingRule struct {
	// cond
	// Enum: [if unless]
	// +kubebuilder:validation:Enum=if;unless;
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// index
	// Required: true
	// +kubebuilder:validation:Optional
	Index *int64 `json:"index"`

	// target server
	// Required: true
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TargetServer string `json:"target_server"`
}

// Validate validates this server switching rule
func (m *ServerSwitchingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serverSwitchingRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverSwitchingRuleTypeCondPropEnum = append(serverSwitchingRuleTypeCondPropEnum, v)
	}
}

const (

	// ServerSwitchingRuleCondIf captures enum value "if"
	ServerSwitchingRuleCondIf string = "if"

	// ServerSwitchingRuleCondUnless captures enum value "unless"
	ServerSwitchingRuleCondUnless string = "unless"
)

// prop value enum
func (m *ServerSwitchingRule) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverSwitchingRuleTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerSwitchingRule) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *ServerSwitchingRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *ServerSwitchingRule) validateTargetServer(formats strfmt.Registry) error {

	if err := validate.RequiredString("target_server", "body", m.TargetServer); err != nil {
		return err
	}

	if err := validate.Pattern("target_server", "body", m.TargetServer, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this server switching rule based on context it is used
func (m *ServerSwitchingRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerSwitchingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerSwitchingRule) UnmarshalBinary(b []byte) error {
	var res ServerSwitchingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
