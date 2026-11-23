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

func TestZonePolicySetManifestGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Manifest.Get(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetManifestGetParams{
			ZoneID:           "zone_id",
			Version:          keycard.String("version"),
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

func TestZonePolicySetManifestApplyWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicySets.Manifest.Apply(
		context.TODO(),
		"policy_set_id",
		keycard.ZonePolicySetManifestApplyParams{
			ZoneID: "zone_id",
			Entries: []keycard.ZonePolicySetManifestApplyParamsEntryUnion{{
				OfZonePolicySetManifestApplysEntryPdpExistingPolicyEntry: &keycard.ZonePolicySetManifestApplyParamsEntryPdpExistingPolicyEntry{
					PolicyID:        "policy_id",
					CedarJson:       map[string]any{},
					CedarRaw:        keycard.String("cedar_raw"),
					PolicyVersionID: keycard.String("policy_version_id"),
				},
			}},
			DryRun:           keycard.Bool(true),
			Activate:         keycard.Bool(true),
			SchemaVersion:    keycard.String("schema_version"),
			IfMatch:          keycard.String("If-Match"),
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
