# ElementTfType

For collection fields (list/set/tuple/map) whose elements are
 primitives (string/number/bool), the TF type of those elements.
 TF_TYPE_UNSPECIFIED for non-collection fields and for collections
 of objects (where `nested_fields` describes the element shape).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ElementTfTypeTfTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ElementTfType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ElementTfTypeTfTypeUnspecified` | TF_TYPE_UNSPECIFIED              |
| `ElementTfTypeTfTypeString`      | TF_TYPE_STRING                   |
| `ElementTfTypeTfTypeNumber`      | TF_TYPE_NUMBER                   |
| `ElementTfTypeTfTypeBool`        | TF_TYPE_BOOL                     |
| `ElementTfTypeTfTypeList`        | TF_TYPE_LIST                     |
| `ElementTfTypeTfTypeSet`         | TF_TYPE_SET                      |
| `ElementTfTypeTfTypeMap`         | TF_TYPE_MAP                      |
| `ElementTfTypeTfTypeObject`      | TF_TYPE_OBJECT                   |
| `ElementTfTypeTfTypeTuple`       | TF_TYPE_TUPLE                    |