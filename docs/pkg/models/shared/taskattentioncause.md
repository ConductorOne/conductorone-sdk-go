# TaskAttentionCause

Why the task cannot progress.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAttentionCauseTaskAttentionCauseUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAttentionCause("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `TaskAttentionCauseTaskAttentionCauseUnspecified`             | TASK_ATTENTION_CAUSE_UNSPECIFIED                              |
| `TaskAttentionCauseTaskAttentionCauseFailed`                  | TASK_ATTENTION_CAUSE_FAILED                                   |
| `TaskAttentionCauseTaskAttentionCauseNoAssignee`              | TASK_ATTENTION_CAUSE_NO_ASSIGNEE                              |
| `TaskAttentionCauseTaskAttentionCauseWaitingOnIntegration`    | TASK_ATTENTION_CAUSE_WAITING_ON_INTEGRATION                   |
| `TaskAttentionCauseTaskAttentionCauseWaitingForAccount`       | TASK_ATTENTION_CAUSE_WAITING_FOR_ACCOUNT                      |
| `TaskAttentionCauseTaskAttentionCauseWaitingOnLock`           | TASK_ATTENTION_CAUSE_WAITING_ON_LOCK                          |
| `TaskAttentionCauseTaskAttentionCauseMultipleAppAccounts`     | TASK_ATTENTION_CAUSE_MULTIPLE_APP_ACCOUNTS                    |
| `TaskAttentionCauseTaskAttentionCauseSubjectDeleted`          | TASK_ATTENTION_CAUSE_SUBJECT_DELETED                          |
| `TaskAttentionCauseTaskAttentionCauseWaitingOnExternalTicket` | TASK_ATTENTION_CAUSE_WAITING_ON_EXTERNAL_TICKET               |