# ApplicableType

BOTH means Agent and gateway. Shared rules use private_data, exfiltration, tool_name, bounded tool_input,
 and user_roles; untrusted_content and agent_id require AGENT applicability.
 Changing applicability revalidates retained rules and cannot invalidate
 existing concrete bindings. Replacing rules in the same update validates
 the merged definition before persistence.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ApplicableTypeClassifierTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ApplicableType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ApplicableTypeClassifierTypeUnspecified` | CLASSIFIER_TYPE_UNSPECIFIED               |
| `ApplicableTypeClassifierTypeAgent`       | CLASSIFIER_TYPE_AGENT                     |
| `ApplicableTypeClassifierTypeGateway`     | CLASSIFIER_TYPE_GATEWAY                   |
| `ApplicableTypeClassifierTypeBoth`        | CLASSIFIER_TYPE_BOTH                      |