# GroupBy

How requests are combined into entries. Defaults to GROUP_BY_API_KEY.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GroupByGroupByUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GroupBy("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `GroupByGroupByUnspecified` | GROUP_BY_UNSPECIFIED        |
| `GroupByGroupByAPIKey`      | GROUP_BY_API_KEY            |
| `GroupByGroupByUser`        | GROUP_BY_USER               |