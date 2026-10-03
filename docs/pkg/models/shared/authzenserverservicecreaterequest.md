# AuthzenServerServiceCreateRequest

AuthzenServerServiceCreateRequest creates an AuthZEN server for an app.
 The new server starts disabled and denies every PDP call until enabled.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `DisplayName`                                                         | `*string`                                                             | :heavy_minus_sign:                                                    | Display name for the server. Defaults to "AuthZEN Server" when empty. |