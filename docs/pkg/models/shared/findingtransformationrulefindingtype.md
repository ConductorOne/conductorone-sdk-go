# FindingTransformationRuleFindingType

The findingType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingTransformationRuleFindingTypeFindingTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingTransformationRuleFindingType("custom_value")
```


## Values

| Name                                                                               | Value                                                                              |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `FindingTransformationRuleFindingTypeFindingTypeUnspecified`                       | FINDING_TYPE_UNSPECIFIED                                                           |
| `FindingTransformationRuleFindingTypeFindingTypeSimilarUsernameMatch`              | FINDING_TYPE_SIMILAR_USERNAME_MATCH                                                |
| `FindingTransformationRuleFindingTypeFindingTypeServiceAccountMisclassification`   | FINDING_TYPE_SERVICE_ACCOUNT_MISCLASSIFICATION                                     |
| `FindingTransformationRuleFindingTypeFindingTypeNhiUnowned`                        | FINDING_TYPE_NHI_UNOWNED                                                           |
| `FindingTransformationRuleFindingTypeFindingTypeServiceAccountUnowned`             | FINDING_TYPE_SERVICE_ACCOUNT_UNOWNED                                               |
| `FindingTransformationRuleFindingTypeFindingTypeDecoyCredentialUsed`               | FINDING_TYPE_DECOY_CREDENTIAL_USED                                                 |
| `FindingTransformationRuleFindingTypeFindingTypeCustom`                            | FINDING_TYPE_CUSTOM                                                                |
| `FindingTransformationRuleFindingTypeFindingTypeConnectorAnomalyDetectionDisabled` | FINDING_TYPE_CONNECTOR_ANOMALY_DETECTION_DISABLED                                  |
| `FindingTransformationRuleFindingTypeFindingTypeDeactivatedOwner`                  | FINDING_TYPE_DEACTIVATED_OWNER                                                     |
| `FindingTransformationRuleFindingTypeFindingTypeUnusedSecret`                      | FINDING_TYPE_UNUSED_SECRET                                                         |
| `FindingTransformationRuleFindingTypeFindingTypeCredentialPubliclyExposed`         | FINDING_TYPE_CREDENTIAL_PUBLICLY_EXPOSED                                           |
| `FindingTransformationRuleFindingTypeFindingTypeDecoyPubliclyExposed`              | FINDING_TYPE_DECOY_PUBLICLY_EXPOSED                                                |
| `FindingTransformationRuleFindingTypeFindingTypeCredentialExpiring`                | FINDING_TYPE_CREDENTIAL_EXPIRING                                                   |
| `FindingTransformationRuleFindingTypeFindingTypeConnectorSyncFailing`              | FINDING_TYPE_CONNECTOR_SYNC_FAILING                                                |
| `FindingTransformationRuleFindingTypeFindingTypeShadowMcp`                         | FINDING_TYPE_SHADOW_MCP                                                            |