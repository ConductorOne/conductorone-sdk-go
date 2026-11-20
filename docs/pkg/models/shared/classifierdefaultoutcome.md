# ClassifierDefaultOutcome

The defaultOutcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClassifierDefaultOutcomeAgentClassifierRuleOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClassifierDefaultOutcome("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ClassifierDefaultOutcomeAgentClassifierRuleOutcomeUnspecified` | AGENT_CLASSIFIER_RULE_OUTCOME_UNSPECIFIED                       |
| `ClassifierDefaultOutcomeAgentClassifierRuleOutcomeAllowed`     | AGENT_CLASSIFIER_RULE_OUTCOME_ALLOWED                           |
| `ClassifierDefaultOutcomeAgentClassifierRuleOutcomeDenied`      | AGENT_CLASSIFIER_RULE_OUTCOME_DENIED                            |