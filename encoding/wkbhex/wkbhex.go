// Package wkbhex implements Well Known Binary encoding and decoding of
// strings.
package wkbhex

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/matoous/goodgeo"
	"github.com/matoous/goodgeo/encoding/wkb"
	"github.com/matoous/goodgeo/encoding/wkbcommon"
)

var (
	// XDR is big endian.
	XDR = wkb.XDR
	// NDR is little endian.
	NDR = wkb.NDR
)

// Encode encodes an arbitrary geometry to a string.
func Encode(g goodgeo.T, byteOrder binary.ByteOrder, opts ...wkbcommon.WKBOption) (string, error) {
	wkb, err := wkb.Marshal(g, byteOrder, opts...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(wkb), nil
}

// Decode decodes an arbitrary geometry from a string.
func Decode(s string, opts ...wkbcommon.WKBOption) (goodgeo.T, error) {
	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return wkb.Unmarshal(data, opts...)
}
