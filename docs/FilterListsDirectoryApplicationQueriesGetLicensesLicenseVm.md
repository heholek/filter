# FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The unique name. | [optional] [readonly] 
**Url** | Pointer to **NullableString** | The URL of the home page. | [optional] [readonly] 
**PermitsModification** | Pointer to **bool** | If the License permits modification. | [optional] [readonly] 
**PermitsDistribution** | Pointer to **bool** | If the License permits distribution. | [optional] [readonly] 
**PermitsCommercialUse** | Pointer to **bool** | If the License permits commercial use. | [optional] [readonly] 
**FilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists released under this License. | [optional] [readonly] 

## Methods

### NewFilterListsDirectoryApplicationQueriesGetLicensesLicenseVm

`func NewFilterListsDirectoryApplicationQueriesGetLicensesLicenseVm() *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm`

NewFilterListsDirectoryApplicationQueriesGetLicensesLicenseVm instantiates a new FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterListsDirectoryApplicationQueriesGetLicensesLicenseVmWithDefaults

`func NewFilterListsDirectoryApplicationQueriesGetLicensesLicenseVmWithDefaults() *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm`

NewFilterListsDirectoryApplicationQueriesGetLicensesLicenseVmWithDefaults instantiates a new FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetPermitsModification

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetPermitsModification() bool`

GetPermitsModification returns the PermitsModification field if non-nil, zero value otherwise.

### GetPermitsModificationOk

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetPermitsModificationOk() (*bool, bool)`

GetPermitsModificationOk returns a tuple with the PermitsModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitsModification

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetPermitsModification(v bool)`

SetPermitsModification sets PermitsModification field to given value.

### HasPermitsModification

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) HasPermitsModification() bool`

HasPermitsModification returns a boolean if a field has been set.

### GetPermitsDistribution

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetPermitsDistribution() bool`

GetPermitsDistribution returns the PermitsDistribution field if non-nil, zero value otherwise.

### GetPermitsDistributionOk

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetPermitsDistributionOk() (*bool, bool)`

GetPermitsDistributionOk returns a tuple with the PermitsDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitsDistribution

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetPermitsDistribution(v bool)`

SetPermitsDistribution sets PermitsDistribution field to given value.

### HasPermitsDistribution

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) HasPermitsDistribution() bool`

HasPermitsDistribution returns a boolean if a field has been set.

### GetPermitsCommercialUse

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetPermitsCommercialUse() bool`

GetPermitsCommercialUse returns the PermitsCommercialUse field if non-nil, zero value otherwise.

### GetPermitsCommercialUseOk

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetPermitsCommercialUseOk() (*bool, bool)`

GetPermitsCommercialUseOk returns a tuple with the PermitsCommercialUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitsCommercialUse

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetPermitsCommercialUse(v bool)`

SetPermitsCommercialUse sets PermitsCommercialUse field to given value.

### HasPermitsCommercialUse

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) HasPermitsCommercialUse() bool`

HasPermitsCommercialUse returns a boolean if a field has been set.

### GetFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetFilterListIds() []int64`

GetFilterListIds returns the FilterListIds field if non-nil, zero value otherwise.

### GetFilterListIdsOk

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) GetFilterListIdsOk() (*[]int64, bool)`

GetFilterListIdsOk returns a tuple with the FilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) SetFilterListIds(v []int64)`

SetFilterListIds sets FilterListIds field to given value.

### HasFilterListIds

`func (o *FilterListsDirectoryApplicationQueriesGetLicensesLicenseVm) HasFilterListIds() bool`

HasFilterListIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


