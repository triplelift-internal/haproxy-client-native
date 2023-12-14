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

// BackendSwitchingRule Backend Switching Rule
//
// HAProxy backend switching rule configuration (corresponds to use_backend directive)
// Example: {"cond":"if","cond_test":"{ req_ssl_sni -i www.example.com }","index":0,"name":"test_backend"}
//
// swagger:model backend_switching_rule
type BackendSwitchingRule struct {
	// cond
	// Enum: [if unless]
	// +kubebuilder:validation:Enum=if;unless;
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// name
	// Required: true
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	Name string `json:"name"`
}

// Validate validates this backend switching rule
func (m *BackendSwitchingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendSwitchingRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendSwitchingRuleTypeCondPropEnum = append(backendSwitchingRuleTypeCondPropEnum, v)
	}
}

const (

	// BackendSwitchingRuleCondIf captures enum value "if"
	BackendSwitchingRuleCondIf string = "if"

	// BackendSwitchingRuleCondUnless captures enum value "unless"
	BackendSwitchingRuleCondUnless string = "unless"
)

// prop value enum
func (m *BackendSwitchingRule) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, backendSwitchingRuleTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BackendSwitchingRule) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *BackendSwitchingRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *BackendSwitchingRule) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backend switching rule based on context it is used
func (m *BackendSwitchingRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackendSwitchingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendSwitchingRule) UnmarshalBinary(b []byte) error {
	var res BackendSwitchingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
