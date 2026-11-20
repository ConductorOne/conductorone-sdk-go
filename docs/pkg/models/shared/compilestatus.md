# CompileStatus

Compile state of the published commit for the function's current
 configuration. Populated by GetFunction only; unset in list responses.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CompileStatusFunctionCompileStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CompileStatus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `CompileStatusFunctionCompileStatusUnspecified` | FUNCTION_COMPILE_STATUS_UNSPECIFIED             |
| `CompileStatusFunctionCompileStatusNotCompiled` | FUNCTION_COMPILE_STATUS_NOT_COMPILED            |
| `CompileStatusFunctionCompileStatusCompiling`   | FUNCTION_COMPILE_STATUS_COMPILING               |
| `CompileStatusFunctionCompileStatusReady`       | FUNCTION_COMPILE_STATUS_READY                   |
| `CompileStatusFunctionCompileStatusFailed`      | FUNCTION_COMPILE_STATUS_FAILED                  |