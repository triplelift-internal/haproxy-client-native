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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPResponseRule HTTP Response Rule
//
// HAProxy HTTP response rule configuration (corresponds to http-response directives)
//
// swagger:model http_response_rule
type HTTPResponseRule struct {

	// return headers
	ReturnHeaders []*HTTPResponseRuleReturnHdrsItems0 `json:"return_hdrs"`

	// acl file
	// Pattern: ^[^\s]+$
	ACLFile string `json:"acl_file,omitempty"`

	// acl keyfmt
	// Pattern: ^[^\s]+$
	ACLKeyfmt string `json:"acl_keyfmt,omitempty"`

	// cache name
	// Pattern: ^[^\s]+$
	CacheName string `json:"cache_name,omitempty"`

	// capture id
	CaptureID *int64 `json:"capture_id,omitempty"`

	// capture sample
	// Pattern: ^[^\s]+$
	CaptureSample string `json:"capture_sample,omitempty"`

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// deny status
	// Maximum: 599
	// Minimum: 200
	DenyStatus *int64 `json:"deny_status,omitempty"`

	// hdr format
	HdrFormat string `json:"hdr_format,omitempty"`

	// hdr match
	HdrMatch string `json:"hdr_match,omitempty"`

	// hdr method
	HdrMethod string `json:"hdr_method,omitempty"`

	// hdr name
	HdrName string `json:"hdr_name,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// log level
	// Enum: [emerg alert crit err warning notice info debug silent]
	LogLevel string `json:"log_level,omitempty"`

	// lua action
	// Pattern: ^[^\s]+$
	LuaAction string `json:"lua_action,omitempty"`

	// lua params
	LuaParams string `json:"lua_params,omitempty"`

	// map file
	// Pattern: ^[^\s]+$
	MapFile string `json:"map_file,omitempty"`

	// map keyfmt
	// Pattern: ^[^\s]+$
	MapKeyfmt string `json:"map_keyfmt,omitempty"`

	// map valuefmt
	// Pattern: ^[^\s]+$
	MapValuefmt string `json:"map_valuefmt,omitempty"`

	// mark value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	MarkValue string `json:"mark_value,omitempty"`

	// nice value
	// Maximum: 1024
	// Minimum: -1024
	NiceValue int64 `json:"nice_value,omitempty"`

	// redir code
	// Enum: [301 302 303 307 308]
	RedirCode *int64 `json:"redir_code,omitempty"`

	// redir option
	RedirOption string `json:"redir_option,omitempty"`

	// redir type
	// Enum: [location prefix scheme]
	RedirType string `json:"redir_type,omitempty"`

	// redir value
	// Pattern: ^[^\s]+$
	RedirValue string `json:"redir_value,omitempty"`

	// return content
	ReturnContent string `json:"return_content,omitempty"`

	// return content format
	// Enum: [default-errorfile errorfile errorfiles file lf-file string lf-string]
	ReturnContentFormat string `json:"return_content_format,omitempty"`

	// return content type
	ReturnContentType *string `json:"return_content_type,omitempty"`

	// return status code
	// Maximum: 599
	// Minimum: 200
	ReturnStatusCode *int64 `json:"return_status_code,omitempty"`

	// sc expr
	ScExpr string `json:"sc_expr,omitempty"`

	// sc id
	ScID int64 `json:"sc_id,omitempty"`

	// sc int
	ScInt *int64 `json:"sc_int,omitempty"`

	// spoe engine
	// Pattern: ^[^\s]+$
	SpoeEngine string `json:"spoe_engine,omitempty"`

	// spoe group
	// Pattern: ^[^\s]+$
	SpoeGroup string `json:"spoe_group,omitempty"`

	// status
	// Maximum: 999
	// Minimum: 100
	Status int64 `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// strict mode
	// Enum: [on off]
	StrictMode string `json:"strict_mode,omitempty"`

	// tos value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	TosValue string `json:"tos_value,omitempty"`

	// track sc0 key
	// Pattern: ^[^\s]+$
	TrackSc0Key string `json:"track-sc0-key,omitempty"`

	// track sc0 table
	// Pattern: ^[^\s]+$
	TrackSc0Table string `json:"track-sc0-table,omitempty"`

	// track sc1 key
	// Pattern: ^[^\s]+$
	TrackSc1Key string `json:"track-sc1-key,omitempty"`

	// track sc1 table
	// Pattern: ^[^\s]+$
	TrackSc1Table string `json:"track-sc1-table,omitempty"`

	// track sc2 key
	// Pattern: ^[^\s]+$
	TrackSc2Key string `json:"track-sc2-key,omitempty"`

	// track sc2 table
	// Pattern: ^[^\s]+$
	TrackSc2Table string `json:"track-sc2-table,omitempty"`

	// type
	// Required: true
	// Enum: [add-acl add-header allow cache-store capture del-acl del-header del-map deny redirect replace-header replace-value return sc-inc-gpc0 sc-inc-gpc1 sc-set-gpt0 send-spoe-group set-header set-log-level set-map set-mark set-nice set-status set-tos set-var set-var-fmt silent-drop strict-mode track-sc0 track-sc1 track-sc2 unset-var wait-for-body time]
	Type string `json:"type"`

	// var expr
	VarExpr string `json:"var_expr,omitempty"`

	// var format
	VarFormat string `json:"var_format,omitempty"`

	// var name
	// Pattern: ^[^\s]+$
	VarName string `json:"var_name,omitempty"`

	// var scope
	// Pattern: ^[^\s]+$
	VarScope string `json:"var_scope,omitempty"`

	// wait at least
	WaitAtLeast *int64 `json:"wait_at_least,omitempty"`

	// wait time
	WaitTime *int64 `json:"wait_time,omitempty"`
}

