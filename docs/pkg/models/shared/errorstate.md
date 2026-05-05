# ErrorState

Error state set when a prepare action fails with a recoverable condition.
 Cleared when the campaign scope is changed.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ErrorStateAccessReviewErrorStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ErrorState("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ErrorStateAccessReviewErrorStateUnspecified`               | ACCESS_REVIEW_ERROR_STATE_UNSPECIFIED                       |
| `ErrorStateAccessReviewErrorStateSelectionQuotaExceedError` | ACCESS_REVIEW_ERROR_STATE_SELECTION_QUOTA_EXCEED_ERROR      |