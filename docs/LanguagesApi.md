# \LanguagesApi

All URIs are relative to *http://filterlists.com/api/directory*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LanguagesGet**](LanguagesApi.md#LanguagesGet) | **Get** /languages | Gets the languages targeted by the FilterLists.



## LanguagesGet

> []FilterListsDirectoryApplicationQueriesGetLanguagesLanguageVm LanguagesGet(ctx).Execute()

Gets the languages targeted by the FilterLists.

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
    resp, r, err := apiClient.LanguagesApi.LanguagesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguagesApi.LanguagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LanguagesGet`: []FilterListsDirectoryApplicationQueriesGetLanguagesLanguageVm
    fmt.Fprintf(os.Stdout, "Response from `LanguagesApi.LanguagesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLanguagesGetRequest struct via the builder pattern


### Return type

[**[]FilterListsDirectoryApplicationQueriesGetLanguagesLanguageVm**](FilterListsDirectoryApplicationQueriesGetLanguagesLanguageVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

