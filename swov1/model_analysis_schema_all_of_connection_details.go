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

// checks if the AnalysisSchemaAllOfConnectionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisSchemaAllOfConnectionDetails{}

// AnalysisSchemaAllOfConnectionDetails struct for AnalysisSchemaAllOfConnectionDetails
type AnalysisSchemaAllOfConnectionDetails struct {
	DurationInMs int32 `json:"durationInMs"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	IpAddress string `json:"ipAddress"`
	Port int32 `json:"port"`
	Protocol string `json:"protocol"`
}

type _AnalysisSchemaAllOfConnectionDetails AnalysisSchemaAllOfConnectionDetails

// NewAnalysisSchemaAllOfConnectionDetails instantiates a new AnalysisSchemaAllOfConnectionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisSchemaAllOfConnectionDetails(durationInMs int32, ipAddress string, port int32, protocol string) *AnalysisSchemaAllOfConnectionDetails {
	this := AnalysisSchemaAllOfConnectionDetails{}
	this.DurationInMs = durationInMs
	this.IpAddress = ipAddress
	this.Port = port
	this.Protocol = protocol
	return &this
}

// NewAnalysisSchemaAllOfConnectionDetailsWithDefaults instantiates a new AnalysisSchemaAllOfConnectionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisSchemaAllOfConnectionDetailsWithDefaults() *AnalysisSchemaAllOfConnectionDetails {
	this := AnalysisSchemaAllOfConnectionDetails{}
	return &this
}

// GetDurationInMs returns the DurationInMs field value
func (o *AnalysisSchemaAllOfConnectionDetails) GetDurationInMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationInMs
}

// GetDurationInMsOk returns a tuple with the DurationInMs field value
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfConnectionDetails) GetDurationInMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationInMs, true
}

// SetDurationInMs sets field value
func (o *AnalysisSchemaAllOfConnectionDetails) SetDurationInMs(v int32) {
	o.DurationInMs = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AnalysisSchemaAllOfConnectionDetails) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfConnectionDetails) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AnalysisSchemaAllOfConnectionDetails) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AnalysisSchemaAllOfConnectionDetails) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIpAddress returns the IpAddress field value
func (o *AnalysisSchemaAllOfConnectionDetails) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfConnectionDetails) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *AnalysisSchemaAllOfConnectionDetails) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetPort returns the Port field value
func (o *AnalysisSchemaAllOfConnectionDetails) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfConnectionDetails) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *AnalysisSchemaAllOfConnectionDetails) SetPort(v int32) {
	o.Port = v
}

// GetProtocol returns the Protocol field value
func (o *AnalysisSchemaAllOfConnectionDetails) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *AnalysisSchemaAllOfConnectionDetails) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *AnalysisSchemaAllOfConnectionDetails) SetProtocol(v string) {
	o.Protocol = v
}

func (o AnalysisSchemaAllOfConnectionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisSchemaAllOfConnectionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["durationInMs"] = o.DurationInMs
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	toSerialize["ipAddress"] = o.IpAddress
	toSerialize["port"] = o.Port
	toSerialize["protocol"] = o.Protocol
	return toSerialize, nil
}

func (o *AnalysisSchemaAllOfConnectionDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"durationInMs",
		"ipAddress",
		"port",
		"protocol",
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

	varAnalysisSchemaAllOfConnectionDetails := _AnalysisSchemaAllOfConnectionDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalysisSchemaAllOfConnectionDetails)

	if err != nil {
		return err
	}

	*o = AnalysisSchemaAllOfConnectionDetails(varAnalysisSchemaAllOfConnectionDetails)

	return err
}

type NullableAnalysisSchemaAllOfConnectionDetails struct {
	value *AnalysisSchemaAllOfConnectionDetails
	isSet bool
}

func (v NullableAnalysisSchemaAllOfConnectionDetails) Get() *AnalysisSchemaAllOfConnectionDetails {
	return v.value
}

func (v *NullableAnalysisSchemaAllOfConnectionDetails) Set(val *AnalysisSchemaAllOfConnectionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisSchemaAllOfConnectionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisSchemaAllOfConnectionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisSchemaAllOfConnectionDetails(val *AnalysisSchemaAllOfConnectionDetails) *NullableAnalysisSchemaAllOfConnectionDetails {
	return &NullableAnalysisSchemaAllOfConnectionDetails{value: val, isSet: true}
}

func (v NullableAnalysisSchemaAllOfConnectionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisSchemaAllOfConnectionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


