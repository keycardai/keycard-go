// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/keycard-api-go/internal/apijson"
	"github.com/stainless-sdks/keycard-api-go/internal/requestconfig"
	"github.com/stainless-sdks/keycard-api-go/option"
	"github.com/stainless-sdks/keycard-api-go/packages/respjson"
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
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/users/%s", query.ZoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of users in the specified zone. Can be filtered by email address.
func (r *ZoneUserService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneUserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/users", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// An authenticated user entity
type User struct {
	// Unique identifier of the user
	ID string `json:"id,required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Email address of the user
	Email string `json:"email,required" format:"email"`
	// Issuer identifier of the identity provider
	Issuer string `json:"issuer,required"`
	// Organization that owns this user
	OrganizationID string `json:"organization_id,required"`
	// Subject identifier from the identity provider
	Subject string `json:"subject,required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Zone this user belongs to
	ZoneID string `json:"zone_id,required"`
	// Date when the user was last authenticated
	AuthenticatedAt string `json:"authenticated_at"`
	// Reference to the identity provider. This field is undefined when the source
	// identity provider is deleted but the user is not deleted.
	ProviderID string `json:"provider_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		Issuer          respjson.Field
		OrganizationID  respjson.Field
		Subject         respjson.Field
		UpdatedAt       respjson.Field
		ZoneID          respjson.Field
		AuthenticatedAt respjson.Field
		ProviderID      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserListResponse struct {
	Items []User `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserGetParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}
