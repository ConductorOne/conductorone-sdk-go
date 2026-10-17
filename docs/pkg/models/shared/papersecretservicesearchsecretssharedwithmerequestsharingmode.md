# PaperSecretServiceSearchSecretsSharedWithMeRequestSharingMode

Filter by sharing mode. Unspecified returns internal secrets; external returns none.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceSearchSecretsSharedWithMeRequestSharingModePaperVaultSharingModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceSearchSecretsSharedWithMeRequestSharingMode("custom_value")
```


## Values

| Name                                                                                            | Value                                                                                           |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `PaperSecretServiceSearchSecretsSharedWithMeRequestSharingModePaperVaultSharingModeUnspecified` | PAPER_VAULT_SHARING_MODE_UNSPECIFIED                                                            |
| `PaperSecretServiceSearchSecretsSharedWithMeRequestSharingModePaperVaultSharingModeInternal`    | PAPER_VAULT_SHARING_MODE_INTERNAL                                                               |
| `PaperSecretServiceSearchSecretsSharedWithMeRequestSharingModePaperVaultSharingModeExternal`    | PAPER_VAULT_SHARING_MODE_EXTERNAL                                                               |