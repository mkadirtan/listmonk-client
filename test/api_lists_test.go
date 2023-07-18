/*
Listmonk

Testing ListsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ListsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ListsApiService CreateList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListsApi.CreateList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsApiService DeleteListById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.ListsApi.DeleteListById(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsApiService GetListById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.ListsApi.GetListById(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsApiService GetLists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListsApi.GetLists(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsApiService UpdateListById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.ListsApi.UpdateListById(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
