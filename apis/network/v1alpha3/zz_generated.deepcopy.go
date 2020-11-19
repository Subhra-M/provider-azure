// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpace) DeepCopyInto(out *AddressSpace) {
	*out = *in
	if in.AddressPrefixes != nil {
		in, out := &in.AddressPrefixes, &out.AddressPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpace.
func (in *AddressSpace) DeepCopy() *AddressSpace {
	if in == nil {
		return nil
	}
	out := new(AddressSpace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSecurityGroup) DeepCopyInto(out *ApplicationSecurityGroup) {
	*out = *in
	out.Properties = in.Properties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSecurityGroup.
func (in *ApplicationSecurityGroup) DeepCopy() *ApplicationSecurityGroup {
	if in == nil {
		return nil
	}
	out := new(ApplicationSecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSecurityGroupPropertiesFormat) DeepCopyInto(out *ApplicationSecurityGroupPropertiesFormat) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSecurityGroupPropertiesFormat.
func (in *ApplicationSecurityGroupPropertiesFormat) DeepCopy() *ApplicationSecurityGroupPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(ApplicationSecurityGroupPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewall) DeepCopyInto(out *AzureFirewall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewall.
func (in *AzureFirewall) DeepCopy() *AzureFirewall {
	if in == nil {
		return nil
	}
	out := new(AzureFirewall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureFirewall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallIPConfiguration) DeepCopyInto(out *AzureFirewallIPConfiguration) {
	*out = *in
	in.AzureFirewallIPConfigurationPropertiesFormat.DeepCopyInto(&out.AzureFirewallIPConfigurationPropertiesFormat)
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallIPConfiguration.
func (in *AzureFirewallIPConfiguration) DeepCopy() *AzureFirewallIPConfiguration {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallIPConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallIPConfigurationPropertiesFormat) DeepCopyInto(out *AzureFirewallIPConfigurationPropertiesFormat) {
	*out = *in
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicIPAddress != nil {
		in, out := &in.PublicIPAddress, &out.PublicIPAddress
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallIPConfigurationPropertiesFormat.
func (in *AzureFirewallIPConfigurationPropertiesFormat) DeepCopy() *AzureFirewallIPConfigurationPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallIPConfigurationPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallList) DeepCopyInto(out *AzureFirewallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureFirewall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallList.
func (in *AzureFirewallList) DeepCopy() *AzureFirewallList {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureFirewallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallNatRule) DeepCopyInto(out *AzureFirewallNatRule) {
	*out = *in
	if in.SourceAddresses != nil {
		in, out := &in.SourceAddresses, &out.SourceAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationAddresses != nil {
		in, out := &in.DestinationAddresses, &out.DestinationAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationPorts != nil {
		in, out := &in.DestinationPorts, &out.DestinationPorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallNatRule.
func (in *AzureFirewallNatRule) DeepCopy() *AzureFirewallNatRule {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallNatRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallNatRuleCollection) DeepCopyInto(out *AzureFirewallNatRuleCollection) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallNatRuleCollection.
func (in *AzureFirewallNatRuleCollection) DeepCopy() *AzureFirewallNatRuleCollection {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallNatRuleCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallNatRuleCollectionProperties) DeepCopyInto(out *AzureFirewallNatRuleCollectionProperties) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AzureFirewallNatRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallNatRuleCollectionProperties.
func (in *AzureFirewallNatRuleCollectionProperties) DeepCopy() *AzureFirewallNatRuleCollectionProperties {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallNatRuleCollectionProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallNetworkRule) DeepCopyInto(out *AzureFirewallNetworkRule) {
	*out = *in
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceAddresses != nil {
		in, out := &in.SourceAddresses, &out.SourceAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationAddresses != nil {
		in, out := &in.DestinationAddresses, &out.DestinationAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationPorts != nil {
		in, out := &in.DestinationPorts, &out.DestinationPorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallNetworkRule.
func (in *AzureFirewallNetworkRule) DeepCopy() *AzureFirewallNetworkRule {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallNetworkRuleCollection) DeepCopyInto(out *AzureFirewallNetworkRuleCollection) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallNetworkRuleCollection.
func (in *AzureFirewallNetworkRuleCollection) DeepCopy() *AzureFirewallNetworkRuleCollection {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallNetworkRuleCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallNetworkRuleCollectionPropertiesFormat) DeepCopyInto(out *AzureFirewallNetworkRuleCollectionPropertiesFormat) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AzureFirewallNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallNetworkRuleCollectionPropertiesFormat.
func (in *AzureFirewallNetworkRuleCollectionPropertiesFormat) DeepCopy() *AzureFirewallNetworkRuleCollectionPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallNetworkRuleCollectionPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallPropertiesFormat) DeepCopyInto(out *AzureFirewallPropertiesFormat) {
	*out = *in
	if in.NatRuleCollections != nil {
		in, out := &in.NatRuleCollections, &out.NatRuleCollections
		*out = new([]AzureFirewallNatRuleCollection)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AzureFirewallNatRuleCollection, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.NetworkRuleCollections != nil {
		in, out := &in.NetworkRuleCollections, &out.NetworkRuleCollections
		*out = new([]AzureFirewallNetworkRuleCollection)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AzureFirewallNetworkRuleCollection, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.IPConfigurations != nil {
		in, out := &in.IPConfigurations, &out.IPConfigurations
		*out = new([]AzureFirewallIPConfiguration)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AzureFirewallIPConfiguration, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.VirtualHub != nil {
		in, out := &in.VirtualHub, &out.VirtualHub
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.FirewallPolicy != nil {
		in, out := &in.FirewallPolicy, &out.FirewallPolicy
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.HubIPAddresses != nil {
		in, out := &in.HubIPAddresses, &out.HubIPAddresses
		*out = new(HubIPAddresses)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallPropertiesFormat.
func (in *AzureFirewallPropertiesFormat) DeepCopy() *AzureFirewallPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallPublicIPAddress) DeepCopyInto(out *AzureFirewallPublicIPAddress) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallPublicIPAddress.
func (in *AzureFirewallPublicIPAddress) DeepCopy() *AzureFirewallPublicIPAddress {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallPublicIPAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallSpec) DeepCopyInto(out *AzureFirewallSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.AzureFirewallPropertiesFormat.DeepCopyInto(&out.AzureFirewallPropertiesFormat)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallSpec.
func (in *AzureFirewallSpec) DeepCopy() *AzureFirewallSpec {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFirewallStatus) DeepCopyInto(out *AzureFirewallStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFirewallStatus.
func (in *AzureFirewallStatus) DeepCopy() *AzureFirewallStatus {
	if in == nil {
		return nil
	}
	out := new(AzureFirewallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubIPAddresses) DeepCopyInto(out *HubIPAddresses) {
	*out = *in
	if in.PublicIPAddresses != nil {
		in, out := &in.PublicIPAddresses, &out.PublicIPAddresses
		*out = new([]AzureFirewallPublicIPAddress)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AzureFirewallPublicIPAddress, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubIPAddresses.
func (in *HubIPAddresses) DeepCopy() *HubIPAddresses {
	if in == nil {
		return nil
	}
	out := new(HubIPAddresses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupList) DeepCopyInto(out *SecurityGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupList.
func (in *SecurityGroupList) DeepCopy() *SecurityGroupList {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupPropertiesFormat) DeepCopyInto(out *SecurityGroupPropertiesFormat) {
	*out = *in
	if in.SecurityRules != nil {
		in, out := &in.SecurityRules, &out.SecurityRules
		*out = new([]SecurityRule)
		if **in != nil {
			in, out := *in, *out
			*out = make([]SecurityRule, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.DefaultSecurityRules != nil {
		in, out := &in.DefaultSecurityRules, &out.DefaultSecurityRules
		*out = new([]SecurityRule)
		if **in != nil {
			in, out := *in, *out
			*out = make([]SecurityRule, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ResourceGUID != nil {
		in, out := &in.ResourceGUID, &out.ResourceGUID
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupPropertiesFormat.
func (in *SecurityGroupPropertiesFormat) DeepCopy() *SecurityGroupPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupSpec) DeepCopyInto(out *SecurityGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.SecurityGroupPropertiesFormat.DeepCopyInto(&out.SecurityGroupPropertiesFormat)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupSpec.
func (in *SecurityGroupSpec) DeepCopy() *SecurityGroupSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupStatus) DeepCopyInto(out *SecurityGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupStatus.
func (in *SecurityGroupStatus) DeepCopy() *SecurityGroupStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityRule) DeepCopyInto(out *SecurityRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityRule.
func (in *SecurityRule) DeepCopy() *SecurityRule {
	if in == nil {
		return nil
	}
	out := new(SecurityRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityRulePropertiesFormat) DeepCopyInto(out *SecurityRulePropertiesFormat) {
	*out = *in
	if in.SourceAddressPrefixes != nil {
		in, out := &in.SourceAddressPrefixes, &out.SourceAddressPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceApplicationSecurityGroups != nil {
		in, out := &in.SourceApplicationSecurityGroups, &out.SourceApplicationSecurityGroups
		*out = make([]ApplicationSecurityGroup, len(*in))
		copy(*out, *in)
	}
	if in.DestinationAddressPrefixes != nil {
		in, out := &in.DestinationAddressPrefixes, &out.DestinationAddressPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationApplicationSecurityGroups != nil {
		in, out := &in.DestinationApplicationSecurityGroups, &out.DestinationApplicationSecurityGroups
		*out = make([]ApplicationSecurityGroup, len(*in))
		copy(*out, *in)
	}
	if in.SourcePortRanges != nil {
		in, out := &in.SourcePortRanges, &out.SourcePortRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationPortRanges != nil {
		in, out := &in.DestinationPortRanges, &out.DestinationPortRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityRulePropertiesFormat.
func (in *SecurityRulePropertiesFormat) DeepCopy() *SecurityRulePropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(SecurityRulePropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceEndpointPropertiesFormat) DeepCopyInto(out *ServiceEndpointPropertiesFormat) {
	*out = *in
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceEndpointPropertiesFormat.
func (in *ServiceEndpointPropertiesFormat) DeepCopy() *ServiceEndpointPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(ServiceEndpointPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource) DeepCopyInto(out *SubResource) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource.
func (in *SubResource) DeepCopy() *SubResource {
	if in == nil {
		return nil
	}
	out := new(SubResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetList) DeepCopyInto(out *SubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetList.
func (in *SubnetList) DeepCopy() *SubnetList {
	if in == nil {
		return nil
	}
	out := new(SubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetPropertiesFormat) DeepCopyInto(out *SubnetPropertiesFormat) {
	*out = *in
	if in.ServiceEndpoints != nil {
		in, out := &in.ServiceEndpoints, &out.ServiceEndpoints
		*out = make([]ServiceEndpointPropertiesFormat, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetPropertiesFormat.
func (in *SubnetPropertiesFormat) DeepCopy() *SubnetPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(SubnetPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.VirtualNetworkNameRef != nil {
		in, out := &in.VirtualNetworkNameRef, &out.VirtualNetworkNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.VirtualNetworkNameSelector != nil {
		in, out := &in.VirtualNetworkNameSelector, &out.VirtualNetworkNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.SubnetPropertiesFormat.DeepCopyInto(&out.SubnetPropertiesFormat)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetwork) DeepCopyInto(out *VirtualNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetwork.
func (in *VirtualNetwork) DeepCopy() *VirtualNetwork {
	if in == nil {
		return nil
	}
	out := new(VirtualNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkList) DeepCopyInto(out *VirtualNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkList.
func (in *VirtualNetworkList) DeepCopy() *VirtualNetworkList {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkPropertiesFormat) DeepCopyInto(out *VirtualNetworkPropertiesFormat) {
	*out = *in
	in.AddressSpace.DeepCopyInto(&out.AddressSpace)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkPropertiesFormat.
func (in *VirtualNetworkPropertiesFormat) DeepCopy() *VirtualNetworkPropertiesFormat {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkPropertiesFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkSpec) DeepCopyInto(out *VirtualNetworkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.VirtualNetworkPropertiesFormat.DeepCopyInto(&out.VirtualNetworkPropertiesFormat)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkSpec.
func (in *VirtualNetworkSpec) DeepCopy() *VirtualNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkStatus) DeepCopyInto(out *VirtualNetworkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkStatus.
func (in *VirtualNetworkStatus) DeepCopy() *VirtualNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkStatus)
	in.DeepCopyInto(out)
	return out
}
