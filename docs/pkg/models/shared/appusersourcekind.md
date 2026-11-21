# AppUserSourceKind

Where the record came from.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AppUserSourceKindKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AppUserSourceKind("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `AppUserSourceKindKindUnspecified` | KIND_UNSPECIFIED                   |
| `AppUserSourceKindKindConnector`   | KIND_CONNECTOR                     |
| `AppUserSourceKindKindExternal`    | KIND_EXTERNAL                      |