# DefaultView

the default view that reviewers will see when they complete their access reviews

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultViewAccessReviewViewTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultView("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `DefaultViewAccessReviewViewTypeUnspecified`  | ACCESS_REVIEW_VIEW_TYPE_UNSPECIFIED           |
| `DefaultViewAccessReviewViewTypeByApp`        | ACCESS_REVIEW_VIEW_TYPE_BY_APP                |
| `DefaultViewAccessReviewViewTypeByUser`       | ACCESS_REVIEW_VIEW_TYPE_BY_USER               |
| `DefaultViewAccessReviewViewTypeUnstructured` | ACCESS_REVIEW_VIEW_TYPE_UNSTRUCTURED          |