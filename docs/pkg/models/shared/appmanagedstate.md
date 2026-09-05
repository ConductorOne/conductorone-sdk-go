# AppManagedState

AppManagedState identifies whether a discovered application is managed.

This message contains a oneof named state. Only a single field of the following list may be set at a time:
  - unmanaged
  - managed



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Managed`                                                                                  | [*shared.AppManagedStateManaged](../../../pkg/models/shared/appmanagedstatemanaged.md)     | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Unmanaged`                                                                                | [*shared.AppManagedStateUnmanaged](../../../pkg/models/shared/appmanagedstateunmanaged.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |