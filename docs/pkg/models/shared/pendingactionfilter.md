# PendingActionFilter

Filter tasks by pending action status. Only applies when exactly one access_review_id is specified.
 Requires the REVIEWS_PENDING_ACTIONS feature flag to be enabled.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PendingActionFilterPendingActionFilterUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PendingActionFilter("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `PendingActionFilterPendingActionFilterUnspecified`    | PENDING_ACTION_FILTER_UNSPECIFIED                      |
| `PendingActionFilterPendingActionFilterWithPending`    | PENDING_ACTION_FILTER_WITH_PENDING                     |
| `PendingActionFilterPendingActionFilterWithoutPending` | PENDING_ACTION_FILTER_WITHOUT_PENDING                  |