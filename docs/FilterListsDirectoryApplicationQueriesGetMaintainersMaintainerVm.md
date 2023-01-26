# FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The identifier. | [optional] 
**Name** | Pointer to **string** | The unique name. | [optional] 
**Url** | Pointer to **NullableString** | The URL of the home page. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The email address. | [optional] 
**TwitterHandle** | Pointer to **NullableString** | The Twitter handle. | [optional] 
**FilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists maintained by this Maintainer. | [optional] 

## Methods

### NewFilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm

`func NewFilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm() *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm`

NewFilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm instantiates a new FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVmWithDefaults

`func NewFilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVmWithDefaults() *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm`

NewFilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVmWithDefaults instantiates a new FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetEmailAddress

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetTwitterHandle

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetTwitterHandle() string`

GetTwitterHandle returns the TwitterHandle field if non-nil, zero value otherwise.

### GetTwitterHandleOk

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetTwitterHandleOk() (*string, bool)`

GetTwitterHandleOk returns a tuple with the TwitterHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterHandle

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetTwitterHandle(v string)`

SetTwitterHandle sets TwitterHandle field to given value.

### HasTwitterHandle

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) HasTwitterHandle() bool`

HasTwitterHandle returns a boolean if a field has been set.

### SetTwitterHandleNil

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetTwitterHandleNil(b bool)`

 SetTwitterHandleNil sets the value for TwitterHandle to be an explicit nil

### UnsetTwitterHandle
`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) UnsetTwitterHandle()`

UnsetTwitterHandle ensures that no value is present for TwitterHandle, not even an explicit nil
### GetFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetFilterListIds() []int64`

GetFilterListIds returns the FilterListIds field if non-nil, zero value otherwise.

### GetFilterListIdsOk

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) GetFilterListIdsOk() (*[]int64, bool)`

GetFilterListIdsOk returns a tuple with the FilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) SetFilterListIds(v []int64)`

SetFilterListIds sets FilterListIds field to given value.

### HasFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetMaintainersMaintainerVm) HasFilterListIds() bool`

HasFilterListIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


