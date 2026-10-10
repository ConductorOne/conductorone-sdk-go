# CredentialKind

The credentialKind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialKindEdgeInferenceCatalogCredentialKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialKind("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `CredentialKindEdgeInferenceCatalogCredentialKindUnspecified`               | EDGE_INFERENCE_CATALOG_CREDENTIAL_KIND_UNSPECIFIED                          |
| `CredentialKindEdgeInferenceCatalogCredentialKindAPIKey`                    | EDGE_INFERENCE_CATALOG_CREDENTIAL_KIND_API_KEY                              |
| `CredentialKindEdgeInferenceCatalogCredentialKindAwsWorkloadIdentityRegion` | EDGE_INFERENCE_CATALOG_CREDENTIAL_KIND_AWS_WORKLOAD_IDENTITY_REGION         |