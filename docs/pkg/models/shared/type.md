# Type

The type field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TypeSelectTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Type("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `TypeSelectTypeUnspecified` | SELECT_TYPE_UNSPECIFIED     |
| `TypeSelectTypeDropdown`    | SELECT_TYPE_DROPDOWN        |
| `TypeSelectTypeRadio`       | SELECT_TYPE_RADIO           |
| `TypeSelectTypeButtons`     | SELECT_TYPE_BUTTONS         |