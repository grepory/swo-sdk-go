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

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

// Error struct for Error
type Error struct {
	// One of a server-defined set of error codes.
	Code string `json:"code"`
	// A human-readable representation of the error.
	Message string `json:"message"`
	// The target of the error.
	Target *string `json:"target,omitempty"`
	Details []Error `json:"details,omitempty"`
	InnerError *InnerError `json:"innerError,omitempty"`
}

type _Error Error

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError(code string, message string) *Error {
	this := Error{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetCode returns the Code field value
func (o *Error) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Error) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Error) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *Error) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Error) SetMessage(v string) {
	o.Message = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Error) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Error) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *Error) SetTarget(v string) {
	o.Target = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Error) GetDetails() []Error {
	if o == nil || IsNil(o.Details) {
		var ret []Error
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetDetailsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Error) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []Error and assigns it to the Details field.
func (o *Error) SetDetails(v []Error) {
	o.Details = v
}

// GetInnerError returns the InnerError field value if set, zero value otherwise.
func (o *Error) GetInnerError() InnerError {
	if o == nil || IsNil(o.InnerError) {
		var ret InnerError
		return ret
	}
	return *o.InnerError
}

// GetInnerErrorOk returns a tuple with the InnerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetInnerErrorOk() (*InnerError, bool) {
	if o == nil || IsNil(o.InnerError) {
		return nil, false
	}
	return o.InnerError, true
}

// HasInnerError returns a boolean if a field has been set.
func (o *Error) HasInnerError() bool {
	if o != nil && !IsNil(o.InnerError) {
		return true
	}

	return false
}

// SetInnerError gets a reference to the given InnerError and assigns it to the InnerError field.
func (o *Error) SetInnerError(v InnerError) {
	o.InnerError = &v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.InnerError) {
		toSerialize["innerError"] = o.InnerError
	}
	return toSerialize, nil
}

func (o *Error) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varError := _Error{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varError)

	if err != nil {
		return err
	}

	*o = Error(varError)

	return err
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

