# EdgeEgressHostVerdictOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeEgressHostVerdictOutcome("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeUnspecified`        | EDGE_EGRESS_EVALUATION_OUTCOME_UNSPECIFIED                                  |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeNotRestricted`      | EDGE_EGRESS_EVALUATION_OUTCOME_NOT_RESTRICTED                               |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeMatchedRule`        | EDGE_EGRESS_EVALUATION_OUTCOME_MATCHED_RULE                                 |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeDefaultDeny`        | EDGE_EGRESS_EVALUATION_OUTCOME_DEFAULT_DENY                                 |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeInvalidHost`        | EDGE_EGRESS_EVALUATION_OUTCOME_INVALID_HOST                                 |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeUnresolvedDenyRule` | EDGE_EGRESS_EVALUATION_OUTCOME_UNRESOLVED_DENY_RULE                         |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeSsrfBlocked`        | EDGE_EGRESS_EVALUATION_OUTCOME_SSRF_BLOCKED                                 |
| `EdgeEgressHostVerdictOutcomeEdgeEgressEvaluationOutcomeDelegatedToAuthzen` | EDGE_EGRESS_EVALUATION_OUTCOME_DELEGATED_TO_AUTHZEN                         |