# RecordType

Not always the step's own type: a step over grants can be narrowed to one
 app, and the app is the record worth naming.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RecordTypeA2UIProvenanceRecordTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RecordType("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `RecordTypeA2UIProvenanceRecordTypeUnspecified`             | A2UI_PROVENANCE_RECORD_TYPE_UNSPECIFIED                     |
| `RecordTypeA2UIProvenanceRecordTypeApp`                     | A2UI_PROVENANCE_RECORD_TYPE_APP                             |
| `RecordTypeA2UIProvenanceRecordTypeUser`                    | A2UI_PROVENANCE_RECORD_TYPE_USER                            |
| `RecordTypeA2UIProvenanceRecordTypeGrant`                   | A2UI_PROVENANCE_RECORD_TYPE_GRANT                           |
| `RecordTypeA2UIProvenanceRecordTypeAppEntitlement`          | A2UI_PROVENANCE_RECORD_TYPE_APP_ENTITLEMENT                 |
| `RecordTypeA2UIProvenanceRecordTypeAppUser`                 | A2UI_PROVENANCE_RECORD_TYPE_APP_USER                        |
| `RecordTypeA2UIProvenanceRecordTypeAppResource`             | A2UI_PROVENANCE_RECORD_TYPE_APP_RESOURCE                    |
| `RecordTypeA2UIProvenanceRecordTypeAppResourceType`         | A2UI_PROVENANCE_RECORD_TYPE_APP_RESOURCE_TYPE               |
| `RecordTypeA2UIProvenanceRecordTypeTask`                    | A2UI_PROVENANCE_RECORD_TYPE_TASK                            |
| `RecordTypeA2UIProvenanceRecordTypePolicy`                  | A2UI_PROVENANCE_RECORD_TYPE_POLICY                          |
| `RecordTypeA2UIProvenanceRecordTypeConnector`               | A2UI_PROVENANCE_RECORD_TYPE_CONNECTOR                       |
| `RecordTypeA2UIProvenanceRecordTypeAccessReview`            | A2UI_PROVENANCE_RECORD_TYPE_ACCESS_REVIEW                   |
| `RecordTypeA2UIProvenanceRecordTypeAccessReviewTemplate`    | A2UI_PROVENANCE_RECORD_TYPE_ACCESS_REVIEW_TEMPLATE          |
| `RecordTypeA2UIProvenanceRecordTypeAccessReviewSelection`   | A2UI_PROVENANCE_RECORD_TYPE_ACCESS_REVIEW_SELECTION         |
| `RecordTypeA2UIProvenanceRecordTypeConflictMonitor`         | A2UI_PROVENANCE_RECORD_TYPE_CONFLICT_MONITOR                |
| `RecordTypeA2UIProvenanceRecordTypeAccessViolation`         | A2UI_PROVENANCE_RECORD_TYPE_ACCESS_VIOLATION                |
| `RecordTypeA2UIProvenanceRecordTypeRequestCatalog`          | A2UI_PROVENANCE_RECORD_TYPE_REQUEST_CATALOG                 |
| `RecordTypeA2UIProvenanceRecordTypeWebhook`                 | A2UI_PROVENANCE_RECORD_TYPE_WEBHOOK                         |
| `RecordTypeA2UIProvenanceRecordTypeDirectory`               | A2UI_PROVENANCE_RECORD_TYPE_DIRECTORY                       |
| `RecordTypeA2UIProvenanceRecordTypeProfileType`             | A2UI_PROVENANCE_RECORD_TYPE_PROFILE_TYPE                    |
| `RecordTypeA2UIProvenanceRecordTypeRoleBinding`             | A2UI_PROVENANCE_RECORD_TYPE_ROLE_BINDING                    |
| `RecordTypeA2UIProvenanceRecordTypeAutomationExecution`     | A2UI_PROVENANCE_RECORD_TYPE_AUTOMATION_EXECUTION            |
| `RecordTypeA2UIProvenanceRecordTypeAutomationExecutionStep` | A2UI_PROVENANCE_RECORD_TYPE_AUTOMATION_EXECUTION_STEP       |
| `RecordTypeA2UIProvenanceRecordTypeFinding`                 | A2UI_PROVENANCE_RECORD_TYPE_FINDING                         |
| `RecordTypeA2UIProvenanceRecordTypeMetric`                  | A2UI_PROVENANCE_RECORD_TYPE_METRIC                          |
| `RecordTypeA2UIProvenanceRecordTypeAutomation`              | A2UI_PROVENANCE_RECORD_TYPE_AUTOMATION                      |
| `RecordTypeA2UIProvenanceRecordTypeGrantHistory`            | A2UI_PROVENANCE_RECORD_TYPE_GRANT_HISTORY                   |
| `RecordTypeA2UIProvenanceRecordTypeGrantReason`             | A2UI_PROVENANCE_RECORD_TYPE_GRANT_REASON                    |
| `RecordTypeA2UIProvenanceRecordTypeAppOwner`                | A2UI_PROVENANCE_RECORD_TYPE_APP_OWNER                       |