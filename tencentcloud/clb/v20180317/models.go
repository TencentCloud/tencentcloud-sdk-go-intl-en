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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AssociateTargetGroupsRequestParams struct {
	// Association array
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

type AssociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Association array
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
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

// Predefined struct for user
type AssociateTargetGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateTargetGroupsResponseParams `json:"Response"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// ID of associated listener
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// ID of associated forwarding rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Protocol type of associated listener, such as HTTP or TCP
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port of associated listener
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Domain name of associated forwarding rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// URL of associated forwarding rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`
}

// Predefined struct for user
type AutoRewriteRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// `HTTPS:443` listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The domain name to be redirected under the listener `HTTPS:443`. If it is left empty, all domain names under the listener `HTTPS:443` will be configured with redirects.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Redirection status code. Valid values: 301, 302, and 307.
	RewriteCodes []*int64 `json:"RewriteCodes,omitnil,omitempty" name:"RewriteCodes"`

	// Whether the matched URL is carried in redirection.
	TakeUrls []*bool `json:"TakeUrls,omitnil,omitempty" name:"TakeUrls"`
}

type AutoRewriteRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// `HTTPS:443` listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The domain name to be redirected under the listener `HTTPS:443`. If it is left empty, all domain names under the listener `HTTPS:443` will be configured with redirects.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Redirection status code. Valid values: 301, 302, and 307.
	RewriteCodes []*int64 `json:"RewriteCodes,omitnil,omitempty" name:"RewriteCodes"`

	// Whether the matched URL is carried in redirection.
	TakeUrls []*bool `json:"TakeUrls,omitnil,omitempty" name:"TakeUrls"`
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

// Predefined struct for user
type AutoRewriteResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AutoRewriteResponse struct {
	*tchttp.BaseResponse
	Response *AutoRewriteResponseParams `json:"Response"`
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
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Unique ID of a real server, which can be obtained from the unInstanceId field in the return of the DescribeInstances API
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Listening port of a real server
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Forwarding weight of a real server. Value range: [0, 100]. Default value: 10.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Public IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// Private IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// Real server instance names
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Bound time of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// Unique ENI ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`
}

type BasicTargetGroupInfo struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Target group name
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`
}

// Predefined struct for user
type BatchDeregisterTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Unbinding targets
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type BatchDeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Unbinding targets
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
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

// Predefined struct for user
type BatchDeregisterTargetsResponseParams struct {
	// IDs of the listeners failed to unbind
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitnil,omitempty" name:"FailListenerIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeregisterTargetsResponseParams `json:"Response"`
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

// Predefined struct for user
type BatchModifyTargetWeightRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of weights to be modified in batches
	ModifyList []*RsWeightRule `json:"ModifyList,omitnil,omitempty" name:"ModifyList"`
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of weights to be modified in batches
	ModifyList []*RsWeightRule `json:"ModifyList,omitnil,omitempty" name:"ModifyList"`
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

// Predefined struct for user
type BatchModifyTargetWeightResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyTargetWeightResponseParams `json:"Response"`
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

// Predefined struct for user
type BatchRegisterTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Binding target
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Binding target
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
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

// Predefined struct for user
type BatchRegisterTargetsResponseParams struct {
	// IDs of the listeners failed to bind. If this is blank, all listeners are bound successfully.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitnil,omitempty" name:"FailListenerIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchRegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchRegisterTargetsResponseParams `json:"Response"`
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
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The port to Bind
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// CVM instance ID. The primary IP of the primary ENI will be bound.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// It is required for binding an IP. It supports an ENI IP or any other private IP. To bind an ENI IP, the ENI should be bound to a CVM instance before being bound to a CLB instance.
	// Note: either `InstanceId` or `EniIp` must be passed in, which is required for binding a dual-stack IPv6 CVM instance.
	EniIp *string `json:"EniIp,omitnil,omitempty" name:"EniIp"`

	// Weight of the CVM instance. Value range: [0, 100]. If it is not specified for binding the instance, 10 will be used by default.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Layer-7 rule ID.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`
}

type BindDetailItem struct {
	// Specifies the ID of CLB to be bound
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Specifies the ID of listener to be bound
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Specifies the domain name to be bound
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Sets the bound rule.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Listener name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Listener protocol.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// URL of the location.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Configuration ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`
}

