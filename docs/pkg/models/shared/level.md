# Level

The level field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.LevelAuthLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Level("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `LevelAuthLevelUnspecified`  | AUTH_LEVEL_UNSPECIFIED       |
| `LevelAuthLevelNone`         | AUTH_LEVEL_NONE              |
| `LevelAuthLevelSingleFactor` | AUTH_LEVEL_SINGLE_FACTOR     |
| `LevelAuthLevelMultiFactor`  | AUTH_LEVEL_MULTI_FACTOR      |
| `LevelAuthLevelPhr`          | AUTH_LEVEL_PHR               |
| `LevelAuthLevelPhrh`         | AUTH_LEVEL_PHRH              |