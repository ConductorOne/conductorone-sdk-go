# ErrorClass

Likely failure family.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ErrorClassSyncErrorClassUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ErrorClass("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `ErrorClassSyncErrorClassUnspecified`         | SYNC_ERROR_CLASS_UNSPECIFIED                  |
| `ErrorClassSyncErrorClassAuthOrScope`         | SYNC_ERROR_CLASS_AUTH_OR_SCOPE                |
| `ErrorClassSyncErrorClassRateLimit`           | SYNC_ERROR_CLASS_RATE_LIMIT                   |
| `ErrorClassSyncErrorClassTimeout`             | SYNC_ERROR_CLASS_TIMEOUT                      |
| `ErrorClassSyncErrorClassUpstreamConsistency` | SYNC_ERROR_CLASS_UPSTREAM_CONSISTENCY         |
| `ErrorClassSyncErrorClassConnectorBug`        | SYNC_ERROR_CLASS_CONNECTOR_BUG                |
| `ErrorClassSyncErrorClassStorage`             | SYNC_ERROR_CLASS_STORAGE                      |
| `ErrorClassSyncErrorClassDataAnomaly`         | SYNC_ERROR_CLASS_DATA_ANOMALY                 |
| `ErrorClassSyncErrorClassRuntimeRefused`      | SYNC_ERROR_CLASS_RUNTIME_REFUSED              |
| `ErrorClassSyncErrorClassDeployment`          | SYNC_ERROR_CLASS_DEPLOYMENT                   |
| `ErrorClassSyncErrorClassUnknown`             | SYNC_ERROR_CLASS_UNKNOWN                      |