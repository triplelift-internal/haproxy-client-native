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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cookie cookie
//
// swagger:model cookie
type Cookie struct {

	// attrs
	Attrs []*Attr `json:"attr,omitempty"`

	// domains
	Domains []*Domain `json:"domain,omitempty"`

	// dynamic
	Dynamic bool `json:"dynamic,omitempty"`

	// httponly
	Httponly bool `json:"httponly,omitempty"`

	// indirect
	Indirect bool `json:"indirect,omitempty"`

	// maxidle
	Maxidle int64 `json:"maxidle,omitempty"`

	// maxlife
	Maxlife int64 `json:"maxlife,omitempty"`

	// name
	// Required: true
	// Pattern: ^[^\s]+$
	Name *string `json:"name"`

	// nocache
	Nocache bool `json:"nocache,omitempty"`

	// postonly
	Postonly bool `json:"postonly,omitempty"`

	// preserve
	Preserve bool `json:"preserve,omitempty"`

	// secure
	Secure bool `json:"secure,omitempty"`

	// type
	// Enum: [rewrite insert prefix]
	// +kubebuilder:validation:Enum=rewrite;insert;prefix;
	Type string `json:"type,omitempty"`
}

// Validate validates this cookie
func (m *Cookie) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cookie) validateAttrs(formats strfmt.Registry) error {
	if swag.IsZero(m.Attrs) { // not required
		return nil
	}

	for i := 0; i < len(m.Attrs); i++ {
		if swag.IsZero(m.Attrs[i]) { // not required
			continue
		}

		if m.Attrs[i] != nil {
			if err := m.Attrs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cookie) validateDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.Domains) { // not required
		return nil
	}

	for i := 0; i < len(m.Domains); i++ {
		if swag.IsZero(m.Domains[i]) { // not required
			continue
		}

		if m.Domains[i] != nil {
			if err := m.Domains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cookie) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var cookieTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rewrite","insert","prefix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cookieTypeTypePropEnum = append(cookieTypeTypePropEnum, v)
	}
}

const (

	// CookieTypeRewrite captures enum value "rewrite"
	CookieTypeRewrite string = "rewrite"

	// CookieTypeInsert captures enum value "insert"
	CookieTypeInsert string = "insert"

	// CookieTypePrefix captures enum value "prefix"
	CookieTypePrefix string = "prefix"
)

// prop value enum
func (m *Cookie) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cookieTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cookie) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cookie based on the context it is used
func (m *Cookie) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cookie) contextValidateAttrs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attrs); i++ {

		if m.Attrs[i] != nil {
			if err := m.Attrs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cookie) contextValidateDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Domains); i++ {

		if m.Domains[i] != nil {
			if err := m.Domains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cookie) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cookie) UnmarshalBinary(b []byte) error {
	var res Cookie
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Attr attr
//
// swagger:model Attr
type Attr struct {

	// value
	// Pattern: ^[^\s]+$
	Value string `json:"value,omitempty"`
}

// Validate validates this attr
func (m *Attr) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Attr) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.Pattern("value", "body", m.Value, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this attr based on context it is used
func (m *Attr) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Attr) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attr) UnmarshalBinary(b []byte) error {
	var res Attr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Domain domain
//
// swagger:model Domain
type Domain struct {

	// value
	// Pattern: ^[^\s]+$
	Value string `json:"value,omitempty"`
}

// Validate validates this domain
func (m *Domain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Domain) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.Pattern("value", "body", m.Value, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain based on context it is used
func (m *Domain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Domain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Domain) UnmarshalBinary(b []byte) error {
	var res Domain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
