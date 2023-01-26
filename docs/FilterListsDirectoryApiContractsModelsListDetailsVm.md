# FilterListsDirectoryApiContractsModelsListDetailsVm

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
**ViewUrls** | Pointer to [**[]FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm**](FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm.md) | The view URLs. | [optional] 
**HomeUrl** | Pointer to **NullableString** | The URL of the home page. | [optional] 
**OnionUrl** | Pointer to **NullableString** | The URL of the Tor / Onion page. | [optional] 
**PolicyUrl** | Pointer to **NullableString** | The URL of the policy/guidelines for the types of rules this FilterList includes. | [optional] 
**SubmissionUrl** | Pointer to **NullableString** | The URL of the submission/contact form for adding rules to this FilterList. | [optional] 
**IssuesUrl** | Pointer to **NullableString** | The URL of the GitHub Issues page. | [optional] 
**ForumUrl** | Pointer to **NullableString** | The URL of the forum page. | [optional] 
**ChatUrl** | Pointer to **NullableString** | The URL of the chat room. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The email address at which the project can be contacted. | [optional] 
**DonateUrl** | Pointer to **NullableString** | The URL at which donations to the project can be made. | [optional] 
**MaintainerIds** | Pointer to **[]int64** | The identifiers of the Maintainers of this FilterList. | [optional] 
**UpstreamFilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists from which this FilterList was forked. | [optional] 
**ForkFilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists that have been forked from this FilterList. | [optional] 
**IncludedInFilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists that include this FilterList. | [optional] 
**IncludesFilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists that this FilterList includes. | [optional] 
**DependencyFilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists that this FilterList depends upon. | [optional] 
**DependentFilterListIds** | Pointer to **[]int64** | The identifiers of the FilterLists dependent upon this FilterList. | [optional] 

## Methods

### NewFilterListsDirectoryApiContractsModelsListDetailsVm

`func NewFilterListsDirectoryApiContractsModelsListDetailsVm() *FilterListsDirectoryApiContractsModelsListDetailsVm`

NewFilterListsDirectoryApiContractsModelsListDetailsVm instantiates a new FilterListsDirectoryApiContractsModelsListDetailsVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterListsDirectoryApiContractsModelsListDetailsVmWithDefaults

`func NewFilterListsDirectoryApiContractsModelsListDetailsVmWithDefaults() *FilterListsDirectoryApiContractsModelsListDetailsVm`

NewFilterListsDirectoryApiContractsModelsListDetailsVmWithDefaults instantiates a new FilterListsDirectoryApiContractsModelsListDetailsVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLicenseId

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLicenseId() int64`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLicenseIdOk() (*int64, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetLicenseId(v int64)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetSyntaxIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSyntaxIds() []int64`

GetSyntaxIds returns the SyntaxIds field if non-nil, zero value otherwise.

### GetSyntaxIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSyntaxIdsOk() (*[]int64, bool)`

GetSyntaxIdsOk returns a tuple with the SyntaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetSyntaxIds(v []int64)`

SetSyntaxIds sets SyntaxIds field to given value.

### HasSyntaxIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasSyntaxIds() bool`

HasSyntaxIds returns a boolean if a field has been set.

### GetLanguageIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLanguageIds() []int64`

GetLanguageIds returns the LanguageIds field if non-nil, zero value otherwise.

### GetLanguageIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLanguageIdsOk() (*[]int64, bool)`

GetLanguageIdsOk returns a tuple with the LanguageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetLanguageIds(v []int64)`

SetLanguageIds sets LanguageIds field to given value.

### HasLanguageIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasLanguageIds() bool`

HasLanguageIds returns a boolean if a field has been set.

### GetTagIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetViewUrls

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetViewUrls() []FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm`

GetViewUrls returns the ViewUrls field if non-nil, zero value otherwise.

### GetViewUrlsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetViewUrlsOk() (*[]FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm, bool)`

GetViewUrlsOk returns a tuple with the ViewUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrls

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetViewUrls(v []FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm)`

SetViewUrls sets ViewUrls field to given value.

### HasViewUrls

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasViewUrls() bool`

HasViewUrls returns a boolean if a field has been set.

### GetHomeUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### SetHomeUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetHomeUrlNil(b bool)`

 SetHomeUrlNil sets the value for HomeUrl to be an explicit nil

### UnsetHomeUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetHomeUrl()`

UnsetHomeUrl ensures that no value is present for HomeUrl, not even an explicit nil
### GetOnionUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetOnionUrl() string`

