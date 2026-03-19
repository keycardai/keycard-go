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

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/apiquery"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// ZoneService contains methods and other services that help with interacting with
// the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneService] method instead.
type ZoneService struct {
	Options                []option.RequestOption
	Applications           ZoneApplicationService
	ApplicationCredentials ZoneApplicationCredentialService
	DelegatedGrants        ZoneDelegatedGrantService
	Providers              ZoneProviderService
	Resources              ZoneResourceService
	Sessions               ZoneSessionService
	UserAgents             ZoneUserAgentService
	Users                  ZoneUserService
	Members                ZoneMemberService
	Secrets                ZoneSecretService
	// Zone-scoped Cedar schema management.
	//
	// The Cedar schema defines the entity model used for authorization decisions. Key
	// entity types and their attributes:
	//
	//   - **Keycard::User** — `email` (String), `groups` (Set of String)
	//   - **Keycard::Application** — `registration_method` (RegistrationMethod entity),
	//     `credential_type` (CredentialType entity)
	//   - **Keycard::RegistrationMethod** — enum entity: `"managed"`, `"dcr"`
	//   - **Keycard::CredentialType** — enum entity: `"token"`, `"password"`,
	//     `"public-key"`, `"url"`, `"public"`
	//   - **Keycard::Resource** — `id` (String), `name` (String), `scopes` (Set of
	//     String)
	//   - **Keycard::Claims** — `email` (String), `groups` (Set of String), plus
	//     arbitrary additional fields
	//
	// Enum-like attributes use Cedar enum entity types (schema version `2026-03-16`+).
	// In policies, reference values as `RegistrationMethod::"managed"` or
	// `CredentialType::"token"`. See the Credentials API spec for the full entity
	// model reference.
	PolicySchemas ZonePolicySchemaService
	// Policy CRUD operations
	Policies ZonePolicyService
	// Policy set CRUD and binding management
	PolicySets ZonePolicySetService
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r ZoneService) {
	r = ZoneService{}
	r.Options = opts
	r.Applications = NewZoneApplicationService(opts...)
	r.ApplicationCredentials = NewZoneApplicationCredentialService(opts...)
	r.DelegatedGrants = NewZoneDelegatedGrantService(opts...)
	r.Providers = NewZoneProviderService(opts...)
	r.Resources = NewZoneResourceService(opts...)
	r.Sessions = NewZoneSessionService(opts...)
	r.UserAgents = NewZoneUserAgentService(opts...)
	r.Users = NewZoneUserService(opts...)
	r.Members = NewZoneMemberService(opts...)
	r.Secrets = NewZoneSecretService(opts...)
	r.PolicySchemas = NewZonePolicySchemaService(opts...)
	r.Policies = NewZonePolicyService(opts...)
	r.PolicySets = NewZonePolicySetService(opts...)
	return
}

