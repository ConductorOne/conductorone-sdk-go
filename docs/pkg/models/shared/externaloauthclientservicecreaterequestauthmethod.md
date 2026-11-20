# ExternalOAuthClientServiceCreateRequestAuthMethod

The authMethod field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExternalOAuthClientServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExternalOAuthClientServiceCreateRequestAuthMethod("custom_value")
```


## Values

| Name                                                                                              | Value                                                                                             |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ExternalOAuthClientServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodUnspecified`       | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_UNSPECIFIED                                                    |
| `ExternalOAuthClientServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodClientSecretBasic` | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_CLIENT_SECRET_BASIC                                            |
| `ExternalOAuthClientServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodClientSecretPost`  | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_CLIENT_SECRET_POST                                             |
| `ExternalOAuthClientServiceCreateRequestAuthMethodExternalOAuthClientAuthMethodPrivateKeyJwt`     | EXTERNAL_O_AUTH_CLIENT_AUTH_METHOD_PRIVATE_KEY_JWT                                                |