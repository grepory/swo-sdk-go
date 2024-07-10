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

// checks if the ListMetricMeasurementsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetricMeasurementsResponse{}

// ListMetricMeasurementsResponse struct for ListMetricMeasurementsResponse
type ListMetricMeasurementsResponse struct {
	Groupings []MeasurementGrouping `json:"groupings"`
	BucketSizeInSeconds int32 `json:"bucketSizeInSeconds"`
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
}

type _ListMetricMeasurementsResponse ListMetricMeasurementsResponse

// NewListMetricMeasurementsResponse instantiates a new ListMetricMeasurementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetricMeasurementsResponse(groupings []MeasurementGrouping, bucketSizeInSeconds int32) *ListMetricMeasurementsResponse {
	this := ListMetricMeasurementsResponse{}
	this.Groupings = groupings
	this.BucketSizeInSeconds = bucketSizeInSeconds
	return &this
}

// NewListMetricMeasurementsResponseWithDefaults instantiates a new ListMetricMeasurementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetricMeasurementsResponseWithDefaults() *ListMetricMeasurementsResponse {
	this := ListMetricMeasurementsResponse{}
	return &this
}

// GetGroupings returns the Groupings field value
func (o *ListMetricMeasurementsResponse) GetGroupings() []MeasurementGrouping {
	if o == nil {
		var ret []MeasurementGrouping
		return ret
	}

	return o.Groupings
}

// GetGroupingsOk returns a tuple with the Groupings field value
// and a boolean to check if the value has been set.
func (o *ListMetricMeasurementsResponse) GetGroupingsOk() ([]MeasurementGrouping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groupings, true
}

// SetGroupings sets field value
func (o *ListMetricMeasurementsResponse) SetGroupings(v []MeasurementGrouping) {
	o.Groupings = v
}

// GetBucketSizeInSeconds returns the BucketSizeInSeconds field value
func (o *ListMetricMeasurementsResponse) GetBucketSizeInSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BucketSizeInSeconds
}

// GetBucketSizeInSecondsOk returns a tuple with the BucketSizeInSeconds field value
// and a boolean to check if the value has been set.
func (o *ListMetricMeasurementsResponse) GetBucketSizeInSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketSizeInSeconds, true
}

// SetBucketSizeInSeconds sets field value
func (o *ListMetricMeasurementsResponse) SetBucketSizeInSeconds(v int32) {
	o.BucketSizeInSeconds = v
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *ListMetricMeasurementsResponse) GetPageInfo() PageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricMeasurementsResponse) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *ListMetricMeasurementsResponse) HasPageInfo() bool {
	if o != nil && !IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PageInfo and assigns it to the PageInfo field.
func (o *ListMetricMeasurementsResponse) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

func (o ListMetricMeasurementsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMetricMeasurementsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupings"] = o.Groupings
	toSerialize["bucketSizeInSeconds"] = o.BucketSizeInSeconds
	if !IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	return toSerialize, nil
}

func (o *ListMetricMeasurementsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupings",
		"bucketSizeInSeconds",
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

	varListMetricMeasurementsResponse := _ListMetricMeasurementsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListMetricMeasurementsResponse)

	if err != nil {
		return err
	}

	*o = ListMetricMeasurementsResponse(varListMetricMeasurementsResponse)

	return err
}

type NullableListMetricMeasurementsResponse struct {
	value *ListMetricMeasurementsResponse
	isSet bool
}

func (v NullableListMetricMeasurementsResponse) Get() *ListMetricMeasurementsResponse {
	return v.value
}

func (v *NullableListMetricMeasurementsResponse) Set(val *ListMetricMeasurementsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetricMeasurementsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetricMeasurementsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetricMeasurementsResponse(val *ListMetricMeasurementsResponse) *NullableListMetricMeasurementsResponse {
	return &NullableListMetricMeasurementsResponse{value: val, isSet: true}
}

func (v NullableListMetricMeasurementsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetricMeasurementsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

