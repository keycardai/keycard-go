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

// Immutable policy version snapshots
//
// ZonePolicyVersionService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicyVersionService] method instead.
type ZonePolicyVersionService struct {
	Options []option.RequestOption
}

// NewZonePolicyVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicyVersionService(opts ...option.RequestOption) (r ZonePolicyVersionService) {
	r = ZonePolicyVersionService{}
	r.Options = opts
	return
}

// Create a new immutable policy version
func (r *ZonePolicyVersionService) New(ctx context.Context, policyID string, params ZonePolicyVersionNewParams, opts ...option.RequestOption) (res *PolicyVersion, err error) {
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
	path := fmt.Sprintf("zones/%s/policies/%s/versions", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a specific policy version
func (r *ZonePolicyVersionService) Get(ctx context.Context, versionID string, params ZonePolicyVersionGetParams, opts ...option.RequestOption) (res *PolicyVersion, err error) {
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
	if params.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policies/%s/versions/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.PolicyID), url.PathEscape(versionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List versions of a policy
func (r *ZonePolicyVersionService) List(ctx context.Context, policyID string, params ZonePolicyVersionListParams, opts ...option.RequestOption) (res *ZonePolicyVersionListResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/policies/%s/versions", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Archive a policy version
func (r *ZonePolicyVersionService) Archive(ctx context.Context, versionID string, params ZonePolicyVersionArchiveParams, opts ...option.RequestOption) (res *PolicyVersion, err error) {
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
	if params.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policies/%s/versions/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.PolicyID), url.PathEscape(versionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type PolicyVersion struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	CreatedBy string    `json:"created_by" api:"required"`
	// Who manages this policy version:
	//
	// - `"platform"` — managed by the Keycard platform (system policy versions).
	// - `"customer"` — managed by the tenant (custom policy versions).
	//
	// Any of "platform", "customer".
	OwnerType PolicyVersionOwnerType `json:"owner_type" api:"required"`
	PolicyID  string                 `json:"policy_id" api:"required"`
	// Schema version this policy was validated against when created.
	SchemaVersion string `json:"schema_version" api:"required"`
	// Hex-encoded content hash
	Sha        string    `json:"sha" api:"required"`
	Version    int64     `json:"version" api:"required"`
	ZoneID     string    `json:"zone_id" api:"required"`
	ArchivedAt time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	ArchivedBy string    `json:"archived_by" api:"nullable"`
	// Cedar policy in JSON representation. Populated when format=json (default).
	CedarJson any `json:"cedar_json" api:"nullable"`
	// Cedar policy in human-readable syntax. Populated when format=cedar.
	CedarRaw string `json:"cedar_raw" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CreatedBy     respjson.Field
		OwnerType     respjson.Field
		PolicyID      respjson.Field
		SchemaVersion respjson.Field
		Sha           respjson.Field
		Version       respjson.Field
		ZoneID        respjson.Field
		ArchivedAt    respjson.Field
		ArchivedBy    respjson.Field
		CedarJson     respjson.Field
		CedarRaw      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyVersion) RawJSON() string { return r.JSON.raw }
func (r *PolicyVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who manages this policy version:
//
// - `"platform"` — managed by the Keycard platform (system policy versions).
// - `"customer"` — managed by the tenant (custom policy versions).
type PolicyVersionOwnerType string

const (
	PolicyVersionOwnerTypePlatform PolicyVersionOwnerType = "platform"
	PolicyVersionOwnerTypeCustomer PolicyVersionOwnerType = "customer"
)

type ZonePolicyVersionListResponse struct {
	Items []PolicyVersion `json:"items" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination ZonePolicyVersionListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type ZonePolicyVersionListResponsePagination struct {
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
func (r ZonePolicyVersionListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyVersionListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyVersionNewParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// Schema version to validate this policy against. Must not be archived.
	SchemaVersion string `json:"schema_version" api:"required"`
	// Cedar policy in human-readable Cedar syntax. Mutually exclusive with cedar_json.
	CedarRaw         param.Opt[string] `json:"cedar_raw,omitzero"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Cedar policy in JSON representation. Mutually exclusive with cedar_raw.
	CedarJson any `json:"cedar_json,omitzero"`
	paramObj
}

func (r ZonePolicyVersionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyVersionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyVersionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyVersionGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	PolicyID         string            `path:"policy_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Policy representation format. `json` returns cedar_json, `cedar` returns
	// cedar_raw.
	//
	// Any of "cedar", "json".
	Format ZonePolicyVersionGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicyVersionGetParams]'s query parameters as
// `url.Values`.
func (r ZonePolicyVersionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Policy representation format. `json` returns cedar_json, `cedar` returns
// cedar_raw.
type ZonePolicyVersionGetParamsFormat string

const (
	ZonePolicyVersionGetParamsFormatCedar ZonePolicyVersionGetParamsFormat = "cedar"
	ZonePolicyVersionGetParamsFormatJson  ZonePolicyVersionGetParamsFormat = "json"
)

type ZonePolicyVersionListParams struct {
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
	// **Deprecated.** Use `expand[]` instead.
	//
	// Opt-in to additional response fields. Still honored for backward compatibility;
	// supplying both `expand` and `expand[]` with disagreeing values returns
	// `400 Bad Request`.
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Policy representation format. `json` returns cedar_json, `cedar` returns
	// cedar_raw.
	//
	// Any of "cedar", "json".
	Format ZonePolicyVersionListParamsFormat `query:"format,omitzero" json:"-"`
	// Sort direction. Default is desc (newest first).
	//
	// Any of "asc", "desc".
	Order ZonePolicyVersionListParamsOrder `query:"order,omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicyVersionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicyVersionListParams]'s query parameters as
// `url.Values`.
func (r ZonePolicyVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Policy representation format. `json` returns cedar_json, `cedar` returns
// cedar_raw.
type ZonePolicyVersionListParamsFormat string

const (
	ZonePolicyVersionListParamsFormatCedar ZonePolicyVersionListParamsFormat = "cedar"
	ZonePolicyVersionListParamsFormatJson  ZonePolicyVersionListParamsFormat = "json"
)

// Sort direction. Default is desc (newest first).
type ZonePolicyVersionListParamsOrder string

const (
	ZonePolicyVersionListParamsOrderAsc  ZonePolicyVersionListParamsOrder = "asc"
	ZonePolicyVersionListParamsOrderDesc ZonePolicyVersionListParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicyVersionListParamsSort string

const (
	ZonePolicyVersionListParamsSortCreatedAt ZonePolicyVersionListParamsSort = "created_at"
)

type ZonePolicyVersionArchiveParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	PolicyID         string            `path:"policy_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
