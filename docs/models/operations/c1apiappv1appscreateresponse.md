# C1APIAppV1AppsCreateResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | HTTP response content type for this operation                         |
| `CreateAppResponse`                                                   | [*shared.CreateAppResponse](../../models/shared/createappresponse.md) | :heavy_minus_sign:                                                    | Returns the new app's values.                                         |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | HTTP response status code for this operation                          |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | Raw HTTP response; suitable for custom response parsing               |