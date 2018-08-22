// Code generated by protoc-gen-validate
// source: envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on HttpConnectionManager with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpConnectionManager) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := HttpConnectionManager_CodecType_name[int32(m.GetCodecType())]; !ok {
		return HttpConnectionManagerValidationError{
			Field:  "CodecType",
			Reason: "value must be one of the defined enum values",
		}
	}

	if len(m.GetStatPrefix()) < 1 {
		return HttpConnectionManagerValidationError{
			Field:  "StatPrefix",
			Reason: "value length must be at least 1 bytes",
		}
	}

	for idx, item := range m.GetHttpFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return HttpConnectionManagerValidationError{
					Field:  fmt.Sprintf("HttpFilters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetAddUserAgent()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "AddUserAgent",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTracing()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "Tracing",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHttpProtocolOptions()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "HttpProtocolOptions",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "Http2ProtocolOptions",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for ServerName

	if v, ok := interface{}(m.GetIdleTimeout()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "IdleTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStreamIdleTimeout()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "StreamIdleTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDrainTimeout()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "DrainTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetAccessLog() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return HttpConnectionManagerValidationError{
					Field:  fmt.Sprintf("AccessLog[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUseRemoteAddress()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "UseRemoteAddress",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for XffNumTrustedHops

	// no validation rules for SkipXffAppend

	// no validation rules for Via

	if v, ok := interface{}(m.GetGenerateRequestId()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "GenerateRequestId",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if _, ok := HttpConnectionManager_ForwardClientCertDetails_name[int32(m.GetForwardClientCertDetails())]; !ok {
		return HttpConnectionManagerValidationError{
			Field:  "ForwardClientCertDetails",
			Reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetSetCurrentClientCertDetails()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManagerValidationError{
				Field:  "SetCurrentClientCertDetails",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for Proxy_100Continue

	// no validation rules for RepresentIpv4RemoteAddressAsIpv4MappedIpv6

	for idx, item := range m.GetUpgradeConfigs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return HttpConnectionManagerValidationError{
					Field:  fmt.Sprintf("UpgradeConfigs[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	switch m.RouteSpecifier.(type) {

	case *HttpConnectionManager_Rds:

		if v, ok := interface{}(m.GetRds()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return HttpConnectionManagerValidationError{
					Field:  "Rds",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *HttpConnectionManager_RouteConfig:

		if v, ok := interface{}(m.GetRouteConfig()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return HttpConnectionManagerValidationError{
					Field:  "RouteConfig",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return HttpConnectionManagerValidationError{
			Field:  "RouteSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// HttpConnectionManagerValidationError is the validation error returned by
// HttpConnectionManager.Validate if the designated constraints aren't met.
type HttpConnectionManagerValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpConnectionManagerValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpConnectionManager.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpConnectionManagerValidationError{}

// Validate checks the field values on Rds with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Rds) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetConfigSource()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return RdsValidationError{
				Field:  "ConfigSource",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if len(m.GetRouteConfigName()) < 1 {
		return RdsValidationError{
			Field:  "RouteConfigName",
			Reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// RdsValidationError is the validation error returned by Rds.Validate if the
// designated constraints aren't met.
type RdsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RdsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRds.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RdsValidationError{}

// Validate checks the field values on HttpFilter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HttpFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return HttpFilterValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpFilterValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDeprecatedV1()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpFilterValidationError{
				Field:  "DeprecatedV1",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HttpFilterValidationError is the validation error returned by
// HttpFilter.Validate if the designated constraints aren't met.
type HttpFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpFilterValidationError{}

// Validate checks the field values on HttpConnectionManager_Tracing with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpConnectionManager_Tracing) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := HttpConnectionManager_Tracing_OperationName_name[int32(m.GetOperationName())]; !ok {
		return HttpConnectionManager_TracingValidationError{
			Field:  "OperationName",
			Reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetClientSampling()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManager_TracingValidationError{
				Field:  "ClientSampling",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRandomSampling()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManager_TracingValidationError{
				Field:  "RandomSampling",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOverallSampling()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManager_TracingValidationError{
				Field:  "OverallSampling",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HttpConnectionManager_TracingValidationError is the validation error
// returned by HttpConnectionManager_Tracing.Validate if the designated
// constraints aren't met.
type HttpConnectionManager_TracingValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpConnectionManager_TracingValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpConnectionManager_Tracing.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpConnectionManager_TracingValidationError{}

// Validate checks the field values on
// HttpConnectionManager_SetCurrentClientCertDetails with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpConnectionManager_SetCurrentClientCertDetails) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSubject()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpConnectionManager_SetCurrentClientCertDetailsValidationError{
				Field:  "Subject",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for Cert

	// no validation rules for Dns

	// no validation rules for Uri

	return nil
}

// HttpConnectionManager_SetCurrentClientCertDetailsValidationError is the
// validation error returned by
// HttpConnectionManager_SetCurrentClientCertDetails.Validate if the
// designated constraints aren't met.
type HttpConnectionManager_SetCurrentClientCertDetailsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpConnectionManager_SetCurrentClientCertDetailsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpConnectionManager_SetCurrentClientCertDetails.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpConnectionManager_SetCurrentClientCertDetailsValidationError{}

// Validate checks the field values on HttpConnectionManager_UpgradeConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *HttpConnectionManager_UpgradeConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UpgradeType

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return HttpConnectionManager_UpgradeConfigValidationError{
					Field:  fmt.Sprintf("Filters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpConnectionManager_UpgradeConfigValidationError is the validation error
// returned by HttpConnectionManager_UpgradeConfig.Validate if the designated
// constraints aren't met.
type HttpConnectionManager_UpgradeConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpConnectionManager_UpgradeConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpConnectionManager_UpgradeConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpConnectionManager_UpgradeConfigValidationError{}

// Validate checks the field values on HttpFilter_DeprecatedV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpFilter_DeprecatedV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	return nil
}

// HttpFilter_DeprecatedV1ValidationError is the validation error returned by
// HttpFilter_DeprecatedV1.Validate if the designated constraints aren't met.
type HttpFilter_DeprecatedV1ValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpFilter_DeprecatedV1ValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpFilter_DeprecatedV1.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpFilter_DeprecatedV1ValidationError{}
