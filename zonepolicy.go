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

// Policy CRUD operations
//
// ZonePolicyService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicyService] method instead.
type ZonePolicyService struct {
	Options []option.RequestOption
	// Immutable policy version snapshots
	Versions ZonePolicyVersionService
}

// NewZonePolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePolicyService(opts ...option.RequestOption) (r ZonePolicyService) {
	r = ZonePolicyService{}
	r.Options = opts
	r.Versions = NewZonePolicyVersionService(opts...)
	return
}

// Create a new policy
func (r *ZonePolicyService) New(ctx context.Context, zoneID string, params ZonePolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
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
	path := fmt.Sprintf("zones/%s/policies", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a policy by ID
func (r *ZonePolicyService) Get(ctx context.Context, policyID string, params ZonePolicyGetParams, opts ...option.RequestOption) (res *Policy, err error) {
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
	path := fmt.Sprintf("zones/%s/policies/%s", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a policy
func (r *ZonePolicyService) Update(ctx context.Context, policyID string, params ZonePolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch.Value)))
	}
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
	path := fmt.Sprintf("zones/%s/policies/%s", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List policies in a zone
func (r *ZonePolicyService) List(ctx context.Context, zoneID string, params ZonePolicyListParams, opts ...option.RequestOption) (res *ZonePolicyListResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/policies", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Archive a policy
func (r *ZonePolicyService) Archive(ctx context.Context, policyID string, params ZonePolicyArchiveParams, opts ...option.RequestOption) (res *Policy, err error) {
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
	path := fmt.Sprintf("zones/%s/policies/%s", url.PathEscape(params.ZoneID), url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Policy struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	CreatedBy string    `json:"created_by" api:"required"`
	Name      string    `json:"name" api:"required"`
	// Who manages this policy:
	//
	// - `"platform"` — managed by the Keycard platform (system policies).
	// - `"customer"` — managed by the tenant (custom policies).
	//
	// Any of "platform", "customer".
	OwnerType   PolicyOwnerType `json:"owner_type" api:"required"`
	UpdatedAt   time.Time       `json:"updated_at" api:"required" format:"date-time"`
	ZoneID      string          `json:"zone_id" api:"required"`
	ArchivedAt  time.Time       `json:"archived_at" api:"nullable" format:"date-time"`
	Description string          `json:"description" api:"nullable"`
	// Human-readable version number of the latest version (e.g., 1, 2, 3)
	LatestVersion   int64  `json:"latest_version" api:"nullable"`
	LatestVersionID string `json:"latest_version_id" api:"nullable"`
	UpdatedBy       string `json:"updated_by" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		Name            respjson.Field
		OwnerType       respjson.Field
		UpdatedAt       respjson.Field
		ZoneID          respjson.Field
		ArchivedAt      respjson.Field
		Description     respjson.Field
		LatestVersion   respjson.Field
		LatestVersionID respjson.Field
		UpdatedBy       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Policy) RawJSON() string { return r.JSON.raw }
func (r *Policy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who manages this policy:
//
// - `"platform"` — managed by the Keycard platform (system policies).
// - `"customer"` — managed by the tenant (custom policies).
type PolicyOwnerType string

const (
	PolicyOwnerTypePlatform PolicyOwnerType = "platform"
	PolicyOwnerTypeCustomer PolicyOwnerType = "customer"
)

type ZonePolicyListResponse struct {
	Items      []Policy                         `json:"items" api:"required"`
	Pagination ZonePolicyListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyListResponsePagination struct {
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
func (r ZonePolicyListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyNewParams struct {
	Name             string            `json:"name" api:"required"`
	Description      param.Opt[string] `json:"description,omitzero"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ZonePolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicyUpdateParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	Description      param.Opt[string] `json:"description,omitzero"`
	Name             param.Opt[string] `json:"name,omitzero"`
	IfMatch          param.Opt[string] `header:"If-Match,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ZonePolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyListParams struct {
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
	Order ZonePolicyListParamsOrder `query:"order,omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicyListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicyListParams]'s query parameters as `url.Values`.
func (r ZonePolicyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction. Default is desc (newest first).
type ZonePolicyListParamsOrder string

const (
	ZonePolicyListParamsOrderAsc  ZonePolicyListParamsOrder = "asc"
	ZonePolicyListParamsOrderDesc ZonePolicyListParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicyListParamsSort string

const (
	ZonePolicyListParamsSortCreatedAt ZonePolicyListParamsSort = "created_at"
)

type ZonePolicyArchiveParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
