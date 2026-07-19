# AccessReviewTemplateServiceCreateRequestScopeType

The scopeType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewTemplateServiceCreateRequestScopeTypeAccessReviewScopeTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewTemplateServiceCreateRequestScopeType("custom_value")
```


## Values

| Name                                                                                      | Value                                                                                     |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `AccessReviewTemplateServiceCreateRequestScopeTypeAccessReviewScopeTypeUnspecified`       | ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED                                                      |
| `AccessReviewTemplateServiceCreateRequestScopeTypeAccessReviewScopeTypeByEntitlements`    | ACCESS_REVIEW_SCOPE_TYPE_BY_ENTITLEMENTS                                                  |
| `AccessReviewTemplateServiceCreateRequestScopeTypeAccessReviewScopeTypeByAccessConflicts` | ACCESS_REVIEW_SCOPE_TYPE_BY_ACCESS_CONFLICTS                                              |
| `AccessReviewTemplateServiceCreateRequestScopeTypeAccessReviewScopeTypeByResource`        | ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE                                                      |
| `AccessReviewTemplateServiceCreateRequestScopeTypeAccessReviewScopeTypeByInheritance`     | ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE                                                   |
| `AccessReviewTemplateServiceCreateRequestScopeTypeAccessReviewScopeTypeByUsers`           | ACCESS_REVIEW_SCOPE_TYPE_BY_USERS                                                         |