# DenialEpisodeState

Whether this denial episode remains open.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialEpisodeStateSpendBlockStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialEpisodeState("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `DenialEpisodeStateSpendBlockStateUnspecified` | SPEND_BLOCK_STATE_UNSPECIFIED                  |
| `DenialEpisodeStateSpendBlockStateOpen`        | SPEND_BLOCK_STATE_OPEN                         |
| `DenialEpisodeStateSpendBlockStateClosed`      | SPEND_BLOCK_STATE_CLOSED                       |