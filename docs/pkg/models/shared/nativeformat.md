# NativeFormat

Client format the variant serves.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.NativeFormatEdgeInferenceNativeFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.NativeFormat("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `NativeFormatEdgeInferenceNativeFormatUnspecified`       | EDGE_INFERENCE_NATIVE_FORMAT_UNSPECIFIED                 |
| `NativeFormatEdgeInferenceNativeFormatOpenaiResponses`   | EDGE_INFERENCE_NATIVE_FORMAT_OPENAI_RESPONSES            |
| `NativeFormatEdgeInferenceNativeFormatOpenaiChat`        | EDGE_INFERENCE_NATIVE_FORMAT_OPENAI_CHAT                 |
| `NativeFormatEdgeInferenceNativeFormatAnthropicMessages` | EDGE_INFERENCE_NATIVE_FORMAT_ANTHROPIC_MESSAGES          |