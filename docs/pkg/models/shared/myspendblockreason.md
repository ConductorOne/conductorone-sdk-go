# MySpendBlockReason

Reason the calls were denied.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MySpendBlockReasonSpendDenyReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MySpendBlockReason("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `MySpendBlockReasonSpendDenyReasonUnspecified`      | SPEND_DENY_REASON_UNSPECIFIED                       |
| `MySpendBlockReasonSpendDenyReasonTenantFrozen`     | SPEND_DENY_REASON_TENANT_FROZEN                     |
| `MySpendBlockReasonSpendDenyReasonSuspendedByAdmin` | SPEND_DENY_REASON_SUSPENDED_BY_ADMIN                |
| `MySpendBlockReasonSpendDenyReasonAppSuspended`     | SPEND_DENY_REASON_APP_SUSPENDED                     |
| `MySpendBlockReasonSpendDenyReasonAppPausedByYou`   | SPEND_DENY_REASON_APP_PAUSED_BY_YOU                 |
| `MySpendBlockReasonSpendDenyReasonNoSupply`         | SPEND_DENY_REASON_NO_SUPPLY                         |