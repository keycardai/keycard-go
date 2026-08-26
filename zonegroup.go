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

// ZoneGroupService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneGroupService] method instead.
type ZoneGroupService struct {
	Options []option.RequestOption
	Members ZoneGroupMemberService
	Roles   ZoneGroupRoleService
}

// NewZoneGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneGroupService(opts ...option.RequestOption) (r ZoneGroupService) {
	r = ZoneGroupService{}
	r.Options = opts
	r.Members = NewZoneGroupMemberService(opts...)
	r.Roles = NewZoneGroupRoleService(opts...)
	return
}

// Creates a group in the zone (managed in Keycard). Groups synced from an external
// directory are created by that directory, not here.
func (r *ZoneGroupService) New(ctx context.Context, zoneID string, body ZoneGroupNewParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a group by ID. Pass `expand[]=member_count` for its member count and
// `expand[]=roles` for the identifiers of its assigned roles.
func (r *ZoneGroupService) Get(ctx context.Context, groupID string, params ZoneGroupGetParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups/%s", url.PathEscape(params.ZoneID), url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Updates a group's name and/or identifier (partial update). A group's source is
// immutable. The name of a group synced from an external directory cannot be
// changed while external sync is enabled for the zone; its identifier can.
func (r *ZoneGroupService) Update(ctx context.Context, groupID string, params ZoneGroupUpdateParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups/%s", url.PathEscape(params.ZoneID), url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of the groups in the specified zone. Use cursor
// pagination via `after`/`before`. Sort: comma-separated field list; prefix with
// `-` for descending (allowed: created_at, name, identifier). Pass
// `expand[]=member_count` to include each group's member count, `expand[]=roles`
// to include the identifiers of the roles assigned to each group, and
// `expand[]=total_count` to include the matching row count. Filter by exact
// identifier via `filter[identifier]` (repeatable, OR'd across values). Search via
// `query[]` (case-insensitive substring match, OR'd across repeated values); it
// matches the group's name and identifier. Pass `filter[id]` (repeatable, max 100)
// to restrict results to a known set of groups — mutually exclusive with
// `after`/`before` (returns 400 if combined). When `filter[id]` is set, `limit` is
// ignored and the response contains every requested group that exists in the zone,
// in a single page. IDs not in the zone are silently omitted.
func (r *ZoneGroupService) List(ctx context.Context, zoneID string, query ZoneGroupListParams, opts ...option.RequestOption) (res *ZoneGroupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a group and its memberships and role assignments. Groups synced from an
// external directory can only be deleted by that directory (after external sync is
// disabled).
func (r *ZoneGroupService) Delete(ctx context.Context, groupID string, body ZoneGroupDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/groups/%s", url.PathEscape(body.ZoneID), url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A zone-scoped group of users, assignable to roles and usable in policies. Roles
// assigned to a group are inherited by its members. `external` is false for groups
// managed in Keycard and true for groups synced from an external directory.
type Group struct {
	// Unique identifier of the group
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the group is synced from an external directory. When true the group is
	// directory-owned and its membership is read-only; when false it is managed in
	// Keycard. Read-only: set by external sync, never by the caller.
	External bool `json:"external" api:"required"`
	// User-specified identifier, unique within the zone. Automatically assigned for
	// groups from an external directory.
	Identifier string `json:"identifier" api:"required"`
	// Human-readable group name
	Name string `json:"name" api:"required"`
	// Organization this group belongs to
	OrganizationID string `json:"organization_id" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zone this group belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// Number of users in the group. Included only when requested via
	// `expand[]=member_count` (group get or list).
	MemberCount int64 `json:"member_count"`
	// Identifiers of the roles assigned to the group; members inherit them. Deduped
	// across scopes. Included only when requested via `expand[]=roles` (group get or
	// list).
	Roles []string `json:"roles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		External       respjson.Field
		Identifier     respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		UpdatedAt      respjson.Field
		ZoneID         respjson.Field
		MemberCount    respjson.Field
		Roles          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for creating a group in Keycard. Groups synced from an external directory
// are created by that directory, not through this endpoint.
//
// The property Name is required.
type GroupCreateParam struct {
	// Human-readable group name
	Name string `json:"name" api:"required" format:"safe-text"`
	// User-specified identifier, unique within the zone. Derived from the name when
	// omitted (a suffix is appended if it collides).
	Identifier param.Opt[string] `json:"identifier,omitzero" format:"safe-text"`
	paramObj
}

func (r GroupCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow GroupCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for updating a group.
type GroupUpdateParam struct {
	// User-specified identifier, unique within the zone.
	Identifier param.Opt[string] `json:"identifier,omitzero" format:"safe-text"`
	// Human-readable group name
	Name param.Opt[string] `json:"name,omitzero" format:"safe-text"`
	paramObj
}

func (r GroupUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow GroupUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupListResponse struct {
	Items []Group `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneGroupListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneGroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneGroupListResponsePagination struct {
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
func (r ZoneGroupListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneGroupListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupNewParams struct {
	// Schema for creating a group in Keycard. Groups synced from an external directory
	// are created by that directory, not through this endpoint.
	GroupCreate GroupCreateParam
	paramObj
}

func (r ZoneGroupNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GroupCreate)
}
func (r *ZoneGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupGetParams struct {
	ZoneID string                        `path:"zoneId" api:"required" json:"-"`
	Expand ZoneGroupGetParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneGroupGetParams]'s query parameters as `url.Values`.
func (r ZoneGroupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupGetParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneGroupGetsExpandString)
	OfZoneGroupGetsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneGroupGetsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGroupGetParamsExpandString string

const (
	ZoneGroupGetParamsExpandStringMemberCount ZoneGroupGetParamsExpandString = "member_count"
	ZoneGroupGetParamsExpandStringRoles       ZoneGroupGetParamsExpandString = "roles"
)

type ZoneGroupUpdateParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Schema for updating a group.
	GroupUpdate GroupUpdateParam
	paramObj
}

func (r ZoneGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GroupUpdate)
}
func (r *ZoneGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Comma-separated sort fields. Prefix with - for descending. Allowed: created_at,
	// name, identifier
	Sort   param.Opt[string]              `query:"sort,omitzero" json:"-"`
	Expand ZoneGroupListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Restrict results to groups with this ID. Repeatable, max 100. Mutually exclusive
	// with after/before.
	FilterID ZoneGroupListParamsFilterIDUnion `query:"filter[id],omitzero" json:"-"`
	// Filter by exact group identifier
	FilterIdentifier ZoneGroupListParamsFilterIdentifierUnion `query:"filter[identifier],omitzero" json:"-"`
	// Search across name and identifier (substring match)
	Query ZoneGroupListParamsQueryUnion `query:"query[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneGroupListParams]'s query parameters as `url.Values`.
func (r ZoneGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneGroupListsExpandString)
	OfZoneGroupListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneGroupListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGroupListParamsExpandString string

const (
	ZoneGroupListParamsExpandStringTotalCount  ZoneGroupListParamsExpandString = "total_count"
	ZoneGroupListParamsExpandStringMemberCount ZoneGroupListParamsExpandString = "member_count"
	ZoneGroupListParamsExpandStringRoles       ZoneGroupListParamsExpandString = "roles"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupListParamsFilterIDUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupListParamsFilterIdentifierUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupListParamsQueryUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGroupDeleteParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}