// Creates a new zone for the authenticated organization. A zone is an isolated
// environment for IAM resources.
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *Zone, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns details of a specific zone by ID
func (r *ZoneService) Get(ctx context.Context, zoneID string, query ZoneGetParams, opts ...option.RequestOption) (res *Zone, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a zone's configuration (partial update)
func (r *ZoneService) Update(ctx context.Context, zoneID string, body ZoneUpdateParams, opts ...option.RequestOption) (res *Zone, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of zones for the authenticated organization
func (r *ZoneService) List(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) (res *ZoneListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a zone and all its associated resources
func (r *ZoneService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// AWS KMS configuration for zone encryption. When not specified, the default
// Keycard Cloud encryption key will be used.
type EncryptionKeyAwsKmsConfig struct {
	// AWS KMS Key ARN for encrypting the zone's data
	Arn string `json:"arn" api:"required"`
	// Any of "aws".
	Type EncryptionKeyAwsKmsConfigType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arn         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EncryptionKeyAwsKmsConfig) RawJSON() string { return r.JSON.raw }
func (r *EncryptionKeyAwsKmsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EncryptionKeyAwsKmsConfig to a
// EncryptionKeyAwsKmsConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EncryptionKeyAwsKmsConfigParam.Overrides()
func (r EncryptionKeyAwsKmsConfig) ToParam() EncryptionKeyAwsKmsConfigParam {
	return param.Override[EncryptionKeyAwsKmsConfigParam](json.RawMessage(r.RawJSON()))
}

type EncryptionKeyAwsKmsConfigType string

const (
	EncryptionKeyAwsKmsConfigTypeAws EncryptionKeyAwsKmsConfigType = "aws"
)

// AWS KMS configuration for zone encryption. When not specified, the default
// Keycard Cloud encryption key will be used.
//
// The properties Arn, Type are required.
type EncryptionKeyAwsKmsConfigParam struct {
	// AWS KMS Key ARN for encrypting the zone's data
	Arn string `json:"arn" api:"required"`
	// Any of "aws".
	Type EncryptionKeyAwsKmsConfigType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r EncryptionKeyAwsKmsConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow EncryptionKeyAwsKmsConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EncryptionKeyAwsKmsConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information
type PageInfoPagination struct {
	// Whether there are more items after the current page
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Whether there are items before the current page
	HasPreviousPage bool `json:"has_previous_page" api:"required"`
	// Cursor pointing to the last item in the current page
	EndCursor string `json:"end_cursor" api:"nullable"`
	// Cursor pointing to the first item in the current page
	StartCursor string `json:"start_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage     respjson.Field
		HasPreviousPage respjson.Field
		EndCursor       respjson.Field
		StartCursor     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageInfoPagination) RawJSON() string { return r.JSON.raw }
func (r *PageInfoPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A zone for organizing resources within an organization
type Zone struct {
	// Unique identifier of the zone
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Organization that owns this zone
	OrganizationID string `json:"organization_id" api:"required"`
	// Protocol configuration for a zone
	Protocols ZoneProtocols `json:"protocols" api:"required"`
	// URL-safe identifier, unique within the zone
	Slug string `json:"slug" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Application ID configured as the default MCP Gateway for the zone
	DefaultMcpGatewayApplicationID string `json:"default_mcp_gateway_application_id"`
	// Resource ID configured as the default resource for the zone
	DefaultResourceID string `json:"default_resource_id"`
	// Human-readable description
	Description string `json:"description" api:"nullable"`
	// AWS KMS configuration for zone encryption. When not specified, the default
	// Keycard Cloud encryption key will be used.
	EncryptionKey EncryptionKeyAwsKmsConfig `json:"encryption_key"`
	// Login flow style for the zone. 'default' uses standard authentication,
	// 'identifier_first' uses identifier-based provider routing.
	//
	// Any of "default", "identifier_first".
	LoginFlow ZoneLoginFlow `json:"login_flow"`
	// Permissions granted to the authenticated principal. Only populated when
	// expand[]=permissions query parameter is provided. Keys are resource types,
	// values are objects mapping action names to boolean values.
	Permissions map[string]map[string]bool `json:"permissions"`
	// Whether the zone requires an invitation for email/password registration, only
	// applies when user_identity_provider_id is not set
	RequiresInvitation bool `json:"requires_invitation"`
	// Provider ID configured for user login
	UserIdentityProviderID string `json:"user_identity_provider_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                             respjson.Field
		CreatedAt                      respjson.Field
		Name                           respjson.Field
		OrganizationID                 respjson.Field
		Protocols                      respjson.Field
		Slug                           respjson.Field
		UpdatedAt                      respjson.Field
		DefaultMcpGatewayApplicationID respjson.Field
		DefaultResourceID              respjson.Field
		Description                    respjson.Field
		EncryptionKey                  respjson.Field
		LoginFlow                      respjson.Field
		Permissions                    respjson.Field
		RequiresInvitation             respjson.Field
		UserIdentityProviderID         respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Zone) RawJSON() string { return r.JSON.raw }
func (r *Zone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol configuration for a zone
type ZoneProtocols struct {
	// OAuth 2.0 protocol configuration for a zone
	Oauth2 ZoneProtocolsOauth2 `json:"oauth2" api:"required"`
	// OpenID Connect protocol configuration for a zone
	Openid ZoneProtocolsOpenid `json:"openid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Oauth2      respjson.Field
		Openid      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneProtocols) RawJSON() string { return r.JSON.raw }
func (r *ZoneProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration for a zone
type ZoneProtocolsOauth2 struct {
	// OAuth 2.0 authorization endpoint
	AuthorizationEndpoint string `json:"authorization_endpoint" api:"required" format:"uri"`
	// OAuth 2.0 Authorization Server Metadata endpoint
	// (.well-known/oauth-authorization-server)
	AuthorizationServerMetadata string `json:"authorization_server_metadata" api:"required" format:"uri"`
	// Whether Dynamic Client Registration is enabled
	DcrEnabled bool `json:"dcr_enabled" api:"required"`
	// OAuth 2.0 issuer identifier
	Issuer string `json:"issuer" api:"required" format:"uri"`
	// JSON Web Key Set endpoint
	JwksUri string `json:"jwks_uri" api:"required" format:"uri"`
	// Whether PKCE is required for authorization code flows
	PkceRequired bool `json:"pkce_required" api:"required"`
	// OAuth 2.0 redirect URI for this zone
	RedirectUri string `json:"redirect_uri" api:"required" format:"uri"`
	// OAuth 2.0 Dynamic Client Registration endpoint
	RegistrationEndpoint string `json:"registration_endpoint" api:"required" format:"uri"`
	// OAuth 2.0 token endpoint
	TokenEndpoint string `json:"token_endpoint" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationEndpoint       respjson.Field
		AuthorizationServerMetadata respjson.Field
		DcrEnabled                  respjson.Field
		Issuer                      respjson.Field
		JwksUri                     respjson.Field
		PkceRequired                respjson.Field
		RedirectUri                 respjson.Field
		RegistrationEndpoint        respjson.Field
		TokenEndpoint               respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneProtocolsOauth2) RawJSON() string { return r.JSON.raw }
func (r *ZoneProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenID Connect protocol configuration for a zone
type ZoneProtocolsOpenid struct {
	// OpenID Connect Provider Configuration endpoint
	// (.well-known/openid-configuration)
	ProviderConfiguration string `json:"provider_configuration" api:"required" format:"uri"`
	// OpenID Connect UserInfo endpoint
	UserinfoEndpoint string `json:"userinfo_endpoint" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProviderConfiguration respjson.Field
		UserinfoEndpoint      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneProtocolsOpenid) RawJSON() string { return r.JSON.raw }
func (r *ZoneProtocolsOpenid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Login flow style for the zone. 'default' uses standard authentication,
// 'identifier_first' uses identifier-based provider routing.
type ZoneLoginFlow string

const (
	ZoneLoginFlowDefault         ZoneLoginFlow = "default"
	ZoneLoginFlowIdentifierFirst ZoneLoginFlow = "identifier_first"
)

type ZoneListResponse struct {
	Items []Zone `json:"items" api:"required"`
	// Pagination information
	PageInfo PageInfoPagination `json:"page_info" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneListResponsePagination `json:"pagination" api:"required"`
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
func (r ZoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneListResponsePagination struct {
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
func (r ZoneListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewParams struct {
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// Assign a default MCP Gateway application to the zone
	DefaultMcpGatewayApplication param.Opt[bool] `json:"default_mcp_gateway_application,omitzero"`
	// Whether the zone requires an invitation for email/password registration, only
	// applies when user_identity_provider_id is not set. Defaults to true.
	RequiresInvitation param.Opt[bool] `json:"requires_invitation,omitzero"`
	// AWS KMS configuration for zone encryption. When not specified, the default
	// Keycard Cloud encryption key will be used.
	EncryptionKey EncryptionKeyAwsKmsConfigParam `json:"encryption_key,omitzero"`
	// Login flow style for the zone. 'default' uses standard authentication,
	// 'identifier_first' uses identifier-based provider routing.
	//
	// Any of "default", "identifier_first".
	LoginFlow ZoneNewParamsLoginFlow `json:"login_flow,omitzero"`
	// Protocol configuration for zone creation
	Protocols ZoneNewParamsProtocols `json:"protocols,omitzero"`
	paramObj
}

func (r ZoneNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Login flow style for the zone. 'default' uses standard authentication,
// 'identifier_first' uses identifier-based provider routing.
type ZoneNewParamsLoginFlow string

const (
	ZoneNewParamsLoginFlowDefault         ZoneNewParamsLoginFlow = "default"
	ZoneNewParamsLoginFlowIdentifierFirst ZoneNewParamsLoginFlow = "identifier_first"
)

// Protocol configuration for zone creation
type ZoneNewParamsProtocols struct {
	// OAuth 2.0 protocol configuration for zone creation
	Oauth2 ZoneNewParamsProtocolsOauth2 `json:"oauth2,omitzero"`
	paramObj
}

func (r ZoneNewParamsProtocols) MarshalJSON() (data []byte, err error) {
	type shadow ZoneNewParamsProtocols
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneNewParamsProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration for zone creation
type ZoneNewParamsProtocolsOauth2 struct {
	// Whether Dynamic Client Registration is enabled
	DcrEnabled param.Opt[bool] `json:"dcr_enabled,omitzero"`
	// Whether PKCE is required for authorization code flows
	PkceRequired param.Opt[bool] `json:"pkce_required,omitzero"`
	paramObj
}

func (r ZoneNewParamsProtocolsOauth2) MarshalJSON() (data []byte, err error) {
	type shadow ZoneNewParamsProtocolsOauth2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneNewParamsProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetParams struct {
	Expand ZoneGetParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneGetParams]'s query parameters as `url.Values`.
func (r ZoneGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneGetParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneGetsExpandString)
	OfZoneGetsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneGetsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneGetParamsExpandString string

const (
	ZoneGetParamsExpandStringPermissions ZoneGetParamsExpandString = "permissions"
)

type ZoneUpdateParams struct {
	// Application ID configured as the default MCP Gateway for the zone (set to null
	// to unset)
	DefaultMcpGatewayApplicationID param.Opt[string] `json:"default_mcp_gateway_application_id,omitzero"`
	// Resource ID to configure as the default resource for the zone (set to null to
	// unset)
	DefaultResourceID param.Opt[string] `json:"default_resource_id,omitzero"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// Provider ID to configure for user login (set to null to unset)
	UserIdentityProviderID param.Opt[string] `json:"user_identity_provider_id,omitzero"`
	// Human-readable name
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether the zone requires an invitation for email/password registration, only
	// applies when user_identity_provider_id is not set
	RequiresInvitation param.Opt[bool] `json:"requires_invitation,omitzero"`
	// AWS KMS configuration for zone encryption update (set to null to remove
	// customer-managed key and revert to default)
	EncryptionKey ZoneUpdateParamsEncryptionKey `json:"encryption_key,omitzero"`
	// Login flow style for the zone. 'default' uses standard authentication,
	// 'identifier_first' uses identifier-based provider routing. Set to null to reset
	// to 'default'.
	//
	// Any of "default", "identifier_first".
	LoginFlow ZoneUpdateParamsLoginFlow `json:"login_flow,omitzero"`
	// Protocol configuration update for a zone (partial update)
	Protocols ZoneUpdateParamsProtocols `json:"protocols,omitzero"`
	paramObj
}

func (r ZoneUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS KMS configuration for zone encryption update (set to null to remove
// customer-managed key and revert to default)
//
// The properties Arn, Type are required.
type ZoneUpdateParamsEncryptionKey struct {
	// AWS KMS Key ARN for encrypting the zone's data
	Arn string `json:"arn" api:"required"`
	// Any of "aws".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ZoneUpdateParamsEncryptionKey) MarshalJSON() (data []byte, err error) {
	type shadow ZoneUpdateParamsEncryptionKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneUpdateParamsEncryptionKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneUpdateParamsEncryptionKey](
		"type", "aws",
	)
}

// Login flow style for the zone. 'default' uses standard authentication,
// 'identifier_first' uses identifier-based provider routing. Set to null to reset
// to 'default'.
type ZoneUpdateParamsLoginFlow string

const (
	ZoneUpdateParamsLoginFlowDefault         ZoneUpdateParamsLoginFlow = "default"
	ZoneUpdateParamsLoginFlowIdentifierFirst ZoneUpdateParamsLoginFlow = "identifier_first"
)

// Protocol configuration update for a zone (partial update)
type ZoneUpdateParamsProtocols struct {
	// OAuth 2.0 protocol configuration update for a zone (partial update)
	Oauth2 ZoneUpdateParamsProtocolsOauth2 `json:"oauth2,omitzero"`
	paramObj
}

func (r ZoneUpdateParamsProtocols) MarshalJSON() (data []byte, err error) {
	type shadow ZoneUpdateParamsProtocols
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneUpdateParamsProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration update for a zone (partial update)
type ZoneUpdateParamsProtocolsOauth2 struct {
	// Whether Dynamic Client Registration is enabled
	DcrEnabled param.Opt[bool] `json:"dcr_enabled,omitzero"`
	// Whether PKCE is required for authorization code flows
	PkceRequired param.Opt[bool] `json:"pkce_required,omitzero"`
	paramObj
}

func (r ZoneUpdateParamsProtocolsOauth2) MarshalJSON() (data []byte, err error) {
	type shadow ZoneUpdateParamsProtocolsOauth2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneUpdateParamsProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]          `query:"limit,omitzero" json:"-"`
	Slug   param.Opt[string]         `query:"slug,omitzero" json:"-"`
	Expand ZoneListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneListParams]'s query parameters as `url.Values`.
func (r ZoneListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneListsExpandString)
	OfZoneListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneListParamsExpandString string

const (
	ZoneListParamsExpandStringTotalCount  ZoneListParamsExpandString = "total_count"
	ZoneListParamsExpandStringPermissions ZoneListParamsExpandString = "permissions"
)
