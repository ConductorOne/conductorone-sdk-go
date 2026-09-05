# DisabledReasonCircuitBreakerPeriod

Snapshot of the period at trip time.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DisabledReasonCircuitBreakerPeriodCircuitBreakerPeriodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DisabledReasonCircuitBreakerPeriod("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `DisabledReasonCircuitBreakerPeriodCircuitBreakerPeriodUnspecified` | CIRCUIT_BREAKER_PERIOD_UNSPECIFIED                                  |
| `DisabledReasonCircuitBreakerPeriodCircuitBreakerPeriodHour`        | CIRCUIT_BREAKER_PERIOD_HOUR                                         |
| `DisabledReasonCircuitBreakerPeriodCircuitBreakerPeriodDay`         | CIRCUIT_BREAKER_PERIOD_DAY                                          |
| `DisabledReasonCircuitBreakerPeriodCircuitBreakerPeriodWeek`        | CIRCUIT_BREAKER_PERIOD_WEEK                                         |
| `DisabledReasonCircuitBreakerPeriodCircuitBreakerPeriodMonth`       | CIRCUIT_BREAKER_PERIOD_MONTH                                        |