// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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

package v20180317

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AssociateTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Association array
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations" list`
}

func (r *AssociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateTargetGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateTargetGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociationItem struct {

	// ID of associated CLB instance
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// ID of associated listener
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// ID of associated forwarding rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Protocol type of associated listener, such as HTTP or TCP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port of associated listener
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Domain name of associated forwarding rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// URL of associated forwarding rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
}

type AutoRewriteRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// HTTPS:443 listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name to be redirected under an HTTPS:443 listener
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`
}

func (r *AutoRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AutoRewriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AutoRewriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AutoRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AutoRewriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Backend struct {

	// Real server type. Value range: CVM, ENI (coming soon)
	Type *string `json:"Type,omitempty" name:"Type"`

	// Unique ID of a real server, which can be obtained from the unInstanceId field in the return of the DescribeInstances API
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Listening port of a real server
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Forwarding weight of a real server. Value range: [0, 100]. Default value: 10.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Public IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// Private IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// Real server instance names
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Bound time of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`

	// Unique ENI ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EniId *string `json:"EniId,omitempty" name:"EniId"`
}

type BasicTargetGroupInfo struct {

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Target group name
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`
}

type BatchDeregisterTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Unbound targets
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets" list`
}

func (r *BatchDeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeregisterTargetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IDs of the listeners failed to unbind
		FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeregisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeregisterTargetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of weights to be modified in batches
	ModifyList []*RsWeightRule `json:"ModifyList,omitempty" name:"ModifyList" list`
}

func (r *BatchModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchModifyTargetWeightRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchModifyTargetWeightResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Binding target
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets" list`
}

func (r *BatchRegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchRegisterTargetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchRegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IDs of the listeners failed to bind. If this is blank, all listeners are bound successfully.
	// Note: This field may return null, indicating that no valid values can be obtained.
		FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchRegisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchRegisterTargetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchTarget struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Binding port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// CVM instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ENI IP
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`

	// Weight of a CVM instance. Value range: [0, 100]. If it is not specified when binding the instance, 10 will be used by default.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Layer-7 rule ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type CertIdRelatedWithLoadBalancers struct {

	// Certificate ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// List of CLB instances associated with certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancers []*LoadBalancer `json:"LoadBalancers,omitempty" name:"LoadBalancers" list`
}

type CertificateInput struct {

	// Authentication type. Value range: UNIDIRECTIONAL (unidirectional authentication), MUTUAL (mutual authentication)
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`

	// ID of a server certificate. If you leave this parameter empty, you must upload the certificate, including CertContent, CertKey, and CertName.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// ID of a client certificate. When the listener adopts mutual authentication (i.e., SSLMode = mutual), if you leave this parameter empty, you must upload the client certificate, including CertCaContent and CertCaName.
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// Name of the uploaded server certificate. If there is no CertId, this parameter is required.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Key of the uploaded server certificate. If there is no CertId, this parameter is required.
	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`

	// Content of the uploaded server certificate. If there is no CertId, this parameter is required.
	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`

	// Name of the uploaded client CA certificate. When SSLMode = mutual, if there is no CertCaId, this parameter is required.
	CertCaName *string `json:"CertCaName,omitempty" name:"CertCaName"`

	// Content of the uploaded client certificate. When SSLMode = mutual, if there is no CertCaId, this parameter is required.
	CertCaContent *string `json:"CertCaContent,omitempty" name:"CertCaContent"`
}

type CertificateOutput struct {

	// Authentication type. Value range: UNIDIRECTIONAL (unidirectional authentication), MUTUAL (mutual authentication)
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`

	// Server certificate ID.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Client certificate ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`
}

type ClassicalHealth struct {

	// Private IP of a real server
	IP *string `json:"IP,omitempty" name:"IP"`

	// Real server port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// CLB listener port
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// Forwarding protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Health check result. Value range: 1 (healthy), 0 (unhealthy)
	HealthStatus *int64 `json:"HealthStatus,omitempty" name:"HealthStatus"`
}

