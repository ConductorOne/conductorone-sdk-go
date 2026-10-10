# EdgeEgressRuleWarningKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeEgressRuleWarningKindEdgeEgressRuleWarningKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeEgressRuleWarningKind("custom_value")
```


## Values

| Name                                                                                | Value                                                                               |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `EdgeEgressRuleWarningKindEdgeEgressRuleWarningKindUnspecified`                     | EDGE_EGRESS_RULE_WARNING_KIND_UNSPECIFIED                                           |
| `EdgeEgressRuleWarningKindEdgeEgressRuleWarningKindAuthzenEvaluateGrantMissing`     | EDGE_EGRESS_RULE_WARNING_KIND_AUTHZEN_EVALUATE_GRANT_MISSING                        |
| `EdgeEgressRuleWarningKindEdgeEgressRuleWarningKindAuthzenNoEdgePrincipal`          | EDGE_EGRESS_RULE_WARNING_KIND_AUTHZEN_NO_EDGE_PRINCIPAL                             |
| `EdgeEgressRuleWarningKindEdgeEgressRuleWarningKindHostDecisionBypassesAuthzenRule` | EDGE_EGRESS_RULE_WARNING_KIND_HOST_DECISION_BYPASSES_AUTHZEN_RULE                   |