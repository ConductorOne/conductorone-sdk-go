# RecordState

The recordState field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RecordStateFindingRecordStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RecordState("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `RecordStateFindingRecordStateUnspecified` | FINDING_RECORD_STATE_UNSPECIFIED           |
| `RecordStateFindingRecordStateActive`      | FINDING_RECORD_STATE_ACTIVE                |
| `RecordStateFindingRecordStateArchived`    | FINDING_RECORD_STATE_ARCHIVED              |