/*
SolarWinds Observability REST API

Testing EntitiesAPIService

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

func Test_swov1_EntitiesAPIService(t *testing.T) {

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)

	t.Run("Test EntitiesAPIService GetEntities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.EntitiesAPI.GetEntities(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitiesAPIService GetEntityById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.EntitiesAPI.GetEntityById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}