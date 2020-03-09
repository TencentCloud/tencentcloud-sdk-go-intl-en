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

package v20180606

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddCdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Business type of acceleration domain name
	// web: static acceleration
	// download: download acceleration
	// media: streaming media VOD acceleration
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Origin server configuration
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Project ID, which is 0 by default, indicating **Default Project**
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP blacklist/whitelist configuration
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Status code cache configuration
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Smart compression configuration
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Bandwidth cap configuration
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range origin-pull configuration
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 origin-pull follow-redirect configuration
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Error code redirect configuration (This feature is in beta test and not fully available yet.)
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// Request header configuration
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Response header configuration
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Download speed configuration
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Node cache key configuration
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Header cache configuration
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// Video dragging configuration
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// Cache expiration time configuration
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cross-border linkage optimization configuration
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Https acceleration configuration
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Timestamp hotlink protection configuration
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// SEO configuration
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Access protocol forced redirect configuration
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer hotlink protection configuration
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Browser cache configuration (This feature is in beta test and not fully available yet.)
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Ipv6 configuration (This feature is in beta test and not fully available yet.)
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Specific configuration for region attributes
	// Applicable to use cases where the configuration of accelerating domain names inside mainland China is inconsistent with the configuration outside mainland China.
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// Domain name acceleration region
	// mainland: acceleration inside mainland China
	// overseas: acceleration outside mainland China
	// global: global acceleration
	// To use overseas acceleration and global acceleration, you need to enable the overseas acceleration service first
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *AddCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AdvanceCacheRule struct {

	// Rule types:
	// all: all files take effect
	// file: specified file suffixes take effect
	// directory: specified paths take effect
	// path: specified absolute paths take effect
	// default: the cache rules when the origin server has not returned max-age
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// Matching content under the corresponding types:
	// For "all", enter an asterisk (*).
	// For "file", enter the suffix, such as jpg, txt.
	// For "directory", enter the path, such as /xxx/test/.
	// For "path", enter the corresponding absolute path, such as /xxx/test.html.
	// For "default", enter "no max-age".
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents" list`

	// Cache expiration time
	// Unit: second. The maximum value is 365 days.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type AdvancedCache struct {

	// Cache expiration rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheRules []*AdvanceCacheRule `json:"CacheRules,omitempty" name:"CacheRules" list`

	// Forced cache configuration
	// on: enabled
	// off: disabled
	// When it is enabled, if the origin server returns no-cache, no-store headers, node caching will still be performed according to the cache expiration rules.
	// It is disabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`

	// Ignore the Set-Cookie header of an origin server
	// on: enabled
	// off: disabled
	// It is disabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitempty" name:"IgnoreSetCookie"`
}

type Authentication struct {

	// Hotlink protection configuration switch
	// on: enabled
	// off: disabled
	// When it is enabled, only one mode needs to be configured. Other modes need to be set to null.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Timestamp hotlink protection mode A configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	TypeA *AuthenticationTypeA `json:"TypeA,omitempty" name:"TypeA"`

	// Timestamp hotlink protection mode B configuration (the mode B backend is being upgraded and the configuration is currently not supported)
	// Note: this field may return null, indicating that no valid values can be obtained.
	TypeB *AuthenticationTypeB `json:"TypeB,omitempty" name:"TypeB"`

	// Timestamp hotlink protection mode C configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	TypeC *AuthenticationTypeC `json:"TypeC,omitempty" name:"TypeC"`

	// Timestamp hotlink protection mode D configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	TypeD *AuthenticationTypeD `json:"TypeD,omitempty" name:"TypeD"`
}

