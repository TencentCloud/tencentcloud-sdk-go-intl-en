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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccelerateMainland struct {
	// Whether to enable Cross-MLC-border acceleration. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AccelerateType struct {
	// Acceleration switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type AccelerationDomain struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Status of the accelerated domain name. Values:
	// <li>`online`: Activated</li>
	// <li>`process`: Being deployed</li>
	// <li>`offline`: Disabled</li>
	// <li>`forbidden`: Blocked</li>
	// <li>`init`: Pending activation</li>
	DomainStatus *string `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

	// Details of the origin.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginDetail *OriginDetail `json:"OriginDetail,omitnil,omitempty" name:"OriginDetail"`

	// Origin-pull protocol configuration. Values:
	// <li>`FOLLOW`: Follow the protocol of origin</li>
	// <li>`HTTP`: Send requests to the origin over HTTP</li>
	// <li>`HTTPS`: Send requests to the origin over HTTPS</li>
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// The port used for HTTP origin-pull requests
	// Note: This field may return·null, indicating that no valid values can be obtained.
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// The port used for HTTPS origin-pull requests
	// Note: This field may return·null, indicating that no valid values can be obtained.
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6 status. Values:
	// <li>`follow`: Follow the IPv6 configuration of the site</li>
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`

	// The CNAME address.
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// Ownership verification status. Values: <li>`pending`: Pending verification</li> <li>`finished`: Verified</li>	
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentificationStatus *string `json:"IdentificationStatus,omitnil,omitempty" name:"IdentificationStatus"`

	// Creation time of the accelerated domain name.
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// Modification time of the accelerated domain name.
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// Information required to verify the ownership of a domain name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// Domain name certificate information
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Certificate *AccelerationDomainCertificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`
}

type AccelerationDomainCertificate struct {
	// Certificate configuration mode. Values: <li>`disable`: Do not configure the certificate;</li><li>`eofreecert`: Use a free certificate provided by EdgeOne; </li><li>`sslcert`: Configure an SSL certificate.</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// List of certificates
	// Note: This field may return·null, indicating that no valid values can be obtained.
	List []*CertificateInfo `json:"List,omitnil,omitempty" name:"List"`
}

type AclCondition struct {
	// Filters: 
	// <li>`host`: Request domain name;</li>
	// <li>`sip`: Client IP;</li>
	// <li>`ua`: User-Agent;</li>
	// <li>`cookie`: Cookie;</li>
	// <li>`cgi`: CGI script;</li>
	// <li>`xff`: XFF header;</li></li>
	// <li>`url`: Request URL;<li></li>
	// <li>`accept`: Request content type;</li>
	// <li>`method`: Request method<;/li>
	// <li>`header`: Request header;</li>
	// <li>`app_proto`: Application layer protocol;</li>
	// <li>`sip_proto`: Network layer protocol;</li>
	// <li>`uabot`: UA rules (only available in custom bot rules);</li>
	// <li>`idcid`: IDC rules (only available in custom bot rules);</li>
	// <li>`sipbot`: Search engine rules (only available in custom bot rules);</li>
	// <li>`portrait`: Client reputation (only available in custom bot rules);</li>
	// <li>`header_seq`: Header sequence (only available in custom bot rules);</li>
	// <li>`hdr`: Request body (only available in custom Web protection rules). </li>
	MatchFrom *string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// The parameter of the field. When `MatchFrom = header`, the key contained in the header can be passed.
	MatchParam *string `json:"MatchParam,omitnil,omitempty" name:"MatchParam"`

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
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// The content to match.
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type AclConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The custom rule.
	AclUserRules []*AclUserRule `json:"AclUserRules,omitnil,omitempty" name:"AclUserRules"`

	// Custom managed rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Customizes []*AclUserRule `json:"Customizes,omitnil,omitempty" name:"Customizes"`
}

type AclUserRule struct {
	// The rule name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// The action. Values:
	// <li>`trans`: Allow</li>
	// <li>`drop`: Block the request</li>
	// <li>`monitor`: Observe</li>
	// <li>`ban`: Block the IP</li>
	// <li>`redirect`: Redirect the request</li>
	// <li>`page`: Return the specified page</li>
	// <li>`alg`: JavaScript challenge</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// The custom rule.
	AclConditions []*AclCondition `json:"AclConditions,omitnil,omitempty" name:"AclConditions"`

	// The rule priority. Value range: 0-100.
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// Rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// The update time, which is only used as an output parameter.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// IP ban duration. Range: 0-2 days. It's required when `Action=ban`. 
	PunishTime *int64 `json:"PunishTime,omitnil,omitempty" name:"PunishTime"`

	// The unit of the IP ban duration. Values:
	// <li>`second`: Second</li>
	// <li>`minutes`: Minute</li>
	// <li>`hour`: Hour</li>Default value: `second`.
	PunishTimeUnit *string `json:"PunishTimeUnit,omitnil,omitempty" name:"PunishTimeUnit"`

	// Name of the custom return page. It's required when `Action=page`.	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// (Disused) ID of the custom return page. The default value is 0, which means that the system default blocking page is used. 
	PageId *int64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// ID of custom response. The ID can be obtained via the `DescribeCustomErrorPages` API. It's required when `Action=page`.	
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`

	// The response code to trigger the return page. It's required when `Action=page`. Value: 100-600. 3xx response codes are not supported. Default value: 567.
	ResponseCode *int64 `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// The redirection URL. It's required when `Action=redirect`.	
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`
}

type Action struct {
	// Common feature operations. The options for this category include:
	// <li> Access URL overriding (AccessUrlRedirect);</li>
	// <li> Origin URL overriding (UpstreamUrlRedirect);</li>
	// <li> QUIC;</li>
	// <li> WebSocket;</li>
	// <li> Video dragging (VideoSeek);</li>
	// <li> Token authentication (Authentication);</li>
	// <li> Custom CacheKey (CacheKey);</li>
	// <li> Node caching TTL (Cache);</li>
	// <li> Browser caching TTL (MaxAge);</li>
	// <li> Offline caching (OfflineCache);</li>
	// <li> Smart routing (SmartRouting);</li>
	// <li> Range-based origin pull (RangeOriginPull);</li>
	// <li> HTTP/2 origin pull (UpstreamHttp2);</li>
	// <li> Host header overriding (HostHeader);</li>
	// <li> Forced HTTPS (ForceRedirect);</li>
	// <li> HTTPS origin pull (OriginPullProtocol);</li>
	// <li> Cache pre-refresh (CachePrefresh);</li>
	// <li> Smart compression (Compression);</li>
	// <li> Hsts;</li>
	// <li> ClientIpHeader;</li>
	// <li> SslTlsSecureConf;</li>
	// <li> OcspStapling;</li>
	// <li> HTTP/2 access (Http2);</li>
	// <li> Redirection during origin pull (UpstreamFollowRedirect);</li>
	// <li> Modifying origin server (Origin);</li>
	// <li> Layer 7 origin pull timeout (HTTPUpstreamTimeout);</li>
	// <li> HTTP response (HttpResponse).</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	NormalAction *NormalAction `json:"NormalAction,omitnil,omitempty" name:"NormalAction"`

	// Feature operation with a request/response header. Features of this type include:
	// <li>`RequestHeader`: HTTP request header modification.</li>
	// <li>`ResponseHeader`: HTTP response header modification.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RewriteAction *RewriteAction `json:"RewriteAction,omitnil,omitempty" name:"RewriteAction"`

	// Feature operation with a status code. Features of this type include:
	// <li>`ErrorPage`: Custom error page.</li>
	// <li>`StatusCodeCache`: Status code cache TTL.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	CodeAction *CodeAction `json:"CodeAction,omitnil,omitempty" name:"CodeAction"`
}

type AdvancedFilter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value of the filtered field.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Whether to enable fuzzy query.
	Fuzzy *bool `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`
}

type AiRule struct {
	// The status of the AI rule engine. Values:
	// <li>`smart_status_close`: Disabled</li>
	// <li>`smart_status_open`: Block</li>
	// <li>`smart_status_observe`: Observe</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type AlgDetectJS struct {
	// Method to validate client behavior.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Proof-of-work strength. Values:
	// <li>`low` (default): Low</li>
	// <li>`middle`: Medium</li>
	// <li>`high`: High</li>
	WorkLevel *string `json:"WorkLevel,omitnil,omitempty" name:"WorkLevel"`

	// Implement a delay before executing JS in milliseconds. Value range: 0-1000. Default value: 500.
	ExecuteMode *int64 `json:"ExecuteMode,omitnil,omitempty" name:"ExecuteMode"`

	// The period threshold for validating the result "Client JS disabled" in seconds. Value range: 5-3600. Default value: 10.
	InvalidStatTime *int64 `json:"InvalidStatTime,omitnil,omitempty" name:"InvalidStatTime"`

	// The number of times for the result "Client JS disabled" occurred in the specified period. Value range: 1-100000000. Default value: 30.
	InvalidThreshold *int64 `json:"InvalidThreshold,omitnil,omitempty" name:"InvalidThreshold"`

	// Client behavior validation results.
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitnil,omitempty" name:"AlgDetectResults"`
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
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The action. Values:
	// <li>`drop`: Block</li>
	// <li>`monitor`: Observe</li>
	// <li>`silence`: Drop w/o response</li>
	// <li>`shortdelay`: Add short latency</li>
	// <li>`longdelay`: Add long latency</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type AlgDetectRule struct {
	// ID of the rule.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Name of the rule.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Whether to enable the rule.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Condition specified for the rule.
	AlgConditions []*AclCondition `json:"AlgConditions,omitnil,omitempty" name:"AlgConditions"`

	// Validate Cookie when the condition is satisfied.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AlgDetectSession *AlgDetectSession `json:"AlgDetectSession,omitnil,omitempty" name:"AlgDetectSession"`

	// Validate client behavior when the condition is satisfied.
	AlgDetectJS []*AlgDetectJS `json:"AlgDetectJS,omitnil,omitempty" name:"AlgDetectJS"`

	// The update time, which is only used as an output parameter.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AlgDetectSession struct {
	// Method to validate Cookie.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The validation mode. Values:
	// <li>`detect`: Validate only</li>
	// <li>`update_detect` (default): Update Cookie and validate</li>
	DetectMode *string `json:"DetectMode,omitnil,omitempty" name:"DetectMode"`

	// Whether to enable Cookie-based session check. The default value is `off`. Values:
	// <li>`off`: Disable</li>
	// <li>`on`: Enable</li>
	SessionAnalyzeSwitch *string `json:"SessionAnalyzeSwitch,omitnil,omitempty" name:"SessionAnalyzeSwitch"`

	// The period threshold for validating the result "No Cookie/Cookie expired" in seconds. Value range: 5-3600. Default value: 10.
	InvalidStatTime *int64 `json:"InvalidStatTime,omitnil,omitempty" name:"InvalidStatTime"`

	// The number of times for the result "No Cookie/Cookie expired" occurred in the specified period. Value range: 1-100000000. Default value: 300.
	InvalidThreshold *int64 `json:"InvalidThreshold,omitnil,omitempty" name:"InvalidThreshold"`

	// Cookie validation results.
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitnil,omitempty" name:"AlgDetectResults"`

	// Cookie-based session check results.
	SessionBehaviors []*AlgDetectResult `json:"SessionBehaviors,omitnil,omitempty" name:"SessionBehaviors"`
}

type AliasDomain struct {
	// The alias domain name.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// Status of the alias domain name. Values:
	// <li>`active`: Activated</li>
	// <li>`pending`: Deploying</li>
	// <li>`conflict`: Reclaimed</li>
	// <li>`stop`: Stopped</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The blocking mode. Values:
	// <li>`0`: Not blocked</li>
	// <li>`11`: Blocked due to regulatory compliance</li>
	// <li>`14`: Blocked due to ICP filing not obtained</li>
	ForbidMode *int64 `json:"ForbidMode,omitnil,omitempty" name:"ForbidMode"`

	// Creation time of the alias domain name.
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// Modification time of the alias domain name.
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type ApplicationProxy struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The domain name or subdomain name when `ProxyType=hostname`.
	// The instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// The proxy type. Values:
	// <li>`hostname`: The proxy is created by subdomain name.</li>
	// <li>`instance`: The proxy is created by instance.</li>
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// The scheduling mode. Values:
	// <li>`ip`: Schedule via Anycast IP.</li>
	// <li>`domain`: Schedule via CNAME.</li>
	PlatType *string `json:"PlatType,omitnil,omitempty" name:"PlatType"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// Default value: overseas.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Whether to enable security protection. Values:
	// <li>`0`: Disable security protection.</li>
	// <li>`1`: Enable security protection.</li>
	SecurityType *int64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// Whether to enable acceleration. Values:
	// <li>`0`: Disable acceleration.</li>
	// <li>`1`: Enable acceleration.</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// The session persistence duration.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// The rule status. Values:
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	// <li>`progress`: Deploying</li>
	// <li>`stopping`: Disabling</li>
	// <li>`fail`: Failed to deploy or disable</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The blocking status of the proxy. Values:
	// <li>`banned`: Blocked</li>
	// <li>`banning`: Blocking</li>
	// <li>`recover`: Unblocked</li>
	// <li>`recovering`: Unblocking</li>
	BanStatus *string `json:"BanStatus,omitnil,omitempty" name:"BanStatus"`

	// Scheduling information.
	ScheduleValue []*string `json:"ScheduleValue,omitnil,omitempty" name:"ScheduleValue"`

	// When `ProxyType=hostname`:
	// This field indicates the unique ID of the subdomain name.
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// The IPv6 access configuration.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// The update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// List of rules.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil,omitempty" name:"ApplicationProxyRules"`

	// Cross-MLC-border acceleration.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type ApplicationProxyRule struct {
	// The protocol. Values:
	// <li>`TCP`: TCP protocol.</li>
	// <li>`UDP`: UDP protocol.</li>
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	// Note that each rule can have up to 20 ports.
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// Origin server type. Valid values:<li>custom: Manually added;</li><li>loadbalancer: Cloud Load Balancer;</li><li>origins: Origin server group.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Details of the origin server:
	// <li>When OriginType is custom, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or OriginValue=["test.com"];</li><li>When OriginType is loadbalancer, it indicates a single Cloud Load Balancer, such as ["lb-xdffsfasdfs"];</li>
	// <li>When OriginType is origins, it requires one and only one element, which represents an origin server group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The rule status. Values:
	// <li>`online`: Enabled.</li>
	// <li>`offline`: Disabled.</li>
	// <li>`progress`: Deploying</li>
	// <li>`stopping`: Disabling</li>
	// <li>`fail`: Failed to deploy or disable</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>Default value: OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>Default value: false
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// Rule tag.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
}

