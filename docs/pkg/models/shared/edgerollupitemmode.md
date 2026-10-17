# EdgeRollupItemMode

The mode field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeRollupItemModeEdgeModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeRollupItemMode("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `EdgeRollupItemModeEdgeModeUnspecified` | EDGE_MODE_UNSPECIFIED                   |
| `EdgeRollupItemModeEdgeModeDisabled`    | EDGE_MODE_DISABLED                      |
| `EdgeRollupItemModeEdgeModeObserve`     | EDGE_MODE_OBSERVE                       |
| `EdgeRollupItemModeEdgeModeEnforce`     | EDGE_MODE_ENFORCE                       |