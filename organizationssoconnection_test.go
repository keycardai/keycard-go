// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/keycardai/keycard-go"
	"github.com/keycardai/keycard-go/internal/testutil"
	"github.com/keycardai/keycard-go/option"
)

func TestOrganizationSSOConnectionGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Organizations.SSOConnection.Get(
		context.TODO(),
		"x",
		keycard.OrganizationSSOConnectionGetParams{
			Expand:           []string{"permissions"},
			XClientRequestID: keycard.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationSSOConnectionUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Organizations.SSOConnection.Update(
		context.TODO(),
		"x",
		keycard.OrganizationSSOConnectionUpdateParams{
			ClientID:     keycard.String("client_id"),
			ClientSecret: keycard.String("client_secret"),
			Identifier:   keycard.String("x"),
			Protocols: keycard.SSOConnectionProtocolParam{
				Oauth2: keycard.SSOConnectionProtocolOauth2Param{
					AuthorizationEndpoint:         keycard.String("https://example.com"),
					CodeChallengeMethodsSupported: []string{"string"},
					JwksUri:                       keycard.String("https://example.com"),
					RegistrationEndpoint:          keycard.String("https://example.com"),
					ScopesSupported:               []string{"string"},
					TokenEndpoint:                 keycard.String("https://example.com"),
				},
				Openid: keycard.SSOConnectionProtocolOpenidParam{
					UserinfoEndpoint: keycard.String("https://example.com"),
				},
			},
			XClientRequestID: keycard.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationSSOConnectionDisableWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	err := client.Organizations.SSOConnection.Disable(
		context.TODO(),
		"x",
		keycard.OrganizationSSOConnectionDisableParams{
			XClientRequestID: keycard.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationSSOConnectionEnableWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Organizations.SSOConnection.Enable(
		context.TODO(),
		"x",
		keycard.OrganizationSSOConnectionEnableParams{
			ClientID:     "client_id",
			Identifier:   "x",
			ClientSecret: keycard.String("client_secret"),
			Protocols: keycard.SSOConnectionProtocolParam{
				Oauth2: keycard.SSOConnectionProtocolOauth2Param{
					AuthorizationEndpoint:         keycard.String("https://example.com"),
					CodeChallengeMethodsSupported: []string{"string"},
					JwksUri:                       keycard.String("https://example.com"),
					RegistrationEndpoint:          keycard.String("https://example.com"),
					ScopesSupported:               []string{"string"},
					TokenEndpoint:                 keycard.String("https://example.com"),
				},
				Openid: keycard.SSOConnectionProtocolOpenidParam{
					UserinfoEndpoint: keycard.String("https://example.com"),
				},
			},
			XClientRequestID: keycard.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *keycard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
