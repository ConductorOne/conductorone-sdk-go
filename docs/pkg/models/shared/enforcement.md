# Enforcement

What the Edge enforces for this rule.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EnforcementEdgeEgressRuleEntitlementIssueEnforcementUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Enforcement("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `EnforcementEdgeEgressRuleEntitlementIssueEnforcementUnspecified`    | EDGE_EGRESS_RULE_ENTITLEMENT_ISSUE_ENFORCEMENT_UNSPECIFIED           |
| `EnforcementEdgeEgressRuleEntitlementIssueEnforcementNotEnforced`    | EDGE_EGRESS_RULE_ENTITLEMENT_ISSUE_ENFORCEMENT_NOT_ENFORCED          |
| `EnforcementEdgeEgressRuleEntitlementIssueEnforcementDeniesEveryone` | EDGE_EGRESS_RULE_ENTITLEMENT_ISSUE_ENFORCEMENT_DENIES_EVERYONE       |