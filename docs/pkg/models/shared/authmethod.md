# AuthMethod

Token-endpoint authentication method. Populated for static clients.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthMethodExternalOAuthClientAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthMethod("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AuthMethodExternalOAuthClientAuthMethodUnspecified`       | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_UNSPECIFIED             |
| `AuthMethodExternalOAuthClientAuthMethodClientSecretBasic` | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_CLIENT_SECRET_BASIC     |
| `AuthMethodExternalOAuthClientAuthMethodClientSecretPost`  | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_CLIENT_SECRET_POST      |
| `AuthMethodExternalOAuthClientAuthMethodPrivateKeyJwt`     | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_PRIVATE_KEY_JWT         |