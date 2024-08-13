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
	"bytes"
	"fmt"
)

// checks if the Asset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Asset{}

// Asset An asset onchain scoped to a particular network, e.g. ETH on base-sepolia, or the USDC ERC20 Token on ethereum-mainnet.
type Asset struct {
	// The ID of the blockchain network
	NetworkId string `json:"network_id"`
	// The ID for the asset on the network
	AssetId string `json:"asset_id"`
	// The number of decimals the asset supports. This is used to convert from atomic units to base units.
	Decimals *int32 `json:"decimals,omitempty"`
	// The optional contract address for the asset. This will be specified for smart contract-based assets, for example ERC20s.
	ContractAddress *string `json:"contract_address,omitempty"`
}

type _Asset Asset

// NewAsset instantiates a new Asset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsset(networkId string, assetId string) *Asset {
	this := Asset{}
	this.NetworkId = networkId
	this.AssetId = assetId
	return &this
}

// NewAssetWithDefaults instantiates a new Asset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithDefaults() *Asset {
	this := Asset{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *Asset) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *Asset) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *Asset) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetAssetId returns the AssetId field value
func (o *Asset) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *Asset) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *Asset) SetAssetId(v string) {
	o.AssetId = v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *Asset) GetDecimals() int32 {
	if o == nil || IsNil(o.Decimals) {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetDecimalsOk() (*int32, bool) {
	if o == nil || IsNil(o.Decimals) {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *Asset) HasDecimals() bool {
	if o != nil && !IsNil(o.Decimals) {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *Asset) SetDecimals(v int32) {
	o.Decimals = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *Asset) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Asset) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *Asset) SetContractAddress(v string) {
	o.ContractAddress = &v
}

func (o Asset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Asset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_id"] = o.NetworkId
	toSerialize["asset_id"] = o.AssetId
	if !IsNil(o.Decimals) {
		toSerialize["decimals"] = o.Decimals
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	return toSerialize, nil
}

func (o *Asset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_id",
		"asset_id",
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

	varAsset := _Asset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAsset)

	if err != nil {
		return err
	}

	*o = Asset(varAsset)

	return err
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


