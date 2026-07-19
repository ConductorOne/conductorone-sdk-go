# PersonalDeviceInput

PersonalDevice is one physical device with its app clients. The device
 identity is a stable thumbprint of the device's root signing key; the root key
 never authenticates an app — each client uses its own key.


## Fields

| Field                                                                                                                                                     | Type                                                                                                                                                      | Required                                                                                                                                                  | Description                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `DisplayName`                                                                                                                                             | `*string`                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                        | The human-friendly device label, defaulted from the first app's name at<br/> registration. Devices are listed sorted by this name. Mutable via<br/> UpdateDevice. |