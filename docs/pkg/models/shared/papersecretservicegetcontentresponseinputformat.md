# PaperSecretServiceGetContentResponseInputFormat

Input format hint for rendering (text secrets only)

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceGetContentResponseInputFormatSecretInputFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceGetContentResponseInputFormat("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `PaperSecretServiceGetContentResponseInputFormatSecretInputFormatUnspecified` | SECRET_INPUT_FORMAT_UNSPECIFIED                                               |
| `PaperSecretServiceGetContentResponseInputFormatSecretInputFormatPlaintext`   | SECRET_INPUT_FORMAT_PLAINTEXT                                                 |
| `PaperSecretServiceGetContentResponseInputFormatSecretInputFormatJSON`        | SECRET_INPUT_FORMAT_JSON                                                      |
| `PaperSecretServiceGetContentResponseInputFormatSecretInputFormatYaml`        | SECRET_INPUT_FORMAT_YAML                                                      |
| `PaperSecretServiceGetContentResponseInputFormatSecretInputFormatKeyValue`    | SECRET_INPUT_FORMAT_KEY_VALUE                                                 |