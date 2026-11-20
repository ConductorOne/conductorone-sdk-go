# ExcludeFindingTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExcludeFindingTypesFindingTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExcludeFindingTypes("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ExcludeFindingTypesFindingTypeUnspecified`                       | FINDING_TYPE_UNSPECIFIED                                          |
| `ExcludeFindingTypesFindingTypeSimilarUsernameMatch`              | FINDING_TYPE_SIMILAR_USERNAME_MATCH                               |
| `ExcludeFindingTypesFindingTypeServiceAccountMisclassification`   | FINDING_TYPE_SERVICE_ACCOUNT_MISCLASSIFICATION                    |
| `ExcludeFindingTypesFindingTypeNhiUnowned`                        | FINDING_TYPE_NHI_UNOWNED                                          |
| `ExcludeFindingTypesFindingTypeServiceAccountUnowned`             | FINDING_TYPE_SERVICE_ACCOUNT_UNOWNED                              |
| `ExcludeFindingTypesFindingTypeDecoyCredentialUsed`               | FINDING_TYPE_DECOY_CREDENTIAL_USED                                |
| `ExcludeFindingTypesFindingTypeCustom`                            | FINDING_TYPE_CUSTOM                                               |
| `ExcludeFindingTypesFindingTypeConnectorAnomalyDetectionDisabled` | FINDING_TYPE_CONNECTOR_ANOMALY_DETECTION_DISABLED                 |
| `ExcludeFindingTypesFindingTypeDeactivatedOwner`                  | FINDING_TYPE_DEACTIVATED_OWNER                                    |
| `ExcludeFindingTypesFindingTypeUnusedSecret`                      | FINDING_TYPE_UNUSED_SECRET                                        |
| `ExcludeFindingTypesFindingTypeCredentialPubliclyExposed`         | FINDING_TYPE_CREDENTIAL_PUBLICLY_EXPOSED                          |
| `ExcludeFindingTypesFindingTypeDecoyPubliclyExposed`              | FINDING_TYPE_DECOY_PUBLICLY_EXPOSED                               |
| `ExcludeFindingTypesFindingTypeCredentialExpiring`                | FINDING_TYPE_CREDENTIAL_EXPIRING                                  |
| `ExcludeFindingTypesFindingTypeConnectorSyncFailing`              | FINDING_TYPE_CONNECTOR_SYNC_FAILING                               |
| `ExcludeFindingTypesFindingTypeShadowMcp`                         | FINDING_TYPE_SHADOW_MCP                                           |
| `ExcludeFindingTypesFindingTypeShadowApp`                         | FINDING_TYPE_SHADOW_APP                                           |
| `ExcludeFindingTypesFindingTypeMcpGatewayToolCallRisk`            | FINDING_TYPE_MCP_GATEWAY_TOOL_CALL_RISK                           |
| `ExcludeFindingTypesFindingTypeEdge`                              | FINDING_TYPE_EDGE                                                 |
| `ExcludeFindingTypesFindingTypeAccessConflict`                    | FINDING_TYPE_ACCESS_CONFLICT                                      |
| `ExcludeFindingTypesFindingTypeConnectorHostOutdated`             | FINDING_TYPE_CONNECTOR_HOST_OUTDATED                              |