# AppEntitlementSearchServiceCountGrantsForUserByAppRequest

CountGrantsForUserByApp request.


## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `AppIds`                                                                                                       | []`string`                                                                                                     | :heavy_minus_sign:                                                                                             | Restrict the count to these applications. Empty counts grants across all<br/> applications the user has access to. |
| `UserID`                                                                                                       | `*string`                                                                                                      | :heavy_minus_sign:                                                                                             | The user whose grants to count.                                                                                |