/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// FlowSecurityPolicyIdentity represents the Identity of the object
var FlowSecurityPolicyIdentity = bambou.Identity{
	Name:     "flowsecuritypolicy",
	Category: "flowsecuritypolicies",
}

// FlowSecurityPoliciesList represents a list of FlowSecurityPolicies
type FlowSecurityPoliciesList []*FlowSecurityPolicy

// FlowSecurityPoliciesAncestor is the interface of an ancestor of a FlowSecurityPolicy must implement.
type FlowSecurityPoliciesAncestor interface {
	FlowSecurityPolicies(*bambou.FetchingInfo) (FlowSecurityPoliciesList, *bambou.Error)
	CreateFlowSecurityPolicies(*FlowSecurityPolicy) *bambou.Error
}

// FlowSecurityPolicy represents the model of a flowsecuritypolicy
type FlowSecurityPolicy struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	Action                         string `json:"action,omitempty"`
	DestinationAddressOverwrite    string `json:"destinationAddressOverwrite,omitempty"`
	FlowID                         string `json:"flowID,omitempty"`
	EntityScope                    string `json:"entityScope,omitempty"`
	SourceAddressOverwrite         string `json:"sourceAddressOverwrite,omitempty"`
	Priority                       int    `json:"priority,omitempty"`
	AssociatedApplicationServiceID string `json:"associatedApplicationServiceID,omitempty"`
	AssociatedNetworkObjectID      string `json:"associatedNetworkObjectID,omitempty"`
	AssociatedNetworkObjectType    string `json:"associatedNetworkObjectType,omitempty"`
	ExternalID                     string `json:"externalID,omitempty"`
}

// NewFlowSecurityPolicy returns a new *FlowSecurityPolicy
func NewFlowSecurityPolicy() *FlowSecurityPolicy {

	return &FlowSecurityPolicy{
		Action: "FORWARD",
	}
}

// Identity returns the Identity of the object.
func (o *FlowSecurityPolicy) Identity() bambou.Identity {

	return FlowSecurityPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowSecurityPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowSecurityPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the FlowSecurityPolicy from the server
func (o *FlowSecurityPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the FlowSecurityPolicy into the server
func (o *FlowSecurityPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the FlowSecurityPolicy from the server
func (o *FlowSecurityPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the FlowSecurityPolicy
func (o *FlowSecurityPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the FlowSecurityPolicy
func (o *FlowSecurityPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the FlowSecurityPolicy
func (o *FlowSecurityPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the FlowSecurityPolicy
func (o *FlowSecurityPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the FlowSecurityPolicy
func (o *FlowSecurityPolicy) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// CreateEventLog creates a new child EventLog under the FlowSecurityPolicy
func (o *FlowSecurityPolicy) CreateEventLog(child *EventLog) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
