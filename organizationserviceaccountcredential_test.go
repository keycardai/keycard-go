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

func TestOrganizationServiceAccountCredentialNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Organizations.ServiceAccounts.Credentials.New(
		context.TODO(),
		"ab3def8hij2klm9opq5rst7uvw",
		keycardapi.OrganizationServiceAccountCredentialNewParams{
			OrganizationID:   "x",
			Name:             "name",
			Description:      keycardapi.String("description"),
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			XRequestID:       keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestOrganizationServiceAccountCredentialGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Organizations.ServiceAccounts.Credentials.Get(
		context.TODO(),
		"ab3def8hij2klm9opq5rst7uvw",
		keycardapi.OrganizationServiceAccountCredentialGetParams{
			OrganizationID:   "x",
			ServiceAccountID: "ab3def8hij2klm9opq5rst7uvw",
			Expand:           []string{"permissions"},
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			XRequestID:       keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestOrganizationServiceAccountCredentialUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Organizations.ServiceAccounts.Credentials.Update(
		context.TODO(),
		"ab3def8hij2klm9opq5rst7uvw",
		keycardapi.OrganizationServiceAccountCredentialUpdateParams{
			OrganizationID:   "x",
			ServiceAccountID: "ab3def8hij2klm9opq5rst7uvw",
			Description:      keycardapi.String("description"),
			Name:             keycardapi.String("name"),
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			XRequestID:       keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestOrganizationServiceAccountCredentialListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Organizations.ServiceAccounts.Credentials.List(
		context.TODO(),
		"ab3def8hij2klm9opq5rst7uvw",
		keycardapi.OrganizationServiceAccountCredentialListParams{
			OrganizationID:   "x",
			After:            keycardapi.String("x"),
			Before:           keycardapi.String("x"),
			Expand:           []string{"permissions"},
			Limit:            keycardapi.Int(1),
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			XRequestID:       keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestOrganizationServiceAccountCredentialDeleteWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	err := client.Organizations.ServiceAccounts.Credentials.Delete(
		context.TODO(),
		"ab3def8hij2klm9opq5rst7uvw",
		keycardapi.OrganizationServiceAccountCredentialDeleteParams{
			OrganizationID:   "x",
			ServiceAccountID: "ab3def8hij2klm9opq5rst7uvw",
			XClientRequestID: keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			XRequestID:       keycardapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
