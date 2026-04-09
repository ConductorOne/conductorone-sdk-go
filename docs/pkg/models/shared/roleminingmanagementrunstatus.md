# RoleMiningManagementRunStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RoleMiningManagementRunStatusRunStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RoleMiningManagementRunStatus("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `RoleMiningManagementRunStatusRunStatusUnspecified` | RUN_STATUS_UNSPECIFIED                              |
| `RoleMiningManagementRunStatusRunStatusRunning`     | RUN_STATUS_RUNNING                                  |
| `RoleMiningManagementRunStatusRunStatusCompleted`   | RUN_STATUS_COMPLETED                                |
| `RoleMiningManagementRunStatusRunStatusFailed`      | RUN_STATUS_FAILED                                   |