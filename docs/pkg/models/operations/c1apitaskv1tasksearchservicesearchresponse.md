# C1APITaskV1TaskSearchServiceSearchResponse


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                | *string*                                                                                     | :heavy_check_mark:                                                                           | HTTP response content type for this operation                                                |
| `StatusCode`                                                                                 | *int*                                                                                        | :heavy_check_mark:                                                                           | HTTP response status code for this operation                                                 |
| `RawResponse`                                                                                | [*http.Response](https://pkg.go.dev/net/http#Response)                                       | :heavy_check_mark:                                                                           | Raw HTTP response; suitable for custom response parsing                                      |
| `TaskSearchResponse`                                                                         | [*shared.TaskSearchResponse](../../../pkg/models/shared/tasksearchresponse.md)               | :heavy_minus_sign:                                                                           | The TaskSearchResponse message contains a list of results and a nextPageToken if applicable. |