// Copyright 2015 Shiguredo Inc. <fuji@shiguredo.jp>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePayload(t *testing.T) {
	assert := assert.New(t)

	ret, err := ParsePayload("hoge")
	assert.Nil(err)
	assert.Equal([]byte("hoge"), ret)

	ret, err = ParsePayload(`\x11\x02\xff`)
	assert.Nil(err)
	assert.Equal([]byte{17, 2, 255}, ret)

	// invalid byte sequence, wrong length
	ret, err = ParsePayload(`\x01\x0211`)
	assert.NotNil(err)
	assert.Equal([]byte{}, ret)

	// invalid byte sequence
	ret, err = ParsePayload(`\x01\xmm`)
	assert.NotNil(err)
	assert.Equal([]byte{1}, ret)
}
