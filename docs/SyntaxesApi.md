# \SyntaxesApi

All URIs are relative to *http://filterlists.com/api/directory*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyntaxesGet**](SyntaxesApi.md#SyntaxesGet) | **Get** /syntaxes | Gets the syntaxes of the FilterLists.



## SyntaxesGet

> []FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm SyntaxesGet(ctx).Execute()

Gets the syntaxes of the FilterLists.

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
    resp, r, err := apiClient.SyntaxesApi.SyntaxesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntaxesApi.SyntaxesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntaxesGet`: []FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm
    fmt.Fprintf(os.Stdout, "Response from `SyntaxesApi.SyntaxesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyntaxesGetRequest struct via the builder pattern


### Return type

[**[]FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm**](FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

