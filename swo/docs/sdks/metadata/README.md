# Metadata
(*Metadata*)

## Overview

### Available Operations

* [ListEntityTypes](#listentitytypes) - List all entity types
* [ListMetricsForEntityType](#listmetricsforentitytype) - List metrics metadata for an entity type

## ListEntityTypes

List all available entity types.

### Example Usage

```go
package main

import(
	"os"
	"github.com/grepory/swo-sdk-go/swo"
	"context"
	"log"
)

func main() {
    s := swo.New(
        swo.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Metadata.ListEntityTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListEntityTypesResponse](../../models/operations/listentitytypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMetricsForEntityType

List metrics metadata for an entity type between a start and end time. If start time and end time unspecified, default will be applied.

### Example Usage

```go
package main

import(
	"os"
	"github.com/grepory/swo-sdk-go/swo"
	"context"
	"github.com/grepory/swo-sdk-go/swo/models/operations"
	"log"
)

func main() {
    s := swo.New(
        swo.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Metadata.ListMetricsForEntityType(ctx, operations.ListMetricsForEntityTypeRequest{
        Type: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListMetricsForEntityTypeRequest](../../models/operations/listmetricsforentitytyperequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListMetricsForEntityTypeResponse](../../models/operations/listmetricsforentitytyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |