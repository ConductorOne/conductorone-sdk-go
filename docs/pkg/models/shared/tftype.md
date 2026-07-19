# TfType

The TF attribute type of the component value.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TfTypeTfTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TfType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `TfTypeTfTypeUnspecified` | TF_TYPE_UNSPECIFIED       |
| `TfTypeTfTypeString`      | TF_TYPE_STRING            |
| `TfTypeTfTypeNumber`      | TF_TYPE_NUMBER            |
| `TfTypeTfTypeBool`        | TF_TYPE_BOOL              |
| `TfTypeTfTypeList`        | TF_TYPE_LIST              |
| `TfTypeTfTypeSet`         | TF_TYPE_SET               |
| `TfTypeTfTypeMap`         | TF_TYPE_MAP               |
| `TfTypeTfTypeObject`      | TF_TYPE_OBJECT            |
| `TfTypeTfTypeTuple`       | TF_TYPE_TUPLE             |