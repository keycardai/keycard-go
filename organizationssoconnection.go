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

// OrganizationSSOConnectionService contains methods and other services that help
// with interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationSSOConnectionService] method instead.
type OrganizationSSOConnectionService struct {
	Options []option.RequestOption
}

// NewOrganizationSSOConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationSSOConnectionService(opts ...option.RequestOption) (r OrganizationSSOConnectionService) {
	r = OrganizationSSOConnectionService{}
	r.Options = opts
	return
}

// Get SSO connection configuration for organization
func (r *OrganizationSSOConnectionService) Get(ctx context.Context, organizationID string, params OrganizationSSOConnectionGetParams, opts ...option.RequestOption) (res *SSOConnection, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/sso-connection", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update SSO connection configuration
func (r *OrganizationSSOConnectionService) Update(ctx context.Context, organizationID string, params OrganizationSSOConnectionUpdateParams, opts ...option.RequestOption) (res *SSOConnection, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/sso-connection", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Disable SSO for organization
func (r *OrganizationSSOConnectionService) Disable(ctx context.Context, organizationID string, body OrganizationSSOConnectionDisableParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", body.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/sso-connection", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Enable SSO for organization
func (r *OrganizationSSOConnectionService) Enable(ctx context.Context, organizationID string, params OrganizationSSOConnectionEnableParams, opts ...option.RequestOption) (res *SSOConnection, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/sso-connection", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// SSO connection configuration for an organization
type SSOConnection struct {
	// Unique identifier for the SSO connection
	ID string `json:"id,required"`
	// OAuth 2.0 client ID
	ClientID string `json:"client_id,required"`
	// Whether a client secret is configured
	ClientSecretSet bool `json:"client_secret_set,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// SSO provider identifier (e.g., issuer URL)
	Identifier string `json:"identifier,required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// Protocol configuration for SSO connection
	Protocols SSOConnectionProtocol `json:"protocols,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ClientID        respjson.Field
		ClientSecretSet respjson.Field
		CreatedAt       respjson.Field
		Identifier      respjson.Field
		UpdatedAt       respjson.Field
		Permissions     respjson.Field
		Protocols       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSOConnection) RawJSON() string { return r.JSON.raw }
func (r *SSOConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol configuration for SSO connection
type SSOConnectionProtocol struct {
	// OAuth 2.0 protocol configuration for SSO connection
	Oauth2 SSOConnectionProtocolOauth2 `json:"oauth2,nullable"`
	// OpenID Connect protocol configuration for SSO connection
	Openid SSOConnectionProtocolOpenid `json:"openid,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Oauth2      respjson.Field
		Openid      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSOConnectionProtocol) RawJSON() string { return r.JSON.raw }
func (r *SSOConnectionProtocol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SSOConnectionProtocol to a SSOConnectionProtocolParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SSOConnectionProtocolParam.Overrides()
func (r SSOConnectionProtocol) ToParam() SSOConnectionProtocolParam {
	return param.Override[SSOConnectionProtocolParam](json.RawMessage(r.RawJSON()))
}

// OAuth 2.0 protocol configuration for SSO connection
type SSOConnectionProtocolOauth2 struct {
	// OAuth 2.0 authorization endpoint
	AuthorizationEndpoint string `json:"authorization_endpoint,nullable" format:"uri"`
	// Supported PKCE code challenge methods
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported,nullable"`
	// JSON Web Key Set endpoint
	JwksUri string `json:"jwks_uri,nullable" format:"uri"`
	// OAuth 2.0 registration endpoint
	RegistrationEndpoint string `json:"registration_endpoint,nullable" format:"uri"`
	// Supported OAuth 2.0 scopes
	ScopesSupported []string `json:"scopes_supported,nullable"`
	// OAuth 2.0 token endpoint
	TokenEndpoint string `json:"token_endpoint,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationEndpoint         respjson.Field
		CodeChallengeMethodsSupported respjson.Field
		JwksUri                       respjson.Field
		RegistrationEndpoint          respjson.Field
		ScopesSupported               respjson.Field
		TokenEndpoint                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSOConnectionProtocolOauth2) RawJSON() string { return r.JSON.raw }
func (r *SSOConnectionProtocolOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenID Connect protocol configuration for SSO connection
type SSOConnectionProtocolOpenid struct {
	// OpenID Connect UserInfo endpoint
	UserinfoEndpoint string `json:"userinfo_endpoint,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserinfoEndpoint respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSOConnectionProtocolOpenid) RawJSON() string { return r.JSON.raw }
func (r *SSOConnectionProtocolOpenid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol configuration for SSO connection
type SSOConnectionProtocolParam struct {
	// OAuth 2.0 protocol configuration for SSO connection
	Oauth2 SSOConnectionProtocolOauth2Param `json:"oauth2,omitzero"`
	// OpenID Connect protocol configuration for SSO connection
	Openid SSOConnectionProtocolOpenidParam `json:"openid,omitzero"`
	paramObj
}

func (r SSOConnectionProtocolParam) MarshalJSON() (data []byte, err error) {
	type shadow SSOConnectionProtocolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSOConnectionProtocolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 protocol configuration for SSO connection
type SSOConnectionProtocolOauth2Param struct {
	// OAuth 2.0 authorization endpoint
	AuthorizationEndpoint param.Opt[string] `json:"authorization_endpoint,omitzero" format:"uri"`
	// JSON Web Key Set endpoint
	JwksUri param.Opt[string] `json:"jwks_uri,omitzero" format:"uri"`
	// OAuth 2.0 registration endpoint
	RegistrationEndpoint param.Opt[string] `json:"registration_endpoint,omitzero" format:"uri"`
	// OAuth 2.0 token endpoint
	TokenEndpoint param.Opt[string] `json:"token_endpoint,omitzero" format:"uri"`
	// Supported PKCE code challenge methods
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported,omitzero"`
	// Supported OAuth 2.0 scopes
	ScopesSupported []string `json:"scopes_supported,omitzero"`
	paramObj
}

func (r SSOConnectionProtocolOauth2Param) MarshalJSON() (data []byte, err error) {
	type shadow SSOConnectionProtocolOauth2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSOConnectionProtocolOauth2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenID Connect protocol configuration for SSO connection
type SSOConnectionProtocolOpenidParam struct {
	// OpenID Connect UserInfo endpoint
	UserinfoEndpoint param.Opt[string] `json:"userinfo_endpoint,omitzero" format:"uri"`
	paramObj
}

func (r SSOConnectionProtocolOpenidParam) MarshalJSON() (data []byte, err error) {
	type shadow SSOConnectionProtocolOpenidParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSOConnectionProtocolOpenidParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationSSOConnectionGetParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationSSOConnectionGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationSSOConnectionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationSSOConnectionUpdateParams struct {
	// OAuth 2.0 client ID (set to null to remove)
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// OAuth 2.0 client secret (set to null to remove)
	ClientSecret param.Opt[string] `json:"client_secret,omitzero"`
	// SSO provider identifier (e.g., issuer URL)
	Identifier       param.Opt[string] `json:"identifier,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Protocol configuration for SSO connection
	Protocols SSOConnectionProtocolParam `json:"protocols,omitzero"`
	paramObj
}

func (r OrganizationSSOConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationSSOConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationSSOConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationSSOConnectionDisableParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type OrganizationSSOConnectionEnableParams struct {
	// OAuth 2.0 client ID
	ClientID string `json:"client_id,required"`
	// SSO provider identifier (e.g., issuer URL)
	Identifier string `json:"identifier,required"`
	// OAuth 2.0 client secret (optional, will be encrypted if provided)
	ClientSecret     param.Opt[string] `json:"client_secret,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Protocol configuration for SSO connection
	Protocols SSOConnectionProtocolParam `json:"protocols,omitzero"`
	paramObj
}

func (r OrganizationSSOConnectionEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationSSOConnectionEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationSSOConnectionEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
