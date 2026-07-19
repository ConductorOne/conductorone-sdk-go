# EventTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EventTypesFindingAuditEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EventTypes("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `EventTypesFindingAuditEventTypeUnspecified`           | FINDING_AUDIT_EVENT_TYPE_UNSPECIFIED                   |
| `EventTypesFindingAuditEventTypeCreated`               | FINDING_AUDIT_EVENT_TYPE_CREATED                       |
| `EventTypesFindingAuditEventTypeStateChanged`          | FINDING_AUDIT_EVENT_TYPE_STATE_CHANGED                 |
| `EventTypesFindingAuditEventTypeSnoozed`               | FINDING_AUDIT_EVENT_TYPE_SNOOZED                       |
| `EventTypesFindingAuditEventTypeSnoozeExpired`         | FINDING_AUDIT_EVENT_TYPE_SNOOZE_EXPIRED                |
| `EventTypesFindingAuditEventTypeRiskAccepted`          | FINDING_AUDIT_EVENT_TYPE_RISK_ACCEPTED                 |
| `EventTypesFindingAuditEventTypeRiskAcceptanceExpired` | FINDING_AUDIT_EVENT_TYPE_RISK_ACCEPTANCE_EXPIRED       |
| `EventTypesFindingAuditEventTypeSuppressed`            | FINDING_AUDIT_EVENT_TYPE_SUPPRESSED                    |
| `EventTypesFindingAuditEventTypeUnsuppressed`          | FINDING_AUDIT_EVENT_TYPE_UNSUPPRESSED                  |
| `EventTypesFindingAuditEventTypeResolved`              | FINDING_AUDIT_EVENT_TYPE_RESOLVED                      |
| `EventTypesFindingAuditEventTypeReopened`              | FINDING_AUDIT_EVENT_TYPE_REOPENED                      |
| `EventTypesFindingAuditEventTypeOwnerChanged`          | FINDING_AUDIT_EVENT_TYPE_OWNER_CHANGED                 |
| `EventTypesFindingAuditEventTypeSeverityOverridden`    | FINDING_AUDIT_EVENT_TYPE_SEVERITY_OVERRIDDEN           |
| `EventTypesFindingAuditEventTypeComment`               | FINDING_AUDIT_EVENT_TYPE_COMMENT                       |
| `EventTypesFindingAuditEventTypeTaskCreated`           | FINDING_AUDIT_EVENT_TYPE_TASK_CREATED                  |
| `EventTypesFindingAuditEventTypeTaskCancelled`         | FINDING_AUDIT_EVENT_TYPE_TASK_CANCELLED                |
| `EventTypesFindingAuditEventTypeEvidenceUpdated`       | FINDING_AUDIT_EVENT_TYPE_EVIDENCE_UPDATED              |
| `EventTypesFindingAuditEventTypeRoutingEvaluated`      | FINDING_AUDIT_EVENT_TYPE_ROUTING_EVALUATED             |
| `EventTypesFindingAuditEventTypeTransformed`           | FINDING_AUDIT_EVENT_TYPE_TRANSFORMED                   |