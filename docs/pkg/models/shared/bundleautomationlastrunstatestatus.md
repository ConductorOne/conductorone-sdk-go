# BundleAutomationLastRunStateStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BundleAutomationLastRunStateStatusBundleAutomationRunStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BundleAutomationLastRunStateStatus("custom_value")
```


## Values

| Name                                                                            | Value                                                                           |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `BundleAutomationLastRunStateStatusBundleAutomationRunStatusUnspecified`        | BUNDLE_AUTOMATION_RUN_STATUS_UNSPECIFIED                                        |
| `BundleAutomationLastRunStateStatusBundleAutomationRunStatusSuccess`            | BUNDLE_AUTOMATION_RUN_STATUS_SUCCESS                                            |
| `BundleAutomationLastRunStateStatusBundleAutomationRunStatusFailure`            | BUNDLE_AUTOMATION_RUN_STATUS_FAILURE                                            |
| `BundleAutomationLastRunStateStatusBundleAutomationRunStatusInProgress`         | BUNDLE_AUTOMATION_RUN_STATUS_IN_PROGRESS                                        |
| `BundleAutomationLastRunStateStatusBundleAutomationRunStatusWaitingForApproval` | BUNDLE_AUTOMATION_RUN_STATUS_WAITING_FOR_APPROVAL                               |