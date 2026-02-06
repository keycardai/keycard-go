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
	"github.com/stainless-sdks/keycard-api-go/shared/constant"
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

func (r *ZoneSecretService) New(ctx context.Context, zoneID string, params ZoneSecretNewParams, opts ...option.RequestOption) (res *ZoneSecretNewResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *ZoneSecretService) Get(ctx context.Context, id string, params ZoneSecretGetParams, opts ...option.RequestOption) (res *ZoneSecretGetResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets/%s", params.ZoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ZoneSecretService) Update(ctx context.Context, id string, params ZoneSecretUpdateParams, opts ...option.RequestOption) (res *ZoneSecretUpdateResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets/%s", params.ZoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

func (r *ZoneSecretService) List(ctx context.Context, zoneID string, params ZoneSecretListParams, opts ...option.RequestOption) (res *[]ZoneSecretListResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

func (r *ZoneSecretService) Delete(ctx context.Context, id string, params ZoneSecretDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secrets/%s", params.ZoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ZoneSecretNewResponse struct {
	// A globally unique opaque identifier
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id,required"`
	// A name for the entity to be displayed in UI
	Name string `json:"name,required"`
	// Any of "token", "password".
	Type      ZoneSecretNewResponseType `json:"type,required"`
	UpdatedAt time.Time                 `json:"updated_at,required" format:"date-time"`
	Version   int64                     `json:"version,required"`
	// A globally unique opaque identifier
	ZoneID string `json:"zone_id,required"`
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
func (r ZoneSecretNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneSecretNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretNewResponseType string

const (
	ZoneSecretNewResponseTypeToken    ZoneSecretNewResponseType = "token"
	ZoneSecretNewResponseTypePassword ZoneSecretNewResponseType = "password"
)

type ZoneSecretGetResponse struct {
	// A globally unique opaque identifier
	ID        string                         `json:"id,required"`
	CreatedAt time.Time                      `json:"created_at,required" format:"date-time"`
	Data      ZoneSecretGetResponseDataUnion `json:"data,required"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id,required"`
	// A name for the entity to be displayed in UI
	Name      string    `json:"name,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	Version   int64     `json:"version,required"`
	// A globally unique opaque identifier
	ZoneID string `json:"zone_id,required"`
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
// [ZoneSecretGetResponseDataToken], [ZoneSecretGetResponseDataPassword].
//
// Use the [ZoneSecretGetResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ZoneSecretGetResponseDataUnion struct {
	// This field is from variant [ZoneSecretGetResponseDataToken].
	Token string `json:"token"`
	// Any of "token", "password".
	Type string `json:"type"`
	// This field is from variant [ZoneSecretGetResponseDataPassword].
	Password string `json:"password"`
	// This field is from variant [ZoneSecretGetResponseDataPassword].
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

func (ZoneSecretGetResponseDataToken) implZoneSecretGetResponseDataUnion()    {}
func (ZoneSecretGetResponseDataPassword) implZoneSecretGetResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ZoneSecretGetResponseDataUnion.AsAny().(type) {
//	case keycardapi.ZoneSecretGetResponseDataToken:
//	case keycardapi.ZoneSecretGetResponseDataPassword:
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

func (u ZoneSecretGetResponseDataUnion) AsToken() (v ZoneSecretGetResponseDataToken) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZoneSecretGetResponseDataUnion) AsPassword() (v ZoneSecretGetResponseDataPassword) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ZoneSecretGetResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ZoneSecretGetResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretGetResponseDataToken struct {
	Token string         `json:"token,required"`
	Type  constant.Token `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneSecretGetResponseDataToken) RawJSON() string { return r.JSON.raw }
func (r *ZoneSecretGetResponseDataToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretGetResponseDataPassword struct {
	Password string            `json:"password,required"`
	Type     constant.Password `json:"type,required"`
	Username string            `json:"username,required"`
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
func (r ZoneSecretGetResponseDataPassword) RawJSON() string { return r.JSON.raw }
func (r *ZoneSecretGetResponseDataPassword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretUpdateResponse struct {
	// A globally unique opaque identifier
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id,required"`
	// A name for the entity to be displayed in UI
	Name string `json:"name,required"`
	// Any of "token", "password".
	Type      ZoneSecretUpdateResponseType `json:"type,required"`
	UpdatedAt time.Time                    `json:"updated_at,required" format:"date-time"`
	Version   int64                        `json:"version,required"`
	// A globally unique opaque identifier
	ZoneID string `json:"zone_id,required"`
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
func (r ZoneSecretUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneSecretUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretUpdateResponseType string

const (
	ZoneSecretUpdateResponseTypeToken    ZoneSecretUpdateResponseType = "token"
	ZoneSecretUpdateResponseTypePassword ZoneSecretUpdateResponseType = "password"
)

type ZoneSecretListResponse struct {
	// A globally unique opaque identifier
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id,required"`
	// A name for the entity to be displayed in UI
	Name string `json:"name,required"`
	// Any of "token", "password".
	Type      ZoneSecretListResponseType `json:"type,required"`
	UpdatedAt time.Time                  `json:"updated_at,required" format:"date-time"`
	Version   int64                      `json:"version,required"`
	// A globally unique opaque identifier
	ZoneID string `json:"zone_id,required"`
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
func (r ZoneSecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneSecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretListResponseType string

const (
	ZoneSecretListResponseTypeToken    ZoneSecretListResponseType = "token"
	ZoneSecretListResponseTypePassword ZoneSecretListResponseType = "password"
)

type ZoneSecretNewParams struct {
	Data ZoneSecretNewParamsDataUnion `json:"data,omitzero,required"`
	// A globally unique opaque identifier
	EntityID string `json:"entity_id,required"`
	// A name for the entity to be displayed in UI
	Name string `json:"name,required"`
	// A description of the entity
	Description      param.Opt[string] `json:"description,omitzero"`
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
	OfToken    *ZoneSecretNewParamsDataToken    `json:",omitzero,inline"`
	OfPassword *ZoneSecretNewParamsDataPassword `json:",omitzero,inline"`
	paramUnion
}

func (u ZoneSecretNewParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfToken, u.OfPassword)
}
func (u *ZoneSecretNewParamsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ZoneSecretNewParamsDataUnion) asAny() any {
	if !param.IsOmitted(u.OfToken) {
		return u.OfToken
	} else if !param.IsOmitted(u.OfPassword) {
		return u.OfPassword
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretNewParamsDataUnion) GetToken() *string {
	if vt := u.OfToken; vt != nil {
		return &vt.Token
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretNewParamsDataUnion) GetPassword() *string {
	if vt := u.OfPassword; vt != nil {
		return &vt.Password
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretNewParamsDataUnion) GetUsername() *string {
	if vt := u.OfPassword; vt != nil {
		return &vt.Username
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretNewParamsDataUnion) GetType() *string {
	if vt := u.OfToken; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPassword; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ZoneSecretNewParamsDataUnion](
		"type",
		apijson.Discriminator[ZoneSecretNewParamsDataToken]("token"),
		apijson.Discriminator[ZoneSecretNewParamsDataPassword]("password"),
	)
}

// The properties Token, Type are required.
type ZoneSecretNewParamsDataToken struct {
	Token string `json:"token,required"`
	// This field can be elided, and will marshal its zero value as "token".
	Type constant.Token `json:"type,required"`
	paramObj
}

func (r ZoneSecretNewParamsDataToken) MarshalJSON() (data []byte, err error) {
	type shadow ZoneSecretNewParamsDataToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneSecretNewParamsDataToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Password, Type, Username are required.
type ZoneSecretNewParamsDataPassword struct {
	Password string `json:"password,required"`
	Username string `json:"username,required"`
	// This field can be elided, and will marshal its zero value as "password".
	Type constant.Password `json:"type,required"`
	paramObj
}

func (r ZoneSecretNewParamsDataPassword) MarshalJSON() (data []byte, err error) {
	type shadow ZoneSecretNewParamsDataPassword
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneSecretNewParamsDataPassword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecretGetParams struct {
	// A globally unique opaque identifier
	ZoneID           string            `path:"zone_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZoneSecretUpdateParams struct {
	// A globally unique opaque identifier
	ZoneID string `path:"zone_id,required" json:"-"`
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
	OfToken    *ZoneSecretUpdateParamsDataToken    `json:",omitzero,inline"`
	OfPassword *ZoneSecretUpdateParamsDataPassword `json:",omitzero,inline"`
	paramUnion
}

func (u ZoneSecretUpdateParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfToken, u.OfPassword)
}
func (u *ZoneSecretUpdateParamsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ZoneSecretUpdateParamsDataUnion) asAny() any {
	if !param.IsOmitted(u.OfToken) {
		return u.OfToken
	} else if !param.IsOmitted(u.OfPassword) {
		return u.OfPassword
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretUpdateParamsDataUnion) GetToken() *string {
	if vt := u.OfToken; vt != nil {
		return &vt.Token
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretUpdateParamsDataUnion) GetPassword() *string {
	if vt := u.OfPassword; vt != nil {
		return &vt.Password
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretUpdateParamsDataUnion) GetUsername() *string {
	if vt := u.OfPassword; vt != nil {
		return &vt.Username
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ZoneSecretUpdateParamsDataUnion) GetType() *string {
	if vt := u.OfToken; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPassword; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ZoneSecretUpdateParamsDataUnion](
		"type",
		apijson.Discriminator[ZoneSecretUpdateParamsDataToken]("token"),
		apijson.Discriminator[ZoneSecretUpdateParamsDataPassword]("password"),
	)
}

// The properties Token, Type are required.
type ZoneSecretUpdateParamsDataToken struct {
	Token string `json:"token,required"`
	// This field can be elided, and will marshal its zero value as "token".
	Type constant.Token `json:"type,required"`
	paramObj
}

func (r ZoneSecretUpdateParamsDataToken) MarshalJSON() (data []byte, err error) {
	type shadow ZoneSecretUpdateParamsDataToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneSecretUpdateParamsDataToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Password, Type, Username are required.
type ZoneSecretUpdateParamsDataPassword struct {
	Password string `json:"password,required"`
	Username string `json:"username,required"`
	// This field can be elided, and will marshal its zero value as "password".
	Type constant.Password `json:"type,required"`
	paramObj
}

func (r ZoneSecretUpdateParamsDataPassword) MarshalJSON() (data []byte, err error) {
	type shadow ZoneSecretUpdateParamsDataPassword
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneSecretUpdateParamsDataPassword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ZoneID           string            `path:"zone_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
