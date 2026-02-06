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

func TestZoneApplicationDependencyGet(t *testing.T) {
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
	_, err := client.Zones.Applications.Dependencies.Get(
		context.TODO(),
		"dependencyId",
		keycardapi.ZoneApplicationDependencyGetParams{
			ZoneID: "zoneId",
			ID:     "id",
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

func TestZoneApplicationDependencyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Applications.Dependencies.List(
		context.TODO(),
		"id",
		keycardapi.ZoneApplicationDependencyListParams{
			ZoneID:        "zoneId",
			Cursor:        keycardapi.String("cursor"),
			Limit:         keycardapi.Int(1),
			WhenAccessing: keycardapi.String("when_accessing"),
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

func TestZoneApplicationDependencyAddWithOptionalParams(t *testing.T) {
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
	err := client.Zones.Applications.Dependencies.Add(
		context.TODO(),
		"dependencyId",
		keycardapi.ZoneApplicationDependencyAddParams{
			ZoneID:        "zoneId",
			ID:            "id",
			WhenAccessing: []string{"string"},
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

func TestZoneApplicationDependencyRemove(t *testing.T) {
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
	err := client.Zones.Applications.Dependencies.Remove(
		context.TODO(),
		"dependencyId",
		keycardapi.ZoneApplicationDependencyRemoveParams{
			ZoneID: "zoneId",
			ID:     "id",
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
