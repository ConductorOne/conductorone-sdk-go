# EdgeInferenceRouteNativeFormat

Accepted native client request format; never inferred from caller headers.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeInferenceRouteNativeFormatEdgeInferenceNativeFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeInferenceRouteNativeFormat("custom_value")
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `EdgeInferenceRouteNativeFormatEdgeInferenceNativeFormatUnspecified`       | EDGE_INFERENCE_NATIVE_FORMAT_UNSPECIFIED                                   |
| `EdgeInferenceRouteNativeFormatEdgeInferenceNativeFormatOpenaiResponses`   | EDGE_INFERENCE_NATIVE_FORMAT_OPENAI_RESPONSES                              |
| `EdgeInferenceRouteNativeFormatEdgeInferenceNativeFormatOpenaiChat`        | EDGE_INFERENCE_NATIVE_FORMAT_OPENAI_CHAT                                   |
| `EdgeInferenceRouteNativeFormatEdgeInferenceNativeFormatAnthropicMessages` | EDGE_INFERENCE_NATIVE_FORMAT_ANTHROPIC_MESSAGES                            |