# BundleAutomationCelEvaluationStateStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BundleAutomationCelEvaluationStateStatusBundleAutomationRunStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BundleAutomationCelEvaluationStateStatus("custom_value")
```


## Values

| Name                                                                                  | Value                                                                                 |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `BundleAutomationCelEvaluationStateStatusBundleAutomationRunStatusUnspecified`        | BUNDLE_AUTOMATION_RUN_STATUS_UNSPECIFIED                                              |
| `BundleAutomationCelEvaluationStateStatusBundleAutomationRunStatusSuccess`            | BUNDLE_AUTOMATION_RUN_STATUS_SUCCESS                                                  |
| `BundleAutomationCelEvaluationStateStatusBundleAutomationRunStatusFailure`            | BUNDLE_AUTOMATION_RUN_STATUS_FAILURE                                                  |
| `BundleAutomationCelEvaluationStateStatusBundleAutomationRunStatusInProgress`         | BUNDLE_AUTOMATION_RUN_STATUS_IN_PROGRESS                                              |
| `BundleAutomationCelEvaluationStateStatusBundleAutomationRunStatusWaitingForApproval` | BUNDLE_AUTOMATION_RUN_STATUS_WAITING_FOR_APPROVAL                                     |