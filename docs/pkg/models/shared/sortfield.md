# SortField

Column to sort by. Unspecified (0) means sort by created_at desc (server default).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SortFieldAutomationSortFieldUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SortField("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `SortFieldAutomationSortFieldUnspecified`        | AUTOMATION_SORT_FIELD_UNSPECIFIED                |
| `SortFieldAutomationSortFieldDisplayName`        | AUTOMATION_SORT_FIELD_DISPLAY_NAME               |
| `SortFieldAutomationSortFieldCreatedAt`          | AUTOMATION_SORT_FIELD_CREATED_AT                 |
| `SortFieldAutomationSortFieldLastExecutedAt`     | AUTOMATION_SORT_FIELD_LAST_EXECUTED_AT           |
| `SortFieldAutomationSortFieldEnabled`            | AUTOMATION_SORT_FIELD_ENABLED                    |
| `SortFieldAutomationSortFieldPrimaryTriggerType` | AUTOMATION_SORT_FIELD_PRIMARY_TRIGGER_TYPE       |