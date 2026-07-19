# SessionPolicyAllowFloorLevel

The minimum assurance level that satisfies this rule.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyAllowFloorLevelAuthLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyAllowFloorLevel("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `SessionPolicyAllowFloorLevelAuthLevelUnspecified`  | AUTH_LEVEL_UNSPECIFIED                              |
| `SessionPolicyAllowFloorLevelAuthLevelNone`         | AUTH_LEVEL_NONE                                     |
| `SessionPolicyAllowFloorLevelAuthLevelSingleFactor` | AUTH_LEVEL_SINGLE_FACTOR                            |
| `SessionPolicyAllowFloorLevelAuthLevelMultiFactor`  | AUTH_LEVEL_MULTI_FACTOR                             |
| `SessionPolicyAllowFloorLevelAuthLevelPhr`          | AUTH_LEVEL_PHR                                      |
| `SessionPolicyAllowFloorLevelAuthLevelPhrh`         | AUTH_LEVEL_PHRH                                     |