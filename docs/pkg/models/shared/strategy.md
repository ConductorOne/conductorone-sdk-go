# Strategy

The strategy field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StrategyEdgeInferenceSelectionStrategyUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Strategy("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `StrategyEdgeInferenceSelectionStrategyUnspecified` | EDGE_INFERENCE_SELECTION_STRATEGY_UNSPECIFIED       |
| `StrategyEdgeInferenceSelectionStrategySingle`      | EDGE_INFERENCE_SELECTION_STRATEGY_SINGLE            |
| `StrategyEdgeInferenceSelectionStrategyFailover`    | EDGE_INFERENCE_SELECTION_STRATEGY_FAILOVER          |
| `StrategyEdgeInferenceSelectionStrategyStageRouter` | EDGE_INFERENCE_SELECTION_STRATEGY_STAGE_ROUTER      |