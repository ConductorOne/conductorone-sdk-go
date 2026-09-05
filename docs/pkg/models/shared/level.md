# Level

The severity of this finding.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.LevelLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Level("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `LevelLevelUnspecified` | LEVEL_UNSPECIFIED       |
| `LevelLevelBlocking`    | LEVEL_BLOCKING          |
| `LevelLevelWarning`     | LEVEL_WARNING           |