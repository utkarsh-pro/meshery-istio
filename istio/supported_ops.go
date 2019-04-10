// Copyright 2019 Layer5.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package istio

type supportedOperation struct {
	// a friendly name
	name string
	// the template file name
	templateName string
	// the app label
	appLabel string
	// // returnLogs specifies if the operation logs should be returned
	// returnLogs bool
}

const (
	customOpCommand        = "custom"
	runVet                 = "istio_vet"
	installIstioCommand    = "istio_install"
	installBookInfoCommand = "install_book_info"
	cbCommand              = "cb1"
)

var supportedOps = map[string]supportedOperation{
	installIstioCommand: {
		name: "Install the latest version of Istio",
		// templateName: "install_istio.tmpl",
	},
	installBookInfoCommand: {
		name: "Install the canonical Book Info Application",
		// templateName: "install_istio.tmpl",
	},
	runVet: {
		name: "Run istio-vet",
		// templateName: "istio_vet.tmpl",
		// appLabel:     "istio-vet",
		// returnLogs:   true,
	},
	cbCommand: {
		name:         "Limit circuit breaker config to one connection",
		templateName: "circuit_breaking.tmpl",
	},
	customOpCommand: {
		name: "Custom YAML",
	},
}