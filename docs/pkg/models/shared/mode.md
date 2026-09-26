# Mode

The mode field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ModeAgentClassifierRuleModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Mode("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `ModeAgentClassifierRuleModeUnspecified` | AGENT_CLASSIFIER_RULE_MODE_UNSPECIFIED   |
| `ModeAgentClassifierRuleModeEnforce`     | AGENT_CLASSIFIER_RULE_MODE_ENFORCE       |
| `ModeAgentClassifierRuleModeObserve`     | AGENT_CLASSIFIER_RULE_MODE_OBSERVE       |
| `ModeAgentClassifierRuleModeDisabled`    | AGENT_CLASSIFIER_RULE_MODE_DISABLED      |