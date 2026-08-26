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

// ZoneGroupMemberService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneGroupMemberService] method instead.
type ZoneGroupMemberService struct {
	Options []option.RequestOption
}

// NewZoneGroupMemberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneGroupMemberService(opts ...option.RequestOption) (r ZoneGroupMemberService) {
	r = ZoneGroupMemberService{}
	r.Options = opts
	return
}

// Returns a paginated list of the group's members. Use cursor pagination via
// `after`/`before`. Pass `expand[]=user` to embed each member's full user record
// and `expand[]=total_count` to include the matching row count. Pass `query[]`
// (repeatable, 1-255 chars) to search members by their user's email or federated
// credential subject (substring match, OR'd across repeated values). Pass
// `filter[id]` (repeatable, max 100) to restrict results to a known set of members
// by user ID — mutually exclusive with `after`/`before` (returns 400 if combined).
// When `filter[id]` is set, `limit` is ignored and the response contains every
// requested member that exists in the group, in a single page. IDs not in the
// group are silently omitted.
func (r *ZoneGroupMemberService) List(ctx context.Context, groupID string, params ZoneGroupMemberListParams, opts ...option.RequestOption) (res *ZoneGroupMemberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups/%s/members", url.PathEscape(params.ZoneID), url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Adds a user to a group managed in Keycard. Membership of externally synced
// groups is not managed manually.
func (r *ZoneGroupMemberService) Add(ctx context.Context, groupID string, params ZoneGroupMemberAddParams, opts ...option.RequestOption) (res *GroupMember, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/groups/%s/members", url.PathEscape(params.ZoneID), url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes a user from a group managed in Keycard. Membership of externally synced
// groups is not managed manually. A member is identified by its user's ID.
func (r *ZoneGroupMemberService) Remove(ctx context.Context, userID string, body ZoneGroupMemberRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	if body.GroupID == "" {
		err = errors.New("missing required groupId parameter")
		return err
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/groups/%s/members/%s", url.PathEscape(body.ZoneID), url.PathEscape(body.GroupID), url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A user's membership in a group
type GroupMember struct {
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of the user
	UserID string `json:"user_id" api:"required"`
	// An authenticated user entity
	User User `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		UserID      respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupMember) RawJSON() string { return r.JSON.raw }
func (r *GroupMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for adding a user to a group
//
// The property UserID is required.
type GroupMemberCreateParam struct {
	// ID of the user to add to the group
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r GroupMemberCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow GroupMemberCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupMemberCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupMemberListResponse struct {
	Items []GroupMember `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneGroupMemberListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGroupMemberListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneGroupMemberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneGroupMemberListResponsePagination struct {
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
func (r ZoneGroupMemberListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneGroupMemberListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupMemberListParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                     `query:"limit,omitzero" json:"-"`
	Expand ZoneGroupMemberListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Restrict results to the member with this user ID. Repeatable, max 100. Mutually
	// exclusive with after/before.
	FilterID ZoneGroupMemberListParamsFilterIDUnion `query:"filter[id],omitzero" json:"-"`
	// Search members by their user's email or federated credential subject (substring
	// match)
	Query ZoneGroupMemberListParamsQueryUnion `query:"query[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneGroupMemberListParams]'s query parameters as
// `url.Values`.
func (r ZoneGroupMemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupMemberListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneGroupMemberListsExpandString)
	OfZoneGroupMemberListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneGroupMemberListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGroupMemberListParamsExpandString string

const (
	ZoneGroupMemberListParamsExpandStringTotalCount ZoneGroupMemberListParamsExpandString = "total_count"
	ZoneGroupMemberListParamsExpandStringUser       ZoneGroupMemberListParamsExpandString = "user"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupMemberListParamsFilterIDUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGroupMemberListParamsQueryUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGroupMemberAddParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Schema for adding a user to a group
	GroupMemberCreate GroupMemberCreateParam
	paramObj
}

func (r ZoneGroupMemberAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GroupMemberCreate)
}
func (r *ZoneGroupMemberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGroupMemberRemoveParams struct {
	ZoneID  string `path:"zoneId" api:"required" json:"-"`
	GroupID string `path:"groupId" api:"required" json:"-"`
	paramObj
}
