# C1APIIamV1RolesGetResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ContentType`                                                              | *string*                                                                   | :heavy_check_mark:                                                         | HTTP response content type for this operation                              |
| `GetRolesResponse`                                                         | [*shared.GetRolesResponse](../../../pkg/models/shared/getrolesresponse.md) | :heavy_minus_sign:                                                         | The GetRolesResponse message contains the retrieved role.                  |
| `StatusCode`                                                               | *int*                                                                      | :heavy_check_mark:                                                         | HTTP response status code for this operation                               |
| `RawResponse`                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                     | :heavy_minus_sign:                                                         | Raw HTTP response; suitable for custom response parsing                    |