# AccessReviewState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewStateAccessReviewStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewState("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `AccessReviewStateAccessReviewStateUnspecified`                   | ACCESS_REVIEW_STATE_UNSPECIFIED                                   |
| `AccessReviewStateAccessReviewStateOpen`                          | ACCESS_REVIEW_STATE_OPEN                                          |
| `AccessReviewStateAccessReviewStateClosed`                        | ACCESS_REVIEW_STATE_CLOSED                                        |
| `AccessReviewStateAccessReviewStatePending`                       | ACCESS_REVIEW_STATE_PENDING                                       |
| `AccessReviewStateAccessReviewStateReview`                        | ACCESS_REVIEW_STATE_REVIEW                                        |
| `AccessReviewStateAccessReviewStatePreparing`                     | ACCESS_REVIEW_STATE_PREPARING                                     |
| `AccessReviewStateAccessReviewStateStarting`                      | ACCESS_REVIEW_STATE_STARTING                                      |
| `AccessReviewStateAccessReviewStateDraft`                         | ACCESS_REVIEW_STATE_DRAFT                                         |
| `AccessReviewStateAccessReviewStateDeleting`                      | ACCESS_REVIEW_STATE_DELETING                                      |
| `AccessReviewStateAccessReviewStateDeleted`                       | ACCESS_REVIEW_STATE_DELETED                                       |
| `AccessReviewStateAccessReviewStateResettingPolicies`             | ACCESS_REVIEW_STATE_RESETTING_POLICIES                            |
| `AccessReviewStateAccessReviewStateCopyingSetupEntitlements`      | ACCESS_REVIEW_STATE_COPYING_SETUP_ENTITLEMENTS                    |
| `AccessReviewStateAccessReviewStateCopyingResourceTypeSelections` | ACCESS_REVIEW_STATE_COPYING_RESOURCE_TYPE_SELECTIONS              |