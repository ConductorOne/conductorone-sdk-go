# FunctionsServiceCreateFunctionRequestFunctionType

The type of function to create. Use FUNCTION_TYPE_ANY for user functions —
 that is the type the Functions UI lists. Do not use any other value.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FunctionsServiceCreateFunctionRequestFunctionTypeFunctionTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FunctionsServiceCreateFunctionRequestFunctionType("custom_value")
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `FunctionsServiceCreateFunctionRequestFunctionTypeFunctionTypeUnspecified` | FUNCTION_TYPE_UNSPECIFIED                                                  |
| `FunctionsServiceCreateFunctionRequestFunctionTypeFunctionTypeAny`         | FUNCTION_TYPE_ANY                                                          |
| `FunctionsServiceCreateFunctionRequestFunctionTypeFunctionTypeCodeMode`    | FUNCTION_TYPE_CODE_MODE                                                    |