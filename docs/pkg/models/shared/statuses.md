# Statuses

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StatusesSecretStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Statuses("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `StatusesSecretStatusUnspecified` | SECRET_STATUS_UNSPECIFIED         |
| `StatusesSecretStatusActive`      | SECRET_STATUS_ACTIVE              |
| `StatusesSecretStatusExpired`     | SECRET_STATUS_EXPIRED             |
| `StatusesSecretStatusBurned`      | SECRET_STATUS_BURNED              |
| `StatusesSecretStatusRevoked`     | SECRET_STATUS_REVOKED             |
| `StatusesSecretStatusDataDeleted` | SECRET_STATUS_DATA_DELETED        |