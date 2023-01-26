/*
FilterLists Directory API

An ASP.NET Core API serving the core FilterList information.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FilterListsDirectoryApiContractsModelsListDetailsVm struct for FilterListsDirectoryApiContractsModelsListDetailsVm
type FilterListsDirectoryApiContractsModelsListDetailsVm struct {
	// The identifier.
	Id *int64 `json:"id,omitempty"`
	// The unique name in title case.
	Name *string `json:"name,omitempty"`
	// The brief description in English (preferably quoted from the project).
	Description NullableString `json:"description,omitempty"`
	// The identifier of the License under which this FilterList is released.
	LicenseId *int64 `json:"licenseId,omitempty"`
	// The identifiers of the Syntaxes implemented by this FilterList.
	SyntaxIds []int64 `json:"syntaxIds,omitempty"`
	// The identifiers of the Languages targeted by this FilterList.
	LanguageIds []int64 `json:"languageIds,omitempty"`
	// The identifiers of the Tags applied to this FilterList.
	TagIds []int64 `json:"tagIds,omitempty"`
	// The view URLs.
	ViewUrls []FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm `json:"viewUrls,omitempty"`
	// The URL of the home page.
	HomeUrl NullableString `json:"homeUrl,omitempty"`
	// The URL of the Tor / Onion page.
	OnionUrl NullableString `json:"onionUrl,omitempty"`
	// The URL of the policy/guidelines for the types of rules this FilterList includes.
	PolicyUrl NullableString `json:"policyUrl,omitempty"`
	// The URL of the submission/contact form for adding rules to this FilterList.
	SubmissionUrl NullableString `json:"submissionUrl,omitempty"`
	// The URL of the GitHub Issues page.
	IssuesUrl NullableString `json:"issuesUrl,omitempty"`
	// The URL of the forum page.
	ForumUrl NullableString `json:"forumUrl,omitempty"`
	// The URL of the chat room.
	ChatUrl NullableString `json:"chatUrl,omitempty"`
	// The email address at which the project can be contacted.
	EmailAddress NullableString `json:"emailAddress,omitempty"`
	// The URL at which donations to the project can be made.
	DonateUrl NullableString `json:"donateUrl,omitempty"`
	// The identifiers of the Maintainers of this FilterList.
	MaintainerIds []int64 `json:"maintainerIds,omitempty"`
	// The identifiers of the FilterLists from which this FilterList was forked.
	UpstreamFilterListIds []int64 `json:"upstreamFilterListIds,omitempty"`
	// The identifiers of the FilterLists that have been forked from this FilterList.
	ForkFilterListIds []int64 `json:"forkFilterListIds,omitempty"`
	// The identifiers of the FilterLists that include this FilterList.
	IncludedInFilterListIds []int64 `json:"includedInFilterListIds,omitempty"`
	// The identifiers of the FilterLists that this FilterList includes.
	IncludesFilterListIds []int64 `json:"includesFilterListIds,omitempty"`
	// The identifiers of the FilterLists that this FilterList depends upon.
	DependencyFilterListIds []int64 `json:"dependencyFilterListIds,omitempty"`
	// The identifiers of the FilterLists dependent upon this FilterList.
	DependentFilterListIds []int64 `json:"dependentFilterListIds,omitempty"`
}

// NewFilterListsDirectoryApiContractsModelsListDetailsVm instantiates a new FilterListsDirectoryApiContractsModelsListDetailsVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterListsDirectoryApiContractsModelsListDetailsVm() *FilterListsDirectoryApiContractsModelsListDetailsVm {
	this := FilterListsDirectoryApiContractsModelsListDetailsVm{}
	return &this
}

// NewFilterListsDirectoryApiContractsModelsListDetailsVmWithDefaults instantiates a new FilterListsDirectoryApiContractsModelsListDetailsVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterListsDirectoryApiContractsModelsListDetailsVmWithDefaults() *FilterListsDirectoryApiContractsModelsListDetailsVm {
	this := FilterListsDirectoryApiContractsModelsListDetailsVm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetDescription() {
	o.Description.Unset()
}

// GetLicenseId returns the LicenseId field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLicenseId() int64 {
	if o == nil || isNil(o.LicenseId) {
		var ret int64
		return ret
	}
	return *o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLicenseIdOk() (*int64, bool) {
	if o == nil || isNil(o.LicenseId) {
    return nil, false
	}
	return o.LicenseId, true
}

// HasLicenseId returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasLicenseId() bool {
	if o != nil && !isNil(o.LicenseId) {
		return true
	}

	return false
}

// SetLicenseId gets a reference to the given int64 and assigns it to the LicenseId field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetLicenseId(v int64) {
	o.LicenseId = &v
}

// GetSyntaxIds returns the SyntaxIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSyntaxIds() []int64 {
	if o == nil || isNil(o.SyntaxIds) {
		var ret []int64
		return ret
	}
	return o.SyntaxIds
}

// GetSyntaxIdsOk returns a tuple with the SyntaxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSyntaxIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.SyntaxIds) {
    return nil, false
	}
	return o.SyntaxIds, true
}

// HasSyntaxIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasSyntaxIds() bool {
	if o != nil && !isNil(o.SyntaxIds) {
		return true
	}

	return false
}

// SetSyntaxIds gets a reference to the given []int64 and assigns it to the SyntaxIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetSyntaxIds(v []int64) {
	o.SyntaxIds = v
}

// GetLanguageIds returns the LanguageIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLanguageIds() []int64 {
	if o == nil || isNil(o.LanguageIds) {
		var ret []int64
		return ret
	}
	return o.LanguageIds
}

// GetLanguageIdsOk returns a tuple with the LanguageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetLanguageIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.LanguageIds) {
    return nil, false
	}
	return o.LanguageIds, true
}

// HasLanguageIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasLanguageIds() bool {
	if o != nil && !isNil(o.LanguageIds) {
		return true
	}

	return false
}

// SetLanguageIds gets a reference to the given []int64 and assigns it to the LanguageIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetLanguageIds(v []int64) {
	o.LanguageIds = v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetTagIds() []int64 {
	if o == nil || isNil(o.TagIds) {
		var ret []int64
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetTagIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.TagIds) {
    return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasTagIds() bool {
	if o != nil && !isNil(o.TagIds) {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int64 and assigns it to the TagIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetTagIds(v []int64) {
	o.TagIds = v
}

// GetViewUrls returns the ViewUrls field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetViewUrls() []FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm {
	if o == nil || isNil(o.ViewUrls) {
		var ret []FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm
		return ret
	}
	return o.ViewUrls
}

// GetViewUrlsOk returns a tuple with the ViewUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetViewUrlsOk() ([]FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm, bool) {
	if o == nil || isNil(o.ViewUrls) {
    return nil, false
	}
	return o.ViewUrls, true
}

// HasViewUrls returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasViewUrls() bool {
	if o != nil && !isNil(o.ViewUrls) {
		return true
	}

	return false
}

// SetViewUrls gets a reference to the given []FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm and assigns it to the ViewUrls field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetViewUrls(v []FilterListsDirectoryApiContractsModelsListDetailsVmViewUrlVm) {
	o.ViewUrls = v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetHomeUrl() string {
	if o == nil || isNil(o.HomeUrl.Get()) {
		var ret string
		return ret
	}
	return *o.HomeUrl.Get()
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetHomeUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.HomeUrl.Get(), o.HomeUrl.IsSet()
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasHomeUrl() bool {
	if o != nil && o.HomeUrl.IsSet() {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given NullableString and assigns it to the HomeUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetHomeUrl(v string) {
	o.HomeUrl.Set(&v)
}
// SetHomeUrlNil sets the value for HomeUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetHomeUrlNil() {
	o.HomeUrl.Set(nil)
}

// UnsetHomeUrl ensures that no value is present for HomeUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetHomeUrl() {
	o.HomeUrl.Unset()
}

// GetOnionUrl returns the OnionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetOnionUrl() string {
	if o == nil || isNil(o.OnionUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OnionUrl.Get()
}

// GetOnionUrlOk returns a tuple with the OnionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetOnionUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.OnionUrl.Get(), o.OnionUrl.IsSet()
}

// HasOnionUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasOnionUrl() bool {
	if o != nil && o.OnionUrl.IsSet() {
		return true
	}

	return false
}

// SetOnionUrl gets a reference to the given NullableString and assigns it to the OnionUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetOnionUrl(v string) {
	o.OnionUrl.Set(&v)
}
// SetOnionUrlNil sets the value for OnionUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetOnionUrlNil() {
	o.OnionUrl.Set(nil)
}

// UnsetOnionUrl ensures that no value is present for OnionUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetOnionUrl() {
	o.OnionUrl.Unset()
}

// GetPolicyUrl returns the PolicyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetPolicyUrl() string {
	if o == nil || isNil(o.PolicyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyUrl.Get()
}

// GetPolicyUrlOk returns a tuple with the PolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetPolicyUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PolicyUrl.Get(), o.PolicyUrl.IsSet()
}

// HasPolicyUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasPolicyUrl() bool {
	if o != nil && o.PolicyUrl.IsSet() {
		return true
	}

	return false
}

// SetPolicyUrl gets a reference to the given NullableString and assigns it to the PolicyUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetPolicyUrl(v string) {
	o.PolicyUrl.Set(&v)
}
// SetPolicyUrlNil sets the value for PolicyUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetPolicyUrlNil() {
	o.PolicyUrl.Set(nil)
}

// UnsetPolicyUrl ensures that no value is present for PolicyUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetPolicyUrl() {
	o.PolicyUrl.Unset()
}

// GetSubmissionUrl returns the SubmissionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSubmissionUrl() string {
	if o == nil || isNil(o.SubmissionUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SubmissionUrl.Get()
}

// GetSubmissionUrlOk returns a tuple with the SubmissionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetSubmissionUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SubmissionUrl.Get(), o.SubmissionUrl.IsSet()
}

// HasSubmissionUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasSubmissionUrl() bool {
	if o != nil && o.SubmissionUrl.IsSet() {
		return true
	}

	return false
}

// SetSubmissionUrl gets a reference to the given NullableString and assigns it to the SubmissionUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetSubmissionUrl(v string) {
	o.SubmissionUrl.Set(&v)
}
// SetSubmissionUrlNil sets the value for SubmissionUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetSubmissionUrlNil() {
	o.SubmissionUrl.Set(nil)
}

// UnsetSubmissionUrl ensures that no value is present for SubmissionUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetSubmissionUrl() {
	o.SubmissionUrl.Unset()
}

// GetIssuesUrl returns the IssuesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIssuesUrl() string {
	if o == nil || isNil(o.IssuesUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IssuesUrl.Get()
}

// GetIssuesUrlOk returns a tuple with the IssuesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIssuesUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.IssuesUrl.Get(), o.IssuesUrl.IsSet()
}

// HasIssuesUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasIssuesUrl() bool {
	if o != nil && o.IssuesUrl.IsSet() {
		return true
	}

	return false
}

// SetIssuesUrl gets a reference to the given NullableString and assigns it to the IssuesUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIssuesUrl(v string) {
	o.IssuesUrl.Set(&v)
}
// SetIssuesUrlNil sets the value for IssuesUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIssuesUrlNil() {
	o.IssuesUrl.Set(nil)
}

// UnsetIssuesUrl ensures that no value is present for IssuesUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetIssuesUrl() {
	o.IssuesUrl.Unset()
}

// GetForumUrl returns the ForumUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForumUrl() string {
	if o == nil || isNil(o.ForumUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ForumUrl.Get()
}

// GetForumUrlOk returns a tuple with the ForumUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForumUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForumUrl.Get(), o.ForumUrl.IsSet()
}

// HasForumUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasForumUrl() bool {
	if o != nil && o.ForumUrl.IsSet() {
		return true
	}

	return false
}

// SetForumUrl gets a reference to the given NullableString and assigns it to the ForumUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetForumUrl(v string) {
	o.ForumUrl.Set(&v)
}
// SetForumUrlNil sets the value for ForumUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetForumUrlNil() {
	o.ForumUrl.Set(nil)
}

// UnsetForumUrl ensures that no value is present for ForumUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetForumUrl() {
	o.ForumUrl.Unset()
}

// GetChatUrl returns the ChatUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetChatUrl() string {
	if o == nil || isNil(o.ChatUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ChatUrl.Get()
}

// GetChatUrlOk returns a tuple with the ChatUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetChatUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChatUrl.Get(), o.ChatUrl.IsSet()
}

// HasChatUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasChatUrl() bool {
	if o != nil && o.ChatUrl.IsSet() {
		return true
	}

	return false
}

// SetChatUrl gets a reference to the given NullableString and assigns it to the ChatUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetChatUrl(v string) {
	o.ChatUrl.Set(&v)
}
// SetChatUrlNil sets the value for ChatUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetChatUrlNil() {
	o.ChatUrl.Set(nil)
}

// UnsetChatUrl ensures that no value is present for ChatUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetChatUrl() {
	o.ChatUrl.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetEmailAddress() string {
	if o == nil || isNil(o.EmailAddress.Get()) {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetEmailAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetDonateUrl returns the DonateUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDonateUrl() string {
	if o == nil || isNil(o.DonateUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DonateUrl.Get()
}

// GetDonateUrlOk returns a tuple with the DonateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDonateUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DonateUrl.Get(), o.DonateUrl.IsSet()
}

// HasDonateUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDonateUrl() bool {
	if o != nil && o.DonateUrl.IsSet() {
		return true
	}

	return false
}

// SetDonateUrl gets a reference to the given NullableString and assigns it to the DonateUrl field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDonateUrl(v string) {
	o.DonateUrl.Set(&v)
}
// SetDonateUrlNil sets the value for DonateUrl to be an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDonateUrlNil() {
	o.DonateUrl.Set(nil)
}

// UnsetDonateUrl ensures that no value is present for DonateUrl, not even an explicit nil
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) UnsetDonateUrl() {
	o.DonateUrl.Unset()
}

// GetMaintainerIds returns the MaintainerIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetMaintainerIds() []int64 {
	if o == nil || isNil(o.MaintainerIds) {
		var ret []int64
		return ret
	}
	return o.MaintainerIds
}

// GetMaintainerIdsOk returns a tuple with the MaintainerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetMaintainerIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.MaintainerIds) {
    return nil, false
	}
	return o.MaintainerIds, true
}

// HasMaintainerIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasMaintainerIds() bool {
	if o != nil && !isNil(o.MaintainerIds) {
		return true
	}

	return false
}

// SetMaintainerIds gets a reference to the given []int64 and assigns it to the MaintainerIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetMaintainerIds(v []int64) {
	o.MaintainerIds = v
}

// GetUpstreamFilterListIds returns the UpstreamFilterListIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetUpstreamFilterListIds() []int64 {
	if o == nil || isNil(o.UpstreamFilterListIds) {
		var ret []int64
		return ret
	}
	return o.UpstreamFilterListIds
}

// GetUpstreamFilterListIdsOk returns a tuple with the UpstreamFilterListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetUpstreamFilterListIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.UpstreamFilterListIds) {
    return nil, false
	}
	return o.UpstreamFilterListIds, true
}

// HasUpstreamFilterListIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasUpstreamFilterListIds() bool {
	if o != nil && !isNil(o.UpstreamFilterListIds) {
		return true
	}

	return false
}

// SetUpstreamFilterListIds gets a reference to the given []int64 and assigns it to the UpstreamFilterListIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetUpstreamFilterListIds(v []int64) {
	o.UpstreamFilterListIds = v
}

// GetForkFilterListIds returns the ForkFilterListIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForkFilterListIds() []int64 {
	if o == nil || isNil(o.ForkFilterListIds) {
		var ret []int64
		return ret
	}
	return o.ForkFilterListIds
}

// GetForkFilterListIdsOk returns a tuple with the ForkFilterListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetForkFilterListIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.ForkFilterListIds) {
    return nil, false
	}
	return o.ForkFilterListIds, true
}

// HasForkFilterListIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasForkFilterListIds() bool {
	if o != nil && !isNil(o.ForkFilterListIds) {
		return true
	}

	return false
}

// SetForkFilterListIds gets a reference to the given []int64 and assigns it to the ForkFilterListIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetForkFilterListIds(v []int64) {
	o.ForkFilterListIds = v
}

// GetIncludedInFilterListIds returns the IncludedInFilterListIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludedInFilterListIds() []int64 {
	if o == nil || isNil(o.IncludedInFilterListIds) {
		var ret []int64
		return ret
	}
	return o.IncludedInFilterListIds
}

// GetIncludedInFilterListIdsOk returns a tuple with the IncludedInFilterListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludedInFilterListIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.IncludedInFilterListIds) {
    return nil, false
	}
	return o.IncludedInFilterListIds, true
}

// HasIncludedInFilterListIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasIncludedInFilterListIds() bool {
	if o != nil && !isNil(o.IncludedInFilterListIds) {
		return true
	}

	return false
}

// SetIncludedInFilterListIds gets a reference to the given []int64 and assigns it to the IncludedInFilterListIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIncludedInFilterListIds(v []int64) {
	o.IncludedInFilterListIds = v
}

// GetIncludesFilterListIds returns the IncludesFilterListIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludesFilterListIds() []int64 {
	if o == nil || isNil(o.IncludesFilterListIds) {
		var ret []int64
		return ret
	}
	return o.IncludesFilterListIds
}

// GetIncludesFilterListIdsOk returns a tuple with the IncludesFilterListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetIncludesFilterListIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.IncludesFilterListIds) {
    return nil, false
	}
	return o.IncludesFilterListIds, true
}

// HasIncludesFilterListIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasIncludesFilterListIds() bool {
	if o != nil && !isNil(o.IncludesFilterListIds) {
		return true
	}

	return false
}

// SetIncludesFilterListIds gets a reference to the given []int64 and assigns it to the IncludesFilterListIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetIncludesFilterListIds(v []int64) {
	o.IncludesFilterListIds = v
}

// GetDependencyFilterListIds returns the DependencyFilterListIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependencyFilterListIds() []int64 {
	if o == nil || isNil(o.DependencyFilterListIds) {
		var ret []int64
		return ret
	}
	return o.DependencyFilterListIds
}

// GetDependencyFilterListIdsOk returns a tuple with the DependencyFilterListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependencyFilterListIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.DependencyFilterListIds) {
    return nil, false
	}
	return o.DependencyFilterListIds, true
}

// HasDependencyFilterListIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDependencyFilterListIds() bool {
	if o != nil && !isNil(o.DependencyFilterListIds) {
		return true
	}

	return false
}

// SetDependencyFilterListIds gets a reference to the given []int64 and assigns it to the DependencyFilterListIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDependencyFilterListIds(v []int64) {
	o.DependencyFilterListIds = v
}

// GetDependentFilterListIds returns the DependentFilterListIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependentFilterListIds() []int64 {
	if o == nil || isNil(o.DependentFilterListIds) {
		var ret []int64
		return ret
	}
	return o.DependentFilterListIds
}

// GetDependentFilterListIdsOk returns a tuple with the DependentFilterListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) GetDependentFilterListIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.DependentFilterListIds) {
    return nil, false
	}
	return o.DependentFilterListIds, true
}

// HasDependentFilterListIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) HasDependentFilterListIds() bool {
	if o != nil && !isNil(o.DependentFilterListIds) {
		return true
	}

	return false
}

// SetDependentFilterListIds gets a reference to the given []int64 and assigns it to the DependentFilterListIds field.
func (o *FilterListsDirectoryApiContractsModelsListDetailsVm) SetDependentFilterListIds(v []int64) {
	o.DependentFilterListIds = v
}

func (o FilterListsDirectoryApiContractsModelsListDetailsVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.LicenseId) {
		toSerialize["licenseId"] = o.LicenseId
	}
	if !isNil(o.SyntaxIds) {
		toSerialize["syntaxIds"] = o.SyntaxIds
	}
	if !isNil(o.LanguageIds) {
		toSerialize["languageIds"] = o.LanguageIds
	}
	if !isNil(o.TagIds) {
		toSerialize["tagIds"] = o.TagIds
	}
	if !isNil(o.ViewUrls) {
		toSerialize["viewUrls"] = o.ViewUrls
	}
	if o.HomeUrl.IsSet() {
		toSerialize["homeUrl"] = o.HomeUrl.Get()
	}
	if o.OnionUrl.IsSet() {
		toSerialize["onionUrl"] = o.OnionUrl.Get()
	}
	if o.PolicyUrl.IsSet() {
		toSerialize["policyUrl"] = o.PolicyUrl.Get()
	}
	if o.SubmissionUrl.IsSet() {
		toSerialize["submissionUrl"] = o.SubmissionUrl.Get()
	}
	if o.IssuesUrl.IsSet() {
		toSerialize["issuesUrl"] = o.IssuesUrl.Get()
	}
	if o.ForumUrl.IsSet() {
		toSerialize["forumUrl"] = o.ForumUrl.Get()
	}
	if o.ChatUrl.IsSet() {
		toSerialize["chatUrl"] = o.ChatUrl.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["emailAddress"] = o.EmailAddress.Get()
	}
	if o.DonateUrl.IsSet() {
		toSerialize["donateUrl"] = o.DonateUrl.Get()
	}
	if !isNil(o.MaintainerIds) {
		toSerialize["maintainerIds"] = o.MaintainerIds
	}
	if !isNil(o.UpstreamFilterListIds) {
		toSerialize["upstreamFilterListIds"] = o.UpstreamFilterListIds
	}
	if !isNil(o.ForkFilterListIds) {
		toSerialize["forkFilterListIds"] = o.ForkFilterListIds
	}
	if !isNil(o.IncludedInFilterListIds) {
		toSerialize["includedInFilterListIds"] = o.IncludedInFilterListIds
	}
	if !isNil(o.IncludesFilterListIds) {
		toSerialize["includesFilterListIds"] = o.IncludesFilterListIds
	}
	if !isNil(o.DependencyFilterListIds) {
		toSerialize["dependencyFilterListIds"] = o.DependencyFilterListIds
	}
	if !isNil(o.DependentFilterListIds) {
		toSerialize["dependentFilterListIds"] = o.DependentFilterListIds
	}
	return json.Marshal(toSerialize)
}

type NullableFilterListsDirectoryApiContractsModelsListDetailsVm struct {
	value *FilterListsDirectoryApiContractsModelsListDetailsVm
	isSet bool
}

func (v NullableFilterListsDirectoryApiContractsModelsListDetailsVm) Get() *FilterListsDirectoryApiContractsModelsListDetailsVm {
	return v.value
}

func (v *NullableFilterListsDirectoryApiContractsModelsListDetailsVm) Set(val *FilterListsDirectoryApiContractsModelsListDetailsVm) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterListsDirectoryApiContractsModelsListDetailsVm) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterListsDirectoryApiContractsModelsListDetailsVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterListsDirectoryApiContractsModelsListDetailsVm(val *FilterListsDirectoryApiContractsModelsListDetailsVm) *NullableFilterListsDirectoryApiContractsModelsListDetailsVm {
	return &NullableFilterListsDirectoryApiContractsModelsListDetailsVm{value: val, isSet: true}
}

func (v NullableFilterListsDirectoryApiContractsModelsListDetailsVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterListsDirectoryApiContractsModelsListDetailsVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


