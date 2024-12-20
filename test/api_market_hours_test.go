/*
Market Data

Testing MarketHoursAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func Test_client_MarketHoursAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MarketHoursAPIService GetMarketHour", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var marketId string

		resp, httpRes, err := apiClient.MarketHoursAPI.GetMarketHour(context.Background(), marketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketHoursAPIService GetMarketHours", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MarketHoursAPI.GetMarketHours(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
