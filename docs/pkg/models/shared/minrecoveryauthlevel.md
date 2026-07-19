# MinRecoveryAuthLevel

The minimum assurance level a recovery ceremony must reach for this policy.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MinRecoveryAuthLevelAuthLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MinRecoveryAuthLevel("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `MinRecoveryAuthLevelAuthLevelUnspecified`  | AUTH_LEVEL_UNSPECIFIED                      |
| `MinRecoveryAuthLevelAuthLevelNone`         | AUTH_LEVEL_NONE                             |
| `MinRecoveryAuthLevelAuthLevelSingleFactor` | AUTH_LEVEL_SINGLE_FACTOR                    |
| `MinRecoveryAuthLevelAuthLevelMultiFactor`  | AUTH_LEVEL_MULTI_FACTOR                     |
| `MinRecoveryAuthLevelAuthLevelPhr`          | AUTH_LEVEL_PHR                              |
| `MinRecoveryAuthLevelAuthLevelPhrh`         | AUTH_LEVEL_PHRH                             |