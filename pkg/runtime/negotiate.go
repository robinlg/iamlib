// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/robinlg/onexblog. The professional
// version of this repository is https://github.com/robinlg/onexblog.

package runtime

import (
	"fmt"

	"github.com/robinlg/iamlib/pkg/json"
)

// NegotiateError is returned when a ClientNegotiator is unable to locate
// a serializer for the requested operation.
type NegotiateError struct {
	ContentType string
	Stream      bool
}

func (e NegotiateError) Error() string {
	if e.Stream {
		return fmt.Sprintf("no stream serializers registered for %s", e.ContentType)
	}
	return fmt.Sprintf("no serializers registered for %s", e.ContentType)
}

type apimachineryClientNegotiator struct{}

var _ ClientNegotiator = &apimachineryClientNegotiator{}

func (n *apimachineryClientNegotiator) Encoder() (Encoder, error) {
	return &apimachineryClientNegotiatorSerializer{}, nil
}

func (n *apimachineryClientNegotiator) Decoder() (Decoder, error) {
	return &apimachineryClientNegotiatorSerializer{}, nil
}

type apimachineryClientNegotiatorSerializer struct{}

var _ Decoder = &apimachineryClientNegotiatorSerializer{}

func (s *apimachineryClientNegotiatorSerializer) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (s *apimachineryClientNegotiatorSerializer) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// NewSimpleClientNegotiator will negotiate for a single serializer. This should only be used
// for testing or when the caller is taking responsibility for setting the GVK on encoded objects.
func NewSimpleClientNegotiator() ClientNegotiator {
	return &apimachineryClientNegotiator{}
}
