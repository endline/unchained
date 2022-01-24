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

// TxSearchResponseResultProof struct for TxSearchResponseResultProof
type TxSearchResponseResultProof struct {
	RootHash string `json:"RootHash"`
	Data string `json:"Data"`
	Proof TxSearchResponseResultProofProof `json:"Proof"`
}

// NewTxSearchResponseResultProof instantiates a new TxSearchResponseResultProof object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxSearchResponseResultProof(rootHash string, data string, proof TxSearchResponseResultProofProof) *TxSearchResponseResultProof {
	this := TxSearchResponseResultProof{}
	this.RootHash = rootHash
	this.Data = data
	this.Proof = proof
	return &this
}

// NewTxSearchResponseResultProofWithDefaults instantiates a new TxSearchResponseResultProof object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxSearchResponseResultProofWithDefaults() *TxSearchResponseResultProof {
	this := TxSearchResponseResultProof{}
	return &this
}

// GetRootHash returns the RootHash field value
func (o *TxSearchResponseResultProof) GetRootHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootHash
}

// GetRootHashOk returns a tuple with the RootHash field value
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultProof) GetRootHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RootHash, true
}

// SetRootHash sets field value
func (o *TxSearchResponseResultProof) SetRootHash(v string) {
	o.RootHash = v
}

// GetData returns the Data field value
func (o *TxSearchResponseResultProof) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultProof) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TxSearchResponseResultProof) SetData(v string) {
	o.Data = v
}

// GetProof returns the Proof field value
func (o *TxSearchResponseResultProof) GetProof() TxSearchResponseResultProofProof {
	if o == nil {
		var ret TxSearchResponseResultProofProof
		return ret
	}

	return o.Proof
}

// GetProofOk returns a tuple with the Proof field value
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultProof) GetProofOk() (*TxSearchResponseResultProofProof, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Proof, true
}

// SetProof sets field value
func (o *TxSearchResponseResultProof) SetProof(v TxSearchResponseResultProofProof) {
	o.Proof = v
}

func (o TxSearchResponseResultProof) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["RootHash"] = o.RootHash
	}
	if true {
		toSerialize["Data"] = o.Data
	}
	if true {
		toSerialize["Proof"] = o.Proof
	}
	return json.Marshal(toSerialize)
}

type NullableTxSearchResponseResultProof struct {
	value *TxSearchResponseResultProof
	isSet bool
}

func (v NullableTxSearchResponseResultProof) Get() *TxSearchResponseResultProof {
	return v.value
}

func (v *NullableTxSearchResponseResultProof) Set(val *TxSearchResponseResultProof) {
	v.value = val
	v.isSet = true
}

func (v NullableTxSearchResponseResultProof) IsSet() bool {
	return v.isSet
}

func (v *NullableTxSearchResponseResultProof) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxSearchResponseResultProof(val *TxSearchResponseResultProof) *NullableTxSearchResponseResultProof {
	return &NullableTxSearchResponseResultProof{value: val, isSet: true}
}

func (v NullableTxSearchResponseResultProof) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxSearchResponseResultProof) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

