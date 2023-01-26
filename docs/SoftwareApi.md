# \SoftwareApi

All URIs are relative to *http://filterlists.com/api/directory*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoftwareGet**](SoftwareApi.md#SoftwareGet) | **Get** /software | Gets the software that subscribes to the FilterLists.



## SoftwareGet

> []FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm SoftwareGet(ctx).Execute()

Gets the software that subscribes to the FilterLists.

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
    resp, r, err := apiClient.SoftwareApi.SoftwareGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareApi.SoftwareGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareGet`: []FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm
    fmt.Fprintf(os.Stdout, "Response from `SoftwareApi.SoftwareGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareGetRequest struct via the builder pattern


### Return type

[**[]FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm**](FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

