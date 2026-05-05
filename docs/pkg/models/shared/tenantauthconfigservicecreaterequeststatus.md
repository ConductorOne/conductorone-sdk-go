# TenantAuthConfigServiceCreateRequestStatus

The initial status of the authentication provider.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TenantAuthConfigServiceCreateRequestStatusAuthConfigStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TenantAuthConfigServiceCreateRequestStatus("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `TenantAuthConfigServiceCreateRequestStatusAuthConfigStatusUnspecified` | AUTH_CONFIG_STATUS_UNSPECIFIED                                          |
| `TenantAuthConfigServiceCreateRequestStatusAuthConfigStatusActive`      | AUTH_CONFIG_STATUS_ACTIVE                                               |
| `TenantAuthConfigServiceCreateRequestStatusAuthConfigStatusDeprecated`  | AUTH_CONFIG_STATUS_DEPRECATED                                           |
| `TenantAuthConfigServiceCreateRequestStatusAuthConfigStatusDisabled`    | AUTH_CONFIG_STATUS_DISABLED                                             |