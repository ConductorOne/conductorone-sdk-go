# DefaultOutcome

The defaultOutcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultOutcomeAgentClassifierRuleOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultOutcome("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `DefaultOutcomeAgentClassifierRuleOutcomeUnspecified` | AGENT_CLASSIFIER_RULE_OUTCOME_UNSPECIFIED             |
| `DefaultOutcomeAgentClassifierRuleOutcomeAllowed`     | AGENT_CLASSIFIER_RULE_OUTCOME_ALLOWED                 |
| `DefaultOutcomeAgentClassifierRuleOutcomeDenied`      | AGENT_CLASSIFIER_RULE_OUTCOME_DENIED                  |