# InputFormat

The inputFormat field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.InputFormatSecretInputFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.InputFormat("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `InputFormatSecretInputFormatUnspecified` | SECRET_INPUT_FORMAT_UNSPECIFIED           |
| `InputFormatSecretInputFormatPlaintext`   | SECRET_INPUT_FORMAT_PLAINTEXT             |
| `InputFormatSecretInputFormatJSON`        | SECRET_INPUT_FORMAT_JSON                  |
| `InputFormatSecretInputFormatYaml`        | SECRET_INPUT_FORMAT_YAML                  |
| `InputFormatSecretInputFormatKeyValue`    | SECRET_INPUT_FORMAT_KEY_VALUE             |