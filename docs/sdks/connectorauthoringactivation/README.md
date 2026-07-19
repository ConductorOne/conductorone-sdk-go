# ConnectorAuthoringActivation

## Overview

### Available Operations

* [ActivateRevision](#activaterevision) - Activate Revision
* [RollbackRevision](#rollbackrevision) - Rollback Revision

## ActivateRevision

ActivateRevision redeems a one-time approval token to activate a built
 connector revision onto its instance connector. It is OWNER-only.
 Double-activate protection is the single-use approval token itself: redeeming
 it is a compare-and-swap that rejects a second redemption of the same token.
 idempotency_key is optional and reserved for a future replay-result cache.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.connector_authoring.v1.ConnectorAuthoringActivationService.ActivateRevision" method="post" path="/api/v1/connector-authoring/activations" -->
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

    res, err := s.ConnectorAuthoringActivation.ActivateRevision(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorAuthoringServiceActivateRevisionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [shared.ConnectorAuthoringServiceActivateRevisionRequest](../../pkg/models/shared/connectorauthoringserviceactivaterevisionrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIConnectorAuthoringV1ConnectorAuthoringActivationServiceActivateRevisionResponse](../../pkg/models/operations/c1apiconnectorauthoringv1connectorauthoringactivationserviceactivaterevisionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RollbackRevision

RollbackRevision redeems a one-time approval token (bound to the rollback
 target's integrity root) to re-point the published + instance pointers at a
 previously activated, still-servable revision under a strictly greater
 activation epoch. It is OWNER-only. The rolled-back-FROM revision's serve
 state is untouched — the pointer move alone stops it serving; permanently
 ending a revision's serve eligibility is a platform kill-switch operation,
 not a tenant API verb.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.connector_authoring.v1.ConnectorAuthoringActivationService.RollbackRevision" method="post" path="/api/v1/connector-authoring/rollbacks" -->
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

    res, err := s.ConnectorAuthoringActivation.RollbackRevision(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorAuthoringServiceRollbackRevisionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [shared.ConnectorAuthoringServiceRollbackRevisionRequest](../../pkg/models/shared/connectorauthoringservicerollbackrevisionrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIConnectorAuthoringV1ConnectorAuthoringActivationServiceRollbackRevisionResponse](../../pkg/models/operations/c1apiconnectorauthoringv1connectorauthoringactivationservicerollbackrevisionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |