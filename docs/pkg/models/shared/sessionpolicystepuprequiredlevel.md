# SessionPolicyStepUpRequiredLevel

The level field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyStepUpRequiredLevelAuthLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyStepUpRequiredLevel("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `SessionPolicyStepUpRequiredLevelAuthLevelUnspecified`  | AUTH_LEVEL_UNSPECIFIED                                  |
| `SessionPolicyStepUpRequiredLevelAuthLevelNone`         | AUTH_LEVEL_NONE                                         |
| `SessionPolicyStepUpRequiredLevelAuthLevelSingleFactor` | AUTH_LEVEL_SINGLE_FACTOR                                |
| `SessionPolicyStepUpRequiredLevelAuthLevelMultiFactor`  | AUTH_LEVEL_MULTI_FACTOR                                 |
| `SessionPolicyStepUpRequiredLevelAuthLevelPhr`          | AUTH_LEVEL_PHR                                          |
| `SessionPolicyStepUpRequiredLevelAuthLevelPhrh`         | AUTH_LEVEL_PHRH                                         |