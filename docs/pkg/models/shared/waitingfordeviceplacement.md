# WaitingForDevicePlacement

Describes a provision step that is paused until the recipient joins the vault's MLS group.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RecipientUserID`                                              | `*string`                                                      | :heavy_minus_sign:                                             | The ID of the user being placed.                               |
| `VaultBoundaryID`                                              | `*string`                                                      | :heavy_minus_sign:                                             | The ID of the vault boundary the recipient is being placed in. |