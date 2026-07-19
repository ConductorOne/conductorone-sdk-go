# Persistence

Whether sessions may persist across browser restarts.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PersistencePersistenceModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Persistence("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `PersistencePersistenceModeUnspecified`     | PERSISTENCE_MODE_UNSPECIFIED                |
| `PersistencePersistenceModeAllowUserChoice` | PERSISTENCE_MODE_ALLOW_USER_CHOICE          |
| `PersistencePersistenceModeAlwaysPersist`   | PERSISTENCE_MODE_ALWAYS_PERSIST             |
| `PersistencePersistenceModeSessionOnly`     | PERSISTENCE_MODE_SESSION_ONLY               |