# PaperSecretServiceCreateExternalRequestInputFormat

For TEXT secrets: hint about the plaintext format (e.g., JSON, YAML, key-value).
 Used by the viewer UI for syntax highlighting. Does not affect encryption.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceCreateExternalRequestInputFormat("custom_value")
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatUnspecified` | SECRET_INPUT_FORMAT_UNSPECIFIED                                                  |
| `PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatPlaintext`   | SECRET_INPUT_FORMAT_PLAINTEXT                                                    |
| `PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatJSON`        | SECRET_INPUT_FORMAT_JSON                                                         |
| `PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatYaml`        | SECRET_INPUT_FORMAT_YAML                                                         |
| `PaperSecretServiceCreateExternalRequestInputFormatSecretInputFormatKeyValue`    | SECRET_INPUT_FORMAT_KEY_VALUE                                                    |