/*
SolarWinds Observability REST API

[Rest API Documentation](https://documentation.solarwinds.com/en/success_center/observability/content/api/api-swagger.htm)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package swov1

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AnalysisSchemaAllOfRedirects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisSchemaAllOfRedirects{}

// AnalysisSchemaAllOfRedirects struct for AnalysisSchemaAllOfRedirects
type AnalysisSchemaAllOfRedirects struct {
	Url string `json:"url"`
	DurationInMs int32 `json:"durationInMs"`
	IpLookupDetails *AnalysisSchemaAllOfIpLookupDetails `json:"ipLookupDetails,omitempty"`
	ConnectionDetails *AnalysisSchemaAllOfConnectionDetails `json:"connectionDetails,omitempty"`
	RequestDetails *HttpMessage `json:"requestDetails,omitempty"`
	ResponseDetails *HttpMessage `json:"responseDetails,omitempty"`
}

type _AnalysisSchemaAllOfRedirects AnalysisSchemaAllOfRedirects

// NewAnalysisSchemaAllOfRedirects instantiates a new AnalysisSchemaAllOfRedirects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisSchemaAllOfRedirects(url string, durationInMs int32) *AnalysisSchemaAllOfRedirects {
	this := AnalysisSchemaAllOfRedirects{}
	this.Url = url
	this.DurationInMs = durationInMs
	return &this
}

// NewAnalysisSchemaAllOfRedirectsWithDefaults instantiates a new AnalysisSchemaAllOfRedirects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisSchemaAllOfRedirectsWithDefaults() *AnalysisSchemaAllOfRedirects {
	this := AnalysisSchemaAllOfRedirects{}
	return &this
}

// GetUrl returns the Url field value
func (o *AnalysisSchemaAllOfRedirects) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfRedirects) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AnalysisSchemaAllOfRedirects) SetUrl(v string) {
	o.Url = v
}

// GetDurationInMs returns the DurationInMs field value
func (o *AnalysisSchemaAllOfRedirects) GetDurationInMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationInMs
}

// GetDurationInMsOk returns a tuple with the DurationInMs field value
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfRedirects) GetDurationInMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationInMs, true
}

// SetDurationInMs sets field value
func (o *AnalysisSchemaAllOfRedirects) SetDurationInMs(v int32) {
	o.DurationInMs = v
}

// GetIpLookupDetails returns the IpLookupDetails field value if set, zero value otherwise.
func (o *AnalysisSchemaAllOfRedirects) GetIpLookupDetails() AnalysisSchemaAllOfIpLookupDetails {
	if o == nil || IsNil(o.IpLookupDetails) {
		var ret AnalysisSchemaAllOfIpLookupDetails
		return ret
	}
	return *o.IpLookupDetails
}

// GetIpLookupDetailsOk returns a tuple with the IpLookupDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfRedirects) GetIpLookupDetailsOk() (*AnalysisSchemaAllOfIpLookupDetails, bool) {
	if o == nil || IsNil(o.IpLookupDetails) {
		return nil, false
	}
	return o.IpLookupDetails, true
}

// HasIpLookupDetails returns a boolean if a field has been set.
func (o *AnalysisSchemaAllOfRedirects) HasIpLookupDetails() bool {
	if o != nil && !IsNil(o.IpLookupDetails) {
		return true
	}

	return false
}

// SetIpLookupDetails gets a reference to the given AnalysisSchemaAllOfIpLookupDetails and assigns it to the IpLookupDetails field.
func (o *AnalysisSchemaAllOfRedirects) SetIpLookupDetails(v AnalysisSchemaAllOfIpLookupDetails) {
	o.IpLookupDetails = &v
}

// GetConnectionDetails returns the ConnectionDetails field value if set, zero value otherwise.
func (o *AnalysisSchemaAllOfRedirects) GetConnectionDetails() AnalysisSchemaAllOfConnectionDetails {
	if o == nil || IsNil(o.ConnectionDetails) {
		var ret AnalysisSchemaAllOfConnectionDetails
		return ret
	}
	return *o.ConnectionDetails
}

// GetConnectionDetailsOk returns a tuple with the ConnectionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfRedirects) GetConnectionDetailsOk() (*AnalysisSchemaAllOfConnectionDetails, bool) {
	if o == nil || IsNil(o.ConnectionDetails) {
		return nil, false
	}
	return o.ConnectionDetails, true
}

// HasConnectionDetails returns a boolean if a field has been set.
func (o *AnalysisSchemaAllOfRedirects) HasConnectionDetails() bool {
	if o != nil && !IsNil(o.ConnectionDetails) {
		return true
	}

	return false
}

// SetConnectionDetails gets a reference to the given AnalysisSchemaAllOfConnectionDetails and assigns it to the ConnectionDetails field.
func (o *AnalysisSchemaAllOfRedirects) SetConnectionDetails(v AnalysisSchemaAllOfConnectionDetails) {
	o.ConnectionDetails = &v
}

// GetRequestDetails returns the RequestDetails field value if set, zero value otherwise.
func (o *AnalysisSchemaAllOfRedirects) GetRequestDetails() HttpMessage {
	if o == nil || IsNil(o.RequestDetails) {
		var ret HttpMessage
		return ret
	}
	return *o.RequestDetails
}

// GetRequestDetailsOk returns a tuple with the RequestDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfRedirects) GetRequestDetailsOk() (*HttpMessage, bool) {
	if o == nil || IsNil(o.RequestDetails) {
		return nil, false
	}
	return o.RequestDetails, true
}

// HasRequestDetails returns a boolean if a field has been set.
func (o *AnalysisSchemaAllOfRedirects) HasRequestDetails() bool {
	if o != nil && !IsNil(o.RequestDetails) {
		return true
	}

	return false
}

// SetRequestDetails gets a reference to the given HttpMessage and assigns it to the RequestDetails field.
func (o *AnalysisSchemaAllOfRedirects) SetRequestDetails(v HttpMessage) {
	o.RequestDetails = &v
}

// GetResponseDetails returns the ResponseDetails field value if set, zero value otherwise.
func (o *AnalysisSchemaAllOfRedirects) GetResponseDetails() HttpMessage {
	if o == nil || IsNil(o.ResponseDetails) {
		var ret HttpMessage
		return ret
	}
	return *o.ResponseDetails
}

// GetResponseDetailsOk returns a tuple with the ResponseDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfRedirects) GetResponseDetailsOk() (*HttpMessage, bool) {
	if o == nil || IsNil(o.ResponseDetails) {
		return nil, false
	}
	return o.ResponseDetails, true
}

// HasResponseDetails returns a boolean if a field has been set.
func (o *AnalysisSchemaAllOfRedirects) HasResponseDetails() bool {
	if o != nil && !IsNil(o.ResponseDetails) {
		return true
	}

	return false
}

// SetResponseDetails gets a reference to the given HttpMessage and assigns it to the ResponseDetails field.
func (o *AnalysisSchemaAllOfRedirects) SetResponseDetails(v HttpMessage) {
	o.ResponseDetails = &v
}

func (o AnalysisSchemaAllOfRedirects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisSchemaAllOfRedirects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["durationInMs"] = o.DurationInMs
	if !IsNil(o.IpLookupDetails) {
		toSerialize["ipLookupDetails"] = o.IpLookupDetails
	}
	if !IsNil(o.ConnectionDetails) {
		toSerialize["connectionDetails"] = o.ConnectionDetails
	}
	if !IsNil(o.RequestDetails) {
		toSerialize["requestDetails"] = o.RequestDetails
	}
	if !IsNil(o.ResponseDetails) {
		toSerialize["responseDetails"] = o.ResponseDetails
	}
	return toSerialize, nil
}

func (o *AnalysisSchemaAllOfRedirects) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"durationInMs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAnalysisSchemaAllOfRedirects := _AnalysisSchemaAllOfRedirects{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalysisSchemaAllOfRedirects)

	if err != nil {
		return err
	}

	*o = AnalysisSchemaAllOfRedirects(varAnalysisSchemaAllOfRedirects)

	return err
}

type NullableAnalysisSchemaAllOfRedirects struct {
	value *AnalysisSchemaAllOfRedirects
	isSet bool
}

func (v NullableAnalysisSchemaAllOfRedirects) Get() *AnalysisSchemaAllOfRedirects {
	return v.value
}

func (v *NullableAnalysisSchemaAllOfRedirects) Set(val *AnalysisSchemaAllOfRedirects) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisSchemaAllOfRedirects) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisSchemaAllOfRedirects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisSchemaAllOfRedirects(val *AnalysisSchemaAllOfRedirects) *NullableAnalysisSchemaAllOfRedirects {
	return &NullableAnalysisSchemaAllOfRedirects{value: val, isSet: true}
}

func (v NullableAnalysisSchemaAllOfRedirects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisSchemaAllOfRedirects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


