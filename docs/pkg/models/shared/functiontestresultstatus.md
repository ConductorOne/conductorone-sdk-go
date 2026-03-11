# FunctionTestResultStatus

The test result status.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FunctionTestResultStatusFunctionTestResultStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FunctionTestResultStatus("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `FunctionTestResultStatusFunctionTestResultStatusUnspecified` | FUNCTION_TEST_RESULT_STATUS_UNSPECIFIED                       |
| `FunctionTestResultStatusFunctionTestResultStatusOk`          | FUNCTION_TEST_RESULT_STATUS_OK                                |
| `FunctionTestResultStatusFunctionTestResultStatusFail`        | FUNCTION_TEST_RESULT_STATUS_FAIL                              |
| `FunctionTestResultStatusFunctionTestResultStatusSkipped`     | FUNCTION_TEST_RESULT_STATUS_SKIPPED                           |