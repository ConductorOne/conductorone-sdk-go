# BundleAutomationCircuitBreakerState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BundleAutomationCircuitBreakerStateCircuitBreakerStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BundleAutomationCircuitBreakerState("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `BundleAutomationCircuitBreakerStateCircuitBreakerStateUnspecified` | CIRCUIT_BREAKER_STATE_UNSPECIFIED                                   |
| `BundleAutomationCircuitBreakerStateCircuitBreakerStateTriggered`   | CIRCUIT_BREAKER_STATE_TRIGGERED                                     |
| `BundleAutomationCircuitBreakerStateCircuitBreakerStateBypass`      | CIRCUIT_BREAKER_STATE_BYPASS                                        |