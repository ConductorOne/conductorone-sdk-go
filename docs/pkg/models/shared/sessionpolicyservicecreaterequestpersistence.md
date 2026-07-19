# SessionPolicyServiceCreateRequestPersistence

The persistence field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyServiceCreateRequestPersistencePersistenceModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyServiceCreateRequestPersistence("custom_value")
```


## Values

| Name                                                                         | Value                                                                        |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `SessionPolicyServiceCreateRequestPersistencePersistenceModeUnspecified`     | PERSISTENCE_MODE_UNSPECIFIED                                                 |
| `SessionPolicyServiceCreateRequestPersistencePersistenceModeAllowUserChoice` | PERSISTENCE_MODE_ALLOW_USER_CHOICE                                           |
| `SessionPolicyServiceCreateRequestPersistencePersistenceModeAlwaysPersist`   | PERSISTENCE_MODE_ALWAYS_PERSIST                                              |
| `SessionPolicyServiceCreateRequestPersistencePersistenceModeSessionOnly`     | PERSISTENCE_MODE_SESSION_ONLY                                                |