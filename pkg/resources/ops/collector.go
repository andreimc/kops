/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ops

import (
	"fmt"

	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/resources"
	"k8s.io/kops/pkg/resources/digitalocean"
	"k8s.io/kops/pkg/resources/gce"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/awsup"
	cloudgce "k8s.io/kops/upup/pkg/fi/cloudup/gce"
	"k8s.io/kops/upup/pkg/fi/cloudup/vsphere"
)

// ListResources collects the resources from the specified cloud
func ListResources(cloud fi.Cloud, clusterName string, region string) (map[string]*resources.Resource, error) {
	switch cloud.ProviderID() {
	case kops.CloudProviderAWS:
		return resources.ListResourcesAWS(cloud.(awsup.AWSCloud), clusterName)
	case kops.CloudProviderDO:
		return digitalocean.ListResources(cloud, clusterName)
	case kops.CloudProviderGCE:
		return gce.ListResourcesGCE(cloud.(cloudgce.GCECloud), clusterName, region)
	case kops.CloudProviderVSphere:
		return resources.ListResourcesVSphere(cloud.(*vsphere.VSphereCloud), clusterName)
	default:
		return nil, fmt.Errorf("delete on clusters on %q not (yet) supported", cloud.ProviderID())
	}
}
