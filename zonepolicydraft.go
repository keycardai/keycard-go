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
	"github.com/keycardai/keycard-go/internal/apiquery"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
)

// Mutable policy authoring workspace
//
// ZonePolicyDraftService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicyDraftService] method instead.
type ZonePolicyDraftService struct {
	Options []option.RequestOption
}

// NewZonePolicyDraftService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePolicyDraftService(opts ...option.RequestOption) (r ZonePolicyDraftService) {
	r = ZonePolicyDraftService{}
	r.Options = opts
	return
}

// Get the draft for a policy
func (r *ZonePolicyDraftService) Get(ctx context.Context, policyID string, params ZonePolicyDraftGetParams, opts ...option.RequestOption) (res *PolicyDraft, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policies/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a policy draft
func (r *ZonePolicyDraftService) Delete(ctx context.Context, policyID string, params ZonePolicyDraftDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/policies/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Create or update a policy draft
func (r *ZonePolicyDraftService) Upsert(ctx context.Context, policyID string, params ZonePolicyDraftUpsertParams, opts ...option.RequestOption) (res *PolicyDraft, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policies/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type ZonePolicyDraftGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Opt-in to additional response fields on a single resource (`user`). Repeatable.
	Expand ZonePolicyDraftGetParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Narrows which Cedar representation the response includes. When omitted, both
	// `cedar_json` and `cedar_raw` are populated. Pass `json` to receive only
	// `cedar_json`, or `cedar` to receive only `cedar_raw`. Callers that don't care
	// about payload size can skip this parameter.
	//
	// Any of "cedar", "json".
	Format ZonePolicyDraftGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicyDraftGetParams]'s query parameters as
// `url.Values`.
func (r ZonePolicyDraftGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZonePolicyDraftGetParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZonePolicyDraftGetsExpandPdpExpandSingleField)
	OfZonePolicyDraftGetsExpandPdpExpandSingleField param.Opt[string] `query:",omitzero,inline"`
	OfZonePolicyDraftGetsExpandArrayItemArray       []string          `query:",omitzero,inline"`
	paramUnion
}

// Opt-in response field requested via `expand[]` on single-resource reads. `user`
// behaves as described on `ExpandField`.
type ZonePolicyDraftGetParamsExpandPdpExpandSingleField string

const (
	ZonePolicyDraftGetParamsExpandPdpExpandSingleFieldUser ZonePolicyDraftGetParamsExpandPdpExpandSingleField = "user"
)

// Narrows which Cedar representation the response includes. When omitted, both
// `cedar_json` and `cedar_raw` are populated. Pass `json` to receive only
// `cedar_json`, or `cedar` to receive only `cedar_raw`. Callers that don't care
// about payload size can skip this parameter.
type ZonePolicyDraftGetParamsFormat string

const (
	ZonePolicyDraftGetParamsFormatCedar ZonePolicyDraftGetParamsFormat = "cedar"
	ZonePolicyDraftGetParamsFormatJson  ZonePolicyDraftGetParamsFormat = "json"
)

type ZonePolicyDraftDeleteParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicyDraftUpsertParams struct {
	ZoneID        string `path:"zone_id" api:"required" json:"-"`
	SchemaVersion string `json:"schema_version" api:"required"`
	// Cedar policy in human-readable Cedar syntax. Mutually exclusive with cedar_json.
	CedarRaw param.Opt[string] `json:"cedar_raw,omitzero"`
	// Whether to validate cedar_json against the schema. Defaults to true.
	ValidateCedar    param.Opt[bool]   `query:"validate_cedar,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Cedar policy in JSON representation. Mutually exclusive with cedar_raw.
	CedarJson any `json:"cedar_json,omitzero"`
	// Response format override. `json` returns cedar_json, `cedar` returns cedar_raw.
	// When omitted, the response matches the input format.
	//
	// Any of "cedar", "json".
	Format ZonePolicyDraftUpsertParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

func (r ZonePolicyDraftUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyDraftUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyDraftUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ZonePolicyDraftUpsertParams]'s query parameters as
// `url.Values`.
func (r ZonePolicyDraftUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response format override. `json` returns cedar_json, `cedar` returns cedar_raw.
// When omitted, the response matches the input format.
type ZonePolicyDraftUpsertParamsFormat string

const (
	ZonePolicyDraftUpsertParamsFormatCedar ZonePolicyDraftUpsertParamsFormat = "cedar"
	ZonePolicyDraftUpsertParamsFormatJson  ZonePolicyDraftUpsertParamsFormat = "json"
)
