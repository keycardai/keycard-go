// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/keycard-api-go/internal/apijson"
	"github.com/stainless-sdks/keycard-api-go/internal/apiquery"
	"github.com/stainless-sdks/keycard-api-go/internal/requestconfig"
	"github.com/stainless-sdks/keycard-api-go/option"
	"github.com/stainless-sdks/keycard-api-go/packages/param"
	"github.com/stainless-sdks/keycard-api-go/packages/respjson"
)

// ZoneMemberService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneMemberService] method instead.
type ZoneMemberService struct {
	Options []option.RequestOption
}

// NewZoneMemberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneMemberService(opts ...option.RequestOption) (r ZoneMemberService) {
	r = ZoneMemberService{}
	r.Options = opts
	return
}

// Returns detailed information about a specific organization user in a zone.
func (r *ZoneMemberService) Get(ctx context.Context, organizationUserID string, query ZoneMemberGetParams, opts ...option.RequestOption) (res *ZoneMember, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if organizationUserID == "" {
		err = errors.New("missing required organizationUserId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/members/%s", query.ZoneID, organizationUserID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the role of an existing zone member. Only organization administrators
// can perform this action.
func (r *ZoneMemberService) Update(ctx context.Context, organizationUserID string, params ZoneMemberUpdateParams, opts ...option.RequestOption) (res *ZoneMember, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if organizationUserID == "" {
		err = errors.New("missing required organizationUserId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/members/%s", params.ZoneID, organizationUserID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Lists all organization users in a zone with their roles and metadata. Supports
// cursor-based pagination.
func (r *ZoneMemberService) List(ctx context.Context, zoneID string, query ZoneMemberListParams, opts ...option.RequestOption) (res *ZoneMemberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/members", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes an organization user's membership from a zone. Only organization
// administrators can perform this action.
func (r *ZoneMemberService) Delete(ctx context.Context, organizationUserID string, body ZoneMemberDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if organizationUserID == "" {
		err = errors.New("missing required organizationUserId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/members/%s", body.ZoneID, organizationUserID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Adds an organization user to a zone with the specified role.
func (r *ZoneMemberService) Add(ctx context.Context, zoneID string, body ZoneMemberAddParams, opts ...option.RequestOption) (res *ZoneMember, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/members", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Represents an organization user's membership in a zone with an assigned role
type ZoneMember struct {
	// Unique identifier of the zone member
	ID string `json:"id,required"`
	// HAL-format hypermedia links for zone member resources
	Links ZoneMember_Links `json:"_links,required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Organization ID that owns the zone
	OrganizationID string `json:"organization_id,required"`
	// Organization user ID of the zone member
	OrganizationUserID string `json:"organization_user_id,required"`
	// Zone role type. zone_manager has full management access, zone_viewer has
	// read-only access.
	//
	// Any of "zone_manager", "zone_viewer".
	Role ZoneRole `json:"role,required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Zone ID the organization user is a member of
	ZoneID string `json:"zone_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Links              respjson.Field
		CreatedAt          respjson.Field
		OrganizationID     respjson.Field
		OrganizationUserID respjson.Field
		Role               respjson.Field
		UpdatedAt          respjson.Field
		ZoneID             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneMember) RawJSON() string { return r.JSON.raw }
func (r *ZoneMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HAL-format hypermedia links for zone member resources
type ZoneMember_Links struct {
	OrganizationUser ZoneMember_LinksOrganizationUser `json:"organization_user,required"`
	Self             ZoneMember_LinksSelf             `json:"self,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrganizationUser respjson.Field
		Self             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneMember_Links) RawJSON() string { return r.JSON.raw }
func (r *ZoneMember_Links) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneMember_LinksOrganizationUser struct {
	// Link to the user resource
	Href string `json:"href,required" format:"uri-reference"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneMember_LinksOrganizationUser) RawJSON() string { return r.JSON.raw }
func (r *ZoneMember_LinksOrganizationUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneMember_LinksSelf struct {
	// Link to this zone member resource
	Href string `json:"href,required" format:"uri-reference"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneMember_LinksSelf) RawJSON() string { return r.JSON.raw }
func (r *ZoneMember_LinksSelf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Zone role type. zone_manager has full management access, zone_viewer has
// read-only access.
type ZoneRole string

const (
	ZoneRoleZoneManager ZoneRole = "zone_manager"
	ZoneRoleZoneViewer  ZoneRole = "zone_viewer"
)

type ZoneMemberListResponse struct {
	Items []ZoneMember `json:"items,required"`
	// Pagination information
	PageInfo PageInfoPagination `json:"page_info,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneMemberListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneMemberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneMemberGetParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}

type ZoneMemberUpdateParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	// Zone role type. zone_manager has full management access, zone_viewer has
	// read-only access.
	//
	// Any of "zone_manager", "zone_viewer".
	Role ZoneRole `json:"role,omitzero,required"`
	paramObj
}

func (r ZoneMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMemberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMemberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneMemberListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of members to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter members by role
	//
	// Any of "zone_manager", "zone_viewer".
	Role ZoneMemberListParamsRole `query:"role,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneMemberListParams]'s query parameters as `url.Values`.
func (r ZoneMemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter members by role
type ZoneMemberListParamsRole string

const (
	ZoneMemberListParamsRoleZoneManager ZoneMemberListParamsRole = "zone_manager"
	ZoneMemberListParamsRoleZoneViewer  ZoneMemberListParamsRole = "zone_viewer"
)

type ZoneMemberDeleteParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}

type ZoneMemberAddParams struct {
	// Organization user ID to add to the zone
	OrganizationUserID string `json:"organization_user_id,required"`
	// Zone role type. zone_manager has full management access, zone_viewer has
	// read-only access.
	//
	// Any of "zone_manager", "zone_viewer".
	Role ZoneRole `json:"role,omitzero,required"`
	paramObj
}

func (r ZoneMemberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMemberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMemberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
