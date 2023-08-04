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

type AccelerateMainland struct {
	// Whether to enable Cross-MLC-border acceleration. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type AccelerateType struct {
	// Acceleration switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type AccelerationDomain struct {
	// Details of the origin.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginDetail *OriginDetail `json:"OriginDetail,omitempty" name:"OriginDetail"`

	// Creation time of the accelerated domain name.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Modification time of the accelerated domain name.
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Status of the accelerated domain name. Values:
	// <li>`online`: Activated</li>
	// <li>`process`: Being deployed</li>
	// <li>`offline`: Disabled</li>
	// <li>`forbidden`: Blocked</li>
	// <li>`init`: Pending activation</li>
	DomainStatus *string `json:"DomainStatus,omitempty" name:"DomainStatus"`

	// The CNAME address.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Ownership verification status. Values: <li>`pending`: Pending verification</li> <li>`finished`: Verified</li>	
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentificationStatus *string `json:"IdentificationStatus,omitempty" name:"IdentificationStatus"`
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
	// <li>`app_proto`: Application layer protocol</li>
	// <li>`sip_proto`: Network layer protocol</li>
	// <li>`uabot`: UA rules (only available in custom bot rules)</li>
	// <li>`idcid`: IDC rules (only available in custom bot rules)</li>
	// <li>`sipbot`: Search engine rules (only available in custom bot rules)</li>
	// <li>`portrait`: Client reputation (only available in custom bot rules)</li>
	// <li>`header_seq`: Header sequence (only available in custom bot rules)</li>
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

	// Custom managed rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Customizes []*AclUserRule `json:"Customizes,omitempty" name:"Customizes"`
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
	// Common operation. Values:
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
	// <li>`SslTlsSecureConf`</li>
	// <li>`OcspStapling`</li>
	// <li>`Http2`: HTTP/2 access</li>
	// <li>`UpstreamFollowRedirect`: Follow origin redirect</li>
	// <li>`Origin`: Origin</li>
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

type AlgDetectJS struct {
	// Method to validate client behavior.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Proof-of-work strength. Values:
	// <li>`low` (default): Low</li>
	// <li>`middle`: Medium</li>
	// <li>`high`: High</li>
	WorkLevel *string `json:"WorkLevel,omitempty" name:"WorkLevel"`

	// Implement a delay before executing JS in milliseconds. Value range: 0-1000. Default value: 500.
	ExecuteMode *int64 `json:"ExecuteMode,omitempty" name:"ExecuteMode"`

	// The period threshold for validating the result "Client JS disabled" in seconds. Value range: 5-3600. Default value: 10.
	InvalidStatTime *int64 `json:"InvalidStatTime,omitempty" name:"InvalidStatTime"`

	// The number of times for the result "Client JS disabled" occurred in the specified period. Value range: 1-100000000. Default value: 30.
	InvalidThreshold *int64 `json:"InvalidThreshold,omitempty" name:"InvalidThreshold"`

	// Client behavior validation results.
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitempty" name:"AlgDetectResults"`
}

type AlgDetectResult struct {
	// The validation result. Values:
	// <li>`invalid`: Invalid Cookie</li>
	// <li>`cookie_empty`: No Cookie/Cookie expired</li>
	// <li>`js_empty`: Client JS disabled</li>
	// <li>`low`: Low-risk session</li>
	// <li>`middle`: Medium-risk session</li>
	// <li>`high`: High-risk session</li>
	// <li>`timeout`: JS validation timed out</li>
	// <li>`not_browser`: Invalid browser</li>
	// <li>`is_bot`: Bot client</li>
	Result *string `json:"Result,omitempty" name:"Result"`

	// The action. Values:
	// <li>`drop`: Block</li>
	// <li>`monitor`: Observe</li>
	// <li>`silence`: Drop w/o response</li>
	// <li>`shortdelay`: Add short latency</li>
	// <li>`longdelay`: Add long latency</li>
	Action *string `json:"Action,omitempty" name:"Action"`
}

type AlgDetectRule struct {
	// ID of the rule.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// Name of the rule.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Whether to enable the rule.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Condition specified for the rule.
	AlgConditions []*AclCondition `json:"AlgConditions,omitempty" name:"AlgConditions"`

	// Validate Cookie when the condition is satisfied.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AlgDetectSession *AlgDetectSession `json:"AlgDetectSession,omitempty" name:"AlgDetectSession"`

	// Validate client behavior when the condition is satisfied.
	AlgDetectJS []*AlgDetectJS `json:"AlgDetectJS,omitempty" name:"AlgDetectJS"`

	// The update time, which is only used as an output parameter.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AlgDetectSession struct {
	// Method to validate Cookie.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The validation mode. Values:
	// <li>`detect`: Validate only</li>
	// <li>`update_detect` (default): Update Cookie and validate</li>
	DetectMode *string `json:"DetectMode,omitempty" name:"DetectMode"`

	// Whether to enable Cookie-based session check. The default value is `off`. Values:
	// <li>`off`: Disable</li>
	// <li>`on`: Enable</li>
	SessionAnalyzeSwitch *string `json:"SessionAnalyzeSwitch,omitempty" name:"SessionAnalyzeSwitch"`

	// The period threshold for validating the result "No Cookie/Cookie expired" in seconds. Value range: 5-3600. Default value: 10.
	InvalidStatTime *int64 `json:"InvalidStatTime,omitempty" name:"InvalidStatTime"`

	// The number of times for the result "No Cookie/Cookie expired" occurred in the specified period. Value range: 1-100000000. Default value: 300.
	InvalidThreshold *int64 `json:"InvalidThreshold,omitempty" name:"InvalidThreshold"`

	// Cookie validation results.
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitempty" name:"AlgDetectResults"`

	// Cookie-based session check results.
	SessionBehaviors []*AlgDetectResult `json:"SessionBehaviors,omitempty" name:"SessionBehaviors"`
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

	// Cross-MLC-border acceleration.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitempty" name:"AccelerateMainland"`
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

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

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

	// Settings of the custom bot rule. If it is null, the settings that were last configured will be used.
	BotUserRules []*BotUserRule `json:"BotUserRules,omitempty" name:"BotUserRules"`

	// Active bot detection rule.
	AlgDetectRule []*AlgDetectRule `json:"AlgDetectRule,omitempty" name:"AlgDetectRule"`

	// Settings of the bot managed rule. It is only used for output.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Customizes []*BotUserRule `json:"Customizes,omitempty" name:"Customizes"`
}

type BotExtendAction struct {
	// Action. Valid values: 
	// <li>`monitor`: Observe;</li>
	// <li>`alg`: JavaScript challenge;</li>
	// <li>`captcha`: Managed challenge;</li>
	// <li>`random`: Actions are executed based on the percentage specified in `ExtendActions`;</li>
	// <li>`silence`: Silence;</li>
	// <li>`shortdelay`: Add short latency;</li>
	// <li>`longdelay`: Add long latency.</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// The probability for triggering the action. Value range: 0-100.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
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

type BotUserRule struct {

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Action. Valid values: 
	// <li>`drop`: Block;</li>
	// <li>`monitor`: Observe;</li>
	// <li>`trans`: Allow;</li>
	// <li>`alg`: JavaScript challenge;</li>
	// <li>`captcha`: Managed challenge;</li>
	// <li>`random`: Random action;</li>
	// <li>`silence`: Silence;</li>
	// <li>`shortdelay`: Add short latency;</li>
	// <li>`longdelay`: Add long latency.</li>
	Action *string `json:"Action,omitempty" name:"Action"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>Default value: `on`
	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`

	// Details of the rule.
	AclConditions []*AclCondition `json:"AclConditions,omitempty" name:"AclConditions"`

	// The rule weight. Value range: 0-100.
	RulePriority *int64 `json:"RulePriority,omitempty" name:"RulePriority"`

	// The rule ID, which is only used as an output parameter.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// [Currently unavailable] Specify the random action and percentage.
	ExtendActions []*BotExtendAction `json:"ExtendActions,omitempty" name:"ExtendActions"`

	// The filter. Values:
	// <li>`sip`: Client IP</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FreqFields []*string `json:"FreqFields,omitempty" name:"FreqFields"`

	// Updated time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The statistical dimension. Values:
	// <li>`source_to_eo`: Responses from the origin server to EdgeOne</li>
	// <li>`client_to_eo`: Requests from the client to EdgeOne</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FreqScope []*string `json:"FreqScope,omitempty" name:"FreqScope"`
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
	// <li>`off`: Disable </li>
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: IgnoreCacheControl is deprecated.
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

	// Request parameter contained in `CacheKey`. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
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
type CheckCnameStatusRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names.
	RecordNames []*string `json:"RecordNames,omitempty" name:"RecordNames"`
}

type CheckCnameStatusRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names.
	RecordNames []*string `json:"RecordNames,omitempty" name:"RecordNames"`
}

func (r *CheckCnameStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCnameStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCnameStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCnameStatusResponseParams struct {
	// List of CNAME statuses.
	CnameStatus []*CnameStatus `json:"CnameStatus,omitempty" name:"CnameStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckCnameStatusResponse struct {
	*tchttp.BaseResponse
	Response *CheckCnameStatusResponseParams `json:"Response"`
}

func (r *CheckCnameStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCnameStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientIpCountry struct {
	// Whether to enable configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Name of the request header that contains the client IP region. It is valid when `Switch=on`. 
	// The default value `EO-Client-IPCountry` is used when it is not specified.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
}

type ClientIpHeader struct {
	// Whether to enable the configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Name of the request header that contains the client IP for origin-pull. 
	// The default value `X-Forwarded-IP` is used when it is not specified. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
}

type CnameStatus struct {
	// The domain name.
	RecordName *string `json:"RecordName,omitempty" name:"RecordName"`

	// The CNAME address.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// The CNAME status. Values:
	// <li>`active`: Activated</li>
	// <li>`moved`: Not activated </li>
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`
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
type CreateAccelerationDomainRequestParams struct {
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitempty" name:"OriginInfo"`
}

type CreateAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitempty" name:"OriginInfo"`
}

func (r *CreateAccelerationDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerationDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainName")
	delete(f, "OriginInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccelerationDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccelerationDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccelerationDomainResponseParams `json:"Response"`
}

func (r *CreateAccelerationDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerationDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// <li>`none`: (Default) Do not configure</li>
	// <li>`hosting`: Managed SSL certificate</li>
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
	// <li>`none`: (Default) Do not configure</li>
	// <li>`hosting`: Managed SSL certificate</li>
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
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
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

	// Ipv6 access configuration. 
	// IPv6 access is disabled if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// The rule details.
	// If this field is not specified, an application proxy rule will not be created.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`

	// Cross-MLC-border acceleration. It is disabled if this parameter is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitempty" name:"AccelerateMainland"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
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

	// Ipv6 access configuration. 
	// IPv6 access is disabled if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// The rule details.
	// If this field is not specified, an application proxy rule will not be created.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitempty" name:"ApplicationProxyRules"`

	// Cross-MLC-border acceleration. It is disabled if this parameter is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitempty" name:"AccelerateMainland"`
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
	delete(f, "AccelerateMainland")
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

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

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

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

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
	delete(f, "SessionPersistTime")
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
	// <li>`ent_global_with_bot`: Enterprise plan that supports content delivery network over the globe and bot management.</li>To get the available plan options for your account, view the output from <a href="https://intl.cloud.tencent.com/document/product/1552/80606?from_cn_redirect=1">DescribeAvailablePlans</a>.
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
	// <li>`ent_global_with_bot`: Enterprise plan that supports content delivery network over the globe and bot management.</li>To get the available plan options for your account, view the output from <a href="https://intl.cloud.tencent.com/document/product/1552/80606?from_cn_redirect=1">DescribeAvailablePlans</a>.
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

	// Resources to be pre-warmed, for example: 
	// http://www.example.com/example.txt 
	// Note: The number of submitted tasks is limited by the quota of the plan. For details, see [Billing Overview](https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1).
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

	// Resources to be pre-warmed, for example: 
	// http://www.example.com/example.txt 
	// Note: The number of submitted tasks is limited by the quota of the plan. For details, see [Billing Overview](https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1).
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

	// Type of cache purging. Values:
	// <li>`purge_url`: Purge by the URL</li>
	// <li>`purge_prefix`: Purge by the directory</li>
	// <li>`purge_host`: Purge by the hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by the cache-tag </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1).
	Type *string `json:"Type,omitempty" name:"Type"`

	// Configures how resources under the directory are purged when `Type = purge_prefix`. Values: <li>`invalidate`: Only resources updated under the directory are purged.</li><li>`delete`: All resources under the directory are purged regardless of whether they are updated. </li>Default value: `invalidate`.
	Method *string `json:"Method,omitempty" name:"Method"`

	// List of cached resources to purge. The format for input depends on the type of cache purging. See examples below for details. <li>By default, non-ASCII characters u200dare escaped based on RFC3986.</li><li>The maximum number of tasks per purging request is determined by the EdgeOne plan. See [Billing Overview (New)](https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1). </li>
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	//
	// Deprecated: EncodeUrl is deprecated.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Type of cache purging. Values:
	// <li>`purge_url`: Purge by the URL</li>
	// <li>`purge_prefix`: Purge by the directory</li>
	// <li>`purge_host`: Purge by the hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by the cache-tag </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1).
	Type *string `json:"Type,omitempty" name:"Type"`

	// Configures how resources under the directory are purged when `Type = purge_prefix`. Values: <li>`invalidate`: Only resources updated under the directory are purged.</li><li>`delete`: All resources under the directory are purged regardless of whether they are updated. </li>Default value: `invalidate`.
	Method *string `json:"Method,omitempty" name:"Method"`

	// List of cached resources to purge. The format for input depends on the type of cache purging. See examples below for details. <li>By default, non-ASCII characters u200dare escaped based on RFC3986.</li><li>The maximum number of tasks per purging request is determined by the EdgeOne plan. See [Billing Overview (New)](https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1). </li>
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
	delete(f, "Method")
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
type CreateSecurityIPGroupRequestParams struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IP group information.
	IPGroup *IPGroup `json:"IPGroup,omitempty" name:"IPGroup"`
}

type CreateSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IP group information.
	IPGroup *IPGroup `json:"IPGroup,omitempty" name:"IPGroup"`
}

func (r *CreateSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "IPGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityIPGroupResponseParams struct {
	// IP group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityIPGroupResponseParams `json:"Response"`
}

func (r *CreateSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneRequestParams struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The access mode. Values:
	// <li> `full`: Access through a name server.</li>
	// <li> `partial`: Access through a CNAME. Before using this access mode, first verify your site with the site verification API (IdentifyZone).<li>`noDomainAccess`: Access without using a domain name. If this value is passed, only the Tags field needs to be set. </li>
	// If not specified, this field uses the default value `full`.
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
	// <li> `full`: Access through a name server.</li>
	// <li> `partial`: Access through a CNAME. Before using this access mode, first verify your site with the site verification API (IdentifyZone).<li>`noDomainAccess`: Access without using a domain name. If this value is passed, only the Tags field needs to be set. </li>
	// If not specified, this field uses the default value `full`.
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

type DDoSAttackEvent struct {
	// The event ID.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// The attack type.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// The attack status.
	AttackStatus *int64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// The maximum attack bandwidth.
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// The peak attack packet rate.
	AttackPacketMaxRate *int64 `json:"AttackPacketMaxRate,omitempty" name:"AttackPacketMaxRate"`

	// The attack start time recorded in seconds.
	AttackStartTime *int64 `json:"AttackStartTime,omitempty" name:"AttackStartTime"`

	// The attack end time recorded in seconds.
	AttackEndTime *int64 `json:"AttackEndTime,omitempty" name:"AttackEndTime"`

	// The DDoS policy ID. 
	// Note: This field may return `null`, indicating that no valid value was found.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The site ID. 
	// Note: This field may return `null`, indicating that no valid value was found.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Geolocation scope. Values: 
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// Note: This field may return `null`, indicating that no valid value was found.
	Area *string `json:"Area,omitempty" name:"Area"`

	// The blocking time of a DDoS attack. 
	// Note: This field may return `null`, indicating that no valid value was found.
	DDoSBlockData []*DDoSBlockData `json:"DDoSBlockData,omitempty" name:"DDoSBlockData"`
}

type DDoSBlockData struct {
	// The start time recorded in UNIX timestamp.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// The end time recorded in UNIX timestamp. `0` indicates the blocking is ongoing.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// The regions blocked.
	BlockArea *string `json:"BlockArea,omitempty" name:"BlockArea"`
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
type DeleteAccelerationDomainsRequestParams struct {
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be deleted.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`

	// Whether to forcibly delete a domain name if it is associated with resources (such as alias domain names and traffic scheduling policies). 
	// <li>`true`: Delete the domain name and all associated resources.</li>
	// <li>`false`: Do not delete the domain name and all associated resources.</li>If it’s not specified, the default value `false` is used.
	Force *bool `json:"Force,omitempty" name:"Force"`
}

type DeleteAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be deleted.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`

	// Whether to forcibly delete a domain name if it is associated with resources (such as alias domain names and traffic scheduling policies). 
	// <li>`true`: Delete the domain name and all associated resources.</li>
	// <li>`false`: Do not delete the domain name and all associated resources.</li>If it’s not specified, the default value `false` is used.
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteAccelerationDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerationDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainNames")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccelerationDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccelerationDomainsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccelerationDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccelerationDomainsResponseParams `json:"Response"`
}

func (r *DeleteAccelerationDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerationDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type DeleteSecurityIPGroupRequestParams struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IP group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

type DeleteSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IP group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityIPGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityIPGroupResponseParams `json:"Response"`
}

func (r *DeleteSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityIPGroupResponse) FromJsonString(s string) error {
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
type DescribeAccelerationDomainsRequestParams struct {
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filters. Each filter can have up to 20 entries. See below for details: 
	// <li>`domain-name`:<br>   <strong>Accelerated domain name</strong><br>   Type: String<br>Required: No 
	// <li>`origin-type`:<br>   <strong>Type of the origin</strong><br>   Type: String<br>   Required: No 
	// <li>`origin`:<br>   <strong>Primary origin</strong><br>   Type: String<br>   Required: No 
	// <li>`backup-origin`:<br>   <strong>Secondary origin</strong><br>   Type: String<br>   Required: No 
	// <li>`domain-cname`:<br>   <strong>Accelerated CNAME</strong><br>   Type: String<br>   Required: No 
	// <li>`share-cname`:<br>   <strong> Shared CNAME</strong><br>   Type: String<br>   Required: No
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// The sorting order. Values:
	// <li>`asc`: Ascending order.</li>
	// <li>`desc`: Descending order.</li>Default value: `asc`.
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// The match mode. Values:
	// <li>`all`: Return all matches.</li>
	// <li>`any`: Return any match.</li>Default value: `all`.
	Match *string `json:"Match,omitempty" name:"Match"`

	// Limit on paginated queries. Default value: 20. Maximum value: 200.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The sorting criteria. Values:
	// <li>`created_on`: Creation time of the accelerated domain name.</li>
	// <li>`domain-name`: Acceleration domain name.</li>
	// </li>Default value: `domain-name`.
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filters. Each filter can have up to 20 entries. See below for details: 
	// <li>`domain-name`:<br>   <strong>Accelerated domain name</strong><br>   Type: String<br>Required: No 
	// <li>`origin-type`:<br>   <strong>Type of the origin</strong><br>   Type: String<br>   Required: No 
	// <li>`origin`:<br>   <strong>Primary origin</strong><br>   Type: String<br>   Required: No 
	// <li>`backup-origin`:<br>   <strong>Secondary origin</strong><br>   Type: String<br>   Required: No 
	// <li>`domain-cname`:<br>   <strong>Accelerated CNAME</strong><br>   Type: String<br>   Required: No 
	// <li>`share-cname`:<br>   <strong> Shared CNAME</strong><br>   Type: String<br>   Required: No
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// The sorting order. Values:
	// <li>`asc`: Ascending order.</li>
	// <li>`desc`: Descending order.</li>Default value: `asc`.
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// The match mode. Values:
	// <li>`all`: Return all matches.</li>
	// <li>`any`: Return any match.</li>Default value: `all`.
	Match *string `json:"Match,omitempty" name:"Match"`

	// Limit on paginated queries. Default value: 20. Maximum value: 200.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The sorting criteria. Values:
	// <li>`created_on`: Creation time of the accelerated domain name.</li>
	// <li>`domain-name`: Acceleration domain name.</li>
	// </li>Default value: `domain-name`.
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAccelerationDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerationDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Direction")
	delete(f, "Match")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerationDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerationDomainsResponseParams struct {
	// Total number of matched accelerated domain names.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of accelerated domain names.
	AccelerationDomains []*AccelerationDomain `json:"AccelerationDomains,omitempty" name:"AccelerationDomains"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccelerationDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccelerationDomainsResponseParams `json:"Response"`
}

func (r *DescribeAccelerationDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerationDomainsResponse) FromJsonString(s string) error {
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
type DescribeDDoSAttackDataRequestParams struct {
	// Start time of the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical metrics.
	// <li>`ddos_attackMaxBandwidth`: Peak attack bandwidth;</li>
	// <li>`ddos_attackMaxPackageRate`: Peak attack packet rate;</li>
	// <li>`ddos_attackBandwidth`: Time-series data of attack bandwidth;</li>
	// <li>`ddos_attackPackageRate`: Time-series data of attack packet rate.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// IDs of DDoS policies to be queried. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day</li>If this field is not specified, the granularity is determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < Period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global </li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical metrics.
	// <li>`ddos_attackMaxBandwidth`: Peak attack bandwidth;</li>
	// <li>`ddos_attackMaxPackageRate`: Peak attack packet rate;</li>
	// <li>`ddos_attackBandwidth`: Time-series data of attack bandwidth;</li>
	// <li>`ddos_attackPackageRate`: Time-series data of attack packet rate.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// IDs of DDoS policies to be queried. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day</li>If this field is not specified, the granularity is determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < Period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global </li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataResponseParams struct {
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of DDoS attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackDataResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventRequestParams struct {
	// Start time of the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of DDoS policy IDs. All policies are selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// (Required) List of sites. No query results are returned if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to display the details.
	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`

	// Geolocation scope. Values: 
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`

	// The sorting field. Values: 
	// <li>`MaxBandWidth`: Peak bandwidth</li>
	// <li>`AttackStartTime` Start time of the attack</li>If this field is not specified, the default value `AttackStartTime` is used.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// The sorting method. Values: 
	// <Li>`asc`: Ascending</li>
	// <li>`desc`: Descending</li>If this field is not specified, the default value `desc` is used.
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
}

type DescribeDDoSAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query period.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of DDoS policy IDs. All policies are selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// (Required) List of sites. No query results are returned if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to display the details.
	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`

	// Geolocation scope. Values: 
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitempty" name:"Area"`

	// The sorting field. Values: 
	// <li>`MaxBandWidth`: Peak bandwidth</li>
	// <li>`AttackStartTime` Start time of the attack</li>If this field is not specified, the default value `AttackStartTime` is used.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// The sorting method. Values: 
	// <Li>`asc`: Ascending</li>
	// <li>`desc`: Descending</li>If this field is not specified, the default value `desc` is used.
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
}

func (r *DescribeDDoSAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ShowDetail")
	delete(f, "Area")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventResponseParams struct {
	// List of DDoS attack data. 
	// Note: This field may return `null`, indicating that no valid value was found.
	Data []*DDoSAttackEvent `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventResponse) FromJsonString(s string) error {
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
type DescribeOriginGroupRequestParams struct {
	// Offset for paginated queries. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Value range: 1-1000. Default value: 10.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filters. Each filter can have up to 20 entries. See below for details:
	// <li>`zone-id`<br>   Filter by the specified <strong>site ID</strong>, such as zone-20hzkd4rdmy0<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`origin-group-id`:<br>   Filter by the specified <strong>origin group ID</strong>, such as origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`origin-group-name`:<br>   Filter by the specified <strong>origin group name</strong><br>   Type: String<br>   Required: No<br>   Fuzzy query: Supported (only one origin group name allowed in a query)</li>
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Offset for paginated queries. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Value range: 1-1000. Default value: 10.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filters. Each filter can have up to 20 entries. See below for details:
	// <li>`zone-id`<br>   Filter by the specified <strong>site ID</strong>, such as zone-20hzkd4rdmy0<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`origin-group-id`:<br>   Filter by the specified <strong>origin group ID</strong>, such as origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported</li><li>`origin-group-name`:<br>   Filter by the specified <strong>origin group name</strong><br>   Type: String<br>   Required: No<br>   Fuzzy query: Supported (only one origin group name allowed in a query)</li>
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
type DescribeOriginProtectionRequestParams struct {
	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`need-update`:<br>   Whether <strong>the intermediate IP update is needed for the site</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Update needed.<br>   `false`: Update not needed.<br></li>
	// <li>`plan-support`:<br>   Whether <strong>origin protection is supported in the plan</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Origin protection supported.<br>   `false`: Origin protection not supported.<br></li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeOriginProtectionRequest struct {
	*tchttp.BaseRequest
	
	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`need-update`:<br>   Whether <strong>the intermediate IP update is needed for the site</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Update needed.<br>   `false`: Update not needed.<br></li>
	// <li>`plan-support`:<br>   Whether <strong>origin protection is supported in the plan</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Origin protection supported.<br>   `false`: Origin protection not supported.<br></li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOriginProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginProtectionResponseParams struct {
	// Origin protection configuration.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OriginProtectionInfo []*OriginProtectionInfo `json:"OriginProtectionInfo,omitempty" name:"OriginProtectionInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOriginProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginProtectionResponseParams `json:"Response"`
}

func (r *DescribeOriginProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The metric to query. Values:
	// <li>`l7Flow_outFlux`: Traffic used for EdegOne responses</li>
	// <li>`l7Flow_inFlux`: Traffic used for EdegOne requests</li>
	// <li>`l7Flow_outBandwidth`: Bandwidth used for EdegOne responses</li>
	// <li>`l7Flow_inBandwidth`: Bandwidth used for EdegOne requests</li>
	// <li>`l7Flow_hit_outFlux`: Traffic used for cache hit</li>
	// <li>`l7Flow_request`: Access requests</li>
	// <li>`l7Flow_flux`: Upstream and downstream traffic used for client access</li>
	// <li>`l7Flow_bandwidth`: Upstream and downstream bandwidth used for client access</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it’s not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`:   All protocols. </li>If it’s not specified, `all` is used. This parameter is not yet available now.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: One day</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>`socket`:<br>u2003u2003 Filter by the specified <strong>HTTP protocol type</strong><br>u2003u2003 Values:<br>u2003u2003 `HTTP`: HTTP protocol;<br>u2003u2003 `HTTPS`: HTTPS protocol;<br>u2003u2003 `QUIC`: QUIC protocol.</li>
	// <li>`tagKey`:<br>u2003u2003 Filter by the specified <strong>tag key</strong></li>
	// <li>`tagValue`<br>u2003u2003 Filter by the specified <strong>tag value</strong></li>
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

	// The metric to query. Values:
	// <li>`l7Flow_outFlux`: Traffic used for EdegOne responses</li>
	// <li>`l7Flow_inFlux`: Traffic used for EdegOne requests</li>
	// <li>`l7Flow_outBandwidth`: Bandwidth used for EdegOne responses</li>
	// <li>`l7Flow_inBandwidth`: Bandwidth used for EdegOne requests</li>
	// <li>`l7Flow_hit_outFlux`: Traffic used for cache hit</li>
	// <li>`l7Flow_request`: Access requests</li>
	// <li>`l7Flow_flux`: Upstream and downstream traffic used for client access</li>
	// <li>`l7Flow_bandwidth`: Upstream and downstream bandwidth used for client access</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it’s not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`:   All protocols. </li>If it’s not specified, `all` is used. This parameter is not yet available now.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: One day</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>`socket`:<br>u2003u2003 Filter by the specified <strong>HTTP protocol type</strong><br>u2003u2003 Values:<br>u2003u2003 `HTTP`: HTTP protocol;<br>u2003u2003 `HTTPS`: HTTPS protocol;<br>u2003u2003 `QUIC`: QUIC protocol.</li>
	// <li>`tagKey`:<br>u2003u2003 Filter by the specified <strong>tag key</strong></li>
	// <li>`tagValue`<br>u2003u2003 Filter by the specified <strong>tag value</strong></li>
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

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`rule-id`:<br>   Filter by the <strong>rule ID</strong><br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`rule-id`:<br>   Filter by the <strong>rule ID</strong><br>   Type: String<br>   Required: No</li>
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
	// If it’s not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>ruleId<br>   Filter by the specified <strong>forwarding rule ID</strong></li>
	// <li>proxyId<br>   Filter by the specified <strong>L4 agent ID</strong></li>
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
	// If it’s not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>ruleId<br>   Filter by the specified <strong>forwarding rule ID</strong></li>
	// <li>proxyId<br>   Filter by the specified <strong>L4 agent ID</strong></li>
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

	// The metric to query. Values:
	// <li>`l7Flow_outFlux`: Traffic used for EdgeOne responses</li>
	// <li>`l7Flow_inFlux`: Traffic used for EdgeOne requests</li>
	// <li>`l7Flow_outBandwidth`: Bandwidth used for EdgeOne responses</li>
	// <li>`l7Flow_inBandwidth`: Bandwidth used for EdgeOne requests</li>
	// <li>`l7Flow_request`: Access requests</li>
	// <li>`l7Flow_flux`: Upstream and downstream traffic used for client access</li>
	// <li>`l7Flow_bandwidth`: Upstream and downstream bandwidth used for client access</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it’s not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`province`:<br>   Filter by the specified <strong>province name</strong>. It’s only available when `Area` is `mainland`. </li>
	// <li>`isp`:<br>   Filter by the specified <strong>ISP</strong>. It’s only available when `Area` is `mainland`.<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others.</li>
	// <li>`domain`:<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>   Filter by the specified <strong>URL Path</strong> (such as `/content` or `content/test.jpg`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`referer`:<br>   Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`resourceType`<br>   Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol version</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0;<br>   `HTTP/1.1`: HTTP 1.1;<br>   `HTTP/2.0`: HTTP 2.0;<br>   `HTTP/3.0`: HTTP 3.0;<br>   `WebSocket`: WebSocket.</li>
	// <li>`socket`:<br>   Filter by the specified <strong>HTTP protocol type</strong><br>   Values:<br>   `HTTP`: HTTP protocol;<br>   `HTTPS`: HTTPS protocol;<br>   `QUIC`: QUIC protocol.</li>
	// <li>`statusCode`:<br>   Filter by the specified <strong> status code</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>  `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`:<br>   Filter by the specified <strong>browser type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>   Filter by the <strong>device type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>   Filter by the <strong>operating system</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Mac OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>   Filter by the <strong>TLS version</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`:<br>   Filter by the specified <strong>IP version</strong>.<br>   Values:<br>   `4`: IPv4;<br>   `6`: IPv6.</li>
	// <li>`tagKey`:<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>`tagValue`<br>   Filter by the specified <strong>tag value</strong></li>
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

	// The metric to query. Values:
	// <li>`l7Flow_outFlux`: Traffic used for EdgeOne responses</li>
	// <li>`l7Flow_inFlux`: Traffic used for EdgeOne requests</li>
	// <li>`l7Flow_outBandwidth`: Bandwidth used for EdgeOne responses</li>
	// <li>`l7Flow_inBandwidth`: Bandwidth used for EdgeOne requests</li>
	// <li>`l7Flow_request`: Access requests</li>
	// <li>`l7Flow_flux`: Upstream and downstream traffic used for client access</li>
	// <li>`l7Flow_bandwidth`: Upstream and downstream bandwidth used for client access</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites
	// Enter the IDs of sites to query. The maximum query period is determined by the <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> of the bound plan. If it’s not specified, all sites are selected by default, and the query period must be within the last 30 days. 
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`province`:<br>   Filter by the specified <strong>province name</strong>. It’s only available when `Area` is `mainland`. </li>
	// <li>`isp`:<br>   Filter by the specified <strong>ISP</strong>. It’s only available when `Area` is `mainland`.<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others.</li>
	// <li>`domain`:<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>   Filter by the specified <strong>URL Path</strong> (such as `/content` or `content/test.jpg`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`referer`:<br>   Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`resourceType`<br>   Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol version</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0;<br>   `HTTP/1.1`: HTTP 1.1;<br>   `HTTP/2.0`: HTTP 2.0;<br>   `HTTP/3.0`: HTTP 3.0;<br>   `WebSocket`: WebSocket.</li>
	// <li>`socket`:<br>   Filter by the specified <strong>HTTP protocol type</strong><br>   Values:<br>   `HTTP`: HTTP protocol;<br>   `HTTPS`: HTTPS protocol;<br>   `QUIC`: QUIC protocol.</li>
	// <li>`statusCode`:<br>   Filter by the specified <strong> status code</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>  `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`:<br>   Filter by the specified <strong>browser type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>   Filter by the <strong>device type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>   Filter by the <strong>operating system</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Mac OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>   Filter by the <strong>TLS version</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`:<br>   Filter by the specified <strong>IP version</strong>.<br>   Values:<br>   `4`: IPv4;<br>   `6`: IPv6.</li>
	// <li>`tagKey`:<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>`tagValue`<br>   Filter by the specified <strong>tag value</strong></li>
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

	// Metrics to query. Valid values: 
	// <li>`l7Flow_outFlux_country`: Query traffic by country/region;</li>
	// <li>`l7Flow_outFlux_statusCode`: Query traffic by status code;</li>
	// <li>`l7Flow_outFlux_domain`: Query traffic by domain;</li>
	// <li>`l7Flow_outFlux_url`: Query traffic by URL;</li>
	// <li>`l7Flow_outFlux_resourceType`: Query traffic by resource type;</li>
	// <li>`l7Flow_outFlux_sip`: Query traffic by source IP;</li>
	// <li>`l7Flow_outFlux_referers`: Query traffic by refer information;</li>
	// <li>`l7Flow_outFlux_ua_device`: Query traffic by device;</li>
	// <li>`l7Flow_outFlux_ua_browser`: Query traffic by browser;</li>
	// <li>`l7Flow_outFlux_us_os`: Query traffic by OS;</li>
	// <li>`l7Flow_request_country`: Query requests by country/region;</li>
	// <li>`l7Flow_request_statusCode`: Query requests by status code;</li>
	// <li>`l7Flow_request_domain`: Query requests by domain;</li>
	// <li>`l7Flow_request_url`: Query requests by URL;</li>
	// <li>`l7Flow_request_resourceType`: Query requests by resource type;</li>
	// <li>`l7Flow_request_sip`: Query requests by source IP;</li>
	// <li>`l7Flow_request_referer`: Query requests by refer information;</li>
	// <li>`l7Flow_request_ua_device`: Query requests by device;</li>
	// <li>`l7Flow_request_ua_browser`: Query requests by browser;</li>
	// <li>`l7Flow_request_us_os`: Query requests by OS.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// (Required) List of sites. No query results are returned if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Queries the top n rows of data. Maximum value: 1000. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`province`:<br>   Filter by the specified <strong>province name</strong>. It’s only available when `Area` is `mainland`.</li>
	// <li>`isp`:<br>   Filter by the specified <strong>ISP</strong>. It’s only available when `Area` is `mainland`.<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others</li>
	// <li>`domain`:<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>   Filter by the specified <strong>URL Path</strong> (such as `/content` or `content/test.jpg`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`referer`:<br>   Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`resourceType`:<br>   Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol version</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0;<br>   `HTTP/1.1`: HTTP 1.1;<br>   `HTTP/2.0`: HTTP 2.0;<br>   `HTTP/3.0`: HTTP 3.0;<br>   `WebSocket`: WebSocket.</li>
	// <li>`socket`<br>   Filter by the specified <strong>HTTP protocol type</strong><br>   Values:<br>   `HTTP`: HTTP protocol;<br>   `HTTPS`: HTTPS protocol;<br>   `QUIC`: QUIC protocol.</li>
	// <li>`statusCode`:<br>   Filter by the specified <strong> status code</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>  In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>  `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`:<br>   Filter by the specified <strong>browser type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>   Filter by the <strong>device type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>   Filter by the <strong>operating system</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Mac OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>   Filter by the <strong>TLS version</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`:<br>   Filter by the specified <strong>IP version</strong>.<br>   Values:<br>   `4`: IPv4;<br>   `6`: IPv6.</li>
	// <li>`tagKey`:<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>`tagValue`:<br>   Filter by the specified <strong>tag value</strong></li>
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

	// Metrics to query. Valid values: 
	// <li>`l7Flow_outFlux_country`: Query traffic by country/region;</li>
	// <li>`l7Flow_outFlux_statusCode`: Query traffic by status code;</li>
	// <li>`l7Flow_outFlux_domain`: Query traffic by domain;</li>
	// <li>`l7Flow_outFlux_url`: Query traffic by URL;</li>
	// <li>`l7Flow_outFlux_resourceType`: Query traffic by resource type;</li>
	// <li>`l7Flow_outFlux_sip`: Query traffic by source IP;</li>
	// <li>`l7Flow_outFlux_referers`: Query traffic by refer information;</li>
	// <li>`l7Flow_outFlux_ua_device`: Query traffic by device;</li>
	// <li>`l7Flow_outFlux_ua_browser`: Query traffic by browser;</li>
	// <li>`l7Flow_outFlux_us_os`: Query traffic by OS;</li>
	// <li>`l7Flow_request_country`: Query requests by country/region;</li>
	// <li>`l7Flow_request_statusCode`: Query requests by status code;</li>
	// <li>`l7Flow_request_domain`: Query requests by domain;</li>
	// <li>`l7Flow_request_url`: Query requests by URL;</li>
	// <li>`l7Flow_request_resourceType`: Query requests by resource type;</li>
	// <li>`l7Flow_request_sip`: Query requests by source IP;</li>
	// <li>`l7Flow_request_referer`: Query requests by refer information;</li>
	// <li>`l7Flow_request_ua_device`: Query requests by device;</li>
	// <li>`l7Flow_request_ua_browser`: Query requests by browser;</li>
	// <li>`l7Flow_request_us_os`: Query requests by OS.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// (Required) List of sites. No query results are returned if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Queries the top n rows of data. Maximum value: 1000. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filters
	// <li>`country`:<br>   Filter by the specified <strong>country code</strong>. <a href="https://en.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a> country codes are used.</li>
	// <li>`province`:<br>   Filter by the specified <strong>province name</strong>. It’s only available when `Area` is `mainland`.</li>
	// <li>`isp`:<br>   Filter by the specified <strong>ISP</strong>. It’s only available when `Area` is `mainland`.<br>   Values: <br>   `2`: CTCC; <br>   `26`: CUCC;<br>   `1046`: CMCC;<br>   `3947`: CTT; <br>   `38`: CERNET; <br>   `43`: GWBN;<br>   `0`: Others</li>
	// <li>`domain`:<br>   Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>   Filter by the specified <strong>URL Path</strong> (such as `/content` or `content/test.jpg`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`referer`:<br>   Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`resourceType`:<br>   Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.</li>
	// <li>`protocol`:<br>   Filter by the specified <strong>HTTP protocol version</strong><br>   Values:<br>   `HTTP/1.0`: HTTP 1.0;<br>   `HTTP/1.1`: HTTP 1.1;<br>   `HTTP/2.0`: HTTP 2.0;<br>   `HTTP/3.0`: HTTP 3.0;<br>   `WebSocket`: WebSocket.</li>
	// <li>`socket`<br>   Filter by the specified <strong>HTTP protocol type</strong><br>   Values:<br>   `HTTP`: HTTP protocol;<br>   `HTTPS`: HTTPS protocol;<br>   `QUIC`: QUIC protocol.</li>
	// <li>`statusCode`:<br>   Filter by the specified <strong> status code</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>  In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>  `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`browserType`:<br>   Filter by the specified <strong>browser type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>  `Firefox`: Firefox browser;<br>   `Chrome`: Chrome browser;<br>   `Safari`: Safari browser;<br>   `MicrosoftEdge`: Microsoft Edge browser;<br>   `IE`: IE browser;<br>   `Opera`: Opera browser;<br>   `QQBrowser`: QQ browser;<br>   `LBBrowser`: LB browser;<br>   `MaxthonBrowser`: Maxthon browser;<br>   `SouGouBrowser`: Sogou browser;<br>  `BIDUBrowser`: Baidu browser;<br>   `TaoBrowser`: Tao browser;<br>   `UBrowser`: UC browser;<br>   `Other`: Other browsers; <br>   `Empty`: The browser type is not specified; <br>   `Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>   Filter by the <strong>device type</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `TV`: TV; <br>   `Tablet`: Tablet;<br>   `Mobile`: Mobile phone;<br>   `Desktop`: Desktop device; <br>   `Other`: Other device;<br>   `Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>   Filter by the <strong>operating system</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values: <br>   `Linux`: Linux OS;<br>   `MacOS`: Mac OS;<br>   `Android`: Android OS;<br>   `IOS`: iOS OS;<br>   `Windows`: Windows OS;<br>   `NetBSD`: NetBSD OS;<br>   `ChromiumOS`: Chromium OS;<br>   `Bot`: Web crawler: <br>   `Other`: Other OS;<br>   `Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>   Filter by the <strong>TLS version</strong>. <br>   When this parameter is specified, the query period must be within the last 30 days. <br>   In this case, the supported <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query period</a> stated when `Zonelds` is specified become invalid.<br>   Values:<br>   `TLS1.0`: TLS 1.0; <br>   `TLS1.1`: TLS 1.1;<br>   `TLS1.2`: TLS 1.2;<br>   `TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`:<br>   Filter by the specified <strong>IP version</strong>.<br>   Values:<br>   `4`: IPv4;<br>   `6`: IPv6.</li>
	// <li>`tagKey`:<br>   Filter by the specified <strong>tag key</strong></li>
	// <li>`tagValue`:<br>   Filter by the specified <strong>tag value</strong></li>
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

	// Top rows of data to query. Maximum value: 1000. Top 10 rows of data are queried if this field is not specified.
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

	// Top rows of data to query. Maximum value: 1000. Top 10 rows of data are queried if this field is not specified.
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

type DiffIPWhitelist struct {
	// The latest intermediate IPs.
	LatestIPWhitelist *IPWhitelist `json:"LatestIPWhitelist,omitempty" name:"LatestIPWhitelist"`

	// The intermediate IPs added to the existing list.
	AddedIPWhitelist *IPWhitelist `json:"AddedIPWhitelist,omitempty" name:"AddedIPWhitelist"`

	// The intermediate IPs removed from the existing list.
	RemovedIPWhitelist *IPWhitelist `json:"RemovedIPWhitelist,omitempty" name:"RemovedIPWhitelist"`

	// The intermediate IPs that remain unchanged.
	NoChangeIPWhitelist *IPWhitelist `json:"NoChangeIPWhitelist,omitempty" name:"NoChangeIPWhitelist"`
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
	// <li>`waf`: Tencent Cloud-managed rules</li>
	// <li>`rate`: Rate limiting rules</li>
	// <li>`acl`: Custom rule</li>
	// <li>`cc`: CC attack defense</li>
	// <li>`bot`: Bot protection</li>
	// Note: this field may return `null`, indicating that no valid value is obtained.
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

type FirstPartConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The transfer period threshold of the first 8 KB. If the threshold is reached, it’s considered a slow attack. Default: `5`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatTime *uint64 `json:"StatTime,omitempty" name:"StatTime"`
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
	// Whether to enable gRPC support. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
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

	// MaxAge (in seconds). The maximum value is 1 day. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
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

	// TLS version. Valid values: 
	// <li>`TLSv1`: TLSv1 version;</li>
	// <li>`TLSV1.1`: TLSv1.1 version;</li>
	// <li>`TLSV1.2`: TLSv1.2 version;</li>
	// <li>`TLSv1.3`: TLSv1.3 version.</li>Only consecutive versions can be enabled at the same time. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
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

	// Cipher suite. Values:
	// <li>`loose-v2023`: Offer the highest compatibility but relatively lower security. It supports TLS 1.0-1.3.</li>
	// <li>`general-v2023`: Keep a balance between the compatibility and security. It supports TLS 1.2-1.3.</li>
	// <li>`strict-v2023`: Provides high security, disabling all insecure cipher suites. It supports TLS 1.2-1.3.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CipherSuite *string `json:"CipherSuite,omitempty" name:"CipherSuite"`
}

type IPGroup struct {
	// Group ID. Enter `0`.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Group name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// IP group information, including IP and IP mask.
	Content []*string `json:"Content,omitempty" name:"Content"`
}

type IPWhitelist struct {
	// List of IPv4 addresses
	IPv4 []*string `json:"IPv4,omitempty" name:"IPv4"`

	// List of IPv6 addresses
	IPv6 []*string `json:"IPv6,omitempty" name:"IPv6"`
}

type Identification struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The subdomain name to be verified. To verify the ownership of a site, leave it blank.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

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

	// A subdomain name under the site. Specify this field if you want to verify the ownership of a subdomain name. Otherwise you can leave it blank.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// A subdomain name under the site. Specify this field if you want to verify the ownership of a subdomain name. Otherwise you can leave it blank.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
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
	delete(f, "Domain")
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

type ImageOptimize struct {
	// Whether to enable configuration. Values: 
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
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

	// Matching method. It defaults to `equal` if it’s left empty.
	// Values: 
	// <li>`is_empty`: The field is empty.</li>
	// <li>`not_exists`: The configuration item does not exist.</li>
	// <li>`include`: Include</li>
	// <li>`not_include`: Do not include</li>
	// <li>`equal`: Equal to</li>
	// <li>`not_equal`: Not equal to</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// The rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// The update time, which is only used as an output parameter.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The rule status. A null value indicates that the rule is enabled. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The rule name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Matching content. It’s not required when `Operator` is `is_emty` or `not_exists`. 
	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
}

type Ipv6 struct {
	// Whether to enable IPv6 access. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
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
type ModifyAccelerationDomainRequestParams struct {
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitempty" name:"OriginInfo"`
}

type ModifyAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitempty" name:"OriginInfo"`
}

func (r *ModifyAccelerationDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainName")
	delete(f, "OriginInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccelerationDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerationDomainResponseParams `json:"Response"`
}

func (r *ModifyAccelerationDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainStatusesRequestParams struct {
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be modified.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`

	// Status of the accelerated domain name. Values:
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Whether to force suspension when the domain name has associated resources (such as alias domain names and traffic scheduling policies). Values:
	// <li>`true`: Suspend the domain name and all associated resources.</li>
	// <li>`true`: Do not suspend the domain name and all associated resources.</li>Default value: `false`.
	Force *bool `json:"Force,omitempty" name:"Force"`
}

type ModifyAccelerationDomainStatusesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be modified.
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`

	// Status of the accelerated domain name. Values:
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Whether to force suspension when the domain name has associated resources (such as alias domain names and traffic scheduling policies). Values:
	// <li>`true`: Suspend the domain name and all associated resources.</li>
	// <li>`true`: Do not suspend the domain name and all associated resources.</li>Default value: `false`.
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *ModifyAccelerationDomainStatusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainStatusesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainNames")
	delete(f, "Status")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerationDomainStatusesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainStatusesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccelerationDomainStatusesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerationDomainStatusesResponseParams `json:"Response"`
}

func (r *ModifyAccelerationDomainStatusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainStatusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// The original configuration will apply if this field is not specified.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>If not specified, this field uses the default value `instance`.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// IPv6 access configuration. The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Cross-MLC-border acceleration. The original configuration will apply if it is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitempty" name:"AccelerateMainland"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Proxy ID.
	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// The original configuration will apply if this field is not specified.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>If not specified, this field uses the default value `instance`.
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`

	// IPv6 access configuration. The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Cross-MLC-border acceleration. The original configuration will apply if it is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitempty" name:"AccelerateMainland"`
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
	delete(f, "AccelerateMainland")
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

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

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

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitempty" name:"SessionPersistTime"`

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
	delete(f, "SessionPersistTime")
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
type ModifyHostsCertificateRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names that the certificate will be attached to.
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Certificate information. Note that only `CertId` is required. If it is not specified, the default certificate will be used.
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitempty" name:"ServerCertInfo"`

	// Whether the certificate is managed by EdgeOne. Values:
	// <li>`apply`: Managed by EdgeOne.</li>
	// <li>`none`: Not managed by EdgeOne.</li>If not specified, this field uses the default value `none`.
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
	// <li>`apply`: Managed by EdgeOne.</li>
	// <li>`none`: Not managed by EdgeOne.</li>If not specified, this field uses the default value `none`.
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
type ModifySecurityIPGroupRequestParams struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IP group configuration.
	IPGroup *IPGroup `json:"IPGroup,omitempty" name:"IPGroup"`

	// Operation type. Valid values: 
	// <li>`append`: Add information of `Content` to `IPGroup`;</li>
	// <li>`remove`: Delete information of `Content` from `IPGroup`;</li>
	// <li>`update`: Replace all information of `IPGroup` and modify the IPGroup name.</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ModifySecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IP group configuration.
	IPGroup *IPGroup `json:"IPGroup,omitempty" name:"IPGroup"`

	// Operation type. Valid values: 
	// <li>`append`: Add information of `Content` to `IPGroup`;</li>
	// <li>`remove`: Delete information of `Content` from `IPGroup`;</li>
	// <li>`update`: Replace all information of `IPGroup` and modify the IPGroup name.</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifySecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "IPGroup")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityIPGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityIPGroupResponseParams `json:"Response"`
}

func (r *ModifySecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityIPGroupResponse) FromJsonString(s string) error {
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
type ModifyZoneRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Access mode of the site. Values:
	// <li> `full`: Access through a name server.</li>
	// <li> `partial`: Access through a CNAME u200drecord. A site using domainless access can only switch to CNAME access. </li>The original configuration applies if this field is not specified.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The custom name servers. The original configuration applies if this field is not specified. It is not allowed to pass this field when a site is connected without using a domain name.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitempty" name:"AliasZoneName"`

	// The region where the site requests access. Values:
	// <li> `global`: Global coverage</li>
	// <li> `mainland`: Chinese mainland</li>
	// <li> `overseas`: Outside the Chinese mainland </li>It is not allowed to pass this field when a site is connected without using a domain name.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Name of the site. This field takes effect only when the site switches from domainless access to CNAME access.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Access mode of the site. Values:
	// <li> `full`: Access through a name server.</li>
	// <li> `partial`: Access through a CNAME u200drecord. A site using domainless access can only switch to CNAME access. </li>The original configuration applies if this field is not specified.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The custom name servers. The original configuration applies if this field is not specified. It is not allowed to pass this field when a site is connected without using a domain name.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitempty" name:"AliasZoneName"`

	// The region where the site requests access. Values:
	// <li> `global`: Global coverage</li>
	// <li> `mainland`: Chinese mainland</li>
	// <li> `overseas`: Outside the Chinese mainland </li>It is not allowed to pass this field when a site is connected without using a domain name.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Name of the site. This field takes effect only when the site switches from domainless access to CNAME access.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
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
	delete(f, "Area")
	delete(f, "ZoneName")
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
	// Site ID to modify.
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

	// QUIC access configuration. 
	// The original configuration will apply if it is not specified.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// POST transport configuration. 
	// The original configuration will apply if it is not specified.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// The smart compression configuration.
	// The original configuration will apply if this field is not specified.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration. 
	// The original configuration will apply if it is not specified.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration. 
	// The original configuration will apply if it is not specified.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS acceleration configuration. 
	// The original configuration will apply if it is not specified.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// The origin server configuration.
	// The original configuration will apply if this field is not specified.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// The smart acceleration configuration.
	// The original configuration will apply if this field is not specified.
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// WebSocket configuration. 
	// The original configuration will apply if it is not specified.
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration. 
	// The original configuration will apply if it is not specified.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// The cache prefresh configuration.
	// The original configuration will apply if this field is not specified.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6 access configuration. 
	// The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull. 
	// The original configuration will apply if it is not specified.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support. 
	// The original configuration will apply if it is not specified.
	Grpc *Grpc `json:"Grpc,omitempty" name:"Grpc"`

	// Image optimization. 
	// It is disabled if this parameter is not specified.
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitempty" name:"ImageOptimize"`

	// Standard debugging configuration.
	StandardDebug *StandardDebug `json:"StandardDebug,omitempty" name:"StandardDebug"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID to modify.
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

	// QUIC access configuration. 
	// The original configuration will apply if it is not specified.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// POST transport configuration. 
	// The original configuration will apply if it is not specified.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// The smart compression configuration.
	// The original configuration will apply if this field is not specified.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration. 
	// The original configuration will apply if it is not specified.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration. 
	// The original configuration will apply if it is not specified.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS acceleration configuration. 
	// The original configuration will apply if it is not specified.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// The origin server configuration.
	// The original configuration will apply if this field is not specified.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// The smart acceleration configuration.
	// The original configuration will apply if this field is not specified.
	SmartRouting *SmartRouting `json:"SmartRouting,omitempty" name:"SmartRouting"`

	// WebSocket configuration. 
	// The original configuration will apply if it is not specified.
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration. 
	// The original configuration will apply if it is not specified.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitempty" name:"ClientIpHeader"`

	// The cache prefresh configuration.
	// The original configuration will apply if this field is not specified.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitempty" name:"CachePrefresh"`

	// Ipv6 access configuration. 
	// The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull. 
	// The original configuration will apply if it is not specified.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support. 
	// The original configuration will apply if it is not specified.
	Grpc *Grpc `json:"Grpc,omitempty" name:"Grpc"`

	// Image optimization. 
	// It is disabled if this parameter is not specified.
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitempty" name:"ImageOptimize"`

	// Standard debugging configuration.
	StandardDebug *StandardDebug `json:"StandardDebug,omitempty" name:"StandardDebug"`
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
	delete(f, "ImageOptimize")
	delete(f, "StandardDebug")
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

	// Whether to allow private access to buckets when `OriginType=cos`. Valid values: 
	// <li>`on`: Private access;</li>
	// <li>`off`: Public access.</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CosPrivateAccess *string `json:"CosPrivateAccess,omitempty" name:"CosPrivateAccess"`
}

type OriginDetail struct {
	// The origin type. Values:
	// <li>`IP_DOMAIN`: IPv4/IPv6 address or domain name</li>
	// <li>`COS`: COS bucket address</li>
	// <li>`ORIGIN_GROUP`: Origin group</li>
	// <li>`AWS_S3`: AWS S3 bucket address</li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The origin address. Enter the origin group ID if `OriginType=ORIGIN_GROUP`.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// ID of the secondary origin group (valid when `OriginType=ORIGIN_GROUP`). If it’s not specified, it indicates that secondary origins are not used.
	BackupOrigin *string `json:"BackupOrigin,omitempty" name:"BackupOrigin"`

	// Name of the primary origin group (valid when `OriginType=ORIGIN_GROUP`).
	OriginGroupName *string `json:"OriginGroupName,omitempty" name:"OriginGroupName"`

	// Name of the secondary origin group (valid when `OriginType=ORIGIN_GROUP` and `BackupOrigin` is specified).
	BackOriginGroupName *string `json:"BackOriginGroupName,omitempty" name:"BackOriginGroupName"`

	// Whether to authenticate access to the private object storage origin (valid when `OriginType=COS/AWS_S3`). Values:
	// <li>`on`: Enable private authentication.</li>
	// <li>`off`: Disable private authentication.</li>
	// If this field is not specified, the default value `off` is used.
	PrivateAccess *string `json:"PrivateAccess,omitempty" name:"PrivateAccess"`

	// The private authentication parameters. This field is valid when `PrivateAccess=on`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitempty" name:"PrivateParameters"`
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

type OriginInfo struct {
	// The origin type. Values:
	// <li>`IP_DOMAIN`: IPv4/IPv6 address or domain name</li>
	// <li>`COS`: COS bucket address </li>
	// <li>`ORIGIN_GROUP`: Origin group </li>
	// <li>`AWS_S3`: AWS S3 bucket address </li>
	// <li>`SPACE`: EdgeOne Shield Space </li>
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// The origin address. Enter the origin group ID if `OriginType=ORIGIN_GROUP`.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// ID of the secondary origin group (valid when `OriginType=ORIGIN_GROUP`). If it’s not specified, it indicates that secondary origins are not used.
	BackupOrigin *string `json:"BackupOrigin,omitempty" name:"BackupOrigin"`

	// Whether to authenticate access to the private object storage origin (valid when `OriginType=COS/AWS_S3`). Values: 
	// <li>`on`: Enable private authentication.</li>
	// <li>`off`: Disable private authentication.</li>If this field is not specified, the default value `off` is used.
	PrivateAccess *string `json:"PrivateAccess,omitempty" name:"PrivateAccess"`

	// The private authentication parameters. This field is valid when `PrivateAccess=on`.
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitempty" name:"PrivateParameters"`
}

type OriginProtectionInfo struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names.
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// List of proxy IDs.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// The existing intermediate IPs.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CurrentIPWhitelist *IPWhitelist `json:"CurrentIPWhitelist,omitempty" name:"CurrentIPWhitelist"`

	// Whether the intermediate IP update is needed for the site. Values:
	// <li>`true`: Update needed;</li>
	// <li>`false`: Update not needed.</li>
	NeedUpdate *bool `json:"NeedUpdate,omitempty" name:"NeedUpdate"`

	// Status of the origin protection configuration. Values:
	// <li>`online`: Origin protection is activated;</li>
	// <li>`offline`: Origin protection is disabled.</li>
	// <li>`nonactivate`: Origin protection is not activated. This value is returned only when the feature is not activated before it’s used.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Whether origin protection is supported in the plan. Values:
	// <li>`true`: Origin protection supported;</li>
	// <li>`false`: Origin protection not supported.</li>
	PlanSupport *bool `json:"PlanSupport,omitempty" name:"PlanSupport"`

	// Differences between the latest and existing intermediate IPs.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiffIPWhitelist *DiffIPWhitelist `json:"DiffIPWhitelist,omitempty" name:"DiffIPWhitelist"`
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
	// Whether to enable POST upload limit (default limit: 32 MB). Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
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
	// Whether to enable QUIC. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
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

	// The custom rate limiting rules. If it is `null`, the previous settings is used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RateLimitCustomizes []*RateLimitUserRule `json:"RateLimitCustomizes,omitempty" name:"RateLimitCustomizes"`
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

	// The rule ID, which is only used as a response parameter.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
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

	// Applicable area. Values:
	// <li>`mainland`: Chinese mainland</li>
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`global`: Global</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// The resource type. Values:
	// <li>`plan`: Plan resources</li>
	// <li>`pay-as-you-go`: Pay-as-you-go resources </li>
	// <li>`value-added`: Value-added resources </li>
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Group *string `json:"Group,omitempty" name:"Group"`

	// The sites that are associated with the current resources.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ZoneNumber *int64 `json:"ZoneNumber,omitempty" name:"ZoneNumber"`
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

	// Slow attack defense configuration. If it is `null`, the previous setting is used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SlowPostConfig *SlowPostConfig `json:"SlowPostConfig,omitempty" name:"SlowPostConfig"`
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

type SlowPostConfig struct {
	// Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Detect slow attacks by the transfer period of the first 8 KB of requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FirstPartConfig *FirstPartConfig `json:"FirstPartConfig,omitempty" name:"FirstPartConfig"`

	// Detect slow attacks by the data rate of the main body (excluding the first 8 KB) of requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SlowRateConfig *SlowRateConfig `json:"SlowRateConfig,omitempty" name:"SlowRateConfig"`

	// The action to taken when a slow attack is detected. Values:
	// <li>`monitor`: Observe</li>
	// <li>`drop`: Block the request</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// ID of the rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
}

type SlowRateConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The sampling interval in seconds. In this way, the first 8 KB of the request is ignored. The rest of data is separated in to multiple parts according to this interval for slow attack measurement.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// The transfer rate threshold in bps. When the transfer rate of a sample is lower than the threshold, it’s considered a slow attack and handled according to the specified `Action`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`
}

type SmartRouting struct {
	// Whether to enable smart acceleration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type StandardDebug struct {
	// Whether to enable standard debugging. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable </li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The client IP to allow. It can be an IPv4/IPv6 address or a CIDR block. If not specified, it means to allow any client IP
	AllowClientIPList []*string `json:"AllowClientIPList,omitempty" name:"AllowClientIPList"`

	// The time when the standard debugging setting expires. If it is exceeded, this feature u200dbecomes invalid.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
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

	// Quota for a resource. Values:
	// <li>`zone`: Quota for sites</li>
	// <li>`custom-rule`: Quota for custom rules</li>
	// <li>`rate-limiting-rule`: Quota for rate limiting rules</li>
	// <li>`l4-proxy-instance`: Quota for L4 proxy instances </li>
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Pack *string `json:"Pack,omitempty" name:"Pack"`

	// ID of the L4 proxy instance.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The protection specification.
	// Values: <li> `cm_30G`: 30 Gbps base protection bandwidth in **Chinese mainland** service area</li><li> `cm_60G`: 60 Gbps base protection bandwidth in **Chinese mainland** service area</li><li> `cm_100G`: 100 Gbps base protection bandwidth in **Chinese mainland** service area</li><li> `anycast_300G`: 300 Gbps Anycast-based protection in **Global (MLC)** service area</li><li> `anycast_unlimited`: Unlimited Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> `cm_30G_anycast_300G`: 30 Gbps base protection bandwidth in **Chinese mainland** service area and 300 Gbps Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> `cm_30G_anycast_unlimited`: 30 Gbps base protection bandwidth in **Chinese mainland** service area and unlimited Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> cm_60G_anycast_300G`: 60 Gbps base protection bandwidth in **Chinese mainland** service area and 300 Gbps Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> cm_60G_anycast_unlimited`: 60 Gbps base protection bandwidth in **Chinese mainland** service area and unlimited Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> `cm_100G_anycast_300G`: 100 Gbps base protection bandwidth in **Chinese mainland** service area and 300 Gbps Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> cm_100G_anycast_unlimited`: 100 Gbps base protection bandwidth in **Chinese mainland** service area and unlimited Anycast-based protection bandwidth in **Global (MLC)** service area </li>
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ProtectionSpecs *string `json:"ProtectionSpecs,omitempty" name:"ProtectionSpecs"`
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
	// Whether to enable HTTP2 origin-pull. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
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

type WafRule struct {
	// Whether to enable managed rules. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// IDs of the rules to be disabled.
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitempty" name:"BlockRuleIDs"`

	// IDs of the rules to be executed in Observe mode.
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitempty" name:"ObserveRuleIDs"`
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

	// Access mode of the site. Values:
	// <li> `full`: Access through a name server.</li>
	// <li> `partial`: Access through a CNAME record.</li>
	// <li> `noDomainAccess`: Access without using a domain name </li>
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

	// Whether it’s a fake site. Valid values: 
	// <li>`0`: Non-fake site;</li>
	// <li>`1`: Fake site.</li>
	IsFake *int64 `json:"IsFake,omitempty" name:"IsFake"`

	// Lock status. Valid values: <li>`enable`: Normal. Modifying is allowed;</li><li>`disable`: Locked. Modifying is not allowed.</li>
	LockStatus *string `json:"LockStatus,omitempty" name:"LockStatus"`
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

	// Image optimization configuration. 
	// Note: This field may return `null`, indicating that no valid value was found.
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitempty" name:"ImageOptimize"`

	// Cross-MLC-border acceleration. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitempty" name:"AccelerateMainland"`

	// Standard debugging configuration.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	StandardDebug *StandardDebug `json:"StandardDebug,omitempty" name:"StandardDebug"`
}