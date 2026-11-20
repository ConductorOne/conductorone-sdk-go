# SyncStatus

Current connector sync status.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SyncStatusSyncStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SyncStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `SyncStatusSyncStatusUnspecified` | SYNC_STATUS_UNSPECIFIED           |
| `SyncStatusSyncStatusRunning`     | SYNC_STATUS_RUNNING               |
| `SyncStatusSyncStatusDone`        | SYNC_STATUS_DONE                  |
| `SyncStatusSyncStatusError`       | SYNC_STATUS_ERROR                 |
| `SyncStatusSyncStatusDisabled`    | SYNC_STATUS_DISABLED              |