# \ListsApi

All URIs are relative to *http://filterlists.com/api/directory*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListsGet**](ListsApi.md#ListsGet) | **Get** /lists | Gets the FilterLists.
[**ListsIdGet**](ListsApi.md#ListsIdGet) | **Get** /lists/{id} | Gets the details of the FilterList.



## ListsGet

> []FilterListsDirectoryApiContractsModelsListVm ListsGet(ctx).Execute()

Gets the FilterLists.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsApi.ListsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsGet`: []FilterListsDirectoryApiContractsModelsListVm
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.ListsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListsGetRequest struct via the builder pattern


### Return type

[**[]FilterListsDirectoryApiContractsModelsListVm**](FilterListsDirectoryApiContractsModelsListVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsIdGet

> FilterListsDirectoryApiContractsModelsListDetailsVm ListsIdGet(ctx, id).Execute()

Gets the details of the FilterList.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(789) // int64 | The identifier of the FilterList.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsApi.ListsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.ListsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsIdGet`: FilterListsDirectoryApiContractsModelsListDetailsVm
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.ListsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the FilterList. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FilterListsDirectoryApiContractsModelsListDetailsVm**](FilterListsDirectoryApiContractsModelsListDetailsVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

