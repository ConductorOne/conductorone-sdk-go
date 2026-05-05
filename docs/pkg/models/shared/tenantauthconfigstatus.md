# TenantAuthConfigStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TenantAuthConfigStatusAuthConfigStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TenantAuthConfigStatus("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `TenantAuthConfigStatusAuthConfigStatusUnspecified` | AUTH_CONFIG_STATUS_UNSPECIFIED                      |
| `TenantAuthConfigStatusAuthConfigStatusActive`      | AUTH_CONFIG_STATUS_ACTIVE                           |
| `TenantAuthConfigStatusAuthConfigStatusDeprecated`  | AUTH_CONFIG_STATUS_DEPRECATED                       |
| `TenantAuthConfigStatusAuthConfigStatusDisabled`    | AUTH_CONFIG_STATUS_DISABLED                         |