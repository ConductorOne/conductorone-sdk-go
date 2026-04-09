# TypeFilter

The typeFilter field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TypeFilterGrantFilterTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TypeFilter("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `TypeFilterGrantFilterTypeUnspecified` | GRANT_FILTER_TYPE_UNSPECIFIED          |
| `TypeFilterGrantFilterTypePermanent`   | GRANT_FILTER_TYPE_PERMANENT            |
| `TypeFilterGrantFilterTypeTemporary`   | GRANT_FILTER_TYPE_TEMPORARY            |