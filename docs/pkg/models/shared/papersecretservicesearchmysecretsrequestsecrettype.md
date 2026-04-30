# PaperSecretServiceSearchMySecretsRequestSecretType

Filter by secret type (optional)

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceSearchMySecretsRequestSecretTypeSecretTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceSearchMySecretsRequestSecretType("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `PaperSecretServiceSearchMySecretsRequestSecretTypeSecretTypeUnspecified` | SECRET_TYPE_UNSPECIFIED                                                   |
| `PaperSecretServiceSearchMySecretsRequestSecretTypeSecretTypeText`        | SECRET_TYPE_TEXT                                                          |
| `PaperSecretServiceSearchMySecretsRequestSecretTypeSecretTypeFile`        | SECRET_TYPE_FILE                                                          |