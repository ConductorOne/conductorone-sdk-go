# AttentionCauses

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AttentionCausesTaskAttentionCauseUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AttentionCauses("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AttentionCausesTaskAttentionCauseUnspecified`             | TASK_ATTENTION_CAUSE_UNSPECIFIED                           |
| `AttentionCausesTaskAttentionCauseFailed`                  | TASK_ATTENTION_CAUSE_FAILED                                |
| `AttentionCausesTaskAttentionCauseNoAssignee`              | TASK_ATTENTION_CAUSE_NO_ASSIGNEE                           |
| `AttentionCausesTaskAttentionCauseWaitingOnIntegration`    | TASK_ATTENTION_CAUSE_WAITING_ON_INTEGRATION                |
| `AttentionCausesTaskAttentionCauseWaitingForAccount`       | TASK_ATTENTION_CAUSE_WAITING_FOR_ACCOUNT                   |
| `AttentionCausesTaskAttentionCauseWaitingOnLock`           | TASK_ATTENTION_CAUSE_WAITING_ON_LOCK                       |
| `AttentionCausesTaskAttentionCauseMultipleAppAccounts`     | TASK_ATTENTION_CAUSE_MULTIPLE_APP_ACCOUNTS                 |
| `AttentionCausesTaskAttentionCauseSubjectDeleted`          | TASK_ATTENTION_CAUSE_SUBJECT_DELETED                       |
| `AttentionCausesTaskAttentionCauseWaitingOnExternalTicket` | TASK_ATTENTION_CAUSE_WAITING_ON_EXTERNAL_TICKET            |