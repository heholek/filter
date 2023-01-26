# FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The identifier. | [optional] 
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | The description. | [optional] 
**Url** | Pointer to **NullableString** | The URL of the home page. | [optional] 
**FilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists implementing this Syntax. | [optional] 
**SoftwareIds** | Pointer to **[]int64** | The identifiers of the Software that supports this Syntax. | [optional] 

## Methods

### NewFilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm

`func NewFilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm() *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm`

NewFilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm instantiates a new FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVmWithDefaults

`func NewFilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVmWithDefaults() *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm`

NewFilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVmWithDefaults instantiates a new FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetFilterListIds() []int64`

GetFilterListIds returns the FilterListIds field if non-nil, zero value otherwise.

### GetFilterListIdsOk

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetFilterListIdsOk() (*[]int64, bool)`

GetFilterListIdsOk returns a tuple with the FilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetFilterListIds(v []int64)`

SetFilterListIds sets FilterListIds field to given value.

### HasFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) HasFilterListIds() bool`

HasFilterListIds returns a boolean if a field has been set.

### GetSoftwareIds

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetSoftwareIds() []int64`

GetSoftwareIds returns the SoftwareIds field if non-nil, zero value otherwise.

### GetSoftwareIdsOk

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) GetSoftwareIdsOk() (*[]int64, bool)`

GetSoftwareIdsOk returns a tuple with the SoftwareIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareIds

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) SetSoftwareIds(v []int64)`

SetSoftwareIds sets SoftwareIds field to given value.

### HasSoftwareIds

`func (o *FilterListsDirectoryApplicationQueriesGetSyntaxesSyntaxVm) HasSoftwareIds() bool`

HasSoftwareIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