GetOnionUrl returns the OnionUrl field if non-nil, zero value otherwise.

### GetOnionUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetOnionUrlOk() (*string, bool)`

GetOnionUrlOk returns a tuple with the OnionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnionUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetOnionUrl(v string)`

SetOnionUrl sets OnionUrl field to given value.

### HasOnionUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasOnionUrl() bool`

HasOnionUrl returns a boolean if a field has been set.

### SetOnionUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetOnionUrlNil(b bool)`

 SetOnionUrlNil sets the value for OnionUrl to be an explicit nil

### UnsetOnionUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetOnionUrl()`

UnsetOnionUrl ensures that no value is present for OnionUrl, not even an explicit nil
### GetPolicyUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetPolicyUrl() string`

GetPolicyUrl returns the PolicyUrl field if non-nil, zero value otherwise.

### GetPolicyUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetPolicyUrlOk() (*string, bool)`

GetPolicyUrlOk returns a tuple with the PolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetPolicyUrl(v string)`

SetPolicyUrl sets PolicyUrl field to given value.

### HasPolicyUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasPolicyUrl() bool`

HasPolicyUrl returns a boolean if a field has been set.

### SetPolicyUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetPolicyUrlNil(b bool)`

 SetPolicyUrlNil sets the value for PolicyUrl to be an explicit nil

### UnsetPolicyUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetPolicyUrl()`

UnsetPolicyUrl ensures that no value is present for PolicyUrl, not even an explicit nil
### GetSubmissionUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSubmissionUrl() string`

GetSubmissionUrl returns the SubmissionUrl field if non-nil, zero value otherwise.

### GetSubmissionUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSubmissionUrlOk() (*string, bool)`

GetSubmissionUrlOk returns a tuple with the SubmissionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetSubmissionUrl(v string)`

SetSubmissionUrl sets SubmissionUrl field to given value.

### HasSubmissionUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasSubmissionUrl() bool`

HasSubmissionUrl returns a boolean if a field has been set.

### SetSubmissionUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetSubmissionUrlNil(b bool)`

 SetSubmissionUrlNil sets the value for SubmissionUrl to be an explicit nil

### UnsetSubmissionUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetSubmissionUrl()`

UnsetSubmissionUrl ensures that no value is present for SubmissionUrl, not even an explicit nil
### GetIssuesUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.

### HasIssuesUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasIssuesUrl() bool`

HasIssuesUrl returns a boolean if a field has been set.

### SetIssuesUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIssuesUrlNil(b bool)`

 SetIssuesUrlNil sets the value for IssuesUrl to be an explicit nil

### UnsetIssuesUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetIssuesUrl()`

UnsetIssuesUrl ensures that no value is present for IssuesUrl, not even an explicit nil
### GetForumUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForumUrl() string`

GetForumUrl returns the ForumUrl field if non-nil, zero value otherwise.

### GetForumUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForumUrlOk() (*string, bool)`

GetForumUrlOk returns a tuple with the ForumUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetForumUrl(v string)`

SetForumUrl sets ForumUrl field to given value.

### HasForumUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasForumUrl() bool`

HasForumUrl returns a boolean if a field has been set.

### SetForumUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetForumUrlNil(b bool)`

 SetForumUrlNil sets the value for ForumUrl to be an explicit nil

### UnsetForumUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetForumUrl()`

UnsetForumUrl ensures that no value is present for ForumUrl, not even an explicit nil
### GetChatUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetChatUrl() string`

GetChatUrl returns the ChatUrl field if non-nil, zero value otherwise.

### GetChatUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetChatUrlOk() (*string, bool)`

GetChatUrlOk returns a tuple with the ChatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetChatUrl(v string)`

SetChatUrl sets ChatUrl field to given value.

### HasChatUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasChatUrl() bool`

HasChatUrl returns a boolean if a field has been set.

### SetChatUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetChatUrlNil(b bool)`

 SetChatUrlNil sets the value for ChatUrl to be an explicit nil

### UnsetChatUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetChatUrl()`

UnsetChatUrl ensures that no value is present for ChatUrl, not even an explicit nil
### GetEmailAddress

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetDonateUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDonateUrl() string`

GetDonateUrl returns the DonateUrl field if non-nil, zero value otherwise.

### GetDonateUrlOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDonateUrlOk() (*string, bool)`

GetDonateUrlOk returns a tuple with the DonateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonateUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDonateUrl(v string)`

SetDonateUrl sets DonateUrl field to given value.

### HasDonateUrl

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDonateUrl() bool`

