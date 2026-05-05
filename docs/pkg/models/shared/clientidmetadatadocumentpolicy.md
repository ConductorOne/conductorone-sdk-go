# ClientIDMetadataDocumentPolicy

Policy for metadata document client_id URLs.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClientIDMetadataDocumentPolicyClientIDMetadataDocumentPolicyUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClientIDMetadataDocumentPolicy("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ClientIDMetadataDocumentPolicyClientIDMetadataDocumentPolicyUnspecified`   | CLIENT_ID_METADATA_DOCUMENT_POLICY_UNSPECIFIED                              |
| `ClientIDMetadataDocumentPolicyClientIDMetadataDocumentPolicyAllowAll`      | CLIENT_ID_METADATA_DOCUMENT_POLICY_ALLOW_ALL                                |
| `ClientIDMetadataDocumentPolicyClientIDMetadataDocumentPolicyRequestable`   | CLIENT_ID_METADATA_DOCUMENT_POLICY_REQUESTABLE                              |
| `ClientIDMetadataDocumentPolicyClientIDMetadataDocumentPolicyAllowlistOnly` | CLIENT_ID_METADATA_DOCUMENT_POLICY_ALLOWLIST_ONLY                           |