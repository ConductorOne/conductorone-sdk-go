# AccessReviewTemplateScopeType

The scopeType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewTemplateScopeTypeAccessReviewScopeTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewTemplateScopeType("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AccessReviewTemplateScopeTypeAccessReviewScopeTypeUnspecified`       | ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED                                  |
| `AccessReviewTemplateScopeTypeAccessReviewScopeTypeByEntitlements`    | ACCESS_REVIEW_SCOPE_TYPE_BY_ENTITLEMENTS                              |
| `AccessReviewTemplateScopeTypeAccessReviewScopeTypeByAccessConflicts` | ACCESS_REVIEW_SCOPE_TYPE_BY_ACCESS_CONFLICTS                          |
| `AccessReviewTemplateScopeTypeAccessReviewScopeTypeByResource`        | ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE                                  |
| `AccessReviewTemplateScopeTypeAccessReviewScopeTypeByInheritance`     | ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE                               |