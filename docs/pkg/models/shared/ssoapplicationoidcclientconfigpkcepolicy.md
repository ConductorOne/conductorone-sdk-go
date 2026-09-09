# SSOApplicationOIDCClientConfigPkcePolicy

PKCE is required by default on create. On update, UNSPECIFIED preserves
 the current policy; set REQUIRED_S256 explicitly to tighten a legacy
 confidential client.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSOApplicationOIDCClientConfigPkcePolicySsoApplicationOidcPkcePolicyUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSOApplicationOIDCClientConfigPkcePolicy("custom_value")
```


## Values

| Name                                                                                        | Value                                                                                       |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `SSOApplicationOIDCClientConfigPkcePolicySsoApplicationOidcPkcePolicyUnspecified`           | SSO_APPLICATION_OIDC_PKCE_POLICY_UNSPECIFIED                                                |
| `SSOApplicationOIDCClientConfigPkcePolicySsoApplicationOidcPkcePolicyRequiredS256`          | SSO_APPLICATION_OIDC_PKCE_POLICY_REQUIRED_S256                                              |
| `SSOApplicationOIDCClientConfigPkcePolicySsoApplicationOidcPkcePolicyAllowMissingForLegacy` | SSO_APPLICATION_OIDC_PKCE_POLICY_ALLOW_MISSING_FOR_LEGACY                                   |