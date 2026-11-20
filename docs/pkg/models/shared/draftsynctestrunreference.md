# DraftSyncTestRunReference

DraftSyncTestRunReference identifies retained draft-test statistics for this connector's installation.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `CatalogID`                                               | `*string`                                                 | :heavy_minus_sign:                                        | Tenant-authored catalog entry that owns the installation. |
| `RevisionID`                                              | `*string`                                                 | :heavy_minus_sign:                                        | Revision used by the draft test.                          |
| `TestRunID`                                               | `*string`                                                 | :heavy_minus_sign:                                        | Draft test run whose execution statistics are requested.  |