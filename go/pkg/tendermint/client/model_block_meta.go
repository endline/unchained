/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BlockMeta struct for BlockMeta
type BlockMeta struct {
	BlockId *BlockID `json:"block_id,omitempty"`
	BlockSize *int32 `json:"block_size,omitempty"`
	Header *BlockHeader `json:"header,omitempty"`
	NumTxs *string `json:"num_txs,omitempty"`
}

// NewBlockMeta instantiates a new BlockMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockMeta() *BlockMeta {
	this := BlockMeta{}
	return &this
}

// NewBlockMetaWithDefaults instantiates a new BlockMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockMetaWithDefaults() *BlockMeta {
	this := BlockMeta{}
	return &this
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *BlockMeta) GetBlockId() BlockID {
	if o == nil || o.BlockId == nil {
		var ret BlockID
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockMeta) GetBlockIdOk() (*BlockID, bool) {
	if o == nil || o.BlockId == nil {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *BlockMeta) HasBlockId() bool {
	if o != nil && o.BlockId != nil {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given BlockID and assigns it to the BlockId field.
func (o *BlockMeta) SetBlockId(v BlockID) {
	o.BlockId = &v
}

// GetBlockSize returns the BlockSize field value if set, zero value otherwise.
func (o *BlockMeta) GetBlockSize() int32 {
	if o == nil || o.BlockSize == nil {
		var ret int32
		return ret
	}
	return *o.BlockSize
}

// GetBlockSizeOk returns a tuple with the BlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockMeta) GetBlockSizeOk() (*int32, bool) {
	if o == nil || o.BlockSize == nil {
		return nil, false
	}
	return o.BlockSize, true
}

// HasBlockSize returns a boolean if a field has been set.
func (o *BlockMeta) HasBlockSize() bool {
	if o != nil && o.BlockSize != nil {
		return true
	}

	return false
}

// SetBlockSize gets a reference to the given int32 and assigns it to the BlockSize field.
func (o *BlockMeta) SetBlockSize(v int32) {
	o.BlockSize = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *BlockMeta) GetHeader() BlockHeader {
	if o == nil || o.Header == nil {
		var ret BlockHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockMeta) GetHeaderOk() (*BlockHeader, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *BlockMeta) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given BlockHeader and assigns it to the Header field.
func (o *BlockMeta) SetHeader(v BlockHeader) {
	o.Header = &v
}

// GetNumTxs returns the NumTxs field value if set, zero value otherwise.
func (o *BlockMeta) GetNumTxs() string {
	if o == nil || o.NumTxs == nil {
		var ret string
		return ret
	}
	return *o.NumTxs
}

// GetNumTxsOk returns a tuple with the NumTxs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockMeta) GetNumTxsOk() (*string, bool) {
	if o == nil || o.NumTxs == nil {
		return nil, false
	}
	return o.NumTxs, true
}

// HasNumTxs returns a boolean if a field has been set.
func (o *BlockMeta) HasNumTxs() bool {
	if o != nil && o.NumTxs != nil {
		return true
	}

	return false
}

// SetNumTxs gets a reference to the given string and assigns it to the NumTxs field.
func (o *BlockMeta) SetNumTxs(v string) {
	o.NumTxs = &v
}

func (o BlockMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockId != nil {
		toSerialize["block_id"] = o.BlockId
	}
	if o.BlockSize != nil {
		toSerialize["block_size"] = o.BlockSize
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.NumTxs != nil {
		toSerialize["num_txs"] = o.NumTxs
	}
	return json.Marshal(toSerialize)
}

type NullableBlockMeta struct {
	value *BlockMeta
	isSet bool
}

func (v NullableBlockMeta) Get() *BlockMeta {
	return v.value
}

func (v *NullableBlockMeta) Set(val *BlockMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockMeta(val *BlockMeta) *NullableBlockMeta {
	return &NullableBlockMeta{value: val, isSet: true}
}

func (v NullableBlockMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

