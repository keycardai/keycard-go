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

// ZoneRoleService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRoleService] method instead.
type ZoneRoleService struct {
	Options []option.RequestOption
}

// NewZoneRoleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRoleService(opts ...option.RequestOption) (r ZoneRoleService) {
	r = ZoneRoleService{}
	r.Options = opts
	return
}

// Creates a new customer-owned role in the specified zone. The owner_type is
// always customer; platform roles are managed by Keycard.
func (r *ZoneRoleService) New(ctx context.Context, zoneID string, body ZoneRoleNewParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/roles", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns details of a specific role by ID
func (r *ZoneRoleService) Get(ctx context.Context, roleID string, query ZoneRoleGetParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if roleID == "" {
		err = errors.New("missing required roleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/roles/%s", url.PathEscape(query.ZoneID), url.PathEscape(roleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a customer-owned role's description. The identifier is immutable, and
// platform-owned roles cannot be modified.
func (r *ZoneRoleService) Update(ctx context.Context, roleID string, params ZoneRoleUpdateParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if roleID == "" {
		err = errors.New("missing required roleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/roles/%s", url.PathEscape(params.ZoneID), url.PathEscape(roleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns the roles defined in the specified zone. The full result set is
// currently returned in a single page; the `after`/`before`/`limit` cursor
// parameters are reserved and not yet enforced, and `pagination` cursors are
// always null.
func (r *ZoneRoleService) List(ctx context.Context, zoneID string, query ZoneRoleListParams, opts ...option.RequestOption) (res *ZoneRoleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/roles", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a customer-owned role. Platform-owned roles cannot be
// deleted, and a role with existing assignments returns 409.
func (r *ZoneRoleService) Delete(ctx context.Context, roleID string, body ZoneRoleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	if roleID == "" {
		err = errors.New("missing required roleId parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/roles/%s", url.PathEscape(body.ZoneID), url.PathEscape(roleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A role that can be assigned to users within a zone.
type Role struct {
	// Unique identifier of the role
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Role identifier: a lowercase slug (letters and digits separated by single
	// hyphens or underscores), unique per owner type within a zone. Role identifiers
	// surface in policy evaluation, so the slug restriction keeps them unambiguous in
	// policy text.
	Identifier string `json:"identifier" api:"required"`
	// Who owns this role. Platform-owned roles are managed by Keycard and cannot be
	// modified or deleted via the API; customer-owned roles are user-created.
	//
	// Any of "platform", "customer".
	OwnerType RoleOwnerType `json:"owner_type" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zone this role belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// Human-readable description
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Identifier  respjson.Field
		OwnerType   respjson.Field
		UpdatedAt   respjson.Field
		ZoneID      respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Role) RawJSON() string { return r.JSON.raw }
func (r *Role) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who owns this role. Platform-owned roles are managed by Keycard and cannot be
// modified or deleted via the API; customer-owned roles are user-created.
type RoleOwnerType string

const (
	RoleOwnerTypePlatform RoleOwnerType = "platform"
	RoleOwnerTypeCustomer RoleOwnerType = "customer"
)

// Schema for creating a new role
//
// The property Identifier is required.
type RoleCreateParam struct {
	// Role identifier: a lowercase slug (letters and digits separated by single
	// hyphens or underscores), unique per owner type within a zone. Role identifiers
	// surface in policy evaluation, so the slug restriction keeps them unambiguous in
	// policy text.
	Identifier string `json:"identifier" api:"required"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r RoleCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow RoleCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoleCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for updating an existing role. The role identifier is immutable.
type RoleUpdateParam struct {
	// Human-readable description (set to null to unset)
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r RoleUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow RoleUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoleUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRoleListResponse struct {
	Items []Role `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneRoleListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneRoleListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneRoleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneRoleListResponsePagination struct {
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
func (r ZoneRoleListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneRoleListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRoleNewParams struct {
	// Schema for creating a new role
	RoleCreate RoleCreateParam
	paramObj
}

func (r ZoneRoleNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RoleCreate)
}
func (r *ZoneRoleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRoleGetParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneRoleUpdateParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Schema for updating an existing role. The role identifier is immutable.
	RoleUpdate RoleUpdateParam
	paramObj
}

func (r ZoneRoleUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RoleUpdate)
}
func (r *ZoneRoleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRoleListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter roles by identifier
	Identifier param.Opt[string] `query:"identifier,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]              `query:"limit,omitzero" json:"-"`
	Expand ZoneRoleListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneRoleListParams]'s query parameters as `url.Values`.
func (r ZoneRoleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneRoleListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneRoleListsExpandString)
	OfZoneRoleListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneRoleListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneRoleListParamsExpandString string

const (
	ZoneRoleListParamsExpandStringTotalCount ZoneRoleListParamsExpandString = "total_count"
)

type ZoneRoleDeleteParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}
