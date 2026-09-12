# BlockingReason

The blockingReason field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BlockingReasonSpendDenyReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BlockingReason("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `BlockingReasonSpendDenyReasonUnspecified`      | SPEND_DENY_REASON_UNSPECIFIED                   |
| `BlockingReasonSpendDenyReasonTenantFrozen`     | SPEND_DENY_REASON_TENANT_FROZEN                 |
| `BlockingReasonSpendDenyReasonSuspendedByAdmin` | SPEND_DENY_REASON_SUSPENDED_BY_ADMIN            |
| `BlockingReasonSpendDenyReasonAppSuspended`     | SPEND_DENY_REASON_APP_SUSPENDED                 |
| `BlockingReasonSpendDenyReasonAppPausedByYou`   | SPEND_DENY_REASON_APP_PAUSED_BY_YOU             |
| `BlockingReasonSpendDenyReasonNoSupply`         | SPEND_DENY_REASON_NO_SUPPLY                     |