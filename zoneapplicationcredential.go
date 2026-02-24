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

// ZoneApplicationCredentialService contains methods and other services that help
// with interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneApplicationCredentialService] method instead.
type ZoneApplicationCredentialService struct {
	Options []option.RequestOption
}

// NewZoneApplicationCredentialService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneApplicationCredentialService(opts ...option.RequestOption) (r ZoneApplicationCredentialService) {
	r = ZoneApplicationCredentialService{}
	r.Options = opts
	return
}

// Creates a new application credential
func (r *ZoneApplicationCredentialService) New(ctx context.Context, zoneID string, body ZoneApplicationCredentialNewParams, opts ...option.RequestOption) (res *ZoneApplicationCredentialNewResponseUnion, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/application-credentials", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details of a specific application credential by ID
func (r *ZoneApplicationCredentialService) Get(ctx context.Context, id string, query ZoneApplicationCredentialGetParams, opts ...option.RequestOption) (res *CredentialUnion, err error) {
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
	path := fmt.Sprintf("zones/%s/application-credentials/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an application credential's configuration
func (r *ZoneApplicationCredentialService) Update(ctx context.Context, id string, params ZoneApplicationCredentialUpdateParams, opts ...option.RequestOption) (res *CredentialUnion, err error) {
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
	path := fmt.Sprintf("zones/%s/application-credentials/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of application credentials in the specified zone
func (r *ZoneApplicationCredentialService) List(ctx context.Context, zoneID string, query ZoneApplicationCredentialListParams, opts ...option.RequestOption) (res *ZoneApplicationCredentialListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/application-credentials", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes an application credential
func (r *ZoneApplicationCredentialService) Delete(ctx context.Context, id string, body ZoneApplicationCredentialDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("zones/%s/application-credentials/%s", url.PathEscape(body.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Common fields shared by all application credential types
type BaseFields struct {
	// Unique identifier of the credential
	ID string `json:"id,required"`
	// ID of the application this credential belongs to
	ApplicationID string `json:"application_id,required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Organization that owns this credential
	OrganizationID string `json:"organization_id,required"`
	// URL-safe identifier, unique within the zone
	Slug string `json:"slug,required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Zone this credential belongs to
	ZoneID string `json:"zone_id,required"`
	// An Application is a software system with an associated identity that can access
	// Resources. It may act on its own behalf (machine-to-machine) or on behalf of a
	// user (delegated access).
	//
	// Deprecated: deprecated
	Application Application `json:"application"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ApplicationID  respjson.Field
		CreatedAt      respjson.Field
		OrganizationID respjson.Field
		Slug           respjson.Field
		UpdatedAt      respjson.Field
		ZoneID         respjson.Field
		Application    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseFields) RawJSON() string { return r.JSON.raw }
func (r *BaseFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialUnion contains all possible properties and values from [Token],
// [Password], [PublicKey], [URL], [Public].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialUnion struct {
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	ID string `json:"id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	ApplicationID string `json:"application_id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	OrganizationID string `json:"organization_id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	Slug string `json:"slug"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	ZoneID string `json:"zone_id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	Application Application `json:"application"`
	Identifier  string      `json:"identifier"`
	// This field is from variant [Token].
	ProviderID string `json:"provider_id"`
	Type       string `json:"type"`
	// This field is from variant [Token].
	Provider Provider `json:"provider"`
	// This field is from variant [Token].
	Subject string `json:"subject"`
	// This field is from variant [Password].
	Password string `json:"password"`
	// This field is from variant [PublicKey].
	JwksUri string `json:"jwks_uri"`
	JSON    struct {
		ID             respjson.Field
		ApplicationID  respjson.Field
		CreatedAt      respjson.Field
		OrganizationID respjson.Field
		Slug           respjson.Field
		UpdatedAt      respjson.Field
		ZoneID         respjson.Field
		Application    respjson.Field
		Identifier     respjson.Field
		ProviderID     respjson.Field
		Type           respjson.Field
		Provider       respjson.Field
		Subject        respjson.Field
		Password       respjson.Field
		JwksUri        respjson.Field
		raw            string
	} `json:"-"`
}

func (u CredentialUnion) AsApplicationCredentialToken() (v Token) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialUnion) AsApplicationCredentialPassword() (v Password) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialUnion) AsApplicationCredentialPublicKey() (v PublicKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialUnion) AsApplicationCredentialURL() (v URL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialUnion) AsApplicationCredentialPublic() (v Public) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Password-based application credential
type Password struct {
	// Username for password credential, also used as OAuth 2.0 client ID
	Identifier string `json:"identifier,required"`
	// Any of "password".
	Type string `json:"type,required"`
	// Password for credential (only returned on creation, store securely), also used
	// as OAuth 2.0 client secret
	Password string `json:"password"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier  respjson.Field
		Type        respjson.Field
		Password    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseFields
}

// Returns the unmodified JSON received from the API
func (r Password) RawJSON() string { return r.JSON.raw }
func (r *Password) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public credential (no secret storage)
type Public struct {
	// Identifier for public credential, also used as OAuth 2.0 client ID
	Identifier string `json:"identifier,required"`
	// Any of "public".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseFields
}

// Returns the unmodified JSON received from the API
func (r Public) RawJSON() string { return r.JSON.raw }
func (r *Public) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public key-based application credential
type PublicKey struct {
	// Client ID for public key credential, also used as OAuth 2.0 client ID
	Identifier string `json:"identifier,required"`
	// JWKS URI to retrieve public keys from
	JwksUri string `json:"jwks_uri,required" format:"uri"`
	// Any of "public-key".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier  respjson.Field
		JwksUri     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseFields
}

// Returns the unmodified JSON received from the API
func (r PublicKey) RawJSON() string { return r.JSON.raw }
func (r *PublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token-based application credential
type Token struct {
	// Identifier for this credential. For token type, this equals the subject value,
	// or '\*' when subject is not specified.
	Identifier string `json:"identifier,required"`
	// ID of the provider issuing tokens verified by this credential
	ProviderID string `json:"provider_id,required"`
	// Any of "token".
	Type string `json:"type,required"`
	// A Provider is a system that supplies access to Resources and allows actors
	// (Users or Applications) to authenticate.
	//
	// Deprecated: deprecated
	Provider Provider `json:"provider"`
	// Subject identifier for the token. When null or omitted, any token from the
	// provider is accepted without checking application-specific claims.
	Subject string `json:"subject,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier  respjson.Field
		ProviderID  respjson.Field
		Type        respjson.Field
		Provider    respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseFields
}

// Returns the unmodified JSON received from the API
func (r Token) RawJSON() string { return r.JSON.raw }
func (r *Token) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL-based application credential
type URL struct {
	// URL of the credential (must be a valid URL)
	Identifier string `json:"identifier,required" format:"uri"`
	// Any of "url".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseFields
}

// Returns the unmodified JSON received from the API
func (r URL) RawJSON() string { return r.JSON.raw }
func (r *URL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ZoneApplicationCredentialNewResponseUnion contains all possible properties and
// values from [Token], [Password], [PublicKey], [URL], [Public].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ZoneApplicationCredentialNewResponseUnion struct {
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	ID string `json:"id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	ApplicationID string `json:"application_id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	OrganizationID string `json:"organization_id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	Slug string `json:"slug"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	ZoneID string `json:"zone_id"`
	// This field is from variant [Token], [Password], [PublicKey], [URL], [Public].
	Application Application `json:"application"`
	Identifier  string      `json:"identifier"`
	// This field is from variant [Token].
	ProviderID string `json:"provider_id"`
	Type       string `json:"type"`
	// This field is from variant [Token].
	Provider Provider `json:"provider"`
	// This field is from variant [Token].
	Subject string `json:"subject"`
	// This field is from variant [Password].
	Password string `json:"password"`
	// This field is from variant [PublicKey].
	JwksUri string `json:"jwks_uri"`
	JSON    struct {
		ID             respjson.Field
		ApplicationID  respjson.Field
		CreatedAt      respjson.Field
		OrganizationID respjson.Field
		Slug           respjson.Field
		UpdatedAt      respjson.Field
		ZoneID         respjson.Field
		Application    respjson.Field
		Identifier     respjson.Field
		ProviderID     respjson.Field
		Type           respjson.Field
		Provider       respjson.Field
		Subject        respjson.Field
		Password       respjson.Field
		JwksUri        respjson.Field
		raw            string
	} `json:"-"`
}

func (u ZoneApplicationCredentialNewResponseUnion) AsApplicationCredentialToken() (v Token) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZoneApplicationCredentialNewResponseUnion) AsApplicationCredentialPassword() (v Password) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZoneApplicationCredentialNewResponseUnion) AsApplicationCredentialPublicKey() (v PublicKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZoneApplicationCredentialNewResponseUnion) AsApplicationCredentialURL() (v URL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZoneApplicationCredentialNewResponseUnion) AsApplicationCredentialPublic() (v Public) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ZoneApplicationCredentialNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ZoneApplicationCredentialNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationCredentialListResponse struct {
	Items []CredentialUnion `json:"items,required"`
	// Pagination information
	PageInfo PageInfoPagination `json:"page_info,required"`
	// Cursor-based pagination metadata
	Pagination ZoneApplicationCredentialListResponsePagination `json:"pagination,required"`
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
func (r ZoneApplicationCredentialListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationCredentialListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneApplicationCredentialListResponsePagination struct {
	// An opaque cursor used for paginating through a list of results
	AfterCursor string `json:"after_cursor,required"`
	// An opaque cursor used for paginating through a list of results
	BeforeCursor string `json:"before_cursor,required"`
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
func (r ZoneApplicationCredentialListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationCredentialListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationCredentialNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Schema
	// for creating a token application credential
	OfApplicationCredentialCreateToken *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateToken `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for creating a password application credential
	OfApplicationCredentialCreatePassword *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePassword `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for creating a public key application credential
	OfApplicationCredentialCreatePublicKey *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublicKey `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for creating a URL application credential
	OfApplicationCredentialCreateURL *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for creating a public application credential
	OfApplicationCredentialCreatePublic *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublic `json:",inline"`

	paramObj
}

func (u ZoneApplicationCredentialNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfApplicationCredentialCreateToken,
		u.OfApplicationCredentialCreatePassword,
		u.OfApplicationCredentialCreatePublicKey,
		u.OfApplicationCredentialCreateURL,
		u.OfApplicationCredentialCreatePublic)
}
func (r *ZoneApplicationCredentialNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for creating a token application credential
//
// The properties ApplicationID, ProviderID, Type are required.
type ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateToken struct {
	// ID of the application this credential belongs to
	ApplicationID string `json:"application_id,required"`
	// ID of the provider issuing tokens this credential verifies
	ProviderID string `json:"provider_id,required"`
	// Any of "token".
	Type string `json:"type,omitzero,required"`
	// Subject identifier for the token. When omitted, any token from the provider is
	// accepted without checking application-specific claims.
	Subject param.Opt[string] `json:"subject,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateToken) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateToken](
		"type", "token",
	)
}

// Schema for creating a password application credential
//
// The properties ApplicationID, Type are required.
type ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePassword struct {
	// ID of the application this credential belongs to
	ApplicationID string `json:"application_id,required"`
	// Any of "password".
	Type string `json:"type,omitzero,required"`
	// Username for password credential, also used as OAuth 2.0 client ID
	// (auto-generated if not provided)
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePassword) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePassword
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePassword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePassword](
		"type", "password",
	)
}

// Schema for creating a public key application credential
//
// The properties ApplicationID, JwksUri, Type are required.
type ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublicKey struct {
	// ID of the application this credential belongs to
	ApplicationID string `json:"application_id,required"`
	// JWKS URI to retrieve public keys from
	JwksUri string `json:"jwks_uri,required" format:"uri"`
	// Any of "public-key".
	Type string `json:"type,omitzero,required"`
	// Client ID for public key credential, also used as OAuth 2.0 client ID
	// (auto-generated if not provided)
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublicKey) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublicKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublicKey](
		"type", "public-key",
	)
}

// Schema for creating a URL application credential
//
// The properties ApplicationID, Identifier, Type are required.
type ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateURL struct {
	// ID of the application this credential belongs to
	ApplicationID string `json:"application_id,required"`
	// URL of the credential (must be a valid URL)
	Identifier string `json:"identifier,required" format:"uri"`
	// Any of "url".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateURL) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateURL](
		"type", "url",
	)
}

// Schema for creating a public application credential
//
// The properties ApplicationID, Type are required.
type ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublic struct {
	// ID of the application this credential belongs to
	ApplicationID string `json:"application_id,required"`
	// Any of "public".
	Type string `json:"type,omitzero,required"`
	// Identifier for public credential, also used as OAuth 2.0 client ID
	// (auto-generated if not provided)
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublic) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreatePublic](
		"type", "public",
	)
}

type ZoneApplicationCredentialGetParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}

type ZoneApplicationCredentialUpdateParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Schema
	// for updating a token credential
	OfTokenCredentialUpdate *ZoneApplicationCredentialUpdateParamsBodyTokenCredentialUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for updating a password credential
	OfPasswordCredentialUpdate *ZoneApplicationCredentialUpdateParamsBodyPasswordCredentialUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for updating a public key credential
	OfPublicKeyCredentialUpdate *ZoneApplicationCredentialUpdateParamsBodyPublicKeyCredentialUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for updating a URL credential
	OfURLCredentialUpdate *ZoneApplicationCredentialUpdateParamsBodyURLCredentialUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Schema
	// for updating a public credential
	OfPublicCredentialUpdate *ZoneApplicationCredentialUpdateParamsBodyPublicCredentialUpdate `json:",inline"`

	paramObj
}

func (u ZoneApplicationCredentialUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTokenCredentialUpdate,
		u.OfPasswordCredentialUpdate,
		u.OfPublicKeyCredentialUpdate,
		u.OfURLCredentialUpdate,
		u.OfPublicCredentialUpdate)
}
func (r *ZoneApplicationCredentialUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for updating a token credential
type ZoneApplicationCredentialUpdateParamsBodyTokenCredentialUpdate struct {
	// Subject identifier for the token. Set to null to unset, which allows any token
	// from the provider to be accepted without checking application-specific claims.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Any of "token".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialUpdateParamsBodyTokenCredentialUpdate) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialUpdateParamsBodyTokenCredentialUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialUpdateParamsBodyTokenCredentialUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialUpdateParamsBodyTokenCredentialUpdate](
		"type", "token",
	)
}

