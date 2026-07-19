# EditorMarkerSeverity

The severity field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EditorMarkerSeverityUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.EditorMarkerSeverity("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `EditorMarkerSeverityUnknown` | UNKNOWN                       |
| `EditorMarkerSeverityHint`    | HINT                          |
| `EditorMarkerSeverityInfo`    | INFO                          |
| `EditorMarkerSeverityWarning` | WARNING                       |
| `EditorMarkerSeverityError`   | ERROR                         |