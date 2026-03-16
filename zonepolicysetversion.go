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

// Immutable policy set manifest snapshots
//
// ZonePolicySetVersionService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicySetVersionService] method instead.
type ZonePolicySetVersionService struct {
	Options []option.RequestOption
}

// NewZonePolicySetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicySetVersionService(opts ...option.RequestOption) (r ZonePolicySetVersionService) {
	r = ZonePolicySetVersionService{}
	r.Options = opts
	return
}

// Validates the manifest, computes SHA, and creates an immutable version snapshot.
func (r *ZonePolicySetVersionService) New(ctx context.Context, policySetID string, params ZonePolicySetVersionNewParams, opts ...option.RequestOption) (res *PolicySetVersion, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/versions", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a specific policy set version
func (r *ZonePolicySetVersionService) Get(ctx context.Context, versionID string, params ZonePolicySetVersionGetParams, opts ...option.RequestOption) (res *PolicySetVersion, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if params.PolicySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/versions/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.PolicySetID), url.PathEscape(versionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set active=true to bind this version as the active zone policy set.
func (r *ZonePolicySetVersionService) Update(ctx context.Context, versionID string, params ZonePolicySetVersionUpdateParams, opts ...option.RequestOption) (res *PolicySetVersion, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if params.PolicySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/versions/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.PolicySetID), url.PathEscape(versionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List versions of a policy set
func (r *ZonePolicySetVersionService) List(ctx context.Context, policySetID string, params ZonePolicySetVersionListParams, opts ...option.RequestOption) (res *ZonePolicySetVersionListResponse, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/versions", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Archive a policy set version
func (r *ZonePolicySetVersionService) Archive(ctx context.Context, versionID string, params ZonePolicySetVersionArchiveParams, opts ...option.RequestOption) (res *PolicySetVersion, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if params.PolicySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/versions/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.PolicySetID), url.PathEscape(versionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the policy versions referenced by this policy set version's manifest as
// a paginated list.
func (r *ZonePolicySetVersionService) ListPolicies(ctx context.Context, versionID string, params ZonePolicySetVersionListPoliciesParams, opts ...option.RequestOption) (res *ZonePolicySetVersionListPoliciesResponse, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if params.PolicySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s/versions/%s/policies", url.PathEscape(params.ZoneID), url.PathEscape(params.PolicySetID), url.PathEscape(versionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type PolicySetVersion struct {
	ID        string            `json:"id" api:"required"`
	CreatedAt time.Time         `json:"created_at" api:"required" format:"date-time"`
	CreatedBy string            `json:"created_by" api:"required"`
	Manifest  PolicySetManifest `json:"manifest" api:"required"`
	// Hex-encoded SHA-256 of the canonicalized manifest
	ManifestSha   string `json:"manifest_sha" api:"required"`
	PolicySetID   string `json:"policy_set_id" api:"required"`
	SchemaVersion string `json:"schema_version" api:"required"`
	Version       int64  `json:"version" api:"required"`
	// Whether this policy set version is currently bound with mode='active'
	Active     bool      `json:"active"`
	ArchivedAt time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	ArchivedBy string    `json:"archived_by" api:"nullable"`
	// JWS Flattened JSON Serialization (RFC 7515 §7.2.2) of a policy set attestation.
	// The protected header carries the signing algorithm and key identifier; the
	// payload is a base64url-encoded AttestationStatement canonicalized per RFC 8785
	// (JCS). Verify using the zone JWKS endpoint (RFC 7517). Currently signed with
	// RS256; future zone key types (e.g. EdDSA) will be indicated by the "alg" header
	// — no envelope changes required.
	Attestation Attestation `json:"attestation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CreatedBy     respjson.Field
		Manifest      respjson.Field
		ManifestSha   respjson.Field
		PolicySetID   respjson.Field
		SchemaVersion respjson.Field
		Version       respjson.Field
		Active        respjson.Field
		ArchivedAt    respjson.Field
		ArchivedBy    respjson.Field
		Attestation   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicySetVersion) RawJSON() string { return r.JSON.raw }
func (r *PolicySetVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetVersionListResponse struct {
	Items      []PolicySetVersion                         `json:"items" api:"required"`
	Pagination ZonePolicySetVersionListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetVersionListResponsePagination struct {
	// Cursor of the last item on the current page. Pass to after for the next page.
	// Null when there is no next page.
	AfterCursor string `json:"after_cursor" api:"required"`
	// Cursor of the first item on the current page. Pass to before for the previous
	// page. Null when there is no previous page.
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items matching the current filters. Only included when
	// expand=total_count is requested.
	TotalCount int64 `json:"total_count" api:"nullable"`
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
func (r ZonePolicySetVersionListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetVersionListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetVersionListPoliciesResponse struct {
	Items      []PolicyVersion                                    `json:"items" api:"required"`
	Pagination ZonePolicySetVersionListPoliciesResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetVersionListPoliciesResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetVersionListPoliciesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetVersionListPoliciesResponsePagination struct {
	// Cursor of the last item on the current page. Pass to after for the next page.
	// Null when there is no next page.
	AfterCursor string `json:"after_cursor" api:"required"`
	// Cursor of the first item on the current page. Pass to before for the previous
	// page. Null when there is no previous page.
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items matching the current filters. Only included when
	// expand=total_count is requested.
	TotalCount int64 `json:"total_count" api:"nullable"`
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
func (r ZonePolicySetVersionListPoliciesResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetVersionListPoliciesResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetVersionNewParams struct {
	ZoneID           string                 `path:"zone_id" api:"required" json:"-"`
	Manifest         PolicySetManifestParam `json:"manifest,omitzero" api:"required"`
	SchemaVersion    string                 `json:"schema_version" api:"required"`
	XAPIVersion      param.Opt[string]      `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string]      `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ZonePolicySetVersionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetVersionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetVersionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetVersionGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	PolicySetID      string            `path:"policy_set_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicySetVersionUpdateParams struct {
	ZoneID      string `path:"zone_id" api:"required" json:"-"`
	PolicySetID string `path:"policy_set_id" api:"required" json:"-"`
	// Must be true. Binds this version as the active zone policy set.
	//
	// Any of true.
	Active           bool              `json:"active,omitzero" api:"required"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ZonePolicySetVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetVersionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetVersionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetVersionListParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// Return items after this cursor (forward pagination). Use after_cursor from a
	// previous response. Mutually exclusive with before.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Return items before this cursor (backward pagination). Use before_cursor from a
	// previous response. Mutually exclusive with after.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Opt-in to additional response fields
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Sort direction. Default is desc (newest first).
	//
	// Any of "asc", "desc".
	Order ZonePolicySetVersionListParamsOrder `query:"order,omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicySetVersionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySetVersionListParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySetVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction. Default is desc (newest first).
type ZonePolicySetVersionListParamsOrder string

const (
	ZonePolicySetVersionListParamsOrderAsc  ZonePolicySetVersionListParamsOrder = "asc"
	ZonePolicySetVersionListParamsOrderDesc ZonePolicySetVersionListParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicySetVersionListParamsSort string

const (
	ZonePolicySetVersionListParamsSortCreatedAt ZonePolicySetVersionListParamsSort = "created_at"
)

type ZonePolicySetVersionArchiveParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	PolicySetID      string            `path:"policy_set_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicySetVersionListPoliciesParams struct {
	ZoneID      string `path:"zone_id" api:"required" json:"-"`
	PolicySetID string `path:"policy_set_id" api:"required" json:"-"`
	// Return items after this cursor (forward pagination). Use after_cursor from a
	// previous response. Mutually exclusive with before.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Return items before this cursor (backward pagination). Use before_cursor from a
	// previous response. Mutually exclusive with after.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Opt-in to additional response fields
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Policy representation format. `json` returns cedar_json, `cedar` returns
	// cedar_raw.
	//
	// Any of "cedar", "json".
	Format ZonePolicySetVersionListPoliciesParamsFormat `query:"format,omitzero" json:"-"`
	// Sort direction. Default is desc (newest first).
	//
	// Any of "asc", "desc".
	Order ZonePolicySetVersionListPoliciesParamsOrder `query:"order,omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicySetVersionListPoliciesParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySetVersionListPoliciesParams]'s query parameters
// as `url.Values`.
func (r ZonePolicySetVersionListPoliciesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Policy representation format. `json` returns cedar_json, `cedar` returns
// cedar_raw.
type ZonePolicySetVersionListPoliciesParamsFormat string

const (
	ZonePolicySetVersionListPoliciesParamsFormatCedar ZonePolicySetVersionListPoliciesParamsFormat = "cedar"
	ZonePolicySetVersionListPoliciesParamsFormatJson  ZonePolicySetVersionListPoliciesParamsFormat = "json"
)

// Sort direction. Default is desc (newest first).
type ZonePolicySetVersionListPoliciesParamsOrder string

const (
	ZonePolicySetVersionListPoliciesParamsOrderAsc  ZonePolicySetVersionListPoliciesParamsOrder = "asc"
	ZonePolicySetVersionListPoliciesParamsOrderDesc ZonePolicySetVersionListPoliciesParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicySetVersionListPoliciesParamsSort string

const (
	ZonePolicySetVersionListPoliciesParamsSortCreatedAt ZonePolicySetVersionListPoliciesParamsSort = "created_at"
)
