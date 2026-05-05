# Sentiment

The sentiment field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SentimentA2UISurfaceFeedbackSentimentUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Sentiment("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `SentimentA2UISurfaceFeedbackSentimentUnspecified` | A2UI_SURFACE_FEEDBACK_SENTIMENT_UNSPECIFIED        |
| `SentimentA2UISurfaceFeedbackSentimentPositive`    | A2UI_SURFACE_FEEDBACK_SENTIMENT_POSITIVE           |
| `SentimentA2UISurfaceFeedbackSentimentNegative`    | A2UI_SURFACE_FEEDBACK_SENTIMENT_NEGATIVE           |