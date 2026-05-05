# Period

Snapshot of the period at trip time.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PeriodCircuitBreakerPeriodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Period("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `PeriodCircuitBreakerPeriodUnspecified` | CIRCUIT_BREAKER_PERIOD_UNSPECIFIED      |
| `PeriodCircuitBreakerPeriodHour`        | CIRCUIT_BREAKER_PERIOD_HOUR             |
| `PeriodCircuitBreakerPeriodDay`         | CIRCUIT_BREAKER_PERIOD_DAY              |
| `PeriodCircuitBreakerPeriodWeek`        | CIRCUIT_BREAKER_PERIOD_WEEK             |
| `PeriodCircuitBreakerPeriodMonth`       | CIRCUIT_BREAKER_PERIOD_MONTH            |