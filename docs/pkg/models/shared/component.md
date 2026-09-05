# Component

Where the finding fits in the parsed document.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ComponentComponentUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Component("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ComponentComponentUnspecified`           | COMPONENT_UNSPECIFIED                     |
| `ComponentComponentDocument`              | COMPONENT_DOCUMENT                        |
| `ComponentComponentEntityID`              | COMPONENT_ENTITY_ID                       |
| `ComponentComponentAcsURL`                | COMPONENT_ACS_URL                         |
| `ComponentComponentNameIDFormat`          | COMPONENT_NAME_ID_FORMAT                  |
| `ComponentComponentSigningCertificate`    | COMPONENT_SIGNING_CERTIFICATE             |
| `ComponentComponentEncryptionCertificate` | COMPONENT_ENCRYPTION_CERTIFICATE          |
| `ComponentComponentRequirement`           | COMPONENT_REQUIREMENT                     |
| `ComponentComponentBinding`               | COMPONENT_BINDING                         |