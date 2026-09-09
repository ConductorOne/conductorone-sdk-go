# SearchDenialsResponse

SearchDenialsResponse contains one page of denial episodes.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Episodes`                                                                | [][shared.DenialEpisode](../../../pkg/models/shared/denialepisode.md)     | :heavy_minus_sign:                                                        | Episodes ordered by most recent refusal, then unique ID, both descending. |
| `NextPageToken`                                                           | `*string`                                                                 | :heavy_minus_sign:                                                        | Token for the next page. Empty when no more episodes remain.              |