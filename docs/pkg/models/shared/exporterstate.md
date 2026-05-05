# ExporterState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExporterStateExportStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExporterState("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ExporterStateExportStateUnspecified` | EXPORT_STATE_UNSPECIFIED              |
| `ExporterStateExportStateExporting`   | EXPORT_STATE_EXPORTING                |
| `ExporterStateExportStateWaiting`     | EXPORT_STATE_WAITING                  |
| `ExporterStateExportStateError`       | EXPORT_STATE_ERROR                    |