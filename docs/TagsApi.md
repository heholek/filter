# \TagsApi

All URIs are relative to *http://filterlists.com/api/directory*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsGet**](TagsApi.md#TagsGet) | **Get** /tags | Gets the tags of the FilterLists.



## TagsGet

> []FilterListsDirectoryApplicationQueriesGetTagsTagVm TagsGet(ctx).Execute()

Gets the tags of the FilterLists.

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
    resp, r, err := apiClient.TagsApi.TagsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TagsGet`: []FilterListsDirectoryApplicationQueriesGetTagsTagVm
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.TagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTagsGetRequest struct via the builder pattern


### Return type

[**[]FilterListsDirectoryApplicationQueriesGetTagsTagVm**](FilterListsDirectoryApplicationQueriesGetTagsTagVm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

