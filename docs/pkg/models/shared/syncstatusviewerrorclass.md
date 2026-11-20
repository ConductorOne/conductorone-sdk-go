# SyncStatusViewErrorClass

Classification of the current error.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SyncStatusViewErrorClassSyncErrorClassUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SyncStatusViewErrorClass("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `SyncStatusViewErrorClassSyncErrorClassUnspecified`         | SYNC_ERROR_CLASS_UNSPECIFIED                                |
| `SyncStatusViewErrorClassSyncErrorClassAuthOrScope`         | SYNC_ERROR_CLASS_AUTH_OR_SCOPE                              |
| `SyncStatusViewErrorClassSyncErrorClassRateLimit`           | SYNC_ERROR_CLASS_RATE_LIMIT                                 |
| `SyncStatusViewErrorClassSyncErrorClassTimeout`             | SYNC_ERROR_CLASS_TIMEOUT                                    |
| `SyncStatusViewErrorClassSyncErrorClassUpstreamConsistency` | SYNC_ERROR_CLASS_UPSTREAM_CONSISTENCY                       |
| `SyncStatusViewErrorClassSyncErrorClassConnectorBug`        | SYNC_ERROR_CLASS_CONNECTOR_BUG                              |
| `SyncStatusViewErrorClassSyncErrorClassStorage`             | SYNC_ERROR_CLASS_STORAGE                                    |
| `SyncStatusViewErrorClassSyncErrorClassDataAnomaly`         | SYNC_ERROR_CLASS_DATA_ANOMALY                               |
| `SyncStatusViewErrorClassSyncErrorClassRuntimeRefused`      | SYNC_ERROR_CLASS_RUNTIME_REFUSED                            |
| `SyncStatusViewErrorClassSyncErrorClassDeployment`          | SYNC_ERROR_CLASS_DEPLOYMENT                                 |
| `SyncStatusViewErrorClassSyncErrorClassUnknown`             | SYNC_ERROR_CLASS_UNKNOWN                                    |