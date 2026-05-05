# CircuitBreakerPeriod

The circuitBreakerPeriod field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CircuitBreakerPeriodCircuitBreakerPeriodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CircuitBreakerPeriod("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `CircuitBreakerPeriodCircuitBreakerPeriodUnspecified` | CIRCUIT_BREAKER_PERIOD_UNSPECIFIED                    |
| `CircuitBreakerPeriodCircuitBreakerPeriodHour`        | CIRCUIT_BREAKER_PERIOD_HOUR                           |
| `CircuitBreakerPeriodCircuitBreakerPeriodDay`         | CIRCUIT_BREAKER_PERIOD_DAY                            |
| `CircuitBreakerPeriodCircuitBreakerPeriodWeek`        | CIRCUIT_BREAKER_PERIOD_WEEK                           |
| `CircuitBreakerPeriodCircuitBreakerPeriodMonth`       | CIRCUIT_BREAKER_PERIOD_MONTH                          |