type AuthenticationTypeA struct {

	// The key for signature calculation
	// Only digits, upper and lower-case letters are allowed. Length limit: 6-32 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature parameter name configuration
	// Only upper and lower-case letters, digits, and underscores (_) are allowed. It cannot start with a digit. Length limit: 1-100 characters.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Signature expiration time settings
	// Unit: second. The maximum value is 31536000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings for authentication/no authentication
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist: indicating that all types apart from the FileExtensions list are authenticated
	// blacklist: indicating that only the types in the FileExtensions list are authenticated
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type AuthenticationTypeB struct {

	// The key for signature calculation
	// Only digits, upper and lower-case letters are allowed. Length limit: 6-32 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature expiration time settings
	// Unit: second. The maximum value is 31536000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings for authentication/no authentication
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist: indicating that all types apart from the FileExtensions list are authenticated
	// blacklist: indicating that only the types in the FileExtensions list are authenticated
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type AuthenticationTypeC struct {

	// The key for signature calculation
	// Only digits, upper and lower-case letters are allowed. Length limit: 6-32 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature expiration time settings
	// Unit: second. The maximum value is 31536000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings for authentication/no authentication
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist: indicating that all types apart from the FileExtensions list are authenticated
	// blacklist: indicating that only the types in the FileExtensions list are authenticated
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type AuthenticationTypeD struct {

	// The key for signature calculation
	// Only digits, upper and lower-case letters are allowed. Length limit: 6-32 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature expiration time settings
	// Unit: second. The maximum value is 31536000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings for authentication/no authentication
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// whitelist: indicating that all types apart from the FileExtensions list are authenticated
	// blacklist: indicating that only the types in the FileExtensions list are authenticated
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Signature parameter name configuration
	// Only upper and lower-case letters, digits, and underscores (_) are allowed. It cannot start with a digit. Length limit: 1-100 characters.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Timestamp parameter name settings
	// Only upper and lower-case letters, digits, and underscores (_) are allowed. It cannot start with a digit. Length limit: 1-100 characters.
	TimeParam *string `json:"TimeParam,omitempty" name:"TimeParam"`

	// Timestamp settings
	// dec: decimal
	// hex: hexadecimal
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
}

type BandwidthAlert struct {

	// Bandwidth cap configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Bandwidth cap threshold (in bps)
	// Note: this field may return null, indicating that no valid values can be obtained.
	BpsThreshold *int64 `json:"BpsThreshold,omitempty" name:"BpsThreshold"`

	// Operation after threshold is reached
	// RESOLVE_DNS_TO_ORIGIN: directly origin-pull. It is only supported for domain names of external origin.
	// RETURN_404: return 404 to all requests.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CounterMeasure *string `json:"CounterMeasure,omitempty" name:"CounterMeasure"`

	// The last time the bandwidth cap threshold was triggered
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" name:"LastTriggerTime"`
}

type BriefDomain struct {

	// Domain name ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tencent Cloud account ID.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// CDN acceleration domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain name CNAME.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Domain name status. pending: under review; rejected: review failed; processing: review succeeded and under deployment; online: enabled; offline: disabled; deleted: deleted.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Domain name business type. web: static acceleration; download: download acceleration; media: streaming media acceleration.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Domain name creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Domain name update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin server configuration details.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Domain name block status, including normal, overdue, quota, malicious, ddos, idle, unlicensed, capping, and readonly.
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// Acceleration region, including mainland, overseas, and global.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Domain name lock status. normal: not locked; mainland: locked in mainland China; overseas: locked outside mainland China; global: locked globally.
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`
}

type Cache struct {

	// Basic cache expiration time configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	SimpleCache *SimpleCache `json:"SimpleCache,omitempty" name:"SimpleCache"`

	// Advanced cache expiration time configuration (This feature is in beta test and not fully available yet.)
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdvancedCache *AdvancedCache `json:"AdvancedCache,omitempty" name:"AdvancedCache"`
}

type CacheKey struct {

	// Whether to enable full-path cache
	// on: enable full-path cache (i.e., disable parameter filter)
	// off: disable full-path cache (i.e., enable parameter filter)
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`
}

type CacheOptResult struct {

	// List of succeeded URLs
	// Note: This field may return null, indicating that no valid values can be obtained.
	SuccessUrls []*string `json:"SuccessUrls,omitempty" name:"SuccessUrls" list`

	// List of failed URLs
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailUrls []*string `json:"FailUrls,omitempty" name:"FailUrls" list`
}

type CappingRule struct {

	// Rule types:
	// all: all files take effect
	// file: specified file suffixes take effect
	// directory: specified paths take effect
	// path: specified absolute paths take effect
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Matching content under the corresponding types for RuleType: 
	// For "all", enter an asterisk (*).
	// For "file", enter the suffix, such as jpg, txt.
	// For "directory", enter the path, such as /xxx/test/.
	// For "path", enter the corresponding absolute path, such as /xxx/test.html.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`

	// Downstream speed value settings (in KB/s)
	KBpsThreshold *int64 `json:"KBpsThreshold,omitempty" name:"KBpsThreshold"`
}

type CdnData struct {

	// Queries the specified metric:
	// flux: traffic (in bytes)
	// bandwidth: bandwidth (in bps)
	// request: number of requests
	// fluxHitRate: traffic hit rate (in %)
	// statusCode: status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx status codes will be returned (in entries)
	// 2XX: Returns the aggregate list of 2xx status codes and the data for status codes starting with 2 (in entries)
	// 3XX: Returns the aggregate list of 3xx status codes and the data for status codes starting with 3 (in entries)
	// 4XX: Returns the aggregate list of 4xx status codes and the data for status codes starting with 4 (in entries)
	// 5XX: Returns the aggregate list of 5xx status codes and the data for status codes starting with 5 (in entries)
	// Alternatively, you can specify a status code for querying.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Detailed data combination
	DetailData []*TimestampData `json:"DetailData,omitempty" name:"DetailData" list`

	// Aggregate data combination
	SummarizedData *SummarizedData `json:"SummarizedData,omitempty" name:"SummarizedData"`
}

type CdnIp struct {

	// IP of the node.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Whether the IP is a Tencent Cloud CDN cache node. `yes`: it is a Tencent Cloud CDN cache node; `no`: it is not.
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// District/country where the node is located. `unknown`: the node location is unknown.
	Location *string `json:"Location,omitempty" name:"Location"`

	// Activation and deactivation history of the node.
	History []*CdnIpHistory `json:"History,omitempty" name:"History" list`

	// Service region of the node. `mainland`: Mainland China; `overseas`: outside Mainland China; `unknown`: unknown
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CdnIpHistory struct {

	// Node status. `online`: activated; `offline`: deactivated
	Status *string `json:"Status,omitempty" name:"Status"`

	// Operation time. If its value is `null`, it means there is no status change record.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Datetime *string `json:"Datetime,omitempty" name:"Datetime"`
}

type ClientCert struct {

	// Client Certificate
	// PEM format, requires Base64 encoding.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Client certificate name
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Certificate expiration time
	// When it is used as an input parameter, it can be left blank.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Certificate issuance time
	// When it is used as an input parameter, it can be left blank.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`
}

type Compatibility struct {

	// Compatibility flag status code.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type Compression struct {

	// Smart compression configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Compression rules array
	// Note: this field may return null, indicating that no valid values can be obtained.
	CompressionRules []*CompressionRule `json:"CompressionRules,omitempty" name:"CompressionRules" list`
}

type CompressionRule struct {

	// true: must be set as true, to enable compression
	// Note: this field may return null, indicating that no valid values can be obtained.
	Compress *bool `json:"Compress,omitempty" name:"Compress"`

	// Compress according to the file suffix type
	// Such as: jpg, txt
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions" list`

	// The minimum file size to trigger compression (in bytes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinLength *int64 `json:"MinLength,omitempty" name:"MinLength"`

	// The maximum file size to trigger compression (in bytes)
	// The maximum value is 30 MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxLength *int64 `json:"MaxLength,omitempty" name:"MaxLength"`

	// File compression algorithm
	// gzip: specifies Gzip compression
	// brotli: It can be enabled when the Gzip compression is specified
	// Note: this field may return null, indicating that no valid values can be obtained.
	Algorithms []*string `json:"Algorithms,omitempty" name:"Algorithms" list`
}

type DeleteCdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name
	// The domain name status should be **Disabled**
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDataRequest struct {
	*tchttp.BaseRequest

	// Queries start time, such as 2018-09-04 10:40:00; the returned result is later than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query end time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the first returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Queries end time, such as 2018-09-04 10:40:00; the returned result is earlier than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the last returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the query metric, which can be:
	// flux: traffic (in bytes)
	// bandwidth: bandwidth (in bps)
	// request: number of requests
	// fluxHitRate: traffic hit rate (in %)
	// statusCode: status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx status codes will be returned (in entries)
	// 2xx: Returns the aggregate list of 2xx status codes and the data for status codes starting with 2 (in entries)
	// 3xx: Returns the aggregate list of 3xx status codes and the data for status codes starting with 3 (in entries)
	// 4xx: Returns the aggregate list of 4xx status codes and the data for status codes starting with 4 (in entries)
	// 5xx: Returns the aggregate list of 5xx status codes and the data for status codes starting with 5 (in entries)
	// It is supported to specify a status code for query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the list of domain names to be queried
	// Up to 30 domain names can be queried at a time
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Time granularity; valid values:
	// `min`: data with 1-minute granularity is returned when the queried period is no longer than 24 hours. This value is not supported if the service region you want to query is outside Mainland China;
	// `5min`: data with 5-minute granularity is returned when the queried period is no longer than 31 days;
	// `hour`: data with 1-hour granularity is returned when the queried period is no longer than 31 days;
	// `day`: data with 1-day granularity is returned when the queried period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) during a multi-domain-name query.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported)
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies an ISP when you query the CDN data within Mainland China. If it is left blank, all ISPs will be queried.
	// To view ISP codes, see [ISP Code Mappings](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// If you have specified an ISP, you cannot specify a province or an IP protocol for the same query.
	Isp *int64 `json:"Isp,omitempty" name:"Isp"`

	// Specifies a province when you query the CDN data within Mainland China. If it is left blank, all provinces will be queried.
	// Specifies a country/region when you query the CDN data outside Mainland China. If it is left blank, all countries/regions will be queried.
	// To view codes of provinces or countries/regions, see [Province Code Mappings](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// If you have specified a province for your query on CDN data within mainland China, you cannot specify an ISP or an IP protocol for the same query.
	District *int64 `json:"District,omitempty" name:"District"`

	// Specifies the protocol to be queried; if you leave it blank, all protocols will be queried.
	// all: All protocols
	// http: specifies the HTTP metric to be queried
	// https: specifies the HTTPS metric to be queried
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Specifies the data source to be queried, which can be seen as the whitelist function.
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Specifies an IP protocol; if it is left blank, all IP protocols will be queried.
	// `all`: All protocols
	// `ipv4`: IPv4
	// `ipv6`: IPv6
	// If the IP protocol to be queried is specified, the district and ISP cannot be specified at the same time
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Specifies a service region. If this value is left blank, CDN data within Mainland China will be queried.
	// `mainland`: specifies to query CDN data within Mainland China;
	// `overseas`: specifies to query CDN data outside Mainland China.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a region type for your query on CDN data outside Mainland China. If this parameter is left blank, data on the service region will be queried. This parameter is valid only when `Area` is `overseas`.
	// `server`: specifies to query data on the service region where Tencent Cloud CDN nodes are located;
	// `client`: specifies to query data on the client region where the request devices are located.
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`
}

func (r *DescribeCdnDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time granularity of the returned data. Specify one of the following during querying:
	// min: 1 minute
	// 5min: 5 minutes
	// hour: 1 hour
	// day: 1 day
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// Returned data details of the specified conditional query
		Data []*ResourceData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDomainLogsRequest struct {
	*tchttp.BaseRequest

	// Specifies a domain name for the query
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Starting time, such as `2019-09-04 00:00:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as `2019-09-04 12:00:00`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paged queries. Default value: 0 (the first page)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paged queries. Default value: 100. Maximum value: 1,000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a region for the query.
	// `mainland`: specifies to return the download link of logs on acceleration within Mainland China;
	// `overseas`: specifies to return the download link of logs on acceleration outside Mainland China;
	// `global`: specifies to return a download link of logs on acceleration within Mainland China and a link of logs on acceleration outside Mainland China.
	// Default value: `mainland`.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeCdnDomainLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDomainLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDomainLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Download link of the log package
		DomainLogs []*DomainLog `json:"DomainLogs,omitempty" name:"DomainLogs" list`

		// Total number of entries obtained
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnDomainLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDomainLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnIpRequest struct {
	*tchttp.BaseRequest

	// List of IPs to be queried
	Ips []*string `json:"Ips,omitempty" name:"Ips" list`
}

func (r *DescribeCdnIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Node ownership details
		Ips []*CdnIp `json:"Ips,omitempty" name:"Ips" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsConfigRequest struct {
	*tchttp.BaseRequest

	// Offset for paged queries. Default value: 0 (the first page).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paged queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, complex type.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters" list`

	// Sorting rules
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeDomainsConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of domain names
		Domains []*DetailDomain `json:"Domains,omitempty" name:"Domains" list`

		// The number of domain names that matched the query conditions
	// Used for paged queries
		TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest

	// Offset for paged queries. Default value: 0 (the first page).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paged queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, complex type.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of domain names
		Domains []*BriefDomain `json:"Domains,omitempty" name:"Domains" list`

		// The number of domain names that matched the query conditions
	// Used for paged queries
		TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpVisitRequest struct {
	*tchttp.BaseRequest

	// Query start time, such as 2018-09-04 10:40:10; the returned result is later than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:10 and the query time granularity is 5 minutes, the time for the first returned entry will be 2018-09-04 10:40:00.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time, such as 2018-09-04 10:40:10; the returned result is earlier than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:10 and the query time granularity is 5 minutes, the time for the last returned entry will be 2018-09-04 10:40:00.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the list of domain names to be queried; up to 30 domain names can be queried at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Time granularity, which can be:
	// 5min: 5 minutes. If the query period is within 24 hours, `5min` will be used by default.
	// day: 1 day. If the query period is longer than 24 hours, `day` will be used by default.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeIpVisitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpVisitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpVisitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time granularity of data statistics, which supports 5min (5 minutes) and day (1 day).
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// Origin-pull data details of each resource.
		Data []*ResourceData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpVisitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpVisitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMapInfoRequest struct {
	*tchttp.BaseRequest

	// Query type:
	// `isp`: queries ISP codes
	// `district`: queries codes of provinces (Mainland China) or countries/regions (outside Mainland China)
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeMapInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMapInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMapInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of mappings.
		MapInfoList []*MapInfo `json:"MapInfoList,omitempty" name:"MapInfoList" list`

		// The relationship between server region ID and sub-region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
		ServerRegionRelation []*RegionMapRelation `json:"ServerRegionRelation,omitempty" name:"ServerRegionRelation" list`

		// The relationship between client region ID and sub-region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
		ClientRegionRelation []*RegionMapRelation `json:"ClientRegionRelation,omitempty" name:"ClientRegionRelation" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMapInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMapInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataRequest struct {
	*tchttp.BaseRequest

	// Query start time, such as 2018-09-04 10:40:00; the returned result is later than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query end time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the first returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time, such as 2018-09-04 10:40:00; the returned result is earlier than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the last returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the query metric, which can be:
	// flux: origin-pull traffic (in bytes)
	// bandwidth: origin-pull bandwidth (in bps)
	// request: number of origin-pull requests
	// failRequest: number of failed origin-pull requests
	// failRate: origin-pull failure rate (in %)
	// statusCode: origin-pull status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx origin-pull status codes will be returned (in entries)
	// 2xx: Returns the aggregate list of 2xx origin-pull status codes and the data for origin-pull status codes starting with 2 (in entries)
	// 3xx: Returns the aggregate list of 3xx origin-pull status codes and the data for origin-pull status codes starting with 3 (in entries)
	// 4xx: Returns the aggregate list of 4xx origin-pull status codes and the data for origin-pull status codes starting with 4 (in entries)
	// 5xx: Returns the aggregate list of 5xx origin-pull status codes and the data for origin-pull status codes starting with 5 (in entries)
	// It is supported to specify a status code for query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the list of domain names to be queried; up to 30 domain names can be queried at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Time granularity; valid values:
	// `min`: data with 1-minute granularity is returned when the queried period is no longer than 24 hours. This value is not supported if the service region you want to query is outside Mainland China;
	// `5min`: data with 5-minute granularity is returned when the queried period is no longer than 31 days;
	// `hour`: data with 1-hour granularity is returned when the queried period is no longer than 31 days;
	// `day`: data with 1-day granularity is returned when the queried period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) when multiple `Domains` are passed in.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported)
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies a service region. If this value is left blank, CDN data within Mainland China will be queried.
	// `mainland`: specifies to query CDN data within Mainland China;
	// `overseas`: specifies to query CDN data outside Mainland China.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeOriginDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOriginDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time granularity of data statistics, which supports min (1 minute), 5min (5 minutes), hour (1 hour), and day (1 day).
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// Origin-pull data details of each resource.
		Data []*ResourceOriginData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOriginDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOriginDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayTypeRequest struct {
	*tchttp.BaseRequest

	// Specifies a service region.
	// `mainland`: queries billing methods within Mainland China;
	// `overseas`: queries billing methods outside Mainland China.
	// Default value: `mainland`.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Billing modes:
	// `flux`: bill-by-traffic
	// `bandwidth`: bill-by-bandwidth
	// When you switch the billing mode for a daily-billing-cycle account, if there is bandwidth usage on the day, this field indicates the billing mode that will take effect on the next day; otherwise, it indicates the billing mode that has already taken effect
		PayType *string `json:"PayType,omitempty" name:"PayType"`

		// Billing cycle:
	// day: daily settlement
	// month: monthly settlement
		BillingCycle *string `json:"BillingCycle,omitempty" name:"BillingCycle"`

		// Billing method:
	// monthMax: billed by the monthly average of daily peak traffic (monthly settlement)
	// day95: billed by the daily 95th percentile bandwidth (monthly settlement)
	// month95: billed by the monthly 95th percentile bandwidth (monthly settlement)
	// sum: billed by the total traffic (daily or monthly settlement)
	// max: billed by the peak bandwidth (daily settlement)
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// Billing method outside Mainland China:
	// `all`: unified billing for all regions
	// `multiple`: separate billing for different regions
		RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

		// Currently billing mode in effect:
	// `flux`: bill-by-traffic
	// `bandwidth`: bill-by-bandwidth
		CurrentPayType *string `json:"CurrentPayType,omitempty" name:"CurrentPayType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest

	// Specifies a purge type:
	// `url`: URL purge record
	// `path`: Directory purge record
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// Specifies the starting time of the period you want to query, such as `2018-08-08 00:00:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Specifies the end time of the period you want to query, such as 2018-08-08 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies a task ID when you want to query by task ID.
	// You must specify either a task ID or a starting time for your query.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Offset for paged queries. Default value: 0 (the first page)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paged queries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// You can filter the results by domain name or a complete URL beginning with `http(s)://`
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Specifies a task state for your query:
	// `fail`: purge failed
	// `done`: purge succeeded
	// `process`: purge in progress
	Status *string `json:"Status,omitempty" name:"Status"`

	// Specifies a purge region for your query:
	// `mainland`: within Mainland China
	// `overseas`: outside Mainland China
	// `global`: global
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Detailed purge record.
	// Note: This field may return null, indicating that no valid values can be obtained.
		PurgeLogs []*PurgeTask `json:"PurgeLogs,omitempty" name:"PurgeLogs" list`

		// Total number of tasks, which is used for pagination.
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushTasksRequest struct {
	*tchttp.BaseRequest

	// Starting time, such as `2018-08-08 00:00:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as `2018-08-08 23:59:59`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies a task ID for your query.
	// You must specify either a task ID or a starting time.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Specifies a keyword for your query. Please enter a domain name or a complete URL beginning with `http(s)://`
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Offset for paged queries. Default value: 0 (the first page)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paged queries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a region for your query:
	// `mainland`: within Mainland China
	// `overseas`: outside Mainland China
	// `global`: global
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a task state for your query:
	// `fail`: prefetch failed
	// `done`: prefetch succeeded
	// `process`: prefetch in progress
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribePushTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Prefetch history
	// Note: This field may return null, indicating that no valid values can be obtained.
		PushLogs []*PushTask `json:"PushLogs,omitempty" name:"PushLogs" list`

		// Total number of tasks, which is used for pagination.
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUrlViolationsRequest struct {
	*tchttp.BaseRequest

	// Offset for paged queries. Default value: 0 (the first page).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paged queries. Default value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specified domain name query
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`
}

func (r *DescribeUrlViolationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUrlViolationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUrlViolationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Details of violating URLs
	// Note: this field may return null, indicating that no valid values can be obtained.
		UrlRecordList []*ViolationUrl `json:"UrlRecordList,omitempty" name:"UrlRecordList" list`

		// Total number of records, which is used for pagination.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUrlViolationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUrlViolationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetailDomain struct {

	// Domain name ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tencent Cloud account ID.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Accelerated domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain name CNAME.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Domain name status. pending: under review; rejected: review failed; processing: review succeeded and under deployment; online: enabled; offline: disabled; deleted: deleted.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Domain name business type. web: static acceleration; download: download acceleration; media: streaming media acceleration.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Domain name creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Domain name update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin server configuration.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP blacklist/whitelist configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Status code cache configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Smart compression configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Bandwidth cap configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range origin-pull configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301 and 302 automatic origin-pull follow-redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Error code redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// Origin-pull request header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Origin server response header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Download speed configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Node cache configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Follows origin server cache header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// Video dragging configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// Cache rules configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cross-border optimization configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// HTTPS configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Timestamp hotlink protection configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// SEO configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Domain name block status, including normal, overdue, quota, malicious, ddos, idle, unlicensed, capping, and readonly.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// Access protocol forced redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Hotlink protection configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Browser cache rules configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// IPv6 configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Whether it is compatible with configurations in old versions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Compatibility *Compatibility `json:"Compatibility,omitempty" name:"Compatibility"`

	// Specific configuration by region.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// Acceleration region, including mainland, overseas, and global.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Domain name lock status. normal: not locked; mainland: locked in mainland China; overseas: locked outside mainland China; global: locked globally.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`

	// Origin-pull timeout configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitempty" name:"OriginPullTimeout"`
}

type DisableCachesRequest struct {
	*tchttp.BaseRequest

	// List of URLs to be blocked
	// Up to 100 entries can be submitted at a time and 3,000 entries per day.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *DisableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCachesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableCachesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Submission result
	// Note: This field may return null, indicating that no valid values can be obtained.
		CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCachesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DomainFilter struct {

	// Filter field name, the list supported is as follows:
	// - origin: master origin server.
	// - domain: domain name.
	// - resourceId: domain name id.
	// - status: domain name status, online, offline, or processing.
	// - serviceType: service type, web, download, or media.
	// - projectId: project ID.
	// - domainType: master origin server type, cname indicates external origin, COS indicates COS origin.
	// - fullUrlCache: full-path cache, which can be on or off.
	// - https: whether to configure HTTPS, which can be on, off or processing.
	// - originPullProtocol: origin-pull protocol type. It supports HTTP, follow, or HTTPS.
	// - tagKey: tag key.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter field value.
	Value []*string `json:"Value,omitempty" name:"Value" list`

	// Whether to enable fuzzy query. Only origin or domain is supported for the filter field name.
	// When fuzzy query is enabled, the maximum Value length is 1. When fuzzy query is disabled, the maximum Value length is 5.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type DomainLog struct {

	// Starting time of the log package
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the log package
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Log package download link
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// Acceleration region corresponding to the log package
	// `mainland`: within Mainland China
	// `overseas`: outside Mainland China
	Area *string `json:"Area,omitempty" name:"Area"`

	// Log package filename
	LogName *string `json:"LogName,omitempty" name:"LogName"`
}

type DownstreamCapping struct {

	// Downstream speed configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Downstream speed limiting rules
	// Note: this field may return null, indicating that no valid values can be obtained.
	CappingRules []*CappingRule `json:"CappingRules,omitempty" name:"CappingRules" list`
}

type EnableCachesRequest struct {
	*tchttp.BaseRequest

	// List of unblocked URLs
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *EnableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCachesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableCachesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result list
	// Note: This field may return null, indicating that no valid values can be obtained.
		CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCachesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ErrorPage struct {

	// Status code redirect configuration switch
	// on: enabled
	// off: disabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Status code redirect rules configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	PageRules []*ErrorPageRule `json:"PageRules,omitempty" name:"PageRules" list`
}

type ErrorPageRule struct {

	// Status code
	// It supports 400, 403, 404, 500.
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// Redirect status code settings
	// It supports 301 or 302.
	RedirectCode *int64 `json:"RedirectCode,omitempty" name:"RedirectCode"`

	// Redirect URL
	// It requires a full redirect path, such as https://www.test.com/error.html.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

type FollowRedirect struct {

	// Origin-pull follow-redirect switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ForceRedirect struct {

	// Access forced redirect configuration switch
	// on: enabled
	// off: disabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Access forced redirect types
	// http: forced HTTP redirect
	// https: forced HTTPS redirect
	// Note: this field may return null, indicating that no valid values can be obtained.
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`

	// Status code returned for forced redirect 
	// It supports 301, 302.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type GetDisableRecordsRequest struct {
	*tchttp.BaseRequest

	// Starting time, such as `2018-12-12 10:24:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as 2018-12-14 10:24:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the URL to be queried
	Url *string `json:"Url,omitempty" name:"Url"`

	// Current URL status
	// disable: The URL remains disabled, and accessing it will return an error 403
	// enable: The URL is enabled (unblocked) and can be normally accessed
	Status *string `json:"Status,omitempty" name:"Status"`

	// Offset for paged queries. Default value: 0 (the first page)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paged queries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetDisableRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDisableRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDisableRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Blocking history
	// Note: This field may return null, indicating that no valid values can be obtained.
		UrlRecordList []*UrlRecord `json:"UrlRecordList,omitempty" name:"UrlRecordList" list`

		// Total number of tasks, which is used for pagination
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDisableRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDisableRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HttpHeaderPathRule struct {

	// HTTP header setting method
	// add: add header. If a header exists, then there will be a repeated header.
	// set: only supports origin-pull header configuration. If a header exists, it will be overwritten. If one does not exist, then the header will be added.
	// del: delete header
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderMode *string `json:"HeaderMode,omitempty" name:"HeaderMode"`

	// HTTP header name. Up to 100 characters can be set.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// HTTP header value. Up to 1000 characters can be set.
	// It is not required when Mode is del
	// It is required when Mode is add/set
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`

	// Rule types:
	// all: all files take effect
	// file: specified file suffixes take effect
	// directory: specified paths take effect
	// path: specified absolute paths take effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Matching content under the corresponding types for RuleType:
	// For "all", enter an asterisk (*).
	// For "file", enter the suffix, such as jpg, txt.
	// For "directory", enter the path, such as /xxx/test/.
	// For "path", enter the corresponding absolute path, such as /xxx/test.html.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`
}

type Https struct {

	// HTTPS configuration switch
	// on: enabled
	// off: disabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// HTTP2 configuration switch
	// on: enabled
	// off: disabled
	// The first time HTTPS acceleration is enabled, it will enable HTTP2 configuration by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// OCSP configuration switch
	// on: enabled
	// off: disabled
	// It is disabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// Client certificate authentication feature
	// on: enabled
	// off: disabled
	// It is disabled by default. If it is enabled, you need to upload the client certificate information. This configuration is in beta test and not fully available yet.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyClient *string `json:"VerifyClient,omitempty" name:"VerifyClient"`

	// Server certificate configuration information
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertInfo *ServerCert `json:"CertInfo,omitempty" name:"CertInfo"`

	// Client certificate configuration information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClientCertInfo *ClientCert `json:"ClientCertInfo,omitempty" name:"ClientCertInfo"`

	// Spdy configuration switch
	// on: enabled
	// off: disabled
	// It is disabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spdy *string `json:"Spdy,omitempty" name:"Spdy"`

	// HTTPS certificate deployment status
	// closed: already closed
	// deploying: being deployed
	// deployed: successfully deployed
	// failed: deployment failed
	// Note: this field may return null, indicating that no valid values can be obtained.
	SslStatus *string `json:"SslStatus,omitempty" name:"SslStatus"`
}

type IpFilter struct {

	// IP blacklist/whitelist configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// IP blacklist/whitelist type
	// whitelist: whitelist
	// blacklist: blacklist
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// IP blacklist/whitelist list
	// It supports IPs in X.X.X.X format, or /8, /16, /24 format IP ranges.
	// Up to 50 whitelists or blacklists can be entered
	// Note: this field may return null, indicating that no valid values can be obtained.
	Filters []*string `json:"Filters,omitempty" name:"Filters" list`
}

type IpFreqLimit struct {

	// IP access limit configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Sets the number limit of request per second
	// 514 will be returned to the requests that exceed the limit
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
}

type Ipv6 struct {

	// Whether to enable the IPv6 feature for a domain name, which can be on or off.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ListTopDataRequest struct {
	*tchttp.BaseRequest

	// Query start date. Example: 2018-09-09.
	// It only supports data query at daily granularity. The date information in the input parameter is the start date.
	// Returns data generated at or after 00:00:00 on the start date.
	// It only supports querying of data within 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end date. Example: 2018-09-10
	// It only supports data query at daily granularity. The date information in the input parameter is the end date.
	// Returns data generated before or at 23:59:59 on the end date.
	// EndTime must be greater than or equal to StartTime
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Objects to be sorted. Valid values:
	// `url`: sorts access URLs with query string parameters included. Supported filters are `flux` and `request`.
	// `path`: sorts access URLs with query string parameters excluded. Supported filters are `flux` and `request`. You need to be whitelisted before using this feature.
	// `district`: sorts provinces or countries/regions. Supported filters are `flux` and `request`.
	// `isp`: sorts ISPs. Supported filters are `flux` and `request`.
	// `host`: sorts domain name access data. Supported filters are `flux`, `request`, `bandwidth`, `fluxHitRate`, `2XX`, `3XX`, `4XX`, `5XX` and `statusCode`.
	// `originHost`: sorts by domain name origin-pull data. Supported filters are `flux`, `request`, `bandwidth`, `origin_2XX`, `origin_3XX`, `oringin_4XX`, `origin_5XX` and `OriginStatusCode`
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Metric name used for sorting:
	// flux: If Metric is `host`, it indicates the access traffic; if Metric is `originHost`, it indicates the origin-pull traffic.
	// bandwidth: If Metric is `host`, it indicates the access bandwidth; if Metric is `originHost`, it indicates the origin-pull bandwidth.
	// request: If Metric is `host`, it indicates the number of access requests; if Metric is `originHost`, it indicates the number of origin-pull requests.
	// fluxHitRate: Average traffic hit rate
	// 2XX: access 2XX status code
	// 3XX: access 3XX status code
	// 4XX: access 4XX status code
	// 5XX: access 5XX status code
	// origin_2XX: origin-pull 2XX status code
	// origin_3XX: origin-pull 3XX status code
	// origin_4XX: origin-pull 4XX status code
	// origin_5XX: origin-pull 5XX status code
	// statusCode: statistics of a specific access status code which is specified in the `Code` parameter.
	// OriginStatusCode: statistics of a specific origin-pull status code which is specified in the `Code` parameter.
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// Specifies the list of domain names to be queried; up to 30 domain names can be queried at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Default value: `false`, indicating that results for all domain names are returned together when you query multiple domain names.
	// If `Metric` is `Url`, `Path`, `District`, or `Isp` and `Filter` is `flux` or `request`, you can set this parameter to `true`, indicating that results for each domain name are returned.
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// When Filter is `statusCode` or `OriginStatusCode`, enter a code to query and sort results.
	Code *string `json:"Code,omitempty" name:"Code"`

	// Specifies a service region for the query. If it is left blank, CDN data within Mainland China will be queried.
	// `mainland`: specifies to query CDN data within Mainland China;
	// `overseas`: specifies to query CDN data outside Mainland China. Supported metrics are `url`, `district`, `host`, and `originHost`. If `Metric` is `originHost`, supported filters are `flux`, `request`, and `bandwidth`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a region type for the query. If it is left blank, data on the service region will be queried. This parameter is only valid when `Area` is `overseas` and `Metric` is `District` or `Host`.
	// `server`: specifies to query data on the service region where Tencent Cloud CDN nodes are located;
	// `client`: specifies to query data on the client region where the request devices are located; if `Metric` is `Host`, supported filters are `flux`, `request`, and `bandwidth`.
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`
}

func (r *ListTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTopDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTopDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Top access data details of each resource
		Data []*TopData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTopDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MainlandConfig struct {

	// Timestamp hotlink protection configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// Bandwidth cap configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Cache rules configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cache configurations.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Smart compression configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Download speed limit configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Error code redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 301 and 302 automatic origin-pull follow-redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Access protocol forced redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// IP blacklist/whitelist configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Browser cache rules configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Origin server configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Cross-border optimization configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Range origin-pull configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// Hotlink protection configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Origin-pull request header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Origin server response header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Follows origin server cache header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// SEO configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Domain name business type. web: static acceleration; download: download acceleration; media: streaming media acceleration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Status code cache configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Video dragging configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`
}

type MapInfo struct {

	// Object ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Object name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type MaxAge struct {

	// Browser cache configuration switch
	// on: enabled
	// off: disabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// MaxAge rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxAgeRules []*MaxAgeRule `json:"MaxAgeRules,omitempty" name:"MaxAgeRules" list`
}

type MaxAgeRule struct {

	// Rule types:
	// all: all files take effect
	// file: specified file suffixes take effect
	// directory: specified paths take effect
	// path: specified absolute paths take effect
	MaxAgeType *string `json:"MaxAgeType,omitempty" name:"MaxAgeType"`

	// Matching content under the corresponding types for MaxAgeType:
	// For "all", enter an asterisk (*).
	// For "file", enter the suffix, such as jpg, txt.
	// For "directory", enter the path, such as /xxx/test/.
	// For "path", enter the corresponding absolute path, such as /xxx/test.html.
	MaxAgeContents []*string `json:"MaxAgeContents,omitempty" name:"MaxAgeContents" list`

	// MaxAge time settings (in seconds)
	MaxAgeTime *int64 `json:"MaxAgeTime,omitempty" name:"MaxAgeTime"`
}

type Origin struct {

	// Master origin server list
	// When modifying the origin server, you need to enter the corresponding OriginType.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Origins []*string `json:"Origins,omitempty" name:"Origins" list`

	// Master origin server type
	// The following types are supported in input parameters:
	// domain: domain name type
	// cos: COS origin
	// ip: IP list is used as origin server
	// ipv6: origin server list is a single IPv6 address
	// ip_ipv6: origin server list is multiple IPv4 addresses and an IPv6 address
	// The following types of output parameters are added:
	// image: cloud Infinite origin
	// ftp: historical FTP origin, which is no longer maintained.
	// When modifying Origins, you need to enter the corresponding OriginType.
	// The IPv6 feature is not fully available yet. To use this feature, you need to apply for it first.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Host header used when pulling the master origin server. If the Host header is not entered, it will be the acceleration domain name by default.
	// If a wildcard domain name is accessed, then the Host header is the sub-domain name during the access by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// When OriginType is COS, you can specify whether to allow access to private buckets.
	// Note: to enable this configuration, you need to authorize the CDN to access this private bucket first.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CosPrivateAccess *string `json:"CosPrivateAccess,omitempty" name:"CosPrivateAccess"`

	// Origin-pull protocol configuration
	// http: forced HTTP origin-pull
	// follow: protocol follow origin-pull
	// https: forced HTTPS origin-pull. It only supports origin server port 443 for origin-pull.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// Backup origin server list
	// When modifying the backup origin server, you need to enter the corresponding BackupOriginType.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BackupOrigins []*string `json:"BackupOrigins,omitempty" name:"BackupOrigins" list`

	// Backup origin server type, which supports the following types:
	// domain: domain name type
	// ip: IP list is used as origin server
	// When modifying BackupOrigins, you need to enter the corresponding BackupOriginType.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BackupOriginType *string `json:"BackupOriginType,omitempty" name:"BackupOriginType"`

	// Host header used when pulling the backup origin server. If the Host header is not entered, it will be the ServerName of master origin server by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BackupServerName *string `json:"BackupServerName,omitempty" name:"BackupServerName"`
}

type OriginPullOptimization struct {

	// Cross-border origin-pull optimization configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Cross-border types
	// OVToCN: origin-pull from outside mainland China to inside mainland China
	// CNToOV: origin-pull from inside mainland China to outside mainland China
	// Note: this field may return null, indicating that no valid values can be obtained.
	OptimizationType *string `json:"OptimizationType,omitempty" name:"OptimizationType"`
}

type OriginPullTimeout struct {

	// The origin-pull connection timeout (in seconds). Valid range: 5-60.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// The origin-pull receipt timeout (in seconds). Valid range: 10-60.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReceiveTimeout *uint64 `json:"ReceiveTimeout,omitempty" name:"ReceiveTimeout"`
}

type OverseaConfig struct {

	// Timestamp hotlink protection configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// Bandwidth cap configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Cache rules configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cache configurations.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Smart compression configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Download speed limit configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Error code redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 301 and 302 automatic origin-pull follow-redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Access protocol forced redirect configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// IP blacklist/whitelist configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Browser cache rules configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Origin server configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Cross-border optimization configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Range origin-pull configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// Hotlink protection configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Origin-pull request header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Origin server response header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Follows origin server cache header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// SEO configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Domain name business type. web: static acceleration; download: download acceleration; media: streaming media acceleration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Status code cache configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Video dragging configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest

	// List of directories. The protocol header such as "http://" or "https://" needs to be included.
	Paths []*string `json:"Paths,omitempty" name:"Paths" list`

	// Purge type:
	// `flush`: purges updated resources
	// `delete`: purges all resources
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgePathCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Purge task ID. Directories submitted in one request share a task ID.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgePathCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgeTask struct {

	// Purge task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Purged URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// Purge task status
	// `fail`: purge failed
	// `done`: purge succeeded
	// `process`: purge in progress
	Status *string `json:"Status,omitempty" name:"Status"`

	// Purge type
	// `url`: URL purge
	// `path`: directory purge
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// Purge method
	// `flush`: purges updated resources; this type is available only for directory purges
	// `delete`: purges all resources
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// Purge task submission time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type PurgeUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// List of URLs. The protocol header such as "http://" or "https://" needs to be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *PurgeUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgeUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Purge task ID. URLs submitted in one request share a task ID.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgeUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushTask struct {

	// Prefetch task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Prefetched URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// Prefetch task status
	// `fail`: prefetch failed
	// `done`: prefetch succeeded
	// `process`: prefetch in progress
	Status *string `json:"Status,omitempty" name:"Status"`

	// Prefetch progress in percentage
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// Prefetch task submission time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Prefetch region
	// `mainland`: within Mainland China
	// `overseas`: outside Mainland China
	// `global`: global
	Area *string `json:"Area,omitempty" name:"Area"`

	// Prefetch task update time
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type PushUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// List of URLs. The protocol header such as "http://" or "https://" needs to be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// Specifies the User-Agent header of an HTTP prefetch request when it is forwarded to the origin server
	// Default value: `TencentCdn`
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Destination region for the prefetch
	// `mainland`: prefetches resources to nodes within Mainland China
	// `overseas`: prefetches resources to nodes outside Mainland China
	// `global`: prefetches resources to global nodes
	// Default value: `mainland`. You can prefetch a URL to nodes in a region provided that CDN service has been enabled for the domain name in the URL in the region.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *PushUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlsCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the submitted task
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PushUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlsCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RangeOriginPull struct {

	// Range GETs configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Referer struct {

	// Referer blacklist/whitelist configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Referer blacklist/whitelist configuration rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	RefererRules []*RefererRule `json:"RefererRules,omitempty" name:"RefererRules" list`
}

type RefererRule struct {

	// Rule types:
	// all: all files take effect
	// file: specified file suffixes take effect
	// directory: specified paths take effect
	// path: specified absolute paths take effect
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Matching content under the corresponding types for RuleType:
	// For "all", enter an asterisk (*).
	// For "file", enter the suffix, such as jpg, txt.
	// For "directory", enter the path, such as /xxx/test/.
	// For "path", enter the corresponding absolute path, such as /xxx/test.html.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`

	// Referer configuration types
	// whitelist: whitelist
	// blacklist: blacklist
	RefererType *string `json:"RefererType,omitempty" name:"RefererType"`

	// Referer content list
	Referers []*string `json:"Referers,omitempty" name:"Referers" list`

	// Whether to allow empty referer
	// true: allow empty referer
	// false: do not allow empty referer
	AllowEmpty *bool `json:"AllowEmpty,omitempty" name:"AllowEmpty"`
}

type RegionMapRelation struct {

	// Region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// List of sub-region IDs
	SubRegionIdList []*int64 `json:"SubRegionIdList,omitempty" name:"SubRegionIdList" list`
}

type RequestHeader struct {

	// Custom request header configuration switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom request header configuration rules
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules" list`
}

type ResourceData struct {

	// Resource name, which is classified as follows based on different query conditions:
	// A specific domain name: This indicates the details of this domain name
	// multiDomains: This indicates the aggregate details of multiple domain names
	// Project ID: This displays the ID of the specifically queried project
	// all: This indicates the details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Data details of a resource
	CdnData []*CdnData `json:"CdnData,omitempty" name:"CdnData" list`
}

type ResourceOriginData struct {

	// Resource name, which is classified as follows based on different query conditions:
	// A specific domain name: This indicates the details of this domain name
	// multiDomains: This indicates the aggregate details of multiple domain names
	// Project ID: This displays the ID of the specifically queried project
	// all: This indicates the details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Origin-pull data details
	OriginData []*CdnData `json:"OriginData,omitempty" name:"OriginData" list`
}

type ResponseHeader struct {

	// Custom response header switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom response header rules
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules" list`
}

type ResponseHeaderCache struct {

	// Origin server header cache switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Seo struct {

	// SEO configuration switch
	// on: enabled
	// off: disabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ServerCert struct {

	// Server certificate ID
	// It is auto-generated when the certificate is being hosted by the SSL Certificate Service
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Server certificate name
	// It is auto-generated when the certificate is being hosted by the SSL Certificate Service
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Server certificate information
	// It is required when uploading an external certificate, which should contain the complete certificate chain.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Server key information
	// It is required when uploading an external certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// Certificate expiration time
	// When it is used as an input parameter, it can be left blank.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Certificate issuance time
	// When it is used as an input parameter, it can be left blank.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// Certificate remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SimpleCache struct {

	// Cache expiration time rules
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheRules []*SimpleCacheRule `json:"CacheRules,omitempty" name:"CacheRules" list`

	// Follows origin server Cache-Control: max-age configurations
	// on: enabled
	// off: disabled
	// If it is enabled, resources that do not match CacheRules rules will be cached by the node according to the max-age value returned by the origin server. Resources that match CacheRules rules will be cached on the node according to the cache expiration time set in CacheRules.
	// It conflicts with CompareMaxAge. They cannot be enabled at the same time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`

	// Forced cache
	// on: enabled
	// off: disabled
	// It is disabled by default. If it is enabled, no-store and no-cache resources returned from the origin server will be cached according to CacheRules rules.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`

	// Ignores the Set-Cookie header of the origin server
	// on: enabled
	// off: disabled
	// It is disabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitempty" name:"IgnoreSetCookie"`

	// Advanced cache expiration configuration. If it is enabled, the max-age value returned by the origin server will be compared with the cache expiration time set in CacheRules, and the smallest value will be cached on the node.
	// on: enabled
	// off: disabled
	// It is disabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	CompareMaxAge *string `json:"CompareMaxAge,omitempty" name:"CompareMaxAge"`
}

type SimpleCacheRule struct {

	// Rule types:
	// all: all files take effect
	// file: specified file suffixes take effect
	// directory: specified paths take effect
	// path: specified absolute paths take effect
	// index: home page
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// Matching content under the corresponding types for CacheType
	// For "all", enter an asterisk (*).
	// For "file", enter the suffix, such as jpg, txt.
	// For "directory", enter the path, such as /xxx/test/.
	// For "path", enter the corresponding absolute path, such as /xxx/test.html.
	// For "index", enter a backslash (/).
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents" list`

	// Cache expiration time settings
	// Unit: second. The maximum value is 365 days.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type Sort struct {

	// Sorting field, which currently supports:
	// createTime, domain name creation time.
	// certExpireTime, certificate expiration time.
	Key *string `json:"Key,omitempty" name:"Key"`

	// asc/desc, which is desc by default.
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

type SpecificConfig struct {

	// Specific configuration for domain name inside mainland China.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mainland *MainlandConfig `json:"Mainland,omitempty" name:"Mainland"`

	// Specific configuration for domain name outside mainland China.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Overseas *OverseaConfig `json:"Overseas,omitempty" name:"Overseas"`
}

type StartCdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name
	// The domain name status should be **Disabled**
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StartCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StatusCodeCache struct {

	// Status code cache expiration configuration switch
	// on: enabled
	// off: disabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Status code cache expiration rules details
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheRules []*StatusCodeCacheRule `json:"CacheRules,omitempty" name:"CacheRules" list`
}

type StatusCodeCacheRule struct {

	// HTTP status code
	// Supports 403 and 404 status codes
	StatusCode *string `json:"StatusCode,omitempty" name:"StatusCode"`

	// Status code cache expiration time (in seconds)
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type StopCdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name
	// The domain name status should be **Enabled**
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StopCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopCdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopCdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SummarizedData struct {

	// Aggregation method, which can be:
	// sum: aggregate summation
	// max: maximum value; in bandwidth mode, the peak bandwidth is calculated based on the aggregate data with 5-minute granularity.
	// avg: average value
	Name *string `json:"Name,omitempty" name:"Name"`

	// Aggregate data value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TimestampData struct {

	// Statistical point in time in forward rounding mode
	// Taking the 5-minute granularity as an example, 13:35:00 indicates that the statistical interval is between 13:35:00 and 13:39:59.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Data value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TopData struct {

	// Resource name, which is classified as follows based on different query conditions:
	// A specific domain name: This indicates the details of this domain name
	// multiDomains: This indicates the aggregate details of multiple domain names
	// Project ID: This displays the ID of the specifically queried project
	// all: This indicates the details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Detailed sorting results
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData" list`
}

type TopDetailData struct {

	// Datatype name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Data value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UpdateDomainConfigRequest struct {
	*tchttp.BaseRequest

	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Origin server configuration
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP blacklist/whitelist configuration
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Status code cache configuration
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Smart compression configuration
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Bandwidth cap configuration
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range origin-pull configuration
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 origin-pull follow-redirect configuration
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Error code redirect configuration (This feature is in beta test and not fully available yet.)
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// Request header configuration
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Response header configuration
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Download speed configuration
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Node cache key configuration
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Header cache configuration
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// Video dragging configuration
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// Cache expiration time configuration
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cross-border linkage optimization configuration
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// HTTPS acceleration configuration
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Timestamp hotlink protection configuration
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// SEO configuration
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Access protocol forced redirect configuration
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer hotlink protection configuration
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Browser cache configuration (This feature is in beta test and not fully available yet.)
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Domain name service type
	// web: static acceleration
	// download: download acceleration
	// media: streaming media VOD acceleration
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Specific configuration for region attributes
	// Applicable to use cases where the configuration of accelerating domain names inside mainland China is inconsistent with the configuration outside mainland China.
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// Domain name acceleration region
	// mainland: acceleration inside mainland China
	// overseas: acceleration outside mainland China
	// global: global acceleration
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *UpdateDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDomainConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDomainConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePayTypeRequest struct {
	*tchttp.BaseRequest

	// Billing region, which can be mainland or overseas.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Billing mode, which can be flux or bandwidth.
	PayType *string `json:"PayType,omitempty" name:"PayType"`
}

func (r *UpdatePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePayTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePayTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UrlRecord struct {

	// Status (disable: blocked; enable: unblocked)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Corresponding URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealUrl *string `json:"RealUrl,omitempty" name:"RealUrl"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type VideoSeek struct {

	// Video dragging switch
	// on: enabled
	// off: disabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ViolationUrl struct {

	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Origin access URL for violating resources
	RealUrl *string `json:"RealUrl,omitempty" name:"RealUrl"`

	// Snapshot path, which is used in the console to show the violating content snapshot.
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// Current status of violating resources
	// forbid: blocked
	// release: unblocked
	// delay: handling delayed
	// reject: appeal dismissed. The status is still blocked.
	// complain: appeal in process
	UrlStatus *string `json:"UrlStatus,omitempty" name:"UrlStatus"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}
