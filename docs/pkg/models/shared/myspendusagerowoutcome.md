# MySpendUsageRowOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MySpendUsageRowOutcomeMySpendUsageOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MySpendUsageRowOutcome("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `MySpendUsageRowOutcomeMySpendUsageOutcomeUnspecified`   | MY_SPEND_USAGE_OUTCOME_UNSPECIFIED                       |
| `MySpendUsageRowOutcomeMySpendUsageOutcomeSettled`       | MY_SPEND_USAGE_OUTCOME_SETTLED                           |
| `MySpendUsageRowOutcomeMySpendUsageOutcomeSettleDropped` | MY_SPEND_USAGE_OUTCOME_SETTLE_DROPPED                    |
| `MySpendUsageRowOutcomeMySpendUsageOutcomeNotGoverned`   | MY_SPEND_USAGE_OUTCOME_NOT_GOVERNED                      |
| `MySpendUsageRowOutcomeMySpendUsageOutcomeZeroPriced`    | MY_SPEND_USAGE_OUTCOME_ZERO_PRICED                       |
| `MySpendUsageRowOutcomeMySpendUsageOutcomeUnpriced`      | MY_SPEND_USAGE_OUTCOME_UNPRICED                          |