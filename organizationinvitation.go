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

// OrganizationInvitationService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationInvitationService] method instead.
type OrganizationInvitationService struct {
	Options []option.RequestOption
}

// NewOrganizationInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationInvitationService(opts ...option.RequestOption) (r OrganizationInvitationService) {
	r = OrganizationInvitationService{}
	r.Options = opts
	return
}

// Create an invitation to join an organization
func (r *OrganizationInvitationService) New(ctx context.Context, organizationID string, params OrganizationInvitationNewParams, opts ...option.RequestOption) (res *Invitation, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/invitations", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List invitations for an organization
func (r *OrganizationInvitationService) List(ctx context.Context, organizationID string, params OrganizationInvitationListParams, opts ...option.RequestOption) (res *OrganizationInvitationListResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/invitations", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete an invitation
func (r *OrganizationInvitationService) Delete(ctx context.Context, invitationID string, params OrganizationInvitationDeleteParams, opts ...option.RequestOption) (err error) {
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
	if invitationID == "" {
		err = errors.New("missing required invitation_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/invitations/%s", url.PathEscape(params.OrganizationID), url.PathEscape(invitationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Invitation struct {
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ID string `json:"id,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// ID of the user who created the invitation
	CreatedBy string `json:"created_by,required"`
	// Email address for the invitation
	Email string `json:"email,required" format:"email"`
	// When the invitation expires
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	OrganizationID string `json:"organization_id,required"`
	// Role that will be assigned when invitation is accepted
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationRole `json:"role,required"`
	// Status of an invitation
	//
	// Any of "pending", "accepted", "expired", "revoked".
	Status InvitationStatus `json:"status,required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Email          respjson.Field
		ExpiresAt      respjson.Field
		OrganizationID respjson.Field
		Role           respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		Permissions    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invitation) RawJSON() string { return r.JSON.raw }
func (r *Invitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of an invitation
type InvitationStatus string

const (
	InvitationStatusPending  InvitationStatus = "pending"
	InvitationStatusAccepted InvitationStatus = "accepted"
	InvitationStatusExpired  InvitationStatus = "expired"
	InvitationStatusRevoked  InvitationStatus = "revoked"
)

type OrganizationInvitationListResponse struct {
	Items []Invitation `json:"items,required"`
	// Pagination information using cursor-based pagination
	PageInfo PageInfoCursor `json:"page_info,required"`
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
func (r OrganizationInvitationListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationInvitationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInvitationNewParams struct {
	// Email address to invite
	Email string `json:"email,required" format:"email"`
	// Role to assign when invitation is accepted
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role             OrganizationRole  `json:"role,omitzero,required"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r OrganizationInvitationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationInvitationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationInvitationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInvitationListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of invitations to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationInvitationListParams]'s query parameters as
// `url.Values`.
func (r OrganizationInvitationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationInvitationDeleteParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
