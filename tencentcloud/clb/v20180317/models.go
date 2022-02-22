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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AssociateTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Association array
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations"`
}

func (r *AssociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Associations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateTargetGroupsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// `HTTPS:443` listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// The domain name to be redirected under the listener `HTTPS:443`. If it is left empty, all domain names under the listener `HTTPS:443` will be configured with redirects.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Redirection status code. Valid values: 301, 302, and 307.
	RewriteCodes []*int64 `json:"RewriteCodes,omitempty" name:"RewriteCodes"`

	// Whether the matched URL is carried in redirection.
	TakeUrls []*bool `json:"TakeUrls,omitempty" name:"TakeUrls"`
}

func (r *AutoRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AutoRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Domains")
	delete(f, "RewriteCodes")
	delete(f, "TakeUrls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AutoRewriteRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AutoRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backend struct {

	// Real server type. Valid values: CVM, ENI.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Unique ID of a real server, which can be obtained from the unInstanceId field in the return of the DescribeInstances API
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Listening port of a real server
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Forwarding weight of a real server. Value range: [0, 100]. Default value: 10.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Public IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// Private IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

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

	// Unbinding targets
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchDeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeregisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IDs of the listeners failed to unbind
		FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeregisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of weights to be modified in batches
	ModifyList []*RsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
}

func (r *BatchModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ModifyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyTargetWeightRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Binding target
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchRegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRegisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchRegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IDs of the listeners failed to bind. If this is blank, all listeners are bound successfully.
	// Note: This field may return null, indicating that no valid values can be obtained.
		FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchRegisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchTarget struct {

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// The port to Bind
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// CVM instance ID. The primary IP of the primary ENI will be bound.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// It is required for binding an IP. It supports an ENI IP or any other private IP. To bind an ENI IP, the ENI should be bound to a CVM instance before being bound to a CLB instance.
	// Note: either `InstanceId` or `EniIp` must be passed in, which is required for binding a dual-stack IPv6 CVM instance.
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`

	// Weight of the CVM instance. Value range: [0, 100]. If it is not specified for binding the instance, 10 will be used by default.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Layer-7 rule ID.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type BindDetailItem struct {

	// Specifies the ID of CLB to be bound
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Specifies the ID of listener to be bound
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Specifies the domain name to be bound
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Sets the bound rule.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Listener name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener protocol.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// URL of the location.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Configuration ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
}

type BlockedIP struct {

	// Blacklisted IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// Blacklisted time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Expiration time
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type CertIdRelatedWithLoadBalancers struct {

	// Certificate ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// List of CLB instances associated with certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancers []*LoadBalancer `json:"LoadBalancers,omitempty" name:"LoadBalancers"`
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
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
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
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// Private IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

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

type CloneLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Clones the name of the CLB instance. The name must be 1-60 characters containing letters, numbers, "-" or "_".
	// Note: if the name of a new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Project ID of the CLB instance, which can be obtained through the [`DescribeProject`](https://intl.cloud.tencent.com/document/product/378/4400?from_cn_redirect=1) API. If this field is not specified, it will default to the default project.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Sets the primary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`, which is applicable only to public network CLB.
	// Note: A primary AZ carries traffic by default, while a secondary AZ does not. It only works when the primary AZ is faulty.
	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`

	// Sets the secondary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	// Note: A secondary AZ carries traffic when the primary AZ fails. 
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" name:"SlaveZoneId"`

	// Specifies an AZ ID for creating a CLB instance, such as `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB network billing mode. This parameter is applicable only to public network CLB instances.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// This parameter is applicable only to public network CLB instances. Valid values: CMCC (China Mobile), CTCC (China Telecom), CUCC (China Unicom). If this parameter is not specified, BGP will be used by default. ISPs supported in a region can be queried with the `DescribeSingleIsp` API. If an ISP is specified, only bill-by-bandwidth-package (BANDWIDTH_PACKAGE) can be used as the network billing mode.
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// Applies for CLB instances for a specified VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Tags a CLB instance when purchasing it
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Dedicated cluster information
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// Bandwidth package ID. If this parameter is specified, the network billing mode (`InternetAccessible.InternetChargeType`) will only support bill-by-bandwidth package (`BANDWIDTH_PACKAGE`).
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// Whether to support binding cross-VPC IPs or cross-region IPs
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// Creates `SnatIp` when the binding IPs of other VPCs feature is enabled
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// ID of the public network CLB dedicated cluster
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Guaranteed performance specification.
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// Tag of the STGW dedicated cluster
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// Availability zones for nearby access of private network CLB instances to distribute traffic
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Unique ID of an EIP, which can only be used when binding the EIP of a private network CLB instance (e.g., `eip-11112222`)
	EipAddressId *string `json:"EipAddressId,omitempty" name:"EipAddressId"`
}

func (r *CloneLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "ProjectId")
	delete(f, "MasterZoneId")
	delete(f, "SlaveZoneId")
	delete(f, "ZoneId")
	delete(f, "InternetAccessible")
	delete(f, "VipIsp")
	delete(f, "Vip")
	delete(f, "Tags")
	delete(f, "ExclusiveCluster")
	delete(f, "BandwidthPackageId")
	delete(f, "SnatPro")
	delete(f, "SnatIps")
	delete(f, "ClusterIds")
	delete(f, "SlaType")
	delete(f, "ClusterTag")
	delete(f, "Zones")
	delete(f, "EipAddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloneLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ConfigListItem struct {

	// Configuration ID.
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`

	// Configuration type.
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Configuration name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// Configuration content.
	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`

	// Creates configuration time.
	CreateTimestamp *string `json:"CreateTimestamp,omitempty" name:"CreateTimestamp"`

	// Modifies configuration time.
	UpdateTimestamp *string `json:"UpdateTimestamp,omitempty" name:"UpdateTimestamp"`
}

type CreateClsLogSetRequest struct {
	*tchttp.BaseRequest

	// Logset retention period in days; max value: 90
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Logset name, which must be unique among all CLS logsets; default value: clb_logset
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// Logset type. Valid values: ACCESS (access logs; default value) and HEALTH (health check logs).
	LogsetType *string `json:"LogsetType,omitempty" name:"LogsetType"`
}

func (r *CreateClsLogSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "LogsetName")
	delete(f, "LogsetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClsLogSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClsLogSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Logset ID.
		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClsLogSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Specifies for which ports to create listeners. Each port corresponds to a new listener.
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// Listener protocol: TCP, UDP, HTTP, HTTPS, or TCP_SSL (which is currently in beta test. If you want to use it, please submit a ticket for application).
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// List of names of the listeners to be created. The array of names and array of ports are in one-to-one correspondence. If you do not want to name them now, you do not need to provide this parameter.
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`

	// Health check parameter, which is applicable only to TCP, UDP, and TCP_SSL listeners.
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

	// Target real server type. `NODE`: binding a general node; `TARGETGROUP`: binding a target group.
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// Session persistence type. Valid values: Normal: the default session persistence type; QUIC_CID: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners. If this field is not specified, the default session persistence type will be used.
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// Whether to enable a persistent connection. This parameter is applicable only to HTTP and HTTPS listeners. Valid values: 0 (disable; default value) and 1 (enable).
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// This parameter is used to specify the end port and is required when creating a port range listener. Only one member can be passed in when inputting the `Ports` parameter, which is used to specify the start port. If you want to try the port range feature, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
	EndPort *uint64 `json:"EndPort,omitempty" name:"EndPort"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`
}

func (r *CreateListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Ports")
	delete(f, "Protocol")
	delete(f, "ListenerNames")
	delete(f, "HealthCheck")
	delete(f, "Certificate")
	delete(f, "SessionExpireTime")
	delete(f, "Scheduler")
	delete(f, "SniSwitch")
	delete(f, "TargetType")
	delete(f, "SessionType")
	delete(f, "KeepaliveEnable")
	delete(f, "EndPort")
	delete(f, "DeregisterTargetRst")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of unique IDs of the created listeners
		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CLB instance type. Valid value: 1 (generic CLB instance).
	Forward *int64 `json:"Forward,omitempty" name:"Forward"`

	// CLB instance name, which takes effect only when only one instance is to be created in the request. It can consist 1 to 60 letters, digits, hyphens (-), or underscores (_).
	// Note: if the name of the new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// Network ID of the target CLB real server, such as `vpc-12345678`, which can be obtained through the [DescribeVpcEx](https://intl.cloud.tencent.com/document/product/215/1372?from_cn_redirect=1) API. If this parameter is not specified, it will default to `DefaultVPC`. This parameter is required for creating a CLB instance.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// A subnet ID must be specified when you purchase a private network CLB instance in a VPC, and the VIP of this instance will be generated in this subnet. This parameter is required for creating a CLB instance.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Project ID of the CLB instance, which can be obtained through the [DescribeProject](https://intl.cloud.tencent.com/document/product/378/4400?from_cn_redirect=1) API. If this parameter is not specified, it will default to the default project.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP version. Valid values: `IPV4` (default), `IPV6` (IPV6 NAT64 version) or `IPv6FullChain` (IPv6 version). This parameter is only for public network CLB instances.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// Number of CLBs to be created. Default value: 1.
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// Sets the primary AZ ID for cross-AZ disaster recovery, such as 100001 or ap-guangzhou-1, which is applicable only to public network CLB.
	// Note: A primary AZ carries traffic, while a secondary AZ does not carry traffic by default and will be used only if the primary AZ becomes unavailable. The platform will automatically select the optimal secondary AZ. The list of primary AZs in a specific region can be queried through the DescribeMasterZones API.
	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`

	// Specifies an AZ ID for creating a CLB instance, such as `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB network billing mode. This parameter is applicable only to public network CLB instances.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// This parameter is applicable only to public network CLB instances. Valid values: CMCC (China Mobile), CTCC (China Telecom), CUCC (China Unicom). If this parameter is not specified, BGP will be used by default. ISPs supported in a region can be queried with the `DescribeSingleIsp` API. If an ISP is specified, only bill-by-bandwidth-package (BANDWIDTH_PACKAGE) can be used as the network billing mode.
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// Tags a CLB instance when purchasing it.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Specifies a VIP for the CLB instance.
	// <ul><li>`VpcId` is optional for creating shared clusters of public network CLB instances. For IPv6 CLB instance type, `SubnetId` is required; for IPv4 and IPv6 NAT64 types, it can be left empty.</li>
	// <li>`VpcId` is optional for creating shared clusters of public network CLB instances. For IPv6 CLB instance type, `SubnetId` is required; for IPv4 and IPv6 NAT64 types, it can be left empty.
	// </li></ul>
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Bandwidth package ID. If this parameter is specified, the network billing mode (`InternetAccessible.InternetChargeType`) will only support bill-by-bandwidth package (`BANDWIDTH_PACKAGE`).
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// Exclusive cluster information. This parameter is required for creating exclusive clusters of CLB instances.
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// Creates an LCU-supported CLB instance
	// <ul><li>To create an LCU-supported CLB, this field is required and the value is `SLA`. LCU-supports CLBs adopt the pay-as-you-go model and their performance is guaranteed.</li>
	// <li>It’s not required for a shared CLB instance.</li></ul>
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Whether Binding IPs of other VPCs feature switch
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// Creates `SnatIp` when the binding IPs of other VPCs feature is enabled
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// Tag for the STGW exclusive cluster.
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// Sets the secondary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	// Note: A secondary AZ will load traffic if the primary AZ has failures. The API `DescribeMasterZones` is used to query the primary and secondary AZ list of a region.
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" name:"SlaveZoneId"`

	// Unique ID of an EIP, which can only be used when binding the EIP of a private network CLB instance. E.g., `eip-11112222`.
	EipAddressId *string `json:"EipAddressId,omitempty" name:"EipAddressId"`

	// Whether to allow CLB traffic to the target group. `true`: allows CLB traffic to the target group and verifies security groups only on CLB; `false`: denies CLB traffic to the target group and verifies security groups on both CLB and backend instances.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerType")
	delete(f, "Forward")
	delete(f, "LoadBalancerName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "AddressIPVersion")
	delete(f, "Number")
	delete(f, "MasterZoneId")
	delete(f, "ZoneId")
	delete(f, "InternetAccessible")
	delete(f, "VipIsp")
	delete(f, "Tags")
	delete(f, "Vip")
	delete(f, "BandwidthPackageId")
	delete(f, "ExclusiveCluster")
	delete(f, "SlaType")
	delete(f, "ClientToken")
	delete(f, "SnatPro")
	delete(f, "SnatIps")
	delete(f, "ClusterTag")
	delete(f, "SlaveZoneId")
	delete(f, "EipAddressId")
	delete(f, "LoadBalancerPassToTarget")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of unique CLB instance IDs.
		LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

		// Order ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a CLB instance, e.g., lb-12345678.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Information of the SNAT IP to be added. You can specify a SNAT IP or use the one automatically assigned by a subnet.
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// Number of SNAT IPs to be added. This parameter is used in conjunction with `SnatIps`. Note that if `Ip` is specified in `SnapIps`, this parameter is not available. It defaults to `1` and the upper limit is `10`.
	Number *uint64 `json:"Number,omitempty" name:"Number"`
}

func (r *CreateLoadBalancerSnatIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerSnatIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SnatIps")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerSnatIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerSnatIpsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoadBalancerSnatIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerSnatIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Information of the new forwarding rule
	Rules []*RuleInput `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of unique IDs of created forwarding rules
		LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

func (r *CreateTargetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupName")
	delete(f, "VpcId")
	delete(f, "Port")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetGroupRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// The number of topic partitions, which changes as partitions are split or merged. Each log topic can have up to 50 partitions. If this parameter is not passed in, 1 partition will be created by default and up to 10 partitions are allowed to be created.
	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// Log type. Valid values: ACCESS (access logs; default value) and HEALTH (health check logs).
	TopicType *string `json:"TopicType,omitempty" name:"TopicType"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "PartitionCount")
	delete(f, "TopicType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log topic ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossTargets struct {

	// VPC ID of the CLB instance
	LocalVpcId *string `json:"LocalVpcId,omitempty" name:"LocalVpcId"`

	// VPC ID of the CVM or ENI instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// IP address of the CVM or ENI instance
	IP *string `json:"IP,omitempty" name:"IP"`

	// VPC name of the CVM or ENI instance
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// ENI ID of the CVM instance
	EniId *string `json:"EniId,omitempty" name:"EniId"`

	// ID of the CVM instance
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Name of the CVM instance
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Region of the CVM or ENI instance
	Region *string `json:"Region,omitempty" name:"Region"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenerRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerListenersRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of listener IDs to delete (20 IDs at most). If this parameter is left empty, all listeners of the CLB instance will be deleted.
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DeleteLoadBalancerListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoadBalancerListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// Array of IDs of the CLB instances to be deleted. Array length limit: 20.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a CLB instance, e.g., lb-12345678.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of the SNAT IP addresses to be deleted
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

func (r *DeleteLoadBalancerSnatIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerSnatIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Ips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerSnatIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerSnatIpsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoadBalancerSnatIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerSnatIpsResponse) FromJsonString(s string) error {
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
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
}

func (r *DeleteRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SourceListenerId")
	delete(f, "TargetListenerId")
	delete(f, "RewriteInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRewriteRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`

	// Domain name of the forwarding rule to be deleted. This parameter does not take effect if LocationIds is specified.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding path of the forwarding rule to be deleted. This parameter does not take effect if LocationIds is specified.
	Url *string `json:"Url,omitempty" name:"Url"`

	// A listener must be configured with a default domain name. If you need to delete the default domain name, you can specify another one as the new default domain name.
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "LocationIds")
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "NewDefaultServerDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`
}

func (r *DeleteTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTargetGroupsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Information of server to be unbound
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

func (r *DeregisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterTargetGroupInstancesRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetsFromClassicalLBRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of real server IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DeregisterTargetsFromClassicalLBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsFromClassicalLBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterTargetsFromClassicalLBRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsFromClassicalLBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID in the format of "lb-12345678"
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Listener ID in the format of "lbl-12345678"
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers to be unbound. Array length limit: 20.
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// Forwarding rule ID in the format of "loc-12345678". When unbinding a server from a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *DeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterTargetsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPListRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Data offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of IPs to be returned. Default value: 100,000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBlockIPListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockIPListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned IPs
		BlockedIPCount *uint64 `json:"BlockedIPCount,omitempty" name:"BlockedIPCount"`

		// Field for getting real client IP
		ClientIPField *string `json:"ClientIPField,omitempty" name:"ClientIPField"`

		// List of IPs added to blocklist 12360
		BlockedIPList []*BlockedIP `json:"BlockedIPList,omitempty" name:"BlockedIPList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIPListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPTaskRequest struct {
	*tchttp.BaseRequest

	// Async task ID returned by the `ModifyBlockIPList` API
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeBlockIPTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockIPTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1: running; 2: failed; 6: succeeded
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIPTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBByInstanceIdRequest struct {
	*tchttp.BaseRequest

	// List of real server IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeClassicalLBByInstanceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBByInstanceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBByInstanceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBByInstanceIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CLB information list
		LoadBalancerInfoList []*ClassicalLoadBalancerInfo `json:"LoadBalancerInfoList,omitempty" name:"LoadBalancerInfoList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBByInstanceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBHealthStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBHealthStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of real server health statuses
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		HealthList []*ClassicalHealth `json:"HealthList,omitempty" name:"HealthList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBHealthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBListenersRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of CLB listener IDs
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// CLB listening protocol. Valid values: TCP, UDP, HTTP, and HTTPS.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// CLB listening port. Value range: 1 - 65535.
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// Listener status. Valid values: 0 (creating) and 1 (running).
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeClassicalLBListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "ListenerPort")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Listener list
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		Listeners []*ClassicalListener `json:"Listeners,omitempty" name:"Listeners"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Real server list
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		Targets []*ClassicalTarget `json:"Targets,omitempty" name:"Targets"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogSetRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClsLogSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClsLogSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClsLogSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Logset ID
		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

		// Health check logset ID
		HealthLogsetId *string `json:"HealthLogsetId,omitempty" name:"HealthLogsetId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsLogSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClsLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossTargetsRequest struct {
	*tchttp.BaseRequest

	// Number of real server lists returned. Default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Starting offset of the real server list returned. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter conditions to query CVMs and ENIs
	// <li> `vpc-id` - String - Required: No - (Filter condition) Filter by VPC ID, such as "vpc-12345678".</li>
	// <li> `ip` - String - Required: No - (Filter condition) Filter by real server IP, such as "192.168.0.1".</li>
	// <li> `listener-id` - String - Required: No - (Filter condition) Filter by listener ID, such as "lbl-12345678".</li>
	// <li> `location-id` - String - Required: No - (Filter condition) Filter by forwarding rule ID of the layer-7 listener, such as "loc-12345678".</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCrossTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of real server lists
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Real server list
		CrossTargetSet []*CrossTargets `json:"CrossTargetSet,omitempty" name:"CrossTargetSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCrossTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomizedConfigAssociateListRequest struct {
	*tchttp.BaseRequest

	// Configuration ID.
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`

	// Start position of the binding list. Default: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of binding lists to pull. Default: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Searches for the domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeCustomizedConfigAssociateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigAssociateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UconfigId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizedConfigAssociateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomizedConfigAssociateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of bound resources
		BindList []*BindDetailItem `json:"BindList,omitempty" name:"BindList"`

		// Total number of bound resources
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomizedConfigAssociateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigAssociateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomizedConfigListRequest struct {
	*tchttp.BaseRequest

	// Configuration type. Valid values: `CLB` (CLB-specific configs), `SERVER` (domain name-specific configs), and `LOCATION` (forwarding rule-specific configs).
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Pagination offset. Default: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results per page. Default: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies the name of configs to query. Fuzzy match is supported.
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// Configuration ID.
	UconfigIds []*string `json:"UconfigIds,omitempty" name:"UconfigIds"`

	// The filters are:
	// <li> loadbalancer-id - String - Required: no - (filter) CLB instance ID, such as "lb-12345678". </li>
	// <li> vip - String - Required: no - (filter) CLB instance VIP, such as "1.1.1.1" and "2204::22:3". </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCustomizedConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConfigName")
	delete(f, "UconfigIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizedConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomizedConfigListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration list.
		ConfigList []*ConfigListItem `json:"ConfigList,omitempty" name:"ConfigList"`

		// Number of configurations.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomizedConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLBListenersRequest struct {
	*tchttp.BaseRequest

	// List of private network IPs to be queried.
	Backends []*LbRsItem `json:"Backends,omitempty" name:"Backends"`
}

func (r *DescribeLBListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLBListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Backends")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLBListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLBListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Listener rule associated with the real server.
		LoadBalancers []*LBItem `json:"LoadBalancers,omitempty" name:"LoadBalancers"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLBListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLBListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of CLB listener IDs to query (100 IDs at most).
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// Type of the listener protocols to be queried. Valid values: TCP, UDP, HTTP, HTTPS, and TCP_SSL.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port of the listeners to be queried
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Listener list
		Listeners []*Listener `json:"Listeners,omitempty" name:"Listeners"`

		// Total number of listeners (with filters of port, protocol, and listener ID applied).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerListByCertIdRequest struct {
	*tchttp.BaseRequest

	// Server or client certificate ID
	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
}

func (r *DescribeLoadBalancerListByCertIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerListByCertIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerListByCertIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerListByCertIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID and list of CLB instances associated with it
		CertSet []*CertIdRelatedWithLoadBalancers `json:"CertSet,omitempty" name:"CertSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancerListByCertIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerListByCertIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLoadBalancerOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of CLB instances
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Number of CLB instances that are running
		RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`

		// Number of CLB instances that are isolated
		IsolationCount *int64 `json:"IsolationCount,omitempty" name:"IsolationCount"`

		// Number of CLB instances that are about to expire
		WillExpireCount *int64 `json:"WillExpireCount,omitempty" name:"WillExpireCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancerOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerTrafficRequest struct {
	*tchttp.BaseRequest

	// CLB instance region. If this parameter is not passed in, CLB instances in all regions will be returned.
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
}

func (r *DescribeLoadBalancerTrafficRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerTrafficRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerTrafficRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerTrafficResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of CLB instances sorted by outbound bandwidth from highest to lowest
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		LoadBalancerTraffic []*LoadBalancerTraffic `json:"LoadBalancerTraffic,omitempty" name:"LoadBalancerTraffic"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancerTrafficResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerTrafficResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersDetailRequest struct {
	*tchttp.BaseRequest

	// Number of CLB instance lists returned. Default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Starting offset of the CLB instance list returned. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// List of fields. Only fields specified will be returned. If it’s left blank, `null` is returned. The fields `LoadBalancerId` and `LoadBalancerName` are added by default. For details about fields, see <a href="https://intl.cloud.tencent.com/document/api/214/30694?from_cn_redirect=1#LoadBalancerDetail">LoadBalancerDetail</a>.
	Fields []*string `json:"Fields,omitempty" name:"Fields"`

	// Target type. Valid values: NODE and GROUP. If the list of fields contains `TargetId`, `TargetAddress`, `TargetPort`, `TargetWeight` and other fields, `Target` of the target group or non-target group must be exported.
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// Filter condition of querying lists describing CLB instance details:
	// <li> loadbalancer-id - String - Required: no - (Filter condition) CLB instance ID, such as "lb-12345678". </li>
	// <li> project-id - String - Required: no - (Filter condition) Project ID, such as "0" and "123".</li>
	// <li> network - String - Required: no - (Filter condition) Network type of the CLB instance, such as "Public" and "Private".</li>
	// <li> vip - String - Required: no - (Filter condition) CLB instance VIP, such as "1.1.1.1" and "2204::22:3". </li>
	// <li> target-ip - String - Required: no - (Filter condition) Private IP of the target real servers, such as"1.1.1.1" and "2203::214:4".</li>
	// <li> vpcid - String - Required: no - (Filter condition) Identifier of the VPC instance to which the CLB instance belongs, such as "vpc-12345678".</li>
	// <li> zone - String - Required: no - (Filter condition) Availability zone where the CLB instance resides, such as "ap-guangzhou-1".</li>
	// <li> tag-key - String - Required: no - (Filter condition) Tag key of the CLB instance, such as "name".</li>
	// <li> tag:* - String - Required: no - (Filter condition) CLB instance tag, followed by tag key after the colon ':'. For example, use {"Name": "tag:name","Values": ["zhangsan", "lisi"]} to filter the tag key “name” with the tag value “zhangsan” and “lisi”.</li>
	// <li> fuzzy-search - String - Required: no - (Filter condition) Fuzzy search for CLB instance VIP and CLB instance name, such as "1.1".</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLoadBalancersDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Fields")
	delete(f, "TargetType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of lists describing CLB instance details.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of CLB instance details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		LoadBalancerDetailSet []*LoadBalancerDetail `json:"LoadBalancerDetailSet,omitempty" name:"LoadBalancerDetailSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancersDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

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
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// Public IP of the real server bound to a CLB.
	BackendPublicIps []*string `json:"BackendPublicIps,omitempty" name:"BackendPublicIps"`

	// Private IP of the real server bound to a CLB.
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`

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

	// Security group ID, e.g., `sg-m1cc****`.
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// Primary AZ ID, e.g., `100001` (Guangzhou Zone 1).
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// Each request can have up to 10 `Filters` and 100 `Filter.Values`. Detailed filter conditions:
	// <li> internet-charge-type - Type: String - Required: No - Filter by CLB network billing mode, including `TRAFFIC_POSTPAID_BY_HOUR`</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "LoadBalancerType")
	delete(f, "Forward")
	delete(f, "LoadBalancerName")
	delete(f, "Domain")
	delete(f, "LoadBalancerVips")
	delete(f, "BackendPublicIps")
	delete(f, "BackendPrivateIps")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "SearchKey")
	delete(f, "ProjectId")
	delete(f, "WithRs")
	delete(f, "VpcId")
	delete(f, "SecurityGroup")
	delete(f, "MasterZone")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of CLB instances that meet the filter criteria. This value is independent of the Limit in the input parameter.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of returned CLB instances.
		LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Quota list
		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRewriteRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of CLB listener IDs
	SourceListenerIds []*string `json:"SourceListenerIds,omitempty" name:"SourceListenerIds"`

	// Array of CLB forwarding rule IDs
	SourceLocationIds []*string `json:"SourceLocationIds,omitempty" name:"SourceLocationIds"`
}

func (r *DescribeRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SourceListenerIds")
	delete(f, "SourceLocationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRewriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of redirection forwarding rules. If there are no redirection rules, an empty array will be returned.
		RewriteSet []*RuleOutput `json:"RewriteSet,omitempty" name:"RewriteSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// Filter. Currently, only filtering by `TargetGroupId`, `BindIP`, or `InstanceId` is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Number of displayed results. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Display offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of results returned for the current query
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of the bound server
		TargetGroupInstanceSet []*TargetGroupBackend `json:"TargetGroupInstanceSet,omitempty" name:"TargetGroupInstanceSet"`

		// The actual total number of bound instances, which is not affected by the setting of `Limit`, `Offset` and the CAM permissions.
		RealCount *uint64 `json:"RealCount,omitempty" name:"RealCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupListRequest struct {
	*tchttp.BaseRequest

	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: `TargetGroupVpcId` and `TargetGroupName`. Target group ID will be used first.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit of the number of displayed results. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTargetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of displayed results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information set of displayed target groups
		TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Target group ID, which is exclusive of `Filters`.
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`

	// Limit of the number of displayed results. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: `TargetGroupVpcId` and `TargetGroupName`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of displayed results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information set of displayed target groups
		TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest

	// List of IDs of CLB instances to be queried
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DescribeTargetHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetHealthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CLB instance list
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitempty" name:"LoadBalancers"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// List of listener IDs (20 IDs at most).
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// Listener protocol type
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of real servers bound to the listener
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		Listeners []*ListenerBackend `json:"Listeners,omitempty" name:"Listeners"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// Request ID, i.e., the RequestId parameter returned by the API.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Order ID.
	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "DealName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Current status of a task. Value range: 0 (succeeded), 1 (failed), 2 (in progress).
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// Array of unique CLB instance IDs.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// Array of rules to be unbound
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations"`
}

func (r *DisassociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Associations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateTargetGroupsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExclusiveCluster struct {

	// Layer-4 dedicated cluster list
	// Note: this field may return null, indicating that no valid values can be obtained.
	L4Clusters []*ClusterItem `json:"L4Clusters,omitempty" name:"L4Clusters"`

	// Layer-7 dedicated cluster list
	// Note: this field may return null, indicating that no valid values can be obtained.
	L7Clusters []*ClusterItem `json:"L7Clusters,omitempty" name:"L7Clusters"`

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
	Values []*string `json:"Values,omitempty" name:"Values"`
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

	// Specifies the source IP for health check. `0`: use the CLB VIP as the source IP; `1`: IP range starting with 100.64 serving as the source IP. Default: `0`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SourceIpType *int64 `json:"SourceIpType,omitempty" name:"SourceIpType"`
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

type LBItem struct {

	// String ID of the CLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// VIP of the CLB instance.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Listener rule.
	Listeners []*ListenerItem `json:"Listeners,omitempty" name:"Listeners"`

	// Region of the CLB instance
	Region *string `json:"Region,omitempty" name:"Region"`
}

type LbRsItem struct {

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Private network IP to be queried, which can be of the CVM or ENI.
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
}

type LbRsTargets struct {

	// Private network IP type, which can be `cvm` or `eni`.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Private network IP of the real server.
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// Port bound to the real server.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// VPC ID of the real server.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Weight of the real server.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
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
	Rules []*RuleOutput `json:"Rules,omitempty" name:"Rules"`

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

	// Session persistence type. Valid values: Normal: the default session persistence type; QUIC_CID: session persistence by QUIC connection ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// Whether a persistent connection is enabled (1: enabled; 0: disabled). This parameter can only be configured in HTTP/HTTPS listeners.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// Only the NAT64 CLB TCP listeners are supported.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Toa *bool `json:"Toa,omitempty" name:"Toa"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`
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
	Rules []*RuleTargets `json:"Rules,omitempty" name:"Rules"`

	// List of real servers bound to a listener (applicable only to TCP/UDP/TCP_SSL listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`

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
	Rules []*RuleHealth `json:"Rules,omitempty" name:"Rules"`
}

type ListenerItem struct {

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener protocol.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Bound rule.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Rules []*RulesItems `json:"Rules,omitempty" name:"Rules"`

	// Object bound to the layer-4 listener.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Targets []*LbRsTargets `json:"Targets,omitempty" name:"Targets"`

	// End port of the listener.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
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
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

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
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// Security group of a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecureGroups []*string `json:"SecureGroups,omitempty" name:"SecureGroups"`

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
	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitempty" name:"BackupZoneSet"`

	// CLB instance isolation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// CLB instance expiration time, which takes effect only for prepaid instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Billing mode of CLB instance. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay as you go).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
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

	// Whether to enable SnatPro.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// `SnatIp` list after SnatPro load balancing is enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// Performance guarantee specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// Whether VIP is blocked
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsBlock *bool `json:"IsBlock,omitempty" name:"IsBlock"`

	// Time blocked or unblocked
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsBlockTime *string `json:"IsBlockTime,omitempty" name:"IsBlockTime"`

	// Whether the IP type is the local BGP
	LocalBgp *bool `json:"LocalBgp,omitempty" name:"LocalBgp"`

	// Dedicated layer-7 tag.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// If the layer-7 listener of an IPv6FullChain CLB instance is enabled, the CLB instance can be bound with an IPv4 and an IPv6 CVM instance simultaneously.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MixIpTarget *bool `json:"MixIpTarget,omitempty" name:"MixIpTarget"`

	// Availability zone of a VPC-based private network CLB instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Whether it is an NFV CLB instance. No returned information: no; l7nfv: yes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NfvInfo *string `json:"NfvInfo,omitempty" name:"NfvInfo"`

	// Health check logset ID of CLB CLS
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	HealthLogSetId *string `json:"HealthLogSetId,omitempty" name:"HealthLogSetId"`

	// Health check log topic ID of CLB CLS
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	HealthLogTopicId *string `json:"HealthLogTopicId,omitempty" name:"HealthLogTopicId"`

	// 
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// CLB attribute
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AttributeFlags []*string `json:"AttributeFlags,omitempty" name:"AttributeFlags"`
}

type LoadBalancerDetail struct {

	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// CLB instance network type:
	// Public: public network; Private: private network.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CLB instance status, including:
	// 0: creating; 1: running.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// CLB instance VIP.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitempty" name:"Address"`

	// IPv6 VIP address of the CLB instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIPv6 *string `json:"AddressIPv6,omitempty" name:"AddressIPv6"`

	// IP version of the CLB instance. Valid values: IPv4, IPv6.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// IPv6 address type of the CLB instance. Valid values: IPv6Nat64, IPv6FullChain.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IPv6Mode *string `json:"IPv6Mode,omitempty" name:"IPv6Mode"`

	// Availability zone where the CLB instance resides.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ISP to which the CLB IP address belongs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIsp *string `json:"AddressIsp,omitempty" name:"AddressIsp"`

	// ID of the VPC instance to which the CLB instance belongs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the project to which the CLB instance belongs. 0: default project.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// CLB instance creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// CLB instance billing mode.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// CLB instance network attribute.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitempty" name:"NetworkAttributes"`

	// Pay-as-you-go attribute of the CLB instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitempty" name:"PrepaidAttributes"`

	// Reserved field, which can be ignored generally.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitempty" name:"ExtraInfo"`

	// Custom configuration ID at the CLB instance level.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// CLB instance tag information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// CLB listener ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener protocol.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener port.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Forwarding rule ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding rule path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// ID of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// Address of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetAddress *string `json:"TargetAddress,omitempty" name:"TargetAddress"`

	// Listening port of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetPort *uint64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// Forwarding weight of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetWeight *uint64 `json:"TargetWeight,omitempty" name:"TargetWeight"`

	// 0: not isolated; 1: isolated.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Isolation *uint64 `json:"Isolation,omitempty" name:"Isolation"`

	// List of the security groups bound to the CLB instance.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// Whether the CLB instance is billed by IP.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LoadBalancerPassToTarget *uint64 `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// Health status of the target real server.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TargetHealth *string `json:"TargetHealth,omitempty" name:"TargetHealth"`
}

type LoadBalancerHealth struct {

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// List of listeners
	// Note: This field may return null, indicating that no valid values can be obtained.
	Listeners []*ListenerHealth `json:"Listeners,omitempty" name:"Listeners"`
}

type LoadBalancerTraffic struct {

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// CLB instance region
	Region *string `json:"Region,omitempty" name:"Region"`

	// CLB instance VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Maximum outbound bandwidth in Mbps
	OutBandwidth *float64 `json:"OutBandwidth,omitempty" name:"OutBandwidth"`
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
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
}

func (r *ManualRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SourceListenerId")
	delete(f, "TargetListenerId")
	delete(f, "RewriteInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManualRewriteRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateClassicalLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// Array of classic CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// Exclusive cluster information
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`
}

func (r *MigrateClassicalLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateClassicalLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "ExclusiveCluster")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateClassicalLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MigrateClassicalLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateClassicalLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateClassicalLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIPListRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// Operation type. Valid values:
	// <li> add_customized_field (sets header initially to enable the blocklist feature)</li>
	// <li> set_customized_field (modifies header)</li>
	// <li> del_customized_field (deletes header)</li>
	// <li> add_blocked (adds IPs to blocklist)</li>
	// <li> del_blocked (deletes IPs from blocklist)</li>
	// <li> flush_blocked (clears blocklist)</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Header field that stores real client IPs
	ClientIPField *string `json:"ClientIPField,omitempty" name:"ClientIPField"`

	// List of blocked IPs. The array can contain up to 200,000 entries in one operation.
	BlockIPList []*string `json:"BlockIPList,omitempty" name:"BlockIPList"`

	// Expiration time in seconds. Default value: 3600
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// IP adding policy. Valid value: fifo (if a blocklist is full, new IPs added to the blocklist will adopt the first-in first-out policy)
	AddStrategy *string `json:"AddStrategy,omitempty" name:"AddStrategy"`
}

func (r *ModifyBlockIPListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIPListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "Type")
	delete(f, "ClientIPField")
	delete(f, "BlockIPList")
	delete(f, "ExpireTime")
	delete(f, "AddStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockIPListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIPListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID
		JodId *string `json:"JodId,omitempty" name:"JodId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockIPListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIPListResponse) FromJsonString(s string) error {
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

	// A listener must be configured with a default domain name. If you need to disable the default domain name, you must specify another one as the new default domain name.
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`
}

func (r *ModifyDomainAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "NewDomain")
	delete(f, "Certificate")
	delete(f, "Http2")
	delete(f, "DefaultServer")
	delete(f, "NewDefaultServerDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainAttributesRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID.
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "NewDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Health check parameter, which is applicable only to TCP, UDP, and TCP_SSL listeners.
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Certificate information. This parameter is applicable only to HTTPS and TCP_SSL listeners.
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// Forwarding method of a listener. Value range: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Whether to enable the SNI feature. This parameter is applicable only to HTTPS listeners. Note: The SNI feature can be enabled but cannot be disabled once enabled.
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// Whether to enable a persistent connection. This parameter is applicable only to HTTP and HTTPS listeners.
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`

	// Session persistence type. `NORMAL`: default session persistence type (L4/L7 session persistence); `QUIC_CID`: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners.
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "SessionExpireTime")
	delete(f, "HealthCheck")
	delete(f, "Certificate")
	delete(f, "Scheduler")
	delete(f, "SniSwitch")
	delete(f, "KeepaliveEnable")
	delete(f, "DeregisterTargetRst")
	delete(f, "SessionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Specifies whether to enable deletion protection.
	DeleteProtect *bool `json:"DeleteProtect,omitempty" name:"DeleteProtect"`
}

func (r *ModifyLoadBalancerAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "TargetRegionInfo")
	delete(f, "InternetChargeInfo")
	delete(f, "LoadBalancerPassToTarget")
	delete(f, "SnatPro")
	delete(f, "DeleteProtect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAttributesRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerSlaRequest struct {
	*tchttp.BaseRequest

	// CLB instance information
	LoadBalancerSla []*SlaUpdateParam `json:"LoadBalancerSla,omitempty" name:"LoadBalancerSla"`
}

func (r *ModifyLoadBalancerSlaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerSlaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerSla")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerSlaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerSlaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancerSlaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerSlaResponse) FromJsonString(s string) error {
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

	// New forwarding path of the forwarding rule. This parameter is not required if the URL does not need to be modified.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Health check information
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Request forwarding method of the rule. Value range: WRR, LEAST_CONN, IP_HASH
	// They represent weighted round robin, least connections, and IP hash, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Session persistence time
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Forwarding protocol between CLB instance and real server. Default value: HTTP. Valid values: HTTP, HTTPS, and TRPC.
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`

	// TRPC callee server route, which is required when `ForwardType` is "TRPC".
	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`

	// TRPC calling service API, which is required when `ForwardType` is "TRPC".
	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "LocationId")
	delete(f, "Url")
	delete(f, "HealthCheck")
	delete(f, "Scheduler")
	delete(f, "SessionExpireTime")
	delete(f, "ForwardType")
	delete(f, "TrpcCallee")
	delete(f, "TrpcFunc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupName")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupAttributeRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupInstancesPortRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify ports
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

func (r *ModifyTargetGroupInstancesPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupInstancesPortRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetGroupInstancesWeightRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify weights
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

func (r *ModifyTargetGroupInstancesWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupInstancesWeightRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// New port of the real server bound to a listener or forwarding rule
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *ModifyTargetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "NewPort")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetPortRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitempty" name:"Url"`

	// List of real servers for which to modify the weights
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// New forwarding weight of a real server. Value range: 0-100. Default value: 10. If the Targets.Weight parameter is set, this parameter will not take effect.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "Targets")
	delete(f, "Weight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetWeightRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Quota struct {

	// Quota name. Valid values:
	// <li> TOTAL_OPEN_CLB_QUOTA: quota of public network CLB instances in the current region</li>
	// <li> TOTAL_INTERNAL_CLB_QUOTA: quota of private network CLB instances in the current region</li>
	// <li> TOTAL_LISTENER_QUOTA: quota of listeners under one CLB instance</li>
	// <li> TOTAL_LISTENER_RULE_QUOTA: quota of forwarding rules under one listener</li>
	// <li> TOTAL_TARGET_BIND_QUOTA: quota of CVM instances can be bound under one forwarding rule</li>
	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`

	// Currently used quantity. If it is `null`, it is meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`

	// Quota limit.
	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type RegisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Server instance array
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

func (r *RegisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterTargetGroupInstancesRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers to be bound. Array length limit: 20.
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target forwarding rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target forwarding rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *RegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterTargetsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsWithClassicalLBRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Real server information
	Targets []*ClassicalTargetInfo `json:"Targets,omitempty" name:"Targets"`
}

func (r *RegisterTargetsWithClassicalLBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetsWithClassicalLBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterTargetsWithClassicalLBRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertForLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OldCertificateId")
	delete(f, "Certificate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceCertForLoadBalancersRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertForLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RewriteLocationMap struct {

	// Source forwarding rule ID
	SourceLocationId *string `json:"SourceLocationId,omitempty" name:"SourceLocationId"`

	// Forwarding rule ID of a redirect target
	TargetLocationId *string `json:"TargetLocationId,omitempty" name:"TargetLocationId"`

	// Redirection status code. Valid values: 301, 302, and 307.
	RewriteCode *int64 `json:"RewriteCode,omitempty" name:"RewriteCode"`

	// Whether the matched URL is carried in redirection. It is required when configuring `RewriteCode`.
	TakeUrl *bool `json:"TakeUrl,omitempty" name:"TakeUrl"`

	// Original domain name of redirection, which must be the corresponding domain name of `SourceLocationId`. It is required when configuring `RewriteCode`.
	SourceDomain *string `json:"SourceDomain,omitempty" name:"SourceDomain"`
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

	// Redirection status code
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RewriteCode *int64 `json:"RewriteCode,omitempty" name:"RewriteCode"`

	// Whether the matched URL is carried in redirection.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TakeUrl *bool `json:"TakeUrl,omitempty" name:"TakeUrl"`

	// Redirection type. Manual: manual redirection; Auto: automatic redirection.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RewriteType *string `json:"RewriteType,omitempty" name:"RewriteType"`
}

type RsWeightRule struct {

	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of real servers whose weights to modify.
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// Forwarding rule ID, which is required only for layer-7 rules.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified
	Url *string `json:"Url,omitempty" name:"Url"`

	// The new forwarding weight of the real server. Value range: [0, 100]. This parameter takes lower precedence than `Weight` in [`Targets`](https://intl.cloud.tencent.com/document/api/214/30694?from_cn_redirect=1#Target), which means that this parameter only takes effect when the `Weight` in `RsWeightRule` is left empty.
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

	// Health status of the real server bound to this rule
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Targets []*TargetHealth `json:"Targets,omitempty" name:"Targets"`
}

type RuleInput struct {

	// Domain name of the forwarding rule. Length: 1-80.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding rule path. Length: 1-200.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Session persistence time in seconds. Value range: 30-3,600. Setting it to 0 indicates that session persistence is disabled.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// Health check information. For more information, please see [Health Check](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1)
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

	// Whether to enable QUIC. Note: QUIC can be enabled only for HTTPS domain names
	Quic *bool `json:"Quic,omitempty" name:"Quic"`
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

	// QUIC status
	// Note: this field may return null, indicating that no valid values can be obtained.
	QuicStatus *string `json:"QuicStatus,omitempty" name:"QuicStatus"`
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
	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`
}

type RulesItems struct {

	// Rule ID.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// Domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Uri
	Url *string `json:"Url,omitempty" name:"Url"`

	// Object bound to the real server.
	Targets []*LbRsTargets `json:"Targets,omitempty" name:"Targets"`
}

type SetLoadBalancerClsLogRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLS logset ID
	// <li>Enter the ID of logset you need to add or update. You can acquire the ID by invoking [DescribeLogsets](https://intl.cloud.tencent.com/document/product/614/56454?from_cn_redirect=1).</li>
	// <li>To delete the log set, set this parameter to `null`.</li>
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// CLS log topic ID
	// <li>Enter the ID of log topic you need to add or update. You can acquire the ID by invoking [DescribeTopics](https://intl.cloud.tencent.com/document/product/614/56454?from_cn_redirect=1).</li>
	// <li>To delete the log set, set this parameter to `null`.</li>
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// Log type:
	// <li>`ACCESS`: access logs</li>
	// <li>`HEALTH`: health check logs</li>
	// Default: `ACCESS`
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

func (r *SetLoadBalancerClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerClsLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LogSetId")
	delete(f, "LogTopicId")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLoadBalancerClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetLoadBalancerClsLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoadBalancerClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerClsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Array of security group IDs. One CLB instance can be bound to up to 50 security groups. If you want to unbind all security groups, you do not need to pass in this parameter, or you can pass in an empty array.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

func (r *SetLoadBalancerSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SecurityGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLoadBalancerSecurityGroupsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *SetSecurityGroupForLoadbalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroup")
	delete(f, "OperationType")
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetSecurityGroupForLoadbalancersRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlaUpdateParam struct {

	// ID of the CLB instance
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// To upgrade to LCU-supported CLB instances. It must be `SLA`.
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`
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
	// Note: this parameter is required when binding a CVM or ENI.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Real server type. Value range: CVM (Cloud Virtual Machine), ENI (Elastic Network Interface). This parameter does not take effect currently as an input parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Unique ID of a CVM instance, which is required when binding a CVM instance. It can be obtained from the `InstanceId` field in the response of the `DescribeInstances` API. It indicates binding the primary IP of the primary ENI.
	// Note: either `InstanceId` or `EniIp` must be passed in.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The new forwarding weight of the real server. Value range: [0, 100]. Default: 10. This parameter takes priority over `Weight` in [`RsWeightRule`](https://intl.cloud.tencent.com/document/api/214/30694?from_cn_redirect=1#RsWeightRule). If it’s left empty, the value of `Weight` in `RsWeightRule` will be used.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// It is required when binding an IP. ENI IPs and other private IPs are supported. To bind an ENI IP, the ENI should be bound to a CVM instance before being bound to a CLB instance.
	// Note: either `InstanceId` or `EniIp` must be passed in. It is required when binding a dual-stack IPv6 CVM instance.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
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
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// Private IP of real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Real server instance name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Real server binding time
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`

	// Unique ENI ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	EniId *string `json:"EniId,omitempty" name:"EniId"`

	// AZ ID of the real server
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
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
	AssociatedRule []*AssociationItem `json:"AssociatedRule,omitempty" name:"AssociatedRule"`
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

	// AZ region, e.g., ap-guangzhou.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ZoneRegion *string `json:"ZoneRegion,omitempty" name:"ZoneRegion"`

	// Whether the AZ is the `LocalZone`, e.g., false.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LocalZone *bool `json:"LocalZone,omitempty" name:"LocalZone"`
}
