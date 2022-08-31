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

type ACLCondition struct {
	// Field to match
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// String to match
	MatchParam *string `json:"MatchParam,omitempty" name:"MatchParam"`

	// Relation between the field and content
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Content to match
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type ACLUserRule struct {
	// Name of the rule
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Action
	Action *string `json:"Action,omitempty" name:"Action"`

	// Status of the rule
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// ACL rule
	Conditions []*ACLCondition `json:"Conditions,omitempty" name:"Conditions"`

	// Priority of the rule
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// ID of the rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// Update time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// IP blocking time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// IP blocking time unit
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// Name of the custom block page
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// ID of the custom block page
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// Redirection URL, which must be a subdomain name of the site
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// Return code configured on the custom block page
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseCode *int64 `json:"ResponseCode,omitempty" name:"ResponseCode"`
}

type AclConfig struct {
	// Switch
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// ACL user rule
	UserRules []*ACLUserRule `json:"UserRules,omitempty" name:"UserRules"`
}

type AiRule struct {
	// `smart_status_close`: Disable; `smart_status_open`: Block;
	// `smart_status_observe`: Observe.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ApplicationProxy struct {
	// ID of the proxy
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Name of the proxy:
	// Domain name or subdomain name when `ProxyType=hostname`
	// Instance name when `ProxyType=instance`
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
	// `hostname`: Create by subdomain name
	// `instance`: Create by instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// When `ProxyType=hostname`:
	// `ProxyName` indicates a specified domain name;
	// `HostId` indicates a unique ID of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	// `origins`: Origin group
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information:
	// When `OriginType=custom`, it indicates one or more origin servers. Example:
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]
	// 
	// When `OriginType=origins`, it indicates an origin group ID. Example:
	// OriginValue=["origin-xxx"]
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

type BotConfig struct {
	// Whether to enable bot security configuration
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Preset rules
	ManagedRule *BotManagedRule `json:"ManagedRule,omitempty" name:"ManagedRule"`

	// Not supported currently
	UaBotRule *BotManagedRule `json:"UaBotRule,omitempty" name:"UaBotRule"`

	// Not supported currently
	IspBotRule *BotManagedRule `json:"IspBotRule,omitempty" name:"IspBotRule"`

	// User portrait rules
	PortraitRule *BotPortraitRule `json:"PortraitRule,omitempty" name:"PortraitRule"`

	// Bot intelligence rules
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntelligenceRule *IntelligenceRule `json:"IntelligenceRule,omitempty" name:"IntelligenceRule"`
}

type BotLog struct {
	// Attack time
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// Attack IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// Domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Request URI
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// Attack type
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Request method
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// Attack content
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// Risk grade
	// Note: This field may return null, indicating that no valid values can be obtained.
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// Rule number
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// IP country/region
	// Note: This field may return null, indicating that no valid values can be obtained.
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// Event ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Processing method
	// Note: This field may return null, indicating that no valid values can be obtained.
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// http_log
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// user agent
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// Detection method
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetectionMethod *string `json:"DetectionMethod,omitempty" name:"DetectionMethod"`

	// Confidence
	// Note: This field may return null, indicating that no valid values can be obtained.
	Confidence *string `json:"Confidence,omitempty" name:"Confidence"`

	// Maliciousness
	// Note: This field may return null, indicating that no valid values can be obtained.
	Maliciousness *string `json:"Maliciousness,omitempty" name:"Maliciousness"`
}

type BotLogData struct {
	// Data set of bot attack logs
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*BotLog `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type BotManagedRule struct {
	// ID of the rule to be enabled
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ManagedIds []*int64 `json:"ManagedIds,omitempty" name:"ManagedIds"`

	// ID of the rule being applied
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// Action of the rule. Values: `drop`; `trans`; `monitor`; `alg`.
	Action *string `json:"Action,omitempty" name:"Action"`

	// The amount of time the IP is blocked
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// Unit of IP blocking time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// Name of the custom block page
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// ID of the custom block page
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// Redirection URL, which must be a subdomain name of your site encoded by URLEncode
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// Response code returned after redirection
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseCode *int64 `json:"ResponseCode,omitempty" name:"ResponseCode"`

	// ID of the rule that is set to allow requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TransManagedIds []*int64 `json:"TransManagedIds,omitempty" name:"TransManagedIds"`

	// ID of the rule that is set to verify requests by JavaScript challenge
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// ID of the rule that is set to verify requests by CAPTCHA
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// ID of the rule that is set to observe requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// ID of the rule that is set to block requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`
}

type BotManagedRuleDetail struct {
	// ID of the rule
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Rule description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Rule type
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// Whether the rule is enabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

type BotPortraitRule struct {
	// ID of the rule being applied
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// ID of the rule that is set to verify requests by JavaScript challenge
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// ID of the rule that is set to verify requests by CAPTCHA
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// ID of the rule that is set to observe requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// ID of the rule that is set to block requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`

	// Feature switch
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CCInterceptEvent struct {
	// Client IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Number of blocks per minute
	// Note: This field may return null, indicating that no valid values can be obtained.
	InterceptNum *int64 `json:"InterceptNum,omitempty" name:"InterceptNum"`

	// Block time in rate-limiting policy per minute in seconds
	InterceptTime *int64 `json:"InterceptTime,omitempty" name:"InterceptTime"`
}

type CCInterceptEventData struct {
	// Data set of attack events
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*CCInterceptEvent `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type CCLog struct {
	// Attack time
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// Attack source IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackSip *string `json:"AttackSip,omitempty" name:"AttackSip"`

	// Attack domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackDomain *string `json:"AttackDomain,omitempty" name:"AttackDomain"`

	// Request URI
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// Number of hits
	// Note: This field may return null, indicating that no valid values can be obtained.
	HitCount *uint64 `json:"HitCount,omitempty" name:"HitCount"`

	// IP country/region
	// Note: This field may return null, indicating that no valid values can be obtained.
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// Event ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Processing method
	// Note: This field may return null, indicating that no valid values can be obtained.
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// http_log
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// Rule number
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// Risk grade
	// Note: This field may return null, indicating that no valid values can be obtained.
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

type CCLogData struct {
	// Data set of CC block logs
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*CCLog `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
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
	// Configuration switch
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Cache prefresh percentage. Values: 1-99
	// Note: This field may return null, indicating that no valid values can be obtained.
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

	// Name of the proxy:
	// Domain name or subdomain name when `ProxyType=hostname`
	// Instance name when `ProxyType=instance`
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
	// `hostname`: Create by subdomain name
	// `instance`: Create by instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Name of the proxy:
	// Domain name or subdomain name when `ProxyType=hostname`
	// Instance name when `ProxyType=instance`
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
	// `hostname`: Create by subdomain name
	// `instance`: Create by instance
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

	// Origin type. Valid values:
	// `custom`: Specified origins
	// `origins`: Origin group
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin information:
	// When `OriginType=custom`, it can include one or more origins in either of the following formats:
	// IP:Port
	// Domain name:Port
	// When `OriginType=origins`, it is an origin group ID.
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

	// Origin type. Valid values:
	// `custom`: Specified origins
	// `origins`: Origin group
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin information:
	// When `OriginType=custom`, it can include one or more origins in either of the following formats:
	// IP:Port
	// Domain name:Port
	// When `OriginType=origins`, it is an origin group ID.
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
type CreateCustomErrorPageRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name of the site
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Name of the file specified to be returned
	Name *string `json:"Name,omitempty" name:"Name"`

	// Content of the custom page
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateCustomErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name of the site
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Name of the file specified to be returned
	Name *string `json:"Name,omitempty" name:"Name"`

	// Content of the custom page
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *CreateCustomErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Name")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomErrorPageResponseParams struct {
	// ID of the custom page
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCustomErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomErrorPageResponseParams `json:"Response"`
}

func (r *CreateCustomErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomErrorPageResponse) FromJsonString(s string) error {
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

	// Subdomain name
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

	// Subdomain name
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
type CreateOriginGroupRequestParams struct {
	// Name of the origin group
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// Origin-pull configuration type. This field is required when `OriginType=self`.
	// `area`: Origin-pull by region
	// `weight`: Origin-pull by weight
	// When `OriginType=third_party/cos`, it can be left empty.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Origin records
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Origin type
	// `self`: Customer origin
	// `third_party`: Third-party origin
	// `cos`: Tencent Cloud COS origin
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Name of the origin group
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// Origin-pull configuration type. This field is required when `OriginType=self`.
	// `area`: Origin-pull by region
	// `weight`: Origin-pull by weight
	// When `OriginType=third_party/cos`, it can be left empty.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Origin records
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Origin type
	// `self`: Customer origin
	// `third_party`: Third-party origin
	// `cos`: Tencent Cloud COS origin
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
}

func (r *CreateOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OriginName")
	delete(f, "Type")
	delete(f, "Record")
	delete(f, "ZoneId")
	delete(f, "OriginType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupResponseParams struct {
	// ID of the newly added origin group
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateOriginGroupResponseParams `json:"Response"`
}

func (r *CreateOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOriginGroupResponse) FromJsonString(s string) error {
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

	// Resource tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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

	// Resource tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "Tags")
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

type DDoSAcl struct {
	// Destination port used as the end port
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// Destination port used as the start port
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// Source port used as the end port
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// Source port used as the start port
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// Protocol. Values: `tcp`, `udp`, and `all`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Action. Values: `drop` (Drop the request); `transmit` (Allow the request); `forward` (Continue to offer protection).
	Action *string `json:"Action,omitempty" name:"Action"`

	// Whether it is a system configuration. Values: `0` (manual configuration); `1` (system configuration).
	Default *int64 `json:"Default,omitempty" name:"Default"`
}

type DDoSAntiPly struct {
	// Enables TCP protocol blocking. `on` (enable); `off` (disable).
	DropTcp *string `json:"DropTcp,omitempty" name:"DropTcp"`

	// Enables UDP protocol blocking. `on` (enable); `off` (disable).
	DropUdp *string `json:"DropUdp,omitempty" name:"DropUdp"`

	// Enables ICMP protocol blocking. `on` (enable); `off` (disable).
	DropIcmp *string `json:"DropIcmp,omitempty" name:"DropIcmp"`

	// Enables blocking for other protocols. `on` (enable); `off` (disable).
	DropOther *string `json:"DropOther,omitempty" name:"DropOther"`

	// Number of new connections the source port can establish. Value range: 0-4294967295.
	SourceCreateLimit *int64 `json:"SourceCreateLimit,omitempty" name:"SourceCreateLimit"`

	// Number of concurrent connections the source port can establish. Value range: 0-4294967295.
	SourceConnectLimit *int64 `json:"SourceConnectLimit,omitempty" name:"SourceConnectLimit"`

	// Number of new connections the destination port can establish. Value range: 0-4294967295.
	DestinationCreateLimit *int64 `json:"DestinationCreateLimit,omitempty" name:"DestinationCreateLimit"`

	// Number of concurrent connections the destination port can establish. Value range: 0-4294967295.
	DestinationConnectLimit *int64 `json:"DestinationConnectLimit,omitempty" name:"DestinationConnectLimit"`

	// Number of abnormal connections allowed. Value range: 0-4294967295.
	AbnormalConnectNum *int64 `json:"AbnormalConnectNum,omitempty" name:"AbnormalConnectNum"`

	// Specifies the ratio of SYN exceptions to trigger alerts. Value range: 0-100
	AbnormalSynRatio *int64 `json:"AbnormalSynRatio,omitempty" name:"AbnormalSynRatio"`

	// Specifies a max number of SYN packets that triggers alarms. Value range: 0-65535
	AbnormalSynNum *int64 `json:"AbnormalSynNum,omitempty" name:"AbnormalSynNum"`

	// Connection timeout period. Value range: 0-65535.
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// Whether to enable null session protection. `0`: Disable; `1`: Enable.
	EmptyConnectProtect *string `json:"EmptyConnectProtect,omitempty" name:"EmptyConnectProtect"`

	// Whether to enable UDP fragmentation. `off`: Disable; `on`: Enable.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UdpShard *string `json:"UdpShard,omitempty" name:"UdpShard"`
}

type DDoSApplication struct {
	// Second-level domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Status of the domain name
	// `init`: NS to be switched
	// `offline`: Site acceleration not enabled with DNS
	// `process`: Deployment in progress
	// `online`: Normal
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Site acceleration switch. `on`: Enable site acceleration; `off`: Disable site acceleration. This field can be used together with `SecurityType`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// Security acceleration switch. `on`: Enable site acceleration; `off`: Disable site acceleration. This field can be used together with `AccelerateType`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SecurityType *string `json:"SecurityType,omitempty" name:"SecurityType"`
}

type DDoSConfig struct {
	// Switch
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSFeaturesFilter struct {
	// Action. `drop`: Drop the request; `transmit`: Allow the request; `drop_block`: Drop the request and block it; `forward`: Continue to offer protection.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Sets how far from the first search position
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// Sets how far from the second search position
	Depth2 *int64 `json:"Depth2,omitempty" name:"Depth2"`

	// End of the destination port
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// Start of the destination port
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// Whether to match string 1 that does not contain all the specified elements
	IsNot *int64 `json:"IsNot,omitempty" name:"IsNot"`

	// Whether to match string 2 that does not contain all the specified elements
	IsNot2 *int64 `json:"IsNot2,omitempty" name:"IsNot2"`

	// Logical operator that combines two conditions. Values: `none`, `and` and `or`. If there is only one condition, pass in `none` for this condition only.
	MatchLogic *string `json:"MatchLogic,omitempty" name:"MatchLogic"`

	// Matching method of the first condition. `pcre`: Regex match; `sunday`: String match.
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// Matching method of the second condition. `pcre`: Regex match; `sunday`: String match.
	MatchType2 *string `json:"MatchType2,omitempty" name:"MatchType2"`

	// Offset from the first search position
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Offset from the second search position
	Offset2 *int64 `json:"Offset2,omitempty" name:"Offset2"`

	// Maximum packet length
	PacketMax *int64 `json:"PacketMax,omitempty" name:"PacketMax"`

	// Minimum packet length
	PacketMin *int64 `json:"PacketMin,omitempty" name:"PacketMin"`

	// Protocol. Values: `tcp`, `udp`, `icmp` and `all`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// End of the source port
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// Start of the source port
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// String in the first condition
	Str *string `json:"Str,omitempty" name:"Str"`

	// String in the second condition
	Str2 *string `json:"Str2,omitempty" name:"Str2"`

	// Layer at which each match starts. Values: `begin_l5`, `no_match`, `begin_l3` and `begin_l4`.
	MatchBegin *string `json:"MatchBegin,omitempty" name:"MatchBegin"`

	// Layer at which each match starts. Values: `begin_l5`, `no_match`, `begin_l3` and `begin_l4`.
	MatchBegin2 *string `json:"MatchBegin2,omitempty" name:"MatchBegin2"`
}

type DDoSGeoIp struct {
	// Region information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RegionId []*int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Whether to remove all settings when empty strings are passed in. Default value: `off` (remove)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DDoSStatusInfo struct {
	// This field is not supported. Value: `off`.
	AiStatus *string `json:"AiStatus,omitempty" name:"AiStatus"`

	// User appid
	Appid *string `json:"Appid,omitempty" name:"Appid"`

	// Protection level. Values: `low`, `middle`, and `high`.
	PlyLevel *string `json:"PlyLevel,omitempty" name:"PlyLevel"`
}

type DDoSUserAllowBlockIP struct {
	// Start IP address in a specific range
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Start mask in a specific range
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mask *int64 `json:"Mask,omitempty" name:"Mask"`

	// IP type. `block`: IP blocklist; `allow`: IP allowlist.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Timestamp
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// End IP address in a specific range
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Ip2 *string `json:"Ip2,omitempty" name:"Ip2"`

	// End mask in a specific range
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mask2 *int64 `json:"Mask2,omitempty" name:"Mask2"`
}

type DDosAttackEvent struct {
	// DDoS policy group ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Attack type (corresponding to interaction event name)
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Attack status
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackStatus *int64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// Maximum attack bandwidth
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// Peak attack packet rate
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackPacketMaxRate *int64 `json:"AttackPacketMaxRate,omitempty" name:"AttackPacketMaxRate"`

	// Attack start time in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackStartTime *int64 `json:"AttackStartTime,omitempty" name:"AttackStartTime"`

	// Attack end time in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackEndTime *int64 `json:"AttackEndTime,omitempty" name:"AttackEndTime"`

	// Event ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Site ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DDosAttackEventData struct {
	// Data set of attack events
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*DDosAttackEvent `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type DDosAttackEventDetailData struct {
	// Attack status
	AttackStatus *int64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// Attack type
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Maximum bandwidth
	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`

	// Maximum packet rate
	PacketMaxRate *int64 `json:"PacketMaxRate,omitempty" name:"PacketMaxRate"`

	// Event ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// DDoS policy group ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DDosAttackSourceEvent struct {
	// Attack source IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackSourceIp *string `json:"AttackSourceIp,omitempty" name:"AttackSourceIp"`

	// Country/Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackRegion *string `json:"AttackRegion,omitempty" name:"AttackRegion"`

	// Accumulative attack traffic
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackFlow *uint64 `json:"AttackFlow,omitempty" name:"AttackFlow"`

	// Accumulative number of attack packets
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackPacketNum *uint64 `json:"AttackPacketNum,omitempty" name:"AttackPacketNum"`
}

type DDosAttackSourceEventData struct {
	// DDoS attack source data set
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*DDosAttackSourceEvent `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type DDosMajorAttackEvent struct {
	// DDoS policy group ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Maximum attack bandwidth
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// Attack time in seconds
	AttackTime *int64 `json:"AttackTime,omitempty" name:"AttackTime"`
}

type DDosMajorAttackEventData struct {
	// `DDosMajorAttackEvent` DDoS attack event
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*DDosMajorAttackEvent `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type DataItem struct {
	// Time
	Time *string `json:"Time,omitempty" name:"Time"`

	// Value
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DdosAcls struct {
	// Whether to remove all settings when empty strings are passed in. Default value: `off` (remove)
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// DDoS port filtering parameters
	Acl []*DDoSAcl `json:"Acl,omitempty" name:"Acl"`
}

type DdosAllowBlock struct {
	// Whether to remove all settings when empty strings are passed in. Default value: `off` (remove)
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Array of objects in blocklist/allowlist configuration
	UserAllowBlockIp []*DDoSUserAllowBlockIP `json:"UserAllowBlockIp,omitempty" name:"UserAllowBlockIp"`
}

type DdosPacketFilter struct {
	// Whether to remove all settings when empty strings are passed in. Default value: `off` (remove)
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Array of objects in feature filtering configuration
	PacketFilter []*DDoSFeaturesFilter `json:"PacketFilter,omitempty" name:"PacketFilter"`
}

type DdosRule struct {
	// DDoS mitigation level
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DdosStatusInfo *DDoSStatusInfo `json:"DdosStatusInfo,omitempty" name:"DdosStatusInfo"`

	// DDoS regional blocking
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DdosGeoIp *DDoSGeoIp `json:"DdosGeoIp,omitempty" name:"DdosGeoIp"`

	// DDoS blocklist/allowlist
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DdosAllowBlock *DdosAllowBlock `json:"DdosAllowBlock,omitempty" name:"DdosAllowBlock"`

	// Protocol blocking and null session protection
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DdosAntiPly *DDoSAntiPly `json:"DdosAntiPly,omitempty" name:"DdosAntiPly"`

	// DDoS feature filtering
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DdosPacketFilter *DdosPacketFilter `json:"DdosPacketFilter,omitempty" name:"DdosPacketFilter"`

	// DDoS port filtering
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DdosAcl *DdosAcls `json:"DdosAcl,omitempty" name:"DdosAcl"`

	// DDoS mitigation switch. `on`: Enable; `off`: Disable.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Whether to enable UDP fragmentation. `on`: Enable; `off`: Disable.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UdpShardOpen *string `json:"UdpShardOpen,omitempty" name:"UdpShardOpen"`
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
type DeleteOriginGroupRequestParams struct {
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DeleteOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DeleteOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OriginId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupResponseParams struct {
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOriginGroupResponseParams `json:"Response"`
}

func (r *DeleteOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOriginGroupResponse) FromJsonString(s string) error {
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

	// Name of the proxy:
	// Domain name or subdomain name when `ProxyType=hostname`
	// Instance name when `ProxyType=instance`
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
	// `hostname`: Create by subdomain name
	// `instance`: Create by instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// When `ProxyType=hostname`:
	// `ProxyName` indicates a specified domain name, such as test.123.com
	// `HostId` indicates a unique ID of the domain name.
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

	// Disused
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`

	// When `PlatType` is `ip`, it indicates the number of proxies that schedule via Anycast IP.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IpCount *uint64 `json:"IpCount,omitempty" name:"IpCount"`

	// When `PlatType` is `domain`, it indicates the number of proxies that schedule via CNAME.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
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
type DescribeBotLogRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Query condition
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeBotLogRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Query condition
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

func (r *DescribeBotLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotLogResponseParams struct {
	// Bot attack data
	Data *BotLogData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotLogResponseParams `json:"Response"`
}

func (r *DescribeBotLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotManagedRulesRequestParams struct {
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/layer-4 proxy
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Total number of pages
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Number of rules per page
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// Rule type. Values: `idcid`, `sipbot` and `uabot`. All rules will be returned if this field is not specified.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

type DescribeBotManagedRulesRequest struct {
	*tchttp.BaseRequest
	
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/layer-4 proxy
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Total number of pages
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Number of rules per page
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// Rule type. Values: `idcid`, `sipbot` and `uabot`. All rules will be returned if this field is not specified.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

func (r *DescribeBotManagedRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotManagedRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Page")
	delete(f, "PerPage")
	delete(f, "RuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotManagedRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotManagedRulesResponseParams struct {
	// Number of bot managed rules returned
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Bot managed rules
	Rules []*BotManagedRuleDetail `json:"Rules,omitempty" name:"Rules"`

	// Total number of bot managed rules
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotManagedRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotManagedRulesResponseParams `json:"Response"`
}

func (r *DescribeBotManagedRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotManagedRulesResponse) FromJsonString(s string) error {
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
type DescribeDDoSPolicyRequestParams struct {
	// Policy group ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Top-level domain name (site)
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Policy group ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Top-level domain name (site)
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSPolicyResponseParams `json:"Response"`
}

func (r *DescribeDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackDataRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metrics
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Port number
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Protocol type. Valid values: tcp, udp, all
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Attack type. Valid values: flood, icmpFlood..., all
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeDDosAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metrics
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Port number
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Protocol type. Valid values: tcp, udp, all
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Attack type. Valid values: flood, icmpFlood..., all
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeDDosAttackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Port")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackDataResponseParams struct {
	// DDoS attack data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned data
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackDataResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventDetailRequestParams struct {
	// Event ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeDDosAttackEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// Event ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeDDosAttackEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventDetailResponseParams struct {
	// DDoS attack event details
	Data *DDosAttackEventDetailData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackEventDetailResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Protocol type. Valid values: {tcp,udp,all}
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Whether to show details. Valid values: Y (yes), N (no).
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

type DescribeDDosAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Protocol type. Valid values: {tcp,udp,all}
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Whether to show details. Valid values: Y (yes), N (no).
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

func (r *DescribeDDosAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "ProtocolType")
	delete(f, "IsShowDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackEventResponseParams struct {
	// DDoS attack event data
	Data *DDosAttackEventData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackSourceEventRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Protocol type. Valid values: {tcp,udp,all}
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

type DescribeDDosAttackSourceEventRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Protocol type. Valid values: {tcp,udp,all}
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

func (r *DescribeDDosAttackSourceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackSourceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "ProtocolType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackSourceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackSourceEventResponseParams struct {
	// DDoS attack source data
	Data *DDosAttackSourceEventData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackSourceEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackSourceEventResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackSourceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackSourceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackTopDataRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter metric
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Number of the top data entries to query. 0: queries all data entries.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Port number
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Protocol type. Valid values: tcp, udp, all
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Attack type. Valid values: flood, icmpFlood..., all
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

type DescribeDDosAttackTopDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter metric
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Number of the top data entries to query. 0: queries all data entries.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Port number
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Protocol type. Valid values: tcp, udp, all
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Attack type. Valid values: flood, icmpFlood..., all
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *DescribeDDosAttackTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Limit")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Port")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosAttackTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosAttackTopDataResponseParams struct {
	// Top N data
	Data []*TopNEntry `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosAttackTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosAttackTopDataResponseParams `json:"Response"`
}

func (r *DescribeDDosAttackTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosAttackTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosMajorAttackEventRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Protocol type. Valid values: {tcp,udp,all}
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

type DescribeDDosMajorAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Protocol type. Valid values: {tcp,udp,all}
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

func (r *DescribeDDosMajorAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosMajorAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ProtocolType")
	delete(f, "ZoneIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDosMajorAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDosMajorAttackEventResponseParams struct {
	// Major DDoS attack event
	Data *DDosMajorAttackEventData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDosMajorAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDosMajorAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDosMajorAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDosMajorAttackEventResponse) FromJsonString(s string) error {
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
type DescribeOriginGroupDetailRequestParams struct {
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeOriginGroupDetailRequest struct {
	*tchttp.BaseRequest
	
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeOriginGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OriginId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupDetailResponseParams struct {
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Origin group name
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// Origin-pull configuration type
	// `area`: Origin-pull by the client IP’s region specified by `Area` in OriginRecord.
	// `weight`: Origin-pull by the weight specified by `Weight` in OriginRecord.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Record
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Site name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Origin type
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Whether the origin group uses layer-4 proxy.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationProxyUsed *bool `json:"ApplicationProxyUsed,omitempty" name:"ApplicationProxyUsed"`

	// Whether the origin group is used for load balancing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancingUsed *bool `json:"LoadBalancingUsed,omitempty" name:"LoadBalancingUsed"`

	// Proxy mode of the load balancing task associated with the origin group.
	// `none`: Not used for load balancing.
	// `dns_only`: Used for DNS-only load balancing.
	// `proxied`: Used for proxied load balancing.
	// `both`: Used for both DNS-only and proxied load balancing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancingUsedType *string `json:"LoadBalancingUsedType,omitempty" name:"LoadBalancingUsedType"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOriginGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginGroupDetailResponseParams `json:"Response"`
}

func (r *DescribeOriginGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupRequestParams struct {
	// Pagination parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameters
	Filters []*OriginFilter `json:"Filters,omitempty" name:"Filters"`

	// Site ID
	// If it’s not specified, all origin groups will be obtained.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Pagination parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameters
	Filters []*OriginFilter `json:"Filters,omitempty" name:"Filters"`

	// Site ID
	// If it’s not specified, all origin groups will be obtained.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupResponseParams struct {
	// Origin group information
	Data []*OriginGroup `json:"Data,omitempty" name:"Data"`

	// Total number of records
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginGroupResponseParams `json:"Response"`
}

func (r *DescribeOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataRequestParams struct {
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l7Flow_outFlux`: Access traffic
	// `l7Flow_request`: Access requests
	// `l7Flow_outBandwidth`: Access bandwidth
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// Time interval. Valid values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// List of `ZoneId` values. This parameter takes effect only when querying in the zone/domain dimension.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of `Domain` values. This parameter takes effect only when querying in the domain dimension.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Protocol type. Valid values: {http,http2,https,all}
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeOverviewL7DataRequest struct {
	*tchttp.BaseRequest
	
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l7Flow_outFlux`: Access traffic
	// `l7Flow_request`: Access requests
	// `l7Flow_outBandwidth`: Access bandwidth
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// Time interval. Valid values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// List of `ZoneId` values. This parameter takes effect only when querying in the zone/domain dimension.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of `Domain` values. This parameter takes effect only when querying in the domain dimension.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Protocol type. Valid values: {http,http2,https,all}
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeOverviewL7DataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewL7DataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewL7DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataResponseParams struct {
	// Query dimension
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time interval
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Detailed data
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewL7DataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewL7DataResponseParams `json:"Response"`
}

func (r *DescribeOverviewL7DataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewL7DataResponse) FromJsonString(s string) error {
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
type DescribeSecurityPolicyListRequestParams struct {
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSecurityPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSecurityPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyListResponseParams struct {
	// List of protected resources
	Entities []*SecurityEntity `json:"Entities,omitempty" name:"Entities"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyListResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyManagedRulesIdRequestParams struct {
	// List of rule IDs
	RuleId []*int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type DescribeSecurityPolicyManagedRulesIdRequest struct {
	*tchttp.BaseRequest
	
	// List of rule IDs
	RuleId []*int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeSecurityPolicyManagedRulesIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyManagedRulesIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyManagedRulesIdResponseParams struct {
	// Total number of returned items
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Managed rule
	Rules []*ManagedRule `json:"Rules,omitempty" name:"Rules"`

	// Total number of returned items
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyManagedRulesIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyManagedRulesIdResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyManagedRulesIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyManagedRulesRequestParams struct {
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/layer-4 proxy
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Total number of pages
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Number of rules per page
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

type DescribeSecurityPolicyManagedRulesRequest struct {
	*tchttp.BaseRequest
	
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/layer-4 proxy
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Total number of pages
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Number of rules per page
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *DescribeSecurityPolicyManagedRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Page")
	delete(f, "PerPage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyManagedRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyManagedRulesResponseParams struct {
	// Number of rules returned
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Managed rule
	Rules []*ManagedRule `json:"Rules,omitempty" name:"Rules"`

	// Total number of rules
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyManagedRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyManagedRulesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyManagedRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyManagedRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRegionsRequestParams struct {

}

type DescribeSecurityPolicyRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecurityPolicyRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRegionsResponseParams struct {
	// Total number of regions
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Region information
	GeoIp []*GeoIp `json:"GeoIp,omitempty" name:"GeoIp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyRegionsResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRequestParams struct {
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Second-level domain name
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Second-level domain name
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

func (r *DescribeSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyResponseParams struct {
	// User ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Second-level domain name
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Security configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Config *SecurityConfig `json:"Config,omitempty" name:"Config"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPortraitRulesRequestParams struct {
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/Application name
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

type DescribeSecurityPortraitRulesRequest struct {
	*tchttp.BaseRequest
	
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/Application name
	Entity *string `json:"Entity,omitempty" name:"Entity"`
}

func (r *DescribeSecurityPortraitRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPortraitRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPortraitRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPortraitRulesResponseParams struct {
	// Number of rules returned in this request
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Bot user profiling rule
	Rules []*PortraitManagedRuleDetail `json:"Rules,omitempty" name:"Rules"`

	// Total number of rules
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPortraitRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPortraitRulesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPortraitRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPortraitRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataRequestParams struct {
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l4Flow_connections`: Access connections
	// `l4Flow_flux`: Access traffic
	// `l4Flow_inFlux`: Inbound traffic
	// `l4Flow_outFlux`: Outbound traffic
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// This field has been disused. Use `ProxyIds` instead.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// This field is not supported currently.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Time interval. Valid values: {min, 5min, hour, day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// This field is not supported currently. Use `Filter` instead.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Supported filters: `proxyd`, `ruleId`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// List of layer-4 proxies
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l4Flow_connections`: Access connections
	// `l4Flow_flux`: Access traffic
	// `l4Flow_inFlux`: Inbound traffic
	// `l4Flow_outFlux`: Outbound traffic
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// This field has been disused. Use `ProxyIds` instead.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// This field is not supported currently.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Time interval. Valid values: {min, 5min, hour, day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// This field is not supported currently. Use `Filter` instead.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Supported filters: `proxyd`, `ruleId`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// List of layer-4 proxies
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`
}

func (r *DescribeTimingL4DataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL4DataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "InstanceIds")
	delete(f, "Protocol")
	delete(f, "Interval")
	delete(f, "RuleId")
	delete(f, "Filters")
	delete(f, "ProxyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL4DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataResponseParams struct {
	// Query dimension
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time interval
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Detailed data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimingL4DataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL4DataResponseParams `json:"Response"`
}

func (r *DescribeTimingL4DataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL4DataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataRequestParams struct {
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l7Flow_outFlux`: Access traffic
	// `l7Flow_request`: Access requests
	// `l7Flow_outBandwidth`: Access bandwidth
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// Time interval. Valid values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Array of `ZoneId` values
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l7Flow_outFlux`: Access traffic
	// `l7Flow_request`: Access requests
	// `l7Flow_outBandwidth`: Access bandwidth
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// Time interval. Valid values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Array of `ZoneId` values
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTimingL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataResponseParams struct {
	// Detailed data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// Query dimension
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time interval
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimingL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeTimingL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataRequestParams struct {
	// Start time of the query (client time in RFC 3339)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Start time of the query (client time in RFC 3339)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l7Cache_outFlux`: Access traffic
	// `l7Cache_request`: Access requests
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// Time interval. Values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter condition:
	// {Key: "cacheType", Value: ["hit"], Operator: "equals"}: Filter by data responded from EdgeOne
	// {Key: "cacheType", Value: ["miss", "dynamic"], Operator: "equals"}: Filter by data responded from the origin server
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTimingL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query (client time in RFC 3339)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Start time of the query (client time in RFC 3339)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported metrics for data query:
	// `l7Cache_outFlux`: Access traffic
	// `l7Cache_request`: Access requests
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// Time interval. Values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter condition:
	// {Key: "cacheType", Value: ["hit"], Operator: "equals"}: Filter by data responded from EdgeOne
	// {Key: "cacheType", Value: ["miss", "dynamic"], Operator: "equals"}: Filter by data responded from the origin server
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTimingL7CacheDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7CacheDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataResponseParams struct {
	// Details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// Metric specified for data query
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time interval
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTimingL7CacheDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL7CacheDataResponseParams `json:"Response"`
}

func (r *DescribeTimingL7CacheDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7CacheDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataRequestParams struct {
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Time series-type access traffic metric
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Top N. 0 indicates to return the full data.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Time interval. Valid values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Array of `ZoneId` values
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// Client time in RFC 3339 format
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Client time in RFC 3339 format
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Time series-type access traffic metric
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Top N. 0 indicates to return the full data.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Time interval. Valid values: {min, 5min, hour, day, week}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Array of `ZoneId` values
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTopL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Limit")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataResponseParams struct {
	// Top detailed data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// Query dimension
	Type *string `json:"Type,omitempty" name:"Type"`

	// Query metric
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeTopL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataRequestParams struct {
	// Start time of the query (client time in RFC 3339)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query (client time in RFC 3339)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metric for time-series data query
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Specifies the number of data records to return. If `0` is passed in, all data is returned.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Time interval. Values: {min, 5min, hour, day, week}. This field is optional.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Array of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter condition
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTopL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query (client time in RFC 3339)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query (client time in RFC 3339)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metric for time-series data query
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Specifies the number of data records to return. If `0` is passed in, all data is returned.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Time interval. Values: {min, 5min, hour, day, week}. This field is optional.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Array of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter condition
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTopL7CacheDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7CacheDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Limit")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataResponseParams struct {
	// Top-ranked data details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// Dimension specified for data query
	Type *string `json:"Type,omitempty" name:"Type"`

	// Metric specified for data query
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopL7CacheDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopL7CacheDataResponseParams `json:"Response"`
}

func (r *DescribeTopL7CacheDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7CacheDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesAttackEventsRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// List of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Whether to show details. Valid values: Y (yes), N (no).
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

type DescribeWebManagedRulesAttackEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// List of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Whether to show details. Valid values: Y (yes), N (no).
	IsShowDetail *string `json:"IsShowDetail,omitempty" name:"IsShowDetail"`
}

func (r *DescribeWebManagedRulesAttackEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesAttackEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "IsShowDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesAttackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesAttackEventsResponseParams struct {
	// Web attack event data
	Data *WebEventData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned data
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesAttackEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesAttackEventsResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesAttackEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesAttackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesDataRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metrics
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Protocol type
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// "webshell" : WebShell detection prevention
	// "oa" : Common OA vulnerability prevention
	// "xss" : XSS attack prevention
	// "xxe" : XXE attack prevention
	// "webscan" : Scanner attack vulnerability prevention
	// "cms" : Common CMS vulnerability prevention
	// "upload" : Malicious file upload attack prevention
	// "sql" : SQL injection attack prevention
	// "cmd_inject": Command/Code injection attack prevention
	// "osc" : Open-source component vulnerability prevention
	// "file_read" : Arbitrary file read prevention
	// "ldap" : LDAP injection attack prevention
	// "other" : Other vulnerability prevention
	// 
	// "all":"All"
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeWebManagedRulesDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metrics
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Protocol type
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// "webshell" : WebShell detection prevention
	// "oa" : Common OA vulnerability prevention
	// "xss" : XSS attack prevention
	// "xxe" : XXE attack prevention
	// "webscan" : Scanner attack vulnerability prevention
	// "cms" : Common CMS vulnerability prevention
	// "upload" : Malicious file upload attack prevention
	// "sql" : SQL injection attack prevention
	// "cmd_inject": Command/Code injection attack prevention
	// "osc" : Open-source component vulnerability prevention
	// "file_read" : Arbitrary file read prevention
	// "ldap" : LDAP injection attack prevention
	// "other" : Other vulnerability prevention
	// 
	// "all":"All"
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeWebManagedRulesDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesDataResponseParams struct {
	// Web attack log entity
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesDataResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesLogRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Query condition
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeWebManagedRulesLogRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Query condition
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

func (r *DescribeWebManagedRulesLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesLogResponseParams struct {
	// Web attack log data
	Data *WebLogData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesLogResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesTopDataRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter metric
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Number of the top data entries to query. 0: queries all data entries.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Port number
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Protocol type. Valid values: tcp, udp, all
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Attack type. Valid values: flood, icmpFlood..., all
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type DescribeWebManagedRulesTopDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter metric
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Number of the top data entries to query. 0: queries all data entries.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Set of DDoS policy group IDs
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Port number
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Protocol type. Valid values: tcp, udp, all
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Attack type. Valid values: flood, icmpFlood..., all
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *DescribeWebManagedRulesTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Limit")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Port")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesTopDataResponseParams struct {
	// Top N data
	Data []*TopNEntry `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesTopDataResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionAttackEventsRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Domain name
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

type DescribeWebProtectionAttackEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Domain name
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

func (r *DescribeWebProtectionAttackEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionAttackEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Domains")
	delete(f, "ZoneIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionAttackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionAttackEventsResponseParams struct {
	// DDoS attack event data
	Data *CCInterceptEventData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionAttackEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionAttackEventsResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionAttackEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionAttackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionDataRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metrics
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Protocol type
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// "webshell" : WebShell detection prevention
	// "oa" : Common OA vulnerability prevention
	// "xss" : XSS attack prevention
	// "xxe" : XXE attack prevention
	// "webscan" : Scanner attack vulnerability prevention
	// "cms" : Common CMS vulnerability prevention
	// "upload" : Malicious file upload attack prevention
	// "sql" : SQL injection attack prevention
	// "cmd_inject": Command/Code injection attack prevention
	// "osc" : Open-source component vulnerability prevention
	// "file_read" : Arbitrary file read prevention
	// "ldap" : LDAP injection attack prevention
	// "other" : Other vulnerability prevention
	// 
	// "all":"All"
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeWebProtectionDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metrics
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of site IDs
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Protocol type
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// "webshell" : WebShell detection prevention
	// "oa" : Common OA vulnerability prevention
	// "xss" : XSS attack prevention
	// "xxe" : XXE attack prevention
	// "webscan" : Scanner attack vulnerability prevention
	// "cms" : Common CMS vulnerability prevention
	// "upload" : Malicious file upload attack prevention
	// "sql" : SQL injection attack prevention
	// "cmd_inject": Command/Code injection attack prevention
	// "osc" : Open-source component vulnerability prevention
	// "file_read" : Arbitrary file read prevention
	// "ldap" : LDAP injection attack prevention
	// "other" : Other vulnerability prevention
	// 
	// "all":"All"
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeWebProtectionDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "ProtocolType")
	delete(f, "AttackType")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionDataResponseParams struct {
	// Data details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Query time granularity. Valid values: {min,5min,hour,day}
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionDataResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionLogRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Query condition
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

type DescribeWebProtectionLogRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of items per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Site set
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Domain name set
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Query condition
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`
}

func (r *DescribeWebProtectionLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionLogResponseParams struct {
	// Block data in rate-limiting policy
	Data *CCLogData `json:"Data,omitempty" name:"Data"`

	// Status. 1: failed; 0: succeeded
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Returned message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionLogResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDDoSPolicyRequestParams struct {
	// ID of the site (top-level domain name)
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeZoneDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site (top-level domain name)
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeZoneDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDDoSPolicyResponseParams struct {
	// User APPID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// DDoS mitigation configuration
	ShieldAreas []*ShieldArea `json:"ShieldAreas,omitempty" name:"ShieldAreas"`

	// Includes the details of all subdomain names
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Domains []*DDoSApplication `json:"Domains,omitempty" name:"Domains"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneDDoSPolicyResponseParams `json:"Response"`
}

func (r *DescribeZoneDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDDoSPolicyResponse) FromJsonString(s string) error {
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

	// Specifies whether to enable CNAME acceleration
	// - `enabled`: Enable
	// - `disabled`: Disable
	CnameSpeedUp *string `json:"CnameSpeedUp,omitempty" name:"CnameSpeedUp"`

	// Ownership verification status of the site when it accesses via CNAME.
	// - `finished`: The site is verified.
	// - `pending`: The site is waiting for verification.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

	// Resource tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`


	Area *string `json:"Area,omitempty" name:"Area"`

	// Billable resource
	// Note: This field may return null, indicating that no valid values can be obtained.
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// Site modification date
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// Site creation date
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// User-defined name server information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// User-defined name server IP information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitempty" name:"VanityNameServersIps"`

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

	// Cache prefresh configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
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

	// Which service is enabled for the domain name.
	// - `lb`: Load balancing
	// - `security`: Security acceleration
	// - `l4`: L4 proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
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

	// Array of site names
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Array of subdomain names
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

	// Array of site names
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Array of subdomain names
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

type Filter struct {
	// Filter dimension
	Key *string `json:"Key,omitempty" name:"Key"`

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Filter dimension value
	Value []*string `json:"Value,omitempty" name:"Value"`
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

type GeoIp struct {
	// Region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Country name
	Country *string `json:"Country,omitempty" name:"Country"`

	// Continent name
	Continent *string `json:"Continent,omitempty" name:"Continent"`

	// Country name in English
	CountryEn *string `json:"CountryEn,omitempty" name:"CountryEn"`

	// Continent name in English
	ContinentEn *string `json:"ContinentEn,omitempty" name:"ContinentEn"`
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

type IntelligenceRule struct {
	// Switch
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Items in a bot intelligence rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*IntelligenceRuleItem `json:"Items,omitempty" name:"Items"`
}

type IntelligenceRuleItem struct {
	// Malicious bot, which is used to tag bad bots
	// Note: This field may return null, indicating that no valid values can be obtained.
	Label *string `json:"Label,omitempty" name:"Label"`

	// Action
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`
}

type IpTableConfig struct {
	// Switch
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// []
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*IpTableRule `json:"Rules,omitempty" name:"Rules"`
}

type IpTableRule struct {
	// Action: `drop` (block), `trans` (allow), `monitor` (observe)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Matches by IP or region. Values: `ip` and `area`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// Matching content
	// Note: This field may return null, indicating that no valid values can be obtained.
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`

	// Rule ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type L7OfflineLog struct {
	// Start time of the log packaging
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// Subdomain name
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

type ManagedRule struct {
	// ID of the rule
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Rule description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Rule type
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// Rule level
	RuleLevelDesc *string `json:"RuleLevelDesc,omitempty" name:"RuleLevelDesc"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Rule status: `block`, `allow`
	Status *string `json:"Status,omitempty" name:"Status"`

	// Tag of the rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleTags []*string `json:"RuleTags,omitempty" name:"RuleTags"`

	// Description of the rule type
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleTypeDesc *string `json:"RuleTypeDesc,omitempty" name:"RuleTypeDesc"`

	// ID of the rule type
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleTypeId *int64 `json:"RuleTypeId,omitempty" name:"RuleTypeId"`
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

	// ID of the proxy
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Name of the proxy:
	// Domain name or subdomain name when `ProxyType=hostname`
	// Instance name when `ProxyType=instance`
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// This parameter is disused.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This parameter is disused.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// Session persistence time. Value range: 30-3600 (in seconds).
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Create by subdomain name
	// `instance`: Create by instance
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// ID of the proxy
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Name of the proxy:
	// Domain name or subdomain name when `ProxyType=hostname`
	// Instance name when `ProxyType=instance`
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// This parameter is disused.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// This parameter is disused.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// Session persistence time. Value range: 30-3600 (in seconds).
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// Specifies how a layer-4 proxy is created.
	// `hostname`: Create by subdomain name
	// `instance`: Create by instance
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
	// ID of the proxy
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
	// `origins`: Origin group
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information:
	// When `OriginType=custom`, it indicates one or more origin servers. Example:
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]
	// 
	// When `OriginType=origins`, it indicates an origin group ID. Example:
	// OriginValue=["origin-xxx"]
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
	// `origins`: Origin group
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information:
	// When `OriginType=custom`, it indicates one or more origin servers. Example:
	// OriginValue=["8.8.8.8:80","9.9.9.9:80"]
	// OriginValue=["test.com:80"]
	// 
	// When `OriginType=origins`, it indicates an origin group ID. Example:
	// OriginValue=["origin-xxx"]
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

	// ID of the proxy
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

	// ID of the proxy
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

	// ID of the proxy
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

	// ID of the proxy
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
	// ID of the proxy
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
type ModifyDDoSPolicyHostRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Second-level domain name
	Host *string `json:"Host,omitempty" name:"Host"`

	// Whether to enable content acceleration. Values: `on` (enable content acceleration), and `off` (disable content acceleration). It can be used together with `SecurityType`.
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// Policy ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Whether to enable security protection. Values: `on` (enable security protection) and `off` (disable security protection). It can be used together with `AccelerateType`.
	SecurityType *string `json:"SecurityType,omitempty" name:"SecurityType"`
}

type ModifyDDoSPolicyHostRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Second-level domain name
	Host *string `json:"Host,omitempty" name:"Host"`

	// Whether to enable content acceleration. Values: `on` (enable content acceleration), and `off` (disable content acceleration). It can be used together with `SecurityType`.
	AccelerateType *string `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// Policy ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Whether to enable security protection. Values: `on` (enable security protection) and `off` (disable security protection). It can be used together with `AccelerateType`.
	SecurityType *string `json:"SecurityType,omitempty" name:"SecurityType"`
}

func (r *ModifyDDoSPolicyHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Host")
	delete(f, "AccelerateType")
	delete(f, "PolicyId")
	delete(f, "SecurityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyHostResponseParams struct {
	// Subdomain name
	Host *string `json:"Host,omitempty" name:"Host"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSPolicyHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSPolicyHostResponseParams `json:"Response"`
}

func (r *ModifyDDoSPolicyHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyRequestParams struct {
	// ID of the policy group
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Detailed DDoS mitigation configuration
	DdosRule *DdosRule `json:"DdosRule,omitempty" name:"DdosRule"`
}

type ModifyDDoSPolicyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the policy group
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Detailed DDoS mitigation configuration
	DdosRule *DdosRule `json:"DdosRule,omitempty" name:"DdosRule"`
}

func (r *ModifyDDoSPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "ZoneId")
	delete(f, "DdosRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSPolicyResponseParams struct {
	// ID of the policy group
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSPolicyResponseParams `json:"Response"`
}

func (r *ModifyDDoSPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSPolicyResponse) FromJsonString(s string) error {
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
type ModifyOriginGroupRequestParams struct {
	// ID of the origin group
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Name of the origin group
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// Origin-pull configuration type. This field is required when `OriginType=self`.
	// `area`: Origin-pull by region
	// `weight`: Origin-pull by weight
	// When `OriginType=third_party/cos`, it can be left empty.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Origin record
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Origin type
	// `self`: Customer origin
	// `third_party`: Third-party origin
	// `cos`: Tencent Cloud COS origin
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
}

type ModifyOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// ID of the origin group
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Name of the origin group
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// Origin-pull configuration type. This field is required when `OriginType=self`.
	// `area`: Origin-pull by region
	// `weight`: Origin-pull by weight
	// When `OriginType=third_party/cos`, it can be left empty.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Origin record
	Record []*OriginRecord `json:"Record,omitempty" name:"Record"`

	// Site ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Origin type
	// `self`: Customer origin
	// `third_party`: Third-party origin
	// `cos`: Tencent Cloud COS origin
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`
}

func (r *ModifyOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OriginId")
	delete(f, "OriginName")
	delete(f, "Type")
	delete(f, "Record")
	delete(f, "ZoneId")
	delete(f, "OriginType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupResponseParams struct {
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOriginGroupResponseParams `json:"Response"`
}

func (r *ModifyOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyRequestParams struct {
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/layer-4 proxy
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Security configuration
	Config *SecurityConfig `json:"Config,omitempty" name:"Config"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain name/layer-4 proxy
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Security configuration
	Config *SecurityConfig `json:"Config,omitempty" name:"Config"`
}

func (r *ModifySecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityPolicyResponseParams `json:"Response"`
}

func (r *ModifySecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyResponse) FromJsonString(s string) error {
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

	// Cache prefresh configuration
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

	// Cache prefresh configuration
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
	// `healthy`: Healthy; `unhealthy`: Unhealthy; `process`: Checking origin.
	Status *string `json:"Status,omitempty" name:"Status"`

	// List of unhealthy origin groups when `Status = unhealthy`
	// Note: This field may return null, indicating that no valid values can be obtained.
	Host []*string `json:"Host,omitempty" name:"Host"`
}

type OriginFilter struct {
	// Field to be filtered. Supported field: name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value of the field
	Value *string `json:"Value,omitempty" name:"Value"`
}

type OriginGroup struct {
	// Origin group ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Origin group name
	OriginName *string `json:"OriginName,omitempty" name:"OriginName"`

	// Origin-pull configuration type
	// `area`: Origin-pull by the client IP’s region specified by `Area` in `OriginRecord`.
	// `weight`: Origin-pull by the weight specified by `Weight` in `OriginRecord`.
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

	// Whether the origin group uses layer-4 proxy.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationProxyUsed *bool `json:"ApplicationProxyUsed,omitempty" name:"ApplicationProxyUsed"`

	// Whether the origin group is used for load balancing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancingUsed *bool `json:"LoadBalancingUsed,omitempty" name:"LoadBalancingUsed"`

	// Origin status 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *OriginCheckOriginStatus `json:"Status,omitempty" name:"Status"`

	// Proxy mode of the load balancing task associated with the origin group.
	// `none`: This origin group is not used for load balancing.
	// `dns_only`: Used for DNS-only load balancing 
	// `proxied`: Used for proxied load balancing
	// `both`: It’s used for both DNS-only and proxied load balancing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadBalancingUsedType *string `json:"LoadBalancingUsedType,omitempty" name:"LoadBalancingUsedType"`
}

type OriginRecord struct {
	// Record value
	Record *string `json:"Record,omitempty" name:"Record"`

	// A specific region when `Type=area`.
	// The default region when `Type` is not specified.
	Area []*string `json:"Area,omitempty" name:"Area"`

	// A specific weight when `Type=weight`.
	// The value range is [1-100].
	// The total weight of multiple origins in an origin group should be 100.
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


	Proto *string `json:"Proto,omitempty" name:"Proto"`
}

type OriginRecordPrivateParameter struct {
	// Name of the private origin authentication parameter.
	// `AccessKeyId`: Access key ID
	// `SecretAccessKey`: Secret access key
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value of the private origin authentication parameter
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PortraitManagedRuleDetail struct {
	// Unique rule ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Rule description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Rule type name: botdb (user profile)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// Rule feature category ID (scanner, bot behavior, etc.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClassificationId *int64 `json:"ClassificationId,omitempty" name:"ClassificationId"`

	// Current rule action status (block, alg, etc.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`
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

type QueryCondition struct {
	// Dimension
	Key *string `json:"Key,omitempty" name:"Key"`

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Dimension value
	Value []*string `json:"Value,omitempty" name:"Value"`
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

type RateLimitConfig struct {
	// Switch
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Rate limit rule
	UserRules []*RateLimitUserRule `json:"UserRules,omitempty" name:"UserRules"`

	// Default template
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Template *RateLimitTemplate `json:"Template,omitempty" name:"Template"`

	// Client filtering
	// Note: This field may return null, indicating that no valid values can be obtained.
	Intelligence *RateLimitIntelligence `json:"Intelligence,omitempty" name:"Intelligence"`
}

type RateLimitIntelligence struct {
	// Whether to enable this feature
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Action. Values: `monitor` (observe), `alg` (JS/Managed challenge)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`
}

type RateLimitTemplate struct {
	// Template name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Template details
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Detail *RateLimitTemplateDetail `json:"Detail,omitempty" name:"Detail"`
}

type RateLimitTemplateDetail struct {
	// Template name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Unique ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Action
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Time it takes to perform the action
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// Request rate threshold
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// Statistical period
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type RateLimitUserRule struct {
	// Rate threshold
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// Data collection time
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Name of the rule
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Action: `monitor` (Observe), `drop` (Block)
	Action *string `json:"Action,omitempty" name:"Action"`

	// Time it takes to perform the action
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// Time unit: second
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// Status of the rule
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// Rule
	Conditions []*ACLCondition `json:"Conditions,omitempty" name:"Conditions"`

	// Priority of the rule
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// ID of the rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// Word filter
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FreqFields []*string `json:"FreqFields,omitempty" name:"FreqFields"`

	// Update time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
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

type Resource struct {
	// Resource ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Billing mode
	// `0`: Pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Effective time
	EnableTime *string `json:"EnableTime,omitempty" name:"EnableTime"`

	// Expiration time
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Status of the plan
	Status *string `json:"Status,omitempty" name:"Status"`

	// Pricing query parameter
	Sv []*Sv `json:"Sv,omitempty" name:"Sv"`

	// Specifies whether to enable auto-renewal
	// `0`: Default
	// `1`: Enable auto-renewal
	// `2`: Disable auto-renewal
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// ID of the plan
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`


	Area *string `json:"Area,omitempty" name:"Area"`
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

type SecEntry struct {
	// Entry key
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Entry value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value []*SecEntryValue `json:"Value,omitempty" name:"Value"`
}

type SecEntryValue struct {
	// Metric name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Metric data details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`

	// Maximum
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// Average
	// Note: This field may return null, indicating that no valid values can be obtained.
	Avg *float64 `json:"Avg,omitempty" name:"Avg"`

	// Sum
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sum *float64 `json:"Sum,omitempty" name:"Sum"`
}

type SecurityConfig struct {
	// WAF configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	WafConfig *WafConfig `json:"WafConfig,omitempty" name:"WafConfig"`

	// Rate limit configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitempty" name:"RateLimitConfig"`

	// DDoS mitigation configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DdosConfig *DDoSConfig `json:"DdosConfig,omitempty" name:"DdosConfig"`

	// ACL configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`

	// Bot configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BotConfig *BotConfig `json:"BotConfig,omitempty" name:"BotConfig"`

	// Switch that controls all web security configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitempty" name:"SwitchConfig"`

	// IP blocklist/allowlist
	// Note: This field may return null, indicating that no valid values can be obtained.
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitempty" name:"IpTableConfig"`
}

type SecurityEntity struct {
	// User APPID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Top-level domain name
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Second-level domain name
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Type of protected resource. Values: `domain` and `application`.
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
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

type ShieldArea struct {
	// ID of the site (top-level domain name)
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Policy ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Type of protected resource. Values: `domain` and `application`.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Layer-4 proxy name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`

	// Layer-7 domain name parameters
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Application []*DDoSApplication `json:"Application,omitempty" name:"Application"`

	// Number of layer-4 TCP forwarding rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TcpNum *int64 `json:"TcpNum,omitempty" name:"TcpNum"`

	// Number of layer-4 UDP forwarding rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UdpNum *int64 `json:"UdpNum,omitempty" name:"UdpNum"`

	// Name of the protected resource
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Whether the shared resource is used. Values: `true` (yes) and `false` (no). The proxy mode cannot be switched when the shared resource is used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Share *bool `json:"Share,omitempty" name:"Share"`
}

type SmartRouting struct {
	// Whether to enable smart acceleration
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Sv struct {
	// Parameter key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SwitchConfig struct {
	// Switch that controls all web security configuration: basic web protection, custom rules, and rate limiting
	WebSwitch *string `json:"WebSwitch,omitempty" name:"WebSwitch"`
}

type Tag struct {
	// Tag key
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
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

type TimingDataItem struct {
	// Second-level timestamp
	// Note: This field may return null, indicating that no valid values can be obtained.
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TimingDataRecord struct {
	// Query dimension value
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// Detailed time series data
	// Note: This field may return null, indicating that no valid values can be obtained.
	TypeValue []*TimingTypeValue `json:"TypeValue,omitempty" name:"TypeValue"`
}

type TimingTypeValue struct {
	// Sum
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sum *int64 `json:"Sum,omitempty" name:"Sum"`

	// Maximum
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// Average
	// Note: This field may return null, indicating that no valid values can be obtained.
	Avg *int64 `json:"Avg,omitempty" name:"Avg"`

	// Metric name
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// This field will be disused soon. Use the `Detail` field instead.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailData []*int64 `json:"DetailData,omitempty" name:"DetailData"`

	// Detailed data
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`
}

type TopDataRecord struct {
	// Query dimension value
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// Top data rankings
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData"`
}

type TopDetailData struct {
	// Field name
	Key *string `json:"Key,omitempty" name:"Key"`

	// Field value
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TopNEntry struct {
	// Entry key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Top N data
	Value []*TopNEntryValue `json:"Value,omitempty" name:"Value"`
}

type TopNEntryValue struct {
	// Entry name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Quantity
	Count *int64 `json:"Count,omitempty" name:"Count"`
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

type WafConfig struct {
	// Switch
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Protection level: `loose`, `normal`, `strict`, `stricter`, `custom`
	Level *string `json:"Level,omitempty" name:"Level"`

	// Mode: `block`, `observe`, `close`
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// WAF rule allowlist/blocklist
	WafRules *WafRule `json:"WafRules,omitempty" name:"WafRules"`

	// AI rule engine
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`
}

type WafRule struct {
	// Blocklist
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitempty" name:"BlockRuleIDs"`

	// Whether the WAF rule is enabled or disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Observe mode
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitempty" name:"ObserveRuleIDs"`
}

type WebAttackEvent struct {
	// Client IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Attack URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackUrl *string `json:"AttackUrl,omitempty" name:"AttackUrl"`

	// Attack time in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackTime *int64 `json:"AttackTime,omitempty" name:"AttackTime"`
}

type WebEventData struct {
	// Data set of attack events
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*WebAttackEvent `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type WebLogData struct {
	// Data
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*WebLogs `json:"List,omitempty" name:"List"`

	// Current page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Number of items per page
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of pages
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type WebLogs struct {
	// Attack content
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// Attack IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// Attack type
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// uuid
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msuuid *string `json:"Msuuid,omitempty" name:"Msuuid"`

	// Request method
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// Request URI
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// Risk grade
	// Note: This field may return null, indicating that no valid values can be obtained.
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// Rule ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// IP country/region
	// Note: This field may return null, indicating that no valid values can be obtained.
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// Event ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Processing method
	// Note: This field may return null, indicating that no valid values can be obtained.
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// http_log
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// user agent
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// Attack time. For consistency considerations, the original parameter `time` was renamed `AttackTime`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`
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

	// Specifies whether to enable CNAME acceleration
	// - `enabled`: Enable
	// - `disabled`: Disable
	// Note: This field may return null, indicating that no valid values can be obtained.
	CnameSpeedUp *string `json:"CnameSpeedUp,omitempty" name:"CnameSpeedUp"`

	// Ownership verification status of the site when it is connected to EdgeOne via CNAME.
	// - `finished`: The site is verified.
	// - `pending`: Verifying the ownership of the site.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

	// Resource tag
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Billable resource
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// Site creation date
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Site modification date
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`


	Area *string `json:"Area,omitempty" name:"Area"`
}

type ZoneFilter struct {
	// Filters by the field name. Vaules:
	// - `name`: Site name.
	// - `status`: Site status.
	// - `tagKey`: Tag key.
	// - `tagValue`: Tag value.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filters by the field value
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Specifies whether to enable fuzzy query. It’s only available when filter name is `name`. If it’s enabled, the length of `Values` must be 1.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}