// Schema for updating a password credential
type ZoneApplicationCredentialUpdateParamsBodyPasswordCredentialUpdate struct {
	// Any of "password".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialUpdateParamsBodyPasswordCredentialUpdate) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialUpdateParamsBodyPasswordCredentialUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialUpdateParamsBodyPasswordCredentialUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialUpdateParamsBodyPasswordCredentialUpdate](
		"type", "password",
	)
}

// Schema for updating a public key credential
type ZoneApplicationCredentialUpdateParamsBodyPublicKeyCredentialUpdate struct {
	// Any of "public-key".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialUpdateParamsBodyPublicKeyCredentialUpdate) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialUpdateParamsBodyPublicKeyCredentialUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialUpdateParamsBodyPublicKeyCredentialUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialUpdateParamsBodyPublicKeyCredentialUpdate](
		"type", "public-key",
	)
}

// Schema for updating a URL credential
type ZoneApplicationCredentialUpdateParamsBodyURLCredentialUpdate struct {
	// URL of the credential (must be a valid URL)
	Identifier param.Opt[string] `json:"identifier,omitzero" format:"uri"`
	// Any of "url".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialUpdateParamsBodyURLCredentialUpdate) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialUpdateParamsBodyURLCredentialUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialUpdateParamsBodyURLCredentialUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialUpdateParamsBodyURLCredentialUpdate](
		"type", "url",
	)
}

