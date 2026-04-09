# GetOnboardingSettingsResponseStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GetOnboardingSettingsResponseStatusOnboardingStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GetOnboardingSettingsResponseStatus("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `GetOnboardingSettingsResponseStatusOnboardingStatusUnspecified` | ONBOARDING_STATUS_UNSPECIFIED                                    |
| `GetOnboardingSettingsResponseStatusOnboardingStatusNotStarted`  | ONBOARDING_STATUS_NOT_STARTED                                    |
| `GetOnboardingSettingsResponseStatusOnboardingStatusInProgress`  | ONBOARDING_STATUS_IN_PROGRESS                                    |
| `GetOnboardingSettingsResponseStatusOnboardingStatusComplete`    | ONBOARDING_STATUS_COMPLETE                                       |
| `GetOnboardingSettingsResponseStatusOnboardingStatusDismissed`   | ONBOARDING_STATUS_DISMISSED                                      |