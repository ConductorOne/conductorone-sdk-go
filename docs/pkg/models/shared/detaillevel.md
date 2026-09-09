# DetailLevel

How much the notification reveals. Defaults to SUMMARY.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DetailLevelFindingNotifyDetailLevelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DetailLevel("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `DetailLevelFindingNotifyDetailLevelUnspecified` | FINDING_NOTIFY_DETAIL_LEVEL_UNSPECIFIED          |
| `DetailLevelFindingNotifyDetailLevelSummary`     | FINDING_NOTIFY_DETAIL_LEVEL_SUMMARY              |
| `DetailLevelFindingNotifyDetailLevelFullDetail`  | FINDING_NOTIFY_DETAIL_LEVEL_FULL_DETAIL          |