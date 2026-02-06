// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/keycard-api-go"
	"github.com/stainless-sdks/keycard-api-go/internal/testutil"
	"github.com/stainless-sdks/keycard-api-go/option"
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
	)
	t.Skip("Prism tests are disabled")
	zones, err := client.Zones.List(context.TODO(), keycardapi.ZoneListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", zones.Items)
}
