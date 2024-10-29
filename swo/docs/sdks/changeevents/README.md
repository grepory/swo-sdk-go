# ChangeEvents
(*ChangeEvents*)

## Overview

### Available Operations

* [CreateChangeEvent](#createchangeevent) - Create an event

## CreateChangeEvent

Create an event

### Example Usage

```go
package main

import(
	"os"
	"github.com/grepory/swo-sdk-go/swo"
	"context"
	"github.com/grepory/swo-sdk-go/swo/models/components"
	"log"
)

func main() {
    s := swo.New(
        swo.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.ChangeEvents.CreateChangeEvent(ctx, components.ChangeEvent{
        Name: "<value>",
        Title: "<value>",
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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.ChangeEvent](../../models/components/changeevent.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.CreateChangeEventResponse](../../models/operations/createchangeeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |