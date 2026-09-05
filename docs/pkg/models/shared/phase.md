# Phase

Durable execution progress and milestone timestamps. Pending refreshes
 expose their current phase; terminal runs retain timings for analysis.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PhaseReportRunPhaseUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Phase("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PhaseReportRunPhaseUnspecified`      | REPORT_RUN_PHASE_UNSPECIFIED          |
| `PhaseReportRunPhaseRequested`        | REPORT_RUN_PHASE_REQUESTED            |
| `PhaseReportRunPhasePreparingScratch` | REPORT_RUN_PHASE_PREPARING_SCRATCH    |
| `PhaseReportRunPhaseRunningFunction`  | REPORT_RUN_PHASE_RUNNING_FUNCTION     |
| `PhaseReportRunPhaseFinalizingOutput` | REPORT_RUN_PHASE_FINALIZING_OUTPUT    |
| `PhaseReportRunPhaseSucceeded`        | REPORT_RUN_PHASE_SUCCEEDED            |
| `PhaseReportRunPhaseFailed`           | REPORT_RUN_PHASE_FAILED               |