type AscriptionInfo struct {

	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// The record type.
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// The record value.
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

type BillingData struct {
	// Time.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Value.
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type BillingDataFilter struct {
	// Parameter name.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Parameter value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type BindSecurityTemplateToEntityRequestParams struct {
	// Site ID of the policy template to be bound to or unbound from.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of domain names to bind to/unbind from a policy template
	Entities []*string `json:"Entities,omitnil,omitempty" name:"Entities"`

	// Action options. Values include:
	// <li>`bind`: Bind the domain names to the specified policy template </li>
	// <li>`unbind-keep-policy`: Unbind a domain name from a policy template and keep the current policy when unbinding</li>
	// <li>`unbind-use-default`: Unbind domain names from policy templates and use default blank policy.</li> Note: Only one domain name can be unbound at one time. When `Operate` is `unbind-keep-policy` or `unbind-use-default`, there can only be one domain name specified in `Entities`.
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// Specifies the ID of the policy template or the site's global policy to be bound or unbound.
	// - To bind to a policy template, or unbind from it, specify the policy template ID.
	// - To bind to the site's global policy, or unbind from it, use the @ZoneLevel@domain parameter value.
	// 
	// Note: After unbinding, the domain name will use an independent policy and rule quota will be calculated separately. Please make sure there is sufficient rule quota before unbinding.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Whether to replace the existing policy template bound with the domain name. Values: 
	// <li>`true`: Replace the template bound to the domain. </li>
	// <li>`false`: Do not replace the template.</li> Note: In this case, the API returns an error if there is already a policy template bound to the specified domain name.
	OverWrite *bool `json:"OverWrite,omitnil,omitempty" name:"OverWrite"`
}

type BindSecurityTemplateToEntityRequest struct {
	*tchttp.BaseRequest
	
	// Site ID of the policy template to be bound to or unbound from.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of domain names to bind to/unbind from a policy template
	Entities []*string `json:"Entities,omitnil,omitempty" name:"Entities"`

	// Action options. Values include:
	// <li>`bind`: Bind the domain names to the specified policy template </li>
	// <li>`unbind-keep-policy`: Unbind a domain name from a policy template and keep the current policy when unbinding</li>
	// <li>`unbind-use-default`: Unbind domain names from policy templates and use default blank policy.</li> Note: Only one domain name can be unbound at one time. When `Operate` is `unbind-keep-policy` or `unbind-use-default`, there can only be one domain name specified in `Entities`.
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// Specifies the ID of the policy template or the site's global policy to be bound or unbound.
	// - To bind to a policy template, or unbind from it, specify the policy template ID.
	// - To bind to the site's global policy, or unbind from it, use the @ZoneLevel@domain parameter value.
	// 
	// Note: After unbinding, the domain name will use an independent policy and rule quota will be calculated separately. Please make sure there is sufficient rule quota before unbinding.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Whether to replace the existing policy template bound with the domain name. Values: 
	// <li>`true`: Replace the template bound to the domain. </li>
	// <li>`false`: Do not replace the template.</li> Note: In this case, the API returns an error if there is already a policy template bound to the specified domain name.
	OverWrite *bool `json:"OverWrite,omitnil,omitempty" name:"OverWrite"`
}

func (r *BindSecurityTemplateToEntityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecurityTemplateToEntityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Entities")
	delete(f, "Operate")
	delete(f, "TemplateId")
	delete(f, "OverWrite")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSecurityTemplateToEntityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSecurityTemplateToEntityResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindSecurityTemplateToEntityResponse struct {
	*tchttp.BaseResponse
	Response *BindSecurityTemplateToEntityResponseParams `json:"Response"`
}

func (r *BindSecurityTemplateToEntityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecurityTemplateToEntityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSharedCNAMEMap struct {
	// The shared CNAME to be bound with or unbound from.
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`

	// Acceleration domains (up to 20).
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`
}

// Predefined struct for user
type BindSharedCNAMERequestParams struct {
	// ID of the site related with the acceleration domain name.	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Action type. Values:
	// <li>`bind`: To bind</li>
	// <li>`unbind`: To unbind</li>
	BindType *string `json:"BindType,omitnil,omitempty" name:"BindType"`

	// Bindings between domain names and a shared CNAME
	BindSharedCNAMEMaps []*BindSharedCNAMEMap `json:"BindSharedCNAMEMaps,omitnil,omitempty" name:"BindSharedCNAMEMaps"`
}

type BindSharedCNAMERequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the acceleration domain name.	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Action type. Values:
	// <li>`bind`: To bind</li>
	// <li>`unbind`: To unbind</li>
	BindType *string `json:"BindType,omitnil,omitempty" name:"BindType"`

	// Bindings between domain names and a shared CNAME
	BindSharedCNAMEMaps []*BindSharedCNAMEMap `json:"BindSharedCNAMEMaps,omitnil,omitempty" name:"BindSharedCNAMEMaps"`
}

func (r *BindSharedCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSharedCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "BindType")
	delete(f, "BindSharedCNAMEMaps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSharedCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSharedCNAMEResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindSharedCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *BindSharedCNAMEResponseParams `json:"Response"`
}

func (r *BindSharedCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSharedCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindZoneToPlanRequestParams struct {
	// ID of the site to be bound.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// ID of the target plan.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type BindZoneToPlanRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site to be bound.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// ID of the target plan.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The settings of the bot managed rule. If it is null, the settings that were last configured will be used.
	BotManagedRule *BotManagedRule `json:"BotManagedRule,omitnil,omitempty" name:"BotManagedRule"`

	// The settings of the client reputation rule. If it is null, the settings that were last configured will be used.
	BotPortraitRule *BotPortraitRule `json:"BotPortraitRule,omitnil,omitempty" name:"BotPortraitRule"`

	// The bot intelligence settings. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntelligenceRule *IntelligenceRule `json:"IntelligenceRule,omitnil,omitempty" name:"IntelligenceRule"`

	// Settings of the custom bot rule. If it is null, the settings that were last configured will be used.
	BotUserRules []*BotUserRule `json:"BotUserRules,omitnil,omitempty" name:"BotUserRules"`

	// Active bot detection rule.
	AlgDetectRule []*AlgDetectRule `json:"AlgDetectRule,omitnil,omitempty" name:"AlgDetectRule"`

	// Settings of the bot managed rule. It is only used for output.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Customizes []*BotUserRule `json:"Customizes,omitnil,omitempty" name:"Customizes"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The probability for triggering the action. Value range: 0-100.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`
}

type BotManagedRule struct {
	// The rule action. Values:
	// <li>`drop`: Block</li>
	// <li>`trans`: Allow</li>
	// <li>`alg`: JavaScript challenge</li>
	// <li>`monitor`: Observe</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// The ID of the rule that applies the "Allow" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransManagedIds []*int64 `json:"TransManagedIds,omitnil,omitempty" name:"TransManagedIds"`

	// The ID of the rule that applies the "JavaScript challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil,omitempty" name:"AlgManagedIds"`

	// The ID of the rule that applies the "Managed challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil,omitempty" name:"CapManagedIds"`

	// The ID of the rule that applies the "Observe" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil,omitempty" name:"MonManagedIds"`

	// The ID of the rule that applies the "Block" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil,omitempty" name:"DropManagedIds"`
}

type BotPortraitRule struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// The ID of the rule that applies the "JavaScript challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil,omitempty" name:"AlgManagedIds"`

	// The ID of the rule that applies the "Managed challenge" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil,omitempty" name:"CapManagedIds"`

	// The ID of the rule that applies the "Observe" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil,omitempty" name:"MonManagedIds"`

	// The ID of the rule that applies the "Block" action.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil,omitempty" name:"DropManagedIds"`
}

type BotUserRule struct {

	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// The action. Values:
	// <li>`drop`: Block the request</li>
	// <li>`monitor`: Observe</li>
	// <li>`trans`: Allow</li>
	// <li>`redirect`: Redirect the request</li>
	// <li>`page`: Return the specified page</li>
	// <li>`alg`: JavaScript challenge</li>
	// <li>`captcha`: Managed challenge</li>
	// <li>`random`: Handle the request randomly by the weight</li>
	// <li>`silence`: Keep the connection but do not response to the client</li>
	// <li>`shortdelay`: Add a short latency period</li>
	// <li>`longdelay`: Add a long latency period</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>Default value: `on`
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// Details of the rule.
	AclConditions []*AclCondition `json:"AclConditions,omitnil,omitempty" name:"AclConditions"`

	// The rule weight. Value range: 0-100.
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// Rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// [Currently unavailable] Specify the random action and percentage.
	ExtendActions []*BotExtendAction `json:"ExtendActions,omitnil,omitempty" name:"ExtendActions"`

	// The filter. Values:
	// <li>`sip`: Client IP</li>
	// This parameter is left empty by default.
	FreqFields []*string `json:"FreqFields,omitnil,omitempty" name:"FreqFields"`

	// The update time, which is only used as an output parameter.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Query scope. Values:
	// <li>`source_to_eo`: (Response) Traffic going from the origin to EdgeOne.</li>
	// <li>`client_to_eo`: (Request) Traffic going from the client to EdgeOne.</li>
	// Default: `source_to_eo`.
	FreqScope []*string `json:"FreqScope,omitnil,omitempty" name:"FreqScope"`

	// Name of the custom return page. It's required when `Action=page`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// ID of custom response. The ID can be obtained via the `DescribeCustomErrorPages` API. It's required when `Action=page`.	
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`

	// The response code to trigger the return page. It's required when `Action=page`. Value: 100-600. 3xx response codes are not supported. Default value: 567.
	ResponseCode *int64 `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// The redirection URL. It's required when `Action=redirect`.
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`
}

type CC struct {
	// WAF switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// ID of the policy
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type CLSTopic struct {
	// The ID of the Tencent Cloud CLS log set.
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// The ID of the Tencent Cloud CLS log topic.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// The region of the Tencent Cloud CLS log set.
	LogSetRegion *string `json:"LogSetRegion,omitnil,omitempty" name:"LogSetRegion"`
}

type Cache struct {
	// Whether to enable cache configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Cache expiration time setting.
	// Unit: second. The maximum value is 365 days.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheTime *int64 `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`

	// Whether to enable force cache. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: IgnoreCacheControl is deprecated.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitnil,omitempty" name:"IgnoreCacheControl"`
}

type CacheConfig struct {
	// Cache configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cache *Cache `json:"Cache,omitnil,omitempty" name:"Cache"`

	// No-cache configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoCache *NoCache `json:"NoCache,omitnil,omitempty" name:"NoCache"`

	// Follows the origin server configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	FollowOrigin *FollowOrigin `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`
}

type CacheKey struct {
	// Whether to enable full-path cache. Values:
	// <li>`on`: Enable full-path cache (i.e., disable Ignore Query String).</li>
	// <li>`off`: Disable full-path cache (i.e., enable Ignore Query String).</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullUrlCache *string `json:"FullUrlCache,omitnil,omitempty" name:"FullUrlCache"`

	// Whether to ignore case in the cache key. Values:
	// <li>`on`: Ignore</li>
	// <li>`off`: Not ignore</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IgnoreCase *string `json:"IgnoreCase,omitnil,omitempty" name:"IgnoreCase"`

	// Request parameter contained in `CacheKey`. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QueryString *QueryString `json:"QueryString,omitnil,omitempty" name:"QueryString"`
}

type CachePrefresh struct {
	// Whether to enable cache prefresh. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The cache prefresh percentage. Values: 1-99
	// Note: This field may return null, indicating that no valid values can be obtained.
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`
}

type CacheTag struct {
	// List of domain names to purge cache for.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type CertificateInfo struct {
	// ID of the server certificate.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Alias of the certificate.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Type of the certificate. Values:
	// <li>`default`: Default certificate</li>
	// <li>`upload`: Specified certificate</li>
	// <li>`managed`: Tencent Cloud-managed certificate</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The certificate expiration time.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Time when the certificate is deployed.
	DeployTime *string `json:"DeployTime,omitnil,omitempty" name:"DeployTime"`

	// Signature algorithm.
	SignAlgo *string `json:"SignAlgo,omitnil,omitempty" name:"SignAlgo"`

	// Status of the certificate. Values:
	// u200c<li>`deployed`: The deployment has completed</li>
	// u200c<li>`processing`: Deployment in progress</li>
	// u200c<li>`applying`: Application in progress</li>
	// u200c<li>`failed`: Application rejected</li>
	// <li>`issued`: Binding failed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type CheckCnameStatusRequestParams struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of accelerated domain names.
	RecordNames []*string `json:"RecordNames,omitnil,omitempty" name:"RecordNames"`
}

type CheckCnameStatusRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of accelerated domain names.
	RecordNames []*string `json:"RecordNames,omitnil,omitempty" name:"RecordNames"`
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
	// CNAME status of accelerated domain names.
	CnameStatus []*CnameStatus `json:"CnameStatus,omitnil,omitempty" name:"CnameStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Name of the request header that contains the client IP region. It is valid when `Switch=on`. 
	// The default value `EO-Client-IPCountry` is used when it is not specified.
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
}

type ClientIpHeader struct {
	// Whether to enable the configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Name of the request header that contains the client IP for origin-pull. 
	// The default value `X-Forwarded-IP` is used when it is not specified. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
}

type CnameStatus struct {
	// The domain name.
	RecordName *string `json:"RecordName,omitnil,omitempty" name:"RecordName"`

	// The CNAME address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// The CNAME status. Values:
	// <li>`active`: Activated</li>
	// <li>`moved`: Not activated </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type CodeAction struct {
	// Feature name. For details, see [DescribeRulesSetting](https://intl.cloud.tencent.com/document/product/1552/80618?from_cn_redirect=1) API
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Operation parameter.
	Parameters []*RuleCodeActionParams `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type Compression struct {
	// Whether to enable smart compression. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Compression algorithm. Values:
	// <li>`brotli`: Brotli algorithm</li>
	// <li>`gzip`: Gzip algorithm</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Algorithms []*string `json:"Algorithms,omitnil,omitempty" name:"Algorithms"`
}

type ConfigGroupVersionInfo struct {
	// Version ID.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Version No.
	VersionNumber *string `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// Configuraration group ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Configuration group type. Valid values: 
	// <li>l7_acceleration: L7 acceleration configuration group. </li>
	// <li>edge_functions: Edge function configuration group. </li>
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Version description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Version status. Valid values: 
	// <li>creating: Being created.</li>
	// <li>inactive: Not effective.</li>
	// <li>active: Effective. </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Version creation time. The time format follows the ISO 8601 standard and is represented in Coordinated Universal Time (UTC).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type CreateAccelerationDomainRequestParams struct {
	// ID of the site related with the acceleration domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Acceleration domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// Origin-pull protocol configuration. Values:
	// <li>`FOLLOW`: Follow the protocol of origin</li>
	// <li>`HTTP`: Send requests to the origin over HTTP</li>
	// <li>`HTTPS`: Send requests to the origin over HTTPS</li>
	// <li>Default: `FOLLOW`</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// Ports for HTTP origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTP`. Port 80 is used if it's not specified. 
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// Ports for HTTPS origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTPS`. Port 443 is used if it's not specified. 
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6 status. Values:
	// <li>`follow`: Follow the IPv6 configuration of the site</li>
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// <li>Default: `follow`</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
}

type CreateAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the acceleration domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Acceleration domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// Origin-pull protocol configuration. Values:
	// <li>`FOLLOW`: Follow the protocol of origin</li>
	// <li>`HTTP`: Send requests to the origin over HTTP</li>
	// <li>`HTTPS`: Send requests to the origin over HTTPS</li>
	// <li>Default: `FOLLOW`</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// Ports for HTTP origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTP`. Port 80 is used if it's not specified. 
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// Ports for HTTPS origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTPS`. Port 443 is used if it's not specified. 
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6 status. Values:
	// <li>`follow`: Follow the IPv6 configuration of the site</li>
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// <li>Default: `follow`</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
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
	delete(f, "OriginProtocol")
	delete(f, "HttpOriginPort")
	delete(f, "HttpsOriginPort")
	delete(f, "IPv6Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccelerationDomainResponseParams struct {
	// Use the information returned by this parameter to verify the ownership of a domain name. For details, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: (Default) Do not configure</li>
	// <li>`hosting`: Managed SSL certificate</li>
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type CreateAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: (Default) Do not configure</li>
	// <li>`hosting`: Managed SSL certificate</li>
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// The scheduling mode. Values:
	// <li>`ip`: Schedule via Anycast IP.</li>
	// <li>`domain`: Schedule via CNAME.</li>
	PlatType *string `json:"PlatType,omitnil,omitempty" name:"PlatType"`

	// Whether to enable security protection. Values:
	// <li>`0`: Disable security protection.</li>
	// <li>`1`: Enable security protection.</li>
	SecurityType *int64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// Whether to enable acceleration. Values:
	// <li>`0`: Disable acceleration.</li>
	// <li>`1`: Enable acceleration.</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// Layer 4 proxy mode, with value: <li>instance: instance mode.</li>If not specified, the default value instance will be used.
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// If not specified, this field uses the default value 600.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// Ipv6 access configuration. 
	// IPv6 access is disabled if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// The rule details.
	// If this field is not specified, an application proxy rule will not be created.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil,omitempty" name:"ApplicationProxyRules"`

	// Cross-MLC-border acceleration. It is disabled if this parameter is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// The scheduling mode. Values:
	// <li>`ip`: Schedule via Anycast IP.</li>
	// <li>`domain`: Schedule via CNAME.</li>
	PlatType *string `json:"PlatType,omitnil,omitempty" name:"PlatType"`

	// Whether to enable security protection. Values:
	// <li>`0`: Disable security protection.</li>
	// <li>`1`: Enable security protection.</li>
	SecurityType *int64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// Whether to enable acceleration. Values:
	// <li>`0`: Disable acceleration.</li>
	// <li>`1`: Enable acceleration.</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// Layer 4 proxy mode, with value: <li>instance: instance mode.</li>If not specified, the default value instance will be used.
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// If not specified, this field uses the default value 600.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// Ipv6 access configuration. 
	// IPv6 access is disabled if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// The rule details.
	// If this field is not specified, an application proxy rule will not be created.
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil,omitempty" name:"ApplicationProxyRules"`

	// Cross-MLC-border acceleration. It is disabled if this parameter is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
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
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// Origin server type. Valid values:<li>custom: Manually added;</li>
	// <li>loadbalancer: Cloud Load Balancer;</li><li>origins: Origin server group.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Details of the origin server:<li>When OriginType is custom, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or OriginValue=["test.com"];</li><li>When OriginType is loadbalancer, it indicates a single Cloud Load Balancer, such as ["lb-xdffsfasdfs"];</li><li>When OriginType is origins, it requires one and only one element, which represents an origin server group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>Default value: OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable.</li>
	// <li>`false`: Disable.</li>Default value: false.
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// Rule tag. This parameter is left empty by default.
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
}

type CreateApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// Origin server type. Valid values:<li>custom: Manually added;</li>
	// <li>loadbalancer: Cloud Load Balancer;</li><li>origins: Origin server group.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Details of the origin server:<li>When OriginType is custom, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or OriginValue=["test.com"];</li><li>When OriginType is loadbalancer, it indicates a single Cloud Load Balancer, such as ["lb-xdffsfasdfs"];</li><li>When OriginType is origins, it requires one and only one element, which represents an origin server group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>Default value: OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable.</li>
	// <li>`false`: Disable.</li>Default value: false.
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// Rule tag. This parameter is left empty by default.
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
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
	delete(f, "RuleTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRuleResponseParams struct {
	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateCLSIndexRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The ID of the real-time log delivery task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CreateCLSIndexRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The ID of the real-time log delivery task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *CreateCLSIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCLSIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSIndexResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCLSIndexResponse struct {
	*tchttp.BaseResponse
	Response *CreateCLSIndexResponseParams `json:"Response"`
}

func (r *CreateCLSIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigGroupVersionRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// GroupId of the version to be created.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Configuration content to be imported. It is required to be in JSON format and encoded in UTF-8. Please refer to the example below for the configuration file content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Version description. The maximum length allowed is 50 characters. This field can be used to provide details about the application scenarios of this version.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateConfigGroupVersionRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// GroupId of the version to be created.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Configuration content to be imported. It is required to be in JSON format and encoded in UTF-8. Please refer to the example below for the configuration file content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Version description. The maximum length allowed is 50 characters. This field can be used to provide details about the application scenarios of this version.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateConfigGroupVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigGroupVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "Content")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigGroupVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigGroupVersionResponseParams struct {
	// Version ID.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigGroupVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigGroupVersionResponseParams `json:"Response"`
}

func (r *CreateConfigGroupVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigGroupVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizeErrorPageRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom response page name. The name must be 2-30 characters long.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Custom response page type, with values:<li>text/html; </li><li>application/json;</li><li>text/plain;</li><li>text/xml.</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// Custom response page description, not exceeding 60 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Custom response page content, not exceeding 2 KB.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CreateCustomizeErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom response page name. The name must be 2-30 characters long.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Custom response page type, with values:<li>text/html; </li><li>application/json;</li><li>text/plain;</li><li>text/xml.</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// Custom response page description, not exceeding 60 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Custom response page content, not exceeding 2 KB.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *CreateCustomizeErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "ContentType")
	delete(f, "Description")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomizeErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizeErrorPageResponseParams struct {
	// Custom response page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomizeErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomizeErrorPageResponseParams `json:"Response"`
}

func (r *CreateCustomizeErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizeErrorPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// Acceleration zone of the Layer 4 proxy instance.<li>mainland: Availability zone in the Chinese mainland;</li><li>overseas: Global availability zone (excluding the Chinese mainland);</li><li>global: Global availability zone.</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li>
	// <li>off: Disable.</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li>
	// <li>off: Disable.</li>
	StaticIp *string `json:"StaticIp,omitnil,omitempty" name:"StaticIp"`

	// Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li>
	// <li>off: Disable.</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
	DDosProtectionConfig *DDosProtectionConfig `json:"DDosProtectionConfig,omitnil,omitempty" name:"DDosProtectionConfig"`
}

type CreateL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// Acceleration zone of the Layer 4 proxy instance.<li>mainland: Availability zone in the Chinese mainland;</li><li>overseas: Global availability zone (excluding the Chinese mainland);</li><li>global: Global availability zone.</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li>
	// <li>off: Disable.</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li>
	// <li>off: Disable.</li>
	StaticIp *string `json:"StaticIp,omitnil,omitempty" name:"StaticIp"`

	// Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li>
	// <li>off: Disable.</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
	DDosProtectionConfig *DDosProtectionConfig `json:"DDosProtectionConfig,omitnil,omitempty" name:"DDosProtectionConfig"`
}

func (r *CreateL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyName")
	delete(f, "Area")
	delete(f, "Ipv6")
	delete(f, "StaticIp")
	delete(f, "AccelerateMainland")
	delete(f, "DDosProtectionConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyResponseParams struct {
	// Layer 4 instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateL4ProxyResponseParams `json:"Response"`
}

func (r *CreateL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyRulesRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rules. A single request supports up to 200 forwarding rules.
	// Note: When L4ProxyRule is used here, Protocol, PortRange, OriginType, OriginValue, and OriginPortRange are required fields; ClientIPPassThroughMode, SessionPersist, SessionPersistTime, and RuleTag are optional fields; do not fill in RuleId and Status.
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

type CreateL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rules. A single request supports up to 200 forwarding rules.
	// Note: When L4ProxyRule is used here, Protocol, PortRange, OriginType, OriginValue, and OriginPortRange are required fields; ClientIPPassThroughMode, SessionPersist, SessionPersistTime, and RuleTag are optional fields; do not fill in RuleId and Status.
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

func (r *CreateL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "L4ProxyRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ProxyRulesResponseParams struct {
	// IDs of newly added forwarding rules, returned as an array.
	L4ProxyRuleIds []*string `json:"L4ProxyRuleIds,omitnil,omitempty" name:"L4ProxyRuleIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateL4ProxyRulesResponseParams `json:"Response"`
}

func (r *CreateL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Origin group name. It can contain 1 to 200 characters ([a-z], [A-Z], [0-9] and [_-]).
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// (Required) Origin group type. Values:
	// <li>`GENERAL`: General origin groups. It supports IPs and domain names. It can be referenced by DNS, Rule Engine, Layer 4 Proxy and General LoadBalancer. </li>
	// <li>`HTTP`: HTTP-specific origin groups. It supports IPs/domain names and object storage buckets. It can be referenced by acceleration domain names, rule engines and HTTP LoadBalancer. It cannot be referenced by L4 proxies. </li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// (Required) Origins in the origin group.
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// Host header used for origin-pull. It only works when `Type=HTTP`. The `HostHeader` specified in `RuleEngine` takes a higher priority over this configuration.
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Origin group name. It can contain 1 to 200 characters ([a-z], [A-Z], [0-9] and [_-]).
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// (Required) Origin group type. Values:
	// <li>`GENERAL`: General origin groups. It supports IPs and domain names. It can be referenced by DNS, Rule Engine, Layer 4 Proxy and General LoadBalancer. </li>
	// <li>`HTTP`: HTTP-specific origin groups. It supports IPs/domain names and object storage buckets. It can be referenced by acceleration domain names, rule engines and HTTP LoadBalancer. It cannot be referenced by L4 proxies. </li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// (Required) Origins in the origin group.
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// Host header used for origin-pull. It only works when `Type=HTTP`. The `HostHeader` specified in `RuleEngine` takes a higher priority over this configuration.
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
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
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Records")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupResponseParams struct {
	// The ID of the origin group.
	OriginGroupId *string `json:"OriginGroupId,omitnil,omitempty" name:"OriginGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

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
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`
}

type CreatePlanForZoneRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

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
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`
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
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// List or order numbers.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreatePlanRequestParams struct {
	// Type of plan subscribed to. Valid values:<li>personal: Personal Edition Plan, prepaid;</li><li>basic: Basic Edition Plan, prepaid;</li><li>standard: Standard Edition Plan, prepaid;</li><li>enterprise: Enterprise Edition Plan, post-paid. </li>When subscribing to a prepaid plan, ensure there is sufficient balance in the account, as insufficient balance will result in a pending payment order.
	// For an overview of billing, refer to [Billing Overview](https://intl.cloud.tencent.com/document/product/1552/94156?from_cn_redirect=1).
	// For differences between plans, refer to [ Comparison of EdgeOne Plan](https://intl.cloud.tencent.com/document/product/1552/94165?from_cn_redirect=1).
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// Whether to automatically use a voucher. Valid values: <li>true: Yes;</li><li>false: No. </li> This parameter is valid only when PlanType is personal, basic, or standard.
	// If this field is not specified, the default value 'false' will be used.
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`

	// Parameters for subscribing to prepaid plans. When PlanType is personal, basic, or standard, this parameter can be optionally set to specify the subscription duration and whether to enable auto-renewal for the plan.
	// If this field is not specified, the default plan duration is 1 month, with auto-renewal disabled.
	PrepaidPlanParam *PrepaidPlanParam `json:"PrepaidPlanParam,omitnil,omitempty" name:"PrepaidPlanParam"`
}

type CreatePlanRequest struct {
	*tchttp.BaseRequest
	
	// Type of plan subscribed to. Valid values:<li>personal: Personal Edition Plan, prepaid;</li><li>basic: Basic Edition Plan, prepaid;</li><li>standard: Standard Edition Plan, prepaid;</li><li>enterprise: Enterprise Edition Plan, post-paid. </li>When subscribing to a prepaid plan, ensure there is sufficient balance in the account, as insufficient balance will result in a pending payment order.
	// For an overview of billing, refer to [Billing Overview](https://intl.cloud.tencent.com/document/product/1552/94156?from_cn_redirect=1).
	// For differences between plans, refer to [ Comparison of EdgeOne Plan](https://intl.cloud.tencent.com/document/product/1552/94165?from_cn_redirect=1).
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// Whether to automatically use a voucher. Valid values: <li>true: Yes;</li><li>false: No. </li> This parameter is valid only when PlanType is personal, basic, or standard.
	// If this field is not specified, the default value 'false' will be used.
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`

	// Parameters for subscribing to prepaid plans. When PlanType is personal, basic, or standard, this parameter can be optionally set to specify the subscription duration and whether to enable auto-renewal for the plan.
	// If this field is not specified, the default plan duration is 1 month, with auto-renewal disabled.
	PrepaidPlanParam *PrepaidPlanParam `json:"PrepaidPlanParam,omitnil,omitempty" name:"PrepaidPlanParam"`
}

func (r *CreatePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanType")
	delete(f, "AutoUseVoucher")
	delete(f, "PrepaidPlanParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanResponseParams struct {
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePlanResponse struct {
	*tchttp.BaseResponse
	Response *CreatePlanResponseParams `json:"Response"`
}

func (r *CreatePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of resources to be preheated. Each element format is similar to the following:
	// http://www.example.com/example.txt. The parameter value is currently required.
	// Note: The number of tasks that can be submitted is limited by the quota of a billing package. For details, see [Billing Overview] (https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1).
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Whether to encode a URL according to RFC3986. Enable this field when the URL contains non-ASCII characters.
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// HTTP header information
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of resources to be preheated. Each element format is similar to the following:
	// http://www.example.com/example.txt. The parameter value is currently required.
	// Note: The number of tasks that can be submitted is limited by the quota of a billing package. For details, see [Billing Overview] (https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1).
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Whether to encode a URL according to RFC3986. Enable this field when the URL contains non-ASCII characters.
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// HTTP header information
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
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
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// List of failed tasks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailedList []*FailReason `json:"FailedList,omitnil,omitempty" name:"FailedList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Type of cache purging. Values:
	// <li>`purge_url`: Purge by the URL</li>
	// <li>`purge_prefix`: Purge by the directory</li>
	// <li>`purge_host`: Purge by the hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by the cache-tag </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Node cache purge method, valid for directory, hostname, and all cache refreshes. Valid values: <li>invalidate: Refreshes only resources that were updated under the directory; </li><li>delete: Refreshes all node resources, regardless of whether they were updated. </li>Default value: invalidate.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// List of resources for which cache is to be purged. Each element format depends on the cache purge type and you can refer to the API examples for details. <li>The number of tasks that can be submitted at a time is limited by the quota of a billing package. For details, see [Billing Overview] (https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1).</li>
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	//
	// Deprecated: EncodeUrl is deprecated.
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// The information attached when the node cache purge type is set to purge_cache_tag.
	CacheTag *CacheTag `json:"CacheTag,omitnil,omitempty" name:"CacheTag"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Type of cache purging. Values:
	// <li>`purge_url`: Purge by the URL</li>
	// <li>`purge_prefix`: Purge by the directory</li>
	// <li>`purge_host`: Purge by the hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by the cache-tag </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Node cache purge method, valid for directory, hostname, and all cache refreshes. Valid values: <li>invalidate: Refreshes only resources that were updated under the directory; </li><li>delete: Refreshes all node resources, regardless of whether they were updated. </li>Default value: invalidate.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// List of resources for which cache is to be purged. Each element format depends on the cache purge type and you can refer to the API examples for details. <li>The number of tasks that can be submitted at a time is limited by the quota of a billing package. For details, see [Billing Overview] (https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1).</li>
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitnil,omitempty" name:"EncodeUrl"`

	// The information attached when the node cache purge type is set to purge_cache_tag.
	CacheTag *CacheTag `json:"CacheTag,omitnil,omitempty" name:"CacheTag"`
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
	delete(f, "CacheTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePurgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskResponseParams struct {
	// ID of the task.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// List of failed tasks and reasons.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailedList []*FailReason `json:"FailedList,omitnil,omitempty" name:"FailedList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateRealtimeLogDeliveryTaskRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The name of the real-time log delivery task, which is a combination of numbers, English letters, - and _, containing up to 200 characters.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// The type of the real-time log delivery task. Valid values:
	// <li>cls: Push to Tencent Cloud CLS;</li>
	// <li>custom_endpoint: Push to a user-defined HTTP(S) address;</li>
	// <li>s3: Push to an AWS S3-compatible bucket address.</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// The list of entities (Layer 7 domains or Layer 4 proxy instances) corresponding to the real-time log delivery task. Valid value examples:
	// <li>Layer 7 domain: domain.example.com;</li>
	// <li>Layer 4 proxy instance: sid-2s69eb5wcms7.</li>
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// The type of data delivery. Valid values:
	// <li>domain: Site acceleration log;</li>
	// <li>application: Layer 4 proxy log;</li>
	// <li>web-rateLiming: Rate limit and CC attack defense log;</li>
	// <li>web-attack: Managed rule log;</li>
	// <li>web-rule: Custom rule log;</li>
	// <li>web-bot: Bot management log.</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// The data delivery area. Valid values:
	// <li>mainland: Within the Chinese mainland;</li>
	// <li>overseas: Global (excluding the Chinese mainland).</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The list of predefined fields for delivery.
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// The list of custom fields for delivery, supporting extracting specified field values from HTTP request headers, response headers, and cookies. Each custom field name must be unique and the maximum number of fields is 200.
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// Log delivery filter conditions. If this field is not filled in, all logs will be delivered.
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// The sampling ratio in permille. Value range: 1 to 1000. For example, 605 represents a sampling ratio of 60.5%. If this field is not filled in, the sampling ratio is 100%.
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// Output format for log delivery. If this field is not specified, the default format is used, which works as follows:
	// <li>When TaskType is 'custom_endpoint', the default format is an array of JSON objects, with each JSON object representing a log entry;</li>
	// <li>When TaskType is 's3', the default format is JSON Lines;</li>Specifically, when TaskType is 'cls', the only allowed value for LogFormat.FormatType is 'json', and other parameters in LogFormat will be ignored. It is recommended not to transfer LogFormat.
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// The configuration information of CLS. This parameter is required when TaskType is cls.
	CLS *CLSTopic `json:"CLS,omitnil,omitempty" name:"CLS"`

	// The configuration information of the custom HTTP service. This parameter is required when TaskType is custom_endpoint.
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// The configuration information of the AWS S3-compatible bucket. This parameter is required when TaskType is s3.
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

type CreateRealtimeLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The name of the real-time log delivery task, which is a combination of numbers, English letters, - and _, containing up to 200 characters.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// The type of the real-time log delivery task. Valid values:
	// <li>cls: Push to Tencent Cloud CLS;</li>
	// <li>custom_endpoint: Push to a user-defined HTTP(S) address;</li>
	// <li>s3: Push to an AWS S3-compatible bucket address.</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// The list of entities (Layer 7 domains or Layer 4 proxy instances) corresponding to the real-time log delivery task. Valid value examples:
	// <li>Layer 7 domain: domain.example.com;</li>
	// <li>Layer 4 proxy instance: sid-2s69eb5wcms7.</li>
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// The type of data delivery. Valid values:
	// <li>domain: Site acceleration log;</li>
	// <li>application: Layer 4 proxy log;</li>
	// <li>web-rateLiming: Rate limit and CC attack defense log;</li>
	// <li>web-attack: Managed rule log;</li>
	// <li>web-rule: Custom rule log;</li>
	// <li>web-bot: Bot management log.</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// The data delivery area. Valid values:
	// <li>mainland: Within the Chinese mainland;</li>
	// <li>overseas: Global (excluding the Chinese mainland).</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The list of predefined fields for delivery.
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// The list of custom fields for delivery, supporting extracting specified field values from HTTP request headers, response headers, and cookies. Each custom field name must be unique and the maximum number of fields is 200.
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// Log delivery filter conditions. If this field is not filled in, all logs will be delivered.
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// The sampling ratio in permille. Value range: 1 to 1000. For example, 605 represents a sampling ratio of 60.5%. If this field is not filled in, the sampling ratio is 100%.
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// Output format for log delivery. If this field is not specified, the default format is used, which works as follows:
	// <li>When TaskType is 'custom_endpoint', the default format is an array of JSON objects, with each JSON object representing a log entry;</li>
	// <li>When TaskType is 's3', the default format is JSON Lines;</li>Specifically, when TaskType is 'cls', the only allowed value for LogFormat.FormatType is 'json', and other parameters in LogFormat will be ignored. It is recommended not to transfer LogFormat.
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// The configuration information of CLS. This parameter is required when TaskType is cls.
	CLS *CLSTopic `json:"CLS,omitnil,omitempty" name:"CLS"`

	// The configuration information of the custom HTTP service. This parameter is required when TaskType is custom_endpoint.
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// The configuration information of the AWS S3-compatible bucket. This parameter is required when TaskType is s3.
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

func (r *CreateRealtimeLogDeliveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRealtimeLogDeliveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskName")
	delete(f, "TaskType")
	delete(f, "EntityList")
	delete(f, "LogType")
	delete(f, "Area")
	delete(f, "Fields")
	delete(f, "CustomFields")
	delete(f, "DeliveryConditions")
	delete(f, "Sample")
	delete(f, "LogFormat")
	delete(f, "CLS")
	delete(f, "CustomEndpoint")
	delete(f, "S3")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRealtimeLogDeliveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRealtimeLogDeliveryTaskResponseParams struct {
	// The ID of the successfully created task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRealtimeLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRealtimeLogDeliveryTaskResponseParams `json:"Response"`
}

func (r *CreateRealtimeLogDeliveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRealtimeLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The rule name (1 to 255 characters)
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The rule name (1 to 255 characters)
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type CreateSecurityIPGroupRequestParams struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP group information.
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`
}

type CreateSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP group information.
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`
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
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateSharedCNAMERequestParams struct {
	// ID of the site to which the shared CNAME belongs.	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Shared CNAME prefix. Enter a valid domain name prefix, such as "test-api" and "test-api.com". A maximum of 50 characters are allowed. 
	// 
	// Complete format of the shared CNAME: '<Custom prefix>+<A 12-character random string in ZoneId>+share.dnse[0-5].com'. 
	// 
	// For example, if the prefix is example.com, EdgeOne will create the shared CNAME: example.com.sai2ig51kaa5.share.dnse2.com.
	SharedCNAMEPrefix *string `json:"SharedCNAMEPrefix,omitnil,omitempty" name:"SharedCNAMEPrefix"`

	// Description. It supports 1-50 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateSharedCNAMERequest struct {
	*tchttp.BaseRequest
	
	// ID of the site to which the shared CNAME belongs.	
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Shared CNAME prefix. Enter a valid domain name prefix, such as "test-api" and "test-api.com". A maximum of 50 characters are allowed. 
	// 
	// Complete format of the shared CNAME: '<Custom prefix>+<A 12-character random string in ZoneId>+share.dnse[0-5].com'. 
	// 
	// For example, if the prefix is example.com, EdgeOne will create the shared CNAME: example.com.sai2ig51kaa5.share.dnse2.com.
	SharedCNAMEPrefix *string `json:"SharedCNAMEPrefix,omitnil,omitempty" name:"SharedCNAMEPrefix"`

	// Description. It supports 1-50 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateSharedCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SharedCNAMEPrefix")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSharedCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSharedCNAMEResponseParams struct {
	// Shared CNAME. Format: '<Custom prefix>+<A 12-character random string in ZoneId>+share.dnse[0-5].com'.
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSharedCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *CreateSharedCNAMEResponseParams `json:"Response"`
}

func (r *CreateSharedCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneRequestParams struct {
	// Site access types. Options include:
	// <li>`partial`: (Default) Access through a CNAME record</li>
	// <li>`full`: Access through a name server</li>
	// <li>`noDomainAccess`: Access without using a domain name</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Site name. For sites connected via CNAME/NS, pass in the secondary domain name (example.com). Leave it blank if the site is connected without a domain name. 
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// The acceleration area of the L7 domain name when `Type` is `partial` or `full`. When Type is `noDomainAccess`, please leave it blank.
	// <li>`global`: Global AZs</li>
	// <li>`mainland`: AZs in the Chinese mainland</li>
	// <li>`overseas`: (Default) AZs outside the Chinese mainland </li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// ID of the plan to which you want to bind the site. If you don't have an EdgeOne plan, purchase one in the EdgeOne console.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// The site alias. It allows up to 20 characters, including [0-9], [a-z], [A-Z] and [-_]. For details, see [Glossary](https://intl.cloud.tencent.com/document/product/1552/70202?from_cn_redirect=1). If you don't want to use it, just leave it blank.
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// Tags of the site. To create tags, go to the [Tag Console](https://console.cloud.tencent.com/tag/taglist).
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to allow duplicate sites. Values:
	// <li>`true`: Duplicate sites are allowed.</li>
	// <li>`false`: Duplicate sites are not allowed.</li>If it is left empty, the default value `false` is used.
	//
	// Deprecated: AllowDuplicates is deprecated.
	AllowDuplicates *bool `json:"AllowDuplicates,omitnil,omitempty" name:"AllowDuplicates"`

	// Whether to skip scanning the existing DNS records of the site. Default value: false.
	//
	// Deprecated: JumpStart is deprecated.
	JumpStart *bool `json:"JumpStart,omitnil,omitempty" name:"JumpStart"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest
	
	// Site access types. Options include:
	// <li>`partial`: (Default) Access through a CNAME record</li>
	// <li>`full`: Access through a name server</li>
	// <li>`noDomainAccess`: Access without using a domain name</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Site name. For sites connected via CNAME/NS, pass in the secondary domain name (example.com). Leave it blank if the site is connected without a domain name. 
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// The acceleration area of the L7 domain name when `Type` is `partial` or `full`. When Type is `noDomainAccess`, please leave it blank.
	// <li>`global`: Global AZs</li>
	// <li>`mainland`: AZs in the Chinese mainland</li>
	// <li>`overseas`: (Default) AZs outside the Chinese mainland </li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// ID of the plan to which you want to bind the site. If you don't have an EdgeOne plan, purchase one in the EdgeOne console.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// The site alias. It allows up to 20 characters, including [0-9], [a-z], [A-Z] and [-_]. For details, see [Glossary](https://intl.cloud.tencent.com/document/product/1552/70202?from_cn_redirect=1). If you don't want to use it, just leave it blank.
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// Tags of the site. To create tags, go to the [Tag Console](https://console.cloud.tencent.com/tag/taglist).
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to allow duplicate sites. Values:
	// <li>`true`: Duplicate sites are allowed.</li>
	// <li>`false`: Duplicate sites are not allowed.</li>If it is left empty, the default value `false` is used.
	AllowDuplicates *bool `json:"AllowDuplicates,omitnil,omitempty" name:"AllowDuplicates"`

	// Whether to skip scanning the existing DNS records of the site. Default value: false.
	JumpStart *bool `json:"JumpStart,omitnil,omitempty" name:"JumpStart"`
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
	delete(f, "Type")
	delete(f, "ZoneName")
	delete(f, "Area")
	delete(f, "PlanId")
	delete(f, "AliasZoneName")
	delete(f, "Tags")
	delete(f, "AllowDuplicates")
	delete(f, "JumpStart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneResponseParams struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Site ownership verification information. After the site is created, you need to complete the ownership verification before the site can serve normally.
	// 
	// If `Type=partial`, add TXT records to your DNS provider or add files to the root DNS server, and then call [VerifyOwnership]() to complete verification. For more information, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
	// 
	// If `Type = full`, switch the DNS server as instructed by [Modifying DNS Server](https://intl.cloud.tencent.com/document/product/1552/90452?from_cn_redirect=1). Then call [VerifyOwnership]() to check the result.
	// 
	// If `Type = noDomainAccess`, leave it blank. No action is required.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CustomEndpoint struct {
	// The address of the custom HTTP interface for real-time log delivery. Currently, only HTTP and HTTPS protocols are supported.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Specifies the custom SecretId for generating an encrypted signature. This parameter is required if the origin server needs authentication.
	AccessId *string `json:"AccessId,omitnil,omitempty" name:"AccessId"`

	// Specifies the custom SecretKey for generating an encrypted signature. This parameter is required if the origin server needs authentication.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// The type of data compression. Valid values:<li>gzip: gzip compression.</li>If this field is not filled in, compression is disabled.
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`

	// The type of the application layer protocol used when POST requests log delivery. Valid values:
	// <li>http: HTTP protocol;</li>
	// <li>https: HTTPS protocol.</li>If this field is not filled in, the protocol type will be parsed from the URL field.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The custom request header carried during log delivery. If the header name you fill in is the default header carried by EdgeOne log delivery such as Content-Type, then the header value you fill in will override the default value. The header value supports referring to a single variable ${batchSize} to obtain the number of log entries included in each POST request.
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type CustomErrorPage struct {
	// Custom error page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom error page name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Custom error page type.
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// Custom error page description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Custom error page content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Custom error page reference.
	References []*ErrorPageReference `json:"References,omitnil,omitempty" name:"References"`
}

type CustomField struct {
	// Extracts data from specified positions in HTTP requests and responses. Valid values:
	// <li>ReqHeader: Extract a specified field value from an HTTP request header;</li>
	// <li>RspHeader: Extract a specified field value from an HTTP response header;</li>
	// <li>cookie: Extract a specified field value from a cookie.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Indicates the name of the parameter from which a value needs to be extracted, such as Accept-Language.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Indicates whether to deliver this field. If not filled in, this field will not be delivered.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type DDoS struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type DDoSAttackEvent struct {
	// The event ID.
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// The attack type.
	AttackType *string `json:"AttackType,omitnil,omitempty" name:"AttackType"`

	// The attack status.
	AttackStatus *int64 `json:"AttackStatus,omitnil,omitempty" name:"AttackStatus"`

	// The maximum attack bandwidth.
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitnil,omitempty" name:"AttackMaxBandWidth"`

	// The peak attack packet rate.
	AttackPacketMaxRate *int64 `json:"AttackPacketMaxRate,omitnil,omitempty" name:"AttackPacketMaxRate"`

	// The attack start time recorded in seconds.
	AttackStartTime *int64 `json:"AttackStartTime,omitnil,omitempty" name:"AttackStartTime"`

	// The attack end time recorded in seconds.
	AttackEndTime *int64 `json:"AttackEndTime,omitnil,omitempty" name:"AttackEndTime"`

	// The DDoS policy ID. 
	// Note: This field may return `null`, indicating that no valid value was found.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The site ID. 
	// Note: This field may return `null`, indicating that no valid value was found.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Geolocation scope. Values: 
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// Note: This field may return `null`, indicating that no valid value was found.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The blocking time of a DDoS attack. 
	// Note: This field may return `null`, indicating that no valid value was found.
	DDoSBlockData []*DDoSBlockData `json:"DDoSBlockData,omitnil,omitempty" name:"DDoSBlockData"`
}

type DDoSBlockData struct {
	// The start time recorded in UNIX timestamp.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time recorded in UNIX timestamp. `0` indicates the blocking is ongoing.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The regions blocked.
	BlockArea *string `json:"BlockArea,omitnil,omitempty" name:"BlockArea"`
}

type DDosProtectionConfig struct {
	// Dedicated anti-DDoS specifications in the Chinese mainland. For details, refer to [Dedicated Anti-DDoS Related Fees](https://intl.cloud.tencent.com/document/product/1552/94162?from_cn_redirect=1).
	// <li>PLATFORM: uses the default protection. Dedicated anti-DDoS is not enabled;</li>
	// <li>BASE30_MAX300: uses dedicated anti-DDoS, which provides 30 Gbps guaranteed protection bandwidth and up to 300 Gbps elastic protection bandwidth;</li>
	// <li>BASE60_MAX600: uses dedicated anti-DDoS, which provides 60 Gbps guaranteed protection bandwidth and up to 600 Gbps elastic protection bandwidth. </li>If this field is not specified, the default value 'PLATFORM' will be used.
	LevelMainland *string `json:"LevelMainland,omitnil,omitempty" name:"LevelMainland"`

	// Configuration of elastic protection bandwidth for exclusive DDoS protection in the Chinese mainland.Valid only when exclusive DDoS protection in the Chinese mainland is enabled (refer to the LevelMainland parameter configuration), and the value has the following limitations:<li>When exclusive DDoS protection is enabled in the Chinese mainland and the 30 Gbps baseline protection bandwidth is used (the LevelMainland parameter value is BASE30_MAX300): the value range is 30 to 300 in Gbps;</li><li>When exclusive DDoS protection is enabled in the Chinese mainland and the 60 Gbps baseline protection bandwidth is used (the LevelMainland parameter value is BASE60_MAX600): the value range is 60 to 600 in Gbps;</li><li>When the default protection of the platform is used (the LevelMainland parameter value is PLATFORM): configuration is not supported, and the value of this parameter is invalid.</li>
	MaxBandwidthMainland *uint64 `json:"MaxBandwidthMainland,omitnil,omitempty" name:"MaxBandwidthMainland"`

	// Dedicated anti-DDoS specifications in global regions (excluding the Chinese mainland).
	// <li>PLATFORM: uses the default protection. Dedicated anti-DDoS is not enabled;</li>
	// <li>ANYCAST300: uses dedicated anti-DDoS, which provides 300 Gbps protection bandwidth;</li>
	// <li>ANYCAST_ALLIN: uses dedicated anti-DDoS, which provides all available protection resources. </li>If this field is not specified, the default value 'PLATFORM' will be used.
	LevelOverseas *string `json:"LevelOverseas,omitnil,omitempty" name:"LevelOverseas"`
}

type DefaultServerCertInfo struct {
	// ID of the server certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Alias of the certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Type of the certificate. Values:
	// <li>`default`: Default certificate;</li>
	// <li>`upload`: Custom certificate;</li>
	// <li>`managed`: Tencent Cloud-managed certificate.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Time when the certificate expires.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Time when the certificate takes effect.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// Common name of the certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`

	// Domain names added to the SAN certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// Deployment status. Values:
	// <li>`processing`: Deployment in progress</li>
	// <li>`deployed`: Deployed</li>
	// <li>`failed`: Deployment failed</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Failure description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Certificate algorithm.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SignAlgo *string `json:"SignAlgo,omitnil,omitempty" name:"SignAlgo"`
}

// Predefined struct for user
type DeleteAccelerationDomainsRequestParams struct {
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be deleted.
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// Whether to forcibly delete a domain name if it is associated with resources (such as alias domain names and traffic scheduling policies). 
	// <li>`true`: Delete the domain name and all associated resources.</li>
	// <li>`false`: Do not delete the domain name and all associated resources.</li>If it’s not specified, the default value `false` is used.
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeleteAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be deleted.
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// Whether to forcibly delete a domain name if it is associated with resources (such as alias domain names and traffic scheduling policies). 
	// <li>`true`: Delete the domain name and all associated resources.</li>
	// <li>`false`: Do not delete the domain name and all associated resources.</li>If it’s not specified, the default value `false` is used.
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The alias domain name to be deleted. If it is left empty, the delete operation is not performed.
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
}

type DeleteAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The alias domain name to be deleted. If it is left empty, the delete operation is not performed.
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

type DeleteApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteCustomErrorPageRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom response page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type DeleteCustomErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom response page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

func (r *DeleteCustomErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomErrorPageResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomErrorPageResponseParams `json:"Response"`
}

func (r *DeleteCustomErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomErrorPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

type DeleteL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

func (r *DeleteL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL4ProxyResponseParams `json:"Response"`
}

func (r *DeleteL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyRulesRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rule IDs. It supports up to 200 forwarding rules at a time.
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rule IDs. It supports up to 200 forwarding rules at a time.
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL4ProxyRulesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL4ProxyRulesResponseParams `json:"Response"`
}

func (r *DeleteL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Origin server group ID. This parameter is required.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Origin server group ID. This parameter is required.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteRealtimeLogDeliveryTaskRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The ID of the real-time log delivery task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteRealtimeLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The ID of the real-time log delivery task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteRealtimeLogDeliveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRealtimeLogDeliveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRealtimeLogDeliveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRealtimeLogDeliveryTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRealtimeLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRealtimeLogDeliveryTaskResponseParams `json:"Response"`
}

func (r *DeleteRealtimeLogDeliveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRealtimeLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IDs of the rules to be deleted.
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IDs of the rules to be deleted.
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteSharedCNAMERequestParams struct {
	// ID of the site to which the shared CNAME belongs.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The shared CNAME to be deleted
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`
}

type DeleteSharedCNAMERequest struct {
	*tchttp.BaseRequest
	
	// ID of the site to which the shared CNAME belongs.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The shared CNAME to be deleted
	SharedCNAME *string `json:"SharedCNAME,omitnil,omitempty" name:"SharedCNAME"`
}

func (r *DeleteSharedCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSharedCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SharedCNAME")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSharedCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSharedCNAMEResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSharedCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSharedCNAMEResponseParams `json:"Response"`
}

func (r *DeleteSharedCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSharedCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DeliveryCondition struct {
	// Log filter conditions. Detailed filter conditions are as follows:
	// <li>EdgeResponseStatusCode: Filter by status code returned to the client by the EdgeOne node.<br>      Supported operators: equal, great, less, great_equal, less_equal<br>      Value range: Any integer greater than or equal to 0</li>
	// <li>OriginResponseStatusCode: Filter by response status code of the origin server.<br>      Supported operators: equal, great, less, great_equal, less_equal<br>      Value range: Any integer greater than or equal to -1</li>
	// <li>SecurityAction: Filter by final action after the request hits the security rule.<br>      Supported operators: equal<br>      Options are as follows:<br>      -: Unknown/unhit<br>      Monitor: Log only<br>      JSChallenge: JavaScript challenge<br>      Deny: Interception<br>      Allow: Allow<br>      BlockIP: IP ban<br>      Redirect: Redirect<br>      ReturnCustomPage: Return to custom page<br>      ManagedChallenge: Managed challenge<br>      Silence: Silence<br>      LongDelay: Response after a long wait<br>      ShortDelay: Response after a short wait</li>
	// <li>SecurityModule: Filter according to the name of the security module that ultimately handles the request. <br>      Supported operators: equal<br>      Options are as follows:<br>      -: Unknown/miss<br>      CustomRule: Web protection - custom rules<br>      RateLimitingCustomRule: Web protection - rate limiting rule<br>      ManagedRule: Web Protection - Managed Rules<br>      L7DDoS: Web Protection - CC Attack Protection<br>      BotManagement: Bot Management - Bot Basic Management<br>      BotClientReputation: Bot Management - Client Reputation Analysis<br>      BotBehaviorAnalysis: Bot Management - Bot Intelligence Analysis<br>      BotCustomRule: Bot Management - Custom Bot Rules<br>      BotActiveDetection: Bot Management - Active Detection</li>
	Conditions []*QueryCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

// Predefined struct for user
type DeployConfigGroupVersionRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Environment ID. Please specify the environment ID to which the version should be released.
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Version information required for release. Multiple versions of different configuration groups can be modified simultaneously, while each group allows modifying only one version at a time.
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// Change description. It is used to describe the content and reasons for this change. A maximum of 100 characters are supported.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type DeployConfigGroupVersionRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Environment ID. Please specify the environment ID to which the version should be released.
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Version information required for release. Multiple versions of different configuration groups can be modified simultaneously, while each group allows modifying only one version at a time.
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// Change description. It is used to describe the content and reasons for this change. A maximum of 100 characters are supported.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *DeployConfigGroupVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployConfigGroupVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "EnvId")
	delete(f, "ConfigGroupVersionInfos")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployConfigGroupVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployConfigGroupVersionResponseParams struct {
	// Release record ID.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeployConfigGroupVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeployConfigGroupVersionResponseParams `json:"Response"`
}

func (r *DeployConfigGroupVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployConfigGroupVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployRecord struct {
	// Details about the released version.
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// Release time. The time format follows the ISO 8601 standard and is represented in Coordinated Universal Time (UTC).
	DeployTime *string `json:"DeployTime,omitnil,omitempty" name:"DeployTime"`

	// Release status. Valid values: 
	// <li>deploying: Being released.</li>
	// <li>failure: Release failed.</li>
	// <li>success: Released successfully. </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Release result information.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Release record ID. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Change description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type DescribeAccelerationDomainsRequestParams struct {
	// ID of the site related with the acceleration domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value of Filters.Values is 20. If it is not specified, all domains related with the specific zone-id are returned.
	// <li>domain-name: Filtering based on the acceleration domain name</li>
	// <li>origin-type: Filtering based on the type of the origin server</li>
	// <li>origin: Filtering based on the primary origin server address</li>
	// <li>backup-origin: Filtering based on the secondary origin server address</li>
	// <li>domain-cname: Filtering based on the CNAME</li>
	// <li>share-cname: Filtering based on the shared CNAME</li>
	// <li>vodeo-sub-app-id: Filtering based on [vodeo sub-application ID]</li>
	// <li>vodeo-distribution-range: Filtering based on [vodeo distribution range]</li>
	// <li>vodeo-bucket-id: Filtering based on [vodeo storage bucket ID];</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. Values include:
	// <li>`created_on`: Creation time of the acceleration domain name</li>
	// <li>`domain-name`: (Default) Acceleration domain name.</li> 
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sort direction. If the field value is number, sort by the numeric value. If the field value is text, sort by the ascill code. Values include:
	// <li>`asc`: Ascending order.</li>
	// <li>`desc`: Descending order.</li> Default value: `asc`.
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// The match mode. Values:
	// <li>`all`: Return all matches.</li>
	// <li>`any`: Return any match.</li>Default value: `all`.
	Match *string `json:"Match,omitnil,omitempty" name:"Match"`
}

type DescribeAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the acceleration domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value of Filters.Values is 20. If it is not specified, all domains related with the specific zone-id are returned.
	// <li>domain-name: Filtering based on the acceleration domain name</li>
	// <li>origin-type: Filtering based on the type of the origin server</li>
	// <li>origin: Filtering based on the primary origin server address</li>
	// <li>backup-origin: Filtering based on the secondary origin server address</li>
	// <li>domain-cname: Filtering based on the CNAME</li>
	// <li>share-cname: Filtering based on the shared CNAME</li>
	// <li>vodeo-sub-app-id: Filtering based on [vodeo sub-application ID]</li>
	// <li>vodeo-distribution-range: Filtering based on [vodeo distribution range]</li>
	// <li>vodeo-bucket-id: Filtering based on [vodeo storage bucket ID];</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. Values include:
	// <li>`created_on`: Creation time of the acceleration domain name</li>
	// <li>`domain-name`: (Default) Acceleration domain name.</li> 
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sort direction. If the field value is number, sort by the numeric value. If the field value is text, sort by the ascill code. Values include:
	// <li>`asc`: Ascending order.</li>
	// <li>`desc`: Descending order.</li> Default value: `asc`.
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// The match mode. Values:
	// <li>`all`: Return all matches.</li>
	// <li>`any`: Return any match.</li>Default value: `all`.
	Match *string `json:"Match,omitnil,omitempty" name:"Match"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "Direction")
	delete(f, "Match")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerationDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerationDomainsResponseParams struct {
	// Total of matched alias domain names.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of all matched acceleration domain names
	AccelerationDomains []*AccelerationDomain `json:"AccelerationDomains,omitnil,omitempty" name:"AccelerationDomains"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value for Filters.Values is 20. The detailed conditions are as follows:
	// <li>target-name: Filter by the target domain name;</li>
	// <li>alias-name: Filter by the alias of the domain name.</li>Fuzzy queries are only supported for the field name alias-name.
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAliasDomainsRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value for Filters.Values is 20. The detailed conditions are as follows:
	// <li>target-name: Filter by the target domain name;</li>
	// <li>alias-name: Filter by the alias of the domain name.</li>Fuzzy queries are only supported for the field name alias-name.
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of the eligible alias domain names.
	AliasDomains []*AliasDomain `json:"AliasDomains,omitnil,omitempty" name:"AliasDomains"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Zone ID. This parameter is required.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filters. Each filter can have up to 20 entries. Details: <li>proxy-id<br>   Filter by the <strong>Proxy ID</strong>, such as: `proxy-ev2sawbwfd`. <br>   Type: String<br>   Required: No</li><li>zone-id<br>   Filter by the <strong>Site ID</strong>, such as `zone-vawer2vadg`. <br>   Type: String<br>   Required: No</li><li>rule-tag<br>   Filter by the <strong>Rule tag</strong>, such as `rule-service-1`. <br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The paginated query offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeApplicationProxiesRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID. This parameter is required.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filters. Each filter can have up to 20 entries. Details: <li>proxy-id<br>   Filter by the <strong>Proxy ID</strong>, such as: `proxy-ev2sawbwfd`. <br>   Type: String<br>   Required: No</li><li>zone-id<br>   Filter by the <strong>Site ID</strong>, such as `zone-vawer2vadg`. <br>   Type: String<br>   Required: No</li><li>rule-tag<br>   Filter by the <strong>Rule tag</strong>, such as `rule-service-1`. <br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The paginated query offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesResponseParams struct {
	// List of application proxies.
	ApplicationProxies []*ApplicationProxy `json:"ApplicationProxies,omitnil,omitempty" name:"ApplicationProxies"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PlanInfo []*PlanInfo `json:"PlanInfo,omitnil,omitempty" name:"PlanInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBillingDataRequestParams struct {
	// Start time of the query.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Zone ID set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Indicator list. Valid values:<li>acc_flux: Content acceleration traffic, in bytes;</li>
	// <li>smt_flux: Smart acceleration traffic, in bytes;</li>
	// <li>l4_flux: Layer 4 acceleration traffic, in bytes;</li>
	// <li>sec_flux: Exclusive DDoS mitigation traffic, in bytes;</li>
	// <li>zxctg_flux: Cross-MLC-border acceleration traffic, in bytes;</li>
	// <li>acc_bandwidth: Content acceleration bandwidth, in bps;</li>
	// <li>smt_bandwidth: Smart acceleration bandwidth, in bps;</li>
	// <li>l4_bandwidth: Layer 4 acceleration bandwidth, in bps;</li>
	// <li>sec_bandwidth: Exclusive DDoS mitigation bandwidth, in bps;</li>
	// <li>zxctg_bandwidth: Cross-MLC-border acceleration bandwidth, in bps;</li><li>sec_request_clean: HTTP/HTTPS requests, in count;</li>
	// <li>smt_request_clean: Smart acceleration requests, in count;</li>
	// <li>quic_request: QUIC requests, in count;</li>
	// <li>bot_request_clean: Bot requests, in count;</li>
	// <li>cls_count: Real-time log delivery log counts, in count;</li>
	// <li>ddos_bandwidth: Exclusive DDoS mitigation elastic bandwidth, in bps.</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Query granularity. Valid values:<li>5min: 5-minute granularity;</li>
	// <li>hour: 1-hour granularity;</li>
	// <li>day: 1-day granularity.</li>
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filter. The detailed filters are as follows:
	// <li>host: Filter by Domain Name. Example: test.example.com.<br></li>
	// <li>proxy-id: Filter by Layer 4 Proxy Instance ID. Example: sid-2rugn89bkla9.<br></li>
	// <li>region-id: Filter by Billing Region.Options are:<br>   CH: Chinese mainland<br>   AF: Africa<br>   AS1: Asia-Pacific Region 1<br>   AS2: Asia-Pacific Region 2<br>   AS3: Asia-Pacific Region 3<br>  EU: Europe<br>   MidEast: Middle East<br>   NA: North America<br>   SA: South America</li>
	Filters []*BillingDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Zone ID set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Indicator list. Valid values:<li>acc_flux: Content acceleration traffic, in bytes;</li>
	// <li>smt_flux: Smart acceleration traffic, in bytes;</li>
	// <li>l4_flux: Layer 4 acceleration traffic, in bytes;</li>
	// <li>sec_flux: Exclusive DDoS mitigation traffic, in bytes;</li>
	// <li>zxctg_flux: Cross-MLC-border acceleration traffic, in bytes;</li>
	// <li>acc_bandwidth: Content acceleration bandwidth, in bps;</li>
	// <li>smt_bandwidth: Smart acceleration bandwidth, in bps;</li>
	// <li>l4_bandwidth: Layer 4 acceleration bandwidth, in bps;</li>
	// <li>sec_bandwidth: Exclusive DDoS mitigation bandwidth, in bps;</li>
	// <li>zxctg_bandwidth: Cross-MLC-border acceleration bandwidth, in bps;</li><li>sec_request_clean: HTTP/HTTPS requests, in count;</li>
	// <li>smt_request_clean: Smart acceleration requests, in count;</li>
	// <li>quic_request: QUIC requests, in count;</li>
	// <li>bot_request_clean: Bot requests, in count;</li>
	// <li>cls_count: Real-time log delivery log counts, in count;</li>
	// <li>ddos_bandwidth: Exclusive DDoS mitigation elastic bandwidth, in bps.</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Query granularity. Valid values:<li>5min: 5-minute granularity;</li>
	// <li>hour: 1-hour granularity;</li>
	// <li>day: 1-day granularity.</li>
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filter. The detailed filters are as follows:
	// <li>host: Filter by Domain Name. Example: test.example.com.<br></li>
	// <li>proxy-id: Filter by Layer 4 Proxy Instance ID. Example: sid-2rugn89bkla9.<br></li>
	// <li>region-id: Filter by Billing Region.Options are:<br>   CH: Chinese mainland<br>   AF: Africa<br>   AS1: Asia-Pacific Region 1<br>   AS2: Asia-Pacific Region 2<br>   AS3: Asia-Pacific Region 3<br>  EU: Europe<br>   MidEast: Middle East<br>   NA: North America<br>   SA: South America</li>
	Filters []*BillingDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeBillingDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "MetricName")
	delete(f, "Interval")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataResponseParams struct {
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*BillingData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingDataResponseParams `json:"Response"`
}

func (r *DescribeBillingDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigGroupVersionDetailRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Version ID.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DescribeConfigGroupVersionDetailRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Version ID.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DescribeConfigGroupVersionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigGroupVersionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigGroupVersionDetailResponseParams struct {
	// Version information.
	ConfigGroupVersionInfo *ConfigGroupVersionInfo `json:"ConfigGroupVersionInfo,omitnil,omitempty" name:"ConfigGroupVersionInfo"`

	// Version file content. It is returned in JSON format.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigGroupVersionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigGroupVersionDetailResponseParams `json:"Response"`
}

func (r *DescribeConfigGroupVersionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigGroupVersionsRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Configuraration group ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Filtering condition. The maximum value of Filters.Values is 20. If this parameter is not specified, all version information for the selected configuration group is returned. Detailed filtering conditions: 
	// <li>version-id: Filter by version ID.</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Paging query offset. The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limited entries in paging queries. The default value is 20 and the maximum value is 100. 
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeConfigGroupVersionsRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Configuraration group ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Filtering condition. The maximum value of Filters.Values is 20. If this parameter is not specified, all version information for the selected configuration group is returned. Detailed filtering conditions: 
	// <li>version-id: Filter by version ID.</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Paging query offset. The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limited entries in paging queries. The default value is 20 and the maximum value is 100. 
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeConfigGroupVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigGroupVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigGroupVersionsResponseParams struct {
	// Total versions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Version information list.
	ConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"ConfigGroupVersionInfos,omitnil,omitempty" name:"ConfigGroupVersionInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigGroupVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigGroupVersionsResponseParams `json:"Response"`
}

func (r *DescribeConfigGroupVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigGroupVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeContentQuotaRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
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
	PurgeQuota []*Quota `json:"PurgeQuota,omitnil,omitempty" name:"PurgeQuota"`

	// Pre-warming quotas.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrefetchQuota []*Quota `json:"PrefetchQuota,omitnil,omitempty" name:"PrefetchQuota"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCustomErrorPagesRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filter criteria. The upper limit for Filters.Values is 20. The detailed filter criteria are as follows:
	// <li>page-id: Filter by page ID;</li>
	// <li>name: Filter by page name;</li>
	// <li>description: Filter by page description;</li>
	// <li>content-type: Filter by page type.</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset of paginated query. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The limit of paginated query. Default value: 20. Maximum value: 1,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCustomErrorPagesRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filter criteria. The upper limit for Filters.Values is 20. The detailed filter criteria are as follows:
	// <li>page-id: Filter by page ID;</li>
	// <li>name: Filter by page name;</li>
	// <li>description: Filter by page description;</li>
	// <li>content-type: Filter by page type.</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The offset of paginated query. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The limit of paginated query. Default value: 20. Maximum value: 1,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCustomErrorPagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomErrorPagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomErrorPagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomErrorPagesResponseParams struct {
	// Total number of custom response pages.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Custom response page data list.
	ErrorPages []*CustomErrorPage `json:"ErrorPages,omitnil,omitempty" name:"ErrorPages"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomErrorPagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomErrorPagesResponseParams `json:"Response"`
}

func (r *DescribeCustomErrorPagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomErrorPagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataRequestParams struct {
	// Start time of the query period.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Statistical metrics.
	// <li>`ddos_attackMaxBandwidth`: Peak attack bandwidth;</li>
	// <li>`ddos_attackMaxPackageRate`: Peak attack packet rate;</li>
	// <li>`ddos_attackBandwidth`: Time-series data of attack bandwidth;</li>
	// <li>`ddos_attackPackageRate`: Time-series data of attack packet rate.</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// IDs of DDoS policies to be queried. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day</li>If this field is not specified, the granularity is determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < Period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global </li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeDDoSAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query period.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Statistical metrics.
	// <li>`ddos_attackMaxBandwidth`: Peak attack bandwidth;</li>
	// <li>`ddos_attackMaxPackageRate`: Peak attack packet rate;</li>
	// <li>`ddos_attackBandwidth`: Time-series data of attack bandwidth;</li>
	// <li>`ddos_attackPackageRate`: Time-series data of attack packet rate.</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// IDs of DDoS policies to be queried. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day</li>If this field is not specified, the granularity is determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < Period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global </li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of DDoS attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Start time. Time range: 30 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. Time range: 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// List of DDoS policy IDs. All policies are selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Parameter to show attack details. If it is configured as false, only the number of attacks is returned without details. If it is configured as true, attack details are returned.
	ShowDetail *bool `json:"ShowDetail,omitnil,omitempty" name:"ShowDetail"`

	// Geolocation scope. Values: 
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The sorting field. Values: 
	// <li>`MaxBandWidth`: Peak bandwidth</li>
	// <li>`AttackStartTime` Start time of the attack</li>If this field is not specified, the default value `AttackStartTime` is used.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The sorting method. Values: 
	// <Li>`asc`: Ascending</li>
	// <li>`desc`: Descending</li>If this field is not specified, the default value `desc` is used.
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeDDoSAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// Start time. Time range: 30 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. Time range: 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// List of DDoS policy IDs. All policies are selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Parameter to show attack details. If it is configured as false, only the number of attacks is returned without details. If it is configured as true, attack details are returned.
	ShowDetail *bool `json:"ShowDetail,omitnil,omitempty" name:"ShowDetail"`

	// Geolocation scope. Values: 
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The sorting field. Values: 
	// <li>`MaxBandWidth`: Peak bandwidth</li>
	// <li>`AttackStartTime` Start time of the attack</li>If this field is not specified, the default value `AttackStartTime` is used.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The sorting method. Values: 
	// <Li>`asc`: Ascending</li>
	// <li>`desc`: Descending</li>If this field is not specified, the default value `desc` is used.
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
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
	Data []*DDoSAttackEvent `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The statistical metric. Values:
	// <li>`ddos_attackFlux_protocol`: Rank protocols by the attack traffic.</li>
	// <li>`ddos_attackPackageNum_protocol`: Rank protocols by the number of attack packets.</li>
	// <li>`ddos_attackNum_attackType`: Rank attack types by the number of attacks.</li>
	// <li>`ddos_attackNum_sregion`: Rank attacker regions by the number of attacks.</li>
	// <li>`ddos_attackFlux_sip`: Rank attacker IPs by the number of attacks.</li>
	// <li>`ddos_attackFlux_sregion`: Rank attacker regions by the number of attacks.</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Site ID set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// The attack type. Values:
	// <li>`flood`: Flood;</li>
	// <li>`icmpFlood`: ICMP flood;</li>
	// <li>`all`: All attack types.</li>This field will be set to the default value `all` if not specified.
	AttackType *string `json:"AttackType,omitnil,omitempty" name:"AttackType"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// The port number.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Queries the top n rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeDDoSAttackTopDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The statistical metric. Values:
	// <li>`ddos_attackFlux_protocol`: Rank protocols by the attack traffic.</li>
	// <li>`ddos_attackPackageNum_protocol`: Rank protocols by the number of attack packets.</li>
	// <li>`ddos_attackNum_attackType`: Rank attack types by the number of attacks.</li>
	// <li>`ddos_attackNum_sregion`: Rank attacker regions by the number of attacks.</li>
	// <li>`ddos_attackFlux_sip`: Rank attacker IPs by the number of attacks.</li>
	// <li>`ddos_attackFlux_sregion`: Rank attacker regions by the number of attacks.</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Site ID set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// The attack type. Values:
	// <li>`flood`: Flood;</li>
	// <li>`icmpFlood`: ICMP flood;</li>
	// <li>`all`: All attack types.</li>This field will be set to the default value `all` if not specified.
	AttackType *string `json:"AttackType,omitnil,omitempty" name:"AttackType"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// The port number.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Queries the top n rows of data. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	Data []*TopEntry `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filter criteria. Each filter criteria can have up to 5 entries.
	// <li>`zone-id`: <br>Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filter criteria. Each filter criteria can have up to 5 entries.
	// <li>`zone-id`: <br>Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of default certificates
	DefaultServerCertInfo []*DefaultServerCertInfo `json:"DefaultServerCertInfo,omitnil,omitempty" name:"DefaultServerCertInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDeployHistoryRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Environment ID.
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Filtering condition. The maximum value of Filters.Values is 20. Detailed filtering conditions: 
	// <li>record-id: Filter by release record ID. </li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDeployHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Environment ID.
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Filtering condition. The maximum value of Filters.Values is 20. Detailed filtering conditions: 
	// <li>record-id: Filter by release record ID. </li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDeployHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "EnvId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeployHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployHistoryResponseParams struct {
	// Total release records.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Release record details.
	Records []*DeployRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeployHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeployHistoryResponseParams `json:"Response"`
}

func (r *DescribeDeployHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// Total environments.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Environment list.
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value for Filters.Values is 20. The detailed conditions are as follows:<li>host: Filter by domain name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeHostsSettingRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value for Filters.Values is 20. The detailed conditions are as follows:<li>host: Filter by domain name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	DetailHosts []*DetailHost `json:"DetailHosts,omitnil,omitempty" name:"DetailHosts"`

	// Number of domain names
	TotalNumber *int64 `json:"TotalNumber,omitnil,omitempty" name:"TotalNumber"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeIPRegionRequestParams struct {
	// List of IPs to be queried, supporting IPV4 and IPV6. Up to 100 can be queried.
	IPs []*string `json:"IPs,omitnil,omitempty" name:"IPs"`
}

type DescribeIPRegionRequest struct {
	*tchttp.BaseRequest
	
	// List of IPs to be queried, supporting IPV4 and IPV6. Up to 100 can be queried.
	IPs []*string `json:"IPs,omitnil,omitempty" name:"IPs"`
}

func (r *DescribeIPRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IPs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPRegionResponseParams struct {
	// List of IP attribution information.
	IPRegionInfo []*IPRegionInfo `json:"IPRegionInfo,omitnil,omitempty" name:"IPRegionInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIPRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPRegionResponseParams `json:"Response"`
}

func (r *DescribeIPRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationsRequestParams struct {
	// Filter conditions. The maximum value for Filters.Values is 20. The detailed conditions are as follows:<li>zone-name: Filter by site name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeIdentificationsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions. The maximum value for Filters.Values is 20. The detailed conditions are as follows:<li>zone-name: Filter by site name.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The site verification information.
	Identifications []*Identification `json:"Identifications,omitnil,omitempty" name:"Identifications"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeL4ProxyRequestParams struct {
	// ID of the zone where the Layer 4 proxy instance belongs.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Paginated query offset. Defaults to 0 when not specified.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria. The upper limit for Filters.Values is 20. If left empty, all Layer 4 proxy instance information under the current zone ID is returned. The detailed filter criteria are as follows: <li>proxy-id: Filters by Layer 4 proxy instance ID;</li>
	// <li>ddos-protection-type: Filters by security protection type;</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the zone where the Layer 4 proxy instance belongs.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Paginated query offset. Defaults to 0 when not specified.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria. The upper limit for Filters.Values is 20. If left empty, all Layer 4 proxy instance information under the current zone ID is returned. The detailed filter criteria are as follows: <li>proxy-id: Filters by Layer 4 proxy instance ID;</li>
	// <li>ddos-protection-type: Filters by security protection type;</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ProxyResponseParams struct {
	// The number of Layer 4 proxy instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of Layer 4 proxy instances.
	L4Proxies []*L4Proxy `json:"L4Proxies,omitnil,omitempty" name:"L4Proxies"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4ProxyResponseParams `json:"Response"`
}

func (r *DescribeL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ProxyRulesRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Paginated query offset. Defaults to 0 when not specified.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria. The upper limit for Filters.Values is 20. All rule information under the current Layer 4 instance will be returned if left empty. The detailed filter criteria are as follows: <li>rule-tag: Filters rules under the Layer 4 proxy instance according to rule tag.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Paginated query offset. Defaults to 0 when not specified.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter criteria. The upper limit for Filters.Values is 20. All rule information under the current Layer 4 instance will be returned if left empty. The detailed filter criteria are as follows: <li>rule-tag: Filters rules under the Layer 4 proxy instance according to rule tag.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ProxyRulesResponseParams struct {
	// The total count of forwarding rules.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of forwarding rules.	
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4ProxyRulesResponseParams `json:"Response"`
}

func (r *DescribeL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupRequestParams struct {
	// (Required) Site ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The paginated query offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Value range: 1-1000. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters. Each filter can have up to 20 entries. See below for details:
	// <li>`origin-group-id`<br>Filter by the <strong>origin group ID</strong>. Format: `origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a`<br>Fuzzy query is not supported</li><li>`origin-group-name`<br>Filter by the <strong>origin group name</strong><br>Fuzzy query is supported. When fuzzy query is used, only one origin groupsource site group name is supported</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// (Required) Site ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The paginated query offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Value range: 1-1000. Default value: 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters. Each filter can have up to 20 entries. See below for details:
	// <li>`origin-group-id`<br>Filter by the <strong>origin group ID</strong>. Format: `origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a`<br>Fuzzy query is not supported</li><li>`origin-group-name`<br>Filter by the <strong>origin group name</strong><br>Fuzzy query is supported. When fuzzy query is used, only one origin groupsource site group name is supported</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ZoneId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Origin group information.
	OriginGroups []*OriginGroup `json:"OriginGroups,omitnil,omitempty" name:"OriginGroups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// Set of site IDs to be queried. This is a required parameter.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`need-update`:<br>   Whether <strong>the intermediate IP update is needed for the site</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Update needed.<br>   `false`: Update not needed.<br></li>
	// <li>`plan-support`:<br>   Whether <strong>origin protection is supported in the plan</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Origin protection supported.<br>   `false`: Origin protection not supported.<br></li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeOriginProtectionRequest struct {
	*tchttp.BaseRequest
	
	// Set of site IDs to be queried. This is a required parameter.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`need-update`:<br>   Whether <strong>the intermediate IP update is needed for the site</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Update needed.<br>   `false`: Update not needed.<br></li>
	// <li>`plan-support`:<br>   Whether <strong>origin protection is supported in the plan</strong>.<br>   Type: String<br>   Required: No<br>   Values:<br>   `true`: Origin protection supported.<br>   `false`: Origin protection not supported.<br></li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	OriginProtectionInfo []*OriginProtectionInfo `json:"OriginProtectionInfo,omitnil,omitempty" name:"OriginProtectionInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The metric to query. Values:
	// <li>`l7Flow_outFlux`: Traffic used for EdegOne responses</li>
	// <li>`l7Flow_inFlux`: Traffic used for EdegOne requests</li>
	// <li>`l7Flow_outBandwidth`: Bandwidth used for EdegOne responses</li>
	// <li>`l7Flow_inBandwidth`: Bandwidth used for EdegOne requests</li>
	// <li>`l7Flow_hit_outFlux`: Traffic used for cache hit</li>
	// <li>`l7Flow_request`: Access requests</li>
	// <li>`l7Flow_flux`: Upstream and downstream traffic used for client access</li>
	// <li>`l7Flow_bandwidth`: Upstream and downstream bandwidth used for client access</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Queried domain name set. This parameter has been discarded.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`: All protocols. </li>If it's not specified, `all` is used. This parameter is not yet available now.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: One day</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filtering condition. The detailed filtering condition key values are as follows: 
	// <li>socket: Filter based on HTTP protocol type. Corresponding value options: <br> HTTP: HTTP protocol; <br> HTTPS: HTTPS protocol; <br> QUIC: QUIC protocol. </li>
	// <li>domains: Filter based on domain name. </li>
	// <li>tagKey: Filter based on Tag Key. </li>
	// <li>tagValue: Filter based on Tag Value. </li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeOverviewL7DataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The metric to query. Values:
	// <li>`l7Flow_outFlux`: Traffic used for EdegOne responses</li>
	// <li>`l7Flow_inFlux`: Traffic used for EdegOne requests</li>
	// <li>`l7Flow_outBandwidth`: Bandwidth used for EdegOne responses</li>
	// <li>`l7Flow_inBandwidth`: Bandwidth used for EdegOne requests</li>
	// <li>`l7Flow_hit_outFlux`: Traffic used for cache hit</li>
	// <li>`l7Flow_request`: Access requests</li>
	// <li>`l7Flow_flux`: Upstream and downstream traffic used for client access</li>
	// <li>`l7Flow_bandwidth`: Upstream and downstream bandwidth used for client access</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Queried domain name set. This parameter has been discarded.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`: All protocols. </li>If it's not specified, `all` is used. This parameter is not yet available now.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: One day</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period <= 1 hour: `min`; <br>1 hour < Period <= 2 days: `5min`; <br>2 days < period <= 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filtering condition. The detailed filtering condition key values are as follows: 
	// <li>socket: Filter based on HTTP protocol type. Corresponding value options: <br> HTTP: HTTP protocol; <br> HTTPS: HTTPS protocol; <br> QUIC: QUIC protocol. </li>
	// <li>domains: Filter based on domain name. </li>
	// <li>tagKey: Filter based on Tag Key. </li>
	// <li>tagValue: Filter based on Tag Value. </li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Note: This field may return 'null', which indicates a failure to obtain a valid value.
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// ZoneId. The parameter is required.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Start time of the query. Either time or job-id is required.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query. Either time or job-id is required.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filtering condition. The maximum value of Filters.Values is 20. Detailed filtering conditions: 
	// <li>job-id 
	// Filter based on task ID.  job-id format: 1379afjk91u32h. Multiple values are not supported, Fuzzy search is not supported. </li>
	// <li>target
	// Filter based on target resource information. target format: http://www.qq.com/1.txt. Multiple values are not supported, Fuzzy search is not supported. </li>
	// <li>domains 
	// Filter based on domain name. domains format: www.qq.com. Fuzzy search is not supported.</li>
	// <li>statuses
	// Filter based on task status.  Fuzzy search is not supported. Options:<br> processing: Processing<br> success: Success<br> failed: Failure<br> timeout: Timeout</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// ZoneId. The parameter is required.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Start time of the query. Either time or job-id is required.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query. Either time or job-id is required.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filtering condition. The maximum value of Filters.Values is 20. Detailed filtering conditions: 
	// <li>job-id 
	// Filter based on task ID.  job-id format: 1379afjk91u32h. Multiple values are not supported, Fuzzy search is not supported. </li>
	// <li>target
	// Filter based on target resource information. target format: http://www.qq.com/1.txt. Multiple values are not supported, Fuzzy search is not supported. </li>
	// <li>domains 
	// Filter based on domain name. domains format: www.qq.com. Fuzzy search is not supported.</li>
	// <li>statuses
	// Filter based on task status.  Fuzzy search is not supported. Options:<br> processing: Processing<br> success: Success<br> failed: Failure<br> timeout: Timeout</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ZoneId")
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of tasks returned.
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// ZoneId. The parameter is required.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Start time of the query. Either time or job-id is required.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query. Either time or job-id is required.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filtering condition. The maximum value of Filters.Values is 20. Detailed filtering conditions: 
	// <li>job-id
	// Filter based on task ID. job-id format: 1379afjk91u32h. Multiple values are not supported, Fuzzy search is not supported;</li>
	// <li>target
	// Filter based on: target resource information. target format: http://www.qq.com/1.txt or tag1. Multiple values are not supported, Fuzzy search is not supported; </li>
	// <li>domains
	// Filter based on domain name. domains format: www.qq.com, Fuzzy search is not supported; </li>
	// <li>statuses
	// Filter based on task status, Fuzzy search is not supported. Options: <br>processing: Processing<br>success: Success<br>failed: Failure<br>timeout: Timeout</li>
	// <li>type
	// Filter based on cleared cache type. Multiple values are not supported, Fuzzy search is not supported. Options:<br> purge_url: URL<br> purge_prefix: Prefix<br> purge_all: All cache content<br> purge_host: Hostname<br> purge_cache_tag: CacheTag</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// ZoneId. The parameter is required.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Start time of the query. Either time or job-id is required.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query. Either time or job-id is required.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filtering condition. The maximum value of Filters.Values is 20. Detailed filtering conditions: 
	// <li>job-id
	// Filter based on task ID. job-id format: 1379afjk91u32h. Multiple values are not supported, Fuzzy search is not supported;</li>
	// <li>target
	// Filter based on: target resource information. target format: http://www.qq.com/1.txt or tag1. Multiple values are not supported, Fuzzy search is not supported; </li>
	// <li>domains
	// Filter based on domain name. domains format: www.qq.com, Fuzzy search is not supported; </li>
	// <li>statuses
	// Filter based on task status, Fuzzy search is not supported. Options: <br>processing: Processing<br>success: Success<br>failed: Failure<br>timeout: Timeout</li>
	// <li>type
	// Filter based on cleared cache type. Multiple values are not supported, Fuzzy search is not supported. Options:<br> purge_url: URL<br> purge_prefix: Prefix<br> purge_all: All cache content<br> purge_host: Hostname<br> purge_cache_tag: CacheTag</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of tasks returned.
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeRealtimeLogDeliveryTasksRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The offset of paginated query. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The limit of paginated query. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value for Filters.Values is 20. If this field is not filled in, all the real-time log delivery task information under the current zone-id will be returned. Detailed filter conditions are as follows:
	// <li>task-id: Filter by real-time log delivery task ID. Fuzzy search is not supported.</li>
	// <li>task-name: Filter by real-time log delivery task name. Fuzzy search is supported, but only one real-time log delivery task name can be filled in for fuzzy search.</li>
	// <li>entity-list: Filter by entity corresponding to the real-time log delivery task. Fuzzy search is not supported. Example values: domain.example.com or sid-2s69eb5wcms7.</li>
	// <li>task-type: Filter by real-time log delivery task type. Fuzzy search is not supported. Optional values:<br>cls: Push to Tencent Cloud CLS;<br>custom_endpoint: Push to a user-defined HTTP(S) address;<br>s3: Push to an AWS S3-compatible bucket address.</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRealtimeLogDeliveryTasksRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The offset of paginated query. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The limit of paginated query. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. The maximum value for Filters.Values is 20. If this field is not filled in, all the real-time log delivery task information under the current zone-id will be returned. Detailed filter conditions are as follows:
	// <li>task-id: Filter by real-time log delivery task ID. Fuzzy search is not supported.</li>
	// <li>task-name: Filter by real-time log delivery task name. Fuzzy search is supported, but only one real-time log delivery task name can be filled in for fuzzy search.</li>
	// <li>entity-list: Filter by entity corresponding to the real-time log delivery task. Fuzzy search is not supported. Example values: domain.example.com or sid-2s69eb5wcms7.</li>
	// <li>task-type: Filter by real-time log delivery task type. Fuzzy search is not supported. Optional values:<br>cls: Push to Tencent Cloud CLS;<br>custom_endpoint: Push to a user-defined HTTP(S) address;<br>s3: Push to an AWS S3-compatible bucket address.</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRealtimeLogDeliveryTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeLogDeliveryTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealtimeLogDeliveryTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealtimeLogDeliveryTasksResponseParams struct {
	// The number of real-time log delivery tasks which match the query conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of all real-time log delivery tasks which match the query conditions.
	RealtimeLogDeliveryTasks []*RealtimeLogDeliveryTask `json:"RealtimeLogDeliveryTasks,omitnil,omitempty" name:"RealtimeLogDeliveryTasks"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRealtimeLogDeliveryTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRealtimeLogDeliveryTasksResponseParams `json:"Response"`
}

func (r *DescribeRealtimeLogDeliveryTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeLogDeliveryTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`rule-id`:<br>   Filter by the <strong>rule ID</strong><br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Filter conditions. Each filter condition can have up to 20 entries. See below for details:
	// <li>`rule-id`:<br>   Filter by the <strong>rule ID</strong><br>   Type: String<br>   Required: No</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of rules. Rules are sorted in order of execution.
	RuleItems []*RuleItem `json:"RuleItems,omitnil,omitempty" name:"RuleItems"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Actions []*RulesSettingAction `json:"Actions,omitnil,omitempty" name:"Actions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSecurityIPGroupInfoRequestParams struct {
	// Site ID, used to specify the query scope.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Maximum number of entries returned in a single response. Default value: 20. Maximum query entries: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The starting entry offset for pagination queries. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSecurityIPGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// Site ID, used to specify the query scope.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Maximum number of entries returned in a single response. Default value: 20. Maximum query entries: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The starting entry offset for pagination queries. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSecurityIPGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityIPGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupInfoResponseParams struct {
	// The number of IP groups that meet the conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Detailed configuration information of the IP group, including the ID, name, and IP/network segment list of each IP group.
	IPGroups []*IPGroup `json:"IPGroups,omitnil,omitempty" name:"IPGroups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityIPGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityIPGroupInfoResponseParams `json:"Response"`
}

func (r *DescribeSecurityIPGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupRequestParams struct {
	// Site ID, used to specify the scope of the queried site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Specifies the ID of a security IP group.
	// <li>When this parameter is provided, only the configuration of the security IP group with the specified ID is queried.</li>
	// <li>When this parameter is not provided, information of all security IP groups under the site is returned.</li>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DescribeSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID, used to specify the scope of the queried site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Specifies the ID of a security IP group.
	// <li>When this parameter is provided, only the configuration of the security IP group with the specified ID is queried.</li>
	// <li>When this parameter is not provided, information of all security IP groups under the site is returned.</li>
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *DescribeSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityIPGroupResponseParams struct {
	// Detailed configuration information of security IP groups, including the ID, name, and IP/IP range list information of each security IP group.
	IPGroups []*IPGroup `json:"IPGroups,omitnil,omitempty" name:"IPGroups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityIPGroupResponseParams `json:"Response"`
}

func (r *DescribeSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTemplateBindingsRequestParams struct {
	// ID of the site to query
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// ID of the policy template to query.
	TemplateId []*string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeSecurityTemplateBindingsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site to query
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// ID of the policy template to query.
	TemplateId []*string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeSecurityTemplateBindingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTemplateBindingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityTemplateBindingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTemplateBindingsResponseParams struct {
	// Bindings of the specified policy template.
	// 
	// When a domain name of a site is bound with the specified policy template, `TemplateScope` includes the `ZoneId` of the related site and the bindings of the domain name. 
	// 
	// Note: If the template is not bound with any domain name, and there is not any existing binding, `TemplateScope=0` is returned.
	// 
	// In the binding list, the same domain name may appear repeatedly in the `EntityStatus` list with different `Status`. For example, when a domain name is being bound to another policy template, it's marked both `online` and `pending`.
	SecurityTemplate []*SecurityTemplateBinding `json:"SecurityTemplate,omitnil,omitempty" name:"SecurityTemplate"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityTemplateBindingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityTemplateBindingsResponseParams `json:"Response"`
}

func (r *DescribeSecurityTemplateBindingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTemplateBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Query indicator. Values: 
	// <li>l4Flow_connections: Number of access connections;</li>
	// <li>l4Flow_flux: Total access traffic;</li>
	// <li>l4Flow_inFlux: Ingress access traffic;</li>
	// <li>l4Flow_outFlux: Egress access traffic. </li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filters
	// <li>ruleId<br>   Filter by the specified <strong>forwarding rule ID</strong></li>
	// <li>proxyId<br>   Filter by the specified <strong>L4 agent ID</strong></li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Query indicator. Values: 
	// <li>l4Flow_connections: Number of access connections;</li>
	// <li>l4Flow_flux: Total access traffic;</li>
	// <li>l4Flow_inFlux: Ingress access traffic;</li>
	// <li>l4Flow_outFlux: Egress access traffic. </li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filters
	// <li>ruleId<br>   Filter by the specified <strong>forwarding rule ID</strong></li>
	// <li>proxyId<br>   Filter by the specified <strong>L4 agent ID</strong></li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of L4 traffic data recorded over time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Indicator list. Values: 
	// <li>l7Flow_outFlux: Edgeone response traffic;</li>
	// <li>l7Flow_inFlux: Edgeone request traffic;</li>
	// <li>l7Flow_outBandwidth: Edgeone response bandwidth;</li>
	// <li>l7Flow_inBandwidth: Edgeone request bandwidth;</li>
	// <li>l7Flow_request: Number of access requests;</li>
	// <li>l7Flow_flux: Uplink + downlink traffic of access requests;< li>
	// <li>l7Flow_bandwidth: Uplink + downlink bandwidth of access requests. </li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filters
	// <li>country<br>Filter by the <strong> Country/Region</strong>. The country/region follows <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> specification. </li>
	// <li>`province`<br>Filter by the <strong>specified province name</strong>. It’s only available when `Area` is `mainland`.</li>
	// <li>`isp`<br>:   Filter by the specified ISP. It’s only available when `Area` is `mainland`.<br>Values: <br>`2`: CTCC; <br>`26`: CUCC; <br>`1046`: CMCC; <br>`3947`: CTT; <br>`38`: CERNET; <br>`43`: GWBN; <br>`0`: Others.</li>
	// <li>`domain`<br>: Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>Filter by the specified <strong>URL path<strong> (such as `/content` or `content/test.jpg`).<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`referer`:<br>Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>If this parameter is specified, the max query period is the last 30 days.<br>The<a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`resourceType`:<br>Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>Note that if this parameter is specified, the max data query period is the last 30 days. <br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">data query scope stated in the specifications of service package</a> related with the `ZoneIds` becomes invalid.</li>
	// <li>`protocol`:<br> Filter  by the specified <strong>HTTP protocol</strong> version <br>Values: <br>`HTTP/1.0`: HTTP 1.0;<br>`HTTP/1.1`: HTTP 1.1;<br>`HTTP/2.0`: HTTP 2.0;<br>`HTTP/3.0`: HTTP 3.0;<br>`WebSocket`: WebSocket.</li>
	// <li>`socket`:<br>Filter by the specified <strong>HTTP protocol</strong> type <br>Values: <br>`HTTP`: HTTP protocol;<br>`HTTPS`: HTTPS protocol;<br>`QUIC`: QUIC protocol.</li>
	// <li>statusCode<br> Filter by [strong> Status Code/strong>]. lt;br> If you only fill in statusCode parameter, you can query data of nearly 30 days at most; <br> If statusCode+Zonelds parameter is filled in at the same time, the supported query data range is the smaller of a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90"> Maximum query range of data analysis/a> and 30 days supported by package. lt;br> The corresponding Value options are as follows: <br> 1XX: Status code of type 1xx; <br> 100:100 status code; <br> 101:101 status code; <br> 102:102 status code; <br> 2XX: Status code of type 2xx; <br> 200:200 status code; <br> 201:201 status code; <br> 202:202 status code; <br> 203:203 status code; <br> 204:204 status code; <br> 205:205 status code; <br> 206:206 status code; <br> 207:207 status code; <br> 3XX: Status code of type 3xx; <br> 300:300 status code; <br> 301:301 status code; <br> 302:302 status code; <br> 303:303 status code; <br> 304:304 status code; <br> 305:305 status code; <br> 307:307 status code; <br> 4XX: Status code of type 4xx; <br> 400:400 status code; <br> 401:401 status code; <br> 402:402 status code; <br> 403:403 status code; <br> 404:404 status code; <br> 405:405 status code; <br> 406:406 status code; <br> 407:407 status code; <br> 408:408 status code; <br> 409:409 status code; <br> 410:410 status code; <br> 411:411 status code; <br> 412:412 status code; <br> 412:413 Status Code; <br> 414:414 status code; <br> 415:415 status code; <br> 416:416 status code; <br> 417:417 status code; <br> 422:422 status code; <br> 423:423 status code; <br> 424:424 status code; <br> 426:426 status code; <br> 451:451 status code; <br> 5XX: Status code of type 5xx; <br> 500:500 status code; <br> 501:501 status code; <br> 502:502 status code; <br> 503:503 status code; <br> 504:504 status code; <br> 505:505 status code; <br> 506:506 status code; <br> 507:507 status code; <br> 510:510 status code; <br> 514:514 status code; <br> 544:544 Status Code.& lt</li>
	// <li>`browserType`:<br>Filter by the specified <strong>browser type</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li><br>Values: <br>`Firefox`: Firefox browser; <br>`Chrome`: Chrome browser; <br>`Safari`: Safari browser; <br>`MicrosoftEdge`: Microsoft Edge browser; <br>`IE`: IE browser; <br>`Opera`: Opera browser; <br>`QQBrowser`: QQ browser; <br>`LBBrowser`: LB browser; <br>`MaxthonBrowser`: Maxthon browser; <br>`SouGouBrowser`: Sogou browser; <br>`BIDUBrowser`: Baidu browser; <br>`TaoBrowser`: Tao browser; <br>`UBrowser`: UC browser; <br>`Other`: Other browsers; <br>`Empty`: The browser type is not specified; <br>`Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>Filter by the <strong>device type</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values:<br>`TV`: TV; <br>`Tablet`: Tablet;<br>`Mobile`: Mobile phone; <br>`Desktop`: Desktop device;<br>`Other`: Other device;<br>`Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>Filter by the <strong>operating system</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`Linux`: Linux OS; <br>`MacOS`: Mac OS;<br>`Android`: Android OS;<br>`IOS`: iOS OS;<br>`Windows`: Windows OS;<br>`NetBSD`: NetBSD OS;<br>`ChromiumOS`: Chromium OS; <br>`Bot`: Web crawler:<br>`Other`: Other OS;   <br>`Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>Filter by the <strong>TLS version</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`TLS1.0`: TLS 1.0;<br>`TLS1.1`: TLS 1.1; <br>`TLS1.2`: TLS 1.2;<br>`TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>Filter by the <strong>specified IP version. <br>Values: <br>`4`: IPv4; <br>`6`: IPv6.
	// <li>`tagKey`<br>Filter by the <strong>Tag Key</strong>. </li>
	// <li>`tagValue`<br>Filter by the <strong>Tag Value</strong>. </li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Indicator list. Values: 
	// <li>l7Flow_outFlux: Edgeone response traffic;</li>
	// <li>l7Flow_inFlux: Edgeone request traffic;</li>
	// <li>l7Flow_outBandwidth: Edgeone response bandwidth;</li>
	// <li>l7Flow_inBandwidth: Edgeone request bandwidth;</li>
	// <li>l7Flow_request: Number of access requests;</li>
	// <li>l7Flow_flux: Uplink + downlink traffic of access requests;< li>
	// <li>l7Flow_bandwidth: Uplink + downlink bandwidth of access requests. </li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The query granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the query period. <br>Period ≤ 1 hour: `min`; <br>1 hour < Period ≤ 2 days: `5min`; <br>2 days < period ≤ 7 days: `hour`; <br>Period > 7 days: `day`.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Filters
	// <li>country<br>Filter by the <strong> Country/Region</strong>. The country/region follows <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> specification. </li>
	// <li>`province`<br>Filter by the <strong>specified province name</strong>. It’s only available when `Area` is `mainland`.</li>
	// <li>`isp`<br>:   Filter by the specified ISP. It’s only available when `Area` is `mainland`.<br>Values: <br>`2`: CTCC; <br>`26`: CUCC; <br>`1046`: CMCC; <br>`3947`: CTT; <br>`38`: CERNET; <br>`43`: GWBN; <br>`0`: Others.</li>
	// <li>`domain`<br>: Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>Filter by the specified <strong>URL path<strong> (such as `/content` or `content/test.jpg`).<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`referer`:<br>Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>If this parameter is specified, the max query period is the last 30 days.<br>The<a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`resourceType`:<br>Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>Note that if this parameter is specified, the max data query period is the last 30 days. <br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">data query scope stated in the specifications of service package</a> related with the `ZoneIds` becomes invalid.</li>
	// <li>`protocol`:<br> Filter  by the specified <strong>HTTP protocol</strong> version <br>Values: <br>`HTTP/1.0`: HTTP 1.0;<br>`HTTP/1.1`: HTTP 1.1;<br>`HTTP/2.0`: HTTP 2.0;<br>`HTTP/3.0`: HTTP 3.0;<br>`WebSocket`: WebSocket.</li>
	// <li>`socket`:<br>Filter by the specified <strong>HTTP protocol</strong> type <br>Values: <br>`HTTP`: HTTP protocol;<br>`HTTPS`: HTTPS protocol;<br>`QUIC`: QUIC protocol.</li>
	// <li>statusCode<br> Filter by [strong> Status Code/strong>]. lt;br> If you only fill in statusCode parameter, you can query data of nearly 30 days at most; <br> If statusCode+Zonelds parameter is filled in at the same time, the supported query data range is the smaller of a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90"> Maximum query range of data analysis/a> and 30 days supported by package. lt;br> The corresponding Value options are as follows: <br> 1XX: Status code of type 1xx; <br> 100:100 status code; <br> 101:101 status code; <br> 102:102 status code; <br> 2XX: Status code of type 2xx; <br> 200:200 status code; <br> 201:201 status code; <br> 202:202 status code; <br> 203:203 status code; <br> 204:204 status code; <br> 205:205 status code; <br> 206:206 status code; <br> 207:207 status code; <br> 3XX: Status code of type 3xx; <br> 300:300 status code; <br> 301:301 status code; <br> 302:302 status code; <br> 303:303 status code; <br> 304:304 status code; <br> 305:305 status code; <br> 307:307 status code; <br> 4XX: Status code of type 4xx; <br> 400:400 status code; <br> 401:401 status code; <br> 402:402 status code; <br> 403:403 status code; <br> 404:404 status code; <br> 405:405 status code; <br> 406:406 status code; <br> 407:407 status code; <br> 408:408 status code; <br> 409:409 status code; <br> 410:410 status code; <br> 411:411 status code; <br> 412:412 status code; <br> 412:413 Status Code; <br> 414:414 status code; <br> 415:415 status code; <br> 416:416 status code; <br> 417:417 status code; <br> 422:422 status code; <br> 423:423 status code; <br> 424:424 status code; <br> 426:426 status code; <br> 451:451 status code; <br> 5XX: Status code of type 5xx; <br> 500:500 status code; <br> 501:501 status code; <br> 502:502 status code; <br> 503:503 status code; <br> 504:504 status code; <br> 505:505 status code; <br> 506:506 status code; <br> 507:507 status code; <br> 510:510 status code; <br> 514:514 status code; <br> 544:544 Status Code.& lt</li>
	// <li>`browserType`:<br>Filter by the specified <strong>browser type</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li><br>Values: <br>`Firefox`: Firefox browser; <br>`Chrome`: Chrome browser; <br>`Safari`: Safari browser; <br>`MicrosoftEdge`: Microsoft Edge browser; <br>`IE`: IE browser; <br>`Opera`: Opera browser; <br>`QQBrowser`: QQ browser; <br>`LBBrowser`: LB browser; <br>`MaxthonBrowser`: Maxthon browser; <br>`SouGouBrowser`: Sogou browser; <br>`BIDUBrowser`: Baidu browser; <br>`TaoBrowser`: Tao browser; <br>`UBrowser`: UC browser; <br>`Other`: Other browsers; <br>`Empty`: The browser type is not specified; <br>`Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>Filter by the <strong>device type</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values:<br>`TV`: TV; <br>`Tablet`: Tablet;<br>`Mobile`: Mobile phone; <br>`Desktop`: Desktop device;<br>`Other`: Other device;<br>`Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>Filter by the <strong>operating system</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`Linux`: Linux OS; <br>`MacOS`: Mac OS;<br>`Android`: Android OS;<br>`IOS`: iOS OS;<br>`Windows`: Windows OS;<br>`NetBSD`: NetBSD OS;<br>`ChromiumOS`: Chromium OS; <br>`Bot`: Web crawler:<br>`Other`: Other OS;   <br>`Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>Filter by the <strong>TLS version</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`TLS1.0`: TLS 1.0;<br>`TLS1.1`: TLS 1.1; <br>`TLS1.2`: TLS 1.2;<br>`TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>Filter by the <strong>specified IP version. <br>Values: <br>`4`: IPv4; <br>`6`: IPv6.
	// <li>`tagKey`<br>Filter by the <strong>Tag Key</strong>. </li>
	// <li>`tagValue`<br>Filter by the <strong>Tag Value</strong>. </li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of L7 traffic data recorded over time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux`: Response traffic.</li>
	// <li>`l7Cache_request`: Response requests.</li>
	// <li>`l7Cache_outBandwidth`: Response bandwidth.</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTimingL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux`: Response traffic.</li>
	// <li>`l7Cache_request`: Response requests.</li>
	// <li>`l7Cache_outBandwidth`: Response bandwidth.</li>
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of cached L7 time-series data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TimingDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

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
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Queries the top N data entries. Maximum value: 1000. Top 10 data entries will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	// <li>`country`<br>Filter by the <strong> Country/Region</strong>. The country/region follows <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> specification. </li>
	// <li>`province`<br>Filter by the <strong>specified province name</strong>. It’s only available when `Area` is `mainland`.</li>
	// <li>`isp`<br>:   Filter by the specified ISP. It’s only available when `Area` is `mainland`.<br>Values: <br>`2`: CTCC; <br>`26`: CUCC; <br>`1046`: CMCC; <br>`3947`: CTT; <br>`38`: CERNET; <br>`43`: GWBN; <br>`0`: Others.</li>
	// <li>`domain`<br>: Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>Filter by the <strong>specified URL Path (such as `/content` or `content/test.jpg`. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`referer`:<br>Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`resourceType`:<br>Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`protocol`:<br> Filter by the specified <strong>HTTP protocol</strong> version <br>Values: <br>`HTTP/1.0`: HTTP 1.0;<br>`HTTP/1.1`: HTTP 1.1;<br>`HTTP/2.0`: HTTP 2.0;<br>`HTTP/3.0`: HTTP 3.0;<br>`WebSocket`: WebSocket.</li>
	// <li>`socket`:<br>Filter by the specified <strong>HTTP protocol type</strong> <br>Values:<br>`HTTP`: HTTP protocol; <br>`HTTPS`: HTTPS protocol;<br>`QUIC`: QUIC protocol.
	// <li>`statusCode`:<br> Filter by the <strong> Status Code</strong><br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values:<br>1XX: Status code of type 1xx <br>100: 100 status code <br>101: 101 status code <br>102: 102 status code <br>2XX: Status code of type 2xx <br>200: 200 status code <br>201: 201 status code <br>202: 202 status code <br>203: 203 status code <br>204: 204 status code <br>205: 205 status code <br>206: 206 status code <br>207: 207 status code <br>3XX: Status code of type 3xx <br>300: 300 status code <br>301: 301 status code <br>302: 302 status code <br>303: 303 status code <br>304: 304 status code <br>305: 305 status code <br>307: 307 status code <br>4XX: Status code of type 4xx <br>400: 400 status code <br>401: 401 status code <br>402: 402 status code <br>403: 403 status code <br>404: 404 status code <br>405: 405 status code <br>406: 406 status code <br>407: 407 status code <br>408: 408 status code <br>409: 409 status code <br>410: 410 status code <br>411: 411 status code <br>412: 412 status code <br>412: 413 Status Code <br>414: 414 status code <br>415: 415 status code <br>416: 416 status code <br>417: 417 status code <br>422: 422 status code <br>423: 423 status code <br>424: 424 status code <br>426: 426 status code <br>451: 451 status code <br>5XX: Status code of type 5xx <br>500: 500 status code <br>501: 501 status code <br>502:502 status code <br>503: 503 status code <br>504: 504 status code <br>505: 505 status code <br>506: 506 status code <br>507: 507 status code <br>510: 510 status code <br>514: 514 status code <br>544: 544 Status Code. </li>
	// <li>`browserType`:<br>Filter by the specified <strong>browser type</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li><br>Values: <br>`Firefox`: Firefox browser; <br>`Chrome`: Chrome browser; <br>`Safari`: Safari browser; <br>`MicrosoftEdge`: Microsoft Edge browser; <br>`IE`: IE browser; <br>`Opera`: Opera browser; <br>`QQBrowser`: QQ browser; <br>`LBBrowser`: LB browser; <br>`MaxthonBrowser`: Maxthon browser; <br>`SouGouBrowser`: Sogou browser; <br>`BIDUBrowser`: Baidu browser; <br>`TaoBrowser`: Tao browser; <br>`UBrowser`: UC browser; <br>`Other`: Other browsers; <br>`Empty`: The browser type is not specified; <br>`Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>Filter by the <strong>device type</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values:<br>`TV`: TV; <br>`Tablet`: Tablet;<br>`Mobile`: Mobile phone; <br>`Desktop`: Desktop device;<br>`Other`: Other device;<br>`Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>Filter by the <strong>operating system</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`Linux`: Linux OS; <br>`MacOS`: Mac OS;<br>`Android`: Android OS;<br>`IOS`: iOS OS;<br>`Windows`: Windows OS;<br>`NetBSD`: NetBSD OS;<br>`ChromiumOS`: Chromium OS; <br>`Bot`: Web crawler:<br>`Other`: Other OS;   <br>`Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>Filter by the <strong>TLS version</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`TLS1.0`: TLS 1.0;<br>`TLS1.1`: TLS 1.1; <br>`TLS1.2`: TLS 1.2;<br>`TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>Filter by the <strong>specified IP version. <br>Values: <br>`4`: IPv4; <br>`6`: IPv6.
	// <li>`tagKey`<br>Filter by the <strong>Tag Key</strong>. </li>
	// <li>`tagValue`<br>Filter by the <strong>Tag Value</strong>. </li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

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
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Queries the top N data entries. Maximum value: 1000. Top 10 data entries will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	// <li>`country`<br>Filter by the <strong> Country/Region</strong>. The country/region follows <a href="https://baike.baidu.com/item/ISO%203166-1/5269555">ISO 3166</a> specification. </li>
	// <li>`province`<br>Filter by the <strong>specified province name</strong>. It’s only available when `Area` is `mainland`.</li>
	// <li>`isp`<br>:   Filter by the specified ISP. It’s only available when `Area` is `mainland`.<br>Values: <br>`2`: CTCC; <br>`26`: CUCC; <br>`1046`: CMCC; <br>`3947`: CTT; <br>`38`: CERNET; <br>`43`: GWBN; <br>`0`: Others.</li>
	// <li>`domain`<br>: Filter by the specified <strong>sub-domain name</strong>, such as `test.example.com`</li>
	// <li>`url`:<br>Filter by the <strong>specified URL Path (such as `/content` or `content/test.jpg`. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`referer`:<br>Filter by the specified <strong>Referer header</strong>, such as `example.com`.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`resourceType`:<br>Filter by the specified <strong>resource file type</strong>, such as `jpg`, `css`. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li>
	// <li>`protocol`:<br> Filter by the specified <strong>HTTP protocol</strong> version <br>Values: <br>`HTTP/1.0`: HTTP 1.0;<br>`HTTP/1.1`: HTTP 1.1;<br>`HTTP/2.0`: HTTP 2.0;<br>`HTTP/3.0`: HTTP 3.0;<br>`WebSocket`: WebSocket.</li>
	// <li>`socket`:<br>Filter by the specified <strong>HTTP protocol type</strong> <br>Values:<br>`HTTP`: HTTP protocol; <br>`HTTPS`: HTTPS protocol;<br>`QUIC`: QUIC protocol.
	// <li>`statusCode`:<br> Filter by the <strong> Status Code</strong><br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values:<br>1XX: Status code of type 1xx <br>100: 100 status code <br>101: 101 status code <br>102: 102 status code <br>2XX: Status code of type 2xx <br>200: 200 status code <br>201: 201 status code <br>202: 202 status code <br>203: 203 status code <br>204: 204 status code <br>205: 205 status code <br>206: 206 status code <br>207: 207 status code <br>3XX: Status code of type 3xx <br>300: 300 status code <br>301: 301 status code <br>302: 302 status code <br>303: 303 status code <br>304: 304 status code <br>305: 305 status code <br>307: 307 status code <br>4XX: Status code of type 4xx <br>400: 400 status code <br>401: 401 status code <br>402: 402 status code <br>403: 403 status code <br>404: 404 status code <br>405: 405 status code <br>406: 406 status code <br>407: 407 status code <br>408: 408 status code <br>409: 409 status code <br>410: 410 status code <br>411: 411 status code <br>412: 412 status code <br>412: 413 Status Code <br>414: 414 status code <br>415: 415 status code <br>416: 416 status code <br>417: 417 status code <br>422: 422 status code <br>423: 423 status code <br>424: 424 status code <br>426: 426 status code <br>451: 451 status code <br>5XX: Status code of type 5xx <br>500: 500 status code <br>501: 501 status code <br>502:502 status code <br>503: 503 status code <br>504: 504 status code <br>505: 505 status code <br>506: 506 status code <br>507: 507 status code <br>510: 510 status code <br>514: 514 status code <br>544: 544 Status Code. </li>
	// <li>`browserType`:<br>Filter by the specified <strong>browser type</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.</li><br>Values: <br>`Firefox`: Firefox browser; <br>`Chrome`: Chrome browser; <br>`Safari`: Safari browser; <br>`MicrosoftEdge`: Microsoft Edge browser; <br>`IE`: IE browser; <br>`Opera`: Opera browser; <br>`QQBrowser`: QQ browser; <br>`LBBrowser`: LB browser; <br>`MaxthonBrowser`: Maxthon browser; <br>`SouGouBrowser`: Sogou browser; <br>`BIDUBrowser`: Baidu browser; <br>`TaoBrowser`: Tao browser; <br>`UBrowser`: UC browser; <br>`Other`: Other browsers; <br>`Empty`: The browser type is not specified; <br>`Bot`: Web crawler.</li>
	// <li>`deviceType`:<br>Filter by the <strong>device type</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values:<br>`TV`: TV; <br>`Tablet`: Tablet;<br>`Mobile`: Mobile phone; <br>`Desktop`: Desktop device;<br>`Other`: Other device;<br>`Empty`: Device type not specified.</li>
	// <li>`operatingSystemType`:<br>Filter by the <strong>operating system</strong>.<br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`Linux`: Linux OS; <br>`MacOS`: Mac OS;<br>`Android`: Android OS;<br>`IOS`: iOS OS;<br>`Windows`: Windows OS;<br>`NetBSD`: NetBSD OS;<br>`ChromiumOS`: Chromium OS; <br>`Bot`: Web crawler:<br>`Other`: Other OS;   <br>`Empty`: The OS is not specified.</li>
	// <li>`tlsVersion`:<br>Filter by the <strong>TLS version</strong>. <br>If this parameter is specified, the max query period is the last 30 days.<br>The <a href="https://intl.cloud.tencent.com/document/product/1552/77380?from_cn_redirect=1#edgeone-.E5.A5.97.E9.A4.90">max data query scope stated in the service package specifications</a> of the site (if `ZoneIds` specified) becomes invalid.<br>Values: <br>`TLS1.0`: TLS 1.0;<br>`TLS1.1`: TLS 1.1; <br>`TLS1.2`: TLS 1.2;<br>`TLS1.3`: TLS 1.3.</li>
	// <li>`ipVersion`<br>Filter by the <strong>specified IP version. <br>Values: <br>`4`: IPv4; <br>`6`: IPv6.
	// <li>`tagKey`<br>Filter by the <strong>Tag Key</strong>. </li>
	// <li>`tagValue`<br>Filter by the <strong>Tag Value</strong>. </li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of top-ranked L7 traffic data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux_domain`: Host/Domain name;</li>
	// <li>`l7Cache_outFlux_url`: URL address;</li>
	// <li>`l7Cache_outFlux_resourceType`: Resource type;</li>
	// <li>`l7Cache_outFlux_statusCode`: Status code.</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of data entries to be queried. The maximum value is 1000. If it is not specified, the value 10 is used by default, indicating that the top 10 data entries.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DescribeTopL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Cache_outFlux_domain`: Host/Domain name;</li>
	// <li>`l7Cache_outFlux_url`: URL address;</li>
	// <li>`l7Cache_outFlux_resourceType`: Resource type;</li>
	// <li>`l7Cache_outFlux_statusCode`: Status code.</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of data entries to be queried. The maximum value is 1000. If it is not specified, the value 10 is used by default, indicating that the top 10 data entries.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. See below for details: 
	// <li>`domain`<br>   Filter by the <strong>sub-domain name</strong>, such as `test.example.com`<br>   Type: String<br>   Required: No</li>
	// <li>`url`<br>   Filter by the <strong>URL</strong>, such as `/content`. The query period cannot exceed 30 days. <br>   Type: String<br>   Required: No</li>
	// <li>`resourceType`<br>   Filter by the <strong>resource file type</strong>, such as `jpg`, `png`. The query period cannot exceed 30 days.<br>  Type: String<br>   Required: No</li>
	// <li>cacheType<br>  Filter by the <strong>cache hit result</strong>.<br> Type: String<br>   Required: No<br>   Values: <br>   `hit`: Cache hit; <br>   `dynamic`: Resource non-cacheable; <br>   `miss`: Cache miss</li>
	// <li>`statusCode`<br>   Filter by the <strong> status code</strong>. The query period  cannot exceed 30 days. <br>   Type: String<br>   Required: No<br>   Values: <br>   `1XX`: All 1xx status codes;<br>   `100`: 100 status code;<br>   `101`: 101 status code;<br>   `102`: 102 status code;<br>   `2XX`: All 2xx status codes;<br>   `200`: 200 status code;<br>   `201`: 201 status code;<br>   `202`: 202 status code;<br>   `203`: 203 status code;<br>   `204`: 204 status code;<br>   `205`: 205 status code;<br>   `206`: 206 status code;<br>   `207`: 207 status code;<br>   `3XX`: All 3xx status codes;<br>   `300`: 300 status code;<br>   `301`: 301 status code;<br>   `302`: 302 status code;<br>   `303`: 303 status code;<br>   `304`: 304 status code;<br>   `305`: 305 status code;<br>   `307`: 307 status code;<br>   `4XX`: All 4xx status codes;<br>   `400`: 400 status code;<br>   `401`: 401 status code;<br>   `402`: 402 status code;<br>   `403`: 403 status code;<br>   `404`: 404 status code;<br>   `405`: 405 status code;<br>   `406`: 406 status code;<br>   `407`: 407 status code;<br>   `408`: 408 status code;<br>   `409`: 409 status code;<br>   `410`: 410 status code;<br>   `411`: 411 status code;<br>   `412`: 412 status code;<br>   `412`: 413 status code;<br>   `414`: 414 status code;<br>   `415`: 415 status code;<br>   `416`: 416 status code;<br>   `417`: 417 status code;<br>  `422`: 422 status code;<br>   `423`: 423 status code;<br>   `424`: 424 status code;<br>   `426`: 426 status code;<br>   `451`: 451 status code;<br>   `5XX`: All 5xx status codes;<br>   `500`: 500 status code;<br>   `501`: 501 status code;<br>   `502`: 502 status code;<br>   `503`: 503 status code;<br>   `504`: 504 status code;<br>   `505`: 505 status code;<br>   `506`: 506 status code;<br>   `507`: 507 status code;<br>   `510`: 510 status code;<br>   `514`: 514 status code;<br>   `544`: 544 status code.</li>
	// <li>`tagKey`:<br>   Filter by the <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>   Filter by the <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Geolocation scope. Values:
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`mainland`: Chinese mainland</li>
	// <li>`global`: Global</li>If this field is not specified, the default value `global` is used.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
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
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of cached L7 top-ranked traffic data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopDataRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribeZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
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
	ZoneSetting *ZoneSetting `json:"ZoneSetting,omitnil,omitempty" name:"ZoneSetting"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. Up to 20 values for each filter. If this parameter is not filled in, the information of all sites under the current account is returned. Detailed filtering conditions are as follows:
	// <li>`zone-name`: Site name </li><li>`zone-id`: Site ID, such as zone-2noz78a8ev6k</li><li>`status`: Site status </li><li>`tag-key`: Tag key </li><li>`tag-value`: Tag value </li>Only `zone-name` supports fuzzy query.
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. Values include:
	// <li>`type`: Connection mode</li>
	// <li>`area`: Acceleration region</li>
	// <li>`create-time`: Creation time</li>
	// <li>`zone-name`: Site name</li>
	// <li>`use-time`: Last used time</li>
	// <li>`active-status` Effective status</li> Default value: `create-time`
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sort direction. If the field value is a number, sort by the numeric value. If the field value is text, sort by the ascill code. Values include:
	// <li>`asc`: From the smallest to largest</li>
	// <li>`desc`: From the largest to smallest.</li>Default value: `desc`
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter conditions. Up to 20 values for each filter. If this parameter is not filled in, the information of all sites under the current account is returned. Detailed filtering conditions are as follows:
	// <li>`zone-name`: Site name </li><li>`zone-id`: Site ID, such as zone-2noz78a8ev6k</li><li>`status`: Site status </li><li>`tag-key`: Tag key </li><li>`tag-value`: Tag value </li>Only `zone-name` supports fuzzy query.
	Filters []*AdvancedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. Values include:
	// <li>`type`: Connection mode</li>
	// <li>`area`: Acceleration region</li>
	// <li>`create-time`: Creation time</li>
	// <li>`zone-name`: Site name</li>
	// <li>`use-time`: Last used time</li>
	// <li>`active-status` Effective status</li> Default value: `create-time`
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sort direction. If the field value is a number, sort by the numeric value. If the field value is text, sort by the ascill code. Values include:
	// <li>`asc`: From the smallest to largest</li>
	// <li>`desc`: From the largest to smallest.</li>Default value: `desc`
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of sites.
	Zones []*Zone `json:"Zones,omitnil,omitempty" name:"Zones"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DestroyPlanRequestParams struct {
	// Plan ID, in the format of edgeone-2wdo315m2y4c.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type DestroyPlanRequest struct {
	*tchttp.BaseRequest
	
	// Plan ID, in the format of edgeone-2wdo315m2y4c.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *DestroyPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPlanResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyPlanResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPlanResponseParams `json:"Response"`
}

func (r *DestroyPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailHost struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The acceleration status. Values:
	// <li>`process`: In progress</li>
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The domain name.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Name of the site
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// The assigned CNAME
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// The resource ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The lock status.
	Lock *int64 `json:"Lock,omitnil,omitempty" name:"Lock"`

	// The domain name status.
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// The acceleration area of the domain name. Values:
	// <li>`global`: Global.</li>
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Outside the Chinese mainland.</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The acceleration type configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccelerateType *AccelerateType `json:"AccelerateType,omitnil,omitempty" name:"AccelerateType"`

	// The HTTPS configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// The cache configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// The origin configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// The security type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecurityType *SecurityType `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// The cache key configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// The smart compression configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// The WAF protection configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Waf *Waf `json:"Waf,omitnil,omitempty" name:"Waf"`

	// The CC protection configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CC *CC `json:"CC,omitnil,omitempty" name:"CC"`

	// DDoS mitigation configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	DDoS *DDoS `json:"DDoS,omitnil,omitempty" name:"DDoS"`

	// The smart routing configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// The IPv6 access configuration item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`
}

type DiffIPWhitelist struct {
	// The latest intermediate IPs.
	LatestIPWhitelist *IPWhitelist `json:"LatestIPWhitelist,omitnil,omitempty" name:"LatestIPWhitelist"`

	// The intermediate IPs added to the existing list.
	AddedIPWhitelist *IPWhitelist `json:"AddedIPWhitelist,omitnil,omitempty" name:"AddedIPWhitelist"`

	// The intermediate IPs removed from the existing list.
	RemovedIPWhitelist *IPWhitelist `json:"RemovedIPWhitelist,omitnil,omitempty" name:"RemovedIPWhitelist"`

	// The intermediate IPs that remain unchanged.
	NoChangeIPWhitelist *IPWhitelist `json:"NoChangeIPWhitelist,omitnil,omitempty" name:"NoChangeIPWhitelist"`
}

type DnsVerification struct {
	// The host record.
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// The record type.
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// The record value.
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

// Predefined struct for user
type DownloadL4LogsRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// List of L4 proxy instance IDs.
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 300.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DownloadL4LogsRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// List of L4 proxy instance IDs.
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 300.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of L4 logs.
	Data []*L4OfflineLog `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Limit on paginated queries. Default value: 20. Maximum value: 300.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ZoneId set. This parameter is required.
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Limit on paginated queries. Default value: 20. Maximum value: 300.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	// Total number of query results.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of L7 logs.
	Data []*L7OfflineLog `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The settings of the block page that applies managed rules. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WafDropPageDetail *DropPageDetail `json:"WafDropPageDetail,omitnil,omitempty" name:"WafDropPageDetail"`

	// The settings of the block page that applies custom rules. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AclDropPageDetail *DropPageDetail `json:"AclDropPageDetail,omitnil,omitempty" name:"AclDropPageDetail"`
}

type DropPageDetail struct {
	// The ID of the block page. Specify `0` to use the default block page. 
	// (Disused) If 0 is passed, the default block page will be used.
	PageId *int64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// The HTTP status code to trigger the block page. Range: 100-600, excluding 3xx codes. Code 566: Requests blocked by managed rules. Code 567: Requests blocked by web security rules (except managed rules).
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// The block page file or URL.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Type of the block page. Values:
	// <li>`page`: Return the specified page.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// ID of custom response. The ID can be obtained via the `DescribeCustomErrorPages` API. It's required when `Type=page`.
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`
}

type EntityStatus struct {
	// Instance name. Only subdomain names are supported.
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// Instance configuration status. Values:
	// <li>`online`: Configuration has taken effect;</li><li>`fail`: Configuration failed;</li><li>`process`: Configuration is being delivered. </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Message returned after the operation completed. 
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type EnvInfo struct {
	// Environment ID.
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Environment type. Valid values: 
	// <li>production: Production environment.</li><li> staging: Test environment. </li>
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// Environment status. Valid values: 
	// <li>creating: Being created.</li>
	// <li>running: The environment is stable, with version changes allowed.</li>
	// <li>version_deploying: The version is currently being deployed, with no more changes allowed. </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Effective scope of the configuration in the current environment. Valid values: 
	// <li>ALL: It takes effect on the entire network when EnvType is set to production.</li>
	// <li>It returns the IP address of the test node for host binding during testing when EnvType is set to staging. </li>
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// For the effective versions of each configuration group in the current environment, there are two possible scenarios based on the value of Status: 
	// <li>When Status is set to version_deploying, the returned value of this field represents the previously effective version. In other words, during the deployment of the new version, the effective version is the one that was in effect before any changes were made.</li>
	// <li>When Status is set to running, the value returned by this field is the currently effective version. </li>
	CurrentConfigGroupVersionInfos []*ConfigGroupVersionInfo `json:"CurrentConfigGroupVersionInfos,omitnil,omitempty" name:"CurrentConfigGroupVersionInfos"`

	// Creation time. The time format follows the ISO 8601 standard and is represented in Coordinated Universal Time (UTC).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time. The time format follows the ISO 8601 standard and is represented in Coordinated Universal Time (UTC).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type ErrorPageReference struct {
	// Referenced business ID, such as the custom block rule ID.
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type ExceptConfig struct {
	// Whether to enable configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The settings of the exception rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptUserRules []*ExceptUserRule `json:"ExceptUserRules,omitnil,omitempty" name:"ExceptUserRules"`
}

type ExceptUserRule struct {
	// The rule name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// The rule action. It only supports the value `skip`, which indicates skipping all managed rules.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// The rule ID, which is automatically created and only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// The update time. If it is null, the current date and time is recorded.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The matching condition.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptUserRuleConditions []*ExceptUserRuleCondition `json:"ExceptUserRuleConditions,omitnil,omitempty" name:"ExceptUserRuleConditions"`

	// The scope to which the exception rule applies.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptUserRuleScope *ExceptUserRuleScope `json:"ExceptUserRuleScope,omitnil,omitempty" name:"ExceptUserRuleScope"`

	// The rule priority. Value range: 0-100. If it is null, it defaults to 0.
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`
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
	MatchFrom *string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// The parameter of the field. Only when `MatchFrom = header`, the key contained in the header can be passed.
	MatchParam *string `json:"MatchParam,omitnil,omitempty" name:"MatchParam"`

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
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// The value of the parameter.
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type ExceptUserRuleScope struct {
	// Exception mode. Values:
	// <li>`complete`: Skip the exception rule for full requests.</li>
	// <li>`partial`: Skip the exception rule for partial requests.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The module to be activated. Values:
	// <li>`waf`: Tencent Cloud-managed rules</li>
	// <li>`rate`: Rate limiting rules</li>
	// <li>`acl`: Custom rule</li>
	// <li>`cc`: CC attack defense</li>
	// <li>`bot`: Bot protection</li>
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Modules []*string `json:"Modules,omitnil,omitempty" name:"Modules"`

	// Module settings of the exception rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PartialModules []*PartialModule `json:"PartialModules,omitnil,omitempty" name:"PartialModules"`

	// Condition settings of the exception rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SkipConditions []*SkipCondition `json:"SkipConditions,omitnil,omitempty" name:"SkipConditions"`
}

type FailReason struct {
	// Failure reason.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// List of resources failed to be processed. 
	Targets []*string `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type FileAscriptionInfo struct {
	// Directory of the verification file.
	IdentifyPath *string `json:"IdentifyPath,omitnil,omitempty" name:"IdentifyPath"`

	// Content of the verification file.
	IdentifyContent *string `json:"IdentifyContent,omitnil,omitempty" name:"IdentifyContent"`
}

type FileVerification struct {
	// EdgeOne obtains the file verification information in the format of "Scheme + Host + URL Path", (e.g. https://www.example.com/.well-known/teo-verification/z12h416twn.txt). This field is the URL path section of the URL you need to create.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Content of the verification file. The contents of this field need to be filled into the text file returned by `Path`.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type Filter struct {
	// Fields to be filtered.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value of the filtered field.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FirstPartConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The transfer period threshold of the first 8 KB. If the threshold is reached, it’s considered a slow attack. Default: `5`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatTime *uint64 `json:"StatTime,omitnil,omitempty" name:"StatTime"`
}

type FollowOrigin struct {
	// Whether to enable the configuration of following the origin server. Valid values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Sets the default cache time when the origin server does not return the Cache-Control header.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DefaultCacheTime *int64 `json:"DefaultCacheTime,omitnil,omitempty" name:"DefaultCacheTime"`

	// Specifies whether to enable cache when the origin server does not return the Cache-Control header.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DefaultCache *string `json:"DefaultCache,omitnil,omitempty" name:"DefaultCache"`

	// Specifies whether to use the default caching policy when Cache-Control is not returned from the origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DefaultCacheStrategy *string `json:"DefaultCacheStrategy,omitnil,omitempty" name:"DefaultCacheStrategy"`
}

type ForceRedirect struct {
	// Whether to enable force HTTPS redirect. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Redirect status code. Values:
	// <li>`301`: 301 redirect</li>
	// <li>`302`: 302 redirect</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil,omitempty" name:"RedirectStatusCode"`
}

type Grpc struct {
	// Whether to enable gRPC support. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type Header struct {
	// HTTP header name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// HTTP header value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Hsts struct {
	// Whether to enable the configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// MaxAge (in seconds). The maximum value is 1 day. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MaxAge *int64 `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// Whether to contain subdomain names. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil,omitempty" name:"IncludeSubDomains"`

	// Whether to enable preloading. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Preload *string `json:"Preload,omitnil,omitempty" name:"Preload"`
}

type Https struct {
	// Whether to enable HTTP2. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Http2 *string `json:"Http2,omitnil,omitempty" name:"Http2"`

	// Whether to enable OCSP. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	OcspStapling *string `json:"OcspStapling,omitnil,omitempty" name:"OcspStapling"`

	// TLS version. Valid values: 
	// <li>`TLSv1`: TLSv1 version;</li>
	// <li>`TLSV1.1`: TLSv1.1 version;</li>
	// <li>`TLSV1.2`: TLSv1.2 version;</li>
	// <li>`TLSv1.3`: TLSv1.3 version.</li>Only consecutive versions can be enabled at the same time. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TlsVersion []*string `json:"TlsVersion,omitnil,omitempty" name:"TlsVersion"`

	// HSTS Configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Hsts *Hsts `json:"Hsts,omitnil,omitempty" name:"Hsts"`

	// The certificate configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertInfo []*ServerCertInfo `json:"CertInfo,omitnil,omitempty" name:"CertInfo"`

	// Whether the certificate is managed by EdgeOne. Values:
	// <li>`apply`: Managed by EdgeOne.</li>
	// <li>`none`: Not managed by EdgeOne.</li>If it is left empty, the default value `none` is used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// The cipher suite, with values:
	// <li>loose-v2023: Provides high compatibility with general security, and supports TLS 1.0-1.3 cipher suites;</li>
	// <li>general-v2023: Provides relatively high compatibility with moderate security, and supports TLS 1.2-1.3 cipher suites;</li>
	// <li>strict-v2023: Provides high security, disables all cipher suites with security risks, and supports TLS 1.2-1.3 cipher suites.</li>
	// Note: This field may return null, which indicates a failure to obtain a valid value.
	CipherSuite *string `json:"CipherSuite,omitnil,omitempty" name:"CipherSuite"`
}

type IPGroup struct {
	// Group ID. Enter `0`.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Group name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// IP group content. Only supports IP and IP mask.
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`
}

type IPRegionInfo struct {
	// IP address, IPV4 or IPV6.
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Whether the IP belongs to an EdgeOne node. Valid values:
	// <li>yes: This IP belongs to an EdgeOne node;</li>
	// <li>no: This IP does not belong to an EdgeOne node.</li>
	IsEdgeOneIP *string `json:"IsEdgeOneIP,omitnil,omitempty" name:"IsEdgeOneIP"`
}

type IPWhitelist struct {
	// List of IPv4 addresses
	IPv4 []*string `json:"IPv4,omitnil,omitempty" name:"IPv4"`

	// List of IPv6 addresses
	IPv6 []*string `json:"IPv6,omitnil,omitempty" name:"IPv6"`
}

type Identification struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// The subdomain name to be verified. To verify the ownership of a site, leave it blank.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The verification status. Values:
	// <li>`pending`: The verification is ongoing.</li>
	// <li>`finished`: The verification completed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Details of the DNS record.
	Ascription *AscriptionInfo `json:"Ascription,omitnil,omitempty" name:"Ascription"`

	// The NS record of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil,omitempty" name:"OriginalNameServers"`

	// Details of the verification file.
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitnil,omitempty" name:"FileAscription"`
}

// Predefined struct for user
type IdentifyZoneRequestParams struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// A subdomain name under the site. Specify this field if you want to verify the ownership of a subdomain name. Otherwise you can leave it blank.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site name.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// A subdomain name under the site. Specify this field if you want to verify the ownership of a subdomain name. Otherwise you can leave it blank.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	Ascription *AscriptionInfo `json:"Ascription,omitnil,omitempty" name:"Ascription"`

	// Details of the verification file.
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitnil,omitempty" name:"FileAscription"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type IncreasePlanQuotaRequestParams struct {
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// The types of new plan quotas available include:<li> site: Number of sites;</li><li> precise_access_control_rule: the number of rules under "Web Protection - Custom Rules - Precision Matching Policy";</li><li> rate_limiting_rule: the number of rules under "Web Protection - Rate Limiting - Precision Rate Limiting Module". </li>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// Number of new quotas. The maximum number of quotas that can be added at one time is 100.
	QuotaNumber *int64 `json:"QuotaNumber,omitnil,omitempty" name:"QuotaNumber"`
}

type IncreasePlanQuotaRequest struct {
	*tchttp.BaseRequest
	
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// The types of new plan quotas available include:<li> site: Number of sites;</li><li> precise_access_control_rule: the number of rules under "Web Protection - Custom Rules - Precision Matching Policy";</li><li> rate_limiting_rule: the number of rules under "Web Protection - Rate Limiting - Precision Rate Limiting Module". </li>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// Number of new quotas. The maximum number of quotas that can be added at one time is 100.
	QuotaNumber *int64 `json:"QuotaNumber,omitnil,omitempty" name:"QuotaNumber"`
}

func (r *IncreasePlanQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IncreasePlanQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "QuotaType")
	delete(f, "QuotaNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IncreasePlanQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IncreasePlanQuotaResponseParams struct {
	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IncreasePlanQuotaResponse struct {
	*tchttp.BaseResponse
	Response *IncreasePlanQuotaResponseParams `json:"Response"`
}

func (r *IncreasePlanQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IncreasePlanQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntelligenceRule struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Items in a bot intelligence rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntelligenceRuleItems []*IntelligenceRuleItem `json:"IntelligenceRuleItems,omitnil,omitempty" name:"IntelligenceRuleItems"`
}

type IntelligenceRuleItem struct {
	// The tag to categorize bots. Values:
	// <li>`evil_bot`: Malicious bot</li>
	// <li>`suspect_bot`: Suspected bot</li>
	// <li>`good_bot`: Good bot</li>
	// <li>`normal`: Normal request</li>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// The action taken on bots. Values
	// <li>`drop`: Block</li>
	// <li>`trans`: Allow</li>
	// <li>`alg`: JavaScript challenge</li>
	// <li>`captcha`: Managed challenge</li>
	// <li>`monitor`: Observe</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type IpTableConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The settings of the basic access control rule. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IpTableRules []*IpTableRule `json:"IpTableRules,omitnil,omitempty" name:"IpTableRules"`
}

type IpTableRule struct {
	// The action. Values:
	// <li>`drop`: Block</li>
	// <li>`trans`: Allow</li>
	// <li>`monitor`: Observe</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The matching dimension. Values:
	// <li>`ip`: Match by IP.</li>
	// <li>`area`: Match by IP region.</li>
	MatchFrom *string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// Matching method. It defaults to `equal` if it’s left empty.
	// Values: 
	// <li>`is_empty`: The field is empty.</li>
	// <li>`not_exists`: The configuration item does not exist.</li>
	// <li>`include`: Include</li>
	// <li>`not_include`: Do not include</li>
	// <li>`equal`: Equal to</li>
	// <li>`not_equal`: Not equal to</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// The rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// The update time, which is only used as an output parameter.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The rule status. A null value indicates that the rule is enabled. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The rule name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Matching content. It’s not required when `Operator` is `is_emty` or `not_exists`. 
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type Ipv6 struct {
	// Whether to enable IPv6 access. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type L4OfflineLog struct {
	// L4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Log query area. Valid values:
	// <li>`mainland`: Chinese mainland;</li>
	// <li>`overseas`: Global (outside the Chinese mainland). </li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Log packet name.
	LogPacketName *string `json:"LogPacketName,omitnil,omitempty" name:"LogPacketName"`

	// Log download address.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// (Disused) Log packaging time. 
	LogTime *int64 `json:"LogTime,omitnil,omitempty" name:"LogTime"`

	// Start time of log packaging.
	LogStartTime *string `json:"LogStartTime,omitnil,omitempty" name:"LogStartTime"`

	// End time of the log package.
	LogEndTime *string `json:"LogEndTime,omitnil,omitempty" name:"LogEndTime"`

	// Log size (in bytes).
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type L4Proxy struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Layer 4 proxy instance name.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// Acceleration zone of the Layer 4 proxy instance.<li>mainland: Availability zone in the Chinese mainland;</li><li>overseas: Global availability zone (excluding the Chinese mainland);</li><li>global: Global availability zone.</li>	
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Access via CNAME.
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// After the fixed IP address is enabled, this value will return the corresponding access IP address; if it is not enabled, this value will be empty.
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// Status of the Layer 4 proxy instance.<li>online: Enabled;</li>
	// <li>offline: Disabled;</li>
	// <li>progress: Deploying;</li>	
	// <li>stopping: Disabling;</li>
	// <li>banned: Blocked;</li>
	// <li>fail: Failed to deploy or disable.</li>	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies whether to enable IPv6 access.<li>on: Enable;</li> <li>off: Disable.</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Specifies whether to enable the fixed IP address.<li>on: Enable;</li> <li>off: Disable.</li>
	StaticIp *string `json:"StaticIp,omitnil,omitempty" name:"StaticIp"`

	// Specifies whether to enable network optimization in the Chinese mainland.<li>on: Enable</li> <li>off: Disable</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// Security protection configuration.
	// Note: This field may return null, indicating that no valid value can be obtained.
	DDosProtectionConfig *DDosProtectionConfig `json:"DDosProtectionConfig,omitnil,omitempty" name:"DDosProtectionConfig"`

	// Number of forwarding rules under the Layer 4 proxy instance.
	L4ProxyRuleCount *int64 `json:"L4ProxyRuleCount,omitnil,omitempty" name:"L4ProxyRuleCount"`

	// Latest modification time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type L4ProxyRule struct {
	// Forwarding rule ID.
	// Note: Do not fill in this parameter when L4ProxyRule is used as an input parameter in CreateL4ProxyRules; it must be filled in when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Forwarding protocol. Valid values:
	// <li>TCP: TCP protocol;</li>
	// <li>UDP: UDP protocol.</li>
	// Note: This parameter must be filled in when L4ProxyRule is used as an input parameter in CreateL4ProxyRules; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Forwarding port, which can be set as follows:
	// <li>A single port, such as 80;</li>
	// <li>A range of ports, such as 81-85, representing ports 81, 82, 83, 84, 85.</li>
	// Note: This parameter must be filled in when L4ProxyRule is used as an input parameter in CreateL4ProxyRules; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	PortRange []*string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// Origin server type. Valid values:
	// <li>IP_DOMAIN: IP/Domain name origin server;</li>
	// <li>ORIGIN_GROUP: Origin server group;</li>
	// <li>LB: Cloud Load Balancer, currently only open to the allowlist.</li>
	// Note: This parameter must be filled in when L4ProxyRule is used as an input parameter in CreateL4ProxyRules; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Origin server address.
	// <li>When OriginType is set to IP_DOMAIN, enter the IP address or domain name, such as 8.8.8.8 or test.com;</li>
	// <li>When OriginType is set to ORIGIN_GROUP, enter the origin server group ID, such as og-537y24vf5b41;</li>
	// <li>When OriginType is set to LB, enter the Cloud Load Balancer instance ID, such as lb-2qwk30xf7s9g.</li>
	// Note: This parameter must be filled in when L4ProxyRule is used as an input parameter in CreateL4ProxyRules; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// Origin server port, which can be set as follows:<li>A single port, such as 80;</li>
	// <li>A range of ports, such as 81-85, representing ports 81, 82, 83, 84, 85. When inputting a range of ports, ensure that the length corresponds with that of the forwarding port range. For example, if the forwarding port range is 80-90, this port range should be 90-100.</li>
	// Note: This parameter must be filled in when L4ProxyRule is used as an input parameter in CreateL4ProxyRules; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	OriginPortRange *string `json:"OriginPortRange,omitnil,omitempty" name:"OriginPortRange"`

	// Transmission of the client IP address. Valid values:<li>TOA: Available only when Protocol=TCP;</li> 
	// <li>PPV1: Transmission via Proxy Protocol V1. Available only when Protocol=TCP;</li>
	// <li>PPV2: Transmission via Proxy Protocol V2;</li> 
	// <li>SPP: Transmission via Simple Proxy Protocol. Available only when Protocol=UDP;</li> 
	// <li>OFF: No transmission.</li>
	// Note: This parameter is optional when L4ProxyRule is used as an input parameter in CreateL4ProxyRules, and if not specified, the default value OFF will be used; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitnil,omitempty" name:"ClientIPPassThroughMode"`

	// Specifies whether to enable session persistence. Valid values:
	// <li>on: Enable;</li>
	// <li>off: Disable.</li>
	// Note: This parameter is optional when L4ProxyRule is used as an input parameter in CreateL4ProxyRules, and if not specified, the default value off will be used; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	SessionPersist *string `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// Session persistence period, with a range of 30-3600, measured in seconds.
	// Note: This parameter is optional when L4ProxyRule is used as an input parameter in CreateL4ProxyRules. It is valid only when SessionPersist is set to on and defaults to 3600 if not specified. It is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// Rule tag. Accepts 1-50 arbitrary characters.
	// Note: This parameter is optional when L4ProxyRule is used as an input parameter in CreateL4ProxyRules; it is optional when L4ProxyRule is used as an input parameter in ModifyL4ProxyRules. If not specified, it will retain its existing value.
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`

	// Rule status. Valid values:<li>online: Enabled;</li>
	// <li>offline: Disabled;</li>
	// <li>progress: Deploying;</li>
	// <li>stopping: Disabling;</li>
	// <li>fail: Failed to deploy or disable.</li>
	// Note: Do not set this parameter when L4ProxyRule is used as an input parameter in CreateL4ProxyRules and ModifyL4ProxyRules.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// BuID.
	BuId *string `json:"BuId,omitnil,omitempty" name:"BuId"`
}

type L7OfflineLog struct {
	// Log domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Log query area. Valid values:
	// <li>`mainland`: Chinese mainland;</li>
	// <li>`overseas`: Global (outside the Chinese mainland). </li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Log packet name.	
	LogPacketName *string `json:"LogPacketName,omitnil,omitempty" name:"LogPacketName"`

	// Log download address.	
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// (Disused) Log packaging time. 
	LogTime *int64 `json:"LogTime,omitnil,omitempty" name:"LogTime"`

	// Start time of log packaging.
	LogStartTime *string `json:"LogStartTime,omitnil,omitempty" name:"LogStartTime"`

	// End time of the log package.
	LogEndTime *string `json:"LogEndTime,omitnil,omitempty" name:"LogEndTime"`

	// Original log size (in bytes).
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type LogFormat struct {
	// Predefined output format type for log delivery. Valid values:
	// <li>json: Use the predefined log output format JSON Lines, where fields in each log record are presented as key-value pairs;</li>
	// <li>csv: Use the predefined log output format csv, where fields in each log record is only presented as field values, excluding field names. </li>
	FormatType *string `json:"FormatType,omitnil,omitempty" name:"FormatType"`

	// A string added before each log delivery batch. Each log delivery batch may contain multiple log records.
	BatchPrefix *string `json:"BatchPrefix,omitnil,omitempty" name:"BatchPrefix"`

	// A string appended after each log delivery batch.
	BatchSuffix *string `json:"BatchSuffix,omitnil,omitempty" name:"BatchSuffix"`

	// A string added before each log record.
	RecordPrefix *string `json:"RecordPrefix,omitnil,omitempty" name:"RecordPrefix"`

	// A string appended after each log record.
	RecordSuffix *string `json:"RecordSuffix,omitnil,omitempty" name:"RecordSuffix"`

	// A string inserted between log records as a separator. Valid values:
	// <li>\n: Line break;</li>
	// <li>\t: Tab character;</li>
	// <li>,: Half-width comma. </li>
	RecordDelimiter *string `json:"RecordDelimiter,omitnil,omitempty" name:"RecordDelimiter"`

	// A string inserted between fields as a separator within a single log record. Valid values:
	// <li>\t: Tab character;</li>
	// <li>,: Half-width comma;</li>
	// <li>;: Half-width semicolon. </li>
	FieldDelimiter *string `json:"FieldDelimiter,omitnil,omitempty" name:"FieldDelimiter"`
}

type MaxAge struct {
	// Whether to follow the origin server. Values:
	// <li>`on`: Follow the origin server and ignore the field MaxAgeTime;</li>
	// <li>`off`: Do not follow the origin server and apply the field MaxAgeTime.</li>
	FollowOrigin *string `json:"FollowOrigin,omitnil,omitempty" name:"FollowOrigin"`

	// Specifies the maximum amount of time (in seconds). The maximum value is 365 days.
	// Note: The value `0` means not to cache.
	MaxAgeTime *int64 `json:"MaxAgeTime,omitnil,omitempty" name:"MaxAgeTime"`
}

// Predefined struct for user
type ModifyAccelerationDomainRequestParams struct {
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// Origin-pull protocol configuration. Values:
	// <li>`FOLLOW`: Follow the protocol of origin</li>
	// <li>`HTTP`: Send requests to the origin over HTTP</li>
	// <li>`HTTPS`: Send requests to the origin over HTTPS</li>
	// <li>The original configuration applies if this field is not specified.</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// Ports for HTTP origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTP`. The original configuration is used if it's not specified.
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// Ports for HTTPS origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTPS`. The original configuration is used if it's not specified.
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6 status. Values:
	// <li>`follow`: Follow the IPv6 configuration of the site</li>
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// <li>The original configuration applies if this field is not specified.</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
}

type ModifyAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Accelerated domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Details of the origin.
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// Origin-pull protocol configuration. Values:
	// <li>`FOLLOW`: Follow the protocol of origin</li>
	// <li>`HTTP`: Send requests to the origin over HTTP</li>
	// <li>`HTTPS`: Send requests to the origin over HTTPS</li>
	// <li>The original configuration applies if this field is not specified.</li>
	OriginProtocol *string `json:"OriginProtocol,omitnil,omitempty" name:"OriginProtocol"`

	// Ports for HTTP origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTP`. The original configuration is used if it's not specified.
	HttpOriginPort *uint64 `json:"HttpOriginPort,omitnil,omitempty" name:"HttpOriginPort"`

	// Ports for HTTPS origin-pull requests. Range: 1-65535. It takes effect when `OriginProtocol=FOLLOW/HTTPS`. The original configuration is used if it's not specified.
	HttpsOriginPort *uint64 `json:"HttpsOriginPort,omitnil,omitempty" name:"HttpsOriginPort"`

	// IPv6 status. Values:
	// <li>`follow`: Follow the IPv6 configuration of the site</li>
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	// <li>The original configuration applies if this field is not specified.</li>
	IPv6Status *string `json:"IPv6Status,omitnil,omitempty" name:"IPv6Status"`
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
	delete(f, "OriginProtocol")
	delete(f, "HttpOriginPort")
	delete(f, "HttpsOriginPort")
	delete(f, "IPv6Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be modified.
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// Status of the accelerated domain name. Values:
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to force suspension when the domain name has associated resources (such as alias domain names and traffic scheduling policies). Values:
	// <li>`true`: Suspend the domain name and all associated resources.</li>
	// <li>`true`: Do not suspend the domain name and all associated resources.</li>Default value: `false`.
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type ModifyAccelerationDomainStatusesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site related with the accelerated domain name.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of accelerated domain names to be modified.
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// Status of the accelerated domain name. Values:
	// <li>`online`: Enabled</li>
	// <li>`offline`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to force suspension when the domain name has associated resources (such as alias domain names and traffic scheduling policies). Values:
	// <li>`true`: Suspend the domain name and all associated resources.</li>
	// <li>`true`: Do not suspend the domain name and all associated resources.</li>Default value: `false`.
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: Off</li>
	// <li>`hosting`: Managed SSL certificate</li>
	// <li>`apply`: Free certificate</li>The original configuration will apply if this field is not specified.
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type ModifyAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The alias domain name.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// The target domain name.
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// Certificate configuration. Values:
	// <li>`none`: Off</li>
	// <li>`hosting`: Managed SSL certificate</li>
	// <li>`apply`: Free certificate</li>The original configuration will apply if this field is not specified.
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// The certificate ID. This field is required when `CertType=hosting`.
	CertId []*string `json:"CertId,omitnil,omitempty" name:"CertId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Status of the alias domain name. Values:
	// <li>`false`: Enable the alias domain name.</li>
	// <li>`true`: Disable the alias domain name.</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`

	// The alias domain name you want to modify its status. If it is left empty, the modify operation is not performed.
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
}

type ModifyAliasDomainStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Status of the alias domain name. Values:
	// <li>`false`: Enable the alias domain name.</li>
	// <li>`true`: Disable the alias domain name.</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`

	// The alias domain name you want to modify its status. If it is left empty, the modify operation is not performed.
	AliasNames []*string `json:"AliasNames,omitnil,omitempty" name:"AliasNames"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// The original configuration will apply if this field is not specified.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// L4 proxy mode. Valid values: 
	// <li>instance: Instance mode. </li>If it is not specified, instance is used by default.
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// IPv6 access configuration. The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Cross-MLC-border acceleration. The original configuration will apply if it is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Domain name or subdomain name when `ProxyType=hostname`; 
	// Instance name when `ProxyType=instance`.
	ProxyName *string `json:"ProxyName,omitnil,omitempty" name:"ProxyName"`

	// The session persistence duration. Value range: 30-3600 (in seconds).
	// The original configuration will apply if this field is not specified.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// L4 proxy mode. Valid values: 
	// <li>instance: Instance mode. </li>If it is not specified, instance is used by default.
	ProxyType *string `json:"ProxyType,omitnil,omitempty" name:"ProxyType"`

	// IPv6 access configuration. The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Cross-MLC-border acceleration. The original configuration will apply if it is not specified.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Origin server type. Valid values:
	// <li>custom: Manually added;</li>
	// <li>origins: Origin server group.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>The original configuration will apply if this field is not specified.
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// Origin server information:
	// <li>When `OriginType=custom`, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or ["test.com"].</li>
	// <li>When `OriginType=origins`, it indicates an origin group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	// 
	// The original configuration will apply if this field is not specified.
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>If not specified, this field uses the default value OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>If it is left empty, the default value `false` is used.
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// Rule tag. The original configuration will apply if it is not specified.
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
}

type ModifyApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Origin server type. Valid values:
	// <li>custom: Manually added;</li>
	// <li>origins: Origin server group.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// The access port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-90</li>
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// The protocol. Values:
	// <li>`TCP`: TCP protocol</li>
	// <li>`UDP`: UDP protocol</li>The original configuration will apply if this field is not specified.
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`

	// Origin server information:
	// <li>When `OriginType=custom`, it indicates one or more origin servers, such as ["8.8.8.8","9.9.9.9"] or ["test.com"].</li>
	// <li>When `OriginType=origins`, it indicates an origin group ID, such as ["origin-537f5b41-162a-11ed-abaa-525400c5da15"].</li>
	// 
	// The original configuration will apply if this field is not specified.
	OriginValue []*string `json:"OriginValue,omitnil,omitempty" name:"OriginValue"`

	// Passes the client IP. Values:
	// <li>`TOA`: Pass the client IP via TOA (available only when `Proto=TCP`).</li>
	// <li>`PPV1`: Pass the client IP via Proxy Protocol V1 (available only when `Proto=TCP`).</li>
	// <li>`PPV2`: Pass the client IP via Proxy Protocol V2.</li>
	// <li>`OFF`: Not pass the client IP.</li>If not specified, this field uses the default value OFF.
	ForwardClientIp *string `json:"ForwardClientIp,omitnil,omitempty" name:"ForwardClientIp"`

	// Whether to enable session persistence. Values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>If it is left empty, the default value `false` is used.
	SessionPersist *bool `json:"SessionPersist,omitnil,omitempty" name:"SessionPersist"`

	// Duration for the persistent session. The value takes effect only when `SessionPersist = true`.
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil,omitempty" name:"SessionPersistTime"`

	// The origin port, which can be:
	// <li>A single port, such as 80</li>
	// <li>A port range, such as 81-82</li>
	OriginPort *string `json:"OriginPort,omitnil,omitempty" name:"OriginPort"`

	// Rule tag. The original configuration will apply if it is not specified.
	RuleTag *string `json:"RuleTag,omitnil,omitempty" name:"RuleTag"`
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
	delete(f, "RuleTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The rule status. Values:
	// <li>`offline`: Disabled</li>
	// <li>`online`: Enabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The rule status. Values:
	// <li>`offline`: Disabled</li>
	// <li>`online`: Enabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The proxy status. Values:
	// <li>`offline`: The proxy is disabled.</li>
	// <li>`online`: The proxy is enabled.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The proxy ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// The proxy status. Values:
	// <li>`offline`: The proxy is disabled.</li>
	// <li>`online`: The proxy is enabled.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyCustomErrorPageRequestParams struct {
	// Custom response page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom response page name. The name must be 2-60 characters long.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Custom response page description, not exceeding 60 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Custom response page type, with values:<li>text/html. </li><li>application/json.</li><li>plain/text.</li><li>text/xml.</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// Custom response page content, not exceeding 2 KB.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ModifyCustomErrorPageRequest struct {
	*tchttp.BaseRequest
	
	// Custom response page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom response page name. The name must be 2-60 characters long.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Custom response page description, not exceeding 60 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Custom response page type, with values:<li>text/html. </li><li>application/json.</li><li>plain/text.</li><li>text/xml.</li>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// Custom response page content, not exceeding 2 KB.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *ModifyCustomErrorPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomErrorPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageId")
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ContentType")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomErrorPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomErrorPageResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomErrorPageResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomErrorPageResponseParams `json:"Response"`
}

func (r *ModifyCustomErrorPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomErrorPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Domain names that you need to modify the certificate configuration
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Certificate configuration mode. Values:
	// <li>`disable`: (Default) Do not configure the certificate</li>
	// <li>`eofreecert`: Use a free certificate provided by EdgeOne</li>
	// <li>`sslcert`: Configure an SSL certificate.</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// ID of the SSL certificate. It takes effect when `mode=sslcert`. To check the certificate ID, go to the [SSL Certificate](https://console.cloud.tencent.com/certoview) console.
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil,omitempty" name:"ServerCertInfo"`

	// Whether the certificate is managed by EdgeOne. Values:
	// <li>`none`: Not managed by EdgeOne</li>
	// <li>`apply`: Managed by EdgeOne</li>
	// Default value: `none`.
	//
	// Deprecated: ApplyType is deprecated.
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Domain names that you need to modify the certificate configuration
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Certificate configuration mode. Values:
	// <li>`disable`: (Default) Do not configure the certificate</li>
	// <li>`eofreecert`: Use a free certificate provided by EdgeOne</li>
	// <li>`sslcert`: Configure an SSL certificate.</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// ID of the SSL certificate. It takes effect when `mode=sslcert`. To check the certificate ID, go to the [SSL Certificate](https://console.cloud.tencent.com/certoview) console.
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil,omitempty" name:"ServerCertInfo"`

	// Whether the certificate is managed by EdgeOne. Values:
	// <li>`none`: Not managed by EdgeOne</li>
	// <li>`apply`: Managed by EdgeOne</li>
	// Default value: `none`.
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`
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
	delete(f, "Mode")
	delete(f, "ServerCertInfo")
	delete(f, "ApplyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyL4ProxyRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Application ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Specifies whether to enable IPv6 access. If this parameter is not filled, this configuration will not be modified. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance] (https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li> 
	// <li>off: Disable.</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Specifies whether to enable network optimization in the Chinese mainland. If this parameter is not filled, this configuration will not be modified. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance] (https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li> 
	// <li>off: Disable.</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

type ModifyL4ProxyRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Application ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Specifies whether to enable IPv6 access. If this parameter is not filled, this configuration will not be modified. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance] (https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li> 
	// <li>off: Disable.</li>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Specifies whether to enable network optimization in the Chinese mainland. If this parameter is not filled, this configuration will not be modified. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance] (https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values:<li>on: Enable;</li> 
	// <li>off: Disable.</li>
	AccelerateMainland *string `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`
}

func (r *ModifyL4ProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Ipv6")
	delete(f, "AccelerateMainland")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rules. A single request supports up to 200 forwarding rules.
	// Note: When L4ProxyRule is used here, RuleId is a required field; Protocol, PortRange, OriginType, OriginValue, OriginPortRange, ClientIPPassThroughMode, SessionPersist, SessionPersistTime, and RuleTag are all optional fields. No modification is made when no value is specified for those fields. Status should not be filled.
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

type ModifyL4ProxyRulesRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rules. A single request supports up to 200 forwarding rules.
	// Note: When L4ProxyRule is used here, RuleId is a required field; Protocol, PortRange, OriginType, OriginValue, OriginPortRange, ClientIPPassThroughMode, SessionPersist, SessionPersistTime, and RuleTag are all optional fields. No modification is made when no value is specified for those fields. Status should not be filled.
	L4ProxyRules []*L4ProxyRule `json:"L4ProxyRules,omitnil,omitempty" name:"L4ProxyRules"`
}

func (r *ModifyL4ProxyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "L4ProxyRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyRulesResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesStatusRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rule IDs. It supports up to 200 forwarding rules at a time.
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Status of forwarding rules. Valid values:
	// <li>online: Enabled;</li>
	// <li>offline: Disabled.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyL4ProxyRulesStatusRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// List of forwarding rule IDs. It supports up to 200 forwarding rules at a time.
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Status of forwarding rules. Valid values:
	// <li>online: Enabled;</li>
	// <li>offline: Disabled.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyL4ProxyRulesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleIds")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyRulesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyRulesStatusResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyRulesStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyRulesStatusResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyRulesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyRulesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyStatusRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Layer 4 proxy instance status. Valid values:<li>online: Enabled;</li>
	// <li>offline: Disabled.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyL4ProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Layer 4 proxy instance ID.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Layer 4 proxy instance status. Valid values:<li>online: Enabled;</li>
	// <li>offline: Disabled.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyL4ProxyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ProxyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ProxyStatusResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyL4ProxyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ProxyStatusResponseParams `json:"Response"`
}

func (r *ModifyL4ProxyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ProxyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupRequestParams struct {
	// Site ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// (Required) Origin group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Origin group name. It can contain 1 to 200 characters ([a-z], [A-Z], [0-9] and [_-]). The original configuration applies if this field is not specified.	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The origin grouptype. Values:
	// <li>`GENERAL`: General origin groups. It supports IPs and domain names. It can be referenced by DNS, Rule Engine, Layer 4 Proxy and General LoadBalancer.</li>
	// <li>`HTTP`: HTTP-specific origin groups. It supports IPs/domain names and object storage buckets. It can be referenced by acceleration domain names, rule engines and HTTP LoadBalancer. It cannot be referenced by L4 proxies. </li>The original configuration is used if it's not specified.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Origin information. The original configuration is used if it's not specified.
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// Host header used for origin-pull. It only works when `Type=HTTP`. If it's not specified, no specific Host header is configured. The `HostHeader` specified in `RuleEngine` takes a higher priority over this configuration. 
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type ModifyOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// (Required) Origin group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Origin group name. It can contain 1 to 200 characters ([a-z], [A-Z], [0-9] and [_-]). The original configuration applies if this field is not specified.	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The origin grouptype. Values:
	// <li>`GENERAL`: General origin groups. It supports IPs and domain names. It can be referenced by DNS, Rule Engine, Layer 4 Proxy and General LoadBalancer.</li>
	// <li>`HTTP`: HTTP-specific origin groups. It supports IPs/domain names and object storage buckets. It can be referenced by acceleration domain names, rule engines and HTTP LoadBalancer. It cannot be referenced by L4 proxies. </li>The original configuration is used if it's not specified.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Origin information. The original configuration is used if it's not specified.
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// Host header used for origin-pull. It only works when `Type=HTTP`. If it's not specified, no specific Host header is configured. The `HostHeader` specified in `RuleEngine` takes a higher priority over this configuration. 
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
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
	delete(f, "GroupId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Records")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyPlanRequestParams struct {
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Auto-renewal configuration item in a prepaid plan. If auto-renewal is enabled, the plan will be automatically renewed one day before it expires. This feature is only available for Personal, Basic, and Standard Edition Plans. If this field is not specified, the original configuration will be retained.
	RenewFlag *RenewFlag `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyPlanRequest struct {
	*tchttp.BaseRequest
	
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Auto-renewal configuration item in a prepaid plan. If auto-renewal is enabled, the plan will be automatically renewed one day before it expires. This feature is only available for Personal, Basic, and Standard Edition Plans. If this field is not specified, the original configuration will be retained.
	RenewFlag *RenewFlag `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *ModifyPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPlanResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPlanResponseParams `json:"Response"`
}

func (r *ModifyPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRealtimeLogDeliveryTaskRequestParams struct {
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The ID of the real-time log delivery task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The name of the real-time log delivery task, which is a combination of numbers, English letters, - and _, containing up to 200 characters. If this field is not filled in, the original configuration will be retained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// The status of the real-time log delivery task. Valid values:
	// <li>`enabled`: Enabled;</li>
	// <li>`disabled`: Disabled.</li>If this field is not filled in, the original configuration will be retained.
	DeliveryStatus *string `json:"DeliveryStatus,omitnil,omitempty" name:"DeliveryStatus"`

	// The list of entities (Layer 7 domains or Layer 4 proxy instances) corresponding to the real-time log delivery task. Valid value examples:
	// <li>Layer 7 domain: domain.example.com;</li>
	// <li>Layer 4 proxy instance: sid-2s69eb5wcms7.</li>If this field is not filled in, the original configuration will be retained.
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// The list of predefined fields for delivery. If this field is not filled in, the original configuration will be retained.
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// The list of custom fields for delivery, supporting extracting specified field values from HTTP request headers, response headers, and cookies. Each custom field name must be unique and the maximum number of fields is 200. If this field is not filled in, the original configuration will be retained.
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// Log delivery filter conditions. If this field is not filled in, all logs will be delivered.
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// The sampling ratio in permille. Value range: 1 to 1000. For example, 605 represents a sampling ratio of 60.5%. If this field is not filled in, the original configuration will be retained.
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// Output format for log delivery. If this field is not specified, the original configuration will be retained.Specifically, when TaskType is set to cls, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to input LogFormat.
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// The configuration information of the custom HTTP service. If this field is not filled in, the original configuration will be retained.
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// The configuration information of the AWS S3-compatible bucket. If this field is not filled in, the original configuration will be retained.
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

type ModifyRealtimeLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// Zone ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The ID of the real-time log delivery task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The name of the real-time log delivery task, which is a combination of numbers, English letters, - and _, containing up to 200 characters. If this field is not filled in, the original configuration will be retained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// The status of the real-time log delivery task. Valid values:
	// <li>`enabled`: Enabled;</li>
	// <li>`disabled`: Disabled.</li>If this field is not filled in, the original configuration will be retained.
	DeliveryStatus *string `json:"DeliveryStatus,omitnil,omitempty" name:"DeliveryStatus"`

	// The list of entities (Layer 7 domains or Layer 4 proxy instances) corresponding to the real-time log delivery task. Valid value examples:
	// <li>Layer 7 domain: domain.example.com;</li>
	// <li>Layer 4 proxy instance: sid-2s69eb5wcms7.</li>If this field is not filled in, the original configuration will be retained.
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// The list of predefined fields for delivery. If this field is not filled in, the original configuration will be retained.
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// The list of custom fields for delivery, supporting extracting specified field values from HTTP request headers, response headers, and cookies. Each custom field name must be unique and the maximum number of fields is 200. If this field is not filled in, the original configuration will be retained.
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// Log delivery filter conditions. If this field is not filled in, all logs will be delivered.
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// The sampling ratio in permille. Value range: 1 to 1000. For example, 605 represents a sampling ratio of 60.5%. If this field is not filled in, the original configuration will be retained.
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// Output format for log delivery. If this field is not specified, the original configuration will be retained.Specifically, when TaskType is set to cls, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to input LogFormat.
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// The configuration information of the custom HTTP service. If this field is not filled in, the original configuration will be retained.
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// The configuration information of the AWS S3-compatible bucket. If this field is not filled in, the original configuration will be retained.
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`
}

func (r *ModifyRealtimeLogDeliveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealtimeLogDeliveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	delete(f, "TaskName")
	delete(f, "DeliveryStatus")
	delete(f, "EntityList")
	delete(f, "Fields")
	delete(f, "CustomFields")
	delete(f, "DeliveryConditions")
	delete(f, "Sample")
	delete(f, "LogFormat")
	delete(f, "CustomEndpoint")
	delete(f, "S3")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRealtimeLogDeliveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRealtimeLogDeliveryTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRealtimeLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRealtimeLogDeliveryTaskResponseParams `json:"Response"`
}

func (r *ModifyRealtimeLogDeliveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealtimeLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The rule name. It is a string that can contain 1–255 characters.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The rule name. It is a string that can contain 1–255 characters.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type ModifySecurityIPGroupRequestParams struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP group configuration.
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`

	// Operation type. Valid values: 
	// <li>`append`: Add information of `Content` to `IPGroup`;</li>
	// <li>`remove`: Delete information of `Content` from `IPGroup`;</li>
	// <li>`update`: Replace all information of `IPGroup` and modify the IPGroup name.</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ModifySecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IP group configuration.
	IPGroup *IPGroup `json:"IPGroup,omitnil,omitempty" name:"IPGroup"`

	// Operation type. Valid values: 
	// <li>`append`: Add information of `Content` to `IPGroup`;</li>
	// <li>`remove`: Delete information of `Content` from `IPGroup`;</li>
	// <li>`update`: Replace all information of `IPGroup` and modify the IPGroup name.</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Security configuration.
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil,omitempty" name:"SecurityConfig"`

	// Subdomain/application name.
	// 
	// Note: When both this parameter and the TemplateId parameter are specified, this parameter will not take effect. Do not specify this parameter and the TemplateId parameter at the same time.
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// Specifies the policy template ID, or the site's global policy.
	// - To configure a policy template, specify the policy template ID.
	// - To configure the site's global policy, use the @ZoneLevel@Domain parameter value.
	// 
	// Note: When this parameter is used, the Entity parameter will not take effect. Do not use this parameter and the Entity parameter at the same time.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Security configuration.
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil,omitempty" name:"SecurityConfig"`

	// Subdomain/application name.
	// 
	// Note: When both this parameter and the TemplateId parameter are specified, this parameter will not take effect. Do not specify this parameter and the TemplateId parameter at the same time.
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// Specifies the policy template ID, or the site's global policy.
	// - To configure a policy template, specify the policy template ID.
	// - To configure the site's global policy, use the @ZoneLevel@Domain parameter value.
	// 
	// Note: When this parameter is used, the Entity parameter will not take effect. Do not use this parameter and the Entity parameter at the same time.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Access mode of the site. Values:
	// <li> `full`: Access through a name server.</li>
	// <li> `partial`: Access through a CNAME record. A site using domainless access can only switch to CNAME access. </li>The original configuration applies if this field is not specified.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The custom name servers. The original configuration applies if this field is not specified. It is not allowed to pass this field when a site is connected without using a domain name.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// The region where the site requests access. Values:
	// <li> `global`: Global coverage</li>
	// <li> `mainland`: Chinese mainland</li>
	// <li> `overseas`: Outside the Chinese mainland </li>It is not allowed to pass this field when a site is connected without using a domain name.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Name of the site. This field takes effect only when the site switches from domainless access to CNAME access.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Access mode of the site. Values:
	// <li> `full`: Access through a name server.</li>
	// <li> `partial`: Access through a CNAME record. A site using domainless access can only switch to CNAME access. </li>The original configuration applies if this field is not specified.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The custom name servers. The original configuration applies if this field is not specified. It is not allowed to pass this field when a site is connected without using a domain name.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// The region where the site requests access. Values:
	// <li> `global`: Global coverage</li>
	// <li> `mainland`: Chinese mainland</li>
	// <li> `overseas`: Outside the Chinese mainland </li>It is not allowed to pass this field when a site is connected without using a domain name.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Name of the site. This field takes effect only when the site switches from domainless access to CNAME access.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Cache expiration time configuration
	// The original configuration will apply if this field is not specified.
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// The node cache key configuration.
	// The original configuration will apply if this field is not specified.
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// The browser cache configuration.
	// The original configuration will apply if this field is not specified.
	MaxAge *MaxAge `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// The offline cache configuration.
	// The original configuration will apply if this field is not specified.
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil,omitempty" name:"OfflineCache"`

	// QUIC access configuration. 
	// The original configuration will apply if it is not specified.
	Quic *Quic `json:"Quic,omitnil,omitempty" name:"Quic"`

	// POST transport configuration. 
	// The original configuration will apply if it is not specified.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil,omitempty" name:"PostMaxSize"`

	// The smart compression configuration.
	// The original configuration will apply if this field is not specified.
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration. 
	// The original configuration will apply if it is not specified.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration. 
	// The original configuration will apply if it is not specified.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// HTTPS acceleration configuration. 
	// The original configuration will apply if it is not specified.
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// The origin server configuration.
	// The original configuration will apply if this field is not specified.
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// The smart acceleration configuration.
	// The original configuration will apply if this field is not specified.
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// WebSocket configuration. 
	// The original configuration will apply if it is not specified.
	WebSocket *WebSocket `json:"WebSocket,omitnil,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration. 
	// The original configuration will apply if it is not specified.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil,omitempty" name:"ClientIpHeader"`

	// The cache prefresh configuration.
	// The original configuration will apply if this field is not specified.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil,omitempty" name:"CachePrefresh"`

	// Ipv6 access configuration. 
	// The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull. 
	// The original configuration will apply if it is not specified.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support. 
	// The original configuration will apply if it is not specified.
	Grpc *Grpc `json:"Grpc,omitnil,omitempty" name:"Grpc"`

	// Image optimization. 
	// It is disabled if this parameter is not specified.
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// Standard debugging configuration.
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// Site ID to modify.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Cache expiration time configuration
	// The original configuration will apply if this field is not specified.
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// The node cache key configuration.
	// The original configuration will apply if this field is not specified.
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// The browser cache configuration.
	// The original configuration will apply if this field is not specified.
	MaxAge *MaxAge `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// The offline cache configuration.
	// The original configuration will apply if this field is not specified.
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil,omitempty" name:"OfflineCache"`

	// QUIC access configuration. 
	// The original configuration will apply if it is not specified.
	Quic *Quic `json:"Quic,omitnil,omitempty" name:"Quic"`

	// POST transport configuration. 
	// The original configuration will apply if it is not specified.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil,omitempty" name:"PostMaxSize"`

	// The smart compression configuration.
	// The original configuration will apply if this field is not specified.
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration. 
	// The original configuration will apply if it is not specified.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration. 
	// The original configuration will apply if it is not specified.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// HTTPS acceleration configuration. 
	// The original configuration will apply if it is not specified.
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// The origin server configuration.
	// The original configuration will apply if this field is not specified.
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// The smart acceleration configuration.
	// The original configuration will apply if this field is not specified.
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// WebSocket configuration. 
	// The original configuration will apply if it is not specified.
	WebSocket *WebSocket `json:"WebSocket,omitnil,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration. 
	// The original configuration will apply if it is not specified.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil,omitempty" name:"ClientIpHeader"`

	// The cache prefresh configuration.
	// The original configuration will apply if this field is not specified.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil,omitempty" name:"CachePrefresh"`

	// Ipv6 access configuration. 
	// The original configuration will apply if it is not specified.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Whether to carry the location information of the client IP during origin-pull. 
	// The original configuration will apply if it is not specified.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support. 
	// The original configuration will apply if it is not specified.
	Grpc *Grpc `json:"Grpc,omitnil,omitempty" name:"Grpc"`

	// Image optimization. 
	// It is disabled if this parameter is not specified.
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// Standard debugging configuration.
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The site status. Values:
	// <li>`false`: Disabled</li>
	// <li>`true`: Enabled</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`
}

type ModifyZoneStatusRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The site status. Values:
	// <li>`false`: Disabled</li>
	// <li>`true`: Enabled</li>
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type NormalAction struct {
	// Feature name. For details, see [DescribeRulesSetting](https://intl.cloud.tencent.com/document/product/1552/80618?from_cn_redirect=1) API
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Parameter
	Parameters []*RuleNormalActionParams `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type NsVerification struct {
	// The DNS server address assigned to the user when connecting a site to EO via NS. You need to switch the NameServer of the domain name to this address.
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`
}

type OfflineCache struct {
	// Whether offline cache is enabled. Valid values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type Origin struct {
	// Primary origin server list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origins []*string `json:"Origins,omitnil,omitempty" name:"Origins"`

	// The list of backup origin servers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupOrigins []*string `json:"BackupOrigins,omitnil,omitempty" name:"BackupOrigins"`

	// Origin-pull protocol configuration. Values:
	// <li>`http`: Force HTTP for origin-pull.</li>
	// <li>`follow`: Follow protocol.</li>
	// <li>`https`: Force HTTPS for origin-pull.</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullProtocol *string `json:"OriginPullProtocol,omitnil,omitempty" name:"OriginPullProtocol"`

	// Whether to allow private access to buckets when `OriginType=cos`. Valid values: 
	// <li>`on`: Private access;</li>
	// <li>`off`: Public access.</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CosPrivateAccess *string `json:"CosPrivateAccess,omitnil,omitempty" name:"CosPrivateAccess"`
}

type OriginDetail struct {
	// Origin server type, with values:
	// <li>IP_DOMAIN: IPv4, IPv6, or domain name type origin server;</li>
	// <li>COS: Tencent Cloud COS origin server;</li>
	// <li>AWS_S3: AWS S3 origin server;</li>
	// <li>ORIGIN_GROUP: origin server group type origin server;</li>
	// <li>VODEO: VOD on EO;</li>
	// <li>VOD: Video on Demand;</li>
	// <li>SPACE: origin server uninstallation. Currently only available to the allowlist;</li>
	// <li>LB: load balancing. Currently only available to the allowlist. </li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Origin server address, which varies according to the value of OriginType:
	// <li>When OriginType = IP_DOMAIN, this parameter should be an IPv4 address, an IPv6 address, or a domain name;</li>
	// <li>When OriginType = COS, this parameter should be the access domain name of the COS bucket;</li>
	// <li>When OriginType = AWS_S3, this parameter should be the access domain name of the S3 bucket;</li>
	// <li>When OriginType = ORIGIN_GROUP, this parameter should be the origin server group ID;</li>
	// <li>When OriginType = VOD, this parameter should be the VOD application ID;</li>
	// <li>When OriginType = VODEO, if VodeoDistributionRange = ALL, then this parameter is "all-buckets-in-vodeo-application"; if VodeoDistributionRange = Bucket, then this parameter is the domain name of the corresponding bucket. </li>
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// Secondary origin group ID. This parameter is valid only when OriginType is ORIGIN_GROUP and a secondary origin group is configured.
	BackupOrigin *string `json:"BackupOrigin,omitnil,omitempty" name:"BackupOrigin"`

	// Primary origin group name. This parameter returns a value when OriginType is ORIGIN_GROUP.
	OriginGroupName *string `json:"OriginGroupName,omitnil,omitempty" name:"OriginGroupName"`

	// Secondary origin group name. This parameter is valid only when OriginType is ORIGIN_GROUP and a secondary origin group is configured.
	BackOriginGroupName *string `json:"BackOriginGroupName,omitnil,omitempty" name:"BackOriginGroupName"`

	// Whether access to the private object storage origin server is allowed. This parameter is valid only when the origin server type OriginType is COS or AWS_S3. Valid values:
	// <li>on: Enable private authentication;</li>
	// <li>off: Disable private authentication. </li>
	// If this field is not specified, the default value 'off' will be used.
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// Private authentication parameter. This parameter is valid only when PrivateAccess is on.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil,omitempty" name:"PrivateParameters"`

	// MO sub-application ID
	VodeoSubAppId *int64 `json:"VodeoSubAppId,omitnil,omitempty" name:"VodeoSubAppId"`

	// MO distribution range. Valid values: <li>All: all</li> <li>Bucket: bucket</li>
	VodeoDistributionRange *string `json:"VodeoDistributionRange,omitnil,omitempty" name:"VodeoDistributionRange"`

	// MO bucket ID, required when the distribution range (DistributionRange) is bucket (Bucket)
	VodeoBucketId *string `json:"VodeoBucketId,omitnil,omitempty" name:"VodeoBucketId"`
}

type OriginGroup struct {
	// The ID of the origin group.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The name of the origin group.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The origin group type. Values:
	// <li>`GENERAL`: General origin group</li>
	// <li>`HTTP`: HTTP-specific origin group</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Details of the origin record.
	Records []*OriginRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// List of instances referencing this origin group.	
	References []*OriginGroupReference `json:"References,omitnil,omitempty" name:"References"`

	// Creation time of the origin group.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The update time of the origin group.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Origin-pull host header
	// Note: This field may return·null, indicating that no valid values can be obtained.
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`
}

type OriginGroupReference struct {
	// Services referencing the origin group. Values:
	// <li>`AccelerationDomain`: Acceleration domain name</li>
	// <li>`RuleEngine`: Rules engine</li>
	// <li>`Loadbalance`: Load balancer</li>
	// <li>`ApplicationProxy`: L4 proxy</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// ID of the instances referencing the origin group
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the instance referencing the origin group
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type OriginInfo struct {
	// Origin server type, with values:
	// <li>IP_DOMAIN: IPv4, IPv6, or domain name type origin server;</li>
	// <li>COS: Tencent Cloud COS origin server;</li>
	// <li>AWS_S3: AWS S3 origin server;</li>
	// <li>ORIGIN_GROUP: origin server group type origin server;</li>
	//  <li>VOD: Video on Demand;</li>
	// <li>SPACE: origin server uninstallation. Currently only available to the allowlist;</li>
	// <li>LB: load balancing. Currently only available to the allowlist. </li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Origin server address, which varies according to the value of OriginType:
	// <li>When OriginType = IP_DOMAIN, fill in an IPv4 address, an IPv6 address, or a domain name;</li>
	// <li>When OriginType = COS, fill in the access domain name of the COS bucket;</li>
	// <li>When OriginType = AWS_S3, fill in the access domain name of the S3 bucket;</li>
	// <li>When OriginType = ORIGIN_GROUP, fill in the origin server group ID;</li>
	// <li>When OriginType = VOD, fill in the VOD application ID;</li>
	// <li>When OriginType = LB, fill in the Cloud Load Balancer instance ID. This feature is currently only available to the allowlist;</li>
	// <li>When OriginType = SPACE, fill in the origin server uninstallation space ID. This feature is currently only available to the allowlist.</li>
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// The ID of the secondary origin group. This parameter is valid only when OriginType is ORIGIN_GROUP. This field indicates the old version capability, which cannot be configured or modified on the control panel after being called. Please submit a ticket if required.
	BackupOrigin *string `json:"BackupOrigin,omitnil,omitempty" name:"BackupOrigin"`

	// Whether access to the private Cloud Object Storage origin server is allowed. This parameter is valid only when OriginType is COS or AWS_S3. Valid values:
	// <li>on: Enable private authentication;</li>
	// <li>off: Disable private authentication.</li>
	// If it is not specified, the default value is off.
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// Private authentication parameter. This parameter is valid only when PrivateAccess is on.
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil,omitempty" name:"PrivateParameters"`

	// VODEO sub-application ID. This parameter is required when OriginType is VODEO.
	VodeoSubAppId *int64 `json:"VodeoSubAppId,omitnil,omitempty" name:"VodeoSubAppId"`

	// VOD on EO distribution range. This parameter is required when OriginType = VODEO. The values are: 
	// <li>All: all buckets under the current application;</li> 
	// <li>Bucket: a specified bucket.</li>
	VodeoDistributionRange *string `json:"VodeoDistributionRange,omitnil,omitempty" name:"VodeoDistributionRange"`

	// VODEO storage bucket ID. This parameter is required when OriginType is VODEO and VodeoDistributionRange is Bucket.
	VodeoBucketId *string `json:"VodeoBucketId,omitnil,omitempty" name:"VodeoBucketId"`
}

type OriginProtectionInfo struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of domain names.
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// List of proxy IDs.
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// The existing intermediate IPs.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CurrentIPWhitelist *IPWhitelist `json:"CurrentIPWhitelist,omitnil,omitempty" name:"CurrentIPWhitelist"`

	// Whether the intermediate IP update is needed for the site. Values:
	// <li>`true`: Update needed;</li>
	// <li>`false`: Update not needed.</li>
	NeedUpdate *bool `json:"NeedUpdate,omitnil,omitempty" name:"NeedUpdate"`

	// Status of the origin protection configuration. Values:
	// <li>`online`: Origin protection is activated;</li>
	// <li>`offline`: Origin protection is disabled.</li>
	// <li>`nonactivate`: Origin protection is not activated. This value is returned only when the feature is not activated before it’s used.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether origin protection is supported in the plan. Values:
	// <li>`true`: Origin protection supported;</li>
	// <li>`false`: Origin protection not supported.</li>
	PlanSupport *bool `json:"PlanSupport,omitnil,omitempty" name:"PlanSupport"`

	// Differences between the latest and existing intermediate IPs.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiffIPWhitelist *DiffIPWhitelist `json:"DiffIPWhitelist,omitnil,omitempty" name:"DiffIPWhitelist"`
}

type OriginRecord struct {
	// The origin record value, which can be an IPv4/IPv6 address or a domain name.
	Record *string `json:"Record,omitnil,omitempty" name:"Record"`

	// The origin type. Values:
	// <li>`IP_DOMAIN`: IPv4/IPv6 address or domain name</li>
	// <li>`COS`: COS bucket address</li>
	// <li>`AWS_S3`: AWS S3 bucket address
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The origin record ID.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Weight of an origin. Range: 0-100. If it is not specified, a random weight is assigned. If `0` is passed in, there is no traffic scheduled to this origin.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Whether to enable private authentication. It is valid when `OriginType=COS/AWS_S3`. Values:
	// <li>`true`: Yes.</li>
	// <li>`false`: No.</li>Default: `false`.
	Private *bool `json:"Private,omitnil,omitempty" name:"Private"`

	// Private authentication parameters. This field is valid when `Private=true`.
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil,omitempty" name:"PrivateParameters"`
}

type OwnershipVerification struct {
	// u200cInformation required for authentication using DNS resolution. It's applicable to sites connected via CNAME. See [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1#7af6ecf8-afca-4e35-8811-b5797ed1bde5).
	//  
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DnsVerification *DnsVerification `json:"DnsVerification,omitnil,omitempty" name:"DnsVerification"`

	// u200cInformation required for verifying via a file. It's applicable to sites connected via CNAMEs. See [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1#7af6ecf8-afca-4e35-8811-b5797ed1bde5).
	//  
	// Note: This field may return·null, indicating that no valid values can be obtained.
	FileVerification *FileVerification `json:"FileVerification,omitnil,omitempty" name:"FileVerification"`

	// u200cInformation required for switching DNS servers. It's applicable to sites connected via NSs. For details, see [Modifying DNS Server](https://intl.cloud.tencent.com/document/product/1552/90452?from_cn_redirect=1).
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NsVerification *NsVerification `json:"NsVerification,omitnil,omitempty" name:"NsVerification"`
}

type PartialModule struct {
	// The module. Values:
	// <li>`waf`: Managed rules</li>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// List of managed rule IDs to be skipped.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Include []*int64 `json:"Include,omitnil,omitempty" name:"Include"`
}

type PlanInfo struct {
	// Settlement currency. Values:
	// <li>`CNY`: Settled by Chinese RMB;</li>
	// <li>`USD`: Settled by US dollars.</li>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Traffic quota of the plan. It includes the traffic for security acceleration, content acceleration and smart acceleration. Unit: byte.
	Flux *uint64 `json:"Flux,omitnil,omitempty" name:"Flux"`

	// Settlement cycle. Values:
	// <li>`y`: Settled by year;</li>
	// <li>`m`: Settled by month;</li>
	// <li>`h`: Settled by hour;</li>
	// <li>`M`: Settled by minute;</li>
	// <li>`s`: Settled by second.</li>
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

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
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// Plan price (in CNY fen/US cent). The price unit depends on the settlement currency.
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// Quota on security acceleration requests
	Request *uint64 `json:"Request,omitnil,omitempty" name:"Request"`

	// Number of sites to be bound to the plan
	SiteNumber *uint64 `json:"SiteNumber,omitnil,omitempty" name:"SiteNumber"`

	// The acceleration region. Values:
	// <li>`mainland`: Chinese mainland</li>
	// <li>`overseas`: Global (Chinese mainland not included)</li>
	// <li>`global`: Global (Chinese mainland included)</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type PostMaxSize struct {
	// Whether to enable POST upload limit (default limit: 32 MB). Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Maximum size. Value range: 1-500 MB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type PrepaidPlanParam struct {
	// Prepaid plan duration, unit: month. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	// 
	// If this field is not specified, the default value '1' will be used.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// The auto-renewal flag for prepaid plan has the following values:
	// <li> on: Enable auto-renewal;</li>
	// <li> off: Disable auto-renewal. </li>
	// If this field is not specified, the default value 'off' will be used. When auto-renewal is enabled, it defaults to renewing for one month.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type PrivateParameter struct {
	// The name of the private authentication parameter. Valid values:
	// <li>AccessKeyId: Access Key ID for authentication;</li>
	// <li>SecretAccessKey: Secret Access Key for authentication;</li>
	// <li>SignatureVersion: Authentication version, v2 or v4;</li>
	// <li>Region: The region of the storage bucket.</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The parameter value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QueryCondition struct {
	// The key of QueryCondition.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The conditional operator. Values:
	// <li>`equals`: Equals</li>
	// <li>`notEquals`: Does not equal</li>
	// <li>`include`: Contains</li>
	// <li>`notInclude`: Does not contain</li>
	// <li>`startWith`: Starts with</li>
	// <li>`notStartWith`: Does not start with</li>
	// <li>`endWith`: Ends with</li>
	// <li>`notEndWith`: Does not end with</li>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// The value of QueryCondition.
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QueryString struct {
	// Whether to use `QueryString` as part of `CacheKey`. Values:
	// <li>`on`: Yes</li>
	// <li>`off`: No</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Specifies how to use query strings in the cache key. Values:
	// <li>`includeCustom`: `Include partial query strings.</li>
	// <li>`excludeCustom`: Exclude partial query strings.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Array of query strings used/excluded
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Quic struct {
	// Whether to enable QUIC. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type Quota struct {

	Batch *int64 `json:"Batch,omitnil,omitempty" name:"Batch"`

	// Daily submission quota limit.
	Daily *int64 `json:"Daily,omitnil,omitempty" name:"Daily"`

	// Remaining daily submission quota.
	DailyAvailable *int64 `json:"DailyAvailable,omitnil,omitempty" name:"DailyAvailable"`

	// Type of cache purging/pre-warming. Values:
	// <li>`purge_prefix`: Purge by prefix</li>
	// <li>`purge_url`: Purge by URL</li>
	// <li>`purge_host`: Purge by hostname</li>
	// <li>`purge_all`: Purge all caches</li>
	// <li>`purge_cache_tag`: Purge by cache tag</li><li>`prefetch_url`: Pre-warm by URL</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type RateLimitConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The settings of the custom rate limiting rule. If it is null, the settings that were last configured will be used.
	RateLimitUserRules []*RateLimitUserRule `json:"RateLimitUserRules,omitnil,omitempty" name:"RateLimitUserRules"`

	// The settings of the rate limiting template. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RateLimitTemplate *RateLimitTemplate `json:"RateLimitTemplate,omitnil,omitempty" name:"RateLimitTemplate"`

	// The client filtering settings. If it is null, the settings that were last configured will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RateLimitIntelligence *RateLimitIntelligence `json:"RateLimitIntelligence,omitnil,omitempty" name:"RateLimitIntelligence"`

	// The custom rate limiting rules. If it is `null`, the previous settings is used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RateLimitCustomizes []*RateLimitUserRule `json:"RateLimitCustomizes,omitnil,omitempty" name:"RateLimitCustomizes"`
}

type RateLimitIntelligence struct {
	// Whether to enable configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Action to be executed. Values:
	// <li>`monitor`: Observe</li>
	// <li>`alg`: Challenge</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The rule ID, which is only used as a response parameter.
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type RateLimitTemplate struct {
	// The mode. Values:
	// <li>`sup_loose`: Super loose</li>
	// <li>`loose`: Loose</li>
	// <li>`emergency`: Emergency</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`close`: Off</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// The action. Values:
	// <li>`alg`: JavaScript challenge</li>
	// <li>`monitor`: Observe</li>If it is left empty, the default value `alg` is used.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The settings of the rate limiting template. It is only used as an output parameter.
	RateLimitTemplateDetail *RateLimitTemplateDetail `json:"RateLimitTemplateDetail,omitnil,omitempty" name:"RateLimitTemplateDetail"`
}

type RateLimitTemplateDetail struct {
	// Note: This field may return null, which indicates a failure to obtain a valid value.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// The unique ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Note: This field may return null, which indicates a failure to obtain a valid value.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Note: This field may return null, which indicates a failure to obtain a valid value.
	PunishTime *int64 `json:"PunishTime,omitnil,omitempty" name:"PunishTime"`

	// The request threshold. Value range: 0-4294967294.
	Threshold *int64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// The statistical period. Value range: 0-120 seconds.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type RateLimitUserRule struct {
	// The request threshold. Value range: 0-4294967294.
	Threshold *int64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// The statistical period. The value can be 10, 20, 30, 40, 50, or 60 seconds.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// The rule name, which consists of only letters, digits, and underscores and cannot start with an underscore.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Action. Values:<li>`monitor`: Observe;</li><li>`drop`: Block;</li><li>`redirect`: Redirect;</li><li>`page`: Return a specific page;</li><li>`alg`: JavaScript challenge. </li>	
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// The amount of time taken to perform the action. Value range: 0 seconds - 2 days.
	PunishTime *int64 `json:"PunishTime,omitnil,omitempty" name:"PunishTime"`

	// The time unit. Values:
	// <li>`second`: Second</li>
	// <li>`minutes`: Minute</li>
	// <li>`hour`: Hour</li>
	PunishTimeUnit *string `json:"PunishTimeUnit,omitnil,omitempty" name:"PunishTimeUnit"`

	// The rule status. Values:
	// <li>`on`: Enabled</li>
	// <li>`off`: Disabled</li>Default value: `on`
	RuleStatus *string `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// The rule details.
	AclConditions []*AclCondition `json:"AclConditions,omitnil,omitempty" name:"AclConditions"`

	// The rule weight. Value range: 0-100.
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// Rule ID, which is only used as an output parameter.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// The filter. Values:
	// <li>`sip`: Client IP</li>
	// This parameter is left empty by default.
	FreqFields []*string `json:"FreqFields,omitnil,omitempty" name:"FreqFields"`

	// Update time. It is only used as a response parameter, and defaults to the current time. 
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Query scope. Values:
	// <li>`source_to_eo`: (Response) Traffic going from the origin to EdgeOne.</li>
	// <li>`client_to_eo`: (Request) Traffic going from the client to EdgeOne.</li>
	// Default: `source_to_eo`.
	FreqScope []*string `json:"FreqScope,omitnil,omitempty" name:"FreqScope"`

	// Name of the custom return page. It's required when `Action=page`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// ID of custom response. The ID can be obtained via the `DescribeCustomErrorPages` API. It's required when `Action=page`.	
	CustomResponseId *string `json:"CustomResponseId,omitnil,omitempty" name:"CustomResponseId"`

	// The response code to trigger the return page. It's required when `Action=page`. Value: 100-600. 3xx response codes are not supported. Default value: 567.
	ResponseCode *int64 `json:"ResponseCode,omitnil,omitempty" name:"ResponseCode"`

	// The redirection URL. It's required when `Action=redirect`.
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`
}

type RealtimeLogDeliveryTask struct {
	// The ID of the real-time log delivery task.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The name of the real-time log delivery task.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// The status of the real-time log delivery task. Valid values:<li>enabled: Enabled;</li><li>disabled: Disabled;</li><li>deleted: Abnormal deletion. Check whether the destination Tencent Cloud CLS log set/log topic has been deleted.</li>
	DeliveryStatus *string `json:"DeliveryStatus,omitnil,omitempty" name:"DeliveryStatus"`

	// The type of the real-time log delivery task. Valid values:<li>cls: Push to Tencent Cloud CLS;</li><li>custom_endpoint: Push to a custom HTTP(S) address;</li><li>s3: Push to an AWS S3-compatible bucket address.</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// The list of entities (Layer 7 domains or Layer 4 proxy instances) corresponding to the real-time log delivery task. Valid value examples:<li>Layer 7 domain: domain.example.com;</li><li>Layer 4 proxy instance: sid-2s69eb5wcms7.</li>
	EntityList []*string `json:"EntityList,omitnil,omitempty" name:"EntityList"`

	// The data delivery type. Valid values: <li>domain: Site acceleration log;</li><li>application: Layer 4 proxy log;</li><li>web-rateLiming: Rate limit and CC attack defense log;</li><li>web-attack: Managed rule log;</li><li>web-rule: Custom rule log;</li><li>web-bot: Bot management log.</li>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// The data delivery area. Valid values:<li>mainland: Within the Chinese mainland;</li><li>overseas: Global (excluding the Chinese mainland).</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The list of predefined fields for delivery.
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// The list of custom fields for delivery.
	CustomFields []*CustomField `json:"CustomFields,omitnil,omitempty" name:"CustomFields"`

	// Log delivery filter conditions.
	DeliveryConditions []*DeliveryCondition `json:"DeliveryConditions,omitnil,omitempty" name:"DeliveryConditions"`

	// The sampling ratio in permille. Value range: 1 to 1000. For example, 605 represents a sampling ratio of 60.5%.
	Sample *uint64 `json:"Sample,omitnil,omitempty" name:"Sample"`

	// Output format for log delivery. When the output parameter is null, the default format is used, which works as follows:
	// <li>When TaskType is 'custom_endpoint', the default format is an array of JSON objects, with each JSON object representing a log entry;</li>
	// <li>When TaskType is 's3', the default format is JSON Lines. </li>
	// Note: This field may return 'null', which indicates a failure to obtain a valid value.
	LogFormat *LogFormat `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	CLS *CLSTopic `json:"CLS,omitnil,omitempty" name:"CLS"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	CustomEndpoint *CustomEndpoint `json:"CustomEndpoint,omitnil,omitempty" name:"CustomEndpoint"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	S3 *S3 `json:"S3,omitnil,omitempty" name:"S3"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type RenewFlag struct {
	// The auto-renewal flag for prepaid plan has the following values:
	// <li> on: Enable auto-renewal;</li>
	// <li> off: Disable auto-renewal. </li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

// Predefined struct for user
type RenewPlanRequestParams struct {
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Renewal plan duration, unit: month. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use a voucher. Valid values: <li>true: Yes;</li><li>false: No. </li> If this field is not specified, the default value 'false' will be used.
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

type RenewPlanRequest struct {
	*tchttp.BaseRequest
	
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Renewal plan duration, unit: month. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use a voucher. Valid values: <li>true: Yes;</li><li>false: No. </li> If this field is not specified, the default value 'false' will be used.
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

func (r *RenewPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Period")
	delete(f, "AutoUseVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewPlanResponseParams struct {
	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewPlanResponse struct {
	*tchttp.BaseResponse
	Response *RenewPlanResponseParams `json:"Response"`
}

func (r *RenewPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// The resource ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Billing mode
	// `0`: Pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The effective time.
	EnableTime *string `json:"EnableTime,omitnil,omitempty" name:"EnableTime"`

	// The expiration time.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// The plan status. Values:
	// <li>`normal`: Normal</li>
	// <li>`isolated`: Isolated</li>
	// <li>`destroyed`: Terminated</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Pricing query parameter
	Sv []*Sv `json:"Sv,omitnil,omitempty" name:"Sv"`

	// Whether to enable auto-renewal. Values:
	// <li>`0`: Default status.</li>
	// <li>`1`: Enable auto-renewal.</li>
	// <li>`2`: Disable auto-renewal.</li>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// ID of the resource associated with the plan.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Applicable area. Values:
	// <li>`mainland`: Chinese mainland</li>
	// <li>`overseas`: Regions outside the Chinese mainland</li>
	// <li>`global`: Global</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The resource type. Values:
	// <li>`plan`: Plan resources</li>
	// <li>`pay-as-you-go`: Pay-as-you-go resources </li>
	// <li>`value-added`: Value-added resources </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// The sites that are associated with the current resources.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneNumber *int64 `json:"ZoneNumber,omitnil,omitempty" name:"ZoneNumber"`

	// Resource tag type. Valid values:
	// <li>vodeo: vodeo resource.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type RewriteAction struct {
	// Feature name. For details, see [DescribeRulesSetting](https://intl.cloud.tencent.com/document/product/1552/80618?from_cn_redirect=1) API
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Parameter
	Parameters []*RuleRewriteActionParams `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type Rule struct {
	// Feature execution conditions.
	// Note: If any condition in the array is met, the feature will run.
	Conditions []*RuleAndConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Feature to be executed.
	// Note: Actions and SubRules cannot both be empty.
	Actions []*Action `json:"Actions,omitnil,omitempty" name:"Actions"`

	// The nested rule.
	// Note: Actions and SubRules cannot both be empty.
	SubRules []*SubRuleItem `json:"SubRules,omitnil,omitempty" name:"SubRules"`
}

type RuleAndConditions struct {
	// Rule engine condition. This condition will be considered met if all items in the array are met.
	Conditions []*RuleCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type RuleChoicePropertiesItem struct {
	// The parameter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The parameter value type.
	// <li>CHOICE: The parameter value can be selected only from `Values`.</li>
	// <li>TOGGLE: The parameter value is of switch type and can be selected from `ChoicesValue`.</li>
	// <li>CUSTOM_NUM: The parameter value is a custom integer.</li>
	// <li>CUSTOM_STRING: The parameter value is a custom string.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Valid parameter values.
	// Note: If `Type` is `CUSTOM_NUM` or `CUSTOM_STRING`, this parameter will be an empty array.
	ChoicesValue []*string `json:"ChoicesValue,omitnil,omitempty" name:"ChoicesValue"`

	// Minimum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Maximum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// Whether multiple values can be selected or entered.
	IsMultiple *bool `json:"IsMultiple,omitnil,omitempty" name:"IsMultiple"`

	// Whether the parameter can be left empty.
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitnil,omitempty" name:"IsAllowEmpty"`

	// Special parameter.
	// <li>NULL: Select `NormalAction` for `RuleAction`. </li>
	// <li>If the member parameter `Id` is `Action`, select `RewirteAction` for `RuleAction`.</li>
	// <li>If the member parameter `Id` is `StatusCode`, select `CodeAction` for `RuleAction`.</li>
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitnil,omitempty" name:"ExtraParameter"`
}

type RuleCodeActionParams struct {
	// The status code.
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// The parameter name. For details, see [DescribeRulesSetting](https://intl.cloud.tencent.com/document/product/1552/80618?from_cn_redirect=1).
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The parameter value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type RuleCondition struct {
	// Operator. Valid values:
	// <li>`equals`: Equals</li>
	// <li>`notEquals`: Does not equal</li>
	// <li>`exist`: Exists</li>
	// <li>`notexist`: Does not exist</li>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Match type. Valid values: <li> filename: File name; </li> <li> extension: File extension; </li> <li> host: Host name; </li> <li> full_url: The complete URL path under the current site, which must include the HTTP protocol, host, and path; </li> <li> url: Request for the URL path under the current site; </li><li> client_country: Client country/region;</li> <li> query_string: The query string of the URL requested under the current site; </li> <li> request_header: HTTP request header; </li><li> client_ip: Client IP address. </li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// The parameter values for match types. It is allowed to pass an empty array only when the match type is query_string or request_header and the operator value is Exist or Does Not Exist. The corresponding match types include:
	// <li> File extension: Extensions like jpg, txt, etc.;</li>
	// <li> File name: For example, foo in foo.jpg;</li>
	// <li> All: All requests for domain names under the site; </li>
	// <li> HOST: The host under the current site, for example, www.maxx55.com;</li>
	// <li> URL Path: Request for the URL path under the current site, for example, /example;</li>
	// <li> URL Full: The complete URL request under the current site, which must include the HTTP protocol, host, and path, for example, https://www.maxx55.cn/example;</li>
	// <li> Client country/region: Country/region codes compliant with the ISO3166 standard;</li>
	// <li> Query string: The parameter values in the query string of the URL requested under the current site, for example, cn and 1 in lang=cn&version=1; </li>
	// <li> HTTP request header: The value of the HTTP request header field, for example, zh-CN,zh;q=0.9 in Accept-Language:zh-CN,zh;q=0.9; </li>
	// <li> Client IP: The client IP address carried by the current request, supporting IPv4, IPv6, and an IP range. </li>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Whether the parameter value is case insensitive. Default value: false.
	IgnoreCase *bool `json:"IgnoreCase,omitnil,omitempty" name:"IgnoreCase"`

	// The parameter name of the match type. This field is required only when `Target=query_string/request_header`.
	// <li>`query_string`: Name of the query string, such as "lang" and "version" in "lang=cn&version=1".</li>
	// <li>`request_header`: Name of the HTTP request header, such as "Accept-Language" in the "Accept-Language:zh-CN,zh;q=0.9" header.</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether the parameter name is case insensitive. Default value: `false`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: IgnoreNameCase is deprecated.
	IgnoreNameCase *bool `json:"IgnoreNameCase,omitnil,omitempty" name:"IgnoreNameCase"`
}

type RuleExtraParameter struct {
	// Parameter name. Valid values:
	// <li>`Action`: Required parameter for HTTP header modification when `RewirteAction` is selected for `RuleAction`.</li>
	// <li>`StatusCode`: Required parameter for the status code feature when `CodeAction` is selected for `RuleAction`.</li>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Parameter value type.
	// <li>`CHOICE`: The parameter value can be selected only from `Values`.</li>
	// <li>`CUSTOM_NUM`: The parameter value is a custom integer.</li>
	// <li>`CUSTOM_STRING`: The parameter value is a custom string.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Valid values.
	// Note: If the value of `Id` is `StatusCode`, values in the array are all integer values. When entering a parameter value, enter the integer value of the string.
	Choices []*string `json:"Choices,omitnil,omitempty" name:"Choices"`
}

type RuleItem struct {
	// The rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The rule name. It is a string that can contain 1–255 characters.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Rule status. Values:
	// <li>`enable`: Enabled</li>
	// <li>`disable`: Disabled</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The rule content.
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// The rule priority. The greater the value, the higher the priority. The minimum value is `1`.
	RulePriority *int64 `json:"RulePriority,omitnil,omitempty" name:"RulePriority"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RuleNormalActionParams struct {
	// The parameter name. For details, see [DescribeRulesSetting](https://intl.cloud.tencent.com/document/product/1552/80618?from_cn_redirect=1).
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The parameter value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type RuleRewriteActionParams struct {
	// Feature parameter name. For details, see [DescribeRulesSetting](https://intl.cloud.tencent.com/document/product/1552/80618?from_cn_redirect=1).
	// <li>`add`: Add the HTTP header.</li>
	// <li>`set`: Rewrite the HTTP header.</li>
	// <li>`del`: Delete the HTTP header.</li>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type RulesProperties struct {
	// Parameter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Minimum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Valid parameter values.
	// Note: If `Type` is `CUSTOM_NUM` or `CUSTOM_STRING`, this parameter will be an empty array.
	ChoicesValue []*string `json:"ChoicesValue,omitnil,omitempty" name:"ChoicesValue"`

	// The parameter value type.
	// <li>`CHOICE`: `If Type=CHOICE`, choose a value in `ChoiceValue`.</li>
	// <li>`TOGGLE`: If `Type=TOGGLE`, choose `on` or `off` from `ChoicesValue`.</li>
	// <li>`OBJECT`: Specify an object. If this is specified, `ChoiceProperties` includes attributes of the specified object. See [Example 2. Create a rule with Type=OBJECT](https://intl.cloud.tencent.com/document/product/1552/80622?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B2-.E5.8F.82.E6.95.B0.E4.B8.BA-OBJECT-.E7.B1.BB.E5.9E.8B.E7.9A.84.E5.88.9B.E5.BB.BA)</li>
	// <li>`CUSTOM_NUM`: (Integer) Custom value.</li>
	// <li>`CUSTOM_STRING`: (String) Custom value.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Maximum value. If both `Min` and `Max` are set to `0`, this parameter does not take effect.
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// Whether multiple values can be selected or entered.
	IsMultiple *bool `json:"IsMultiple,omitnil,omitempty" name:"IsMultiple"`

	// Whether the parameter can be left empty.
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitnil,omitempty" name:"IsAllowEmpty"`

	// Associated configuration parameters of this parameter, which are required for API call.
	// Note: This parameter will be an empty array if no special parameters are added as optional parameters.
	ChoiceProperties []*RuleChoicePropertiesItem `json:"ChoiceProperties,omitnil,omitempty" name:"ChoiceProperties"`

	// <li>NULL: No special parameters when `NormalAction` is selected for `RuleAction`.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitnil,omitempty" name:"ExtraParameter"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Parameter information
	Properties []*RulesProperties `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type S3 struct {
	// The URL without bucket name or path, for example: `https://storage.googleapis.com`, `https://s3.ap-northeast-2.amazonaws.com`, and `https://cos.ap-nanjing.myqcloud.com`.
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// The region where the bucket is located, for example: `ap-northeast-2`.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The bucket name and log storage directory, for example: `your_bucket_name/EO-logs/`. If the directory does not exist in the bucket, it will be created automatically.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// The Access Key ID used to access the bucket.
	AccessId *string `json:"AccessId,omitnil,omitempty" name:"AccessId"`

	// The secret key used to access the bucket.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// The data compress type. Valid values:<li>gzip: gzip compression.</li>If this field is not filled in, compression is disabled.
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

type SecEntry struct {
	// The query dimension value.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The details.
	Value []*SecEntryValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type SecEntryValue struct {
	// The metric name.
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// The time-series data details.
	Detail []*TimingDataItem `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The maximum value.
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// The average value.
	Avg *float64 `json:"Avg,omitnil,omitempty" name:"Avg"`

	// Sum
	Sum *float64 `json:"Sum,omitnil,omitempty" name:"Sum"`
}

type SecurityConfig struct {
	// Managed rule. If the parameter is null or not filled, the configuration last set will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	WafConfig *WafConfig `json:"WafConfig,omitnil,omitempty" name:"WafConfig"`

	// Rate limiting. If the parameter is null or not filled, the configuration last set will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// Custom rule. If the parameter is null or not filled, the configuration last set will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AclConfig *AclConfig `json:"AclConfig,omitnil,omitempty" name:"AclConfig"`

	// Bot configuration. If the parameter is null or not filled, the configuration last set will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	BotConfig *BotConfig `json:"BotConfig,omitnil,omitempty" name:"BotConfig"`

	// Switch setting of the 7-layer protection. If the parameter is null or not filled, the configuration last set will be used by default.Note: This field may return null, indicating that no valid value can be obtained.
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitnil,omitempty" name:"SwitchConfig"`

	// Basic access control. If the parameter is null or not filled, the configuration last set will be used by default.
	// Note: This field may return null, indicating that no valid value can be obtained.
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitnil,omitempty" name:"IpTableConfig"`

	// Exception rule configuration. If the parameter is null or not filled, the configuration last set will be used by default.Note: This field may return null, indicating that no valid value can be obtained.
	ExceptConfig *ExceptConfig `json:"ExceptConfig,omitnil,omitempty" name:"ExceptConfig"`

	// Custom block page settings. If the parameter is null or not filled, the configuration last set will be used by default.Note: This field may return null, indicating that no valid value can be obtained.
	DropPageConfig *DropPageConfig `json:"DropPageConfig,omitnil,omitempty" name:"DropPageConfig"`

	// Security template settings
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TemplateConfig *TemplateConfig `json:"TemplateConfig,omitnil,omitempty" name:"TemplateConfig"`

	// Settings for slow attack defense. If the parameter is null or not filled, the configuration last set will be used by default.Note: This field may return null, indicating that no valid value can be obtained.
	SlowPostConfig *SlowPostConfig `json:"SlowPostConfig,omitnil,omitempty" name:"SlowPostConfig"`
}

type SecurityTemplateBinding struct {
	// template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Binding status of the template.
	TemplateScope []*TemplateScope `json:"TemplateScope,omitnil,omitempty" name:"TemplateScope"`
}

type SecurityType struct {
	// Whether to enable the security type setting. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ServerCertInfo struct {
	// ID of the server certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Alias of the certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Type of the certificate. Values:
	// u200c<li>`default`: Default certificate</li>
	// u200c<li>`upload`: Custom certificate</li>
	// u200c<li>`managed`: Tencent Cloud-managed certificate</li>
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Time when the certificate expires.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Time when the certificate is deployed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeployTime *string `json:"DeployTime,omitnil,omitempty" name:"DeployTime"`

	// Signature algorithm.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SignAlgo *string `json:"SignAlgo,omitnil,omitempty" name:"SignAlgo"`

	// Domain name of the certificate.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`
}

type SkipCondition struct {
	// The field type. Values:
	// <li>`header_fields`: HTTP request header</li>
	// <li>`cookie`: HTTP request cookie</li>
	// <li>`query_string`: Query string in the HTTP request URL</li>
	// <li>`uri`: HTTP request URI</li>
	// <li>`body_raw`: HTTP request body</li>
	// <li>`body_json`: JSON HTTP request body</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The specific field. Values:
	// <li>`args`: Query parameter in the URI, such as "?name1=jack&age=12"</li>
	// <li>`path`: Partial path in the URI, such as "/path/to/resource.jpg"</li>
	// <li>`full`: Full path in the URI, such as "example.com/path/to/resource.jpg?name1=jack&age=12"</li>
	// <li>`upload_filename`: File path segment</li>
	// <li>`keys`: All keys</li>
	// <li>`values`: Values of all keys</li>
	// <li>`key_value`: Key and its value</li>
	Selector *string `json:"Selector,omitnil,omitempty" name:"Selector"`

	// The match method used to match the key. Values:
	// <li>`equal`: Exact match</li>
	// <li>`wildcard`: Wildcard match (only asterisks)</li>
	MatchFromType *string `json:"MatchFromType,omitnil,omitempty" name:"MatchFromType"`

	// The value that matches the key.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MatchFrom []*string `json:"MatchFrom,omitnil,omitempty" name:"MatchFrom"`

	// The match method used to match the content.
	// <li>`equal`: Exact match</li>
	// <li>`wildcard`: Wildcard match (only asterisks)</li>
	MatchContentType *string `json:"MatchContentType,omitnil,omitempty" name:"MatchContentType"`

	// The value that matches the content.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MatchContent []*string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`
}

type SlowPostConfig struct {
	// Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Detect slow attacks by the transfer period of the first 8 KB of requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FirstPartConfig *FirstPartConfig `json:"FirstPartConfig,omitnil,omitempty" name:"FirstPartConfig"`

	// Detect slow attacks by the data rate of the main body (excluding the first 8 KB) of requests
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SlowRateConfig *SlowRateConfig `json:"SlowRateConfig,omitnil,omitempty" name:"SlowRateConfig"`

	// The action to taken when a slow attack is detected. Values:
	// <li>`monitor`: Observe</li>
	// <li>`drop`: Block the request</li>
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// ID of the rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type SlowRateConfig struct {
	// Switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The sampling interval in seconds. In this way, the first 8 KB of the request is ignored. The rest of data is separated in to multiple parts according to this interval for slow attack measurement.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// The transfer rate threshold in bps. When the transfer rate of a sample is lower than the threshold, it’s considered a slow attack and handled according to the specified `Action`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type SmartRouting struct {
	// Whether to enable smart acceleration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type StandardDebug struct {
	// Whether to enable standard debugging. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable </li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The client IP to allow. It can be an IPv4/IPv6 address or a CIDR block. If not specified, it means to allow any client IP
	AllowClientIPList []*string `json:"AllowClientIPList,omitnil,omitempty" name:"AllowClientIPList"`

	// The time when the standard debugging setting expires. If it is exceeded, this feature becomes invalid.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type SubRule struct {
	// The condition that determines if a feature should run.
	// Note: If any condition in the array is met, the feature will run.
	Conditions []*RuleAndConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// The feature to be executed.
	Actions []*Action `json:"Actions,omitnil,omitempty" name:"Actions"`
}

type SubRuleItem struct {
	// Nested rule settings
	Rules []*SubRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Tag of the rule.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type Sv struct {
	// The parameter key.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The parameter value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Quota for a resource. Values:
	// <li>`zone`: Quota for sites</li>
	// <li>`custom-rule`: Quota for custom rules</li>
	// <li>`rate-limiting-rule`: Quota for rate limiting rules</li>
	// <li>`l4-proxy-instance`: Quota for L4 proxy instances </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pack *string `json:"Pack,omitnil,omitempty" name:"Pack"`

	// ID of the L4 proxy instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The protection specification.
	// Values: <li> `cm_30G`: 30 Gbps base protection bandwidth in **Chinese mainland** service area</li><li> `cm_60G`: 60 Gbps base protection bandwidth in **Chinese mainland** service area</li><li> `cm_100G`: 100 Gbps base protection bandwidth in **Chinese mainland** service area</li><li> `anycast_300G`: 300 Gbps Anycast-based protection in **Global (MLC)** service area</li><li> `anycast_unlimited`: Unlimited Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> `cm_30G_anycast_300G`: 30 Gbps base protection bandwidth in **Chinese mainland** service area and 300 Gbps Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> `cm_30G_anycast_unlimited`: 30 Gbps base protection bandwidth in **Chinese mainland** service area and unlimited Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> cm_60G_anycast_300G`: 60 Gbps base protection bandwidth in **Chinese mainland** service area and 300 Gbps Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> cm_60G_anycast_unlimited`: 60 Gbps base protection bandwidth in **Chinese mainland** service area and unlimited Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> `cm_100G_anycast_300G`: 100 Gbps base protection bandwidth in **Chinese mainland** service area and 300 Gbps Anycast-based protection bandwidth in **Global (MLC)** service area</li><li> cm_100G_anycast_unlimited`: 100 Gbps base protection bandwidth in **Chinese mainland** service area and unlimited Anycast-based protection bandwidth in **Global (MLC)** service area </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProtectionSpecs *string `json:"ProtectionSpecs,omitnil,omitempty" name:"ProtectionSpecs"`
}

type SwitchConfig struct {
	// Whether to enable web protection. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>It does not affect DDoS and bot configuration.
	WebSwitch *string `json:"WebSwitch,omitnil,omitempty" name:"WebSwitch"`
}

type Tag struct {
	// The tag key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// The tag value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Task struct {
	// ID of the task.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Resource.
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// Type of the task.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Node cache purge method, with values:
	// <li>invalidate: Marks as expired. A back-to-origin validation is triggered upon user request, sending an HTTP conditional request with If-None-Match and If-Modified-Since headers. If the origin server responds with 200, the node will fetch new resources from the origin and update the cache; if the origin server responds with 304, the cache will not be updated;</li>
	// <li>delete: Directly deletes the node's cache, triggering a resource fetch from the origin upon user request.</li>
	// Note: This field may return null, which indicates a failure to obtain a valid value.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// Status. Valid values:
	// <li>processing: Processing;</li>
	// <li>success: Succeeded;</li>
	// <li>failed: Failed;</li>
	// <li>timeout: Timed out. </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Creation time of the task.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Completion time of the task.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TemplateConfig struct {
	// The template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// The template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

type TemplateScope struct {
	// ID of the site.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// List of instance statuses
	// Note: This field may return·null, indicating that no valid values can be obtained.
	EntityStatus []*EntityStatus `json:"EntityStatus,omitnil,omitempty" name:"EntityStatus"`
}

type TimingDataItem struct {
	// The query time recorded in seconds using UNIX timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// The value.
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TimingDataRecord struct {
	// The query dimension value.
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`

	// Detailed time series data
	TypeValue []*TimingTypeValue `json:"TypeValue,omitnil,omitempty" name:"TypeValue"`
}

type TimingTypeValue struct {
	// Sum.
	Sum *int64 `json:"Sum,omitnil,omitempty" name:"Sum"`

	// The maximum value.
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// The average value.
	Avg *int64 `json:"Avg,omitnil,omitempty" name:"Avg"`

	// Metric name.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*TimingDataItem `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type TopDataRecord struct {
	// The query dimension value.
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`

	// Top data rankings
	DetailData []*TopDetailData `json:"DetailData,omitnil,omitempty" name:"DetailData"`
}

type TopDetailData struct {
	// The field name.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The field value.
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TopEntry struct {
	// The query dimension value.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The details.
	Value []*TopEntryValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type TopEntryValue struct {
	// The item name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The number of items.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type UpgradePlanRequestParams struct {
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Target plan version for upgrade. Valid values: <li>basic: Basic Edition Plan;</li><li>standard: Standard Edition Plan. </li>
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// Whether to automatically use a voucher. Valid values: <li>true: Yes;</li><li>false: No. </li> If this field is not specified, the default value 'false' will be used.
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

type UpgradePlanRequest struct {
	*tchttp.BaseRequest
	
	// Plan ID, formatted as edgeone-2unuvzjmmn2q.
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Target plan version for upgrade. Valid values: <li>basic: Basic Edition Plan;</li><li>standard: Standard Edition Plan. </li>
	PlanType *string `json:"PlanType,omitnil,omitempty" name:"PlanType"`

	// Whether to automatically use a voucher. Valid values: <li>true: Yes;</li><li>false: No. </li> If this field is not specified, the default value 'false' will be used.
	AutoUseVoucher *string `json:"AutoUseVoucher,omitnil,omitempty" name:"AutoUseVoucher"`
}

func (r *UpgradePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "PlanType")
	delete(f, "AutoUseVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradePlanResponseParams struct {
	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradePlanResponse struct {
	*tchttp.BaseResponse
	Response *UpgradePlanResponseParams `json:"Response"`
}

func (r *UpgradePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpstreamHttp2 struct {
	// Whether to enable HTTP2 origin-pull. Valid values: 
	// <li>`on`: Enable;</li>
	// <li>`off`: Disable.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type VanityNameServers struct {
	// Whether to enable custom name servers. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// List of custom name servers
	Servers []*string `json:"Servers,omitnil,omitempty" name:"Servers"`
}

type VanityNameServersIps struct {
	// Custom name of the name server
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// IPv4 address of the custom name server
	IPv4 *string `json:"IPv4,omitnil,omitempty" name:"IPv4"`
}

// Predefined struct for user
type VerifyOwnershipRequestParams struct {
	// Site or acceleration domain name
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type VerifyOwnershipRequest struct {
	*tchttp.BaseRequest
	
	// Site or acceleration domain name
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *VerifyOwnershipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyOwnershipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyOwnershipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyOwnershipResponseParams struct {
	// Result of ownership verification
	// <li>`success`: Verification passed</li>
	// <li>`fail`: Verification failed</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// When the ownership verification result is `fail`, this field returns the reason of failure.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyOwnershipResponse struct {
	*tchttp.BaseResponse
	Response *VerifyOwnershipResponseParams `json:"Response"`
}

func (r *VerifyOwnershipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyOwnershipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Waf struct {
	// Whether to enable WAF. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// ID of the policy
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type WafConfig struct {
	// Whether to enable WAF configuration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>The configuration can be modified even when it is disabled.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The protection level. Values:
	// <li>`loose`: Loose</li>
	// <li>`normal`: Moderate</li>
	// <li>`strict`: Strict</li>
	// <li>`stricter`: Super strict</li>
	// <li>`custom`: Custom</li>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// The WAF global mode. Values:
	// <li>`block`: Block globally</li>
	// <li>`observe`: Observe globally</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// The settings of the managed rule. If it is null, the settings that were last configured will be used.
	WafRule *WafRule `json:"WafRule,omitnil,omitempty" name:"WafRule"`

	// The setting of the AI rule engine. If it is null, the setting that was last configured will be used.
	AiRule *AiRule `json:"AiRule,omitnil,omitempty" name:"AiRule"`
}

type WafRule struct {
	// Whether to enable managed rules. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// IDs of the rules to be disabled.
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitnil,omitempty" name:"BlockRuleIDs"`

	// IDs of the rules to be executed in Observe mode.
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitnil,omitempty" name:"ObserveRuleIDs"`
}

type WebSocket struct {
	// Whether to enable WebSocket connection timeout. Values:
	// <li>`on`: The field "Timeout" can be configured.</li>
	// <li>`off`: The field "Timeout" is fixed to 15 seconds.</li>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// The timeout period in seconds. Maximum value: 120.
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type Zone struct {
	// Site ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// List of name servers used by the site
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil,omitempty" name:"OriginalNameServers"`

	// The list of name servers assigned by Tencent Cloud.
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`

	// The site status. Values:
	// u200c<li>`active`: The name server is switched to EdgeOne.</li>
	// u200c<li>`pending`: The name server is not switched.</li>
	// u200c<li>`moved`: The name server is changed to other service providers.</li>
	// u200c<li>`deactivated`: The site is blocked.</li>
	// <li>`initializing`: The site is not bound with any plan. </li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Site connection method, valid values:
	// <li>full: Connect via NS;</li>
	// <li>partial: Connect via CNAME;</li>
	// <li>noDomainAccess: Connect without using a domain name;</li>
	// <li>vodeo: Connect by default when VODEO is enabled.</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether the site is disabled.
	Paused *bool `json:"Paused,omitnil,omitempty" name:"Paused"`

	// Whether CNAME acceleration is enabled. Values:
	// <li>`enabled`: Enabled</li>
	// <li>`disabled`: Disabled</li>
	CnameSpeedUp *string `json:"CnameSpeedUp,omitnil,omitempty" name:"CnameSpeedUp"`

	// CNAME record access status. Values:
	// <li>`finished`: The site is verified.</li>
	// <li>`pending`: The site is being verified.</li>
	CnameStatus *string `json:"CnameStatus,omitnil,omitempty" name:"CnameStatus"`

	// The list of resource tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The list of billable resources.
	Resources []*Resource `json:"Resources,omitnil,omitempty" name:"Resources"`

	// The creation time of the site.
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// The modification date of the site.
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// The site access region. Values:
	// <li>`global`: Global.</li>
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Outside the Chinese mainland.</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// The custom name server information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil,omitempty" name:"VanityNameServers"`

	// The custom name server IP information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitnil,omitempty" name:"VanityNameServersIps"`

	// Status of the proxy. Values:
	// <li>`active`: Enabled</li>
	// <li>`inactive`: Not activated</li>
	// <li>`paused`: Disabled</li>
	ActiveStatus *string `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// The site alias. It can be up to 20 characters consisting of digits, letters, hyphens (-) and underscores (_).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AliasZoneName *string `json:"AliasZoneName,omitnil,omitempty" name:"AliasZoneName"`

	// Whether it’s a fake site. Valid values: 
	// <li>`0`: Non-fake site;</li>
	// <li>`1`: Fake site.</li>
	IsFake *int64 `json:"IsFake,omitnil,omitempty" name:"IsFake"`

	// Lock status. Values: <li>`enable`: Normal. Modification is allowed.</li><li>`disable`: Locked. Modification is not allowed.</li><li>`plan_migrate`: Adjusting the plan. Modification is not allowed.</li> 
	LockStatus *string `json:"LockStatus,omitnil,omitempty" name:"LockStatus"`

	// Ownership verification information
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`
}

type ZoneSetting struct {
	// Name of the site
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Site acceleration region. Values:
	// <li>`mainland`: Acceleration in the Chinese mainland.</li>
	// <li>`overseas`: Acceleration outside the Chinese mainland.</li>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Node cache key configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitnil,omitempty" name:"CacheKey"`

	// The QUIC access configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Quic *Quic `json:"Quic,omitnil,omitempty" name:"Quic"`

	// The POST transport configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil,omitempty" name:"PostMaxSize"`

	// Smart compression configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Compression *Compression `json:"Compression,omitnil,omitempty" name:"Compression"`

	// HTTP2 origin-pull configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil,omitempty" name:"UpstreamHttp2"`

	// Force HTTPS redirect configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// Cache expiration time configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil,omitempty" name:"CacheConfig"`

	// Origin server configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origin *Origin `json:"Origin,omitnil,omitempty" name:"Origin"`

	// Smart acceleration configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil,omitempty" name:"SmartRouting"`

	// Browser cache configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// The offline cache configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil,omitempty" name:"OfflineCache"`

	// WebSocket configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WebSocket *WebSocket `json:"WebSocket,omitnil,omitempty" name:"WebSocket"`

	// Origin-pull client IP header configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil,omitempty" name:"ClientIpHeader"`

	// Cache prefresh configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil,omitempty" name:"CachePrefresh"`

	// IPv6 access configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ipv6 *Ipv6 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// HTTPS acceleration configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitnil,omitempty" name:"Https"`

	// Whether to carry the location information of the client IP during origin-pull.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil,omitempty" name:"ClientIpCountry"`

	// Configuration of gRPC support
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Grpc *Grpc `json:"Grpc,omitnil,omitempty" name:"Grpc"`

	// Image optimization configuration. 
	// Note: This field may return `null`, indicating that no valid value was found.
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil,omitempty" name:"ImageOptimize"`

	// Cross-MLC-border acceleration. 
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil,omitempty" name:"AccelerateMainland"`

	// Standard debugging configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil,omitempty" name:"StandardDebug"`
}