HasDonateUrl returns a boolean if a field has been set.

### SetDonateUrlNil

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDonateUrlNil(b bool)`

 SetDonateUrlNil sets the value for DonateUrl to be an explicit nil

### UnsetDonateUrl
`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetDonateUrl()`

UnsetDonateUrl ensures that no value is present for DonateUrl, not even an explicit nil
### GetMaintainerIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetMaintainerIds() []int64`

GetMaintainerIds returns the MaintainerIds field if non-nil, zero value otherwise.

### GetMaintainerIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetMaintainerIdsOk() (*[]int64, bool)`

GetMaintainerIdsOk returns a tuple with the MaintainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetMaintainerIds(v []int64)`

SetMaintainerIds sets MaintainerIds field to given value.

### HasMaintainerIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasMaintainerIds() bool`

HasMaintainerIds returns a boolean if a field has been set.

### GetUpstreamFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetUpstreamFilterListIds() []int64`

GetUpstreamFilterListIds returns the UpstreamFilterListIds field if non-nil, zero value otherwise.

### GetUpstreamFilterListIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetUpstreamFilterListIdsOk() (*[]int64, bool)`

GetUpstreamFilterListIdsOk returns a tuple with the UpstreamFilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetUpstreamFilterListIds(v []int64)`

SetUpstreamFilterListIds sets UpstreamFilterListIds field to given value.

### HasUpstreamFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasUpstreamFilterListIds() bool`

HasUpstreamFilterListIds returns a boolean if a field has been set.

### GetForkFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForkFilterListIds() []int64`

GetForkFilterListIds returns the ForkFilterListIds field if non-nil, zero value otherwise.

### GetForkFilterListIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForkFilterListIdsOk() (*[]int64, bool)`

GetForkFilterListIdsOk returns a tuple with the ForkFilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetForkFilterListIds(v []int64)`

SetForkFilterListIds sets ForkFilterListIds field to given value.

### HasForkFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasForkFilterListIds() bool`

HasForkFilterListIds returns a boolean if a field has been set.

### GetIncludedInFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludedInFilterListIds() []int64`

GetIncludedInFilterListIds returns the IncludedInFilterListIds field if non-nil, zero value otherwise.

### GetIncludedInFilterListIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludedInFilterListIdsOk() (*[]int64, bool)`

GetIncludedInFilterListIdsOk returns a tuple with the IncludedInFilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedInFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIncludedInFilterListIds(v []int64)`

SetIncludedInFilterListIds sets IncludedInFilterListIds field to given value.

### HasIncludedInFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasIncludedInFilterListIds() bool`

HasIncludedInFilterListIds returns a boolean if a field has been set.

### GetIncludesFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludesFilterListIds() []int64`

GetIncludesFilterListIds returns the IncludesFilterListIds field if non-nil, zero value otherwise.

### GetIncludesFilterListIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludesFilterListIdsOk() (*[]int64, bool)`

GetIncludesFilterListIdsOk returns a tuple with the IncludesFilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIncludesFilterListIds(v []int64)`

SetIncludesFilterListIds sets IncludesFilterListIds field to given value.

### HasIncludesFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasIncludesFilterListIds() bool`

HasIncludesFilterListIds returns a boolean if a field has been set.

### GetDependencyFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependencyFilterListIds() []int64`

GetDependencyFilterListIds returns the DependencyFilterListIds field if non-nil, zero value otherwise.

### GetDependencyFilterListIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependencyFilterListIdsOk() (*[]int64, bool)`

GetDependencyFilterListIdsOk returns a tuple with the DependencyFilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDependencyFilterListIds(v []int64)`

SetDependencyFilterListIds sets DependencyFilterListIds field to given value.

### HasDependencyFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDependencyFilterListIds() bool`

HasDependencyFilterListIds returns a boolean if a field has been set.

### GetDependentFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependentFilterListIds() []int64`

GetDependentFilterListIds returns the DependentFilterListIds field if non-nil, zero value otherwise.

### GetDependentFilterListIdsOk

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependentFilterListIdsOk() (*[]int64, bool)`

GetDependentFilterListIdsOk returns a tuple with the DependentFilterListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDependentFilterListIds(v []int64)`

SetDependentFilterListIds sets DependentFilterListIds field to given value.

### HasDependentFilterListIds

`func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDependentFilterListIds() bool`

HasDependentFilterListIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


