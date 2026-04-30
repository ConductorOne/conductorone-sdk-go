# PaperSecretStatus

Computed status

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretStatusSecretStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `PaperSecretStatusSecretStatusUnspecified` | SECRET_STATUS_UNSPECIFIED                  |
| `PaperSecretStatusSecretStatusActive`      | SECRET_STATUS_ACTIVE                       |
| `PaperSecretStatusSecretStatusExpired`     | SECRET_STATUS_EXPIRED                      |
| `PaperSecretStatusSecretStatusBurned`      | SECRET_STATUS_BURNED                       |
| `PaperSecretStatusSecretStatusRevoked`     | SECRET_STATUS_REVOKED                      |
| `PaperSecretStatusSecretStatusDataDeleted` | SECRET_STATUS_DATA_DELETED                 |