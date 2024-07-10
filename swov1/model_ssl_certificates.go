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

// checks if the SslCertificates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SslCertificates{}

// SslCertificates struct for SslCertificates
type SslCertificates struct {
	Type string `json:"type"`
	Status string `json:"status"`
	MinTimeValid int32 `json:"minTimeValid"`
	Certificates *SslCertificatesAllOfCertificates `json:"certificates,omitempty"`
}

type _SslCertificates SslCertificates

// NewSslCertificates instantiates a new SslCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSslCertificates(type_ string, status string, minTimeValid int32) *SslCertificates {
	this := SslCertificates{}
	this.Type = type_
	this.Status = status
	this.MinTimeValid = minTimeValid
	return &this
}

// NewSslCertificatesWithDefaults instantiates a new SslCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSslCertificatesWithDefaults() *SslCertificates {
	this := SslCertificates{}
	return &this
}

// GetType returns the Type field value
func (o *SslCertificates) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SslCertificates) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SslCertificates) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *SslCertificates) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SslCertificates) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SslCertificates) SetStatus(v string) {
	o.Status = v
}

// GetMinTimeValid returns the MinTimeValid field value
func (o *SslCertificates) GetMinTimeValid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinTimeValid
}

// GetMinTimeValidOk returns a tuple with the MinTimeValid field value
// and a boolean to check if the value has been set.
func (o *SslCertificates) GetMinTimeValidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinTimeValid, true
}

// SetMinTimeValid sets field value
func (o *SslCertificates) SetMinTimeValid(v int32) {
	o.MinTimeValid = v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *SslCertificates) GetCertificates() SslCertificatesAllOfCertificates {
	if o == nil || IsNil(o.Certificates) {
		var ret SslCertificatesAllOfCertificates
		return ret
	}
	return *o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SslCertificates) GetCertificatesOk() (*SslCertificatesAllOfCertificates, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *SslCertificates) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given SslCertificatesAllOfCertificates and assigns it to the Certificates field.
func (o *SslCertificates) SetCertificates(v SslCertificatesAllOfCertificates) {
	o.Certificates = &v
}

func (o SslCertificates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SslCertificates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	toSerialize["minTimeValid"] = o.MinTimeValid
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	return toSerialize, nil
}

func (o *SslCertificates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"status",
		"minTimeValid",
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

	varSslCertificates := _SslCertificates{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSslCertificates)

	if err != nil {
		return err
	}

	*o = SslCertificates(varSslCertificates)

	return err
}

type NullableSslCertificates struct {
	value *SslCertificates
	isSet bool
}

func (v NullableSslCertificates) Get() *SslCertificates {
	return v.value
}

func (v *NullableSslCertificates) Set(val *SslCertificates) {
	v.value = val
	v.isSet = true
}

func (v NullableSslCertificates) IsSet() bool {
	return v.isSet
}

func (v *NullableSslCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSslCertificates(val *SslCertificates) *NullableSslCertificates {
	return &NullableSslCertificates{value: val, isSet: true}
}

func (v NullableSslCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSslCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


