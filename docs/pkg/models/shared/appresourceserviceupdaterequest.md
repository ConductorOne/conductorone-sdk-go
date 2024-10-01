# AppResourceServiceUpdateRequest

The AppResourceServiceUpdateRequest message.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `AppResource`                                                                                 | [*shared.AppResourceInput](../../../pkg/models/shared/appresourceinput.md)                    | :heavy_minus_sign:                                                                            | The app resource message is a single resource that can have entitlements.                     |
| `AppResourceExpandMask`                                                                       | [*shared.AppResourceExpandMask](../../../pkg/models/shared/appresourceexpandmask.md)          | :heavy_minus_sign:                                                                            | The app resource expand mask lets you get information about related objects from the request. |
| `UpdateMask`                                                                                  | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           |