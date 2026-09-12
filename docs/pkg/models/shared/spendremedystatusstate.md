# SpendRemedyStatusState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SpendRemedyStatusStateSpendRemedyStatusStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SpendRemedyStatusState("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `SpendRemedyStatusStateSpendRemedyStatusStateUnspecified`                 | SPEND_REMEDY_STATUS_STATE_UNSPECIFIED                                     |
| `SpendRemedyStatusStateSpendRemedyStatusStatePending`                     | SPEND_REMEDY_STATUS_STATE_PENDING                                         |
| `SpendRemedyStatusStateSpendRemedyStatusStateRejected`                    | SPEND_REMEDY_STATUS_STATE_REJECTED                                        |
| `SpendRemedyStatusStateSpendRemedyStatusStateCancelled`                   | SPEND_REMEDY_STATUS_STATE_CANCELLED                                       |
| `SpendRemedyStatusStateSpendRemedyStatusStateExpired`                     | SPEND_REMEDY_STATUS_STATE_EXPIRED                                         |
| `SpendRemedyStatusStateSpendRemedyStatusStateApprovedAwaitingFulfillment` | SPEND_REMEDY_STATUS_STATE_APPROVED_AWAITING_FULFILLMENT                   |
| `SpendRemedyStatusStateSpendRemedyStatusStateFulfillmentFailed`           | SPEND_REMEDY_STATUS_STATE_FULFILLMENT_FAILED                              |
| `SpendRemedyStatusStateSpendRemedyStatusStateAppliedRetryRequired`        | SPEND_REMEDY_STATUS_STATE_APPLIED_RETRY_REQUIRED                          |
| `SpendRemedyStatusStateSpendRemedyStatusStateAppliedStillBlocked`         | SPEND_REMEDY_STATUS_STATE_APPLIED_STILL_BLOCKED                           |
| `SpendRemedyStatusStateSpendRemedyStatusStateStale`                       | SPEND_REMEDY_STATUS_STATE_STALE                                           |
| `SpendRemedyStatusStateSpendRemedyStatusStateUnavailable`                 | SPEND_REMEDY_STATUS_STATE_UNAVAILABLE                                     |