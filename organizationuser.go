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
func (r *OrganizationUserService) Get(ctx context.Context, userID string, params OrganizationUserGetParams, opts ...option.RequestOption) (res *OrganizationUser, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", url.PathEscape(params.OrganizationID), url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update user status in an organization
func (r *OrganizationUserService) Update(ctx context.Context, userID string, params OrganizationUserUpdateParams, opts ...option.RequestOption) (res *OrganizationUser, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", url.PathEscape(params.OrganizationID), url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List users in an organization
func (r *OrganizationUserService) List(ctx context.Context, organizationID string, params OrganizationUserListParams, opts ...option.RequestOption) (res *OrganizationUserListResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Remove a user from an organization
func (r *OrganizationUserService) Delete(ctx context.Context, userID string, params OrganizationUserDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", url.PathEscape(params.OrganizationID), url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// User's role in the organization
type OrganizationRole string

const (
	OrganizationRoleOrgAdmin  OrganizationRole = "org_admin"
	OrganizationRoleOrgMember OrganizationRole = "org_member"
	OrganizationRoleOrgViewer OrganizationRole = "org_viewer"
)

// Status of organization membership
type OrganizationStatus string

const (
	OrganizationStatusActive   OrganizationStatus = "active"
	OrganizationStatusDisabled OrganizationStatus = "disabled"
)

type OrganizationUser struct {
	// The keycard account ID
	ID string `json:"id" api:"required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User's role in the organization
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationRole `json:"role" api:"required"`
	// Identity provider issuer
	Source string `json:"source" api:"required" format:"uri"`
	// Status of organization membership
	//
	// Any of "active", "disabled".
	Status OrganizationStatus `json:"status" api:"required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
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
func (r OrganizationUser) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserListResponse struct {
	Items []OrganizationUser `json:"items" api:"required"`
	// Pagination information using cursor-based pagination
	PageInfo PageInfoCursor `json:"page_info" api:"required"`
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

type OrganizationUserGetParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id" api:"required" json:"-"`
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
	OrganizationID   string            `path:"organization_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// New role for the user in the organization
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationRole `json:"role,omitzero"`
	// New status for the user in the organization
	//
	// Any of "active", "disabled".
	Status OrganizationStatus `json:"status,omitzero"`
	paramObj
}

func (r OrganizationUserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	Role OrganizationRole `query:"role,omitzero" json:"-"`
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

type OrganizationUserDeleteParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
