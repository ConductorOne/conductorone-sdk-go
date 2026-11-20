# OAuthAuthorizationServiceCreateRequestAccessMode

The accessMode field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OAuthAuthorizationServiceCreateRequestAccessModeOAuthAuthorizationAccessModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.OAuthAuthorizationServiceCreateRequestAccessMode("custom_value")
```


## Values

| Name                                                                                         | Value                                                                                        |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `OAuthAuthorizationServiceCreateRequestAccessModeOAuthAuthorizationAccessModeUnspecified`    | O_AUTH_AUTHORIZATION_ACCESS_MODE_UNSPECIFIED                                                 |
| `OAuthAuthorizationServiceCreateRequestAccessModeOAuthAuthorizationAccessModeAllTenantUsers` | O_AUTH_AUTHORIZATION_ACCESS_MODE_ALL_TENANT_USERS                                            |
| `OAuthAuthorizationServiceCreateRequestAccessModeOAuthAuthorizationAccessModeEntitledUsers`  | O_AUTH_AUTHORIZATION_ACCESS_MODE_ENTITLED_USERS                                              |