# WorkloadFederationServiceCreateProviderRequestWellKnownProvider

Well-known provider type. Required -- UNSPECIFIED is rejected.
 When set to a named source, the backend validates issuer_url consistency.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.WorkloadFederationServiceCreateProviderRequestWellKnownProviderWellKnownWorkloadProviderUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.WorkloadFederationServiceCreateProviderRequestWellKnownProvider("custom_value")
```


## Values

| Name                                                                                                     | Value                                                                                                    |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `WorkloadFederationServiceCreateProviderRequestWellKnownProviderWellKnownWorkloadProviderUnspecified`    | WELL_KNOWN_WORKLOAD_PROVIDER_UNSPECIFIED                                                                 |
| `WorkloadFederationServiceCreateProviderRequestWellKnownProviderWellKnownWorkloadProviderCustom`         | WELL_KNOWN_WORKLOAD_PROVIDER_CUSTOM                                                                      |
| `WorkloadFederationServiceCreateProviderRequestWellKnownProviderWellKnownWorkloadProviderGithubActions`  | WELL_KNOWN_WORKLOAD_PROVIDER_GITHUB_ACTIONS                                                              |
| `WorkloadFederationServiceCreateProviderRequestWellKnownProviderWellKnownWorkloadProviderGitlabCi`       | WELL_KNOWN_WORKLOAD_PROVIDER_GITLAB_CI                                                                   |
| `WorkloadFederationServiceCreateProviderRequestWellKnownProviderWellKnownWorkloadProviderHcpTerraform`   | WELL_KNOWN_WORKLOAD_PROVIDER_HCP_TERRAFORM                                                               |
| `WorkloadFederationServiceCreateProviderRequestWellKnownProviderWellKnownWorkloadProviderAwsIamOutbound` | WELL_KNOWN_WORKLOAD_PROVIDER_AWS_IAM_OUTBOUND                                                            |