# AuthoringEventType

Authored transition type, populated for an authored revision event.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthoringEventTypeAuthoredEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthoringEventType("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AuthoringEventTypeAuthoredEventTypeUnspecified`              | AUTHORED_EVENT_TYPE_UNSPECIFIED                               |
| `AuthoringEventTypeAuthoredEventTypeRevisionActivated`        | AUTHORED_EVENT_TYPE_REVISION_ACTIVATED                        |
| `AuthoringEventTypeAuthoredEventTypeRolledBack`               | AUTHORED_EVENT_TYPE_ROLLED_BACK                               |
| `AuthoringEventTypeAuthoredEventTypeSuspended`                | AUTHORED_EVENT_TYPE_SUSPENDED                                 |
| `AuthoringEventTypeAuthoredEventTypeResumed`                  | AUTHORED_EVENT_TYPE_RESUMED                                   |
| `AuthoringEventTypeAuthoredEventTypeRuntimeDeploymentChanged` | AUTHORED_EVENT_TYPE_RUNTIME_DEPLOYMENT_CHANGED                |
| `AuthoringEventTypeAuthoredEventTypeInstallationLost`         | AUTHORED_EVENT_TYPE_INSTALLATION_LOST                         |
| `AuthoringEventTypeAuthoredEventTypeYanked`                   | AUTHORED_EVENT_TYPE_YANKED                                    |