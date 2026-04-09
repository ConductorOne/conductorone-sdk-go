# PaperSecretServiceCreateInternalRequestInputFormat

For TEXT secrets: hint about the plaintext format (e.g., JSON, YAML, key-value).
 Used by the viewer UI for syntax highlighting. Does not affect encryption.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceCreateInternalRequestInputFormat("custom_value")
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatUnspecified` | SECRET_INPUT_FORMAT_UNSPECIFIED                                                  |
| `PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatPlaintext`   | SECRET_INPUT_FORMAT_PLAINTEXT                                                    |
| `PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatJSON`        | SECRET_INPUT_FORMAT_JSON                                                         |
| `PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatYaml`        | SECRET_INPUT_FORMAT_YAML                                                         |
| `PaperSecretServiceCreateInternalRequestInputFormatSecretInputFormatKeyValue`    | SECRET_INPUT_FORMAT_KEY_VALUE                                                    |