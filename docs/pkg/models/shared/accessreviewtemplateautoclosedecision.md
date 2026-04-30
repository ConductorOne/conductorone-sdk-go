# AccessReviewTemplateAutoCloseDecision

The autoCloseDecision field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewTemplateAutoCloseDecisionCloseDecisionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewTemplateAutoCloseDecision("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AccessReviewTemplateAutoCloseDecisionCloseDecisionUnspecified` | CLOSE_DECISION_UNSPECIFIED                                      |
| `AccessReviewTemplateAutoCloseDecisionCloseDecisionRevoked`     | CLOSE_DECISION_REVOKED                                          |
| `AccessReviewTemplateAutoCloseDecisionCloseDecisionSkip`        | CLOSE_DECISION_SKIP                                             |
| `AccessReviewTemplateAutoCloseDecisionCloseDecisionNoAction`    | CLOSE_DECISION_NO_ACTION                                        |