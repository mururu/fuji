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

package inidef

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
)

type AnyError interface{}

type Error string

func (e Error) Error() string {
	return string(e)
}

// NilOrString defines the value is nil or empty
type NilOrString interface{}

func IsNil(str NilOrString) bool {
	if str == nil {
		return true
	}
	return false
}
func String(str NilOrString) string {
	stringValue, ok := str.(string)
	if ok == false {
		return ("nil")
	}
	return stringValue
}

// ValidMqttPublishTopic validates the Topic is validate or not
// This is used with validator packages.
func ValidMqttPublishTopic(v interface{}, param string) error {
	str := reflect.ValueOf(v)
	if str.Kind() != reflect.String {
		return errors.New("ValidMqttPublishTopic only validates strings")
	}
	if !utf8.ValidString(str.String()) {
		return errors.New("not a valid UTF8 string")
	}
	reu0 := regexp.MustCompile("\u0000")
	if reu0.FindString(str.String()) != "" {
		return errors.New("Topic SHALL NOT include \\U0000 character")
	}
	rewild := regexp.MustCompile("[+#]+")
	if rewild.FindString(str.String()) != "" {
		return errors.New("SHALL NOT MQTT pub-topic include wildard character")
	}
	return nil
}

type Config struct {
	GatewayName string
	BrokerNames []string

	Sections []ConfigSection
}

type ConfigSection struct {
	Title string
	Type  string
	Name  string
	Arg   string

	Values map[string]string
}

// Loadini loads ini format file from confPath arg and returns []ConfigSection.
// ConfigSection has a Type, Name and arg.
// example:
// [broker "sango"]
// [broker "sango/1"]
// [broker "sango/2"]
//
// ret = [
//   ConfigSection{Type: "broker", Name: "sango"},
//   ConfigSection{Type: "broker", Name: "sango", Arg: "1"},
//   ConfigSection{Type: "broker", Name: "sango", Arg: "2"},
// ]
func LoadConfig(confPath string) (Config, error) {
	dat, err := ioutil.ReadFile(confPath)
	if err != nil {
		return Config{}, err
	}

	return LoadConfigByte(dat)
}

// LoadConfigByte returnes []ConfigSection from []byte.
// This is invoked from LoadConfig.
func LoadConfigByte(conf []byte) (Config, error) {
	cfg, err := ini.Load(conf)

	config := Config{}

	var sections []ConfigSection
	var bn []string

	for _, section := range cfg.Sections() {
		key := section.Name()

		k := strings.Fields(key)
		if len(k) > 2 {
			log.Errorf("invalid section(space), %v", k)
			continue
		}
		rt := ConfigSection{
			Title:  key,
			Type:   k[0],
			Values: section.KeysHash(),
		}
		if rt.Type == "gateway" {
			name, ok := rt.Values["name"]
			if !ok {
				return config, fmt.Errorf("gateway has not name")
			}
			config.GatewayName = name
		}

		// type only
		if len(k) == 1 {
			sections = append(sections, rt)
			continue
		}

		// parse name and args
		t := strings.TrimFunc(k[1], func(c rune) bool {
			if c == '"' {
				return true
			}
			return false
		})
		tt := strings.Split(t, "/")
		if len(tt) > 2 {
			log.Errorf("invalid section(slash), %v", t)
			continue
		}
		rt.Name = tt[0]
		if len(tt) == 2 { // if args exists, store it
			rt.Arg = tt[1]
		}
		if rt.Type == "broker" {
			bn = append(bn, rt.Name)
		}
		sections = append(sections, rt)
	}

	config.Sections = sections
	config.BrokerNames = bn

	return config, err
}
