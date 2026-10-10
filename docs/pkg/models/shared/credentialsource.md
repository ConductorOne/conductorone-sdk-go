# CredentialSource

The credentialSource field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialSourceEdgeInferenceCredentialSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialSource("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `CredentialSourceEdgeInferenceCredentialSourceUnspecified`         | EDGE_INFERENCE_CREDENTIAL_SOURCE_UNSPECIFIED                       |
| `CredentialSourceEdgeInferenceCredentialSourceTenantVault`         | EDGE_INFERENCE_CREDENTIAL_SOURCE_TENANT_VAULT                      |
| `CredentialSourceEdgeInferenceCredentialSourceAwsWorkloadIdentity` | EDGE_INFERENCE_CREDENTIAL_SOURCE_AWS_WORKLOAD_IDENTITY             |