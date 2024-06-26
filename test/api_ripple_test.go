/*
moon-vault-api

Testing RippleAPIService

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

func Test_moonsdk_RippleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RippleAPIService CreateRippleAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RippleAPI.CreateRippleAccount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RippleAPIService GetRippleAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.RippleAPI.GetRippleAccount(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RippleAPIService ListRippleAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RippleAPI.ListRippleAccounts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RippleAPIService SignRippleTransaction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.RippleAPI.SignRippleTransaction(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
