// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"github.com/keycardai/keycard-go/option"
)

// ZoneCatalogTaskService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCatalogTaskService] method instead.
type ZoneCatalogTaskService struct {
	Options []option.RequestOption
}

// NewZoneCatalogTaskService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCatalogTaskService(opts ...option.RequestOption) (r ZoneCatalogTaskService) {
	r = ZoneCatalogTaskService{}
	r.Options = opts
	return
}
