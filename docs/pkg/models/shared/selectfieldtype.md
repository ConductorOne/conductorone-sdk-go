# SelectFieldType

The type field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SelectFieldTypeSelectTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SelectFieldType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `SelectFieldTypeSelectTypeUnspecified` | SELECT_TYPE_UNSPECIFIED                |
| `SelectFieldTypeSelectTypeDropdown`    | SELECT_TYPE_DROPDOWN                   |
| `SelectFieldTypeSelectTypeRadio`       | SELECT_TYPE_RADIO                      |
| `SelectFieldTypeSelectTypeButtons`     | SELECT_TYPE_BUTTONS                    |