# TaskAuditViewEventType

The eventType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAuditViewEventTypeTaskAuditEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAuditViewEventType("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `TaskAuditViewEventTypeTaskAuditEventTypeUnspecified` | TASK_AUDIT_EVENT_TYPE_UNSPECIFIED                     |
| `TaskAuditViewEventTypeTaskAuditEventTypeNeutral`     | TASK_AUDIT_EVENT_TYPE_NEUTRAL                         |
| `TaskAuditViewEventTypeTaskAuditEventTypeError`       | TASK_AUDIT_EVENT_TYPE_ERROR                           |