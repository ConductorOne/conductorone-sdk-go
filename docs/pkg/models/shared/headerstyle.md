# HeaderStyle

The headerStyle field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.HeaderStyleProviderCredentialHeaderStyleUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.HeaderStyle("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `HeaderStyleProviderCredentialHeaderStyleUnspecified`         | PROVIDER_CREDENTIAL_HEADER_STYLE_UNSPECIFIED                  |
| `HeaderStyleProviderCredentialHeaderStyleXAPIKey`             | PROVIDER_CREDENTIAL_HEADER_STYLE_X_API_KEY                    |
| `HeaderStyleProviderCredentialHeaderStyleAuthorizationBearer` | PROVIDER_CREDENTIAL_HEADER_STYLE_AUTHORIZATION_BEARER         |