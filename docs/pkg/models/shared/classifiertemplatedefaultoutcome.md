# ClassifierTemplateDefaultOutcome

The defaultOutcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClassifierTemplateDefaultOutcomeAgentClassifierRuleOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClassifierTemplateDefaultOutcome("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ClassifierTemplateDefaultOutcomeAgentClassifierRuleOutcomeUnspecified` | AGENT_CLASSIFIER_RULE_OUTCOME_UNSPECIFIED                               |
| `ClassifierTemplateDefaultOutcomeAgentClassifierRuleOutcomeAllowed`     | AGENT_CLASSIFIER_RULE_OUTCOME_ALLOWED                                   |
| `ClassifierTemplateDefaultOutcomeAgentClassifierRuleOutcomeDenied`      | AGENT_CLASSIFIER_RULE_OUTCOME_DENIED                                    |