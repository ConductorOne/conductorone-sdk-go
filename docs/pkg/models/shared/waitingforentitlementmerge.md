# WaitingForEntitlementMerge

Describes a provision step that is paused until the target entitlement, created ahead of connector sync with a Baton match ID, is merged with its connector-synced counterpart.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AppEntitlementID`                                    | `*string`                                             | :heavy_minus_sign:                                    | The ID of the entitlement being waited on.            |
| `AppID`                                               | `*string`                                             | :heavy_minus_sign:                                    | The ID of the app the awaited entitlement belongs to. |