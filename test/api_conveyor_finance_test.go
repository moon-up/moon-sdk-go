/*
moon-vault-api

Testing ConveyorFinanceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package moonsdk

import (
	"context"
	"testing"

	openapiclient "github.com/moon-up/moon-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_moonsdk_ConveyorFinanceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConveyorFinanceAPIService Swap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ConveyorFinanceAPI.Swap(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
