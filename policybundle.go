// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"github.com/keycardai/keycard-go/option"
)

// PolicyBundleService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyBundleService] method instead.
type PolicyBundleService struct {
	Options []option.RequestOption
}

// NewPolicyBundleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPolicyBundleService(opts ...option.RequestOption) (r PolicyBundleService) {
	r = PolicyBundleService{}
	r.Options = opts
	return
}
