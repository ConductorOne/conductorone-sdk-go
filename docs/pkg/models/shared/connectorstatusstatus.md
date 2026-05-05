# ConnectorStatusStatus

The status of the connector sync.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ConnectorStatusStatusSyncStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ConnectorStatusStatus("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `ConnectorStatusStatusSyncStatusUnspecified` | SYNC_STATUS_UNSPECIFIED                      |
| `ConnectorStatusStatusSyncStatusRunning`     | SYNC_STATUS_RUNNING                          |
| `ConnectorStatusStatusSyncStatusDone`        | SYNC_STATUS_DONE                             |
| `ConnectorStatusStatusSyncStatusError`       | SYNC_STATUS_ERROR                            |
| `ConnectorStatusStatusSyncStatusDisabled`    | SYNC_STATUS_DISABLED                         |