// Copyright Monax Industries Limited
// SPDX-License-Identifier: Apache-2.0

package rpc

import (
	"io"
)

// Used for rpc request and response data.
type Codec interface {
	EncodeBytes(interface{}) ([]byte, error)
	Encode(interface{}, io.Writer) error
	DecodeBytes(interface{}, []byte) error
	Decode(interface{}, io.Reader) error
}
