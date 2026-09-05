# FunctionType

The functionType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FunctionTypeFunctionTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FunctionType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `FunctionTypeFunctionTypeUnspecified` | FUNCTION_TYPE_UNSPECIFIED             |
| `FunctionTypeFunctionTypeAny`         | FUNCTION_TYPE_ANY                     |
| `FunctionTypeFunctionTypeCodeMode`    | FUNCTION_TYPE_CODE_MODE               |
| `FunctionTypeFunctionTypeConnector`   | FUNCTION_TYPE_CONNECTOR               |