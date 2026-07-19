# TFFieldMappingTFType

The Terraform attribute type. For collections of structured
 objects, the element shape is in `nested_fields`. For collections
 of primitives, the element type is in `element_tf_type`.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TFFieldMappingTFTypeTfTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TFFieldMappingTFType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `TFFieldMappingTFTypeTfTypeUnspecified` | TF_TYPE_UNSPECIFIED                     |
| `TFFieldMappingTFTypeTfTypeString`      | TF_TYPE_STRING                          |
| `TFFieldMappingTFTypeTfTypeNumber`      | TF_TYPE_NUMBER                          |
| `TFFieldMappingTFTypeTfTypeBool`        | TF_TYPE_BOOL                            |
| `TFFieldMappingTFTypeTfTypeList`        | TF_TYPE_LIST                            |
| `TFFieldMappingTFTypeTfTypeSet`         | TF_TYPE_SET                             |
| `TFFieldMappingTFTypeTfTypeMap`         | TF_TYPE_MAP                             |
| `TFFieldMappingTFTypeTfTypeObject`      | TF_TYPE_OBJECT                          |
| `TFFieldMappingTFTypeTfTypeTuple`       | TF_TYPE_TUPLE                           |