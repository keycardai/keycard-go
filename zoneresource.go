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
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// ZoneResourceService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneResourceService] method instead.
type ZoneResourceService struct {
	Options []option.RequestOption
}

// NewZoneResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneResourceService(opts ...option.RequestOption) (r ZoneResourceService) {
	r = ZoneResourceService{}
	r.Options = opts
	return
}

// Creates a new Resource - a system that exposes protected information or
// functionality requiring authentication
func (r *ZoneResourceService) New(ctx context.Context, zoneID string, body ZoneResourceNewParams, opts ...option.RequestOption) (res *Resource, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/resources", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns details of a specific Resource by ID
func (r *ZoneResourceService) Get(ctx context.Context, id string, query ZoneResourceGetParams, opts ...option.RequestOption) (res *Resource, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/resources/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a Resource's configuration and metadata
func (r *ZoneResourceService) Update(ctx context.Context, id string, params ZoneResourceUpdateParams, opts ...option.RequestOption) (res *Resource, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/resources/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a list of resources in the specified zone
func (r *ZoneResourceService) List(ctx context.Context, zoneID string, query ZoneResourceListParams, opts ...option.RequestOption) (res *ZoneResourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/resources", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a Resource
func (r *ZoneResourceService) Delete(ctx context.Context, id string, body ZoneResourceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/resources/%s", url.PathEscape(body.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ZoneResourceListResponse struct {
	Items []Resource `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneResourceListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneResourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneResourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneResourceListResponsePagination struct {
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
func (r ZoneResourceListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneResourceListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneResourceNewParams struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier" api:"required"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the application that provides this resource
	ApplicationID param.Opt[string] `json:"application_id,omitzero"`
	// ID of the credential provider to associate with the resource
	CredentialProviderID param.Opt[string] `json:"credential_provider_id,omitzero"`
	// The expected type of client for this credential. Native clients must use
	// localhost URLs for redirect_uris or URIs with custom schemes. Web clients must
	// use https URLs and must not use localhost as the hostname.
	//
	// Any of "native", "web".
	ApplicationType ZoneResourceNewParamsApplicationType `json:"application_type,omitzero"`
	// Entity metadata
	Metadata MetadataParam `json:"metadata,omitzero"`
	// Scopes supported by the resource
	Scopes []string `json:"scopes,omitzero"`
	paramObj
}

func (r ZoneResourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneResourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneResourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The expected type of client for this credential. Native clients must use
// localhost URLs for redirect_uris or URIs with custom schemes. Web clients must
// use https URLs and must not use localhost as the hostname.
type ZoneResourceNewParamsApplicationType string

const (
	ZoneResourceNewParamsApplicationTypeNative ZoneResourceNewParamsApplicationType = "native"
	ZoneResourceNewParamsApplicationTypeWeb    ZoneResourceNewParamsApplicationType = "web"
)

type ZoneResourceGetParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneResourceUpdateParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// ID of the application that provides this resource (set to null to unset)
	ApplicationID param.Opt[string] `json:"application_id,omitzero"`
	// ID of the credential provider to associate with the resource (set to null to
	// unset)
	CredentialProviderID param.Opt[string] `json:"credential_provider_id,omitzero"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// User specified identifier, unique within the zone
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	// Human-readable name
	Name param.Opt[string] `json:"name,omitzero"`
	// Entity metadata (set to null or {} to remove metadata)
	Metadata MetadataUpdateParam `json:"metadata,omitzero"`
	// Scopes supported by the resource (set to null to unset)
	Scopes []string `json:"scopes,omitzero"`
	// The expected type of client for this credential. Native clients must use
	// localhost URLs for redirect_uris or URIs with custom schemes. Web clients must
	// use https URLs and must not use localhost as the hostname.
	//
	// Any of "native", "web".
	ApplicationType ZoneResourceUpdateParamsApplicationType `json:"application_type,omitzero"`
	paramObj
}

func (r ZoneResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneResourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneResourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The expected type of client for this credential. Native clients must use
// localhost URLs for redirect_uris or URIs with custom schemes. Web clients must
// use https URLs and must not use localhost as the hostname.
type ZoneResourceUpdateParamsApplicationType string

const (
	ZoneResourceUpdateParamsApplicationTypeNative ZoneResourceUpdateParamsApplicationType = "native"
	ZoneResourceUpdateParamsApplicationTypeWeb    ZoneResourceUpdateParamsApplicationType = "web"
)

type ZoneResourceListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter resources by credential provider ID
	CredentialProviderID param.Opt[string] `query:"credentialProviderId,omitzero" json:"-"`
	// Filter resources by identifier
	Identifier param.Opt[string] `query:"identifier,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                  `query:"limit,omitzero" json:"-"`
	Slug   param.Opt[string]                 `query:"slug,omitzero" json:"-"`
	Expand ZoneResourceListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneResourceListParams]'s query parameters as `url.Values`.
func (r ZoneResourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneResourceListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneResourceListsExpandString)
	OfZoneResourceListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneResourceListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneResourceListParamsExpandString string

const (
	ZoneResourceListParamsExpandStringTotalCount ZoneResourceListParamsExpandString = "total_count"
)

type ZoneResourceDeleteParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}
