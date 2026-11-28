# Requirement

Connection requirements to match. Unspecified matches any requirements.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequirementVirtualMcpServerRequirementFilterUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Requirement("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `RequirementVirtualMcpServerRequirementFilterUnspecified` | VIRTUAL_MCP_SERVER_REQUIREMENT_FILTER_UNSPECIFIED         |
| `RequirementVirtualMcpServerRequirementFilterNone`        | VIRTUAL_MCP_SERVER_REQUIREMENT_FILTER_NONE                |
| `RequirementVirtualMcpServerRequirementFilterDpopOnly`    | VIRTUAL_MCP_SERVER_REQUIREMENT_FILTER_DPOP_ONLY           |
| `RequirementVirtualMcpServerRequirementFilterDeviceTrust` | VIRTUAL_MCP_SERVER_REQUIREMENT_FILTER_DEVICE_TRUST        |