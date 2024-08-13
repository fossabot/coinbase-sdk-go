/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
Contact: yuga.cohler@coinbase.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the FetchStakingRewardsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchStakingRewardsRequest{}

// FetchStakingRewardsRequest struct for FetchStakingRewardsRequest
type FetchStakingRewardsRequest struct {
	// The ID of the blockchain network
	NetworkId string `json:"network_id"`
	// The ID of the asset for which the staking rewards are being fetched
	AssetId string `json:"asset_id"`
	// The onchain addresses for which the staking rewards are being fetched
	AddressIds []string `json:"address_ids"`
	// The start time of this reward period
	StartTime time.Time `json:"start_time"`
	// The end time of this reward period
	EndTime time.Time `json:"end_time"`
	Format StakingRewardFormat `json:"format"`
}

type _FetchStakingRewardsRequest FetchStakingRewardsRequest

// NewFetchStakingRewardsRequest instantiates a new FetchStakingRewardsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchStakingRewardsRequest(networkId string, assetId string, addressIds []string, startTime time.Time, endTime time.Time, format StakingRewardFormat) *FetchStakingRewardsRequest {
	this := FetchStakingRewardsRequest{}
	this.NetworkId = networkId
	this.AssetId = assetId
	this.AddressIds = addressIds
	this.StartTime = startTime
	this.EndTime = endTime
	this.Format = format
	return &this
}

// NewFetchStakingRewardsRequestWithDefaults instantiates a new FetchStakingRewardsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchStakingRewardsRequestWithDefaults() *FetchStakingRewardsRequest {
	this := FetchStakingRewardsRequest{}
	var format StakingRewardFormat = STAKINGREWARDFORMAT_USD
	this.Format = format
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *FetchStakingRewardsRequest) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewardsRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *FetchStakingRewardsRequest) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetAssetId returns the AssetId field value
func (o *FetchStakingRewardsRequest) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewardsRequest) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *FetchStakingRewardsRequest) SetAssetId(v string) {
	o.AssetId = v
}

// GetAddressIds returns the AddressIds field value
func (o *FetchStakingRewardsRequest) GetAddressIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AddressIds
}

// GetAddressIdsOk returns a tuple with the AddressIds field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewardsRequest) GetAddressIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressIds, true
}

// SetAddressIds sets field value
func (o *FetchStakingRewardsRequest) SetAddressIds(v []string) {
	o.AddressIds = v
}

// GetStartTime returns the StartTime field value
func (o *FetchStakingRewardsRequest) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewardsRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *FetchStakingRewardsRequest) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *FetchStakingRewardsRequest) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewardsRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *FetchStakingRewardsRequest) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetFormat returns the Format field value
func (o *FetchStakingRewardsRequest) GetFormat() StakingRewardFormat {
	if o == nil {
		var ret StakingRewardFormat
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewardsRequest) GetFormatOk() (*StakingRewardFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *FetchStakingRewardsRequest) SetFormat(v StakingRewardFormat) {
	o.Format = v
}

func (o FetchStakingRewardsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchStakingRewardsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_id"] = o.NetworkId
	toSerialize["asset_id"] = o.AssetId
	toSerialize["address_ids"] = o.AddressIds
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	toSerialize["format"] = o.Format
	return toSerialize, nil
}

func (o *FetchStakingRewardsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_id",
		"asset_id",
		"address_ids",
		"start_time",
		"end_time",
		"format",
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

	varFetchStakingRewardsRequest := _FetchStakingRewardsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchStakingRewardsRequest)

	if err != nil {
		return err
	}

	*o = FetchStakingRewardsRequest(varFetchStakingRewardsRequest)

	return err
}

type NullableFetchStakingRewardsRequest struct {
	value *FetchStakingRewardsRequest
	isSet bool
}

func (v NullableFetchStakingRewardsRequest) Get() *FetchStakingRewardsRequest {
	return v.value
}

func (v *NullableFetchStakingRewardsRequest) Set(val *FetchStakingRewardsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchStakingRewardsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchStakingRewardsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchStakingRewardsRequest(val *FetchStakingRewardsRequest) *NullableFetchStakingRewardsRequest {
	return &NullableFetchStakingRewardsRequest{value: val, isSet: true}
}

func (v NullableFetchStakingRewardsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchStakingRewardsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


