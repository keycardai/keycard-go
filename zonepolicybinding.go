// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/apiquery"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// Policy set binding management. Active slots are replace-only by design: once
// bound (at zone bootstrap or on a principal's first policy check), an active slot
// is never emptied — evaluation fails closed, so an empty active slot would deny
// all access. Replace the bound version rather than removing the binding. DELETE
// applies to shadow bindings only.
//
// ZonePolicyBindingService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicyBindingService] method instead.
type ZonePolicyBindingService struct {
	Options []option.RequestOption
}

// NewZonePolicyBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicyBindingService(opts ...option.RequestOption) (r ZonePolicyBindingService) {
	r = ZonePolicyBindingService{}
	r.Options = opts
	return
}

// Returns a specific binding by its stable identifier.
func (r *ZonePolicyBindingService) Get(ctx context.Context, bindingID string, params ZonePolicyBindingGetParams, opts ...option.RequestOption) (res *PolicySetBinding, err error) {
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
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-bindings/%s", url.PathEscape(params.ZoneID), url.PathEscape(bindingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns bindings for the zone, filtered by target and mode.
//
// `filter[target_type]` defaults to `zone` and `filter[mode]` defaults to `shadow`
// when omitted (preserving the endpoint's historical shadow-only listing); pass
// explicit values to widen. The deprecated bare `scope_type` and `mode` parameters
// are honored as equivalents and suppress the respective defaults; supplying a
// bare parameter and its `filter[...]` form with different values returns 400.
// With an effective target type set of exactly `zone`, `filter[target_id]` values
// other than the zone ID return 400.
func (r *ZonePolicyBindingService) List(ctx context.Context, zoneID string, params ZonePolicyBindingListParams, opts ...option.RequestOption) (res *ZonePolicyBindingListResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/policy-bindings", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Removes a binding by its identifier, stopping evaluation for that slot.
func (r *ZonePolicyBindingService) Delete(ctx context.Context, bindingID string, params ZonePolicyBindingDeleteParams, opts ...option.RequestOption) (err error) {
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
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/policy-bindings/%s", url.PathEscape(params.ZoneID), url.PathEscape(bindingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns shadow evaluation metrics for a specific binding. Use query parameters
// to change scope:
//
//   - `scope=zone` — zone-level aggregate across all pairings
//   - `psv_id=<id>` — metrics for a specific policy set version (in any role),
//     including per-pairing breakdown
func (r *ZonePolicyBindingService) GetMetrics(ctx context.Context, bindingID string, params ZonePolicyBindingGetMetricsParams, opts ...option.RequestOption) (res *ZonePolicyBindingGetMetricsResponse, err error) {
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
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-bindings/%s/metrics", url.PathEscape(params.ZoneID), url.PathEscape(bindingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Upserts a binding for a slot identified by (scope_type, scope_target_id, mode).
// If a binding already exists for the slot, it is replaced. The binding ID is
// stable per slot and survives replacements.
func (r *ZonePolicyBindingService) Set(ctx context.Context, zoneID string, params ZonePolicyBindingSetParams, opts ...option.RequestOption) (res *PolicySetBinding, err error) {
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
	path := fmt.Sprintf("zones/%s/policy-bindings", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type PolicySetBinding struct {
	// Binding identifier (stable per slot)
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Binding mode
	//
	// Any of "active", "shadow".
	Mode PolicySetBindingMode `json:"mode" api:"required"`
	// Public ID of the bound policy set
	PolicySetID string `json:"policy_set_id" api:"required"`
	// Public ID of the bound policy set version
	PolicySetVersionID string `json:"policy_set_version_id" api:"required"`
	// **Deprecated.** Use `target_id` instead. Carries the same value.
	//
	// Deprecated: deprecated
	ScopeTargetID string `json:"scope_target_id" api:"required"`
	// **Deprecated.** Use `target_type` instead. Carries the same value.
	//
	// Any of "zone".
	//
	// Deprecated: deprecated
	ScopeType PolicySetBindingScopeType `json:"scope_type" api:"required"`
	// Target entity ID. Equals zone_id for zone-targeted bindings.
	TargetID string `json:"target_id" api:"required"`
	// What this binding targets
	//
	// Any of "zone", "user".
	TargetType PolicySetBindingTargetType `json:"target_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Mode               respjson.Field
		PolicySetID        respjson.Field
		PolicySetVersionID respjson.Field
		ScopeTargetID      respjson.Field
		ScopeType          respjson.Field
		TargetID           respjson.Field
		TargetType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicySetBinding) RawJSON() string { return r.JSON.raw }
func (r *PolicySetBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binding mode
type PolicySetBindingMode string

const (
	PolicySetBindingModeActive PolicySetBindingMode = "active"
	PolicySetBindingModeShadow PolicySetBindingMode = "shadow"
)

// **Deprecated.** Use `target_type` instead. Carries the same value.
type PolicySetBindingScopeType string

const (
	PolicySetBindingScopeTypeZone PolicySetBindingScopeType = "zone"
)

// What this binding targets
type PolicySetBindingTargetType string

const (
	PolicySetBindingTargetTypeZone PolicySetBindingTargetType = "zone"
	PolicySetBindingTargetTypeUser PolicySetBindingTargetType = "user"
)

type ZonePolicyBindingListResponse struct {
	Items []PolicySetBinding `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyBindingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyBindingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyBindingGetMetricsResponse struct {
	AgreeCount                 int64     `json:"agree_count" api:"required"`
	ShadowAllowActiveDenyCount int64     `json:"shadow_allow_active_deny_count" api:"required"`
	ShadowDenyActiveAllowCount int64     `json:"shadow_deny_active_allow_count" api:"required"`
	Since                      time.Time `json:"since" api:"required" format:"date-time"`
	TotalEvaluations           int64     `json:"total_evaluations" api:"required"`
	// Fraction of evaluations where active and shadow decisions agreed (0.0–1.0)
	AgreementRate       float64 `json:"agreement_rate"`
	AvgShadowDurationMs int64   `json:"avg_shadow_duration_ms"`
	// Per-pairing breakdown. Present when querying by PSV.
	Pairings               []ZonePolicyBindingGetMetricsResponsePairing `json:"pairings"`
	StddevShadowDurationMs int64                                        `json:"stddev_shadow_duration_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreeCount                 respjson.Field
		ShadowAllowActiveDenyCount respjson.Field
		ShadowDenyActiveAllowCount respjson.Field
		Since                      respjson.Field
		TotalEvaluations           respjson.Field
		AgreementRate              respjson.Field
		AvgShadowDurationMs        respjson.Field
		Pairings                   respjson.Field
		StddevShadowDurationMs     respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyBindingGetMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyBindingGetMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyBindingGetMetricsResponsePairing struct {
	ActivePsvID string    `json:"active_psv_id" api:"required"`
	AgreeCount  int64     `json:"agree_count" api:"required"`
	FirstEval   time.Time `json:"first_eval" api:"required" format:"date-time"`
	LastEval    time.Time `json:"last_eval" api:"required" format:"date-time"`
	ShadowPsvID string    `json:"shadow_psv_id" api:"required"`
	Total       int64     `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivePsvID respjson.Field
		AgreeCount  respjson.Field
		FirstEval   respjson.Field
		LastEval    respjson.Field
		ShadowPsvID respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyBindingGetMetricsResponsePairing) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyBindingGetMetricsResponsePairing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyBindingGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicyBindingListParams struct {
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Filter on binding `mode`. Repeatable; repeated instances OR across values. See
	// `FilterValues` in the shared spec for the full wire convention.
	//
	// Allowed values: `active`, `shadow`. Unknown values return 400 with the list of
	// allowed values. Comma-separated single values are rejected with a 400 pointing
	// at the repeated-parameter OR form.
	//
	// **Defaults to `shadow`** when omitted (and the deprecated bare `mode` parameter
	// is absent), preserving the endpoint's historical shadow-only listing. Pass
	// `filter[mode]=active` (or both values) to widen.
	FilterMode []string `query:"filter[mode],omitzero" json:"-"`
	// Filter on `target_id`. Repeatable; repeated instances OR across values. See
	// `FilterValues` in the shared spec for the full wire convention.
	//
	// Zone-targeted elements carry the zone ID as their `target_id`. When the
	// effective `filter[target_type]` set is exactly `zone` (including via its
	// default), any `filter[target_id]` value other than the path `zone_id` returns
	// `400 Bad Request`.
	FilterTargetID []string `query:"filter[target_id],omitzero" json:"-"`
	// Filter on `target_type`. Repeatable; repeated instances OR across values. See
	// `FilterValues` in the shared spec for the full wire convention.
	//
	// Allowed values: `zone`, `user` (`resource` and `session` are reserved and not
	// yet accepted). Unknown values return 400 with the list of allowed values.
	// Comma-separated single values are rejected with a 400 pointing at the
	// repeated-parameter OR form.
	//
	// **Defaults to `zone`** when omitted (and no deprecated equivalent parameter is
	// supplied), so listings exclude principal-scoped elements unless explicitly
	// widened. On `listPolicies` the default is skipped when `filter[id]` is present,
	// so by-ID fetches resolve regardless of target.
	//
	// Note: the allowed-value enum is enforced in the handler (not as an OpenAPI
	// `items.enum`) so the server can return a targeted error for the comma-AND form
	// instead of a generic "not in allowed values" response.
	FilterTargetType []string `query:"filter[target_type],omitzero" json:"-"`
	// **Deprecated.** Use `filter[mode]` instead. Still honored for backward
	// compatibility but only accepts `shadow`; other values return 400. Conflicts with
	// `filter[mode]` return 400.
	//
	// Any of "shadow".
	Mode ZonePolicyBindingListParamsMode `query:"mode,omitzero" json:"-"`
	// **Deprecated.** Use `filter[target_type]` instead. Still honored for backward
	// compatibility but only accepts `zone`; other values return 400. Conflicts with
	// `filter[target_type]` return 400.
	//
	// Any of "zone".
	ScopeType ZonePolicyBindingListParamsScopeType `query:"scope_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicyBindingListParams]'s query parameters as
// `url.Values`.
func (r ZonePolicyBindingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// **Deprecated.** Use `filter[mode]` instead. Still honored for backward
// compatibility but only accepts `shadow`; other values return 400. Conflicts with
// `filter[mode]` return 400.
type ZonePolicyBindingListParamsMode string

const (
	ZonePolicyBindingListParamsModeShadow ZonePolicyBindingListParamsMode = "shadow"
)

// **Deprecated.** Use `filter[target_type]` instead. Still honored for backward
// compatibility but only accepts `zone`; other values return 400. Conflicts with
// `filter[target_type]` return 400.
type ZonePolicyBindingListParamsScopeType string

const (
	ZonePolicyBindingListParamsScopeTypeZone ZonePolicyBindingListParamsScopeType = "zone"
)

type ZonePolicyBindingDeleteParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicyBindingGetMetricsParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// Policy set to filter metrics for. Required when psv_id is provided.
	PolicySetID param.Opt[string] `query:"policy_set_id,omitzero" json:"-"`
	// Filter to a specific policy set version (in any role). When provided, response
	// includes per-pairing breakdown. Requires policy_set_id.
	PsvID param.Opt[string] `query:"psv_id,omitzero" json:"-"`
	// Start of the time window (default 24h ago)
	Since            param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	XAPIVersion      param.Opt[string]    `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string]    `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Metrics scope. `binding` (default) returns metrics for this binding. `zone`
	// returns zone-level aggregate across all pairings.
	//
	// Any of "binding", "zone".
	Scope ZonePolicyBindingGetMetricsParamsScope `query:"scope,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicyBindingGetMetricsParams]'s query parameters as
// `url.Values`.
func (r ZonePolicyBindingGetMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metrics scope. `binding` (default) returns metrics for this binding. `zone`
// returns zone-level aggregate across all pairings.
type ZonePolicyBindingGetMetricsParamsScope string

const (
	ZonePolicyBindingGetMetricsParamsScopeBinding ZonePolicyBindingGetMetricsParamsScope = "binding"
	ZonePolicyBindingGetMetricsParamsScopeZone    ZonePolicyBindingGetMetricsParamsScope = "zone"
)

type ZonePolicyBindingSetParams struct {
	// Binding mode
	//
	// Any of "shadow".
	Mode ZonePolicyBindingSetParamsMode `json:"mode,omitzero" api:"required"`
	// Public ID of the policy set to bind
	PolicySetID string `json:"policy_set_id" api:"required"`
	// Public ID of the policy set version to bind
	PolicySetVersionID string `json:"policy_set_version_id" api:"required"`
	// **Deprecated.** Use `target_id` instead.
	ScopeTargetID param.Opt[string] `json:"scope_target_id,omitzero"`
	// Defaults to zone_id for zone targets. Required for other targets.
	TargetID         param.Opt[string] `json:"target_id,omitzero"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// **Deprecated.** Use `target_type` instead.
	//
	// Any of "zone".
	ScopeType ZonePolicyBindingSetParamsScopeType `json:"scope_type,omitzero"`
	// What the binding targets. Only `zone` is supported for direct binding writes.
	//
	// Any of "zone".
	TargetType ZonePolicyBindingSetParamsTargetType `json:"target_type,omitzero"`
	paramObj
}

func (r ZonePolicyBindingSetParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyBindingSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyBindingSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binding mode
type ZonePolicyBindingSetParamsMode string

const (
	ZonePolicyBindingSetParamsModeShadow ZonePolicyBindingSetParamsMode = "shadow"
)

// **Deprecated.** Use `target_type` instead.
type ZonePolicyBindingSetParamsScopeType string

const (
	ZonePolicyBindingSetParamsScopeTypeZone ZonePolicyBindingSetParamsScopeType = "zone"
)

// What the binding targets. Only `zone` is supported for direct binding writes.
type ZonePolicyBindingSetParamsTargetType string

const (
	ZonePolicyBindingSetParamsTargetTypeZone ZonePolicyBindingSetParamsTargetType = "zone"
)
