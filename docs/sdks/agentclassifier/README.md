# AgentClassifier

## Overview

### Available Operations

* [GetAgentClassifierPolicy](#getagentclassifierpolicy) - Get Agent Classifier Policy
* [UpdateAgentClassifierPolicy](#updateagentclassifierpolicy) - Update Agent Classifier Policy

## GetAgentClassifierPolicy

Get returns the tenant's classifier policy with its ordered rules. A tenant
 that has never written a policy receives an empty policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.AgentClassifierService.GetAgentClassifierPolicy" method="get" path="/api/v1/settings/ai-governance/agent-classifier-policy" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AgentClassifier.GetAgentClassifierPolicy(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAgentClassifierPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.C1APIAiGovernanceV1AgentClassifierServiceGetAgentClassifierPolicyResponse](../../pkg/models/operations/c1apiaigovernancev1agentclassifierservicegetagentclassifierpolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAgentClassifierPolicy

Update replaces the policy fields named in update_mask. When `rules` is in
 the mask, the supplied list replaces the entire ordered cascade; rules
 without an ID receive one, and the policy is materialized on first write.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.ai_governance.v1.AgentClassifierService.UpdateAgentClassifierPolicy" method="post" path="/api/v1/settings/ai-governance/agent-classifier-policy" -->
```go
package main

import(
	"context"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := conductoronesdkgo.New(
        conductoronesdkgo.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AgentClassifier.UpdateAgentClassifierPolicy(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAgentClassifierPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [shared.UpdateAgentClassifierPolicyRequest](../../pkg/models/shared/updateagentclassifierpolicyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.C1APIAiGovernanceV1AgentClassifierServiceUpdateAgentClassifierPolicyResponse](../../pkg/models/operations/c1apiaigovernancev1agentclassifierserviceupdateagentclassifierpolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |