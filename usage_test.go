// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/keycardai/keycard-go"
	"github.com/keycardai/keycard-go/internal/testutil"
	"github.com/keycardai/keycard-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycardapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	t.Skip("Mock server tests are disabled")
	zones, err := client.Zones.List(context.TODO(), keycardapi.ZoneListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", zones.Items)
}
