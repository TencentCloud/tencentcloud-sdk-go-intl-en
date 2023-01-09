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

package v20220901

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccelerateType struct {
	// Acceleration switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type AclCondition struct {
	// The field to match. Values:
	// <li>`host`: Request domain name</li>
	// <li>`sip`: Client IP</li>
	// <li>`ua`: User-Agent</li>
	// <li>`cookie`: Cookie</li>
	// <li>`cgi`: CGI script</li>
	// <li>`xff`: XFF header</li>
	// <li>`url`: Request URL</li>
	// <li>`accept`: Request content type</li>
	// <li>`method`: Request method</li>
	// <li>`header`: Request header</li>
	// <li>`sip_proto`: Network layer protocol</li>
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// The parameter of the field. When `MatchFrom = header`, the key contained in the header can be passed.
	MatchParam *string `json:"MatchParam,omitempty" name:"MatchParam"`

	// The logical operator. Values:
	// <li>`equal`: Value equals</li>
	// <li>`not_equal`: Value not equals</li>
	// <li>`include`: String contains</li>
	// <li>`not_include`: String not contains</li>
	// <li>`match`: IP matches</li>
	// <li>`not_match`: IP not matches</li>
	// <li>`include_area`: Regions contain</li>
	// <li>`is_empty`: Value left empty</li>
	// <li>`not_exists`: Key fields not exist</li>
	// <li>`regexp`: Regex matches</li>
	// <li>`len_gt`: Value greater than</li>
	// <li>`len_lt`: Value smaller than</li>
	// <li>`len_eq`: Value equals</li>
	// <li>`match_prefix`: Prefix matches</li>
	// <li>`match_suffix`: Suffix matches</li>
	// <li>`wildcard`: Wildcard</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// The content to match.
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type AclConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The custom rule.
	AclUserRules []*AclUserRule `json:"AclUserRules,omitempty" name:"AclUserRules"`
}

type AclUserRule struct {
	// The rule name.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// The rule action. Values:
	// <li>`trans`: Allow the request.</li>
	// <li>`drop`: Block the request.</li>
	// <li>`monitor`: Observe the request.</li>
	// <li>`ban`: Block the IP.</li>
	// <li>`redirect`: Redirect the request.</li>
	// <li>`page`: Return the specified page.</li>
	// <li>`alg`: Verify the request by Javascript challenge.</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// The custom rule.
	AclConditions []*AclCondition `json:"AclConditions,omitempty" name:"AclConditions"`

	// The rule priority. Value range: 0-100.
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// The rule ID, which is only used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// The update time, which is only used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The IP blocking duration. Value range: 0 seconds - 2 days. Default value: 0 seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// The unit of the IP blocking duration. Values:
	// <li>`second`: Second</li>
	// <li>`minutes`: Minute</li>
	// <li>`hour`: Hour</li>Default value: second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// The name of the custom page, which defaults to an empty string.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The ID of the custom page, which defaults to 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// The redirection URL, which must be a subdomain name of the site. It defaults to an empty string.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// The response code returned after redirection, which defaults to 0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResponseCode *int64 `json:"ResponseCode,omitempty" name:"ResponseCode"`
}

type Action struct {
	// Common feature operation. Features of this type include:
	// <li>`AccessUrlRedirect`: Access URL rewrite</li>
	// <li>`UpstreamUrlRedirect`: Origin-pull URL rewrite</li>
	// <li>`QUIC`: QUIC</li>
	// <li>`WebSocket`: WebSocket</li>
	// <li>`VideoSeek`: Video dragging</li>
	// <li>`Authentication`: Token authentication</li>
	// <li>`CacheKey`: Custom cache key</li>
	// <li>`Cache`: Node cache TTL</li>
	// <li>`MaxAge`: Browser cache TTL</li>
	// <li>`OfflineCache`: Offline cache</li>
	// <li>`SmartRouting`: Smart acceleration</li>
	// <li>`RangeOriginPull`: Range GETs</li>
	// <li>`UpstreamHttp2`: HTTP/2 forwarding</li>
	// <li>`HostHeader`: Host header rewrite</li>
	// <li>`ForceRedirect`: Force HTTPS</li>
	// <li>`OriginPullProtocol`: Origin-pull HTTPS</li>
	// <li>`CachePrefresh`: Cache prefresh</li>
	// <li>`Compression`: Smart compression</li>
	// <li>`Hsts`</li>
	// <li>`ClientIpHeader`</li>
	// <li>`TlsVersion`</li>
	// <li>`OcspStapling`</li>
	// <li>`Http2`: HTTP/2 access</li>
	// <li>`UpstreamFollowRedirect: Follow origin redirect</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NormalAction *NormalAction `json:"NormalAction,omitempty" name:"NormalAction"`

	// Feature operation with a request/response header. Features of this type include:
	// <li>`RequestHeader`: HTTP request header modification.</li>
	// <li>`ResponseHeader`: HTTP response header modification.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RewriteAction *RewriteAction `json:"RewriteAction,omitempty" name:"RewriteAction"`

	// Feature operation with a status code. Features of this type include:
	// <li>`ErrorPage`: Custom error page.</li>
	// <li>`StatusCodeCache`: Status code cache TTL.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	CodeAction *CodeAction `json:"CodeAction,omitempty" name:"CodeAction"`
}

type AdvancedFilter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value of the filtered field.
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Whether to enable fuzzy query.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type AiRule struct {
	// The status of the AI rule engine. Values:
	// <li>`smart_status_close`: Disabled</li>
	// <li>`smart_status_open`: Block</li>
	// <li>`smart_status_observe`: Observe</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type AliasDomain struct {
	// The alias domain name.
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`

	// Status of the alias domain name. Values:
	// <li>`active`: Activated</li>
	// <li>`pending`: Deploying</li>
	// <li>`conflict`: Reclaimed</li>
	// <li>`stop`: Stopped</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The blocking mode. Values:
	// <li>`0`: Not blocked</li>
	// <li>`11`: Blocked due to regulatory compliance</li>
	// <li>`14`: Blocked due to ICP filing not obtained</li>
	ForbidMode *int64 `json:"ForbidMode,omitempty" name:"ForbidMode"`

	// Creation time of the alias domain name.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Modification time of the alias domain name.
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`
}

type ApplicationProxy struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The domain name or subdomain name when `ProxyType=hostname`.
	// The instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// The scheduling mode. Values:
	// <li>`ip`: Schedule via Anycast IP.</li>
	// <li>`domain`: Schedule via CNAME.</li>
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// Default value: overseas.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether to enable security protection. Values:
	// <li>`0`: Disable security protection.</li>
	// <li>`1`: Enable security protection.</li>
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// Whether to enable acceleration. Values:
	// <li>`0`: Disable acceleration.</li>
	// <li>`1`: Enable acceleration.</li>
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// The session persistence duration.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// The rule status. Values:
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	// <li>`progress`: Deploying</li>
	// <li>`stopping`: Disabling</li>
	// <li>`fail`: Failed to deploy or disable</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The blocking status of the proxy. Values:
	// <li>`banned`: Blocked</li>
	// <li>`banning`: Blocking</li>
	// <li>`recover`: Unblocked</li>
	// <li>`recovering`: Unblocking</li>
	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`

	// Scheduling information.
	ScheduleValue []*string `json:"ScheduleValue,omitempty" name:"ScheduleValue"`

	// When `ProxyType=hostname`:
	// This field indicates the unique ID of the subdomain name.
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// The IPv6 access configuration.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// The update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// List of rules.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`
}

type ApplicationProxyRule struct {
	// The protocol. Values:
	// <li>`TCP`: TCP protocol.</li>
	// <li>`UDP`: UDP protocol.</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	// Note that each rule can have up to 20 ports.
	Port []*string `json:"Port,omitempty" name:"Port"`

	// The origin type. Values:
	// <li>`custom`: Specified origins</li>
	// <li>`origins`: Origin group</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information:
	// <li>When `OriginType=custom`, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or ["test.com"].</li>
	// <li>When `OriginType=origins`, it indicates an origin group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The rule status. Values:
	// <li>`online`: Enabled.</li>
	// <li>`offline`: Disabled.</li>
	// <li>`progress`: Deploying</li>
	// <li>`stopping`: Disabling</li>
	// <li>`fail`: Failed to deploy or disable</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>Default value: OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>Default value: false
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitempty" name:"OriginPort"`
}

type AscriptionInfo struct {

	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// The record type.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// The record value.
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
}

