// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/client-native/v5/config-parser/parsers/tcp"
)

func TestRequeststcp(t *testing.T) {
	tests := map[string]bool{
		`tcp-request content capture str("DNS resolution failure") len 32 unless dns_successful`: true,
		"tcp-request content accept":                                                   true,
		"tcp-request content accept if !HTTP":                                          true,
		"tcp-request content reject":                                                   true,
		"tcp-request content reject if !HTTP":                                          true,
		"tcp-request content capture req.payload(0,6) len 6":                           true,
		"tcp-request content capture req.payload(0,6) len 6 if !HTTP":                  true,
		"tcp-request content do-resolve(txn.myip,mydns,ipv6) capture.req.hdr(0),lower": true,
		"tcp-request content do-resolve(txn.myip,mydns) capture.req.hdr(0),lower":      true,
		"tcp-request content set-priority-class int(1)":                                true,
		"tcp-request content set-priority-class int(1) if some_check":                  true,
		"tcp-request content set-priority-offset int(10)":                              true,
		"tcp-request content set-priority-offset int(10) if some_check":                true,
		"tcp-request content track-sc0 src":                                            true,
		"tcp-request content track-sc0 src if some_check":                              true,
		"tcp-request content track-sc1 src":                                            true,
		"tcp-request content track-sc1 src if some_check":                              true,
		"tcp-request content track-sc2 src":                                            true,
		"tcp-request content track-sc2 src if some_check":                              true,
		"tcp-request content track-sc0 src table foo":                                  true,
		"tcp-request content track-sc0 src table foo if some_check":                    true,
		"tcp-request content track-sc1 src table foo":                                  true,
		"tcp-request content track-sc1 src table foo if some_check":                    true,
		"tcp-request content track-sc2 src table foo":                                  true,
		"tcp-request content track-sc2 src table foo if some_check":                    true,
		"tcp-request content track-sc5 src":                                            true,
		"tcp-request content track-sc5 src if some_check":                              true,
		"tcp-request content track-sc5 src table foo":                                  true,
		"tcp-request content track-sc5 src table foo if some_check":                    true,
		"tcp-request content sc-inc-gpc(1,2)":                                          true,
		"tcp-request content sc-inc-gpc(1,2) if is-error":                              true,
		"tcp-request content sc-inc-gpc0(2)":                                           true,
		"tcp-request content sc-inc-gpc0(2) if is-error":                               true,
		"tcp-request content sc-inc-gpc1(2)":                                           true,
		"tcp-request content sc-inc-gpc1(2) if is-error":                               true,
		"tcp-request content sc-set-gpt(x,9) 1337 if exceeds_limit":                    true,
		"tcp-request content sc-set-gpt0(0) 1337":                                      true,
		"tcp-request content sc-set-gpt0(0) 1337 if exceeds_limit":                     true,
		"tcp-request content sc-add-gpc(1,2) 1":                                        true,
		"tcp-request content sc-add-gpc(1,2) 1 if is-error":                            true,
		"tcp-request content set-dst ipv4(10.0.0.1)":                                   true,
		"tcp-request content set-var(sess.src) src":                                    true,
		"tcp-request content set-var(sess.dn) ssl_c_s_dn":                              true,
		"tcp-request content set-var-fmt(sess.src) src":                                true,
		"tcp-request content set-var-fmt(sess.dn) ssl_c_s_dn":                          true,
		"tcp-request content unset-var(sess.src)":                                      true,
		"tcp-request content unset-var(sess.dn)":                                       true,
		"tcp-request content silent-drop":                                              true,
		"tcp-request content silent-drop if !HTTP":                                     true,
		"tcp-request content silent-drop rst-ttl 1":                                    true,
		"tcp-request content silent-drop rst-ttl 1 if { src,table_http_req_rate(ratelimits.agg),sub(txn.rate_limit) ge 1000 }":    true,
		"tcp-request content send-spoe-group engine group":                                                                        true,
		"tcp-request content use-service lua.deny":                                                                                true,
		"tcp-request content use-service lua.deny if !HTTP":                                                                       true,
		"tcp-request content lua.foo":                                                                                             true,
		"tcp-request content lua.foo param if !HTTP":                                                                              true,
		"tcp-request content lua.foo param param1":                                                                                true,
		"tcp-request connection accept":                                                                                           true,
		"tcp-request connection accept if !HTTP":                                                                                  true,
		"tcp-request connection reject":                                                                                           true,
		"tcp-request connection reject if !HTTP":                                                                                  true,
		"tcp-request connection expect-proxy layer4 if { src -f proxies.lst }":                                                    true,
		"tcp-request connection expect-netscaler-cip layer4":                                                                      true,
		"tcp-request connection expect-netscaler-cip layer4 if TRUE":                                                              true,
		"tcp-request connection capture req.payload(0,6) len 6":                                                                   true,
		"tcp-request connection track-sc0 src":                                                                                    true,
		"tcp-request connection track-sc0 src if some_check":                                                                      true,
		"tcp-request connection track-sc1 src":                                                                                    true,
		"tcp-request connection track-sc1 src if some_check":                                                                      true,
		"tcp-request connection track-sc2 src":                                                                                    true,
		"tcp-request connection track-sc2 src if some_check":                                                                      true,
		"tcp-request connection track-sc0 src table foo":                                                                          true,
		"tcp-request connection track-sc0 src table foo if some_check":                                                            true,
		"tcp-request connection track-sc1 src table foo":                                                                          true,
		"tcp-request connection track-sc1 src table foo if some_check":                                                            true,
		"tcp-request connection track-sc2 src table foo":                                                                          true,
		"tcp-request connection track-sc2 src table foo if some_check":                                                            true,
		"tcp-request connection track-sc5 src":                                                                                    true,
		"tcp-request connection track-sc5 src if some_check":                                                                      true,
		"tcp-request connection track-sc5 src table foo":                                                                          true,
		"tcp-request connection track-sc5 src table foo if some_check":                                                            true,
		"tcp-request connection sc-add-gpc(1,2) 1":                                                                                true,
		"tcp-request connection sc-add-gpc(1,2) 1 if is-error":                                                                    true,
		"tcp-request connection sc-inc-gpc(1,2)":                                                                                  true,
		"tcp-request connection sc-inc-gpc(1,2) if is-error":                                                                      true,
		"tcp-request connection sc-inc-gpc0(2)":                                                                                   true,
		"tcp-request connection sc-inc-gpc0(2) if is-error":                                                                       true,
		"tcp-request connection sc-inc-gpc1(2)":                                                                                   true,
		"tcp-request connection sc-inc-gpc1(2) if is-error":                                                                       true,
		"tcp-request connection sc-set-gpt(scx,44) 1337 if exceeds_limit":                                                         true,
		"tcp-request connection sc-set-gpt0(0) 1337":                                                                              true,
		"tcp-request connection sc-set-gpt0(0) 1337 if exceeds_limit":                                                             true,
		"tcp-request connection set-src src,ipmask(24)":                                                                           true,
		"tcp-request connection set-src src,ipmask(24) if some_check":                                                             true,
		"tcp-request connection set-src hdr(x-forwarded-for)":                                                                     true,
		"tcp-request connection set-src hdr(x-forwarded-for) if some_check":                                                       true,
		"tcp-request connection silent-drop":                                                                                      true,
		"tcp-request connection silent-drop if !HTTP":                                                                             true,
		"tcp-request connection silent-drop rst-ttl 1":                                                                            true,
		"tcp-request connection silent-drop rst-ttl 1 if { src,table_http_req_rate(ratelimits.agg),sub(txn.rate_limit) ge 1000 }": true,
		"tcp-request connection lua.foo":                                                                                          true,
		"tcp-request connection lua.foo param if !HTTP":                                                                           true,
		"tcp-request connection lua.foo param param1":                                                                             true,
		"tcp-request session accept":                                                                                              true,
		"tcp-request session accept if !HTTP":                                                                                     true,
		"tcp-request session reject":                                                                                              true,
		"tcp-request session reject if !HTTP":                                                                                     true,
		"tcp-request session track-sc0 src":                                                                                       true,
		"tcp-request session track-sc0 src if some_check":                                                                         true,
		"tcp-request session track-sc1 src":                                                                                       true,
		"tcp-request session track-sc1 src if some_check":                                                                         true,
		"tcp-request session track-sc2 src":                                                                                       true,
		"tcp-request session track-sc2 src if some_check":                                                                         true,
		"tcp-request session track-sc0 src table foo":                                                                             true,
		"tcp-request session track-sc0 src table foo if some_check":                                                               true,
		"tcp-request session track-sc1 src table foo":                                                                             true,
		"tcp-request session track-sc1 src table foo if some_check":                                                               true,
		"tcp-request session track-sc2 src table foo":                                                                             true,
		"tcp-request session track-sc2 src table foo if some_check":                                                               true,
		"tcp-request session track-sc5 src":                                                                                       true,
		"tcp-request session track-sc5 src if some_check":                                                                         true,
		"tcp-request session track-sc5 src table foo":                                                                             true,
		"tcp-request session track-sc5 src table foo if some_check":                                                               true,
		"tcp-request session sc-add-gpc(1,2) 1":                                                                                   true,
		"tcp-request session sc-add-gpc(1,2) 1 if is-error":                                                                       true,
		"tcp-request session sc-inc-gpc(1,2)":                                                                                     true,
		"tcp-request session sc-inc-gpc(1,2) if is-error":                                                                         true,
		"tcp-request session sc-inc-gpc0(2)":                                                                                      true,
		"tcp-request session sc-inc-gpc0(2) if is-error":                                                                          true,
		"tcp-request session sc-inc-gpc1(2)":                                                                                      true,
		"tcp-request session sc-inc-gpc1(2) if is-error":                                                                          true,
		"tcp-request session sc-set-gpt(sc5,1) 1337 if exceeds_limit":                                                             true,
		"tcp-request session sc-set-gpt0(0) 1337":                                                                                 true,
		"tcp-request session sc-set-gpt0(0) 1337 if exceeds_limit":                                                                true,
		"tcp-request session set-var(sess.src) src":                                                                               true,
		"tcp-request session set-var(sess.dn) ssl_c_s_dn":                                                                         true,
		"tcp-request session set-var-fmt(sess.src) src":                                                                           true,
		"tcp-request session set-var-fmt(sess.dn) ssl_c_s_dn":                                                                     true,
		"tcp-request session unset-var(sess.src)":                                                                                 true,
		"tcp-request session unset-var(sess.dn)":                                                                                  true,
		"tcp-request session silent-drop":                                                                                         true,
		"tcp-request session silent-drop if !HTTP":                                                                                true,
		"tcp-request session silent-drop rst-ttl 1":                                                                               true,
		"tcp-request session silent-drop rst-ttl 1 if { src,table_http_req_rate(ratelimits.agg),sub(txn.rate_limit) ge 1000 }": true,
		"tcp-request session attach-srv srv1":                                   true,
		"tcp-request session attach-srv srv1 name example.com":                  true,
		"tcp-request session attach-srv srv1 name example.com if exceeds_limit": true,
		"tcp-request content set-bandwidth-limit my-limit":                      true,
		"tcp-request content set-bandwidth-limit my-limit limit 1m period 10s":  true,
		"tcp-request content set-bandwidth-limit my-limit period 10s":           true,
		"tcp-request content set-bandwidth-limit my-limit limit 1m":             true,
		"tcp-request content set-log-level silent":                              true,
		"tcp-request content set-log-level silent if FALSE":                     true,
		"tcp-request content set-mark 20":                                       true,
		"tcp-request content set-mark 0x1Ab if FALSE":                           true,
		"tcp-request connection set-mark 20":                                    true,
		"tcp-request connection set-mark 0x1Ab if FALSE":                        true,
		"tcp-request connection set-src-port hdr(port)":                         true,
		"tcp-request connection set-src-port hdr(port) if FALSE":                true,
		"tcp-request content set-src-port hdr(port)":                            true,
		"tcp-request content set-src-port hdr(port) if FALSE":                   true,
		"tcp-request content set-tos 0 if FALSE":                                true,
		"tcp-request content set-tos 0":                                         true,
		"tcp-request connection set-tos 0 if FALSE":                             true,
		"tcp-request connection set-tos 0":                                      true,
		"tcp-request connection set-var-fmt(txn.ip_port) %%[dst]:%%[dst_port]":  true,
		"tcp-request content set-nice 0 if FALSE":                               true,
		"tcp-request content set-nice 0":                                        true,
		"tcp-request content switch-mode http":                                  true,
		"tcp-request content switch-mode http if FALSE":                         true,
		"tcp-request content switch-mode http proto my-proto":                   true,
		"tcp-request content sc-inc-gpc":                                        false,
		"tcp-request content sc-inc-gpc0":                                       false,
		"tcp-request content sc-inc-gpc1":                                       false,
		"tcp-request content sc-add-gpc":                                        false,
		"tcp-request silent-drop rst-ttl":                                       false,
		"tcp-request connection sc-add-gpc":                                     false,
		"tcp-request connection sc-inc-gpc":                                     false,
		"tcp-request connection sc-inc-gpc0":                                    false,
		"tcp-request connection sc-inc-gpc1":                                    false,
		"tcp-request connection set-src":                                        false,
		"tcp-request connection silent-drop rst-ttl":                            false,
		"tcp-request session sc-add-gpc":                                        false,
		"tcp-request session sc-inc-gpc":                                        false,
		"tcp-request session sc-inc-gpc0":                                       false,
		"tcp-request session sc-inc-gpc1":                                       false,
		"tcp-request session silent-drop rst-ttl":                               false,
		"tcp-request session attach-srv":                                        false,
		"tcp-request session attach-srv srv1 name":                              false,
		"tcp-request session attach-srv srv1 if":                                false,
		"tcp-request session attach-srv srv1 name example.com unless":           false,
		"tcp-request":            false,
		"tcp-request content":    false,
		"tcp-request connection": false,
		"tcp-request content do-resolve(txn.myip) capture.req.hdr(0),lower": false,
		"tcp-request session":                                               false,
		"tcp-request content lua.":                                          false,
		"tcp-request content lua. param":                                    false,
		"tcp-request connection lua.":                                       false,
		"tcp-request connection lua. param":                                 false,
		"tcp-request content track-sc0 src table":                           false,
		"tcp-request content track-sc0 src table if some_check":             false,
		"tcp-request content track-sc1 src table":                           false,
		"tcp-request content track-sc1 src table if some_check":             false,
		"tcp-request content track-sc2 src table":                           false,
		"tcp-request content track-sc2 src table if some_check":             false,
		"tcp-request connection track-sc0 src table":                        false,
		"tcp-request connection track-sc0 src table if some_check":          false,
		"tcp-request connection track-sc1 src table":                        false,
		"tcp-request connection track-sc1 src table if some_check":          false,
		"tcp-request connection track-sc2 src table":                        false,
		"tcp-request connection track-sc2 src table if some_check":          false,
		"tcp-request session track-sc0 src table":                           false,
		"tcp-request session track-sc0 src table if some_check":             false,
		"tcp-request session track-sc1 src table":                           false,
		"tcp-request session track-sc1 src table if some_check":             false,
		"tcp-request session track-sc2 src table":                           false,
		"tcp-request session track-sc2 src table if some_check":             false,
		"tcp-request content track-sc5 src table":                           false,
		"tcp-request content track-sc5 src table if some_check":             false,
		"tcp-request connection track-sc5 src table":                        false,
		"tcp-request connection track-sc5 src table if some_check":          false,
		"tcp-request session track-sc5 src table":                           false,
		"tcp-request session track-sc5 src table if some_check":             false,
		"tcp-request content set-bandwidth-limit my-limit limit":            false,
		"tcp-request content set-bandwidth-limit my-limit period":           false,
		"tcp-request content set-bandwidth-limit my-limit 10s":              false,
		"tcp-request content set-bandwidth-limit my-limit period 10s limit": false,
		"tcp-request content set-bandwidth-limit my-limit limit period 10s": false,
		"tcp-request content set-log-level":                                 false,
		"tcp-request connection set-mark":                                   false,
		"tcp-request content set-mark":                                      false,
		"tcp-request connection set-src-port":                               false,
		"tcp-request content set-src-port":                                  false,
		"tcp-request connection set-tos":                                    false,
		"tcp-request content set-tos":                                       false,
		"tcp-request content set-nice":                                      false,
		"tcp-request content switch-mode":                                   false,
		"tcp-request content switch-mode tcp":                               false,
		"tcp-request content switch-mode http proto":                        false,
		"---":     false,
		"--- ---": false,
	}
	parser := &tcp.Requests{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
			line := strings.TrimSpace(command)
			lines := strings.SplitN(line, "\n", -1)
			var err error
			parser.Init()
			if len(lines) > 1 {
				for _, line = range lines {
					line = strings.TrimSpace(line)
					if err = ProcessLine(line, parser); err != nil {
						break
					}
				}
			} else {
				err = ProcessLine(line, parser)
			}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
