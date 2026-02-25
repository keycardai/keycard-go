// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/keycardlabs/keycard-go/internal/apijson"
	"github.com/keycardlabs/keycard-go/internal/requestconfig"
	"github.com/keycardlabs/keycard-go/option"
	"github.com/keycardlabs/keycard-go/packages/param"
)

// ServiceAccountTokenService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceAccountTokenService] method instead.
type ServiceAccountTokenService struct {
	Options []option.RequestOption
}

// NewServiceAccountTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewServiceAccountTokenService(opts ...option.RequestOption) (r ServiceAccountTokenService) {
	r = ServiceAccountTokenService{}
	r.Options = opts
	return
}

// Exchange service account credentials for organization-scoped M2M token
func (r *ServiceAccountTokenService) New(ctx context.Context, params ServiceAccountTokenNewParams, opts ...option.RequestOption) (res *TokenResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "service-account-token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ServiceAccountTokenNewParams struct {
	// Service account client ID
	ClientID string `json:"client_id" api:"required"`
	// Service account client secret
	ClientSecret string `json:"client_secret" api:"required"`
	// OAuth 2.0 grant type (must be "client_credentials")
	//
	// Any of "client_credentials".
	GrantType        ServiceAccountTokenNewParamsGrantType `json:"grant_type,omitzero" api:"required"`
	XClientRequestID param.Opt[string]                     `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ServiceAccountTokenNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ServiceAccountTokenNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ServiceAccountTokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 grant type (must be "client_credentials")
type ServiceAccountTokenNewParamsGrantType string

const (
	ServiceAccountTokenNewParamsGrantTypeClientCredentials ServiceAccountTokenNewParamsGrantType = "client_credentials"
)
