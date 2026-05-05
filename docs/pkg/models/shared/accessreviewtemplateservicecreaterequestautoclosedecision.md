# AccessReviewTemplateServiceCreateRequestAutoCloseDecision

The autoCloseDecision field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewTemplateServiceCreateRequestAutoCloseDecisionCloseDecisionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewTemplateServiceCreateRequestAutoCloseDecision("custom_value")
```


## Values

| Name                                                                                | Value                                                                               |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `AccessReviewTemplateServiceCreateRequestAutoCloseDecisionCloseDecisionUnspecified` | CLOSE_DECISION_UNSPECIFIED                                                          |
| `AccessReviewTemplateServiceCreateRequestAutoCloseDecisionCloseDecisionRevoked`     | CLOSE_DECISION_REVOKED                                                              |
| `AccessReviewTemplateServiceCreateRequestAutoCloseDecisionCloseDecisionSkip`        | CLOSE_DECISION_SKIP                                                                 |
| `AccessReviewTemplateServiceCreateRequestAutoCloseDecisionCloseDecisionNoAction`    | CLOSE_DECISION_NO_ACTION                                                            |