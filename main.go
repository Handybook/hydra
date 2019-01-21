/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

//go:generate swagger generate spec

package main

import (
	"github.com/ory/hydra/cmd"
	"github.com/ory/x/profilex"
	"github.com/newrelic/go-agent"
	"os"
)

func main() {
	defer profilex.Profile().Stop()

	config := newrelic.NewConfig("hydra", "bcc28e299818f07c28b71493458b89c072e274b9")
	config.Enabled = os.Getenv("KUBE_ENVIRONMENT") == "production"
	_, err := newrelic.NewApplication(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute()
}
