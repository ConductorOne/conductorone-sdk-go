# EffectType

The effectType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EffectTypeConnectorActionEffectTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EffectType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `EffectTypeConnectorActionEffectTypeUnspecified` | CONNECTOR_ACTION_EFFECT_TYPE_UNSPECIFIED         |
| `EffectTypeConnectorActionEffectTypeGrant`       | CONNECTOR_ACTION_EFFECT_TYPE_GRANT               |
| `EffectTypeConnectorActionEffectTypeRevoke`      | CONNECTOR_ACTION_EFFECT_TYPE_REVOKE              |