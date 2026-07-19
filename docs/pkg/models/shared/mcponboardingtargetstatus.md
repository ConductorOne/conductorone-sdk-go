# McpOnboardingTargetStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.McpOnboardingTargetStatusMcpOnboardingTargetStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.McpOnboardingTargetStatus("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `McpOnboardingTargetStatusMcpOnboardingTargetStatusUnspecified` | MCP_ONBOARDING_TARGET_STATUS_UNSPECIFIED                        |
| `McpOnboardingTargetStatusMcpOnboardingTargetStatusPending`     | MCP_ONBOARDING_TARGET_STATUS_PENDING                            |
| `McpOnboardingTargetStatusMcpOnboardingTargetStatusDone`        | MCP_ONBOARDING_TARGET_STATUS_DONE                               |
| `McpOnboardingTargetStatusMcpOnboardingTargetStatusSkipped`     | MCP_ONBOARDING_TARGET_STATUS_SKIPPED                            |