# SyncRunSummaryErrorClass

Classification of the latest recorded error.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SyncRunSummaryErrorClassSyncErrorClassUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SyncRunSummaryErrorClass("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `SyncRunSummaryErrorClassSyncErrorClassUnspecified`         | SYNC_ERROR_CLASS_UNSPECIFIED                                |
| `SyncRunSummaryErrorClassSyncErrorClassAuthOrScope`         | SYNC_ERROR_CLASS_AUTH_OR_SCOPE                              |
| `SyncRunSummaryErrorClassSyncErrorClassRateLimit`           | SYNC_ERROR_CLASS_RATE_LIMIT                                 |
| `SyncRunSummaryErrorClassSyncErrorClassTimeout`             | SYNC_ERROR_CLASS_TIMEOUT                                    |
| `SyncRunSummaryErrorClassSyncErrorClassUpstreamConsistency` | SYNC_ERROR_CLASS_UPSTREAM_CONSISTENCY                       |
| `SyncRunSummaryErrorClassSyncErrorClassConnectorBug`        | SYNC_ERROR_CLASS_CONNECTOR_BUG                              |
| `SyncRunSummaryErrorClassSyncErrorClassStorage`             | SYNC_ERROR_CLASS_STORAGE                                    |
| `SyncRunSummaryErrorClassSyncErrorClassDataAnomaly`         | SYNC_ERROR_CLASS_DATA_ANOMALY                               |
| `SyncRunSummaryErrorClassSyncErrorClassRuntimeRefused`      | SYNC_ERROR_CLASS_RUNTIME_REFUSED                            |
| `SyncRunSummaryErrorClassSyncErrorClassDeployment`          | SYNC_ERROR_CLASS_DEPLOYMENT                                 |
| `SyncRunSummaryErrorClassSyncErrorClassUnknown`             | SYNC_ERROR_CLASS_UNKNOWN                                    |