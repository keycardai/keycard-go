// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/apiquery"
	shimjson "github.com/keycardai/keycard-go/internal/encoding/json"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// ZoneGroupRoleService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneGroupRoleService] method instead.
type ZoneGroupRoleService struct {
	Options []option.RequestOption
}

// NewZoneGroupRoleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneGroupRoleService(opts ...option.RequestOption) (r ZoneGroupRoleService) {
	r = ZoneGroupRoleService{}
	r.Options = opts
	return
}

// Returns the roles assigned to the group. Members inherit these roles. Returns
// the shared role-assignment shape with `principal_type` set to `group`. Use
// cursor pagination via `after`/`before`; pass `expand[]=total_count` to include
// the matching row count. Pass `filter[id]` (repeatable, max 100) to restrict
// results to a known set of role assignments, mutually exclusive with
// `after`/`before` (returns 400 if combined). When `filter[id]` is set, `limit` is
// ignored and the response contains every requested assignment that exists on the
// group, in a single page. IDs not on the group are silently omitted.
func (r *ZoneGroupRoleService) List(ctx context.Context, groupID string, params ZoneGroupRoleListParams, opts ...option.RequestOption) (res *ZoneGroupRoleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups/%s/roles", url.PathEscape(params.ZoneID), url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Assigns a role to the group; members inherit it. Provide role_id, or
// role_identifier with owner_type. Returns the shared role-assignment shape with
// `principal_type` set to `group`.
func (r *ZoneGroupRoleService) Add(ctx context.Context, groupID string, params ZoneGroupRoleAddParams, opts ...option.RequestOption) (res *RoleAssignment, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups/%s/roles", url.PathEscape(params.ZoneID), url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Revokes a role from the group. Provide the same (scope_type, scope_id) pair the
// grant was created with, or omit both to revoke the unscoped grant.
func (r *ZoneGroupRoleService) Remove(ctx context.Context, roleID string, params ZoneGroupRoleRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	if params.GroupID == "" {
		err = errors.New("missing required groupId parameter")
		return err
	}
	if roleID == "" {
		err = errors.New("missing required roleId parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/groups/%s/roles/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.GroupID), url.PathEscape(roleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

type ZoneGroupRoleListResponse struct {
	Items []RoleAssignment `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneGroupRoleListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGroupRoleListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneGroupRoleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneGroupRoleListResponsePagination struct {
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
func (r ZoneGroupRoleListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneGroupRoleListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupRoleListParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                   `query:"limit,omitzero" json:"-"`
	Expand ZoneGroupRoleListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Restrict results to the role assignment with this ID. Repeatable, max 100.
	// Mutually exclusive with after/before.
	FilterID ZoneGroupRoleListParamsFilterIDUnion `query:"filter[id],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneGroupRoleListParams]'s query parameters as
// `url.Values`.
func (r ZoneGroupRoleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupRoleListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneGroupRoleListsExpandString)
	OfZoneGroupRoleListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneGroupRoleListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGroupRoleListParamsExpandString string

const (
	ZoneGroupRoleListParamsExpandStringTotalCount ZoneGroupRoleListParamsExpandString = "total_count"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupRoleListParamsFilterIDUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGroupRoleAddParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Schema for assigning a role to a principal. Provide exactly one of role_id or
	// role_identifier. When role_identifier is used, owner_type is required to
	// disambiguate roles that share an identifier across owner types; owner_type must
	// be omitted when role_id is used.
	RoleAssignmentCreate RoleAssignmentCreateParam
	paramObj
}

func (r ZoneGroupRoleAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RoleAssignmentCreate)
}
func (r *ZoneGroupRoleAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupRoleRemoveParams struct {
	ZoneID  string `path:"zoneId" api:"required" json:"-"`
	GroupID string `path:"groupId" api:"required" json:"-"`
	// Scope target of the grant to revoke. Provide together with scope_type.
	ScopeID param.Opt[string] `query:"scope_id,omitzero" json:"-"`
	// Scope kind of the grant to revoke. Provide together with scope_id.
	ScopeType param.Opt[string] `query:"scope_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneGroupRoleRemoveParams]'s query parameters as
// `url.Values`.
func (r ZoneGroupRoleRemoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
