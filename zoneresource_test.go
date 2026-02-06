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

func TestZoneResourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Resources.New(
		context.TODO(),
		"zoneId",
		keycardapi.ZoneResourceNewParams{
			Identifier:           "x",
			Name:                 "x",
			ApplicationID:        keycardapi.String("application_id"),
			CredentialProviderID: keycardapi.String("credential_provider_id"),
			Description:          keycardapi.String("description"),
			Metadata: keycardapi.MetadataParam{
				DocsURL: keycardapi.String("https://example.com"),
			},
			Scopes: []string{"string"},
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

func TestZoneResourceGet(t *testing.T) {
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
	_, err := client.Zones.Resources.Get(
		context.TODO(),
		"id",
		keycardapi.ZoneResourceGetParams{
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

func TestZoneResourceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Resources.Update(
		context.TODO(),
		"id",
		keycardapi.ZoneResourceUpdateParams{
			ZoneID:               "zoneId",
			ApplicationID:        keycardapi.String("application_id"),
			CredentialProviderID: keycardapi.String("credential_provider_id"),
			Description:          keycardapi.String("description"),
			Identifier:           keycardapi.String("x"),
			Metadata: keycardapi.MetadataUpdateParam{
				DocsURL: keycardapi.String("https://example.com"),
			},
			Name:   keycardapi.String("x"),
			Scopes: []string{"x"},
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

func TestZoneResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Resources.List(
		context.TODO(),
		"zoneId",
		keycardapi.ZoneResourceListParams{
			CredentialProviderID: keycardapi.String("credentialProviderId"),
			Identifier:           keycardapi.String("identifier"),
			Slug:                 keycardapi.String("slug"),
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

func TestZoneResourceDelete(t *testing.T) {
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
	err := client.Zones.Resources.Delete(
		context.TODO(),
		"id",
		keycardapi.ZoneResourceDeleteParams{
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
