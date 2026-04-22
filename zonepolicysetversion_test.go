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

func TestZonePolicySetVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Versions.New(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetVersionNewParams{
			ZoneID: "zone_id",
			Manifest: keycard.PolicySetManifestParam{
				Entries: []keycard.PolicySetManifestEntryParam{{
					PolicyID:        "policy_id",
					PolicyVersionID: "policy_version_id",
					Sha:             keycard.String("sha"),
				}},
			},
			SchemaVersion:    "schema_version",
			XAPIVersion:      keycard.String("X-API-Version"),
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

func TestZonePolicySetVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Versions.Get(
		context.TODO(),
		"version_id",
		keycard.ZonePolicySetVersionGetParams{
			ZoneID:           "zone_id",
			PolicySetID:      "policy_set_id",
			XAPIVersion:      keycard.String("X-API-Version"),
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

func TestZonePolicySetVersionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Versions.Update(
		context.TODO(),
		"version_id",
		keycard.ZonePolicySetVersionUpdateParams{
			ZoneID:           "zone_id",
			PolicySetID:      "policy_set_id",
			Active:           true,
			XAPIVersion:      keycard.String("X-API-Version"),
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

func TestZonePolicySetVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Versions.List(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetVersionListParams{
			ZoneID:           "zone_id",
			After:            keycard.String("x"),
			Before:           keycard.String("x"),
			Expand:           []string{"total_count"},
			Limit:            keycard.Int(1),
			Order:            keycard.ZonePolicySetVersionListParamsOrderAsc,
			Sort:             keycard.ZonePolicySetVersionListParamsSortCreatedAt,
			XAPIVersion:      keycard.String("X-API-Version"),
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

func TestZonePolicySetVersionArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Versions.Archive(
		context.TODO(),
		"version_id",
		keycard.ZonePolicySetVersionArchiveParams{
			ZoneID:           "zone_id",
			PolicySetID:      "policy_set_id",
			XAPIVersion:      keycard.String("X-API-Version"),
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

func TestZonePolicySetVersionListPoliciesWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Versions.ListPolicies(
		context.TODO(),
		"version_id",
		keycard.ZonePolicySetVersionListPoliciesParams{
			ZoneID:           "zone_id",
			PolicySetID:      "policy_set_id",
			After:            keycard.String("x"),
			Before:           keycard.String("x"),
			Expand:           []string{"total_count"},
			Format:           keycard.ZonePolicySetVersionListPoliciesParamsFormatCedar,
			Limit:            keycard.Int(1),
			Order:            keycard.ZonePolicySetVersionListPoliciesParamsOrderAsc,
			Sort:             keycard.ZonePolicySetVersionListPoliciesParamsSortCreatedAt,
			XAPIVersion:      keycard.String("X-API-Version"),
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
