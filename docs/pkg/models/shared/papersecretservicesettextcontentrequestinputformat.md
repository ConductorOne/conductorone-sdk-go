# PaperSecretServiceSetTextContentRequestInputFormat

Input format hint for the viewer UI when the secret is decrypted.
 Does not affect encryption — this is metadata only.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceSetTextContentRequestInputFormat("custom_value")
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatUnspecified` | SECRET_INPUT_FORMAT_UNSPECIFIED                                                  |
| `PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatPlaintext`   | SECRET_INPUT_FORMAT_PLAINTEXT                                                    |
| `PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatJSON`        | SECRET_INPUT_FORMAT_JSON                                                         |
| `PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatYaml`        | SECRET_INPUT_FORMAT_YAML                                                         |
| `PaperSecretServiceSetTextContentRequestInputFormatSecretInputFormatKeyValue`    | SECRET_INPUT_FORMAT_KEY_VALUE                                                    |