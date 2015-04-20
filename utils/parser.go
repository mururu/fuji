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
	"fmt"
	"strconv"
	"strings"
)

// parsePayload parse payload setting.
// ex:
//   payload = \x00\x03\xff
func ParsePayload(arg string) ([]byte, error) {
	if strings.Contains(arg, `\x`) {
		if len(arg)%4 != 0 { // something wrong
			return []byte{}, fmt.Errorf("invalid binary is specified")
		}
		ret := make([]byte, 0, len(arg)/4)
		for i := 0; i < len(arg); i += 4 {
			if arg[i:i+2] != `\x` {
				return ret, fmt.Errorf("invalid prefix, %v", arg[i:i+4])
			}
			tmp := arg[i+2 : i+4] // get only 01 from \x01
			b, err := strconv.ParseInt(tmp, 16, 64)
			if err != nil { // could not parse
				return ret, fmt.Errorf("could not parse binary payload, %v", tmp)
			}
			ret = append(ret, byte(b))
		}
		return ret, nil
	}
	return []byte(arg), nil
}
