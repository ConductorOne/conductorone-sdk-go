# AccessMode

The accessMode field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessModeOAuthAuthorizationAccessModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessMode("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AccessModeOAuthAuthorizationAccessModeUnspecified`    | O_AUTH_AUTHORIZATION_ACCESS_MODE_UNSPECIFIED           |
| `AccessModeOAuthAuthorizationAccessModeAllTenantUsers` | O_AUTH_AUTHORIZATION_ACCESS_MODE_ALL_TENANT_USERS      |
| `AccessModeOAuthAuthorizationAccessModeEntitledUsers`  | O_AUTH_AUTHORIZATION_ACCESS_MODE_ENTITLED_USERS        |