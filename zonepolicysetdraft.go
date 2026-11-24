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
	"github.com/keycardai/keycard-go/packages/respjson"
)

// Mutable policy set manifest workspace
//
// ZonePolicySetDraftService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicySetDraftService] method instead.
type ZonePolicySetDraftService struct {
	Options []option.RequestOption
}

// NewZonePolicySetDraftService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicySetDraftService(opts ...option.RequestOption) (r ZonePolicySetDraftService) {
	r = ZonePolicySetDraftService{}
	r.Options = opts
	return
}

// Safe draft creation that never overwrites an existing draft. If a draft already
// exists for this policy set, it is returned as-is (200 OK). Three body modes
// control how a new draft is initialized: (1) empty body — hydrate from the latest
// non-archived version, (2) { base_version_id } — hydrate from a specific version,
// (3) full body with manifest and schema_version — create with explicit content.
// base_version_id and manifest are mutually exclusive; providing both is rejected.
func (r *ZonePolicySetDraftService) New(ctx context.Context, policySetID string, params ZonePolicySetDraftNewParams, opts ...option.RequestOption) (res *PolicySetDraft, err error) {
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
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get the draft for a policy set
func (r *ZonePolicySetDraftService) Get(ctx context.Context, policySetID string, params ZonePolicySetDraftGetParams, opts ...option.RequestOption) (res *PolicySetDraft, err error) {
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
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Applies partial updates to an existing draft. Supports updating metadata fields
// (name, description, schema_version) and manifest operations (set, add, remove
// entries).
func (r *ZonePolicySetDraftService) Update(ctx context.Context, policySetID string, params ZonePolicySetDraftUpdateParams, opts ...option.RequestOption) (res *PolicySetDraft, err error) {
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
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete a policy set draft
func (r *ZonePolicySetDraftService) Delete(ctx context.Context, policySetID string, params ZonePolicySetDraftDeleteParams, opts ...option.RequestOption) (err error) {
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
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns the policy versions referenced by the draft's manifest as a paginated
// list. Mirrors the published
// `/policy-sets/{policy_set_id}/versions/{version_id}/policies` endpoint but
// sources entries from the policy set's mutable draft. Archived policy versions
// referenced by the manifest are still returned.
func (r *ZonePolicySetDraftService) ListPolicies(ctx context.Context, policySetID string, params ZonePolicySetDraftListPoliciesParams, opts ...option.RequestOption) (res *ZonePolicySetDraftListPoliciesResponse, err error) {
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
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/draft/policies", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Full replacement of the draft state. Creates the draft if it does not exist,
// replaces it if it does. Standard REST PUT idempotency — the caller provides the
// complete desired state.
func (r *ZonePolicySetDraftService) Upsert(ctx context.Context, policySetID string, params ZonePolicySetDraftUpsertParams, opts ...option.RequestOption) (res *PolicySetDraft, err error) {
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
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type ZonePolicySetDraftListPoliciesResponse struct {
	Items []PolicyVersion `json:"items" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination ZonePolicySetDraftListPoliciesResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetDraftListPoliciesResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetDraftListPoliciesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type ZonePolicySetDraftListPoliciesResponsePagination struct {
	// An opaque cursor used for paginating through a list of results
	AfterCursor string `json:"after_cursor" api:"required"`
	// An opaque cursor used for paginating through a list of results
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items across all pages. Only present when the request includes
	// ?expand[]=total_count.
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterCursor  respjson.Field
		BeforeCursor respjson.Field
		TotalCount   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetDraftListPoliciesResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetDraftListPoliciesResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetDraftNewParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// ID of a policy set version to hydrate from. Mutually exclusive with manifest —
	// the server rejects requests containing both fields.
	BaseVersionID    param.Opt[string]      `json:"base_version_id,omitzero"`
	Description      param.Opt[string]      `json:"description,omitzero"`
	Name             param.Opt[string]      `json:"name,omitzero"`
	SchemaVersion    param.Opt[string]      `json:"schema_version,omitzero"`
	XAPIVersion      param.Opt[string]      `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string]      `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	Manifest         PolicySetManifestParam `json:"manifest,omitzero"`
	paramObj
}

func (r ZonePolicySetDraftNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetDraftNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetDraftNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetDraftGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Opt-in to additional response fields on a single resource (`user`). Repeatable.
	Expand ZonePolicySetDraftGetParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySetDraftGetParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySetDraftGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZonePolicySetDraftGetParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZonePolicySetDraftGetsExpandPdpExpandSingleField)
	OfZonePolicySetDraftGetsExpandPdpExpandSingleField param.Opt[string] `query:",omitzero,inline"`
	OfZonePolicySetDraftGetsExpandArrayItemArray       []string          `query:",omitzero,inline"`
	paramUnion
}

// Opt-in response field requested via `expand[]` on single-resource reads. `user`
// behaves as described on `ExpandField`.
type ZonePolicySetDraftGetParamsExpandPdpExpandSingleField string

const (
	ZonePolicySetDraftGetParamsExpandPdpExpandSingleFieldUser ZonePolicySetDraftGetParamsExpandPdpExpandSingleField = "user"
)

type ZonePolicySetDraftUpdateParams struct {
	ZoneID             string                                            `path:"zone_id" api:"required" json:"-"`
	Description        param.Opt[string]                                 `json:"description,omitzero"`
	Name               param.Opt[string]                                 `json:"name,omitzero"`
	SchemaVersion      param.Opt[string]                                 `json:"schema_version,omitzero"`
	XAPIVersion        param.Opt[string]                                 `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID   param.Opt[string]                                 `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	ManifestOperations []ZonePolicySetDraftUpdateParamsManifestOperation `json:"manifest_operations,omitzero"`
	paramObj
}

func (r ZonePolicySetDraftUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetDraftUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetDraftUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Op, PolicyID are required.
type ZonePolicySetDraftUpdateParamsManifestOperation struct {
	// set: Replace policy_version_id for an existing entry matching policy_id. add:
	// Append a new entry with the given policy_id and policy_version_id. remove:
	// Remove the entry matching policy_id.
	//
	// Any of "set", "add", "remove".
	Op       string `json:"op,omitzero" api:"required"`
	PolicyID string `json:"policy_id" api:"required"`
	// Required for set and add operations. Ignored for remove.
	PolicyVersionID param.Opt[string] `json:"policy_version_id,omitzero"`
	paramObj
}

func (r ZonePolicySetDraftUpdateParamsManifestOperation) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetDraftUpdateParamsManifestOperation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetDraftUpdateParamsManifestOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZonePolicySetDraftUpdateParamsManifestOperation](
		"op", "set", "add", "remove",
	)
}

type ZonePolicySetDraftDeleteParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicySetDraftListPoliciesParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// Cursor for forward pagination. Returned in `Pagination.after_cursor`. Mutually
	// exclusive with `before`.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returned in `Pagination.before_cursor`. Mutually
	// exclusive with `after`.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return per page.
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Opt-in to additional response fields (`total_count`, `user`). Repeatable;
	// matches the `expand[]` convention used across the Keycard API.
	Expand ZonePolicySetDraftListPoliciesParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Narrows which Cedar representation the response includes. When omitted, both
	// `cedar_json` and `cedar_raw` are populated. Pass `json` to receive only
	// `cedar_json`, or `cedar` to receive only `cedar_raw`. Callers that don't care
	// about payload size can skip this parameter.
	//
	// Any of "cedar", "json".
	Format ZonePolicySetDraftListPoliciesParamsFormat `query:"format,omitzero" json:"-"`
	// Sort direction. Default is desc (newest first).
	//
	// Any of "asc", "desc".
	Order ZonePolicySetDraftListPoliciesParamsOrder `query:"order,omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicySetDraftListPoliciesParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySetDraftListPoliciesParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySetDraftListPoliciesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZonePolicySetDraftListPoliciesParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZonePolicySetDraftListPoliciessExpandPdpExpandField)
	OfZonePolicySetDraftListPoliciessExpandPdpExpandField param.Opt[string] `query:",omitzero,inline"`
	OfZonePolicySetDraftListPoliciessExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Opt-in response field requested via `expand[]` on list endpoints. `total_count`
// adds `pagination.total_count`; `user` adds a `*_user` object beside every
// `created_by`, `updated_by` and `archived_by` field that resolves to an
// organization user.
type ZonePolicySetDraftListPoliciesParamsExpandPdpExpandField string

const (
	ZonePolicySetDraftListPoliciesParamsExpandPdpExpandFieldTotalCount ZonePolicySetDraftListPoliciesParamsExpandPdpExpandField = "total_count"
	ZonePolicySetDraftListPoliciesParamsExpandPdpExpandFieldUser       ZonePolicySetDraftListPoliciesParamsExpandPdpExpandField = "user"
)

// Narrows which Cedar representation the response includes. When omitted, both
// `cedar_json` and `cedar_raw` are populated. Pass `json` to receive only
// `cedar_json`, or `cedar` to receive only `cedar_raw`. Callers that don't care
// about payload size can skip this parameter.
type ZonePolicySetDraftListPoliciesParamsFormat string

const (
	ZonePolicySetDraftListPoliciesParamsFormatCedar ZonePolicySetDraftListPoliciesParamsFormat = "cedar"
	ZonePolicySetDraftListPoliciesParamsFormatJson  ZonePolicySetDraftListPoliciesParamsFormat = "json"
)

// Sort direction. Default is desc (newest first).
type ZonePolicySetDraftListPoliciesParamsOrder string

const (
	ZonePolicySetDraftListPoliciesParamsOrderAsc  ZonePolicySetDraftListPoliciesParamsOrder = "asc"
	ZonePolicySetDraftListPoliciesParamsOrderDesc ZonePolicySetDraftListPoliciesParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicySetDraftListPoliciesParamsSort string

const (
	ZonePolicySetDraftListPoliciesParamsSortCreatedAt ZonePolicySetDraftListPoliciesParamsSort = "created_at"
)

type ZonePolicySetDraftUpsertParams struct {
	ZoneID           string                 `path:"zone_id" api:"required" json:"-"`
	Manifest         PolicySetManifestParam `json:"manifest,omitzero" api:"required"`
	SchemaVersion    string                 `json:"schema_version" api:"required"`
	Description      param.Opt[string]      `json:"description,omitzero"`
	Name             param.Opt[string]      `json:"name,omitzero"`
	XAPIVersion      param.Opt[string]      `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string]      `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ZonePolicySetDraftUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetDraftUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetDraftUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
