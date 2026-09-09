# MySpendBlockState

Whether this denial episode remains open.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MySpendBlockStateSpendBlockStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MySpendBlockState("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `MySpendBlockStateSpendBlockStateUnspecified` | SPEND_BLOCK_STATE_UNSPECIFIED                 |
| `MySpendBlockStateSpendBlockStateOpen`        | SPEND_BLOCK_STATE_OPEN                        |
| `MySpendBlockStateSpendBlockStateClosed`      | SPEND_BLOCK_STATE_CLOSED                      |