type ClassicalListener struct {

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// CLB listener port
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// Backend forwarding port of a listener
	InstancePort *int64 `json:"InstancePort,omitempty" name:"InstancePort"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener protocol type
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Session persistence time
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// Whether health check is enabled. 1: enabled; 0: disabled.
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// Response timeout period
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// Check interval
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// Health threshold
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// Unhealthy threshold
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// A request balancing method for HTTP and HTTPS listeners of a public network classic CLB. wrr means weighted round robin, while ip_hash means consistent hashing based on source IPs of access requests.
	HttpHash *string `json:"HttpHash,omitempty" name:"HttpHash"`

	// Health check return code for HTTP and HTTPS listeners of a public network classic CLB. For more information, see the explanation of the field in the listener creating API.
	HttpCode *int64 `json:"HttpCode,omitempty" name:"HttpCode"`

	// Health check path for HTTP and HTTPS listeners of a public network classic CLB
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// Authentication method for an HTTPS listener of a public network classic CLB
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`

	// Server certificate ID for an HTTPS listener of a public network classic CLB
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Client certificate ID for an HTTPS listener of a public network classic CLB
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// Listener status. Value range: 0 (creating), 1 (running)
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ClassicalLoadBalancerInfo struct {

	// Real server ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of CLB instance IDs
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`
}

type ClassicalTarget struct {

	// Real server type. Value range: CVM, ENI (coming soon)
	Type *string `json:"Type,omitempty" name:"Type"`

	// Unique ID of a real server, which can be obtained from the unInstanceId field in the return of the DescribeInstances API
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Forwarding weight of a real server. Value range: [0, 100]. Default value: 10.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Public IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// Private IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// Real server instance names
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Real server status
	// 1: failed; 2: running; 3: creating; 4: shut down; 5: returned; 6: returning; 7: restarting; 8: starting; 9: shutting down; 10: resetting password; 11: formatting; 12: making image; 13: setting bandwidth; 14: reinstalling system; 19: upgrading; 21: hot migrating
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunFlag *int64 `json:"RunFlag,omitempty" name:"RunFlag"`
}

type ClassicalTargetInfo struct {

	// Real server ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Weight. Value range: [0, 100]
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type ClusterItem struct {

	// Unique cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster AZ, such as ap-guangzhou-1
	// Note: this field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Specifies for which ports to create listeners. Each port corresponds to a new listener
	Ports []*int64 `json:"Ports,omitempty" name:"Ports" list`

	// Listener protocol: TCP, UDP, HTTP, HTTPS, or TCP_SSL (which is currently in beta test. If you want to use it, please submit a ticket for application)
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// List of names of the listeners to be created. The array of names and array of ports are in one-to-one correspondence. If you do not want to name them now, you do not need to provide this parameter.
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames" list`

	// Health check parameter, which is applicable only to TCP/UDP/TCP_SSL listeners
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Certificate information. This parameter is applicable only to TCP_SSL listeners and HTTPS listeners with the SNI feature not enabled.
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// Session persistence time in seconds. Value range: 30-3,600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Forwarding method of a listener. Value range: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR. This parameter is applicable only to TCP/UDP/TCP_SSL listeners.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Whether to enable the SNI feature. This parameter is applicable only to HTTPS listeners
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`
}

func (r *CreateListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateListenerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of unique IDs of the created listeners
		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateListenerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CLB instance type. 1: generic CLB instance. Currently, only 1 can be passed in
	Forward *int64 `json:"Forward,omitempty" name:"Forward"`

	// CLB instance name, which takes effect only when an instance is created. Rule: 1-50 letters, digits, dashes (-), or underscores (_).
	// Note: If this name is the same as the name of an existing CLB instance in the system, the system will automatically generate a name for this newly created instance.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Network ID of the backend target server of CLB, which can be obtained through the DescribeVpcEx API. If this parameter is not passed in, it will default to a basic network ("0").
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// A subnet ID must be specified when you purchase a private network CLB instance in a VPC, and the VIP of this instance will be generated in this subnet.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the DescribeProject API. If this parameter is not passed in, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP version. Valid values: IPv4, IPv6, IPv6FullChain. Default value: IPv4. This parameter is applicable only to public network CLB instances.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// Number of CLBs to be created. Default value: 1.
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// Sets the primary AZ ID for cross-AZ disaster recovery, such as 100001 or ap-guangzhou-1, which is applicable only to public network CLB.
	// Note: A primary AZ carries traffic, while a secondary AZ does not carry traffic by default and will be used only if the primary AZ becomes unavailable. The platform will automatically select the optimal secondary AZ. The list of primary AZs in a specific region can be queried through the DescribeMasterZones API.
	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`

	// Specifies an AZ ID for creating a CLB instance, such as ap-guangzhou-1, which is applicable only to public network CLB.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB network billing mode. This parameter is applicable only to public network CLB instances.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// This parameter is applicable only to public network CLB instances. Valid values: CMCC (China Mobile), CTCC (China Telecom), CUCC (China Unicom). If this parameter is not specified, BGP will be used by default. ISPs supported in a region can be queried with the `DescribeSingleIsp` API. If an ISP is specified, only bill-by-bandwidth-package (BANDWIDTH_PACKAGE) can be used as the network billing mode.
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// Tags a CLB instance when purchasing it
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of unique CLB instance IDs.
		LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Information of the new forwarding rule
	Rules []*RuleInput `json:"Rules,omitempty" name:"Rules" list`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of unique IDs of created forwarding rules
		LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTargetGroupRequest struct {
	*tchttp.BaseRequest

	// Target group name (up to 50 characters)
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// `vpcid` attribute of a target group. If this parameter is left empty, the default VPC will be used.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Default port of a target group, which can be used for subsequently added servers.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Real server bound to a target group
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances" list`
}

func (r *CreateTargetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTargetGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTargetGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID generated after target group creation
		TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTargetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTargetGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// ID of the listener to be deleted
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteListenerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteListenerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// Array of IDs of the CLB instances to be deleted. Array length limit: 20
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRewriteRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Source listener ID
	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`

	// Target listener ID
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// Redirection relationship between forwarding rules
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos" list`
}

func (r *DeleteRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRewriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRewriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRewriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Array of IDs of the forwarding rules to be deleted
	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds" list`

	// Domain name of the forwarding rule to be deleted. This parameter does not take effect if LocationIds is specified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding path of the forwarding rule to be deleted. This parameter does not take effect if LocationIds is specified
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds" list`
}

func (r *DeleteTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTargetGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTargetGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Information of server to be unbound
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances" list`
}

func (r *DeregisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterTargetGroupInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeregisterTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterTargetGroupInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetsFromClassicalLBRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of real server instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DeregisterTargetsFromClassicalLBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterTargetsFromClassicalLBRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetsFromClassicalLBResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeregisterTargetsFromClassicalLBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterTargetsFromClassicalLBResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID in the format of lb-12345678
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Listener ID in the format of lbl-12345678
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers to be unbound. Array length limit: 20
	Targets []*Target `json:"Targets,omitempty" name:"Targets" list`

	// Forwarding rule ID in the format of loc-12345678. When unbinding a server from a layer-7 forwarding rule, you must provide either this parameter or Domain+Url
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *DeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterTargetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeregisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterTargetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBByInstanceIdRequest struct {
	*tchttp.BaseRequest

	// List of real server IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeClassicalLBByInstanceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBByInstanceIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBByInstanceIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CLB information list
		LoadBalancerInfoList []*ClassicalLoadBalancerInfo `json:"LoadBalancerInfoList,omitempty" name:"LoadBalancerInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBByInstanceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBByInstanceIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBHealthStatusRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DescribeClassicalLBHealthStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBHealthStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of real server health statuses
	// Note: This field may return null, indicating that no valid values can be obtained.
		HealthList []*ClassicalHealth `json:"HealthList,omitempty" name:"HealthList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBHealthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBHealthStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBListenersRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of CLB listener IDs
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds" list`

	// CLB listening protocol. Value range: TCP, UDP, HTTP, HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// CLB listening port. Value range: [1-65535]
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// Listener status. Value range: 0 (creating), 1 (running)
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeClassicalLBListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBListenersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of listeners
	// Note: This field may return null, indicating that no valid values can be obtained.
		Listeners []*ClassicalListener `json:"Listeners,omitempty" name:"Listeners" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBListenersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

func (r *DescribeClassicalLBTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBTargetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of real servers
	// Note: This field may return null, indicating that no valid values can be obtained.
		Targets []*ClassicalTarget `json:"Targets,omitempty" name:"Targets" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicalLBTargetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of IDs of the CLB listeners to be queried
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds" list`

	// Type of the listener protocol to be queried. Value range: TCP, UDP, HTTP, HTTPS, TCP_SSL
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port of the listener to be queried
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeListenersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of listeners
		Listeners []*Listener `json:"Listeners,omitempty" name:"Listeners" list`

		// Total number of listeners
	// Note: this field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeListenersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerListByCertIdRequest struct {
	*tchttp.BaseRequest

	// Server or client certificate ID
	CertIds []*string `json:"CertIds,omitempty" name:"CertIds" list`
}

func (r *DescribeLoadBalancerListByCertIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoadBalancerListByCertIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerListByCertIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID and list of CLB instances associated with it
		CertSet []*CertIdRelatedWithLoadBalancers `json:"CertSet,omitempty" name:"CertSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancerListByCertIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoadBalancerListByCertIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`

	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CLB instance type. 1: generic CLB instance; 0: classic CLB instance
	Forward *int64 `json:"Forward,omitempty" name:"Forward"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Domain name assigned to a CLB instance by Tencent Cloud. This parameter is meaningful only for the public network classic CLB.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// VIP address of a CLB instance (there can be multiple addresses)
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips" list`

	// Public IP of the real server bound to a CLB.
	BackendPublicIps []*string `json:"BackendPublicIps,omitempty" name:"BackendPublicIps" list`

	// Private IP of the real server bound to a CLB.
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps" list`

	// Data offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned CLB instances. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Sort by parameter. Value range: LoadBalancerName, CreateTime, Domain, LoadBalancerType.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 1: reverse; 0: sequential. Default value: reverse by creation time |
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// Search field which fuzzy matches name, domain name, or VIP.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the DescribeProject API.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Whether a CLB instance is bound to a real server. 0: no; 1: yes; -1: query all.
	WithRs *int64 `json:"WithRs,omitempty" name:"WithRs"`

	// VPC where a CLB instance resides, such as vpc-bhqkbhdx.
	// Basic network does not support queries by VpcId.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Security group ID, such as sg-m1cc9123
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// Master AZ, such as "100001" (Guangzhou Zone 1)
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of CLB instances that meet the filter criteria. This value is independent of the Limit in the input parameter.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of returned CLB instances.
		LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRewriteRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of CLB listener IDs
	SourceListenerIds []*string `json:"SourceListenerIds,omitempty" name:"SourceListenerIds" list`

	// Array of CLB forwarding rules
	SourceLocationIds []*string `json:"SourceLocationIds,omitempty" name:"SourceLocationIds" list`
}

func (r *DescribeRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRewriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRewriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of redirection forwarding rules. If there are no redirection rules, an empty array will be returned
		RewriteSet []*RuleOutput `json:"RewriteSet,omitempty" name:"RewriteSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRewriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// Filter. Currently, only filtering by `TargetGroupId`, `BindIP`, or `InstanceId` is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of displayed results. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Display offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetGroupInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of results in current query
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of the bound server
		TargetGroupInstanceSet []*TargetGroupBackend `json:"TargetGroupInstanceSet,omitempty" name:"TargetGroupInstanceSet" list`

		// Actual statistics, which are not affected by `Limit`, `Offset`, and `CAM`.
		RealCount *uint64 `json:"RealCount,omitempty" name:"RealCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetGroupInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupListRequest struct {
	*tchttp.BaseRequest

	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds" list`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: TargetGroupVpcId, TargetGroupName. Target group ID will be used first.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit of the number of displayed results. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTargetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of displayed results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information set of displayed target groups
		TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Target group ID, which is exclusive of `Filters`.
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds" list`

	// Limit of the number of displayed results. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: TargetGroupVpcId, TargetGroupName
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of displayed results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information set of displayed target groups
		TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest

	// List of IDs of CLB instances to be queried
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`
}

func (r *DescribeTargetHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetHealthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetHealthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of CLB instances
	// Note: This field may return null, indicating that no valid values can be obtained.
		LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitempty" name:"LoadBalancers" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetHealthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of listener IDs
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds" list`

	// Listener protocol type
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of real servers bound to the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
		Listeners []*ListenerBackend `json:"Listeners,omitempty" name:"Listeners" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTargetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// Request ID, i.e., the RequestId parameter returned by the API
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Current status of a task. Value range: 0 (succeeded), 1 (failed), 2 (in progress).
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Array of rules to be unbound
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations" list`
}

func (r *DisassociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateTargetGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateTargetGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExclusiveCluster struct {

	// Layer-4 dedicated cluster list
	// Note: this field may return null, indicating that no valid values can be obtained.
	L4Clusters []*ClusterItem `json:"L4Clusters,omitempty" name:"L4Clusters" list`

	// Layer-7 dedicated cluster list
	// Note: this field may return null, indicating that no valid values can be obtained.
	L7Clusters []*ClusterItem `json:"L7Clusters,omitempty" name:"L7Clusters" list`

	// vpcgw cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassicalCluster *ClusterItem `json:"ClassicalCluster,omitempty" name:"ClassicalCluster"`
}

type ExtraInfo struct {

	// Whether to enable VIP direct connection
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZhiTong *bool `json:"ZhiTong,omitempty" name:"ZhiTong"`

	// TgwGroup name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TgwGroupName *string `json:"TgwGroupName,omitempty" name:"TgwGroupName"`
}

type Filter struct {

	// Filter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter value array
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type HealthCheck struct {

	// Whether to enable health check. 1: enable; 0: disable.
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// Health check response timeout period in seconds (applicable only to layer-4 listeners). Value range: 2-60. Default value: 2. This parameter should be less than the check interval.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// Health check interval in seconds. Value range: 5-300. Default value: 5.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// Health threshold. Default value: 3, indicating that if a forward is found healthy three consecutive times, it is considered to be normal. Value range: 2-10
	// Note: This field may return null, indicating that no valid values can be obtained.
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// Unhealthy threshold. Default value: 3, indicating that if a forward is found unhealthy three consecutive times, it is considered to be exceptional. Value range: 2-10
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnHealthNum *int64 `json:"UnHealthNum,omitempty" name:"UnHealthNum"`

	// Health check status code (applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners). Value range: 1-31. Default value: 31.
	// 1 means that the return value of 1xx after detection means healthy, 2 for returning 2xx for healthy, 4 for returning 3xx for healthy, 8 for returning 4xx for healthy, and 16 for returning 5xx for healthy. If you want multiple return codes to represent healthy, sum up the corresponding values. Note: The HTTP health check mode of TCP listeners only supports specifying one kind of health check status code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpCode *int64 `json:"HttpCode,omitempty" name:"HttpCode"`

	// Health check path (applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners).
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// Health check domain name (applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners).
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// Health check method (applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners). Value range: HEAD, GET. Default value: HEAD.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" name:"HttpCheckMethod"`

	// Health check port (a custom check parameter), which is the port of the real server by default. Unless you want to specify a port, it is recommended to leave it empty. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckPort *int64 `json:"CheckPort,omitempty" name:"CheckPort"`

	// Health check protocol (a custom check parameter), which is required if the value of CheckType is CUSTOM. This parameter represents the input format of the health check. Value range: HEX, TEXT. If the value is HEX, the characters of SendContext and RecvContext can only be selected from 0123456789ABCDEF and the length must be an even number. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ContextType *string `json:"ContextType,omitempty" name:"ContextType"`

	// Health check protocol (a custom check parameter), which is required if the value of CheckType is CUSTOM. This parameter represents the content of the request sent by the health check. Only ASCII visible characters are allowed, and the maximum length is 500. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	SendContext *string `json:"SendContext,omitempty" name:"SendContext"`

	// Health check protocol (a custom check parameter), which is required if the value of CheckType is CUSTOM. This parameter represents the result returned by the health check. Only ASCII visible characters are allowed, and the maximum length is 500. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecvContext *string `json:"RecvContext,omitempty" name:"RecvContext"`

	// Health check protocol (a custom check parameter). Value range: TCP, HTTP, CUSTOM (applicable only to TCP/UDP listeners, where UDP listeners only support CUSTOM. If custom health check is used, this parameter is required).
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`

	// Health check protocol (a custom check parameter), which is required if the value of CheckType is HTTP. This parameter represents the HTTP version of the real server. Value range: HTTP/1.0, HTTP/1.1. (Applicable only to TCP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpVersion *string `json:"HttpVersion,omitempty" name:"HttpVersion"`
}

type InternetAccessible struct {

	// TRAFFIC_POSTPAID_BY_HOUR: hourly pay-as-you-go by traffic; BANDWIDTH_POSTPAID_BY_HOUR: hourly pay-as-you-go by bandwidth;
	// BANDWIDTH_PACKAGE: billed by bandwidth package (currently, this method is supported only if the ISP is specified)
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// Maximum outbound bandwidth in Mbps, which applies only to public network CLB. Value range: 0-65,535. Default value: 10.
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Bandwidth package type, such as SINGLEISP
	// Note: This field may return null, indicating that no valid values can be obtained.
	BandwidthpkgSubType *string `json:"BandwidthpkgSubType,omitempty" name:"BandwidthpkgSubType"`
}

type LBChargePrepaid struct {

	// Renewal type. AUTO_RENEW: automatic renewal; MANUAL_RENEW: manual renewal
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Cycle, indicating the number of months (reserved field)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type Listener struct {

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Information of certificates bound to the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
	Certificate *CertificateOutput `json:"Certificate,omitempty" name:"Certificate"`

	// Health check information of the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Request scheduling method
	// Note: This field may return null, indicating that no valid values can be obtained.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Session persistence time
	// Note: This field may return null, indicating that no valid values can be obtained.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Whether to enable the SNI feature (this parameter is only meaningful for HTTPS listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// All forwarding rules under a listener (this parameter is meaningful only for HTTP/HTTPS listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*RuleOutput `json:"Rules,omitempty" name:"Rules" list`

	// Listener name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// End port of a port range
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`

	// Real server type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// Basic information of a bound target group. This field will be returned when a target group is bound to a listener.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitempty" name:"TargetGroup"`
}

type ListenerBackend struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Information of rules under a listener (applicable only to HTTP/HTTPS listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*RuleTargets `json:"Rules,omitempty" name:"Rules" list`

	// List of real servers bound to a listener (applicable only to TCP/UDP/TCP_SSL listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets []*Backend `json:"Targets,omitempty" name:"Targets" list`

	// Ending port in port range if port range is supported; 0 if port range is not supported
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
}

type ListenerHealth struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener protocol
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// List of forwarding rules of the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*RuleHealth `json:"Rules,omitempty" name:"Rules" list`
}

type LoadBalancer struct {

	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CLB type identifier. Value range: 1 (CLB); 0 (classic CLB).
	Forward *uint64 `json:"Forward,omitempty" name:"Forward"`

	// CLB instance domain name. This field is provided only to public network classic CLB instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// List of VIPs of a CLB instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips" list`

	// CLB instance status, including:
	// 0: creating; 1: running.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// CLB instance creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last status change time of a CLB instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusTime *string `json:"StatusTime,omitempty" name:"StatusTime"`

	// ID of the project to which a CLB instance belongs. 0: default project.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Protective CLB identifier. Value range: 1 (protective), 0 (non-protective).
	// Note: This field may return null, indicating that no valid values can be obtained.
	OpenBgp *uint64 `json:"OpenBgp,omitempty" name:"OpenBgp"`

	// SNAT is enabled for all private network classic CLB created before December 2016.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Snat *bool `json:"Snat,omitempty" name:"Snat"`

	// 0: not isolated; 1: isolated.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Isolation *uint64 `json:"Isolation,omitempty" name:"Isolation"`

	// Log information. Only the public network CLB that have HTTP or HTTPS listeners can generate logs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Log *string `json:"Log,omitempty" name:"Log"`

	// Subnet where a CLB instance resides (meaningful only for private network VPC CLB)
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CLB instance tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// Security group of a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecureGroups []*string `json:"SecureGroups,omitempty" name:"SecureGroups" list`

	// Basic information of a backend server bound to a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitempty" name:"TargetRegionInfo"`

	// Anycast CLB publishing region. For non-anycast CLB, this field returns an empty string.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`

	// IP version. Valid values: ipv4, ipv6
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// VPC ID in a numeric form
	// Note: This field may return null, indicating that no valid values can be obtained.
	NumericalVpcId *uint64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`

	// ISP to which a CLB IP address belongs
	// Note: This field may return null, indicating that no valid values can be obtained.
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// Primary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterZone *ZoneInfo `json:"MasterZone,omitempty" name:"MasterZone"`

	// Secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitempty" name:"BackupZoneSet" list`

	// CLB instance isolation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// CLB instance expiration time, which takes effect only for prepaid instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// CLB instance billing mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// CLB instance network attributes
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitempty" name:"NetworkAttributes"`

	// Prepaid billing attributes of a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitempty" name:"PrepaidAttributes"`

	// Logset ID of CLB Log Service (CLS)
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// Log topic ID of CLB Log Service (CLS)
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// IPv6 address of a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddressIPv6 *string `json:"AddressIPv6,omitempty" name:"AddressIPv6"`

	// Reserved field which can be ignored generally.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitempty" name:"ExtraInfo"`

	// Whether an Anti-DDoS Pro instance can be bound
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDDos *bool `json:"IsDDos,omitempty" name:"IsDDos"`

	// Custom configuration ID at the CLB instance level
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// Whether a real server opens the traffic from a CLB instance to the internet
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// Private network dedicated cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// This field is meaningful only when the IP address version is `ipv6`. Valid values: IPv6Nat64, IPv6FullChain
	// Note: this field may return null, indicating that no valid values can be obtained.
	IPv6Mode *string `json:"IPv6Mode,omitempty" name:"IPv6Mode"`

	// Whether to enable SnatPro
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// SnatIp list after SnatPro load balancing is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps" list`

	// Performance guarantee specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// Whether VIP is blocked
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsBlock *bool `json:"IsBlock,omitempty" name:"IsBlock"`

	// 
	IsBlockTime *string `json:"IsBlockTime,omitempty" name:"IsBlockTime"`
}

type LoadBalancerHealth struct {

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// List of listeners
	// Note: This field may return null, indicating that no valid values can be obtained.
	Listeners []*ListenerHealth `json:"Listeners,omitempty" name:"Listeners" list`
}

type ManualRewriteRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Source listener ID
	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`

	// Target listener ID
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// Redirection relationship between forwarding rules
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos" list`
}

func (r *ManualRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManualRewriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManualRewriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManualRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManualRewriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainAttributesRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name, which must be under a created forwarding rule.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// New domain name
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`

	// Domain name certificate information. Note: This is only applicable to SNI-enabled listeners.
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// Whether to enable HTTP/2. Note: HTTP/2 can be enabled only for HTTPS domain names.
	Http2 *bool `json:"Http2,omitempty" name:"Http2"`

	// Whether to set this domain name as the default domain name. Note: Only one default domain name can be set under one listener.
	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`
}

func (r *ModifyDomainAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDomainAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDomainAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Legacy domain name under a listener.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// New domain name. 	Length limit: 1-120. There are three usage formats: non-regular expression, wildcard, and regular expression. A non-regular expression can only contain letters, digits, "-", and ".". In a wildcard, "*" can only be at the beginning or the end. A regular expressions must begin with a "~".
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`
}

func (r *ModifyDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// New listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Session persistence time in seconds. Value range: 30-3,600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Health check parameter, which is applicable only to TCP/UDP/TCP_SSL listeners.
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Certificate information. This parameter is applicable only to HTTPS/TCP_SSL listeners.
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// Forwarding method of a listener. Value range: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Whether to enable the SNI feature. This parameter is applicable only to HTTPS listeners. Note: The SNI feature can be enabled but cannot be disabled once enabled.
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyListenerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyListenerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest

	// Unique CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Region information of the real server bound to a CLB.
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitempty" name:"TargetRegionInfo"`

	// Network billing parameter
	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`

	// Whether the target opens traffic from CLB to the internet. If yes (true), only security groups on CLB will be verified; if no (false), security groups on both CLB and backend instance need to be verified.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// Whether to enable SnatPro
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`
}

func (r *ModifyLoadBalancerAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoadBalancerAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// This parameter can be used to query whether CLB billing mode switch is successful.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancerAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoadBalancerAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// ID of the forwarding rule to be modified.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// New forwarding path of the forwarding rule. This parameter is not required if the URL does not need to be modified
	Url *string `json:"Url,omitempty" name:"Url"`

	// Health check information
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Request forwarding method of the rule. Value range: WRR, LEAST_CONN, IP_HASH
	// They represent weighted round robin, least connections, and IP hash, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Session persistence time
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Forwarding protocol between CLB instance and real server. Value range: HTTP, HTTPS. Default value: HTTP
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// New name of target group
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// New default port of target group
	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyTargetGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupInstancesPortRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify port
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances" list`
}

func (r *ModifyTargetGroupInstancesPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetGroupInstancesPortRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupInstancesPortResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetGroupInstancesPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetGroupInstancesPortResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupInstancesWeightRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify weight
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances" list`
}

func (r *ModifyTargetGroupInstancesWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetGroupInstancesWeightRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupInstancesWeightResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetGroupInstancesWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetGroupInstancesWeightResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the ports
	Targets []*Target `json:"Targets,omitempty" name:"Targets" list`

	// New port of the real server bound to a listener or forwarding rule
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *ModifyTargetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetPortRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetPortResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetPortResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified
	Url *string `json:"Url,omitempty" name:"Url"`

	// List of real servers for which to modify the weight
	Targets []*Target `json:"Targets,omitempty" name:"Targets" list`

	// New forwarding weight of a real server. Value range: 0-100. Default value: 10. If the Targets.Weight parameter is set, this parameter will not take effect.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetWeightRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTargetWeightResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Server instance array
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances" list`
}

func (r *RegisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterTargetGroupInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterTargetGroupInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers to be bound. Array length limit: 20
	Targets []*Target `json:"Targets,omitempty" name:"Targets" list`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target forwarding rule domain name. This parameter does not take effect if LocationId is specified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target forwarding rule URL. This parameter does not take effect if LocationId is specified
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *RegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterTargetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterTargetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsWithClassicalLBRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Real server information
	Targets []*ClassicalTargetInfo `json:"Targets,omitempty" name:"Targets" list`
}

func (r *RegisterTargetsWithClassicalLBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterTargetsWithClassicalLBRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsWithClassicalLBResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterTargetsWithClassicalLBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterTargetsWithClassicalLBResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceCertForLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// ID of the certificate to be replaced, which can be a server certificate or a client certificate.
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// Information such as the content of the new certificate
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
}

func (r *ReplaceCertForLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceCertForLoadBalancersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceCertForLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceCertForLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceCertForLoadBalancersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RewriteLocationMap struct {

	// Source forwarding rule ID
	SourceLocationId *string `json:"SourceLocationId,omitempty" name:"SourceLocationId"`

	// Forwarding rule ID of a redirect target
	TargetLocationId *string `json:"TargetLocationId,omitempty" name:"TargetLocationId"`
}

type RewriteTarget struct {

	// Listener ID of a redirect target
	// Note: This field may return null, indicating that there is no redirection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// Forwarding rule ID of a redirect target
	// Note: This field may return null, indicating that there is no redirection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetLocationId *string `json:"TargetLocationId,omitempty" name:"TargetLocationId"`
}

type RsWeightRule struct {

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the weight
	Targets []*Target `json:"Targets,omitempty" name:"Targets" list`

	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified
	Url *string `json:"Url,omitempty" name:"Url"`

	// New forwarding weight of a real server. Value range: 0-100.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type RuleHealth struct {

	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding rule Url
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Health check status of the real server bound to this rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets []*TargetHealth `json:"Targets,omitempty" name:"Targets" list`
}

type RuleInput struct {

	// Domain name of the forwarding rule. Length: 1-80.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding rule path. Length: 1-200.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Session persistence time in seconds. Value range: 30-3,600. Setting it to 0 indicates that session persistence is disabled.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Health check information
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Certificate information
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// Request forwarding method of the rule. Value range: WRR, LEAST_CONN, IP_HASH
	// They represent weighted round robin, least connections, and IP hash, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Forwarding protocol between the CLB instance and real server. Currently, HTTP/HTTPS/TRPC are supported.
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`

	// Whether to set this domain name as the default domain name. Note: Only one default domain name can be set under one listener.
	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`

	// Whether to enable HTTP/2. Note: HTTP/2 can be enabled only for HTTPS domain names.
	Http2 *bool `json:"Http2,omitempty" name:"Http2"`

	// Target real server type. NODE: binding a general node; TARGETGROUP: binding a target group.
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// TRPC callee server route, which is required when `ForwardType` is `TRPC`.
	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`

	// TRPC calling service API, which is required when `ForwardType` is `TRPC`.
	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`
}

type RuleOutput struct {

	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding rule path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Session persistence time
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Health check information
	// Note: This field may return null, indicating that no valid values can be obtained.
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Certificate information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Certificate *CertificateOutput `json:"Certificate,omitempty" name:"Certificate"`

	// Request forwarding method of the rule
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// ID of the listener to which the forwarding rule belongs
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Redirect target information of a forwarding rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	RewriteTarget *RewriteTarget `json:"RewriteTarget,omitempty" name:"RewriteTarget"`

	// Whether to enable gzip
	HttpGzip *bool `json:"HttpGzip,omitempty" name:"HttpGzip"`

	// Whether the forwarding rule is automatically created
	BeAutoCreated *bool `json:"BeAutoCreated,omitempty" name:"BeAutoCreated"`

	// Whether to use as the default domain name
	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`

	// Whether to enable Http2
	Http2 *bool `json:"Http2,omitempty" name:"Http2"`

	// Forwarding protocol between CLB and real server
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`

	// Forwarding rule creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Real server type
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// Basic information of a bound target group. This field will be returned if a target group is bound to a rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitempty" name:"TargetGroup"`

	// WAF instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	WafDomainId *string `json:"WafDomainId,omitempty" name:"WafDomainId"`

	// TRPC callee server route, which is valid when `ForwardType` is `TRPC`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`

	// TRPC calling service API, which is valid when `ForwardType` is `TRPC`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`
}

type RuleTargets struct {

	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding rule path.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Real server information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets []*Backend `json:"Targets,omitempty" name:"Targets" list`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of security group IDs. One CLB instance can be bound to up to 50 security groups. If you want to unbind all security groups, you do not need to pass in this parameter, or you can pass in an empty array.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups" list`
}

func (r *SetLoadBalancerSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetLoadBalancerSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetLoadBalancerSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoadBalancerSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetLoadBalancerSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetSecurityGroupForLoadbalancersRequest struct {
	*tchttp.BaseRequest

	// Security group ID, such as sg-12345678
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// ADD: bind a security group;
	// DEL: unbind a security group
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// Array of CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`
}

func (r *SetSecurityGroupForLoadbalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetSecurityGroupForLoadbalancersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetSecurityGroupForLoadbalancersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSecurityGroupForLoadbalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetSecurityGroupForLoadbalancersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SnatIp struct {

	// Unique VPC subnet ID, such as `subnet-12345678`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// IP address, such as 192.168.0.1
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type TagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Target struct {

	// Listening port of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Real server type. Value range: CVM (Cloud Virtual Machine), ENI (Elastic Network Interface). This parameter does not take effect currently as an input parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Unique ID of a CVM instance, which needs to be passed in when binding a CVM instance and can be obtained from the InstanceId field in the return of the DescribeInstances API.
	// Note: Either InstanceId or EniIp must be passed in.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Forwarding weight of a real server. Value range: [0, 100]. Default value: 10.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// This parameter must be passed in when you bind an ENI, which represents the IP address of the ENI. The ENI has to be bound to a CVM instance first before it can be bound to a CLB instance. Note: Either InstanceId or EniIp must be passed in. To bind an ENI, you need to submit a ticket for application first.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`
}

type TargetGroupAssociation struct {

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type TargetGroupBackend struct {

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Real server type. Valid values: CVM, ENI (coming soon)
	Type *string `json:"Type,omitempty" name:"Type"`

	// Unique real server ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Listening port of real server
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Forwarding weight of real server. Value range: [0, 100]. Default value: 10.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// Public IP of real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// Private IP of real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// Real server instance name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Real server binding time
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`

	// Unique ENI ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	EniId *string `json:"EniId,omitempty" name:"EniId"`
}

type TargetGroupInfo struct {

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// `vpcid` of target group
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Target group name
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// Default port of target group
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Target group creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Target group modification time
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// Array of associated rules
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociatedRule []*AssociationItem `json:"AssociatedRule,omitempty" name:"AssociatedRule" list`
}

type TargetGroupInstance struct {

	// Private IP of target group instance
	BindIP *string `json:"BindIP,omitempty" name:"BindIP"`

	// Port of target group instance
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Weight of target group instance
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// New port of target group instance
	NewPort *uint64 `json:"NewPort,omitempty" name:"NewPort"`
}

type TargetHealth struct {

	// Private IP of the target
	IP *string `json:"IP,omitempty" name:"IP"`

	// Port bound to the target
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Current health status. true: healthy; false: unhealthy.
	HealthStatus *bool `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// Instance ID of the target, such as ins-12345678
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Detailed information of the current health status. Alive: healthy; Dead: exceptional; Unknown: check not started/checking/unknown status.
	HealthStatusDetial *string `json:"HealthStatusDetial,omitempty" name:"HealthStatusDetial"`
}

type TargetRegionInfo struct {

	// Region of the target, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Network of the target, which is in the format of vpc-abcd1234 for VPC or 0 for basic network
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ZoneInfo struct {

	// Unique AZ ID in a numeric form, such as 100001
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Unique AZ ID in a string form, such as ap-guangzhou-1
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ name, such as Guangzhou Zone 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}
