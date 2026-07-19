# McpOnboardingStatus

The current status of the AIAM MCP onboarding briefing, tracked
 independently of `status`.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.McpOnboardingStatusMcpOnboardingStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.McpOnboardingStatus("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `McpOnboardingStatusMcpOnboardingStatusUnspecified` | MCP_ONBOARDING_STATUS_UNSPECIFIED                   |
| `McpOnboardingStatusMcpOnboardingStatusNotStarted`  | MCP_ONBOARDING_STATUS_NOT_STARTED                   |
| `McpOnboardingStatusMcpOnboardingStatusInProgress`  | MCP_ONBOARDING_STATUS_IN_PROGRESS                   |
| `McpOnboardingStatusMcpOnboardingStatusComplete`    | MCP_ONBOARDING_STATUS_COMPLETE                      |
| `McpOnboardingStatusMcpOnboardingStatusDismissed`   | MCP_ONBOARDING_STATUS_DISMISSED                     |