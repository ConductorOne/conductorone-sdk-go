# FunctionNetworkConfigMode

The mode field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FunctionNetworkConfigModeFunctionNetworkModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FunctionNetworkConfigMode("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `FunctionNetworkConfigModeFunctionNetworkModeUnspecified` | FUNCTION_NETWORK_MODE_UNSPECIFIED                         |
| `FunctionNetworkConfigModeFunctionNetworkModeLegacy`      | FUNCTION_NETWORK_MODE_LEGACY                              |
| `FunctionNetworkConfigModeFunctionNetworkModeEdge`        | FUNCTION_NETWORK_MODE_EDGE                                |