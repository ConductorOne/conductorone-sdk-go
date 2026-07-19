# DecoyCredentialUsedTypeKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DecoyCredentialUsedTypeKindDecoyCredentialKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DecoyCredentialUsedTypeKind("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `DecoyCredentialUsedTypeKindDecoyCredentialKindUnspecified`          | DECOY_CREDENTIAL_KIND_UNSPECIFIED                                    |
| `DecoyCredentialUsedTypeKindDecoyCredentialKindUserClientCredential` | DECOY_CREDENTIAL_KIND_USER_CLIENT_CREDENTIAL                         |
| `DecoyCredentialUsedTypeKindDecoyCredentialKindConnectorClient`      | DECOY_CREDENTIAL_KIND_CONNECTOR_CLIENT                               |
| `DecoyCredentialUsedTypeKindDecoyCredentialKindWorkloadFederation`   | DECOY_CREDENTIAL_KIND_WORKLOAD_FEDERATION                            |
| `DecoyCredentialUsedTypeKindDecoyCredentialKindAccessToken`          | DECOY_CREDENTIAL_KIND_ACCESS_TOKEN                                   |