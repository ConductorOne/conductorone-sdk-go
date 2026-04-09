# SecretType

The secretType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SecretTypeSecretTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SecretType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `SecretTypeSecretTypeUnspecified` | SECRET_TYPE_UNSPECIFIED           |
| `SecretTypeSecretTypeText`        | SECRET_TYPE_TEXT                  |
| `SecretTypeSecretTypeFile`        | SECRET_TYPE_FILE                  |