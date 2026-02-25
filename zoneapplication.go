// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/keycardlabs/keycard-go/internal/apijson"
	"github.com/keycardlabs/keycard-go/internal/apiquery"
	"github.com/keycardlabs/keycard-go/internal/requestconfig"
	"github.com/keycardlabs/keycard-go/option"
	"github.com/keycardlabs/keycard-go/packages/param"
	"github.com/keycardlabs/keycard-go/packages/respjson"
)

// ZoneApplicationService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneApplicationService] method instead.
type ZoneApplicationService struct {
	Options      []option.RequestOption
	Dependencies ZoneApplicationDependencyService
}

// NewZoneApplicationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneApplicationService(opts ...option.RequestOption) (r ZoneApplicationService) {
	r = ZoneApplicationService{}
	r.Options = opts
	r.Dependencies = NewZoneApplicationDependencyService(opts...)
	return
}

// Creates a new Application - a software system with an identity that can access
// Resources
func (r *ZoneApplicationService) New(ctx context.Context, zoneID string, body ZoneApplicationNewParams, opts ...option.RequestOption) (res *Application, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details of a specific Application by ID
func (r *ZoneApplicationService) Get(ctx context.Context, id string, query ZoneApplicationGetParams, opts ...option.RequestOption) (res *Application, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an Application's configuration and metadata
func (r *ZoneApplicationService) Update(ctx context.Context, id string, params ZoneApplicationUpdateParams, opts ...option.RequestOption) (res *Application, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of applications in the specified zone
func (r *ZoneApplicationService) List(ctx context.Context, zoneID string, query ZoneApplicationListParams, opts ...option.RequestOption) (res *ZoneApplicationListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes an application
func (r *ZoneApplicationService) Delete(ctx context.Context, id string, body ZoneApplicationDeleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s", url.PathEscape(body.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns a list of application credentials for a specific application
func (r *ZoneApplicationService) ListCredentials(ctx context.Context, id string, params ZoneApplicationListCredentialsParams, opts ...option.RequestOption) (res *ZoneApplicationListCredentialsResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s/application-credentials", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns a list of resources provided by an application
func (r *ZoneApplicationService) ListResources(ctx context.Context, id string, params ZoneApplicationListResourcesParams, opts ...option.RequestOption) (res *ZoneApplicationListResourcesResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s/resources", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// An Application is a software system with an associated identity that can access
// Resources. It may act on its own behalf (machine-to-machine) or on behalf of a
// user (delegated access).
type Application struct {
	// Unique identifier of the application
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Number of resource dependencies
	DependenciesCount int64 `json:"dependencies_count" api:"required"`
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier" api:"required"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Organization that owns this application
	OrganizationID string `json:"organization_id" api:"required"`
	// URL-safe identifier, unique within the zone
	Slug string `json:"slug" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zone this application belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// Human-readable description
	Description string `json:"description" api:"nullable"`
	// Entity metadata
	Metadata Metadata `json:"metadata"`
	// Protocol-specific configuration
	Protocols ApplicationProtocols `json:"protocols" api:"nullable"`
	// Traits of the application
	Traits []ApplicationTrait `json:"traits" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		DependenciesCount respjson.Field
		Identifier        respjson.Field
		Name              respjson.Field
		OrganizationID    respjson.Field
		Slug              respjson.Field
		UpdatedAt         respjson.Field
		ZoneID            respjson.Field
		Description       respjson.Field
		Metadata          respjson.Field
		Protocols         respjson.Field
		Traits            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Application) RawJSON() string { return r.JSON.raw }
func (r *Application) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol-specific configuration
type ApplicationProtocols struct {
	// OAuth 2.0 protocol configuration
	Oauth2 ApplicationProtocolsOauth2 `json:"oauth2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Oauth2      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationProtocols) RawJSON() string { return r.JSON.raw }
func (r *ApplicationProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration
type ApplicationProtocolsOauth2 struct {
	// OAuth 2.0 post-logout redirect URIs for this application
	PostLogoutRedirectUris []string `json:"post_logout_redirect_uris" api:"nullable" format:"uri"`
	// OAuth 2.0 redirect URIs for this application
	RedirectUris []string `json:"redirect_uris" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PostLogoutRedirectUris respjson.Field
		RedirectUris           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationProtocolsOauth2) RawJSON() string { return r.JSON.raw }
func (r *ApplicationProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Traits ascribe behaviors and characteristics to an application, which may
// activate trait-specific user experiences, workflows, or other system behaviors
type ApplicationTrait string

const (
	ApplicationTraitGateway     ApplicationTrait = "gateway"
	ApplicationTraitMcpProvider ApplicationTrait = "mcp-provider"
)

// Entity metadata
type Metadata struct {
	// Documentation URL
	DocsURL string `json:"docs_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocsURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metadata) RawJSON() string { return r.JSON.raw }
func (r *Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Metadata to a MetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MetadataParam.Overrides()
func (r Metadata) ToParam() MetadataParam {
	return param.Override[MetadataParam](json.RawMessage(r.RawJSON()))
}

// Entity metadata
type MetadataParam struct {
	// Documentation URL
	DocsURL param.Opt[string] `json:"docs_url,omitzero" format:"uri"`
	paramObj
}

func (r MetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow MetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entity metadata (set to null or {} to remove metadata)
type MetadataUpdateParam struct {
	// Documentation URL (set to null to unset)
	DocsURL param.Opt[string] `json:"docs_url,omitzero" format:"uri"`
	paramObj
}

func (r MetadataUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow MetadataUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetadataUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationListResponse struct {
	Items []Application `json:"items" api:"required"`
	// Pagination information
	PageInfo PageInfoPagination `json:"page_info" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneApplicationListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneApplicationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneApplicationListResponsePagination struct {
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
func (r ZoneApplicationListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationListCredentialsResponse struct {
	Items []CredentialUnion `json:"items" api:"required"`
	// Pagination information
	PageInfo PageInfoPagination `json:"page_info" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneApplicationListCredentialsResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneApplicationListCredentialsResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationListCredentialsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneApplicationListCredentialsResponsePagination struct {
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
func (r ZoneApplicationListCredentialsResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationListCredentialsResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationListResourcesResponse struct {
	Items []Resource `json:"items" api:"required"`
	// Pagination information
	PageInfo PageInfoPagination `json:"page_info" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneApplicationListResourcesResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneApplicationListResourcesResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationListResourcesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneApplicationListResourcesResponsePagination struct {
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
func (r ZoneApplicationListResourcesResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationListResourcesResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationNewParams struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier" api:"required"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// Dependencies of the application
	Dependencies []ZoneApplicationNewParamsDependency `json:"dependencies,omitzero"`
	// Entity metadata
	Metadata MetadataParam `json:"metadata,omitzero"`
	// Protocol-specific configuration for application creation
	Protocols ZoneApplicationNewParamsProtocols `json:"protocols,omitzero"`
	// Traits of the application
	Traits []ApplicationTrait `json:"traits,omitzero"`
	paramObj
}

func (r ZoneApplicationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type ZoneApplicationNewParamsDependency struct {
	// Resource identifier
	ID   string            `json:"id" api:"required"`
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ZoneApplicationNewParamsDependency) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationNewParamsDependency
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationNewParamsDependency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol-specific configuration for application creation
type ZoneApplicationNewParamsProtocols struct {
	// OAuth 2.0 protocol configuration for application creation
	Oauth2 ZoneApplicationNewParamsProtocolsOauth2 `json:"oauth2,omitzero"`
	paramObj
}

func (r ZoneApplicationNewParamsProtocols) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationNewParamsProtocols
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationNewParamsProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration for application creation
type ZoneApplicationNewParamsProtocolsOauth2 struct {
	// OAuth 2.0 post-logout redirect URIs for this application
	PostLogoutRedirectUris []string `json:"post_logout_redirect_uris,omitzero" format:"uri"`
	// OAuth 2.0 redirect URIs for this application
	RedirectUris []string `json:"redirect_uris,omitzero" format:"uri"`
	paramObj
}

func (r ZoneApplicationNewParamsProtocolsOauth2) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationNewParamsProtocolsOauth2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationNewParamsProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationGetParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneApplicationUpdateParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// User specified identifier, unique within the zone
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	// Human-readable name
	Name param.Opt[string] `json:"name,omitzero"`
	// Entity metadata (set to null or {} to remove metadata)
	Metadata MetadataUpdateParam `json:"metadata,omitzero"`
	// Protocol-specific configuration for application update
	Protocols ZoneApplicationUpdateParamsProtocols `json:"protocols,omitzero"`
	// Traits of the application
	Traits []ApplicationTrait `json:"traits,omitzero"`
	paramObj
}

func (r ZoneApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol-specific configuration for application update
type ZoneApplicationUpdateParamsProtocols struct {
	// OAuth 2.0 protocol configuration for application update
	Oauth2 ZoneApplicationUpdateParamsProtocolsOauth2 `json:"oauth2,omitzero"`
	paramObj
}

func (r ZoneApplicationUpdateParamsProtocols) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationUpdateParamsProtocols
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationUpdateParamsProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration for application update
type ZoneApplicationUpdateParamsProtocolsOauth2 struct {
	// OAuth 2.0 post-logout redirect URIs for this application (set to null or [] to
	// unset)
	PostLogoutRedirectUris []string `json:"post_logout_redirect_uris,omitzero" format:"uri"`
	// OAuth 2.0 redirect URIs for this application (set to null or [] to unset)
	RedirectUris []string `json:"redirect_uris,omitzero" format:"uri"`
	paramObj
}

func (r ZoneApplicationUpdateParamsProtocolsOauth2) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationUpdateParamsProtocolsOauth2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationUpdateParamsProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before     param.Opt[string] `query:"before,omitzero" json:"-"`
	Cursor     param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Identifier param.Opt[string] `query:"identifier,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                     `query:"limit,omitzero" json:"-"`
	Slug   param.Opt[string]                    `query:"slug,omitzero" json:"-"`
	Expand ZoneApplicationListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Filter by traits (OR matching - returns applications with any of the specified
	// traits)
	Traits []ApplicationTrait `query:"traits,omitzero" json:"-"`
	// Filter by traits (AND matching - returns applications with all of the specified
	// traits)
	TraitsAll []ApplicationTrait `query:"traits[all],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationListParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneApplicationListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneApplicationListsExpandString)
	OfZoneApplicationListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneApplicationListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneApplicationListParamsExpandString string

const (
	ZoneApplicationListParamsExpandStringTotalCount ZoneApplicationListParamsExpandString = "total_count"
)

type ZoneApplicationDeleteParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneApplicationListCredentialsParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                                `query:"limit,omitzero" json:"-"`
	Expand ZoneApplicationListCredentialsParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationListCredentialsParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationListCredentialsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneApplicationListCredentialsParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneApplicationListCredentialssExpandString)
	OfZoneApplicationListCredentialssExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneApplicationListCredentialssExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneApplicationListCredentialsParamsExpandString string

const (
	ZoneApplicationListCredentialsParamsExpandStringTotalCount ZoneApplicationListCredentialsParamsExpandString = "total_count"
)

type ZoneApplicationListResourcesParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                              `query:"limit,omitzero" json:"-"`
	Expand ZoneApplicationListResourcesParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationListResourcesParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationListResourcesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneApplicationListResourcesParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneApplicationListResourcessExpandString)
	OfZoneApplicationListResourcessExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneApplicationListResourcessExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneApplicationListResourcesParamsExpandString string

const (
	ZoneApplicationListResourcesParamsExpandStringTotalCount ZoneApplicationListResourcesParamsExpandString = "total_count"
)
