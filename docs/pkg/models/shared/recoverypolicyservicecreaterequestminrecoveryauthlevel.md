# RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevel

The minimum assurance level a recovery ceremony must reach.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevelAuthLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevel("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevelAuthLevelUnspecified`  | AUTH_LEVEL_UNSPECIFIED                                                        |
| `RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevelAuthLevelNone`         | AUTH_LEVEL_NONE                                                               |
| `RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevelAuthLevelSingleFactor` | AUTH_LEVEL_SINGLE_FACTOR                                                      |
| `RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevelAuthLevelMultiFactor`  | AUTH_LEVEL_MULTI_FACTOR                                                       |
| `RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevelAuthLevelPhr`          | AUTH_LEVEL_PHR                                                                |
| `RecoveryPolicyServiceCreateRequestMinRecoveryAuthLevelAuthLevelPhrh`         | AUTH_LEVEL_PHRH                                                               |