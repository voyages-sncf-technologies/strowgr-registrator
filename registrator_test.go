/*
 *  Copyright (C) 2016 VSCT
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */
package main

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	registrator "github.com/voyages-sncf-technologies/strowgr-registrator/internal"
	"reflect"
	"runtime"
	"testing"
)

func instanceFixture() *registrator.RegisterCommand {
	instance := registrator.NewInstance()
	instance.Header.Application = "Test"
	instance.Header.Platform = "TST"
	instance.Server.BackendId = "BACK"
	instance.Server.Port = "1234"
	instance.Server.Ip = "1.2.3.4"

	return instance
}

func containerJSONFixture() types.ContainerJSON {
	return types.ContainerJSON{
		ContainerJSONBase: &types.ContainerJSONBase{
			Name: "test",
		},
	}
}

func TestDefaultNamingStrategy(t *testing.T) {

	strategyFunc := defaultNamingStrategy

	expected := "1_2_3_4_test_1234"
	result := strategyFunc(containerJSONFixture(), instanceFixture())
	AssertEquals(t, expected, result)
}

func TestNamingStrategy(t *testing.T) {

	strategyFunc := defaultNamingStrategy

	expected := "1_2_3_4_test_1234"
	result := strategyFunc(containerJSONFixture(), instanceFixture())
	AssertEquals(t, expected, result)
}

func TestContainerNamingStrategySelector_default(t *testing.T) {
	var config = &container.Config{
		Labels: map[string]string{
			"registrator.id_generator": "pouet",
		},
	}

	result := getNamingStrategy(config)
	expected := defaultNamingStrategy
	AssertEquals(t, GetFunctionName(expected), GetFunctionName(result))
}

func TestContainerNamingStrategySelector_container(t *testing.T) {
	var config = &container.Config{
		Labels: map[string]string{
			"registrator.id_generator": "container_name",
		},
	}

	result := getNamingStrategy(config)
	expected := containerNamingStrategy
	AssertEquals(t, GetFunctionName(expected), GetFunctionName(result))
}

func AssertEquals(t *testing.T, expected interface{}, result interface{}) {
	if result != expected {
		t.Logf("Expected '%s', got '%s'", expected, result)
		t.Fail()
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
