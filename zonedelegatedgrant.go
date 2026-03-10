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

// ZoneDelegatedGrantService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDelegatedGrantService] method instead.
type ZoneDelegatedGrantService struct {
	Options []option.RequestOption
}

// NewZoneDelegatedGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDelegatedGrantService(opts ...option.RequestOption) (r ZoneDelegatedGrantService) {
	r = ZoneDelegatedGrantService{}
	r.Options = opts
	return
}

// Returns details of a specific delegated grant by grant ID
func (r *ZoneDelegatedGrantService) Get(ctx context.Context, id string, query ZoneDelegatedGrantGetParams, opts ...option.RequestOption) (res *Grant, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/delegated-grants/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revokes an active delegated grant
func (r *ZoneDelegatedGrantService) Update(ctx context.Context, id string, params ZoneDelegatedGrantUpdateParams, opts ...option.RequestOption) (res *Grant, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/delegated-grants/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of delegated grants in the specified zone. Can be filtered by
// user, resource, or status.
func (r *ZoneDelegatedGrantService) List(ctx context.Context, zoneID string, query ZoneDelegatedGrantListParams, opts ...option.RequestOption) (res *ZoneDelegatedGrantListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/delegated-grants", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently revokes a delegated grant, removing the user's access to the
// protected resource
func (r *ZoneDelegatedGrantService) Delete(ctx context.Context, id string, body ZoneDelegatedGrantDeleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/delegated-grants/%s", url.PathEscape(body.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// User authorization for a resource to be accessed on their behalf. The grant
// links the user, resource, and the provider that issued the grant.
type Grant struct {
	// Unique identifier of the delegated grant
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Date when grant expires
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Organization that owns this grant
	OrganizationID string `json:"organization_id" api:"required"`
	// ID of the provider that issued this grant
	ProviderID string `json:"provider_id" api:"required"`
	// Indicates whether a refresh token is stored for this grant. Grants with refresh
	// tokens can be refreshed even after access token expiration.
	RefreshTokenSet bool `json:"refresh_token_set" api:"required"`
	// ID of resource receiving grant
	ResourceID string `json:"resource_id" api:"required"`
	// Granted OAuth scopes
	Scopes []string `json:"scopes" api:"required"`
	// Any of "active", "expired", "revoked".
	Status GrantStatus `json:"status" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Reference to the user granting permission
	UserID string `json:"user_id" api:"required"`
	// Zone this grant belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// Whether the grant is currently active (deprecated - use status instead)
	//
	// Deprecated: deprecated
	Active bool `json:"active"`
	// A Provider is a system that supplies access to Resources and allows actors
	// (Users or Applications) to authenticate.
	//
	// Deprecated: deprecated
	Provider Provider `json:"provider"`
	// Timestamp when this grant's tokens were last refreshed. Omitted if grant was
	// never refreshed.
	RefreshedAt time.Time `json:"refreshed_at" format:"date-time"`
	// A Resource is a system that exposes protected information or functionality. It
	// requires authentication of the requesting actor, which may be a user or
	// application, before allowing access.
	//
	// Deprecated: deprecated
	Resource Resource `json:"resource"`
	// An authenticated user entity
	//
	// Deprecated: deprecated
	User User `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		ExpiresAt       respjson.Field
		OrganizationID  respjson.Field
		ProviderID      respjson.Field
		RefreshTokenSet respjson.Field
		ResourceID      respjson.Field
		Scopes          respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		UserID          respjson.Field
		ZoneID          respjson.Field
		Active          respjson.Field
		Provider        respjson.Field
		RefreshedAt     respjson.Field
		Resource        respjson.Field
		User            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Grant) RawJSON() string { return r.JSON.raw }
func (r *Grant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GrantStatus string

const (
	GrantStatusActive  GrantStatus = "active"
	GrantStatusExpired GrantStatus = "expired"
	GrantStatusRevoked GrantStatus = "revoked"
)

type ZoneDelegatedGrantListResponse struct {
	Items []Grant `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneDelegatedGrantListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneDelegatedGrantListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneDelegatedGrantListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneDelegatedGrantListResponsePagination struct {
	// An opaque cursor used for paginating through a list of results
	AfterCursor string `json:"after_cursor" api:"required"`
	// An opaque cursor used for paginating through a list of results
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items matching the query. Only included when
	// expand[]=total_count is requested.
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
func (r ZoneDelegatedGrantListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneDelegatedGrantListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDelegatedGrantGetParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneDelegatedGrantUpdateParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Any of "revoked".
	Status ZoneDelegatedGrantUpdateParamsStatus `json:"status,omitzero" api:"required"`
	paramObj
}

func (r ZoneDelegatedGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneDelegatedGrantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneDelegatedGrantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDelegatedGrantUpdateParamsStatus string

const (
	ZoneDelegatedGrantUpdateParamsStatusRevoked ZoneDelegatedGrantUpdateParamsStatus = "revoked"
)

type ZoneDelegatedGrantListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by resource ID
	ResourceID param.Opt[string] `query:"resource_id,omitzero" json:"-"`
	// Filter by user ID
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Any of "true".
	Active ZoneDelegatedGrantListParamsActive      `query:"active,omitzero" json:"-"`
	Expand ZoneDelegatedGrantListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Any of "active", "expired", "revoked".
	Status ZoneDelegatedGrantListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneDelegatedGrantListParams]'s query parameters as
// `url.Values`.
func (r ZoneDelegatedGrantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneDelegatedGrantListParamsActive string

const (
	ZoneDelegatedGrantListParamsActiveTrue ZoneDelegatedGrantListParamsActive = "true"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneDelegatedGrantListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneDelegatedGrantListsExpandString)
	OfZoneDelegatedGrantListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneDelegatedGrantListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneDelegatedGrantListParamsExpandString string

const (
	ZoneDelegatedGrantListParamsExpandStringTotalCount ZoneDelegatedGrantListParamsExpandString = "total_count"
)

type ZoneDelegatedGrantListParamsStatus string

const (
	ZoneDelegatedGrantListParamsStatusActive  ZoneDelegatedGrantListParamsStatus = "active"
	ZoneDelegatedGrantListParamsStatusExpired ZoneDelegatedGrantListParamsStatus = "expired"
	ZoneDelegatedGrantListParamsStatusRevoked ZoneDelegatedGrantListParamsStatus = "revoked"
)

type ZoneDelegatedGrantDeleteParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}
