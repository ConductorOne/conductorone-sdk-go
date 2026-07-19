# SearchAutomationsRequestSortField

Column to sort by. Unspecified (0) means sort by created_at desc (server default).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SearchAutomationsRequestSortFieldAutomationSortFieldUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SearchAutomationsRequestSortField("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `SearchAutomationsRequestSortFieldAutomationSortFieldUnspecified`        | AUTOMATION_SORT_FIELD_UNSPECIFIED                                        |
| `SearchAutomationsRequestSortFieldAutomationSortFieldDisplayName`        | AUTOMATION_SORT_FIELD_DISPLAY_NAME                                       |
| `SearchAutomationsRequestSortFieldAutomationSortFieldCreatedAt`          | AUTOMATION_SORT_FIELD_CREATED_AT                                         |
| `SearchAutomationsRequestSortFieldAutomationSortFieldLastExecutedAt`     | AUTOMATION_SORT_FIELD_LAST_EXECUTED_AT                                   |
| `SearchAutomationsRequestSortFieldAutomationSortFieldEnabled`            | AUTOMATION_SORT_FIELD_ENABLED                                            |
| `SearchAutomationsRequestSortFieldAutomationSortFieldPrimaryTriggerType` | AUTOMATION_SORT_FIELD_PRIMARY_TRIGGER_TYPE                               |