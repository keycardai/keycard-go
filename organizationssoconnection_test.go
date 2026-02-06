// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/keycard-api-go"
	"github.com/stainless-sdks/keycard-api-go/internal/testutil"
	"github.com/stainless-sdks/keycard-api-go/option"
)

func TestOrganizationSSOConnectionGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycardapi.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Organizations.SSOConnection.Get(
		context.TODO(),
		"x",
		keycardapi.OrganizationSSOConnectionGetParams{
			Expand:           []string{"permissions"},
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationSSOConnectionUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycardapi.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Organizations.SSOConnection.Update(
		context.TODO(),
		"x",
		keycardapi.OrganizationSSOConnectionUpdateParams{
			ClientID:     keycardapi.String("client_id"),
			ClientSecret: keycardapi.String("client_secret"),
			Identifier:   keycardapi.String("x"),
			Protocols: keycardapi.OrganizationSSOConnectionUpdateParamsProtocols{
				Oauth2: keycardapi.OrganizationSSOConnectionUpdateParamsProtocolsOauth2{
					AuthorizationEndpoint:         keycardapi.String("https://example.com"),
					CodeChallengeMethodsSupported: []string{"string"},
					JwksUri:                       keycardapi.String("https://example.com"),
					RegistrationEndpoint:          keycardapi.String("https://example.com"),
					ScopesSupported:               []string{"string"},
					TokenEndpoint:                 keycardapi.String("https://example.com"),
				},
				Openid: keycardapi.OrganizationSSOConnectionUpdateParamsProtocolsOpenid{
					UserinfoEndpoint: keycardapi.String("https://example.com"),
				},
			},
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationSSOConnectionDisableWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycardapi.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Organizations.SSOConnection.Disable(
		context.TODO(),
		"x",
		keycardapi.OrganizationSSOConnectionDisableParams{
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationSSOConnectionEnableWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycardapi.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Organizations.SSOConnection.Enable(
		context.TODO(),
		"x",
		keycardapi.OrganizationSSOConnectionEnableParams{
			ClientID:     "client_id",
			Identifier:   "x",
			ClientSecret: keycardapi.String("client_secret"),
			Protocols: keycardapi.OrganizationSSOConnectionEnableParamsProtocols{
				Oauth2: keycardapi.OrganizationSSOConnectionEnableParamsProtocolsOauth2{
					AuthorizationEndpoint:         keycardapi.String("https://example.com"),
					CodeChallengeMethodsSupported: []string{"string"},
					JwksUri:                       keycardapi.String("https://example.com"),
					RegistrationEndpoint:          keycardapi.String("https://example.com"),
					ScopesSupported:               []string{"string"},
					TokenEndpoint:                 keycardapi.String("https://example.com"),
				},
				Openid: keycardapi.OrganizationSSOConnectionEnableParamsProtocolsOpenid{
					UserinfoEndpoint: keycardapi.String("https://example.com"),
				},
			},
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
