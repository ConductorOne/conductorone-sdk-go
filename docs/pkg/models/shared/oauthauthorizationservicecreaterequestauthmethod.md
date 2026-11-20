# OAuthAuthorizationServiceCreateRequestAuthMethod

The authMethod field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OAuthAuthorizationServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.OAuthAuthorizationServiceCreateRequestAuthMethod("custom_value")
```


## Values

| Name                                                                                             | Value                                                                                            |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `OAuthAuthorizationServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodUnspecified`       | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_UNSPECIFIED                                                   |
| `OAuthAuthorizationServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodClientSecretBasic` | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_CLIENT_SECRET_BASIC                                           |
| `OAuthAuthorizationServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodClientSecretPost`  | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_CLIENT_SECRET_POST                                            |
| `OAuthAuthorizationServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodPrivateKeyJwt`     | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_PRIVATE_KEY_JWT                                               |