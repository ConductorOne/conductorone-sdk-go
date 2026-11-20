# SyncStatusViewStatus

Current sync status.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SyncStatusViewStatusSyncStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SyncStatusViewStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `SyncStatusViewStatusSyncStatusUnspecified` | SYNC_STATUS_UNSPECIFIED                     |
| `SyncStatusViewStatusSyncStatusRunning`     | SYNC_STATUS_RUNNING                         |
| `SyncStatusViewStatusSyncStatusDone`        | SYNC_STATUS_DONE                            |
| `SyncStatusViewStatusSyncStatusError`       | SYNC_STATUS_ERROR                           |
| `SyncStatusViewStatusSyncStatusDisabled`    | SYNC_STATUS_DISABLED                        |