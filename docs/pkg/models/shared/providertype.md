# ProviderType

The providerType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ProviderTypeProviderTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ProviderType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ProviderTypeProviderTypeUnspecified` | PROVIDER_TYPE_UNSPECIFIED             |
| `ProviderTypeProviderTypeOauth2`      | PROVIDER_TYPE_OAUTH2                  |
| `ProviderTypeProviderTypeMicrosoft`   | PROVIDER_TYPE_MICROSOFT               |