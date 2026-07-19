# UpdateOnboardingSettingsRequestStatus

The new onboarding status to set. UNSPECIFIED leaves the core onboarding
 status unchanged (set mcp_onboarding_status alone to retire the AIAM
 briefing without touching the core wizard).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UpdateOnboardingSettingsRequestStatusOnboardingStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UpdateOnboardingSettingsRequestStatus("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `UpdateOnboardingSettingsRequestStatusOnboardingStatusUnspecified` | ONBOARDING_STATUS_UNSPECIFIED                                      |
| `UpdateOnboardingSettingsRequestStatusOnboardingStatusNotStarted`  | ONBOARDING_STATUS_NOT_STARTED                                      |
| `UpdateOnboardingSettingsRequestStatusOnboardingStatusInProgress`  | ONBOARDING_STATUS_IN_PROGRESS                                      |
| `UpdateOnboardingSettingsRequestStatusOnboardingStatusComplete`    | ONBOARDING_STATUS_COMPLETE                                         |
| `UpdateOnboardingSettingsRequestStatusOnboardingStatusDismissed`   | ONBOARDING_STATUS_DISMISSED                                        |