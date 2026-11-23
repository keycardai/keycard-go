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

func TestZonePolicyEvaluationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicyEvaluations.New(
		context.TODO(),
		"zone_id",
		keycard.ZonePolicyEvaluationNewParams{
			Action: "action",
			PolicySource: keycard.ZonePolicyEvaluationNewParamsPolicySource{
				Manifest: keycard.ZonePolicyEvaluationNewParamsPolicySourceManifest{
					Entries: []keycard.PolicySetManifestEntryParam{{
						PolicyID:        "policy_id",
						PolicyVersionID: "policy_version_id",
						Sha:             keycard.String("sha"),
					}},
					SchemaVersion: "schema_version",
				},
				Policy: keycard.ZonePolicyEvaluationNewParamsPolicySourcePolicy{
					PolicyID:      "policy_id",
					SchemaVersion: keycard.String("schema_version"),
					VersionID:     keycard.String("version_id"),
				},
				PolicySet: keycard.ZonePolicyEvaluationNewParamsPolicySourcePolicySet{
					PolicySetID: "policy_set_id",
					VersionID:   keycard.String("version_id"),
				},
				Raw: keycard.ZonePolicyEvaluationNewParamsPolicySourceRaw{
					Cedar:         "cedar",
					SchemaVersion: "schema_version",
					Validate:      keycard.Bool(true),
				},
			},
			PrincipalID:       "principal_id",
			PrincipalType:     keycard.ZonePolicyEvaluationNewParamsPrincipalTypeUser,
			ResourceID:        "resource_id",
			ActingPrincipalID: keycard.String("acting_principal_id"),
			AuthorityType:     keycard.ZonePolicyEvaluationNewParamsAuthorityTypeDelegation,
			Context: keycard.ZonePolicyEvaluationNewParamsContext{
				Claims: keycard.ZonePolicyEvaluationNewParamsContextClaims{
					ActingPrincipal: keycard.ZonePolicyEvaluationNewParamsContextClaimsActingPrincipal{
						Email:  keycard.String("email"),
						Groups: []string{"string"},
						IssuerClaims: map[string]any{
							"foo": "bar",
						},
					},
					Actor: keycard.ZonePolicyEvaluationNewParamsContextClaimsActor{
						Email:  keycard.String("email"),
						Groups: []string{"string"},
						IssuerClaims: map[string]any{
							"foo": "bar",
						},
					},
					Subject: keycard.ZonePolicyEvaluationNewParamsContextClaimsSubject{
						Email:  keycard.String("email"),
						Groups: []string{"string"},
						IssuerClaims: map[string]any{
							"foo": "bar",
						},
					},
				},
				Scopes:    []string{"string"},
				SessionID: keycard.String("session_id"),
			},
			SubjectID:        keycard.String("subject_id"),
			SubjectType:      keycard.ZonePolicyEvaluationNewParamsSubjectTypeUser,
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

func TestZonePolicyEvaluationListRequestsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PolicyEvaluations.ListRequests(
		context.TODO(),
		"zone_id",
		keycard.ZonePolicyEvaluationListRequestsParams{
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Expand: keycard.ZonePolicyEvaluationListRequestsParamsExpandUnion{
				OfZonePolicyEvaluationListRequestssExpandPdpExpandField: keycard.String("total_count"),
			},
			FilterActiveDecision: []string{"string"},
			FilterAuthorityType:  []string{"string"},
			FilterHasShadow:      keycard.Bool(true),
			FilterPrincipalID:    []string{"string"},
			FilterPrincipalType:  []string{"string"},
			FilterResourceID:     []string{"string"},
			FilterSubjectID:      []string{"string"},
			FilterSubjectType:    []string{"string"},
			Limit:                keycard.Int(1),
			Order:                keycard.ZonePolicyEvaluationListRequestsParamsOrderAsc,
			QueryAction:          []string{"x"},
			Sort:                 keycard.ZonePolicyEvaluationListRequestsParamsSortCreatedAt,
			XAPIVersion:          keycard.String("X-API-Version"),
			XClientRequestID:     keycard.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
