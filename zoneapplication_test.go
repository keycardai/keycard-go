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

func TestZoneApplicationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Applications.New(
		context.TODO(),
		"zoneId",
		keycard.ZoneApplicationNewParams{
			Identifier: "x",
			Name:       "x",
			Dependencies: []keycard.ZoneApplicationNewParamsDependency{{
				ID:   "id",
				Type: keycard.String("type"),
			}},
			Description: keycard.String("description"),
			Metadata: keycard.MetadataParam{
				DocsURL: keycard.String("https://example.com"),
			},
			Protocols: keycard.ZoneApplicationNewParamsProtocols{
				Oauth2: keycard.ZoneApplicationNewParamsProtocolsOauth2{
					PostLogoutRedirectUris: []string{"https://example.com"},
					RedirectUris:           []string{"https://example.com"},
				},
			},
			Traits: []keycard.ApplicationTrait{keycard.ApplicationTraitGateway},
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

func TestZoneApplicationGet(t *testing.T) {
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
	_, err := client.Zones.Applications.Get(
		context.TODO(),
		"id",
		keycard.ZoneApplicationGetParams{
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

func TestZoneApplicationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Applications.Update(
		context.TODO(),
		"id",
		keycard.ZoneApplicationUpdateParams{
			ZoneID:      "zoneId",
			Description: keycard.String("description"),
			Identifier:  keycard.String("x"),
			Metadata: keycard.MetadataUpdateParam{
				DocsURL: keycard.String("https://example.com"),
			},
			Name: keycard.String("x"),
			Protocols: keycard.ZoneApplicationUpdateParamsProtocols{
				Oauth2: keycard.ZoneApplicationUpdateParamsProtocolsOauth2{
					PostLogoutRedirectUris: []string{"https://example.com"},
					RedirectUris:           []string{"https://example.com"},
				},
			},
			Traits: []keycard.ApplicationTrait{keycard.ApplicationTraitGateway},
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

func TestZoneApplicationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Applications.List(
		context.TODO(),
		"zoneId",
		keycard.ZoneApplicationListParams{
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Cursor: keycard.String("cursor"),
			Expand: keycard.ZoneApplicationListParamsExpandUnion{
				OfZoneApplicationListsExpandString: keycard.String("total_count"),
			},
			Identifier: keycard.String("identifier"),
			Limit:      keycard.Int(1),
			Slug:       keycard.String("slug"),
			Traits:     []keycard.ApplicationTrait{keycard.ApplicationTraitGateway},
			TraitsAll:  []keycard.ApplicationTrait{keycard.ApplicationTraitGateway},
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

func TestZoneApplicationDelete(t *testing.T) {
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
	err := client.Zones.Applications.Delete(
		context.TODO(),
		"id",
		keycard.ZoneApplicationDeleteParams{
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

func TestZoneApplicationListCredentialsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Applications.ListCredentials(
		context.TODO(),
		"id",
		keycard.ZoneApplicationListCredentialsParams{
			ZoneID: "zoneId",
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Cursor: keycard.String("cursor"),
			Expand: keycard.ZoneApplicationListCredentialsParamsExpandUnion{
				OfZoneApplicationListCredentialssExpandString: keycard.String("total_count"),
			},
			Limit: keycard.Int(1),
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

func TestZoneApplicationListResourcesWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Applications.ListResources(
		context.TODO(),
		"id",
		keycard.ZoneApplicationListResourcesParams{
			ZoneID: "zoneId",
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Cursor: keycard.String("cursor"),
			Expand: keycard.ZoneApplicationListResourcesParamsExpandUnion{
				OfZoneApplicationListResourcessExpandString: keycard.String("total_count"),
			},
			Limit: keycard.Int(1),
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
