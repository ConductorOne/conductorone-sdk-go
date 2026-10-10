# Outcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OutcomeAgentClassifierRuleOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Outcome("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `OutcomeAgentClassifierRuleOutcomeUnspecified` | AGENT_CLASSIFIER_RULE_OUTCOME_UNSPECIFIED      |
| `OutcomeAgentClassifierRuleOutcomeAllowed`     | AGENT_CLASSIFIER_RULE_OUTCOME_ALLOWED          |
| `OutcomeAgentClassifierRuleOutcomeDenied`      | AGENT_CLASSIFIER_RULE_OUTCOME_DENIED           |