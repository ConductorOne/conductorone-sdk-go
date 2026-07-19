# FloorLevel

The minimum assurance level that satisfies this rule. Required on enforced
 Allow rules.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FloorLevelAuthLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FloorLevel("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `FloorLevelAuthLevelUnspecified`  | AUTH_LEVEL_UNSPECIFIED            |
| `FloorLevelAuthLevelNone`         | AUTH_LEVEL_NONE                   |
| `FloorLevelAuthLevelSingleFactor` | AUTH_LEVEL_SINGLE_FACTOR          |
| `FloorLevelAuthLevelMultiFactor`  | AUTH_LEVEL_MULTI_FACTOR           |
| `FloorLevelAuthLevelPhr`          | AUTH_LEVEL_PHR                    |
| `FloorLevelAuthLevelPhrh`         | AUTH_LEVEL_PHRH                   |