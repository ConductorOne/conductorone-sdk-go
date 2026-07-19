# UpdateOnboardingSettingsResponseMcpOnboardingStatus

The updated AIAM MCP onboarding status.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UpdateOnboardingSettingsResponseMcpOnboardingStatusMcpOnboardingStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UpdateOnboardingSettingsResponseMcpOnboardingStatus("custom_value")
```


## Values

| Name                                                                                | Value                                                                               |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `UpdateOnboardingSettingsResponseMcpOnboardingStatusMcpOnboardingStatusUnspecified` | MCP_ONBOARDING_STATUS_UNSPECIFIED                                                   |
| `UpdateOnboardingSettingsResponseMcpOnboardingStatusMcpOnboardingStatusNotStarted`  | MCP_ONBOARDING_STATUS_NOT_STARTED                                                   |
| `UpdateOnboardingSettingsResponseMcpOnboardingStatusMcpOnboardingStatusInProgress`  | MCP_ONBOARDING_STATUS_IN_PROGRESS                                                   |
| `UpdateOnboardingSettingsResponseMcpOnboardingStatusMcpOnboardingStatusComplete`    | MCP_ONBOARDING_STATUS_COMPLETE                                                      |
| `UpdateOnboardingSettingsResponseMcpOnboardingStatusMcpOnboardingStatusDismissed`   | MCP_ONBOARDING_STATUS_DISMISSED                                                     |