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

// checks if the Validation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Validation{}

// Validation struct for Validation
type Validation struct {
	Type string `json:"type"`
	Status string `json:"status"`
}

type _Validation Validation

// NewValidation instantiates a new Validation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidation(type_ string, status string) *Validation {
	this := Validation{}
	this.Type = type_
	this.Status = status
	return &this
}

// NewValidationWithDefaults instantiates a new Validation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationWithDefaults() *Validation {
	this := Validation{}
	return &this
}

// GetType returns the Type field value
func (o *Validation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Validation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Validation) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *Validation) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Validation) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Validation) SetStatus(v string) {
	o.Status = v
}

func (o Validation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Validation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *Validation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"status",
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

	varValidation := _Validation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValidation)

	if err != nil {
		return err
	}

	*o = Validation(varValidation)

	return err
}

type NullableValidation struct {
	value *Validation
	isSet bool
}

func (v NullableValidation) Get() *Validation {
	return v.value
}

func (v *NullableValidation) Set(val *Validation) {
	v.value = val
	v.isSet = true
}

func (v NullableValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidation(val *Validation) *NullableValidation {
	return &NullableValidation{value: val, isSet: true}
}

func (v NullableValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


