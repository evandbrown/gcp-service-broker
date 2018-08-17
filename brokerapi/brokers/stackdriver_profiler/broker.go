// Copyright 2018 Google Inc.
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

package stackdriver_profiler

import (
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/broker_base"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/models"
	"github.com/GoogleCloudPlatform/gcp-service-broker/utils"
	"github.com/pivotal-cf/brokerapi"
)

type StackdriverProfilerBroker struct {
	broker_base.BrokerBase
}

// No-op, no service is required for the profiler
func (b *StackdriverProfilerBroker) Provision(instanceId string, details brokerapi.ProvisionDetails, plan models.ServicePlan) (models.ServiceInstanceDetails, error) {
	return models.ServiceInstanceDetails{}, nil
}

// No-op, no service is required for the profiler
func (b *StackdriverProfilerBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails) error {
	return nil
}

// Creates a service account with access to Stackdriver Profiler
func (b *StackdriverProfilerBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (models.ServiceBindingCredentials, error) {
	out, err := utils.SetParameter(details.RawParameters, "role", "cloudprofiler.agent")
	if err != nil {
		return models.ServiceBindingCredentials{}, err
	}
	details.RawParameters = out

	// Create account
	newBinding, err := b.AccountManager.CreateCredentials(instanceID, bindingID, details, models.ServiceInstanceDetails{})

	if err != nil {
		return models.ServiceBindingCredentials{}, err
	}

	return newBinding, nil
}
