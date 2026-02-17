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

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/apiquery"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// ZoneSessionService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSessionService] method instead.
type ZoneSessionService struct {
	Options []option.RequestOption
}

// NewZoneSessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSessionService(opts ...option.RequestOption) (r ZoneSessionService) {
	r = ZoneSessionService{}
	r.Options = opts
	return
}

// Returns details of a specific session by session ID
func (r *ZoneSessionService) Get(ctx context.Context, id string, query ZoneSessionGetParams, opts ...option.RequestOption) (res *SessionUnion, err error) {
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
	path := fmt.Sprintf("zones/%s/sessions/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revokes an active session
func (r *ZoneSessionService) Update(ctx context.Context, id string, params ZoneSessionUpdateParams, opts ...option.RequestOption) (res *SessionUnion, err error) {
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
	path := fmt.Sprintf("zones/%s/sessions/%s", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns entry sessions in the specified zone. Entry sessions are app user
// sessions with an initiator that are roots or direct children of a root user
// session. Can be filtered by session type, status, and user.
func (r *ZoneSessionService) List(ctx context.Context, zoneID string, query ZoneSessionListParams, opts ...option.RequestOption) (res *ZoneSessionListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/sessions", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes a session, effectively logging out the user or application
func (r *ZoneSessionService) Delete(ctx context.Context, id string, body ZoneSessionDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("zones/%s/sessions/%s", url.PathEscape(body.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// SessionUnion contains all possible properties and values from
// [SessionUserSessionType], [SessionApplicationSessionType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SessionUnion struct {
	SessionType string `json:"session_type"`
	// This field is from variant [SessionUserSessionType].
	UserID string `json:"user_id"`
	ID     string `json:"id"`
	Active bool   `json:"active"`
	// This field is from variant [SessionUserSessionType].
	Application     Application `json:"application"`
	ApplicationID   string      `json:"application_id"`
	AuthenticatedAt time.Time   `json:"authenticated_at"`
	CreatedAt       time.Time   `json:"created_at"`
	ExpiresAt       time.Time   `json:"expires_at"`
	Issuer          string      `json:"issuer"`
	// This field is a union of [SessionUserSessionTypeMetadata],
	// [SessionApplicationSessionTypeMetadata]
	Metadata       SessionUnionMetadata `json:"metadata"`
	OrganizationID string               `json:"organization_id"`
	ParentID       string               `json:"parent_id"`
	ProviderID     string               `json:"provider_id"`
	SessionData    any                  `json:"session_data"`
	Status         string               `json:"status"`
	Subject        string               `json:"subject"`
	UpdatedAt      time.Time            `json:"updated_at"`
	// This field is from variant [SessionUserSessionType].
	User User `json:"user"`
	// This field is from variant [SessionUserSessionType].
	UserAgent UserAgent `json:"user_agent"`
	// This field is from variant [SessionUserSessionType].
	UserAgentID string `json:"user_agent_id"`
	ZoneID      string `json:"zone_id"`
	JSON        struct {
		SessionType     respjson.Field
		UserID          respjson.Field
		ID              respjson.Field
		Active          respjson.Field
		Application     respjson.Field
		ApplicationID   respjson.Field
		AuthenticatedAt respjson.Field
		CreatedAt       respjson.Field
		ExpiresAt       respjson.Field
		Issuer          respjson.Field
		Metadata        respjson.Field
		OrganizationID  respjson.Field
		ParentID        respjson.Field
		ProviderID      respjson.Field
		SessionData     respjson.Field
		Status          respjson.Field
		Subject         respjson.Field
		UpdatedAt       respjson.Field
		User            respjson.Field
		UserAgent       respjson.Field
		UserAgentID     respjson.Field
		ZoneID          respjson.Field
		raw             string
	} `json:"-"`
}

func (u SessionUnion) AsUserSessionType() (v SessionUserSessionType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SessionUnion) AsApplicationSessionType() (v SessionApplicationSessionType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SessionUnion) RawJSON() string { return u.JSON.raw }

func (r *SessionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SessionUnionMetadata is an implicit subunion of [SessionUnion].
// SessionUnionMetadata provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [SessionUnion].
type SessionUnionMetadata struct {
	Name string `json:"name"`
	JSON struct {
		Name respjson.Field
		raw  string
	} `json:"-"`
}

func (r *SessionUnionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User session type-specific fields
type SessionUserSessionType struct {
	// Any of "user".
	SessionType string `json:"session_type,required"`
	// User ID
	UserID string `json:"user_id,required"`
	// Session ID
	ID string `json:"id"`
	// Whether the session is currently active (deprecated - use status instead)
	//
	// Deprecated: deprecated
	Active bool `json:"active"`
	// An Application is a software system with an associated identity that can access
	// Resources. It may act on its own behalf (machine-to-machine) or on behalf of a
	// user (delegated access).
	//
	// Deprecated: deprecated
	Application Application `json:"application"`
	// Application ID that initiated this session
	ApplicationID string `json:"application_id"`
	// Date when the session was authenticated
	AuthenticatedAt time.Time `json:"authenticated_at" format:"date-time"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Date when session expires
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Issuer URL from IdP
	Issuer string `json:"issuer" format:"uri"`
	// Session metadata
	Metadata SessionUserSessionTypeMetadata `json:"metadata"`
	// Organization that owns this session
	OrganizationID string `json:"organization_id"`
	// Parent session ID for hierarchical sessions (user sessions only). When null,
	// this is a web session - a top-level session initiated directly by a user. When
	// set, this is a child session derived from the parent, used for token refresh or
	// delegation. Application sessions cannot have parents.
	ParentID string `json:"parent_id"`
	// Provider ID
	ProviderID string `json:"provider_id"`
	// Session claims data (ID token claims for users, application claims for
	// applications)
	SessionData map[string]any `json:"session_data"`
	// Any of "active", "expired", "revoked".
	Status string `json:"status"`
	// Subject claim from IdP
	Subject string `json:"subject"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// An authenticated user entity
	//
	// Deprecated: deprecated
	User User `json:"user"`
	// A User Agent represents a user agent (browser, desktop app, CLI tool) that can
	// initiate user sessions via OAuth 2.0 Dynamic Client Registration.
	//
	// Deprecated: deprecated
	UserAgent UserAgent `json:"user_agent"`
	// User agent ID (browser/client) that initiated this session
	UserAgentID string `json:"user_agent_id"`
	// Zone this session belongs to
	ZoneID string `json:"zone_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SessionType     respjson.Field
		UserID          respjson.Field
		ID              respjson.Field
		Active          respjson.Field
		Application     respjson.Field
		ApplicationID   respjson.Field
		AuthenticatedAt respjson.Field
		CreatedAt       respjson.Field
		ExpiresAt       respjson.Field
		Issuer          respjson.Field
		Metadata        respjson.Field
		OrganizationID  respjson.Field
		ParentID        respjson.Field
		ProviderID      respjson.Field
		SessionData     respjson.Field
		Status          respjson.Field
		Subject         respjson.Field
		UpdatedAt       respjson.Field
		User            respjson.Field
		UserAgent       respjson.Field
		UserAgentID     respjson.Field
		ZoneID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionUserSessionType) RawJSON() string { return r.JSON.raw }
func (r *SessionUserSessionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Session metadata
type SessionUserSessionTypeMetadata struct {
	// Name of the initiating application or user agent
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionUserSessionTypeMetadata) RawJSON() string { return r.JSON.raw }
func (r *SessionUserSessionTypeMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Application session type-specific fields
type SessionApplicationSessionType struct {
	// Application ID that initiated this session
	ApplicationID string `json:"application_id,required"`
	// Any of "application".
	SessionType string `json:"session_type,required"`
	// Session ID
	ID string `json:"id"`
	// Whether the session is currently active (deprecated - use status instead)
	//
	// Deprecated: deprecated
	Active bool `json:"active"`
	// An Application is a software system with an associated identity that can access
	// Resources. It may act on its own behalf (machine-to-machine) or on behalf of a
	// user (delegated access).
	//
	// Deprecated: deprecated
	Application Application `json:"application"`
	// Date when the session was authenticated
	AuthenticatedAt time.Time `json:"authenticated_at" format:"date-time"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Date when session expires
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Issuer URL from IdP
	Issuer string `json:"issuer" format:"uri"`
	// Session metadata
	Metadata SessionApplicationSessionTypeMetadata `json:"metadata"`
	// Organization that owns this session
	OrganizationID string `json:"organization_id"`
	// Parent session ID for hierarchical sessions (user sessions only). When null,
	// this is a web session - a top-level session initiated directly by a user. When
	// set, this is a child session derived from the parent, used for token refresh or
	// delegation. Application sessions cannot have parents.
	ParentID string `json:"parent_id"`
	// Provider ID
	ProviderID string `json:"provider_id"`
	// Session claims data (ID token claims for users, application claims for
	// applications)
	SessionData map[string]any `json:"session_data"`
	// Any of "active", "expired", "revoked".
	Status string `json:"status"`
	// Subject claim from IdP
	Subject string `json:"subject"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Zone this session belongs to
	ZoneID string `json:"zone_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationID   respjson.Field
		SessionType     respjson.Field
		ID              respjson.Field
		Active          respjson.Field
		Application     respjson.Field
		AuthenticatedAt respjson.Field
		CreatedAt       respjson.Field
		ExpiresAt       respjson.Field
		Issuer          respjson.Field
		Metadata        respjson.Field
		OrganizationID  respjson.Field
		ParentID        respjson.Field
		ProviderID      respjson.Field
		SessionData     respjson.Field
		Status          respjson.Field
		Subject         respjson.Field
		UpdatedAt       respjson.Field
		ZoneID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionApplicationSessionType) RawJSON() string { return r.JSON.raw }
func (r *SessionApplicationSessionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Session metadata
type SessionApplicationSessionTypeMetadata struct {
	// Name of the initiating application or user agent
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionApplicationSessionTypeMetadata) RawJSON() string { return r.JSON.raw }
func (r *SessionApplicationSessionTypeMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSessionListResponse struct {
	Items []SessionUnion `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneSessionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneSessionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSessionGetParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}

type ZoneSessionUpdateParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	// Any of "revoked".
	Status ZoneSessionUpdateParamsStatus `json:"status,omitzero,required"`
	paramObj
}

func (r ZoneSessionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneSessionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneSessionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSessionUpdateParamsStatus string

const (
	ZoneSessionUpdateParamsStatusRevoked ZoneSessionUpdateParamsStatus = "revoked"
)

type ZoneSessionListParams struct {
	// Filter by user ID
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Any of "true".
	Active ZoneSessionListParamsActive `query:"active,omitzero" json:"-"`
	// Any of "user", "application".
	SessionType ZoneSessionListParamsSessionType `query:"session_type,omitzero" json:"-"`
	// Any of "active", "expired", "revoked".
	Status ZoneSessionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneSessionListParams]'s query parameters as `url.Values`.
func (r ZoneSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSessionListParamsActive string

const (
	ZoneSessionListParamsActiveTrue ZoneSessionListParamsActive = "true"
)

type ZoneSessionListParamsSessionType string

const (
	ZoneSessionListParamsSessionTypeUser        ZoneSessionListParamsSessionType = "user"
	ZoneSessionListParamsSessionTypeApplication ZoneSessionListParamsSessionType = "application"
)

type ZoneSessionListParamsStatus string

const (
	ZoneSessionListParamsStatusActive  ZoneSessionListParamsStatus = "active"
	ZoneSessionListParamsStatusExpired ZoneSessionListParamsStatus = "expired"
	ZoneSessionListParamsStatusRevoked ZoneSessionListParamsStatus = "revoked"
)

type ZoneSessionDeleteParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}
