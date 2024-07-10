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

// checks if the ListEntityTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEntityTypesResponse{}

// ListEntityTypesResponse struct for ListEntityTypesResponse
type ListEntityTypesResponse struct {
	// Array of entity types
	Types []string `json:"types"`
}

type _ListEntityTypesResponse ListEntityTypesResponse

// NewListEntityTypesResponse instantiates a new ListEntityTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEntityTypesResponse(types []string) *ListEntityTypesResponse {
	this := ListEntityTypesResponse{}
	this.Types = types
	return &this
}

// NewListEntityTypesResponseWithDefaults instantiates a new ListEntityTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEntityTypesResponseWithDefaults() *ListEntityTypesResponse {
	this := ListEntityTypesResponse{}
	return &this
}

// GetTypes returns the Types field value
func (o *ListEntityTypesResponse) GetTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *ListEntityTypesResponse) GetTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *ListEntityTypesResponse) SetTypes(v []string) {
	o.Types = v
}

func (o ListEntityTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEntityTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["types"] = o.Types
	return toSerialize, nil
}

func (o *ListEntityTypesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"types",
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

	varListEntityTypesResponse := _ListEntityTypesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListEntityTypesResponse)

	if err != nil {
		return err
	}

	*o = ListEntityTypesResponse(varListEntityTypesResponse)

	return err
}

type NullableListEntityTypesResponse struct {
	value *ListEntityTypesResponse
	isSet bool
}

func (v NullableListEntityTypesResponse) Get() *ListEntityTypesResponse {
	return v.value
}

func (v *NullableListEntityTypesResponse) Set(val *ListEntityTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListEntityTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListEntityTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEntityTypesResponse(val *ListEntityTypesResponse) *NullableListEntityTypesResponse {
	return &NullableListEntityTypesResponse{value: val, isSet: true}
}

func (v NullableListEntityTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEntityTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


