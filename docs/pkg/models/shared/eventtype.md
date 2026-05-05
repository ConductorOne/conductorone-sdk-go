# EventType

The eventType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EventTypeGrantEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EventType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `EventTypeGrantEventTypeUnspecified` | GRANT_EVENT_TYPE_UNSPECIFIED         |
| `EventTypeGrantEventTypeAdded`       | GRANT_EVENT_TYPE_ADDED               |
| `EventTypeGrantEventTypeRemoved`     | GRANT_EVENT_TYPE_REMOVED             |