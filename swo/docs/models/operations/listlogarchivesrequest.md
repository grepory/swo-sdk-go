# ListLogArchivesRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `StartTime`                                                        | *string*                                                           | :heavy_check_mark:                                                 | Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ |
| `EndTime`                                                          | *string*                                                           | :heavy_check_mark:                                                 | Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ |
| `PageSize`                                                         | **int*                                                             | :heavy_minus_sign:                                                 | Number of items in a response page                                 |
| `SkipToken`                                                        | **string*                                                          | :heavy_minus_sign:                                                 | Token for the requested page                                       |