// Predefined struct for user
type BindZoneToPlanRequestParams struct {
	// ID of the site to be bound.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// ID of the target plan.
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

type BindZoneToPlanRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site to be bound.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// ID of the target plan.
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

func (r *BindZoneToPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindZoneToPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindZoneToPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindZoneToPlanResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindZoneToPlanResponse struct {
	*tchttp.BaseResponse
	Response *BindZoneToPlanResponseParams `json:"Response"`
}

func (r *BindZoneToPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindZoneToPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotConfig struct {
	// Whether to enable bot security. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The settings of the bot managed rule. If it is null, the settings that were last configured will be used.
	BotManagedRule *BotManagedRule `json:"BotManagedRule,omitempty" name:"BotManagedRule"`

	// The settings of the client reputation rule. If it is null, the settings that were last configured will be used.
	BotPortraitRule *BotPortraitRule `json:"BotPortraitRule,omitempty" name:"BotPortraitRule"`

	// The bot intelligence settings. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntelligenceRule *IntelligenceRule `json:"IntelligenceRule,omitempty" name:"IntelligenceRule"`
}

type BotManagedRule struct {
	// The rule action. Values:
	// <li>`drop`: Block</li>
	// <li>`trans`: Allow</li>
	// <li>`alg`: JavaScript challenge</li>
	// <li>`monitor`: Observe</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// The rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// The ID of the rule that applies the "Allow" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransManagedIds []*int64 `json:"TransManagedIds,omitempty" name:"TransManagedIds"`

	// The ID of the rule that applies the "JavaScript challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// The ID of the rule that applies the "Managed challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// The ID of the rule that applies the "Observe" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// The ID of the rule that applies the "Block" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`
}

type BotPortraitRule struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// The ID of the rule that applies the "JavaScript challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitempty" name:"AlgManagedIds"`

	// The ID of the rule that applies the "Managed challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CapManagedIds []*int64 `json:"CapManagedIds,omitempty" name:"CapManagedIds"`

	// The ID of the rule that applies the "Observe" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MonManagedIds []*int64 `json:"MonManagedIds,omitempty" name:"MonManagedIds"`

	// The ID of the rule that applies the "Block" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DropManagedIds []*int64 `json:"DropManagedIds,omitempty" name:"DropManagedIds"`
}

type CC struct {
	// WAF switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// ID of the policy
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type Cache struct {
	// Whether to enable cache configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Cache expiration time setting.
	// Unit: second. The maximum value is 365 days.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`

	// Whether to enable force cache. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`
}

type CacheConfig struct {
	// Cache configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// No-cache configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoCache *NoCache `json:"NoCache,omitempty" name:"NoCache"`

	// Follows the origin server configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	FollowOrigin *FollowOrigin `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type CacheKey struct {
	// Whether to enable full-path cache. Values:
	// <li>`on`: Enable full-path cache (i.e., disable Ignore Query String).</li>
	// <li>`off`: Disable full-path cache (i.e., enable Ignore Query String).</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`

	// Whether to ignore case in the cache key. Values:
	// <li>`on`: Ignore</li>
	// <li>`off`: Not ignore</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IgnoreCase *string `json:"IgnoreCase,omitempty" name:"IgnoreCase"`

	// Request parameter contained in `CacheKey`
	// Note: This field may return null, indicating that no valid values can be obtained.
	QueryString *QueryString `json:"QueryString,omitempty" name:"QueryString"`
}

type CachePrefresh struct {
	// Whether to enable cache prefresh. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The cache prefresh percentage. Values: 1-99
	// Note: This field may return null, indicating that no valid values can be obtained.
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`
}

// Predefined struct for user
type CheckCertificateRequestParams struct {
	// Content of the certificate.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Content of the private key.
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type CheckCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Content of the certificate.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Content of the private key.
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

type ClientIpCountry struct {
	// Whether to enable configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Header name of ClientIpCountry. This field is valid only when `Switch=on`.
	// If it is left empty, the default value `EO-Client-IPCountry` will be used.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
}

type ClientIpHeader struct {
	// Whether to enable the configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The name of the HTTP header that contains the client IP, which is used for forwarding.
	// If this field is not specified, the default value `X-Forwarded-IP` will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
}

type ClientRule struct {
	// The client IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// The rule type.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// The rule ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// The rule description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The blocking status. Values:
	// <li>`block`: Block;</li>
	// <li>`allow`: Allow.</li>
	IpStatus *string `json:"IpStatus,omitempty" name:"IpStatus"`

	// The blocking time recorded in UNIX timestamp.
	BlockTime *int64 `json:"BlockTime,omitempty" name:"BlockTime"`

	// The data entry ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ClsLogTopicInfo struct {
	// Name of the task.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Name of the site.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// ID of the log topic.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Type of the task.
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// Retention period of the log topic.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Whether the log topic is enabled.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Whether the log topic is deleted.
	Deleted *string `json:"Deleted,omitempty" name:"Deleted"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Target location. Values:
	// <li>`cls`: Ship logs to CLS;</li>
	// <li>`custom_enpoint`: Ship logs to a custom address.</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// Region of the logset.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// ID of the site.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland;</li>
	// <li>`overseas`: Global (outside the Chinese mainland).</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// Type of the shipping task. Values:
	// <li>`cls`: Ship logs to CLS.</li>
	// <li>`custom_endpoint`: Ship logs to custom APIs.</li>
	LogSetType *string `json:"LogSetType,omitempty" name:"LogSetType"`
}

type CodeAction struct {
	// Feature name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) API to view the requirements for entering the feature name.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Operation parameter.
	Parameters []*RuleCodeActionParams `json:"Parameters,omitempty" name:"Parameters"`
}

type Compression struct {
	// Whether to enable smart compression. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Compression algorithm. Values:
	// <li>`brotli`: Brotli algorithm</li>
	// <li>`gzip`: Gzip algorithm</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Algorithms []*string `json:"Algorithms,omitempty" name:"Algorithms"`
}

// Predefined struct for user
type CreateAliasDomainRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: Off</li>
	// <li>`hosting`: Managed SSL certificate</li>
	// <li>`apply`: Free certificate</li>Default value: none
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitempty" name:"CertId"`
}

type CreateAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: Off</li>
	// <li>`hosting`: Managed SSL certificate</li>
	// <li>`apply`: Free certificate</li>Default value: none
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitempty" name:"CertId"`
}

func (r *CreateAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasName")
	delete(f, "TargetName")
	delete(f, "CertType")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAliasDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAliasDomainResponseParams `json:"Response"`
}

func (r *CreateAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// When `ProxyType=hostname`, this field indicates a domain name or subdomain name.
	// When `ProxyType=instance`, it indicates a proxy instance.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// The scheduling mode. Values:
	// <li>`ip`: Schedule via Anycast IP.</li>
	// <li>`domain`: Schedule via CNAME.</li>
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// Whether to enable security protection. Values:
	// <li>`0`: Disable security protection.</li>
	// <li>`1`: Enable security protection.</li>
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// Whether to enable acceleration. Values:
	// <li>`0`: Disable acceleration.</li>
	// <li>`1`: Enable acceleration.</li>
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>If not specified, this field uses the default value `instance`.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// If not specified, this field uses the default value 600.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// The IPv6 access configuration.
	// If this field is not specified, IPv6 access will be disabled.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// The rule details.
	// If this field is not specified, an application proxy rule will not be created.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// When `ProxyType=hostname`, this field indicates a domain name or subdomain name.
	// When `ProxyType=instance`, it indicates a proxy instance.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// The scheduling mode. Values:
	// <li>`ip`: Schedule via Anycast IP.</li>
	// <li>`domain`: Schedule via CNAME.</li>
	PlatType *string `json:"PlatType,omitempty" name:"PlatType"`

	// Whether to enable security protection. Values:
	// <li>`0`: Disable security protection.</li>
	// <li>`1`: Enable security protection.</li>
	SecurityType *int64 `json:"SecurityType,omitempty" name:"SecurityType"`

	// Whether to enable acceleration. Values:
	// <li>`0`: Disable acceleration.</li>
	// <li>`1`: Enable acceleration.</li>
	AccelerateType *int64 `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>If not specified, this field uses the default value `instance`.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// If not specified, this field uses the default value 600.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// The IPv6 access configuration.
	// If this field is not specified, IPv6 access will be disabled.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// The rule details.
	// If this field is not specified, an application proxy rule will not be created.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`
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
	delete(f, "ProxyName")
	delete(f, "PlatType")
	delete(f, "SecurityType")
	delete(f, "AccelerateType")
	delete(f, "ProxyType")
	delete(f, "SessionPersistTime")
	delete(f, "Ipv6")
	delete(f, "ApplicationProxyRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyResponseParams struct {
	// The L4 application proxy ID.
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitempty" name:"Port"`

	// The origin type. Values:
	// <li>`custom`: Specified origins</li>
	// <li>`origins`: Origin group</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information:
	// <li>When `OriginType=custom`, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or ["test.com"].</li>
	// <li>When `OriginType=origins`, it indicates an origin group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>Default value: OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable.</li>
	// <li>`false`: Disable.</li>Default value: false.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitempty" name:"OriginPort"`
}

type CreateApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitempty" name:"Port"`

	// The origin type. Values:
	// <li>`custom`: Specified origins</li>
	// <li>`origins`: Origin group</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Origin server information:
	// <li>When `OriginType=custom`, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or ["test.com"].</li>
	// <li>When `OriginType=origins`, it indicates an origin group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>Default value: OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable.</li>
	// <li>`false`: Disable.</li>Default value: false.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitempty" name:"OriginPort"`
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
	delete(f, "OriginPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRuleResponseParams struct {
	// The rule ID.
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
type CreateCredentialRequestParams struct {

}

type CreateCredentialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCredentialResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateCredentialResponseParams `json:"Response"`
}

func (r *CreateCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The origin type. Values:
	// <li>`self`: Customer origin</li>
	// <li>`third_party`: Third-party origin</li>
	// <li>`cos`: Tencent Cloud COS origin</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The name of the origin group.
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// The origin configuration type when `OriginType=self`. Values:
	// <li>`area`: Configure by region.</li>
	// <li>`weight`: Configure by weight.</li>
	// <li>`proto`: Configure by HTTP protocol.</li>When `OriginType=third_party/cos`, leave this field empty.
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// Details of the origin record.
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`

	// The origin domain. This field can be specified only when `OriginType=self`.
	HostHeader *string `json:"HostHeader,omitempty" name:"HostHeader"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The origin type. Values:
	// <li>`self`: Customer origin</li>
	// <li>`third_party`: Third-party origin</li>
	// <li>`cos`: Tencent Cloud COS origin</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The name of the origin group.
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// The origin configuration type when `OriginType=self`. Values:
	// <li>`area`: Configure by region.</li>
	// <li>`weight`: Configure by weight.</li>
	// <li>`proto`: Configure by HTTP protocol.</li>When `OriginType=third_party/cos`, leave this field empty.
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// Details of the origin record.
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`

	// The origin domain. This field can be specified only when `OriginType=self`.
	HostHeader *string `json:"HostHeader,omitempty" name:"HostHeader"`
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
	delete(f, "ZoneId")
	delete(f, "OriginType")
	delete(f, "OriginGroupName")
	delete(f, "ConfigurationType")
	delete(f, "OriginRecords")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupResponseParams struct {
	// The ID of the origin group.
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

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
type CreatePlanForZoneRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The plan option. Values:
	// <li>`sta`: Standard plan that supports content delivery network outside the Chinese mainland.</li>
	// <li>`sta_with_bot`: Standard plan that supports content delivery network outside the Chinese mainland and bot management.</li>
	// <li>`sta_cm`: Standard plan that supports content delivery network inside the Chinese mainland.</li>
	// <li>`sta_cm_with_bot`: Standard plan that supports content delivery network inside the Chinese mainland and bot management.</li>
	// <li>`sta_global`: Standard plan that supports content delivery network over the globe.</li>
	// <li>`sta_global_with_bot`: Standard plan that supports content delivery network over the globe and bot management.</li>
	// <li>`ent`: Enterprise plan that supports content delivery network outside the Chinese mainland.</li>
	// <li>`ent_with_bot`: Enterprise plan that supports content delivery network outside the Chinese mainland and bot management.</li>
	// <li>`ent_cm`: Enterprise plan that supports content delivery network inside the Chinese mainland.</li>
	// <li>`ent_cm_with_bot`: Enterprise plan that supports content delivery network inside the Chinese mainland and bot management.</li>
	// <li>`ent_global`: Enterprise plan that supports content delivery network over the globe.</li>
	// <li>`ent_global_with_bot`: Enterprise plan that supports content delivery network over the globe and bot management.</li>To get the available plan options for your account, view the output from <a href="https://tcloud4api.woa.com/document/product/1657/80124?!preview&!document=1">DescribeAvailablePlans</a>.
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
}

type CreatePlanForZoneRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The plan option. Values:
	// <li>`sta`: Standard plan that supports content delivery network outside the Chinese mainland.</li>
	// <li>`sta_with_bot`: Standard plan that supports content delivery network outside the Chinese mainland and bot management.</li>
	// <li>`sta_cm`: Standard plan that supports content delivery network inside the Chinese mainland.</li>
	// <li>`sta_cm_with_bot`: Standard plan that supports content delivery network inside the Chinese mainland and bot management.</li>
	// <li>`sta_global`: Standard plan that supports content delivery network over the globe.</li>
	// <li>`sta_global_with_bot`: Standard plan that supports content delivery network over the globe and bot management.</li>
	// <li>`ent`: Enterprise plan that supports content delivery network outside the Chinese mainland.</li>
	// <li>`ent_with_bot`: Enterprise plan that supports content delivery network outside the Chinese mainland and bot management.</li>
	// <li>`ent_cm`: Enterprise plan that supports content delivery network inside the Chinese mainland.</li>
	// <li>`ent_cm_with_bot`: Enterprise plan that supports content delivery network inside the Chinese mainland and bot management.</li>
	// <li>`ent_global`: Enterprise plan that supports content delivery network over the globe.</li>
	// <li>`ent_global_with_bot`: Enterprise plan that supports content delivery network over the globe and bot management.</li>To get the available plan options for your account, view the output from <a href="https://tcloud4api.woa.com/document/product/1657/80124?!preview&!document=1">DescribeAvailablePlans</a>.
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
}

func (r *CreatePlanForZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanForZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PlanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePlanForZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanForZoneResponseParams struct {
	// List of purchased resources.
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`

	// List or order numbers.
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePlanForZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreatePlanForZoneResponseParams `json:"Response"`
}

func (r *CreatePlanForZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanForZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of resources to be pre-warmed, for example:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Whether to encode a URL according to RFC3986. Enable this field when the URL contains non-ASCII characters.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// HTTP header information
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of resources to be pre-warmed, for example:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Whether to encode a URL according to RFC3986. Enable this field when the URL contains non-ASCII characters.
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
	// ID of the task.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// List of failed tasks.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Mode of cache purging. Values:
	// <li>`purge_url`: Purge by URL</li>
	// <li>`purge_prefix`: Purge by prefix</li>
	// <li>`purge_host`: Purge by hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by cache tag</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target resource to be purged, which depends on the `Type` field.
	// 1. When `Type = purge_host`:
	// Enter the hostname, such as www.example.com and foo.bar.example.com.
	// 2. When `Type = purge_prefix`:
	// Enter the prefix, such as http://www.example.com/example.
	// 3. When `Type = purge_url`:
	// Enter the URL, such as https://www.example.com/example.jpg.
	// 4. When `Type = purge_all`:
	// This field can be left empty.
	// 5. When `Type = purge_cache_tag`:
	// Enter the cache tag, such as tag1.
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Mode of cache purging. Values:
	// <li>`purge_url`: Purge by URL</li>
	// <li>`purge_prefix`: Purge by prefix</li>
	// <li>`purge_host`: Purge by hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by cache tag</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target resource to be purged, which depends on the `Type` field.
	// 1. When `Type = purge_host`:
	// Enter the hostname, such as www.example.com and foo.bar.example.com.
	// 2. When `Type = purge_prefix`:
	// Enter the prefix, such as http://www.example.com/example.
	// 3. When `Type = purge_url`:
	// Enter the URL, such as https://www.example.com/example.jpg.
	// 4. When `Type = purge_all`:
	// This field can be left empty.
	// 5. When `Type = purge_cache_tag`:
	// Enter the cache tag, such as tag1.
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
	// ID of the task.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// List of failed tasks and reasons.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
type CreateReplayTaskRequestParams struct {
	// List of replay task IDs.
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

type CreateReplayTaskRequest struct {
	*tchttp.BaseRequest
	
	// List of replay task IDs.
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *CreateReplayTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplayTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReplayTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplayTaskResponseParams struct {
	// ID of the task.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// List of failed tasks and reasons.
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateReplayTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateReplayTaskResponseParams `json:"Response"`
}

func (r *CreateReplayTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplayTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The rule name (1 to 255 characters)
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The rule name (1 to 255 characters)
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "ZoneId")
	delete(f, "RuleName")
	delete(f, "Status")
	delete(f, "Rules")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateSpeedTestingRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The subdomain name to test.
	Host *string `json:"Host,omitempty" name:"Host"`
}

type CreateSpeedTestingRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The subdomain name to test.
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *CreateSpeedTestingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpeedTestingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSpeedTestingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpeedTestingResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSpeedTestingResponse struct {
	*tchttp.BaseResponse
	Response *CreateSpeedTestingResponseParams `json:"Response"`
}

func (r *CreateSpeedTestingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpeedTestingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneRequestParams struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The access mode. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record. Note that you should verify your site with the IdentifyZone API before starting site access.</li>If it is left empty, the default value `full` is used.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Whether to skip scanning the existing DNS records of the site. Default value: false.
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`

	// The resource tag.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to allow duplicate sites. Values:
	// <li>`true`: Duplicate sites are allowed.</li>
	// <li>`false`: Duplicate sites are not allowed.</li>If it is left empty, the default value `false` is used.
	AllowDuplicates *bool `json:"AllowDuplicates,omitempty" name:"AllowDuplicates"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitempty" name:"AliasZoneName"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The access mode. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record. Note that you should verify your site with the IdentifyZone API before starting site access.</li>If it is left empty, the default value `full` is used.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Whether to skip scanning the existing DNS records of the site. Default value: false.
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`

	// The resource tag.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to allow duplicate sites. Values:
	// <li>`true`: Duplicate sites are allowed.</li>
	// <li>`false`: Duplicate sites are not allowed.</li>If it is left empty, the default value `false` is used.
	AllowDuplicates *bool `json:"AllowDuplicates,omitempty" name:"AllowDuplicates"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitempty" name:"AliasZoneName"`
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
	delete(f, "ZoneName")
	delete(f, "Type")
	delete(f, "JumpStart")
	delete(f, "Tags")
	delete(f, "AllowDuplicates")
	delete(f, "AliasZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneResponseParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

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

type DDoS struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type DefaultServerCertInfo struct {
	// ID of the server certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Alias of the certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Type of the certificate. Values:
	// <li>`default`: Default certificate;</li>
	// <li>`upload`: Custom certificate;</li>
	// <li>`managed`: Tencent Cloud-managed certificate.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time when the certificate expires.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Time when the certificate takes effect.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// Common name of the certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CommonName *string `json:"CommonName,omitempty" name:"CommonName"`

	// Domain names added to the SAN certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// Deployment status. Values:
	// <li>`processing`: Deployment in progress</li>
	// <li>`deployed`: Deployed</li>
	// <li>`failed`: Deployment failed</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Failure description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Certificate algorithm.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SignAlgo *string `json:"SignAlgo,omitempty" name:"SignAlgo"`
}

// Predefined struct for user
type DeleteAliasDomainRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The alias domain name to be deleted. If it is left empty, the delete operation is not performed.
	AliasNames []*string `json:"AliasNames,omitempty" name:"AliasNames"`
}

type DeleteAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The alias domain name to be deleted. If it is left empty, the delete operation is not performed.
	AliasNames []*string `json:"AliasNames,omitempty" name:"AliasNames"`
}

func (r *DeleteAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAliasDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAliasDomainResponseParams `json:"Response"`
}

func (r *DeleteAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
}

type DeleteApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The rule ID.
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
type DeleteOriginGroupRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The ID of the origin group.
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`
}

type DeleteOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The ID of the origin group.
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`
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
	delete(f, "ZoneId")
	delete(f, "OriginGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupResponseParams struct {
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
type DeleteRulesRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IDs of the rules to be deleted.
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

type DeleteRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IDs of the rules to be deleted.
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *DeleteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRulesResponseParams `json:"Response"`
}

func (r *DeleteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneResponseParams struct {
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
type DescribeAddableEntityListRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Type of the shipping entity. Values:
	// <li>`domain`: L7 acceleration logs;</li>
	// <li>`application`: L4 acceleration logs;</li>
	// <li>`web-rateLiming`: Rate limiting logs;</li>
	// <li>`web-attack`: Web security logs;</li>
	// <li>`web-rule`: Custom rule logs;</li>
	// <li>`web-bot`: Bot management logs.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
}

type DescribeAddableEntityListRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Type of the shipping entity. Values:
	// <li>`domain`: L7 acceleration logs;</li>
	// <li>`application`: L4 acceleration logs;</li>
	// <li>`web-rateLiming`: Rate limiting logs;</li>
	// <li>`web-attack`: Web security logs;</li>
	// <li>`web-rule`: Custom rule logs;</li>
	// <li>`web-bot`: Bot management logs.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`
}

func (r *DescribeAddableEntityListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddableEntityListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "EntityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddableEntityListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddableEntityListResponseParams struct {
	// Total number of query results.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of available shipping entities.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAddableEntityListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddableEntityListResponseParams `json:"Response"`
}

func (r *DescribeAddableEntityListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddableEntityListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasDomainsRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`target-name`:<br>   Filter by <strong>target domain name</strong><br>   Type: String<br>   Required: No</li><li>`alias-name`:<br>   Filter by <strong>alias domain name</strong><br>   Type: String<br>   Required: No</li>Only `alias-name` supports fuzzy query.
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAliasDomainsRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`target-name`:<br>   Filter by <strong>target domain name</strong><br>   Type: String<br>   Required: No</li><li>`alias-name`:<br>   Filter by <strong>alias domain name</strong><br>   Type: String<br>   Required: No</li>Only `alias-name` supports fuzzy query.
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAliasDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAliasDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasDomainsResponseParams struct {
	// Total eligible alias domain names.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Information of the eligible alias domain names.
	AliasDomains []*AliasDomain `json:"AliasDomains,omitempty" name:"AliasDomains"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAliasDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAliasDomainsResponseParams `json:"Response"`
}

func (r *DescribeAliasDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesRequestParams struct {
	// The paginated query offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries. <li>`proxy-id`:<br>   Filter by <strong>proxy ID</strong>, such as proxy-ev2sawbwfd<br>   Type: String<br>   Required: No</li><li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-vawer2vadg<br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApplicationProxiesRequest struct {
	*tchttp.BaseRequest
	
	// The paginated query offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries. <li>`proxy-id`:<br>   Filter by <strong>proxy ID</strong>, such as proxy-ev2sawbwfd<br>   Type: String<br>   Required: No</li><li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-vawer2vadg<br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApplicationProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesResponseParams struct {
	// List of application proxies.
	ApplicationProxies []*ApplicationProxy `json:"ApplicationProxies,omitempty" name:"ApplicationProxies"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationProxiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProxiesResponseParams `json:"Response"`
}

func (r *DescribeApplicationProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailablePlansRequestParams struct {

}

type DescribeAvailablePlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAvailablePlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailablePlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailablePlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailablePlansResponseParams struct {
	// Plans available for the current user
	// Note: This field may return null, indicating that no valid values can be obtained.
	PlanInfo []*PlanInfo `json:"PlanInfo,omitempty" name:"PlanInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAvailablePlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailablePlansResponseParams `json:"Response"`
}

func (r *DescribeAvailablePlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailablePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientRuleListRequestParams struct {
	// The ID of the site to be queried.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The subdomain name to be queried.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Rule type. Values:
	// <li>`acl`: Custom rules;</li>
	// <li>`rate`: Rate limiting rules.</li>All rules will be queried if this field is not specified.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// The rule ID.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// The client IP.
	SourceClientIp *string `json:"SourceClientIp,omitempty" name:"SourceClientIp"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeClientRuleListRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the site to be queried.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The subdomain name to be queried.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Rule type. Values:
	// <li>`acl`: Custom rules;</li>
	// <li>`rate`: Rate limiting rules.</li>All rules will be queried if this field is not specified.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// The rule ID.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// The client IP.
	SourceClientIp *string `json:"SourceClientIp,omitempty" name:"SourceClientIp"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeClientRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Domain")
	delete(f, "RuleType")
	delete(f, "RuleId")
	delete(f, "SourceClientIp")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientRuleListResponseParams struct {
	// The blocked client information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*ClientRule `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClientRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientRuleListResponseParams `json:"Response"`
}

func (r *DescribeClientRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeContentQuotaRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeContentQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaResponseParams struct {
	// Purging quotas.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PurgeQuota []*Quota `json:"PurgeQuota,omitempty" name:"PurgeQuota"`

	// Pre-warming quotas.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrefetchQuota []*Quota `json:"PrefetchQuota,omitempty" name:"PrefetchQuota"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContentQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContentQuotaResponseParams `json:"Response"`
}

func (r *DescribeContentQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackTopDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The statistical metric. Values:
	// <li>`ddos_attackFlux_protocol`: Rank protocols by the attack traffic.</li>
	// <li>`ddos_attackPackageNum_protocol`: Rank protocols by the number of attack packets.</li>
	// <li>`ddos_attackNum_attackType`: Rank attack types by the number of attacks.</li>
	// <li>`ddos_attackNum_sregion`: Rank attacker regions by the number of attacks.</li>
	// <li>`ddos_attackFlux_sip`: Rank attacker IPs by the number of attacks.</li>
	// <li>`ddos_attackFlux_sregion`: Rank attacker regions by the number of attacks.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// List of site IDs to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The attack type. Values:
	// <li>`flood`: Flood;</li>
	// <li>`icmpFlood`: ICMP flood;</li>
	// <li>`all`: All attack types.</li>This field will be set to the default value `all` if not specified.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// The port number.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Queries the top n rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackTopDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The statistical metric. Values:
	// <li>`ddos_attackFlux_protocol`: Rank protocols by the attack traffic.</li>
	// <li>`ddos_attackPackageNum_protocol`: Rank protocols by the number of attack packets.</li>
	// <li>`ddos_attackNum_attackType`: Rank attack types by the number of attacks.</li>
	// <li>`ddos_attackNum_sregion`: Rank attacker regions by the number of attacks.</li>
	// <li>`ddos_attackFlux_sip`: Rank attacker IPs by the number of attacks.</li>
	// <li>`ddos_attackFlux_sregion`: Rank attacker regions by the number of attacks.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// List of site IDs to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The attack type. Values:
	// <li>`flood`: Flood;</li>
	// <li>`icmpFlood`: ICMP flood;</li>
	// <li>`all`: All attack types.</li>This field will be set to the default value `all` if not specified.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// The port number.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Queries the top n rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "AttackType")
	delete(f, "ProtocolType")
	delete(f, "Port")
	delete(f, "Limit")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackTopDataResponseParams struct {
	// The list of top-ranked DDoS attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopEntry `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackTopDataResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesRequestParams struct {
	// Filter criteria. Each filter criteria can have up to 5 entries.
	// <li>`zone-id`: <br>Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: `0`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// Filter criteria. Each filter criteria can have up to 5 entries.
	// <li>`zone-id`: <br>Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: `0`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
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
	DefaultServerCertInfo []*DefaultServerCertInfo `json:"DefaultServerCertInfo,omitempty" name:"DefaultServerCertInfo"`

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
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone`:<br>   Filter by <strong>site name</strong>, such as tencent.com (up to one entry)<br>   Type: String<br>   Required: No
	// <li>`host`:<br>   Filter by <strong>domain name</strong>, such as test.tencent.com (up to one entry)<br>   Type: String<br>   Required: No
	// <li>`type`:<br>   Filter by <strong>DNS record type</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `A`: A record<br>   `AAAA`: AAAA record<br>   `CNAME`: CNAME record<br>   `MX`: MX record<br>   `TXT`: TXT record<br>   `NS`: NS record<br>   `SRV`: SRV record<br>   `CAA`: CAA record
	// <li>`code`:<br>   Filter by <strong>DNS status code</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `NoError`: Success<br>   `NXDomain`: Not found the request domain<br>   `NotImp`: Not supported request type<br>   `Refused`: The domain name server refuses to execute the request for policy reasons
	// <li>`area`:<br>   Filter by <strong>DNS region</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `Asia`<br>   `Europe`<br>   `Africa`<br>   `Oceania`<br>   `Americas`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>This field will be set to the default value `min` if not specified.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type DescribeDnsDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone`:<br>   Filter by <strong>site name</strong>, such as tencent.com (up to one entry)<br>   Type: String<br>   Required: No
	// <li>`host`:<br>   Filter by <strong>domain name</strong>, such as test.tencent.com (up to one entry)<br>   Type: String<br>   Required: No
	// <li>`type`:<br>   Filter by <strong>DNS record type</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `A`: A record<br>   `AAAA`: AAAA record<br>   `CNAME`: CNAME record<br>   `MX`: MX record<br>   `TXT`: TXT record<br>   `NS`: NS record<br>   `SRV`: SRV record<br>   `CAA`: CAA record
	// <li>`code`:<br>   Filter by <strong>DNS status code</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `NoError`: Success<br>   `NXDomain`: Not found the request domain<br>   `NotImp`: Not supported request type<br>   `Refused`: The domain name server refuses to execute the request for policy reasons
	// <li>`area`:<br>   Filter by <strong>DNS region</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `Asia`<br>   `Europe`<br>   `Africa`<br>   `Oceania`<br>   `Americas`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>This field will be set to the default value `min` if not specified.
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
	// DNS statistics.
	Data []*DnsData `json:"Data,omitempty" name:"Data"`

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
type DescribeHostsSettingRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`host`:<br>   Filter by <strong>domain name </strong><br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeHostsSettingRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`host`:<br>   Filter by <strong>domain name </strong><br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingResponseParams struct {
	// List of domain names.
	DetailHosts []*DetailHost `json:"DetailHosts,omitempty" name:"DetailHosts"`

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
type DescribeIdentificationsRequestParams struct {
	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-name`: <br>Filter by <strong>site name</strong><br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeIdentificationsRequest struct {
	*tchttp.BaseRequest
	
	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-name`: <br>Filter by <strong>site name</strong><br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIdentificationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdentificationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationsResponseParams struct {
	// Number of eligible sites.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The site verification information.
	Identifications []*Identification `json:"Identifications,omitempty" name:"Identifications"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIdentificationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdentificationsResponseParams `json:"Response"`
}

func (r *DescribeIdentificationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSetsRequestParams struct {
	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// Name of the logset.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`
}

type DescribeLogSetsRequest struct {
	*tchttp.BaseRequest
	
	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// Name of the logset.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`
}

func (r *DescribeLogSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogSetRegion")
	delete(f, "LogSetId")
	delete(f, "LogSetName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogSetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSetsResponseParams struct {
	// List of logsets.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogSetList []*LogSetInfo `json:"LogSetList,omitempty" name:"LogSetList"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogSetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogSetsResponseParams `json:"Response"`
}

func (r *DescribeLogSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTopicTasksRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeLogTopicTasksRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLogTopicTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogTopicTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTopicTasksResponseParams struct {
	// List of shipping tasks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicList []*ClsLogTopicInfo `json:"TopicList,omitempty" name:"TopicList"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogTopicTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogTopicTasksResponseParams `json:"Response"`
}

func (r *DescribeLogTopicTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupRequestParams struct {
	// Offset for paginated queries. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Value range: 1-1000. Default value: 10.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-id`<br>   Filter by <strong>site ID</strong>, such as zone-20hzkd4rdmy0<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`origin-group-id`:<br>   Filter by <strong>origin group ID</strong>, such as origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`origin-group-name`:<br>   Filter by <strong>origin group name</strong><br>   Type: String<br>   Required: No<br>   Fuzzy query: Supported (only one origin group name allowed in a query)
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Offset for paginated queries. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Value range: 1-1000. Default value: 10.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-id`<br>   Filter by <strong>site ID</strong>, such as zone-20hzkd4rdmy0<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`origin-group-id`:<br>   Filter by <strong>origin group ID</strong>, such as origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`origin-group-name`:<br>   Filter by <strong>origin group name</strong><br>   Type: String<br>   Required: No<br>   Fuzzy query: Supported (only one origin group name allowed in a query)
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupResponseParams struct {
	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Origin group information.
	OriginGroups []*OriginGroup `json:"OriginGroups,omitempty" name:"OriginGroups"`

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
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Flow_outFlux`: Access traffic;</li>
	// <li>`l7Flow_request`: Access requests;</li>
	// <li>`l7Flow_outBandwidth`: Access bandwidth.</li>
	// <li>`l7Flow_hit_outFlux`: Cache hit traffic.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`:   All protocols. </li>If it's not specified, `all` is used. This parameter is not yet available now.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: One day</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>tagKey<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>tagValue<br>   Filter by the specified <strong>tag value</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeOverviewL7DataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Flow_outFlux`: Access traffic;</li>
	// <li>`l7Flow_request`: Access requests;</li>
	// <li>`l7Flow_outBandwidth`: Access bandwidth.</li>
	// <li>`l7Flow_hit_outFlux`: Cache hit traffic.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`:   All protocols. </li>If it's not specified, `all` is used. This parameter is not yet available now.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: One day</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>tagKey<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>tagValue<br>   Filter by the specified <strong>tag value</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Protocol")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewL7DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataResponseParams struct {
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of L7 traffic summary statistics recorded over time.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	// Start time of the query.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`job-id`:<br>   Filter by <strong>task ID</strong>, such as 1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`target`:<br>   Filter by <strong>target resource</strong>, such as http://www.qq.com/1.txt (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`domains`:<br>   Filter by <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`statuses`:<br>   Filter by <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: The task is in progress.<br>   `success`: The task succeeded.<br>   `failed`: The task failed.<br>   `timeout`: The task timed out.</li>
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`job-id`:<br>   Filter by <strong>task ID</strong>, such as 1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`target`:<br>   Filter by <strong>target resource</strong>, such as http://www.qq.com/1.txt (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`domains`:<br>   Filter by <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`statuses`:<br>   Filter by <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: The task is in progress.<br>   `success`: The task succeeded.<br>   `failed`: The task failed.<br>   `timeout`: The task timed out.</li>
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksResponseParams struct {
	// Total entries that match the specified query condition.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of tasks returned.
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
	// Disused. Use "zone-id" in "Filters" instead.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Start time of the query.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries. <li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-xxx (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`job-id`:<br>   Filter by <strong>task ID</strong>, such as 1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`target`:<br>   Filter by <strong>target resource</strong>, such as http://www.qq.com/1.txt and tag1<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`domains`:<br>   Filter by <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`statuses`:<br>   Filter by <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: The task is in progress.<br>   `success`: The task succeeded.<br>   `failed`: The task failed.<br>   `timeout`: The task timed out.<li>`type`:<br>   Filter by <strong>purging mode</strong> (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `purge_url`: Purge by URL.<br>   `purge_prefix`: Purge by prefix.<br>   `purge_all`: Purge all caches.<br>   `purge_host`: Purge by hostname.<br>   `purge_cache_tag`: Purge by cache tag.</li>
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// Disused. Use "zone-id" in "Filters" instead.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Start time of the query.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries. <li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-xxx (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`job-id`:<br>   Filter by <strong>task ID</strong>, such as 1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`target`:<br>   Filter by <strong>target resource</strong>, such as http://www.qq.com/1.txt and tag1<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`domains`:<br>   Filter by <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`statuses`:<br>   Filter by <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: The task is in progress.<br>   `success`: The task succeeded.<br>   `failed`: The task failed.<br>   `timeout`: The task timed out.<li>`type`:<br>   Filter by <strong>purging mode</strong> (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `purge_url`: Purge by URL.<br>   `purge_prefix`: Purge by prefix.<br>   `purge_all`: Purge all caches.<br>   `purge_host`: Purge by hostname.<br>   `purge_cache_tag`: Purge by cache tag.</li>
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "ZoneId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// Total entries that match the specified query condition.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of tasks returned.
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
type DescribeRulesRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`rule-id`:<br>   Filter by the <strong>rule ID</strong><br>   Type: string<br>   Required: No
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`rule-id`:<br>   Filter by the <strong>rule ID</strong><br>   Type: string<br>   Required: No
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesResponseParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of rules. Rules are sorted in order of execution.
	RuleItems []*RuleItem `json:"RuleItems,omitempty" name:"RuleItems"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesResponseParams `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesSettingRequestParams struct {

}

type DescribeRulesSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRulesSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesSettingResponseParams struct {
	// List of the settings of the rule engine that can be used for request match and their detailed recommended configuration information.
	Actions []*RulesSettingAction `json:"Actions,omitempty" name:"Actions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRulesSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesSettingResponseParams `json:"Response"`
}

func (r *DescribeRulesSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSingleL7AnalysisDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Flow_singleIpRequest`: Number of requests from a single IP.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`domain`<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0<br>   `HTTP/1.1`: HTTP 1.1<br>   `HTTP/2.0`: HTTP 2.0<br>   `HTTP/3.0`: HTTP 3.0<br>   `WebSocket`: WebSocket</li>
	// <li>tagKey<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>tagValue<br>   Filter by the specified <strong>tag value</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeSingleL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Flow_singleIpRequest`: Number of requests from a single IP.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`domain`<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0<br>   `HTTP/1.1`: HTTP 1.1<br>   `HTTP/2.0`: HTTP 2.0<br>   `HTTP/3.0`: HTTP 3.0<br>   `WebSocket`: WebSocket</li>
	// <li>tagKey<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>tagValue<br>   Filter by the specified <strong>tag value</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeSingleL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSingleL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSingleL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSingleL7AnalysisDataResponseParams struct {
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of L7 dimensional data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SingleDataRecord `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSingleL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSingleL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeSingleL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSingleL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingDetailsRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSpeedTestingDetailsRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSpeedTestingDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpeedTestingDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingDetailsResponseParams struct {
	// The site’s load speed across regions.
	SpeedTestingDetailData *SpeedTestingDetailData `json:"SpeedTestingDetailData,omitempty" name:"SpeedTestingDetailData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpeedTestingDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpeedTestingDetailsResponseParams `json:"Response"`
}

func (r *DescribeSpeedTestingDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingMetricDataRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSpeedTestingMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSpeedTestingMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpeedTestingMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingMetricDataResponseParams struct {
	// The site test metrics.
	SpeedTestingMetricData *SpeedTestingMetricData `json:"SpeedTestingMetricData,omitempty" name:"SpeedTestingMetricData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpeedTestingMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpeedTestingMetricDataResponseParams `json:"Response"`
}

func (r *DescribeSpeedTestingMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingQuotaRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeSpeedTestingQuotaRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeSpeedTestingQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpeedTestingQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpeedTestingQuotaResponseParams struct {
	// The quota limit on site tests.
	SpeedTestingQuota *SpeedTestingQuota `json:"SpeedTestingQuota,omitempty" name:"SpeedTestingQuota"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpeedTestingQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpeedTestingQuotaResponseParams `json:"Response"`
}

func (r *DescribeSpeedTestingQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpeedTestingQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metric to query. Values:
	// <li>`l4Flow_connections`: Access connections;</li>
	// <li>`l4Flow_flux`: Access traffic;</li>
	// <li>`l4Flow_inFlux`: Inbound traffic;</li>
	// <li>`l4Flow_outFlux`: Outbound traffic;</li>
	// <li>`l4Flow_outPkt`: Outbound packets.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>ruleId<br>   Filter by the specified <strong>forwarding rule ID</strong></li>
	// <li>proxyId<br>   Filter by the specified <strong>L4 agent ID</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metric to query. Values:
	// <li>`l4Flow_connections`: Access connections;</li>
	// <li>`l4Flow_flux`: Access traffic;</li>
	// <li>`l4Flow_inFlux`: Inbound traffic;</li>
	// <li>`l4Flow_outFlux`: Outbound traffic;</li>
	// <li>`l4Flow_outPkt`: Outbound packets.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>ruleId<br>   Filter by the specified <strong>forwarding rule ID</strong></li>
	// <li>proxyId<br>   Filter by the specified <strong>L4 agent ID</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ProxyIds")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL4DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataResponseParams struct {
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of L4 traffic data recorded over time.
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
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The list of metrics. Values:
	// <li>`l7Flow_outFlux`: Access traffic;</li>
	// <li>`l7Flow_request`: Access requests;</li>
	// <li>`l7Flow_outBandwidth`: Access bandwidth.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`province`<br>   Filter by the specified <strong>province name</strong>. It's only available when `Area` is `mainland`. </li>
	// <li>`isp`<br>   Filter by the specified <strong>ISP</strong>. It's only available when `Area` is `mainland`.<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others</li>
	// <li>`domain`<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`<br>   Filter by the specified <strong>URL Path</strong> (such as `/content` or `content/test.jpg`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`referer`<br>   Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`resourceType`<br>   Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0<br>   `HTTP/1.1`: HTTP 1.1<br>   `HTTP/2.0`: HTTP 2.0<br>   `HTTP/3.0`: HTTP 3.0<br>   `WebSocket`: WebSocket</li>
	// <li>`statusCode`<br>   Filter by the specified <strong> status code</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>  `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`<br>   Filter by the specified <strong>browser type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`<br>   Filter by the <strong>device type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`<br>   Filter by the <strong>operating system</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Mac OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`<br>   Filter by the <strong>TLS version</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>   Filter by the specified <strong>IP version</strong>.<br>   Values:<br>   `4`: IPv4;<br>   `6`: IPv6.</li>
	// <li>tagKey<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>tagValue<br>   Filter by the specified <strong>tag value</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The list of metrics. Values:
	// <li>`l7Flow_outFlux`: Access traffic;</li>
	// <li>`l7Flow_request`: Access requests;</li>
	// <li>`l7Flow_outBandwidth`: Access bandwidth.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it's not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`province`<br>   Filter by the specified <strong>province name</strong>. It's only available when `Area` is `mainland`. </li>
	// <li>`isp`<br>   Filter by the specified <strong>ISP</strong>. It's only available when `Area` is `mainland`.<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others</li>
	// <li>`domain`<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`<br>   Filter by the specified <strong>URL Path</strong> (such as `/content` or `content/test.jpg`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`referer`<br>   Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`resourceType`<br>   Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0<br>   `HTTP/1.1`: HTTP 1.1<br>   `HTTP/2.0`: HTTP 2.0<br>   `HTTP/3.0`: HTTP 3.0<br>   `WebSocket`: WebSocket</li>
	// <li>`statusCode`<br>   Filter by the specified <strong> status code</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>  `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`<br>   Filter by the specified <strong>browser type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`<br>   Filter by the <strong>device type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`<br>   Filter by the <strong>operating system</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Mac OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`<br>   Filter by the <strong>TLS version</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>   Filter by the specified <strong>IP version</strong>.<br>   Values:<br>   `4`: IPv4;<br>   `6`: IPv6.</li>
	// <li>tagKey<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>tagValue<br>   Filter by the specified <strong>tag value</strong></li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataResponseParams struct {
	// The list of L7 traffic data recorded over time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux`: Response traffic.</li>
	// <li>`l7Cache_request`: Response requests.</li>
	// <li>`l7Cache_outBandwidth`: Response bandwidth.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTimingL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux`: Response traffic.</li>
	// <li>`l7Cache_request`: Response requests.</li>
	// <li>`l7Cache_outBandwidth`: Response bandwidth.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataResponseParams struct {
	// The list of cached L7 time-series data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Flow_outFlux_country`: Country the request came from;</li>
	// <li>`l7Flow_outFlux_statusCode`: Status code of the request;</li>
	// <li>`l7Flow_outFlux_domain`: Domain name of the request;</li>
	// <li>`l7Flow_outFlux_url`: URL of the request;</li>
	// <li>`l7Flow_outFlux_resourceType`: Resource type;</li>
	// <li>`l7Flow_outFlux_sip`: Client IP;</li>
	// <li>`l7Flow_outFlux_referers`: Refer header;</li>
	// <li>`l7Flow_outFlux_ua_device`: Device type;</li>
	// <li>`l7Flow_outFlux_ua_browser`: Browser type;</li>
	// <li>`l7Flow_outFlux_us_os`: OS type;</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Queries the top n rows of data. Maximum value: 1000. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter conditions. See below for details: 
	// <li>`country`:<br>   Filter by the <strong>country/region code</strong>. <a href="https://zh.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.<br>   Type: String<br>   Required: No</li>
	// <li>`province`<br>   Filter by the <strong>province name</strong>. It’s only available when `Area` is `mainland`. <br>   Type: String<br>   Required: No</li>
	// <li>`isp`<br>   Filter by the <strong>ISP</strong>. It’s only available when `Area` is `mainland`.<br>   Type: String<br>   Required: No<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others</li>
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. Separate multiple URLs with semicolons. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`referer`<br>   Filter by the <strong>Referer header</strong>, such as `example.com`. The query period cannot exceed 30 days.<br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>`protocol`:<br>   Filter by the <strong>HTTP protocol</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `HTTP/1.0`: HTTP 1.0<br>   `HTTP/1.1`: HTTP 1.1<br>   `HTTP/2.0`: HTTP 2.0<br>   `HTTP/3.0`: HTTP 3.0<br>   `WebSocket`: WebSocket</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`<br>   Filter by the <strong>browser type</strong>. The query period cannot exceed 30 days. <br>   Type:  String<br>   Required:  No<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`<br>   Filter by the <strong>device type</strong>. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`<br>   Filter by the <strong>operating system</strong>. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Max OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`<br>   Filter by the <strong>TLS version</strong>. The query period cannot exceed 30 days.<br>   Type: String<br>   Required: No<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>   Filter by the <strong>IP version</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `4`: IPv4；<br>   `6`: IPv6.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Flow_outFlux_country`: Country the request came from;</li>
	// <li>`l7Flow_outFlux_statusCode`: Status code of the request;</li>
	// <li>`l7Flow_outFlux_domain`: Domain name of the request;</li>
	// <li>`l7Flow_outFlux_url`: URL of the request;</li>
	// <li>`l7Flow_outFlux_resourceType`: Resource type;</li>
	// <li>`l7Flow_outFlux_sip`: Client IP;</li>
	// <li>`l7Flow_outFlux_referers`: Refer header;</li>
	// <li>`l7Flow_outFlux_ua_device`: Device type;</li>
	// <li>`l7Flow_outFlux_ua_browser`: Browser type;</li>
	// <li>`l7Flow_outFlux_us_os`: OS type;</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Queries the top n rows of data. Maximum value: 1000. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter conditions. See below for details: 
	// <li>`country`:<br>   Filter by the <strong>country/region code</strong>. <a href="https://zh.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.<br>   Type: String<br>   Required: No</li>
	// <li>`province`<br>   Filter by the <strong>province name</strong>. It’s only available when `Area` is `mainland`. <br>   Type: String<br>   Required: No</li>
	// <li>`isp`<br>   Filter by the <strong>ISP</strong>. It’s only available when `Area` is `mainland`.<br>   Type: String<br>   Required: No<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others</li>
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. Separate multiple URLs with semicolons. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`referer`<br>   Filter by the <strong>Referer header</strong>, such as `example.com`. The query period cannot exceed 30 days.<br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>`protocol`:<br>   Filter by the <strong>HTTP protocol</strong><br>   Type: String<br>   Required: No<br>   Values:<br>   `HTTP/1.0`: HTTP 1.0<br>   `HTTP/1.1`: HTTP 1.1<br>   `HTTP/2.0`: HTTP 2.0<br>   `HTTP/3.0`: HTTP 3.0<br>   `WebSocket`: WebSocket</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`<br>   Filter by the <strong>browser type</strong>. The query period cannot exceed 30 days. <br>   Type:  String<br>   Required:  No<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`<br>   Filter by the <strong>device type</strong>. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`<br>   Filter by the <strong>operating system</strong>. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Max OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`<br>   Filter by the <strong>TLS version</strong>. The query period cannot exceed 30 days.<br>   Type: String<br>   Required: No<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>   Filter by the <strong>IP version</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `4`: IPv4；<br>   `6`: IPv6.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataResponseParams struct {
	// The list of top-ranked L7 traffic data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux_domain`: Host/Domain name;</li>
	// <li>`l7Cache_outFlux_url`: URL address;</li>
	// <li>`l7Cache_outFlux_resourceType`: Resource type;</li>
	// <li>`l7Cache_outFlux_statusCode`: Status code.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Queries the top rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeTopL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux_domain`: Host/Domain name;</li>
	// <li>`l7Cache_outFlux_url`: URL address;</li>
	// <li>`l7Cache_outFlux_resourceType`: Resource type;</li>
	// <li>`l7Cache_outFlux_statusCode`: Status code.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Queries the top rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataResponseParams struct {
	// The list of cached L7 top-ranked traffic data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopDataRecord `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeWebManagedRulesDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`waf_interceptNum`: Requests blocked by WAF.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebManagedRulesDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`waf_interceptNum`: Requests blocked by WAF.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "QueryCondition")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesDataResponseParams struct {
	// The list of WAF attack data recorded over time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeWebManagedRulesHitRuleDetailRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebManagedRulesHitRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebManagedRulesHitRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesHitRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesHitRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesHitRuleDetailResponseParams struct {
	// The hit rule information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecHitRuleInfo `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebManagedRulesHitRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebManagedRulesHitRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeWebManagedRulesHitRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebManagedRulesHitRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesLogRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`attackType`: Attack type;</li>
	// <li>`riskLevel`: Risk level;</li>
	// <li>`action`: Action;</li>
	// <li>`ruleId`: Rule ID;</li>
	// <li>`sipCountryCode`: Country code of the attacker IP;</li>
	// <li>`attackIp`: Attacker IP;</li>
	// <li>`oriDomain`: Attacked subdomain name;</li>
	// <li>`eventId`: Event ID;</li>
	// <li>`ua`: User agent;</li>
	// <li>`requestMethod`: Request method;</li>
	// <li>`uri`: Uniform resource identifier.</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebManagedRulesLogRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`attackType`: Attack type;</li>
	// <li>`riskLevel`: Risk level;</li>
	// <li>`action`: Action;</li>
	// <li>`ruleId`: Rule ID;</li>
	// <li>`sipCountryCode`: Country code of the attacker IP;</li>
	// <li>`attackIp`: Attacker IP;</li>
	// <li>`oriDomain`: Attacked subdomain name;</li>
	// <li>`eventId`: Event ID;</li>
	// <li>`ua`: User agent;</li>
	// <li>`requestMethod`: Request method;</li>
	// <li>`uri`: Uniform resource identifier.</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebManagedRulesLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebManagedRulesLogResponseParams struct {
	// The list of web log data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*WebLogs `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeWebProtectionClientIpListRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionClientIpListRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebProtectionClientIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionClientIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionClientIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionClientIpListResponseParams struct {
	// The list of CC attacker IPs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecClientIp `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionClientIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionClientIpListResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionClientIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionClientIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metrics to query. Values:
	// <li>`ccRate_interceptNum`: Requests restricted by the rate limiting rules;</li>
	// <li>`ccAcl_interceptNum`: Requests restricted by the custom rules.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Metrics to query. Values:
	// <li>`ccRate_interceptNum`: Requests restricted by the rate limiting rules;</li>
	// <li>`ccAcl_interceptNum`: Requests restricted by the custom rules.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionDataResponseParams struct {
	// The list of CC protection data recorded over time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeWebProtectionHitRuleDetailRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The rule type. Values:
	// <li>`rate`: Rate limiting rules;</li>
	// <li>`acl`: Custom rules.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionHitRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The rule type. Values:
	// <li>`rate`: Rate limiting rules;</li>
	// <li>`acl`: Custom rules.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebProtectionHitRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionHitRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EntityType")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "QueryCondition")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionHitRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionHitRuleDetailResponseParams struct {
	// The list of hit CC protection rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecHitRuleInfo `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionHitRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionHitRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionHitRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionHitRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionTopDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`ccRate_requestNum_url`: Top-ranked URLs by rate limiting requests.</li>
	// <li>`ccRate_cipRequestNum_region`: Top-ranked client IPs by rate limiting requests.</li>
	// <li>`ccAcl_requestNum_url`: Top-ranked URLs by custom rule requests.</li>
	// <li>`ccAcl_requestNum_cip`: Top-ranked client IPs by custom rule execution requests.</li>
	// <li>`ccAcl_cipRequestNum_region`: Top-ranked clients by custom rule execution requests.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries the top n rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeWebProtectionTopDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`ccRate_requestNum_url`: Top-ranked URLs by rate limiting requests.</li>
	// <li>`ccRate_cipRequestNum_region`: Top-ranked client IPs by rate limiting requests.</li>
	// <li>`ccAcl_requestNum_url`: Top-ranked URLs by custom rule requests.</li>
	// <li>`ccAcl_requestNum_cip`: Top-ranked client IPs by custom rule execution requests.</li>
	// <li>`ccAcl_cipRequestNum_region`: Top-ranked clients by custom rule execution requests.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries the top n rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeWebProtectionTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Interval")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionTopDataResponseParams struct {
	// The list of top-ranked CC protection data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopEntry `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebProtectionTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebProtectionTopDataResponseParams `json:"Response"`
}

func (r *DescribeWebProtectionTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebProtectionTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneSettingRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
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
	// The site configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneSetting *ZoneSetting `json:"ZoneSetting,omitempty" name:"ZoneSetting"`

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
	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-name`:<br>   Filter by <strong>site name</strong><br>   Type: String<br>   Required: No</li><li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No</li><li>`status`:<br>   Filter by <strong>site status</strong><br>   Type: String<br>   Required: No</li><li>`tag-key`:<br>   Filter by <strong>tag key</strong><br>   Type: String<br>   Required: No</li><li>`tag-value`:<br>   Filter by <strong>tag value</strong><br>   Type: String<br>   Required: No</li>Only `zone-name` supports fuzzy query.
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// The sorting field. Values:
	// <li>`type`: Access mode</li>
	// <li>`area`: Acceleration region</li>
	// <li>`create-time`: Creation date</li>
	// <li>`zone-name`: Site name</li>
	// <li>`use-time`: Last used date</li>
	// <li>`active-status`: Activation status</li>If it is left empty, the default value `create-time` is used.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The sorting direction. Values:
	// <li>`asc`: From smallest to largest</li>
	// <li>`desc`: From largest to smallest</li>If it is left empty, the default value `desc` is used.
	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-name`:<br>   Filter by <strong>site name</strong><br>   Type: String<br>   Required: No</li><li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No</li><li>`status`:<br>   Filter by <strong>site status</strong><br>   Type: String<br>   Required: No</li><li>`tag-key`:<br>   Filter by <strong>tag key</strong><br>   Type: String<br>   Required: No</li><li>`tag-value`:<br>   Filter by <strong>tag value</strong><br>   Type: String<br>   Required: No</li>Only `zone-name` supports fuzzy query.
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// The sorting field. Values:
	// <li>`type`: Access mode</li>
	// <li>`area`: Acceleration region</li>
	// <li>`create-time`: Creation date</li>
	// <li>`zone-name`: Site name</li>
	// <li>`use-time`: Last used date</li>
	// <li>`active-status`: Activation status</li>If it is left empty, the default value `create-time` is used.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The sorting direction. Values:
	// <li>`asc`: From smallest to largest</li>
	// <li>`desc`: From largest to smallest</li>If it is left empty, the default value `desc` is used.
	Direction *string `json:"Direction,omitempty" name:"Direction"`
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
	delete(f, "Order")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Number of eligible sites.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Details of sites
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The acceleration status. Values:
	// <li>`process`: In progress</li>
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The domain name.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Name of the site
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The assigned CNAME
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// The resource ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The lock status.
	Lock *int64 `json:"Lock,omitempty" name:"Lock"`

	// The domain name status.
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// The acceleration area of the domain name. Values:
	// <li>`global`: Global.</li>
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Outside the Chinese mainland.</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// The acceleration type configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccelerateType *AccelerateType `json:"AccelerateType,omitempty" name:"AccelerateType"`

	// The HTTPS configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// The cache configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// The origin configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// The security type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecurityType *SecurityType `json:"SecurityType,omitempty" name:"SecurityType"`

	// The cache key configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// The smart compression configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// The WAF protection configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Waf *Waf `json:"Waf,omitempty" name:"Waf"`

	// The CC protection configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CC *CC `json:"CC,omitempty" name:"CC"`

	// DDoS mitigation configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	DDoS *DDoS `json:"DDoS,omitempty" name:"DDoS"`

	// The smart routing configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// The IPv6 access configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`
}

type DistrictStatistics struct {
	// The ISO 3166-2 Alpha-2 country code. For the list of country codes, see [ISO 3166-2](https://zh.m.wikipedia.org/zh-hans/ISO_3166-2).
	Alpha2 *string `json:"Alpha2,omitempty" name:"Alpha2"`

	// The overall load time, in milliseconds.
	LoadTime *int64 `json:"LoadTime,omitempty" name:"LoadTime"`
}

type DnsData struct {
	// The time.
	Time *string `json:"Time,omitempty" name:"Time"`

	// The value.
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type DownloadL4LogsRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DownloadL4LogsRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DownloadL4LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL4LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "ProxyIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL4LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL4LogsResponseParams struct {
	// The list of L4 log data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*L4OfflineLog `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadL4LogsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadL4LogsResponseParams `json:"Response"`
}

func (r *DownloadL4LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL4LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL7LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsResponseParams struct {
	// The list of L7 log data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*L7OfflineLog `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

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

type DropPageConfig struct {
	// Whether to enable configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The settings of the block page that applies managed rules. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WafDropPageDetail *DropPageDetail `json:"WafDropPageDetail,omitempty" name:"WafDropPageDetail"`

	// The settings of the block page that applies custom rules. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AclDropPageDetail *DropPageDetail `json:"AclDropPageDetail,omitempty" name:"AclDropPageDetail"`
}

type DropPageDetail struct {
	// The ID of the block page, which can be obtained from the CreateSecurityDropPage API.
	// If 0 is passed, the default block page will be used.
	PageId *int64 `json:"PageId,omitempty" name:"PageId"`

	// The HTTP status code of the block page. Value range: 100-600.
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// The block page file or URL.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of the block page. Values:
	// <li>`file`: Block page file</li>
	// <li>`url`: Block page URL</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type ExceptConfig struct {
	// Whether to enable configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The settings of the exception rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptUserRules []*ExceptUserRule `json:"ExceptUserRules,omitempty" name:"ExceptUserRules"`
}

type ExceptUserRule struct {
	// The rule name.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// The rule action. It only supports the value `skip`, which indicates skipping all managed rules.
	Action *string `json:"Action,omitempty" name:"Action"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// The rule ID, which is automatically created and only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// The update time. If it is null, the current date and time is recorded.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The matching condition.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptUserRuleConditions []*ExceptUserRuleCondition `json:"ExceptUserRuleConditions,omitempty" name:"ExceptUserRuleConditions"`

	// The scope to which the exception rule applies.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptUserRuleScope *ExceptUserRuleScope `json:"ExceptUserRuleScope,omitempty" name:"ExceptUserRuleScope"`

	// The rule priority. Value range: 0-100. If it is null, it defaults to 0.
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`
}

type ExceptUserRuleCondition struct {
	// The field to match. Values:
	// <li>`host`: Request domain name</li>
	// <li>`sip`: Client IP</li>
	// <li>`ua`: User-Agent</li>
	// <li>`cookie`: Cookie</li>
	// <li>`cgi`: CGI script</li>
	// <li>`xff`: XFF header</li>
	// <li>`url`: Request URL</li>
	// <li>`accept`: Request content type</li>
	// <li>`method`: Request method</li>
	// <li>`header`: Request header</li>
	// <li>`sip_proto`: Network layer protocol</li>
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// The parameter of the field. Only when `MatchFrom = header`, the key contained in the header can be passed.
	MatchParam *string `json:"MatchParam,omitempty" name:"MatchParam"`

	// The logical operator. Values:
	// <li>`equal`: String equals</li>
	// <li>`not_equal`: Value not equals</li>
	// <li>`include`: String contains</li>
	// <li>`not_include`: String not contains</li>
	// <li>`match`: IP matches</li>
	// <li>`not_match`: IP not matches</li>
	// <li>`include_area`: Regions contain</li>
	// <li>`is_empty`: Value left empty</li>
	// <li>`not_exists`: Key fields not exist</li>
	// <li>`regexp`: Regex matches</li>
	// <li>`len_gt`: Value greater than</li>
	// <li>`len_lt`: Value smaller than</li>
	// <li>`len_eq`: Value equals</li>
	// <li>`match_prefix`: Prefix matches</li>
	// <li>`match_suffix`: Suffix matches</li>
	// <li>`wildcard`: Wildcard</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// The value of the parameter.
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type ExceptUserRuleScope struct {
	// Exception mode. Values:
	// <li>`complete`: Skip the exception rule for full requests.</li>
	// <li>`partial`: Skip the exception rule for partial requests.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// The module to be activated. Values:
	// <li>`waf`: Managed rules</li>
	// <li>`cc`: Rate limiting rules</li>
	// <li>`bot`: bot protection</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Modules []*string `json:"Modules,omitempty" name:"Modules"`

	// Module settings of the exception rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PartialModules []*PartialModule `json:"PartialModules,omitempty" name:"PartialModules"`

	// Condition settings of the exception rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SkipConditions []*SkipCondition `json:"SkipConditions,omitempty" name:"SkipConditions"`
}

type FailReason struct {
	// Failure reason.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// List of resources failed to be processed. 
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type FileAscriptionInfo struct {
	// Directory of the verification file.
	IdentifyPath *string `json:"IdentifyPath,omitempty" name:"IdentifyPath"`

	// Content of the verification file.
	IdentifyContent *string `json:"IdentifyContent,omitempty" name:"IdentifyContent"`
}

type Filter struct {
	// Fields to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value of the filtered field.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FollowOrigin struct {
	// Whether to enable the configuration of following the origin server. Valid values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Sets the default cache time when the origin server does not return the Cache-Control header.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DefaultCacheTime *int64 `json:"DefaultCacheTime,omitempty" name:"DefaultCacheTime"`

	// Specifies whether to enable cache when the origin server does not return the Cache-Control header.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DefaultCache *string `json:"DefaultCache,omitempty" name:"DefaultCache"`

	// Specifies whether to use the default caching policy when Cache-Control is not returned from the origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DefaultCacheStrategy *string `json:"DefaultCacheStrategy,omitempty" name:"DefaultCacheStrategy"`
}

type ForceRedirect struct {
	// Whether to enable force HTTPS redirect. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Redirect status code. Values:
	// <li>`301`: 301 redirect</li>
	// <li>`302`: 302 redirect</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type Grpc struct {
	// Whether to enable gRPC support
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Header struct {
	// HTTP header name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// HTTP header value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Hsts struct {
	// Whether to enable the configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The MaxAge value in seconds. Maximum value: `86400` (one day)
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxAge *int64 `json:"MaxAge,omitempty" name:"MaxAge"`

	// Whether to contain subdomain names. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncludeSubDomains *string `json:"IncludeSubDomains,omitempty" name:"IncludeSubDomains"`

	// Whether to enable preloading. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Preload *string `json:"Preload,omitempty" name:"Preload"`
}

type Https struct {
	// Whether to enable HTTP2. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// Whether to enable OCSP. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// TLS version. Values:
	// <li>`TLSv1`: TLSv1 version</li>
	// <li>`TLSV1.1`: TLSv1.1 version</li>
	// <li>`TLSV1.2`: TLSv1.2 version</li>
	// <li>`TLSv1.3`: TLSv1.3 version</li>Only consecutive versions can be enabled at the same time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TlsVersion []*string `json:"TlsVersion,omitempty" name:"TlsVersion"`

	// HSTS Configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Hsts *Hsts `json:"Hsts,omitempty" name:"Hsts"`

	// The certificate configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertInfo []*ServerCertInfo `json:"CertInfo,omitempty" name:"CertInfo"`

	// Whether the certificate is managed by EdgeOne. Values:
	// <li>`apply`: Managed by EdgeOne.</li>
	// <li>`none`: Not managed by EdgeOne.</li>If it is left empty, the default value `none` is used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ApplyType *string `json:"ApplyType,omitempty" name:"ApplyType"`
}

type Identification struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The verification status. Values:
	// <li>`pending`: The verification is ongoing.</li>
	// <li>`finished`: The verification completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Details of the DNS record.
	Ascription *AscriptionInfo `json:"Ascription,omitempty" name:"Ascription"`

	// The NS record of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// Details of the verification file.
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitempty" name:"FileAscription"`
}

// Predefined struct for user
type IdentifyZoneRequestParams struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
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
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdentifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdentifyZoneResponseParams struct {
	// Details of the DNS record.
	Ascription *AscriptionInfo `json:"Ascription,omitempty" name:"Ascription"`

	// Details of the verification file.
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitempty" name:"FileAscription"`

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

type IntelligenceRule struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Items in a bot intelligence rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntelligenceRuleItems []*IntelligenceRuleItem `json:"IntelligenceRuleItems,omitempty" name:"IntelligenceRuleItems"`
}

type IntelligenceRuleItem struct {
	// The tag to categorize bots. Values:
	// <li>`evil_bot`: Malicious bot</li>
	// <li>`suspect_bot`: Suspected bot</li>
	// <li>`good_bot`: Good bot</li>
	// <li>`normal`: Normal request</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// The action taken on bots. Values
	// <li>`drop`: Block</li>
	// <li>`trans`: Allow</li>
	// <li>`alg`: JavaScript challenge</li>
	// <li>`captcha`: Managed challenge</li>
	// <li>`monitor`: Observe</li>
	Action *string `json:"Action,omitempty" name:"Action"`
}

type IpTableConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The settings of the basic access control rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IpTableRules []*IpTableRule `json:"IpTableRules,omitempty" name:"IpTableRules"`
}

type IpTableRule struct {
	// The action. Values:
	// <li>`drop`: Block</li>
	// <li>`trans`: Allow</li>
	// <li>`monitor`: Observe</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// The matching dimension. Values:
	// <li>`ip`: Match by IP.</li>
	// <li>`area`: Match by IP region.</li>
	MatchFrom *string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// The matching content.
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`

	// The rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// The update time, which is only used as an output parameter.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The rule status. A null value indicates that the rule is enabled. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type Ipv6 struct {
	// Whether to enable IPv6 access. Values:
	// <li>`on`: Enable IPv6 access.</li>
	// <li>`off`: Disable IPv6 access.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type L4OfflineLog struct {
	// The start time of the log packaging.
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// The L4 proxy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The log size, in bytes.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// The download address.
	Url *string `json:"Url,omitempty" name:"Url"`

	// The log package name.
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`

	// The acceleration region. Values:
	// <li>`mainland`: Chinese mainland;</li>
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type L7OfflineLog struct {
	// Start time of the log packaging
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// The subdomain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Log size, in bytes.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Download address
	Url *string `json:"Url,omitempty" name:"Url"`

	// Log package name
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland;</li>
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type LogSetInfo struct {
	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// Name of the logset.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// Whether the logset is deleted. Values:
	// <li>`no`: The logset is not deleted;</li>
	// <li>`yes`: The logset is deleted.</li>
	Deleted *string `json:"Deleted,omitempty" name:"Deleted"`
}

type MaxAge struct {
	// Whether to follow the origin server. Values:
	// <li>`on`: Follow the origin server and ignore the field MaxAgeTime;</li>
	// <li>`off`: Do not follow the origin server and apply the field MaxAgeTime.</li>
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`

	// Specifies the maximum amount of time (in seconds). The maximum value is 365 days.
	// Note: The value `0` means not to cache.
	MaxAgeTime *int64 `json:"MaxAgeTime,omitempty" name:"MaxAgeTime"`
}

// Predefined struct for user
type ModifyAliasDomainRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: Off</li>
	// <li>`hosting`: Managed SSL certificate</li>
	// <li>`apply`: Free certificate</li>The original configuration will apply if this field is not specified.
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitempty" name:"CertId"`
}

type ModifyAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: Off</li>
	// <li>`hosting`: Managed SSL certificate</li>
	// <li>`apply`: Free certificate</li>The original configuration will apply if this field is not specified.
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitempty" name:"CertId"`
}

func (r *ModifyAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasName")
	delete(f, "TargetName")
	delete(f, "CertType")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAliasDomainResponseParams `json:"Response"`
}

func (r *ModifyAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainStatusRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Status of the alias domain name. Values:
	// <li>`false`: Enable the alias domain name.</li>
	// <li>`true`: Disable the alias domain name.</li>
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// The alias domain name you want to modify its status. If it is left empty, the modify operation is not performed.
	AliasNames []*string `json:"AliasNames,omitempty" name:"AliasNames"`
}

type ModifyAliasDomainStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Status of the alias domain name. Values:
	// <li>`false`: Enable the alias domain name.</li>
	// <li>`true`: Disable the alias domain name.</li>
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// The alias domain name you want to modify its status. If it is left empty, the modify operation is not performed.
	AliasNames []*string `json:"AliasNames,omitempty" name:"AliasNames"`
}

func (r *ModifyAliasDomainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Paused")
	delete(f, "AliasNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAliasDomainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAliasDomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAliasDomainStatusResponseParams `json:"Response"`
}

func (r *ModifyAliasDomainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The domain name or subdomain name when `ProxyType=hostname`.
	// The instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// The original configuration will apply if this field is not specified.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>If not specified, this field uses the default value `instance`.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// The IPv6 access configuration. The original configuration will apply if this field is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The domain name or subdomain name when `ProxyType=hostname`.
	// The instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// The original configuration will apply if this field is not specified.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>If not specified, this field uses the default value `instance`.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// The IPv6 access configuration. The original configuration will apply if this field is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`
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
	delete(f, "SessionPersistTime")
	delete(f, "ProxyType")
	delete(f, "Ipv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyResponseParams struct {
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The origin type. Values:
	// <li>`custom`: Specified origins</li>
	// <li>`origins`: Origin group</li></li>The original configuration will apply if this field is not specified.
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitempty" name:"Port"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>The original configuration will apply if this field is not specified.
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Origin server information:
	// <li>When `OriginType=custom`, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or ["test.com"].</li>
	// <li>When `OriginType=origins`, it indicates an origin group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	// 
	// The original configuration will apply if this field is not specified.
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>If not specified, this field uses the default value OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>If it is left empty, the default value `false` is used.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitempty" name:"OriginPort"`
}

type ModifyApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The origin type. Values:
	// <li>`custom`: Specified origins</li>
	// <li>`origins`: Origin group</li></li>The original configuration will apply if this field is not specified.
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitempty" name:"Port"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>The original configuration will apply if this field is not specified.
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Origin server information:
	// <li>When `OriginType=custom`, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or ["test.com"].</li>
	// <li>When `OriginType=origins`, it indicates an origin group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	// 
	// The original configuration will apply if this field is not specified.
	OriginValue []*string `json:"OriginValue,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>If not specified, this field uses the default value OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>If it is left empty, the default value `false` is used.
	SessionPersist *bool `json:"SessionPersist,omitempty" name:"SessionPersist"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitempty" name:"OriginPort"`
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
	delete(f, "OriginType")
	delete(f, "Port")
	delete(f, "Proto")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	delete(f, "OriginPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleResponseParams struct {
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The rule status. Values:
	// <li>`offline`: Disabled</li>
	// <li>`online`: Enabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The rule status. Values:
	// <li>`offline`: Disabled</li>
	// <li>`online`: Enabled</li>
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The proxy status. Values:
	// <li>`offline`: The proxy is disabled.</li>
	// <li>`online`: The proxy is enabled.</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// The proxy status. Values:
	// <li>`offline`: The proxy is disabled.</li>
	// <li>`online`: The proxy is enabled.</li>
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
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// ID of the certificate.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Status of the certificate. Values:
	// <li>`deployed`: The certificate is deployed;</li>
	// <li>`disabled`: The certificate is disabled.</li>When a deployment fails, you can try again.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyDefaultCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// ID of the certificate.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Status of the certificate. Values:
	// <li>`deployed`: The certificate is deployed;</li>
	// <li>`disabled`: The certificate is disabled.</li>When a deployment fails, you can try again.
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
type ModifyHostsCertificateRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names that the certificate will be attached to.
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Certificate information. Note that only `CertId` is required. If it is not specified, the default certificate will be used.
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitempty" name:"ServerCertInfo"`

	// Whether the certificate is managed by EdgeOne. Values:
	// <li>`apply`: Managed by EdgeOne</li>
	// <li>`none`: Not managed by EdgeOne</li>If it is left empty, the default value `apply` is used.
	ApplyType *string `json:"ApplyType,omitempty" name:"ApplyType"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names that the certificate will be attached to.
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Certificate information. Note that only `CertId` is required. If it is not specified, the default certificate will be used.
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitempty" name:"ServerCertInfo"`

	// Whether the certificate is managed by EdgeOne. Values:
	// <li>`apply`: Managed by EdgeOne</li>
	// <li>`none`: Not managed by EdgeOne</li>If it is left empty, the default value `apply` is used.
	ApplyType *string `json:"ApplyType,omitempty" name:"ApplyType"`
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
	delete(f, "ServerCertInfo")
	delete(f, "ApplyType")
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
type ModifyOriginGroupRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The ID of the origin group.
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// The origin type. Values:
	// <li>`self`: Customer origin</li>
	// <li>`third_party`: Third-party origin</li>
	// <li>`cos`: Tencent Cloud COS origin</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The name of the origin group.
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// The origin configuration type when `OriginType=self`. Values:
	// <li>`area`: Configure by region.</li>
	// <li>`weight`: Configure by weight.</li>
	// <li>`proto`: Configure by HTTP protocol.</li> When `OriginType=third_party/cos`, leave this field empty.
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// Details of the origin record.
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`

	// The origin domain. This field can be specified only when `OriginType=self`.
	// If it is left empty, the existing configuration is used.
	HostHeader *string `json:"HostHeader,omitempty" name:"HostHeader"`
}

type ModifyOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The ID of the origin group.
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// The origin type. Values:
	// <li>`self`: Customer origin</li>
	// <li>`third_party`: Third-party origin</li>
	// <li>`cos`: Tencent Cloud COS origin</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The name of the origin group.
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// The origin configuration type when `OriginType=self`. Values:
	// <li>`area`: Configure by region.</li>
	// <li>`weight`: Configure by weight.</li>
	// <li>`proto`: Configure by HTTP protocol.</li> When `OriginType=third_party/cos`, leave this field empty.
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// Details of the origin record.
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`

	// The origin domain. This field can be specified only when `OriginType=self`.
	// If it is left empty, the existing configuration is used.
	HostHeader *string `json:"HostHeader,omitempty" name:"HostHeader"`
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
	delete(f, "ZoneId")
	delete(f, "OriginGroupId")
	delete(f, "OriginType")
	delete(f, "OriginGroupName")
	delete(f, "ConfigurationType")
	delete(f, "OriginRecords")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupResponseParams struct {
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
type ModifyRulePriorityRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Order of rule IDs. If there are multiple rules, they will be executed in order from top to bottom.
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

type ModifyRulePriorityRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Order of rule IDs. If there are multiple rules, they will be executed in order from top to bottom.
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *ModifyRulePriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRulePriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRulePriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRulePriorityResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRulePriorityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRulePriorityResponseParams `json:"Response"`
}

func (r *ModifyRulePriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRulePriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The rule name. It is a string that can contain 1–255 characters.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The rule name. It is a string that can contain 1–255 characters.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "ZoneId")
	delete(f, "RuleName")
	delete(f, "Rules")
	delete(f, "RuleId")
	delete(f, "Status")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifySecurityPolicyRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Security configuration.
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitempty" name:"SecurityConfig"`

	// The subdomain name/L4 proxy. You must specify either "Entity" or "TemplateId".
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// The template ID. You must specify either this field or "Entity".
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Security configuration.
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitempty" name:"SecurityConfig"`

	// The subdomain name/L4 proxy. You must specify either "Entity" or "TemplateId".
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// The template ID. You must specify either this field or "Entity".
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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
	delete(f, "SecurityConfig")
	delete(f, "Entity")
	delete(f, "TemplateId")
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
type ModifySecurityWafGroupPolicyRequestParams struct {
	// The site ID. You must specify either "ZoneId+Entity" or "TemplateId".
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The subdomain name. You must specify either "ZoneId+Entity" or "TemplateId". 
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>If not specified, it defaults to the setting that was last configured.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The rule level. Values:
	// <li>`loose`: Loose</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`stricter`: Super strict</li>
	// <li>`custom`: Custom</li>If not specified, it defaults to the setting that was last configured.
	Level *string `json:"Level,omitempty" name:"Level"`

	// The rule action. Values:
	// <li>`block`: Block</li>
	// <li>`observe`: Observe</li>If not specified, it defaults to the setting that was last configured.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The settings of the managed rule. If not specified, it defaults to the settings that were last configured.
	WafRules *WafRule `json:"WafRules,omitempty" name:"WafRules"`

	// The settings of the AI rule engine. If not specified, it defaults to the settings that were last configured.
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`

	// The settings of the managed rule group. If not specified, it defaults to the settings that were last configured.
	WafGroups []*WafGroup `json:"WafGroups,omitempty" name:"WafGroups"`

	// The template ID. You must specify either this field or "ZoneId+Entity".
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type ModifySecurityWafGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The site ID. You must specify either "ZoneId+Entity" or "TemplateId".
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The subdomain name. You must specify either "ZoneId+Entity" or "TemplateId". 
	Entity *string `json:"Entity,omitempty" name:"Entity"`

	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>If not specified, it defaults to the setting that was last configured.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The rule level. Values:
	// <li>`loose`: Loose</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`stricter`: Super strict</li>
	// <li>`custom`: Custom</li>If not specified, it defaults to the setting that was last configured.
	Level *string `json:"Level,omitempty" name:"Level"`

	// The rule action. Values:
	// <li>`block`: Block</li>
	// <li>`observe`: Observe</li>If not specified, it defaults to the setting that was last configured.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The settings of the managed rule. If not specified, it defaults to the settings that were last configured.
	WafRules *WafRule `json:"WafRules,omitempty" name:"WafRules"`

	// The settings of the AI rule engine. If not specified, it defaults to the settings that were last configured.
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`

	// The settings of the managed rule group. If not specified, it defaults to the settings that were last configured.
	WafGroups []*WafGroup `json:"WafGroups,omitempty" name:"WafGroups"`

	// The template ID. You must specify either this field or "ZoneId+Entity".
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *ModifySecurityWafGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityWafGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entity")
	delete(f, "Switch")
	delete(f, "Level")
	delete(f, "Mode")
	delete(f, "WafRules")
	delete(f, "AiRule")
	delete(f, "WafGroups")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityWafGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityWafGroupPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityWafGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityWafGroupPolicyResponseParams `json:"Response"`
}

func (r *ModifySecurityWafGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityWafGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site access method. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record.</li>The original configuration will apply if this field is not specified.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The custom name servers. If this field is not specified, the default name servers will be used.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitempty" name:"AliasZoneName"`
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site access method. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record.</li>The original configuration will apply if this field is not specified.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The custom name servers. If this field is not specified, the default name servers will be used.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitempty" name:"AliasZoneName"`
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
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "VanityNameServers")
	delete(f, "AliasZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneResponseParams struct {
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
	// The site ID to be modified.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Cache expiration time configuration
	// The original configuration will apply if this field is not specified.
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// The node cache key configuration.
	// The original configuration will apply if this field is not specified.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// The browser cache configuration.
	// The original configuration will apply if this field is not specified.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// The offline cache configuration.
	// The original configuration will apply if this field is not specified.
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// The QUIC access configuration.
	// The original configuration will apply if this field is not specified.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// The POST transport configuration.
	// The original configuration will apply if this field is not specified.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// The smart compression configuration.
	// The original configuration will apply if this field is not specified.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// The HTTP2 origin-pull configuration.
	// The original configuration will apply if this field is not specified.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// The force HTTPS redirect configuration.
	// The original configuration will apply if this field is not specified.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// The HTTPS acceleration configuration.
	// The original configuration will apply if this field is not specified.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// The origin server configuration.
	// The original configuration will apply if this field is not specified.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// The smart acceleration configuration.
	// The original configuration will apply if this field is not specified.
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// The WebSocket configuration.
	// The original configuration will apply if this field is not specified.
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// The origin-pull client IP header configuration.
	// The original configuration will apply if this field is not specified.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// The cache prefresh configuration.
	// The original configuration will apply if this field is not specified.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// The IPv6 access configuration.
	// The original configuration will apply if this field is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull.
	// The original configuration will apply if this field is not specified.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support
	// The original configuration will apply if this field is not specified.
	Grpc *Grpc `json:"Grpc,omitempty" name:"Grpc"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// The site ID to be modified.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Cache expiration time configuration
	// The original configuration will apply if this field is not specified.
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// The node cache key configuration.
	// The original configuration will apply if this field is not specified.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// The browser cache configuration.
	// The original configuration will apply if this field is not specified.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// The offline cache configuration.
	// The original configuration will apply if this field is not specified.
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// The QUIC access configuration.
	// The original configuration will apply if this field is not specified.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// The POST transport configuration.
	// The original configuration will apply if this field is not specified.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// The smart compression configuration.
	// The original configuration will apply if this field is not specified.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// The HTTP2 origin-pull configuration.
	// The original configuration will apply if this field is not specified.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// The force HTTPS redirect configuration.
	// The original configuration will apply if this field is not specified.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// The HTTPS acceleration configuration.
	// The original configuration will apply if this field is not specified.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// The origin server configuration.
	// The original configuration will apply if this field is not specified.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// The smart acceleration configuration.
	// The original configuration will apply if this field is not specified.
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// The WebSocket configuration.
	// The original configuration will apply if this field is not specified.
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// The origin-pull client IP header configuration.
	// The original configuration will apply if this field is not specified.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// The cache prefresh configuration.
	// The original configuration will apply if this field is not specified.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// The IPv6 access configuration.
	// The original configuration will apply if this field is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull.
	// The original configuration will apply if this field is not specified.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support
	// The original configuration will apply if this field is not specified.
	Grpc *Grpc `json:"Grpc,omitempty" name:"Grpc"`
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
	delete(f, "CacheConfig")
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
	delete(f, "Ipv6")
	delete(f, "ClientIpCountry")
	delete(f, "Grpc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingResponseParams struct {
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site status. Values:
	// <li>`false`: Disabled</li>
	// <li>`true`: Enabled</li>
	Paused *bool `json:"Paused,omitempty" name:"Paused"`
}

type ModifyZoneStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site status. Values:
	// <li>`false`: Disabled</li>
	// <li>`true`: Enabled</li>
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
	delete(f, "ZoneId")
	delete(f, "Paused")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneStatusResponseParams struct {
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

type NoCache struct {
	// Whether to enable no-cache configuration. Valid values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type NormalAction struct {
	// Feature name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) API to view the requirements for entering the feature name.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Parameter
	Parameters []*RuleNormalActionParams `json:"Parameters,omitempty" name:"Parameters"`
}

type OfflineCache struct {
	// Whether offline cache is enabled. Valid values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OptimizeAction struct {
	// The optimization metric. Values:
	// <li>`Http2`</li>
	// <li>`Http3`</li>
	// <li>`Brotli`</li>
	Name *string `json:"Name,omitempty" name:"Name"`

	// The network environment.
	Connectivity *string `json:"Connectivity,omitempty" name:"Connectivity"`

	// The estimated load time, in milliseconds.
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// The estimated improvement ratio, in %.
	Ratio *int64 `json:"Ratio,omitempty" name:"Ratio"`
}

type Origin struct {
	// Primary origin server list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origins []*string `json:"Origins,omitempty" name:"Origins"`

	// The list of backup origin servers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupOrigins []*string `json:"BackupOrigins,omitempty" name:"BackupOrigins"`

	// Origin-pull protocol configuration. Values:
	// <li>`http`: Force HTTP for origin-pull.</li>
	// <li>`follow`: Follow protocol.</li>
	// <li>`https`: Force HTTPS for origin-pull.</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// Whether to allow private access to buckets when `OriginType=cos`. Values:
	// <li>`on`: Allow private access.</li>
	// <li>`off`: Allow public access.</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CosPrivateAccess *string `json:"CosPrivateAccess,omitempty" name:"CosPrivateAccess"`
}

type OriginGroup struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The ID of the origin group.
	OriginGroupId *string `json:"OriginGroupId,omitempty" name:"OriginGroupId"`

	// The origin type. Values:
	// <li>`self`: Customer origin</li>
	// <li>`third_party`: Third-party origin</li>
	// <li>`cos`: Tencent Cloud COS origin</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The name of the origin group.
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// The origin configuration type when `OriginType=self`. Values:
	// <li>`area`: Configure by region.</li>
	// <li>`weight`: Configure by weight.</li>
	// <li>`proto`: Configure by HTTP protocol.</li>When `OriginType=third_party/cos`, leave this field empty.
	ConfigurationType *string `json:"ConfigurationType,omitempty" name:"ConfigurationType"`

	// The origin record information.
	OriginRecords []*OriginRecord `json:"OriginRecords,omitempty" name:"OriginRecords"`

	// The update time of the origin group.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The origin domain when `OriginType=self`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HostHeader *string `json:"HostHeader,omitempty" name:"HostHeader"`
}

type OriginRecord struct {
	// The origin record value, which can be an IPv4/IPv6 address or a domain name.
	Record *string `json:"Record,omitempty" name:"Record"`

	// The origin record ID.
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// The origin port. Value rang: 1-65535.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// The weight when `ConfigurationType=weight`.
	// If 0 or no value is passed, the weight of each origin in a group will be 0 or left empty, indicating that origin-pull is performed by round-robin.
	// If a value between 1-100 is passed, the total weight of multiple origins in a group should be 100, indicating that origin-pull is performed by weight.
	// The weight when `ConfigurationType=proto`.
	// If 0 or no value is passed, the weight of each origin in a group will be 0 or left empty, indicating that origin-pull is performed by round-robin.
	// If a value between 1-100 is passed, the total weight of multiple origins with the same protocol in a group should be 100, indicating that origin-pull is performed by weight.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// The origin protocol when `ConfigurationType=proto`, indicating that origin-pull is performed by protocol.
	// <li>`http`: HTTP protocol</li>
	// <li>`https`: HTTPS protocol</li>
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// The region when `ConfigurationType=area`, which is specified by country code (ISO 3166 alpha-2) or continent code. If not specified, it indicates all regions. Supported continent codes:
	// <li>`Asia`</li>
	// <li>`Europe`</li>
	// <li>`Africa`</li>
	// <li>`Oceania`</li>
	// <li>`Americas`</li>An origin group must have at least one origin configured with all regions.
	Area []*string `json:"Area,omitempty" name:"Area"`

	// It is valid only when `OriginType=third_part`.
	// Whether the origin group is private. Values:
	// <li>`true`: Yes.</li>
	// <li>`false`: No.</li>If not specified, it defaults to false.
	Private *bool `json:"Private,omitempty" name:"Private"`

	// The authentication parameter, which is used when `Private=true`.
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitempty" name:"PrivateParameters"`
}

type PartialModule struct {
	// The module. Values:
	// <li>`waf`: Managed rules</li>
	Module *string `json:"Module,omitempty" name:"Module"`

	// List of managed rule IDs to be skipped.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Include []*int64 `json:"Include,omitempty" name:"Include"`
}

type PlanInfo struct {
	// Settlement currency. Values:
	// <li>`CNY`: Settled by Chinese RMB;</li>
	// <li>`USD`: Settled by US dollars.</li>
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Traffic quota of the plan. It includes the traffic for security acceleration, content acceleration and smart acceleration. Unit: byte.
	Flux *uint64 `json:"Flux,omitempty" name:"Flux"`

	// Settlement cycle. Values:
	// <li>`y`: Settled by year;</li>
	// <li>`m`: Settled by month;</li>
	// <li>`h`: Settled by hour;</li>
	// <li>`M`: Settled by minute;</li>
	// <li>`s`: Settled by second.</li>
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// The plan option. Values:
	// <li>`sta`: Standard plan that supports content delivery network outside the Chinese mainland.</li>
	// <li>`sta_with_bot`: Standard plan that supports content delivery network outside the Chinese mainland and bot management.</li>
	// <li>`sta_cm`: Standard plan that supports content delivery network inside the Chinese mainland.</li>
	// <li>`sta_cm_with_bot`: Standard plan that supports content delivery network inside the Chinese mainland and bot management.</li>
	// <li>`sta`: Standard plan that supports content delivery network over the globe.</li>
	// <li>`sta_global_with_bot`: Standard plan that supports content delivery network over the globe and bot management.</li>
	// <li>`ent`: Enterprise plan that supports content delivery network outside the Chinese mainland.</li>
	// <li>`ent_with_bot`: Enterprise plan that supports content delivery network outside the Chinese mainland and bot management.</li>
	// <li>`ent_cm`: Enterprise plan that supports content delivery network inside the Chinese mainland.</li>
	// <li>`ent_cm_with_bot`: Enterprise plan that supports content delivery network inside the Chinese mainland and bot management.</li>
	// <li>`ent_global`: Enterprise plan that supports content delivery network over the globe.</li>
	// <li>`ent_global_with_bot`: Enterprise plan that supports content delivery network over the globe and bot management.</li>
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`

	// Plan price (in CNY fen/US cent). The price unit depends on the settlement currency.
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// Quota on security acceleration requests
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// Number of sites to be bound to the plan
	SiteNumber *uint64 `json:"SiteNumber,omitempty" name:"SiteNumber"`

	// The acceleration region. Values:
	// <li>`mainland`: Chinese mainland</li>
	// <li>`overseas`: Global (Chinese mainland not included)</li>
	// <li>`global`: Global (Chinese mainland included)</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type PostMaxSize struct {
	// Whether to enable POST upload limit (default limit: 32 MB). Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Maximum size. Value range: 1-500 MB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type PrivateParameter struct {
	// The parameter name. Values
	// <li>`AccessKeyId`: Access Key ID</li>
	// <li>`SecretAccessKey`: Secret Access Key</li>
	Name *string `json:"Name,omitempty" name:"Name"`

	// The parameter value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type QueryCondition struct {
	// The key of QueryCondition.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The conditional operator. Values:
	// <li>`equals`: Equals</li>
	// <li>`notEquals`: Does not equal</li>
	// <li>`include`: Contains</li>
	// <li>`notInclude`: Does not contain</li>
	// <li>`startWith`: Starts with</li>
	// <li>`notStartWith`: Does not start with</li>
	// <li>`endWith`: Ends with</li>
	// <li>`notEndWith`: Does not end with</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// The value of QueryCondition.
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type QueryString struct {
	// Whether to use `QueryString` as part of `CacheKey`. Values:
	// <li>`on`: Yes</li>
	// <li>`off`: No</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Specifies how to use query strings in the cache key. Values:
	// <li>`includeCustom`: `Include partial query strings.</li>
	// <li>`excludeCustom`: Exclude partial query strings.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Array of query strings used/excluded
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Quic struct {
	// Whether to enable QUIC. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Quota struct {

	Batch *int64 `json:"Batch,omitempty" name:"Batch"`

	// Daily submission quota limit.
	Daily *int64 `json:"Daily,omitempty" name:"Daily"`

	// Remaining daily submission quota.
	DailyAvailable *int64 `json:"DailyAvailable,omitempty" name:"DailyAvailable"`

	// Type of cache purging/pre-warming. Values:
	// <li>`purge_prefix`: Purge by prefix</li>
	// <li>`purge_url`: Purge by URL</li>
	// <li>`purge_host`: Purge by hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by cache tag</li><li>`prefetch_url`: Pre-warm by URL</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type RateLimitConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The settings of the custom rate limiting rule. If it is null, the settings that were last configured will be used.
	RateLimitUserRules []*RateLimitUserRule `json:"RateLimitUserRules,omitempty" name:"RateLimitUserRules"`

	// The settings of the rate limiting template. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RateLimitTemplate *RateLimitTemplate `json:"RateLimitTemplate,omitempty" name:"RateLimitTemplate"`

	// The client filtering settings. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RateLimitIntelligence *RateLimitIntelligence `json:"RateLimitIntelligence,omitempty" name:"RateLimitIntelligence"`
}

type RateLimitIntelligence struct {
	// Whether to enable configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Action to be executed. Values:
	// <li>`monitor`: Observe</li>
	// <li>`alg`: Challenge</li>
	Action *string `json:"Action,omitempty" name:"Action"`
}

type RateLimitTemplate struct {
	// The mode. Values:
	// <li>`sup_loose`: Super loose</li>
	// <li>`loose`: Loose</li>
	// <li>`emergency`: Emergency</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`close`: Off</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The action. Values:
	// <li>`alg`: JavaScript challenge</li>
	// <li>`monitor`: Observe</li>If it is left empty, the default value `alg` is used.
	Action *string `json:"Action,omitempty" name:"Action"`

	// The settings of the rate limiting template. It is only used as an output parameter.
	RateLimitTemplateDetail *RateLimitTemplateDetail `json:"RateLimitTemplateDetail,omitempty" name:"RateLimitTemplateDetail"`
}

type RateLimitTemplateDetail struct {
	// The mode. Values:
	// <li>`sup_loose`: Super loose</li>
	// <li>`loose`: Loose</li>
	// <li>`emergency`: Emergency</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`close`: Off</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The unique ID.
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// The action. Values:
	// <li>`alg`: JavaScript challenge</li>
	// <li>`monitor`: Observe</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// The blocking duration, in seconds. Value range: 0-172800.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// The request threshold. Value range: 0-4294967294.
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// The statistical period. Value range: 0-120 seconds.
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type RateLimitUserRule struct {
	// The request threshold. Value range: 0-4294967294.
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// The statistical period. The value can be 10, 20, 30, 40, 50, or 60 seconds.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// The rule name, which consists of only letters, digits, and underscores and cannot start with an underscore.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// The action. Values:
	// <li>`monitor`: Observe</li>
	// <li>`drop`: Block</li>
	// <li>`alg`: JavaScript challenge</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// The amount of time taken to perform the action. Value range: 0 seconds - 2 days.
	PunishTime *int64 `json:"PunishTime,omitempty" name:"PunishTime"`

	// The time unit. Values:
	// <li>`second`: Second</li>
	// <li>`minutes`: Minute</li>
	// <li>`hour`: Hour</li>
	PunishTimeUnit *string `json:"PunishTimeUnit,omitempty" name:"PunishTimeUnit"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>Default value: on
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// The rule details.
	AclConditions []*AclCondition `json:"AclConditions,omitempty" name:"AclConditions"`

	// The rule weight. Value range: 0-100.
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// The rule ID, which is only used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// The filter. Values:
	// <li>`sip`: Client IP</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	FreqFields []*string `json:"FreqFields,omitempty" name:"FreqFields"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The statistical dimension. Values:
	// <li>`source_to_eo`: Responses from the origin server to EdgeOne</li>
	// <li>`client_to_eo`: Requests from the client to EdgeOne</li>
	// Note: A null value indicates responses from the origin server to EdgeOne are recorded.
	FreqScope []*string `json:"FreqScope,omitempty" name:"FreqScope"`
}

// Predefined struct for user
type ReclaimAliasDomainRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type ReclaimAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *ReclaimAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReclaimAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReclaimAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReclaimAliasDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReclaimAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *ReclaimAliasDomainResponseParams `json:"Response"`
}

func (r *ReclaimAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReclaimAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReclaimZoneRequestParams struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type ReclaimZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
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
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReclaimZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReclaimZoneResponseParams struct {
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
	// The resource ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Billing mode
	// `0`: Pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The effective time.
	EnableTime *string `json:"EnableTime,omitempty" name:"EnableTime"`

	// The expiration time.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// The plan status. Values:
	// <li>`normal`: Normal</li>
	// <li>`isolated`: Isolated</li>
	// <li>`destroyed`: Terminated</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Pricing query parameter
	Sv []*Sv `json:"Sv,omitempty" name:"Sv"`

	// Whether to enable auto-renewal. Values:
	// <li>`0`: Default status.</li>
	// <li>`1`: Enable auto-renewal.</li>
	// <li>`2`: Disable auto-renewal.</li>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// ID of the resource associated with the plan.
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// The region. Values:
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Outside the Chinese mainland.</li>
	Area *string `json:"Area,omitempty" name:"Area"`
}

type RewriteAction struct {
	// Feature name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) API to view the requirements for entering the feature name.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Parameter
	Parameters []*RuleRewriteActionParams `json:"Parameters,omitempty" name:"Parameters"`
}

type Rule struct {
	// Feature to be executed.
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`

	// Feature execution conditions.
	// Note: If any condition in the array is met, the feature will run.
	Conditions []*RuleAndConditions `json:"Conditions,omitempty" name:"Conditions"`

	// The nested rule.
	SubRules []*SubRuleItem `json:"SubRules,omitempty" name:"SubRules"`
}

type RuleAndConditions struct {
	// Rule engine condition. This condition will be considered met if all items in the array are met.
	Conditions []*RuleCondition `json:"Conditions,omitempty" name:"Conditions"`
}

type RuleChoicePropertiesItem struct {
	// The parameter name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The parameter value type.
	// <li>CHOICE: The parameter value can be selected only from `Values`.</li>
	// <li>TOGGLE: The parameter value is of switch type and can be selected from `ChoicesValue`.</li>
	// <li>CUSTOM_NUM: The parameter value is a custom integer.</li>
	// <li>CUSTOM_STRING: The parameter value is a custom string.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Valid parameter values.
	// Note: If `Type` is `CUSTOM_NUM` or `CUSTOM_STRING`, this parameter will be an empty array.
	ChoicesValue []*string `json:"ChoicesValue,omitempty" name:"ChoicesValue"`

	// Minimum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// Maximum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// Whether multiple values can be selected or entered.
	IsMultiple *bool `json:"IsMultiple,omitempty" name:"IsMultiple"`

	// Whether the parameter can be left empty.
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitempty" name:"IsAllowEmpty"`

	// Special parameter.
	// <li>NULL: Select `NormalAction` for `RuleAction`. </li>
	// <li>If the member parameter `Id` is `Action`, select `RewirteAction` for `RuleAction`.</li>
	// <li>If the member parameter `Id` is `StatusCode`, select `CodeAction` for `RuleAction`.</li>
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitempty" name:"ExtraParameter"`
}

type RuleCodeActionParams struct {
	// The status code.
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// The parameter name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) API to view the requirements for entering the parameter name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The parameter value.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RuleCondition struct {
	// Operator. Valid values:
	// <li>`equals`: Equals</li>
	// <li>`notEquals`: Does not equal</li>
	// <li>`exist`: Exists</li>
	// <li>`notexist`: Does not exist</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// The match type. Values:
	// <li>`filename`: File name</li>
	// <li>`extension`: File extension</li>
	// <li>`host`: Host</li>
	// <li>`full_url`: Full URL, which indicates the complete URL path under the current site and must contain the HTTP protocol, host, and path.</li>
	// <li>`url`: Partial URL under the current site</li><li>`client_country`: Country/Region of the client</li>
	// <li>`query_string`: Query string in the request URL</li>
	// <li>`request_header`: HTTP request header</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// The parameter value of the match type. It can be an empty string only when `Target=query string/request header` and `Operator=exist/notexist`.
	// <li>When `Target=extension`, enter the file extension, such as "jpg" and "txt".</li>
	// <li>When `Target=filename`, enter the file name, such as "foo" in "foo.jpg".</li>
	// <li>When `Target=all`, it indicates any site request.</li>
	// <li>When `Target=host`, enter the host under the current site, such as "www.maxx55.com".</li>
	// <li>When `Target=url`, enter the partial URL path under the current site, such as "/example".</li>
	// <li>When `Target=full_url`, enter the complete URL under the current site. It must contain the HTTP protocol, host, and path, such as "https://www.maxx55.cn/example".</li>
	// <li>When `Target=client_country`, enter the ISO-3166 country/region code.</li>
	// <li>When `Target=query_string`, enter the value of the query string, such as "cn" and "1" in "lang=cn&version=1".</li>
	// <li>When `Target=request_header`, enter the HTTP request header value, such as "zh-CN,zh;q=0.9" in the "Accept-Language:zh-CN,zh;q=0.9" header.</li>
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Whether the parameter value is case insensitive. Default value: false.
	IgnoreCase *bool `json:"IgnoreCase,omitempty" name:"IgnoreCase"`

	// The parameter name of the match type. This field is required only when `Target=query_string/request_header`.
	// <li>`query_string`: Name of the query string, such as "lang" and "version" in "lang=cn&version=1".</li>
	// <li>`request_header`: Name of the HTTP request header, such as "Accept-Language" in the "Accept-Language:zh-CN,zh;q=0.9" header.</li>
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether the parameter name is case insensitive. Default value: `false`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IgnoreNameCase *bool `json:"IgnoreNameCase,omitempty" name:"IgnoreNameCase"`
}

type RuleExtraParameter struct {
	// Parameter name. Valid values:
	// <li>`Action`: Required parameter for HTTP header modification when `RewirteAction` is selected for `RuleAction`.</li>
	// <li>`StatusCode`: Required parameter for the status code feature when `CodeAction` is selected for `RuleAction`.</li>
	Id *string `json:"Id,omitempty" name:"Id"`

	// Parameter value type.
	// <li>`CHOICE`: The parameter value can be selected only from `Values`.</li>
	// <li>`CUSTOM_NUM`: The parameter value is a custom integer.</li>
	// <li>`CUSTOM_STRING`: The parameter value is a custom string.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Valid values.
	// Note: If the value of `Id` is `StatusCode`, values in the array are all integer values. When entering a parameter value, enter the integer value of the string.
	Choices []*string `json:"Choices,omitempty" name:"Choices"`
}

type RuleItem struct {
	// The rule ID.
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// The rule name. It is a string that can contain 1–255 characters.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`

	// The rule priority. The greater the value, the higher the priority. The minimum value is `1`.
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type RuleNormalActionParams struct {
	// Parameter name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) API to view the requirements for entering the parameter name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The parameter value.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RuleRewriteActionParams struct {
	// Feature parameter name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) API to view the requirements for entering the parameter name, which has three values:
	// <li>add: Add the HTTP header.</li>
	// <li>set: Rewrite the HTTP header.</li>
	// <li>del: Delete the HTTP header.</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter value
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type RulesProperties struct {
	// Parameter name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Minimum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// Valid parameter values.
	// Note: If `Type` is `CUSTOM_NUM` or `CUSTOM_STRING`, this parameter will be an empty array.
	ChoicesValue []*string `json:"ChoicesValue,omitempty" name:"ChoicesValue"`

	// Parameter value type.
	// <li>`CHOICE`: The parameter value can be selected only from `ChoicesValue`.</li>
	// <li>`TOGGLE`: The parameter value is of switch type and can be selected from `ChoicesValue`.</li>
	// <li>`OBJECT`: The parameter value is of object type, and `ChoiceProperties` indicates the attributes associated with the object type.</li>
	// <li>`CUSTOM_NUM`: Custom integer</li>
	// <li>`CUSTOM_STRING`: Custom string.</li>Note: If `OBJECT` is selected, refer to [Example 2. Create a rule with parameters of OBJECT type](https://tcloud4api.woa.com/document/product/1657/79382?!preview&!document=1).
	Type *string `json:"Type,omitempty" name:"Type"`

	// Maximum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// Whether multiple values can be selected or entered.
	IsMultiple *bool `json:"IsMultiple,omitempty" name:"IsMultiple"`

	// Whether the parameter can be left empty.
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitempty" name:"IsAllowEmpty"`

	// Associated configuration parameters of this parameter, which are required for API call.
	// Note: This parameter will be an empty array if no special parameters are added as optional parameters.
	ChoiceProperties []*RuleChoicePropertiesItem `json:"ChoiceProperties,omitempty" name:"ChoiceProperties"`

	// <li>NULL: No special parameters when `NormalAction` is selected for `RuleAction`.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitempty" name:"ExtraParameter"`
}

type RulesSettingAction struct {
	// Feature name. Valid values:
	// <li>Access URL rewrite (`AccessUrlRedirect`).</li>
	// <li>Origin-pull URL rewrite (`UpstreamUrlRedirect`).</li>
	// <li>Custom error page
	// (`ErrorPage`).</li>
	// <li>QUIC (`QUIC`).</li>
	// <li>WebSocket (`WebSocket`).</li>
	// <li>Video dragging (`VideoSeek`).</li>
	// <li>Token authentication (`Authentication`).</li>
	// <li>`CacheKey`: Custom cache key.</li>
	// <li>`Cache`: Node cache TTL.</li>
	// <li>`MaxAge`: Browser cache TTL.</li>
	// <li>`OfflineCache`: Offline cache.</li>
	// <li>`SmartRouting`: Smart acceleration.</li>
	// <li>`RangeOriginPull`: Range GETs.</li>
	// <li>`UpstreamHttp2`: HTTP/2 forwarding.</li>
	// <li>`HostHeader`: Host header rewrite.</li>
	// <li>`ForceRedirect`: Force HTTPS.</li>
	// <li>`OriginPullProtocol`: Origin-pull HTTPS.</li>
	// <li>`CachePrefresh`: Cache prefresh.</li>
	// <li>`Compression`: Smart compression.</li>
	// <li>`RequestHeader`: HTTP request header modification.</li>
	// <li>HTTP response header modification (`ResponseHeader`).</li>
	// <li>Status code cache TTL (`StatusCodeCache`).</li>
	// <li>`Hsts`.</li>
	// <li>`ClientIpHeader`.</li>
	// <li>`TlsVersion`.</li>
	// <li>`OcspStapling`.</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// Parameter information
	Properties []*RulesProperties `json:"Properties,omitempty" name:"Properties"`
}

type SecClientIp struct {
	// IP of the client.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Maximum QPS.
	RequestMaxQps *int64 `json:"RequestMaxQps,omitempty" name:"RequestMaxQps"`

	// Number of requests.
	RequestNum *int64 `json:"RequestNum,omitempty" name:"RequestNum"`
}

type SecEntry struct {
	// The query dimension value.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The details.
	Value []*SecEntryValue `json:"Value,omitempty" name:"Value"`
}

type SecEntryValue struct {
	// The metric name.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// The time-series data details.
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`

	// The maximum value.
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// The average value.
	Avg *float64 `json:"Avg,omitempty" name:"Avg"`

	// Sum
	Sum *float64 `json:"Sum,omitempty" name:"Sum"`
}

type SecHitRuleInfo struct {
	// The rule ID.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// The rule type.
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// Action. Values:
	// <li>`trans`: Allow;</li>
	// <li>`alg`: Algorithm challenge;</li>
	// <li>`drop`: Discard;</li>
	// <li>`ban`: Block the source IP;</li>
	// <li>`redirect`: Redirect;</li>
	// <li>`page`: Return to the specified page;</li>
	// <li>`monitor`: Observe.</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// The hit time recorded in seconds using UNIX timestamp.
	HitTime *int64 `json:"HitTime,omitempty" name:"HitTime"`

	// The number of requests.
	RequestNum *int64 `json:"RequestNum,omitempty" name:"RequestNum"`

	// The rule description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The subdomain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The bot tag. Values:
	// <li>`evil_bot`: Malicious bot</li>
	// <li>`suspect_bot`: Suspected bot</li>
	// <li>`good_bot`: Good bot</li>
	// <li>`normal`: Normal request</li>
	// <li>`none`: Uncategorized</li>
	BotLabel *string `json:"BotLabel,omitempty" name:"BotLabel"`
}

type SecRuleRelatedInfo struct {
	// The rule ID.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Action. Values:
	// <li>`trans`: Allow;</li>
	// <li>`alg`: Algorithm challenge;</li>
	// <li>`drop`: Discard;</li>
	// <li>`ban`: Block the source IP;</li>
	// <li>`redirect`: Redirect;</li>
	// <li>`page`: Return to the specified page;</li>
	// <li>`monitor`: Observe.</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// Risk level (only found in WAF logs). Values:
	// <li>`high risk`: High risk;</li>
	// <li>`middle risk`: Middle risk;</li>
	// <li>`low risk`: Low risk;</li>
	// <li>`unkonw`: Unknown.</li>
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// Rule level. Values:
	// <li>`normal`: Moderate.</li>
	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`

	// Rule description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The rule type.
	RuleTypeName *string `json:"RuleTypeName,omitempty" name:"RuleTypeName"`

	// The attack content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`
}

type SecurityConfig struct {
	// The settings of the managed rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WafConfig *WafConfig `json:"WafConfig,omitempty" name:"WafConfig"`

	// The settings of the rate limiting rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitempty" name:"RateLimitConfig"`

	// The settings of the custom rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`

	// The settings of the bot configuration. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BotConfig *BotConfig `json:"BotConfig,omitempty" name:"BotConfig"`

	// The switch setting of the layer-7 protection. If it is null, the setting that was last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitempty" name:"SwitchConfig"`

	// The settings of the basic access control rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitempty" name:"IpTableConfig"`

	// The settings of the exception rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptConfig *ExceptConfig `json:"ExceptConfig,omitempty" name:"ExceptConfig"`

	// The settings of the custom block page. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DropPageConfig *DropPageConfig `json:"DropPageConfig,omitempty" name:"DropPageConfig"`

	// Security template settings
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TemplateConfig *TemplateConfig `json:"TemplateConfig,omitempty" name:"TemplateConfig"`
}

type SecurityType struct {
	// Whether to enable the security type setting. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ServerCertInfo struct {
	// ID of the server certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Alias of the certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Type of the certificate. Values:
	// <li>`default`: Default certificate</lil>
	// <li>`upload`: Specified certificate</li>
	// <li>`managed`: Tencent Cloud-managed certificate</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Time when the certificate expires.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Time when the certificate is deployed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// Signature algorithm.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SignAlgo *string `json:"SignAlgo,omitempty" name:"SignAlgo"`

	// Domain name of the certificate.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CommonName *string `json:"CommonName,omitempty" name:"CommonName"`
}

type SingleDataRecord struct {
	// The query dimension value.
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// Value of the metric under the query dimension.
	TypeValue []*SingleTypeValue `json:"TypeValue,omitempty" name:"TypeValue"`
}

type SingleTypeValue struct {
	// The metric name.
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// The metric value.
	DetailData *int64 `json:"DetailData,omitempty" name:"DetailData"`
}

type SkipCondition struct {
	// The field type. Values:
	// <li>`header_fields`: HTTP request header</li>
	// <li>`cookie`: HTTP request cookie</li>
	// <li>`query_string`: Query string in the HTTP request URL</li>
	// <li>`uri`: HTTP request URI</li>
	// <li>`body_raw`: HTTP request body</li>
	// <li>`body_json`: JSON HTTP request body</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// The specific field. Values:
	// <li>`args`: Query parameter in the URI, such as "?name1=jack&age=12"</li>
	// <li>`path`: Partial path in the URI, such as "/path/to/resource.jpg"</li>
	// <li>`full`: Full path in the URI, such as "example.com/path/to/resource.jpg?name1=jack&age=12"</li>
	// <li>`upload_filename`: File path segment</li>
	// <li>`keys`: All keys</li>
	// <li>`values`: Values of all keys</li>
	// <li>`key_value`: Key and its value</li>
	Selector *string `json:"Selector,omitempty" name:"Selector"`

	// The match method used to match the key. Values:
	// <li>`equal`: Exact match</li>
	// <li>`wildcard`: Wildcard match (only asterisks)</li>
	MatchFromType *string `json:"MatchFromType,omitempty" name:"MatchFromType"`

	// The value that matches the key.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MatchFrom []*string `json:"MatchFrom,omitempty" name:"MatchFrom"`

	// The match method used to match the content.
	// <li>`equal`: Exact match</li>
	// <li>`wildcard`: Wildcard match (only asterisks)</li>
	MatchContentType *string `json:"MatchContentType,omitempty" name:"MatchContentType"`

	// The value that matches the content.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MatchContent []*string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type SmartRouting struct {
	// Whether to enable smart acceleration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SpeedTestingConfig struct {
	// The task type. Values:
	// <li>`1`: Page performance</li>
	// <li>`2`: File uploads</li>
	// <li>`3`: File downloads</li>
	// <li>`4`: Port performance</li>
	// <li>`5`: Network quality</li>
	// <li>`6`: Audio/Video experience</li>
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// The URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// The user agent.
	UA *string `json:"UA,omitempty" name:"UA"`

	// The network type.
	Connectivity *string `json:"Connectivity,omitempty" name:"Connectivity"`
}

type SpeedTestingDetailData struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The site performance across regions.
	DistrictStatistics []*DistrictStatistics `json:"DistrictStatistics,omitempty" name:"DistrictStatistics"`
}

type SpeedTestingInfo struct {
	// The task status. Values:
	// <li>`200`: The task completed.</li>
	// <li>`100`: The task is running.</li>
	// <li>`503`: The task failed.</li>
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// ID of the site test task.
	TestId *string `json:"TestId,omitempty" name:"TestId"`

	// The settings of the site test task.
	SpeedTestingConfig *SpeedTestingConfig `json:"SpeedTestingConfig,omitempty" name:"SpeedTestingConfig"`

	// The site test result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpeedTestingStatistics *SpeedTestingStatistics `json:"SpeedTestingStatistics,omitempty" name:"SpeedTestingStatistics"`
}

type SpeedTestingMetricData struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The origin information.
	OriginSpeedTestingInfo []*SpeedTestingInfo `json:"OriginSpeedTestingInfo,omitempty" name:"OriginSpeedTestingInfo"`

	// The EdgeOne information.
	ProxySpeedTestingInfo []*SpeedTestingInfo `json:"ProxySpeedTestingInfo,omitempty" name:"ProxySpeedTestingInfo"`

	// The site status.
	SpeedTestingStatus *SpeedTestingStatus `json:"SpeedTestingStatus,omitempty" name:"SpeedTestingStatus"`

	// The optimization suggestions.
	OptimizeAction []*OptimizeAction `json:"OptimizeAction,omitempty" name:"OptimizeAction"`

	// The EdgeOne load time, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyLoadTime *int64 `json:"ProxyLoadTime,omitempty" name:"ProxyLoadTime"`

	// The origin load time, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginLoadTime *int64 `json:"OriginLoadTime,omitempty" name:"OriginLoadTime"`
}

type SpeedTestingQuota struct {
	// The total number of site tests.
	TotalTestRuns *int64 `json:"TotalTestRuns,omitempty" name:"TotalTestRuns"`

	// The number of available site tests.
	AvailableTestRuns *int64 `json:"AvailableTestRuns,omitempty" name:"AvailableTestRuns"`
}

type SpeedTestingStatistics struct {
	// Last contentful paint, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstContentfulPaint *int64 `json:"FirstContentfulPaint,omitempty" name:"FirstContentfulPaint"`

	// Full content load, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstMeaningfulPaint *int64 `json:"FirstMeaningfulPaint,omitempty" name:"FirstMeaningfulPaint"`

	// Average download speed, in KB/s.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OverallDownloadSpeed *float64 `json:"OverallDownloadSpeed,omitempty" name:"OverallDownloadSpeed"`

	// Rendering time, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenderTime *int64 `json:"RenderTime,omitempty" name:"RenderTime"`

	// DOM content loaded, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DocumentFinishTime *int64 `json:"DocumentFinishTime,omitempty" name:"DocumentFinishTime"`

	// Average TCP connection, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TcpConnectionTime *int64 `json:"TcpConnectionTime,omitempty" name:"TcpConnectionTime"`

	// Average backend response, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResponseTime *int64 `json:"ResponseTime,omitempty" name:"ResponseTime"`

	// Average DOM content download, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileDownloadTime *int64 `json:"FileDownloadTime,omitempty" name:"FileDownloadTime"`

	// Load time, in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LoadTime *int64 `json:"LoadTime,omitempty" name:"LoadTime"`
}

type SpeedTestingStatus struct {
	// The URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Whether the URL uses HTTPS.
	Tls *bool `json:"Tls,omitempty" name:"Tls"`

	// Creation time of the task.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// The task status. Values:
	// <li>`200`: The task completed.</li>
	// <li>`100`: The task is running.</li>
	// <li>`503`: The task failed./li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// The user agent.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UA *string `json:"UA,omitempty" name:"UA"`

	// The network environment.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Connectivity *string `json:"Connectivity,omitempty" name:"Connectivity"`

	// Whether the URL is reachable. Values:
	// <li>`true`: Yes</li>
	// <li>`false`: No</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Reachable *bool `json:"Reachable,omitempty" name:"Reachable"`

	// Whether the URL connection timed out. Values:
	// <li>`true`: Yes</li>
	// <li>`false`: No</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimedOut *bool `json:"TimedOut,omitempty" name:"TimedOut"`
}

type SubRule struct {
	// The condition that determines if a feature should run.
	// Note: If any condition in the array is met, the feature will run.
	Conditions []*RuleAndConditions `json:"Conditions,omitempty" name:"Conditions"`

	// The feature to be executed.
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`
}

type SubRuleItem struct {
	// Nested rule settings
	Rules []*SubRule `json:"Rules,omitempty" name:"Rules"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type Sv struct {
	// The parameter key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The parameter value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SwitchConfig struct {
	// Whether to enable web protection. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>It does not affect DDoS and bot configuration.
	WebSwitch *string `json:"WebSwitch,omitempty" name:"WebSwitch"`
}

type Tag struct {
	// The tag key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// The tag value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {
	// ID of the task.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Status of the task.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Resource.
	Target *string `json:"Target,omitempty" name:"Target"`

	// Type of the task.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Creation time of the task.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Completion time of the task.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TemplateConfig struct {
	// The template ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
}

type TimingDataItem struct {
	// The query time recorded in seconds using UNIX timestamp.
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The value.
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TimingDataRecord struct {
	// The query dimension value.
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// Detailed time series data
	TypeValue []*TimingTypeValue `json:"TypeValue,omitempty" name:"TypeValue"`
}

type TimingTypeValue struct {
	// Sum.
	Sum *int64 `json:"Sum,omitempty" name:"Sum"`

	// The maximum value.
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// The average value.
	Avg *int64 `json:"Avg,omitempty" name:"Avg"`

	// Metric name.
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*TimingDataItem `json:"Detail,omitempty" name:"Detail"`
}

type TopDataRecord struct {
	// The query dimension value.
	TypeKey *string `json:"TypeKey,omitempty" name:"TypeKey"`

	// Top data rankings
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData"`
}

type TopDetailData struct {
	// The field name.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The field value.
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TopEntry struct {
	// The query dimension value.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The details.
	Value []*TopEntryValue `json:"Value,omitempty" name:"Value"`
}

type TopEntryValue struct {
	// The item name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The number of items.
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type UpstreamHttp2 struct {
	// Whether to enable HTTP2 origin-pull. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type VanityNameServers struct {
	// Whether to enable custom name servers. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// List of custom name servers
	Servers []*string `json:"Servers,omitempty" name:"Servers"`
}

type VanityNameServersIps struct {
	// Custom name of the name server
	Name *string `json:"Name,omitempty" name:"Name"`

	// IPv4 address of the custom name server
	IPv4 *string `json:"IPv4,omitempty" name:"IPv4"`
}

type Waf struct {
	// Whether to enable WAF. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// ID of the policy
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type WafConfig struct {
	// Whether to enable WAF configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>The configuration can be modified even when it is disabled.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The protection level. Values:
	// <li>`loose`: Loose</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`stricter`: Super strict</li>
	// <li>`custom`: Custom</li>
	Level *string `json:"Level,omitempty" name:"Level"`

	// The WAF global mode. Values:
	// <li>`block`: Block globally</li>
	// <li>`observe`: Observe globally</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The settings of the managed rule. If it is null, the settings that were last configured will be used.
	WafRule *WafRule `json:"WafRule,omitempty" name:"WafRule"`

	// The setting of the AI rule engine. If it is null, the setting that was last configured will be used.
	AiRule *AiRule `json:"AiRule,omitempty" name:"AiRule"`
}

type WafGroup struct {
	// Action to be executed. Values:
	// <li>`block`: Block</li>
	// <li>`observe: Observe</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// The protection level. Values:
	// <li>`loose`: Loose</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`stricter`: Super strict</li>
	// <li>`custom`: Custom</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Level *string `json:"Level,omitempty" name:"Level"`

	// ID of the rule type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`
}

type WafRule struct {
	// Whether to enable managed rules. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// IDs of the managed rules in the Block mode. You can obtain more details from [DescribeSecurityGroupManagedRules](https://tcloud4api.woa.com/document/product/1657/80807?!preview&!document=1).
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitempty" name:"BlockRuleIDs"`

	// IDs of the managed rules in the Observe mode. You can obtain more details from [DescribeSecurityGroupManagedRules](https://tcloud4api.woa.com/document/product/1657/80807?!preview&!document=1).
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitempty" name:"ObserveRuleIDs"`
}

type WebLogs struct {
	// The attack event ID.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// The attacker IP.
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// The attacked subdomain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The HTTP log content.
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// The country code of the attacker IP, which is defined in ISO-3166 alpha-2. For the list of country codes, see [ISO-3166](https://git.woa.com/edgeone/iso-3166/blob/master/all/all.json).
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// The attack time recorded in seconds using UNIX timestamp.
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// The request address.
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// The attack content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// The security rule information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleDetailList []*SecRuleRelatedInfo `json:"RuleDetailList,omitempty" name:"RuleDetailList"`

	// The request type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReqMethod *string `json:"ReqMethod,omitempty" name:"ReqMethod"`
}

type WebSocket struct {
	// Whether to enable WebSocket connection timeout. Values:
	// <li>`on`: The field "Timeout" can be configured.</li>
	// <li>`off`: The field "Timeout" is fixed to 15 seconds.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The timeout period in seconds. Maximum value: 120.
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type Zone struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// List of name servers used by the site
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// The list of name servers assigned by Tencent Cloud.
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// The site status. Values:
	// <li>`active`: The name server is switched.</li>
	// <li>`pending`: The name server is not switched.</li>
	// <li>`moved`: The name server is moved.</li>
	// <li>`deactivated`: The site is blocked.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The site access method. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Whether the site is disabled.
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// Whether CNAME acceleration is enabled. Values:
	// <li>`enabled`: Enabled</li>
	// <li>`disabled`: Disabled</li>
	CnameSpeedUp *string `json:"CnameSpeedUp,omitempty" name:"CnameSpeedUp"`

	// CNAME record access status. Values:
	// <li>`finished`: The site is verified.</li>
	// <li>`pending`: The site is being verified.</li>
	CnameStatus *string `json:"CnameStatus,omitempty" name:"CnameStatus"`

	// The list of resource tags.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The list of billable resources.
	Resources []*Resource `json:"Resources,omitempty" name:"Resources"`

	// The creation time of the site.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// The modification date of the site.
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// The site access region. Values:
	// <li>`global`: Global.</li>
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Outside the Chinese mainland.</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// The custom name server information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// The custom name server IP information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitempty" name:"VanityNameServersIps"`

	// Status of the proxy. Values:
	// <li>`active`: Enabled</li>
	// <li>`inactive`: Not activated</li>
	// <li>`paused`: Disabled</li>
	ActiveStatus *string `json:"ActiveStatus,omitempty" name:"ActiveStatus"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AliasZoneName *string `json:"AliasZoneName,omitempty" name:"AliasZoneName"`
}

type ZoneSetting struct {
	// Name of the site
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Site acceleration region. Values:
	// <li>`mainland`: Acceleration in the Chinese mainland.</li>
	// <li>`overseas`: Acceleration outside the Chinese mainland.</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// Node cache key configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// The QUIC access configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// The POST transport configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// Smart compression configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Cache expiration time configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`

	// Origin server configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Smart acceleration configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// Browser cache configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// The offline cache configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// WebSocket configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// Cache prefresh configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// IPv6 access configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// HTTPS acceleration configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Whether to carry the location information of the client IP during origin-pull.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Grpc *Grpc `json:"Grpc,omitempty" name:"Grpc"`
}