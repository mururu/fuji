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

package message

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TopicString_Validate(t *testing.T) {
	assert := assert.New(t)

	// normal
	t0 := TopicString{Str: "topicprefix/gateway/device/type"}
	assert.Nil(t0.Validate())
	t1 := TopicString{Str: "a"}
	assert.Nil(t1.Validate())

	// topic too long
	t2 := TopicString{Str: strings.Repeat("a", 1000000)}
	assert.NotNil(t2.Validate())

	// topic has # wild card for publish
	t3 := TopicString{Str: "topicprefix/#/device/type"}
	assert.NotNil(t3.Validate())

	// topic has + wild card for publish
	t4 := TopicString{Str: "topicprefix/+/device/type"}
	assert.NotNil(t4.Validate())

	// topic has null string
	t5 := TopicString{
		Str: fmt.Sprintf("topicprefix/gatewayname%c/device/type", '\u0000'),
	}
	assert.NotNil(t5.Validate())
}
