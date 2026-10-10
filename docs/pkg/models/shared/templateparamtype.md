# TemplateParamType

The type field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TemplateParamTypeTemplateParamTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TemplateParamType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `TemplateParamTypeTemplateParamTypeUnspecified`  | TEMPLATE_PARAM_TYPE_UNSPECIFIED                  |
| `TemplateParamTypeTemplateParamTypeGrantPolicy`  | TEMPLATE_PARAM_TYPE_GRANT_POLICY                 |
| `TemplateParamTypeTemplateParamTypeHookFunction` | TEMPLATE_PARAM_TYPE_HOOK_FUNCTION                |
| `TemplateParamTypeTemplateParamTypeString`       | TEMPLATE_PARAM_TYPE_STRING                       |