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

	"github.com/keycardlabs/keycard-go/internal/apijson"
	"github.com/keycardlabs/keycard-go/internal/apiquery"
	"github.com/keycardlabs/keycard-go/internal/requestconfig"
	"github.com/keycardlabs/keycard-go/option"
	"github.com/keycardlabs/keycard-go/packages/param"
	"github.com/keycardlabs/keycard-go/packages/respjson"
)

// ZoneProviderService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneProviderService] method instead.
type ZoneProviderService struct {
	Options []option.RequestOption
}

// NewZoneProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneProviderService(opts ...option.RequestOption) (r ZoneProviderService) {
	r = ZoneProviderService{}
	r.Options = opts
	return
}

// Creates a new Provider - a system that supplies access to Resources and allows
// actors to authenticate
func (r *ZoneProviderService) New(ctx context.Context, zoneID string, body ZoneProviderNewParams, opts ...option.RequestOption) (res *Provider, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/providers", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details of a specific Provider by ID
func (r *ZoneProviderService) Get(ctx context.Context, id string, query ZoneProviderGetParams, opts ...option.RequestOption) (res *Provider, err error) {
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
	path := fmt.Sprintf("zones/%s/providers/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Provider's configuration and metadata
func (r *ZoneProviderService) Update(ctx context.Context, id string, params ZoneProviderUpdateParams, opts ...option.RequestOption) (res *Provider, err error) {
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
	path := fmt.Sprintf("zones/%s/providers/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of providers in the specified zone
func (r *ZoneProviderService) List(ctx context.Context, zoneID string, query ZoneProviderListParams, opts ...option.RequestOption) (res *ZoneProviderListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/providers", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes a provider
func (r *ZoneProviderService) Delete(ctx context.Context, id string, body ZoneProviderDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("zones/%s/providers/%s", url.PathEscape(body.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A Provider is a system that supplies access to Resources and allows actors
// (Users or Applications) to authenticate.
type Provider struct {
	// Unique identifier of the provider
	ID string `json:"id,required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier,required"`
	// Human-readable name
	Name string `json:"name,required"`
	// Organization that owns this provider
	OrganizationID string `json:"organization_id,required"`
	// URL-safe identifier, unique within the zone
	Slug string `json:"slug,required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Zone this provider belongs to
	ZoneID string `json:"zone_id,required"`
	// OAuth 2.0 client identifier
	ClientID string `json:"client_id,nullable"`
	// Indicates whether a client secret is configured
	ClientSecretSet bool `json:"client_secret_set"`
	// Human-readable description
	Description string `json:"description,nullable"`
	// Provider metadata
	Metadata any `json:"metadata,nullable"`
	// Protocol-specific configuration
	Protocols ProviderProtocols `json:"protocols,nullable"`
	// Any of "external", "keycard-vault", "keycard-sts".
	Type ProviderType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Identifier      respjson.Field
		Name            respjson.Field
		OrganizationID  respjson.Field
		Slug            respjson.Field
		UpdatedAt       respjson.Field
		ZoneID          respjson.Field
		ClientID        respjson.Field
		ClientSecretSet respjson.Field
		Description     respjson.Field
		Metadata        respjson.Field
		Protocols       respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Provider) RawJSON() string { return r.JSON.raw }
func (r *Provider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol-specific configuration
type ProviderProtocols struct {
	// OAuth 2.0 protocol configuration
	Oauth2 ProviderProtocolsOauth2 `json:"oauth2,nullable"`
	// OpenID Connect protocol configuration
	Openid ProviderProtocolsOpenid `json:"openid,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Oauth2      respjson.Field
		Openid      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderProtocols) RawJSON() string { return r.JSON.raw }
func (r *ProviderProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration
type ProviderProtocolsOauth2 struct {
	AuthorizationEndpoint string `json:"authorization_endpoint,nullable" format:"uri"`
	// Whether to include the resource parameter in authorization requests.
	AuthorizationResourceEnabled bool `json:"authorization_resource_enabled,nullable"`
	// The resource parameter value to include in authorization requests. Defaults to
	// "resource" when authorization_resource_enabled is true.
	AuthorizationResourceParameter string   `json:"authorization_resource_parameter,nullable"`
	CodeChallengeMethodsSupported  []string `json:"code_challenge_methods_supported,nullable"`
	JwksUri                        string   `json:"jwks_uri,nullable" format:"uri"`
	RegistrationEndpoint           string   `json:"registration_endpoint,nullable" format:"uri"`
	// The query parameter name for scopes in authorization requests. Defaults to
	// "scope". Slack v2 uses "user_scope".
	ScopeParameter string `json:"scope_parameter,nullable"`
	// The separator character for scope values. Defaults to " " (space). Slack v2 uses
	// ",".
	ScopeSeparator  string   `json:"scope_separator,nullable"`
	ScopesSupported []string `json:"scopes_supported,nullable"`
	TokenEndpoint   string   `json:"token_endpoint,nullable" format:"uri"`
	// Dot-separated path to the access token in the token response body. Defaults to
	// "access_token". Slack v2 uses "authed_user.access_token".
	TokenResponseAccessTokenPointer string `json:"token_response_access_token_pointer,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationEndpoint           respjson.Field
		AuthorizationResourceEnabled    respjson.Field
		AuthorizationResourceParameter  respjson.Field
		CodeChallengeMethodsSupported   respjson.Field
		JwksUri                         respjson.Field
		RegistrationEndpoint            respjson.Field
		ScopeParameter                  respjson.Field
		ScopeSeparator                  respjson.Field
		ScopesSupported                 respjson.Field
		TokenEndpoint                   respjson.Field
		TokenResponseAccessTokenPointer respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderProtocolsOauth2) RawJSON() string { return r.JSON.raw }
func (r *ProviderProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenID Connect protocol configuration
type ProviderProtocolsOpenid struct {
	UserinfoEndpoint string `json:"userinfo_endpoint,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserinfoEndpoint respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderProtocolsOpenid) RawJSON() string { return r.JSON.raw }
func (r *ProviderProtocolsOpenid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderType string

const (
	ProviderTypeExternal     ProviderType = "external"
	ProviderTypeKeycardVault ProviderType = "keycard-vault"
	ProviderTypeKeycardSts   ProviderType = "keycard-sts"
)

type ZoneProviderListResponse struct {
	Items []Provider `json:"items,required"`
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
func (r ZoneProviderListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneProviderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneProviderNewParams struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier,required"`
	// Human-readable name
	Name string `json:"name,required"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// OAuth 2.0 client identifier
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// OAuth 2.0 client secret (will be encrypted and stored securely)
	ClientSecret param.Opt[string] `json:"client_secret,omitzero"`
	// Provider metadata
	Metadata any `json:"metadata,omitzero"`
	// Protocol-specific configuration for provider creation
	Protocols ZoneProviderNewParamsProtocols `json:"protocols,omitzero"`
	paramObj
}

func (r ZoneProviderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol-specific configuration for provider creation
type ZoneProviderNewParamsProtocols struct {
	// OAuth 2.0 protocol configuration for provider creation
	Oauth2 ZoneProviderNewParamsProtocolsOauth2 `json:"oauth2,omitzero"`
	// OpenID Connect protocol configuration for provider creation
	Openid ZoneProviderNewParamsProtocolsOpenid `json:"openid,omitzero"`
	paramObj
}

func (r ZoneProviderNewParamsProtocols) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderNewParamsProtocols
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderNewParamsProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration for provider creation
type ZoneProviderNewParamsProtocolsOauth2 struct {
	AuthorizationEndpoint param.Opt[string] `json:"authorization_endpoint,omitzero" format:"uri"`
	// Whether to include the resource parameter in authorization requests.
	AuthorizationResourceEnabled param.Opt[bool] `json:"authorization_resource_enabled,omitzero"`
	// The resource parameter value to include in authorization requests. Defaults to
	// "resource" when authorization_resource_enabled is true.
	AuthorizationResourceParameter param.Opt[string] `json:"authorization_resource_parameter,omitzero"`
	JwksUri                        param.Opt[string] `json:"jwks_uri,omitzero" format:"uri"`
	RegistrationEndpoint           param.Opt[string] `json:"registration_endpoint,omitzero" format:"uri"`
	// The query parameter name for scopes in authorization requests. Defaults to
	// "scope". Slack v2 uses "user_scope".
	ScopeParameter param.Opt[string] `json:"scope_parameter,omitzero"`
	// The separator character for scope values. Defaults to " " (space). Slack v2 uses
	// ",".
	ScopeSeparator param.Opt[string] `json:"scope_separator,omitzero"`
	TokenEndpoint  param.Opt[string] `json:"token_endpoint,omitzero" format:"uri"`
	// Dot-separated path to the access token in the token response body. Defaults to
	// "access_token". Slack v2 uses "authed_user.access_token".
	TokenResponseAccessTokenPointer param.Opt[string] `json:"token_response_access_token_pointer,omitzero"`
	CodeChallengeMethodsSupported   []string          `json:"code_challenge_methods_supported,omitzero"`
	ScopesSupported                 []string          `json:"scopes_supported,omitzero"`
	paramObj
}

func (r ZoneProviderNewParamsProtocolsOauth2) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderNewParamsProtocolsOauth2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderNewParamsProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenID Connect protocol configuration for provider creation
type ZoneProviderNewParamsProtocolsOpenid struct {
	UserinfoEndpoint param.Opt[string] `json:"userinfo_endpoint,omitzero" format:"uri"`
	paramObj
}

func (r ZoneProviderNewParamsProtocolsOpenid) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderNewParamsProtocolsOpenid
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderNewParamsProtocolsOpenid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneProviderGetParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}

type ZoneProviderUpdateParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	// OAuth 2.0 client identifier. Set to null to remove.
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// OAuth 2.0 client secret (will be encrypted and stored securely). Set to null to
	// remove.
	ClientSecret param.Opt[string] `json:"client_secret,omitzero"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// User specified identifier, unique within the zone
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	// Human-readable name
	Name param.Opt[string] `json:"name,omitzero"`
	// Provider metadata. Set to null to remove all metadata.
	Metadata any `json:"metadata,omitzero"`
	// Protocol-specific configuration. Set to null to remove all protocols.
	Protocols ZoneProviderUpdateParamsProtocols `json:"protocols,omitzero"`
	paramObj
}

func (r ZoneProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol-specific configuration. Set to null to remove all protocols.
type ZoneProviderUpdateParamsProtocols struct {
	// OAuth 2.0 protocol configuration. Set to null to remove all OAuth2 config.
	Oauth2 ZoneProviderUpdateParamsProtocolsOauth2 `json:"oauth2,omitzero"`
	// OpenID Connect protocol configuration. Set to null to remove all OpenID config.
	Openid ZoneProviderUpdateParamsProtocolsOpenid `json:"openid,omitzero"`
	paramObj
}

func (r ZoneProviderUpdateParamsProtocols) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderUpdateParamsProtocols
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderUpdateParamsProtocols) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration. Set to null to remove all OAuth2 config.
type ZoneProviderUpdateParamsProtocolsOauth2 struct {
	AuthorizationEndpoint param.Opt[string] `json:"authorization_endpoint,omitzero" format:"uri"`
	// Whether to include the resource parameter in authorization requests. Set to null
	// to unset.
	AuthorizationResourceEnabled param.Opt[bool] `json:"authorization_resource_enabled,omitzero"`
	// The resource parameter value to include in authorization requests. Defaults to
	// "resource" when authorization_resource_enabled is true. Set to null to unset.
	AuthorizationResourceParameter param.Opt[string] `json:"authorization_resource_parameter,omitzero"`
	JwksUri                        param.Opt[string] `json:"jwks_uri,omitzero" format:"uri"`
	RegistrationEndpoint           param.Opt[string] `json:"registration_endpoint,omitzero" format:"uri"`
	// The query parameter name for scopes in authorization requests. Defaults to
	// "scope". Set to null to unset.
	ScopeParameter param.Opt[string] `json:"scope_parameter,omitzero"`
	// The separator character for scope values. Defaults to " " (space). Set to null
	// to unset.
	ScopeSeparator param.Opt[string] `json:"scope_separator,omitzero"`
	TokenEndpoint  param.Opt[string] `json:"token_endpoint,omitzero" format:"uri"`
	// Dot-separated path to the access token in the token response body. Defaults to
	// "access_token". Set to null to unset.
	TokenResponseAccessTokenPointer param.Opt[string] `json:"token_response_access_token_pointer,omitzero"`
	CodeChallengeMethodsSupported   []string          `json:"code_challenge_methods_supported,omitzero"`
	ScopesSupported                 []string          `json:"scopes_supported,omitzero"`
	paramObj
}

func (r ZoneProviderUpdateParamsProtocolsOauth2) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderUpdateParamsProtocolsOauth2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderUpdateParamsProtocolsOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenID Connect protocol configuration. Set to null to remove all OpenID config.
type ZoneProviderUpdateParamsProtocolsOpenid struct {
	UserinfoEndpoint param.Opt[string] `json:"userinfo_endpoint,omitzero" format:"uri"`
	paramObj
}

func (r ZoneProviderUpdateParamsProtocolsOpenid) MarshalJSON() (data []byte, err error) {
	type shadow ZoneProviderUpdateParamsProtocolsOpenid
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneProviderUpdateParamsProtocolsOpenid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneProviderListParams struct {
	Cursor     param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Identifier param.Opt[string] `query:"identifier,omitzero" json:"-"`
	Limit      param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Slug       param.Opt[string] `query:"slug,omitzero" json:"-"`
	// Any of "external", "keycard-vault", "keycard-sts".
	Type ZoneProviderListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneProviderListParams]'s query parameters as `url.Values`.
func (r ZoneProviderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneProviderListParamsType string

const (
	ZoneProviderListParamsTypeExternal     ZoneProviderListParamsType = "external"
	ZoneProviderListParamsTypeKeycardVault ZoneProviderListParamsType = "keycard-vault"
	ZoneProviderListParamsTypeKeycardSts   ZoneProviderListParamsType = "keycard-sts"
)

type ZoneProviderDeleteParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}
