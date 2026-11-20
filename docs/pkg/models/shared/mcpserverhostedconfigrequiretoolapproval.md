# ~~MCPServerHostedConfigRequireToolApproval~~

Deprecated and ignored: set tool_enablement on the register request instead.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerHostedConfigRequireToolApprovalOptionalBoolUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerHostedConfigRequireToolApproval("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `MCPServerHostedConfigRequireToolApprovalOptionalBoolUnspecified` | OPTIONAL_BOOL_UNSPECIFIED                                         |
| `MCPServerHostedConfigRequireToolApprovalOptionalBoolTrue`        | OPTIONAL_BOOL_TRUE                                                |
| `MCPServerHostedConfigRequireToolApprovalOptionalBoolFalse`       | OPTIONAL_BOOL_FALSE                                               |