# MCPAccessProfileServiceSearchRequestableConnectorsResponse

MCPAccessProfileServiceSearchRequestableConnectorsResponse returns one page
 of card-ready requestable-connector entries.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `List`                                                                                      | [][shared.RequestableConnectorView](../../../pkg/models/shared/requestableconnectorview.md) | :heavy_minus_sign:                                                                          | The page of connector cards.                                                                |
| `NextPageToken`                                                                             | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Token for next page.                                                                        |