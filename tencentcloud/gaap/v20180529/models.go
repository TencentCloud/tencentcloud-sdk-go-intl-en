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

package v20180529

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccessConfiguration struct {

	// Acceleration region.
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// Connection bandwidth upper limit in Mbps.
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Concurrent connection upper limit in 10,000 connections, which indicates the allowed number of concurrently online connections.
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`
}

type AccessRegionDetial struct {

	// Region ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Region name in Chinese or English
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Value array of the available concurrence
	ConcurrentList []*int64 `json:"ConcurrentList,omitempty" name:"ConcurrentList" list`

	// Value array of the available bandwidth
	BandwidthList []*int64 `json:"BandwidthList,omitempty" name:"BandwidthList" list`
}

type AccessRegionDomainConf struct {

	// Region ID.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Region/country code for the nearest access, which can be obtained via the DescribeCountryAreaMapping API.
	NationCountryInnerList []*string `json:"NationCountryInnerList,omitempty" name:"NationCountryInnerList" list`
}

type AddRealServersRequest struct {
	*tchttp.BaseRequest

	// Project ID corresponding to origin server
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP or domain name corresponding to origin server
	RealServerIP []*string `json:"RealServerIP,omitempty" name:"RealServerIP" list`

	// Origin server name
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// Tag list
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`
}

