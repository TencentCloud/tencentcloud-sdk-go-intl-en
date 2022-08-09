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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessControl struct {
	// Whether to enable request header and request URL access control. Valid values: on, off
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Request header and request URL access rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AccessControlRules []*AccessControlRule `json:"AccessControlRules,omitempty" name:"AccessControlRules"`

	// Returns status code
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
}

type AccessControlRule struct {
	// requestHeader: access control over request header
	// url: access control over access URL
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Blocked content
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleContent *string `json:"RuleContent,omitempty" name:"RuleContent"`

	// on: regular match
	// off: exact match
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// This parameter is required only if `RuleType` is `requestHeader`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleHeader *string `json:"RuleHeader,omitempty" name:"RuleHeader"`
}

// Predefined struct for user
type AddCLSTopicDomainsRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Region configuration for domains
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type AddCLSTopicDomainsRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Region configuration for domains
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *AddCLSTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCLSTopicDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "DomainAreaConfigs")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCLSTopicDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCLSTopicDomainsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddCLSTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *AddCLSTopicDomainsResponseParams `json:"Response"`
}

func (r *AddCLSTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCLSTopicDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceCacheRule struct {
	// Rule types:
	// `all`: Apply to all files.
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	// `default`: the cache rules when the origin server has not returned max-age
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// Content for each CacheType:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// For `default`, enter "no max-age".
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents"`

	// Cache expiration time
	// Unit: second. The maximum value is 365 days.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type AdvanceConfig struct {
	// Advanced configuration name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether to support advanced configuration
	// `on`: Supported
	// `off`: Not supported
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type AdvanceHttps struct {
	// Custom TLS data switch
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CustomTlsStatus *string `json:"CustomTlsStatus,omitempty" name:"CustomTlsStatus"`

	// TLS version settings. Valid values: `TLSv1`, `TLSV1.1`, `TLSV1.2`, and `TLSv1.3`. Only consecutive versions can be enabled at the same time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TlsVersion []*string `json:"TlsVersion,omitempty" name:"TlsVersion"`

	// Custom encryption suite
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Cipher *string `json:"Cipher,omitempty" name:"Cipher"`

	// Origin-pull verification status
	// `off`: Disables origin-pull verification
	// `oneWay`: Only verify the origin
	// `twoWay`: Enables two-way origin-pull verification
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VerifyOriginType *string `json:"VerifyOriginType,omitempty" name:"VerifyOriginType"`

	// Configuration information of the origin-pull certificate
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CertInfo *ServerCert `json:"CertInfo,omitempty" name:"CertInfo"`

	// Configuration information of the origin server certificate
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	OriginCertInfo *ClientCert `json:"OriginCertInfo,omitempty" name:"OriginCertInfo"`
}

type AdvancedAuthentication struct {
	// Hotlink protection configuration switch (which can be on or off). If it is enabled, only one mode can and must be configured, while other modes are null.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Timestamp hotlink protection advanced configuration mode A
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TypeA *AdvancedAuthenticationTypeA `json:"TypeA,omitempty" name:"TypeA"`

	// Timestamp hotlink protection advanced configuration mode B
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TypeB *AdvancedAuthenticationTypeB `json:"TypeB,omitempty" name:"TypeB"`

	// Timestamp hotlink protection advanced configuration mode C
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TypeC *AdvancedAuthenticationTypeC `json:"TypeC,omitempty" name:"TypeC"`

	// Timestamp hotlink protection advanced configuration mode D
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TypeD *AdvancedAuthenticationTypeD `json:"TypeD,omitempty" name:"TypeD"`

	// Timestamp hotlink protection advanced configuration mode E
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TypeE *AdvancedAuthenticationTypeE `json:"TypeE,omitempty" name:"TypeE"`

	// Timestamp hotlink protection advanced configuration mode F
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TypeF *AdvancedAuthenticationTypeF `json:"TypeF,omitempty" name:"TypeF"`
}

type AdvancedAuthenticationTypeA struct {
	// Key used for signature calculation, allowing 6 to 32 bytes of letters and digits.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Time field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	TimeParam *string `json:"TimeParam,omitempty" name:"TimeParam"`

	// Expiration time in seconds
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Whether the expiration time parameter is required
	ExpireTimeRequired *bool `json:"ExpireTimeRequired,omitempty" name:"ExpireTimeRequired"`

	// URL composition, e.g., `${private_key}${schema}${host}${full_uri}`.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Time format. Valid values: dec (decimal), hex (hexadecimal).
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// Status code returned when the authentication failed
	FailCode *int64 `json:"FailCode,omitempty" name:"FailCode"`

	// Status code returned when the URL expired
	ExpireCode *int64 `json:"ExpireCode,omitempty" name:"ExpireCode"`

	// List of URLs to be authenticated
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// Reserved field
	Transformation *int64 `json:"Transformation,omitempty" name:"Transformation"`
}

type AdvancedAuthenticationTypeB struct {
	// Alpha key name
	KeyAlpha *string `json:"KeyAlpha,omitempty" name:"KeyAlpha"`

	// Beta key name
	KeyBeta *string `json:"KeyBeta,omitempty" name:"KeyBeta"`

	// Gamma key name
	KeyGamma *string `json:"KeyGamma,omitempty" name:"KeyGamma"`

	// Signature field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Time field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	TimeParam *string `json:"TimeParam,omitempty" name:"TimeParam"`

	// Expiration time in seconds
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Time format. Valid values: dec (decimal), hex (hexadecimal).
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// Status code returned when the authentication failed
	FailCode *int64 `json:"FailCode,omitempty" name:"FailCode"`

	// Status code returned when the URL expired
	ExpireCode *int64 `json:"ExpireCode,omitempty" name:"ExpireCode"`

	// List of URLs to be authenticated
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`
}

type AdvancedAuthenticationTypeC struct {
	// Access key
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// Authentication key
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type AdvancedAuthenticationTypeD struct {
	// Key used for signature calculation, allowing 6 to 32 bytes of letters and digits.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Alternative key used for authentication after the authentication key (`SecretKey`) failed
	BackupSecretKey *string `json:"BackupSecretKey,omitempty" name:"BackupSecretKey"`

	// Signature field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Time field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	TimeParam *string `json:"TimeParam,omitempty" name:"TimeParam"`

	// Expiration time in seconds
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Time format. Valid values: dec (decimal), hex (hexadecimal).
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
}

type AdvancedAuthenticationTypeE struct {
	// Key used for signature calculation, allowing 6 to 32 bytes of letters and digits.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// ACL signature field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AclSignParam *string `json:"AclSignParam,omitempty" name:"AclSignParam"`

	// Start time field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StartTimeParam *string `json:"StartTimeParam,omitempty" name:"StartTimeParam"`

	// Expiration time field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExpireTimeParam *string `json:"ExpireTimeParam,omitempty" name:"ExpireTimeParam"`

	// Time format (dec)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
}

type AdvancedAuthenticationTypeF struct {
	// Signature field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Time field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimeParam *string `json:"TimeParam,omitempty" name:"TimeParam"`

	// Transaction field name in the URI string, starting with a letter, and consisting of letters, digits, and underscores.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TransactionParam *string `json:"TransactionParam,omitempty" name:"TransactionParam"`

	// CMK used for signature calculation, allowing 6 to 32 bytes of letters and digits.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Alternative key used for signature calculation, which is used after the CMK fails in authentication. It allows 6 to 32 bytes of letters and digits.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BackupSecretKey *string `json:"BackupSecretKey,omitempty" name:"BackupSecretKey"`
}

type AdvancedCCRules struct {
	// Rule name
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Detection duration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DetectionTime *uint64 `json:"DetectionTime,omitempty" name:"DetectionTime"`

	// Detection frequency threshold
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FrequencyLimit *uint64 `json:"FrequencyLimit,omitempty" name:"FrequencyLimit"`

	// Whether to enable IP penalty. Valid values: `on` and `off`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PunishmentSwitch *string `json:"PunishmentSwitch,omitempty" name:"PunishmentSwitch"`

	// IP penalty duration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PunishmentTime *uint64 `json:"PunishmentTime,omitempty" name:"PunishmentTime"`

	// Action. Valid values: `intercept` and `redirect`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// A redirection URL used when Action is `redirect`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// Layer-7 rule configuration for CC frequency limiting
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Configure []*ScdnSevenLayerRules `json:"Configure,omitempty" name:"Configure"`
}

type AdvancedCache struct {
	// Cache expiration rule
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CacheRules []*AdvanceCacheRule `json:"CacheRules,omitempty" name:"CacheRules"`

	// Forced cache configuration
	// on: enabled
	// off: disabled
	// When this is enabled, if the origin server returns no-cache, no-store headers, node caching will still be performed according to the cache expiration rules.
	// This is disabled by default
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`

	// Whether to ignore the header and body on cache nodes if the origin server returns the header `Set-Cookie`.
	// `on`: Ignore; do not cache the header and body.
	// `off`: Do not ignore; follow the custom cache rules of cache nodes.
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitempty" name:"IgnoreSetCookie"`
}

type AdvancedScdnAclGroup struct {
	// Rule name
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Specific configurations
	Configure []*AdvancedScdnAclRule `json:"Configure,omitempty" name:"Configure"`

	// Action. Valid values: `intercept` and `redirect`.
	Result *string `json:"Result,omitempty" name:"Result"`

	// Whether the rule is activated. Valid values: `active` and `inactive`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error page configuration
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`
}

type AdvancedScdnAclRule struct {
	// Keyword. Valid values:
	// `protocol`: HTTP protocol
	// `httpVersion`: HTTP version
	// `method`: request method
	// `ip`: requester IP
	// `ipAsn`: ASN of the requester IP
	// `ipCountry`: country/region of the requester IP
	// `ipArea`: region of the requester IP
	// `xForwardFor`: X-Forward-For request header
	// `directory`: Path
	// `index`: Homepage
	// `path`: Full path of a file
	// `file`: File extension
	// `param`: Request parameter
	// `referer`: Referer request header
	// `cookie`: Cookie request header
	// `userAgent`: User-Agent request header
	// `head`: Custom request header
	MatchKey *string `json:"MatchKey,omitempty" name:"MatchKey"`

	// Logical operator. Valid values:
	// `exclude`: The keyword is not included
	// `include`: The keyword is included
	// `notequal`: Not the same as the keyword
	// `equal`: The same as the keyword
	// `matching`: The prefix is matched
	// `null`: Empty or does not exist
	LogicOperator *string `json:"LogicOperator,omitempty" name:"LogicOperator"`

	// Match value
	// When `MatchKey` is `protocol`,
	// Values: `HTTP` and `HTTPS`.
	// 
	// When `MatchKey` is `httpVersion`,
	// Values: `HTTP/1.0`, `HTTP/1.1`, `HTTP/1.2`, `HTTP/2`, and `HTTP/3`.
	// 
	// When `MatchKey` is `method`,
	// Values: `HEAD`, `GET`, `POST`, `PUT`, `OPTIONS`, `TRACE`, `DELETE`, `PATCH` and `CONNECT`.
	// 
	// When `MatchKey` is `ipCountry`, valid values include:
	// `OTHER`: Other areas
	// `VE`: Venezuela
	// `UY`: Uruguay
	// `SR`: Suriname
	// `PY`: Paraguay
	// `PE`: Peru
	// `GY`: Guyana
	// `EC`: Ecuador
	// `CO`: Colombia
	// `CL`: Chile
	// `BR`: Brazil
	// `BO`: Bolivia
	// `AR`: Argentina
	// `NZ`: New Zealand
	// `WS`: Samoa
	// `VU`: Vanuatu
	// `TV`: Tuvalu
	// `TO`: Tonga
	// `TK`: Tokelau
	// `PW`: Palau
	// `NU`: Niue
	// `NR`: Nauru
	// `KI`: Kiribati
	// `GU`: Guam
	// `FM`: Micronesia
	// `AU`: Australia
	// `US`: United States
	// `PR`: Puerto Rico
	// `DO`: Dominican Republic
	// `CR`: Costa Rica
	// `AS`: American Samoa
	// `AG`: Antigua and Barbuda
	// `PA`: Panama
	// `NI`: Nicaragua
	// `MX`: Mexico
	// `JM`: Jamaica
	// `HT`: Haiti
	// `HN`: Honduras
	// `GT`: Guatemala
	// `GP`: Guadeloupe
	// `GL`: Greenland
	// `GD`: Grenada
	// `CU`: Cuba
	// `CA`: Canada
	// `BZ`: Belize
	// `BS`: Bahamas
	// `BM`: Bermuda
	// `BB`: Barbados
	// `AW`: Aruba
	// `AI`: Anguilla
	// `VA`: Vatican
	// `SK`: Slovakia
	// `RU`: Russia
	// `GB`: United Kingdom
	// `CZ`: Czech Republic
	// `UA`: Ukraine
	// `TR`: Turkey
	// `SI`: Slovenia
	// `SE`: Sweden
	// `RS`: Republic of Serbia
	// `RO`: Romania
	// `PT`: Portugal
	// `PL`: Poland
	// `NO`: Norway
	// `NL`: Netherlands
	// `MT`: Malta
	// `MK`: Macedonia
	// `ME`: Montenegro
	// `MD`: Moldova
	// `MC`: Monaco
	// `LV`: Latvia
	// `LU`: Luxembourg
	// `LT`: Lithuania
	// `LI`: Liechtenstein
	// `KZ`: Kazakhstan
	// `IT`: Italy
	// `IS`: Iceland
	// `IE`: Ireland
	// `HU`: Hungary
	// `HR`: Croatia
	// `GR`: Greece
	// `GI`: Gibraltar
	// `GG`: Guernsey
	// `GE`: Georgia
	// `FR`: France
	// `FI`: Finland
	// `ES`: Spain
	// `EE`: Estonia
	// `DK`: Denmark
	// `DE`: Germany
	// `CY`: Cyprus
	// `CH`: Switzerland
	// `BY`: Belarus
	// `BG`: Bulgaria
	// `BE`: Belgium
	// `AZ`: Azerbaijan
	// `AT`: Austria
	// `AM`: Armenia
	// `AL`: Albania
	// `AD`: Andorra
	// `TL`: East Timor
	// `SY`: Syria
	// `SA`: Saudi Arabia
	// `PS`: Palestine
	// `LK`: Sri Lanka
	// `LK`: Sri Lanka
	// `KP`: North Korea
	// `KG`: Kyrgyzstan
	// `HK`: Hong Kong, China
	// `BN`: Brunei
	// `BD`: Bangladesh
	// `AE`: United Arab Emirates
	// `YE`: Yemen
	// `VN`: Vietnam
	// `UZ`: Uzbekistan
	// `TW`: Taiwan, China
	// `TM`: Turkmenistan
	// `TJ`: Tajikistan
	// `TH`: Thailand
	// `SG`: Singapore
	// `QA`: Qatar
	// `PK`: Pakistan
	// `PH`: Philippines
	// `OM`: Oman
	// `NP`: Nepal
	// `MY`: Malaysia
	// `MV`: Maldives
	// `MO`: Macao, China
	// `MN`: Mongolia
	// `MM`: Myanmar
	// `LB`: Lebanon
	// `KW`: Kuwait
	// `KR`: South Korea
	// `KH`: Cambodia
	// `JP`: Japan
	// `JO`: Jordan
	// `IR`: Iran
	// `IQ`: Iraq
	// `IN`: India
	// `IL`: Israel
	// `ID`: Indonesia
	// `CN`: China
	// `BT`: Bhutan
	// `BH`: Bahrain
	// `AF`: Afghanistan
	// `LY`: Libya
	// `CD`: Democratic Republic of the Congo
	// `RE`: La Réunion
	// `SZ`: Swaziland
	// `ZW`: Zimbabwe
	// `ZM`: Zambia
	// `YT`: Mayotte
	// `UG`: Uganda
	// `TZ`: Tanzania
	// `TN`: Tunisia
	// `TG`: Togo
	// `TD`: Chad
	// `SO`: Somalia
	// `SN`: Senegal
	// `SD`: Sudan
	// `SC`: Seychelles
	// `RW`: Rwanda
	// `NG`: Nigeria
	// `NE`: Niger
	// `NA`: Namibia
	// `MZ`: Mozambique
	// `MW`: Malawi
	// `MU`: Mauritius
	// `MR`: Mauritania
	// `ML`: Mali
	// `MG`: Madagascar
	// `MA`: Morocco
	// `LS`: Lesotho
	// `LR`: Liberia
	// `KM`: Comoros
	// `KE`: Kenya
	// `GN`: Guinea
	// `GM`: Gambia
	// `GH`: Ghana
	// `GA`: Gabon
	// `ET`: Ethiopia
	// `ER`: Eritrea
	// `EG`: Egypt
	// `DZ`: Algeria
	// `DJ`: Djibouti
	// `CM`: Cameroon
	// `CG`: Republic of the Congo
	// `BW`: Botswana
	// `BJ`: Benin
	// `BI`: Burundi
	// `AO`: Angola
	// 
	// When MatchKey is `ipArea`, valid values include:
	// `OTHER`: other areas
	// `AS`: Asia
	// `EU`: Europe
	// `AN`: Antarctica
	// `AF`: Africa
	// `OC`: Oceania
	// `NA`: North America
	// `SA`: South America
	// 
	// When MatchKey is `index`,
	// valid value is `/;/index.html`.
	MatchValue []*string `json:"MatchValue,omitempty" name:"MatchValue"`

	// Whether to distinguish uppercase or lowercase letters. `true`: case sensitive; `false`: case insensitive.
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// This field is required when `MatchKey` is `param` or `cookie`. For `param`, it indicates a key value of the request parameter if MatchKey is `param`, while a key value of the Cookie request header if MatchKey is `cookie`.
	MatchKeyParam *string `json:"MatchKeyParam,omitempty" name:"MatchKeyParam"`
}

type Authentication struct {
	// Hotlink protection configuration switch
	// `on`: Enable
	// `off`: Disable
	// When this is enabled, one mode needs to be configured. Other modes need to be set to null.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Timestamp hotlink protection mode A configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TypeA *AuthenticationTypeA `json:"TypeA,omitempty" name:"TypeA"`

	// Timestamp hotlink protection mode B configuration (mode B is being upgraded and is currently not supported)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TypeB *AuthenticationTypeB `json:"TypeB,omitempty" name:"TypeB"`

	// Timestamp hotlink protection mode C configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TypeC *AuthenticationTypeC `json:"TypeC,omitempty" name:"TypeC"`

	// Timestamp hotlink protection mode D configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TypeD *AuthenticationTypeD `json:"TypeD,omitempty" name:"TypeD"`
}

type AuthenticationTypeA struct {
	// The key for signature calculation
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature parameter name
	// Only upper and lower-case letters, digits, and underscores (_) are allowed. It cannot start with a digit. Length limit: 1-100 characters.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Signature expiration time
	// Unit: second. The maximum value is 630720000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings determining if authentication should be performed
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions"`

	// `whitelist`: All file types apart from the FileExtensions list are authenticated.
	// `blacklist`: Only the file types in the FileExtensions list are authenticated.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Backup key, which is used to calculate a signature.
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackupSecretKey *string `json:"BackupSecretKey,omitempty" name:"BackupSecretKey"`
}

type AuthenticationTypeB struct {
	// The key for signature calculation
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature expiration time
	// Unit: second. The maximum value is 630720000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings determining if authentication should be performed
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions"`

	// whitelist: indicates that all file types apart from the FileExtensions list are authenticated
	// blacklist: indicates that only the file types in the FileExtensions list are authenticated
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Backup key, which is used to calculate a signature.
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackupSecretKey *string `json:"BackupSecretKey,omitempty" name:"BackupSecretKey"`
}

type AuthenticationTypeC struct {
	// The key for signature calculation
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature expiration time
	// Unit: second. The maximum value is 630720000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings determining if authentication should be performed
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions"`

	// `whitelist`: All file types apart from the FileExtensions list are authenticated.
	// `blacklist`: Only the file types in the FileExtensions list are authenticated.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Timestamp settings
	// `dec`: Decimal
	// `hex`: Hexadecimal
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// Backup key, which is used to calculate a signature.
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackupSecretKey *string `json:"BackupSecretKey,omitempty" name:"BackupSecretKey"`
}

type AuthenticationTypeD struct {
	// The key for signature calculation
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Signature expiration time
	// Unit: second. The maximum value is 630720000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// File extension list settings determining if authentication should be performed
	// If it contains an asterisk (*), this indicates all files.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions"`

	// `whitelist`: All file types apart from the FileExtensions list are authenticated.
	// `blacklist`: Only the file types in the FileExtensions list are authenticated.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Signature parameter name
	// Only upper and lower-case letters, digits, and underscores (_) are allowed. It cannot start with a digit. Length limit: 1-100 characters.
	SignParam *string `json:"SignParam,omitempty" name:"SignParam"`

	// Timestamp parameter name
	// Only upper and lower-case letters, digits, and underscores (_) are allowed. It cannot start with a digit. Length limit: 1-100 characters.
	TimeParam *string `json:"TimeParam,omitempty" name:"TimeParam"`

	// Timestamp settings
	// `dec`: Decimal
	// `hex`: Hexadecimal
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// Backup key, which is used to calculate a signature.
	// 6-32 characters. Only digits and letters are allowed. 
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackupSecretKey *string `json:"BackupSecretKey,omitempty" name:"BackupSecretKey"`
}

type AwsPrivateAccess struct {
	// Switch, which can be set to on or off.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Access ID.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// Key.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Region.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// BucketName
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

type BandwidthAlert struct {
	// Specifies whether to enable the bandwidth cap
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// The upper limit of bandwidth usage (in bps) or traffic usage (in bytes).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BpsThreshold *int64 `json:"BpsThreshold,omitempty" name:"BpsThreshold"`

	// Action taken when threshold is reached
	// `RESOLVE_DNS_TO_ORIGIN`: Requests will be forwarded to the origin server. This is only supported for domain names of external origin.
	// `RETURN_404`: A 404 error will be returned for all requests.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CounterMeasure *string `json:"CounterMeasure,omitempty" name:"CounterMeasure"`

	// The last time when the usage upper limit in the Chinese mainland was reached
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" name:"LastTriggerTime"`

	// Indicates whether to trigger alerts when the upper limit is reached
	// `on`: Enable
	// `off`: Disable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AlertSwitch *string `json:"AlertSwitch,omitempty" name:"AlertSwitch"`

	// Triggers alarms when the ratio of bandwidth or traffic usage to the usage upper limit reaches the specified value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AlertPercentage *int64 `json:"AlertPercentage,omitempty" name:"AlertPercentage"`

	// The last time when the usage outside the Chinese mainland reached the upper limit
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LastTriggerTimeOverseas *string `json:"LastTriggerTimeOverseas,omitempty" name:"LastTriggerTimeOverseas"`

	// Dimension of the usage limit
	// `bandwidth`: Bandwidth
	// `flux`: Traffic
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Usage limit configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatisticItems []*StatisticItem `json:"StatisticItems,omitempty" name:"StatisticItems"`
}

type BotCookie struct {
	// Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Rule type, which can only be `all` currently.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Rule value. Valid value: `*`.
	RuleValue []*string `json:"RuleValue,omitempty" name:"RuleValue"`

	// Action. Valid values: `monitor`, `intercept`, `redirect`, and `captcha`.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Redirection target page
	// Note: This field may return null, indicating that no valid values can be obtained.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type BotJavaScript struct {
	// Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Rule type, which can only be `file` currently.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Rule value. Valid values: `html` and `htm`.
	RuleValue []*string `json:"RuleValue,omitempty" name:"RuleValue"`

	// Action. Valid values: `monitor`, `intercept`, `redirect`, and `captcha`.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Redirection target page
	// Note: This field may return null, indicating that no valid values can be obtained.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type BriefDomain struct {
	// Domain name ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tencent Cloud account ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Acceleration domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// CNAME address of domain name
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Acceleration service status
	// `rejected`: The domain name is rejected due to expiration/deregistration of its ICP filing
	// `processing`: Deploying
	// `online`: Activated
	// `offline`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`

	// Project ID, which can be viewed on the Tencent Cloud project management page
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Domain name service type
	// `web`: Static acceleration
	// `download`: Download acceleration
	// `media`: Streaming media VOD acceleration
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Domain name creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Domain name update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin server configuration details.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Domain name block status
	// `normal`: Normal
	// `overdue`: The domain name has been disabled due to account arrears. The acceleration service can be resumed after the account is topped up.
	// `malicious`: The acceleration service has been forcibly disabled due to detection of malicious behavior.
	// `ddos`: The acceleration service has been disabled due to large-scale DDoS attacks to the domain name
	// `idle`: No operations or data has been detected for more than 90 days. The domain name is determined to be inactive which automatically disables the acceleration service. You can restart the service.
	// `unlicensed`: The acceleration service has been automatically disabled as the domain name has no ICP filing or its ICP filing is deregistered. Service can be resumed after an ICP filing is obtained.
	// `capping`: The configured upper limit for bandwidth has been reached.
	// `readonly`: The domain name has a special configuration and has been locked.
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// Acceleration region
	// `mainland`: Acceleration inside the Chinese mainland
	// `overseas`: Acceleration outside the Chinese mainland
	// `global`: Acceleration over the globe
	Area *string `json:"Area,omitempty" name:"Area"`

	// Domain name lock status
	// `normal`: Not locked
	// `mainland`: Locked in the Chinese mainland
	// overseas: Locked outside the Chinese mainland
	// global: Locked globally
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`

	// Product of the domain name, either `cdn` or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Primary domain name
	ParentHost *string `json:"ParentHost,omitempty" name:"ParentHost"`
}

type Cache struct {
	// Basic cache expiration time configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SimpleCache *SimpleCache `json:"SimpleCache,omitempty" name:"SimpleCache"`

	// (Disused) Advanced cache validity configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AdvancedCache *AdvancedCache `json:"AdvancedCache,omitempty" name:"AdvancedCache"`

	// Advanced path cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleCache []*RuleCache `json:"RuleCache,omitempty" name:"RuleCache"`
}

type CacheConfig struct {
	// `on`: Enable
	// `off`: Disable
	HeuristicCacheTimeSwitch *string `json:"HeuristicCacheTimeSwitch,omitempty" name:"HeuristicCacheTimeSwitch"`

	// Unit: Second
	HeuristicCacheTime *int64 `json:"HeuristicCacheTime,omitempty" name:"HeuristicCacheTime"`
}

type CacheConfigCache struct {
	// Cache configuration switch
	// on: enable
	// off: disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Cache expiration time settings
	// Unit: second. The maximum value is 365 days.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`

	// Advanced cache expiration configuration. If this is enabled, the max-age value returned by the origin server will be compared with the cache expiration time set in CacheRules, and the smallest value will be cached on the node.
	// `on`: Enable
	// `off`: Disable
	// This is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CompareMaxAge *string `json:"CompareMaxAge,omitempty" name:"CompareMaxAge"`

	// Force cache
	// `on`: Enable
	// `off`: Disable
	// This is disabled by default. If enabled, the `no-store` and `no-cache` resources returned from the origin server will be cached according to `CacheRules` rules.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`

	// Whether to ignore the header and body on cache nodes if the origin server returns the header `Set-Cookie`.
	// `on`: Ignore; do not cache the header and body.
	// `off`: Do not ignore; follow the custom cache rules of cache nodes.
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitempty" name:"IgnoreSetCookie"`
}

type CacheConfigFollowOrigin struct {
	// Follow origin server switch configuration
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Heuristic cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HeuristicCache *HeuristicCache `json:"HeuristicCache,omitempty" name:"HeuristicCache"`
}

type CacheConfigNoCache struct {
	// No-cache configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: this field may return null, indicating that no valid value is obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Always forwards to the origin server for verification
	// `on`: Enable
	// `off`: Disable
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Revalidate *string `json:"Revalidate,omitempty" name:"Revalidate"`
}

type CacheKey struct {
	// Whether to enable full-path cache
	// `on`: Enable full-path cache (i.e., disable Ignore Query String)
	// `off`: Disable full-path cache (i.e., enable Ignore Query String)
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`

	// Specifies whether the cache key is case sensitive
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreCase *string `json:"IgnoreCase,omitempty" name:"IgnoreCase"`

	// Request parameter contained in `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	QueryString *QueryStringKey `json:"QueryString,omitempty" name:"QueryString"`

	// Cookie contained in `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cookie *CookieKey `json:"Cookie,omitempty" name:"Cookie"`

	// Request header contained in `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Header *HeaderKey `json:"Header,omitempty" name:"Header"`

	// Custom string contained in `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheTag *CacheTagKey `json:"CacheTag,omitempty" name:"CacheTag"`

	// Request protocol contained in `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Scheme *SchemeKey `json:"Scheme,omitempty" name:"Scheme"`

	// Path-specific cache key configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	KeyRules []*KeyRule `json:"KeyRules,omitempty" name:"KeyRules"`
}

type CacheOptResult struct {
	// List of succeeded URLs
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SuccessUrls []*string `json:"SuccessUrls,omitempty" name:"SuccessUrls"`

	// List of failed URLs
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FailUrls []*string `json:"FailUrls,omitempty" name:"FailUrls"`
}

type CacheTagKey struct {
	// Whether to use `CacheTag` as part of `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Value of custom `CacheTag`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CappingRule struct {
	// Rule types:
	// `all`: Apply to all files.
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each `RuleType`: 
	// For `all`, enter a wildcard `*`.
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// Downstream speed value settings (in KB/s)
	KBpsThreshold *int64 `json:"KBpsThreshold,omitempty" name:"KBpsThreshold"`
}

type CdnData struct {
	// Queries by the specified metric:
	// `flux`: Traffic (in bytes)
	// `bandwidth`: Bandwidth (in bps)
	// `request`: Number of requests
	// `fluxHitRate`: Traffic hit rate (in %)
	// `statusCode`: Status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx status codes will be returned (in entries)
	// `2XX`: Returns the aggregate list of 2xx status codes and the data for status codes starting with 2 (in entries)
	// `3XX`: Returns the aggregate list of 3xx status codes and the data for status codes starting with 3 (in entries)
	// `4XX`: Returns the aggregate list of 4xx status codes and the data for status codes starting with 4 (in entries)
	// `5XX`: Returns the aggregate list of 5xx status codes and the data for status codes starting with 5 (in entries)
	// You can also specify a status code for querying.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Detailed data combination
	DetailData []*TimestampData `json:"DetailData,omitempty" name:"DetailData"`

	// Aggregate data combination
	SummarizedData *SummarizedData `json:"SummarizedData,omitempty" name:"SummarizedData"`
}

type CdnIp struct {
	// IP to be queried
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// IP ownership:
	// yes: Tencent Cloud CDN node
	// no: non-Tencent Cloud CDN node
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Node district/country
	// unknown: unknown node location
	Location *string `json:"Location,omitempty" name:"Location"`

	// Activation and deactivation history of the node.
	History []*CdnIpHistory `json:"History,omitempty" name:"History"`

	// Node region
	// `mainland`: Acceleration nodes inside the Chinese mainland
	// `overseas`: Acceleration nodes outside the Chinese mainland
	// `unknown`: Service region unknown
	Area *string `json:"Area,omitempty" name:"Area"`

	// City where the nodes reside
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	City *string `json:"City,omitempty" name:"City"`
}

type CdnIpHistory struct {
	// Operation type
	// `online`: Nodes activated
	// `offline`: Nodes deactivated
	Status *string `json:"Status,omitempty" name:"Status"`

	// Operation time corresponding to operation type
	// If its value is `null`, it means there is no status change record.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Datetime *string `json:"Datetime,omitempty" name:"Datetime"`
}

type ClientCert struct {
	// Client certificate
	// PEM format, requires Base64 encoding.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Client certificate name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Certificate expiration time
	// When this is used as an input parameter, it can be left blank.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Certificate issuance time
	// When this is used as an input parameter, it can be left blank.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`
}

type ClsLogObject struct {
	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Log time
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// Log content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Capture path
	Filename *string `json:"Filename,omitempty" name:"Filename"`

	// Log source device
	Source *string `json:"Source,omitempty" name:"Source"`
}

type ClsSearchLogs struct {
	// Cursor for more search results
	Context *string `json:"Context,omitempty" name:"Context"`

	// Whether all search results have been returned
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// Log content information
	Results []*ClsLogObject `json:"Results,omitempty" name:"Results"`
}

type Compatibility struct {
	// Compatibility flag status code.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type Compression struct {
	// Smart compression configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Compression rules array
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CompressionRules []*CompressionRule `json:"CompressionRules,omitempty" name:"CompressionRules"`
}

type CompressionRule struct {
	// true: must be set as true, enables compression
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Compress *bool `json:"Compress,omitempty" name:"Compress"`

	// The minimum file size to trigger compression (in bytes)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MinLength *int64 `json:"MinLength,omitempty" name:"MinLength"`

	// The maximum file size to trigger compression (in bytes).
	// The maximum value is 30 MB.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxLength *int64 `json:"MaxLength,omitempty" name:"MaxLength"`

	// File compression algorithm
	// `gzip`: Gzip compression
	// `brotli`: Brotli compression
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Algorithms []*string `json:"Algorithms,omitempty" name:"Algorithms"`

	// Compress based on file suffix.
	// File suffixes such as jpg and txt are supported.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FileExtensions []*string `json:"FileExtensions,omitempty" name:"FileExtensions"`

	// Rule types:
	// `all`: Apply to all files.
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	// `contentType`: Apply when the `ContentType` is specified.
	// If this field is specified, `FileExtensions` does not take effect.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each `CacheType`:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// For `contentType`, enter `text/html`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`
}

type CookieKey struct {
	// Whether to use `Cookie` as part of `CacheKey`. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Used cookies (separated by ';')
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type CreateClsLogTopicRequestParams struct {
	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Domain name region information
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs"`
}

type CreateClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Domain name region information
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs"`
}

func (r *CreateClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "LogsetId")
	delete(f, "Channel")
	delete(f, "DomainAreaConfigs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClsLogTopicResponseParams struct {
	// Topic ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateClsLogTopicResponseParams `json:"Response"`
}

func (r *CreateClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnFailedLogTaskRequestParams struct {
	// ID of the failed task to retry
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Region. Valid values: `mainland` and `overseas`.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CreateScdnFailedLogTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the failed task to retry
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Region. Valid values: `mainland` and `overseas`.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateScdnFailedLogTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnFailedLogTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScdnFailedLogTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnFailedLogTaskResponseParams struct {
	// Creation result. 
	// 0: Creation succeeded
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateScdnFailedLogTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateScdnFailedLogTaskResponseParams `json:"Response"`
}

func (r *CreateScdnFailedLogTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnFailedLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCdnDomainRequestParams struct {
	// Domain name
	// The domain name status should be `Disabled`
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DeleteCdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain name
	// The domain name status should be `Disabled`
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCdnDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCdnDomainResponseParams `json:"Response"`
}

func (r *DeleteCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClsLogTopicRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type DeleteClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *DeleteClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "LogsetId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClsLogTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClsLogTopicResponseParams `json:"Response"`
}

func (r *DeleteClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataRequestParams struct {
	// Start time of the query, e.g., 2018-09-04 10:40:00.
	// The specified start time will be rounded down based on the granularity parameter `Interval`. For example, if you set the start time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query, e.g. 2018-09-04 10:40:00.
	// The specified end time will be rounded down based on the granularity parameter `Interval`. For example, if you set the end time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Time granularity, which can be:
	// `min`: 1-minute granularity. The query period cannot exceed 24 hours.
	// `5min`: 5-minute granularity. The query range cannot exceed 31 days.
	// `hour`: 1-hour granularity. The query period cannot exceed 31 days.
	// `day`: 1-day granularity. The query period cannot exceed 31 days.
	// 
	// `min` is not supported if the `Area` field is `overseas`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Domain name whose billing data is to be queried
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Specifies the project ID to be queried. [Check project ID in the console](https://console.cloud.tencent.com/project)
	// If the `Domain` parameter is passed in, the `Proejct` parameter is ignored. Only the billing data of the specified domain name is returned. 
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Acceleration region whose billing data is to be queried:
	// `mainland`: Regions within the Chinese mainland
	// `overseas`: Regions outside the Chinese mainland
	// If this parameter is left empty, `mainland` will be used by default
	Area *string `json:"Area,omitempty" name:"Area"`

	// Country/region to be queried if `Area` is `overseas`
	// To view codes of provinces or countries/regions, see [Province Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E7.9C.81.E4.BB.BD.E6.98.A0.E5.B0.84)
	// If this parameter is left empty, all countries/regions will be queried
	District *int64 `json:"District,omitempty" name:"District"`

	// Billing statistics type
	// `flux`: Bill by traffic
	// `bandwidth`: Bill by bandwidth
	// Default value: `bandwidth`
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`


	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query, e.g., 2018-09-04 10:40:00.
	// The specified start time will be rounded down based on the granularity parameter `Interval`. For example, if you set the start time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query, e.g. 2018-09-04 10:40:00.
	// The specified end time will be rounded down based on the granularity parameter `Interval`. For example, if you set the end time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Time granularity, which can be:
	// `min`: 1-minute granularity. The query period cannot exceed 24 hours.
	// `5min`: 5-minute granularity. The query range cannot exceed 31 days.
	// `hour`: 1-hour granularity. The query period cannot exceed 31 days.
	// `day`: 1-day granularity. The query period cannot exceed 31 days.
	// 
	// `min` is not supported if the `Area` field is `overseas`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Domain name whose billing data is to be queried
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Specifies the project ID to be queried. [Check project ID in the console](https://console.cloud.tencent.com/project)
	// If the `Domain` parameter is passed in, the `Proejct` parameter is ignored. Only the billing data of the specified domain name is returned. 
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Acceleration region whose billing data is to be queried:
	// `mainland`: Regions within the Chinese mainland
	// `overseas`: Regions outside the Chinese mainland
	// If this parameter is left empty, `mainland` will be used by default
	Area *string `json:"Area,omitempty" name:"Area"`

	// Country/region to be queried if `Area` is `overseas`
	// To view codes of provinces or countries/regions, see [Province Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E7.9C.81.E4.BB.BD.E6.98.A0.E5.B0.84)
	// If this parameter is left empty, all countries/regions will be queried
	District *int64 `json:"District,omitempty" name:"District"`

	// Billing statistics type
	// `flux`: Bill by traffic
	// `bandwidth`: Bill by bandwidth
	// Default value: `bandwidth`
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`

	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
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
	delete(f, "Domain")
	delete(f, "Project")
	delete(f, "Area")
	delete(f, "District")
	delete(f, "Metric")
	delete(f, "Product")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataResponseParams struct {
	// Time granularity, which is specified by the parameter passed in during the query:
	// `min`: 1 minute
	// `5min`: 5 minutes
	// `hour`: 1 hour
	// `day`: 1 day
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Data details
	Data []*ResourceBillingData `json:"Data,omitempty" name:"Data"`

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
type DescribeCdnDataRequestParams struct {
	// Start time of the query, e.g., 2018-09-04 10:40:00.
	// The specified start time will be rounded down based on the granularity parameter `Interval`. For example, if you set the start time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query, e.g. 2018-09-04 10:40:00.
	// The specified end time will be rounded down based on the granularity parameter `Interval`. For example, if you set the end time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the metric to query, which can be:
	// `flux`: Traffic (in bytes)
	// `fluxIn`: Upstream traffic (in bytes), only used for the `ecdn` product
	// `fluxOut`: Downstream traffic (in bytes), only used for the `ecdn` product
	// `bandwidth`: Bandwidth (in bps)
	// `bandwidthIn`: Upstream bandwidth (in bps), only used for the `ecdn` product
	// `bandwidthOut`: Downstream bandwidth (in bps), only used for the `ecdn` product
	// `request`: Number of requests
	// `hitRequest`: Number of hit requests
	// `requestHitRate`: Request hit rate (in % with two decimal digits)
	// `hitFlux`: Hit traffic (in bytes)
	// `fluxHitRate`: Traffic hit rate (in % with two decimal digits)
	// `statusCode`: Status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx status codes will be returned (in entries)
	// `2xx`: Returns the aggregate list of 2xx status codes and the data for status codes starting with 2 (in entries)
	// `3xx`: Returns the aggregate list of 3xx status codes and the data for status codes starting with 3 (in entries)
	// `4xx`: Returns the aggregate list of 4xx status codes and the data for status codes starting with 4 (in entries)
	// `5xx`: Returns the aggregate list of 5xx status codes and the data for status codes starting with 5 (in entries)
	// Specifies the status code to query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the list of domain names to be queried
	// You can specify one or more domain names.
	// Up to 30 domain names can be queried in one request.
	// If this parameter is not specified, it means to query all domain names under the current account.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried. [Check project ID in the console](https://console.cloud.tencent.com/project)
	// Note that `Project` will be ignored if `Domains` is specified.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Sampling interval. The available options vary for different query period. See below: 
	// `min`: Return data with 1-minute granularity. It’s available when the query period is  within 24 hours and `Area` is `mainland`.
	// `5min`: Return data with 5-minute granularity. It’s available when the query period is within 31 days.
	// `hour`: Return data with 1-hour granularity. It’s available when the query period is within 31 days.
	// `day`: Return data with 1-day granularity. It’s available when the query period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) during a multi-domain-name query.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported).
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies an ISP when you query the CDN data within the Chinese mainland. If this is left blank, all ISPs will be queried.
	// To view ISP codes, see [ISP Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// Note that only one of `District`, `Isp` and `IpProtocol` can be specified.
	Isp *int64 `json:"Isp,omitempty" name:"Isp"`

	// Specifies a province when you query the CDN data within the Chinese mainland. If this is left blank, all provinces will be queried.
	// Specifies a country/region when you query the CDN data outside the Chinese mainland. If this is left blank, all countries/regions will be queried.
	// To view codes of provinces or countries/regions, see [Province Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8).
	// When `Area` is `mainland`, you can query by the province. Note that only one of `District`, `Isp` and `IpProtocol` can be specified.
	District *int64 `json:"District,omitempty" name:"District"`

	// Specifies the protocol to be queried; if you leave it blank, all protocols will be queried.
	// `all`: All protocols
	// `http`: Query HTTP data
	// `https`: Query HTTPS data
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Specifies the data source to be queried. It’s only open to beta users now. 
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Specifies the IP protocol to be queried. If it’s not specified, data of all IP protocols are returned.
	// `all`: All protocols
	// `ipv4`: Query IPv4 data
	// `ipv6`: Query IPv6 data
	// If `IpProtocol` is specified, `District` parameter can not be specified at the same time.
	// Note: `ipv4` and `ipv6` are only available to beta users. 
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Specifies the service area. If it’s not specified, CDN data of the Chinese mainland are returned.
	// `mainland`: Query CDN data in the Chinese mainland.
	// `overseas`: Query CDN data outside the Chinese mainland.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specify whether to query by the region of the server or client. This parameter is valid only when `Area` is `overseas`.
	// `server`: Query by the location of server (Tencent Cloud CDN nodes)
	// `client`: Query by the location of the client (where the request devices are located)
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Specifies a time zone to query. The default time zone is UTC+08:00.
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

type DescribeCdnDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query, e.g., 2018-09-04 10:40:00.
	// The specified start time will be rounded down based on the granularity parameter `Interval`. For example, if you set the start time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query, e.g. 2018-09-04 10:40:00.
	// The specified end time will be rounded down based on the granularity parameter `Interval`. For example, if you set the end time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the metric to query, which can be:
	// `flux`: Traffic (in bytes)
	// `fluxIn`: Upstream traffic (in bytes), only used for the `ecdn` product
	// `fluxOut`: Downstream traffic (in bytes), only used for the `ecdn` product
	// `bandwidth`: Bandwidth (in bps)
	// `bandwidthIn`: Upstream bandwidth (in bps), only used for the `ecdn` product
	// `bandwidthOut`: Downstream bandwidth (in bps), only used for the `ecdn` product
	// `request`: Number of requests
	// `hitRequest`: Number of hit requests
	// `requestHitRate`: Request hit rate (in % with two decimal digits)
	// `hitFlux`: Hit traffic (in bytes)
	// `fluxHitRate`: Traffic hit rate (in % with two decimal digits)
	// `statusCode`: Status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx status codes will be returned (in entries)
	// `2xx`: Returns the aggregate list of 2xx status codes and the data for status codes starting with 2 (in entries)
	// `3xx`: Returns the aggregate list of 3xx status codes and the data for status codes starting with 3 (in entries)
	// `4xx`: Returns the aggregate list of 4xx status codes and the data for status codes starting with 4 (in entries)
	// `5xx`: Returns the aggregate list of 5xx status codes and the data for status codes starting with 5 (in entries)
	// Specifies the status code to query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the list of domain names to be queried
	// You can specify one or more domain names.
	// Up to 30 domain names can be queried in one request.
	// If this parameter is not specified, it means to query all domain names under the current account.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried. [Check project ID in the console](https://console.cloud.tencent.com/project)
	// Note that `Project` will be ignored if `Domains` is specified.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Sampling interval. The available options vary for different query period. See below: 
	// `min`: Return data with 1-minute granularity. It’s available when the query period is  within 24 hours and `Area` is `mainland`.
	// `5min`: Return data with 5-minute granularity. It’s available when the query period is within 31 days.
	// `hour`: Return data with 1-hour granularity. It’s available when the query period is within 31 days.
	// `day`: Return data with 1-day granularity. It’s available when the query period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) during a multi-domain-name query.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported).
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies an ISP when you query the CDN data within the Chinese mainland. If this is left blank, all ISPs will be queried.
	// To view ISP codes, see [ISP Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// Note that only one of `District`, `Isp` and `IpProtocol` can be specified.
	Isp *int64 `json:"Isp,omitempty" name:"Isp"`

	// Specifies a province when you query the CDN data within the Chinese mainland. If this is left blank, all provinces will be queried.
	// Specifies a country/region when you query the CDN data outside the Chinese mainland. If this is left blank, all countries/regions will be queried.
	// To view codes of provinces or countries/regions, see [Province Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8).
	// When `Area` is `mainland`, you can query by the province. Note that only one of `District`, `Isp` and `IpProtocol` can be specified.
	District *int64 `json:"District,omitempty" name:"District"`

	// Specifies the protocol to be queried; if you leave it blank, all protocols will be queried.
	// `all`: All protocols
	// `http`: Query HTTP data
	// `https`: Query HTTPS data
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Specifies the data source to be queried. It’s only open to beta users now. 
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Specifies the IP protocol to be queried. If it’s not specified, data of all IP protocols are returned.
	// `all`: All protocols
	// `ipv4`: Query IPv4 data
	// `ipv6`: Query IPv6 data
	// If `IpProtocol` is specified, `District` parameter can not be specified at the same time.
	// Note: `ipv4` and `ipv6` are only available to beta users. 
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Specifies the service area. If it’s not specified, CDN data of the Chinese mainland are returned.
	// `mainland`: Query CDN data in the Chinese mainland.
	// `overseas`: Query CDN data outside the Chinese mainland.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specify whether to query by the region of the server or client. This parameter is valid only when `Area` is `overseas`.
	// `server`: Query by the location of server (Tencent Cloud CDN nodes)
	// `client`: Query by the location of the client (where the request devices are located)
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Specifies a time zone to query. The default time zone is UTC+08:00.
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

func (r *DescribeCdnDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	delete(f, "Detail")
	delete(f, "Isp")
	delete(f, "District")
	delete(f, "Protocol")
	delete(f, "DataSource")
	delete(f, "IpProtocol")
	delete(f, "Area")
	delete(f, "AreaType")
	delete(f, "Product")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnDataResponseParams struct {
	// Time granularity of the returned data. 
	// `min`: 1 minute
	// `5min`: 5 minutes
	// `hour`: 1 hour
	// `day`: 1 day
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Returned data details of the specified conditional query
	Data []*ResourceData `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCdnDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnDataResponseParams `json:"Response"`
}

func (r *DescribeCdnDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnDomainLogsRequestParams struct {
	// Specifies a domain name for the query
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Starting time, such as `2019-09-04 00:00:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as `2019-09-04 12:00:00`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1,000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a region for the query.
	// `mainland`: specifies to return the download link of logs on acceleration within Mainland China;
	// `overseas`: specifies to return the download link of logs on acceleration outside Mainland China;
	// `global`: specifies to return a download link of logs on acceleration within Mainland China and a link of logs on acceleration outside Mainland China.
	// Default value: `mainland`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// The type of log to be downloaded.
	// access: access logs
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type DescribeCdnDomainLogsRequest struct {
	*tchttp.BaseRequest
	
	// Specifies a domain name for the query
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Starting time, such as `2019-09-04 00:00:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as `2019-09-04 12:00:00`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1,000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a region for the query.
	// `mainland`: specifies to return the download link of logs on acceleration within Mainland China;
	// `overseas`: specifies to return the download link of logs on acceleration outside Mainland China;
	// `global`: specifies to return a download link of logs on acceleration within Mainland China and a link of logs on acceleration outside Mainland China.
	// Default value: `mainland`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// The type of log to be downloaded.
	// access: access logs
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

func (r *DescribeCdnDomainLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDomainLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnDomainLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnDomainLogsResponseParams struct {
	// Download link of the log package
	DomainLogs []*DomainLog `json:"DomainLogs,omitempty" name:"DomainLogs"`

	// Total number of entries obtained
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCdnDomainLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnDomainLogsResponseParams `json:"Response"`
}

func (r *DescribeCdnDomainLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDomainLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnIpRequestParams struct {
	// List of IPs to be queried
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

type DescribeCdnIpRequest struct {
	*tchttp.BaseRequest
	
	// List of IPs to be queried
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

func (r *DescribeCdnIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnIpResponseParams struct {
	// Node ownership details
	Ips []*CdnIp `json:"Ips,omitempty" name:"Ips"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCdnIpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnIpResponseParams `json:"Response"`
}

func (r *DescribeCdnIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnOriginIpRequestParams struct {

}

type DescribeCdnOriginIpRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCdnOriginIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnOriginIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnOriginIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnOriginIpResponseParams struct {
	// Intermediate node IP details
	Ips []*OriginIp `json:"Ips,omitempty" name:"Ips"`

	// Number of intermediate node IPs
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCdnOriginIpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnOriginIpResponseParams `json:"Response"`
}

func (r *DescribeCdnOriginIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnOriginIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertDomainsRequestParams struct {
	// Base64-encoded string of certificate in PEM format
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// Managed certificate ID. `Cert` and `CertId` cannot be both empty. If they’re both filled in, `CerID` prevails.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Product of the domain name, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeCertDomainsRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded string of certificate in PEM format
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// Managed certificate ID. `Cert` and `CertId` cannot be both empty. If they’re both filled in, `CerID` prevails.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Product of the domain name, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeCertDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cert")
	delete(f, "CertId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertDomainsResponseParams struct {
	// List of domain names connected to CDN
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// List of CDN domain names with certificates configured
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertifiedDomains []*string `json:"CertifiedDomains,omitempty" name:"CertifiedDomains"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCertDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertDomainsResponseParams `json:"Response"`
}

func (r *DescribeCertDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsConfigRequestParams struct {
	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, complex type.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting rules
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

type DescribeDomainsConfigRequest struct {
	*tchttp.BaseRequest
	
	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, complex type.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`

	// Sorting rules
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeDomainsConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsConfigResponseParams struct {
	// List of domain names
	Domains []*DetailDomain `json:"Domains,omitempty" name:"Domains"`

	// Number of domain names that match the specified query conditions
	// Used for paginated queries
	TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainsConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsConfigResponseParams `json:"Response"`
}

func (r *DescribeDomainsConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsRequestParams struct {
	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, which supports complex type.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, which supports complex type.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsResponseParams struct {
	// List of domain names
	Domains []*BriefDomain `json:"Domains,omitempty" name:"Domains"`

	// The number of domain names that matched the query conditions
	// Used for paginated queries
	TotalNumber *int64 `json:"TotalNumber,omitempty" name:"TotalNumber"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsResponseParams `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpStatusRequestParams struct {
	// Acceleration domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Node type.
	// `edge`: Edge server
	// `last`: Intermediate server
	// If this parameter is left empty, edge server information will be returned by default
	Layer *string `json:"Layer,omitempty" name:"Layer"`

	// Specifies a region to query.
	// `mainland`: Nodes in the Chinese mainland
	// `overseas`: Nodes outside the Chinese mainland
	// `global`: Global nodes
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether to return a value as an IP range
	Segment *bool `json:"Segment,omitempty" name:"Segment"`
}

type DescribeIpStatusRequest struct {
	*tchttp.BaseRequest
	
	// Acceleration domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Node type.
	// `edge`: Edge server
	// `last`: Intermediate server
	// If this parameter is left empty, edge server information will be returned by default
	Layer *string `json:"Layer,omitempty" name:"Layer"`

	// Specifies a region to query.
	// `mainland`: Nodes in the Chinese mainland
	// `overseas`: Nodes outside the Chinese mainland
	// `global`: Global nodes
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether to return a value as an IP range
	Segment *bool `json:"Segment,omitempty" name:"Segment"`
}

func (r *DescribeIpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Layer")
	delete(f, "Area")
	delete(f, "Segment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpStatusResponseParams struct {
	// Node list
	Ips []*IpStatus `json:"Ips,omitempty" name:"Ips"`

	// Total number of nodes
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpStatusResponseParams `json:"Response"`
}

func (r *DescribeIpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpVisitRequestParams struct {
	// Query start time, such as 2018-09-04 10:40:10; the returned result is later than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:10 and the query time granularity is 5 minutes, the time for the first returned entry will be 2018-09-04 10:40:00.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time, such as 2018-09-04 10:40:10; the returned result is earlier than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:10 and the query time granularity is 5 minutes, the time for the last returned entry will be 2018-09-04 10:40:00.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the list of domain names to be queried; up to 30 domain names can be queried at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Sampling interval in minutes. The available options vary for different query period. See below: 
	// 5min: 5 minutes. If the query period is within 24 hours, `5min` will be used by default.
	// day: 1 day. If the query period is longer than 24 hours, `day` will be used by default.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
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
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Sampling interval in minutes. The available options vary for different query period. See below: 
	// 5min: 5 minutes. If the query period is within 24 hours, `5min` will be used by default.
	// day: 1 day. If the query period is longer than 24 hours, `day` will be used by default.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeIpVisitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpVisitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpVisitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpVisitResponseParams struct {
	// Time granularity of data statistics, which supports 5min (5 minutes) and day (1 day).
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Origin-pull data details of each resource.
	Data []*ResourceData `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpVisitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpVisitResponseParams `json:"Response"`
}

func (r *DescribeIpVisitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpVisitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMapInfoRequestParams struct {
	// Query type:
	// `isp`: queries ISP codes
	// `district`: queries codes of provinces (Mainland China) or countries/regions (outside Mainland China)
	Name *string `json:"Name,omitempty" name:"Name"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMapInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMapInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMapInfoResponseParams struct {
	// Array of mappings.
	MapInfoList []*MapInfo `json:"MapInfoList,omitempty" name:"MapInfoList"`

	// Mapping relationship between server region ID and sub-region ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ServerRegionRelation []*RegionMapRelation `json:"ServerRegionRelation,omitempty" name:"ServerRegionRelation"`

	// Mapping relationship between client region ID and sub-region ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClientRegionRelation []*RegionMapRelation `json:"ClientRegionRelation,omitempty" name:"ClientRegionRelation"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMapInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMapInfoResponseParams `json:"Response"`
}

func (r *DescribeMapInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMapInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginDataRequestParams struct {
	// Start time of the query, e.g., 2018-09-04 10:40:00.
	// The specified start time will be rounded down based on the granularity parameter `Interval`. For example, if you set the start time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query, e.g. 2018-09-04 10:40:00.
	// The specified end time will be rounded down based on the granularity parameter `Interval`. For example, if you set the end time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the metric to query, which can be:
	// `flux`: Origin-pull traffic (in bytes)
	// `bandwidth`: Origin-pull bandwidth (in bps)
	// `request`: Number of origin-pull requests
	// `failRequest`: Number of failed origin-pull requests
	// `failRate`: Origin-pull failure rate (in %)
	// `statusCode`: Origin-pull status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx origin-pull status codes will be returned (in entries)
	// `2xx`: Returns the aggregate list of 2xx origin-pull status codes and the data for origin-pull status codes starting with 2 (in entries)
	// `3xx`: Returns the aggregate list of 3xx origin-pull status codes and the data for origin-pull status codes starting with 3 (in entries)
	// `4xx`: Returns the aggregate list of 4xx origin-pull status codes and the data for origin-pull status codes starting with 4 (in entries)
	// `5xx`: Returns the aggregate list of 5xx origin-pull status codes and the data for origin-pull status codes starting with 5 (in entries)
	// It is supported to specify a status code for query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the list of domain names to query. You can query up to 30 domain names at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried. [Check project ID in the console](https://console.cloud.tencent.com/project)
	// If the domain name is not specified, the specified project will be queried. Up to 30 acceleration domain names can be queried at a time.
	// If the domain name information is specified, this parameter can be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Time granularity, which can be:
	// `min`: Return data with 1-minute granularity. It’s available when the query period is  within 24 hours and `Area` is `mainland`.
	// `5min`: Return data with 5-minute granularity. It’s available when the query period is within 31 days.
	// `hour`: Return data with 1-hour granularity. It’s available when the query period is within 31 days.
	// `day`: Return data with 1-day granularity. It’s available when the query period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) when multiple `Domains` are passed in.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported)
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies the service region. If this value is left blank, it means to query CDN data within the Chinese mainland.
	// `mainland`: Query CDN data in the Chinese mainland.
	// `overseas`: Query CDN data outside the Chinese mainland.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a time zone to query. The default time zone is UTC+08:00.
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

type DescribeOriginDataRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query, e.g., 2018-09-04 10:40:00.
	// The specified start time will be rounded down based on the granularity parameter `Interval`. For example, if you set the start time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query, e.g. 2018-09-04 10:40:00.
	// The specified end time will be rounded down based on the granularity parameter `Interval`. For example, if you set the end time to 2018-09-04 10:40:00 with 1-hour granularity, the time will be rounded down to 2018-09-04 10:00:00.
	// The period between the start time and end time can be up to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the metric to query, which can be:
	// `flux`: Origin-pull traffic (in bytes)
	// `bandwidth`: Origin-pull bandwidth (in bps)
	// `request`: Number of origin-pull requests
	// `failRequest`: Number of failed origin-pull requests
	// `failRate`: Origin-pull failure rate (in %)
	// `statusCode`: Origin-pull status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx origin-pull status codes will be returned (in entries)
	// `2xx`: Returns the aggregate list of 2xx origin-pull status codes and the data for origin-pull status codes starting with 2 (in entries)
	// `3xx`: Returns the aggregate list of 3xx origin-pull status codes and the data for origin-pull status codes starting with 3 (in entries)
	// `4xx`: Returns the aggregate list of 4xx origin-pull status codes and the data for origin-pull status codes starting with 4 (in entries)
	// `5xx`: Returns the aggregate list of 5xx origin-pull status codes and the data for origin-pull status codes starting with 5 (in entries)
	// It is supported to specify a status code for query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the list of domain names to query. You can query up to 30 domain names at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried. [Check project ID in the console](https://console.cloud.tencent.com/project)
	// If the domain name is not specified, the specified project will be queried. Up to 30 acceleration domain names can be queried at a time.
	// If the domain name information is specified, this parameter can be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Time granularity, which can be:
	// `min`: Return data with 1-minute granularity. It’s available when the query period is  within 24 hours and `Area` is `mainland`.
	// `5min`: Return data with 5-minute granularity. It’s available when the query period is within 31 days.
	// `hour`: Return data with 1-hour granularity. It’s available when the query period is within 31 days.
	// `day`: Return data with 1-day granularity. It’s available when the query period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) when multiple `Domains` are passed in.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported)
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies the service region. If this value is left blank, it means to query CDN data within the Chinese mainland.
	// `mainland`: Query CDN data in the Chinese mainland.
	// `overseas`: Query CDN data outside the Chinese mainland.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a time zone to query. The default time zone is UTC+08:00.
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

func (r *DescribeOriginDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	delete(f, "Detail")
	delete(f, "Area")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginDataResponseParams struct {
	// Time granularity of data statistics, which supports `min` (1 minute), `5min` (5 minutes), `hour` (1 hour), and `day` (1 day).
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Origin-pull data details of each resource.
	Data []*ResourceOriginData `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOriginDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginDataResponseParams `json:"Response"`
}

func (r *DescribeOriginDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePayTypeRequestParams struct {
	// Specifies a service region.
	// `mainland`: queries billing methods within Mainland China;
	// `overseas`: queries billing methods outside Mainland China.
	// Default value: `mainland`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribePayTypeRequest struct {
	*tchttp.BaseRequest
	
	// Specifies a service region.
	// `mainland`: queries billing methods within Mainland China;
	// `overseas`: queries billing methods outside Mainland China.
	// Default value: `mainland`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePayTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Area")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePayTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePayTypeResponseParams struct {
	// Billing modes:
	// `flux`: Bill by traffic
	// `bandwidth`: Bill by bandwidth
	// `request`: Bill by the number of requests
	// `flux_sep`: Disused field
	// `bandwidth_sep`: Disused field
	// When you change a daily billing mode to another, and there is network usage on the day of change, this field shows the new billing mode, which takes effect from the next day. If there is no network usage on the day of change, this field shows the new billing mode directly.
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// Billing cycle:
	// `day`: Daily
	// `month`: Monthly 
	// `hour`: Hourly 
	BillingCycle *string `json:"BillingCycle,omitempty" name:"BillingCycle"`

	// `monthMax`: Billed by the monthly average of daily peak traffic (monthly settlement)
	// `day95`: Billed by the daily 95th percentile bandwidth (monthly settlement)
	// `month95`: Billed by the monthly 95th percentile bandwidth (monthly settlement)
	// `sum`: Billed by the total traffic/total requests (daily or monthly settlement)
	// `max`: Billed by the peak bandwidth (daily settlement)
	StatType *string `json:"StatType,omitempty" name:"StatType"`

	// Billing method for regions outside the Chinese mainland:
	// `all`: Unified billing for all regions
	// `multiple`: Region-specific billing
	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

	// The current billing mode in effect:
	// `flux`: Billed by traffic
	// `bandwidth`: Billed by bandwidth
	// `request`: Billed by the number of requests
	// `flux_sep`: Disused field
	// `bandwidth_sep`: Disused field
	CurrentPayType *string `json:"CurrentPayType,omitempty" name:"CurrentPayType"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribePayTypeResponseParams `json:"Response"`
}

func (r *DescribePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePayTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeQuotaRequestParams struct {

}

type DescribePurgeQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePurgeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeQuotaResponseParams struct {
	// URL purge usage and quota.
	UrlPurge []*Quota `json:"UrlPurge,omitempty" name:"UrlPurge"`

	// Directory purge usage and quota.
	PathPurge []*Quota `json:"PathPurge,omitempty" name:"PathPurge"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePurgeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeQuotaResponseParams `json:"Response"`
}

func (r *DescribePurgeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksRequestParams struct {
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

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20
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

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PurgeType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Keyword")
	delete(f, "Status")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// Detailed purge record.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PurgeLogs []*PurgeTask `json:"PurgeLogs,omitempty" name:"PurgeLogs"`

	// Total number of tasks, which is used for pagination.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribePushQuotaRequestParams struct {

}

type DescribePushQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePushQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePushQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushQuotaResponseParams struct {
	// URL prefetch usage and quota.
	UrlPush []*Quota `json:"UrlPush,omitempty" name:"UrlPush"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePushQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribePushQuotaResponseParams `json:"Response"`
}

func (r *DescribePushQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushTasksRequestParams struct {
	// Starting time, such as `2018-08-08 00:00:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as `2018-08-08 23:59:59`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies a task ID for your query.
	// You must specify either a task ID or a starting time.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Specifies a keyword for your query. Please enter a domain name or a complete URL beginning with `http(s)://`
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a region for your query:
	// `mainland`: Chinese mainland
	// `overseas`: Outside the Chinese mainland
	// `global`: Globe
	Area *string `json:"Area,omitempty" name:"Area"`

	// Queries the status of a specified task
	// `fail`: Prefetch failed
	// `done`: Prefetch succeeded
	// `process`: Prefetch in progress
	// `invalid`: Invalid prefetch with 4XX/5XX status code returned from the origin server
	Status *string `json:"Status,omitempty" name:"Status"`
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

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies a region for your query:
	// `mainland`: Chinese mainland
	// `overseas`: Outside the Chinese mainland
	// `global`: Globe
	Area *string `json:"Area,omitempty" name:"Area"`

	// Queries the status of a specified task
	// `fail`: Prefetch failed
	// `done`: Prefetch succeeded
	// `process`: Prefetch in progress
	// `invalid`: Invalid prefetch with 4XX/5XX status code returned from the origin server
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribePushTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Keyword")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePushTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushTasksResponseParams struct {
	// Prefetch history
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PushLogs []*PushTask `json:"PushLogs,omitempty" name:"PushLogs"`

	// Total number of tasks, which is used for pagination.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePushTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePushTasksResponseParams `json:"Response"`
}

func (r *DescribePushTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportDataRequestParams struct {
	// Query the start time in the format of `yyyy-MM-dd`
	// If the report type is `daily`, the start time and end time must be of the same day.
	// If the report type is `weekly`, the start time must be Monday and the end time must be the Sunday of the same week.
	// If the report type is `monthly`, the start time must be the first day of the month and the end time must be the last day of the same month.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query the end time in the format of `yyyy-MM-dd`
	// If the report type is `daily`, the start time and end time must be of the same day.
	// If the report type is `weekly`, the start time must be Monday and the end time must be the Sunday of the same week.
	// If the report type is `monthly`, the start time must be the first day of the month and the end time must be the last day of the same month.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Report type
	// daily: daily report
	// weekly: weekly report (Monday to Sunday)
	// monthly: monthly report (calendar month)
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// Domain name acceleration region
	// `mainland`: Regions within the Chinese mainland
	// `overseas`: Regions outside the Chinese mainland
	Area *string `json:"Area,omitempty" name:"Area"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of data entries. Default value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filters by project ID
	Project *int64 `json:"Project,omitempty" name:"Project"`
}

type DescribeReportDataRequest struct {
	*tchttp.BaseRequest
	
	// Query the start time in the format of `yyyy-MM-dd`
	// If the report type is `daily`, the start time and end time must be of the same day.
	// If the report type is `weekly`, the start time must be Monday and the end time must be the Sunday of the same week.
	// If the report type is `monthly`, the start time must be the first day of the month and the end time must be the last day of the same month.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query the end time in the format of `yyyy-MM-dd`
	// If the report type is `daily`, the start time and end time must be of the same day.
	// If the report type is `weekly`, the start time must be Monday and the end time must be the Sunday of the same week.
	// If the report type is `monthly`, the start time must be the first day of the month and the end time must be the last day of the same month.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Report type
	// daily: daily report
	// weekly: weekly report (Monday to Sunday)
	// monthly: monthly report (calendar month)
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// Domain name acceleration region
	// `mainland`: Regions within the Chinese mainland
	// `overseas`: Regions outside the Chinese mainland
	Area *string `json:"Area,omitempty" name:"Area"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of data entries. Default value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filters by project ID
	Project *int64 `json:"Project,omitempty" name:"Project"`
}

func (r *DescribeReportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ReportType")
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Project")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportDataResponseParams struct {
	// Domain name-level data details.
	DomainReport []*ReportData `json:"DomainReport,omitempty" name:"DomainReport"`

	// Project-level data details
	ProjectReport []*ReportData `json:"ProjectReport,omitempty" name:"ProjectReport"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReportDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportDataResponseParams `json:"Response"`
}

func (r *DescribeReportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlViolationsRequestParams struct {
	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specified domain name query
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type DescribeUrlViolationsRequest struct {
	*tchttp.BaseRequest
	
	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specified domain name query
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *DescribeUrlViolationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlViolationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUrlViolationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlViolationsResponseParams struct {
	// Details of URLs in violation
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UrlRecordList []*ViolationUrl `json:"UrlRecordList,omitempty" name:"UrlRecordList"`

	// Total number of records, which is used for pagination.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUrlViolationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUrlViolationsResponseParams `json:"Response"`
}

func (r *DescribeUrlViolationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlViolationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailDomain struct {
	// Domain name ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tencent Cloud account ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Accelerated domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// CNAME address of domain name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Acceleration service status
	// `rejected`: The domain name is rejected due to expiration/deregistration of its ICP filing
	// `processing`: Deploying
	// `online`: Enabled
	// `offline`: Disabled
	Status *string `json:"Status,omitempty" name:"Status"`

	// Project ID, which can be viewed on the Tencent Cloud project management page
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Acceleration domain name service type
	// `web`: Webpage file downloads
	// `download`: Large file downloads
	// `media`: Audio and video on demand acceleration
	// `hybrid`: Dynamic and static content acceleration
	// `dynamic`: Dynamic content acceleration
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Domain name creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Domain name update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin server configuration.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP blocklist/allowlist configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Status code cache configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Smart compression configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Bandwidth cap configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range GETs configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 origin-pull follow-redirect configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Custom error page configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// Custom request header configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Custom response header configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Single-link downstream speed limit configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Configuration of cache with/without parameter
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Origin server header cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// Video dragging configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// Node cache expiration rule configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cross-border linkage optimization configuration (in beta)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// HTTPS Acceleration Configuration Guide
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Timestamp hotlink protection configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// SEO configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Domain name block status
	// `normal`: Normal
	// `overdue`: The domain name has been disabled due to account arrears. The acceleration service can be resumed after the account is topped up.
	// `malicious`: The acceleration service has been forcibly disabled due to detection of malicious behavior.
	// `ddos`: The acceleration service has been disabled due to large-scale DDoS attacks to the domain name
	// `idle`: No operations or data has been detected for more than 90 days. The domain name is determined to be inactive which automatically disables the acceleration service. You can restart the service.
	// `unlicensed`: The acceleration service has been automatically disabled as the domain name has no ICP filing or its ICP filing is deregistered. Service can be resumed after an ICP filing is obtained.
	// `capping`: The configured upper limit for bandwidth has been reached.
	// `readonly`: The domain name has a special configuration and has been locked.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// Access protocol forced redirect configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Browser cache expiration rule configuration (in beta)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// IPv6 origin-pull configuration (in beta)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Ipv6 *Ipv6 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Backwards compatibility configuration (compatibility field for internal use)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Compatibility *Compatibility `json:"Compatibility,omitempty" name:"Compatibility"`

	// Region-specific configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// Acceleration region
	// `mainland`: Acceleration inside the Chinese mainland
	// `overseas`: Acceleration outside the Chinese mainland
	// `global`: Acceleration over the globe
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Domain name lock status
	// `normal`: Not locked
	// `mainland`: Locked in the Chinese mainland
	// `overseas`: Locked outside the Chinese mainland
	// global: Locked globally
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`

	// Origin-pull timeout configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitempty" name:"OriginPullTimeout"`

	// S3 bucket origin access authentication configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitempty" name:"AwsPrivateAccess"`

	// SCDN configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitempty" name:"SecurityConfig"`

	// Image optimization configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ImageOptimization *ImageOptimization `json:"ImageOptimization,omitempty" name:"ImageOptimization"`

	// UA blocklist/allowlist configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UserAgentFilter *UserAgentFilter `json:"UserAgentFilter,omitempty" name:"UserAgentFilter"`

	// Access control
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AccessControl *AccessControl `json:"AccessControl,omitempty" name:"AccessControl"`

	// Whether to support advanced configuration items
	// `on`: Supported
	// `off`: Not supported
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Advance *string `json:"Advance,omitempty" name:"Advance"`

	// URL redirect configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UrlRedirect *UrlRedirect `json:"UrlRedirect,omitempty" name:"UrlRedirect"`

	// Access port configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AccessPort []*int64 `json:"AccessPort,omitempty" name:"AccessPort"`

	// Tag configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// Timestamp hotlink protection advanced configuration (allowlist feature)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AdvancedAuthentication *AdvancedAuthentication `json:"AdvancedAuthentication,omitempty" name:"AdvancedAuthentication"`

	// Origin-pull authentication advanced configuration (allowlist feature)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginAuthentication *OriginAuthentication `json:"OriginAuthentication,omitempty" name:"OriginAuthentication"`

	// IPv6 access configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitempty" name:"Ipv6Access"`

	// Advanced configuration settings
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AdvanceSet []*AdvanceConfig `json:"AdvanceSet,omitempty" name:"AdvanceSet"`

	// Offline cache (only available to beta users)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// Merging origin-pull requests (only available to beta users)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginCombine *OriginCombine `json:"OriginCombine,omitempty" name:"OriginCombine"`

	// POST request configuration item
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PostMaxSize *PostSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// QUIC configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// Access authentication for OSS origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitempty" name:"OssPrivateAccess"`

	// WebSocket configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Remote authentication configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RemoteAuthentication *RemoteAuthentication `json:"RemoteAuthentication,omitempty" name:"RemoteAuthentication"`

	// Shared CNAME configuration (only available to beta users)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ShareCname *ShareCname `json:"ShareCname,omitempty" name:"ShareCname"`

	// Rule engine
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleEngine *RuleEngine `json:"RuleEngine,omitempty" name:"RuleEngine"`

	// Primary domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ParentHost *string `json:"ParentHost,omitempty" name:"ParentHost"`

	// Access authentication for Huawei Cloud OBS origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitempty" name:"HwPrivateAccess"`

	// Access authentication for QiNiu Cloud Kodo origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitempty" name:"QnPrivateAccess"`
}

// Predefined struct for user
type DisableCachesRequestParams struct {
	// List of URLs to be blocked (URLs must contain `http://` or `https://`).
	// Up to 100 entries can be submitted at a time and 3,000 entries per day.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

type DisableCachesRequest struct {
	*tchttp.BaseRequest
	
	// List of URLs to be blocked (URLs must contain `http://` or `https://`).
	// Up to 100 entries can be submitted at a time and 3,000 entries per day.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

func (r *DisableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCachesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableCachesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCachesResponseParams struct {
	// Submission result
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

	// Task ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableCachesResponse struct {
	*tchttp.BaseResponse
	Response *DisableCachesResponseParams `json:"Response"`
}

func (r *DisableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClsLogTopicRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type DisableClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *DisableClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClsLogTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *DisableClsLogTopicResponseParams `json:"Response"`
}

func (r *DisableClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAreaConfig struct {
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Region list, where the element can be `mainland`/`overseas`
	Area []*string `json:"Area,omitempty" name:"Area"`
}

type DomainFilter struct {
	// Filters by the field name, which includes:
	// - `origin`: Primary origin server.
	// - `domain`: Domain name.
	// - `resourceId`: Domain name ID.
	// - `status`: Domain name status. Valid values: `online`, `offline`, and `processing`.
	// - `serviceType`: Service type. Valid values: `web`, `download`, `media`, `hybrid` and `dynamic`.
	// - `projectId`: Project ID.
	// - `domainType`: Primary origin server type. Valid values: `cname` (customer origin), `COS` (COS origin), and `third_party` (third-party object storage origin).
	// - `fullUrlCache`: Whether to enable full-path cache, which can be `on` or `off`.
	// - `https`: Whether to configure HTTPS, which can be `on`, `off` or `processing`.
	// - `originPullProtocol`: Origin-pull protocol type, which can be `http`, `follow`, or `https`.
	// - `tagKey`: Tag key.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter field value.
	Value []*string `json:"Value,omitempty" name:"Value"`

	// Whether to enable fuzzy query. Only `origin` or `domain` is supported for the filter field name.
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
	// `mainland`: Within the Chinese mainland
	// `overseas`: Outside the Chinese mainland
	Area *string `json:"Area,omitempty" name:"Area"`

	// Log package filename
	LogName *string `json:"LogName,omitempty" name:"LogName"`
}

type DownstreamCapping struct {
	// Downstream speed configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Downstream speed limiting rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CappingRules []*CappingRule `json:"CappingRules,omitempty" name:"CappingRules"`
}

// Predefined struct for user
type EnableCachesRequestParams struct {
	// List of unblocked URLs
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// URL blocking date
	Date *string `json:"Date,omitempty" name:"Date"`
}

type EnableCachesRequest struct {
	*tchttp.BaseRequest
	
	// List of unblocked URLs
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// URL blocking date
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *EnableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCachesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableCachesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableCachesResponseParams struct {
	// Result list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

	// Task ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableCachesResponse struct {
	*tchttp.BaseResponse
	Response *EnableCachesResponseParams `json:"Response"`
}

func (r *EnableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClsLogTopicRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type EnableClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *EnableClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClsLogTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *EnableClsLogTopicResponseParams `json:"Response"`
}

func (r *EnableClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorPage struct {
	// Status code redirect configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Status code redirect rules configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PageRules []*ErrorPageRule `json:"PageRules,omitempty" name:"PageRules"`
}

type ErrorPageRule struct {
	// Status code
	// Supports 400, 403, 404, 500.
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// Redirect status code settings
	// Supports 301 or 302.
	RedirectCode *int64 `json:"RedirectCode,omitempty" name:"RedirectCode"`

	// Redirect URL
	// Requires a full redirect path, such as https://www.test.com/error.html.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

type ExtraLogset struct {
	// Logset information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Logset *LogSetInfo `json:"Logset,omitempty" name:"Logset"`

	// Log topic information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics"`
}

type FollowRedirect struct {
	// Origin-pull follow-redirect switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`


	RedirectConfig *RedirectConfig `json:"RedirectConfig,omitempty" name:"RedirectConfig"`
}

type ForceRedirect struct {
	// Access forced redirect configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Access forced redirect types
	// http: forced HTTP redirect
	// https: forced HTTPS redirect
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`

	// Status code returned for forced redirect 
	// Supports 301, 302.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`

	// Whether to return the newly added header during force redirection
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CarryHeaders *string `json:"CarryHeaders,omitempty" name:"CarryHeaders"`
}

// Predefined struct for user
type GetDisableRecordsRequestParams struct {
	// Specifies the URL to be queried
	Url *string `json:"Url,omitempty" name:"Url"`

	// Starting time, such as `2018-12-12 10:24:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as `2018-12-14 10:24:00`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Current URL status
	// disable: The URL remains disabled, and accessing it will return an error 403
	// enable: The URL is enabled (unblocked) and can be normally accessed
	Status *string `json:"Status,omitempty" name:"Status"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination limit. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Task ID. The task ID and start time cannot be both left empty.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type GetDisableRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the URL to be queried
	Url *string `json:"Url,omitempty" name:"Url"`

	// Starting time, such as `2018-12-12 10:24:00`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as `2018-12-14 10:24:00`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Current URL status
	// disable: The URL remains disabled, and accessing it will return an error 403
	// enable: The URL is enabled (unblocked) and can be normally accessed
	Status *string `json:"Status,omitempty" name:"Status"`

	// Offset for paginated queries. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination limit. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Task ID. The task ID and start time cannot be both left empty.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GetDisableRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDisableRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDisableRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDisableRecordsResponseParams struct {
	// Blocking history
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UrlRecordList []*UrlRecord `json:"UrlRecordList,omitempty" name:"UrlRecordList"`

	// Total number of tasks, which is used for pagination.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDisableRecordsResponse struct {
	*tchttp.BaseResponse
	Response *GetDisableRecordsResponseParams `json:"Response"`
}

func (r *GetDisableRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDisableRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GuetzliAdapter struct {
	// Switch. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type HeaderKey struct {
	// Whether to use it as part of `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Array of headers that make up the `CacheKey` (separated by ';')
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type HeuristicCache struct {
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Heuristic cache validity configuration
	CacheConfig *CacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`
}

type Hsts struct {
	// Whether to enable. Valid values: on, off.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// `MaxAge` value.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAge *int64 `json:"MaxAge,omitempty" name:"MaxAge"`

	// Whether to include subdomain names. Valid values: on, off.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IncludeSubDomains *string `json:"IncludeSubDomains,omitempty" name:"IncludeSubDomains"`
}

type HttpHeaderPathRule struct {
	// HTTP header setting methods
	// `set`: sets a value for an existing header parameter, a new header parameter, or multiple header parameters. Multiple header parameters will be merged into one.
	// `del`: deletes a header parameter.
	// `add`: adds a header parameter. By default, you can repeat the same action to add the same header parameter, which may affect browser response. Please consider the set operation first.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HeaderMode *string `json:"HeaderMode,omitempty" name:"HeaderMode"`

	// HTTP header name. Up to 100 characters can be set.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// HTTP header value. Up to 1000 characters can be set.
	// Not required when Mode is del
	// Required when Mode is add/set
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`

	// Rule types:
	// `all`: Apply to all files.
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each `RuleType`:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`
}

type HttpHeaderRule struct {
	// HTTP header setting method. Valid values: `add` (add header), `set` (set header) or `del` (delete header).
	HeaderMode *string `json:"HeaderMode,omitempty" name:"HeaderMode"`

	// HTTP header name
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// HTTP header value
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
}

type Https struct {
	// HTTPS configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// HTTP2 configuration switch
	// `on`: Enable
	// `off`: Disable
	// Enabling HTTPS acceleration for the first time will enable HTTP2 configuration by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// OCSP configuration switch
	// `on`: Enable
	// `off`: Disable
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// Client certificate authentication feature
	// `on`: Enable
	// `off`: Disable
	// This is disabled by default. The client certificate information is needed when enabled. This is still in beta and not generally available yet.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VerifyClient *string `json:"VerifyClient,omitempty" name:"VerifyClient"`

	// Server certificate configuration information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertInfo *ServerCert `json:"CertInfo,omitempty" name:"CertInfo"`

	// Client certificate configuration information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClientCertInfo *ClientCert `json:"ClientCertInfo,omitempty" name:"ClientCertInfo"`

	// Spdy configuration switch
	// `on`: Enable
	// `off`: Disable
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Spdy *string `json:"Spdy,omitempty" name:"Spdy"`

	// HTTPS certificate deployment status
	// closed: already closed
	// deploying: in deployment
	// deployed: successfully deployed
	// failed: deployment failed
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SslStatus *string `json:"SslStatus,omitempty" name:"SslStatus"`

	// HSTS configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Hsts *Hsts `json:"Hsts,omitempty" name:"Hsts"`

	// TLS version settings, which only support certain advanced domain names. Valid values: `TLSv1`, `TLSV1.1`, `TLSV1.2`, and `TLSv1.3`. Only consecutive versions can be enabled at the same time.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TlsVersion []*string `json:"TlsVersion,omitempty" name:"TlsVersion"`
}

type HwPrivateAccess struct {
	// Whether to enable access authentication. Valid values: `on`, `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Access ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// Key
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// BucketName
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

type ImageOptimization struct {
	// `WebpAdapter` configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	WebpAdapter *WebpAdapter `json:"WebpAdapter,omitempty" name:"WebpAdapter"`

	// `TpgAdapter` configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TpgAdapter *TpgAdapter `json:"TpgAdapter,omitempty" name:"TpgAdapter"`

	// `GuetzliAdapter` configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitempty" name:"GuetzliAdapter"`
}

type IpFilter struct {
	// IP blocklist/allowlist configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// IP blocklist/allowlist type
	// `whitelist`: IP allowlist
	// `blacklist`: IP blocklist
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// IP blocklist/allowlist
	// Supports IPs in X.X.X.X format, or IP ranges in /8, /16, /24 format.
	// Up to 50 whitelists or blacklists can be entered
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Filters []*string `json:"Filters,omitempty" name:"Filters"`

	// IP blocklist/allowlist path-based configuration. This feature is only available to selected beta customers.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FilterRules []*IpFilterPathRule `json:"FilterRules,omitempty" name:"FilterRules"`

	// HTTP code returned when the IP allowlist/blocklist verification fails
	// Valid values: 400-499
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
}

type IpFilterPathRule struct {
	// IP blocklist/allowlist type
	// `whitelist`: allowlist IPs
	// `blacklist`: blocklist IPs
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// IP blocklist/allowlist list
	// Supports IPs in X.X.X.X format, or /8, /16, /24 format IP ranges.
	// Up to 50 allowlists or blocklists can be entered.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Filters []*string `json:"Filters,omitempty" name:"Filters"`

	// Rule types:
	// `all`: Effective for all files
	// `file`: Effective for specified file suffixes
	// `directory`: Effective for specified paths
	// `path`: Effective for specified absolute paths
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each RuleType:
	// For `all`, enter an asterisk (*).
	// For `file`, enter the suffix, such as jpg, txt.
	// For `directory`, enter the path, such as /xxx/test/.
	// For `path`, enter the corresponding absolute path, such as /xxx/test.html.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`
}

type IpFreqLimit struct {
	// IP access limit configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Sets the limited number of requests per second
	// 514 will be returned for requests that exceed the limit
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
}

type IpStatus struct {
	// Node IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Region of the node
	District *string `json:"District,omitempty" name:"District"`

	// ISP of the node
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// City of the node
	City *string `json:"City,omitempty" name:"City"`

	// Status of the node
	// `online`: The node is active and scheduling normally.
	// `offline`: The node is inactive.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type Ipv6 struct {
	// Whether to enable the IPv6 feature for a domain name. Values include `on` or `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Ipv6Access struct {
	// Whether to enable the IPv6 feature for a domain name. Values include `on` or `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type KeyRule struct {
	// Content for each CacheType:
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter the path, such as /xxx/test/.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// For `index`, enter a backslash (/).
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// Rule types:
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	// `index`: home page
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Whether to enable full-path cache
	// `on`: Enable full-path cache (i.e., disable Ignore Query String).
	// `off`: Disable full-path cache (i.e., enable Ignore Query String).
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`

	// Whether caches are case insensitive
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreCase *string `json:"IgnoreCase,omitempty" name:"IgnoreCase"`

	// Request parameter contained in `CacheKey`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	QueryString *RuleQueryString `json:"QueryString,omitempty" name:"QueryString"`

	// Path cache key tag, the value "user" is passed.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleTag *string `json:"RuleTag,omitempty" name:"RuleTag"`
}

// Predefined struct for user
type ListClsLogTopicsRequestParams struct {
	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type ListClsLogTopicsRequest struct {
	*tchttp.BaseRequest
	
	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *ListClsLogTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsLogTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListClsLogTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClsLogTopicsResponseParams struct {
	// Information of logsets in the Shanghai region
	Logset *LogSetInfo `json:"Logset,omitempty" name:"Logset"`

	// Information of log topics in the Shanghai region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics"`

	// Information on logsets in regions except Shanghai
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExtraLogset []*ExtraLogset `json:"ExtraLogset,omitempty" name:"ExtraLogset"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListClsLogTopicsResponse struct {
	*tchttp.BaseResponse
	Response *ListClsLogTopicsResponseParams `json:"Response"`
}

func (r *ListClsLogTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsLogTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClsTopicDomainsRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type ListClsTopicDomainsRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *ListClsTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsTopicDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListClsTopicDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClsTopicDomainsResponseParams struct {
	// Developer ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// Channel
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Domain name region configuration, which may contain deleted domain names. If this is to be used in `ManageClsTopicDomains` API, you need to exclude deleted domain names by using the `ListCdnDomains` API.
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs"`

	// Log topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Last modified time of log topic
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListClsTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *ListClsTopicDomainsResponseParams `json:"Response"`
}

func (r *ListClsTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsTopicDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopDataRequestParams struct {
	// Query start time in the format of `yyyy-MM-dd HH:mm:ss`
	// Only supports data query at daily granularity. The date in the input parameter is used as the start date.
	// If the specified start date is greater than 00:00:00, it will be rounded down to 00:00:00 on the date. For example, if `StartTime` is 2018-09-04 10:40:00, it will be rounded down to 2018-09-04 00:00:00.
	// Only data from the last 90 days will be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of `yyyy-MM-dd HH:mm:ss`
	// Only supports data query at daily granularity. The date in the input parameter is used as the end date.
	// If the specified end date is smaller than 23:59:59, it will be rounded up to 23:59:59 on the date. For example, if `EndTime` is 2018-09-05 22:40:00, it will be rounded up to 2018-09-05 23:59:59.
	// `EndTime` must be later than or equal to `StartTime`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Objects to be sorted. Valid values:
	// `url`: Sort by access URL (URLs carrying no parameters). Supported filters are `flux` and `request`.
	// `district`: sorts provinces or countries/regions. Supported filters are `flux` and `request`.
	// `isp`: sorts ISPs. Supported filters are `flux` and `request`.
	// `host`: Sort by domain name access data. Supported filters are `flux`, `request`, `bandwidth`, `fluxHitRate`, and `statusCode` (2XX, 3XX, 4XX, 5XX).
	// `originHost`: Sort by domain name origin-pull data. Supported filters are `flux`, `request`, `bandwidth`, and `OriginStatusCode` (origin_2XX, origin_3XX, origin_4XX, origin_5XX).
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
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// The sorted results of all domain names are returned by default (false) during a multi-domain-name query
	// If `Metric` is `url`, `path`, `district`, or `isp` and `Filter` is `flux` or `request`, it can be set to `true` to return the sorted results of each domain.
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// When Filter is `statusCode` or `OriginStatusCode`, enter a code to query and sort results.
	Code *string `json:"Code,omitempty" name:"Code"`

	// Specifies the service region. If this value is left blank, it means to query CDN data within the Chinese mainland.
	// `mainland`: Query CDN data in the Chinese mainland.
	// `overseas`: Query CDN data outside the Chinese mainland. Supported metrics are `url`, `district`, `host`, and `originHost`. If `Metric` is `originHost`, supported filters are `flux`, `request`, and `bandwidth`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a region type for the query. If it is left blank, data of the service region will be queried. This parameter is only valid when `Area` is `overseas` and `Metric` is `district` or `host`.
	// `server`: Query by the location of server (Tencent Cloud CDN nodes).
	// `client`: Query data of the client region where the request devices are located; if `Metric` is `host`, supported filters are `flux`, `request`, and `bandwidth`.
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Returns the first N data entries. The default value is 100 if this parameter is not specified, whereas 1000 if `Metric` is `url`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type ListTopDataRequest struct {
	*tchttp.BaseRequest
	
	// Query start time in the format of `yyyy-MM-dd HH:mm:ss`
	// Only supports data query at daily granularity. The date in the input parameter is used as the start date.
	// If the specified start date is greater than 00:00:00, it will be rounded down to 00:00:00 on the date. For example, if `StartTime` is 2018-09-04 10:40:00, it will be rounded down to 2018-09-04 00:00:00.
	// Only data from the last 90 days will be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of `yyyy-MM-dd HH:mm:ss`
	// Only supports data query at daily granularity. The date in the input parameter is used as the end date.
	// If the specified end date is smaller than 23:59:59, it will be rounded up to 23:59:59 on the date. For example, if `EndTime` is 2018-09-05 22:40:00, it will be rounded up to 2018-09-05 23:59:59.
	// `EndTime` must be later than or equal to `StartTime`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Objects to be sorted. Valid values:
	// `url`: Sort by access URL (URLs carrying no parameters). Supported filters are `flux` and `request`.
	// `district`: sorts provinces or countries/regions. Supported filters are `flux` and `request`.
	// `isp`: sorts ISPs. Supported filters are `flux` and `request`.
	// `host`: Sort by domain name access data. Supported filters are `flux`, `request`, `bandwidth`, `fluxHitRate`, and `statusCode` (2XX, 3XX, 4XX, 5XX).
	// `originHost`: Sort by domain name origin-pull data. Supported filters are `flux`, `request`, `bandwidth`, and `OriginStatusCode` (origin_2XX, origin_3XX, origin_4XX, origin_5XX).
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
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// The sorted results of all domain names are returned by default (false) during a multi-domain-name query
	// If `Metric` is `url`, `path`, `district`, or `isp` and `Filter` is `flux` or `request`, it can be set to `true` to return the sorted results of each domain.
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// When Filter is `statusCode` or `OriginStatusCode`, enter a code to query and sort results.
	Code *string `json:"Code,omitempty" name:"Code"`

	// Specifies the service region. If this value is left blank, it means to query CDN data within the Chinese mainland.
	// `mainland`: Query CDN data in the Chinese mainland.
	// `overseas`: Query CDN data outside the Chinese mainland. Supported metrics are `url`, `district`, `host`, and `originHost`. If `Metric` is `originHost`, supported filters are `flux`, `request`, and `bandwidth`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a region type for the query. If it is left blank, data of the service region will be queried. This parameter is only valid when `Area` is `overseas` and `Metric` is `district` or `host`.
	// `server`: Query by the location of server (Tencent Cloud CDN nodes).
	// `client`: Query data of the client region where the request devices are located; if `Metric` is `host`, supported filters are `flux`, `request`, and `bandwidth`.
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Returns the first N data entries. The default value is 100 if this parameter is not specified, whereas 1000 if `Metric` is `url`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Filter")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Detail")
	delete(f, "Code")
	delete(f, "Area")
	delete(f, "AreaType")
	delete(f, "Product")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopDataResponseParams struct {
	// Top access data details of each resource
	Data []*TopData `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListTopDataResponse struct {
	*tchttp.BaseResponse
	Response *ListTopDataResponseParams `json:"Response"`
}

func (r *ListTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogSetInfo struct {
	// Developer ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// Channel
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// Whether it is the default logset
	IsDefault *uint64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// Log retention period in days
	LogsetSavePeriod *uint64 `json:"LogsetSavePeriod,omitempty" name:"LogsetSavePeriod"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Whether the logset has been removed from CLS
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Deleted *string `json:"Deleted,omitempty" name:"Deleted"`

	// Whether English is used in this region
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RegionEn *string `json:"RegionEn,omitempty" name:"RegionEn"`
}

type MainlandConfig struct {
	// Timestamp hotlink protection configuration.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// Bandwidth cap configuration.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Cache rule configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cache configurations.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Smart compression configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Download speed limit configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Error code redirect configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 301 and 302 automatic origin-pull follow-redirect configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Force redirect by access protocol.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// IP blocklist/allowlist configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limiting configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Browser cache rules configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Origin server configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Cross-border optimization configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Range GETs configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// Hotlink protection configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Origin-pull request header configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Origin server response header configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Follows origin server cache header configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// SEO configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Domain name service type. Values: `web` (static acceleration); `download` (download acceleration); `media` (streaming media acceleration).
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Status code cache configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Video dragging configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// Access authentication for S3 origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitempty" name:"AwsPrivateAccess"`

	// Access authentication for OSS origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitempty" name:"OssPrivateAccess"`

	// Access authentication for Huawei Cloud OBS origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitempty" name:"HwPrivateAccess"`

	// Access authentication for QiNiu Cloud Kodo origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitempty" name:"QnPrivateAccess"`
}

// Predefined struct for user
type ManageClsTopicDomainsRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Domain name region configuration. Note: if this field is empty, it means to unbind all domain names from the corresponding topic
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs"`
}

type ManageClsTopicDomainsRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Domain name region configuration. Note: if this field is empty, it means to unbind all domain names from the corresponding topic
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitempty" name:"DomainAreaConfigs"`
}

func (r *ManageClsTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageClsTopicDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	delete(f, "DomainAreaConfigs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageClsTopicDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageClsTopicDomainsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManageClsTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *ManageClsTopicDomainsResponseParams `json:"Response"`
}

func (r *ManageClsTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageClsTopicDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MapInfo struct {
	// Object ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Object name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type MaxAge struct {
	// Browser cache configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// MaxAge rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAgeRules []*MaxAgeRule `json:"MaxAgeRules,omitempty" name:"MaxAgeRules"`
}

type MaxAgeRule struct {
	// Rule types:
	// `all`: effective for all files.
	// `file`: effective for specified file suffixes.
	// `directory`: effective for specified paths.
	// `path`: effective for specified absolute paths.
	// `index`: effective for specified homepages.
	MaxAgeType *string `json:"MaxAgeType,omitempty" name:"MaxAgeType"`

	// Content for each `MaxAgeType`:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter the suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter the path, e.g., `/xxx/test/`.
	// For `path`, enter the absolute path, e.g., `/xxx/test.html`.
	// For `index`, enter a forward slash `/`.
	// Note: The rule `all` cannot be deleted. It follows origin by default and can be modified.
	MaxAgeContents []*string `json:"MaxAgeContents,omitempty" name:"MaxAgeContents"`

	// MaxAge time (in seconds)
	// Note: The value `0` means not to cache.
	MaxAgeTime *int64 `json:"MaxAgeTime,omitempty" name:"MaxAgeTime"`

	// Whether to follow the origin server. Valid values: `on` and `off`. If it's on, `MaxAgeTime` is ignored.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type OfflineCache struct {
	// Whether to enable offline cache. Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Origin struct {
	// Primary origin server list
	// When modifying the origin server, you need to enter the corresponding OriginType.
	// Note: This field may return `null`, indicating that no valid value was found.
	Origins []*string `json:"Origins,omitempty" name:"Origins"`

	// Primary origin server type
	// The following types are supported for input parameters:
	// `domain`: domain name
	// `domainv6`: IPv6 domain name
	// cos: COS origin
	// `ip`: IP address
	// ipv6: origin server list is a single IPv6 address
	// `ip_ipv6`: multiple IPv4 addresses and one IPv6 address
	// `ip_domain`: IP addresses and domain names (only available to beta users)
	// `ip_domainv6`: Multiple IPv4 addresses and one IPv6 domain name
	// `ipv6_domain`: multiple IPv6 addresses and one domain name
	// `ipv6_domainv6`: Multiple IPv6 addresses and one IPv6 domain name
	// `domain_domainv6`: Multiple IPv4 domain names and one IPv6 domain name
	// `ip_ipv6_domain`: multiple IPv4 and IPv6 addresses and one domain name
	// `ip_ipv6_domainv6`: Multiple IPv4 and IPv6 addresses and one IPv6 domain name
	// `ip_domain_domainv6`: Multiple IPv4 addresses and IPv4 domain names and one IPv6 domain name
	// `ipv6_domain_domainv6`: Multiple IPv4 domain names and IPv6 addresses and one IPv6 domain name
	// `ip_ipv6_domain_domainv6`: Multiple IPv4 and IPv6 addresses and IPv4 domain names and one IPv6 domain name
	// The following types of output parameters are added:
	// image: Cloud Infinite origin
	// ftp: legacy FTP origin, which is no longer maintained.
	// When modifying `Origins`, you need to enter the corresponding OriginType.
	// The IPv6 feature is not generally available yet. Please send in a whitelist application to use this feature.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// It is required when a COS origin or third-party origin is used for acceleration.
	// Host header used when accessing the primary origin server. If it is left empty, the acceleration domain name will be used by default.
	// If a wildcard domain name is accessed, then the sub-domain name during the access will be used by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// When OriginType is COS, you can specify if access to private buckets is allowed.
	// Note: To enable this configuration, you need to first grant CDN access to the private bucket. Values: `on` and `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CosPrivateAccess *string `json:"CosPrivateAccess,omitempty" name:"CosPrivateAccess"`

	// Origin-pull protocol configuration
	// http: forced HTTP origin-pull
	// follow: protocol follow origin-pull
	// https: forced HTTPS origin-pull. This only supports origin server port 443 for origin-pull.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// Backup origin server list
	// When modifying the backup origin server, you need to enter the corresponding BackupOriginType.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackupOrigins []*string `json:"BackupOrigins,omitempty" name:"BackupOrigins"`

	// Backup origin server type, which supports the following types:
	// `domain`: Domain name
	// `ip`: IP address
	// When modifying BackupOrigins, you need to enter the corresponding BackupOriginType.
	// The following backup origin servers are only available to beta users. Submit an application if you want to become a beta user.
	// `ipv6_domain`: Multiple IPv6 addresses and one domain name
	// `ip_ipv6`: Multiple IPv4 addresses and one IPv6 address
	// `ipv6_domain`: Multiple IPv6 addresses and one domain name
	// `ip_ipv6_domain`: Multiple IPv4 and IPv6 addresses and one domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackupOriginType *string `json:"BackupOriginType,omitempty" name:"BackupOriginType"`

	// Host header used when accessing the backup origin server. If it is left empty, the `ServerName` of primary origin server will be used by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackupServerName *string `json:"BackupServerName,omitempty" name:"BackupServerName"`

	// Origin-pull path
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BasePath *string `json:"BasePath,omitempty" name:"BasePath"`

	// Origin-pull path rewriting configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PathRules []*PathRule `json:"PathRules,omitempty" name:"PathRules"`

	// Path-based origin-pull configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PathBasedOrigin []*PathBasedOriginRule `json:"PathBasedOrigin,omitempty" name:"PathBasedOrigin"`

	// HTTPS advanced origin-pull configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AdvanceHttps *AdvanceHttps `json:"AdvanceHttps,omitempty" name:"AdvanceHttps"`

	// Object storage vendor
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginCompany *string `json:"OriginCompany,omitempty" name:"OriginCompany"`
}

type OriginAuthentication struct {
	// Authentication switch, which can be on or off.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Authentication type configuration A
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TypeA *OriginAuthenticationTypeA `json:"TypeA,omitempty" name:"TypeA"`
}

type OriginAuthenticationTypeA struct {
	// Key used for signature calculation, allowing 6 to 32 bytes of letters and digits.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type OriginCombine struct {
	// Whether to enable the merging pull requests feature. Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OriginIp struct {
	// Intermediate IP range/intermediate IP. The IP range information is returned by default.
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type OriginPullOptimization struct {
	// Cross-border origin-pull optimization configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Cross-border types
	// OVToCN: origin-pull from outside mainland China to inside mainland China
	// CNToOV: origin-pull from inside mainland China to outside mainland China
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OptimizationType *string `json:"OptimizationType,omitempty" name:"OptimizationType"`
}

type OriginPullTimeout struct {
	// The origin-pull connection timeout (in seconds). Valid range: 5-60.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`

	// The origin-pull receipt timeout (in seconds). Valid range: 10-60.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ReceiveTimeout *uint64 `json:"ReceiveTimeout,omitempty" name:"ReceiveTimeout"`
}

type OssPrivateAccess struct {
	// Whether to enable access authentication. Valid values: `on`, `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Access ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// Key.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Region
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// BucketName
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

type OverseaConfig struct {
	// Timestamp hotlink protection configuration.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Authentication *Authentication `json:"Authentication,omitempty" name:"Authentication"`

	// Bandwidth cap configuration.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Cache rule configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Cache configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Smart compression configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Download speed limit configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitempty" name:"DownstreamCapping"`

	// Error code redirect configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ErrorPage *ErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// 301 and 302 automatic origin-pull follow-redirect configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Protocol redirect configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// HTTPS configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// IP blocklist/allowlist configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Browser cache rules configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Origin server configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Cross-border optimization configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitempty" name:"OriginPullOptimization"`

	// Range GETs configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// Hotlink protection configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Origin-pull request header configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RequestHeader *RequestHeader `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// Origin server response header configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Follows origin server cache header configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitempty" name:"ResponseHeaderCache"`

	// SEO configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Seo *Seo `json:"Seo,omitempty" name:"Seo"`

	// Domain name service type. Values: `web` (static acceleration); `download` (download acceleration); `media` (streaming media acceleration).
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Status code cache configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Video dragging configuration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VideoSeek *VideoSeek `json:"VideoSeek,omitempty" name:"VideoSeek"`

	// Access authentication for S3 origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitempty" name:"AwsPrivateAccess"`

	// Access authentication for OSS origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitempty" name:"OssPrivateAccess"`

	// Access authentication for Huawei Cloud OBS origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitempty" name:"HwPrivateAccess"`

	// Access authentication for QiNiu Cloud Kodo origin
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitempty" name:"QnPrivateAccess"`
}

type PathBasedOriginRule struct {
	// Rule types:
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	// `index`: Apply to specified homepages.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each `RuleType`:
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// For `index`, enter a forward slash `/`.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// Origin server list. Domain name and IPv4 addresses are supported.
	Origin []*string `json:"Origin,omitempty" name:"Origin"`
}

type PathRule struct {
	// Whether to enable wildcard match (`*`).
	// `false`: disabled
	// `true`: enabled
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Regex *bool `json:"Regex,omitempty" name:"Regex"`

	// Matched URL. Only URLs are supported, while parameters are not. The exact match is used by default. If wildcard match is enabled, up to 5 wildcards are supported. The URL can contain up to 1,024 characters.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Origin server when the path matches. COS origin with private read/write is not supported. The default origin server will be used by default when this field is left empty.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// Origin server host header when the path matches. The default `ServerName` will be used by default when this field is left empty.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// Region of the origin server. Valid values: `CN` and `OV`.
	// `CN`: Within the Chinese mainland
	// `OV`: Outside the Chinese mainland
	// Default value: `CN`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	OriginArea *string `json:"OriginArea,omitempty" name:"OriginArea"`

	// Origin server URI path when the path matches, starting with `/` and excluding parameters. The path can contain up to 1,024 characters. The wildcards in the match path can be respectively captured using `$1`, `$2`, `$3`, `$4`, and `$5`. Up to 10 values can be captured.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ForwardUri *string `json:"ForwardUri,omitempty" name:"ForwardUri"`

	// Origin-pull header setting when the path matches.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RequestHeaders []*HttpHeaderRule `json:"RequestHeaders,omitempty" name:"RequestHeaders"`

	// When `Regex` is `false`, this parameter should be `true`.
	// `false`: Disabled
	// `true`: enabled
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FullMatch *bool `json:"FullMatch,omitempty" name:"FullMatch"`
}

type PostSize struct {
	// Limit the size of POST requests. The default value is 32 MB.
	// off: Disable
	// on: Enable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Maximum size. Value range: 1 MB to 200 MB.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

// Predefined struct for user
type PurgePathCacheRequestParams struct {
	// List of directories. The protocol header such as "http://" or "https://" needs to be included.
	Paths []*string `json:"Paths,omitempty" name:"Paths"`

	// Purge type:
	// `flush`: purges updated resources
	// `delete`: purges all resources
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// Whether to encode Chinese characters before purge.
	UrlEncode *bool `json:"UrlEncode,omitempty" name:"UrlEncode"`

	// Region to purge
	// The acceleration region of the acceleration domain name will be purged if this parameter is not passed in.
	// If `mainland` is passed in, only the content cached on nodes in the Chinese mainland will be purged.
	// If `overseas` is passed in, only the content cached on nodes outside the Chinese mainland will be purged.
	// The specified region to purge should match the domain name’s acceleration region.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest
	
	// List of directories. The protocol header such as "http://" or "https://" needs to be included.
	Paths []*string `json:"Paths,omitempty" name:"Paths"`

	// Purge type:
	// `flush`: purges updated resources
	// `delete`: purges all resources
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// Whether to encode Chinese characters before purge.
	UrlEncode *bool `json:"UrlEncode,omitempty" name:"UrlEncode"`

	// Region to purge
	// The acceleration region of the acceleration domain name will be purged if this parameter is not passed in.
	// If `mainland` is passed in, only the content cached on nodes in the Chinese mainland will be purged.
	// If `overseas` is passed in, only the content cached on nodes outside the Chinese mainland will be purged.
	// The specified region to purge should match the domain name’s acceleration region.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Paths")
	delete(f, "FlushType")
	delete(f, "UrlEncode")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurgePathCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurgePathCacheResponseParams struct {
	// Purge task ID. Directories submitted in one request share a task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *PurgePathCacheResponseParams `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PurgeTask struct {
	// Purge task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Purged URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// Purge task status
	// `fail`: Purge failed
	// `done`: Purge succeeded
	// `process`: Purge in progress
	Status *string `json:"Status,omitempty" name:"Status"`

	// Purge type
	// `url`: URL purge
	// `path`: directory purge
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// Purge method
	// `flush`: purges updated resources; this type is available only for directory purges
	// `delete`: Purge all resources
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// Purge task submission time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type PurgeUrlsCacheRequestParams struct {
	// List of URLs. The protocol header such as "http://" or "https://" needs to be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// Purging region
	// The acceleration region of the acceleration domain name will be purged if this parameter is not passed in.
	// If `mainland` is passed in, only the content cached on nodes in the Chinese mainland will be purged.
	// If `overseas` is passed in, only the content cached on nodes outside the Chinese mainland will be purged.
	// The specified purging region should match the domain name acceleration region.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether to encode Chinese characters for purge
	UrlEncode *bool `json:"UrlEncode,omitempty" name:"UrlEncode"`
}

type PurgeUrlsCacheRequest struct {
	*tchttp.BaseRequest
	
	// List of URLs. The protocol header such as "http://" or "https://" needs to be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// Purging region
	// The acceleration region of the acceleration domain name will be purged if this parameter is not passed in.
	// If `mainland` is passed in, only the content cached on nodes in the Chinese mainland will be purged.
	// If `overseas` is passed in, only the content cached on nodes outside the Chinese mainland will be purged.
	// The specified purging region should match the domain name acceleration region.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether to encode Chinese characters for purge
	UrlEncode *bool `json:"UrlEncode,omitempty" name:"UrlEncode"`
}

func (r *PurgeUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgeUrlsCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "Area")
	delete(f, "UrlEncode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurgeUrlsCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurgeUrlsCacheResponseParams struct {
	// Purge task ID. URLs submitted in one request share a task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PurgeUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *PurgeUrlsCacheResponseParams `json:"Response"`
}

func (r *PurgeUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgeUrlsCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PushTask struct {
	// Prefetch task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Prefetched URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// Prefetch task status
	// `fail`: Prefetch failed
	// `done`: Prefetch succeeded
	// `process`: Prefetch in progress
	// `invalid`: Invalid prefetch with 4XX/5XX status code returned from the origin server
	Status *string `json:"Status,omitempty" name:"Status"`

	// Prefetch progress in percentage
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// Prefetch task submission time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Prefetch region
	// `mainland`: Within the Chinese mainland
	// `overseas`: Outside the Chinese mainland
	// `global`: Globe
	Area *string `json:"Area,omitempty" name:"Area"`

	// Prefetch task update time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type PushUrlsCacheRequestParams struct {
	// List of URLs. The protocol header such as "http://" or "https://" needs to be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// Specifies the User-Agent header of an HTTP prefetch request when it is forwarded to the origin server
	// Default value: `TencentCdn`
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Destination region for the prefetch
	// `mainland`: prefetches resources to nodes within Mainland China
	// `overseas`: prefetches resources to nodes outside Mainland China
	// `global`: prefetches resources to global nodes
	// Default value: `mainland`. You can prefetch a URL to nodes in a region provided that CDN service has been enabled for the domain name in the URL in the region.
	Area *string `json:"Area,omitempty" name:"Area"`

	// By default, prefetch for regions in the Chinese mainland is performed onto the intermediate nodes, while prefetch for regions outside the Chinese mainland is performed onto the edge nodes and the traffic generated will be billed.
	// If this parameter is `middle` or left empty, prefetch will be performed onto the intermediate node.
	Layer *string `json:"Layer,omitempty" name:"Layer"`

	// Whether to recursively resolve the M3U8 index file and prefetch the TS shards in it.
	// Notes:
	// 1. This feature requires that the M3U8 index file can be directly requested and obtained.
	// 2. In the M3U8 index file, currently only the TS shards at the first to the third level can be recursively resolved.
	// 3. Prefetching the TS shards obtained through recursive resolution consumes the daily prefetch quota. If the usage exceeds the quota, the feature will be disabled and TS shards will not be prefetched.
	ParseM3U8 *bool `json:"ParseM3U8,omitempty" name:"ParseM3U8"`

	// Specifies whether to disable Range GETs.
	// Notes:
	// This feature is in beta test.
	DisableRange *bool `json:"DisableRange,omitempty" name:"DisableRange"`
}

type PushUrlsCacheRequest struct {
	*tchttp.BaseRequest
	
	// List of URLs. The protocol header such as "http://" or "https://" needs to be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// Specifies the User-Agent header of an HTTP prefetch request when it is forwarded to the origin server
	// Default value: `TencentCdn`
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Destination region for the prefetch
	// `mainland`: prefetches resources to nodes within Mainland China
	// `overseas`: prefetches resources to nodes outside Mainland China
	// `global`: prefetches resources to global nodes
	// Default value: `mainland`. You can prefetch a URL to nodes in a region provided that CDN service has been enabled for the domain name in the URL in the region.
	Area *string `json:"Area,omitempty" name:"Area"`

	// By default, prefetch for regions in the Chinese mainland is performed onto the intermediate nodes, while prefetch for regions outside the Chinese mainland is performed onto the edge nodes and the traffic generated will be billed.
	// If this parameter is `middle` or left empty, prefetch will be performed onto the intermediate node.
	Layer *string `json:"Layer,omitempty" name:"Layer"`

	// Whether to recursively resolve the M3U8 index file and prefetch the TS shards in it.
	// Notes:
	// 1. This feature requires that the M3U8 index file can be directly requested and obtained.
	// 2. In the M3U8 index file, currently only the TS shards at the first to the third level can be recursively resolved.
	// 3. Prefetching the TS shards obtained through recursive resolution consumes the daily prefetch quota. If the usage exceeds the quota, the feature will be disabled and TS shards will not be prefetched.
	ParseM3U8 *bool `json:"ParseM3U8,omitempty" name:"ParseM3U8"`

	// Specifies whether to disable Range GETs.
	// Notes:
	// This feature is in beta test.
	DisableRange *bool `json:"DisableRange,omitempty" name:"DisableRange"`
}

func (r *PushUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlsCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "UserAgent")
	delete(f, "Area")
	delete(f, "Layer")
	delete(f, "ParseM3U8")
	delete(f, "DisableRange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushUrlsCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushUrlsCacheResponseParams struct {
	// ID of the submitted task
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PushUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *PushUrlsCacheResponseParams `json:"Response"`
}

func (r *PushUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlsCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QnPrivateAccess struct {
	// Switch. Valid values: `on`, `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Access ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// Key
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type QueryStringKey struct {
	// Whether to use `QueryString` as part of `CacheKey`. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Whether to sort again
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Reorder *string `json:"Reorder,omitempty" name:"Reorder"`

	// Include/exclude query parameters. Valid values: `includeAll`, `excludeAll`, `includeCustom`, `excludeAll`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Array of included/excluded query strings (separated by ';')
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Quic struct {
	// Whether to enable QUIC
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Quota struct {
	// Quota limit for one batch submission request.
	Batch *int64 `json:"Batch,omitempty" name:"Batch"`

	// Daily submission quota limit.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Remaining daily submission quota.
	Available *int64 `json:"Available,omitempty" name:"Available"`

	// Quota region.
	Area *string `json:"Area,omitempty" name:"Area"`
}

type RangeOriginPull struct {
	// Range GETs configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Range GETs configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RangeRules []*RangeOriginPullRule `json:"RangeRules,omitempty" name:"RangeRules"`
}

type RangeOriginPullRule struct {
	// Specifies whether Range GETs is enabled
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Rule types:
	// `file`: effective for specified file suffixes.
	// `directory`: effective for specified paths.
	// `path`: effective for specified absolute paths.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each `RuleType`:
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`
}

type RedirectConfig struct {

	Switch *string `json:"Switch,omitempty" name:"Switch"`


	FollowRedirectHost *string `json:"FollowRedirectHost,omitempty" name:"FollowRedirectHost"`


	FollowRedirectBackupHost *string `json:"FollowRedirectBackupHost,omitempty" name:"FollowRedirectBackupHost"`
}

type Referer struct {
	// Referer blacklist/whitelist configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Referer blacklist/whitelist configuration rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RefererRules []*RefererRule `json:"RefererRules,omitempty" name:"RefererRules"`
}

type RefererRule struct {
	// Rule types:
	// `all`: Apply to all files.
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each `RuleType`:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// Referer configuration types
	// `whitelist`: Allowlist
	// `blacklist`: Blocklist
	RefererType *string `json:"RefererType,omitempty" name:"RefererType"`

	// Referer content list
	Referers []*string `json:"Referers,omitempty" name:"Referers"`

	// Whether to allow empty referer
	// `true`: Allow empty referer when `RefererType = whitelist`.
	// `false`: Reject empty refer when `RefererType = blacklist`.
	AllowEmpty *bool `json:"AllowEmpty,omitempty" name:"AllowEmpty"`
}

type RegionMapRelation struct {
	// Region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// List of sub-region IDs
	SubRegionIdList []*int64 `json:"SubRegionIdList,omitempty" name:"SubRegionIdList"`
}

type RemoteAuthentication struct {
	// Remote authentication switch
	// `on`: enable
	// `off`: disable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Remote authentication rule configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RemoteAuthenticationRules []*RemoteAuthenticationRule `json:"RemoteAuthenticationRules,omitempty" name:"RemoteAuthenticationRules"`

	// Remote authentication server
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Server *string `json:"Server,omitempty" name:"Server"`
}

type RemoteAuthenticationRule struct {
	// Remote authentication server
	// The server configured in `RemoteAutherntication` is used by default.
	Server *string `json:"Server,omitempty" name:"Server"`

	// HTTP method used by the remote authentication server. Valid values: `get`, `post`, `head`, and `all`. 
	// `all`: the remote authentication server follows the client request method.
	// Default: `all`
	AuthMethod *string `json:"AuthMethod,omitempty" name:"AuthMethod"`

	// Rule types:
	// `all`: apply to all files
	// `file`: apply to files with the specified suffixes
	// `directory`: apply to the specified directories
	// `path`: apply to the specified absolute paths
	// Default: `all`.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Content for each `RuleType`:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// For `index`, enter a forward slash `/`.
	// Default: `*`
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// Timeout period of the remote authentication server. Unit: ms.
	// Value range: [1, 30,000]
	// Default: 20000
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" name:"AuthTimeout"`

	// Whether to deny or allow the request when the remote authentication server is timed out:
	// `RETURN_200`: the request is allowed when the remote authentication server is timed out.
	// `RETURN_403`: the request is denied when the remote authentication server is timed out.
	// Default: `RETURN_200`
	AuthTimeoutAction *string `json:"AuthTimeoutAction,omitempty" name:"AuthTimeoutAction"`
}

type ReportData struct {
	// Project ID/domain name ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Project name/domain name.
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Total traffic/max bandwidth in bytes and bps, respectively.
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// Percentage of individual resource out of all resources.
	Percentage *float64 `json:"Percentage,omitempty" name:"Percentage"`

	// Total billable traffic/max billable bandwidth in bytes and bps, respectively.
	BillingValue *int64 `json:"BillingValue,omitempty" name:"BillingValue"`

	// Percentage of billable amount out of total amount.
	BillingPercentage *float64 `json:"BillingPercentage,omitempty" name:"BillingPercentage"`
}

type RequestHeader struct {
	// Custom request header configuration switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom request header configuration rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules"`
}

type ResourceBillingData struct {
	// Resource name, which is classified as follows based on different query filters:
	// When a domain name is specified: Details of the domain name
	// `multiDomains`: Aggregated details of multiple domain names
	// A specific project ID: ID of the specifically queried project
	// `all`: Details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Billing data details
	BillingData []*CdnData `json:"BillingData,omitempty" name:"BillingData"`
}

type ResourceData struct {
	// Resource name. 
	// A single domain name: Queries domain name details by a domain name. The details of the domain name will be displayed when the passed parameter `detail` is `true`.
	// Multiple domain names: Queries domain name details by multiple domain names. The aggregated details of the domain names will be displayed.
	// Project ID: Queries domain name details by a project ID. The aggregated details of the domain names of the project will be displayed.
	// `all`: Account-level data, which is aggregated details of all domain names of an account.
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Data details of a resource
	CdnData []*CdnData `json:"CdnData,omitempty" name:"CdnData"`
}

type ResourceOriginData struct {
	// Resource name, which is classified as follows based on different query filters:
	// A specific domain name: Details of the domain name
	// `multiDomains`: Aggregated details of multiple domain names
	// Project ID: ID of the specifically queried project
	// `all`: Details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Origin-pull data details
	OriginData []*CdnData `json:"OriginData,omitempty" name:"OriginData"`
}

type ResponseHeader struct {
	// Custom response header switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom response header rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules"`
}

type ResponseHeaderCache struct {
	// Origin server header cache switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Revalidate struct {
	// Whether to always forward to the origin server for verification. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Forwards to the origin server for verification only for specific request path
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`
}

type RuleCache struct {
	// Content for each `CacheType`:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter the suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter the path, e.g., `/xxx/test/`.
	// For `path`, enter the absolute path, e.g., `/xxx/test.html`.
	// For `index`, enter a forward slash `/`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// Rule types:
	// `all`: effective for all files.
	// `file`: effective for specified file suffixes.
	// `directory`: effective for specified paths.
	// `path`: effective for specified absolute paths.
	// `index`: homepage.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheConfig *RuleCacheConfig `json:"CacheConfig,omitempty" name:"CacheConfig"`
}

type RuleCacheConfig struct {
	// Cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Cache *CacheConfigCache `json:"Cache,omitempty" name:"Cache"`

	// No cache configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NoCache *CacheConfigNoCache `json:"NoCache,omitempty" name:"NoCache"`

	// Follows the origin server configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowOrigin *CacheConfigFollowOrigin `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type RuleEngine struct {
	// Specifies whether to enable rule engine
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Rule
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitempty" name:"Content"`
}

type RuleQueryString struct {
	// Whether to use `QueryString` as part of `CacheKey`. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// `includeCustom` will retain partial query strings
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Array of included/excluded query strings (separated by ';')
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ScdnAclConfig struct {
	// Whether to enable. Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// This field is disused. Please use `AdvancedScriptData` instead.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ScriptData []*ScdnAclGroup `json:"ScriptData,omitempty" name:"ScriptData"`

	// Error page configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// ACL rule group, which is required when the access control is on.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AdvancedScriptData []*AdvancedScdnAclGroup `json:"AdvancedScriptData,omitempty" name:"AdvancedScriptData"`
}

type ScdnAclGroup struct {
	// Rule name
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Specific configurations
	Configure []*ScdnAclRule `json:"Configure,omitempty" name:"Configure"`

	// Action. Valid values: `intercept` and `redirect`.
	Result *string `json:"Result,omitempty" name:"Result"`

	// Whether the rule is activated. Valid values: `active` and `inactive`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Error page configuration
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`
}

type ScdnAclRule struct {
	// Keyword
	MatchKey *string `json:"MatchKey,omitempty" name:"MatchKey"`

	// Logical operator. Valid values:
	LogiOperator *string `json:"LogiOperator,omitempty" name:"LogiOperator"`

	// Matched value
	MatchValue *string `json:"MatchValue,omitempty" name:"MatchValue"`
}

type ScdnBotConfig struct {
	// Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Bot cookie policy
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BotCookie []*BotCookie `json:"BotCookie,omitempty" name:"BotCookie"`

	// Bot JS policy
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BotJavaScript []*BotJavaScript `json:"BotJavaScript,omitempty" name:"BotJavaScript"`
}

type ScdnCCRules struct {
	// Rule types:
	// `all`: effective for all files.
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	// `index`: effective for web homepages and root directories.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Rule value (blocking condition)
	RuleValue []*string `json:"RuleValue,omitempty" name:"RuleValue"`

	// IP access limit rule
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Detection granularity
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DetectionTime *uint64 `json:"DetectionTime,omitempty" name:"DetectionTime"`

	// Frequency threshold
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FrequencyLimit *uint64 `json:"FrequencyLimit,omitempty" name:"FrequencyLimit"`

	// Whether to block or redirect requests from suspicious IPs. Valid values: `on` and `off`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PunishmentSwitch *string `json:"PunishmentSwitch,omitempty" name:"PunishmentSwitch"`

	// Suspicious IP restriction duration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PunishmentTime *uint64 `json:"PunishmentTime,omitempty" name:"PunishmentTime"`

	// Action. Valid values: `intercept` and `redirect`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitempty" name:"Action"`

	// The redirection target URL used when the `Action` is `redirect`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

type ScdnConfig struct {
	// Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom CC attack defense rule
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Rules []*ScdnCCRules `json:"Rules,omitempty" name:"Rules"`

	// Advanced custom CC attack defense rule
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AdvancedRules []*AdvancedCCRules `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
}

type ScdnDdosConfig struct {
	// Whether to enable DDoS defense. Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ScdnErrorPage struct {
	// Status code
	// `403` is passed in when the action is `intercept`.
	// `301` is passed in when the action is `redirect`.
	RedirectCode *int64 `json:"RedirectCode,omitempty" name:"RedirectCode"`

	// URL to be redirected
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

type ScdnSevenLayerRules struct {
	// Whether values are case sensitive
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// Rule types:
	// `protocol`: protocol. Valid values: `HTTP` and `HTTPS`.
	// `method`: request method. Valid values: `HEAD`, `GET`, `POST`, `PUT`, `OPTIONS`, `TRACE`, `DELETE`, `PATCH` and `CONNECT`.
	// `all`: domain name. The matching content is `*` and cannot be edited.
	// `ip`: IP in CIDR format.
	// `directory`: path starting with a slash (/). You can specify a directory or specific path using up to 128 characters.
	// `index`: default homepage, which is specified by `/;/index.html` and cannot be edited.
	// `path`: full path of the file, such as `/acb/test.png`. Wildcard is supported, such as `/abc/*.jpg`.
	// `file`: file extension, such as `jpg`, `png` and `css`.
	// `param`: request parameter. The value can contain up to 512 characters.
	// `referer`: Referer. The value can contain up to 512 characters.
	// `cookie`: Cookie. The value can contain up to 512 characters.
	// `user-agent`: User-Agent. The value can contain up to 512 characters.
	// `head`: custom header. The value can contain up to 512 characters. If the matching content is blank or does not exist, enter the matching parameter directly.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Logical operator, which connects the relation between RuleType and RuleValue. Valid values:
	// `exclude`: the rule value is not contained. 
	// `include`: the rule value is contained. 
	// `notequal`: the rule value is not equal to the specified rule type. 
	// `equal`: the rule value is equal to the specified rule type. 
	// `matching`: the rule value matches with the prefix of the specified rule type.
	// `null`: the rule value is empty or does not exist.
	LogicOperator *string `json:"LogicOperator,omitempty" name:"LogicOperator"`

	// Rule value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RuleValue []*string `json:"RuleValue,omitempty" name:"RuleValue"`

	// Matched parameter. Only request parameters, Cookie, and custom request headers have a value.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RuleParam *string `json:"RuleParam,omitempty" name:"RuleParam"`
}

type ScdnWafConfig struct {
	// Whether to enable WAF. Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// WAF protection mode. Valid values: `intercept` and `observe`. Default value: `intercept`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Redirection error page
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitempty" name:"ErrorPage"`

	// Whether to enable Web shell blocking. Valid values: `on` and `off`. Default value: `off`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WebShellSwitch *string `json:"WebShellSwitch,omitempty" name:"WebShellSwitch"`

	// Attack blocking rules
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Rules []*ScdnWafRule `json:"Rules,omitempty" name:"Rules"`

	// WAF rule level. Valid values: 100, 200, and 300.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// WAF sub-rule switch
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubRuleSwitch []*WafSubRuleStatus `json:"SubRuleSwitch,omitempty" name:"SubRuleSwitch"`
}

type ScdnWafRule struct {
	// Attack type
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// Defense action. Valid value: `observe`.
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

type SchemeKey struct {
	// Whether to use the scheme as part of the cache key. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

// Predefined struct for user
type SearchClsLogRequestParams struct {
	// ID of logset to be queried
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// List of IDs of log topics to be queried, separated by commas
	TopicIds *string `json:"TopicIds,omitempty" name:"TopicIds"`

	// Query start time in the format of YYYY-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of YYYY-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of logs to be returned at a time. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Query statement. For more details, see [https://intl.cloud.tencent.com/document/product/614/16982?from_cn_redirect=1].
	Query *string `json:"Query,omitempty" name:"Query"`

	// This field is used when loading more results. Pass through the last `context` value returned to get more log content. Up to 10,000 logs can be obtained through the cursor. Please narrow down the time range as much as possible.
	Context *string `json:"Context,omitempty" name:"Context"`

	// Sorting by log time. Valid values: asc (ascending), desc (descending). Default value: desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type SearchClsLogRequest struct {
	*tchttp.BaseRequest
	
	// ID of logset to be queried
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// List of IDs of log topics to be queried, separated by commas
	TopicIds *string `json:"TopicIds,omitempty" name:"TopicIds"`

	// Query start time in the format of YYYY-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of YYYY-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of logs to be returned at a time. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Specifies whether to access CDN or ECDN. Valid values: `cdn` (default) and `ecdn`.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Query statement. For more details, see [https://intl.cloud.tencent.com/document/product/614/16982?from_cn_redirect=1].
	Query *string `json:"Query,omitempty" name:"Query"`

	// This field is used when loading more results. Pass through the last `context` value returned to get more log content. Up to 10,000 logs can be obtained through the cursor. Please narrow down the time range as much as possible.
	Context *string `json:"Context,omitempty" name:"Context"`

	// Sorting by log time. Valid values: asc (ascending), desc (descending). Default value: desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *SearchClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Channel")
	delete(f, "Query")
	delete(f, "Context")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogResponseParams struct {
	// Query results
	Logs *ClsSearchLogs `json:"Logs,omitempty" name:"Logs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchClsLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchClsLogResponseParams `json:"Response"`
}

func (r *SearchClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityConfig struct {
	// on|off
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type Seo struct {
	// SEO configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ServerCert struct {
	// Server certificate ID, which is auto-generated when the certificate is being managed by the SSL Certificate Service
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Server certificate name
	// This is auto-generated when the certificate is being hosted by the SSL Certificate Service
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Server certificate information
	// This is required when uploading an external certificate, which should contain the complete certificate chain.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Server key information
	// This is required when uploading an external certificate.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// Time when the certificate expires
	// Can be left blank when used as an input parameter
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Certificate issuance time
	// Can be left blank when used as an input parameter
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// Certificate remarks
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Certificate source
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	From *string `json:"From,omitempty" name:"From"`
}

type ShareCname struct {
	// Specifies whether to enable Shared CNAME. If it is set to `off`, the default CNAME is used. If it is set to `on`, a shared CNAME is used.
	// 
	// * ShareCname is only available to beta users. To use this feature, please submit a ticket for application.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Shared CNAME to be configured
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`
}

type SimpleCache struct {
	// Cache expiration time rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheRules []*SimpleCacheRule `json:"CacheRules,omitempty" name:"CacheRules"`

	// Follows origin server Cache-Control: max-age configurations
	// `on`: Enable
	// `off`: Disable
	// If this is enabled, resources that do not match CacheRules rules will be cached by the node according to the max-age value returned by the origin server. Resources that match CacheRules rules will be cached on the node according to the cache expiration time set in CacheRules.
	// This conflicts with CompareMaxAge. The two cannot be enabled at the same time.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`

	// Forced cache
	// `on`: Enable
	// `off`: Disable
	// This is disabled by default. If enabled, the `no-store` and `no-cache` resources returned from the origin server will be cached according to `CacheRules` rules.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitempty" name:"IgnoreCacheControl"`

	// Ignores the Set-Cookie header of the origin server
	// `on`: Enable
	// `off`: Disable
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitempty" name:"IgnoreSetCookie"`

	// Advanced cache expiration configuration. If this is enabled, the max-age value returned by the origin server will be compared with the cache expiration time set in CacheRules, and the smallest value will be cached on the node.
	// `on`: Enable
	// `off`: Disable
	// It is disabled by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CompareMaxAge *string `json:"CompareMaxAge,omitempty" name:"CompareMaxAge"`

	// Always forwards to the origin server for verification
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Revalidate *Revalidate `json:"Revalidate,omitempty" name:"Revalidate"`
}

type SimpleCacheRule struct {
	// Rule types:
	// `all`: Apply to all files.
	// `file`: Apply to files with the specified suffixes.
	// `directory`: Apply to specified paths.
	// `path`: Apply to specified absolute paths.
	// index: home page
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// Content for each `CacheType`:
	// For `all`, enter a wildcard `*`.
	// For `file`, enter a suffix, e.g., `jpg` or `txt`.
	// For `directory`, enter a path, e.g., `/xxx/test/`.
	// For `path`, enter an absolute path, e.g., `/xxx/test.html`.
	// For `index`, enter a forward slash `/`.
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents"`

	// Cache expiration time settings
	// Unit: second. The maximum value is 365 days.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type Sort struct {
	// Fields that can be sorted. Currently supports:
	// `createTime`: domain name creation time.
	// `certExpireTime`: certificate expiration time.
	// Default value: createTime.
	Key *string `json:"Key,omitempty" name:"Key"`

	// asc/desc. Default value: desc.
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

type SpecificConfig struct {
	// Specific configuration for domain name inside mainland China.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mainland *MainlandConfig `json:"Mainland,omitempty" name:"Mainland"`

	// Specific configuration for domain name outside mainland China.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Overseas *OverseaConfig `json:"Overseas,omitempty" name:"Overseas"`
}

// Predefined struct for user
type StartCdnDomainRequestParams struct {
	// Domain name
	// The domain name status should be `Disabled`
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type StartCdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain name
	// The domain name status should be `Disabled`
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StartCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCdnDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StartCdnDomainResponseParams `json:"Response"`
}

func (r *StartCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticItem struct {
	// Type of usage limit. `total`: Cumulative usage; `moment`: Instantaneous usage.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Unblocking time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UnBlockTime *uint64 `json:"UnBlockTime,omitempty" name:"UnBlockTime"`

	// Bandwidth/Traffic threshold
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BpsThreshold *uint64 `json:"BpsThreshold,omitempty" name:"BpsThreshold"`

	// Specifies how to disable CDN service when the threshold is exceeded. `RETURN_404`: Return 404; `RESOLVE_DNS_TO_ORIGIN`: Forward to origin server.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CounterMeasure *string `json:"CounterMeasure,omitempty" name:"CounterMeasure"`

	// Threshold (in percentage) that triggers alarms
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AlertPercentage *uint64 `json:"AlertPercentage,omitempty" name:"AlertPercentage"`

	// Whether to enable the alarm threshold trigger. Values: `on`, `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AlertSwitch *string `json:"AlertSwitch,omitempty" name:"AlertSwitch"`

	// Metric type. `flux`: Traffic; `bandwidth`: Bandwidth.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Metric *string `json:"Metric,omitempty" name:"Metric"`


	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`

	// Whether to enable usage limit configuration. Values: `on`, `off`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type StatusCodeCache struct {
	// Status code cache expiration configuration switch
	// `on`: Enable
	// `off`: Disable
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Status code cache expiration rules details
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CacheRules []*StatusCodeCacheRule `json:"CacheRules,omitempty" name:"CacheRules"`
}

type StatusCodeCacheRule struct {
	// HTTP status code
	// Supports 403 and 404 status codes
	StatusCode *string `json:"StatusCode,omitempty" name:"StatusCode"`

	// Status code cache expiration time (in seconds)
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

// Predefined struct for user
type StopCdnDomainRequestParams struct {
	// Domain name
	// The domain name status should be **Enabled**
	Domain *string `json:"Domain,omitempty" name:"Domain"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCdnDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StopCdnDomainResponseParams `json:"Response"`
}

func (r *StopCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummarizedData struct {
	// Aggregation method, which can be:
	// `sum`: Aggregate summation
	// `max`: Maximum value. In bandwidth mode, the peak bandwidth is calculated based on the data aggregated in 5 minutes.
	// `avg`: Average value
	Name *string `json:"Name,omitempty" name:"Name"`

	// Aggregated value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type Tag struct {
	// Tag key
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TimestampData struct {
	// The start point of the sampling period. 
	// For example, if the time is set to 13:35:00, and `interval` is `5min`, the data returned is collected between 13:35:00 and 13:39:59
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
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData"`
}

type TopDetailData struct {
	// Datatype name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Data value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TopicInfo struct {
	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Whether to enable publishing
	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`

	// Creation time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Either `cdn` or `ecdn`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Whether the logset has been removed from CLS
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Deleted *string `json:"Deleted,omitempty" name:"Deleted"`
}

type TpgAdapter struct {
	// Switch. Valid values: `on`, `off`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

// Predefined struct for user
type UpdateDomainConfigRequestParams struct {
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Origin server configuration
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP blocklist/allowlist configuration
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Status code cache configuration
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Smart compression configuration
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Bandwidth cap configuration
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range GETs configuration
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 origin-pull follow-redirect configuration
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Error code redirect configuration (This feature is in beta and not generally available yet.)
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

	// Protocol redirect configuration
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer configuration
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Browser cache configuration (This feature is in beta and not generally available yet.)
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Specific-region special configuration
	// Applicable to cases where the acceleration domain name configuration differs for regions in and outside the Chinese mainland.
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// Domain name service type
	// `web`: Static acceleration
	// `download`: Download acceleration
	// `media`: Streaming media VOD acceleration
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Domain name acceleration region
	// `mainland`: Acceleration inside the Chinese mainland
	// `overseas`: Acceleration outside the Chinese mainland
	// `global`: Acceleration over the globe
	// After switching to global acceleration, configurations of the domain name will be deployed to the region inside or outside the Chinese mainland. The deployment will take some time as this domain name has special settings.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Origin-pull timeout configuration
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitempty" name:"OriginPullTimeout"`

	// Access authentication for S3 origin
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitempty" name:"AwsPrivateAccess"`

	// UA blocklist/allowlist configuration
	UserAgentFilter *UserAgentFilter `json:"UserAgentFilter,omitempty" name:"UserAgentFilter"`

	// Access control
	AccessControl *AccessControl `json:"AccessControl,omitempty" name:"AccessControl"`

	// URL rewriting configuration
	UrlRedirect *UrlRedirect `json:"UrlRedirect,omitempty" name:"UrlRedirect"`

	// Access port configuration
	AccessPort []*int64 `json:"AccessPort,omitempty" name:"AccessPort"`

	// Timestamp hotlink protection advanced configuration (allowlist feature)
	AdvancedAuthentication *AdvancedAuthentication `json:"AdvancedAuthentication,omitempty" name:"AdvancedAuthentication"`

	// Origin-pull authentication advanced configuration (allowlist feature)
	OriginAuthentication *OriginAuthentication `json:"OriginAuthentication,omitempty" name:"OriginAuthentication"`

	// IPv6 access configuration
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitempty" name:"Ipv6Access"`

	// Offline cache
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// Merging pull requests
	OriginCombine *OriginCombine `json:"OriginCombine,omitempty" name:"OriginCombine"`

	// Post transport configuration
	PostMaxSize *PostSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// QUIC access, which is a paid service. You can check the product document and Billing Overview for more information.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// Access authentication for OSS origin
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitempty" name:"OssPrivateAccess"`

	// WebSocket configuration
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Remote authentication configuration
	RemoteAuthentication *RemoteAuthentication `json:"RemoteAuthentication,omitempty" name:"RemoteAuthentication"`

	// Shared CNAME configuration (only available to beta users)
	ShareCname *ShareCname `json:"ShareCname,omitempty" name:"ShareCname"`

	// Access authentication for Huawei Cloud OBS origin
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitempty" name:"HwPrivateAccess"`

	// Access authentication for QiNiu Cloud Kodo origin
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitempty" name:"QnPrivateAccess"`
}

type UpdateDomainConfigRequest struct {
	*tchttp.BaseRequest
	
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Origin server configuration
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP blocklist/allowlist configuration
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Status code cache configuration
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitempty" name:"StatusCodeCache"`

	// Smart compression configuration
	Compression *Compression `json:"Compression,omitempty" name:"Compression"`

	// Bandwidth cap configuration
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitempty" name:"BandwidthAlert"`

	// Range GETs configuration
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitempty" name:"RangeOriginPull"`

	// 301/302 origin-pull follow-redirect configuration
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitempty" name:"FollowRedirect"`

	// Error code redirect configuration (This feature is in beta and not generally available yet.)
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

	// Protocol redirect configuration
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Referer configuration
	Referer *Referer `json:"Referer,omitempty" name:"Referer"`

	// Browser cache configuration (This feature is in beta and not generally available yet.)
	MaxAge *MaxAge `json:"MaxAge,omitempty" name:"MaxAge"`

	// Specific-region special configuration
	// Applicable to cases where the acceleration domain name configuration differs for regions in and outside the Chinese mainland.
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitempty" name:"SpecificConfig"`

	// Domain name service type
	// `web`: Static acceleration
	// `download`: Download acceleration
	// `media`: Streaming media VOD acceleration
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Domain name acceleration region
	// `mainland`: Acceleration inside the Chinese mainland
	// `overseas`: Acceleration outside the Chinese mainland
	// `global`: Acceleration over the globe
	// After switching to global acceleration, configurations of the domain name will be deployed to the region inside or outside the Chinese mainland. The deployment will take some time as this domain name has special settings.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Origin-pull timeout configuration
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitempty" name:"OriginPullTimeout"`

	// Access authentication for S3 origin
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitempty" name:"AwsPrivateAccess"`

	// UA blocklist/allowlist configuration
	UserAgentFilter *UserAgentFilter `json:"UserAgentFilter,omitempty" name:"UserAgentFilter"`

	// Access control
	AccessControl *AccessControl `json:"AccessControl,omitempty" name:"AccessControl"`

	// URL rewriting configuration
	UrlRedirect *UrlRedirect `json:"UrlRedirect,omitempty" name:"UrlRedirect"`

	// Access port configuration
	AccessPort []*int64 `json:"AccessPort,omitempty" name:"AccessPort"`

	// Timestamp hotlink protection advanced configuration (allowlist feature)
	AdvancedAuthentication *AdvancedAuthentication `json:"AdvancedAuthentication,omitempty" name:"AdvancedAuthentication"`

	// Origin-pull authentication advanced configuration (allowlist feature)
	OriginAuthentication *OriginAuthentication `json:"OriginAuthentication,omitempty" name:"OriginAuthentication"`

	// IPv6 access configuration
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitempty" name:"Ipv6Access"`

	// Offline cache
	OfflineCache *OfflineCache `json:"OfflineCache,omitempty" name:"OfflineCache"`

	// Merging pull requests
	OriginCombine *OriginCombine `json:"OriginCombine,omitempty" name:"OriginCombine"`

	// Post transport configuration
	PostMaxSize *PostSize `json:"PostMaxSize,omitempty" name:"PostMaxSize"`

	// QUIC access, which is a paid service. You can check the product document and Billing Overview for more information.
	Quic *Quic `json:"Quic,omitempty" name:"Quic"`

	// Access authentication for OSS origin
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitempty" name:"OssPrivateAccess"`

	// WebSocket configuration
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`

	// Remote authentication configuration
	RemoteAuthentication *RemoteAuthentication `json:"RemoteAuthentication,omitempty" name:"RemoteAuthentication"`

	// Shared CNAME configuration (only available to beta users)
	ShareCname *ShareCname `json:"ShareCname,omitempty" name:"ShareCname"`

	// Access authentication for Huawei Cloud OBS origin
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitempty" name:"HwPrivateAccess"`

	// Access authentication for QiNiu Cloud Kodo origin
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitempty" name:"QnPrivateAccess"`
}

func (r *UpdateDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "ProjectId")
	delete(f, "Origin")
	delete(f, "IpFilter")
	delete(f, "IpFreqLimit")
	delete(f, "StatusCodeCache")
	delete(f, "Compression")
	delete(f, "BandwidthAlert")
	delete(f, "RangeOriginPull")
	delete(f, "FollowRedirect")
	delete(f, "ErrorPage")
	delete(f, "RequestHeader")
	delete(f, "ResponseHeader")
	delete(f, "DownstreamCapping")
	delete(f, "CacheKey")
	delete(f, "ResponseHeaderCache")
	delete(f, "VideoSeek")
	delete(f, "Cache")
	delete(f, "OriginPullOptimization")
	delete(f, "Https")
	delete(f, "Authentication")
	delete(f, "Seo")
	delete(f, "ForceRedirect")
	delete(f, "Referer")
	delete(f, "MaxAge")
	delete(f, "SpecificConfig")
	delete(f, "ServiceType")
	delete(f, "Area")
	delete(f, "OriginPullTimeout")
	delete(f, "AwsPrivateAccess")
	delete(f, "UserAgentFilter")
	delete(f, "AccessControl")
	delete(f, "UrlRedirect")
	delete(f, "AccessPort")
	delete(f, "AdvancedAuthentication")
	delete(f, "OriginAuthentication")
	delete(f, "Ipv6Access")
	delete(f, "OfflineCache")
	delete(f, "OriginCombine")
	delete(f, "PostMaxSize")
	delete(f, "Quic")
	delete(f, "OssPrivateAccess")
	delete(f, "WebSocket")
	delete(f, "RemoteAuthentication")
	delete(f, "ShareCname")
	delete(f, "HwPrivateAccess")
	delete(f, "QnPrivateAccess")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDomainConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDomainConfigResponseParams `json:"Response"`
}

func (r *UpdateDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePayTypeRequestParams struct {
	// Billing region, which can be `mainland` or `overseas`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Billing mode, which can be `flux` or `bandwidth`.
	PayType *string `json:"PayType,omitempty" name:"PayType"`
}

type UpdatePayTypeRequest struct {
	*tchttp.BaseRequest
	
	// Billing region, which can be `mainland` or `overseas`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Billing mode, which can be `flux` or `bandwidth`.
	PayType *string `json:"PayType,omitempty" name:"PayType"`
}

func (r *UpdatePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePayTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Area")
	delete(f, "PayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePayTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePayTypeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdatePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePayTypeResponseParams `json:"Response"`
}

func (r *UpdatePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePayTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScdnDomainRequestParams struct {
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// WAF configuration
	Waf *ScdnWafConfig `json:"Waf,omitempty" name:"Waf"`

	// Custom defense policy configuration
	Acl *ScdnAclConfig `json:"Acl,omitempty" name:"Acl"`

	// CC attack defense configurations. CC attack defense is enabled by default.
	CC *ScdnConfig `json:"CC,omitempty" name:"CC"`

	// DDoS defense configuration. DDoS defense is enabled by default.
	Ddos *ScdnDdosConfig `json:"Ddos,omitempty" name:"Ddos"`

	// Bot defense configuration
	Bot *ScdnBotConfig `json:"Bot,omitempty" name:"Bot"`
}

type UpdateScdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// WAF configuration
	Waf *ScdnWafConfig `json:"Waf,omitempty" name:"Waf"`

	// Custom defense policy configuration
	Acl *ScdnAclConfig `json:"Acl,omitempty" name:"Acl"`

	// CC attack defense configurations. CC attack defense is enabled by default.
	CC *ScdnConfig `json:"CC,omitempty" name:"CC"`

	// DDoS defense configuration. DDoS defense is enabled by default.
	Ddos *ScdnDdosConfig `json:"Ddos,omitempty" name:"Ddos"`

	// Bot defense configuration
	Bot *ScdnBotConfig `json:"Bot,omitempty" name:"Bot"`
}

func (r *UpdateScdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Waf")
	delete(f, "Acl")
	delete(f, "CC")
	delete(f, "Ddos")
	delete(f, "Bot")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateScdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScdnDomainResponseParams struct {
	// Result of the request. `Success` indicates that the configurations are updated.
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateScdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *UpdateScdnDomainResponseParams `json:"Response"`
}

func (r *UpdateScdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UrlRecord struct {
	// Status. `disable`: Blocked; `enable`: Unblocked.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Corresponding URL
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RealUrl *string `json:"RealUrl,omitempty" name:"RealUrl"`

	// Creation time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UrlRedirect struct {
	// Whether to enable URL rewriting
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Rule of URL rewriting rule, which is required if `Switch` is `on`. There can be up to 10 rules.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PathRules []*UrlRedirectRule `json:"PathRules,omitempty" name:"PathRules"`
}

type UrlRedirectRule struct {
	// Redirect status code. Valid values: 301, 302
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`

	// URL to be matched. Only URLs are supported, while parameters are not. The exact match is used by default. In regex match, up to 5 wildcards `*` are supported. The URL can contain up to 1,024 characters.
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// Target URL, starting with `/` and excluding parameters. The path can contain up to 1,024 characters. The wildcards in the matching path can be respectively captured using `$1`, `$2`, `$3`, `$4`, and `$5`. Up to 10 values can be captured.
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// Target host. It should be a standard domain name starting with `http://` or `https://`. If it is left empty, “http://[current domain name]” will be used by default.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RedirectHost *string `json:"RedirectHost,omitempty" name:"RedirectHost"`

	// Whether to use full-path matching or arbitrary matching
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FullMatch *bool `json:"FullMatch,omitempty" name:"FullMatch"`
}

type UserAgentFilter struct {
	// Switch. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// UA blacklist/whitelist effect rule list
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FilterRules []*UserAgentFilterRule `json:"FilterRules,omitempty" name:"FilterRules"`
}

type UserAgentFilterRule struct {
	// Effective access path type
	// `all`: All access paths are effective
	// `file`: Effective by file extension
	// `directory`: Effective by directory
	// `path`: Effective by full access path
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Effective access paths
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`

	// `UserAgent` list
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UserAgents []*string `json:"UserAgents,omitempty" name:"UserAgents"`

	// Blocklist or allowlist. Valid values: `blacklist`, `whitelist`
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type VideoSeek struct {
	// Video dragging switch
	// `on`: Enable
	// `off`: Disable
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ViolationUrl struct {
	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Origin access URL of the resource in violation
	RealUrl *string `json:"RealUrl,omitempty" name:"RealUrl"`

	// Snapshot path. This is used to display a snapshot of the content in violation on the console.
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// Current status of the resources in violation
	// `forbid`: Blocked
	// `release`: Unblocked
	// `delay`: Processing delayed
	// `reject`: Appeal dismissed. The status is still blocked.
	// `complain`: Appeal in process
	UrlStatus *string `json:"UrlStatus,omitempty" name:"UrlStatus"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WafSubRuleStatus struct {
	// Sub-rule status. Valid values: `on` and `off`.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// List of rule IDs
	SubIds []*int64 `json:"SubIds,omitempty" name:"SubIds"`
}

type WebSocket struct {
	// Whether to enable custom WebSocket timeout setting. When it’s `off`: WebSocket connection is supported, and the default timeout period is 15 seconds. To change the timeout period, please set it to `on`.
	// 
	// * WebSocket is an ECDN feature. You can enable it in the ECDN domain name configuration.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Sets timeout period in seconds. Maximum value: 300
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type WebpAdapter struct {
	// Switch. Valid values: on, off
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}