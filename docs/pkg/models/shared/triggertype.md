# TriggerType

How this run was initiated (e.g., manual, scheduled).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TriggerTypeTriggerTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TriggerType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `TriggerTypeTriggerTypeUnspecified`      | TRIGGER_TYPE_UNSPECIFIED                 |
| `TriggerTypeTriggerTypeManual`           | TRIGGER_TYPE_MANUAL                      |
| `TriggerTypeTriggerTypeUpliftCompletion` | TRIGGER_TYPE_UPLIFT_COMPLETION           |
| `TriggerTypeTriggerTypeScheduled`        | TRIGGER_TYPE_SCHEDULED                   |
| `TriggerTypeTriggerTypeDirectoryMerge`   | TRIGGER_TYPE_DIRECTORY_MERGE             |