# HistoryActorKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.HistoryActorKindActorKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.HistoryActorKind("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `HistoryActorKindActorKindUnspecified` | ACTOR_KIND_UNSPECIFIED                 |
| `HistoryActorKindActorKindAPI`         | ACTOR_KIND_API                         |
| `HistoryActorKindActorKindSlack`       | ACTOR_KIND_SLACK                       |
| `HistoryActorKindActorKindMsteams`     | ACTOR_KIND_MSTEAMS                     |
| `HistoryActorKindActorKindJiraCloud`   | ACTOR_KIND_JIRA_CLOUD                  |
| `HistoryActorKindActorKindInternal`    | ACTOR_KIND_INTERNAL                    |
| `HistoryActorKindActorKindSupport`     | ACTOR_KIND_SUPPORT                     |
| `HistoryActorKindActorKindWorkflow`    | ACTOR_KIND_WORKFLOW                    |