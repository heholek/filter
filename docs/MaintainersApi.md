# \MaintainersApi

All URIs are relative to *http://filterlists.com/api/directory*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaintainersGet**](MaintainersApi.md#MaintainersGet) | **Get** /maintainers | Gets the maintainers of the FilterLists.



## MaintainersGet

> []FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm MaintainersGet(ctx).Execute()

Gets the maintainers of the FilterLists.

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
    resp, r, err := apiClient.MaintainersApi.MaintainersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainersApi.MaintainersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaintainersGet`: []FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm
    fmt.Fprintf(os.Stdout, "Response from `MaintainersApi.MaintainersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMaintainersGetRequest struct via the builder pattern


### Return type

[**[]FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm**](FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

