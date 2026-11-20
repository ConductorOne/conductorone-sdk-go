# EdgeServiceCheckInferenceVaultAccessResponseReason

The reason field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeServiceCheckInferenceVaultAccessResponseReasonEdgeInferenceVaultAccessReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeServiceCheckInferenceVaultAccessResponseReason("custom_value")
```


## Values

| Name                                                                                                 | Value                                                                                                |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `EdgeServiceCheckInferenceVaultAccessResponseReasonEdgeInferenceVaultAccessReasonUnspecified`        | EDGE_INFERENCE_VAULT_ACCESS_REASON_UNSPECIFIED                                                       |
| `EdgeServiceCheckInferenceVaultAccessResponseReasonEdgeInferenceVaultAccessReasonGranted`            | EDGE_INFERENCE_VAULT_ACCESS_REASON_GRANTED                                                           |
| `EdgeServiceCheckInferenceVaultAccessResponseReasonEdgeInferenceVaultAccessReasonNotGranted`         | EDGE_INFERENCE_VAULT_ACCESS_REASON_NOT_GRANTED                                                       |
| `EdgeServiceCheckInferenceVaultAccessResponseReasonEdgeInferenceVaultAccessReasonNoRuntimePrincipal` | EDGE_INFERENCE_VAULT_ACCESS_REASON_NO_RUNTIME_PRINCIPAL                                              |
| `EdgeServiceCheckInferenceVaultAccessResponseReasonEdgeInferenceVaultAccessReasonVaultNotFound`      | EDGE_INFERENCE_VAULT_ACCESS_REASON_VAULT_NOT_FOUND                                                   |