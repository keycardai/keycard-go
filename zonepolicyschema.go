// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/apiquery"
	shimjson "github.com/keycardai/keycard-go/internal/encoding/json"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

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
//
// ZonePolicySchemaService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicySchemaService] method instead.
type ZonePolicySchemaService struct {
	Options []option.RequestOption
}

// NewZonePolicySchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicySchemaService(opts ...option.RequestOption) (r ZonePolicySchemaService) {
	r = ZonePolicySchemaService{}
	r.Options = opts
	return
}

// Get a policy schema by version
func (r *ZonePolicySchemaService) Get(ctx context.Context, version string, params ZonePolicySchemaGetParams, opts ...option.RequestOption) (res *SchemaVersionWithZoneInfo, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-schemas/%s", url.PathEscape(params.ZoneID), url.PathEscape(version))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List policy schemas
func (r *ZonePolicySchemaService) List(ctx context.Context, zoneID string, params ZonePolicySchemaListParams, opts ...option.RequestOption) (res *ZonePolicySchemaListResponse, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-schemas", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Set the default policy schema for a zone
func (r *ZonePolicySchemaService) SetDefault(ctx context.Context, version string, params ZonePolicySchemaSetDefaultParams, opts ...option.RequestOption) (res *SchemaVersionWithZoneInfo, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-schemas/%s", url.PathEscape(params.ZoneID), url.PathEscape(version))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// A versioned Cedar schema that defines the entity model, actions, and context
// shape used for policy evaluation. The schema contains the valid entity types
// (User, Application, Resource), their attributes, and the allowed attribute
// values. See the Credentials API spec for a full reference of entity attributes
// and valid values.
type SchemaVersion struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Controls what can be done with this schema version:
	//
	//   - `"active"` - new policy versions can be created and validated against it.
	//   - `"deprecated"` - superseded by a newer version but still accepts new policy
	//     versions.
	//   - `"archived"` - closed to new policy versions. Existing policy set versions
	//     pinned to this schema still evaluate normally.
	//
	// Any of "active", "deprecated", "archived".
	Status     SchemaVersionStatus `json:"status" api:"required"`
	UpdatedAt  time.Time           `json:"updated_at" api:"required" format:"date-time"`
	Version    string              `json:"version" api:"required"`
	ArchivedAt time.Time           `json:"archived_at" api:"nullable" format:"date-time"`
	// Cedar schema in human-readable syntax. Populated when format=cedar.
	CedarSchema string `json:"cedar_schema" api:"nullable"`
	// Cedar schema as JSON object. Populated when format=json (default).
	CedarSchemaJson any       `json:"cedar_schema_json" api:"nullable"`
	DeprecatedAt    time.Time `json:"deprecated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Version         respjson.Field
		ArchivedAt      respjson.Field
		CedarSchema     respjson.Field
		CedarSchemaJson respjson.Field
		DeprecatedAt    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchemaVersion) RawJSON() string { return r.JSON.raw }
func (r *SchemaVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls what can be done with this schema version:
//
//   - `"active"` - new policy versions can be created and validated against it.
//   - `"deprecated"` - superseded by a newer version but still accepts new policy
//     versions.
//   - `"archived"` - closed to new policy versions. Existing policy set versions
//     pinned to this schema still evaluate normally.
type SchemaVersionStatus string

const (
	SchemaVersionStatusActive     SchemaVersionStatus = "active"
	SchemaVersionStatusDeprecated SchemaVersionStatus = "deprecated"
	SchemaVersionStatusArchived   SchemaVersionStatus = "archived"
)

// A versioned Cedar schema that defines the entity model, actions, and context
// shape used for policy evaluation. The schema contains the valid entity types
// (User, Application, Resource), their attributes, and the allowed attribute
// values. See the Credentials API spec for a full reference of entity attributes
// and valid values.
type SchemaVersionWithZoneInfo struct {
	// Whether this is the zone's default schema. Clients use this to pre-select which
	// schema to write policies against. Has no effect on evaluation.
	IsDefault bool `json:"is_default" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsDefault   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	SchemaVersion
}

// Returns the unmodified JSON received from the API
func (r SchemaVersionWithZoneInfo) RawJSON() string { return r.JSON.raw }
func (r *SchemaVersionWithZoneInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySchemaListResponse struct {
	Items []SchemaVersionWithZoneInfo `json:"items" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination ZonePolicySchemaListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySchemaListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySchemaListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type ZonePolicySchemaListResponsePagination struct {
	// An opaque cursor used for paginating through a list of results
	AfterCursor string `json:"after_cursor" api:"required"`
	// An opaque cursor used for paginating through a list of results
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items across all pages. Only present when the request includes
	// ?expand[]=total_count.
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
func (r ZonePolicySchemaListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySchemaListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySchemaGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Schema representation format. `cedar` returns human-readable Cedar syntax in
	// `cedar_schema`, `json` returns Cedar JSON schema object in `cedar_schema_json`.
	//
	// Any of "cedar", "json".
	Format ZonePolicySchemaGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySchemaGetParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySchemaGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Schema representation format. `cedar` returns human-readable Cedar syntax in
// `cedar_schema`, `json` returns Cedar JSON schema object in `cedar_schema_json`.
type ZonePolicySchemaGetParamsFormat string

const (
	ZonePolicySchemaGetParamsFormatCedar ZonePolicySchemaGetParamsFormat = "cedar"
	ZonePolicySchemaGetParamsFormatJson  ZonePolicySchemaGetParamsFormat = "json"
)

type ZonePolicySchemaListParams struct {
	// Cursor for forward pagination. Returned in `Pagination.after_cursor`. Mutually
	// exclusive with `before`.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returned in `Pagination.before_cursor`. Mutually
	// exclusive with `after`.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter schemas by default status. When `true`, returns only the zone's default
	// schema. When `false`, returns only non-default schemas. Omit to return all
	// schemas.
	FilterDefault param.Opt[bool] `query:"filter[default],omitzero" json:"-"`
	// **Deprecated.** Use `filter[default]` instead.
	//
	// Filter schemas by default status. When `true`, returns only the zone's default
	// schema. When `false`, returns only non-default schemas. Omit to return all
	// schemas.
	//
	// Still honored for backward compatibility. Supplying both `is_default` and
	// `filter[default]` with conflicting values returns `400 Bad Request`.
	IsDefault param.Opt[bool] `query:"is_default,omitzero" json:"-"`
	// Maximum number of items to return per page.
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// **Deprecated.** Use `expand[]` instead.
	//
	// Opt-in to additional response fields. Still honored for backward compatibility;
	// supplying both `expand` and `expand[]` with disagreeing values returns
	// `400 Bad Request`.
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Schema representation format. `cedar` returns human-readable Cedar syntax in
	// `cedar_schema`, `json` returns Cedar JSON schema object in `cedar_schema_json`.
	//
	// Any of "cedar", "json".
	Format ZonePolicySchemaListParamsFormat `query:"format,omitzero" json:"-"`
	// Sort direction. Default is desc (newest first).
	//
	// Any of "asc", "desc".
	Order ZonePolicySchemaListParamsOrder `query:"order,omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicySchemaListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySchemaListParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySchemaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Schema representation format. `cedar` returns human-readable Cedar syntax in
// `cedar_schema`, `json` returns Cedar JSON schema object in `cedar_schema_json`.
type ZonePolicySchemaListParamsFormat string

const (
	ZonePolicySchemaListParamsFormatCedar ZonePolicySchemaListParamsFormat = "cedar"
	ZonePolicySchemaListParamsFormatJson  ZonePolicySchemaListParamsFormat = "json"
)

// Sort direction. Default is desc (newest first).
type ZonePolicySchemaListParamsOrder string

const (
	ZonePolicySchemaListParamsOrderAsc  ZonePolicySchemaListParamsOrder = "asc"
	ZonePolicySchemaListParamsOrderDesc ZonePolicySchemaListParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicySchemaListParamsSort string

const (
	ZonePolicySchemaListParamsSortCreatedAt ZonePolicySchemaListParamsSort = "created_at"
)

type ZonePolicySchemaSetDefaultParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	Body             any
	paramObj
}

func (r ZonePolicySchemaSetDefaultParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ZonePolicySchemaSetDefaultParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
