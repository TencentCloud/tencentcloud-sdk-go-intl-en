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

type Action struct {
	// Common feature operation. Features of this type include:
	// <li>`AccessUrlRedirect`: Access URL rewrite.</li>
	// <li>`UpstreamUrlRedirect`: Origin-pull URL rewrite.</li>
	// <li>`QUIC`: QUIC.</li>
	// <li>`WebSocket`: WebSocket.</li>
	// <li>`VideoSeek`: Video dragging.</li>
	// <li>`Authentication`: Token authentication.</li>
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
	// <li>`Hsts`.</li>
	// <li>`ClientIpHeader`.</li>
	// <li>`TlsVersion`.</li>
	// <li>`OcspStapling`.</li>
	// <li>`Http2`: HTTP/2 access.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	// The name of the field to filter.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Values of the filtered field.
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Whether to enable fuzzy query.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type AscriptionInfo struct {

	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// The record type.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// The record value.
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
}

type BillingDataFilter struct {
	// Parameter name. Valid values:
	// `zone`: Site name
	// `host`: Domain name
	// `proxy`: L4 proxy
	// `plan`: Plan type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type BotLog struct {
	// The attack time recorded in seconds using UNIX timestamp.
	AttackTime *uint64 `json:"AttackTime,omitempty" name:"AttackTime"`

	// The attacker IP.
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// The attacked domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The URI.
	RequestUri *string `json:"RequestUri,omitempty" name:"RequestUri"`

	// Attack type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Request method.
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// The attack content.
	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`

	// The attack level.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// The rule ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// The country code of the attacker IP, which is defined in ISO-3166 alpha-2. For the list of country codes, see [ISO-3166](https://git.woa.com/edgeone/iso-3166/blob/master/all/all.json).
	SipCountryCode *string `json:"SipCountryCode,omitempty" name:"SipCountryCode"`

	// The attack event ID.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// The processing method.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DisposalMethod *string `json:"DisposalMethod,omitempty" name:"DisposalMethod"`

	// The HTTP log.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpLog *string `json:"HttpLog,omitempty" name:"HttpLog"`

	// The user agent.
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// The detection method.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetectionMethod *string `json:"DetectionMethod,omitempty" name:"DetectionMethod"`

	// The credibility level.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Confidence *string `json:"Confidence,omitempty" name:"Confidence"`

	// Maliciousness
	// Note: This field may return null, indicating that no valid values can be obtained.
	Maliciousness *string `json:"Maliciousness,omitempty" name:"Maliciousness"`

	// The security rule information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleDetailList []*SecRuleRelatedInfo `json:"RuleDetailList,omitempty" name:"RuleDetailList"`

	// The bot tag.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Label *string `json:"Label,omitempty" name:"Label"`
}

type CC struct {
	// WAF switch. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// ID of the policy
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type CCInterceptEvent struct {
	// The client IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// The requests per minute that are blocked.
	InterceptNum *int64 `json:"InterceptNum,omitempty" name:"InterceptNum"`

	// Block time in seconds.
	InterceptTime *int64 `json:"InterceptTime,omitempty" name:"InterceptTime"`
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

	Switch *string `json:"Switch,omitempty" name:"Switch"`


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
type CreateDnsRecordRequestParams struct {
	// The site ID of the DNS record.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The DNS record type. Values:
	// <li>`A`: Point a domain name to an IPv4 address, such as 8.8.8.8.</li>
	// <li>`AAAA`: Point a domain name to an IPv6 address.</li>
	// <li>`MX`: It is used for email servers. The record value and priority parameters are provided by email service providers. If there are multiple MX records, the lower the priority value, the higher the priority.</li>
	// <li>`CNAME`: Point a domain name to another domain name that can be resolved to an IP address.</li>
	// <li>`TXT`: Identify and describe a domain name. It is usually used for domain verification and as SPF records (for anti-spam).</li>
	// <li>`NS`: If you need to authorize a subdomain name to another DNS service provider for DNS resolution, you need to add an NS record. You cannot add an NS record for a root domain name.</li>
	// <li>`CAA`: Specify CAs to issue certificates for sites.</li>
	// <li>`SRV`: Identify a service used by a server. It is commonly used in Microsoft directory management.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// The DNS record name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The DNS record content.
	Content *string `json:"Content,omitempty" name:"Content"`

	// The proxy mode. Values:
	// <li>`dns_only`: Only DNS</li>
	// <li>`proxied`: Proxied</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// TTL (in seconds). The smaller the value, the faster the record changes take effect. Default value: 300
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// Specifies a value in the range 1–50 when you make changes to the MX records. A smaller value indicates higher priority. Note that the default value 0 will be used if this field is not specified.
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type CreateDnsRecordRequest struct {
	*tchttp.BaseRequest
	
	// The site ID of the DNS record.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The DNS record type. Values:
	// <li>`A`: Point a domain name to an IPv4 address, such as 8.8.8.8.</li>
	// <li>`AAAA`: Point a domain name to an IPv6 address.</li>
	// <li>`MX`: It is used for email servers. The record value and priority parameters are provided by email service providers. If there are multiple MX records, the lower the priority value, the higher the priority.</li>
	// <li>`CNAME`: Point a domain name to another domain name that can be resolved to an IP address.</li>
	// <li>`TXT`: Identify and describe a domain name. It is usually used for domain verification and as SPF records (for anti-spam).</li>
	// <li>`NS`: If you need to authorize a subdomain name to another DNS service provider for DNS resolution, you need to add an NS record. You cannot add an NS record for a root domain name.</li>
	// <li>`CAA`: Specify CAs to issue certificates for sites.</li>
	// <li>`SRV`: Identify a service used by a server. It is commonly used in Microsoft directory management.</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// The DNS record name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The DNS record content.
	Content *string `json:"Content,omitempty" name:"Content"`

	// The proxy mode. Values:
	// <li>`dns_only`: Only DNS</li>
	// <li>`proxied`: Proxied</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// TTL (in seconds). The smaller the value, the faster the record changes take effect. Default value: 300
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// Specifies a value in the range 1–50 when you make changes to the MX records. A smaller value indicates higher priority. Note that the default value 0 will be used if this field is not specified.
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
	delete(f, "TTL")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDnsRecordResponseParams struct {
	// The DNS record ID.
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

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
type CreateLogSetRequestParams struct {
	// Name of the logset.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

type CreateLogSetRequest struct {
	*tchttp.BaseRequest
	
	// Name of the logset.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

func (r *CreateLogSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogSetName")
	delete(f, "LogSetRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogSetResponseParams struct {
	// ID of the logset created.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLogSetResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogSetResponseParams `json:"Response"`
}

func (r *CreateLogSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogTopicTaskRequestParams struct {
	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// Topic name of the logset.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Name of the shipping task.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Name of the site.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Type of the shipping entity. Values:
	// <li>`domain`: L7 acceleration logs;</li>
	// <li>`application`: L4 acceleration logs;</li>
	// <li>`web-rateLiming`: Rate limiting logs;</li>
	// <li>`web-attack`: Web security logs;</li>
	// <li>`web-rule`: Custom rule logs;</li>
	// <li>`web-bot`: Bot management logs.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// Retention period of the log topic. Value range: 1–366 (in days).
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// List of shipping entities.
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Global (outside the Chinese mainland).</li> If this field is not specified, the acceleration region will be determined based on the user’s region.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CreateLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// Topic name of the logset.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Name of the shipping task.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Name of the site.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Type of the shipping entity. Values:
	// <li>`domain`: L7 acceleration logs;</li>
	// <li>`application`: L4 acceleration logs;</li>
	// <li>`web-rateLiming`: Rate limiting logs;</li>
	// <li>`web-attack`: Web security logs;</li>
	// <li>`web-rule`: Custom rule logs;</li>
	// <li>`web-bot`: Bot management logs.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// Retention period of the log topic. Value range: 1–366 (in days).
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// List of shipping entities.
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland.</li>
	// <li>`overseas`: Global (outside the Chinese mainland).</li> If this field is not specified, the acceleration region will be determined based on the user’s region.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogSetId")
	delete(f, "LogSetRegion")
	delete(f, "TopicName")
	delete(f, "TaskName")
	delete(f, "ZoneId")
	delete(f, "ZoneName")
	delete(f, "EntityType")
	delete(f, "Period")
	delete(f, "EntityList")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogTopicTaskResponseParams struct {
	// ID of the log topic created.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogTopicTaskResponseParams `json:"Response"`
}

func (r *CreateLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogTopicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanForZoneRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Plan options available for purchase. Values:
	// <li>`sta`: Standard plan for global areas except Chinese mainland;</li>
	// <li>`sta_with_bot`: Standard plan for global areas except Chinese mainland, with extra bot management;</li>
	// <li>`sta_cm`: Standard plan for Chinese mainland;</li>
	// <li>`sta_cm_with_bot`: Standard plan for Chinese mainland, with extra bot management;</li>
	// <li>`ent`: Enterprise plan for global areas except Chinese mainland;</li>
	// <li>`ent_with_bot`: Enterprise plan for global areas except Chinese mainland, with extra bot management;</li>
	// <li>`ent_cm`: Enterprise plan for Chinese mainland;</li>
	// <li>`ent_cm_with_bot`: Enterprise plan for Chinese mainland, with extra bot management.</li>To get the available plan options for your account, view the output from <a href="https://tcloud4api.woa.com/document/product/1657/80124?!preview&!document=1">DescribeAvailablePlans</a>.
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
}

type CreatePlanForZoneRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Plan options available for purchase. Values:
	// <li>`sta`: Standard plan for global areas except Chinese mainland;</li>
	// <li>`sta_with_bot`: Standard plan for global areas except Chinese mainland, with extra bot management;</li>
	// <li>`sta_cm`: Standard plan for Chinese mainland;</li>
	// <li>`sta_cm_with_bot`: Standard plan for Chinese mainland, with extra bot management;</li>
	// <li>`ent`: Enterprise plan for global areas except Chinese mainland;</li>
	// <li>`ent_with_bot`: Enterprise plan for global areas except Chinese mainland, with extra bot management;</li>
	// <li>`ent_cm`: Enterprise plan for Chinese mainland;</li>
	// <li>`ent_cm_with_bot`: Enterprise plan for Chinese mainland, with extra bot management.</li>To get the available plan options for your account, view the output from <a href="https://tcloud4api.woa.com/document/product/1657/80124?!preview&!document=1">DescribeAvailablePlans</a>.
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

	// Purging mode. Values:
	// <li>`purge_url`: Purge URLs;</li>
	// <li>`purge_prefix`: Purge prefixes;</li>
	// <li>`purge_host`: Purge hostnames;</li>
	// <li>`purge_all`: Purge all caches.</li>
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
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Purging mode. Values:
	// <li>`purge_url`: Purge URLs;</li>
	// <li>`purge_prefix`: Purge prefixes;</li>
	// <li>`purge_host`: Purge hostnames;</li>
	// <li>`purge_all`: Purge all caches.</li>
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
type CreateZoneRequestParams struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The access mode. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record.</li>This field will be set to the default value `full` if not specified.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Whether to skip scanning the existing DNS records of the site. Default value: false.
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`

	// The resource tag.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest
	
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The access mode. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record.</li>This field will be set to the default value `full` if not specified.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Whether to skip scanning the existing DNS records of the site. Default value: false.
	JumpStart *bool `json:"JumpStart,omitempty" name:"JumpStart"`

	// The resource tag.
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
	delete(f, "ZoneName")
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
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// ID of the site.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DDoSAttackEventDetailData struct {
	// The attack status. Values:
	// <li>`1`: The attack is being observed;</li>
	// <li>`2`: The attack started;</li>
	// <li>`3`: The attack ended.</li>
	AttackStatus *int64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// The attack type.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// The end time.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// The start time.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// The maximum bandwidth.
	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`

	// The maximum packet rate.
	PacketMaxRate *int64 `json:"PacketMaxRate,omitempty" name:"PacketMaxRate"`

	// The event ID.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// The DDoS policy ID.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DDoSAttackSourceEvent struct {
	// The attacker IP.
	AttackSourceIp *string `json:"AttackSourceIp,omitempty" name:"AttackSourceIp"`

	// The country or region.
	AttackRegion *string `json:"AttackRegion,omitempty" name:"AttackRegion"`

	// The accumulative attack traffic.
	AttackFlow *uint64 `json:"AttackFlow,omitempty" name:"AttackFlow"`

	// The accumulative attack packets.
	AttackPacketNum *uint64 `json:"AttackPacketNum,omitempty" name:"AttackPacketNum"`
}

type DDoSBlockData struct {
	// The start time recorded in UNIX timestamp.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// The start time recorded in UNIX timestamp.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DDoSMajorAttackEvent struct {
	// The DDoS policy ID.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The maximum attack bandwidth.
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitempty" name:"AttackMaxBandWidth"`

	// The attack time recorded in seconds using UNIX timestamp.
	AttackTime *int64 `json:"AttackTime,omitempty" name:"AttackTime"`
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
	// <li>`processing`: Deployment in progress;</li>
	// <li>`deployed`: Deployed.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Failure description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`


	SignAlgo *string `json:"SignAlgo,omitempty" name:"SignAlgo"`
}

// Predefined struct for user
type DeleteDnsRecordsRequestParams struct {
	// The site ID of the DNS record to be deleted.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The ID of the DNS record to be deleted.
	DnsRecordIds []*string `json:"DnsRecordIds,omitempty" name:"DnsRecordIds"`
}

type DeleteDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// The site ID of the DNS record to be deleted.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The ID of the DNS record to be deleted.
	DnsRecordIds []*string `json:"DnsRecordIds,omitempty" name:"DnsRecordIds"`
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
	delete(f, "DnsRecordIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDnsRecordsResponseParams struct {
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
type DeleteLogTopicTaskRequestParams struct {
	// ID of the shipping task to be deleted.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

type DeleteLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the shipping task to be deleted.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`
}

func (r *DeleteLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "LogSetRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogTopicTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogTopicTaskResponseParams `json:"Response"`
}

func (r *DeleteLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogTopicTaskResponse) FromJsonString(s string) error {
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
type DescribeBillingDataRequestParams struct {
	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Time granularity. Values:
	// <ul>
	// <li>`min`: One minute</li>
	// <li>`5min`: Five minutes</li>
	// <li>`hour`: One hour</li>
	// <li>`day`: One day</li>
	// </ul>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Metric item. Values:
	// <ul>
	// <li>`acc_flux`: Content acceleration traffic;</li>
	// <li>`quic_request`: QUIC requests;</li>
	// <li>`sec_flux`: Security traffic;</li>
	// <li>`sec_request_clean`: Clean security requests.</li>
	// </ul>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Filter item. Values:
	// <ul>
	// <li>`zone`: Site;</li>
	// <li>`plan`: Service plan;</li>
	// <li>`service`: L4 or L7;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	// </ul>
	Filters []*BillingDataFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Time granularity. Values:
	// <ul>
	// <li>`min`: One minute</li>
	// <li>`5min`: Five minutes</li>
	// <li>`hour`: One hour</li>
	// <li>`day`: One day</li>
	// </ul>
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Metric item. Values:
	// <ul>
	// <li>`acc_flux`: Content acceleration traffic;</li>
	// <li>`quic_request`: QUIC requests;</li>
	// <li>`sec_flux`: Security traffic;</li>
	// <li>`sec_request_clean`: Clean security requests.</li>
	// </ul>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Filter item. Values:
	// <ul>
	// <li>`zone`: Site;</li>
	// <li>`plan`: Service plan;</li>
	// <li>`service`: L4 or L7;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	// </ul>
	Filters []*BillingDataFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Interval")
	delete(f, "MetricName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataResponseParams struct {
	// Data of the sampling point
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DnsData `json:"Data,omitempty" name:"Data"`

	// Time granularity of sampling
	// Note: This field may return null, indicating that no valid values can be obtained.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBotClientIpListRequestParams struct {
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
	// <li>`overseas`: Global (outside the Chinese mainland).</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeBotClientIpListRequest struct {
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
	// <li>`overseas`: Global (outside the Chinese mainland).</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeBotClientIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotClientIpListRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotClientIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotClientIpListResponseParams struct {
	// The list of client IP data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecClientIp `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotClientIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotClientIpListResponseParams `json:"Response"`
}

func (r *DescribeBotClientIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotClientIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`bot_interceptNum`: Blocked bot requests;</li>
	// <li>`bot_noneRequestNum`: Uncategorized bot requests;</li>
	// <li>`bot_maliciousRequestNum`: Malicious bot requests;</li>
	// <li>`bot_suspectedRequestNum`: Suspected bot requests;</li>
	// <li>`bot_friendlyRequestNum`: Friendly bot requests;</li>
	// <li>`bot_normalRequestNum`: Normal bot requests.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

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

type DescribeBotDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`bot_interceptNum`: Blocked bot requests;</li>
	// <li>`bot_noneRequestNum`: Uncategorized bot requests;</li>
	// <li>`bot_maliciousRequestNum`: Malicious bot requests;</li>
	// <li>`bot_suspectedRequestNum`: Suspected bot requests;</li>
	// <li>`bot_friendlyRequestNum`: Friendly bot requests;</li>
	// <li>`bot_normalRequestNum`: Normal bot requests.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

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

func (r *DescribeBotDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "Domains")
	delete(f, "ZoneIds")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotDataResponseParams struct {
	// The list of bot attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotDataResponseParams `json:"Response"`
}

func (r *DescribeBotDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotHitRuleDetailRequestParams struct {
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

type DescribeBotHitRuleDetailRequest struct {
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

func (r *DescribeBotHitRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotHitRuleDetailRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotHitRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotHitRuleDetailResponseParams struct {
	// The hit rule information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecHitRuleInfo `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotHitRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotHitRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeBotHitRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotHitRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotLogRequestParams struct {
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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	// <li>`sipCountryCode`: The country code of the attacker IP;</li>
	// <li>`attackIp`: Attacker IP;</li>
	// <li>`ruleId`: Rule ID;</li>
	// <li>`eventId`: The event ID;</li>
	// <li>`ua`: User agent;</li>
	// <li>`requestMethod`: Request method;</li>
	// <li>`uri`: Uniform resource identifier.</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeBotLogRequest struct {
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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`action`: The action;</li>
	// <li>`sipCountryCode`: The country code of the attacker IP;</li>
	// <li>`attackIp`: Attacker IP;</li>
	// <li>`ruleId`: Rule ID;</li>
	// <li>`eventId`: The event ID;</li>
	// <li>`ua`: User agent;</li>
	// <li>`requestMethod`: Request method;</li>
	// <li>`uri`: Uniform resource identifier.</li>
	QueryCondition []*QueryCondition `json:"QueryCondition,omitempty" name:"QueryCondition"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotLogResponseParams struct {
	// The list of bot attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*BotLog `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeBotTopDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`bot_requestNum_labelType`: Top-ranked tag types by bot requests.</li>
	// <li>`bot_requestNum_url`: Top-ranked URLs by bot requests.</li>
	// <li>`bot_cipRequestNum_region`: Top-ranked client IPs by bot requests.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries the top rows of data. Maximum value: 1000. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

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

type DescribeBotTopDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`bot_requestNum_labelType`: Top-ranked tag types by bot requests.</li>
	// <li>`bot_requestNum_url`: Top-ranked URLs by bot requests.</li>
	// <li>`bot_cipRequestNum_region`: Top-ranked client IPs by bot requests.</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries the top rows of data. Maximum value: 1000. Top 10 rows of data will be queried if this field is not specified.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

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

func (r *DescribeBotTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Interval")
	delete(f, "QueryCondition")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotTopDataResponseParams struct {
	// The list of top-ranked bot attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*TopEntry `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotTopDataResponseParams `json:"Response"`
}

func (r *DescribeBotTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotTopDataResponse) FromJsonString(s string) error {
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
type DescribeDDoSAttackDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`ddos_attackMaxBandwidth`: Peak attack bandwidth;</li>
	// <li>`ddos_attackMaxPackageRate`: Peak attack packet rate;</li>
	// <li>`ddos_attackBandwidth`: Attack bandwidth;</li>
	// <li>`ddos_attackPackageRate`: Attack packet rate.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// The port number.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// The attack type. Values:
	// <li>`flood`: Flood;</li>
	// <li>`icmpFlood`: ICMP flood;</li>
	// <li>`all`: All attack types.</li>This field will be set to the default value `all` if not specified.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

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

type DescribeDDoSAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// List of statistical metric. Values:
	// <li>`ddos_attackMaxBandwidth`: Peak attack bandwidth;</li>
	// <li>`ddos_attackMaxPackageRate`: Peak attack packet rate;</li>
	// <li>`ddos_attackBandwidth`: Attack bandwidth;</li>
	// <li>`ddos_attackPackageRate`: Attack packet rate.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// The port number.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// The attack type. Values:
	// <li>`flood`: Flood;</li>
	// <li>`icmpFlood`: ICMP flood;</li>
	// <li>`all`: All attack types.</li>This field will be set to the default value `all` if not specified.
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

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
	delete(f, "Port")
	delete(f, "AttackType")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "ProtocolType")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataResponseParams struct {
	// List of DDoS attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*SecEntry `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeDDoSAttackEventDetailRequestParams struct {
	// The event ID to be queried.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// The event ID to be queried.
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventDetailResponseParams struct {
	// The details of a DDoS attack event.
	Data *DDoSAttackEventDetailData `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackEventDetailResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to display the details.
	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to display the details.
	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
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
	delete(f, "ProtocolType")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ShowDetail")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventResponseParams struct {
	// The list of DDoS attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
type DescribeDDoSAttackSourceEventRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSAttackSourceEventRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSAttackSourceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProtocolType")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackSourceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackSourceEventResponseParams struct {
	// The list of DDoS attacker data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DDoSAttackSourceEvent `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSAttackSourceEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackSourceEventResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackSourceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackSourceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackTopDataRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The statistical metric. Values:
	// <li>`ddos_attackFlux_protocol`: Top-ranked protocols by DDoS attack traffic.</li>
	// <li>`ddos_attackPackageNum_protocol`: Top-ranked protocols by DDoS attack packets.</li>
	// <li>`ddos_attackNum_attackType`: Top-ranked attack types by DDoS attacks.</li>
	// <li>`ddos_attackNum_sregion`: Top-ranked attack source regions by DDoS attacks.</li>
	// <li>`ddos_attackFlux_sip`: Top-ranked attacker IPs by DDoS attack traffic.</li>
	// <li>`ddos_attackFlux_sregion`: Top-ranked attack source regions by DDoS attack traffic.</li>
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
	// <li>`ddos_attackFlux_protocol`: Top-ranked protocols by DDoS attack traffic.</li>
	// <li>`ddos_attackPackageNum_protocol`: Top-ranked protocols by DDoS attack packets.</li>
	// <li>`ddos_attackNum_attackType`: Top-ranked attack types by DDoS attacks.</li>
	// <li>`ddos_attackNum_sregion`: Top-ranked attack source regions by DDoS attacks.</li>
	// <li>`ddos_attackFlux_sip`: Top-ranked attacker IPs by DDoS attack traffic.</li>
	// <li>`ddos_attackFlux_sregion`: Top-ranked attack source regions by DDoS attack traffic.</li>
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
type DescribeDDoSBlockListRequestParams struct {
	// The start time of the attack event.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The list of attack events.
	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSBlockListRequest struct {
	*tchttp.BaseRequest
	
	// The start time of the attack event.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The list of attack events.
	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EventIds")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSBlockListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSBlockListResponseParams struct {
	// The blocking time of a DDoS attack.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DDoSBlockData `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSBlockListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSBlockListResponseParams `json:"Response"`
}

func (r *DescribeDDoSBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSMajorAttackEventRequestParams struct {
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeDDoSMajorAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies sites by ID. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The list of DDoS policy IDs to be specified. All policies will be selected if this field is not specified.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" name:"PolicyIds"`

	// The protocol type. Values:
	// <li>`tcp`: TCP protocol;</li>
	// <li>`udp`: UDP protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Limit on paginated queries. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeDDoSMajorAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSMajorAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "ProtocolType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSMajorAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSMajorAttackEventResponseParams struct {
	// The list of large DDoS attack data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DDoSMajorAttackEvent `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSMajorAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSMajorAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDoSMajorAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSMajorAttackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesRequestParams struct {
	// Filter criteria. Each filter criteria can have up to 5 entries.
	// <li>`zone-id`: <br>Filter by <strong>site ID</strong>. Format: zone-xxx
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset for paginated queries. Default value: `0`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// Filter criteria. Each filter criteria can have up to 5 entries.
	// <li>`zone-id`: <br>Filter by <strong>site ID</strong>. Format: zone-xxx
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
type DescribeDnsRecordsRequestParams struct {
	// The site ID of the DNS record. All sites’ DNS records will be returned if this field is not specified.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`record-id`:<br>   Filter by <strong>DNS record ID</strong>, such as record-1a8df68z<br>   Type: String<br>Required: No
	// <li>`record-name`:<br>   Filter by <strong>DNS record name</strong><br>   Type: String<br>Required: No
	// <li>`record-type`:<br>   Filter by <strong>DNS record type</strong><br>   Type: String<br>Required: No<br>   Values:<br>   `A`: Point a domain name to an IPv4 address, such as 8.8.8.8.<br>   `AAAA`: Point a domain name to an IPv6 address.<br>   `CNAME`: Point a domain name to another domain name that can be resolved to an IP address.<br>   `TXT`: Identify and describe a domain name. It is usually used for domain verification and as SPF records (for anti-spam).<br>   `NS`: If you need to authorize a subdomain name to another DNS service provider for DNS resolution, you need to add an NS record. You cannot add an NS record for a root domain name.<br>   `CAA`: Specify CAs to issue certificates for sites.<br>   `SRV`: Identify a service used by a server. It is commonly used in Microsoft directory management.<br>  `MX`: Specify the mail server for receiving emails.
	// <li>`mode`:<br>   Filter by <strong>proxy mode</strong><br>   Type: String<br>Required: No<br>   Values:<br>   `dns_only`: Only DNS<br>   `proxied`: Proxied
	// <li>`ttl`:<br>   Filter by <strong>TTL</strong><br>   Type: String<br>Required: No
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// The sorting order. Values:
	// <li>`ASC`: Ascending order</li>
	// <li>`desc`: Descending order</li> Default value: asc
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// The match mode. Values:
	// <li>`all`: Return all records that match the specified filter.</li>
	// <li>`any`: Return any record that matches the specified filter.</li>Default value: all.
	Match *string `json:"Match,omitempty" name:"Match"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The sorting criteria. Values:
	// <li>`content`: DNS record content.</li>
	// <li>`created_on`: Creation time of the DNS record.</li>
	// <li>`mode`: Proxy mode.</li>
	// <li>`record-name`: DNS record name.</li>
	// <li>`ttl`: DNS TTL.</li>
	// <li>`record-type`: DNS record type.</li>If this field is not specified, the DNS records are sorted based on `record-type` and `recrod-name`.
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeDnsRecordsRequest struct {
	*tchttp.BaseRequest
	
	// The site ID of the DNS record. All sites’ DNS records will be returned if this field is not specified.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`record-id`:<br>   Filter by <strong>DNS record ID</strong>, such as record-1a8df68z<br>   Type: String<br>Required: No
	// <li>`record-name`:<br>   Filter by <strong>DNS record name</strong><br>   Type: String<br>Required: No
	// <li>`record-type`:<br>   Filter by <strong>DNS record type</strong><br>   Type: String<br>Required: No<br>   Values:<br>   `A`: Point a domain name to an IPv4 address, such as 8.8.8.8.<br>   `AAAA`: Point a domain name to an IPv6 address.<br>   `CNAME`: Point a domain name to another domain name that can be resolved to an IP address.<br>   `TXT`: Identify and describe a domain name. It is usually used for domain verification and as SPF records (for anti-spam).<br>   `NS`: If you need to authorize a subdomain name to another DNS service provider for DNS resolution, you need to add an NS record. You cannot add an NS record for a root domain name.<br>   `CAA`: Specify CAs to issue certificates for sites.<br>   `SRV`: Identify a service used by a server. It is commonly used in Microsoft directory management.<br>  `MX`: Specify the mail server for receiving emails.
	// <li>`mode`:<br>   Filter by <strong>proxy mode</strong><br>   Type: String<br>Required: No<br>   Values:<br>   `dns_only`: Only DNS<br>   `proxied`: Proxied
	// <li>`ttl`:<br>   Filter by <strong>TTL</strong><br>   Type: String<br>Required: No
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`

	// The sorting order. Values:
	// <li>`ASC`: Ascending order</li>
	// <li>`desc`: Descending order</li> Default value: asc
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// The match mode. Values:
	// <li>`all`: Return all records that match the specified filter.</li>
	// <li>`any`: Return any record that matches the specified filter.</li>Default value: all.
	Match *string `json:"Match,omitempty" name:"Match"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The sorting criteria. Values:
	// <li>`content`: DNS record content.</li>
	// <li>`created_on`: Creation time of the DNS record.</li>
	// <li>`mode`: Proxy mode.</li>
	// <li>`record-name`: DNS record name.</li>
	// <li>`ttl`: DNS TTL.</li>
	// <li>`record-type`: DNS record type.</li>If this field is not specified, the DNS records are sorted based on `record-type` and `recrod-name`.
	Order *string `json:"Order,omitempty" name:"Order"`
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
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Direction")
	delete(f, "Match")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnsRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnsRecordsResponseParams struct {
	// Total number of DNS records.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of DNS records
	DnsRecords []*DnsRecord `json:"DnsRecords,omitempty" name:"DnsRecords"`

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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeDnssecRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDnssecResponseParams struct {
	// The DNSSEC status. Values:
	// <li>`enabled`: Enabled</li>
	// <li>`disabled`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The DNSSEC information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DnssecInfo *DnssecInfo `json:"DnssecInfo,omitempty" name:"DnssecInfo"`

	// The update time of the site information.
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
type DescribeHostsSettingRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Offset for paginated queries. Default value: 0. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`host`:<br>   Filter by <strong>domain name </strong><br>   Type: String<br>   Required: No
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
	// <li>`host`:<br>   Filter by <strong>domain name </strong><br>   Type: String<br>   Required: No
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
	// <li>`zone-name`: <br>Filter by <strong>site name</strong><br>   Type: String<br>   Required: No
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeIdentificationsRequest struct {
	*tchttp.BaseRequest
	
	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-name`: <br>Filter by <strong>site name</strong><br>   Type: String<br>   Required: No
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
type DescribeLogTopicTaskDetailRequestParams struct {
	// ID of the shipping task.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeLogTopicTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// ID of the shipping task.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeLogTopicTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogTopicTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTopicTaskDetailResponseParams struct {
	// The shipping task details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogTopicDetailInfo *LogTopicDetailInfo `json:"LogTopicDetailInfo,omitempty" name:"LogTopicDetailInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogTopicTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogTopicTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeLogTopicTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTopicTaskDetailResponse) FromJsonString(s string) error {
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

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

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

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`tagKey`:<br>   Filter by <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>  Filter by <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`
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

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of subdomain names to be queried. All subdomain names will be selected if this field is not specified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The protocol type. Values:
	// <li>`http`: HTTP protocol;</li>
	// <li>`https`: HTTPS protocol;</li>
	// <li>`http2`: HTTP2 protocol;</li>
	// <li>`all`: All protocol types.</li>This field will be set to the default value `all` if not specified.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

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

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`tagKey`:<br>   Filter by <strong>tag key</strong><br>   Type: String<br>   Required: No</li>
	// <li>`tagValue`<br>  Filter by <strong>tag value</strong><br>   Type: String<br>   Required: No</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Area")
	delete(f, "Filters")
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
	// <li>`zone-id`:<br>   Filter by the <strong>site ID</strong>, such as zone-1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`job-id`:<br>   Filter by <strong>task ID</strong>, such as 1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`target`:<br>   Filter by <strong>target resource</strong>, such as http://www.qq.com/1.txt (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`domains`:<br>   Filter by <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`statuses`:<br>   Filter by <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: The task is in progress.<br>   `success`: The task succeeded.<br>   `failed`: The task failed.<br>   `timeout`: The task timed out.
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
	// <li>`zone-id`:<br>   Filter by the <strong>site ID</strong>, such as zone-1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`job-id`:<br>   Filter by <strong>task ID</strong>, such as 1379afjk91u32h (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`target`:<br>   Filter by <strong>target resource</strong>, such as http://www.qq.com/1.txt (up to one entry)<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`domains`:<br>   Filter by <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`statuses`:<br>   Filter by <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: The task is in progress.<br>   `success`: The task succeeded.<br>   `failed`: The task failed.<br>   `timeout`: The task timed out.
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
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Start time of the query.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`job-id`:<br> Filter by the <strong>Task ID</strong>, such as 1379afjk91u32h. Only one ID can be specified.<br>Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`target`:<br>   Filter by the <strong>resource address</strong>, such as http://www.qq.com/1.txt. Only one entry allowed.<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`domains`:<br>   Filter by the <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`statuses`:<br>   Filter by the <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: Tasks in progress<br>   `success`: Succeeded tasks<br>   `failed`: Failed tasks<br>   `timeout`: Timed-out tasks<li>`type`:<br>   Filter by the <strong>purging mode</strong>. Only one value allowed.<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `purge_url`: Purge URLs.<br>   `purge_prefix`: Purge prefixes.<br>   `purge_all`: Purge all caches.<br>   `purge_host`: Purge hostnames.
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Start time of the query.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`job-id`:<br> Filter by the <strong>Task ID</strong>, such as 1379afjk91u32h. Only one ID can be specified.<br>Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`target`:<br>   Filter by the <strong>resource address</strong>, such as http://www.qq.com/1.txt. Only one entry allowed.<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`domains`:<br>   Filter by the <strong>domain name</strong>, such as www.qq.com<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<li>`statuses`:<br>   Filter by the <strong>task status</strong><br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `processing`: Tasks in progress<br>   `success`: Succeeded tasks<br>   `failed`: Failed tasks<br>   `timeout`: Timed-out tasks<li>`type`:<br>   Filter by the <strong>purging mode</strong>. Only one value allowed.<br>   Type: String<br>   Required: No<br>   Fuzzy query: Not supported<br>   Values:<br>   `purge_url`: Purge URLs.<br>   `purge_prefix`: Purge prefixes.<br>   `purge_all`: Purge all caches.<br>   `purge_host`: Purge hostnames.
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

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`country`: Country/Region;</li>
	// <li>`domain`: Domain name;</li>
	// <li>`protocol`: Protocol type;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

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

type DescribeSingleL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// The start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query metric. Values:
	// <li>`l7Flow_singleIpRequest`: Number of requests from a single IP.</li>
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`country`: Country/Region;</li>
	// <li>`domain`: Domain name;</li>
	// <li>`protocol`: Protocol type;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

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

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`ruleId`: Filter by rule ID;</li>
	// <li>`proxyId`: Filter by connection ID.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
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

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// List of L4 proxy IDs. All L4 proxies will be selected if this field is not specified.
	ProxyIds []*string `json:"ProxyIds,omitempty" name:"ProxyIds"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`ruleId`: Filter by rule ID;</li>
	// <li>`proxyId`: Filter by connection ID.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
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

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`country`: Country/Region;</li>
	// <li>`domain`: Domain name;</li>
	// <li>`protocol`: Protocol type;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`statusCode`: Status code;</li>
	// <li>`browserType`: Browser type;</li>
	// <li>`deviceType`: Device type;</li>
	// <li>`operatingSystemType`: OS type;</li>
	// <li>`tlsVersion`: TLS version;</li>
	// <li>`url`: URL address;</li>
	// <li>`referer`: Refer header;</li>
	// <li>`ipVersion`: IP version;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
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

	// List of sites to be queried. All sites will be selected if this field is not specified.
	ZoneIds []*string `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minute;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`country`: Country/Region;</li>
	// <li>`domain`: Domain name;</li>
	// <li>`protocol`: Protocol type;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`statusCode`: Status code;</li>
	// <li>`browserType`: Browser type;</li>
	// <li>`deviceType`: Device type;</li>
	// <li>`operatingSystemType`: OS type;</li>
	// <li>`tlsVersion`: TLS version;</li>
	// <li>`url`: URL address;</li>
	// <li>`referer`: Refer header;</li>
	// <li>`ipVersion`: IP version;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`cacheType`: Cache type;</li>
	// <li>`domain`: Host/domain name;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`url`: URL address;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`cacheType`: Cache type;</li>
	// <li>`domain`: Host/domain name;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`url`: URL address;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`country`: Country/Region;</li>
	// <li>`domain`: Domain name;</li>
	// <li>`protocol`: Protocol type;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`statusCode`: Status code;</li>
	// <li>`browserType`: Browser type;</li>
	// <li>`deviceType`: Device type;</li>
	// <li>`operatingSystemType`: OS type;</li>
	// <li>`tlsVersion`: TLS version;</li>
	// <li>`url`: URL address;</li>
	// <li>`referer`: Refer header;</li>
	// <li>`ipVersion`: IP version;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`country`: Country/Region;</li>
	// <li>`domain`: Domain name;</li>
	// <li>`protocol`: Protocol type;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`statusCode`: Status code;</li>
	// <li>`browserType`: Browser type;</li>
	// <li>`deviceType`: Device type;</li>
	// <li>`operatingSystemType`: OS type;</li>
	// <li>`tlsVersion`: TLS version;</li>
	// <li>`url`: URL address;</li>
	// <li>`referer`: Refer header;</li>
	// <li>`ipVersion`: IP version;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`cacheType`: Cache type;</li>
	// <li>`domain`: Host/domain name;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`url`: URL address;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
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

	// The key of the parameter QueryCondition, which is used to specify a filter. Values:
	// <li>`cacheType`: Cache type;</li>
	// <li>`domain`: Host/domain name;</li>
	// <li>`resourceType`: Resource type;</li>
	// <li>`url`: URL address;</li>
	// <li>`tagKey`: Tag key;</li>
	// <li>`tagValue`: Tag value.</li>
	Filters []*QueryCondition `json:"Filters,omitempty" name:"Filters"`

	// The query time granularity. Values:
	// <li>`min`: 1 minute;</li>
	// <li>`5min`: 5 minutes;</li>
	// <li>`hour`: 1 hour;</li>
	// <li>`day`: 1 day.</li>If this field is not specified, the granularity will be determined based on the interval between the start time and end time as follows: 1-minute granularity applies for a 1-hour interval, 5-minute granularity for a 2-day interval, 1-hour granularity for a 7-day interval, and 1-day granularity for an interval of over 7 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Data storage region. Values:
	// <li>`overseas`: Global (outside the Chinese mainland);</li>
	// <li>`mainland`: Chinese mainland.</li>If this field is not specified, the data storage region will be determined based on the user’s location.
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
type DescribeWebProtectionAttackEventsRequestParams struct {
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

type DescribeWebProtectionAttackEventsRequest struct {
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
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebProtectionAttackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebProtectionAttackEventsResponseParams struct {
	// The list of CC attack events.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*CCInterceptEvent `json:"Data,omitempty" name:"Data"`

	// Total number of query results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
	// <li>`zone-name`:<br>   Filter by <strong>site name</strong><br>   Type: String<br>   Required: No<li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No<li>`status`:<br>   Filter by <strong>site status</strong><br>   Type: String<br>   Required: No<li>`tag-key`:<br>   Filter by <strong>tag key</strong><br>   Type: String<br>   Required: No<li>`tag-value`:<br>   Filter by <strong>tag value</strong><br>   Type: String<br>   Required: No<li>`Fuzzy`:<br>   Filter by <strong>values in fuzzy query</strong> (only `zone-name` allowed). Values limit: 1<br>   Type: Boolean<br>   Required: No<br>   Default value: false
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// The page offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The paginated query limit. Default value: 20. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter criteria. Each filter criteria can have up to 20 entries.
	// <li>`zone-name`:<br>   Filter by <strong>site name</strong><br>   Type: String<br>   Required: No<li>`zone-id`:<br>   Filter by <strong>site ID</strong>, such as zone-xxx<br>   Type: String<br>   Required: No<li>`status`:<br>   Filter by <strong>site status</strong><br>   Type: String<br>   Required: No<li>`tag-key`:<br>   Filter by <strong>tag key</strong><br>   Type: String<br>   Required: No<li>`tag-value`:<br>   Filter by <strong>tag value</strong><br>   Type: String<br>   Required: No<li>`Fuzzy`:<br>   Filter by <strong>values in fuzzy query</strong> (only `zone-name` allowed). Values limit: 1<br>   Type: Boolean<br>   Required: No<br>   Default value: false
	Filters []*AdvancedFilter `json:"Filters,omitempty" name:"Filters"`
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


	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`
}

type DnsData struct {
	// The time.
	Time *string `json:"Time,omitempty" name:"Time"`

	// The value.
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DnsRecord struct {
	// The record ID.
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

	// The DNS record type. Values:
	// <li>`A`: Point a domain name to an IPv4 address, such as 8.8.8.8.</li>
	// <li>`AAAA`: Point a domain name to an IPv6 address.</li>
	// <li>`MX`: It is used for email servers. The record value and priority parameters are provided by email service providers. If there are multiple MX records, the lower the priority value, the higher the priority.</li>
	// <li>`CNAME`: Point a domain name to another domain name that can be resolved to an IP address.</li>
	// <li>`TXT`: Identify and describe a domain name. It is usually used for domain verification and as SPF records (for anti-spam).</li>
	// <li>`NS`: If you need to authorize a subdomain name to another DNS service provider for DNS resolution, you need to add an NS record. You cannot add an NS record for a root domain name.</li>
	// <li>`CAA`: Specify CAs to issue certificates for sites.</li>
	// <li>`SRV`: Identify a service used by a server. It is commonly used in Microsoft directory management.</li>
	DnsRecordType *string `json:"DnsRecordType,omitempty" name:"DnsRecordType"`

	// The record name.
	DnsRecordName *string `json:"DnsRecordName,omitempty" name:"DnsRecordName"`

	// The record value.
	Content *string `json:"Content,omitempty" name:"Content"`

	// The proxy mode. Values:
	// <li>`dns_only`: Only DNS</li>
	// <li>`proxied`: Proxied</li>
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// TTL (in seconds). The smaller the value, the faster the record changes take effect.
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// The MX record priority. The smaller the value, the higher the priority.
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// The creation time.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// The modification time.
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`

	// The lock status of the domain name.
	Locked *bool `json:"Locked,omitempty" name:"Locked"`

	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The DNS record status. Values:
	// <li>`active`: Activated</li>
	// <li>`pending`: Deactivated</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The CNAME address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// The service used by the domain name. Values:
	// <li>`lb`: Load balancing</li>
	// <li>`security`: Security protection</li>
	// <li>`l4`: L4 proxy</li>
	DomainStatus []*string `json:"DomainStatus,omitempty" name:"DomainStatus"`
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

type FailReason struct {
	// Failure reason.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// List of resources failed to be processed.
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type FileAscriptionInfo struct {

	IdentifyPath *string `json:"IdentifyPath,omitempty" name:"IdentifyPath"`


	IdentifyContent *string `json:"IdentifyContent,omitempty" name:"IdentifyContent"`
}

type Filter struct {
	// The name of the field to filter.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value of the filtered field.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FollowOrigin struct {
	// Whether to enable the configuration of following the origin server. Valid values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
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

type Header struct {
	// HTTP header name
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
}

type Identification struct {
	// The site name.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// The verification status. Values:
	// <li>`pending`: The verification is ongoing.</li>
	// <li>`finished`: The verification completed.</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// The site ownership information.
	Ascription *AscriptionInfo `json:"Ascription,omitempty" name:"Ascription"`

	// The NS record of the domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`


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
	// The site ownership information.
	Ascription *AscriptionInfo `json:"Ascription,omitempty" name:"Ascription"`


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

type LogTopicDetailInfo struct {
	// Name of the shipping task.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// Type of the shipping task.
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// List of tasks.
	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`

	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// Name of the logset.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// Topic ID of the shipping task.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic name of the shipping task.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Retention period of the shipping task topic. Unit: day.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Whether the shipping task is enabled.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Creation time in the format of YYYY-mm-dd HH:MM:SS.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland;</li>
	// <li>`overseas`: Global (outside the Chinese mainland).</li>
	Area *string `json:"Area,omitempty" name:"Area"`

	// ID of the site.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Name of the site.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Whether the shipping task is deleted. Values:
	// <li>`yes`: The shipping task is deleted;</li>
	// <li>`no`: The shipping task is not deleted.</li>
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
type ModifyDnsRecordRequestParams struct {
	// The record ID.
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The DNS record type. Values:
	// <li>`A`: Point a domain name to an IPv4 address, such as 8.8.8.8.</li>
	// <li>`AAAA`: Point a domain name to an IPv6 address.</li>
	// <li>`MX`: It is used for email servers. The record value and priority parameters are provided by email service providers. If there are multiple MX records, the lower the priority value, the higher the priority.</li>
	// <li>`CNAME`: Point a domain name to another domain name that can be resolved to an IP address.</li>
	// <li>`TXT`: Identify and describe a domain name. It is usually used for domain verification and as SPF records (for anti-spam).</li>
	// <li>`NS`: If you need to authorize a subdomain name to another DNS service provider for DNS resolution, you need to add an NS record. You cannot add an NS record for a root domain name.</li>
	// <li>`CAA`: Specify CAs to issue certificates for sites.</li>
	// <li>`SRV`: Identify a service used by a server. It is commonly used in Microsoft directory management.</li>
	DnsRecordType *string `json:"DnsRecordType,omitempty" name:"DnsRecordType"`

	// The record name, which consists of the host record and site name. Note that the original configuration will be used if this field is not specified.
	DnsRecordName *string `json:"DnsRecordName,omitempty" name:"DnsRecordName"`

	// The record content. Note that the original configuration will be used if this field is not specified.
	Content *string `json:"Content,omitempty" name:"Content"`

	// TTL (in seconds). The smaller the value, the faster the record changes take effect. Default value: 300. Note that the original configuration will be used if this field is not specified.
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// Specifies a value in the range 1–50 when you make changes to the MX records. A smaller value indicates higher priority. Note that the default value 0 will be used if this field is not specified.
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// The proxy mode. Values:
	// <li>`dns_only`: Only DNS</li>
	// <li>`proxied`: Proxied</li></li>The original configuration will apply if this field is not specified.
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ModifyDnsRecordRequest struct {
	*tchttp.BaseRequest
	
	// The record ID.
	DnsRecordId *string `json:"DnsRecordId,omitempty" name:"DnsRecordId"`

	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The DNS record type. Values:
	// <li>`A`: Point a domain name to an IPv4 address, such as 8.8.8.8.</li>
	// <li>`AAAA`: Point a domain name to an IPv6 address.</li>
	// <li>`MX`: It is used for email servers. The record value and priority parameters are provided by email service providers. If there are multiple MX records, the lower the priority value, the higher the priority.</li>
	// <li>`CNAME`: Point a domain name to another domain name that can be resolved to an IP address.</li>
	// <li>`TXT`: Identify and describe a domain name. It is usually used for domain verification and as SPF records (for anti-spam).</li>
	// <li>`NS`: If you need to authorize a subdomain name to another DNS service provider for DNS resolution, you need to add an NS record. You cannot add an NS record for a root domain name.</li>
	// <li>`CAA`: Specify CAs to issue certificates for sites.</li>
	// <li>`SRV`: Identify a service used by a server. It is commonly used in Microsoft directory management.</li>
	DnsRecordType *string `json:"DnsRecordType,omitempty" name:"DnsRecordType"`

	// The record name, which consists of the host record and site name. Note that the original configuration will be used if this field is not specified.
	DnsRecordName *string `json:"DnsRecordName,omitempty" name:"DnsRecordName"`

	// The record content. Note that the original configuration will be used if this field is not specified.
	Content *string `json:"Content,omitempty" name:"Content"`

	// TTL (in seconds). The smaller the value, the faster the record changes take effect. Default value: 300. Note that the original configuration will be used if this field is not specified.
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// Specifies a value in the range 1–50 when you make changes to the MX records. A smaller value indicates higher priority. Note that the default value 0 will be used if this field is not specified.
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// The proxy mode. Values:
	// <li>`dns_only`: Only DNS</li>
	// <li>`proxied`: Proxied</li></li>The original configuration will apply if this field is not specified.
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
	delete(f, "DnsRecordId")
	delete(f, "ZoneId")
	delete(f, "DnsRecordType")
	delete(f, "DnsRecordName")
	delete(f, "Content")
	delete(f, "TTL")
	delete(f, "Priority")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnsRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnsRecordResponseParams struct {
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The DNSSEC status. Values:
	// <li>`enabled`: Enabled</li>
	// <li>`disabled`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyDnssecRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The DNSSEC status. Values:
	// <li>`enabled`: Enabled</li>
	// <li>`disabled`: Disabled</li>
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
	delete(f, "ZoneId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDnssecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDnssecResponseParams struct {
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
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names that the certificate will be attached to.
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Certificate information. Note that only `CertId` is required. If it is not specified, the default certificate will be used.
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitempty" name:"ServerCertInfo"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names that the certificate will be attached to.
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Certificate information. Note that only `CertId` is required. If it is not specified, the default certificate will be used.
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitempty" name:"ServerCertInfo"`
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
type ModifyLogTopicTaskRequestParams struct {
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// ID of the log topic.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Type of the shipping entity. Values:
	// <li>`domain`: L7 acceleration logs;</li>
	// <li>`application`: L4 acceleration logs;</li>
	// <li>`web-rateLiming`: Rate limiting logs;</li>
	// <li>`web-attack`: Web security logs;</li>
	// <li>`web-rule`: Custom rule logs;</li>
	// <li>`web-bot`: Bot management logs.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// Name of the shipping task.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// The new topic name. If you do not specify this field, no changes will be made.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// The new logset name.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// The retention period of the updated logset.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// List of shipping entities to be deleted.
	DropEntityList []*string `json:"DropEntityList,omitempty" name:"DropEntityList"`

	// List of shipping entities to be added.
	AddedEntityList []*string `json:"AddedEntityList,omitempty" name:"AddedEntityList"`
}

type ModifyLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Region of the logset.
	LogSetRegion *string `json:"LogSetRegion,omitempty" name:"LogSetRegion"`

	// ID of the logset.
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// ID of the log topic.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Type of the shipping entity. Values:
	// <li>`domain`: L7 acceleration logs;</li>
	// <li>`application`: L4 acceleration logs;</li>
	// <li>`web-rateLiming`: Rate limiting logs;</li>
	// <li>`web-attack`: Web security logs;</li>
	// <li>`web-rule`: Custom rule logs;</li>
	// <li>`web-bot`: Bot management logs.</li>
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// Name of the shipping task.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// The new topic name. If you do not specify this field, no changes will be made.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// The new logset name.
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// The retention period of the updated logset.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// List of shipping entities to be deleted.
	DropEntityList []*string `json:"DropEntityList,omitempty" name:"DropEntityList"`

	// List of shipping entities to be added.
	AddedEntityList []*string `json:"AddedEntityList,omitempty" name:"AddedEntityList"`
}

func (r *ModifyLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "LogSetRegion")
	delete(f, "LogSetId")
	delete(f, "TopicId")
	delete(f, "EntityType")
	delete(f, "TaskName")
	delete(f, "TopicName")
	delete(f, "LogSetName")
	delete(f, "Period")
	delete(f, "DropEntityList")
	delete(f, "AddedEntityList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogTopicTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogTopicTaskResponseParams `json:"Response"`
}

func (r *ModifyLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogTopicTaskResponse) FromJsonString(s string) error {
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
type ModifyZoneCnameSpeedUpRequestParams struct {
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The CNAME acceleration status. Values:
	// <li>`enabled`: Enabled</li>
	// <li>`disabled`: Disabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyZoneCnameSpeedUpRequest struct {
	*tchttp.BaseRequest
	
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The CNAME acceleration status. Values:
	// <li>`enabled`: Enabled</li>
	// <li>`disabled`: Disabled</li>
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
	delete(f, "ZoneId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneCnameSpeedUpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCnameSpeedUpResponseParams struct {
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
	// The site ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// The site access method. Values:
	// <li>`full`: Access through a name server.</li>
	// <li>`partial`: Access through a CNAME record.</li>The original configuration will apply if this field is not specified.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The custom name servers. If this field is not specified, the default name servers will be used.
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitempty" name:"VanityNameServers"`
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


	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`
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

	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`
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
	// <li>`https`: Force HTTPS for origin-pull. This only supports port 443 on the origin server.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// When OriginType is COS, you can specify if access to private buckets is allowed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosPrivateAccess *string `json:"CosPrivateAccess,omitempty" name:"CosPrivateAccess"`
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

	// Plan option. Values:
	// <li>`sta`: Standard plan that supports content delivery network outside Chinese mainland;</li>
	// <li>`sta_with_bot`: Standard plan that supports content delivery network outside Chinese mainland and bot management;</li>
	// <li>`sta_cm`: Standard plan that supports content delivery network inside Chinese mainland;</li>
	// <li>`sta_cm_with_bot`: Standard plan that supports content delivery network inside Chinese mainland and bot management;</li>
	// <li>`ent`: Enterprise plan that supports content delivery network outside Chinese mainland;</li>
	// <li>`ent_with_bot`: Enterprise plan that supports content delivery network outside Chinese mainland and bot management;</li>
	// <li>`ent_cm`: Enterprise plan that supports content delivery network inside Chinese mainland;</li>
	// <li>`ent_cm_with_bot`: Enterprise plan that supports content delivery network inside Chinese mainland and bot management.</li>
	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`

	// Plan price (in CNY fen/US cent). The price unit depends on the settlement currency.
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// Quota on security acceleration requests
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// Number of sites to be bound to the plan
	SiteNumber *uint64 `json:"SiteNumber,omitempty" name:"SiteNumber"`

	// Acceleration region. Values:
	// <li>`mainland`: Chinese mainland;</li>
	// <li>`overseas`: Global (Chinese mainland not included).</li>
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

type QueryCondition struct {
	// The key of QueryCondition.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The conditional operator. Values:
	// <li>`equals`: Equal to;</li>
	// <li>`notEquals`: Not equal to;</li>
	// <li>`include`: Contain;</li>
	// <li>`notInclude`: Not contain;</li>
	// <li>`startWith`: Start with;</li>
	// <li>`notStartWith`: Not start with;</li>
	// <li>`endWith`: End with;</li>
	// <li>`notEndWith`: Not end with.</li>
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

	// Quota type. Values:
	// <li>`purge_prefix`: Purge prefixes;</li>
	// <li>`purge_url`: Purge URLs;</li>
	// <li>`purge_host`: Purge hostnames;</li>
	// <li>`purge_all`: Purge all caches.</li>
	Type *string `json:"Type,omitempty" name:"Type"`
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
	// Feature execution conditions.
	// Note: If any condition in the array is met, the feature will run.
	Conditions []*RuleAndConditions `json:"Conditions,omitempty" name:"Conditions"`

	// Feature to be executed.
	Actions []*Action `json:"Actions,omitempty" name:"Actions"`
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
	// <li>equal: Equal to.</li>
	// <li>notequal: Not equal to.</li>
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Match type. Valid values:
	// <li>`host`: All</li>
	// <li>`filename`: File name</li>
	// <li>`extension`: File extension</li>
	// <li>`host`: HOST: .</li>
	// <li>`full_url`: The full URL of the current site. It must contain the HTTP protocol, host, and path.</li>
	// <li>`url`: The URL path of the current site.</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// Parameter values of the match type. Each match type has the following valid values:
	// <li>`Target=extension`:  The extension of the file, such as `jpg` and `txt`.</li>
	// <li>`Target=filename`: The file name without the extension.</li>
	// <li>`Target=host`: Values can be `all` 
	// or a host, such as `www.maxx55.com`.</li>
	// <li>`Target=url`: A URL request path under the current site, such as `/example`.</li>
	// <li>`Target=full_url`: A complete URL request under the current site. It must contain the protocol, host, and path, such as `https://www.maxx55.cn/example`.</li>
	Values []*string `json:"Values,omitempty" name:"Values"`
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
	// <li>`default`: Default certificate;</lil>
	// <li>`upload`: Custom certificate;</li>
	// <li>`managed`: Tencent Cloud-managed certificate.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
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

type SmartRouting struct {
	// Whether to enable smart acceleration. Values:
	// <li>`on`: Enable</li>
	// <li>`off`: Disable</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Sv struct {
	// The parameter key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The parameter value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type SwitchLogTopicTaskRequestParams struct {
	// Topic ID of the shipping task.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Whether to enable the shipping task. Values:
	// <li>`true`: Enable the shipping task;</li>
	// <li>`false`: Disable the shipping task.</li>
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
}

type SwitchLogTopicTaskRequest struct {
	*tchttp.BaseRequest
	
	// Topic ID of the shipping task.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Whether to enable the shipping task. Values:
	// <li>`true`: Enable the shipping task;</li>
	// <li>`false`: Disable the shipping task.</li>
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
}

func (r *SwitchLogTopicTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchLogTopicTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "IsOpen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchLogTopicTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchLogTopicTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchLogTopicTaskResponse struct {
	*tchttp.BaseResponse
	Response *SwitchLogTopicTaskResponseParams `json:"Response"`
}

func (r *SwitchLogTopicTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchLogTopicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Whether CNAME flattening is enabled. Valid values:
	// <li>`enabled`: Enabled.</li>
	// <li>`disabled`: Disabled.</li>
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


	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitempty" name:"ClientIpCountry"`
}