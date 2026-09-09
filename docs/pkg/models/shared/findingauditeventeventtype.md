# FindingAuditEventEventType

The eventType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingAuditEventEventTypeFindingAuditEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingAuditEventEventType("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `FindingAuditEventEventTypeFindingAuditEventTypeUnspecified`           | FINDING_AUDIT_EVENT_TYPE_UNSPECIFIED                                   |
| `FindingAuditEventEventTypeFindingAuditEventTypeCreated`               | FINDING_AUDIT_EVENT_TYPE_CREATED                                       |
| `FindingAuditEventEventTypeFindingAuditEventTypeStateChanged`          | FINDING_AUDIT_EVENT_TYPE_STATE_CHANGED                                 |
| `FindingAuditEventEventTypeFindingAuditEventTypeSnoozed`               | FINDING_AUDIT_EVENT_TYPE_SNOOZED                                       |
| `FindingAuditEventEventTypeFindingAuditEventTypeSnoozeExpired`         | FINDING_AUDIT_EVENT_TYPE_SNOOZE_EXPIRED                                |
| `FindingAuditEventEventTypeFindingAuditEventTypeRiskAccepted`          | FINDING_AUDIT_EVENT_TYPE_RISK_ACCEPTED                                 |
| `FindingAuditEventEventTypeFindingAuditEventTypeRiskAcceptanceExpired` | FINDING_AUDIT_EVENT_TYPE_RISK_ACCEPTANCE_EXPIRED                       |
| `FindingAuditEventEventTypeFindingAuditEventTypeSuppressed`            | FINDING_AUDIT_EVENT_TYPE_SUPPRESSED                                    |
| `FindingAuditEventEventTypeFindingAuditEventTypeUnsuppressed`          | FINDING_AUDIT_EVENT_TYPE_UNSUPPRESSED                                  |
| `FindingAuditEventEventTypeFindingAuditEventTypeResolved`              | FINDING_AUDIT_EVENT_TYPE_RESOLVED                                      |
| `FindingAuditEventEventTypeFindingAuditEventTypeReopened`              | FINDING_AUDIT_EVENT_TYPE_REOPENED                                      |
| `FindingAuditEventEventTypeFindingAuditEventTypeOwnerChanged`          | FINDING_AUDIT_EVENT_TYPE_OWNER_CHANGED                                 |
| `FindingAuditEventEventTypeFindingAuditEventTypeSeverityOverridden`    | FINDING_AUDIT_EVENT_TYPE_SEVERITY_OVERRIDDEN                           |
| `FindingAuditEventEventTypeFindingAuditEventTypeComment`               | FINDING_AUDIT_EVENT_TYPE_COMMENT                                       |
| `FindingAuditEventEventTypeFindingAuditEventTypeTaskCreated`           | FINDING_AUDIT_EVENT_TYPE_TASK_CREATED                                  |
| `FindingAuditEventEventTypeFindingAuditEventTypeTaskCancelled`         | FINDING_AUDIT_EVENT_TYPE_TASK_CANCELLED                                |
| `FindingAuditEventEventTypeFindingAuditEventTypeEvidenceUpdated`       | FINDING_AUDIT_EVENT_TYPE_EVIDENCE_UPDATED                              |
| `FindingAuditEventEventTypeFindingAuditEventTypeRoutingEvaluated`      | FINDING_AUDIT_EVENT_TYPE_ROUTING_EVALUATED                             |
| `FindingAuditEventEventTypeFindingAuditEventTypeTransformed`           | FINDING_AUDIT_EVENT_TYPE_TRANSFORMED                                   |
| `FindingAuditEventEventTypeFindingAuditEventTypeReprocessRequested`    | FINDING_AUDIT_EVENT_TYPE_REPROCESS_REQUESTED                           |
| `FindingAuditEventEventTypeFindingAuditEventTypeReprocessCompleted`    | FINDING_AUDIT_EVENT_TYPE_REPROCESS_COMPLETED                           |
| `FindingAuditEventEventTypeFindingAuditEventTypeAssigneeChanged`       | FINDING_AUDIT_EVENT_TYPE_ASSIGNEE_CHANGED                              |