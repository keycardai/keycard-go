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
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// InvitationService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvitationService] method instead.
type InvitationService struct {
	Options []option.RequestOption
}

// NewInvitationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvitationService(opts ...option.RequestOption) (r InvitationService) {
	r = InvitationService{}
	r.Options = opts
	return
}

// View invitation details by token without consuming the token
func (r *InvitationService) Get(ctx context.Context, token string, query InvitationGetParams, opts ...option.RequestOption) (res *InvitationGetResponse, err error) {
	if !param.IsOmitted(query.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", query.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("invitations/%s", url.PathEscape(token))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Accept and consume an invitation token to join the organization
func (r *InvitationService) Accept(ctx context.Context, token string, body InvitationAcceptParams, opts ...option.RequestOption) (res *InvitationAcceptResponse, err error) {
	if !param.IsOmitted(body.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", body.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("invitations/%s/accept", url.PathEscape(token))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Public invitation details viewable by token
type InvitationGetResponse struct {
	// Name of the user who sent the invitation
	CreatedByName string `json:"created_by_name,required"`
	// Email address for the invitation
	Email string `json:"email,required" format:"email"`
	// When the invitation expires
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Name of the organization being invited to
	OrganizationName string `json:"organization_name,required"`
	// Role that will be assigned when invitation is accepted
	//
	// Any of "org_admin", "org_member", "org_viewer".
	Role OrganizationRole `json:"role,required"`
	// Status of an invitation
	//
	// Any of "pending", "accepted", "expired", "revoked".
	Status InvitationStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedByName    respjson.Field
		Email            respjson.Field
		ExpiresAt        respjson.Field
		OrganizationName respjson.Field
		Role             respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvitationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InvitationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of accepting an invitation
type InvitationAcceptResponse struct {
	// ID of the organization joined
	OrganizationID string `json:"organization_id,required"`
	// Name of the organization joined
	OrganizationName string `json:"organization_name,required"`
	// Whether the invitation was successfully accepted
	Success bool `json:"success,required"`
	// ID of the user who accepted the invitation
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrganizationID   respjson.Field
		OrganizationName respjson.Field
		Success          respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvitationAcceptResponse) RawJSON() string { return r.JSON.raw }
func (r *InvitationAcceptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvitationGetParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type InvitationAcceptParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
