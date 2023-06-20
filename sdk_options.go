package conductoroneapi

import (
	"net/url"
	"strings"
)

// WithTenantSmart takes a URL, tenant name, or full tenant URL and sets the tenant domain for the SDk
// Valid input examples: insulator.conductor.one, insulator, https://insulator.conductor.one,
func WithTenantSmart(tenant string) SDKOption {
	foundUrl, err := url.Parse(tenant)
	if err == nil {
		if !strings.HasPrefix(tenant, "http") {
			foundUrl.Scheme = "https://" + foundUrl.Scheme
		}
		return WithServerURL(foundUrl.String())
	}

	return WithTenantDomain(tenant)
}
