# TBEgressRuleMode

The mode field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TBEgressRuleModeTbEgressModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TBEgressRuleMode("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `TBEgressRuleModeTbEgressModeUnspecified` | TB_EGRESS_MODE_UNSPECIFIED                |
| `TBEgressRuleModeTbEgressModeEnforce`     | TB_EGRESS_MODE_ENFORCE                    |
| `TBEgressRuleModeTbEgressModeObserve`     | TB_EGRESS_MODE_OBSERVE                    |
| `TBEgressRuleModeTbEgressModeDisabled`    | TB_EGRESS_MODE_DISABLED                   |