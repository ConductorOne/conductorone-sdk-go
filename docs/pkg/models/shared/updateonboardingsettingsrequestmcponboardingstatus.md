# UpdateOnboardingSettingsRequestMcpOnboardingStatus

The new MCP onboarding status to set. Omit (or UNSPECIFIED) to leave it
 unchanged. Setting NOT_STARTED restarts the briefing and clears the stored
 mcp_onboarding_goal and mcp_onboarding_targets, unless this same request also
 sets them (those win).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UpdateOnboardingSettingsRequestMcpOnboardingStatusMcpOnboardingStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UpdateOnboardingSettingsRequestMcpOnboardingStatus("custom_value")
```


## Values

| Name                                                                               | Value                                                                              |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `UpdateOnboardingSettingsRequestMcpOnboardingStatusMcpOnboardingStatusUnspecified` | MCP_ONBOARDING_STATUS_UNSPECIFIED                                                  |
| `UpdateOnboardingSettingsRequestMcpOnboardingStatusMcpOnboardingStatusNotStarted`  | MCP_ONBOARDING_STATUS_NOT_STARTED                                                  |
| `UpdateOnboardingSettingsRequestMcpOnboardingStatusMcpOnboardingStatusInProgress`  | MCP_ONBOARDING_STATUS_IN_PROGRESS                                                  |
| `UpdateOnboardingSettingsRequestMcpOnboardingStatusMcpOnboardingStatusComplete`    | MCP_ONBOARDING_STATUS_COMPLETE                                                     |
| `UpdateOnboardingSettingsRequestMcpOnboardingStatusMcpOnboardingStatusDismissed`   | MCP_ONBOARDING_STATUS_DISMISSED                                                    |