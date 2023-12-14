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

// SpoeTransaction SPOE configuration transaction
//
// SPOE configuration transaction
// Example: {"_version":2,"id":"273e3385-2d0c-4fb1-aa27-93cbb31ff203","status":"in_progress"}
//
// swagger:model spoe_transaction
type SpoeTransaction struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// id
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	ID string `json:"id,omitempty"`

	// status
	// Enum: [failed in_progress success]
	// +kubebuilder:validation:Enum=failed;in_progress;success;
	Status string `json:"status,omitempty"`
}

// Validate validates this spoe transaction
func (m *SpoeTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpoeTransaction) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var spoeTransactionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["failed","in_progress","success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		spoeTransactionTypeStatusPropEnum = append(spoeTransactionTypeStatusPropEnum, v)
	}
}

const (

	// SpoeTransactionStatusFailed captures enum value "failed"
	SpoeTransactionStatusFailed string = "failed"

	// SpoeTransactionStatusInProgress captures enum value "in_progress"
	SpoeTransactionStatusInProgress string = "in_progress"

	// SpoeTransactionStatusSuccess captures enum value "success"
	SpoeTransactionStatusSuccess string = "success"
)

// prop value enum
func (m *SpoeTransaction) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, spoeTransactionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SpoeTransaction) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this spoe transaction based on context it is used
func (m *SpoeTransaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpoeTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpoeTransaction) UnmarshalBinary(b []byte) error {
	var res SpoeTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
