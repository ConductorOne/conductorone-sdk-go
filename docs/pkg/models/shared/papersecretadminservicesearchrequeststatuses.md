# PaperSecretAdminServiceSearchRequestStatuses

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretAdminServiceSearchRequestStatusesSecretStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretAdminServiceSearchRequestStatuses("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `PaperSecretAdminServiceSearchRequestStatusesSecretStatusUnspecified` | SECRET_STATUS_UNSPECIFIED                                             |
| `PaperSecretAdminServiceSearchRequestStatusesSecretStatusActive`      | SECRET_STATUS_ACTIVE                                                  |
| `PaperSecretAdminServiceSearchRequestStatusesSecretStatusExpired`     | SECRET_STATUS_EXPIRED                                                 |
| `PaperSecretAdminServiceSearchRequestStatusesSecretStatusBurned`      | SECRET_STATUS_BURNED                                                  |
| `PaperSecretAdminServiceSearchRequestStatusesSecretStatusRevoked`     | SECRET_STATUS_REVOKED                                                 |
| `PaperSecretAdminServiceSearchRequestStatusesSecretStatusDataDeleted` | SECRET_STATUS_DATA_DELETED                                            |