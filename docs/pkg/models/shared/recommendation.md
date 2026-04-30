# Recommendation

The recommendation field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RecommendationInsightRecommendationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Recommendation("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `RecommendationInsightRecommendationUnspecified` | INSIGHT_RECOMMENDATION_UNSPECIFIED               |
| `RecommendationInsightRecommendationApprove`     | INSIGHT_RECOMMENDATION_APPROVE                   |
| `RecommendationInsightRecommendationDeny`        | INSIGHT_RECOMMENDATION_DENY                      |
| `RecommendationInsightRecommendationReview`      | INSIGHT_RECOMMENDATION_REVIEW                    |