# TenantAuthConfigProviderType

Provider type (read-only after creation — provider config determines type).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TenantAuthConfigProviderTypeAuthConfigProviderTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TenantAuthConfigProviderType("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeUnspecified` | AUTH_CONFIG_PROVIDER_TYPE_UNSPECIFIED                           |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeGoogle`      | AUTH_CONFIG_PROVIDER_TYPE_GOOGLE                                |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeMicrosoft`   | AUTH_CONFIG_PROVIDER_TYPE_MICROSOFT                             |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeOkta`        | AUTH_CONFIG_PROVIDER_TYPE_OKTA                                  |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeOnelogin`    | AUTH_CONFIG_PROVIDER_TYPE_ONELOGIN                              |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeJumpcloud`   | AUTH_CONFIG_PROVIDER_TYPE_JUMPCLOUD                             |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypePingone`     | AUTH_CONFIG_PROVIDER_TYPE_PINGONE                               |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeOidc`        | AUTH_CONFIG_PROVIDER_TYPE_OIDC                                  |
| `TenantAuthConfigProviderTypeAuthConfigProviderTypeC1Local`     | AUTH_CONFIG_PROVIDER_TYPE_C1_LOCAL                              |