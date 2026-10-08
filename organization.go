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

// OrganizationService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options         []option.RequestOption
	Users           OrganizationUserService
	Invitations     OrganizationInvitationService
	ServiceAccounts OrganizationServiceAccountService
	SSOConnection   OrganizationSSOConnectionService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.Options = opts
	r.Users = NewOrganizationUserService(opts...)
	r.Invitations = NewOrganizationInvitationService(opts...)
	r.ServiceAccounts = NewOrganizationServiceAccountService(opts...)
	r.SSOConnection = NewOrganizationSSOConnectionService(opts...)
	return
}

func (r *OrganizationService) New(ctx context.Context, params OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get organization by ID or label
func (r *OrganizationService) Get(ctx context.Context, organizationID string, params OrganizationGetParams, opts ...option.RequestOption) (res *Organization, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Update organization details
func (r *OrganizationService) Update(ctx context.Context, organizationID string, params OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List organizations for the current user
func (r *OrganizationService) List(ctx context.Context, params OrganizationListParams, opts ...option.RequestOption) (res *OrganizationListResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Deletes the organization and all zones.
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, body OrganizationDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", body.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return err
	}
	path := fmt.Sprintf("organizations/%s", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Organization struct {
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ID string `json:"id" api:"required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A domain name segment for the entity, often derived from the name.
	Label string `json:"label" api:"required"`
	// A name for the entity to be displayed in UI
	Name string `json:"name" api:"required"`
	// Whether SSO is enabled for this organization
	SSOEnabled bool `json:"sso_enabled" api:"required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Identifier of the zone containing this organization's users and their
	// permissions and roles in the organization.
	ZoneID string `json:"zone_id" api:"required"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Label       respjson.Field
		Name        respjson.Field
		SSOEnabled  respjson.Field
		UpdatedAt   respjson.Field
		ZoneID      respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Organization) RawJSON() string { return r.JSON.raw }
func (r *Organization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information using cursor-based pagination
type PageInfoCursor struct {
	// Whether there are more items after the current page
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Whether there are more items before the current page
	HasPrevPage bool `json:"has_prev_page" api:"required"`
	// Cursor pointing to the last item in the current page
	EndCursor string `json:"end_cursor"`
	// Cursor pointing to the first item in the current page
	StartCursor string `json:"start_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		HasPrevPage respjson.Field
		EndCursor   respjson.Field
		StartCursor respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageInfoCursor) RawJSON() string { return r.JSON.raw }
func (r *PageInfoCursor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListResponse struct {
	Items []Organization `json:"items" api:"required"`
	// Pagination information using cursor-based pagination
	//
	// Deprecated: deprecated
	PageInfo PageInfoCursor `json:"page_info" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination OrganizationListResponsePagination `json:"pagination" api:"required"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		Pagination  respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type OrganizationListResponsePagination struct {
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
func (r OrganizationListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationNewParams struct {
	// Organization name
	Name             param.Opt[string] `json:"name,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationGetParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Supports "permissions" to include the
	// permissions field with the caller's permissions for the resource. For the
	// service account and service account credential list operations, "total_count"
	// populates pagination.total_count with the number of items matching the same
	// filters as the list (excluding cursor and limit). Other operations ignore expand
	// values they do not use.
	Expand OrganizationGetParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationGetParams]'s query parameters as `url.Values`.
func (r OrganizationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OrganizationGetParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfOrganizationGetsExpandString)
	OfOrganizationGetsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfOrganizationGetsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type OrganizationGetParamsExpandString string

const (
	OrganizationGetParamsExpandStringPermissions OrganizationGetParamsExpandString = "permissions"
	OrganizationGetParamsExpandStringTotalCount  OrganizationGetParamsExpandString = "total_count"
)

type OrganizationUpdateParams struct {
	// Organization name
	Name             param.Opt[string] `json:"name,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of organizations to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Supports "permissions" to include the
	// permissions field with the caller's permissions for the resource. For the
	// service account and service account credential list operations, "total_count"
	// populates pagination.total_count with the number of items matching the same
	// filters as the list (excluding cursor and limit). Other operations ignore expand
	// values they do not use.
	Expand OrganizationListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OrganizationListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfOrganizationListsExpandString)
	OfOrganizationListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfOrganizationListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type OrganizationListParamsExpandString string

const (
	OrganizationListParamsExpandStringPermissions OrganizationListParamsExpandString = "permissions"
	OrganizationListParamsExpandStringTotalCount  OrganizationListParamsExpandString = "total_count"
)

type OrganizationDeleteParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
