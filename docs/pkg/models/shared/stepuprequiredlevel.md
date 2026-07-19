# StepUpRequiredLevel

The assurance level the step-up must reach.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StepUpRequiredLevelAuthLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StepUpRequiredLevel("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `StepUpRequiredLevelAuthLevelUnspecified`  | AUTH_LEVEL_UNSPECIFIED                     |
| `StepUpRequiredLevelAuthLevelNone`         | AUTH_LEVEL_NONE                            |
| `StepUpRequiredLevelAuthLevelSingleFactor` | AUTH_LEVEL_SINGLE_FACTOR                   |
| `StepUpRequiredLevelAuthLevelMultiFactor`  | AUTH_LEVEL_MULTI_FACTOR                    |
| `StepUpRequiredLevelAuthLevelPhr`          | AUTH_LEVEL_PHR                             |
| `StepUpRequiredLevelAuthLevelPhrh`         | AUTH_LEVEL_PHRH                            |