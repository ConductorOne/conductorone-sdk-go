# AccessReviewServiceCreateRequestScopeType

The type of scoping method for the campaign (e.g., by entitlements, by access conflicts, or by resource).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewServiceCreateRequestScopeType("custom_value")
```


## Values

| Name                                                                              | Value                                                                             |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeUnspecified`       | ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED                                              |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByEntitlements`    | ACCESS_REVIEW_SCOPE_TYPE_BY_ENTITLEMENTS                                          |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByAccessConflicts` | ACCESS_REVIEW_SCOPE_TYPE_BY_ACCESS_CONFLICTS                                      |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByResource`        | ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE                                              |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByInheritance`     | ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE                                           |