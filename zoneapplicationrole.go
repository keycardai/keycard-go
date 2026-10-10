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

// ZoneApplicationRoleService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneApplicationRoleService] method instead.
type ZoneApplicationRoleService struct {
	Options []option.RequestOption
}

// NewZoneApplicationRoleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneApplicationRoleService(opts ...option.RequestOption) (r ZoneApplicationRoleService) {
	r = ZoneApplicationRoleService{}
	r.Options = opts
	return
}

// Returns the roles assigned to the specified application within the zone. The
// full result set is currently returned in a single page; the
// `after`/`before`/`limit` cursor parameters are reserved and not yet enforced,
// and `pagination` cursors are always null.
func (r *ZoneApplicationRoleService) List(ctx context.Context, applicationID string, params ZoneApplicationRoleListParams, opts ...option.RequestOption) (res *ZoneApplicationRoleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/applications/%s/roles", url.PathEscape(params.ZoneID), url.PathEscape(applicationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Assigns a role to the application. Provide exactly one of role_id or
// role_identifier; when role_identifier is used, owner_type is required to
// disambiguate roles that share an identifier across owner types (and must be
// omitted with role_id). An optional (scope_type, scope_id) pair scopes the grant;
// only platform roles on the org zone may carry a scope, and a `zone` scope must
// reference a different zone in the same organization.
func (r *ZoneApplicationRoleService) Assign(ctx context.Context, applicationID string, params ZoneApplicationRoleAssignParams, opts ...option.RequestOption) (res *RoleAssignment, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/applications/%s/roles", url.PathEscape(params.ZoneID), url.PathEscape(applicationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Revokes a role from the application. Provide the same (scope_type, scope_id)
// pair the grant was created with, or omit both to revoke the unscoped grant.
func (r *ZoneApplicationRoleService) Revoke(ctx context.Context, roleID string, params ZoneApplicationRoleRevokeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	if params.ApplicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return err
	}
	if roleID == "" {
		err = errors.New("missing required roleId parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/applications/%s/roles/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.ApplicationID), url.PathEscape(roleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

type ZoneApplicationRoleListResponse struct {
	Items []RoleAssignment `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneApplicationRoleListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneApplicationRoleListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationRoleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneApplicationRoleListResponsePagination struct {
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
func (r ZoneApplicationRoleListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationRoleListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationRoleListParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                         `query:"limit,omitzero" json:"-"`
	Expand ZoneApplicationRoleListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationRoleListParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationRoleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneApplicationRoleListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneApplicationRoleListsExpandString)
	OfZoneApplicationRoleListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneApplicationRoleListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneApplicationRoleListParamsExpandString string

const (
	ZoneApplicationRoleListParamsExpandStringTotalCount ZoneApplicationRoleListParamsExpandString = "total_count"
)

type ZoneApplicationRoleAssignParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Schema for assigning a role to a principal. Provide exactly one of role_id or
	// role_identifier. When role_identifier is used, owner_type is required to
	// disambiguate roles that share an identifier across owner types; owner_type must
	// be omitted when role_id is used.
	RoleAssignmentCreate RoleAssignmentCreateParam
	paramObj
}

func (r ZoneApplicationRoleAssignParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RoleAssignmentCreate)
}
func (r *ZoneApplicationRoleAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationRoleRevokeParams struct {
	ZoneID        string `path:"zoneId" api:"required" json:"-"`
	ApplicationID string `path:"applicationId" api:"required" json:"-"`
	// Scope target of the grant to revoke. Provide together with scope_type.
	ScopeID param.Opt[string] `query:"scope_id,omitzero" json:"-"`
	// Scope kind of the grant to revoke. Provide together with scope_id.
	ScopeType param.Opt[string] `query:"scope_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationRoleRevokeParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationRoleRevokeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
