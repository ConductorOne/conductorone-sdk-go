# DeactivatedOwnerTypeSource

The source field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DeactivatedOwnerTypeSourceDeactivatedOwnerSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DeactivatedOwnerTypeSource("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `DeactivatedOwnerTypeSourceDeactivatedOwnerSourceUnspecified`         | DEACTIVATED_OWNER_SOURCE_UNSPECIFIED                                  |
| `DeactivatedOwnerTypeSourceDeactivatedOwnerSourceIdentityCorrelation` | DEACTIVATED_OWNER_SOURCE_IDENTITY_CORRELATION                         |
| `DeactivatedOwnerTypeSourceDeactivatedOwnerSourceOwnershipAssigned`   | DEACTIVATED_OWNER_SOURCE_OWNERSHIP_ASSIGNED                           |
| `DeactivatedOwnerTypeSourceDeactivatedOwnerSourceSecretRunAsIdentity` | DEACTIVATED_OWNER_SOURCE_SECRET_RUN_AS_IDENTITY                       |