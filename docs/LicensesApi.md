# \LicensesApi

All URIs are relative to *http://filterlists.com/api/directory*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LicensesGet**](LicensesApi.md#LicensesGet) | **Get** /licenses | Gets the licenses applied to the FilterLists.



## LicensesGet

> []FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm LicensesGet(ctx).Execute()

Gets the licenses applied to the FilterLists.

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
    resp, r, err := apiClient.LicensesApi.LicensesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.LicensesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensesGet`: []FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.LicensesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLicensesGetRequest struct via the builder pattern


### Return type

[**[]FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm**](FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

