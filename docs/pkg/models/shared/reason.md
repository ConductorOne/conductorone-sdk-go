# Reason

The reason field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ReasonDeactivatedOwnerReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Reason("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `ReasonDeactivatedOwnerReasonUnspecified`        | DEACTIVATED_OWNER_REASON_UNSPECIFIED             |
| `ReasonDeactivatedOwnerReasonUserDeleted`        | DEACTIVATED_OWNER_REASON_USER_DELETED            |
| `ReasonDeactivatedOwnerReasonUserDisabled`       | DEACTIVATED_OWNER_REASON_USER_DISABLED           |
| `ReasonDeactivatedOwnerReasonEmploymentInactive` | DEACTIVATED_OWNER_REASON_EMPLOYMENT_INACTIVE     |