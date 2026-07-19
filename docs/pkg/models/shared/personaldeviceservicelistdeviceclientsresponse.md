# PersonalDeviceServiceListDeviceClientsResponse

The PersonalDeviceServiceListDeviceClientsResponse message.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Clients`                                                                           | [][shared.PersonalDeviceClient](../../../pkg/models/shared/personaldeviceclient.md) | :heavy_minus_sign:                                                                  | The app clients registered on the device.                                           |
| `NextPageToken`                                                                     | `*string`                                                                           | :heavy_minus_sign:                                                                  | A token to retrieve the next page of results, or empty if there are no more.        |