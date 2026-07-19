# PolicyRuleMode

Whether the rule is live, evaluated-only, or skipped.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PolicyRuleModePolicyRuleModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PolicyRuleMode("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `PolicyRuleModePolicyRuleModeUnspecified` | POLICY_RULE_MODE_UNSPECIFIED              |
| `PolicyRuleModePolicyRuleModeEnforce`     | POLICY_RULE_MODE_ENFORCE                  |
| `PolicyRuleModePolicyRuleModeObserve`     | POLICY_RULE_MODE_OBSERVE                  |
| `PolicyRuleModePolicyRuleModeDisabled`    | POLICY_RULE_MODE_DISABLED                 |