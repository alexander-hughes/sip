/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package services

import (
	airshipv1 "sipcluster/pkg/api/v1"
)

type AuthHost struct {
	Service
}

func newAuthHost(infraCfg airshipv1.InfraConfig) InfrastructureService {
	authhost := &AuthHost{
		Service: Service{
			serviceName: airshipv1.AuthHostService,
			config:      infraCfg,
		},
	}
	return authhost
}
