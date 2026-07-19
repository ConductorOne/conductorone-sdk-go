# SessionPolicyPolicyRuleMode

The mode field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyPolicyRuleModePolicyRuleModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyPolicyRuleMode("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `SessionPolicyPolicyRuleModePolicyRuleModeUnspecified` | POLICY_RULE_MODE_UNSPECIFIED                           |
| `SessionPolicyPolicyRuleModePolicyRuleModeEnforce`     | POLICY_RULE_MODE_ENFORCE                               |
| `SessionPolicyPolicyRuleModePolicyRuleModeObserve`     | POLICY_RULE_MODE_OBSERVE                               |
| `SessionPolicyPolicyRuleModePolicyRuleModeDisabled`    | POLICY_RULE_MODE_DISABLED                              |