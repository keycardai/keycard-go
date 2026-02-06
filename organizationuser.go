// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/keycard-api-go/internal/apijson"
	"github.com/stainless-sdks/keycard-api-go/internal/apiquery"
	"github.com/stainless-sdks/keycard-api-go/internal/requestconfig"
	"github.com/stainless-sdks/keycard-api-go/option"
	"github.com/stainless-sdks/keycard-api-go/packages/param"
	"github.com/stainless-sdks/keycard-api-go/packages/respjson"
)

// OrganizationUserService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationUserService] method instead.
type OrganizationUserService struct {
	Options []option.RequestOption
}

// NewOrganizationUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationUserService(opts ...option.RequestOption) (r OrganizationUserService) {
	r = OrganizationUserService{}
	r.Options = opts
	return
}

// Get a specific user in an organization
func (r *OrganizationUserService) Get(ctx context.Context, userID string, params OrganizationUserGetParams, opts ...option.RequestOption) (res *OrganizationUserGetResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", params.OrganizationID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update user status in an organization
func (r *OrganizationUserService) Update(ctx context.Context, userID string, params OrganizationUserUpdateParams, opts ...option.RequestOption) (res *OrganizationUserUpdateResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", params.OrganizationID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List users in an organization
func (r *OrganizationUserService) List(ctx context.Context, organizationID string, params OrganizationUserListParams, opts ...option.RequestOption) (res *OrganizationUserListResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Remove a user from an organization
func (r *OrganizationUserService) Delete(ctx context.Context, userID string, params OrganizationUserDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", params.OrganizationID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type OrganizationUserGetResponse struct {
	// The keycard account ID
	ID string `json:"id,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User's role in the organization
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationUserGetResponseRole `json:"role,required"`
	// Identity provider issuer
	Source string `json:"source,required" format:"uri"`
	// Status of organization membership
	//
	// Any of "active", "disabled".
	Status OrganizationUserGetResponseStatus `json:"status,required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User email address
	Email string `json:"email" format:"email"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Role        respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Email       respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's role in the organization
type OrganizationUserGetResponseRole string

const (
	OrganizationUserGetResponseRoleOrgAdmin  OrganizationUserGetResponseRole = "org_admin"
	OrganizationUserGetResponseRoleOrgMember OrganizationUserGetResponseRole = "org_member"
	OrganizationUserGetResponseRoleOrgViewer OrganizationUserGetResponseRole = "org_viewer"
)

// Status of organization membership
type OrganizationUserGetResponseStatus string

const (
	OrganizationUserGetResponseStatusActive   OrganizationUserGetResponseStatus = "active"
	OrganizationUserGetResponseStatusDisabled OrganizationUserGetResponseStatus = "disabled"
)

type OrganizationUserUpdateResponse struct {
	// The keycard account ID
	ID string `json:"id,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User's role in the organization
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationUserUpdateResponseRole `json:"role,required"`
	// Identity provider issuer
	Source string `json:"source,required" format:"uri"`
	// Status of organization membership
	//
	// Any of "active", "disabled".
	Status OrganizationUserUpdateResponseStatus `json:"status,required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User email address
	Email string `json:"email" format:"email"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Role        respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Email       respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's role in the organization
type OrganizationUserUpdateResponseRole string

const (
	OrganizationUserUpdateResponseRoleOrgAdmin  OrganizationUserUpdateResponseRole = "org_admin"
	OrganizationUserUpdateResponseRoleOrgMember OrganizationUserUpdateResponseRole = "org_member"
	OrganizationUserUpdateResponseRoleOrgViewer OrganizationUserUpdateResponseRole = "org_viewer"
)

// Status of organization membership
type OrganizationUserUpdateResponseStatus string

const (
	OrganizationUserUpdateResponseStatusActive   OrganizationUserUpdateResponseStatus = "active"
	OrganizationUserUpdateResponseStatusDisabled OrganizationUserUpdateResponseStatus = "disabled"
)

type OrganizationUserListResponse struct {
	Items []OrganizationUserListResponseItem `json:"items,required"`
	// Pagination information using cursor-based pagination
	PageInfo OrganizationUserListResponsePageInfo `json:"page_info,required"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserListResponseItem struct {
	// The keycard account ID
	ID string `json:"id,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User's role in the organization
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role string `json:"role,required"`
	// Identity provider issuer
	Source string `json:"source,required" format:"uri"`
	// Status of organization membership
	//
	// Any of "active", "disabled".
	Status string `json:"status,required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User email address
	Email string `json:"email" format:"email"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Role        respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Email       respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information using cursor-based pagination
type OrganizationUserListResponsePageInfo struct {
	// Whether there are more items after the current page
	HasNextPage bool `json:"has_next_page,required"`
	// Whether there are more items before the current page
	HasPrevPage bool `json:"has_prev_page,required"`
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
func (r OrganizationUserListResponsePageInfo) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserListResponsePageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserGetParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationUserGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationUserGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationUserUpdateParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// New role for the user in the organization
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationUserUpdateParamsRole `json:"role,omitzero"`
	// New status for the user in the organization
	//
	// Any of "active", "disabled".
	Status OrganizationUserUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r OrganizationUserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New role for the user in the organization
type OrganizationUserUpdateParamsRole string

const (
	OrganizationUserUpdateParamsRoleOrgAdmin  OrganizationUserUpdateParamsRole = "org_admin"
	OrganizationUserUpdateParamsRoleOrgMember OrganizationUserUpdateParamsRole = "org_member"
	OrganizationUserUpdateParamsRoleOrgViewer OrganizationUserUpdateParamsRole = "org_viewer"
)

// New status for the user in the organization
type OrganizationUserUpdateParamsStatus string

const (
	OrganizationUserUpdateParamsStatusActive   OrganizationUserUpdateParamsStatus = "active"
	OrganizationUserUpdateParamsStatusDisabled OrganizationUserUpdateParamsStatus = "disabled"
)

type OrganizationUserListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of users to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Filter users by role
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationUserListParamsRole `query:"role,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationUserListParams]'s query parameters as
// `url.Values`.
func (r OrganizationUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter users by role
type OrganizationUserListParamsRole string

const (
	OrganizationUserListParamsRoleOrgAdmin  OrganizationUserListParamsRole = "org_admin"
	OrganizationUserListParamsRoleOrgMember OrganizationUserListParamsRole = "org_member"
	OrganizationUserListParamsRoleOrgViewer OrganizationUserListParamsRole = "org_viewer"
)

type OrganizationUserDeleteParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
