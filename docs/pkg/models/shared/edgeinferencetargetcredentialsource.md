# EdgeInferenceTargetCredentialSource

The credentialSource field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeInferenceTargetCredentialSourceEdgeInferenceCredentialSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeInferenceTargetCredentialSource("custom_value")
```


## Values

| Name                                                                                  | Value                                                                                 |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `EdgeInferenceTargetCredentialSourceEdgeInferenceCredentialSourceUnspecified`         | EDGE_INFERENCE_CREDENTIAL_SOURCE_UNSPECIFIED                                          |
| `EdgeInferenceTargetCredentialSourceEdgeInferenceCredentialSourceTenantVault`         | EDGE_INFERENCE_CREDENTIAL_SOURCE_TENANT_VAULT                                         |
| `EdgeInferenceTargetCredentialSourceEdgeInferenceCredentialSourceAwsWorkloadIdentity` | EDGE_INFERENCE_CREDENTIAL_SOURCE_AWS_WORKLOAD_IDENTITY                                |