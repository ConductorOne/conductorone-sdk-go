# VaultType

Vault type selector (default: PAPER_VAULT for backward compatibility)

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.VaultTypeStoreCredentialVaultTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.VaultType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `VaultTypeStoreCredentialVaultTypeUnspecified` | STORE_CREDENTIAL_VAULT_TYPE_UNSPECIFIED        |
| `VaultTypeStoreCredentialVaultTypePaperVault`  | STORE_CREDENTIAL_VAULT_TYPE_PAPER_VAULT        |
| `VaultTypeStoreCredentialVaultTypeAppVault`    | STORE_CREDENTIAL_VAULT_TYPE_APP_VAULT          |