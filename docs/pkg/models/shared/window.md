# Window

How the limit counts requests.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.WindowWindowUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Window("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `WindowWindowUnspecified` | WINDOW_UNSPECIFIED        |
| `WindowWindowPerHour`     | WINDOW_PER_HOUR           |
| `WindowWindowConcurrent`  | WINDOW_CONCURRENT         |