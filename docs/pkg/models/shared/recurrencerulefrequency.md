# RecurrenceRuleFrequency

Frequency of the recurrence: FREQUENCY_DAILY, FREQUENCY_WEEKLY, FREQUENCY_MONTHLY, or FREQUENCY_YEARLY.
 Use FREQUENCY_NONE for a non-recurring schedule.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RecurrenceRuleFrequencyFrequencyUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RecurrenceRuleFrequency("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `RecurrenceRuleFrequencyFrequencyUnspecified` | FREQUENCY_UNSPECIFIED                         |
| `RecurrenceRuleFrequencyFrequencyNone`        | FREQUENCY_NONE                                |
| `RecurrenceRuleFrequencyFrequencyDaily`       | FREQUENCY_DAILY                               |
| `RecurrenceRuleFrequencyFrequencyWeekly`      | FREQUENCY_WEEKLY                              |
| `RecurrenceRuleFrequencyFrequencyMonthly`     | FREQUENCY_MONTHLY                             |
| `RecurrenceRuleFrequencyFrequencyYearly`      | FREQUENCY_YEARLY                              |