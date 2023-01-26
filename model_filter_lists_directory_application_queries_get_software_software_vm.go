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

// FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm struct for FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm
type FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm struct {
	// The identifier.
	Id *int64 `json:"id,omitempty"`
	// The unique name.
	Name *string `json:"name,omitempty"`
	// The description.
	Description NullableString `json:"description,omitempty"`
	// The URL of the home page.
	HomeUrl NullableString `json:"homeUrl,omitempty"`
	// The URL of the Software download.
	DownloadUrl NullableString `json:"downloadUrl,omitempty"`
	// If the Software supports the abp: URL scheme to click-to-subscribe to a FilterList.
	SupportsAbpUrlScheme *bool `json:"supportsAbpUrlScheme,omitempty"`
	// The identifiers of the Syntaxes that this Software supports.
	SyntaxIds []int64 `json:"syntaxIds,omitempty"`
}

// NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm instantiates a new FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm() *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm {
	this := FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm{}
	return &this
}

// NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVmWithDefaults instantiates a new FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVmWithDefaults() *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm {
	this := FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) UnsetDescription() {
	o.Description.Unset()
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetHomeUrl() string {
	if o == nil || isNil(o.HomeUrl.Get()) {
		var ret string
		return ret
	}
	return *o.HomeUrl.Get()
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetHomeUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.HomeUrl.Get(), o.HomeUrl.IsSet()
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasHomeUrl() bool {
	if o != nil && o.HomeUrl.IsSet() {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given NullableString and assigns it to the HomeUrl field.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetHomeUrl(v string) {
	o.HomeUrl.Set(&v)
}
// SetHomeUrlNil sets the value for HomeUrl to be an explicit nil
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetHomeUrlNil() {
	o.HomeUrl.Set(nil)
}

// UnsetHomeUrl ensures that no value is present for HomeUrl, not even an explicit nil
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) UnsetHomeUrl() {
	o.HomeUrl.Unset()
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDownloadUrl() string {
	if o == nil || isNil(o.DownloadUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl.IsSet() {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given NullableString and assigns it to the DownloadUrl field.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}
// SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetDownloadUrlNil() {
	o.DownloadUrl.Set(nil)
}

// UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) UnsetDownloadUrl() {
	o.DownloadUrl.Unset()
}

// GetSupportsAbpUrlScheme returns the SupportsAbpUrlScheme field value if set, zero value otherwise.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSupportsAbpUrlScheme() bool {
	if o == nil || isNil(o.SupportsAbpUrlScheme) {
		var ret bool
		return ret
	}
	return *o.SupportsAbpUrlScheme
}

// GetSupportsAbpUrlSchemeOk returns a tuple with the SupportsAbpUrlScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSupportsAbpUrlSchemeOk() (*bool, bool) {
	if o == nil || isNil(o.SupportsAbpUrlScheme) {
    return nil, false
	}
	return o.SupportsAbpUrlScheme, true
}

// HasSupportsAbpUrlScheme returns a boolean if a field has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasSupportsAbpUrlScheme() bool {
	if o != nil && !isNil(o.SupportsAbpUrlScheme) {
		return true
	}

	return false
}

// SetSupportsAbpUrlScheme gets a reference to the given bool and assigns it to the SupportsAbpUrlScheme field.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetSupportsAbpUrlScheme(v bool) {
	o.SupportsAbpUrlScheme = &v
}

// GetSyntaxIds returns the SyntaxIds field value if set, zero value otherwise.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSyntaxIds() []int64 {
	if o == nil || isNil(o.SyntaxIds) {
		var ret []int64
		return ret
	}
	return o.SyntaxIds
}

// GetSyntaxIdsOk returns a tuple with the SyntaxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) GetSyntaxIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.SyntaxIds) {
    return nil, false
	}
	return o.SyntaxIds, true
}

// HasSyntaxIds returns a boolean if a field has been set.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) HasSyntaxIds() bool {
	if o != nil && !isNil(o.SyntaxIds) {
		return true
	}

	return false
}

// SetSyntaxIds gets a reference to the given []int64 and assigns it to the SyntaxIds field.
func (o *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) SetSyntaxIds(v []int64) {
	o.SyntaxIds = v
}

func (o FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) MarshalJSON() ([]byte, error) {
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
	if o.HomeUrl.IsSet() {
		toSerialize["homeUrl"] = o.HomeUrl.Get()
	}
	if o.DownloadUrl.IsSet() {
		toSerialize["downloadUrl"] = o.DownloadUrl.Get()
	}
	if !isNil(o.SupportsAbpUrlScheme) {
		toSerialize["supportsAbpUrlScheme"] = o.SupportsAbpUrlScheme
	}
	if !isNil(o.SyntaxIds) {
		toSerialize["syntaxIds"] = o.SyntaxIds
	}
	return json.Marshal(toSerialize)
}

type NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm struct {
	value *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm
	isSet bool
}

func (v NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) Get() *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm {
	return v.value
}

func (v *NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) Set(val *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm(val *FilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) *NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm {
	return &NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm{value: val, isSet: true}
}

func (v NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterListsDirectoryApplicationQueriesGetSoftwareSoftwareVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


