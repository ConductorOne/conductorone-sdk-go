# DirectoryMergeConfig

DirectoryMergeConfig configures how AppUsers from this directory are matched to C1 Users.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `MatchCases`                                                                              | [][shared.DirectoryMergeMatchCase](../../../pkg/models/shared/directorymergematchcase.md) | :heavy_minus_sign:                                                                        | Ordered list of match cases evaluated in sequence. First match wins.                      |