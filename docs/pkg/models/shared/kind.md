# Kind

The type of credential.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.KindAPIKeyKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Kind("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `KindAPIKeyKindUnspecified`        | API_KEY_KIND_UNSPECIFIED           |
| `KindAPIKeyKindPersonal`           | API_KEY_KIND_PERSONAL              |
| `KindAPIKeyKindDevice`             | API_KEY_KIND_DEVICE                |
| `KindAPIKeyKindServicePrincipal`   | API_KEY_KIND_SERVICE_PRINCIPAL     |
| `KindAPIKeyKindWorkloadFederation` | API_KEY_KIND_WORKLOAD_FEDERATION   |
| `KindAPIKeyKindConnectedApp`       | API_KEY_KIND_CONNECTED_APP         |
| `KindAPIKeyKindConnector`          | API_KEY_KIND_CONNECTOR             |
| `KindAPIKeyKindFunction`           | API_KEY_KIND_FUNCTION              |
| `KindAPIKeyKindOtherSystem`        | API_KEY_KIND_OTHER_SYSTEM          |