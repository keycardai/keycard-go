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

func TestZoneApplicationCredentialNewWithOptionalParams(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Zones.ApplicationCredentials.New(
		context.TODO(),
		"zoneId",
		keycard.ZoneApplicationCredentialNewParams{
			OfApplicationCredentialCreateToken: &keycard.ZoneApplicationCredentialNewParamsBodyApplicationCredentialCreateToken{
				ApplicationID: "application_id",
				ProviderID:    "provider_id",
				Type:          "token",
				Subject:       keycard.String("subject"),
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

func TestZoneApplicationCredentialGet(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Zones.ApplicationCredentials.Get(
		context.TODO(),
		"id",
		keycard.ZoneApplicationCredentialGetParams{
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

func TestZoneApplicationCredentialUpdateWithOptionalParams(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Zones.ApplicationCredentials.Update(
		context.TODO(),
		"id",
		keycard.ZoneApplicationCredentialUpdateParams{
			ZoneID: "zoneId",
			OfTokenCredentialUpdate: &keycard.ZoneApplicationCredentialUpdateParamsBodyTokenCredentialUpdate{
				Subject: keycard.String("subject"),
				Type:    "token",
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

func TestZoneApplicationCredentialListWithOptionalParams(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Zones.ApplicationCredentials.List(
		context.TODO(),
		"zoneId",
		keycard.ZoneApplicationCredentialListParams{
			After:         keycard.String("x"),
			ApplicationID: keycard.String("applicationId"),
			Before:        keycard.String("x"),
			Expand: keycard.ZoneApplicationCredentialListParamsExpandUnion{
				OfZoneApplicationCredentialListsExpandString: keycard.String("total_count"),
			},
			FilterOwnerTypeNe: keycard.ZoneApplicationCredentialListParamsFilterOwnerTypeNePlatform,
			FilterTraitsNe: keycard.ZoneApplicationCredentialListParamsFilterTraitsNeUnion{
				OfString: keycard.String("string"),
			},
			FilterType: keycard.ZoneApplicationCredentialListParamsFilterTypeUnion{
				OfZoneApplicationCredentialListsFilterTypeString: keycard.String("token"),
			},
			Limit: keycard.Int(1),
			Query: keycard.ZoneApplicationCredentialListParamsQueryUnion{
				OfString: keycard.String("x"),
			},
			QueryIdentifier: keycard.ZoneApplicationCredentialListParamsQueryIdentifierUnion{
				OfString: keycard.String("x"),
			},
			QueryProviderName: keycard.ZoneApplicationCredentialListParamsQueryProviderNameUnion{
				OfString: keycard.String("x"),
			},
			Slug: keycard.String("slug"),
			Sort: keycard.String("-created_at, \u000b\f-created_at,\r\r \t\u000b\n\r-created_at,\n\n\f\t-created_at,\n\u000b\r \r\fcreated_at,\n\t\t\n\t\f\f\ncreated_at,\n\u000b\u000b  \n\r\r -created_at,\f \u000b\u000b\f\t\n\n -created_at"),
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

func TestZoneApplicationCredentialDelete(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	err := client.Zones.ApplicationCredentials.Delete(
		context.TODO(),
		"id",
		keycard.ZoneApplicationCredentialDeleteParams{
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
