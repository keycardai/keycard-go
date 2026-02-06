// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi

import (
	"context"
	"encoding/json"
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
	McpGateways            ZoneMcpGatewayService
	Providers              ZoneProviderService
	Resources              ZoneResourceService
	Sessions               ZoneSessionService
	UserAgents             ZoneUserAgentService
	Users                  ZoneUserService
	Members                ZoneMemberService
	Secrets                ZoneSecretService
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
	r.McpGateways = NewZoneMcpGatewayService(opts...)
	r.Providers = NewZoneProviderService(opts...)
	r.Resources = NewZoneResourceService(opts...)
	r.Sessions = NewZoneSessionService(opts...)
	r.UserAgents = NewZoneUserAgentService(opts...)
	r.Users = NewZoneUserService(opts...)
	r.Members = NewZoneMemberService(opts...)
	r.Secrets = NewZoneSecretService(opts...)
	return
}

// Creates a new zone for the authenticated organization. A zone is an isolated
// environment for IAM resources.
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *Zone, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details of a specific zone by ID
func (r *ZoneService) Get(ctx context.Context, zoneID string, query ZoneGetParams, opts ...option.RequestOption) (res *Zone, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates a zone's configuration (partial update)
func (r *ZoneService) Update(ctx context.Context, zoneID string, body ZoneUpdateParams, opts ...option.RequestOption) (res *Zone, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of zones for the authenticated organization
func (r *ZoneService) List(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) (res *ZoneListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes a zone and all its associated resources
func (r *ZoneService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Removes downstream resource, dependency, and optionally upstream
// resource/provider
func (r *ZoneService) DeleteMcpServer(ctx context.Context, downstreamID string, body ZoneDeleteMcpServerParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if downstreamID == "" {
		err = errors.New("missing required downstreamId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/mcp-servers/%s", body.ZoneID, downstreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns aggregated access records per session-resource pair. At least one of
// user_id, session_id, or resource_id must be provided. By default when filtering
// by user_id, returns sessions with an initiator (application or user agent). Use
// has_initiator=true to explicitly filter to sessions with an initiator.
func (r *ZoneService) ListSessionResourceAccess(ctx context.Context, zoneID string, query ZoneListSessionResourceAccessParams, opts ...option.RequestOption) (res *ZoneListSessionResourceAccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/session-resource-access", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// AWS KMS configuration for zone encryption. When not specified, the default
// Keycard Cloud encryption key will be used.
type EncryptionKeyAwsKmsConfig struct {
	// AWS KMS Key ARN for encrypting the zone's data
	Arn string `json:"arn,required"`
	// Any of "aws".
	Type EncryptionKeyAwsKmsConfigType `json:"type,required"`
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
	Arn string `json:"arn,required"`
	// Any of "aws".
	Type EncryptionKeyAwsKmsConfigType `json:"type,omitzero,required"`
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
	HasNextPage bool `json:"has_next_page,required"`
	// Whether there are items before the current page
	HasPreviousPage bool `json:"has_previous_page,required"`
	// Cursor pointing to the last item in the current page
	EndCursor string `json:"end_cursor,nullable"`
	// Cursor pointing to the first item in the current page
	StartCursor string `json:"start_cursor,nullable"`
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
	ID string `json:"id,required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name
	Name string `json:"name,required"`
	// Organization that owns this zone
	OrganizationID string `json:"organization_id,required"`
	// Protocol configuration for a zone
	Protocols ZoneProtocols `json:"protocols,required"`
	// URL-safe identifier, unique within the zone
	Slug string `json:"slug,required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Application ID configured as the default MCP Gateway for the zone
	DefaultMcpGatewayApplicationID string `json:"default_mcp_gateway_application_id"`
	// Human-readable description
	Description string `json:"description,nullable"`
	// Whether directory open signups are enabled for the zone, only applies when
	// user_identity_provider_id is not set
	DirectoryOpenSignupsEnabled bool `json:"directory_open_signups_enabled"`
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
		Description                    respjson.Field
		DirectoryOpenSignupsEnabled    respjson.Field
		EncryptionKey                  respjson.Field
		LoginFlow                      respjson.Field
		Permissions                    respjson.Field
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
	Oauth2 ZoneProtocolsOauth2 `json:"oauth2,required"`
	// OpenID Connect protocol configuration for a zone
	Openid ZoneProtocolsOpenid `json:"openid,required"`
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
	AuthorizationEndpoint string `json:"authorization_endpoint,required" format:"uri"`
	// OAuth 2.0 Authorization Server Metadata endpoint
	// (.well-known/oauth-authorization-server)
	AuthorizationServerMetadata string `json:"authorization_server_metadata,required" format:"uri"`
	// Whether Dynamic Client Registration is enabled
	DcrEnabled bool `json:"dcr_enabled,required"`
	// OAuth 2.0 issuer identifier
	Issuer string `json:"issuer,required" format:"uri"`
	// JSON Web Key Set endpoint
	JwksUri string `json:"jwks_uri,required" format:"uri"`
	// Whether PKCE is required for authorization code flows
	PkceRequired bool `json:"pkce_required,required"`
	// OAuth 2.0 redirect URI for this zone
	RedirectUri string `json:"redirect_uri,required" format:"uri"`
	// OAuth 2.0 Dynamic Client Registration endpoint
	RegistrationEndpoint string `json:"registration_endpoint,required" format:"uri"`
	// OAuth 2.0 token endpoint
	TokenEndpoint string `json:"token_endpoint,required" format:"uri"`
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
	ProviderConfiguration string `json:"provider_configuration,required" format:"uri"`
	// OpenID Connect UserInfo endpoint
	UserinfoEndpoint string `json:"userinfo_endpoint,required" format:"uri"`
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
	Items []Zone `json:"items,required"`
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
func (r ZoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListSessionResourceAccessResponse struct {
	Items []ZoneListSessionResourceAccessResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListSessionResourceAccessResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneListSessionResourceAccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregated record of session-resource access events
type ZoneListSessionResourceAccessResponseItem struct {
	// When access first occurred
	FirstAccessedAt time.Time `json:"first_accessed_at,required" format:"date-time"`
	// When access most recently occurred
	LastAccessedAt time.Time `json:"last_accessed_at,required" format:"date-time"`
	// Organization ID
	OrganizationID string `json:"organization_id,required"`
	// Resource ID
	ResourceID string `json:"resource_id,required"`
	// Session ID
	SessionID string `json:"session_id,required"`
	// Total number of access events for this session-resource pair
	TotalAccessCount int64 `json:"total_access_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstAccessedAt  respjson.Field
		LastAccessedAt   respjson.Field
		OrganizationID   respjson.Field
		ResourceID       respjson.Field
		SessionID        respjson.Field
		TotalAccessCount respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListSessionResourceAccessResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ZoneListSessionResourceAccessResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewParams struct {
	// Human-readable name
	Name string `json:"name,required"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// Assign a default MCP Gateway application to the zone
	DefaultMcpGatewayApplication param.Opt[bool] `json:"default_mcp_gateway_application,omitzero"`
	// Whether directory open signups are enabled for the zone, only applies when
	// user_identity_provider_id is not set
	DirectoryOpenSignupsEnabled param.Opt[bool] `json:"directory_open_signups_enabled,omitzero"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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

func (u *ZoneGetParamsExpandUnion) asAny() any {
	if !param.IsOmitted(u.OfZoneGetsExpandString) {
		return &u.OfZoneGetsExpandString
	} else if !param.IsOmitted(u.OfZoneGetsExpandArrayItemArray) {
		return &u.OfZoneGetsExpandArrayItemArray
	}
	return nil
}

type ZoneGetParamsExpandString string

const (
	ZoneGetParamsExpandStringPermissions ZoneGetParamsExpandString = "permissions"
)

type ZoneUpdateParams struct {
	// Application ID configured as the default MCP Gateway for the zone (set to null
	// to unset)
	DefaultMcpGatewayApplicationID param.Opt[string] `json:"default_mcp_gateway_application_id,omitzero"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// Provider ID to configure for user login (set to null to unset)
	UserIdentityProviderID param.Opt[string] `json:"user_identity_provider_id,omitzero"`
	// Whether directory open signups are enabled for the zone, only applies when
	// user_identity_provider_id is not set
	DirectoryOpenSignupsEnabled param.Opt[bool] `json:"directory_open_signups_enabled,omitzero"`
	// Human-readable name
	Name param.Opt[string] `json:"name,omitzero"`
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
	Arn string `json:"arn,required"`
	// Any of "aws".
	Type string `json:"type,omitzero,required"`
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
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Slug   param.Opt[string] `query:"slug,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneListParams]'s query parameters as `url.Values`.
func (r ZoneListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneDeleteMcpServerParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}

type ZoneListSessionResourceAccessParams struct {
	// Filter by resource ID
	ResourceID param.Opt[string] `query:"resource_id,omitzero" json:"-"`
	// Filter by session ID
	SessionID param.Opt[string] `query:"session_id,omitzero" json:"-"`
	// Filter by user ID
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Filter sessions that have an initiator (application_id OR user_agent_id is set).
	//
	// Any of "true".
	HasInitiator ZoneListSessionResourceAccessParamsHasInitiator `query:"has_initiator,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneListSessionResourceAccessParams]'s query parameters as
// `url.Values`.
func (r ZoneListSessionResourceAccessParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter sessions that have an initiator (application_id OR user_agent_id is set).
type ZoneListSessionResourceAccessParamsHasInitiator string

const (
	ZoneListSessionResourceAccessParamsHasInitiatorTrue ZoneListSessionResourceAccessParamsHasInitiator = "true"
)
