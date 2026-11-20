# ArtifactVersionState

Publication outcome. PENDING and FAILED are visible only to current Owners.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ArtifactVersionStateArtifactVersionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ArtifactVersionState("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `ArtifactVersionStateArtifactVersionStateUnspecified` | ARTIFACT_VERSION_STATE_UNSPECIFIED                    |
| `ArtifactVersionStateArtifactVersionStatePending`     | ARTIFACT_VERSION_STATE_PENDING                        |
| `ArtifactVersionStateArtifactVersionStateReady`       | ARTIFACT_VERSION_STATE_READY                          |
| `ArtifactVersionStateArtifactVersionStateFailed`      | ARTIFACT_VERSION_STATE_FAILED                         |