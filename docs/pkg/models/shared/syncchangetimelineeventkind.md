# SyncChangeTimelineEventKind

Classification of the observed metadata change.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SyncChangeTimelineEventKindEventKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SyncChangeTimelineEventKind("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `SyncChangeTimelineEventKindEventKindUnspecified`              | EVENT_KIND_UNSPECIFIED                                         |
| `SyncChangeTimelineEventKindEventKindConfigUpdated`            | EVENT_KIND_CONFIG_UPDATED                                      |
| `SyncChangeTimelineEventKindEventKindSyncResourceTypesUpdated` | EVENT_KIND_SYNC_RESOURCE_TYPES_UPDATED                         |
| `SyncChangeTimelineEventKindEventKindReleaseVersionChanged`    | EVENT_KIND_RELEASE_VERSION_CHANGED                             |
| `SyncChangeTimelineEventKindEventKindAuthoredRevisionEvent`    | EVENT_KIND_AUTHORED_REVISION_EVENT                             |
| `SyncChangeTimelineEventKindEventKindActivationBindingSealed`  | EVENT_KIND_ACTIVATION_BINDING_SEALED                           |
| `SyncChangeTimelineEventKindEventKindCurrentSyncDisabled`      | EVENT_KIND_CURRENT_SYNC_DISABLED                               |