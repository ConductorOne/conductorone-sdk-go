# UpdateOnboardingSettingsResponseStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UpdateOnboardingSettingsResponseStatusOnboardingStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UpdateOnboardingSettingsResponseStatus("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `UpdateOnboardingSettingsResponseStatusOnboardingStatusUnspecified` | ONBOARDING_STATUS_UNSPECIFIED                                       |
| `UpdateOnboardingSettingsResponseStatusOnboardingStatusNotStarted`  | ONBOARDING_STATUS_NOT_STARTED                                       |
| `UpdateOnboardingSettingsResponseStatusOnboardingStatusInProgress`  | ONBOARDING_STATUS_IN_PROGRESS                                       |
| `UpdateOnboardingSettingsResponseStatusOnboardingStatusComplete`    | ONBOARDING_STATUS_COMPLETE                                          |
| `UpdateOnboardingSettingsResponseStatusOnboardingStatusDismissed`   | ONBOARDING_STATUS_DISMISSED                                         |