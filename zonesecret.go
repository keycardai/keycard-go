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

// ZoneSecretService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSecretService] method instead.
type ZoneSecretService struct {
	Options []option.RequestOption
}

// NewZoneSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSecretService(opts ...option.RequestOption) (r ZoneSecretService) {
	r = ZoneSecretService{}
	r.Options = opts
	return
}

func (r *ZoneSecretService) New(ctx context.Context, zoneID string, params ZoneSecretNewParams, opts ...option.RequestOption) (res *Secret, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithVaultAPIBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *ZoneSecretService) Get(ctx context.Context, id string, params ZoneSecretGetParams, opts ...option.RequestOption) (res *ZoneSecretGetResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithVaultAPIBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ZoneSecretService) Update(ctx context.Context, id string, params ZoneSecretUpdateParams, opts ...option.RequestOption) (res *Secret, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithVaultAPIBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

func (r *ZoneSecretService) List(ctx context.Context, zoneID string, params ZoneSecretListParams, opts ...option.RequestOption) (res *[]Secret, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithVaultAPIBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

func (r *ZoneSecretService) Delete(ctx context.Context, id string, params ZoneSecretDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithVaultAPIBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Secret struct {
	// A globally unique opaque identifier
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id" api:"required"`
	// A name for the entity to be displayed in UI
	Name string `json:"name" api:"required"`
	// Any of "token", "password".
	Type      SecretType `json:"type" api:"required"`
	UpdatedAt time.Time  `json:"updated_at" api:"required" format:"date-time"`
	Version   int64      `json:"version" api:"required"`
	// A globally unique opaque identifier
	ZoneID string `json:"zone_id" api:"required"`
	// A description of the entity
	Description string `json:"description"`
	// A JSON object containing arbitrary metadata. Metadata will not be encrypted.
	Metadata any `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EntityID    respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Version     respjson.Field
		ZoneID      respjson.Field
		Description respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Secret) RawJSON() string { return r.JSON.raw }
func (r *Secret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretType string

const (
	SecretTypeToken    SecretType = "token"
	SecretTypePassword SecretType = "password"
)

type SecretPasswordFields struct {
	Password string `json:"password" api:"required"`
	// Any of "password".
	Type     SecretPasswordFieldsType `json:"type" api:"required"`
	Username string                   `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Password    respjson.Field
		Type        respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecretPasswordFields) RawJSON() string { return r.JSON.raw }
func (r *SecretPasswordFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SecretPasswordFields to a SecretPasswordFieldsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SecretPasswordFieldsParam.Overrides()
func (r SecretPasswordFields) ToParam() SecretPasswordFieldsParam {
	return param.Override[SecretPasswordFieldsParam](json.RawMessage(r.RawJSON()))
}

type SecretPasswordFieldsType string

const (
	SecretPasswordFieldsTypePassword SecretPasswordFieldsType = "password"
)

// The properties Password, Type, Username are required.
type SecretPasswordFieldsParam struct {
	Password string `json:"password" api:"required"`
	// Any of "password".
	Type     SecretPasswordFieldsType `json:"type,omitzero" api:"required"`
	Username string                   `json:"username" api:"required"`
	paramObj
}

func (r SecretPasswordFieldsParam) MarshalJSON() (data []byte, err error) {
	type shadow SecretPasswordFieldsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecretPasswordFieldsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretTokenFields struct {
	Token string `json:"token" api:"required"`
	// Any of "token".
	Type SecretTokenFieldsType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecretTokenFields) RawJSON() string { return r.JSON.raw }
func (r *SecretTokenFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SecretTokenFields to a SecretTokenFieldsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SecretTokenFieldsParam.Overrides()
func (r SecretTokenFields) ToParam() SecretTokenFieldsParam {
	return param.Override[SecretTokenFieldsParam](json.RawMessage(r.RawJSON()))
}

type SecretTokenFieldsType string

const (
	SecretTokenFieldsTypeToken SecretTokenFieldsType = "token"
)

// The properties Token, Type are required.
type SecretTokenFieldsParam struct {
	Token string `json:"token" api:"required"`
	// Any of "token".
	Type SecretTokenFieldsType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r SecretTokenFieldsParam) MarshalJSON() (data []byte, err error) {
	type shadow SecretTokenFieldsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecretTokenFieldsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretGetResponse struct {
	// A globally unique opaque identifier
	ID        string                         `json:"id" api:"required"`
	CreatedAt time.Time                      `json:"created_at" api:"required" format:"date-time"`
	Data      ZoneSecretGetResponseDataUnion `json:"data" api:"required"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id" api:"required"`
	// A name for the entity to be displayed in UI
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	Version   int64     `json:"version" api:"required"`
	// A globally unique opaque identifier
	ZoneID string `json:"zone_id" api:"required"`
	// A description of the entity
	Description string `json:"description"`
	// A JSON object containing arbitrary metadata. Metadata will not be encrypted.
	Metadata any `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Data        respjson.Field
		EntityID    respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		Version     respjson.Field
		ZoneID      respjson.Field
		Description respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneSecretGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneSecretGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ZoneSecretGetResponseDataUnion contains all possible properties and values from
// [SecretTokenFields], [SecretPasswordFields].
//
// Use the [ZoneSecretGetResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ZoneSecretGetResponseDataUnion struct {
	// This field is from variant [SecretTokenFields].
	Token string `json:"token"`
	// Any of "token", "password".
	Type string `json:"type"`
	// This field is from variant [SecretPasswordFields].
	Password string `json:"password"`
	// This field is from variant [SecretPasswordFields].
	Username string `json:"username"`
	JSON     struct {
		Token    respjson.Field
		Type     respjson.Field
		Password respjson.Field
		Username respjson.Field
		raw      string
	} `json:"-"`
}

// anyZoneSecretGetResponseData is implemented by each variant of
// [ZoneSecretGetResponseDataUnion] to add type safety for the return type of
// [ZoneSecretGetResponseDataUnion.AsAny]
type anyZoneSecretGetResponseData interface {
	implZoneSecretGetResponseDataUnion()
}

func (SecretTokenFields) implZoneSecretGetResponseDataUnion()    {}
func (SecretPasswordFields) implZoneSecretGetResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ZoneSecretGetResponseDataUnion.AsAny().(type) {
//	case keycard.SecretTokenFields:
//	case keycard.SecretPasswordFields:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ZoneSecretGetResponseDataUnion) AsAny() anyZoneSecretGetResponseData {
	switch u.Type {
	case "token":
		return u.AsToken()
	case "password":
		return u.AsPassword()
	}
	return nil
}

func (u ZoneSecretGetResponseDataUnion) AsToken() (v SecretTokenFields) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZoneSecretGetResponseDataUnion) AsPassword() (v SecretPasswordFields) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ZoneSecretGetResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ZoneSecretGetResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretNewParams struct {
	Data ZoneSecretNewParamsDataUnion `json:"data,omitzero" api:"required"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id" api:"required"`
	// A name for the entity to be displayed in UI
	Name string `json:"name" api:"required"`
	// A description of the entity
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional zone ID. This field is provided for API compatibility but is ignored
	// during processing. The zone ID is derived from the path parameter
	// (/zones/{zone_id}/secrets) and takes precedence.
	ZoneID           param.Opt[string] `json:"zone_id,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// A JSON object containing arbitrary metadata. Metadata will not be encrypted.
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r ZoneSecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneSecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneSecretNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneSecretNewParamsDataUnion struct {
	OfToken    *SecretTokenFieldsParam    `json:",omitzero,inline"`
	OfPassword *SecretPasswordFieldsParam `json:",omitzero,inline"`
	paramUnion
}

func (u ZoneSecretNewParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfToken, u.OfPassword)
}
func (u *ZoneSecretNewParamsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ZoneSecretNewParamsDataUnion](
		"type",
		apijson.Discriminator[SecretTokenFieldsParam]("token"),
		apijson.Discriminator[SecretPasswordFieldsParam]("password"),
	)
}

type ZoneSecretGetParams struct {
	// A globally unique opaque identifier
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZoneSecretUpdateParams struct {
	// A globally unique opaque identifier
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// A description of the entity
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the entity to be displayed in UI
	Name             param.Opt[string]               `json:"name,omitzero"`
	XClientRequestID param.Opt[string]               `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	Data             ZoneSecretUpdateParamsDataUnion `json:"data,omitzero"`
	// A JSON object containing arbitrary metadata. Metadata will not be encrypted.
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r ZoneSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneSecretUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneSecretUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneSecretUpdateParamsDataUnion struct {
	OfToken    *SecretTokenFieldsParam    `json:",omitzero,inline"`
	OfPassword *SecretPasswordFieldsParam `json:",omitzero,inline"`
	paramUnion
}

func (u ZoneSecretUpdateParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfToken, u.OfPassword)
}
func (u *ZoneSecretUpdateParamsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ZoneSecretUpdateParamsDataUnion](
		"type",
		apijson.Discriminator[SecretTokenFieldsParam]("token"),
		apijson.Discriminator[SecretPasswordFieldsParam]("password"),
	)
}

type ZoneSecretListParams struct {
	// The entity to list all secrets for
	EntityID         param.Opt[string] `query:"entity_id,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// The type of secrets to list
	//
	// Any of "token", "password".
	Type ZoneSecretListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneSecretListParams]'s query parameters as `url.Values`.
func (r ZoneSecretListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of secrets to list
type ZoneSecretListParamsType string

const (
	ZoneSecretListParamsTypeToken    ZoneSecretListParamsType = "token"
	ZoneSecretListParamsTypePassword ZoneSecretListParamsType = "password"
)

type ZoneSecretDeleteParams struct {
	// A globally unique opaque identifier
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
