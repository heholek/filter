# FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The identifier. | [optional] 
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | The description. | [optional] 
**HomeUrl** | Pointer to **NullableString** | The URL of the home page. | [optional] 
**DownloadUrl** | Pointer to **NullableString** | The URL of the Software download. | [optional] 
**SupportsAbpUrlScheme** | Pointer to **bool** | If the Software supports the abp: URL scheme to click-to-subscribe to a FilterList. | [optional] 
**SyntaxIds** | Pointer to **[]int64** | The identifiers of the Syntaxes that this Software supports. | [optional] 

## Methods

### NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm

`func NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm() *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm`

NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm instantiates a new FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVmWithDefaults

`func NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVmWithDefaults() *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm`

NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVmWithDefaults instantiates a new FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHomeUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### SetHomeUrlNil

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetHomeUrlNil(b bool)`

 SetHomeUrlNil sets the value for HomeUrl to be an explicit nil

### UnsetHomeUrl
`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) UnsetHomeUrl()`

UnsetHomeUrl ensures that no value is present for HomeUrl, not even an explicit nil
### GetDownloadUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### SetDownloadUrlNil

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetSupportsAbpUrlScheme

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSupportsAbpUrlScheme() bool`

GetSupportsAbpUrlScheme returns the SupportsAbpUrlScheme field if non-nil, zero value otherwise.

### GetSupportsAbpUrlSchemeOk

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSupportsAbpUrlSchemeOk() (*bool, bool)`

GetSupportsAbpUrlSchemeOk returns a tuple with the SupportsAbpUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAbpUrlScheme

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetSupportsAbpUrlScheme(v bool)`

SetSupportsAbpUrlScheme sets SupportsAbpUrlScheme field to given value.

### HasSupportsAbpUrlScheme

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasSupportsAbpUrlScheme() bool`

HasSupportsAbpUrlScheme returns a boolean if a field has been set.

### GetSyntaxIds

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSyntaxIds() []int64`

GetSyntaxIds returns the SyntaxIds field if non-nil, zero value otherwise.

### GetSyntaxIdsOk

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSyntaxIdsOk() (*[]int64, bool)`

GetSyntaxIdsOk returns a tuple with the SyntaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxIds

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetSyntaxIds(v []int64)`

SetSyntaxIds sets SyntaxIds field to given value.

### HasSyntaxIds

`func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasSyntaxIds() bool`

HasSyntaxIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


