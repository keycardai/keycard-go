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

func TestZoneProviderNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Zones.Providers.New(
		context.TODO(),
		"zoneId",
		keycard.ZoneProviderNewParams{
			Identifier:   "x",
			Name:         "x",
			ClientID:     keycard.String("client_id"),
			ClientSecret: keycard.String("client_secret"),
			Description:  keycard.String("description"),
			Metadata:     map[string]any{},
			Protocols: keycard.ZoneProviderNewParamsProtocols{
				Oauth2: keycard.ZoneProviderNewParamsProtocolsOauth2{
					AuthorizationEndpoint: keycard.String("https://example.com"),
					AuthorizationParameters: map[string]string{
						"foo": "string",
					},
					AuthorizationResourceEnabled:    keycard.Bool(true),
					AuthorizationResourceParameter:  keycard.String("authorization_resource_parameter"),
					CodeChallengeMethodsSupported:   []string{"string"},
					Issuer:                          keycard.String("https://example.com"),
					JwksUri:                         keycard.String("https://example.com"),
					RegistrationEndpoint:            keycard.String("https://example.com"),
					ScopeParameter:                  keycard.String("scope_parameter"),
					ScopeSeparator:                  keycard.String("scope_separator"),
					ScopesSupported:                 []string{"string"},
					TokenEndpoint:                   keycard.String("https://example.com"),
					TokenResponseAccessTokenPointer: keycard.String("token_response_access_token_pointer"),
				},
				Openid: keycard.ZoneProviderNewParamsProtocolsOpenid{
					UserinfoEndpoint: keycard.String("https://example.com"),
				},
			},
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

func TestZoneProviderGet(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Zones.Providers.Get(
		context.TODO(),
		"id",
		keycard.ZoneProviderGetParams{
			ZoneID: "zoneId",
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

func TestZoneProviderUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Zones.Providers.Update(
		context.TODO(),
		"id",
		keycard.ZoneProviderUpdateParams{
			ZoneID:       "zoneId",
			ClientID:     keycard.String("client_id"),
			ClientSecret: keycard.String("client_secret"),
			Description:  keycard.String("description"),
			Identifier:   keycard.String("x"),
			Metadata:     map[string]any{},
			Name:         keycard.String("x"),
			Protocols: keycard.ZoneProviderUpdateParamsProtocols{
				Oauth2: keycard.ZoneProviderUpdateParamsProtocolsOauth2{
					AuthorizationEndpoint: keycard.String("https://example.com"),
					AuthorizationParameters: map[string]string{
						"foo": "string",
					},
					AuthorizationResourceEnabled:    keycard.Bool(true),
					AuthorizationResourceParameter:  keycard.String("authorization_resource_parameter"),
					CodeChallengeMethodsSupported:   []string{"string"},
					Issuer:                          keycard.String("https://example.com"),
					JwksUri:                         keycard.String("https://example.com"),
					RegistrationEndpoint:            keycard.String("https://example.com"),
					ScopeParameter:                  keycard.String("scope_parameter"),
					ScopeSeparator:                  keycard.String("scope_separator"),
					ScopesSupported:                 []string{"string"},
					TokenEndpoint:                   keycard.String("https://example.com"),
					TokenResponseAccessTokenPointer: keycard.String("token_response_access_token_pointer"),
				},
				Openid: keycard.ZoneProviderUpdateParamsProtocolsOpenid{
					UserinfoEndpoint: keycard.String("https://example.com"),
				},
			},
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

func TestZoneProviderListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Zones.Providers.List(
		context.TODO(),
		"zoneId",
		keycard.ZoneProviderListParams{
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Cursor: keycard.String("cursor"),
			Expand: keycard.ZoneProviderListParamsExpandUnion{
				OfZoneProviderListsExpandString: keycard.String("total_count"),
			},
			Identifier: keycard.String("identifier"),
			Limit:      keycard.Int(1),
			Slug:       keycard.String("slug"),
			Type:       keycard.ZoneProviderListParamsTypeExternal,
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

func TestZoneProviderDelete(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	err := client.Zones.Providers.Delete(
		context.TODO(),
		"id",
		keycard.ZoneProviderDeleteParams{
			ZoneID: "zoneId",
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
