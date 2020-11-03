/*
Copyright 2019 The Crossplane Authors.

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
package network

import (
	networkmgmt "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
	azure "github.com/crossplane/provider-azure/pkg/clients"
	"reflect"
)

// UpdateAzureFirewallStatusFromAzure updates the status related to the external
// Azure Firewall in the AzureFirewallStatus
func UpdateAzureFirewallStatusFromAzure(v *v1alpha3.AzureFirewall, az networkmgmt.AzureFirewall) {

	v.Status.State = toStringProvisioningState(az.ProvisioningState)
	v.Status.ID = azure.ToString(az.ID)
	v.Status.Etag = azure.ToString(az.Etag)
	v.Status.Type = azure.ToString(az.Type)
}

func toStringProvisioningState(provisioningState networkmgmt.ProvisioningState) string {
	return string(provisioningState)
}

func NewFireWallParameters(v *v1alpha3.AzureFirewall) networkmgmt.AzureFirewall {
	return networkmgmt.AzureFirewall{
		Location: azure.ToStringPtr(v.Spec.Location),
		Name:     azure.ToStringPtr(v.Spec.Name),
		ID:       azure.ToStringPtr(v.Spec.ID),
		Etag:     azure.ToStringPtr(v.Spec.Etag),
		Type:     azure.ToStringPtr(v.Spec.Type),
		Zones:    azure.ToStringArrayPtr(v.Spec.Zones),
		Tags:     azure.ToStringPtrMap(v.Spec.Tags),
		AzureFirewallPropertiesFormat: &networkmgmt.AzureFirewallPropertiesFormat{
			//ApplicationRuleCollections: nil,
			//NatRuleCollections:         nil,
			//NetworkRuleCollections:     nil,
			IPConfigurations:  setIPConfigurations(v.Spec.IPConfigurations),
			ProvisioningState: networkmgmt.ProvisioningState(v.Spec.ProvisioningState),
			ThreatIntelMode:   networkmgmt.AzureFirewallThreatIntelMode(v.Spec.ThreatIntelMode),
			VirtualHub:        setSubResource(v.Spec.VirtualHub),
			FirewallPolicy:    setSubResource(v.Spec.FirewallPolicy),
			HubIPAddresses:    setHubIpAddresses(v.Spec.HubIPAddresses),
		},
	}
}

func setHubIpAddresses(addresses *v1alpha3.HubIPAddresses) *networkmgmt.HubIPAddresses {
	var hubIpAddresses = new(networkmgmt.HubIPAddresses)
	if nil != addresses {
		hubIpAddresses.PrivateIPAddress = addresses.PrivateIPAddress
		for _, publicIpAddress := range *addresses.PublicIPAddresses {
			var azureFirewallPublicIPAddress = networkmgmt.AzureFirewallPublicIPAddress{}
			azureFirewallPublicIPAddress.Address = publicIpAddress.Address
			*hubIpAddresses.PublicIPAddresses = append(*hubIpAddresses.PublicIPAddresses, azureFirewallPublicIPAddress)
		}
	}
	return hubIpAddresses
}

func setIPConfigurations(configurations *[]v1alpha3.AzureFirewallIPConfiguration) *[]networkmgmt.AzureFirewallIPConfiguration {
	var azipc = new([]networkmgmt.AzureFirewallIPConfiguration)
	for _, c := range *configurations {
		var config = networkmgmt.AzureFirewallIPConfiguration{}
		config.Etag = c.Etag
		config.ID = c.ID
		config.Name = c.Name
		config.AzureFirewallIPConfigurationPropertiesFormat = new(networkmgmt.AzureFirewallIPConfigurationPropertiesFormat)
		if c.AzureFirewallIPConfigurationPropertiesFormat.PrivateIPAddress != nil {
			config.PrivateIPAddress = c.AzureFirewallIPConfigurationPropertiesFormat.PrivateIPAddress
		}
		if c.AzureFirewallIPConfigurationPropertiesFormat.ProvisioningState != nil {
			config.ProvisioningState = networkmgmt.ProvisioningState(*c.AzureFirewallIPConfigurationPropertiesFormat.ProvisioningState)
		}

		if nil != setSubResource(c.AzureFirewallIPConfigurationPropertiesFormat.PublicIPAddress) {
			config.AzureFirewallIPConfigurationPropertiesFormat.PublicIPAddress = setSubResource(c.AzureFirewallIPConfigurationPropertiesFormat.PublicIPAddress)
		}
		if nil != setSubResource(c.AzureFirewallIPConfigurationPropertiesFormat.Subnet) {
			config.AzureFirewallIPConfigurationPropertiesFormat.Subnet = setSubResource(c.AzureFirewallIPConfigurationPropertiesFormat.Subnet)
		}
		*azipc = append(*azipc, config)
	}
	return azipc
}

func setSubResource(sr *v1alpha3.SubResource) *networkmgmt.SubResource {
	if nil != sr {
		if nil != sr.ID {
			var subResource = new(networkmgmt.SubResource)
			subResource.ID = sr.ID
			return subResource
		}
	}
	return nil
}

func AzureFirewallNeedsUpdate(firewall *v1alpha3.AzureFirewall, az networkmgmt.AzureFirewall) bool {

	if !reflect.DeepEqual(firewall.Name, az.Name) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.Location, az.Location) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.Zones, az.Zones) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.Etag, az.Etag) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.FirewallPolicy, az.FirewallPolicy) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.HubIPAddresses, az.HubIPAddresses) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.VirtualHub, az.VirtualHub) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.Type, az.Type) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.ThreatIntelMode, az.ThreatIntelMode) {
		return true
	}
	if !reflect.DeepEqual(firewall.Spec.Tags, az.Tags) {
		return true
	}
	//TODO: Azure firewall rules needed to added here once completed with structs

	return false
}

// NewSecurityGroupParameters returns an Azure SecurityGroup object from a Security Group Spec
func NewAzureFirewallParameters(v *v1alpha3.AzureFirewall) networkmgmt.AzureFirewall {
	return networkmgmt.AzureFirewall{
		Zones:    azure.ToStringArrayPtr(v.Spec.Zones),
		Etag:     azure.ToStringPtr(v.Spec.Etag),
		ID:       azure.ToStringPtr(v.Spec.ID),
		Name:     azure.ToStringPtr(v.Name),
		Type:     azure.ToStringPtr(v.Spec.Type),
		Location: azure.ToStringPtr(v.Spec.Location),
		Tags:     azure.ToStringPtrMap(v.Spec.Tags),
		AzureFirewallPropertiesFormat: &networkmgmt.AzureFirewallPropertiesFormat{
			//ApplicationRuleCollections: nil,
			//NatRuleCollections:         nil,
			//NetworkRuleCollections:     nil,
			IPConfigurations:  setIPConfigurations(v.Spec.IPConfigurations),
			ProvisioningState: networkmgmt.ProvisioningState(v.Spec.ProvisioningState),
			ThreatIntelMode:   networkmgmt.AzureFirewallThreatIntelMode(v.Spec.ThreatIntelMode),
			VirtualHub:        setSubResource(v.Spec.VirtualHub),
			FirewallPolicy:    setSubResource(v.Spec.FirewallPolicy),
			HubIPAddresses:    setHubIpAddresses(v.Spec.HubIPAddresses),
		},
	}
}