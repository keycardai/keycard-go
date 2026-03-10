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
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get organization by ID or label
func (r *OrganizationService) Get(ctx context.Context, organizationID string, params OrganizationGetParams, opts ...option.RequestOption) (res *Organization, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
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
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
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
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Exchange user token for organization-scoped M2M token
func (r *OrganizationService) ExchangeToken(ctx context.Context, organizationID string, body OrganizationExchangeTokenParams, opts ...option.RequestOption) (res *TokenResponse, err error) {
	if !param.IsOmitted(body.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", body.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/token", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// List unified view of users and invitations in an organization
func (r *OrganizationService) ListIdentities(ctx context.Context, organizationID string, params OrganizationListIdentitiesParams, opts ...option.RequestOption) (res *OrganizationListIdentitiesResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/identities", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns the list of available roles in the system for the organization. This
// includes both organization-level roles (e.g., org_admin, org_member) and
// zone-level roles (e.g., zone_manager, zone_viewer).
//
// Each role includes:
//
// - `name`: Internal identifier (e.g., org_admin, zone_manager)
// - `label`: Human-readable display name (e.g., Organization Administrator)
// - `scope`: Whether the role applies at organization or zone level
func (r *OrganizationService) ListRoles(ctx context.Context, organizationID string, params OrganizationListRolesParams, opts ...option.RequestOption) (res *OrganizationListRolesResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/roles", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
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

// The scope at which a role can be assigned.
//
// - organization: Roles that apply at the organization level (e.g., org_admin)
// - zone: Roles that apply at the zone level (e.g., zone_manager)
type RoleScope string

const (
	RoleScopeOrganization RoleScope = "organization"
	RoleScopeZone         RoleScope = "zone"
)

// OAuth2-style token response for M2M tokens
type TokenResponse struct {
	// The M2M access token
	AccessToken string `json:"access_token" api:"required"`
	// Token type (always "Bearer")
	TokenType string `json:"token_type" api:"required"`
	// Token expiration time in seconds
	ExpiresIn int64 `json:"expires_in"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		TokenType   respjson.Field
		ExpiresIn   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenResponse) RawJSON() string { return r.JSON.raw }
func (r *TokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListResponse struct {
	Items []Organization `json:"items" api:"required"`
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
func (r OrganizationListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of identities (users and invitations) in an organization
type OrganizationListIdentitiesResponse struct {
	Items []OrganizationListIdentitiesResponseItem `json:"items" api:"required"`
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
func (r OrganizationListIdentitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListIdentitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified view of users and invitations in an organization
type OrganizationListIdentitiesResponseItem struct {
	// The identity ID (user or invitation)
	ID string `json:"id" api:"required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address of the identity
	Email string `json:"email" api:"required" format:"email"`
	// Role in the organization
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationRole `json:"role" api:"required"`
	// Identity provider issuer
	Source string `json:"source" api:"required" format:"uri"`
	// Status of the identity (OrganizationStatus for users, InvitationStatus for
	// invitations)
	//
	// Any of "active", "disabled", "pending", "accepted", "expired", "revoked".
	Status OrganizationListIdentitiesResponseItemStatus `json:"status" api:"required"`
	// Type of identity (user or invitation)
	//
	// Any of "user", "invitation".
	Type string `json:"type" api:"required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		Role        respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationListIdentitiesResponseItem) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListIdentitiesResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the identity (OrganizationStatus for users, InvitationStatus for
// invitations)
type OrganizationListIdentitiesResponseItemStatus string

const (
	OrganizationListIdentitiesResponseItemStatusActive   OrganizationListIdentitiesResponseItemStatus = "active"
	OrganizationListIdentitiesResponseItemStatusDisabled OrganizationListIdentitiesResponseItemStatus = "disabled"
	OrganizationListIdentitiesResponseItemStatusPending  OrganizationListIdentitiesResponseItemStatus = "pending"
	OrganizationListIdentitiesResponseItemStatusAccepted OrganizationListIdentitiesResponseItemStatus = "accepted"
	OrganizationListIdentitiesResponseItemStatusExpired  OrganizationListIdentitiesResponseItemStatus = "expired"
	OrganizationListIdentitiesResponseItemStatusRevoked  OrganizationListIdentitiesResponseItemStatus = "revoked"
)

// List of available roles
type OrganizationListRolesResponse struct {
	// List of roles
	Items []OrganizationListRolesResponseItem `json:"items" api:"required"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationListRolesResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListRolesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A role definition that can be assigned to users
type OrganizationListRolesResponseItem struct {
	// Detailed description of the role and its permissions
	Description string `json:"description" api:"required"`
	// Human-readable display name for the role
	Label string `json:"label" api:"required"`
	// Internal identifier for the role (e.g., org_admin, zone_manager)
	Name string `json:"name" api:"required"`
	// The scope at which this role can be assigned (organization or zone)
	//
	// Any of "organization", "zone".
	Scope RoleScope `json:"scope" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Label       respjson.Field
		Name        respjson.Field
		Scope       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationListRolesResponseItem) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListRolesResponseItem) UnmarshalJSON(data []byte) error {
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
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationGetParams]'s query parameters as `url.Values`.
func (r OrganizationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationExchangeTokenParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type OrganizationListIdentitiesParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of identities to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Filter identities by role
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationRole `query:"role,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListIdentitiesParams]'s query parameters as
// `url.Values`.
func (r OrganizationListIdentitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationListRolesParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Filter roles by scope (organization or zone level)
	//
	// Any of "organization", "zone".
	Scope RoleScope `query:"scope,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListRolesParams]'s query parameters as
// `url.Values`.
func (r OrganizationListRolesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
