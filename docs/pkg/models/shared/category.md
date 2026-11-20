# Category

The category field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CategoryEdgeFindingCategoryUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Category("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CategoryEdgeFindingCategoryUnspecified`     | EDGE_FINDING_CATEGORY_UNSPECIFIED            |
| `CategoryEdgeFindingCategoryDestinationRisk` | EDGE_FINDING_CATEGORY_DESTINATION_RISK       |
| `CategoryEdgeFindingCategoryPolicyHealth`    | EDGE_FINDING_CATEGORY_POLICY_HEALTH          |
| `CategoryEdgeFindingCategoryDataVolume`      | EDGE_FINDING_CATEGORY_DATA_VOLUME            |
| `CategoryEdgeFindingCategoryIdentity`        | EDGE_FINDING_CATEGORY_IDENTITY               |
| `CategoryEdgeFindingCategoryAiUsage`         | EDGE_FINDING_CATEGORY_AI_USAGE               |
| `CategoryEdgeFindingCategoryTransport`       | EDGE_FINDING_CATEGORY_TRANSPORT              |
| `CategoryEdgeFindingCategoryBehavior`        | EDGE_FINDING_CATEGORY_BEHAVIOR               |
| `CategoryEdgeFindingCategoryEdgeHealth`      | EDGE_FINDING_CATEGORY_EDGE_HEALTH            |