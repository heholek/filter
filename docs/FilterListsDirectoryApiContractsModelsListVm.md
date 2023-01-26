# FilterListsDirectoryApiContractsModelsListVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The identifier. | [optional] 
**Name** | Pointer to **string** | The unique name in title case. | [optional] 
**Description** | Pointer to **NullableString** | The brief description in English (preferably quoted from the project). | [optional] 
**LicenseId** | Pointer to **int64** | The identifier of the License under which this FilterList is released. | [optional] 
**SyntaxIds** | Pointer to **[]int64** | The identifiers of the Syntaxes implemented by this FilterList. | [optional] 
**LanguageIds** | Pointer to **[]int64** | The identifiers of the Languages targeted by this FilterList. | [optional] 
**TagIds** | Pointer to **[]int64** | The identifiers of the Tags applied to this FilterList. | [optional] 
**PrimaryViewUrl** | Pointer to **NullableString** | The primary view URL. | [optional] 
**MaintainerIds** | Pointer to **[]int64** | The identifiers of the Maintainers of this FilterList. | [optional] 

## Methods

### NewFilterListsDirectoryApiContractsModelsListVm

`func NewFilterListsDirectoryApiContractsModelsListVm() *FilterListsDirectoryApiContractsModelsListVm`

NewFilterListsDirectoryApiContractsModelsListVm instantiates a new FilterListsDirectoryApiContractsModelsListVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterListsDirectoryApiContractsModelsListVmWithDefaults

`func NewFilterListsDirectoryApiContractsModelsListVmWithDefaults() *FilterListsDirectoryApiContractsModelsListVm`

NewFilterListsDirectoryApiContractsModelsListVmWithDefaults instantiates a new FilterListsDirectoryApiContractsModelsListVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FilterListsDirectoryApiContractsModelsListVm) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLicenseId

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetLicenseId() int64`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetLicenseIdOk() (*int64, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetLicenseId(v int64)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetSyntaxIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetSyntaxIds() []int64`

GetSyntaxIds returns the SyntaxIds field if non-nil, zero value otherwise.

### GetSyntaxIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetSyntaxIdsOk() (*[]int64, bool)`

GetSyntaxIdsOk returns a tuple with the SyntaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetSyntaxIds(v []int64)`

SetSyntaxIds sets SyntaxIds field to given value.

### HasSyntaxIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasSyntaxIds() bool`

HasSyntaxIds returns a boolean if a field has been set.

### GetLanguageIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetLanguageIds() []int64`

GetLanguageIds returns the LanguageIds field if non-nil, zero value otherwise.

### GetLanguageIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetLanguageIdsOk() (*[]int64, bool)`

GetLanguageIdsOk returns a tuple with the LanguageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetLanguageIds(v []int64)`

SetLanguageIds sets LanguageIds field to given value.

### HasLanguageIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasLanguageIds() bool`

HasLanguageIds returns a boolean if a field has been set.

### GetTagIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetPrimaryViewUrl

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetPrimaryViewUrl() string`

GetPrimaryViewUrl returns the PrimaryViewUrl field if non-nil, zero value otherwise.

### GetPrimaryViewUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetPrimaryViewUrlOk() (*string, bool)`

GetPrimaryViewUrlOk returns a tuple with the PrimaryViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryViewUrl

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetPrimaryViewUrl(v string)`

SetPrimaryViewUrl sets PrimaryViewUrl field to given value.

### HasPrimaryViewUrl

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasPrimaryViewUrl() bool`

HasPrimaryViewUrl returns a boolean if a field has been set.

### SetPrimaryViewUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetPrimaryViewUrlNil(b bool)`

 SetPrimaryViewUrlNil sets the value for PrimaryViewUrl to be an explicit nil

### UnsetPrimaryViewUrl
`func (o *FilterListsDirectoryApiContractsModelsListVm) UnsetPrimaryViewUrl()`

UnsetPrimaryViewUrl ensures that no value is present for PrimaryViewUrl, not even an explicit nil
### GetMaintainerIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetMaintainerIds() []int64`

GetMaintainerIds returns the MaintainerIds field if non-nil, zero value otherwise.

### GetMaintainerIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListVm) GetMaintainerIdsOk() (*[]int64, bool)`

GetMaintainerIdsOk returns a tuple with the MaintainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) SetMaintainerIds(v []int64)`

SetMaintainerIds sets MaintainerIds field to given value.

### HasMaintainerIds

`func (o *FilterListsDirectoryApiContractsModelsListVm) HasMaintainerIds() bool`

HasMaintainerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


