# DenialFiltersState

Optional current episode state to match.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialFiltersStateSpendBlockStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialFiltersState("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `DenialFiltersStateSpendBlockStateUnspecified` | SPEND_BLOCK_STATE_UNSPECIFIED                  |
| `DenialFiltersStateSpendBlockStateOpen`        | SPEND_BLOCK_STATE_OPEN                         |
| `DenialFiltersStateSpendBlockStateClosed`      | SPEND_BLOCK_STATE_CLOSED                       |