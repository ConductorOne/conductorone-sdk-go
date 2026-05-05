# Severity

The severity field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SeverityUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.Severity("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `SeverityUnknown` | UNKNOWN           |
| `SeverityHint`    | HINT              |
| `SeverityInfo`    | INFO              |
| `SeverityWarning` | WARNING           |
| `SeverityError`   | ERROR             |