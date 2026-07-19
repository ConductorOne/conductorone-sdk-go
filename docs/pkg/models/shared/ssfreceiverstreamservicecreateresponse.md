# SSFReceiverStreamServiceCreateResponse

SSFReceiverStreamServiceCreateResponse returns the created stream and the push auth token in plaintext.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `PushAuthTokenPlaintext`                                                     | `*string`                                                                    | :heavy_minus_sign:                                                           | Push auth token returned in plaintext ONLY on create.                        |
| `SsfReceiverStream`                                                          | [*shared.SSFReceiverStream](../../../pkg/models/shared/ssfreceiverstream.md) | :heavy_minus_sign:                                                           | N/A                                                                          |