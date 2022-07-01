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

package v20220106

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ApplicationProxy struct {
	// Instance ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Instance name
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// Scheduling mode:
	// `ip`: Anycast IP
	// `domain`: CNAME
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// `0`: Disable security protection; `1`: Enable security protection.
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// `0`: Disable acceleration; `1`: Enable acceleration.
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// This field is moved to `Rule.ForwardClientIp`.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This field is moved to `Rule.SessionPersist`.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// Rule list
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// Status:
	// `online`: Enable
	// `offline`: Disable
	// `progress`: Deploying
	// `stopping`: Disabling
	// `fail`: Deployment/Disabling failed
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Scheduling information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

	// Update time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Site ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Session persistence duration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Subdomain name
	// `instance`: Instance
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// ID of the layer-7 domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type ApplicationProxyRule struct {
	// Protocol. Valid values: `TCP` and `UDP`.
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Port. Valid values:
	// `80`: Port 80
	// `81-90`: Port range 81-90
	Port []*string `json:"Port,omitempty" name:"Port"`

	// Origin server type. Valid values:
	// `custom`: Specified origins
	// `origins`: An origin group
	// `load_balancing`: A load balancer
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information.
	// When `OriginType=custom`, this field value indicates multiple origin servers in either of the following formats:
	// IP:Port
	// Domain name:Port
	// When `OriginType=origins`, it indicates the origin group ID.
	//  
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Rule ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Status:
	// `online`: Enable
	// `offline`: Disable
	// `progress`: Deploying
	// `stopping`: Disabling
	// `fail`: Deployment/Disabling failed
	Status *string `json:"Status,omitempty" name:"Status"`

	// Passes the client IP. When `Proto=TCP`, valid values:
	// `TOA`: Pass the client IP via TOA
	// `PPV1`: Pass the client IP via Proxy Protocol V1
	// `PPV2`: Pass the client IP via Proxy Protocol V2
	// `OFF`: Do not pass the client IP.
	// When `Proto=UDP`, valid values:
	// `PPV2`: Pass the client IP via Proxy Protocol V2
	// `OFF`: Do not pass the client IP.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Specifies whether to enable session persistence
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type CacheConfig struct {
	// Cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cache *CacheConfigCache `json:"Cache,omitempty" name:"Cache"`

	// No-cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NoCache *CacheConfigNoCache `json:"NoCache,omitempty" name:"NoCache"`

	// Follows the origin server configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowOrigin *CacheConfigFollowOrigin `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type CacheConfigCache struct {
	// Cache configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Cache expiration time settings
	// Unit: second. The maximum value is 365 days.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`

	// Specifies whether to enable force cache
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`
}

type CacheConfigFollowOrigin struct {
	// Specifies whether to follow the origin server configuration
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CacheConfigNoCache struct {
	// Whether to cache the configuration
	// `on`: Do not cache
	// `off`: Cache
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CacheKey struct {
	// Specifies whether to enable full-path cache
	// `on`: Enable full-path cache (i.e., disable Ignore Query String)
	// `off`: Disable full-path cache (i.e., enable Ignore Query String)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`

	// Specifies whether the cache key is case sensitive
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreCase *string `json:"IgnoreCase,omitempty" name:"IgnoreCase"`

	// Request parameter contained in `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	QueryString *QueryString `json:"QueryString,omitempty" name:"QueryString"`
}

type CachePrefresh struct {

	Switch *string `json:"Switch,omitempty" name:"Switch"`


	Percent *int64 `json:"Percent,omitempty" name:"Percent"`
}

type CertFilter struct {
	// Filters by the field name. Values:
	//  - `host`: Domain name
	//  - `certId`: Certificate ID
	//  - `certAlias`: Certificate alias
	//  - `certType: default`: Default certificate; `upload`: External certificate; `managed`: Tencent Cloud certificate.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filters by the field value
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Specifies whether to enable fuzzy query, which only supports the `host` field.
	// If it is enabled, the length of `Value` must be 1.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type CertSort struct {
	// Fields that can be sorted. Values:
	// `createTime`: Domain name creation time
	// `certExpireTime`: Certificate expiration time
	// `certDeployTime`: Certificate deployment time
	Key *string `json:"Key,omitempty" name:"Key"`

	// Sorting order. Valid values: `asc` and `desc` (default).
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

// Predefined struct for user
type CheckCertificateRequestParams struct {
	// Certificate
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Private key
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type CheckCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Private key
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *CheckCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Certificate")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCertificateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CheckCertificateResponseParams `json:"Response"`
}

func (r *CheckCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientIp struct {
	// Specifies whether to enable client IP header
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Name of the origin-pull client IP request header
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
}

type CnameStatus struct {
	// Record name
	Name *string `json:"Name,omitempty" name:"Name"`

	// CNAME address
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Status
	// `active`: Activated
	// `moved`: Not activated
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type Compression struct {
	// Whether to enable Smart compression
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

// Predefined struct for user
type CreateApplicationProxyRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Layer-4 proxy name
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// Scheduling mode. Values:
	// `ip`: Anycast IP
	// `domain`: CNAME
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// `0`: Disable security protection; `1`: Enable security protection.
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// `0`: Disable acceleration; `1`: Enable acceleration.
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// This field is moved to `Rule.ForwardClientIp`.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This field is moved to `Rule.SessionPersist`.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// Rule details
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// Session persistence duration. Value range: 30-3600 (in seconds).
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Subdomain name
	// `instance`: Instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Layer-4 proxy name
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// Scheduling mode. Values:
	// `ip`: Anycast IP
	// `domain`: CNAME
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// `0`: Disable security protection; `1`: Enable security protection.
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// `0`: Disable acceleration; `1`: Enable acceleration.
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// This field is moved to `Rule.ForwardClientIp`.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This field is moved to `Rule.SessionPersist`.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// Rule details
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// Session persistence duration. Value range: 30-3600 (in seconds).
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Subdomain name
	// `instance`: Instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

func (r *CreateApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ZoneName")
	delete(f, "ProxyName")
	delete(f, "PlatType")
	delete(f, "SecurityType")
	delete(f, "AccelerateType")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	delete(f, "Rule")
	delete(f, "SessionPersistTime")
	delete(f, "ProxyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyResponseParams struct {
	// Layer-4 application proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyResponseParams `json:"Response"`
}

func (r *CreateApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRuleRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Protocol. Valid values: `TCP` and `UDP`.
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Port. Valid values:
	// `80`: Port 80
	// `81-90`: Port range 81-90
	Port []*string `json:"Port,omitempty" name:"Port"`

	// Origin server type. Valid values:
	// `custom`: Specified origins
	// `origins`: An origin group
	// `load_balancing`: A load balancer
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information.
	// When `OriginType=custom`, this field value indicates multiple origin servers in either of the following formats:
	// IP:Port
	// Domain name:Port
	// When `OriginType=origins`, it indicates the origin group ID.
	//  
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. When `Proto=TCP`, valid values:
	// `TOA`: Pass the client IP via TOA
	// `PPV1`: Pass the client IP via Proxy Protocol V1
	// `PPV2`: Pass the client IP via Proxy Protocol V2
	// `OFF`: Do not pass the client IP.
	// When `Proto=UDP`, valid values:
	// `PPV2`: Pass the client IP via Proxy Protocol V2
	// `OFF`: Do not pass the client IP.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Specifies whether to enable session persistence 
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type CreateApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Protocol. Valid values: `TCP` and `UDP`.
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Port. Valid values:
	// `80`: Port 80
	// `81-90`: Port range 81-90
	Port []*string `json:"Port,omitempty" name:"Port"`

	// Origin server type. Valid values:
	// `custom`: Specified origins
	// `origins`: An origin group
	// `load_balancing`: A load balancer
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information.
	// When `OriginType=custom`, this field value indicates multiple origin servers in either of the following formats:
	// IP:Port
	// Domain name:Port
	// When `OriginType=origins`, it indicates the origin group ID.
	//  
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. When `Proto=TCP`, valid values:
	// `TOA`: Pass the client IP via TOA
	// `PPV1`: Pass the client IP via Proxy Protocol V1
	// `PPV2`: Pass the client IP via Proxy Protocol V2
	// `OFF`: Do not pass the client IP.
	// When `Proto=UDP`, valid values:
	// `PPV2`: Pass the client IP via Proxy Protocol V2
	// `OFF`: Do not pass the client IP.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Specifies whether to enable session persistence 
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

func (r *CreateApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Proto")
	delete(f, "Port")
	delete(f, "OriginType")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyRuleResponseParams `json:"Response"`
}

func (r *CreateApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRulesRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule list
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`
}

type CreateApplicationProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule list
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *CreateApplicationProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRulesResponseParams struct {
	// Array of rule IDs
	RuleId []*string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyRulesResponseParams `json:"Response"`
}

func (r *CreateApplicationProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDnsRecordRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Record content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Proxy mode. Valid values: `dns_only`, `cdn_only`, and `secure_cdn`.
	Mode *string `json:"Mode,omitempty" name:"Mode"`


	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type CreateDnsRecordRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Record content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Proxy mode. Valid values: `dns_only`, `cdn_only`, and `secure_cdn`.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *CreateDnsRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDnsRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "Mode")
	delete(f, "Ttl")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDnsRecordResponseParams struct {
	// Record ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Record type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Record content
	Content *string `json:"Content,omitempty" name:"Content"`


	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Proxy mode
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Resolution status. Valid values:
	// `active`: Activated
	// `pending`: Not activated
	Status *string `json:"Status,omitempty" name:"Status"`

	// Whether the DNS record is locked
	Locked *bool `json:"Locked,omitempty" name:"Locked"`

	// Creation time
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Modification time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// CNAME address
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateDnsRecordResponseParams `json:"Response"`
}

func (r *CreateDnsRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancingRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name. You can use @ to represent the root domain.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Proxy mode. Valid values:
	// `dns_only`: Only DNS
	// `proxied`: Enable proxy
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of the origin group used
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// Indicates DNS TTL time when `Type=dns_only` 
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

type CreateLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name. You can use @ to represent the root domain.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Proxy mode. Valid values:
	// `dns_only`: Only DNS
	// `proxied`: Enable proxy
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of the origin group used
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// Indicates DNS TTL time when `Type=dns_only` 
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *CreateLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Host")
	delete(f, "Type")
	delete(f, "OriginId")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancingResponseParams struct {
	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancingResponseParams `json:"Response"`
}

func (r *CreateLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of resources to be pre-warmed, for example:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to encode the URL
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// HTTP header information
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of resources to be pre-warmed, for example:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to encode the URL
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// HTTP header information
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

func (r *CreatePrefetchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrefetchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskResponseParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// List of failed tasks
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrefetchTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrefetchTaskResponseParams `json:"Response"`
}

func (r *CreatePrefetchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Type of the purging task. Values:
	// - `purge_url`: Purge by the URL
	// - `purge_prefix`: Purge by the prefix
	// - `purge_host`: Purge by the Hostname
	// - `purge_all`: Purge all cached contents
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target resource to be purged, which depends on the `Type` field.
	// 1. When `Type = purge_host`:
	// Hostnames are purged, such as www.example.com and foo.bar.example.com.
	// 2. When `Type = purge_prefix`:
	// Prefixes are purged, such as http://www.example.com/example.
	// 3. When `Type = purge_url`:
	// URLs are purged, such as https://www.example.com/example.jpg.
	// 4. When `Type = purge_all`: All types of resources are purged.
	// `Targets` is not a required field.
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Type of the purging task. Values:
	// - `purge_url`: Purge by the URL
	// - `purge_prefix`: Purge by the prefix
	// - `purge_host`: Purge by the Hostname
	// - `purge_all`: Purge all cached contents
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target resource to be purged, which depends on the `Type` field.
	// 1. When `Type = purge_host`:
	// Hostnames are purged, such as www.example.com and foo.bar.example.com.
	// 2. When `Type = purge_prefix`:
	// Prefixes are purged, such as http://www.example.com/example.
	// 3. When `Type = purge_url`:
	// URLs are purged, such as https://www.example.com/example.jpg.
	// 4. When `Type = purge_all`: All types of resources are purged.
	// `Targets` is not a required field.
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

func (r *CreatePurgeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePurgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskResponseParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// List of failed tasks and reasons
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePurgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePurgeTaskResponseParams `json:"Response"`
}

func (r *CreatePurgeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneRequestParams struct {
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Access mode. Valid values:
	// - `full` (default): Access via NS
	// - `partial`: Access via CNAME
	Type *string `json:"Type,omitempty" name:"Type"`

	// Specifies whether to skip resolution record scanning
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest
	
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Access mode. Valid values:
	// - `full` (default): Access via NS
	// - `partial`: Access via CNAME
	Type *string `json:"Type,omitempty" name:"Type"`

	// Specifies whether to skip resolution record scanning
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`
}

func (r *CreateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "JumpStart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Specifies how the site is connected to EdgeOne.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Site status
	// - `pending`: The name server is not switched.
	// - `active`: The name server is switched to another assigned.
	// - `moved`: Move the NS out of Tencent Cloud
	Status *string `json:"Status,omitempty" name:"Status"`

	// List of name servers
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// List of name servers assigned to users
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// Site creation date
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Site update time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreateZoneResponseParams `json:"Response"`
}

func (r *CreateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataItem struct {
	// Time
	Time *string `json:"Time,omitempty" name:"Time"`

	// Value
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DefaultServerCertInfo struct {
	// Server certificate ID, which is the ID of the default certificate. If you choose to upload an external certificate for SSL certificate management, a certificate ID will be generated.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Certificate alias
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Certificate type. Valid values:
	// `default`: Default certificate
	// `upload`: External certificate
	// `managed`: Tencent Cloud managed certificate
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time when the certificate expires
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Time when the certificate takes effect
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// Certificate common name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CommonName *string `json:"CommonName,omitempty" name:"CommonName"`

	// Domain names added to the SAN certificate
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// Certificate status. Valid values:
	// `applying`: Application in progress
	// `failed`: Application failed
	// `processing`: Deploying certificate
	// `deployed`: Certificate deployed
	// `disabled`: Certificate disabled
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Returns a message to display failure causes when `Status=failed`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`
}

// Predefined struct for user
type DeleteApplicationProxyRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type DeleteApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DeleteApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyResponseParams struct {
	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationProxyResponseParams `json:"Response"`
}

func (r *DeleteApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyRuleRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationProxyRuleResponseParams `json:"Response"`
}

func (r *DeleteApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDnsRecordsRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

type DeleteDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDnsRecordsResponseParams struct {
	// Record ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDnsRecordsResponseParams `json:"Response"`
}

func (r *DeleteDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancingRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
}

type DeleteLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
}

func (r *DeleteLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancingResponseParams struct {
	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancingResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneRequestParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteZoneResponse struct {
	*tchttp.BaseResponse
	Response *DeleteZoneResponseParams `json:"Response"`
}

func (r *DeleteZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxyDetailRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type DescribeApplicationProxyDetailRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

func (r *DescribeApplicationProxyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxyDetailResponseParams struct {
	// Instance ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Instance name
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// Proxy mode. Valid values:
	// `ip`: Anycast IP
	// `domain`: CNAME
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// `0`: Disable security protection; `1`: Enable security protection.
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// `0`: Disable acceleration; `1`: Enable acceleration.
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// This field is moved to `Rule.ForwardClientIp`.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This field is moved to `Rule.SessionPersist`.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// List of rules
	Rule []*ApplicationProxyRule `json:"Rule,omitempty" name:"Rule"`

	// Status. Valid values:
	// `online`: Enable
	// `offline`: Disable
	// `progress`: Deploying
	Status *string `json:"Status,omitempty" name:"Status"`

	// Scheduling information
	ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Session persistence time
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Subdomain name
	// `instance`: Instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// ID of the layer-7 domain name
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationProxyDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProxyDetailResponseParams `json:"Response"`
}

func (r *DescribeApplicationProxyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxyRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Pagination parameter
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Pagination parameter
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxyResponseParams struct {
	// List of data
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Data []*ApplicationProxy `json:"Data,omitempty" name:"Data"`

	// Total number of records
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Indicates the number of instances that can be created by the site when `ZoneId` is specified
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`


	IpCount *uint64 `json:"IpCount,omitempty" name:"IpCount"`


	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProxyResponseParams `json:"Response"`
}

func (r *DescribeApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCnameStatusRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type DescribeCnameStatusRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DescribeCnameStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCnameStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCnameStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCnameStatusResponseParams struct {
	// List of CNAME statuses
	Status []*CnameStatus `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCnameStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCnameStatusResponseParams `json:"Response"`
}

func (r *DescribeCnameStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCnameStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeDefaultCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesResponseParams struct {
	// Total number of certificates
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of default certificates
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertInfo []*DefaultServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDefaultCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultCertificatesResponseParams `json:"Response"`
}

func (r *DescribeDefaultCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsDataRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter parameters
	Filters []*DnsDataFilter `json:"Filters,omitempty" name:"Filters"`

	// Time granularity. The default value is `min`. The server can adapt to the time granularity specified.
	// Valid values:
	// `min`: 1 minute
	// `5min`: 5 minutes
	// `hour`: 1 hour
	// `day`: 1 day
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeDnsDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter parameters
	Filters []*DnsDataFilter `json:"Filters,omitempty" name:"Filters"`

	// Time granularity. The default value is `min`. The server can adapt to the time granularity specified.
	// Valid values:
	// `min`: 1 minute
	// `5min`: 5 minutes
	// `hour`: 1 hour
	// `day`: 1 day
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeDnsDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsDataResponseParams struct {
	// DNS request data
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Data []*DataItem `json:"Data,omitempty" name:"Data"`

	// Interval
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDnsDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDnsDataResponseParams `json:"Response"`
}

func (r *DescribeDnsDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsRecordsRequestParams struct {
	// Query filter
	Filters []*DnsRecordFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorts the order
	Order *string `json:"Order,omitempty" name:"Order"`

	// Valid values: `asc`, and `desc`.
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// Valid values: `all`, and `any`.
	Match *string `json:"Match,omitempty" name:"Match"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Query filter
	Filters []*DnsRecordFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorts the order
	Order *string `json:"Order,omitempty" name:"Order"`

	// Valid values: `asc`, and `desc`.
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// Valid values: `all`, and `any`.
	Match *string `json:"Match,omitempty" name:"Match"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "Direction")
	delete(f, "Match")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsRecordsResponseParams struct {
	// Used for paginated query by total count
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of DNS records
	Records []*DnsRecord `json:"Records,omitempty" name:"Records"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDnsRecordsResponseParams `json:"Response"`
}

func (r *DescribeDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnssecRequestParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeDnssecRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDnssecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnssecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnssecResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// DNSSEC status. Valid values:
	// - `enabled`: Enabled
	// - `disabled`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`


	Dnssec *DnssecInfo `json:"Dnssec,omitempty" name:"Dnssec"`

	// Modification time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDnssecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDnssecResponseParams `json:"Response"`
}

func (r *DescribeDnssecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDnssecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsCertificateRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query filter
	Filters []*CertFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting order
	Sort *CertSort `json:"Sort,omitempty" name:"Sort"`
}

type DescribeHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query filter
	Filters []*CertFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting order
	Sort *CertSort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeHostsCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsCertificateResponseParams struct {
	// Used for paginated query by total count
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of certificate configurations for domain names
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Hosts []*HostCertSetting `json:"Hosts,omitempty" name:"Hosts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostsCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsCertificateResponseParams `json:"Response"`
}

func (r *DescribeHostsCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a domain name for the query
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
}

type DescribeHostsSettingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a domain name for the query
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
}

func (r *DescribeHostsSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Hosts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingResponseParams struct {
	// List of domain names
	Hosts []*DetailHost `json:"Hosts,omitempty" name:"Hosts"`

	// Number of domain names
	TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostsSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsSettingResponseParams `json:"Response"`
}

func (r *DescribeHostsSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationRequestParams struct {
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeIdentificationRequest struct {
	*tchttp.BaseRequest
	
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeIdentificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdentificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationResponseParams struct {
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Verification status. Valid values:
	// - `pending`: Verifying
	// - `finished`: The site is verified.
	Status *string `json:"Status,omitempty" name:"Status"`


	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// Record type
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Record value
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// NS records of the domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIdentificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdentificationResponseParams `json:"Response"`
}

func (r *DescribeIdentificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancingDetailRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
}

type DescribeLoadBalancingDetailRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`
}

func (r *DescribeLoadBalancingDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancingDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancingDetailResponseParams struct {
	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name. You can use @ to represent the root domain.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Proxy mode. Valid values:
	// `dns_only`: Only DNS
	// `proxied`: Enable proxy
	Type *string `json:"Type,omitempty" name:"Type"`

	// Indicates DNS TTL time when `Type=dns_only`
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// ID of the origin group used
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// Information of the origin server used
	Origin []*OriginGroup `json:"Origin,omitempty" name:"Origin"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Status of the task
	Status *string `json:"Status,omitempty" name:"Status"`

	// Schedules domain names
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancingDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancingDetailResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancingDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancingRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Pagination parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Ignore query string parameter
	Host *string `json:"Host,omitempty" name:"Host"`

	// Specifies whether the `Host` parameter supports fuzzy match
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type DescribeLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Pagination parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Ignore query string parameter
	Host *string `json:"Host,omitempty" name:"Host"`

	// Specifies whether the `Host` parameter supports fuzzy match
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

func (r *DescribeLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Host")
	delete(f, "Fuzzy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancingResponseParams struct {
	// Total number of records
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// CLB information
	Data []*LoadBalancing `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancingResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Resources queried
	Target *string `json:"Target,omitempty" name:"Target"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Resources queried
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePrefetchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksResponseParams struct {
	// Total entries that match the specified query condition
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of tasks returned
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrefetchTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrefetchTasksResponseParams `json:"Response"`
}

func (r *DescribePrefetchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Type of the purging task
	Type *string `json:"Type,omitempty" name:"Type"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries content
	Target *string `json:"Target,omitempty" name:"Target"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Type of the purging task
	Type *string `json:"Type,omitempty" name:"Type"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries content
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// Total entries that match the specified query condition
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of tasks returned
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeTasksResponseParams `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDetailsRequestParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeZoneDetailsRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeZoneDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDetailsResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of name servers used
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// List of name servers assigned to users by Tencent Cloud
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// Site status
	// - `active`: The name server is switched.
	// - `pending`: The name server is not switched.
	// - `moved`: The name server is moved.
	// - `deactivated`: The name server is blocked.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Specifies how the site is connected to EdgeOne.
	// - `full`: The site is connected via name server.
	// - `partial`: The site is connected via CNAME.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Indicates whether the site is disabled
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// Site creation date
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Site modification date
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// User-defined name server information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// User-defined name server IP information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitempty" name:"VanityNameServersIps"`

	// Specifies whether to enable CNAME acceleration
	// - `enabled`: Enable
	// - `disabled`: Disable
	CnameSpeedUp *string `json:"CnameSpeedUp,omitempty" name:"CnameSpeedUp"`

	// Ownership verification status of the site when it accesses via CNAME.
	// - `finished`: The site is verified.
	// - `pending`: The site is waiting for verification.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneDetailsResponseParams `json:"Response"`
}

func (r *DescribeZoneDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneSettingRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeZoneSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneSettingResponseParams struct {
	// Cache expiration time configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

	// Node cache key configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Browser cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Offline cache
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// QUIC access
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// POST transport configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// Smart compression configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS acceleration configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Origin server configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Dynamic acceleration configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Domain name of the site
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// WebSocket configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`


	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneSettingResponseParams `json:"Response"`
}

func (r *DescribeZoneSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// Pagination parameter, which specifies the offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter, which specifies the number of sites returned in each page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, which supports complex type.
	Filters []*ZoneFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// Pagination parameter, which specifies the offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter, which specifies the number of sites returned in each page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, which supports complex type.
	Filters []*ZoneFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Number of sites that match the specified conditions
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Details of sites
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Zones []*Zone `json:"Zones,omitempty" name:"Zones"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailHost struct {
	// Tencent Cloud account ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Acceleration service status
	// `process`: Deploying
	// `online`: Enabled
	// `offline`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`

	// Domain name
	Host *string `json:"Host,omitempty" name:"Host"`
}

type DnsDataFilter struct {
	// Parameter name. Valid values:
	// `zone`: Site name
	// `host`: Domain name
	// `type`: DNS resolution type
	// `code`: DNS response code
	// `area`: Region of the resolution server
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter value
	// When `Name=area`, valid values:
	// `Asia`
	// `Europe`
	// `Africa`
	// `Oceania`
	// `Americas`
	// 
	// When `Name=code`, valid values:
	// `NoError`: Successful response.
	// `NXDomain`: Non-existent domain in the request. It is only valid when the response is from the authoritative name server.
	// `NotImp`: Request type not supported.
	// `Refused`: The name server refuses to perform the requested operation for policy reasons.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Parameter value
	// When `Name=area`, valid values:
	// `Asia`
	// `Europe`
	// `Africa`
	// `Oceania`
	// `Americas`
	// 
	// When `Name=code`, valid values:
	// `NoError`: Successful response.
	// `NXDomain`: Non-existent domain in the request. It is only valid when the response is from the authoritative name server.
	// `NotImp`: Request type not supported.
	// `Refused`: The name server refuses to perform the requested operation for policy reasons.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DnsRecord struct {
	// Record ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Record type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Host record
	Name *string `json:"Name,omitempty" name:"Name"`

	// Record value
	Content *string `json:"Content,omitempty" name:"Content"`

	// Proxy mode
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// TTL value
	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// Priority
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Creation time
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Modification time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// Domain name lock
	Locked *bool `json:"Locked,omitempty" name:"Locked"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Resolution status
	// `active`: Activated
	// `pending`: Not activated
	Status *string `json:"Status,omitempty" name:"Status"`

	// CNAME address
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Specifies whether to enable load balancing, layer-4 proxy, or security protection for the domain name.
	// - `lb`: Load balancing.
	// - `security`: Security protection.
	// - `l4`: Layer-4 proxy.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DomainStatus []*string `json:"DomainStatus,omitempty" name:"DomainStatus"`
}

type DnsRecordFilter struct {
	// Filters by the field name. Vaules:
	// - `name`: Site name.
	// - `status`: Site status.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filters by the field value
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Specifies whether to enable fuzzy query. It’s only available when the filter name is `name`. If it’s enabled, the length of `Values` must be 1.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type DnssecInfo struct {
	// Flag
	Flags *int64 `json:"Flags,omitempty" name:"Flags"`

	// Encryption algorithm
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Encryption type
	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`

	// Digest type
	DigestType *string `json:"DigestType,omitempty" name:"DigestType"`

	// Digest algorithm
	DigestAlgorithm *string `json:"DigestAlgorithm,omitempty" name:"DigestAlgorithm"`

	// Digest message
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// DS record value
	DS *string `json:"DS,omitempty" name:"DS"`

	// Key tag
	KeyTag *int64 `json:"KeyTag,omitempty" name:"KeyTag"`

	// Public key
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

// Predefined struct for user
type DownloadL7LogsRequestParams struct {
	// Start time. It must conform to the RFC3339 standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time. It must conform to the RFC3339 standard.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// List of sites
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// List of domain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest
	
	// Start time. It must conform to the RFC3339 standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time. It must conform to the RFC3339 standard.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// List of sites
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// List of domain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *DownloadL7LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Zones")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL7LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsResponseParams struct {
	// Layer-7 offline log data
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Data []*L7OfflineLog `json:"Data,omitempty" name:"Data"`

	// Page size
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Page number
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Total number of pages
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of entries
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadL7LogsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadL7LogsResponseParams `json:"Response"`
}

func (r *DownloadL7LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailReason struct {
	// Failure reason
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// List of resources failed to be processed. 
	//  
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type ForceRedirect struct {
	// Force redirect configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Redirection status code
	// 301
	// 302
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type Header struct {
	// HTTP header name
	Name *string `json:"Name,omitempty" name:"Name"`

	// HTTP header value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type HostCertSetting struct {
	// Domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Server certificate configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`
}

type Hsts struct {
	// Specifies whether to enable. Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// `MaxAge` value.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAge *int64 `json:"MaxAge,omitempty" name:"MaxAge"`

	// Specifies whether to include subdomain names. Valid values: `on` and `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IncludeSubDomains *string `json:"IncludeSubDomains,omitempty" name:"IncludeSubDomains"`

	// Specifies whether to preload. Valid values: `on` and `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Preload *string `json:"Preload,omitempty" name:"Preload"`
}

type Https struct {
	// HTTP2 configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// OCSP configuration switch
	// `on`: Enable
	// `off`: Disable
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// TLS version settings. Valid values: `TLSv1`, `TLSV1.1`, `TLSV1.2`, and `TLSv1.3`. Only consecutive versions can be enabled at the same time.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TlsVersion []*string `json:"TlsVersion,omitempty" name:"TlsVersion"`

	// HSTS Configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Hsts *Hsts `json:"Hsts,omitempty" name:"Hsts"`
}

// Predefined struct for user
type IdentifyZoneRequestParams struct {
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *IdentifyZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdentifyZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdentifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdentifyZoneResponseParams struct {
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`


	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// Record type
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Record value
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IdentifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *IdentifyZoneResponseParams `json:"Response"`
}

func (r *IdentifyZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdentifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportDnsRecordsRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// File content
	File *string `json:"File,omitempty" name:"File"`
}

type ImportDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// File content
	File *string `json:"File,omitempty" name:"File"`
}

func (r *ImportDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "File")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportDnsRecordsResponseParams struct {
	// Record ID
	Ids []*string `json:"Ids,omitempty" name:"Ids"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *ImportDnsRecordsResponseParams `json:"Response"`
}

func (r *ImportDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type L7OfflineLog struct {
	// Start time of the log packaging
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// Site name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Log size, in bytes
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Download address
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Log package name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`
}

type LoadBalancing struct {
	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name. You can use @ to represent the root domain.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Proxy mode. Valid values:
	// `dns_only`: Only DNS
	// `proxied`: Enable proxy
	Type *string `json:"Type,omitempty" name:"Type"`

	// Indicates DNS TTL time when `Type=dns_only`
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// ID of the origin group used
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// Information of the origin server used
	Origin []*OriginGroup `json:"Origin,omitempty" name:"Origin"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Schedules domain names
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`
}

type MaxAge struct {
	// Specifies the max age of the cache (in seconds). The maximum value is 365 days.
	// Note: the value `0` means not to cache.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAgeTime *int64 `json:"MaxAgeTime,omitempty" name:"MaxAgeTime"`

	// Specifies whether to follow the max cache age of the origin server. Valid values: `on` and `off`. If it's on, `MaxAgeTime` is ignored.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

// Predefined struct for user
type ModifyApplicationProxyRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Layer-4 proxy name
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// This parameter is disused.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This parameter is disused.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// Session persistence time. Value range: 30-3600 (in seconds).
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Subdomain name
	// `instance`: Instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Layer-4 proxy name
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// This parameter is disused.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This parameter is disused.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// Session persistence time. Value range: 30-3600 (in seconds).
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Subdomain name
	// `instance`: Instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

func (r *ModifyApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "ProxyName")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	delete(f, "SessionPersistTime")
	delete(f, "ProxyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyResponseParams struct {
	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Protocol. Valid values: `TCP` and `UDP`.
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Port. Valid values:
	// `80`: Port 80
	// `81-90`: Port range 81-90
	Port []*string `json:"Port,omitempty" name:"Port"`

	// Origin server type. Valid values:
	// `custom`: Specified origins
	// `origins`: An origin group
	// `load_balancing`: A load balancer
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information.
	// When `OriginType=custom`, this field value indicates multiple origin servers in either of the following formats:
	// IP:Port
	// Domain name: Port
	// When `OriginType=origins`, it indicates the origin group ID.
	//  
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. When `Proto=TCP`, valid values:
	// `TOA`: Pass the client IP via TOA.
	// `PPV1`: Pass the client IP via Proxy Protocol V1.
	// `PPV2`: Pass the client IP via Proxy Protocol V2.
	// `OFF`: Do not pass the client IP.
	// When `Proto=UDP`, valid values:
	// `PPV2`: Pass the client IP via Proxy Protocol V2.
	// `OFF`: Do not pass the client IP.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Specifies whether to enable session persistence
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

type ModifyApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Protocol. Valid values: `TCP` and `UDP`.
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Port. Valid values:
	// `80`: Port 80
	// `81-90`: Port range 81-90
	Port []*string `json:"Port,omitempty" name:"Port"`

	// Origin server type. Valid values:
	// `custom`: Specified origins
	// `origins`: An origin group
	// `load_balancing`: A load balancer
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information.
	// When `OriginType=custom`, this field value indicates multiple origin servers in either of the following formats:
	// IP:Port
	// Domain name: Port
	// When `OriginType=origins`, it indicates the origin group ID.
	//  
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. When `Proto=TCP`, valid values:
	// `TOA`: Pass the client IP via TOA.
	// `PPV1`: Pass the client IP via Proxy Protocol V1.
	// `PPV2`: Pass the client IP via Proxy Protocol V2.
	// `OFF`: Do not pass the client IP.
	// When `Proto=UDP`, valid values:
	// `PPV2`: Pass the client IP via Proxy Protocol V2.
	// `OFF`: Do not pass the client IP.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Specifies whether to enable session persistence
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`
}

func (r *ModifyApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	delete(f, "Proto")
	delete(f, "Port")
	delete(f, "OriginType")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyRuleResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleStatusRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Status
	// `offline`: Disabled
	// `online`: Enabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Status
	// `offline`: Disabled
	// `online`: Enabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyApplicationProxyRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleStatusResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyStatusRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Status
	// `offline`: Disabled
	// `online`: Enabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Status
	// `offline`: Disabled
	// `online`: Enabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyApplicationProxyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyStatusResponseParams struct {
	// Layer-4 proxy ID
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationProxyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyStatusResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultCertificateRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Certificate ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Certificate status
	// `deployed`: The certificate is deployed.
	// `disabled`: The certificate is disabled.
	// If the deployment fails, you can pass in `Status = deployed` again.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyDefaultCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Certificate ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Certificate status
	// `deployed`: The certificate is deployed.
	// `disabled`: The certificate is disabled.
	// If the deployment fails, you can pass in `Status = deployed` again.
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDefaultCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "CertId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDefaultCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultCertificateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDefaultCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDefaultCertificateResponseParams `json:"Response"`
}

func (r *ModifyDefaultCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordRequestParams struct {
	// Record ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Record content
	Content *string `json:"Content,omitempty" name:"Content"`


	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Proxy mode
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ModifyDnsRecordRequest struct {
	*tchttp.BaseRequest
	
	// Record ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Record content
	Content *string `json:"Content,omitempty" name:"Content"`

	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Proxy mode
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyDnsRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "Ttl")
	delete(f, "Priority")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordResponseParams struct {
	// Record ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Record type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Record content
	Content *string `json:"Content,omitempty" name:"Content"`


	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Proxy mode
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Resolution status
	Status *string `json:"Status,omitempty" name:"Status"`

	// CNAME address
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Whether the DNS record is locked
	Locked *bool `json:"Locked,omitempty" name:"Locked"`

	// Creation time
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Modification time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDnsRecordResponseParams `json:"Response"`
}

func (r *ModifyDnsRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnssecRequestParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// DNSSEC status
	// - `enabled`: Enabled
	// - `disabled`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyDnssecRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// DNSSEC status
	// - `enabled`: Enabled
	// - `disabled`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDnssecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnssecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnssecResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// DNSSEC status.
	// - `enabled`: Enabled
	// - `disabled`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`

	// DNSSEC information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Dnssec *DnssecInfo `json:"Dnssec,omitempty" name:"Dnssec"`

	// Modification time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDnssecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDnssecResponseParams `json:"Response"`
}

func (r *ModifyDnssecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDnssecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Domain name that the certificate will be attached to
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Certificate information. Note that only `CertId` is required. If it is not specified, the default certificate will be used.
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Domain name that the certificate will be attached to
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Certificate information. Note that only `CertId` is required. If it is not specified, the default certificate will be used.
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`
}

func (r *ModifyHostsCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Hosts")
	delete(f, "CertInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHostsCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostsCertificateResponseParams `json:"Response"`
}

func (r *ModifyHostsCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancingRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// Proxy mode.
	// `dns_only`: Only DNS
	// `proxied`: Enable proxy
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of the origin group used
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// Indicates DNS TTL time when `Type=dns_only`
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

type ModifyLoadBalancingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// Proxy mode.
	// `dns_only`: Only DNS
	// `proxied`: Enable proxy
	Type *string `json:"Type,omitempty" name:"Type"`

	// ID of the origin group used
	OriginId []*string `json:"OriginId,omitempty" name:"OriginId"`

	// Indicates DNS TTL time when `Type=dns_only`
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *ModifyLoadBalancingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	delete(f, "Type")
	delete(f, "OriginId")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancingResponseParams struct {
	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancingResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancingStatusRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// Status.
	// `online`: Enabled
	// `offline`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyLoadBalancingStatusRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// Status.
	// `online`: Enabled
	// `offline`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyLoadBalancingStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LoadBalancingId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancingStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancingStatusResponseParams struct {
	// CLB instance ID
	LoadBalancingId *string `json:"LoadBalancingId,omitempty" name:"LoadBalancingId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancingStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancingStatusResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancingStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancingStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCnameSpeedUpRequestParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CNAME acceleration status.
	// - `enabled`: Enabled
	// - `disabled`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyZoneCnameSpeedUpRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// CNAME acceleration status.
	// - `enabled`: Enabled
	// - `disabled`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyZoneCnameSpeedUpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCnameSpeedUpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneCnameSpeedUpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCnameSpeedUpResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// CNAME acceleration status.
	// - `enabled`: Enabled
	// - `disabled`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`

	// Update time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneCnameSpeedUpResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneCnameSpeedUpResponseParams `json:"Response"`
}

func (r *ModifyZoneCnameSpeedUpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCnameSpeedUpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneRequestParams struct {
	// Site ID, which is used to identify the site.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Specifies how the site is connected to EdgeOne.
	// - `full`: Connect via the name server.
	// - `partial`: Connect via the CNAME.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Custom site information
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// Site ID, which is used to identify the site.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Specifies how the site is connected to EdgeOne.
	// - `full`: Connect via the name server.
	// - `partial`: Connect via the CNAME.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Custom site information
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`
}

func (r *ModifyZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Type")
	delete(f, "VanityNameServers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Name server used by the site
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// Site status.
	// - `pending`: The name server is not connected.
	// - `active`: The name server is connected.
	// - `moved`: The name server is moved.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Specifies how the site is connected to EdgeOne.
	// - `full`: Connect via the name server.
	// - `partial`: Connect via the CNAME.
	Type *string `json:"Type,omitempty" name:"Type"`

	// List of name servers assigned by Tencent Cloud
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// Creation time
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Modification time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// CNAME access status.
	// - `finished`: Ownership of the site is verified.
	// - `pending`: Verifying the ownership of the site.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneResponseParams `json:"Response"`
}

func (r *ModifyZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingRequestParams struct {
	// ID of the site to be modified
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Cache expiration time
	Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

	// Node cache key
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Browser cache configuration
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Offline cache
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// QUIC access
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// Maximum size of files transferred over POST request
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// Smart compression configuration
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS acceleration configuration
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Origin server configuration
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Smart acceleration configuration
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// WebSocket configuration
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration
	ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`


	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site to be modified
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Cache expiration time
	Cache *CacheConfig `json:"Cache,omitempty" name:"Cache"`

	// Node cache key
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Browser cache configuration
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Offline cache
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// QUIC access
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// Maximum size of files transferred over POST request
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// Smart compression configuration
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS acceleration configuration
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Origin server configuration
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Smart acceleration configuration
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// WebSocket configuration
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration
	ClientIpHeader *ClientIp `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`
}

func (r *ModifyZoneSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Cache")
	delete(f, "CacheKey")
	delete(f, "MaxAge")
	delete(f, "OfflineCache")
	delete(f, "Quic")
	delete(f, "PostMaxSize")
	delete(f, "Compression")
	delete(f, "UpstreamHttp2")
	delete(f, "ForceRedirect")
	delete(f, "Https")
	delete(f, "Origin")
	delete(f, "SmartRouting")
	delete(f, "WebSocket")
	delete(f, "ClientIpHeader")
	delete(f, "CachePrefresh")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingResponseParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneSettingResponseParams `json:"Response"`
}

func (r *ModifyZoneSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneStatusRequestParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site status.
	// - `false`: Enable the site.
	// - `true`: Disable the site.
	Paused *bool `json:"Paused,omitempty" name:"Paused"`
}

type ModifyZoneStatusRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site status.
	// - `false`: Enable the site.
	// - `true`: Disable the site.
	Paused *bool `json:"Paused,omitempty" name:"Paused"`
}

func (r *ModifyZoneStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Paused")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneStatusResponseParams struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Site status.
	// - `false`: Enable the site.
	// - `true`: Disable the site.
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// Update time
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyZoneStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneStatusResponseParams `json:"Response"`
}

func (r *ModifyZoneStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineCache struct {
	// Whether to enable offline cache. Valid values: `on` and `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Origin struct {
	// Origin-pull protocol.
	// `http`: Switch HTTPS requests to HTTP
	// `follow`: Follow the protocol of the request.
	// `https`: Switch HTTP requests to HTTPS. This only supports port 443 on the origin server.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`
}

type OriginCheckOriginStatus struct {

	Status *string `json:"Status,omitempty" name:"Status"`


	Host []*string `json:"Host,omitempty" name:"Host"`
}

type OriginGroup struct {
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Origin group name
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// Origin server type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Origin server type
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Whether the origin group is used for layer-4 proxy
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ApplicationProxyUsed *bool `json:"ApplicationProxyUsed,omitempty" name:"ApplicationProxyUsed"`

	// Whether the origin group is used for load balancing
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	LoadBalancingUsed *bool `json:"LoadBalancingUsed,omitempty" name:"LoadBalancingUsed"`


	Status *OriginCheckOriginStatus `json:"Status,omitempty" name:"Status"`


	LoadBalancingUsedType *string `json:"LoadBalancingUsedType,omitempty" name:"LoadBalancingUsedType"`
}

type OriginRecord struct {
	// Record value
	Record *string `json:"Record,omitempty" name:"Record"`

	// Region of the origin group. It’s available when the origin group `Type` is `area`. 
	// If it’s left empty, it means to use the default region.
	Area []*string `json:"Area,omitempty" name:"Area"`

	// The weight of the origin group. It’s available when the `Type` is `weight`.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// Port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Record ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// Specifies whether to run private origin authentication.
	// It is valid only when `OriginType=third_part`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Private *bool `json:"Private,omitempty" name:"Private"`

	// Private origin parameter.
	// It is valid only when `Private=true`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PrivateParameter []*OriginRecordPrivateParameter `json:"PrivateParameter,omitempty" name:"PrivateParameter"`
}

type OriginRecordPrivateParameter struct {
	// Name of the private origin authentication parameter.
	// `AccessKeyId`: Access key ID
	// `SecretAccessKey`: Secret access key
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value of the private origin authentication parameter
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PostMaxSize struct {
	// Specifies whether to enable custom setting of the maximum file size. 
	// `off`: Disable. In this case, the max size defaults to 32 MB.
	// `on`: Enable. You can set a custom max size.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Maximum size. Value range: 1-500 MB.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type QueryString struct {
	// Whether to use `QueryString` as part of `CacheKey`. Valid values: `on` and `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// `includeCustom`: Include the specified query strings.
	// `excludeCustom`: Exclude the specified query strings.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Array of query strings used/excluded
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Quic struct {
	// Whether to enable QUIC
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

// Predefined struct for user
type ReclaimZoneRequestParams struct {
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ReclaimZoneRequest struct {
	*tchttp.BaseRequest
	
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ReclaimZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReclaimZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReclaimZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReclaimZoneResponseParams struct {
	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReclaimZoneResponse struct {
	*tchttp.BaseResponse
	Response *ReclaimZoneResponseParams `json:"Response"`
}

func (r *ReclaimZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReclaimZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanDnsRecordsRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type ScanDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ScanDnsRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanDnsRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanDnsRecordsResponseParams struct {
	// Scan status
	// - `doing`: Scanning
	// - `done`: Scanned
	Status *string `json:"Status,omitempty" name:"Status"`

	// Number of DNS records added after scanning
	RecordsAdded *int64 `json:"RecordsAdded,omitempty" name:"RecordsAdded"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanDnsRecordsResponse struct {
	*tchttp.BaseResponse
	Response *ScanDnsRecordsResponseParams `json:"Response"`
}

func (r *ScanDnsRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanDnsRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerCertInfo struct {
	// Server certificate ID, which is the ID of the default certificate. If you choose to upload an external certificate for SSL certificate management, a certificate ID will be generated.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Alias of the certificate
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Certificate type.
	// `default`: Default certificate
	// `upload`: External certificate
	// `managed`: Tencent Cloud managed certificate
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time when the certificate expires
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Certificate deployment time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// Certificate deployment status.
	// `processing`: Deploying
	// `deployed`: Deployed
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type SmartRouting struct {
	// Whether to enable smart acceleration
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Task struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Status of the task
	Status *string `json:"Status,omitempty" name:"Status"`

	// Resource
	Target *string `json:"Target,omitempty" name:"Target"`

	// Task type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task completion time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UpstreamHttp2 struct {
	// Whether to enable HTTP2 origin-pull
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type VanityNameServers struct {
	// Whether to enable the custom name server
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// List of custom name servers
	Servers []*string `json:"Servers,omitempty" name:"Servers"`
}

type VanityNameServersIps struct {
	// Name of the custom name server
	Name *string `json:"Name,omitempty" name:"Name"`

	// IPv4 address of the custom name server
	IPv4 *string `json:"IPv4,omitempty" name:"IPv4"`
}

type WebSocket struct {
	// Whether to enable custom WebSocket timeout setting. When it’s `off`: it means to keep the default WebSocket connection timeout period, which is 15 seconds. To change the timeout period, please set it to `on`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Sets timeout period in seconds. Maximum value: 120
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type Zone struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of name servers used by the site
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// List of name servers assigned by Tencent Cloud
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// Site status
	// - `active`: The name server is switched.
	// - `pending`: The name server is not switched.
	// - `moved`: The name server is moved.
	// - `deactivated`: The name server is blocked.
	Status *string `json:"Status,omitempty" name:"Status"`

	// How the site is connected to EdgeOne.
	// - `full`: The site is connected via name server.
	// - `partial`: The site is connected via CNAME.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Indicates whether the site is disabled
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// Site creation date
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Site modification date
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// Ownership verification status of the site when it is connected to EdgeOne via CNAME.
	// - `finished`: The site is verified.
	// - `pending`: Verifying the ownership of the site.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`
}

type ZoneFilter struct {
	// Filters by the field name. Vaules:
	// - `name`: Site name.
	// - `status`: Site status.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filters by the field value
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Specifies whether to enable fuzzy query. It’s only available when filter name is `name`. If it’s enabled, the length of `Values` must be 1.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}