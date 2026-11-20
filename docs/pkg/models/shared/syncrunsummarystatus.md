# SyncRunSummaryStatus

Recorded lifecycle state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SyncRunSummaryStatusSyncLifecycleStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SyncRunSummaryStatus("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `SyncRunSummaryStatusSyncLifecycleStatusUnspecified`   | SYNC_LIFECYCLE_STATUS_UNSPECIFIED                      |
| `SyncRunSummaryStatusSyncLifecycleStatusSyncing`       | SYNC_LIFECYCLE_STATUS_SYNCING                          |
| `SyncRunSummaryStatusSyncLifecycleStatusErrored`       | SYNC_LIFECYCLE_STATUS_ERRORED                          |
| `SyncRunSummaryStatusSyncLifecycleStatusCompleted`     | SYNC_LIFECYCLE_STATUS_COMPLETED                        |
| `SyncRunSummaryStatusSyncLifecycleStatusNotConnected`  | SYNC_LIFECYCLE_STATUS_NOT_CONNECTED                    |
| `SyncRunSummaryStatusSyncLifecycleStatusSkipped`       | SYNC_LIFECYCLE_STATUS_SKIPPED                          |
| `SyncRunSummaryStatusSyncLifecycleStatusErroredNoData` | SYNC_LIFECYCLE_STATUS_ERRORED_NO_DATA                  |