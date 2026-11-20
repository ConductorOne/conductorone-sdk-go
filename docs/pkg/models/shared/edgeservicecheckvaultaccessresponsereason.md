# EdgeServiceCheckVaultAccessResponseReason

Explanation of the current grant state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeServiceCheckVaultAccessResponseReasonEdgeVaultAccessReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeServiceCheckVaultAccessResponseReason("custom_value")
```


## Values

| Name                                                                               | Value                                                                              |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `EdgeServiceCheckVaultAccessResponseReasonEdgeVaultAccessReasonUnspecified`        | EDGE_VAULT_ACCESS_REASON_UNSPECIFIED                                               |
| `EdgeServiceCheckVaultAccessResponseReasonEdgeVaultAccessReasonGranted`            | EDGE_VAULT_ACCESS_REASON_GRANTED                                                   |
| `EdgeServiceCheckVaultAccessResponseReasonEdgeVaultAccessReasonNotGranted`         | EDGE_VAULT_ACCESS_REASON_NOT_GRANTED                                               |
| `EdgeServiceCheckVaultAccessResponseReasonEdgeVaultAccessReasonNoRuntimePrincipal` | EDGE_VAULT_ACCESS_REASON_NO_RUNTIME_PRINCIPAL                                      |
| `EdgeServiceCheckVaultAccessResponseReasonEdgeVaultAccessReasonVaultNotFound`      | EDGE_VAULT_ACCESS_REASON_VAULT_NOT_FOUND                                           |