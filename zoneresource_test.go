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

func TestZoneResourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Resources.New(
		context.TODO(),
		"zoneId",
		keycard.ZoneResourceNewParams{
			Identifier:           "x",
			Name:                 "x",
			ApplicationID:        keycard.String("application_id"),
			ApplicationType:      keycard.ZoneResourceNewParamsApplicationTypeNative,
			CredentialProviderID: keycard.String("credential_provider_id"),
			Description:          keycard.String("description"),
			Metadata: keycard.MetadataParam{
				DocsURL: keycard.String("https://example.com"),
			},
			Scopes: []string{"string"},
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

func TestZoneResourceGet(t *testing.T) {
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
	_, err := client.Zones.Resources.Get(
		context.TODO(),
		"id",
		keycard.ZoneResourceGetParams{
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

func TestZoneResourceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Resources.Update(
		context.TODO(),
		"id",
		keycard.ZoneResourceUpdateParams{
			ZoneID:               "zoneId",
			ApplicationID:        keycard.String("application_id"),
			ApplicationType:      keycard.ZoneResourceUpdateParamsApplicationTypeNative,
			CredentialProviderID: keycard.String("credential_provider_id"),
			Description:          keycard.String("description"),
			Identifier:           keycard.String("x"),
			Metadata: keycard.MetadataUpdateParam{
				DocsURL: keycard.String("https://example.com"),
			},
			Name:   keycard.String("x"),
			Scopes: []string{"x"},
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

func TestZoneResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Resources.List(
		context.TODO(),
		"zoneId",
		keycard.ZoneResourceListParams{
			After:                keycard.String("x"),
			Before:               keycard.String("x"),
			CredentialProviderID: keycard.String("credentialProviderId"),
			Expand: keycard.ZoneResourceListParamsExpandUnion{
				OfZoneResourceListsExpandString: keycard.String("total_count"),
			},
			Identifier: keycard.String("identifier"),
			Limit:      keycard.Int(1),
			Slug:       keycard.String("slug"),
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

func TestZoneResourceDelete(t *testing.T) {
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
	err := client.Zones.Resources.Delete(
		context.TODO(),
		"id",
		keycard.ZoneResourceDeleteParams{
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