type BlockedIP struct {
	// Blacklisted IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Blacklisted time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Expiration time
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type CertIdRelatedWithLoadBalancers struct {
	// Certificate ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// List of CLB instances associated with certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancers []*LoadBalancer `json:"LoadBalancers,omitnil,omitempty" name:"LoadBalancers"`
}

type CertInfo struct {
	// ID of the certificate. If it's not specified, `CertContent` and `CertKey` are required. For a server certificate, you also need to specify `CertName`. 
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Name of the uploaded certificate. It's required if `CertId` is not specified.
	CertName *string `json:"CertName,omitnil,omitempty" name:"CertName"`

	// Public key of the uploaded certificate. This is required if `CertId` is not specified.
	CertContent *string `json:"CertContent,omitnil,omitempty" name:"CertContent"`

	// Private key of the uploaded server certificate. This is required if `CertId` is not specified.
	CertKey *string `json:"CertKey,omitnil,omitempty" name:"CertKey"`
}

type CertificateInput struct {
	// Authentication type. Value range: UNIDIRECTIONAL (unidirectional authentication), MUTUAL (mutual authentication)
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// ID of a server certificate. If you leave this parameter empty, you must upload the certificate, including CertContent, CertKey, and CertName.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// ID of a client certificate. When the listener adopts mutual authentication (i.e., SSLMode = mutual), if you leave this parameter empty, you must upload the client certificate, including CertCaContent and CertCaName.
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// Name of the uploaded server certificate. If there is no CertId, this parameter is required.
	CertName *string `json:"CertName,omitnil,omitempty" name:"CertName"`

	// Key of the uploaded server certificate. If there is no CertId, this parameter is required.
	CertKey *string `json:"CertKey,omitnil,omitempty" name:"CertKey"`

	// Content of the uploaded server certificate. If there is no CertId, this parameter is required.
	CertContent *string `json:"CertContent,omitnil,omitempty" name:"CertContent"`

	// Name of the uploaded client CA certificate. When SSLMode = mutual, if there is no CertCaId, this parameter is required.
	CertCaName *string `json:"CertCaName,omitnil,omitempty" name:"CertCaName"`

	// Content of the uploaded client certificate. When SSLMode = mutual, if there is no CertCaId, this parameter is required.
	CertCaContent *string `json:"CertCaContent,omitnil,omitempty" name:"CertCaContent"`
}

type CertificateOutput struct {
	// Authentication type. Value range: UNIDIRECTIONAL (unidirectional authentication), MUTUAL (mutual authentication)
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// Server certificate ID.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Client certificate ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// IDs of extra server certificates
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExtCertIds []*string `json:"ExtCertIds,omitnil,omitempty" name:"ExtCertIds"`
}

type ClassicalHealth struct {
	// Private IP of a real server
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Real server port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// CLB listener port
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// Forwarding protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Health check result. Value range: 1 (healthy), 0 (unhealthy)
	HealthStatus *int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`
}

type ClassicalListener struct {
	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// CLB listener port
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// Backend forwarding port of a listener
	InstancePort *int64 `json:"InstancePort,omitnil,omitempty" name:"InstancePort"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Listener protocol type
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Session persistence time
	SessionExpire *int64 `json:"SessionExpire,omitnil,omitempty" name:"SessionExpire"`

	// Whether health check is enabled. 1: enabled; 0: disabled.
	HealthSwitch *int64 `json:"HealthSwitch,omitnil,omitempty" name:"HealthSwitch"`

	// Response timeout period
	TimeOut *int64 `json:"TimeOut,omitnil,omitempty" name:"TimeOut"`

	// Check interval
	IntervalTime *int64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// Health threshold
	HealthNum *int64 `json:"HealthNum,omitnil,omitempty" name:"HealthNum"`

	// Unhealthy threshold
	UnhealthNum *int64 `json:"UnhealthNum,omitnil,omitempty" name:"UnhealthNum"`

	// A request balancing method for HTTP and HTTPS listeners of a public network classic CLB. wrr means weighted round robin, while ip_hash means consistent hashing based on source IPs of access requests.
	HttpHash *string `json:"HttpHash,omitnil,omitempty" name:"HttpHash"`

	// Health check return code for HTTP and HTTPS listeners of a public network classic CLB. For more information, see the explanation of the field in the listener creating API.
	HttpCode *int64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// Health check path for HTTP and HTTPS listeners of a public network classic CLB
	HttpCheckPath *string `json:"HttpCheckPath,omitnil,omitempty" name:"HttpCheckPath"`

	// Authentication method for an HTTPS listener of a public network classic CLB
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// Server certificate ID for an HTTPS listener of a public network classic CLB
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Client certificate ID for an HTTPS listener of a public network classic CLB
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// Listener status. Value range: 0 (creating), 1 (running)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ClassicalLoadBalancerInfo struct {
	// Real server ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of CLB instance IDs
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type ClassicalTarget struct {
	// Real server type. Value range: CVM, ENI (coming soon)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Unique ID of a real server, which can be obtained from the unInstanceId field in the return of the DescribeInstances API
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Forwarding weight of a real server. Value range: [0, 100]. Default value: 10.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Public IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// Private IP of a real server
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// Real server instance names
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Real server status
	// 1: failed; 2: running; 3: creating; 4: shut down; 5: returned; 6: returning; 7: restarting; 8: starting; 9: shutting down; 10: resetting password; 11: formatting; 12: making image; 13: setting bandwidth; 14: reinstalling system; 19: upgrading; 21: hot migrating
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunFlag *int64 `json:"RunFlag,omitnil,omitempty" name:"RunFlag"`
}

type ClassicalTargetInfo struct {
	// Real server ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Weight. Value range: [0, 100]
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type CloneLoadBalancerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Clones the name of the CLB instance. The name must be 1-60 characters containing letters, numbers, "-" or "_".
	// Note: if the name of a new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the `DescribeProject` API. If this parameter is not passed in, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Sets the primary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`, which is applicable only to public network CLB.
	// Note: A primary AZ loads traffic, while a secondary AZ does not load traffic by default and will be used only if the primary AZ becomes unavailable. The platform will automatically select the optimal secondary AZ. You can use the `DescribeResource` API to query the primary AZ list of a region.
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// Specifies the secondary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`. It is applicable only to public network CLB.
	// Note: A secondary AZ will load traffic if the primary AZ is faulty. You can use the `DescribeMasterZones` API to query the primary and secondary AZ list of a region.
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// Specifies an AZ ID for creating a CLB instance, such as `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CLB network billing mode. This parameter is applicable only to public network CLB instances.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// ISP of VIP. Values: `CMCC` (China Mobile), `CUCC` (China Unicom) and `CTCC` (China Telecom). You need to activate static single-line IPs. This feature is in beta and is only available in Guangzhou, Shanghai, Nanjing, Jinan, Hangzhou, Fuzhou, Beijing, Shijiazhuang, Wuhan, Changsha, Chengdu and Chongqing regions. To try it out, please contact your sales rep. If it's specified, the network billing mode must be `BANDWIDTH_PACKAGE`. If it's not specified, BGP is used by default. To query ISPs supported in a region, please use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1). 
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// Applies for CLB instances for a specified VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Tags a CLB instance when purchasing it
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Dedicated cluster information
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// Bandwidth package ID. If this parameter is specified, the network billing mode (`InternetAccessible.InternetChargeType`) will only support bill-by-bandwidth package (`BANDWIDTH_PACKAGE`).
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// Whether to support binding cross-VPC IPs or cross-region IPs
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// Creates `SnatIp` when the binding IPs of other VPCs feature is enabled
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// ID of the public network CLB dedicated cluster
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Specification of the LCU-supported instance.
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// Tag of the STGW dedicated cluster
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// Availability zones for nearby access of private network CLB instances to distribute traffic
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Unique ID of an EIP, which can only be used when binding the EIP of a private network CLB instance (e.g., `eip-11112222`)
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`
}

type CloneLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Clones the name of the CLB instance. The name must be 1-60 characters containing letters, numbers, "-" or "_".
	// Note: if the name of a new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the `DescribeProject` API. If this parameter is not passed in, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Sets the primary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`, which is applicable only to public network CLB.
	// Note: A primary AZ loads traffic, while a secondary AZ does not load traffic by default and will be used only if the primary AZ becomes unavailable. The platform will automatically select the optimal secondary AZ. You can use the `DescribeResource` API to query the primary AZ list of a region.
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// Specifies the secondary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`. It is applicable only to public network CLB.
	// Note: A secondary AZ will load traffic if the primary AZ is faulty. You can use the `DescribeMasterZones` API to query the primary and secondary AZ list of a region.
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// Specifies an AZ ID for creating a CLB instance, such as `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CLB network billing mode. This parameter is applicable only to public network CLB instances.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// ISP of VIP. Values: `CMCC` (China Mobile), `CUCC` (China Unicom) and `CTCC` (China Telecom). You need to activate static single-line IPs. This feature is in beta and is only available in Guangzhou, Shanghai, Nanjing, Jinan, Hangzhou, Fuzhou, Beijing, Shijiazhuang, Wuhan, Changsha, Chengdu and Chongqing regions. To try it out, please contact your sales rep. If it's specified, the network billing mode must be `BANDWIDTH_PACKAGE`. If it's not specified, BGP is used by default. To query ISPs supported in a region, please use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1). 
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// Applies for CLB instances for a specified VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Tags a CLB instance when purchasing it
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Dedicated cluster information
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// Bandwidth package ID. If this parameter is specified, the network billing mode (`InternetAccessible.InternetChargeType`) will only support bill-by-bandwidth package (`BANDWIDTH_PACKAGE`).
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// Whether to support binding cross-VPC IPs or cross-region IPs
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// Creates `SnatIp` when the binding IPs of other VPCs feature is enabled
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// ID of the public network CLB dedicated cluster
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Specification of the LCU-supported instance.
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// Tag of the STGW dedicated cluster
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// Availability zones for nearby access of private network CLB instances to distribute traffic
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Unique ID of an EIP, which can only be used when binding the EIP of a private network CLB instance (e.g., `eip-11112222`)
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`
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

// Predefined struct for user
type CloneLoadBalancerResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CloneLoadBalancerResponseParams `json:"Response"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Cluster AZ, such as ap-guangzhou-1
	// Note: this field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ConfigListItem struct {
	// Configuration ID.
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// Configuration type.
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// Configuration name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// Configuration content.
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// Creates configuration time.
	CreateTimestamp *string `json:"CreateTimestamp,omitnil,omitempty" name:"CreateTimestamp"`

	// Modifies configuration time.
	UpdateTimestamp *string `json:"UpdateTimestamp,omitnil,omitempty" name:"UpdateTimestamp"`
}

// Predefined struct for user
type CreateClsLogSetRequestParams struct {
	// Logset name, which must be unique among all CLS logsets; default value: clb_logset
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Logset retention period (in days)
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Logset type. Valid values: ACCESS (access logs; default value) and HEALTH (health check logs).
	LogsetType *string `json:"LogsetType,omitnil,omitempty" name:"LogsetType"`
}

type CreateClsLogSetRequest struct {
	*tchttp.BaseRequest
	
	// Logset name, which must be unique among all CLS logsets; default value: clb_logset
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Logset retention period (in days)
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Logset type. Valid values: ACCESS (access logs; default value) and HEALTH (health check logs).
	LogsetType *string `json:"LogsetType,omitnil,omitempty" name:"LogsetType"`
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
	delete(f, "LogsetName")
	delete(f, "Period")
	delete(f, "LogsetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClsLogSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClsLogSetResponseParams struct {
	// Logset ID.
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClsLogSetResponse struct {
	*tchttp.BaseResponse
	Response *CreateClsLogSetResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateListenerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Specifies for which ports to create listeners. Each port corresponds to a new listener.
	Ports []*int64 `json:"Ports,omitnil,omitempty" name:"Ports"`

	// Listener protocol. Values: TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// List of names of the listeners to be created. The array of names and array of ports are in one-to-one correspondence. If you do not want to name them now, you do not need to provide this parameter.
	ListenerNames []*string `json:"ListenerNames,omitnil,omitempty" name:"ListenerNames"`

	// Health check parameter. It is only applicable only to TCP, UDP, TCP_SSL and QUIC listeners.
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Certificate information. This parameter is only applicable to TCP_SSL listeners and HTTPS listeners with the SNI feature not enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Session persistence time in seconds. Value range: 30-3,600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Listener forwarding mode. Values: `WRR` (weighted round robin) and `LEAST_CONN` (least connections). 
	// Default value: `WRR`. This parameter is only applicable to TCP, UDP, TCP_SSL and QUIC listeners.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Whether to enable the SNI feature. This parameter is applicable only to HTTPS listeners
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// Target real server type. `NODE`: binding a general node; `TARGETGROUP`: binding a target group.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Session persistence type. Valid values: Normal: the default session persistence type; QUIC_CID: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners. If this field is not specified, the default session persistence type will be used.
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// Whether to enable a persistent connection. This parameter is applicable only to HTTP and HTTPS listeners. Valid values: 0 (disable; default value) and 1 (enable).
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// This parameter is used to specify the end port and is required when creating a port range listener. Only one member can be passed in when inputting the `Ports` parameter, which is used to specify the start port. If you want to try the port range feature, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
	EndPort *uint64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// Certificate information. You can specify multiple server-side certificates with different algorithm types. This parameter is only applicable to HTTPS listeners with the SNI feature not enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// Maximum number of concurrent listener connections. It's available for TCP/UDP/TCP_SSL/QUIC listeners. If it's set to `-1` or not specified, the listener speed is not limited. 
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// Maximum number of new listener connections. It's available for TCP/UDP/TCP_SSL/QUIC listeners. If it's set to `-1` or not specified, the listener speed is not limited. 
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// Connection idle timeout period (in seconds). It's only available to TCP listeners. Value range: 300-900 for shared and dedicated instances; 300-2000 for LCU-supported CLB instances. It defaults to 900. To set a period longer than 2000 seconds (up to 3600 seconds), please submit a [submit](https://console.cloud.tencent.com/workorder/category). 
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`


	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Specifies for which ports to create listeners. Each port corresponds to a new listener.
	Ports []*int64 `json:"Ports,omitnil,omitempty" name:"Ports"`

	// Listener protocol. Values: TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// List of names of the listeners to be created. The array of names and array of ports are in one-to-one correspondence. If you do not want to name them now, you do not need to provide this parameter.
	ListenerNames []*string `json:"ListenerNames,omitnil,omitempty" name:"ListenerNames"`

	// Health check parameter. It is only applicable only to TCP, UDP, TCP_SSL and QUIC listeners.
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Certificate information. This parameter is only applicable to TCP_SSL listeners and HTTPS listeners with the SNI feature not enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Session persistence time in seconds. Value range: 30-3,600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Listener forwarding mode. Values: `WRR` (weighted round robin) and `LEAST_CONN` (least connections). 
	// Default value: `WRR`. This parameter is only applicable to TCP, UDP, TCP_SSL and QUIC listeners.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Whether to enable the SNI feature. This parameter is applicable only to HTTPS listeners
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// Target real server type. `NODE`: binding a general node; `TARGETGROUP`: binding a target group.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Session persistence type. Valid values: Normal: the default session persistence type; QUIC_CID: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners. If this field is not specified, the default session persistence type will be used.
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// Whether to enable a persistent connection. This parameter is applicable only to HTTP and HTTPS listeners. Valid values: 0 (disable; default value) and 1 (enable).
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// This parameter is used to specify the end port and is required when creating a port range listener. Only one member can be passed in when inputting the `Ports` parameter, which is used to specify the start port. If you want to try the port range feature, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
	EndPort *uint64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// Certificate information. You can specify multiple server-side certificates with different algorithm types. This parameter is only applicable to HTTPS listeners with the SNI feature not enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// Maximum number of concurrent listener connections. It's available for TCP/UDP/TCP_SSL/QUIC listeners. If it's set to `-1` or not specified, the listener speed is not limited. 
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// Maximum number of new listener connections. It's available for TCP/UDP/TCP_SSL/QUIC listeners. If it's set to `-1` or not specified, the listener speed is not limited. 
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// Connection idle timeout period (in seconds). It's only available to TCP listeners. Value range: 300-900 for shared and dedicated instances; 300-2000 for LCU-supported CLB instances. It defaults to 900. To set a period longer than 2000 seconds (up to 3600 seconds), please submit a [submit](https://console.cloud.tencent.com/workorder/category). 
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`

	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`
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
	delete(f, "MultiCertInfo")
	delete(f, "MaxConn")
	delete(f, "MaxCps")
	delete(f, "IdleConnectTimeout")
	delete(f, "SnatEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// Array of unique IDs of the created listeners
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *CreateListenerResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLoadBalancerRequestParams struct {
	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// CLB instance type. Valid value: 1 (generic CLB instance).
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// CLB instance name, which takes effect only when only one instance is to be created in the request. It can consist 1 to 60 letters, digits, hyphens (-), or underscores (_).
	// Note: if the name of the new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// Network ID of the target device on the CLB backend, such as `vpc-12345678`, which can be obtained through the `DescribeVpcEx` API. If this parameter is not entered, `DefaultVPC` is used by default. This parameter is required when creating a private network instance.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// A subnet ID must be specified when you purchase a private network CLB instance in a VPC, and the VIP of this instance will be generated in this subnet. This parameter is required for creating a CLB instance.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the `DescribeProject` API. If this parameter is not entered, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// It's only applicable to public network CLB instances. IP version. Values: `IPV4`, `IPV6` and `IPv6FullChain` (case-insensitive). Default: `IPV4`. Note: `IPV6` indicates IPv6 NAT64, while `IPv6FullChain` indicates IPv6. 
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// Number of CLBs to be created. Default value: 1.
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// ID of the primary AZ for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`. It's only available to public CLB instances. 
	// Note: The traffic only goes to the primary AZ in normal cases. The secondary AZ is used only when the primary AZ is unavailable. To query the list of primary AZs in a region, use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1).
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// Specifies an AZ ID for creating a CLB instance, such as `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// It only works on LCU-supported instances on private networks and all instances on public networks.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// ISP of VIP. Values: `CMCC` (China Mobile), `CUCC` (China Unicom) and `CTCC` (China Telecom). You need to activate static single-line IPs. This feature is in beta and is only available in Guangzhou, Shanghai, Nanjing, Jinan, Hangzhou, Fuzhou, Beijing, Shijiazhuang, Wuhan, Changsha, Chengdu and Chongqing regions. To try it out, please contact your sales rep. If it's specified, the network billing mode must be `BANDWIDTH_PACKAGE`. If it's not specified, BGP is used by default. To query ISPs supported in a region, please use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1). 
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// Tags the CLB instance when purchasing it. Up to 20 tag key value pairs are supported.
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	// Note: If the specified VIP is occupied or is not within the IP range of the specified VPC subnet, you cannot use the VIP to create a CLB instance in a private network or an IPv6 BGP CLB instance in a public network.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Bandwidth package ID. If this parameter is specified, the network billing mode (`InternetAccessible.InternetChargeType`) will only support bill-by-bandwidth package (`BANDWIDTH_PACKAGE`).
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// Information about the dedicated CLB instance. You must specify this parameter when you create a dedicated CLB instance in a private network.
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// Specification of LCU-supported instance.
	// <ul><li>This parameter is required to create LCU-supported instances. Values: <ul><li>`SLA`: Super Large 4. When you have activated Super Large models, `SLA` refers to Super Large 4.</li><li>`clb.c2.medium`: Standard</li><li>`clb.c3.small`: Advanced 1</li><li>`clb.c3.medium`: Advanced 1</li><li>`clb.c4.small`: Super Large 1</li><li>`clb.c4.medium`: Super Large 2</li><li>`clb.c4.large`: Super Large 3</li><li>`clb.c4.xlarge`: Super Large 4</li> For Super Large 2 and above models, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).</ul></li><li> This parameter is not required for creating shared instances.</li></ul>For more details, see [Instance Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Whether Binding IPs of other VPCs feature switch
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// Creates `SnatIp` when the binding IPs of other VPCs feature is enabled
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// Tag for the STGW exclusive cluster.
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// Specifies the secondary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`. It is applicable only to public network CLB.
	// Note: The traffic only goes to the secondary AZ when the primary AZ is unavailable. You can query the list of primary and secondary AZ of a region by calling [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1).
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// Unique ID of an EIP, which can only be used when binding the EIP of a private network CLB instance. E.g., `eip-11112222`.
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`

	// Whether to allow CLB traffic to the target group. `true`: allows CLB traffic to the target group and verifies security groups only on CLB; `false`: denies CLB traffic to the target group and verifies security groups on both CLB and backend instances.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// Upgrades to domain name-based CLB
	DynamicVip *bool `json:"DynamicVip,omitnil,omitempty" name:"DynamicVip"`

	// Network egress point
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// CLB instance type. Valid value: 1 (generic CLB instance).
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// CLB instance name, which takes effect only when only one instance is to be created in the request. It can consist 1 to 60 letters, digits, hyphens (-), or underscores (_).
	// Note: if the name of the new CLB instance already exists, a default name will be generated automatically.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// Network ID of the target device on the CLB backend, such as `vpc-12345678`, which can be obtained through the `DescribeVpcEx` API. If this parameter is not entered, `DefaultVPC` is used by default. This parameter is required when creating a private network instance.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// A subnet ID must be specified when you purchase a private network CLB instance in a VPC, and the VIP of this instance will be generated in this subnet. This parameter is required for creating a CLB instance.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the `DescribeProject` API. If this parameter is not entered, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// It's only applicable to public network CLB instances. IP version. Values: `IPV4`, `IPV6` and `IPv6FullChain` (case-insensitive). Default: `IPV4`. Note: `IPV6` indicates IPv6 NAT64, while `IPv6FullChain` indicates IPv6. 
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// Number of CLBs to be created. Default value: 1.
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// ID of the primary AZ for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`. It's only available to public CLB instances. 
	// Note: The traffic only goes to the primary AZ in normal cases. The secondary AZ is used only when the primary AZ is unavailable. To query the list of primary AZs in a region, use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1).
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// Specifies an AZ ID for creating a CLB instance, such as `ap-guangzhou-1`, which is applicable only to public network CLB instances.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// It only works on LCU-supported instances on private networks and all instances on public networks.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// ISP of VIP. Values: `CMCC` (China Mobile), `CUCC` (China Unicom) and `CTCC` (China Telecom). You need to activate static single-line IPs. This feature is in beta and is only available in Guangzhou, Shanghai, Nanjing, Jinan, Hangzhou, Fuzhou, Beijing, Shijiazhuang, Wuhan, Changsha, Chengdu and Chongqing regions. To try it out, please contact your sales rep. If it's specified, the network billing mode must be `BANDWIDTH_PACKAGE`. If it's not specified, BGP is used by default. To query ISPs supported in a region, please use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1). 
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// Tags the CLB instance when purchasing it. Up to 20 tag key value pairs are supported.
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	// Note: If the specified VIP is occupied or is not within the IP range of the specified VPC subnet, you cannot use the VIP to create a CLB instance in a private network or an IPv6 BGP CLB instance in a public network.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Bandwidth package ID. If this parameter is specified, the network billing mode (`InternetAccessible.InternetChargeType`) will only support bill-by-bandwidth package (`BANDWIDTH_PACKAGE`).
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// Information about the dedicated CLB instance. You must specify this parameter when you create a dedicated CLB instance in a private network.
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// Specification of LCU-supported instance.
	// <ul><li>This parameter is required to create LCU-supported instances. Values: <ul><li>`SLA`: Super Large 4. When you have activated Super Large models, `SLA` refers to Super Large 4.</li><li>`clb.c2.medium`: Standard</li><li>`clb.c3.small`: Advanced 1</li><li>`clb.c3.medium`: Advanced 1</li><li>`clb.c4.small`: Super Large 1</li><li>`clb.c4.medium`: Super Large 2</li><li>`clb.c4.large`: Super Large 3</li><li>`clb.c4.xlarge`: Super Large 4</li> For Super Large 2 and above models, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).</ul></li><li> This parameter is not required for creating shared instances.</li></ul>For more details, see [Instance Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Whether Binding IPs of other VPCs feature switch
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// Creates `SnatIp` when the binding IPs of other VPCs feature is enabled
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// Tag for the STGW exclusive cluster.
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// Specifies the secondary AZ ID for cross-AZ disaster recovery, such as `100001` or `ap-guangzhou-1`. It is applicable only to public network CLB.
	// Note: The traffic only goes to the secondary AZ when the primary AZ is unavailable. You can query the list of primary and secondary AZ of a region by calling [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1).
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// Unique ID of an EIP, which can only be used when binding the EIP of a private network CLB instance. E.g., `eip-11112222`.
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`

	// Whether to allow CLB traffic to the target group. `true`: allows CLB traffic to the target group and verifies security groups only on CLB; `false`: denies CLB traffic to the target group and verifies security groups on both CLB and backend instances.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// Upgrades to domain name-based CLB
	DynamicVip *bool `json:"DynamicVip,omitnil,omitempty" name:"DynamicVip"`

	// Network egress point
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`
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
	delete(f, "DynamicVip")
	delete(f, "Egress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerResponseParams struct {
	// Array of unique CLB instance IDs.
	// This field may return `null` in some cases, such as there is delay during instance creation. You can query the IDs of the created instances by invoking `DescribeTaskStatus` with the `RequestId` or `DealName` returned by this API.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Order ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancerResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLoadBalancerSnatIpsRequestParams struct {
	// Unique ID of a CLB instance, e.g., lb-12345678.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Information of the SNAT IP to be added. You can specify a SNAT IP or use the one automatically assigned by a subnet.
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// Number of SNAT IPs to be added. This parameter is used in conjunction with `SnatIps`. Note that if `Ip` is specified in `SnapIps`, this parameter is not available. It defaults to `1` and the upper limit is `10`.
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type CreateLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a CLB instance, e.g., lb-12345678.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Information of the SNAT IP to be added. You can specify a SNAT IP or use the one automatically assigned by a subnet.
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// Number of SNAT IPs to be added. This parameter is used in conjunction with `SnatIps`. Note that if `Ip` is specified in `SnapIps`, this parameter is not available. It defaults to `1` and the upper limit is `10`.
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`
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

// Predefined struct for user
type CreateLoadBalancerSnatIpsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoadBalancerSnatIpsResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancerSnatIpsResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRuleRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Information of the new forwarding rule
	Rules []*RuleInput `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Information of the new forwarding rule
	Rules []*RuleInput `json:"Rules,omitnil,omitempty" name:"Rules"`
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

// Predefined struct for user
type CreateRuleResponseParams struct {
	// Array of unique IDs of created forwarding rules
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateTargetGroupRequestParams struct {
	// Target group name (up to 50 characters)
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// `vpcid` attribute of a target group. If this parameter is left empty, the default VPC will be used.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Default port of a target group, which can be used for subsequently added servers.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Real server bound to a target group
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type CreateTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// Target group name (up to 50 characters)
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// `vpcid` attribute of a target group. If this parameter is left empty, the default VPC will be used.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Default port of a target group, which can be used for subsequently added servers.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Real server bound to a target group
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
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

// Predefined struct for user
type CreateTargetGroupResponseParams struct {
	// ID generated after target group creation
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTargetGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateTargetGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateTopicRequestParams struct {
	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// The number of topic partitions, which changes as partitions are split or merged. Each log topic can have up to 50 partitions. If this parameter is not passed in, 1 partition will be created by default and up to 10 partitions are allowed to be created.
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// Log type. Valid values: ACCESS (access logs; default value) and HEALTH (health check logs).
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// Logset retention period (in days). Default: 30 days.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Log topic storage type. Valid values: `hot` (STANDARD storage); `cold` (IA storage). Default value: `hot`.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// The number of topic partitions, which changes as partitions are split or merged. Each log topic can have up to 50 partitions. If this parameter is not passed in, 1 partition will be created by default and up to 10 partitions are allowed to be created.
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// Log type. Valid values: ACCESS (access logs; default value) and HEALTH (health check logs).
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// Logset retention period (in days). Default: 30 days.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Log topic storage type. Valid values: `hot` (STANDARD storage); `cold` (IA storage). Default value: `hot`.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
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
	delete(f, "Period")
	delete(f, "StorageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
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
	LocalVpcId *string `json:"LocalVpcId,omitnil,omitempty" name:"LocalVpcId"`

	// VPC ID of the CVM or ENI instance
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// IP address of the CVM or ENI instance
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// VPC name of the CVM or ENI instance
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// ENI ID of the CVM instance
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`

	// ID of the CVM instance
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the CVM instance
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Region of the CVM or ENI instance
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// ID of the listener to be deleted
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// ID of the listener to be deleted
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
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

// Predefined struct for user
type DeleteListenerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteListenerResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLoadBalancerListenersRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of listener IDs to delete (20 IDs at most). If this parameter is left empty, all listeners of the CLB instance will be deleted.
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`
}

type DeleteLoadBalancerListenersRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of listener IDs to delete (20 IDs at most). If this parameter is left empty, all listeners of the CLB instance will be deleted.
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`
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

// Predefined struct for user
type DeleteLoadBalancerListenersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerListenersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerListenersResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLoadBalancerRequestParams struct {
	// Array of IDs of the CLB instances to be deleted. Array length limit: 20.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// Array of IDs of the CLB instances to be deleted. Array length limit: 20.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
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

// Predefined struct for user
type DeleteLoadBalancerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLoadBalancerSnatIpsRequestParams struct {
	// Unique ID of a CLB instance, e.g., lb-12345678.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of the SNAT IP addresses to be deleted
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`
}

type DeleteLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a CLB instance, e.g., lb-12345678.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of the SNAT IP addresses to be deleted
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`
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

// Predefined struct for user
type DeleteLoadBalancerSnatIpsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerSnatIpsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerSnatIpsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRewriteRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Source listener ID
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// Target listener ID
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// Redirection relationship between forwarding rules
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
}

type DeleteRewriteRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Source listener ID
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// Target listener ID
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// Redirection relationship between forwarding rules
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
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

// Predefined struct for user
type DeleteRewriteResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRewriteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRewriteResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRuleRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Array of IDs of the forwarding rules to be deleted
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`

	// The domain name associated with the forwarding rule to delete. If the rule is associated with multiple domain names, specify any one of them.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The forwarding path of the forwarding rule to delete.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Specifies a new default domain name for the listener. This field is used when the original default domain name is disabled. If there are multiple domain names, specify one of them.
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Array of IDs of the forwarding rules to be deleted
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`

	// The domain name associated with the forwarding rule to delete. If the rule is associated with multiple domain names, specify any one of them.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The forwarding path of the forwarding rule to delete.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Specifies a new default domain name for the listener. This field is used when the original default domain name is disabled. If there are multiple domain names, specify one of them.
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`
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

// Predefined struct for user
type DeleteRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteTargetGroupsRequestParams struct {
	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
}

type DeleteTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
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

// Predefined struct for user
type DeleteTargetGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTargetGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeregisterFunctionTargetsRequestParams struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of functions to be unbound
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// The ID of target forwarding rule. To unbind a function from an L7 forwarding rule, either `LocationId` or `Domain+Url` is required. 
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the target forwarding rule. It is ignored if `LocationId` is specified. 
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// URL of the target forwarding rule. It is ignored if `LocationId` is specified. 
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DeregisterFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of functions to be unbound
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// The ID of target forwarding rule. To unbind a function from an L7 forwarding rule, either `LocationId` or `Domain+Url` is required. 
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the target forwarding rule. It is ignored if `LocationId` is specified. 
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// URL of the target forwarding rule. It is ignored if `LocationId` is specified. 
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *DeregisterFunctionTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterFunctionTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "FunctionTargets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterFunctionTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterFunctionTargetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterFunctionTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterFunctionTargetsResponseParams `json:"Response"`
}

func (r *DeregisterFunctionTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterFunctionTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetGroupInstancesRequestParams struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Information of server to be unbound
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type DeregisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Information of server to be unbound
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
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

// Predefined struct for user
type DeregisterTargetGroupInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterTargetGroupInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeregisterTargetsFromClassicalLBRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of real server IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DeregisterTargetsFromClassicalLBRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of real server IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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

// Predefined struct for user
type DeregisterTargetsFromClassicalLBResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterTargetsFromClassicalLBResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterTargetsFromClassicalLBResponseParams `json:"Response"`
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

// Predefined struct for user
type DeregisterTargetsRequestParams struct {
	// CLB instance ID in the format of "lb-12345678"
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Listener ID in the format of "lbl-12345678"
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of real servers to be unbound. Array length limit: 20.
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Forwarding rule ID in the format of "loc-12345678". When unbinding a server from a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID in the format of "lb-12345678"
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Listener ID in the format of "lbl-12345678"
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of real servers to be unbound. Array length limit: 20.
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Forwarding rule ID in the format of "loc-12345678". When unbinding a server from a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
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

// Predefined struct for user
type DeregisterTargetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterTargetsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBlockIPListRequestParams struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Data offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of IPs to be returned. Default value: 100,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBlockIPListRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Data offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of IPs to be returned. Default value: 100,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeBlockIPListResponseParams struct {
	// Number of returned IPs
	BlockedIPCount *uint64 `json:"BlockedIPCount,omitnil,omitempty" name:"BlockedIPCount"`

	// Field for getting real client IP
	ClientIPField *string `json:"ClientIPField,omitnil,omitempty" name:"ClientIPField"`

	// List of IPs added to blocklist 12360
	BlockedIPList []*BlockedIP `json:"BlockedIPList,omitnil,omitempty" name:"BlockedIPList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBlockIPListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockIPListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBlockIPTaskRequestParams struct {
	// Async task ID returned by the `ModifyBlockIPList` API
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeBlockIPTaskRequest struct {
	*tchttp.BaseRequest
	
	// Async task ID returned by the `ModifyBlockIPList` API
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DescribeBlockIPTaskResponseParams struct {
	// 1: running; 2: failed; 6: succeeded
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBlockIPTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockIPTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClassicalLBByInstanceIdRequestParams struct {
	// List of real server IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeClassicalLBByInstanceIdRequest struct {
	*tchttp.BaseRequest
	
	// List of real server IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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

// Predefined struct for user
type DescribeClassicalLBByInstanceIdResponseParams struct {
	// CLB information list
	LoadBalancerInfoList []*ClassicalLoadBalancerInfo `json:"LoadBalancerInfoList,omitnil,omitempty" name:"LoadBalancerInfoList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBByInstanceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBByInstanceIdResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClassicalLBHealthStatusRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

type DescribeClassicalLBHealthStatusRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
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

// Predefined struct for user
type DescribeClassicalLBHealthStatusResponseParams struct {
	// List of real server health statuses
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HealthList []*ClassicalHealth `json:"HealthList,omitnil,omitempty" name:"HealthList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBHealthStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClassicalLBListenersRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of CLB listener IDs
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// CLB listening protocol. Valid values: TCP, UDP, HTTP, and HTTPS.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// CLB listening port. Value range: 1 - 65535.
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// Listener status. Valid values: 0 (creating) and 1 (running).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeClassicalLBListenersRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of CLB listener IDs
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// CLB listening protocol. Valid values: TCP, UDP, HTTP, and HTTPS.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// CLB listening port. Value range: 1 - 65535.
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// Listener status. Valid values: 0 (creating) and 1 (running).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
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

// Predefined struct for user
type DescribeClassicalLBListenersResponseParams struct {
	// Listener list
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Listeners []*ClassicalListener `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBListenersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClassicalLBTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type DescribeClassicalLBTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
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

// Predefined struct for user
type DescribeClassicalLBTargetsResponseParams struct {
	// Real server list
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Targets []*ClassicalTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBTargetsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClsLogSetRequestParams struct {

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

// Predefined struct for user
type DescribeClsLogSetResponseParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Health check logset ID
	HealthLogsetId *string `json:"HealthLogsetId,omitnil,omitempty" name:"HealthLogsetId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClsLogSetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClsLogSetResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCrossTargetsRequestParams struct {
	// Number of real server lists returned. Default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting offset of the real server list returned. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter conditions to query CVMs and ENIs
	// <li> `vpc-id` - String - Required: No - (Filter condition) Filter by VPC ID, such as "vpc-12345678".</li>
	// <li> `ip` - String - Required: No - (Filter condition) Filter by real server IP, such as "192.168.0.1".</li>
	// <li> `listener-id` - String - Required: No - (Filter condition) Filter by listener ID, such as "lbl-12345678".</li>
	// <li> `location-id` - String - Required: No - (Filter condition) Filter by forwarding rule ID of the layer-7 listener, such as "loc-12345678".</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCrossTargetsRequest struct {
	*tchttp.BaseRequest
	
	// Number of real server lists returned. Default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting offset of the real server list returned. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter conditions to query CVMs and ENIs
	// <li> `vpc-id` - String - Required: No - (Filter condition) Filter by VPC ID, such as "vpc-12345678".</li>
	// <li> `ip` - String - Required: No - (Filter condition) Filter by real server IP, such as "192.168.0.1".</li>
	// <li> `listener-id` - String - Required: No - (Filter condition) Filter by listener ID, such as "lbl-12345678".</li>
	// <li> `location-id` - String - Required: No - (Filter condition) Filter by forwarding rule ID of the layer-7 listener, such as "loc-12345678".</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeCrossTargetsResponseParams struct {
	// Total number of real server lists
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Real server list
	CrossTargetSet []*CrossTargets `json:"CrossTargetSet,omitnil,omitempty" name:"CrossTargetSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCrossTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossTargetsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCustomizedConfigAssociateListRequestParams struct {
	// Configuration ID.
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// Start position of the binding list. Default: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of binding lists to pull. Default: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searches for the domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeCustomizedConfigAssociateListRequest struct {
	*tchttp.BaseRequest
	
	// Configuration ID.
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// Start position of the binding list. Default: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of binding lists to pull. Default: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searches for the domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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

// Predefined struct for user
type DescribeCustomizedConfigAssociateListResponseParams struct {
	// List of bound resources
	BindList []*BindDetailItem `json:"BindList,omitnil,omitempty" name:"BindList"`

	// Total number of bound resources
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomizedConfigAssociateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizedConfigAssociateListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCustomizedConfigListRequestParams struct {
	// Configuration type. Valid values: `CLB` (CLB-specific configs), `SERVER` (domain name-specific configs), and `LOCATION` (forwarding rule-specific configs).
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// Pagination offset. Default: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page. Default: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies the name of configs to query. Fuzzy match is supported.
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// Configuration ID.
	UconfigIds []*string `json:"UconfigIds,omitnil,omitempty" name:"UconfigIds"`

	// The filters are:
	// <li> loadbalancer-id - String - Required: no - (filter) CLB instance ID, such as "lb-12345678". </li>
	// <li> vip - String - Required: no - (filter) CLB instance VIP, such as "1.1.1.1" and "2204::22:3". </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCustomizedConfigListRequest struct {
	*tchttp.BaseRequest
	
	// Configuration type. Valid values: `CLB` (CLB-specific configs), `SERVER` (domain name-specific configs), and `LOCATION` (forwarding rule-specific configs).
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// Pagination offset. Default: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page. Default: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies the name of configs to query. Fuzzy match is supported.
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// Configuration ID.
	UconfigIds []*string `json:"UconfigIds,omitnil,omitempty" name:"UconfigIds"`

	// The filters are:
	// <li> loadbalancer-id - String - Required: no - (filter) CLB instance ID, such as "lb-12345678". </li>
	// <li> vip - String - Required: no - (filter) CLB instance VIP, such as "1.1.1.1" and "2204::22:3". </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeCustomizedConfigListResponseParams struct {
	// Configuration list.
	ConfigList []*ConfigListItem `json:"ConfigList,omitnil,omitempty" name:"ConfigList"`

	// Number of configurations.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomizedConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizedConfigListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeIdleLoadBalancersRequestParams struct {
	// Data offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned CLB instances. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// CLB instance region
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
}

type DescribeIdleLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// Data offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned CLB instances. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// CLB instance region
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
}

func (r *DescribeIdleLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdleLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "LoadBalancerRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdleLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdleLoadBalancersResponseParams struct {
	// List of idle CLBs
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IdleLoadBalancers []*IdleLoadBalancer `json:"IdleLoadBalancers,omitnil,omitempty" name:"IdleLoadBalancers"`

	// Total number of idle CLB instances
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIdleLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdleLoadBalancersResponseParams `json:"Response"`
}

func (r *DescribeIdleLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdleLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLBListenersRequestParams struct {
	// List of private network IPs to be queried.
	Backends []*LbRsItem `json:"Backends,omitnil,omitempty" name:"Backends"`
}

type DescribeLBListenersRequest struct {
	*tchttp.BaseRequest
	
	// List of private network IPs to be queried.
	Backends []*LbRsItem `json:"Backends,omitnil,omitempty" name:"Backends"`
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

// Predefined struct for user
type DescribeLBListenersResponseParams struct {
	// Listener rule associated with the real server.
	LoadBalancers []*LBItem `json:"LoadBalancers,omitnil,omitempty" name:"LoadBalancers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLBListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLBListenersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeListenersRequestParams struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of CLB listener IDs to query (100 IDs at most).
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// Type of the listener protocols to be queried. Values: TCP`, `UDP`, `HTTP`, `HTTPS`, `TCP_SSL` and `QUIC`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port of the listeners to be queried
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of CLB listener IDs to query (100 IDs at most).
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// Type of the listener protocols to be queried. Values: TCP`, `UDP`, `HTTP`, `HTTPS`, `TCP_SSL` and `QUIC`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Port of the listeners to be queried
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
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

// Predefined struct for user
type DescribeListenersResponseParams struct {
	// Listener list
	Listeners []*Listener `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// Total number of listeners (with filters of port, protocol, and listener ID applied).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLoadBalancerListByCertIdRequestParams struct {
	// Server or client certificate ID
	CertIds []*string `json:"CertIds,omitnil,omitempty" name:"CertIds"`
}

type DescribeLoadBalancerListByCertIdRequest struct {
	*tchttp.BaseRequest
	
	// Server or client certificate ID
	CertIds []*string `json:"CertIds,omitnil,omitempty" name:"CertIds"`
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

// Predefined struct for user
type DescribeLoadBalancerListByCertIdResponseParams struct {
	// Certificate ID and list of CLB instances associated with it
	CertSet []*CertIdRelatedWithLoadBalancers `json:"CertSet,omitnil,omitempty" name:"CertSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerListByCertIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerListByCertIdResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLoadBalancerOverviewRequestParams struct {

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

// Predefined struct for user
type DescribeLoadBalancerOverviewResponseParams struct {
	// Total number of CLB instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Number of CLB instances that are running
	RunningCount *int64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// Number of CLB instances that are isolated
	IsolationCount *int64 `json:"IsolationCount,omitnil,omitempty" name:"IsolationCount"`

	// Number of CLB instances that are about to expire
	WillExpireCount *int64 `json:"WillExpireCount,omitnil,omitempty" name:"WillExpireCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerOverviewResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLoadBalancerTrafficRequestParams struct {
	// CLB instance region. If this parameter is not passed in, CLB instances in all regions will be returned.
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
}

type DescribeLoadBalancerTrafficRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance region. If this parameter is not passed in, CLB instances in all regions will be returned.
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
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

// Predefined struct for user
type DescribeLoadBalancerTrafficResponseParams struct {
	// Information of CLB instances sorted by outbound bandwidth from highest to lowest
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LoadBalancerTraffic []*LoadBalancerTraffic `json:"LoadBalancerTraffic,omitnil,omitempty" name:"LoadBalancerTraffic"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerTrafficResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerTrafficResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLoadBalancersDetailRequestParams struct {
	// Number of CLB instance lists returned. Default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting offset of the CLB instance list returned. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// List of fields. Only fields specified will be returned. If it’s left blank, `null` is returned. The fields `LoadBalancerId` and `LoadBalancerName` are added by default. For details about fields, see <a href="https://intl.cloud.tencent.com/document/api/214/30694?from_cn_redirect=1#LoadBalancerDetail">LoadBalancerDetail</a>.
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Target type. Valid values: NODE and GROUP. If the list of fields contains `TargetId`, `TargetAddress`, `TargetPort`, `TargetWeight` and other fields, `Target` of the target group or non-target group must be exported.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

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
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLoadBalancersDetailRequest struct {
	*tchttp.BaseRequest
	
	// Number of CLB instance lists returned. Default value: 20; maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting offset of the CLB instance list returned. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// List of fields. Only fields specified will be returned. If it’s left blank, `null` is returned. The fields `LoadBalancerId` and `LoadBalancerName` are added by default. For details about fields, see <a href="https://intl.cloud.tencent.com/document/api/214/30694?from_cn_redirect=1#LoadBalancerDetail">LoadBalancerDetail</a>.
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Target type. Valid values: NODE and GROUP. If the list of fields contains `TargetId`, `TargetAddress`, `TargetPort`, `TargetWeight` and other fields, `Target` of the target group or non-target group must be exported.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

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
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeLoadBalancersDetailResponseParams struct {
	// Total number of lists describing CLB instance details.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of CLB instance details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerDetailSet []*LoadBalancerDetail `json:"LoadBalancerDetailSet,omitnil,omitempty" name:"LoadBalancerDetailSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancersDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancersDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLoadBalancersRequestParams struct {
	// CLB instance IDs. There can be up to 20 IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// CLB instance type. 1: generic CLB instance; 0: classic CLB instance
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// The domain name that Tencent Cloud assigned for the CLB instance.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// VIP address of a CLB instance (there can be multiple addresses)
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// Public IPs of the backend services bound with the load balancer. Only the public IPs of CVMs are supported now.
	BackendPublicIps []*string `json:"BackendPublicIps,omitnil,omitempty" name:"BackendPublicIps"`

	// Private IPs of the backend services bound with the load balancer. Only the private IPs of CVMs are supported now.
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitnil,omitempty" name:"BackendPrivateIps"`

	// Data offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned CLB instances. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sort by parameter. Value range: LoadBalancerName, CreateTime, Domain, LoadBalancerType.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 1: reverse; 0: sequential. Default value: reverse by creation time |
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Search field which fuzzy matches name, domain name, or VIP.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the DescribeProject API.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether a CLB instance is bound to a real server. 0: no; 1: yes; -1: query all.
	WithRs *int64 `json:"WithRs,omitnil,omitempty" name:"WithRs"`

	// VPC where a CLB instance resides, such as vpc-bhqkbhdx.
	// Basic network does not support queries by VpcId.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Security group ID, e.g., `sg-m1cc****`.
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Primary AZ ID, e.g., `100001` (Guangzhou Zone 1).
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// Each request can have up to 10 `Filters` and 100 `Filter.Values`. Detailed filter conditions:
	// <li> internet-charge-type - Type: String - Required: No - Filter by CLB network billing mode, including `TRAFFIC_POSTPAID_BY_HOUR`</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance IDs. There can be up to 20 IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// CLB instance type. 1: generic CLB instance; 0: classic CLB instance
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// The domain name that Tencent Cloud assigned for the CLB instance.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// VIP address of a CLB instance (there can be multiple addresses)
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// Public IPs of the backend services bound with the load balancer. Only the public IPs of CVMs are supported now.
	BackendPublicIps []*string `json:"BackendPublicIps,omitnil,omitempty" name:"BackendPublicIps"`

	// Private IPs of the backend services bound with the load balancer. Only the private IPs of CVMs are supported now.
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitnil,omitempty" name:"BackendPrivateIps"`

	// Data offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned CLB instances. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sort by parameter. Value range: LoadBalancerName, CreateTime, Domain, LoadBalancerType.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 1: reverse; 0: sequential. Default value: reverse by creation time |
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Search field which fuzzy matches name, domain name, or VIP.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// ID of the project to which a CLB instance belongs, which can be obtained through the DescribeProject API.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether a CLB instance is bound to a real server. 0: no; 1: yes; -1: query all.
	WithRs *int64 `json:"WithRs,omitnil,omitempty" name:"WithRs"`

	// VPC where a CLB instance resides, such as vpc-bhqkbhdx.
	// Basic network does not support queries by VpcId.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Security group ID, e.g., `sg-m1cc****`.
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Primary AZ ID, e.g., `100001` (Guangzhou Zone 1).
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// Each request can have up to 10 `Filters` and 100 `Filter.Values`. Detailed filter conditions:
	// <li> internet-charge-type - Type: String - Required: No - Filter by CLB network billing mode, including `TRAFFIC_POSTPAID_BY_HOUR`</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeLoadBalancersResponseParams struct {
	// Total number of CLB instances that meet the filter criteria. This value is independent of the Limit in the input parameter.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Array of returned CLB instances.
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitnil,omitempty" name:"LoadBalancerSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeQuotaRequestParams struct {

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

// Predefined struct for user
type DescribeQuotaResponseParams struct {
	// Quota list
	QuotaSet []*Quota `json:"QuotaSet,omitnil,omitempty" name:"QuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeResourcesRequestParams struct {
	// Number of returned AZ resources. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting offset of the returned AZ resource list. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter to query the list of AZ resources as detailed below:
	// <li> `zone` - String - Optional - Filter by AZ, such as "ap-guangzhou-1".</li>
	// <li> `isp` -- String - Optional - Filter by the ISP. Values: `BGP`, `CMCC`, `CUCC` and `CTCC`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned AZ resources. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting offset of the returned AZ resource list. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter to query the list of AZ resources as detailed below:
	// <li> `zone` - String - Optional - Filter by AZ, such as "ap-guangzhou-1".</li>
	// <li> `isp` -- String - Optional - Filter by the ISP. Values: `BGP`, `CMCC`, `CUCC` and `CTCC`.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesResponseParams struct {
	// List of resources supported by the AZ
	ZoneResourceSet []*ZoneResource `json:"ZoneResourceSet,omitnil,omitempty" name:"ZoneResourceSet"`

	// Number of entries in the AZ resource list.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesResponseParams `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRewriteRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of CLB listener IDs
	SourceListenerIds []*string `json:"SourceListenerIds,omitnil,omitempty" name:"SourceListenerIds"`

	// Array of CLB forwarding rule IDs
	SourceLocationIds []*string `json:"SourceLocationIds,omitnil,omitempty" name:"SourceLocationIds"`
}

type DescribeRewriteRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of CLB listener IDs
	SourceListenerIds []*string `json:"SourceListenerIds,omitnil,omitempty" name:"SourceListenerIds"`

	// Array of CLB forwarding rule IDs
	SourceLocationIds []*string `json:"SourceLocationIds,omitnil,omitempty" name:"SourceLocationIds"`
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

// Predefined struct for user
type DescribeRewriteResponseParams struct {
	// Array of redirection forwarding rules. If there are no redirection rules, an empty array will be returned.
	RewriteSet []*RuleOutput `json:"RewriteSet,omitnil,omitempty" name:"RewriteSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRewriteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRewriteResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTargetGroupInstancesRequestParams struct {
	// Filter. Currently, only filtering by `TargetGroupId`, `BindIP`, or `InstanceId` is supported.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of displayed results. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Display offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Filter. Currently, only filtering by `TargetGroupId`, `BindIP`, or `InstanceId` is supported.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of displayed results. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Display offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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

// Predefined struct for user
type DescribeTargetGroupInstancesResponseParams struct {
	// Number of results returned for the current query
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of the bound server
	TargetGroupInstanceSet []*TargetGroupBackend `json:"TargetGroupInstanceSet,omitnil,omitempty" name:"TargetGroupInstanceSet"`

	// The actual total number of bound instances, which is not affected by the setting of `Limit`, `Offset` and the CAM permissions.
	RealCount *uint64 `json:"RealCount,omitnil,omitempty" name:"RealCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTargetGroupListRequestParams struct {
	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: `TargetGroupVpcId` and `TargetGroupName`. Target group ID will be used first.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit of the number of displayed results. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTargetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID array
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: `TargetGroupVpcId` and `TargetGroupName`. Target group ID will be used first.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit of the number of displayed results. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeTargetGroupListResponseParams struct {
	// Number of displayed results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information set of displayed target groups
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitnil,omitempty" name:"TargetGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTargetGroupsRequestParams struct {
	// Target group ID, which is exclusive of `Filters`.
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// Limit of the number of displayed results. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: `TargetGroupVpcId` and `TargetGroupName`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID, which is exclusive of `Filters`.
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// Limit of the number of displayed results. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Starting display offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter array, which is exclusive of `TargetGroupIds`. Valid values: `TargetGroupVpcId` and `TargetGroupName`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeTargetGroupsResponseParams struct {
	// Number of displayed results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information set of displayed target groups
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitnil,omitempty" name:"TargetGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTargetHealthRequestParams struct {
	// List of IDs of CLB instances to be queried
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of CLB instances to be queried
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
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

// Predefined struct for user
type DescribeTargetHealthResponseParams struct {
	// CLB instance list
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitnil,omitempty" name:"LoadBalancers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetHealthResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTargetsRequestParams struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of listener IDs (20 IDs at most).
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// Listener protocol type
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Query the list of backend services associated with a load balancer
	// <li> `location-id` - String - Optional - Rule ID, such as "loc-12345678".</li>
	// <li> `private-ip-address` - String - Optional - Backend service private IP, such as `172.16.1.1`</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// List of listener IDs (20 IDs at most).
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// Listener protocol type
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Query the list of backend services associated with a load balancer
	// <li> `location-id` - String - Optional - Rule ID, such as "loc-12345678".</li>
	// <li> `private-ip-address` - String - Optional - Backend service private IP, such as `172.16.1.1`</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetsResponseParams struct {
	// Information of real servers bound to the listener
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Listeners []*ListenerBackend `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// Request ID, i.e., the RequestId parameter returned by the API.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Order ID.
	// Note: Either `TaskId` or `DealName` is required.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Request ID, i.e., the RequestId parameter returned by the API.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Order ID.
	// Note: Either `TaskId` or `DealName` is required.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
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

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// Current status of a task. Value range: 0 (succeeded), 1 (failed), 2 (in progress).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Array of unique CLB instance IDs.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DisassociateTargetGroupsRequestParams struct {
	// Array of rules to be unbound
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

type DisassociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Array of rules to be unbound
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
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

// Predefined struct for user
type DisassociateTargetGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateTargetGroupsResponseParams `json:"Response"`
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
	L4Clusters []*ClusterItem `json:"L4Clusters,omitnil,omitempty" name:"L4Clusters"`

	// Layer-7 dedicated cluster list
	// Note: this field may return null, indicating that no valid values can be obtained.
	L7Clusters []*ClusterItem `json:"L7Clusters,omitnil,omitempty" name:"L7Clusters"`

	// vpcgw cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClassicalCluster *ClusterItem `json:"ClassicalCluster,omitnil,omitempty" name:"ClassicalCluster"`
}

type ExtraInfo struct {
	// Whether to enable VIP direct connection
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZhiTong *bool `json:"ZhiTong,omitnil,omitempty" name:"ZhiTong"`

	// TgwGroup name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TgwGroupName *string `json:"TgwGroupName,omitnil,omitempty" name:"TgwGroupName"`
}

type Filter struct {
	// Filter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value array
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FunctionInfo struct {
	// Function namespace
	FunctionNamespace *string `json:"FunctionNamespace,omitnil,omitempty" name:"FunctionNamespace"`

	// Function name
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// Function version name or alias
	FunctionQualifier *string `json:"FunctionQualifier,omitnil,omitempty" name:"FunctionQualifier"`

	// Function qualifier type. Values: `VERSION`, `ALIAS`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FunctionQualifierType *string `json:"FunctionQualifierType,omitnil,omitempty" name:"FunctionQualifierType"`
}

type FunctionTarget struct {
	// SCF related information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Function *FunctionInfo `json:"Function,omitnil,omitempty" name:"Function"`

	// Weight
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type HealthCheck struct {
	// Whether to enable health check. 1: enable; 0: disable.
	HealthSwitch *int64 `json:"HealthSwitch,omitnil,omitempty" name:"HealthSwitch"`

	// Health check response timeout period in seconds (applicable only to layer-4 listeners). Value range: 2-60. Default value: 2. This parameter should be less than the check interval.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeOut *int64 `json:"TimeOut,omitnil,omitempty" name:"TimeOut"`

	// Health check probing interval period. It defaults to `5`. For IPv4 CLB instances, the range is 2-300. For IPv6 CLB instances, the range is 5-300. Unit: second
	// Note: For some IPv4 CLB instances created long ago, the range is 5-300.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntervalTime *int64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// Health threshold. Default value: 3, indicating that if a forward is found healthy three consecutive times, it is considered to be normal. Value range: 2-10
	// Note: This field may return null, indicating that no valid values can be obtained.
	HealthNum *int64 `json:"HealthNum,omitnil,omitempty" name:"HealthNum"`

	// Unhealthy threshold. Default value: 3, indicating that if a forward is found unhealthy three consecutive times, it is considered to be exceptional. Value range: 2-10
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnHealthNum *int64 `json:"UnHealthNum,omitnil,omitempty" name:"UnHealthNum"`

	// Health check status code (applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners). Value range: 1-31. Default value: 31.
	// `1`: Returns code 1xx for healthy status. `2`: Returns code 2xx for healthy status. `4`: Returns code 3xx for healthy status. `8`: Returns code 4xx for healthy status. `16`: Returns code 5xx for healthy status. If you want multiple return codes to represent healthy, sum up the corresponding values. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HttpCode *int64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// Health check path (applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners).
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpCheckPath *string `json:"HttpCheckPath,omitnil,omitempty" name:"HttpCheckPath"`

	// Health check domain name. It's applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners. It's required for HTTP health check of TCP listeners.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpCheckDomain *string `json:"HttpCheckDomain,omitnil,omitempty" name:"HttpCheckDomain"`

	// Health check method (applicable only to HTTP/HTTPS forwarding rules and HTTP health checks of TCP listeners). Value range: HEAD, GET. Default value: HEAD.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpCheckMethod *string `json:"HttpCheckMethod,omitnil,omitempty" name:"HttpCheckMethod"`

	// Health check port (a custom check parameter), which is the port of the real server by default. Unless you want to specify a port, it is recommended to leave it empty. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckPort *int64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// Health check protocol (a custom check parameter), which is required if the value of CheckType is CUSTOM. This parameter represents the input format of the health check. Value range: HEX, TEXT. If the value is HEX, the characters of SendContext and RecvContext can only be selected from 0123456789ABCDEF and the length must be an even number. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// Health check protocol (a custom check parameter), which is required if the value of CheckType is CUSTOM. This parameter represents the content of the request sent by the health check. Only ASCII visible characters are allowed, and the maximum length is 500. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	SendContext *string `json:"SendContext,omitnil,omitempty" name:"SendContext"`

	// Health check protocol (a custom check parameter), which is required if the value of CheckType is CUSTOM. This parameter represents the result returned by the health check. Only ASCII visible characters are allowed, and the maximum length is 500. (Applicable only to TCP/UDP listeners.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecvContext *string `json:"RecvContext,omitnil,omitempty" name:"RecvContext"`

	// Health check protocol. Values: `TCP`, `HTTP`, `HTTPS`, `GRPC`, `PING`, and `CUSTOM`. UDP listeners support `PING`/`CUSTOM`. TCP listener support `TCP`/`HTTP`/`CUSTOM`. TCP_SSL and QUIC listeners support `TCP`/`HTTP`. HTTP rules support `HTTP`/`GRPC. HTTPS rules support `HTTP`/`HTTPS`/`GRPC`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// HTTP version. HTTP version of the backend service. Values: `HTTP/1.0`, HTTP/1.1`. It is only applicable to TCP listeners, and is required when `CheckType`=`HTTP`. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`

	// Specifies the type of health check source IP. `0` (default): CLB VIP. `1`: 100.64 IP range.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceIpType *int64 `json:"SourceIpType,omitnil,omitempty" name:"SourceIpType"`

	// GRPC health check status code, which is only applicable to rules with GRPC as the backend forwarding protocol. It can be a single number (such as `20`), multiple numbers (such as `20,25`) or a range (such as `0-99`). The default value is `12`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExtendedCode *string `json:"ExtendedCode,omitnil,omitempty" name:"ExtendedCode"`
}

type IdleLoadBalancer struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// CLB instance region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CLB instance VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// The reason why the load balancer is considered idle. `NO_RULES`: No rules configured. `NO_RS`: The rules are not associated with servers.
	IdleReason *string `json:"IdleReason,omitnil,omitempty" name:"IdleReason"`

	// CLB instance status, including:
	// `0`: Creating; `1`: Running.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// CLB type. Value range: `1` (CLB); `0` (classic CLB).
	Forward *uint64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// The load balancing hostname.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

// Predefined struct for user
type InquiryPriceCreateLoadBalancerRequestParams struct {
	// Network type of the CLB to query. `OPEN`: Public network; `INTERNAL`: Private network is intranet type
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// The billing mode to query. `POSTPAID`:Pay as you go
	LoadBalancerChargeType *string `json:"LoadBalancerChargeType,omitnil,omitempty" name:"LoadBalancerChargeType"`

	// Reserved field
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`

	// The network billing mode to query 
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Number of CLB instances to query. Default value: 1.
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Availability zone in the format of "ap-guangzhou-1"
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// To query the price of monthly subscribed LCU-supported instances, specify the instance specification in this parameter, such as `clb.c3.small`. For PAYG instances, use `SLA`.
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// IP version. Valid values: `IPV4` (default), `IPV6` (IPV6 NAT64 version) or `IPv6FullChain` (IPv6 version). 
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// ISP of VIP. Values: `CMCC` (China Mobile), `CUCC` (China Unicom) and `CTCC` (China Telecom). You need to activate static single-line IPs. This feature is in beta and is only available in Guangzhou, Shanghai, Nanjing, Jinan, Hangzhou, Fuzhou, Beijing, Shijiazhuang, Wuhan, Changsha, Chengdu and Chongqing regions. To try it out, please contact your sales rep. If it's specified, the network billing mode must be `BANDWIDTH_PACKAGE`. If it's not specified, BGP is used by default. To query ISPs supported in a region, please use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1). 
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`
}

type InquiryPriceCreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// Network type of the CLB to query. `OPEN`: Public network; `INTERNAL`: Private network is intranet type
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// The billing mode to query. `POSTPAID`:Pay as you go
	LoadBalancerChargeType *string `json:"LoadBalancerChargeType,omitnil,omitempty" name:"LoadBalancerChargeType"`

	// Reserved field
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`

	// The network billing mode to query 
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Number of CLB instances to query. Default value: 1.
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Availability zone in the format of "ap-guangzhou-1"
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// To query the price of monthly subscribed LCU-supported instances, specify the instance specification in this parameter, such as `clb.c3.small`. For PAYG instances, use `SLA`.
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// IP version. Valid values: `IPV4` (default), `IPV6` (IPV6 NAT64 version) or `IPv6FullChain` (IPv6 version). 
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// ISP of VIP. Values: `CMCC` (China Mobile), `CUCC` (China Unicom) and `CTCC` (China Telecom). You need to activate static single-line IPs. This feature is in beta and is only available in Guangzhou, Shanghai, Nanjing, Jinan, Hangzhou, Fuzhou, Beijing, Shijiazhuang, Wuhan, Changsha, Chengdu and Chongqing regions. To try it out, please contact your sales rep. If it's specified, the network billing mode must be `BANDWIDTH_PACKAGE`. If it's not specified, BGP is used by default. To query ISPs supported in a region, please use [DescribeResources](https://intl.cloud.tencent.com/document/api/214/70213?from_cn_redirect=1). 
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`
}

func (r *InquiryPriceCreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerType")
	delete(f, "LoadBalancerChargeType")
	delete(f, "LoadBalancerChargePrepaid")
	delete(f, "InternetAccessible")
	delete(f, "GoodsNum")
	delete(f, "ZoneId")
	delete(f, "SlaType")
	delete(f, "AddressIPVersion")
	delete(f, "VipIsp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateLoadBalancerResponseParams struct {
	// Price of the instance with the specified configurations.
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceModifyLoadBalancerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// New bandwidth bandwidth specification
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`
}

type InquiryPriceModifyLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// New bandwidth bandwidth specification
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`
}

func (r *InquiryPriceModifyLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceModifyLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "InternetAccessible")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceModifyLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceModifyLoadBalancerResponseParams struct {
	// Pricing information
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceModifyLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceModifyLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceModifyLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceModifyLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRefundLoadBalancerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type InquiryPriceRefundLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

func (r *InquiryPriceRefundLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRefundLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRefundLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRefundLoadBalancerResponseParams struct {
	// Price of the instance with the specified configurations.
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRefundLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRefundLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceRefundLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRefundLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewLoadBalancerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Renewal period
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`
}

type InquiryPriceRenewLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Renewal period
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`
}

func (r *InquiryPriceRenewLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewLoadBalancerResponseParams struct {
	// Price to renew
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternetAccessible struct {
	// TRAFFIC_POSTPAID_BY_HOUR: hourly pay-as-you-go by traffic; BANDWIDTH_POSTPAID_BY_HOUR: hourly pay-as-you-go by bandwidth;
	// BANDWIDTH_PACKAGE: billed by bandwidth package (currently, this method is supported only if the ISP is specified)
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// Maximum outgoing bandwidth in Mbps. It works on LCU-supported instances on private networks and all instances on public networks.
	// - For shared and dedicated CLB instances on public networks, the range is 1Mbps-2048Mbps.
	// - For all LCU-supported CLB instances:
	//   - It defaults to General LCU-supported instance. SLA corresponds to Super Large 1, and the range of maximum outgoing bandwidth is 1 Mbps - 10240 Mbps.
	//   - If you have enabled Super Large specification, the range of maximum outgoing bandwidth is 1 Mbps - 61440 Mbps Super Large LCU-supported specification is in beta now. To join the beta, [submit a ticket](https://console.cloud.tencent.com/workorder/category).
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Bandwidth package type, such as SINGLEISP
	// Note: This field may return null, indicating that no valid values can be obtained.
	BandwidthpkgSubType *string `json:"BandwidthpkgSubType,omitnil,omitempty" name:"BandwidthpkgSubType"`
}

type ItemPrice struct {
	// PAYG unit price, in USD.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Subsequent billing unit. Value Range: 
	// `HOUR`: Calculate the cost by hour. It's available when "InternetChargeType=POSTPAID_BY_HOUR".
	// `GB`: Calculate the cost by traffic in GB. It's available when "InternetChargeType=TRAFFIC_POSTPAID_BY_HOUR".
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// Reserved field
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Reserved field
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// Discount unit price of a pay-as-you-go instance, in USD.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// Discount. For example, 20.0 indicates 80% off.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`
}

type LBChargePrepaid struct {
	// Renewal type. AUTO_RENEW: automatic renewal; MANUAL_RENEW: manual renewal
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Cycle, indicating the number of months (reserved field)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type LBItem struct {
	// String ID of the CLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// VIP of the CLB instance.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Listener rule.
	Listeners []*ListenerItem `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// Region of the CLB instance
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type LbRsItem struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Private network IP to be queried, which can be of the CVM or ENI.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`
}

type LbRsTargets struct {
	// Private network IP type, which can be `cvm` or `eni`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Private network IP of the real server.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Port bound to the real server.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// VPC ID of the real server.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Weight of the real server.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type Listener struct {
	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Listener protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Information of certificates bound to the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
	Certificate *CertificateOutput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Health check information of the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Request scheduling method
	// Note: This field may return null, indicating that no valid values can be obtained.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Session persistence time
	// Note: This field may return null, indicating that no valid values can be obtained.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Whether to enable SNI. `1`: Enable; `0`: Do not enable. This parameter is only meaningful for HTTPS listeners.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// All forwarding rules under a listener (this parameter is meaningful only for HTTP/HTTPS listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*RuleOutput `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Listener name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Listener creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// End port of a port range
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndPort *int64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`

	// Real server type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Basic information of a bound target group. This field will be returned when a target group is bound to a listener.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitnil,omitempty" name:"TargetGroup"`

	// Session persistence type. Valid values: Normal: the default session persistence type; QUIC_CID: session persistence by QUIC connection ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// Whether a persistent connection is enabled (1: enabled; 0: disabled). This parameter can only be configured in HTTP/HTTPS listeners.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// Only the NAT64 CLB TCP listeners are supported.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Toa *bool `json:"Toa,omitnil,omitempty" name:"Toa"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// Attribute of listener
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AttrFlags []*string `json:"AttrFlags,omitnil,omitempty" name:"AttrFlags"`

	// List of bound target groups
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TargetGroupList []*BasicTargetGroupInfo `json:"TargetGroupList,omitnil,omitempty" name:"TargetGroupList"`

	// Maximum number of concurrent listener connections. If it’s set to `-1`, the listener speed is not limited. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// Maximum number of new listener connections. If it’s set to `-1`, the listener speed is not limited. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// Connection idle timeout period (in seconds). It’s only available to TCP listeners. Value range: 300-900 for shared and dedicated instances; 300-2000 for LCU-supported CLB instances. It defaults to 900.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`
}

type ListenerBackend struct {
	// Listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Listener protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Information of rules under a listener (applicable only to HTTP/HTTPS listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*RuleTargets `json:"Rules,omitnil,omitempty" name:"Rules"`

	// List of real servers bound to a listener (applicable only to TCP/UDP/TCP_SSL listeners)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets []*Backend `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Ending port in port range if port range is supported; 0 if port range is not supported
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndPort *int64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`
}

type ListenerHealth struct {
	// Listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Listener name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Listener protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// List of forwarding rules of the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*RuleHealth `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ListenerItem struct {
	// Listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Listener protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Bound rule.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Rules []*RulesItems `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Object bound to the layer-4 listener.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Targets []*LbRsTargets `json:"Targets,omitnil,omitempty" name:"Targets"`

	// End port of the listener.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndPort *int64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`
}

type LoadBalancer struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// CLB instance network type:
	// OPEN: public network; INTERNAL: private network.
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// CLB type identifier. Value range: 1 (CLB); 0 (classic CLB).
	Forward *uint64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// Domain name of the CLB instance. It is only available for public classic CLBs. This parameter will be discontinued soon. Please use `LoadBalancerDomain` instead.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// List of VIPs of a CLB instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// CLB instance status, including:
	// 0: creating; 1: running.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// CLB instance creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last status change time of a CLB instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusTime *string `json:"StatusTime,omitnil,omitempty" name:"StatusTime"`

	// ID of the project to which a CLB instance belongs. 0: default project.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Protective CLB identifier. Value range: 1 (protective), 0 (non-protective).
	// Note: This field may return null, indicating that no valid values can be obtained.
	OpenBgp *uint64 `json:"OpenBgp,omitnil,omitempty" name:"OpenBgp"`

	// SNAT is enabled for all private network classic CLB created before December 2016.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Snat *bool `json:"Snat,omitnil,omitempty" name:"Snat"`

	// 0: not isolated; 1: isolated.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Isolation *uint64 `json:"Isolation,omitnil,omitempty" name:"Isolation"`

	// Log information. Only the public network CLB that have HTTP or HTTPS listeners can generate logs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// Subnet where a CLB instance resides (meaningful only for private network VPC CLB)
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// CLB instance tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Security group of a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecureGroups []*string `json:"SecureGroups,omitnil,omitempty" name:"SecureGroups"`

	// Basic information of a backend server bound to a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitnil,omitempty" name:"TargetRegionInfo"`

	// Anycast CLB publishing region. For non-anycast CLB, this field returns an empty string.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnycastZone *string `json:"AnycastZone,omitnil,omitempty" name:"AnycastZone"`

	// IP version. Valid values: ipv4, ipv6
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// VPC ID in a numeric form
	// Note: This field may return null, indicating that no valid values can be obtained.
	NumericalVpcId *uint64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`

	// ISP to which a CLB IP address belongs
	// Note: This field may return null, indicating that no valid values can be obtained.
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// Primary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterZone *ZoneInfo `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// Secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitnil,omitempty" name:"BackupZoneSet"`

	// CLB instance isolation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// CLB instance expiration time, which takes effect only for prepaid instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Billing mode of CLB instance. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay as you go).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// CLB instance network attributes
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitnil,omitempty" name:"NetworkAttributes"`

	// Prepaid billing attributes of a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitnil,omitempty" name:"PrepaidAttributes"`

	// Logset ID of CLB Log Service (CLS)
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// Log topic ID of CLB Log Service (CLS)
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// IPv6 address of a CLB instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddressIPv6 *string `json:"AddressIPv6,omitnil,omitempty" name:"AddressIPv6"`

	// Reserved field which can be ignored generally.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// Whether an Anti-DDoS Pro instance can be bound
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDDos *bool `json:"IsDDos,omitnil,omitempty" name:"IsDDos"`

	// Custom configuration ID at the CLB instance level
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Whether a real server opens the traffic from a CLB instance to the internet
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// Private network dedicated cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// This field is meaningful only when the IP address version is `ipv6`. Valid values: IPv6Nat64, IPv6FullChain
	// Note: this field may return null, indicating that no valid values can be obtained.
	IPv6Mode *string `json:"IPv6Mode,omitnil,omitempty" name:"IPv6Mode"`

	// Whether to enable SnatPro.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// `SnatIp` list after SnatPro load balancing is enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// Specification of the LCU-supported instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// Whether VIP is blocked
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsBlock *bool `json:"IsBlock,omitnil,omitempty" name:"IsBlock"`

	// Time blocked or unblocked
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsBlockTime *string `json:"IsBlockTime,omitnil,omitempty" name:"IsBlockTime"`

	// Whether the IP type is the local BGP
	LocalBgp *bool `json:"LocalBgp,omitnil,omitempty" name:"LocalBgp"`

	// Dedicated layer-7 tag.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// If the layer-7 listener of an IPv6FullChain CLB instance is enabled, the CLB instance can be bound with an IPv4 and an IPv6 CVM instance simultaneously.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MixIpTarget *bool `json:"MixIpTarget,omitnil,omitempty" name:"MixIpTarget"`

	// Availability zone of a VPC-based private network CLB instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Whether it is an NFV CLB instance. No returned information: no; l7nfv: yes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NfvInfo *string `json:"NfvInfo,omitnil,omitempty" name:"NfvInfo"`

	// Health check logset ID of CLB CLS
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	HealthLogSetId *string `json:"HealthLogSetId,omitnil,omitempty" name:"HealthLogSetId"`

	// Health check log topic ID of CLB CLS
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	HealthLogTopicId *string `json:"HealthLogTopicId,omitnil,omitempty" name:"HealthLogTopicId"`

	// Cluster ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// CLB attribute
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AttributeFlags []*string `json:"AttributeFlags,omitnil,omitempty" name:"AttributeFlags"`

	// Domain name of the CLB instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`

	// Network egress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`
}

type LoadBalancerDetail struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// CLB instance network type:
	// Public: public network; Private: private network.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// CLB instance status, including:
	// 0: creating; 1: running.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// CLB instance VIP.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// IPv6 VIP address of the CLB instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIPv6 *string `json:"AddressIPv6,omitnil,omitempty" name:"AddressIPv6"`

	// IP version of the CLB instance. Valid values: IPv4, IPv6.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// IPv6 address type of the CLB instance. Valid values: IPv6Nat64, IPv6FullChain.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IPv6Mode *string `json:"IPv6Mode,omitnil,omitempty" name:"IPv6Mode"`

	// Availability zone where the CLB instance resides.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// ISP to which the CLB IP address belongs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddressIsp *string `json:"AddressIsp,omitnil,omitempty" name:"AddressIsp"`

	// ID of the VPC instance to which the CLB instance belongs.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the project to which the CLB instance belongs. 0: default project.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// CLB instance creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// CLB instance billing mode.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// CLB instance network attribute.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitnil,omitempty" name:"NetworkAttributes"`

	// Pay-as-you-go attribute of the CLB instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitnil,omitempty" name:"PrepaidAttributes"`

	// Reserved field, which can be ignored generally.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// Custom configuration IDs of CLB instances. Multiple IDs must be separated by commas (,).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// CLB instance tag information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CLB listener ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Listener protocol.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listener port.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Forwarding rule ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Forwarding rule path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// ID of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// Address of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// Listening port of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetPort *uint64 `json:"TargetPort,omitnil,omitempty" name:"TargetPort"`

	// Forwarding weight of target real servers.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetWeight *uint64 `json:"TargetWeight,omitnil,omitempty" name:"TargetWeight"`

	// 0: not isolated; 1: isolated.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Isolation *uint64 `json:"Isolation,omitnil,omitempty" name:"Isolation"`

	// List of the security groups bound to the CLB instance.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Whether the CLB instance is billed by IP.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LoadBalancerPassToTarget *uint64 `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// Health status of the target real server.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TargetHealth *string `json:"TargetHealth,omitnil,omitempty" name:"TargetHealth"`

	// List o domain names associated with the forwarding rule
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Domains *string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// The secondary zone of multi-AZ CLB instance
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SlaveZone []*string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// The AZ of private CLB instance. This is only available for beta users.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Whether to enable SNI. `1`: Enable; `0`: Do not enable. This parameter is only meaningful for HTTPS listeners.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// Domain name of the CLB instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`

	// Network egress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`
}

type LoadBalancerHealth struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// List of listeners
	// Note: This field may return null, indicating that no valid values can be obtained.
	Listeners []*ListenerHealth `json:"Listeners,omitnil,omitempty" name:"Listeners"`
}

type LoadBalancerTraffic struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// CLB instance region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CLB instance VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Maximum outbound bandwidth in Mbps
	OutBandwidth *float64 `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// CLB domain name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

// Predefined struct for user
type ManualRewriteRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Source listener ID
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// Target listener ID
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// Redirection relationship between forwarding rules
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
}

type ManualRewriteRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Source listener ID
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// Target listener ID
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// Redirection relationship between forwarding rules
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
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

// Predefined struct for user
type ManualRewriteResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManualRewriteResponse struct {
	*tchttp.BaseResponse
	Response *ManualRewriteResponseParams `json:"Response"`
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

// Predefined struct for user
type MigrateClassicalLoadBalancersRequestParams struct {
	// Array of classic CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Exclusive cluster information
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`
}

type MigrateClassicalLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// Array of classic CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Exclusive cluster information
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`
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

// Predefined struct for user
type MigrateClassicalLoadBalancersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MigrateClassicalLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *MigrateClassicalLoadBalancersResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyBlockIPListRequestParams struct {
	// CLB instance ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Operation type. Valid values:
	// <li> add_customized_field (sets header initially to enable the blocklist feature)</li>
	// <li> set_customized_field (modifies header)</li>
	// <li> del_customized_field (deletes header)</li>
	// <li> add_blocked (adds IPs to blocklist)</li>
	// <li> del_blocked (deletes IPs from blocklist)</li>
	// <li> flush_blocked (clears blocklist)</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Header field that stores real client IPs
	ClientIPField *string `json:"ClientIPField,omitnil,omitempty" name:"ClientIPField"`

	// List of blocked IPs. The array can contain up to 200,000 entries in one operation.
	BlockIPList []*string `json:"BlockIPList,omitnil,omitempty" name:"BlockIPList"`

	// Expiration time in seconds. Default value: 3600
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// IP adding policy. Valid value: fifo (if a blocklist is full, new IPs added to the blocklist will adopt the first-in first-out policy)
	AddStrategy *string `json:"AddStrategy,omitnil,omitempty" name:"AddStrategy"`
}

type ModifyBlockIPListRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Operation type. Valid values:
	// <li> add_customized_field (sets header initially to enable the blocklist feature)</li>
	// <li> set_customized_field (modifies header)</li>
	// <li> del_customized_field (deletes header)</li>
	// <li> add_blocked (adds IPs to blocklist)</li>
	// <li> del_blocked (deletes IPs from blocklist)</li>
	// <li> flush_blocked (clears blocklist)</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Header field that stores real client IPs
	ClientIPField *string `json:"ClientIPField,omitnil,omitempty" name:"ClientIPField"`

	// List of blocked IPs. The array can contain up to 200,000 entries in one operation.
	BlockIPList []*string `json:"BlockIPList,omitnil,omitempty" name:"BlockIPList"`

	// Expiration time in seconds. Default value: 3600
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// IP adding policy. Valid value: fifo (if a blocklist is full, new IPs added to the blocklist will adopt the first-in first-out policy)
	AddStrategy *string `json:"AddStrategy,omitnil,omitempty" name:"AddStrategy"`
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

// Predefined struct for user
type ModifyBlockIPListResponseParams struct {
	// Async task ID
	JodId *string `json:"JodId,omitnil,omitempty" name:"JodId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBlockIPListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockIPListResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDomainAttributesRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The domain name, which must be associated with an existing forwarding rule. If there are multiple domain names, you only need to specify one.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The one domain name to modify. `NewDomain` and `NewDomains` can not be both specified.
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`

	// Certificate information of the domain name. It is only applicable to listeners with SNI enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Whether to enable HTTP/2. Note: HTTP/2 can be enabled only for HTTPS domain names.
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// Whether to set this domain name as the default domain name. Note: Only one default domain name can be set under one listener.
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// Whether to enable QUIC. Note: QUIC can be enabled only for HTTPS domain names.
	Quic *bool `json:"Quic,omitnil,omitempty" name:"Quic"`

	// Specifies a new default domain name for the listener. This field is used when the original default domain name is disabled. If there are multiple domain names, specify one of them.
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`

	// The new domain names to modify. `NewDomain` and `NewDomains` can not be both specified.
	NewDomains []*string `json:"NewDomains,omitnil,omitempty" name:"NewDomains"`

	// Certificate information of the domain name. It is only applicable to listeners with SNI enabled. You can specify multiple server-side certificates with different algorithm types. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`
}

type ModifyDomainAttributesRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The domain name, which must be associated with an existing forwarding rule. If there are multiple domain names, you only need to specify one.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The one domain name to modify. `NewDomain` and `NewDomains` can not be both specified.
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`

	// Certificate information of the domain name. It is only applicable to listeners with SNI enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Whether to enable HTTP/2. Note: HTTP/2 can be enabled only for HTTPS domain names.
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// Whether to set this domain name as the default domain name. Note: Only one default domain name can be set under one listener.
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// Whether to enable QUIC. Note: QUIC can be enabled only for HTTPS domain names.
	Quic *bool `json:"Quic,omitnil,omitempty" name:"Quic"`

	// Specifies a new default domain name for the listener. This field is used when the original default domain name is disabled. If there are multiple domain names, specify one of them.
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`

	// The new domain names to modify. `NewDomain` and `NewDomains` can not be both specified.
	NewDomains []*string `json:"NewDomains,omitnil,omitempty" name:"NewDomains"`

	// Certificate information of the domain name. It is only applicable to listeners with SNI enabled. You can specify multiple server-side certificates with different algorithm types. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`
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
	delete(f, "Quic")
	delete(f, "NewDefaultServerDomain")
	delete(f, "NewDomains")
	delete(f, "MultiCertInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDomainRequestParams struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Legacy domain name under a listener.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// New domain name. 	Length limit: 1-120. There are three usage formats: non-regular expression, wildcard, and regular expression. A non-regular expression can only contain letters, digits, "-", and ".". In a wildcard, "*" can only be at the beginning or the end. A regular expressions must begin with a "~".
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Legacy domain name under a listener.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// New domain name. 	Length limit: 1-120. There are three usage formats: non-regular expression, wildcard, and regular expression. A non-regular expression can only contain letters, digits, "-", and ".". In a wildcard, "*" can only be at the beginning or the end. A regular expressions must begin with a "~".
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`
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

// Predefined struct for user
type ModifyDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyFunctionTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The backend cloud functions to modify
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or `Domain`+`Url`.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if `LocationId` is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if `LocationId` is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ModifyFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The backend cloud functions to modify
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or `Domain`+`Url`.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if `LocationId` is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if `LocationId` is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *ModifyFunctionTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "FunctionTargets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFunctionTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionTargetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFunctionTargetsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFunctionTargetsResponseParams `json:"Response"`
}

func (r *ModifyFunctionTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// New listener name
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Session persistence time in seconds. Value range: 30-3,600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Health check parameter. It is only applicable only to TCP, UDP, TCP_SSL and QUIC listeners.
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Certificate information. This parameter is only applicable to HTTPS/TCP_SSL/QUIC listeners. `Certificate` and `MultiCertInfo` cannot be both specified.
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Forwarding method of a listener. Value range: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Whether to enable the SNI feature. This parameter is applicable only to HTTPS listeners. Note: The SNI feature can be enabled but cannot be disabled once enabled.
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// Target backend type. `NODE`: A single node; `TARGETGROUP`: A target group.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Whether to enable a persistent connection. This parameter is applicable only to HTTP and HTTPS listeners.
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// Session persistence type. `NORMAL`: default session persistence type (L4/L7 session persistence); `QUIC_CID`: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners.
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// Certificate information. You can specify multiple server-side certificates with different algorithm types. This parameter is only applicable to HTTPS listeners with the SNI feature not enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// The maximum number of concurrent connections at the listener level. This parameter takes effect only on LCU-supported instances and TCP/UDP/TCP_SSL/QUIC listeners. Value range: 1 to the maximum concurrency of the instance. -1 indicates that no limit is set on concurrent connections.
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// The maximum number of new connections at the listener level. This parameter takes effect only on LCU-supported instances and TCP/UDP/TCP_SSL/QUIC listeners. Value range: 1 to the maximum number of new connections of the instance. -1 indicates that no limit is set on concurrent connections.
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// Connection idle timeout period (in seconds). It’s only available to TCP listeners. Value range: 300-900 for shared and dedicated instances; 300-2000 for LCU-supported CLB instances. It defaults to 900. To set a period longer than 2000 seconds (up to 3600 seconds), please submit a [submit](https://console.cloud.tencent.com/workorder/category). 
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// New listener name
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Session persistence time in seconds. Value range: 30-3,600. The default value is 0, indicating that session persistence is not enabled. This parameter is applicable only to TCP/UDP listeners.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Health check parameter. It is only applicable only to TCP, UDP, TCP_SSL and QUIC listeners.
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Certificate information. This parameter is only applicable to HTTPS/TCP_SSL/QUIC listeners. `Certificate` and `MultiCertInfo` cannot be both specified.
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Forwarding method of a listener. Value range: WRR, LEAST_CONN.
	// They represent weighted round robin and least connections, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Whether to enable the SNI feature. This parameter is applicable only to HTTPS listeners. Note: The SNI feature can be enabled but cannot be disabled once enabled.
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// Target backend type. `NODE`: A single node; `TARGETGROUP`: A target group.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Whether to enable a persistent connection. This parameter is applicable only to HTTP and HTTPS listeners.
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// Whether to send the TCP RST packet to the client when unbinding a real server. This parameter is applicable to TCP listeners only.
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// Session persistence type. `NORMAL`: default session persistence type (L4/L7 session persistence); `QUIC_CID`: session persistence by QUIC connection ID. The `QUIC_CID` value can only be configured in UDP listeners.
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// Certificate information. You can specify multiple server-side certificates with different algorithm types. This parameter is only applicable to HTTPS listeners with the SNI feature not enabled. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// The maximum number of concurrent connections at the listener level. This parameter takes effect only on LCU-supported instances and TCP/UDP/TCP_SSL/QUIC listeners. Value range: 1 to the maximum concurrency of the instance. -1 indicates that no limit is set on concurrent connections.
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// The maximum number of new connections at the listener level. This parameter takes effect only on LCU-supported instances and TCP/UDP/TCP_SSL/QUIC listeners. Value range: 1 to the maximum number of new connections of the instance. -1 indicates that no limit is set on concurrent connections.
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// Connection idle timeout period (in seconds). It’s only available to TCP listeners. Value range: 300-900 for shared and dedicated instances; 300-2000 for LCU-supported CLB instances. It defaults to 900. To set a period longer than 2000 seconds (up to 3600 seconds), please submit a [submit](https://console.cloud.tencent.com/workorder/category). 
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`
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
	delete(f, "TargetType")
	delete(f, "KeepaliveEnable")
	delete(f, "DeregisterTargetRst")
	delete(f, "SessionType")
	delete(f, "MultiCertInfo")
	delete(f, "MaxConn")
	delete(f, "MaxCps")
	delete(f, "IdleConnectTimeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyListenerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyListenerResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLoadBalancerAttributesRequestParams struct {
	// Unique CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// The backend service information of cross-region binding 1.0
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitnil,omitempty" name:"TargetRegionInfo"`

	// Network billing parameter
	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitnil,omitempty" name:"InternetChargeInfo"`

	// Whether the target opens traffic from CLB to the internet. If yes (true), only security groups on CLB will be verified; if no (false), security groups on both CLB and backend instance need to be verified.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// Whether to enable cross-region binding 2.0
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// Specifies whether to enable deletion protection.
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`

	// Modifies the second-level domain name of CLB from mycloud.com to tencentclb.com. Note that the sub-domain names will be changed as well. After the modification, mycloud.com will be invalidated. 
	ModifyClassicDomain *bool `json:"ModifyClassicDomain,omitnil,omitempty" name:"ModifyClassicDomain"`
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Unique CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// The backend service information of cross-region binding 1.0
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitnil,omitempty" name:"TargetRegionInfo"`

	// Network billing parameter
	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitnil,omitempty" name:"InternetChargeInfo"`

	// Whether the target opens traffic from CLB to the internet. If yes (true), only security groups on CLB will be verified; if no (false), security groups on both CLB and backend instance need to be verified.
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// Whether to enable cross-region binding 2.0
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// Specifies whether to enable deletion protection.
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`

	// Modifies the second-level domain name of CLB from mycloud.com to tencentclb.com. Note that the sub-domain names will be changed as well. After the modification, mycloud.com will be invalidated. 
	ModifyClassicDomain *bool `json:"ModifyClassicDomain,omitnil,omitempty" name:"ModifyClassicDomain"`
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
	delete(f, "ModifyClassicDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesResponseParams struct {
	// This parameter can be used to query whether CLB billing mode switch is successful.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLoadBalancerSlaRequestParams struct {
	// CLB instance information.
	LoadBalancerSla []*SlaUpdateParam `json:"LoadBalancerSla,omitnil,omitempty" name:"LoadBalancerSla"`
}

type ModifyLoadBalancerSlaRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance information.
	LoadBalancerSla []*SlaUpdateParam `json:"LoadBalancerSla,omitnil,omitempty" name:"LoadBalancerSla"`
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

// Predefined struct for user
type ModifyLoadBalancerSlaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerSlaResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerSlaResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLoadBalancersProjectRequestParams struct {
	// IDs of CLB instances ID(s).
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyLoadBalancersProjectRequest struct {
	*tchttp.BaseRequest
	
	// IDs of CLB instances ID(s).
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyLoadBalancersProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancersProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancersProjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancersProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancersProjectResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancersProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// ID of the forwarding rule to be modified.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// New forwarding path of the forwarding rule. This parameter is not required if the URL does not need to be modified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Health check information
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Request forwarding method of the rule. Value range: WRR, LEAST_CONN, IP_HASH
	// They represent weighted round robin, least connections, and IP hash, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Session persistence time
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Forwarding protocol between CLB instance and real server. Default value: HTTP. Valid values: HTTP, HTTPS, and TRPC.
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// TRPC callee server route, which is required when `ForwardType` is "TRPC". This is now only for internal usage.
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC calling service API, which is required when `ForwardType` is "TRPC". This is now only for internal usage.
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// ID of the forwarding rule to be modified.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// New forwarding path of the forwarding rule. This parameter is not required if the URL does not need to be modified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Health check information
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Request forwarding method of the rule. Value range: WRR, LEAST_CONN, IP_HASH
	// They represent weighted round robin, least connections, and IP hash, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Session persistence time
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Forwarding protocol between CLB instance and real server. Default value: HTTP. Valid values: HTTP, HTTPS, and TRPC.
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// TRPC callee server route, which is required when `ForwardType` is "TRPC". This is now only for internal usage.
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC calling service API, which is required when `ForwardType` is "TRPC". This is now only for internal usage.
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`
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

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTargetGroupAttributeRequestParams struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// New name of target group
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// New default port of target group
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type ModifyTargetGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// New name of target group
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// New default port of target group
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`
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

// Predefined struct for user
type ModifyTargetGroupAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTargetGroupInstancesPortRequestParams struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify ports
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type ModifyTargetGroupInstancesPortRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify ports
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
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

// Predefined struct for user
type ModifyTargetGroupInstancesPortResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupInstancesPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupInstancesPortResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTargetGroupInstancesWeightRequestParams struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify weights
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type ModifyTargetGroupInstancesWeightRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Array of servers for which to modify weights
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
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

// Predefined struct for user
type ModifyTargetGroupInstancesWeightResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupInstancesWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupInstancesWeightResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTargetPortRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the ports
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// New port of the real server bound to a listener or forwarding rule
	NewPort *int64 `json:"NewPort,omitnil,omitempty" name:"NewPort"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of real servers for which to modify the ports
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// New port of the real server bound to a listener or forwarding rule
	NewPort *int64 `json:"NewPort,omitnil,omitempty" name:"NewPort"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
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

// Predefined struct for user
type ModifyTargetPortResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetPortResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTargetWeightRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// List of real servers for which to modify the weights
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// New forwarding weight of a real server. Value range: 0-100. Default value: 10. If the Targets.Weight parameter is set, this parameter will not take effect.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// List of real servers for which to modify the weights
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// New forwarding weight of a real server. Value range: 0-100. Default value: 10. If the Targets.Weight parameter is set, this parameter will not take effect.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
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

// Predefined struct for user
type ModifyTargetWeightResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetWeightResponseParams `json:"Response"`
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

type MultiCertInfo struct {
	// Authentication type. Values: `UNIDIRECTIONAL` (one-way authentication), `MUTUAL` (two-way authentication)
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// List of listener or rule certificates. One-way and two-way authentication are supported. Only one certificate can be specified for one algorithm. If `SSLMode` is `MUTUAL` (two-way authentication), at least one CA certificate is required. 
	CertList []*CertInfo `json:"CertList,omitnil,omitempty" name:"CertList"`
}

type Price struct {
	// Instance price.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Network price.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitnil,omitempty" name:"BandwidthPrice"`

	// LCU price.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LcuPrice *ItemPrice `json:"LcuPrice,omitnil,omitempty" name:"LcuPrice"`
}

type Quota struct {
	// Quota name. Valid values:
	// <li> `TOTAL_OPEN_CLB_QUOTA`: Quota of public network CLB instances in the current region</li>
	// <li> `TOTAL_INTERNAL_CLB_QUOTA`: Quota of private network CLB instances in the current region</li>
	// <li> `TOTAL_LISTENER_QUOTA`: Quota of listeners under one CLB instance</li>
	// <li> `TOTAL_LISTENER_RULE_QUOTA`: Quota of forwarding rules under one listener</li>
	// <li> `TOTAL_TARGET_BIND_QUOTA`: Quota of CVM instances can be bound under one forwarding rule</li>
	// <li> `TOTAL_SNAP_IP_QUOTA`: Quota of SNAT IPs for cross-region binding 2.0 under one CLB instance </li>
	// <li> `TOTAL_ISP_CLB_QUOTA`: Quota of triple-ISP (CMCC/CUCC/CTCC) CLB instances in the current region</li>
	QuotaId *string `json:"QuotaId,omitnil,omitempty" name:"QuotaId"`

	// Currently used quantity. If it is `null`, it is meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	QuotaCurrent *int64 `json:"QuotaCurrent,omitnil,omitempty" name:"QuotaCurrent"`

	// Quota limit.
	QuotaLimit *int64 `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`
}

// Predefined struct for user
type RegisterFunctionTargetsRequestParams struct {
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// SCF functions to be bound.
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// ID of the target forwarding rule. To bind an SCF function to a L7 forwarding rule, this parameter or `Domain+Url` is required.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the target forwarding rule. It is ignored if `LocationId` is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// URL of the target forwarding rule. It is ignored if `LocationId` is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type RegisterFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// SCF functions to be bound.
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// ID of the target forwarding rule. To bind an SCF function to a L7 forwarding rule, this parameter or `Domain+Url` is required.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the target forwarding rule. It is ignored if `LocationId` is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// URL of the target forwarding rule. It is ignored if `LocationId` is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *RegisterFunctionTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterFunctionTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "FunctionTargets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterFunctionTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterFunctionTargetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterFunctionTargetsResponse struct {
	*tchttp.BaseResponse
	Response *RegisterFunctionTargetsResponseParams `json:"Response"`
}

func (r *RegisterFunctionTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterFunctionTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetGroupInstancesRequestParams struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Server instance array
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type RegisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Server instance array
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
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

// Predefined struct for user
type RegisterTargetGroupInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RegisterTargetGroupInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type RegisterTargetsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of real servers to be bound. Array length limit: 20.
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target forwarding rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target forwarding rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type RegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of real servers to be bound. Array length limit: 20.
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Forwarding rule ID. When binding a real server to a layer-7 forwarding rule, you must provide either this parameter or Domain+Url.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target forwarding rule domain name. This parameter does not take effect if LocationId is specified.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target forwarding rule URL. This parameter does not take effect if LocationId is specified.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
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

// Predefined struct for user
type RegisterTargetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *RegisterTargetsResponseParams `json:"Response"`
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

// Predefined struct for user
type RegisterTargetsWithClassicalLBRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Real server information
	Targets []*ClassicalTargetInfo `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type RegisterTargetsWithClassicalLBRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Real server information
	Targets []*ClassicalTargetInfo `json:"Targets,omitnil,omitempty" name:"Targets"`
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

// Predefined struct for user
type RegisterTargetsWithClassicalLBResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterTargetsWithClassicalLBResponse struct {
	*tchttp.BaseResponse
	Response *RegisterTargetsWithClassicalLBResponseParams `json:"Response"`
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

// Predefined struct for user
type ReplaceCertForLoadBalancersRequestParams struct {
	// ID of the certificate to be replaced, which can be a server certificate or a client certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Information such as the content of the new certificate
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`
}

type ReplaceCertForLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// ID of the certificate to be replaced, which can be a server certificate or a client certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Information such as the content of the new certificate
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`
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

// Predefined struct for user
type ReplaceCertForLoadBalancersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplaceCertForLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceCertForLoadBalancersResponseParams `json:"Response"`
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

type Resource struct {
	// Specific ISP resource information, Vaules: `CMCC`, `CUCC`, `CTCC`, `BGP`, and `INTERNAL`.
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// ISP information, such as `CMCC`, `CUCC`, `CTCC`, `BGP`, and `INTERNAL`.
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// Available resources
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AvailabilitySet []*ResourceAvailability `json:"AvailabilitySet,omitnil,omitempty" name:"AvailabilitySet"`

	// ISP Type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TypeSet []*TypeInfo `json:"TypeSet,omitnil,omitempty" name:"TypeSet"`
}

type ResourceAvailability struct {
	// Specific ISP resource information. Values: `CMCC`, `CUCC`, `CTCC`, `BGP`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether the resource is available. Values: `Available`, `Unavailable`
	Availability *string `json:"Availability,omitnil,omitempty" name:"Availability"`
}

type RewriteLocationMap struct {
	// Source forwarding rule ID
	SourceLocationId *string `json:"SourceLocationId,omitnil,omitempty" name:"SourceLocationId"`

	// ID of the forwarding rule of the destination
	TargetLocationId *string `json:"TargetLocationId,omitnil,omitempty" name:"TargetLocationId"`

	// Redirection status code. Valid values: 301, 302, and 307.
	RewriteCode *int64 `json:"RewriteCode,omitnil,omitempty" name:"RewriteCode"`

	// Whether the matched URL is carried in redirection. It is required when configuring `RewriteCode`.
	TakeUrl *bool `json:"TakeUrl,omitnil,omitempty" name:"TakeUrl"`

	// Original domain name of redirection, which must be the corresponding domain name of `SourceLocationId`. It is required when configuring `RewriteCode`.
	SourceDomain *string `json:"SourceDomain,omitnil,omitempty" name:"SourceDomain"`
}

type RewriteTarget struct {
	// Listener ID of a redirect target
	// Note: This field may return null, indicating that there is no redirection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// Forwarding rule ID of a redirect target
	// Note: This field may return null, indicating that there is no redirection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetLocationId *string `json:"TargetLocationId,omitnil,omitempty" name:"TargetLocationId"`

	// Redirection status code
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RewriteCode *int64 `json:"RewriteCode,omitnil,omitempty" name:"RewriteCode"`

	// Whether the matched URL is carried in redirection.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TakeUrl *bool `json:"TakeUrl,omitnil,omitempty" name:"TakeUrl"`

	// Redirection type. Manual: manual redirection; Auto: automatic redirection.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RewriteType *string `json:"RewriteType,omitnil,omitempty" name:"RewriteType"`
}

type RsWeightRule struct {
	// CLB listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of real servers whose weights to modify.
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Forwarding rule ID, which is required only for layer-7 rules.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Target rule domain name. This parameter does not take effect if LocationId is specified
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Target rule URL. This parameter does not take effect if LocationId is specified
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// The new forwarding weight of the real server. Value range: [0, 100]. This parameter takes lower precedence than `Weight` in [`Targets`](https://intl.cloud.tencent.com/document/api/214/30694?from_cn_redirect=1#Target), which means that this parameter only takes effect when the `Weight` in `RsWeightRule` is left empty.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type RuleHealth struct {
	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Forwarding rule Url
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Health status of the real server bound to this rule
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Targets []*TargetHealth `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type RuleInput struct {
	// Forwarding rule path. Length: 1-200.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// The domain name associated with the forwarding rule. It can contain 1-80 characters. Only one domain name can be entered. If you need to enter multiple domain names, use `Domains`.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Session persistence time in seconds. Value range: 30-3,600. Setting it to 0 indicates that session persistence is disabled.
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Health check information. For more information, please see [Health Check](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1)
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Certificate information. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Request forwarding method of the rule. Value range: WRR, LEAST_CONN, IP_HASH
	// They represent weighted round robin, least connections, and IP hash, respectively. Default value: WRR.
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// Forwarding protocol between the CLB instance and backend service. Values: `HTTP`, `HTTPS`, `GRPC` and `TRPC` (only for internal usage). It defaults to `HTTP`.
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// Whether to set this domain name as the default domain name. Note: Only one default domain name can be set under one listener.
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// Whether to enable HTTP/2. Note: HTTP/2 can be enabled only for HTTPS domain names.
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// Target real server type. NODE: binding a general node; TARGETGROUP: binding a target group.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// TRPC callee server route, which is required when `ForwardType` is "TRPC". This is now only for internal usage.
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC calling service API, which is required when `ForwardType` is "TRPC". This is now only for internal usage.
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`

	// Whether to enable QUIC. Note: QUIC can be enabled only for HTTPS domain names
	Quic *bool `json:"Quic,omitnil,omitempty" name:"Quic"`

	// The domain name associated with the forwarding rule. Each contain 1-80 characters. If you only need to enter one domain name, use `Domain` instead.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Certificate information. You can specify multiple server-side certificates with different algorithm types. `Certificate` and `MultiCertInfo` cannot be specified at the same time. 
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`
}

type RuleOutput struct {
	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Forwarding rule path.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Session persistence time
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// Health check information
	// Note: This field may return null, indicating that no valid values can be obtained.
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// Certificate information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Certificate *CertificateOutput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Request forwarding method of the rule
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// ID of the listener to which the forwarding rule belongs
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Redirect target information of a forwarding rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	RewriteTarget *RewriteTarget `json:"RewriteTarget,omitnil,omitempty" name:"RewriteTarget"`

	// Whether to enable gzip
	HttpGzip *bool `json:"HttpGzip,omitnil,omitempty" name:"HttpGzip"`

	// Whether the forwarding rule is automatically created
	BeAutoCreated *bool `json:"BeAutoCreated,omitnil,omitempty" name:"BeAutoCreated"`

	// Whether to use as the default domain name
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// Whether to enable Http2
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// Forwarding protocol between CLB and real server
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// Forwarding rule creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Real server type
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Basic information of a bound target group. This field will be returned if a target group is bound to a rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitnil,omitempty" name:"TargetGroup"`

	// WAF instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	WafDomainId *string `json:"WafDomainId,omitnil,omitempty" name:"WafDomainId"`

	// TRPC callee server route, which is valid when `ForwardType` is `TRPC`. This is now only for internal usage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC calling service API, which is valid when `ForwardType` is `TRPC`. This is now only for internal usage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`

	// QUIC status
	// Note: this field may return null, indicating that no valid values can be obtained.
	QuicStatus *string `json:"QuicStatus,omitnil,omitempty" name:"QuicStatus"`

	// List of domain names associated with the forwarding rule
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// List of bound target groups
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TargetGroupList []*BasicTargetGroupInfo `json:"TargetGroupList,omitnil,omitempty" name:"TargetGroupList"`
}

type RuleTargets struct {
	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name of the forwarding rule
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Forwarding rule path.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Real server information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets []*Backend `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Information about backend SCF functions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`
}

type RulesItems struct {
	// Rule ID.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// Domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Uri
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Object bound to the real server.
	Targets []*LbRsTargets `json:"Targets,omitnil,omitempty" name:"Targets"`
}

// Predefined struct for user
type SetCustomizedConfigForLoadBalancerRequestParams struct {
	// Operation type: `ADD`, `DELETE`, `UPDATE`, `BIND`, `UNBIND`
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// This field is required except for creating custom configurations, such as "pz-1234abcd".
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// This field is required when creating or modifying custom configurations.
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// This field is required when creating or renaming custom configurations.
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// This field is required when binding/unbinding resources.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type SetCustomizedConfigForLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// Operation type: `ADD`, `DELETE`, `UPDATE`, `BIND`, `UNBIND`
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// This field is required except for creating custom configurations, such as "pz-1234abcd".
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// This field is required when creating or modifying custom configurations.
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// This field is required when creating or renaming custom configurations.
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// This field is required when binding/unbinding resources.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

func (r *SetCustomizedConfigForLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCustomizedConfigForLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "UconfigId")
	delete(f, "ConfigContent")
	delete(f, "ConfigName")
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCustomizedConfigForLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCustomizedConfigForLoadBalancerResponseParams struct {
	// Configuration ID, such as "pz-1234abcd"
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetCustomizedConfigForLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *SetCustomizedConfigForLoadBalancerResponseParams `json:"Response"`
}

func (r *SetCustomizedConfigForLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCustomizedConfigForLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerClsLogRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLS logset ID
	// <li>Enter the ID of logset you need to add or update. You can acquire the ID by invoking [DescribeLogsets](https://intl.cloud.tencent.com/document/product/614/56454?from_cn_redirect=1).</li>
	// <li>To delete the log set, set this parameter to `null`.</li>
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// CLS log topic ID
	// <li>Enter the ID of log topic you need to add or update. You can acquire the ID by invoking [DescribeTopics](https://intl.cloud.tencent.com/document/product/614/56454?from_cn_redirect=1).</li>
	// <li>To delete the log set, set this parameter to `null`.</li>
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// Log type:
	// <li>`ACCESS`: access logs</li>
	// <li>`HEALTH`: health check logs</li>
	// Default: `ACCESS`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type SetLoadBalancerClsLogRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLS logset ID
	// <li>Enter the ID of logset you need to add or update. You can acquire the ID by invoking [DescribeLogsets](https://intl.cloud.tencent.com/document/product/614/56454?from_cn_redirect=1).</li>
	// <li>To delete the log set, set this parameter to `null`.</li>
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// CLS log topic ID
	// <li>Enter the ID of log topic you need to add or update. You can acquire the ID by invoking [DescribeTopics](https://intl.cloud.tencent.com/document/product/614/56454?from_cn_redirect=1).</li>
	// <li>To delete the log set, set this parameter to `null`.</li>
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// Log type:
	// <li>`ACCESS`: access logs</li>
	// <li>`HEALTH`: health check logs</li>
	// Default: `ACCESS`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
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

// Predefined struct for user
type SetLoadBalancerClsLogResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetLoadBalancerClsLogResponse struct {
	*tchttp.BaseResponse
	Response *SetLoadBalancerClsLogResponseParams `json:"Response"`
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

// Predefined struct for user
type SetLoadBalancerSecurityGroupsRequestParams struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of security group IDs. One CLB instance can be bound to up to 50 security groups. If you want to unbind all security groups, you do not need to pass in this parameter, or you can pass in an empty array.
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Array of security group IDs. One CLB instance can be bound to up to 50 security groups. If you want to unbind all security groups, you do not need to pass in this parameter, or you can pass in an empty array.
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`
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

// Predefined struct for user
type SetLoadBalancerSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetLoadBalancerSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *SetLoadBalancerSecurityGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type SetSecurityGroupForLoadbalancersRequestParams struct {
	// Security group ID, such as sg-12345678
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// ADD: bind a security group;
	// DEL: unbind a security group
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Array of CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type SetSecurityGroupForLoadbalancersRequest struct {
	*tchttp.BaseRequest
	
	// Security group ID, such as sg-12345678
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// ADD: bind a security group;
	// DEL: unbind a security group
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Array of CLB instance IDs
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
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

// Predefined struct for user
type SetSecurityGroupForLoadbalancersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetSecurityGroupForLoadbalancersResponse struct {
	*tchttp.BaseResponse
	Response *SetSecurityGroupForLoadbalancersResponseParams `json:"Response"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// LCU-supported instance specification. Value:
	// <li>`SLA`: If you have activated Super Large LCU-supported instances, `SLA` indicates Super Large 4.</li>
	// <li>`clb.c2.medium`: Standard</li>
	// <li>`clb.c3.small`: Advanced 1</li>
	// <li>`clb.c3.medium`: Advanced 2</li>
	// <li>`clb.c4.small`: Super Large 1</li>
	// <li>`clb.c4.medium`: Super Large 2</li>
	// <li>`clb.c4.large`: Super Large 3</li>
	// <li>`clb.c4.xlarge`: Super Large 4</li> For Super Large 2 and above specifications, please [submit a ticket](https://console.cloud.tencent.com/workorder/category). For more specifications, see [Specifications Comparison](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1)
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`
}

type SnatIp struct {
	// Unique VPC subnet ID, such as `subnet-12345678`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// IP address, such as 192.168.0.1
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type SpecAvailability struct {
	// Specification type
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecType *string `json:"SpecType,omitnil,omitempty" name:"SpecType"`

	// Specification availability
	// Note: This field may return null, indicating that no valid values can be obtained.
	Availability *string `json:"Availability,omitnil,omitempty" name:"Availability"`
}

type TagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Target struct {
	// Listening port of a real server
	// Note: this parameter is required when binding a CVM or ENI.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Real server type. Value range: CVM (Cloud Virtual Machine), ENI (Elastic Network Interface). This parameter does not take effect currently as an input parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Unique ID of a CVM instance, which is required when binding a CVM instance. It can be obtained from the `InstanceId` field in the response of the `DescribeInstances` API. It indicates binding the primary IP of the primary ENI.
	// Note: Either `InstanceId` or `EniIp` can be passed in.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The new forwarding weight of the real server. Value range: [0, 100]. Default: 10. This parameter takes priority over `Weight` in [`RsWeightRule`](https://intl.cloud.tencent.com/document/api/214/30694?from_cn_redirect=1#RsWeightRule). If it’s left empty, the value of `Weight` in `RsWeightRule` will be used.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// It is required when binding an IP. ENI IPs and other private IPs are supported. To bind an ENI IP, the ENI should be bound to a CVM instance before being bound to a CLB instance.
	// Note: Either `InstanceId` or `EniIp` can be passed in. `EniIp` is required in a cross-region binding or when the dual-stack IPV6 CVM is bound.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EniIp *string `json:"EniIp,omitnil,omitempty" name:"EniIp"`
}

type TargetGroupAssociation struct {
	// CLB instance ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Forwarding rule ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`
}

type TargetGroupBackend struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Real server type. Valid values: CVM, ENI (coming soon)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Unique real server ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Listening port of real server
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Forwarding weight of real server. Value range: [0, 100]. Default value: 10.
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Public IP of real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// Private IP of real server
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// Real server instance name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Real server binding time
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// Unique ENI ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`

	// AZ ID of the real server
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type TargetGroupInfo struct {
	// Target group ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// `vpcid` of target group
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Target group name
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// Default port of target group
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Target group creation time
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Target group modification time
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// Array of associated rules
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociatedRule []*AssociationItem `json:"AssociatedRule,omitnil,omitempty" name:"AssociatedRule"`
}

type TargetGroupInstance struct {
	// Private IP of target group instance
	BindIP *string `json:"BindIP,omitnil,omitempty" name:"BindIP"`

	// Port of target group instance
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Weight of target group instance
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// New port of target group instance
	NewPort *uint64 `json:"NewPort,omitnil,omitempty" name:"NewPort"`
}

type TargetHealth struct {
	// Private IP of the target
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Port bound to the target
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Current health status. true: healthy; false: unhealthy.
	HealthStatus *bool `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// Instance ID of the target, such as ins-12345678
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// Detailed information about the current health status. Alive: healthy; Dead: exceptional; Unknown: check not started/checking/unknown status.
	HealthStatusDetail *string `json:"HealthStatusDetail,omitnil,omitempty" name:"HealthStatusDetail"`

	// (**This parameter will be disused soon. Please use `HealthStatusDetail` instead.**) Details of the current health status. Values: `Alive` (healthy), `Dead` (abnormal), `Unknown` (Health check not started/checking/unknown status)
	//
	// Deprecated: HealthStatusDetial is deprecated.
	HealthStatusDetial *string `json:"HealthStatusDetial,omitnil,omitempty" name:"HealthStatusDetial"`
}

type TargetRegionInfo struct {
	// Region of the target, such as ap-guangzhou
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Network of the target, which is in the format of vpc-abcd1234 for VPC or 0 for basic network
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type TypeInfo struct {
	// ISP Type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Specification availability
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecAvailabilitySet []*SpecAvailability `json:"SpecAvailabilitySet,omitnil,omitempty" name:"SpecAvailabilitySet"`
}

type ZoneInfo struct {
	// Unique AZ ID in a numeric form, such as 100001
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Unique AZ ID in a string form, such as ap-guangzhou-1
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// AZ name, such as Guangzhou Zone 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// AZ region, e.g., ap-guangzhou.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ZoneRegion *string `json:"ZoneRegion,omitnil,omitempty" name:"ZoneRegion"`

	// Whether the AZ is the `LocalZone`, e.g., false.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LocalZone *bool `json:"LocalZone,omitnil,omitempty" name:"LocalZone"`

	// Whether the AZ is an edge zone. Values: `true`, `false`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EdgeZone *bool `json:"EdgeZone,omitnil,omitempty" name:"EdgeZone"`
}

type ZoneResource struct {
	// Primary AZ, such as "ap-guangzhou-1".
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// List of resources
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceSet []*Resource `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// Secondary AZ, such as "ap-guangzhou-2". 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// IP version. Values: `IPv4`, `IPv6`, and `IPv6_Nat`.
	IPVersion *string `json:"IPVersion,omitnil,omitempty" name:"IPVersion"`

	// Region of the AZ, such as `ap-guangzhou`.
	ZoneRegion *string `json:"ZoneRegion,omitnil,omitempty" name:"ZoneRegion"`

	// Whether the AZ is a `LocalZone`. Values: `true`, `false`.
	LocalZone *bool `json:"LocalZone,omitnil,omitempty" name:"LocalZone"`

	// Type of resources in the zone. Values: `SHARED`, `EXCLUSIVE`
	ZoneResourceType *string `json:"ZoneResourceType,omitnil,omitempty" name:"ZoneResourceType"`

	// Whether the AZ is an edge zone. Values: `true`, `false`.
	EdgeZone *bool `json:"EdgeZone,omitnil,omitempty" name:"EdgeZone"`

	// Network egress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`
}