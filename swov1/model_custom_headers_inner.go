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

// checks if the CustomHeadersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomHeadersInner{}

// CustomHeadersInner struct for CustomHeadersInner
type CustomHeadersInner struct {
	// Name of a request header. Must contain only characters allowed by RFC: a-z, A-Z, 0-9, - and _. 
	Name string `json:"name"`
	// Value of a request header.
	Value string `json:"value"`
}

type _CustomHeadersInner CustomHeadersInner

// NewCustomHeadersInner instantiates a new CustomHeadersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomHeadersInner(name string, value string) *CustomHeadersInner {
	this := CustomHeadersInner{}
	this.Name = name
	this.Value = value
	return &this
}

// NewCustomHeadersInnerWithDefaults instantiates a new CustomHeadersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomHeadersInnerWithDefaults() *CustomHeadersInner {
	this := CustomHeadersInner{}
	return &this
}

// GetName returns the Name field value
func (o *CustomHeadersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomHeadersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomHeadersInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *CustomHeadersInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CustomHeadersInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CustomHeadersInner) SetValue(v string) {
	o.Value = v
}

func (o CustomHeadersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomHeadersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CustomHeadersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varCustomHeadersInner := _CustomHeadersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomHeadersInner)

	if err != nil {
		return err
	}

	*o = CustomHeadersInner(varCustomHeadersInner)

	return err
}

type NullableCustomHeadersInner struct {
	value *CustomHeadersInner
	isSet bool
}

func (v NullableCustomHeadersInner) Get() *CustomHeadersInner {
	return v.value
}

func (v *NullableCustomHeadersInner) Set(val *CustomHeadersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomHeadersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomHeadersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomHeadersInner(val *CustomHeadersInner) *NullableCustomHeadersInner {
	return &NullableCustomHeadersInner{value: val, isSet: true}
}

func (v NullableCustomHeadersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomHeadersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


