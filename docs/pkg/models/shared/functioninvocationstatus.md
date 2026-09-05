# FunctionInvocationStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FunctionInvocationStatusFunctionInvocationStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FunctionInvocationStatus("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `FunctionInvocationStatusFunctionInvocationStatusUnspecified`           | FUNCTION_INVOCATION_STATUS_UNSPECIFIED                                  |
| `FunctionInvocationStatusFunctionInvocationStatusPending`               | FUNCTION_INVOCATION_STATUS_PENDING                                      |
| `FunctionInvocationStatusFunctionInvocationStatusRunning`               | FUNCTION_INVOCATION_STATUS_RUNNING                                      |
| `FunctionInvocationStatusFunctionInvocationStatusSuccess`               | FUNCTION_INVOCATION_STATUS_SUCCESS                                      |
| `FunctionInvocationStatusFunctionInvocationStatusError`                 | FUNCTION_INVOCATION_STATUS_ERROR                                        |
| `FunctionInvocationStatusFunctionInvocationStatusCancellationRequested` | FUNCTION_INVOCATION_STATUS_CANCELLATION_REQUESTED                       |
| `FunctionInvocationStatusFunctionInvocationStatusCancelled`             | FUNCTION_INVOCATION_STATUS_CANCELLED                                    |
| `FunctionInvocationStatusFunctionInvocationStatusUnknown`               | FUNCTION_INVOCATION_STATUS_UNKNOWN                                      |