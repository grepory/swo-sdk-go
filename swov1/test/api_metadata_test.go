/*
SolarWinds Observability REST API

Testing MetadataAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package swov1

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	swoclient "github.com/solarwinds/swo-sdk-go/swov1"
)

func Test_swov1_MetadataAPIService(t *testing.T) {

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)

	t.Run("Test MetadataAPIService ListEntityTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetadataAPI.ListEntityTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataAPIService ListMetricsForEntityType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.MetadataAPI.ListMetricsForEntityType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}