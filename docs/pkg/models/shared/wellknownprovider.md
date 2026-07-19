# WellKnownProvider

Well-known provider type. Drives UX (wizard presets, docs, icons).
 Set at creation time, immutable.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.WellKnownProviderWellKnownWorkloadProviderUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.WellKnownProvider("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `WellKnownProviderWellKnownWorkloadProviderUnspecified`    | WELL_KNOWN_WORKLOAD_PROVIDER_UNSPECIFIED                   |
| `WellKnownProviderWellKnownWorkloadProviderCustom`         | WELL_KNOWN_WORKLOAD_PROVIDER_CUSTOM                        |
| `WellKnownProviderWellKnownWorkloadProviderGithubActions`  | WELL_KNOWN_WORKLOAD_PROVIDER_GITHUB_ACTIONS                |
| `WellKnownProviderWellKnownWorkloadProviderGitlabCi`       | WELL_KNOWN_WORKLOAD_PROVIDER_GITLAB_CI                     |
| `WellKnownProviderWellKnownWorkloadProviderHcpTerraform`   | WELL_KNOWN_WORKLOAD_PROVIDER_HCP_TERRAFORM                 |
| `WellKnownProviderWellKnownWorkloadProviderAwsIamOutbound` | WELL_KNOWN_WORKLOAD_PROVIDER_AWS_IAM_OUTBOUND              |
| `WellKnownProviderWellKnownWorkloadProviderSpiffe`         | WELL_KNOWN_WORKLOAD_PROVIDER_SPIFFE                        |