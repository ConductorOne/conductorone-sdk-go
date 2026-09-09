# SetProviderCredentialRequestHeaderStyle

The headerStyle field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SetProviderCredentialRequestHeaderStyleProviderCredentialHeaderStyleUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SetProviderCredentialRequestHeaderStyle("custom_value")
```


## Values

| Name                                                                                      | Value                                                                                     |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `SetProviderCredentialRequestHeaderStyleProviderCredentialHeaderStyleUnspecified`         | PROVIDER_CREDENTIAL_HEADER_STYLE_UNSPECIFIED                                              |
| `SetProviderCredentialRequestHeaderStyleProviderCredentialHeaderStyleXAPIKey`             | PROVIDER_CREDENTIAL_HEADER_STYLE_X_API_KEY                                                |
| `SetProviderCredentialRequestHeaderStyleProviderCredentialHeaderStyleAuthorizationBearer` | PROVIDER_CREDENTIAL_HEADER_STYLE_AUTHORIZATION_BEARER                                     |