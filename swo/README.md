# github.com/grepory/swo-sdk-go/swo

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/grepory/swo-sdk-go/swo* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/grepory/swo-sdk-go/swo&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/swo/swo). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Pagination](#pagination)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/grepory/swo-sdk-go/swo
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [ChangeEvents](docs/sdks/changeevents/README.md)

* [CreateChangeEvent](docs/sdks/changeevents/README.md#createchangeevent) - Create an event

### [Entities](docs/sdks/entities/README.md)

* [ListEntities](docs/sdks/entities/README.md#listentities) - Get a list of entities
* [GetEntityByID](docs/sdks/entities/README.md#getentitybyid) - Get an entity by ID

### [Logs](docs/sdks/logs/README.md)

* [SearchLogs](docs/sdks/logs/README.md#searchlogs) - Search logs
* [ListLogArchives](docs/sdks/logs/README.md#listlogarchives) - Retrieve a list of log archives

### [Metadata](docs/sdks/metadata/README.md)

* [ListEntityTypes](docs/sdks/metadata/README.md#listentitytypes) - List all entity types
* [ListMetricsForEntityType](docs/sdks/metadata/README.md#listmetricsforentitytype) - List metrics metadata for an entity type

### [Metrics](docs/sdks/metrics/README.md)

* [ListMetrics](docs/sdks/metrics/README.md#listmetrics) - List metrics
* [GetMetricByName](docs/sdks/metrics/README.md#getmetricbyname) - Get metric info
* [ListMetricAttributes](docs/sdks/metrics/README.md#listmetricattributes) - List metric attribute names
* [ListMetricAttributeValues](docs/sdks/metrics/README.md#listmetricattributevalues) - List metric attribute values
* [ListMetricMeasurements](docs/sdks/metrics/README.md#listmetricmeasurements) - List metric measurement values, grouped by attributes, filtered by the filter


### [Tokens](docs/sdks/tokens/README.md)

* [CreateToken](docs/sdks/tokens/README.md#createtoken) - Create ingestion token

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	"github.com/grepory/swo-sdk-go/swo"
	"github.com/grepory/swo-sdk-go/swo/models/operations"
	"log"
	"os"
)

func main() {
	s := swo.New(
		swo.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Entities.ListEntities(ctx, operations.ListEntitiesRequest{
		Type: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}

	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/grepory/swo-sdk-go/swo"
	"github.com/grepory/swo-sdk-go/swo/models/components"
	"github.com/grepory/swo-sdk-go/swo/retry"
	"log"
	"models/operations"
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
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/grepory/swo-sdk-go/swo"
	"github.com/grepory/swo-sdk-go/swo/models/components"
	"github.com/grepory/swo-sdk-go/swo/retry"
	"log"
	"os"
)

func main() {
	s := swo.New(
		swo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateChangeEvent` function may return the following errors:

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/grepory/swo-sdk-go/swo"
	"github.com/grepory/swo-sdk-go/swo/models/components"
	"github.com/grepory/swo-sdk-go/swo/models/sdkerrors"
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

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.{region}.cloud.solarwinds.com/` | `region` (default is `na-01`) |

#### Example

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
		swo.WithServerIndex(0),
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

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithRegion string`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		swo.WithServerURL("https://api.{region}.cloud.solarwinds.com/"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                 | Type                 | Scheme               | Environment Variable |
| -------------------- | -------------------- | -------------------- | -------------------- |
| `BearerAuth`         | http                 | HTTP Bearer          | `SWO_BEARER_AUTH`    |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/grepory/swo-sdk-go/swo&utm_campaign=go)