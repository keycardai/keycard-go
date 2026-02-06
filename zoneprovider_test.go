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

func TestZoneProviderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Providers.New(
		context.TODO(),
		"zoneId",
		keycardapi.ZoneProviderNewParams{
			Identifier:   "x",
			Name:         "x",
			ClientID:     keycardapi.String("client_id"),
			ClientSecret: keycardapi.String("client_secret"),
			Description:  keycardapi.String("description"),
			Metadata:     map[string]any{},
			Protocols: keycardapi.ZoneProviderNewParamsProtocols{
				Oauth2: keycardapi.ZoneProviderNewParamsProtocolsOauth2{
					AuthorizationEndpoint:          keycardapi.String("https://example.com"),
					AuthorizationResourceEnabled:   keycardapi.Bool(true),
					AuthorizationResourceParameter: keycardapi.String("authorization_resource_parameter"),
					CodeChallengeMethodsSupported:  []string{"string"},
					JwksUri:                        keycardapi.String("https://example.com"),
					RegistrationEndpoint:           keycardapi.String("https://example.com"),
					ScopesSupported:                []string{"string"},
					TokenEndpoint:                  keycardapi.String("https://example.com"),
				},
				Openid: keycardapi.ZoneProviderNewParamsProtocolsOpenid{
					UserinfoEndpoint: keycardapi.String("https://example.com"),
				},
			},
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

func TestZoneProviderGet(t *testing.T) {
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
	_, err := client.Zones.Providers.Get(
		context.TODO(),
		"id",
		keycardapi.ZoneProviderGetParams{
			ZoneID: "zoneId",
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

func TestZoneProviderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Providers.Update(
		context.TODO(),
		"id",
		keycardapi.ZoneProviderUpdateParams{
			ZoneID:       "zoneId",
			ClientID:     keycardapi.String("client_id"),
			ClientSecret: keycardapi.String("client_secret"),
			Description:  keycardapi.String("description"),
			Identifier:   keycardapi.String("x"),
			Metadata:     map[string]any{},
			Name:         keycardapi.String("x"),
			Protocols: keycardapi.ZoneProviderUpdateParamsProtocols{
				Oauth2: keycardapi.ZoneProviderUpdateParamsProtocolsOauth2{
					AuthorizationEndpoint:          keycardapi.String("https://example.com"),
					AuthorizationResourceEnabled:   keycardapi.Bool(true),
					AuthorizationResourceParameter: keycardapi.String("authorization_resource_parameter"),
					CodeChallengeMethodsSupported:  []string{"string"},
					JwksUri:                        keycardapi.String("https://example.com"),
					RegistrationEndpoint:           keycardapi.String("https://example.com"),
					ScopesSupported:                []string{"string"},
					TokenEndpoint:                  keycardapi.String("https://example.com"),
				},
				Openid: keycardapi.ZoneProviderUpdateParamsProtocolsOpenid{
					UserinfoEndpoint: keycardapi.String("https://example.com"),
				},
			},
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

func TestZoneProviderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Providers.List(
		context.TODO(),
		"zoneId",
		keycardapi.ZoneProviderListParams{
			Cursor:     keycardapi.String("cursor"),
			Identifier: keycardapi.String("identifier"),
			Limit:      keycardapi.Int(1),
			Slug:       keycardapi.String("slug"),
			Type:       keycardapi.ZoneProviderListParamsTypeExternal,
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

func TestZoneProviderDelete(t *testing.T) {
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
	err := client.Zones.Providers.Delete(
		context.TODO(),
		"id",
		keycardapi.ZoneProviderDeleteParams{
			ZoneID: "zoneId",
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
