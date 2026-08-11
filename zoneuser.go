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

// ZoneUserService contains methods and other services that help with interacting
// with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneUserService] method instead.
type ZoneUserService struct {
	Options []option.RequestOption
}

// NewZoneUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneUserService(opts ...option.RequestOption) (r ZoneUserService) {
	r = ZoneUserService{}
	r.Options = opts
	return
}

// Returns details of a specific user by user ID
func (r *ZoneUserService) Get(ctx context.Context, id string, query ZoneUserGetParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/users/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a list of users in the specified zone.
//
// Note: cursor pagination, search, and sort are not yet enabled for all zones.
// Where they are not enabled, the response returns all users in the zone (capped
// at 100) in `items`, with `after_cursor` and `before_cursor` set to `null` and
// `total_count` of `0`; `filter[email]` and `filter[identifier]` are still
// applied, while the pagination, search, and sort parameters below are accepted
// but ignored.
//
// Use cursor pagination via `after`/`before`. Sort: comma-separated field list;
// prefix with `-` for descending. Use `expand[]=total_count` to include the
// matching row count, `expand[]=session_count` to include per-user session counts,
// `expand[]=grant_count` to include per-user delegated-grant counts,
// `expand[]=role-assignments` to include each user's structured role grants,
// `expand[]=credentials` to include each user's authentication credentials (each
// with its `provider_id`), and `expand[]=credentials.provider` to additionally
// inline the full identity provider on each federation credential. Filter by exact
// email via `filter[email]` and by exact identifier via `filter[identifier]`;
// search via `query[email]` / `query[subject]` / `query[]` (substring match, OR'd
// across repeated values). `query[]` matches against email and federation
// credential subject. Pass `filter[id]` (repeatable, max 100) to restrict results
// to a known set of users — mutually exclusive with `after`/`before` (returns 400
// if combined). When `filter[id]` is set, `limit` is ignored and the response
// contains every requested user that exists in the zone, in a single page. IDs not
// in the zone are silently omitted.
func (r *ZoneUserService) List(ctx context.Context, zoneID string, query ZoneUserListParams, opts ...option.RequestOption) (res *ZoneUserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/users", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// An authenticated user entity
type User struct {
	// Unique identifier of the user
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address of the user
	Email string `json:"email" api:"required" format:"email"`
	// Whether the email address has been verified
	EmailVerified bool `json:"email_verified" api:"required"`
	// Zone-scoped user identifier. Defaults to the user's Keycard ID. When the
	// provider has user_identifier_claim configured, the value is set from that claim
	// at user creation time.
	Identifier string `json:"identifier" api:"required"`
	// Organization that owns this user
	OrganizationID string `json:"organization_id" api:"required"`
	// Status of the user. Disabled users cannot authenticate.
	//
	// Any of "active", "disabled".
	Status UserStatus `json:"status" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zone this user belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// Date when the user was last authenticated
	AuthenticatedAt string `json:"authenticated_at"`
	// Authentication credentials for this user, each carrying its identity provider
	// for federation credentials. Populated only when `expand[]=credentials` is set on
	// the listing endpoint.
	Credentials []UserCredentialUnion `json:"credentials"`
	// Delegated-grant count for this user. Populated only when `expand[]=grant_count`
	// is set on the listing endpoint.
	GrantCount int64 `json:"grant_count"`
	// Issuer identifier of the identity provider
	Issuer string `json:"issuer"`
	// Reference to the identity provider. This field is undefined when the source
	// identity provider is deleted but the user is not deleted.
	ProviderID string `json:"provider_id"`
	// Role grants for this user within the zone. Populated only when
	// `expand[]=role-assignments` is set on the listing endpoint.
	RoleAssignments []UserRoleAssignment `json:"role_assignments"`
	// Session count for this user. Populated only when `expand[]=session_count` is set
	// on the listing endpoint.
	SessionCount int64 `json:"session_count"`
	// Subject identifier from the identity provider
	Subject string `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		EmailVerified   respjson.Field
		Identifier      respjson.Field
		OrganizationID  respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		ZoneID          respjson.Field
		AuthenticatedAt respjson.Field
		Credentials     respjson.Field
		GrantCount      respjson.Field
		Issuer          respjson.Field
		ProviderID      respjson.Field
		RoleAssignments respjson.Field
		SessionCount    respjson.Field
		Subject         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the user. Disabled users cannot authenticate.
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusDisabled UserStatus = "disabled"
)

// UserCredentialUnion contains all possible properties and values from
// [UserCredentialUserCredentialFederation],
// [UserCredentialUserCredentialPassword].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserCredentialUnion struct {
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [UserCredentialUserCredentialFederation].
	ProviderID string    `json:"provider_id"`
	Type       string    `json:"type"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant [UserCredentialUserCredentialFederation].
	Issuer string `json:"issuer"`
	// This field is from variant [UserCredentialUserCredentialFederation].
	Provider Provider `json:"provider"`
	// This field is from variant [UserCredentialUserCredentialFederation].
	Subject string `json:"subject"`
	JSON    struct {
		CreatedAt  respjson.Field
		ProviderID respjson.Field
		Type       respjson.Field
		UpdatedAt  respjson.Field
		Issuer     respjson.Field
		Provider   respjson.Field
		Subject    respjson.Field
		raw        string
	} `json:"-"`
}

func (u UserCredentialUnion) AsUserCredentialFederation() (v UserCredentialUserCredentialFederation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserCredentialUnion) AsUserCredentialPassword() (v UserCredentialUserCredentialPassword) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserCredentialUnion) RawJSON() string { return u.JSON.raw }

func (r *UserCredentialUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Federation credential: the user authenticates through an identity provider.
type UserCredentialUserCredentialFederation struct {
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of the identity provider backing this credential. `null` when the source
	// provider has been deleted.
	ProviderID string `json:"provider_id" api:"required"`
	// Any of "federation".
	Type string `json:"type" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Issuer identifier of the identity provider.
	Issuer string `json:"issuer"`
	// A Provider is a system that supplies access to Resources and allows actors
	// (Users or Applications) to authenticate.
	Provider Provider `json:"provider"`
	// Subject identifier from the identity provider.
	Subject string `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		ProviderID  respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Issuer      respjson.Field
		Provider    respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserCredentialUserCredentialFederation) RawJSON() string { return r.JSON.raw }
func (r *UserCredentialUserCredentialFederation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Password credential: the user authenticates with email and password. The email
// lives on the user.
type UserCredentialUserCredentialPassword struct {
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "password".
	Type string `json:"type" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserCredentialUserCredentialPassword) RawJSON() string { return r.JSON.raw }
func (r *UserCredentialUserCredentialPassword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A role granted to a user within a zone.
type UserRoleAssignment struct {
	// ID of the assigned role
	RoleID string `json:"role_id" api:"required"`
	// Role identifier: a lowercase slug (letters and digits separated by single
	// hyphens or underscores), unique per owner type within a zone. Role identifiers
	// surface in policy evaluation, so the slug restriction keeps them unambiguous in
	// policy text.
	RoleIdentifier string `json:"role_identifier" api:"required"`
	// Owner type of the granted role. Disambiguates roles that share an identifier
	// across owner types.
	//
	// Any of "platform", "customer".
	RoleOwnerType string `json:"role_owner_type" api:"required"`
	// The resource this grant is scoped to, or null when the grant is unscoped
	// (applies to the owning zone itself).
	Scope UserRoleAssignmentScope `json:"scope" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RoleID         respjson.Field
		RoleIdentifier respjson.Field
		RoleOwnerType  respjson.Field
		Scope          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRoleAssignment) RawJSON() string { return r.JSON.raw }
func (r *UserRoleAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The resource this grant is scoped to, or null when the grant is unscoped
// (applies to the owning zone itself).
type UserRoleAssignmentScope struct {
	// The ID of the scoped resource.
	ID string `json:"id" api:"required"`
	// The kind of resource this grant is scoped to (e.g. `zone`).
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRoleAssignmentScope) RawJSON() string { return r.JSON.raw }
func (r *UserRoleAssignmentScope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserListResponse struct {
	Items []User `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneUserListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneUserListResponsePagination struct {
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
func (r ZoneUserListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserGetParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneUserListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Comma-separated sort fields. Prefix with - for descending. Allowed: email,
	// authenticated_at
	Sort   param.Opt[string]             `query:"sort,omitzero" json:"-"`
	Expand ZoneUserListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Filter by exact email address
	FilterEmail ZoneUserListParamsFilterEmailUnion `query:"filter[email],omitzero" format:"email" json:"-"`
	// Restrict results to users with this publicId. Repeatable, max 100. Mutually
	// exclusive with after/before.
	FilterID ZoneUserListParamsFilterIDUnion `query:"filter[id],omitzero" json:"-"`
	// Filter by exact user identifier
	FilterIdentifier ZoneUserListParamsFilterIdentifierUnion `query:"filter[identifier],omitzero" json:"-"`
	// Search across email and credential subject (substring match)
	Query ZoneUserListParamsQueryUnion `query:"query[],omitzero" json:"-"`
	// Search by email (substring match)
	QueryEmail ZoneUserListParamsQueryEmailUnion `query:"query[email],omitzero" json:"-"`
	// Search by federated credential subject (substring match)
	QuerySubject ZoneUserListParamsQuerySubjectUnion `query:"query[subject],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneUserListParams]'s query parameters as `url.Values`.
func (r ZoneUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneUserListsExpandString)
	OfZoneUserListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneUserListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneUserListParamsExpandString string

const (
	ZoneUserListParamsExpandStringTotalCount          ZoneUserListParamsExpandString = "total_count"
	ZoneUserListParamsExpandStringSessionCount        ZoneUserListParamsExpandString = "session_count"
	ZoneUserListParamsExpandStringGrantCount          ZoneUserListParamsExpandString = "grant_count"
	ZoneUserListParamsExpandStringRoleAssignments     ZoneUserListParamsExpandString = "role-assignments"
	ZoneUserListParamsExpandStringCredentials         ZoneUserListParamsExpandString = "credentials"
	ZoneUserListParamsExpandStringCredentialsProvider ZoneUserListParamsExpandString = "credentials.provider"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsFilterEmailUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsFilterIDUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsFilterIdentifierUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsQueryUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsQueryEmailUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserListParamsQuerySubjectUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}
