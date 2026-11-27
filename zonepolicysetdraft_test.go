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

func TestZonePolicySetDraftNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Draft.New(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetDraftNewParams{
			ZoneID:        "zone_id",
			BaseVersionID: keycard.String("base_version_id"),
			Description:   keycard.String("description"),
			Manifest: keycard.PolicySetManifestParam{
				Entries: []keycard.PolicySetManifestEntryParam{{
					PolicyID:        "policy_id",
					PolicyVersionID: "policy_version_id",
					Sha:             keycard.String("sha"),
				}},
			},
			Name:             keycard.String("name"),
			SchemaVersion:    keycard.String("schema_version"),
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

func TestZonePolicySetDraftGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Draft.Get(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetDraftGetParams{
			ZoneID: "zone_id",
			Expand: keycard.ZonePolicySetDraftGetParamsExpandUnion{
				OfZonePolicySetDraftGetsExpandPdpExpandSingleField: keycard.String("user"),
			},
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

func TestZonePolicySetDraftUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Draft.Update(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetDraftUpdateParams{
			ZoneID:      "zone_id",
			Description: keycard.String("description"),
			ManifestOperations: []keycard.ZonePolicySetDraftUpdateParamsManifestOperation{{
				Op:              "set",
				PolicyID:        "policy_id",
				PolicyVersionID: keycard.String("policy_version_id"),
			}},
			Name:             keycard.String("name"),
			SchemaVersion:    keycard.String("schema_version"),
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

func TestZonePolicySetDraftDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Zones.PolicySets.Draft.Delete(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetDraftDeleteParams{
			ZoneID:           "zone_id",
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

func TestZonePolicySetDraftListPoliciesWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Draft.ListPolicies(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetDraftListPoliciesParams{
			ZoneID: "zone_id",
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Expand: keycard.ZonePolicySetDraftListPoliciesParamsExpandUnion{
				OfZonePolicySetDraftListPoliciessExpandPdpExpandField: keycard.String("total_count"),
			},
			Format:           keycard.ZonePolicySetDraftListPoliciesParamsFormatCedar,
			Limit:            keycard.Int(1),
			Order:            keycard.ZonePolicySetDraftListPoliciesParamsOrderAsc,
			Sort:             keycard.ZonePolicySetDraftListPoliciesParamsSortCreatedAt,
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

func TestZonePolicySetDraftUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Draft.Upsert(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetDraftUpsertParams{
			ZoneID: "zone_id",
			Manifest: keycard.PolicySetManifestParam{
				Entries: []keycard.PolicySetManifestEntryParam{{
					PolicyID:        "policy_id",
					PolicyVersionID: "policy_version_id",
					Sha:             keycard.String("sha"),
				}},
			},
			SchemaVersion:    "schema_version",
			Description:      keycard.String("description"),
			Name:             keycard.String("name"),
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
