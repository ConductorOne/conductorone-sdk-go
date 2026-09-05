# GrantSourceFilter

Restricts the step to grants of either DIRECT (grants the user holds directly,
 including grants that are also inherited) or UNSPECIFIED (all grants).
 Composes with every inclusion mode, including inclusion_list_cel.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GrantSourceFilterGrantSourceFilterUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GrantSourceFilter("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `GrantSourceFilterGrantSourceFilterUnspecified` | GRANT_SOURCE_FILTER_UNSPECIFIED                 |
| `GrantSourceFilterGrantSourceFilterDirect`      | GRANT_SOURCE_FILTER_DIRECT                      |