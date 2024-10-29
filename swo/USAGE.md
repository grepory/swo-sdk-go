<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/grepory/swo-sdk-go/swo"
	"github.com/grepory/swo-sdk-go/swo/models/components"
	"log"
	"os"
)

func main() {
	s := swo.New(
		swo.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
	)

	ctx := context.Background()
	res, err := s.ChangeEvents.CreateChangeEvent(ctx, components.ChangeEvent{
		Name:  "<value>",
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
<!-- End SDK Example Usage [usage] -->