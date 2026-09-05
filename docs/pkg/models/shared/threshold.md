# Threshold

Deny (or flag) when the judge scores at or above this level. Unspecified =
 HIGH.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ThresholdPromptInjectionThresholdUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Threshold("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `ThresholdPromptInjectionThresholdUnspecified` | PROMPT_INJECTION_THRESHOLD_UNSPECIFIED         |
| `ThresholdPromptInjectionThresholdLow`         | PROMPT_INJECTION_THRESHOLD_LOW                 |
| `ThresholdPromptInjectionThresholdMedium`      | PROMPT_INJECTION_THRESHOLD_MEDIUM              |
| `ThresholdPromptInjectionThresholdHigh`        | PROMPT_INJECTION_THRESHOLD_HIGH                |