func (r *AddRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RealServerIP")
	delete(f, "RealServerName")
	delete(f, "TagSet")
	if len(f) > 0 {
		return errors.New("AddRealServersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Origin server information list
		RealServerSet []*NewRealServer `json:"RealServerSet,omitempty" name:"RealServerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthPriceGradient struct {

	// Bandwidth range.
	BandwidthRange []*int64 `json:"BandwidthRange,omitempty" name:"BandwidthRange" list`

	// Bandwidth unit price within the bandwidth range. Unit: CNY/Mbps/day.
	BandwidthUnitPrice *float64 `json:"BandwidthUnitPrice,omitempty" name:"BandwidthUnitPrice"`

	// Discounted bandwidth price in CNY/Mbps/day.
	DiscountBandwidthUnitPrice *float64 `json:"DiscountBandwidthUnitPrice,omitempty" name:"DiscountBandwidthUnitPrice"`
}

type BindListenerRealServersRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of origin servers to be bound. If the origin server scheduling policy type of this listener is weighted round robin, you need to enter the `RealServerWeight`, i.e., the origin server weight. If this field is left empty or for other scheduling types, the default origin server weight is 1.
	RealServerBindSet []*RealServerBindSetReq `json:"RealServerBindSet,omitempty" name:"RealServerBindSet" list`
}

func (r *BindListenerRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindListenerRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "RealServerBindSet")
	if len(f) > 0 {
		return errors.New("BindListenerRealServersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindListenerRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindListenerRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindListenerRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindRealServer struct {

	// Origin server ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// Origin server IP or domain name
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// Origin server weight
	RealServerWeight *int64 `json:"RealServerWeight,omitempty" name:"RealServerWeight"`

	// Origin server health check status. Valid values:
	// 0: normal;
	// 1: exceptional.
	// If health check is not enabled, this status will always be normal.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RealServerStatus *int64 `json:"RealServerStatus,omitempty" name:"RealServerStatus"`

	// Origin server port number
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerPort *int64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// If the origin server is a domain name, the domain name will be resolved to one or multiple IPs. This field indicates the exceptional IP list.
	DownIPList []*string `json:"DownIPList,omitempty" name:"DownIPList" list`
}

type BindRealServerInfo struct {

	// Origin server IP or domain name
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// Origin server ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// Origin server name
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`
}

type BindRuleRealServersRequest struct {
	*tchttp.BaseRequest

	// Forwarding rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// An information list of the origin servers to bind.
	// If there are origin servers bound already, they will be replaced by this new origin server list.
	// If this field is empty, it indicates unbinding all origin servers of this rule.
	// If the origin server scheduling policy type of this rule is weighted round robin, you need to enter `RealServerWeight`, i.e., the origin server weight. If this field is left empty or for other scheduling types, the default origin server weight is 1.
	RealServerBindSet []*RealServerBindSetReq `json:"RealServerBindSet,omitempty" name:"RealServerBindSet" list`
}

func (r *BindRuleRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRuleRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RealServerBindSet")
	if len(f) > 0 {
		return errors.New("BindRuleRealServersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindRuleRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindRuleRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRuleRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Certificate struct {

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Certificate name; It's an old parameter, please switch to CertificateAlias.
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`

	// Certificate type.
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// Certificate name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// Certificate creation time in the format of UNIX timestamp, indicating the number of seconds that have elapsed since January 1, 1970 (midnight in UTC/GMT).
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Certificate effective time in the format of UNIX timestamp, indicating the number of seconds that have elapsed since January 1, 1970 (midnight in UTC/GMT).
	// Note: This field may return null, indicating that no valid values can be obtained.
	BeginTime *uint64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// Certificate expiration time in the format of UNIX timestamp, indicating the number of seconds that have elapsed since January 1, 1970 (midnight in UTC/GMT).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Common name of the certificate issuer.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuerCN *string `json:"IssuerCN,omitempty" name:"IssuerCN"`

	// Common name of the certificate subject.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubjectCN *string `json:"SubjectCN,omitempty" name:"SubjectCN"`
}

type CertificateAliasInfo struct {

	// Certificate ID.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Certificate alias.
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`
}

type CertificateDetail struct {

	// Certificate ID.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Certificate type.
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// Certificate name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// Certificate content.
	CertificateContent *string `json:"CertificateContent,omitempty" name:"CertificateContent"`

	// Key content. This field will be returned if the certificate type is the SSL certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertificateKey *string `json:"CertificateKey,omitempty" name:"CertificateKey"`

	// Creation time in the format of UNIX timestamp, indicating the number of seconds that have elapsed since January 1, 1970 (midnight in UTC/GMT).
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Certificate effective time in the format of UNIX timestamp, indicating the number of seconds that have elapsed since January 1, 1970 (midnight in UTC/GMT).
	// Note: This field may return null, indicating that no valid values can be obtained.
	BeginTime *uint64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// Certificate expiration time in the format of UNIX timestamp, indicating the number of seconds that have elapsed since January 1, 1970 (midnight in UTC/GMT).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Common name of the certificate's issuer.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IssuerCN *string `json:"IssuerCN,omitempty" name:"IssuerCN"`

	// Common name of the certificate subject.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubjectCN *string `json:"SubjectCN,omitempty" name:"SubjectCN"`
}

type CheckProxyCreateRequest struct {
	*tchttp.BaseRequest

	// Access (acceleration) region of the connection. The value can be obtained via the DescribeAccessRegionsByDestRegion API.
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// Origin server region of the connection. The value can be obtained via the DescribeDestRegions API.
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// Connection bandwidth cap. Unit: Mbps.
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Connection concurrence cap, which indicates the maximum number of simultaneous online connections. Unit: 10,000 connections.
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// Connection group ID that needs to be entered when a connection is created in a connection group
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CheckProxyCreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckProxyCreateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRegion")
	delete(f, "RealServerRegion")
	delete(f, "Bandwidth")
	delete(f, "Concurrent")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("CheckProxyCreateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckProxyCreateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Queries whether a connection with the specified configuration can be created. 1: yes; 0: no.
		CheckFlag *uint64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckProxyCreateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckProxyCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxiesRequest struct {
	*tchttp.BaseRequest

	// Connection instance ID; It's an old parameter, please switch to ProxyIds.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	// For more information, please see How to Ensure Idempotence.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Connection instance ID; It's a new parameter.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds" list`
}

func (r *CloseProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return errors.New("CloseProxiesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Only the running connection instance ID lists can be enabled.
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet" list`

		// ID list of connection instances failed to be enabled.
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxyGroupRequest struct {
	*tchttp.BaseRequest

	// Connection group instance ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CloseProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("CloseProxyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of IDs of the connection instances that are not running, which cannot be enabled.
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet" list`

		// List of IDs of the connection instances failed to be enabled.
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// Connection ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Security group policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CloseSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return errors.New("CloseSecurityPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async Process ID. Using DescribeAsyncTaskStatus to query process and status.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountryAreaMap struct {

	// Country name.
	NationCountryName *string `json:"NationCountryName,omitempty" name:"NationCountryName"`

	// Country code.
	NationCountryInnerCode *string `json:"NationCountryInnerCode,omitempty" name:"NationCountryInnerCode"`

	// Region name.
	GeographicalZoneName *string `json:"GeographicalZoneName,omitempty" name:"GeographicalZoneName"`

	// Region code.
	GeographicalZoneInnerCode *string `json:"GeographicalZoneInnerCode,omitempty" name:"GeographicalZoneInnerCode"`

	// Continent name.
	ContinentName *string `json:"ContinentName,omitempty" name:"ContinentName"`

	// Continent code.
	ContinentInnerCode *string `json:"ContinentInnerCode,omitempty" name:"ContinentInnerCode"`
}

type CreateCertificateRequest struct {
	*tchttp.BaseRequest

	// Certificate type. Where:
	// 0: basic authentication configuration;
	// 1: indicates client CA certificate;
	// 2: server SSL certificate;
	// 3: origin server CA certificate;
	// 4: connection SSL certificate.
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// Certificate content. URL encoding. Where:
	// If the certificate type is basic authentication, enter username/password pair for this parameter. Format: 'username:password', for example, root:FSGdT. The password is `htpasswd` or `openssl`, for example, openssl passwd -crypt 123456.
	// When the certificate type is CA/SSL certificate, enter the certificate content for this parameter in the format of `pem`.
	CertificateContent *string `json:"CertificateContent,omitempty" name:"CertificateContent"`

	// Certificate name
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// Key content. URL encoding. This parameter is required only when the certificate type is SSL certificate. The format is `pem`.
	CertificateKey *string `json:"CertificateKey,omitempty" name:"CertificateKey"`
}

func (r *CreateCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateType")
	delete(f, "CertificateContent")
	delete(f, "CertificateAlias")
	delete(f, "CertificateKey")
	if len(f) > 0 {
		return errors.New("CreateCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainErrorPageInfoRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Original error code
	ErrorNos []*int64 `json:"ErrorNos,omitempty" name:"ErrorNos" list`

	// New response packet
	Body *string `json:"Body,omitempty" name:"Body"`

	// New error code
	NewErrorNo *int64 `json:"NewErrorNo,omitempty" name:"NewErrorNo"`

	// Response header to be deleted
	ClearHeaders []*string `json:"ClearHeaders,omitempty" name:"ClearHeaders" list`

	// Response header to be set
	SetHeaders []*HttpHeaderParam `json:"SetHeaders,omitempty" name:"SetHeaders" list`
}

func (r *CreateDomainErrorPageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainErrorPageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "ErrorNos")
	delete(f, "Body")
	delete(f, "NewErrorNo")
	delete(f, "ClearHeaders")
	delete(f, "SetHeaders")
	if len(f) > 0 {
		return errors.New("CreateDomainErrorPageInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainErrorPageInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration ID of a custom error response
		ErrorPageId *string `json:"ErrorPageId,omitempty" name:"ErrorPageId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainErrorPageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainErrorPageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name to be created. Each listener supports up to 100 domain names.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Server certificate, which is used for the HTTPS interaction between client and GAAP.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Client CA certificate, which is used for the HTTPS interaction between client and GAAP.
	// This field is required only when the mutual authentication method is adopted.
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// Client CA certificate, which is used for the HTTPS interaction between the client and GAAP.
	// This field or the `ClientCertificateId` field is required for mutual authentication only.
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds" list`
}

func (r *CreateDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	if len(f) > 0 {
		return errors.New("CreateDomainRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPListenerRequest struct {
	*tchttp.BaseRequest

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener port, which is based on the listeners of same transport layer protocol (TCP or UDP). The port must be unique.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Connection ID, which cannot be set together with `GroupId` at the same time. A listener will be created for the corresponding connection.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Connection group ID, which cannot be set together with `ProxyId` at the same time. A listener will be created for the corresponding connection group.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateHTTPListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("CreateHTTPListenerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Created listener ID
		ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHTTPListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPSListenerRequest struct {
	*tchttp.BaseRequest

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener port, which is based on the listeners of same transport layer protocol (TCP or UDP). The port must be unique.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Server certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Protocol types of the forwarding from acceleration connection to origin server: HTTP | HTTPS
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// Connection ID, which cannot be set together with `GroupId` at the same time. A listener will be created for the corresponding connection.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Authentication type, where:
	// 0: one-way authentication;
	// 1: mutual authentication.
	// The one-way authentication is used by default.
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// Client CA certificate ID, which is required only when the mutual authentication is adopted.
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// IDs of multiple new client CA certificates. This field or the `ClientCertificateId` field is required for mutual authentication only.
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds" list`

	// Connection group ID, which cannot be set together with `ProxyId` at the same time. A listener will be created for the corresponding connection group.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateHTTPSListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPSListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "CertificateId")
	delete(f, "ForwardProtocol")
	delete(f, "ProxyId")
	delete(f, "AuthType")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("CreateHTTPSListenerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHTTPSListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Created listener ID
		ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHTTPSListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPSListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupDomainRequest struct {
	*tchttp.BaseRequest

	// Connection group ID of the domain name to be enabled.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateProxyGroupDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("CreateProxyGroupDomainRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection group ID.
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProxyGroupDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupRequest struct {
	*tchttp.BaseRequest

	// Project ID of connection group
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Alias of connection group
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Origin server region; Reference API: DescribeDestRegions; It returnes the `RegionId` of the parameter `RegionDetail`.
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// Tag list
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// List of acceleration regions, including their names, bandwidth, and concurrence configuration.
	AccessRegionSet []*AccessConfiguration `json:"AccessRegionSet,omitempty" name:"AccessRegionSet" list`
}

func (r *CreateProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "GroupName")
	delete(f, "RealServerRegion")
	delete(f, "TagSet")
	delete(f, "AccessRegionSet")
	if len(f) > 0 {
		return errors.New("CreateProxyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection Group ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyRequest struct {
	*tchttp.BaseRequest

	// Project ID of connection.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Connection name.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// Access region.
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// Connection bandwidth cap. Unit: Mbps.
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Connection concurrence cap, which indicates the maximum number of simultaneous online connections. Unit: 10,000 connections.
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// Origin server region. If GroupId exists, the origin server region is the one of connection group, and this field is not required. If GroupId does not exist, this field is reuqired.
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	// For more information, please see How to Ensure Idempotence.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Connection group ID. This parameter is required when the connection is created in the connection group. Otherwise, this field is ignored.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// List of tags to be added for connection.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// ID of the replicated connection. Only a running connection can be replicated.
	// The connection is to be replicated if this parameter is set.
	ClonedProxyId *string `json:"ClonedProxyId,omitempty" name:"ClonedProxyId"`

	// Billing mode (0: bill-by-bandwidth, 1: bill-by-traffic. Default value: bill-by-bandwidth)
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`
}

func (r *CreateProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProxyName")
	delete(f, "AccessRegion")
	delete(f, "Bandwidth")
	delete(f, "Concurrent")
	delete(f, "RealServerRegion")
	delete(f, "ClientToken")
	delete(f, "GroupId")
	delete(f, "TagSet")
	delete(f, "ClonedProxyId")
	delete(f, "BillingType")
	if len(f) > 0 {
		return errors.New("CreateProxyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID of connection.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest

	// Layer-7 listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name of the forwarding rule
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Path of the forwarding rule
	Path *string `json:"Path,omitempty" name:"Path"`

	// The origin server type of the forwarding rule, which supports IP and DOMAIN types.
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// Forwarding rules of origin server, which supports round robin (rr), weighted round robin (wrr), and least connections (lc).
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Whether the health check is enabled for rules. 1: enabled; 0: disabled.
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Parameters related to origin server health check
	CheckParams *RuleCheckParams `json:"CheckParams,omitempty" name:"CheckParams"`

	// Protocol types of the forwarding from acceleration connection to origin server, which supports HTTP or HTTPS.
	// If this field is not passed in, it indicates that the ForwardProtocol of the corresponding listener will be used.
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// Remote host to which the acceleration connection forwards. If this parameter is not specified, the default host will be used, i.e., the host with which the client initiates HTTP requests.
	ForwardHost *string `json:"ForwardHost,omitempty" name:"ForwardHost"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "Path")
	delete(f, "RealServerType")
	delete(f, "Scheduler")
	delete(f, "HealthCheck")
	delete(f, "CheckParams")
	delete(f, "ForwardProtocol")
	delete(f, "ForwardHost")
	if len(f) > 0 {
		return errors.New("CreateRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The ID of the successfully created forwarding rule
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// Default policy: ACCEPT or DROP
	DefaultAction *string `json:"DefaultAction,omitempty" name:"DefaultAction"`

	// Acceleration connection ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Connection group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultAction")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("CreateSecurityPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security policy ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRulesRequest struct {
	*tchttp.BaseRequest

	// Security policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// List of access rules
	RuleList []*SecurityPolicyRuleIn `json:"RuleList,omitempty" name:"RuleList" list`
}

func (r *CreateSecurityRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "RuleList")
	if len(f) > 0 {
		return errors.New("CreateSecurityRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of rule IDs
		RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTCPListenersRequest struct {
	*tchttp.BaseRequest

	// Listener name.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// List of listener ports.
	Ports []*uint64 `json:"Ports,omitempty" name:"Ports" list`

	// Origin server scheduling policy of listeners, which supports round robin (rr), weighted round robin (wrr), and least connections (lc).
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Whether origin server has the health check enabled. 1: enabled; 0: disabled. UDP listeners do not support health check.
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// The origin server type of listeners, supporting IP or DOMAIN type. The DOMAIN origin servers do not support the weighted round robin.
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// Connection ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Connection group ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Time interval of origin server health check (unit: seconds). Value range: [5, 300].
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// Response timeout of origin server health check (unit: seconds). Value range: [2, 60]. The timeout value shall be less than the time interval for health check DelayLoop.
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// List of origin server ports, which only supports the listeners of version 1.0 and connection group.
	RealServerPorts []*uint64 `json:"RealServerPorts,omitempty" name:"RealServerPorts" list`

	// Listener methods of getting client IPs. 0: TOA; 1: Proxy Protocol.
	ClientIPMethod *int64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod"`

	// Whether to enable the primary/secondary origin server mode. Valid values: 1 (enable) and 0 (disable). It cannot be enabled for domain name origin servers.
	FailoverSwitch *int64 `json:"FailoverSwitch,omitempty" name:"FailoverSwitch"`
}

func (r *CreateTCPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTCPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Ports")
	delete(f, "Scheduler")
	delete(f, "HealthCheck")
	delete(f, "RealServerType")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	delete(f, "DelayLoop")
	delete(f, "ConnectTimeout")
	delete(f, "RealServerPorts")
	delete(f, "ClientIPMethod")
	delete(f, "FailoverSwitch")
	if len(f) > 0 {
		return errors.New("CreateTCPListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTCPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returns the listener ID
		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTCPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTCPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUDPListenersRequest struct {
	*tchttp.BaseRequest

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// List of listener ports
	Ports []*uint64 `json:"Ports,omitempty" name:"Ports" list`

	// Origin server scheduling policy of listeners, which supports round robin (rr), weighted round robin (wrr), and least connections (lc).
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Origin server type of listeners, which supports IP or DOMAIN type.
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// Connection ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Connection group ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// List of origin server ports, which only supports the listeners of version 1.0 and connection group.
	RealServerPorts []*uint64 `json:"RealServerPorts,omitempty" name:"RealServerPorts" list`
}

func (r *CreateUDPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUDPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerName")
	delete(f, "Ports")
	delete(f, "Scheduler")
	delete(f, "RealServerType")
	delete(f, "ProxyId")
	delete(f, "GroupId")
	delete(f, "RealServerPorts")
	if len(f) > 0 {
		return errors.New("CreateUDPListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUDPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returns the listener ID
		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUDPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUDPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest

	// ID of the certificate to be deleted.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DeleteCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("DeleteCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainErrorPageInfoRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a custom error page. For more information, please see the response to CreateDomainErrorPageInfo.
	ErrorPageId *string `json:"ErrorPageId,omitempty" name:"ErrorPageId"`
}

func (r *DeleteDomainErrorPageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainErrorPageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ErrorPageId")
	if len(f) > 0 {
		return errors.New("DeleteDomainErrorPageInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainErrorPageInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainErrorPageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainErrorPageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name to be deleted
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Whether to make a forced deletion of forwarding rules that have been bound to origin servers. 0: no; 1: yes.
	// When not making a forced deletion, if there are rules bound to origin servers, they will not be deleted.
	Force *uint64 `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "Force")
	if len(f) > 0 {
		return errors.New("DeleteDomainRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenersRequest struct {
	*tchttp.BaseRequest

	// ID list of listeners to be deleted
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds" list`

	// Whether to allow a forced deletion of listeners that have been bound to origin servers. 1: allowed; 0: not allow.
	Force *uint64 `json:"Force,omitempty" name:"Force"`

	// Connection group ID; Either this parameter or `GroupId` must be set, but you cannot set both.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Connection ID; Either this parameter or `GroupId` must be set, but you cannot set both.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DeleteListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerIds")
	delete(f, "Force")
	delete(f, "GroupId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return errors.New("DeleteListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID list of listeners failed to be deleted
		OperationFailedListenerSet []*string `json:"OperationFailedListenerSet,omitempty" name:"OperationFailedListenerSet" list`

		// ID list of listeners deleted successfully
		OperationSucceedListenerSet []*string `json:"OperationSucceedListenerSet,omitempty" name:"OperationSucceedListenerSet" list`

		// ID list of invalid listeners. For example: the listener does not exist, or the instance corresponding to the listener does not match.
		InvalidStatusListenerSet []*string `json:"InvalidStatusListenerSet,omitempty" name:"InvalidStatusListenerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProxyGroupRequest struct {
	*tchttp.BaseRequest

	// ID of the connection group to be deleted.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Whether to enable forced deletion. Valid values:
	// 0: no;
	// 1: yes.
	// Default value: 0. If there is a connection or listener/rule bound to an origin server in the connection group and `Force` is 0, the operation will return a failure.
	Force *uint64 `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Force")
	if len(f) > 0 {
		return errors.New("DeleteProxyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest

	// Layer-7 listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Forwarding rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Whether to make a forced deletion of forwarding rules that have been bound to origin servers. 0: no; 1: yes.
	Force *uint64 `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "RuleId")
	delete(f, "Force")
	if len(f) > 0 {
		return errors.New("DeleteRuleRequest has unknown keys!")
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return errors.New("DeleteSecurityPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityRulesRequest struct {
	*tchttp.BaseRequest

	// Security policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// List of access rule IDs
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList" list`
}

func (r *DeleteSecurityRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return errors.New("DeleteSecurityRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsByDestRegionRequest struct {
	*tchttp.BaseRequest

	// Origin server region: the DescribeDestRegions API returns the value of `RegionId` field of `DestRegionSet`.
	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`
}

func (r *DescribeAccessRegionsByDestRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsByDestRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DestRegion")
	if len(f) > 0 {
		return errors.New("DescribeAccessRegionsByDestRegionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsByDestRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of available acceleration regions
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of available acceleration region information
		AccessRegionSet []*AccessRegionDetial `json:"AccessRegionSet,omitempty" name:"AccessRegionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessRegionsByDestRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsByDestRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccessRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeAccessRegionsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total quantity of acceleration regions
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Acceleration region details list
		AccessRegionSet []*RegionDetail `json:"AccessRegionSet,omitempty" name:"AccessRegionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateDetailRequest struct {
	*tchttp.BaseRequest

	// Certificate ID.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("DescribeCertificateDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate Details.
		CertificateDetail *CertificateDetail `json:"CertificateDetail,omitempty" name:"CertificateDetail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest

	// Certificate type. Where:
	// 0: basic authentication configuration;
	// 1: client CA certificate;
	// 2: server SSL certificate;
	// 3: origin server CA certificate;
	// 4: connection SSL certificate.
	// -1: all types.
	// The default value is -1.
	CertificateType *int64 `json:"CertificateType,omitempty" name:"CertificateType"`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity limit. The default value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeCertificatesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Server certificate list, which includes certificate ID and certificate name.
		CertificateSet []*Certificate `json:"CertificateSet,omitempty" name:"CertificateSet" list`

		// Total quantity of server certificates that match the query conditions.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCountryAreaMappingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCountryAreaMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCountryAreaMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeCountryAreaMappingRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCountryAreaMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Country/region code mapping table
		CountryAreaMappingList []*CountryAreaMap `json:"CountryAreaMappingList,omitempty" name:"CountryAreaMappingList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCountryAreaMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCountryAreaMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDestRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDestRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDestRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeDestRegionsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDestRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of origin server regions
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of origin server region details
		DestRegionSet []*RegionDetail `json:"DestRegionSet,omitempty" name:"DestRegionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDestRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDestRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoByIdsRequest struct {
	*tchttp.BaseRequest

	// List of custom error IDs. Up to 10 IDs are supported
	ErrorPageIds []*string `json:"ErrorPageIds,omitempty" name:"ErrorPageIds" list`
}

func (r *DescribeDomainErrorPageInfoByIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoByIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ErrorPageIds")
	if len(f) > 0 {
		return errors.New("DescribeDomainErrorPageInfoByIdsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoByIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration set of custom error responses
	// Note: this field may return null, indicating that no valid values can be obtained.
		ErrorPageSet []*DomainErrorPageInfo `json:"ErrorPageSet,omitempty" name:"ErrorPageSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainErrorPageInfoByIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoByIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainErrorPageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	if len(f) > 0 {
		return errors.New("DescribeDomainErrorPageInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainErrorPageInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration set of a custom error response
	// Note: This field may return null, indicating that no valid values can be obtained.
		ErrorPageSet []*DomainErrorPageInfo `json:"ErrorPageSet,omitempty" name:"ErrorPageSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainErrorPageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainErrorPageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAndStatisticsProxyRequest struct {
	*tchttp.BaseRequest

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeGroupAndStatisticsProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupAndStatisticsProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("DescribeGroupAndStatisticsProxyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAndStatisticsProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of connection groups that the statistics can be derived from
		GroupSet []*GroupStatisticsInfo `json:"GroupSet,omitempty" name:"GroupSet" list`

		// Connection group quantity
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupAndStatisticsProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupAndStatisticsProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupDomainConfigRequest struct {
	*tchttp.BaseRequest

	// Connection group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DescribeGroupDomainConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Nearest access configuration list of domain name resolution.
		AccessRegionList []*DomainAccessRegionDict `json:"AccessRegionList,omitempty" name:"AccessRegionList" list`

		// Default accesses Ip.
		DefaultDnsIp *string `json:"DefaultDnsIp,omitempty" name:"DefaultDnsIp"`

		// Connection group ID.
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// Total number of configuration of access regions.
		AccessRegionCount *int64 `json:"AccessRegionCount,omitempty" name:"AccessRegionCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPListenersRequest struct {
	*tchttp.BaseRequest

	// Connection ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Filter condition. Exact query by listener IDs.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Filter condition. Exact query by listener names.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Filter condition. Exact query by listener ports.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity limit. The default value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter condition. It supports fuzzy query by ports or listener names. This parameter cannot be used with `ListenerName` or `Port`.
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Connection group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeHTTPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchValue")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DescribeHTTPListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Quantity of listeners
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// HTTP listener list
		ListenerSet []*HTTPListener `json:"ListenerSet,omitempty" name:"ListenerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHTTPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPSListenersRequest struct {
	*tchttp.BaseRequest

	// Filter condition. Connection ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Filter condition. Exact query by listener IDs.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Filter condition. Exact query by listener names.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Filter condition. Exact query by listener ports.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Offset. The default value is 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity limit. The default value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter condition. It supports fuzzy query by ports or listener names.
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Connection group ID as a filter
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeHTTPSListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPSListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchValue")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DescribeHTTPSListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHTTPSListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Quantity of listeners
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// HTTPS listener list
		ListenerSet []*HTTPSListener `json:"ListenerSet,omitempty" name:"ListenerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHTTPSListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPSListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerRealServersRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DescribeListenerRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	if len(f) > 0 {
		return errors.New("DescribeListenerRealServersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of origin servers that can be bound
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// An information list of origin servers
		RealServerSet []*RealServer `json:"RealServerSet,omitempty" name:"RealServerSet" list`

		// Number of bound origin servers
		BindRealServerTotalCount *uint64 `json:"BindRealServerTotalCount,omitempty" name:"BindRealServerTotalCount"`

		// Information list of bound origin servers
		BindRealServerSet []*BindRealServer `json:"BindRealServerSet,omitempty" name:"BindRealServerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenerRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerStatisticsRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical metric name list. It supports:["InBandwidth", "OutBandwidth", "Concurrent", "InPackets", "OutPackets"]
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`

	// Monitoring granularity. It currently supports: 300, 3,600, and 86,400. Unit: seconds.
	// Time range: <= 1 day, supported minimum granularity: 300 seconds;
	// Time range: <= 7 days, supported minimum granularity:3,600 seconds;
	// Time range: > 7 days, supported minimum granularity:86,400 seconds;
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeListenerStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Granularity")
	if len(f) > 0 {
		return errors.New("DescribeListenerStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection group statistics
		StatisticsData []*MetricStatisticsInfo `json:"StatisticsData,omitempty" name:"StatisticsData" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenerStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxiesRequest struct {
	*tchttp.BaseRequest

	// Queries by one or multiple instance IDs. The upper limit on the number of instances for each request is 100. This parameter does not support specifying InstanceIds and Filters at the same time. It's an old parameter, please switch to ProxyIds.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. The default value is 20, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter conditions.   
	// The upper limit on Filters for each request is 10, and the upper limit on Filter.Values is 5. This parameter does not support specifying InstanceIds and Filters at the same time. 
	// ProjectId - String - Required: No - Filter by a project ID.    
	// AccessRegion - String - Required: No - Filter by an access region.    
	// RealServerRegion - String - Required: No - Filter by an origin server region.
	// GroupId - String - Required: No - Filter by a connection group ID.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Queries by one or multiple instance IDs. The upper limit on the number of instances for each request is 100. This parameter does not support specifying InstanceIds and Filters at the same time. It's a new parameter, and replaces InstanceIds.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds" list`

	// Tag list. If this field exists, the list of the resources with the tag will be pulled.
	// It supports up to 5 tags. If there are two or more tags, the connections tagged any of them will be pulled.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// When this field is 1, only not-grouped connections are pulled.
	// When this field is 0, only grouped connections are pulled.
	// When this field does not exist, all connections are pulled, including both not-grouped and grouped connections.
	Independent *int64 `json:"Independent,omitempty" name:"Independent"`
}

func (r *DescribeProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ProxyIds")
	delete(f, "TagSet")
	delete(f, "Independent")
	if len(f) > 0 {
		return errors.New("DescribeProxiesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of connections.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Connection instance information list; It's an old parameter, please switch to ProxySet.
		InstanceSet []*ProxyInfo `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// Connection instance information list; It's a new parameter.
		ProxySet []*ProxyInfo `json:"ProxySet,omitempty" name:"ProxySet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxiesStatusRequest struct {
	*tchttp.BaseRequest

	// Connection ID list; It's an old parameter, please switch to ProxyIds.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Connection ID list; It's a new parameter.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds" list`
}

func (r *DescribeProxiesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return errors.New("DescribeProxiesStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxiesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection status list.
		InstanceStatusSet []*ProxyStatus `json:"InstanceStatusSet,omitempty" name:"InstanceStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxiesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyAndStatisticsListenersRequest struct {
	*tchttp.BaseRequest

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProxyAndStatisticsListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyAndStatisticsListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("DescribeProxyAndStatisticsListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyAndStatisticsListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of connections that the statistics can be derived from
		ProxySet []*ProxySimpleInfo `json:"ProxySet,omitempty" name:"ProxySet" list`

		// Quantity of connections
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyAndStatisticsListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyAndStatisticsListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyDetailRequest struct {
	*tchttp.BaseRequest

	// Connection ID to be queried.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DescribeProxyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	if len(f) > 0 {
		return errors.New("DescribeProxyDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection details
		ProxyDetail *ProxyInfo `json:"ProxyDetail,omitempty" name:"ProxyDetail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupDetailsRequest struct {
	*tchttp.BaseRequest

	// Connection group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeProxyGroupDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DescribeProxyGroupDetailsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection group details
		ProxyGroupDetail *ProxyGroupDetail `json:"ProxyGroupDetail,omitempty" name:"ProxyGroupDetail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyGroupDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupListRequest struct {
	*tchttp.BaseRequest

	// Offset. The default value is 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. The default value is 20. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Project ID. Value range:
	// -1: all projects of this user
	// 0: default project
	// Other values: specified project
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Tag list. If this field exists, the list of the resources with the tag will be pulled.
	// It supports up to 5 tags. If there are two or more tags, the connection groups tagged any of them will be pulled.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// Filter conditions.   
	// The limit on Filter.Values of each request is 5.
	// RealServerRegion - String - Required: No - Filter by origin server region; Refer to the RegionId in the results returned by DescribeDestRegions API.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeProxyGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	delete(f, "TagSet")
	delete(f, "Filters")
	if len(f) > 0 {
		return errors.New("DescribeProxyGroupListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of connection groups.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of connection groups.
	// Note: This field may return null, indicating that no valid values can be obtained.
		ProxyGroupList []*ProxyGroupInfo `json:"ProxyGroupList,omitempty" name:"ProxyGroupList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupStatisticsRequest struct {
	*tchttp.BaseRequest

	// Connection group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical metric name list. Values: InBandwidth (inbound bandwidth); OutBandwidth (outbound bandwidth); Concurrent (concurrence); InPackets (inbound packets); OutPackets (outbound packets).
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`

	// Monitoring granularity. It currently supports: 60, 300, 3,600, 86,400. Unit: seconds.
	// Time range: <= 1 day, supported minimum granularity: 60 seconds;
	// Time range: <= 7 days, supported minimum granularity: 3,600 seconds;
	// Time range: <= 30 days, supported minimum granularity: 86,400 seconds;
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeProxyGroupStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Granularity")
	if len(f) > 0 {
		return errors.New("DescribeProxyGroupStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyGroupStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection group statistics
		StatisticsData []*MetricStatisticsInfo `json:"StatisticsData,omitempty" name:"StatisticsData" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyGroupStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyGroupStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyStatisticsRequest struct {
	*tchttp.BaseRequest

	// Connection ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Start time (2019-03-25 12:00:00)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (2019-03-25 12:00:00)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical metric name list. Valid values: `InBandwidth` (inbound bandwidth); `OutBandwidth` (outbound bandwidth); Concurrent (concurrence); `InPackets` (inbound packets); `OutPackets` (outbound packets); `PacketLoss` (packet loss rate); `Latency` (latency); `HttpQPS` (the number of HTTP requests); `HttpsQPS` (the number of HTTPS requests).
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`

	// Monitoring granularity. It currently supports: 60, 300, 3,600, and 86,400. Unit: seconds.
	// Time range: <= 1 day, supported minimum granularity: 60 seconds;
	// Time range: <= 7 days, supported minimum granularity: 3,600 seconds;
	// Time range: <= 30 days, supported minimum granularity: 86,400 seconds;
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeProxyStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Granularity")
	if len(f) > 0 {
		return errors.New("DescribeProxyStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxyStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection statistics
		StatisticsData []*MetricStatisticsInfo `json:"StatisticsData,omitempty" name:"StatisticsData" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxyStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServerStatisticsRequest struct {
	*tchttp.BaseRequest

	// Origin server ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Layer-7 rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Statistics duration. Unit: hours. It only supports querying statistics for the past 1, 3, 6, 12, and 24 hours.
	WithinTime *uint64 `json:"WithinTime,omitempty" name:"WithinTime"`

	// Statistics start time, such as `2020-08-19 00:00:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Statistics end time, such as `2020-08-19 23:59:59`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistics granularity in seconds. Only 1-minute (60-second) and 5-minute (300-second) granularities are supported.
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeRealServerStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServerStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerId")
	delete(f, "ListenerId")
	delete(f, "RuleId")
	delete(f, "WithinTime")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Granularity")
	if len(f) > 0 {
		return errors.New("DescribeRealServerStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServerStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Origin server status statistics of specified listener
		StatisticsData []*StatisticsDataInfo `json:"StatisticsData,omitempty" name:"StatisticsData" list`

		// Status statistics of multiple origin servers
		RsStatisticsData []*MetricStatisticsInfo `json:"RsStatisticsData,omitempty" name:"RsStatisticsData" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealServerStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServerStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersRequest struct {
	*tchttp.BaseRequest

	// Queries the project ID to which the origin server belongs. -1: all projects.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Origin server IP or domain name to be queried. The fuzzy match is supported.
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// Offset, which is 0 by default.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity of values to return. The default value is 20 and the maximum value is 50.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list. If this field exists, the list of the resources with the tag will be pulled.
	// It supports up to 5 tags. If there are two or more tags, the origin servers tagged any of them will be pulled.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// Filter conditions. The value of the `name` of the `filter` (RealServerName, RealServerIP)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SearchValue")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagSet")
	delete(f, "Filters")
	if len(f) > 0 {
		return errors.New("DescribeRealServersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// An information list of origin server
		RealServerSet []*BindRealServerInfo `json:"RealServerSet,omitempty" name:"RealServerSet" list`

		// The quantity of origin servers
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersStatusRequest struct {
	*tchttp.BaseRequest

	// List of origin server IDs
	RealServerIds []*string `json:"RealServerIds,omitempty" name:"RealServerIds" list`
}

func (r *DescribeRealServersStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerIds")
	if len(f) > 0 {
		return errors.New("DescribeRealServersStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealServersStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Quantity of origin server query results returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Binding status list of origin servers
		RealServerStatusSet []*RealServerStatus `json:"RealServerStatusSet,omitempty" name:"RealServerStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealServersStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealServersStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionAndPriceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionAndPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionAndPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeRegionAndPriceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionAndPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of origin server regions
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of origin server region details
		DestRegionSet []*RegionDetail `json:"DestRegionSet,omitempty" name:"DestRegionSet" list`

		// Connection bandwidth price gradient
		BandwidthUnitPrice []*BandwidthPriceGradient `json:"BandwidthUnitPrice,omitempty" name:"BandwidthUnitPrice" list`

		// Currency type of bandwidth price:
	// CNY (Chinese Yuan)
	// USD (United States Dollar)
		Currency *string `json:"Currency,omitempty" name:"Currency"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionAndPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionAndPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagRequest struct {
	*tchttp.BaseRequest

	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource type, including:
	// Proxy (connection);
	// ProxyGroup (connection group);
	// RealServer (origin server).
	// If this field is not specified, all resources with the tag will be queried.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeResourcesByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return errors.New("DescribeResourcesByTagRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total resources
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Resource list corresponding to the tag
		ResourceSet []*TagResourceInfo `json:"ResourceSet,omitempty" name:"ResourceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleRealServersRequest struct {
	*tchttp.BaseRequest

	// Forwarding rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of values to be returned. The default value is 20. Maximum is 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRuleRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeRuleRealServersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Quantity of origin servers that can be bound
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information list of origin servers that can be bound
		RealServerSet []*RealServer `json:"RealServerSet,omitempty" name:"RealServerSet" list`

		// Quantity of bound origin servers
		BindRealServerTotalCount *uint64 `json:"BindRealServerTotalCount,omitempty" name:"BindRealServerTotalCount"`

		// Bound origin server information list
		BindRealServerSet []*BindRealServer `json:"BindRealServerSet,omitempty" name:"BindRealServerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesByRuleIdsRequest struct {
	*tchttp.BaseRequest

	// List of rule IDs. Up to 10 rules are supported.
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds" list`
}

func (r *DescribeRulesByRuleIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesByRuleIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	if len(f) > 0 {
		return errors.New("DescribeRulesByRuleIdsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesByRuleIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of returned rules.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of returned rules.
		RuleSet []*RuleInfo `json:"RuleSet,omitempty" name:"RuleSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesByRuleIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesByRuleIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest

	// Layer-7 listener ID.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	if len(f) > 0 {
		return errors.New("DescribeRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule information list classified by domain name type
		DomainRuleSet []*DomainRuleSet `json:"DomainRuleSet,omitempty" name:"DomainRuleSet" list`

		// Total quantity of domain names under this listener
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityPolicyDetailRequest struct {
	*tchttp.BaseRequest

	// Security policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeSecurityPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return errors.New("DescribeSecurityPolicyDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

		// Security policy status:
	// BOUND (security policies enabled)
	// UNBIND (security policies disabled)
	// BINDING (enabling security policies)
	// UNBINDING (disabling security policies)
		Status *string `json:"Status,omitempty" name:"Status"`

		// Default policy: ACCEPT or DROP.
		DefaultAction *string `json:"DefaultAction,omitempty" name:"DefaultAction"`

		// Policy ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// List of rules
		RuleList []*SecurityPolicyRuleOut `json:"RuleList,omitempty" name:"RuleList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityRulesRequest struct {
	*tchttp.BaseRequest

	// List of security rule IDs. Up to 20 security rules are supported.
	SecurityRuleIds []*string `json:"SecurityRuleIds,omitempty" name:"SecurityRuleIds" list`
}

func (r *DescribeSecurityRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityRuleIds")
	if len(f) > 0 {
		return errors.New("DescribeSecurityRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of returned security rules.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of returned security rules.
		SecurityRuleSet []*SecurityPolicyRuleOut `json:"SecurityRuleSet,omitempty" name:"SecurityRuleSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTCPListenersRequest struct {
	*tchttp.BaseRequest

	// Connection ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Filter condition. Exact query by listener IDs.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Filter condition. Exact query by listener names.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Filter condition. Exact query by listener ports.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity limit. The default value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Connection group ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Filter condition. It supports fuzzy query by ports or listener names. This parameter cannot be used with `ListenerName` or `Port`.
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeTCPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTCPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return errors.New("DescribeTCPListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTCPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total quantity of listeners that matches the conditions
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// TCP listener list
		ListenerSet []*TCPListener `json:"ListenerSet,omitempty" name:"ListenerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTCPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTCPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUDPListenersRequest struct {
	*tchttp.BaseRequest

	// Connection ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Filter condition. Exact query by listener IDs.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Filter condition. Exact query by listener names.
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Filter condition. Exact query by listener ports.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity limit. The default value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Connection group ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Filter condition. It supports fuzzy query by ports or listener names. This parameter cannot be used with `ListenerName` or `Port`.
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeUDPListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUDPListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "Port")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return errors.New("DescribeUDPListenersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUDPListenersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Quantity of listeners
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// UDP listener list
		ListenerSet []*UDPListener `json:"ListenerSet,omitempty" name:"ListenerSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUDPListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUDPListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyProxiesRequest struct {
	*tchttp.BaseRequest

	// The identifier for forced deletion
	// 1: this connection list is deleted forcibly regardless of whether the origin server has been bound.
	// 0: this connection list cannot be deleted if the origin server has been bound.
	// If this identifier is 0, the deletion can be performed only when all the connections have not been bound to any origin servers.
	Force *int64 `json:"Force,omitempty" name:"Force"`

	// List of connection instance IDs; It's an old parameter, please switch to ProxyIds.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	// For more information, please see How to Ensure Idempotence.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// List of connection instance IDs; It's a new parameter.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds" list`
}

func (r *DestroyProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Force")
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return errors.New("DestroyProxiesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DestroyProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID list of connection instances that cannot be terminated.
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet" list`

		// ID list of connection instances that failed to be terminated.
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAccessRegionDict struct {

	// Nearest access region
	NationCountryInnerList []*NationCountryInnerInfo `json:"NationCountryInnerList,omitempty" name:"NationCountryInnerList" list`

	// Acceleration region connection list
	ProxyList []*ProxyIdDict `json:"ProxyList,omitempty" name:"ProxyList" list`

	// Acceleration region ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Acceleration region internal code
	GeographicalZoneInnerCode *string `json:"GeographicalZoneInnerCode,omitempty" name:"GeographicalZoneInnerCode"`

	// Internal code of the continent to which the acceleration region belongs
	ContinentInnerCode *string `json:"ContinentInnerCode,omitempty" name:"ContinentInnerCode"`

	// Acceleration region alias
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DomainErrorPageInfo struct {

	// Configuration ID of a custom error response
	ErrorPageId *string `json:"ErrorPageId,omitempty" name:"ErrorPageId"`

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Original error code
	ErrorNos []*int64 `json:"ErrorNos,omitempty" name:"ErrorNos" list`

	// New error code
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewErrorNo *int64 `json:"NewErrorNo,omitempty" name:"NewErrorNo"`

	// Response header to be cleared
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClearHeaders []*string `json:"ClearHeaders,omitempty" name:"ClearHeaders" list`

	// Response header to be set
	// Note: This field may return null, indicating that no valid values can be obtained.
	SetHeaders []*HttpHeaderParam `json:"SetHeaders,omitempty" name:"SetHeaders" list`

	// Configured response body (excluding HTTP header)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Body *string `json:"Body,omitempty" name:"Body"`

	// Rule status. 0: success
	// Note: this field may return null, indicating that no valid value is obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DomainRuleSet struct {

	// Forwarding rule domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Forwarding rule list of the domain name.
	RuleSet []*RuleInfo `json:"RuleSet,omitempty" name:"RuleSet" list`

	// Server certificate ID of the domain. When it is `default`, it indicates that the default certificate will be used (i.e., the certificate configured for the listener).
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Server certificate name of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// Client certificate ID of the domain. When it is `default`, it indicates that the default certificate will be used (i.e., the certificate configured for the listener).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// Client certificate name of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientCertificateAlias *string `json:"ClientCertificateAlias,omitempty" name:"ClientCertificateAlias"`

	// Basic authentication configuration ID of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BasicAuthConfId *string `json:"BasicAuthConfId,omitempty" name:"BasicAuthConfId"`

	// Basic authentication status:
	// 0: disabled;
	// 1: enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BasicAuth *int64 `json:"BasicAuth,omitempty" name:"BasicAuth"`

	// Basic authentication configuration name of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BasicAuthConfAlias *string `json:"BasicAuthConfAlias,omitempty" name:"BasicAuthConfAlias"`

	// Origin server authentication certificate ID of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerCertificateId *string `json:"RealServerCertificateId,omitempty" name:"RealServerCertificateId"`

	// Origin server authentication status:
	// 0: disabled;
	// 1: enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerAuth *int64 `json:"RealServerAuth,omitempty" name:"RealServerAuth"`

	// Origin server authentication certificate name of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerCertificateAlias *string `json:"RealServerCertificateAlias,omitempty" name:"RealServerCertificateAlias"`

	// Connection authentication certificate ID of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GaapCertificateId *string `json:"GaapCertificateId,omitempty" name:"GaapCertificateId"`

	// Connection authentication status:
	// 0: disabled;
	// 1: enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GaapAuth *int64 `json:"GaapAuth,omitempty" name:"GaapAuth"`

	// Connection authentication certificate name of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GaapCertificateAlias *string `json:"GaapCertificateAlias,omitempty" name:"GaapCertificateAlias"`

	// Origin server authentication domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerCertificateDomain *string `json:"RealServerCertificateDomain,omitempty" name:"RealServerCertificateDomain"`

	// Returns IDs and aliases of multiple certificates when there are multiple client certificates.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolyClientCertificateAliasInfo []*CertificateAliasInfo `json:"PolyClientCertificateAliasInfo,omitempty" name:"PolyClientCertificateAliasInfo" list`

	// Returns IDs and aliases of multiple certificates when there are multiple origin certificates.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolyRealServerCertificateAliasInfo []*CertificateAliasInfo `json:"PolyRealServerCertificateAliasInfo,omitempty" name:"PolyRealServerCertificateAliasInfo" list`

	// Domain name status.
	// 0: running;
	// 1: changing;
	// 2: deleting.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DomainStatus *uint64 `json:"DomainStatus,omitempty" name:"DomainStatus"`
}

type Filter struct {

	// Filter conditions
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type GroupStatisticsInfo struct {

	// Connection group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Connection group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// List of connections of a connection group
	ProxySet []*ProxySimpleInfo `json:"ProxySet,omitempty" name:"ProxySet" list`
}

type HTTPListener struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Listener creation time; using UNIX timestamp.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Listener protocol. Valid values: HTTP, HTTPS. The value `HTTP` is used for this structure
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener status:
	// 0: running;
	// 1: creating;
	// 2: terminating;
	// 3: adjusting origin server;
	// 4: modifying configuration.
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`
}

type HTTPSListener struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Listener protocol. Valid values: HTTP, HTTPS. The value `HTTPS` is used for this structure
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener status:
	// 0: running;
	// 1: creating;
	// 2: terminating;
	// 3: adjusting origin server;
	// 4: modifying configuration.
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`

	// Server SSL certificate ID of the listener
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Protocol used in the forwarding from connections to origin servers
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// Listener creation time; using UNIX timestamp.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Server SSL certificate alias
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`

	// Client CA certificate ID of the listener
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// Listener authentication mode. Valid values:
	// 0: one-way authentication;
	// 1: mutual authentication.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`

	// Client CA certificate alias
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientCertificateAlias *string `json:"ClientCertificateAlias,omitempty" name:"ClientCertificateAlias"`

	// Alias information of multiple client CA certificates.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolyClientCertificateAliasInfo []*CertificateAliasInfo `json:"PolyClientCertificateAliasInfo,omitempty" name:"PolyClientCertificateAliasInfo" list`
}

type HttpHeaderParam struct {

	// HTTP header name
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// HTTP header value
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
}

type InquiryPriceCreateProxyRequest struct {
	*tchttp.BaseRequest

	// Acceleration region name.
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// Connection bandwidth cap. Unit: Mbps.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Origin server region name. It's an old parameter, please switch to RealServerRegion.
	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`

	// Upper limit of connection concurrence, which indicates a number of simultaneous online connections. Unit: 10,000 connections. It's an old parameter, please switch to Concurrent.
	Concurrency *int64 `json:"Concurrency,omitempty" name:"Concurrency"`

	// Origin server region name; It's a new parameter.
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// Upper limit of connection concurrence, which indicates a number of simultaneous online connections. Unit: 10,000 connections. It's a new parameter.
	Concurrent *int64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// Billing mode. Valid values: 0: bill-by-bandwidth (default value); 1: bill-by-traffic.
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`
}

func (r *InquiryPriceCreateProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRegion")
	delete(f, "Bandwidth")
	delete(f, "DestRegion")
	delete(f, "Concurrency")
	delete(f, "RealServerRegion")
	delete(f, "Concurrent")
	delete(f, "BillingType")
	if len(f) > 0 {
		return errors.New("InquiryPriceCreateProxyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateProxyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Basic price of connection in USD/day.
		ProxyDailyPrice *float64 `json:"ProxyDailyPrice,omitempty" name:"ProxyDailyPrice"`

		// Tiered price of connection bandwidth.
	// Note: this field may return null, indicating that no valid values can be obtained.
		BandwidthUnitPrice []*BandwidthPriceGradient `json:"BandwidthUnitPrice,omitempty" name:"BandwidthUnitPrice" list`

		// Discounted basic price of connection in USD/day.
		DiscountProxyDailyPrice *float64 `json:"DiscountProxyDailyPrice,omitempty" name:"DiscountProxyDailyPrice"`

		// Currency, which supports CNY, USD, etc.
		Currency *string `json:"Currency,omitempty" name:"Currency"`

		// Connection traffic price in USD/GB.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FlowUnitPrice *float64 `json:"FlowUnitPrice,omitempty" name:"FlowUnitPrice"`

		// Discounted connection traffic price in USD/GB.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DiscountFlowUnitPrice *float64 `json:"DiscountFlowUnitPrice,omitempty" name:"DiscountFlowUnitPrice"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListenerInfo struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listening port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Listener protocol type
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type MetricStatisticsInfo struct {

	// Metric name
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Metric statistics
	MetricData []*StatisticsDataInfo `json:"MetricData,omitempty" name:"MetricData" list`
}

type ModifyCertificateAttributesRequest struct {
	*tchttp.BaseRequest

	// Certificate ID.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Certificate name. Up to 50 characters.
	CertificateAlias *string `json:"CertificateAlias,omitempty" name:"CertificateAlias"`
}

func (r *ModifyCertificateAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "CertificateAlias")
	if len(f) > 0 {
		return errors.New("ModifyCertificateAttributesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateRequest struct {
	*tchttp.BaseRequest

	// Listener instance ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Domain name whose certificate needs to be modified
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// New server certificate ID:
	// If CertificateId=default, using the listener certificate.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// New client certificate ID:
	// If ClientCertificateId=default, using the listener certificate.
	// This parameter is required only when the mutual authentication is adopted.
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// List of new IDs of multiple client certificates, where:
	// This parameter or the `ClientCertificateId` parameter is required for mutual authentication only.
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds" list`
}

func (r *ModifyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	if len(f) > 0 {
		return errors.New("ModifyCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest

	// Layer-7 listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Original domain name information
	OldDomain *string `json:"OldDomain,omitempty" name:"OldDomain"`

	// New domain name information
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`

	// Server SSL certificate ID. It's only applicable to the connections of version 3.0:
	// If this field is not passed in, the original certificate will be used;
	// If this field is passed in, and CertificateId=default, the listener certificate will be used;
	// For other cases, the certificate specified by CertificateId will be used.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Client CA certificate ID. It's only applicable to the connections of version 3.0:
	// If this field is not passed in, the original certificate will be used;
	// If this field is passed in, and ClientCertificateId=default, the listener certificate will be used;
	// For other cases, the certificate specified by ClientCertificateId will be used.
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// Client CA certificate ID. It is only applicable to connections on version 3.0, where:
	// If this field and `ClientCertificateId` are not included, the original certificate will be used;
	// If this field is included, and ClientCertificateId=default, then the listener certificate will be used;
	// In other cases, the certificate specified by `ClientCertificateId` or `PolyClientCertificateIds` will be used.
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds" list`
}

func (r *ModifyDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "OldDomain")
	delete(f, "NewDomain")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	if len(f) > 0 {
		return errors.New("ModifyDomainRequest has unknown keys!")
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupDomainConfigRequest struct {
	*tchttp.BaseRequest

	// Connection group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Default access IP or domain name of domain name resolution
	DefaultDnsIp *string `json:"DefaultDnsIp,omitempty" name:"DefaultDnsIp"`

	// Nearest access region configuration.
	AccessRegionList []*AccessRegionDomainConf `json:"AccessRegionList,omitempty" name:"AccessRegionList" list`
}

func (r *ModifyGroupDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DefaultDnsIp")
	delete(f, "AccessRegionList")
	if len(f) > 0 {
		return errors.New("ModifyGroupDomainConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// Listener ID to be modified
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// New listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Connection ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *ModifyHTTPListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return errors.New("ModifyHTTPListenerAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHTTPListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPSListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Connection ID. This field is required if using a single connection listener.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// New listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Type of the protocol used in the forwarding from connections to origin servers
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// New listener server certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// New listener client certificate ID
	ClientCertificateId *string `json:"ClientCertificateId,omitempty" name:"ClientCertificateId"`

	// Client certificate ID of the listener after modification, which is a new field.
	PolyClientCertificateIds []*string `json:"PolyClientCertificateIds,omitempty" name:"PolyClientCertificateIds" list`
}

func (r *ModifyHTTPSListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPSListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "ProxyId")
	delete(f, "ListenerName")
	delete(f, "ForwardProtocol")
	delete(f, "CertificateId")
	delete(f, "ClientCertificateId")
	delete(f, "PolyClientCertificateIds")
	if len(f) > 0 {
		return errors.New("ModifyHTTPSListenerAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHTTPSListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHTTPSListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPSListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesAttributeRequest struct {
	*tchttp.BaseRequest

	// ID of one or multiple connections to be operated; It's an old parameter, please switch to ProxyIds.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Connection name. Up to 30 characters.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	// For more information, please see How to Ensure Idempotence.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// ID of one or multiple connections to be operated; It's a new parameter.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds" list`
}

func (r *ModifyProxiesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProxyName")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return errors.New("ModifyProxiesAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxiesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesProjectRequest struct {
	*tchttp.BaseRequest

	// The target project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID of one or multiple connections to be operated; It's an old parameter, please switch to ProxyIds.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	// For more information, please see How to Ensure Idempotence.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// ID of one or multiple connections to be operated; It's a new parameter.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds" list`
}

func (r *ModifyProxiesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return errors.New("ModifyProxiesProjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxiesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxiesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxiesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyConfigurationRequest struct {
	*tchttp.BaseRequest

	// Connection instance ID; It's an old parameter, please switch to ProxyId.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Target bandwidth. Unit: Mbps.
	// Bandwidth or Concurrent must be set. Use the DescribeAccessRegionsByDestRegion API to obtain the value range.
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Target concurrence value. Unit: 10,000 connections.
	// Bandwidth or Concurrent must be set. Use the DescribeAccessRegionsByDestRegion API to obtain the value range.
	Concurrent *uint64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	// For more information, please see How to Ensure Idempotence.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Connection instance ID; It's a new parameter.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Billing mode (0: bill-by-bandwidth, 1: bill-by-traffic. Default value: bill-by-bandwidth)
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`
}

func (r *ModifyProxyConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "Concurrent")
	delete(f, "ClientToken")
	delete(f, "ProxyId")
	delete(f, "BillingType")
	if len(f) > 0 {
		return errors.New("ModifyProxyConfigurationRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxyConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// ID of the connection group to be modified.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// New connection group name. Up to 30 characters. The extra characters will be truncated.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyProxyGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("ModifyProxyGroupAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProxyGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProxyGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerNameRequest struct {
	*tchttp.BaseRequest

	// Origin server name
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// Origin server ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`
}

func (r *ModifyRealServerNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealServerNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerName")
	delete(f, "RealServerId")
	if len(f) > 0 {
		return errors.New("ModifyRealServerNameRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRealServerNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealServerNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleAttributeRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Forwarding rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Scheduling policy:
	// rr: round robin;
	// wrr: weighted round robin;
	// lc: least connections.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Whether to enable the origin server health check:
	// 1: enable;
	// 0: disable.
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Health check configuration parameters
	CheckParams *RuleCheckParams `json:"CheckParams,omitempty" name:"CheckParams"`

	// Forwarding rule path
	Path *string `json:"Path,omitempty" name:"Path"`

	// Protocol types of the forwarding from acceleration connection to origin server, which supports default, HTTP and HTTPS.
	// If `ForwardProtocol=default`, the `ForwardProtocol` of the listener will be used.
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// The `host` carried in the request forwarded from the acceleration connection to the origin server.
	// If `ForwardHost=default`, the domain name of rule will be used. For other cases, the value set in this field will be used.
	ForwardHost *string `json:"ForwardHost,omitempty" name:"ForwardHost"`
}

func (r *ModifyRuleAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "RuleId")
	delete(f, "Scheduler")
	delete(f, "HealthCheck")
	delete(f, "CheckParams")
	delete(f, "Path")
	delete(f, "ForwardProtocol")
	delete(f, "ForwardHost")
	if len(f) > 0 {
		return errors.New("ModifyRuleAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityRuleRequest struct {
	*tchttp.BaseRequest

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Rule name: up to 30 characters. The extra characters will be truncated.
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// Security policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Security rule action
	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`

	// A CIDR IP address associated with the rule
	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`

	// Protocol type
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port range. Valid values:
	// A single port: 80
	// Multiple ports: 80 and 443
	// Consecutive ports: 3306-20000
	// All ports: ALL
	DestPortRange *string `json:"DestPortRange,omitempty" name:"DestPortRange"`
}

func (r *ModifySecurityRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "AliasName")
	delete(f, "PolicyId")
	delete(f, "RuleAction")
	delete(f, "SourceCidr")
	delete(f, "Protocol")
	delete(f, "DestPortRange")
	if len(f) > 0 {
		return errors.New("ModifySecurityRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTCPListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Connection group ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Connection ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Origin server scheduling policy of listeners, which supports round robin (rr), weighted round robin (wrr), and least connections (lc).
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Time interval of origin server health check (unit: seconds). Value range: [5, 300].
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// Response timeout of origin server health check (unit: seconds). Value range: [2, 60]. The timeout value shall be less than the time interval for health check DelayLoop.
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// Whether to enable health check. 1: enable; 0: disable.
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Whether to enable the primary/secondary origin server mode. Valid values: 1 (enable) and 0 (disable). It cannot be enabled for domain name origin servers.
	FailoverSwitch *uint64 `json:"FailoverSwitch,omitempty" name:"FailoverSwitch"`
}

func (r *ModifyTCPListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTCPListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "GroupId")
	delete(f, "ProxyId")
	delete(f, "ListenerName")
	delete(f, "Scheduler")
	delete(f, "DelayLoop")
	delete(f, "ConnectTimeout")
	delete(f, "HealthCheck")
	delete(f, "FailoverSwitch")
	if len(f) > 0 {
		return errors.New("ModifyTCPListenerAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTCPListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTCPListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTCPListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUDPListenerAttributeRequest struct {
	*tchttp.BaseRequest

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Connection group ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Connection ID; Either `ProxyId` or `GroupId` must be set, but you cannot set both.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Origin server scheduling policy of listeners
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
}

func (r *ModifyUDPListenerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUDPListenerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "GroupId")
	delete(f, "ProxyId")
	delete(f, "ListenerName")
	delete(f, "Scheduler")
	if len(f) > 0 {
		return errors.New("ModifyUDPListenerAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUDPListenerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUDPListenerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUDPListenerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NationCountryInnerInfo struct {

	// Country name
	NationCountryName *string `json:"NationCountryName,omitempty" name:"NationCountryName"`

	// Country internal code
	NationCountryInnerCode *string `json:"NationCountryInnerCode,omitempty" name:"NationCountryInnerCode"`
}

type NewRealServer struct {

	// Origin server ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// Origin server IP or domain name
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`
}

type OpenProxiesRequest struct {
	*tchttp.BaseRequest

	// List of connection instance IDs; It's an old parameter, please switch to ProxyIds.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	// For more information, please see How to Ensure Idempotence.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// List of connection instance IDs; It's a new parameter.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds" list`
}

func (r *OpenProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ClientToken")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return errors.New("OpenProxiesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenProxiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The connection instance ID list cannot be enabled if it's not disabled.
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet" list`

		// ID list of connection instances failed to be enabled.
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProxyGroupRequest struct {
	*tchttp.BaseRequest

	// Connection group instance ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *OpenProxyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("OpenProxyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenProxyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of IDs of the connection instances that are not disabled, which cannot be enabled.
		InvalidStatusInstanceSet []*string `json:"InvalidStatusInstanceSet,omitempty" name:"InvalidStatusInstanceSet" list`

		// List of IDs of the connection instances failed to be enabled.
		OperationFailedInstanceSet []*string `json:"OperationFailedInstanceSet,omitempty" name:"OperationFailedInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProxyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProxyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// ID of the connections requiring enabled security policies.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Security policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *OpenSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return errors.New("OpenSecurityPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async Process ID. Using DescribeAsyncTaskStatus to query process and status.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyGroupDetail struct {

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of connections in connection group
	ProxyNum *int64 `json:"ProxyNum,omitempty" name:"ProxyNum"`

	// Connection group status:
	// 0: running normally;
	// 1: creating;
	// 4: terminating;
	// 11: migrating;
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Owner UIN
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Creation UIN
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// Connection name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Default IP of domain name resolution for connection groups
	DnsDefaultIp *string `json:"DnsDefaultIp,omitempty" name:"DnsDefaultIp"`

	// Connection group domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Target region
	RealServerRegionInfo *RegionDetail `json:"RealServerRegionInfo,omitempty" name:"RealServerRegionInfo"`

	// Whether it is an old connection group, i.e., those created before August 3, 2018.
	IsOldGroup *bool `json:"IsOldGroup,omitempty" name:"IsOldGroup"`

	// Connection group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// Security policy ID. This field exists if security policies are set.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Connection group version
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Describes how the connection obtains client IPs. 0: TOA; 1: Proxy Protocol.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClientIPMethod []*int64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod" list`
}

type ProxyGroupInfo struct {

	// Connection group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Connection group domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Connection group name
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Target region
	RealServerRegionInfo *RegionDetail `json:"RealServerRegionInfo,omitempty" name:"RealServerRegionInfo"`

	// Connection group status.
	// Where:
	// 0: running;
	// 1: creating;
	// 4: terminating;
	// 11: connection migrating.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Tag list.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// Connection group version
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether the connection group contains a Microsoft connection
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProxyType *uint64 `json:"ProxyType,omitempty" name:"ProxyType"`
}

type ProxyIdDict struct {

	// Connection ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type ProxyInfo struct {

	// Connection instance ID; It's an old parameter, please switch to ProxyId.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Creation time in the format of UNIX timestamp, indicating the number of seconds that have elapsed since January 1, 1970 (midnight in UTC/GMT).
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Connection name.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// Access region.
	AccessRegion *string `json:"AccessRegion,omitempty" name:"AccessRegion"`

	// Origin server region.
	RealServerRegion *string `json:"RealServerRegion,omitempty" name:"RealServerRegion"`

	// Bandwidth. Unit: Mbps.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Concurrence. Unit: requests/second.
	Concurrent *int64 `json:"Concurrent,omitempty" name:"Concurrent"`

	// Connection status:
	// RUNNING: running;
	// CREATING: creating;
	// DESTROYING: terminating;
	// OPENING: enabling;
	// CLOSING: disabling;
	// CLOSED: disabled;
	// ADJUSTING: adjusting configuration
	// ISOLATING: isolating (it's triggered when the account is in arrears);
	// ISOLATED: isolated (it's triggered when the account is in arrears);
	// UNKNOWN: unknown status.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Accessed domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Accessed IP.
	IP *string `json:"IP,omitempty" name:"IP"`

	// Connection versions: 1.0, 2.0, 3.0.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Connection instance ID; It's a new parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// 1: this connection is expandable; 0: this connection is not expandable.
	Scalarable *int64 `json:"Scalarable,omitempty" name:"Scalarable"`

	// Supported protocol types.
	SupportProtocols []*string `json:"SupportProtocols,omitempty" name:"SupportProtocols" list`

	// Connection group ID. This field exists if a connection belongs to a connection group.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Security policy ID. This field exists if security policies are configured.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Access region details, including region ID and region name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessRegionInfo *RegionDetail `json:"AccessRegionInfo,omitempty" name:"AccessRegionInfo"`

	// Origin server region details, including region ID and region name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerRegionInfo *RegionDetail `json:"RealServerRegionInfo,omitempty" name:"RealServerRegionInfo"`

	// Forwarding IP of the connection
	ForwardIP *string `json:"ForwardIP,omitempty" name:"ForwardIP"`

	// Tag list. This field is an empty list if no tags exist.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSet []*TagPair `json:"TagSet,omitempty" name:"TagSet" list`

	// Whether security groups are supported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SupportSecurity *int64 `json:"SupportSecurity,omitempty" name:"SupportSecurity"`

	// Billing mode. 0: bill-by-bandwidth; 1: bill-by-traffic.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BillingType *int64 `json:"BillingType,omitempty" name:"BillingType"`

	// List of domain names associated with resolution record
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelatedGlobalDomains []*string `json:"RelatedGlobalDomains,omitempty" name:"RelatedGlobalDomains" list`

	// Configuration change time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifyConfigTime *uint64 `json:"ModifyConfigTime,omitempty" name:"ModifyConfigTime"`

	// Connection type. 104: SILVER connection.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProxyType *uint64 `json:"ProxyType,omitempty" name:"ProxyType"`

	// Describes how the connection obtains client IPs. 0: TOA; 1: Proxy Protocol.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClientIPMethod []*int64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod" list`
}

type ProxySimpleInfo struct {

	// Connection ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Connection name
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// Listener list
	ListenerList []*ListenerInfo `json:"ListenerList,omitempty" name:"ListenerList" list`
}

type ProxyStatus struct {

	// Connection instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Connection status.
	// Valid values:
	// RUNNING: running;
	// CREATING: creating;
	// DESTROYING: terminating;
	// OPENING: enabling;
	// CLOSING: disabling;
	// CLOSED: disabled;
	// ADJUSTING: adjusting configuration;
	// ISOLATING: isolating;
	// ISOLATED: isolated;
	// UNKNOWN: unknown status.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type RealServer struct {

	// Origin server IP or domain name
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// Origin server ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// Origin server name
	RealServerName *string `json:"RealServerName,omitempty" name:"RealServerName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type RealServerBindSetReq struct {

	// Origin server ID
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// Origin server port
	RealServerPort *uint64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// Origin server IP
	RealServerIP *string `json:"RealServerIP,omitempty" name:"RealServerIP"`

	// Origin server weight
	RealServerWeight *uint64 `json:"RealServerWeight,omitempty" name:"RealServerWeight"`

	// Origin server role: master (primary origin server); slave (secondary origin server). This parameter is applicable when the primary/secondary origin server mode is enabled for a TCP listener.
	RealServerFailoverRole *string `json:"RealServerFailoverRole,omitempty" name:"RealServerFailoverRole"`
}

type RealServerStatus struct {

	// Origin server ID.
	RealServerId *string `json:"RealServerId,omitempty" name:"RealServerId"`

	// 0: not bound, 1: bound to rule or listener.
	BindStatus *int64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// ID of the connection bound to this origin server. This string is empty if they are not bound.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type RegionDetail struct {

	// Region ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Region name in Chinese or English
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type RemoveRealServersRequest struct {
	*tchttp.BaseRequest

	// List of origin server IDs
	RealServerIds []*string `json:"RealServerIds,omitempty" name:"RealServerIds" list`
}

func (r *RemoveRealServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveRealServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealServerIds")
	if len(f) > 0 {
		return errors.New("RemoveRealServersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveRealServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveRealServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveRealServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleCheckParams struct {

	// Time interval of health check
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// Response timeout of health check
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// Check path of health check
	Path *string `json:"Path,omitempty" name:"Path"`

	// Health check method: GET/HEAD
	Method *string `json:"Method,omitempty" name:"Method"`

	// Return code indicting normal origin servers. Value range: [100, 200, 300, 400, 500]
	StatusCode []*uint64 `json:"StatusCode,omitempty" name:"StatusCode" list`

	// Domain name to be performed health check
	// You cannot modify this parameter when calling ModifyRuleAttribute API.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Origin server failure check frequency
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FailedCountInter *uint64 `json:"FailedCountInter,omitempty" name:"FailedCountInter"`

	// Origin server health check threshold. All requests to the origin server will be blocked once the threshold is exceeded.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FailedThreshold *uint64 `json:"FailedThreshold,omitempty" name:"FailedThreshold"`

	// Duration to block requests targeting the origin server after a failed health check
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BlockInter *uint64 `json:"BlockInter,omitempty" name:"BlockInter"`
}

type RuleInfo struct {

	// Rule information
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Listener information
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Rule domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Rule path
	Path *string `json:"Path,omitempty" name:"Path"`

	// Origin server type
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// Forwarding policy of the origin server
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Whether health check is enabled. 1: enabled, 0: disabled
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Rule status. 0: running, 1: creating, 2: terminating, 3: binding/unbinding origin server, 4: updating configuration
	RuleStatus *uint64 `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// Health check parameters
	CheckParams *RuleCheckParams `json:"CheckParams,omitempty" name:"CheckParams"`

	// Bound origin server information
	RealServerSet []*BindRealServer `json:"RealServerSet,omitempty" name:"RealServerSet" list`

	// Origin server service status. 0: exceptional, 1: normal
	// If health check is not enabled, this status will always be normal.
	// As long as one origin server is exceptional, this status will be exceptional. Please view `RealServerSet` for the status of specific origin servers.
	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// The `host` carried in the request forwarded from the connection to the origin server. `default` indicates directly forwarding the received 'host'.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ForwardHost *string `json:"ForwardHost,omitempty" name:"ForwardHost"`
}

type SecurityPolicyRuleIn struct {

	// Source IP or IP range of the request.
	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`

	// Policy: Allow (ACCEPT) or reject (DROP).
	Action *string `json:"Action,omitempty" name:"Action"`

	// Rule alias
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// Protocol: TCP or UDP. ALL indicates all protocols.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Target port. Formatting examples:
	// Single port: 80
	// Multiple ports: 80, 443
	// Consecutive ports: 3306-20000
	// All ports: ALL
	DestPortRange *string `json:"DestPortRange,omitempty" name:"DestPortRange"`
}

type SecurityPolicyRuleOut struct {

	// Policy: Allow (ACCEPT) or reject (DROP).
	Action *string `json:"Action,omitempty" name:"Action"`

	// Source IP or IP range of the request.
	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`

	// Rule alias
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// Target port range
	// Note: This field may return null, indicating that no valid values can be obtained.
	DestPortRange *string `json:"DestPortRange,omitempty" name:"DestPortRange"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Protocol type to be matched (TCP/UDP)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Security policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type SetAuthenticationRequest struct {
	*tchttp.BaseRequest

	// Listener ID.
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// The domain name requiring advanced configuration, i.e., the domain name of the listener's forwarding rules.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Whether to enable the basic authentication:
	// 0: disable basic authentication;
	// 1: enable basic authentication.
	// The default value is 0.
	BasicAuth *int64 `json:"BasicAuth,omitempty" name:"BasicAuth"`

	// Whether to enable the connection authentication, which is for the origin server to authenticate GAAP.
	// 0: disable;
	// 1: enable.
	// The default value is 0.
	GaapAuth *int64 `json:"GaapAuth,omitempty" name:"GaapAuth"`

	// Whether to enable the origin server authentication, which is for GAAP to authenticate the server.
	// 0: disable;
	// 1: enable.
	// The default value is 0.
	RealServerAuth *int64 `json:"RealServerAuth,omitempty" name:"RealServerAuth"`

	// Basic authentication configuration ID, which is obtained from the certificate management page.
	BasicAuthConfId *string `json:"BasicAuthConfId,omitempty" name:"BasicAuthConfId"`

	// Connection SSL certificate ID, which is obtained from the certificate management page.
	GaapCertificateId *string `json:"GaapCertificateId,omitempty" name:"GaapCertificateId"`

	// CA certificate ID of the origin server, which is obtained from the certificate management page. When authenticating the origin server, enter this parameter or the `RealServerCertificateIds` parameter.
	RealServerCertificateId *string `json:"RealServerCertificateId,omitempty" name:"RealServerCertificateId"`

	// Domain name of the origin server certificate.
	RealServerCertificateDomain *string `json:"RealServerCertificateDomain,omitempty" name:"RealServerCertificateDomain"`

	// CA certificate IDs of multiple origin servers, which are obtained from the certificate management page. When authenticating the origin servers, enter this parameter or the `RealServerCertificateId` parameter.
	PolyRealServerCertificateIds []*string `json:"PolyRealServerCertificateIds,omitempty" name:"PolyRealServerCertificateIds" list`
}

func (r *SetAuthenticationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAuthenticationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "BasicAuth")
	delete(f, "GaapAuth")
	delete(f, "RealServerAuth")
	delete(f, "BasicAuthConfId")
	delete(f, "GaapCertificateId")
	delete(f, "RealServerCertificateId")
	delete(f, "RealServerCertificateDomain")
	delete(f, "PolyRealServerCertificateIds")
	if len(f) > 0 {
		return errors.New("SetAuthenticationRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAuthenticationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAuthenticationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAuthenticationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticsDataInfo struct {

	// Corresponding time point
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Statistics value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type TCPListener struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Origin server port, which is only valid for the connections of version 1.0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerPort *uint64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// Type of the origin server bound to listeners
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// Listener protocol: TCP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener status. Valid values:
	// 0: running;
	// 1: creating;
	// 2: terminating;
	// 3: adjusting origin server;
	// 4: adjusting configuration.
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`

	// Origin server access policy of listener. Valid values:
	// rr: round robin;
	// wrr: weighted round robin;
	// lc: least connection.
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Response timeout of origin server health check (unit: seconds).
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// Time interval of origin server health check (unit: seconds).
	DelayLoop *uint64 `json:"DelayLoop,omitempty" name:"DelayLoop"`

	// Whether health check is enabled for listener. Valid values:
	// 0: disabled;
	// 1: enabled
	HealthCheck *uint64 `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Status of origin server bound to listener. Valid values:
	// 0: exceptional;
	// 1: normal.
	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// Information of the origin server bound to listeners
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerSet []*BindRealServer `json:"RealServerSet,omitempty" name:"RealServerSet" list`

	// Listener creation time; using UNIX timestamp.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Describes how the listener obtains client IPs. 0: TOA; 1: Proxy Protocol.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClientIPMethod *uint64 `json:"ClientIPMethod,omitempty" name:"ClientIPMethod"`
}

type TagPair struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagResourceInfo struct {

	// Resource types:
	// `Proxy`: connections;
	// `ProxyGroup`: connection groups;
	// `RealServer`: origin servers.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type UDPListener struct {

	// Listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// Listener name
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// Listener port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Origin server port, which is only valid for the connections or connection groups of version 1.0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealServerPort *uint64 `json:"RealServerPort,omitempty" name:"RealServerPort"`

	// Type of the origin server bound to listeners
	RealServerType *string `json:"RealServerType,omitempty" name:"RealServerType"`

	// Listener protocol: UDP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Listener status. Valid values:
	// 0: running;
	// 1: creating;
	// 2: terminating;
	// 3: adjusting origin server;
	// 4: adjusting configuration.
	ListenerStatus *uint64 `json:"ListenerStatus,omitempty" name:"ListenerStatus"`

	// Origin server access policy of listeners
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Status of origin server bound to listener. 0: normal, 1: exceptional IP, 2: exceptional domain name resolution
	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// Information of the origin server bound to listeners
	RealServerSet []*BindRealServer `json:"RealServerSet,omitempty" name:"RealServerSet" list`

	// Listener creation time; using UNIX timestamp.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}
