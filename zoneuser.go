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

// ZoneUserService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneUserService] method instead.
type ZoneUserService struct {
	Options []option.RequestOption
}

// NewZoneUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneUserService(opts ...option.RequestOption) (r ZoneUserService) {
	r = ZoneUserService{}
	r.Options = opts
	return
}

// Returns details of a specific user by user ID
func (r *ZoneUserService) Get(ctx context.Context, id string, query ZoneUserGetParams, opts ...option.RequestOption) (res *User, err error) {
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
	path := fmt.Sprintf("zones/%s/users/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of users in the specified zone. Can be filtered by email address.
func (r *ZoneUserService) List(ctx context.Context, zoneID string, query ZoneUserListParams, opts ...option.RequestOption) (res *ZoneUserListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/users", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An authenticated user entity
type User struct {
	// Unique identifier of the user
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address of the user
	Email string `json:"email" api:"required" format:"email"`
	// Whether the email address has been verified
	EmailVerified bool `json:"email_verified" api:"required"`
	// Organization that owns this user
	OrganizationID string `json:"organization_id" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zone this user belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// Date when the user was last authenticated
	AuthenticatedAt string `json:"authenticated_at"`
	// Issuer identifier of the identity provider
	Issuer string `json:"issuer"`
	// Reference to the identity provider. This field is undefined when the source
	// identity provider is deleted but the user is not deleted.
	ProviderID string `json:"provider_id"`
	// Subject identifier from the identity provider
	Subject string `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		EmailVerified   respjson.Field
		OrganizationID  respjson.Field
		UpdatedAt       respjson.Field
		ZoneID          respjson.Field
		AuthenticatedAt respjson.Field
		Issuer          respjson.Field
		ProviderID      respjson.Field
		Subject         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserListResponse struct {
	Items []User `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneUserListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneUserListResponsePagination struct {
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
func (r ZoneUserListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserGetParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneUserListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]              `query:"limit,omitzero" json:"-"`
	Expand ZoneUserListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneUserListParams]'s query parameters as `url.Values`.
func (r ZoneUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneUserListsExpandString)
	OfZoneUserListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneUserListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneUserListParamsExpandString string

const (
	ZoneUserListParamsExpandStringTotalCount ZoneUserListParamsExpandString = "total_count"
)
