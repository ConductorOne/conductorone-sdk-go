# TerminalReason

Structured terminal reason, unspecified for legacy results.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TerminalReasonCustomAnalysisTerminalReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TerminalReason("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `TerminalReasonCustomAnalysisTerminalReasonUnspecified` | CUSTOM_ANALYSIS_TERMINAL_REASON_UNSPECIFIED             |
| `TerminalReasonCustomAnalysisTerminalReasonError`       | CUSTOM_ANALYSIS_TERMINAL_REASON_ERROR                   |
| `TerminalReasonCustomAnalysisTerminalReasonSuperseded`  | CUSTOM_ANALYSIS_TERMINAL_REASON_SUPERSEDED              |