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
	shimjson "github.com/keycardai/keycard-go/internal/encoding/json"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// ZoneUserRoleService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneUserRoleService] method instead.
type ZoneUserRoleService struct {
	Options []option.RequestOption
}

// NewZoneUserRoleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneUserRoleService(opts ...option.RequestOption) (r ZoneUserRoleService) {
	r = ZoneUserRoleService{}
	r.Options = opts
	return
}

// Returns the roles assigned to the specified user within the zone. The full
// result set is currently returned in a single page; the `after`/`before`/`limit`
// cursor parameters are reserved and not yet enforced, and `pagination` cursors
// are always null.
func (r *ZoneUserRoleService) List(ctx context.Context, userID string, params ZoneUserRoleListParams, opts ...option.RequestOption) (res *ZoneUserRoleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/users/%s/roles", url.PathEscape(params.ZoneID), url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Assigns a role to the user. Provide exactly one of role_id or role_identifier;
// when role_identifier is used, owner_type is required to disambiguate roles that
// share an identifier across owner types (and must be omitted with role_id). An
// optional (scope_type, scope_id) pair scopes the grant; only platform roles on
// the org zone may carry a scope, and a `zone` scope must reference a different
// zone in the same organization.
func (r *ZoneUserRoleService) Assign(ctx context.Context, userID string, params ZoneUserRoleAssignParams, opts ...option.RequestOption) (res *RoleAssignment, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/users/%s/roles", url.PathEscape(params.ZoneID), url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Revokes a role from the user. Provide the same (scope_type, scope_id) pair the
// grant was created with, or omit both to revoke the unscoped grant.
func (r *ZoneUserRoleService) Revoke(ctx context.Context, roleID string, params ZoneUserRoleRevokeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	if params.UserID == "" {
		err = errors.New("missing required userId parameter")
		return err
	}
	if roleID == "" {
		err = errors.New("missing required roleId parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/users/%s/roles/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.UserID), url.PathEscape(roleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Represents a role assigned to a principal within a zone
type RoleAssignment struct {
	// Unique identifier of the role assignment
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of the principal the role is assigned to (a user, application, or group ID).
	PrincipalID string `json:"principal_id" api:"required"`
	// The kind of principal the role is assigned to: `user`, `application`, or
	// `group`. A role assigned to a `group` is inherited by that group's members.
	PrincipalType string `json:"principal_type" api:"required"`
	// ID of the assigned role
	RoleID string `json:"role_id" api:"required"`
	// Role identifier: a lowercase slug (letters and digits separated by single
	// hyphens or underscores), unique per owner type within a zone. Role identifiers
	// surface in policy evaluation, so the slug restriction keeps them unambiguous in
	// policy text.
	RoleIdentifier string `json:"role_identifier" api:"required"`
	// Owner type of the assigned role. Disambiguates roles that share an identifier
	// across owner types.
	//
	// Any of "platform", "customer".
	RoleOwnerType RoleAssignmentRoleOwnerType `json:"role_owner_type" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zone this assignment belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// The ID of the scoped resource. Null when the assignment is unscoped.
	ScopeID string `json:"scope_id" api:"nullable"`
	// The kind of resource this grant is scoped to (e.g. `zone`). Null when the
	// assignment is unscoped (applies to the owning zone itself).
	ScopeType string `json:"scope_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		PrincipalID    respjson.Field
		PrincipalType  respjson.Field
		RoleID         respjson.Field
		RoleIdentifier respjson.Field
		RoleOwnerType  respjson.Field
		UpdatedAt      respjson.Field
		ZoneID         respjson.Field
		ScopeID        respjson.Field
		ScopeType      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoleAssignment) RawJSON() string { return r.JSON.raw }
func (r *RoleAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Owner type of the assigned role. Disambiguates roles that share an identifier
// across owner types.
type RoleAssignmentRoleOwnerType string

const (
	RoleAssignmentRoleOwnerTypePlatform RoleAssignmentRoleOwnerType = "platform"
	RoleAssignmentRoleOwnerTypeCustomer RoleAssignmentRoleOwnerType = "customer"
)

// Schema for assigning a role to a principal. Provide exactly one of role_id or
// role_identifier. When role_identifier is used, owner_type is required to
// disambiguate roles that share an identifier across owner types; owner_type must
// be omitted when role_id is used.
type RoleAssignmentCreateParam struct {
	// ID of the role to assign. Provide exactly one of role_id or role_identifier;
	// owner_type must be omitted when role_id is used.
	RoleID param.Opt[string] `json:"role_id,omitzero"`
	// Role identifier: a lowercase slug (letters and digits separated by single
	// hyphens or underscores), unique per owner type within a zone. Role identifiers
	// surface in policy evaluation, so the slug restriction keeps them unambiguous in
	// policy text.
	RoleIdentifier param.Opt[string] `json:"role_identifier,omitzero"`
	// The ID of the resource to scope the grant to. Provide together with scope_type,
	// or omit both for an unscoped assignment. When scope_type is `zone`, this must
	// reference a different zone in the same organization.
	ScopeID param.Opt[string] `json:"scope_id,omitzero"`
	// The kind of resource to scope the grant to (e.g. `zone`). Provide together with
	// scope_id, or omit both for an unscoped assignment (applies to the owning zone
	// itself). Only platform roles on the org zone may carry a scope.
	ScopeType param.Opt[string] `json:"scope_type,omitzero"`
	// Owner type of the role to assign. Required with role_identifier (an identifier
	// is unique only per owner type); must be omitted with role_id.
	//
	// Any of "platform", "customer".
	OwnerType RoleAssignmentCreateOwnerType `json:"owner_type,omitzero"`
	paramObj
}

func (r RoleAssignmentCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow RoleAssignmentCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoleAssignmentCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Owner type of the role to assign. Required with role_identifier (an identifier
// is unique only per owner type); must be omitted with role_id.
type RoleAssignmentCreateOwnerType string

const (
	RoleAssignmentCreateOwnerTypePlatform RoleAssignmentCreateOwnerType = "platform"
	RoleAssignmentCreateOwnerTypeCustomer RoleAssignmentCreateOwnerType = "customer"
)

type ZoneUserRoleListResponse struct {
	Items []RoleAssignment `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneUserRoleListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneUserRoleListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserRoleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneUserRoleListResponsePagination struct {
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
func (r ZoneUserRoleListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserRoleListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserRoleListParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                  `query:"limit,omitzero" json:"-"`
	Expand ZoneUserRoleListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneUserRoleListParams]'s query parameters as `url.Values`.
func (r ZoneUserRoleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserRoleListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneUserRoleListsExpandString)
	OfZoneUserRoleListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneUserRoleListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneUserRoleListParamsExpandString string

const (
	ZoneUserRoleListParamsExpandStringTotalCount ZoneUserRoleListParamsExpandString = "total_count"
)

type ZoneUserRoleAssignParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Schema for assigning a role to a principal. Provide exactly one of role_id or
	// role_identifier. When role_identifier is used, owner_type is required to
	// disambiguate roles that share an identifier across owner types; owner_type must
	// be omitted when role_id is used.
	RoleAssignmentCreate RoleAssignmentCreateParam
	paramObj
}

func (r ZoneUserRoleAssignParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RoleAssignmentCreate)
}
func (r *ZoneUserRoleAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserRoleRevokeParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	UserID string `path:"userId" api:"required" json:"-"`
	// Scope target of the grant to revoke. Provide together with scope_type.
	ScopeID param.Opt[string] `query:"scope_id,omitzero" json:"-"`
	// Scope kind of the grant to revoke. Provide together with scope_id.
	ScopeType param.Opt[string] `query:"scope_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneUserRoleRevokeParams]'s query parameters as
// `url.Values`.
func (r ZoneUserRoleRevokeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
