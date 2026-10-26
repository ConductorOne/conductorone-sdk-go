# EdgeEgressRuleProblemKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeEgressRuleProblemKind("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindUnspecified`             | EDGE_EGRESS_RULE_PROBLEM_KIND_UNSPECIFIED                                   |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindTooManyRules`            | EDGE_EGRESS_RULE_PROBLEM_KIND_TOO_MANY_RULES                                |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindInvalidPattern`          | EDGE_EGRESS_RULE_PROBLEM_KIND_INVALID_PATTERN                               |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindUnknownEntitlement`      | EDGE_EGRESS_RULE_PROBLEM_KIND_UNKNOWN_ENTITLEMENT                           |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindInvalidBodyLimit`        | EDGE_EGRESS_RULE_PROBLEM_KIND_INVALID_BODY_LIMIT                            |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindAuthzenServerRequired`   | EDGE_EGRESS_RULE_PROBLEM_KIND_AUTHZEN_SERVER_REQUIRED                       |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindUnexpectedAuthzenServer` | EDGE_EGRESS_RULE_PROBLEM_KIND_UNEXPECTED_AUTHZEN_SERVER                     |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindUnknownAuthzenServer`    | EDGE_EGRESS_RULE_PROBLEM_KIND_UNKNOWN_AUTHZEN_SERVER                        |
| `EdgeEgressRuleProblemKindEdgeEgressRuleProblemKindInactiveAuthzenServer`   | EDGE_EGRESS_RULE_PROBLEM_KIND_INACTIVE_AUTHZEN_SERVER                       |