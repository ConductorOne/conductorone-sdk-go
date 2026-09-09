# ClawLLMRateSnapshotSource

The source field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClawLLMRateSnapshotSourceClawLlmRateSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClawLLMRateSnapshotSource("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ClawLLMRateSnapshotSourceClawLlmRateSourceUnspecified`      | CLAW_LLM_RATE_SOURCE_UNSPECIFIED                             |
| `ClawLLMRateSnapshotSourceClawLlmRateSourceListPrice`        | CLAW_LLM_RATE_SOURCE_LIST_PRICE                              |
| `ClawLLMRateSnapshotSourceClawLlmRateSourceNegotiated`       | CLAW_LLM_RATE_SOURCE_NEGOTIATED                              |
| `ClawLLMRateSnapshotSourceClawLlmRateSourceOperatorOverride` | CLAW_LLM_RATE_SOURCE_OPERATOR_OVERRIDE                       |
| `ClawLLMRateSnapshotSourceClawLlmRateSourceCustomerSupplied` | CLAW_LLM_RATE_SOURCE_CUSTOMER_SUPPLIED                       |