// Validate validates this http response rule
func (m *HTTPResponseRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReturnHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateACLFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateACLKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCacheName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureSample(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDenyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapValuefmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNiceValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnContentFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrictMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTosValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPResponseRule) validateReturnHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.ReturnHeaders); i++ {
		if swag.IsZero(m.ReturnHeaders[i]) { // not required
			continue
		}

		if m.ReturnHeaders[i] != nil {
			if err := m.ReturnHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("return_hdrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HTTPResponseRule) validateACLFile(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLFile) { // not required
		return nil
	}

	if err := validate.Pattern("acl_file", "body", string(m.ACLFile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateACLKeyfmt(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("acl_keyfmt", "body", string(m.ACLKeyfmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateCacheName(formats strfmt.Registry) error {

	if swag.IsZero(m.CacheName) { // not required
		return nil
	}

	if err := validate.Pattern("cache_name", "body", string(m.CacheName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateCaptureSample(formats strfmt.Registry) error {

	if swag.IsZero(m.CaptureSample) { // not required
		return nil
	}

	if err := validate.Pattern("capture_sample", "body", string(m.CaptureSample), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeCondPropEnum = append(httpResponseRuleTypeCondPropEnum, v)
	}
}

const (

	// HTTPResponseRuleCondIf captures enum value "if"
	HTTPResponseRuleCondIf string = "if"

	// HTTPResponseRuleCondUnless captures enum value "unless"
	HTTPResponseRuleCondUnless string = "unless"
)

// prop value enum
func (m *HTTPResponseRule) validateCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateDenyStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.DenyStatus) { // not required
		return nil
	}

	if err := validate.MinimumInt("deny_status", "body", int64(*m.DenyStatus), 200, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("deny_status", "body", int64(*m.DenyStatus), 599, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug","silent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeLogLevelPropEnum = append(httpResponseRuleTypeLogLevelPropEnum, v)
	}
}

const (

	// HTTPResponseRuleLogLevelEmerg captures enum value "emerg"
	HTTPResponseRuleLogLevelEmerg string = "emerg"

	// HTTPResponseRuleLogLevelAlert captures enum value "alert"
	HTTPResponseRuleLogLevelAlert string = "alert"

	// HTTPResponseRuleLogLevelCrit captures enum value "crit"
	HTTPResponseRuleLogLevelCrit string = "crit"

	// HTTPResponseRuleLogLevelErr captures enum value "err"
	HTTPResponseRuleLogLevelErr string = "err"

	// HTTPResponseRuleLogLevelWarning captures enum value "warning"
	HTTPResponseRuleLogLevelWarning string = "warning"

	// HTTPResponseRuleLogLevelNotice captures enum value "notice"
	HTTPResponseRuleLogLevelNotice string = "notice"

	// HTTPResponseRuleLogLevelInfo captures enum value "info"
	HTTPResponseRuleLogLevelInfo string = "info"

	// HTTPResponseRuleLogLevelDebug captures enum value "debug"
	HTTPResponseRuleLogLevelDebug string = "debug"

	// HTTPResponseRuleLogLevelSilent captures enum value "silent"
	HTTPResponseRuleLogLevelSilent string = "silent"
)

// prop value enum
func (m *HTTPResponseRule) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeLogLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateLogLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogLevelEnum("log_level", "body", m.LogLevel); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateLuaAction(formats strfmt.Registry) error {

	if swag.IsZero(m.LuaAction) { // not required
		return nil
	}

	if err := validate.Pattern("lua_action", "body", string(m.LuaAction), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapFile(formats strfmt.Registry) error {

	if swag.IsZero(m.MapFile) { // not required
		return nil
	}

	if err := validate.Pattern("map_file", "body", string(m.MapFile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapKeyfmt(formats strfmt.Registry) error {

	if swag.IsZero(m.MapKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_keyfmt", "body", string(m.MapKeyfmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapValuefmt(formats strfmt.Registry) error {

	if swag.IsZero(m.MapValuefmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_valuefmt", "body", string(m.MapValuefmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMarkValue(formats strfmt.Registry) error {

	if swag.IsZero(m.MarkValue) { // not required
		return nil
	}

	if err := validate.Pattern("mark_value", "body", string(m.MarkValue), `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateNiceValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NiceValue) { // not required
		return nil
	}

	if err := validate.MinimumInt("nice_value", "body", int64(m.NiceValue), -1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nice_value", "body", int64(m.NiceValue), 1024, false); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeRedirCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[301,302,303,307,308]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeRedirCodePropEnum = append(httpResponseRuleTypeRedirCodePropEnum, v)
	}
}

// prop value enum
func (m *HTTPResponseRule) validateRedirCodeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeRedirCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateRedirCode(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirCodeEnum("redir_code", "body", *m.RedirCode); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeRedirTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["location","prefix","scheme"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeRedirTypePropEnum = append(httpResponseRuleTypeRedirTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleRedirTypeLocation captures enum value "location"
	HTTPResponseRuleRedirTypeLocation string = "location"

	// HTTPResponseRuleRedirTypePrefix captures enum value "prefix"
	HTTPResponseRuleRedirTypePrefix string = "prefix"

	// HTTPResponseRuleRedirTypeScheme captures enum value "scheme"
	HTTPResponseRuleRedirTypeScheme string = "scheme"
)

// prop value enum
func (m *HTTPResponseRule) validateRedirTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeRedirTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateRedirType(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirTypeEnum("redir_type", "body", m.RedirType); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateRedirValue(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirValue) { // not required
		return nil
	}

	if err := validate.Pattern("redir_value", "body", string(m.RedirValue), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeReturnContentFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default-errorfile","errorfile","errorfiles","file","lf-file","string","lf-string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeReturnContentFormatPropEnum = append(httpResponseRuleTypeReturnContentFormatPropEnum, v)
	}
}

const (

	// HTTPResponseRuleReturnContentFormatDefaultErrorfile captures enum value "default-errorfile"
	HTTPResponseRuleReturnContentFormatDefaultErrorfile string = "default-errorfile"

	// HTTPResponseRuleReturnContentFormatErrorfile captures enum value "errorfile"
	HTTPResponseRuleReturnContentFormatErrorfile string = "errorfile"

	// HTTPResponseRuleReturnContentFormatErrorfiles captures enum value "errorfiles"
	HTTPResponseRuleReturnContentFormatErrorfiles string = "errorfiles"

	// HTTPResponseRuleReturnContentFormatFile captures enum value "file"
	HTTPResponseRuleReturnContentFormatFile string = "file"

	// HTTPResponseRuleReturnContentFormatLfFile captures enum value "lf-file"
	HTTPResponseRuleReturnContentFormatLfFile string = "lf-file"

	// HTTPResponseRuleReturnContentFormatString captures enum value "string"
	HTTPResponseRuleReturnContentFormatString string = "string"

	// HTTPResponseRuleReturnContentFormatLfString captures enum value "lf-string"
	HTTPResponseRuleReturnContentFormatLfString string = "lf-string"
)

// prop value enum
func (m *HTTPResponseRule) validateReturnContentFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeReturnContentFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateReturnContentFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnContentFormat) { // not required
		return nil
	}

	// value enum
	if err := m.validateReturnContentFormatEnum("return_content_format", "body", m.ReturnContentFormat); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateReturnStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnStatusCode) { // not required
		return nil
	}

	if err := validate.MinimumInt("return_status_code", "body", int64(*m.ReturnStatusCode), 200, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("return_status_code", "body", int64(*m.ReturnStatusCode), 599, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeEngine(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeEngine) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_engine", "body", string(m.SpoeEngine), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeGroup) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_group", "body", string(m.SpoeGroup), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MinimumInt("status", "body", int64(m.Status), 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", int64(m.Status), 999, false); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeStrictModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeStrictModePropEnum = append(httpResponseRuleTypeStrictModePropEnum, v)
	}
}

const (

	// HTTPResponseRuleStrictModeOn captures enum value "on"
	HTTPResponseRuleStrictModeOn string = "on"

	// HTTPResponseRuleStrictModeOff captures enum value "off"
	HTTPResponseRuleStrictModeOff string = "off"
)

// prop value enum
func (m *HTTPResponseRule) validateStrictModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeStrictModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateStrictMode(formats strfmt.Registry) error {

	if swag.IsZero(m.StrictMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrictModeEnum("strict_mode", "body", m.StrictMode); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTosValue(formats strfmt.Registry) error {

	if swag.IsZero(m.TosValue) { // not required
		return nil
	}

	if err := validate.Pattern("tos_value", "body", string(m.TosValue), `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc0Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc0Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-key", "body", string(m.TrackSc0Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc0Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc0Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-table", "body", string(m.TrackSc0Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc1Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc1Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-key", "body", string(m.TrackSc1Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc1Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc1Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-table", "body", string(m.TrackSc1Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc2Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc2Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-key", "body", string(m.TrackSc2Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc2Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc2Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-table", "body", string(m.TrackSc2Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add-acl","add-header","allow","cache-store","capture","del-acl","del-header","del-map","deny","redirect","replace-header","replace-value","return","sc-inc-gpc0","sc-inc-gpc1","sc-set-gpt0","send-spoe-group","set-header","set-log-level","set-map","set-mark","set-nice","set-status","set-tos","set-var","set-var-fmt","silent-drop","strict-mode","track-sc0","track-sc1","track-sc2","unset-var","wait-for-body time"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeTypePropEnum = append(httpResponseRuleTypeTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleTypeAddACL captures enum value "add-acl"
	HTTPResponseRuleTypeAddACL string = "add-acl"

	// HTTPResponseRuleTypeAddHeader captures enum value "add-header"
	HTTPResponseRuleTypeAddHeader string = "add-header"

	// HTTPResponseRuleTypeAllow captures enum value "allow"
	HTTPResponseRuleTypeAllow string = "allow"

	// HTTPResponseRuleTypeCacheStore captures enum value "cache-store"
	HTTPResponseRuleTypeCacheStore string = "cache-store"

	// HTTPResponseRuleTypeCapture captures enum value "capture"
	HTTPResponseRuleTypeCapture string = "capture"

	// HTTPResponseRuleTypeDelACL captures enum value "del-acl"
	HTTPResponseRuleTypeDelACL string = "del-acl"

	// HTTPResponseRuleTypeDelHeader captures enum value "del-header"
	HTTPResponseRuleTypeDelHeader string = "del-header"

	// HTTPResponseRuleTypeDelMap captures enum value "del-map"
	HTTPResponseRuleTypeDelMap string = "del-map"

	// HTTPResponseRuleTypeDeny captures enum value "deny"
	HTTPResponseRuleTypeDeny string = "deny"

	// HTTPResponseRuleTypeRedirect captures enum value "redirect"
	HTTPResponseRuleTypeRedirect string = "redirect"

	// HTTPResponseRuleTypeReplaceHeader captures enum value "replace-header"
	HTTPResponseRuleTypeReplaceHeader string = "replace-header"

	// HTTPResponseRuleTypeReplaceValue captures enum value "replace-value"
	HTTPResponseRuleTypeReplaceValue string = "replace-value"

	// HTTPResponseRuleTypeReturn captures enum value "return"
	HTTPResponseRuleTypeReturn string = "return"

	// HTTPResponseRuleTypeScIncGpc0 captures enum value "sc-inc-gpc0"
	HTTPResponseRuleTypeScIncGpc0 string = "sc-inc-gpc0"

	// HTTPResponseRuleTypeScIncGpc1 captures enum value "sc-inc-gpc1"
	HTTPResponseRuleTypeScIncGpc1 string = "sc-inc-gpc1"

	// HTTPResponseRuleTypeScSetGpt0 captures enum value "sc-set-gpt0"
	HTTPResponseRuleTypeScSetGpt0 string = "sc-set-gpt0"

	// HTTPResponseRuleTypeSendSpoeGroup captures enum value "send-spoe-group"
	HTTPResponseRuleTypeSendSpoeGroup string = "send-spoe-group"

	// HTTPResponseRuleTypeSetHeader captures enum value "set-header"
	HTTPResponseRuleTypeSetHeader string = "set-header"

	// HTTPResponseRuleTypeSetLogLevel captures enum value "set-log-level"
	HTTPResponseRuleTypeSetLogLevel string = "set-log-level"

	// HTTPResponseRuleTypeSetMap captures enum value "set-map"
	HTTPResponseRuleTypeSetMap string = "set-map"

	// HTTPResponseRuleTypeSetMark captures enum value "set-mark"
	HTTPResponseRuleTypeSetMark string = "set-mark"

	// HTTPResponseRuleTypeSetNice captures enum value "set-nice"
	HTTPResponseRuleTypeSetNice string = "set-nice"

	// HTTPResponseRuleTypeSetStatus captures enum value "set-status"
	HTTPResponseRuleTypeSetStatus string = "set-status"

	// HTTPResponseRuleTypeSetTos captures enum value "set-tos"
	HTTPResponseRuleTypeSetTos string = "set-tos"

	// HTTPResponseRuleTypeSetVar captures enum value "set-var"
	HTTPResponseRuleTypeSetVar string = "set-var"

	// HTTPResponseRuleTypeSetVarFmt captures enum value "set-var-fmt"
	HTTPResponseRuleTypeSetVarFmt string = "set-var-fmt"

	// HTTPResponseRuleTypeSilentDrop captures enum value "silent-drop"
	HTTPResponseRuleTypeSilentDrop string = "silent-drop"

	// HTTPResponseRuleTypeStrictMode captures enum value "strict-mode"
	HTTPResponseRuleTypeStrictMode string = "strict-mode"

	// HTTPResponseRuleTypeTrackSc0 captures enum value "track-sc0"
	HTTPResponseRuleTypeTrackSc0 string = "track-sc0"

	// HTTPResponseRuleTypeTrackSc1 captures enum value "track-sc1"
	HTTPResponseRuleTypeTrackSc1 string = "track-sc1"

	// HTTPResponseRuleTypeTrackSc2 captures enum value "track-sc2"
	HTTPResponseRuleTypeTrackSc2 string = "track-sc2"

	// HTTPResponseRuleTypeUnsetVar captures enum value "unset-var"
	HTTPResponseRuleTypeUnsetVar string = "unset-var"

	// HTTPResponseRuleTypeWaitForBodyTime captures enum value "wait-for-body time"
	HTTPResponseRuleTypeWaitForBodyTime string = "wait-for-body time"
)

// prop value enum
func (m *HTTPResponseRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarName(formats strfmt.Registry) error {

	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", string(m.VarName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarScope(formats strfmt.Registry) error {

	if swag.IsZero(m.VarScope) { // not required
		return nil
	}

	if err := validate.Pattern("var_scope", "body", string(m.VarScope), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPResponseRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPResponseRule) UnmarshalBinary(b []byte) error {
	var res HTTPResponseRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HTTPResponseRuleReturnHdrsItems0 HTTP response rule return hdrs items0
//
// swagger:model HTTPResponseRuleReturnHdrsItems0
type HTTPResponseRuleReturnHdrsItems0 struct {

	// fmt
	// Required: true
	Fmt *string `json:"fmt"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this HTTP response rule return hdrs items0
func (m *HTTPResponseRuleReturnHdrsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFmt(formats); err != nil {
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

func (m *HTTPResponseRuleReturnHdrsItems0) validateFmt(formats strfmt.Registry) error {

	if err := validate.Required("fmt", "body", m.Fmt); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRuleReturnHdrsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPResponseRuleReturnHdrsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPResponseRuleReturnHdrsItems0) UnmarshalBinary(b []byte) error {
	var res HTTPResponseRuleReturnHdrsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
