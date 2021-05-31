// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// Package gcompress provides kinds of compression algorithms for binary/bytes data.
package gcompress

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
)

// Zlib compresses <data> with zlib algorithm.
func Zlib(data []byte) ([]byte, error) {
	if data == nil || len(data) < 13 {
		return data, nil
	}
	var in bytes.Buffer
	var err error
	w := zlib.NewWriter(&in)
	if _, err = w.Write(data); err != nil {
		return nil, err
	}
	if err = w.Close(); err != nil {
		return in.Bytes(), err
	}
	return in.Bytes(), nil
}

// UnZlib decompresses <data> with zlib algorithm.
func UnZlib(data []byte) ([]byte, error) {
	if data == nil || len(data) < 13 {
		return data, nil
	}

	b := bytes.NewReader(data)
	var out bytes.Buffer
	var err error
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(&out, r); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// Gzip compresses <data> with gzip algorithm.
func Gzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	var err error
	zip := gzip.NewWriter(&buf)
	_, err = zip.Write(data)
	if err != nil {
		return nil, err
	}
	if err = zip.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnGzip decompresses <data> with gzip algorithm.
func UnGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	content := bytes.NewReader(data)
	zipData, err := gzip.NewReader(content)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(&buf, zipData); err != nil {
		return nil, err
	}
	if err = zipData.Close(); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
