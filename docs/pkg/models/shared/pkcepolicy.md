# PkcePolicy

Effective PKCE policy.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PkcePolicySsoApplicationOidcPkcePolicyUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PkcePolicy("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `PkcePolicySsoApplicationOidcPkcePolicyUnspecified`           | SSO_APPLICATION_OIDC_PKCE_POLICY_UNSPECIFIED                  |
| `PkcePolicySsoApplicationOidcPkcePolicyRequiredS256`          | SSO_APPLICATION_OIDC_PKCE_POLICY_REQUIRED_S256                |
| `PkcePolicySsoApplicationOidcPkcePolicyAllowMissingForLegacy` | SSO_APPLICATION_OIDC_PKCE_POLICY_ALLOW_MISSING_FOR_LEGACY     |