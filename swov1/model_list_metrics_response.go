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

// checks if the ListMetricsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetricsResponse{}

// ListMetricsResponse struct for ListMetricsResponse
type ListMetricsResponse struct {
	MetricsInfo []GetMetricResponse `json:"metricsInfo"`
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
}

type _ListMetricsResponse ListMetricsResponse

// NewListMetricsResponse instantiates a new ListMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetricsResponse(metricsInfo []GetMetricResponse) *ListMetricsResponse {
	this := ListMetricsResponse{}
	this.MetricsInfo = metricsInfo
	return &this
}

// NewListMetricsResponseWithDefaults instantiates a new ListMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetricsResponseWithDefaults() *ListMetricsResponse {
	this := ListMetricsResponse{}
	return &this
}

// GetMetricsInfo returns the MetricsInfo field value
func (o *ListMetricsResponse) GetMetricsInfo() []GetMetricResponse {
	if o == nil {
		var ret []GetMetricResponse
		return ret
	}

	return o.MetricsInfo
}

// GetMetricsInfoOk returns a tuple with the MetricsInfo field value
// and a boolean to check if the value has been set.
func (o *ListMetricsResponse) GetMetricsInfoOk() ([]GetMetricResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricsInfo, true
}

// SetMetricsInfo sets field value
func (o *ListMetricsResponse) SetMetricsInfo(v []GetMetricResponse) {
	o.MetricsInfo = v
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *ListMetricsResponse) GetPageInfo() PageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricsResponse) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *ListMetricsResponse) HasPageInfo() bool {
	if o != nil && !IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PageInfo and assigns it to the PageInfo field.
func (o *ListMetricsResponse) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

func (o ListMetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMetricsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricsInfo"] = o.MetricsInfo
	if !IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	return toSerialize, nil
}

func (o *ListMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricsInfo",
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

	varListMetricsResponse := _ListMetricsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListMetricsResponse)

	if err != nil {
		return err
	}

	*o = ListMetricsResponse(varListMetricsResponse)

	return err
}

type NullableListMetricsResponse struct {
	value *ListMetricsResponse
	isSet bool
}

func (v NullableListMetricsResponse) Get() *ListMetricsResponse {
	return v.value
}

func (v *NullableListMetricsResponse) Set(val *ListMetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetricsResponse(val *ListMetricsResponse) *NullableListMetricsResponse {
	return &NullableListMetricsResponse{value: val, isSet: true}
}

func (v NullableListMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


