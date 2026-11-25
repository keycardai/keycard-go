// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// Stateless Cedar policy validation
//
// ZonePolicyValidationService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicyValidationService] method instead.
type ZonePolicyValidationService struct {
	Options []option.RequestOption
}

// NewZonePolicyValidationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicyValidationService(opts ...option.RequestOption) (r ZonePolicyValidationService) {
	r = ZonePolicyValidationService{}
	r.Options = opts
	return
}

// Validate Cedar policy content against a schema
func (r *ZonePolicyValidationService) New(ctx context.Context, zoneID string, params ZonePolicyValidationNewParams, opts ...option.RequestOption) (res *ZonePolicyValidationNewResponse, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-validations", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type ZonePolicyValidationNewResponse struct {
	// Validation errors. Empty array when valid.
	Errors []ZonePolicyValidationNewResponseError `json:"errors" api:"required"`
	// Schema version validated against (echo back).
	SchemaVersion string `json:"schema_version" api:"required"`
	// True if the policy passes schema validation with no errors.
	Valid bool `json:"valid" api:"required"`
	// Validation warnings. May be non-empty even when valid is true. References to
	// actions absent from the schema are reported here as per-occurrence warnings with
	// code `unknown_action` (partitioned out of errors: they do not make valid false).
	Warnings []ZonePolicyValidationNewResponseWarning `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors        respjson.Field
		SchemaVersion respjson.Field
		Valid         respjson.Field
		Warnings      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyValidationNewResponseError struct {
	// Human-readable description of the validation issue.
	Message string `json:"message" api:"required"`
	// Machine-readable kind of the diagnostic (e.g. `unknown_action`, `parse_error`).
	// Optional and additive; absent when the finding has no specific code.
	Code string `json:"code"`
	// Optional structured details keyed by the diagnostic `code`, so consumers need
	// not parse `message`. For `unknown_action`: `{"action_uid": "..."}`.
	Data map[string]string `json:"data"`
	// Optional advisory text, rendered by clients as secondary text below the primary
	// message.
	Help string `json:"help"`
	// Half-open source range [start, end) locating a diagnostic in the submitted
	// policy text. Optional: present only when the engine can map the finding to a
	// span. Populated for parse/syntax errors and for semantic findings when the
	// policy was submitted as `cedar_raw`; absent for whole-policy semantic findings
	// (no single span) and for policies submitted as `cedar_json` (the JSON AST
	// carries no source locations).
	Range ZonePolicyValidationNewResponseErrorRange `json:"range"`
	// Severity of the diagnostic. Optional and additive; older clients that only read
	// `message` can ignore it. Diagnostics under `errors` are `error`; those under
	// `warnings` are `warning`.
	//
	// Any of "error", "warning".
	Severity string `json:"severity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Code        respjson.Field
		Data        respjson.Field
		Help        respjson.Field
		Range       respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Half-open source range [start, end) locating a diagnostic in the submitted
// policy text. Optional: present only when the engine can map the finding to a
// span. Populated for parse/syntax errors and for semantic findings when the
// policy was submitted as `cedar_raw`; absent for whole-policy semantic findings
// (no single span) and for policies submitted as `cedar_json` (the JSON AST
// carries no source locations).
type ZonePolicyValidationNewResponseErrorRange struct {
	// A position in the submitted policy text.
	End ZonePolicyValidationNewResponseErrorRangeEnd `json:"end" api:"required"`
	// A position in the submitted policy text.
	Start ZonePolicyValidationNewResponseErrorRangeStart `json:"start" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseErrorRange) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseErrorRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A position in the submitted policy text.
type ZonePolicyValidationNewResponseErrorRangeEnd struct {
	// 1-based column number (counts Unicode scalar values).
	Column int64 `json:"column" api:"required"`
	// 1-based line number.
	Line int64 `json:"line" api:"required"`
	// 0-based byte offset into the UTF-8 policy text.
	Offset int64 `json:"offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Line        respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseErrorRangeEnd) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseErrorRangeEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A position in the submitted policy text.
type ZonePolicyValidationNewResponseErrorRangeStart struct {
	// 1-based column number (counts Unicode scalar values).
	Column int64 `json:"column" api:"required"`
	// 1-based line number.
	Line int64 `json:"line" api:"required"`
	// 0-based byte offset into the UTF-8 policy text.
	Offset int64 `json:"offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Line        respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseErrorRangeStart) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseErrorRangeStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyValidationNewResponseWarning struct {
	// Human-readable description of the validation issue.
	Message string `json:"message" api:"required"`
	// Machine-readable kind of the diagnostic (e.g. `unknown_action`, `parse_error`).
	// Optional and additive; absent when the finding has no specific code.
	Code string `json:"code"`
	// Optional structured details keyed by the diagnostic `code`, so consumers need
	// not parse `message`. For `unknown_action`: `{"action_uid": "..."}`.
	Data map[string]string `json:"data"`
	// Optional advisory text, rendered by clients as secondary text below the primary
	// message.
	Help string `json:"help"`
	// Half-open source range [start, end) locating a diagnostic in the submitted
	// policy text. Optional: present only when the engine can map the finding to a
	// span. Populated for parse/syntax errors and for semantic findings when the
	// policy was submitted as `cedar_raw`; absent for whole-policy semantic findings
	// (no single span) and for policies submitted as `cedar_json` (the JSON AST
	// carries no source locations).
	Range ZonePolicyValidationNewResponseWarningRange `json:"range"`
	// Severity of the diagnostic. Optional and additive; older clients that only read
	// `message` can ignore it. Diagnostics under `errors` are `error`; those under
	// `warnings` are `warning`.
	//
	// Any of "error", "warning".
	Severity string `json:"severity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Code        respjson.Field
		Data        respjson.Field
		Help        respjson.Field
		Range       respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Half-open source range [start, end) locating a diagnostic in the submitted
// policy text. Optional: present only when the engine can map the finding to a
// span. Populated for parse/syntax errors and for semantic findings when the
// policy was submitted as `cedar_raw`; absent for whole-policy semantic findings
// (no single span) and for policies submitted as `cedar_json` (the JSON AST
// carries no source locations).
type ZonePolicyValidationNewResponseWarningRange struct {
	// A position in the submitted policy text.
	End ZonePolicyValidationNewResponseWarningRangeEnd `json:"end" api:"required"`
	// A position in the submitted policy text.
	Start ZonePolicyValidationNewResponseWarningRangeStart `json:"start" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseWarningRange) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseWarningRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A position in the submitted policy text.
type ZonePolicyValidationNewResponseWarningRangeEnd struct {
	// 1-based column number (counts Unicode scalar values).
	Column int64 `json:"column" api:"required"`
	// 1-based line number.
	Line int64 `json:"line" api:"required"`
	// 0-based byte offset into the UTF-8 policy text.
	Offset int64 `json:"offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Line        respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseWarningRangeEnd) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseWarningRangeEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A position in the submitted policy text.
type ZonePolicyValidationNewResponseWarningRangeStart struct {
	// 1-based column number (counts Unicode scalar values).
	Column int64 `json:"column" api:"required"`
	// 1-based line number.
	Line int64 `json:"line" api:"required"`
	// 0-based byte offset into the UTF-8 policy text.
	Offset int64 `json:"offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Line        respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyValidationNewResponseWarningRangeStart) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyValidationNewResponseWarningRangeStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyValidationNewParams struct {
	// Schema version to validate against. Must not be archived.
	SchemaVersion string `json:"schema_version" api:"required"`
	// Cedar policy in human-readable Cedar syntax. Mutually exclusive with cedar_json.
	CedarRaw         param.Opt[string] `json:"cedar_raw,omitzero"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Cedar policy in JSON representation. Mutually exclusive with cedar_raw.
	CedarJson any `json:"cedar_json,omitzero"`
	paramObj
}

func (r ZonePolicyValidationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyValidationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyValidationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
