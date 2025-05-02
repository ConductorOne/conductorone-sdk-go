# UnenrollFromAllAccessProfiles

The UnenrollFromAllAccessProfiles message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `CatalogIds`                                                                         | []*string*                                                                           | :heavy_minus_sign:                                                                   | Optional list of catalog IDs to unenroll from. If empty, unenroll from all catalogs. |
| `UserIdsCel`                                                                         | **string*                                                                            | :heavy_minus_sign:                                                                   | The userIdsCel field.                                                                |
| `UserRefs`                                                                           | [][shared.UserRef](../../../pkg/models/shared/userref.md)                            | :heavy_minus_sign:                                                                   | The userRefs field.                                                                  |