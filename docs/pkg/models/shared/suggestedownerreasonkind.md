# SuggestedOwnerReasonKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SuggestedOwnerReasonKindSuggestedOwnerReasonKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SuggestedOwnerReasonKind("custom_value")
```


## Values

| Name                                                                         | Value                                                                        |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `SuggestedOwnerReasonKindSuggestedOwnerReasonKindUnspecified`                | SUGGESTED_OWNER_REASON_KIND_UNSPECIFIED                                      |
| `SuggestedOwnerReasonKindSuggestedOwnerReasonKindConnectorReportedOwner`     | SUGGESTED_OWNER_REASON_KIND_CONNECTOR_REPORTED_OWNER                         |
| `SuggestedOwnerReasonKindSuggestedOwnerReasonKindGrantOnResource`            | SUGGESTED_OWNER_REASON_KIND_GRANT_ON_RESOURCE                                |
| `SuggestedOwnerReasonKindSuggestedOwnerReasonKindCredentialCreator`          | SUGGESTED_OWNER_REASON_KIND_CREDENTIAL_CREATOR                               |
| `SuggestedOwnerReasonKindSuggestedOwnerReasonKindSiblingNhiOwner`            | SUGGESTED_OWNER_REASON_KIND_SIBLING_NHI_OWNER                                |
| `SuggestedOwnerReasonKindSuggestedOwnerReasonKindAppOwner`                   | SUGGESTED_OWNER_REASON_KIND_APP_OWNER                                        |
| `SuggestedOwnerReasonKindSuggestedOwnerReasonKindManagerOfInactiveCandidate` | SUGGESTED_OWNER_REASON_KIND_MANAGER_OF_INACTIVE_CANDIDATE                    |