// Schema for updating a public credential
type ZoneApplicationCredentialUpdateParamsBodyPublicCredentialUpdate struct {
	// Identifier for public credential, also used as OAuth 2.0 client ID
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	// Any of "public".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ZoneApplicationCredentialUpdateParamsBodyPublicCredentialUpdate) MarshalJSON() (data []byte, err error) {
	type shadow ZoneApplicationCredentialUpdateParamsBodyPublicCredentialUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneApplicationCredentialUpdateParamsBodyPublicCredentialUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneApplicationCredentialUpdateParamsBodyPublicCredentialUpdate](
		"type", "public",
	)
}

type ZoneApplicationCredentialListParams struct {
	// Cursor for forward pagination
	After         param.Opt[string] `query:"after,omitzero" json:"-"`
	ApplicationID param.Opt[string] `query:"applicationId,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                               `query:"limit,omitzero" json:"-"`
	Slug   param.Opt[string]                              `query:"slug,omitzero" json:"-"`
	Expand ZoneApplicationCredentialListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationCredentialListParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationCredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneApplicationCredentialListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneApplicationCredentialListsExpandString)
	OfZoneApplicationCredentialListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneApplicationCredentialListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneApplicationCredentialListParamsExpandString string

const (
	ZoneApplicationCredentialListParamsExpandStringTotalCount ZoneApplicationCredentialListParamsExpandString = "total_count"
)

type ZoneApplicationCredentialDeleteParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}
