// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20190722

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type CreateCaptchaInfoInternationalRequestParams struct {
	// <p>Captcha name</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Client type</p><p>Enumeration values:</p><ul><li>web: For web scenario</li><li>android: For Android client</li><li>ios: For iOS client</li></ul><p>Default value: web</p>
	ChannelInfo *string `json:"ChannelInfo,omitnil,omitempty" name:"ChannelInfo"`

	// <p>Verification level</p><p>Enumeration values:</p><ul><li>1: Experience-oriented</li><li>2: Balanced</li><li>3: Security-focused</li></ul><p>Default value: 1</p>
	VerifyRank *string `json:"VerifyRank,omitnil,omitempty" name:"VerifyRank"`

	// <p>Validation type</p><p>Enumeration values:</p><ul><li>1: Invisible verification. UserSetCapType input 1, DisableInvisibleSwitch must be 2</li><li>2: Slide</li><li>8: Graphical</li><li>9: Voice</li></ul>
	UserSetCapType *string `json:"UserSetCapType,omitnil,omitempty" name:"UserSetCapType"`

	// <p>Interception mode</p><p>Enumeration values:</p><ul><li>block: interception mode</li><li>notify: perception mode</li></ul><p>Default value: notify</p>
	DefendMode *string `json:"DefendMode,omitnil,omitempty" name:"DefendMode"`

	// <p>Resource tag, key&amp;value format</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Verification mechanism</p><p>Enumeration values:</p><ul><li>0: One-Click Verification</li><li>1: Always verify</li><li>2: Invisible verification. DisableInvisibleSwitch input 2, UserSetCapType must be 1</li></ul>
	DisableInvisibleSwitch *string `json:"DisableInvisibleSwitch,omitnil,omitempty" name:"DisableInvisibleSwitch"`

	// <p>web domain name</p><p>Only valid when ChannelInfo is web</p>
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// <p>app BundleId</p><p>Only valid when ChannelInfo is ios</p>
	VerifyBundleId *string `json:"VerifyBundleId,omitnil,omitempty" name:"VerifyBundleId"`

	// <p>app package</p><p>Only valid when ChannelInfo is android</p>
	VerifyPackage *string `json:"VerifyPackage,omitnil,omitempty" name:"VerifyPackage"`

	// <p>Whether to enable captcha encryption. 0: Off. 1: On</p>
	CheckAppidSwitch *int64 `json:"CheckAppidSwitch,omitnil,omitempty" name:"CheckAppidSwitch"`

	// <p>Whether to enable non-repeating IV</p><p>Enumeration values:</p><ul><li>0: Off</li><li>1: On</li></ul><p>Input 1 is allowed only when CheckAppidSwitch is 1</p>
	CheckIvSwitch *int64 `json:"CheckIvSwitch,omitnil,omitempty" name:"CheckIvSwitch"`

	// <p>Checkbox display method</p><p>Enumeration values:</p><ul><li>0: simplified version</li><li>1: basic version</li><li>2: invisible version</li></ul>
	CheckBoxStyle *string `json:"CheckBoxStyle,omitnil,omitempty" name:"CheckBoxStyle"`
}

type CreateCaptchaInfoInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>Captcha name</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Client type</p><p>Enumeration values:</p><ul><li>web: For web scenario</li><li>android: For Android client</li><li>ios: For iOS client</li></ul><p>Default value: web</p>
	ChannelInfo *string `json:"ChannelInfo,omitnil,omitempty" name:"ChannelInfo"`

	// <p>Verification level</p><p>Enumeration values:</p><ul><li>1: Experience-oriented</li><li>2: Balanced</li><li>3: Security-focused</li></ul><p>Default value: 1</p>
	VerifyRank *string `json:"VerifyRank,omitnil,omitempty" name:"VerifyRank"`

	// <p>Validation type</p><p>Enumeration values:</p><ul><li>1: Invisible verification. UserSetCapType input 1, DisableInvisibleSwitch must be 2</li><li>2: Slide</li><li>8: Graphical</li><li>9: Voice</li></ul>
	UserSetCapType *string `json:"UserSetCapType,omitnil,omitempty" name:"UserSetCapType"`

	// <p>Interception mode</p><p>Enumeration values:</p><ul><li>block: interception mode</li><li>notify: perception mode</li></ul><p>Default value: notify</p>
	DefendMode *string `json:"DefendMode,omitnil,omitempty" name:"DefendMode"`

	// <p>Resource tag, key&amp;value format</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Verification mechanism</p><p>Enumeration values:</p><ul><li>0: One-Click Verification</li><li>1: Always verify</li><li>2: Invisible verification. DisableInvisibleSwitch input 2, UserSetCapType must be 1</li></ul>
	DisableInvisibleSwitch *string `json:"DisableInvisibleSwitch,omitnil,omitempty" name:"DisableInvisibleSwitch"`

	// <p>web domain name</p><p>Only valid when ChannelInfo is web</p>
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// <p>app BundleId</p><p>Only valid when ChannelInfo is ios</p>
	VerifyBundleId *string `json:"VerifyBundleId,omitnil,omitempty" name:"VerifyBundleId"`

	// <p>app package</p><p>Only valid when ChannelInfo is android</p>
	VerifyPackage *string `json:"VerifyPackage,omitnil,omitempty" name:"VerifyPackage"`

	// <p>Whether to enable captcha encryption. 0: Off. 1: On</p>
	CheckAppidSwitch *int64 `json:"CheckAppidSwitch,omitnil,omitempty" name:"CheckAppidSwitch"`

	// <p>Whether to enable non-repeating IV</p><p>Enumeration values:</p><ul><li>0: Off</li><li>1: On</li></ul><p>Input 1 is allowed only when CheckAppidSwitch is 1</p>
	CheckIvSwitch *int64 `json:"CheckIvSwitch,omitnil,omitempty" name:"CheckIvSwitch"`

	// <p>Checkbox display method</p><p>Enumeration values:</p><ul><li>0: simplified version</li><li>1: basic version</li><li>2: invisible version</li></ul>
	CheckBoxStyle *string `json:"CheckBoxStyle,omitnil,omitempty" name:"CheckBoxStyle"`
}

func (r *CreateCaptchaInfoInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCaptchaInfoInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "ChannelInfo")
	delete(f, "VerifyRank")
	delete(f, "UserSetCapType")
	delete(f, "DefendMode")
	delete(f, "Tags")
	delete(f, "DisableInvisibleSwitch")
	delete(f, "VerifyDomain")
	delete(f, "VerifyBundleId")
	delete(f, "VerifyPackage")
	delete(f, "CheckAppidSwitch")
	delete(f, "CheckIvSwitch")
	delete(f, "CheckBoxStyle")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCaptchaInfoInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCaptchaInfoInternationalResponseParams struct {
	// <p>Result data.</p>
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Captcha status code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha information</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCaptchaInfoInternationalResponse struct {
	*tchttp.BaseResponse
	Response *CreateCaptchaInfoInternationalResponseParams `json:"Response"`
}

func (r *CreateCaptchaInfoInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCaptchaInfoInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIpWhiteListInternationalRequestParams struct {
	// <p>ip allowlist name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>ip data</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Remark information.</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateIpWhiteListInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>ip allowlist name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>ip data</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Remark information.</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateIpWhiteListInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIpWhiteListInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "CaptchaAppid")
	delete(f, "Ip")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIpWhiteListInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIpWhiteListInternationalResponseParams struct {
	// <p>Result data</p>
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>IP allowlist resource id</p>
	IdList []*int64 `json:"IdList,omitnil,omitempty" name:"IdList"`

	// <p>Captcha status code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha information</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIpWhiteListInternationalResponse struct {
	*tchttp.BaseResponse
	Response *CreateIpWhiteListInternationalResponseParams `json:"Response"`
}

func (r *CreateIpWhiteListInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIpWhiteListInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpWhiteListInternationalRequestParams struct {
	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>Record number</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteIpWhiteListInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>Record number</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteIpWhiteListInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpWhiteListInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppid")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIpWhiteListInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpWhiteListInternationalResponseParams struct {
	// <p>Result data</p>
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Captcha status code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha info</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIpWhiteListInternationalResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIpWhiteListInternationalResponseParams `json:"Response"`
}

func (r *DeleteIpWhiteListInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpWhiteListInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaConsoleDataInternational struct {
	// <p>Paginated data array.</p>
	DataList []*DescribeCaptchaConsoleSubDataInternational `json:"DataList,omitnil,omitempty" name:"DataList"`

	// <p>Total pages</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>Current page</p>
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`
}

type DescribeCaptchaConsoleSubDataInternational struct {
	// <p>Verification code id</p>
	CaptchaAppId *int64 `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// <p>Verification name</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Separate domain names with commas</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Random key</p>
	EncryptKey *string `json:"EncryptKey,omitnil,omitempty" name:"EncryptKey"`

	// <p>Verification scenario</p><p>Enumeration values:</p><ul><li>1: Account</li><li>2: SMS</li><li>3: Promotion</li><li>4: Comment</li><li>5: Data protection</li><li>6: Other</li></ul>
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// <p>Validation type</p><p>Enumeration values:</p><ul><li>1: Invisible verification. UserSetCapType input 1, DisableInvisibleSwitch must be 2</li><li>2: Sliding puzzle</li><li>8: Graphical point selection</li><li>9: Voice verification</li></ul>
	UserSetCapType *int64 `json:"UserSetCapType,omitnil,omitempty" name:"UserSetCapType"`

	// <p>Intelligent verification-free</p><p>Enumeration values:</p><ul><li>0: disable</li><li>1: enable</li></ul>
	NoVerifyRule *int64 `json:"NoVerifyRule,omitnil,omitempty" name:"NoVerifyRule"`

	// <p>Language</p><p>Enumeration values:</p><ul><li>1: Self adaptive</li><li>2052: Simplified</li><li>1028: Traditional</li><li>1033: English</li></ul>
	CaptchaLanguage *string `json:"CaptchaLanguage,omitnil,omitempty" name:"CaptchaLanguage"`

	// <p>Verification level</p><p>Enumeration values:</p><ul><li>1: Experience-oriented</li><li>2: Balanced</li><li>3: Security-focused</li></ul><p>Default value: 1</p>
	VerifyRank *int64 `json:"VerifyRank,omitnil,omitempty" name:"VerifyRank"`

	// <p>Client type</p><p>Enumeration values:</p><ul><li>web: For web scenario usage</li><li>android: For Android client usage</li><li>ios: For iOS client usage</li></ul>
	ChannelInfo *string `json:"ChannelInfo,omitnil,omitempty" name:"ChannelInfo"`

	// <p>Interception mode</p><p>Enumeration values:</p><ul><li>block: interception mode</li><li>notify: perception mode</li></ul><p>Default value: notify</p>
	DefendMode *string `json:"DefendMode,omitnil,omitempty" name:"DefendMode"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Whether to enable captchaAppid encryption</p><p>Enumeration values:</p><ul><li>0: Off</li><li>1: On</li></ul>
	CheckAppidSwitch *int64 `json:"CheckAppidSwitch,omitnil,omitempty" name:"CheckAppidSwitch"`

	// <p>Resource tag.</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Whether to enable non-repeating IV</p><p>Enumeration values:</p><ul><li>0: Disabled</li><li>1: Enabled</li></ul>
	CheckIvSwitch *int64 `json:"CheckIvSwitch,omitnil,omitempty" name:"CheckIvSwitch"`

	// <p>Verification mechanism</p><p>Enumeration values:</p><ul><li>0: One-Click Verification</li><li>1: Always verify</li><li>2: Invisible verification. DisableInvisibleSwitch input 2, UserSetCapType must be 1</li></ul>
	DisableInvisibleSwitch *string `json:"DisableInvisibleSwitch,omitnil,omitempty" name:"DisableInvisibleSwitch"`

	// <p>Web domain name</p><p>Valid only when ChannelInfo is web</p>
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// <p>app BundleId</p><p>Valid only when ChannelInfo is ios</p>
	VerifyBundleId *string `json:"VerifyBundleId,omitnil,omitempty" name:"VerifyBundleId"`

	// <p>app package</p><p>Only valid when ChannelInfo is android</p>
	VerifyPackage *string `json:"VerifyPackage,omitnil,omitempty" name:"VerifyPackage"`

	// <p>Checkbox display method</p><p>Enumeration values:</p><ul><li>0: simplified version</li><li>1: basic version</li><li>2: invisible version</li></ul>
	CheckBoxStyle *string `json:"CheckBoxStyle,omitnil,omitempty" name:"CheckBoxStyle"`

	// <p>Customer type</p><p>Enumeration values:</p><ul><li>0: General user</li><li>1: waf</li><li>2: EO</li></ul>
	CustomerType *string `json:"CustomerType,omitnil,omitempty" name:"CustomerType"`
}

// Predefined struct for user
type DescribeCaptchaInfoListInternationalRequestParams struct {
	// <p>Pagination parameter - page number</p>
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// <p>Pagination parameters - number of records per page</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query parameter - Behavior verification type</p><p>Enumeration values:</p><ul><li>1: Invisible verification</li><li>2: Slide verification</li><li>8: Graphical verification</li><li>9: Voice verification</li></ul>
	UserSetCapTypeArr []*string `json:"UserSetCapTypeArr,omitnil,omitempty" name:"UserSetCapTypeArr"`

	// <p>Query parameter - risk control level</p><p>Enumeration values:</p><ul><li>1: Experience-oriented</li><li>2: Balanced</li><li>3: Security-focused</li></ul>
	VerifyRankArr []*string `json:"VerifyRankArr,omitnil,omitempty" name:"VerifyRankArr"`

	// <p>Query parameter - client multiple selection</p><p>Enumeration values:</p><ul><li>web:</li><li>ios </li><li>android</li></ul>
	ChannelInfoArr []*string `json:"ChannelInfoArr,omitnil,omitempty" name:"ChannelInfoArr"`

	// <p>Query parameter -Captcha appid</p>
	CaptchaAppId *string `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// <p>Query parameter - Captcha name</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Sorting parameter</p><p>Input limits: desc: in descending order by creation time; asc: in ascending order by creation time</p>
	OrderBy *OrderByInternational `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeCaptchaInfoListInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>Pagination parameter - page number</p>
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// <p>Pagination parameters - number of records per page</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query parameter - Behavior verification type</p><p>Enumeration values:</p><ul><li>1: Invisible verification</li><li>2: Slide verification</li><li>8: Graphical verification</li><li>9: Voice verification</li></ul>
	UserSetCapTypeArr []*string `json:"UserSetCapTypeArr,omitnil,omitempty" name:"UserSetCapTypeArr"`

	// <p>Query parameter - risk control level</p><p>Enumeration values:</p><ul><li>1: Experience-oriented</li><li>2: Balanced</li><li>3: Security-focused</li></ul>
	VerifyRankArr []*string `json:"VerifyRankArr,omitnil,omitempty" name:"VerifyRankArr"`

	// <p>Query parameter - client multiple selection</p><p>Enumeration values:</p><ul><li>web:</li><li>ios </li><li>android</li></ul>
	ChannelInfoArr []*string `json:"ChannelInfoArr,omitnil,omitempty" name:"ChannelInfoArr"`

	// <p>Query parameter -Captcha appid</p>
	CaptchaAppId *string `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// <p>Query parameter - Captcha name</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Sorting parameter</p><p>Input limits: desc: in descending order by creation time; asc: in ascending order by creation time</p>
	OrderBy *OrderByInternational `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *DescribeCaptchaInfoListInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaInfoListInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "UserSetCapTypeArr")
	delete(f, "VerifyRankArr")
	delete(f, "ChannelInfoArr")
	delete(f, "CaptchaAppId")
	delete(f, "AppName")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaInfoListInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaInfoListInternationalResponseParams struct {
	// <p>Data block after paging query.</p>
	Data *DescribeCaptchaConsoleDataInternational `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Captcha response code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha information</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCaptchaInfoListInternationalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaInfoListInternationalResponseParams `json:"Response"`
}

func (r *DescribeCaptchaInfoListInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaInfoListInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaIpWhiteListDataNew struct {
	// <p>Data list.</p>
	DataList []*DescribeCaptchaWhiteListItem `json:"DataList,omitnil,omitempty" name:"DataList"`

	// <p>Total number of records</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>Page number.</p>
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`
}

// Predefined struct for user
type DescribeCaptchaResultRequestParams struct {
	// <p>Fixed value: 9.</p>
	CaptchaType *uint64 `json:"CaptchaType,omitnil,omitempty" name:"CaptchaType"`

	// <p>User verification ticket returned by the frontend callback function</p>
	Ticket *string `json:"Ticket,omitnil,omitempty" name:"Ticket"`

	// <p>Public network IP of the verification</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>Random string returned by the frontend callback function</p>
	Randstr *string `json:"Randstr,omitnil,omitempty" name:"Randstr"`

	// <p>Captcha appId. Log in to the <a href="https://console.cloud.tencent.com/captcha/graphical">verification code console</a>. In the [Key] column of the verification list, you can see CaptchaAppId.</p>
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// <p>Captcha application key. Log in to the <a href="https://console.cloud.tencent.com/captcha/graphical">verification code console</a>, and view AppSecretKey in the [Key] column of the verification list. AppSecretKey is a key for server-side verification of verification code tickets. Keep it confidential and do not leak it to third parties.</p>
	AppSecretKey *string `json:"AppSecretKey,omitnil,omitempty" name:"AppSecretKey"`

	// <p>Reserved field</p>
	BusinessId *uint64 `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// <p>Reserved field</p>
	SceneId *uint64 `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// <p>mac address or unique device identifier</p>
	MacAddress *string `json:"MacAddress,omitnil,omitempty" name:"MacAddress"`

	// <p>Mobile device number</p>
	Imei *string `json:"Imei,omitnil,omitempty" name:"Imei"`

	// <p>Whether to return the time when the frontend obtains the verification code. Value: 1: need to return</p>
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitnil,omitempty" name:"NeedGetCaptchaTime"`
}

type DescribeCaptchaResultRequest struct {
	*tchttp.BaseRequest
	
	// <p>Fixed value: 9.</p>
	CaptchaType *uint64 `json:"CaptchaType,omitnil,omitempty" name:"CaptchaType"`

	// <p>User verification ticket returned by the frontend callback function</p>
	Ticket *string `json:"Ticket,omitnil,omitempty" name:"Ticket"`

	// <p>Public network IP of the verification</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>Random string returned by the frontend callback function</p>
	Randstr *string `json:"Randstr,omitnil,omitempty" name:"Randstr"`

	// <p>Captcha appId. Log in to the <a href="https://console.cloud.tencent.com/captcha/graphical">verification code console</a>. In the [Key] column of the verification list, you can see CaptchaAppId.</p>
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// <p>Captcha application key. Log in to the <a href="https://console.cloud.tencent.com/captcha/graphical">verification code console</a>, and view AppSecretKey in the [Key] column of the verification list. AppSecretKey is a key for server-side verification of verification code tickets. Keep it confidential and do not leak it to third parties.</p>
	AppSecretKey *string `json:"AppSecretKey,omitnil,omitempty" name:"AppSecretKey"`

	// <p>Reserved field</p>
	BusinessId *uint64 `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// <p>Reserved field</p>
	SceneId *uint64 `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// <p>mac address or unique device identifier</p>
	MacAddress *string `json:"MacAddress,omitnil,omitempty" name:"MacAddress"`

	// <p>Mobile device number</p>
	Imei *string `json:"Imei,omitnil,omitempty" name:"Imei"`

	// <p>Whether to return the time when the frontend obtains the verification code. Value: 1: need to return</p>
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitnil,omitempty" name:"NeedGetCaptchaTime"`
}

func (r *DescribeCaptchaResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaType")
	delete(f, "Ticket")
	delete(f, "UserIp")
	delete(f, "Randstr")
	delete(f, "CaptchaAppId")
	delete(f, "AppSecretKey")
	delete(f, "BusinessId")
	delete(f, "SceneId")
	delete(f, "MacAddress")
	delete(f, "Imei")
	delete(f, "NeedGetCaptchaTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaResultResponseParams struct {
	// <p>1 OK verification passed<br>7 captcha no match The passed-in Randstr is invalid. Please check whether Randstr is consistent with the frontend Randstr.<br>8 ticket expired The passed-in ticket has expired (ticket valid period is 5 minutes). Please generate a new ticket and Randstr for verification.<br>9 ticket reused The passed-in ticket has been reused. Please generate a new ticket and Randstr for verification.<br>15 decrypt fail The passed-in ticket is invalid. Please check whether ticket is consistent with the frontend ticket.<br>16 appid-ticket mismatch The passed-in CaptchaAppId is incorrect. Please check whether CaptchaAppId is consistent with the frontend CaptchaAppId, and ensure that CaptchaAppId is obtained from [Verification Management] -> [Basic Configuration] in the verification code console.<br>21 diff Bill verification exception. Possible reasons: (1) If the ticket contains the trerror prefix, it is generally because the user has a poor network connection, which causes frontend automatic disaster recovery and generates a disaster recovery ticket. The business side can skip or post-process it based on needs. (2) If the ticket does not contain the trerror prefix, it is because the verification code risk control system has detected a security risk in the request. The business side can block it based on needs.<br>100 appid-secretkey-ticket mismatch Parameter validation error. (1) Please check whether CaptchaAppId and AppSecretKey are correct. CaptchaAppId and AppSecretKey need to be obtained from [Verification Management] > [Basic Configuration] in the verification code console. (2) Please check whether the passed-in ticket is generated by the passed-in CaptchaAppId.</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Description and error message</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// <p>In non-perception mode, this parameter returns the verification result:<br>EvilLevel=0: The request is not malicious<br>EvilLevel=100: The request is malicious</p>
	EvilLevel *int64 `json:"EvilLevel,omitnil,omitempty" name:"EvilLevel"`

	// <p>Frontend verification code retrieval time, Timestamp Format</p>
	GetCaptchaTime *int64 `json:"GetCaptchaTime,omitnil,omitempty" name:"GetCaptchaTime"`

	// <p>Interception type</p>
	EvilBitmap *int64 `json:"EvilBitmap,omitnil,omitempty" name:"EvilBitmap"`

	// <p>Time of submitting the verification code</p>
	SubmitCaptchaTime *int64 `json:"SubmitCaptchaTime,omitnil,omitempty" name:"SubmitCaptchaTime"`

	// <p>Device risk category</p>
	DeviceRiskCategory *string `json:"DeviceRiskCategory,omitnil,omitempty" name:"DeviceRiskCategory"`

	// <p>Verification code score</p><p>Value range: 0-100. Higher scores indicate higher risk</p>
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCaptchaResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaResultResponseParams `json:"Response"`
}

func (r *DescribeCaptchaResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCaptchaWhiteListItem struct {
	// <p>No.</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Allowlist name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Bind captcha</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>ip address</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Status. 0: Ip allowlisted; 1: cancel allowlisting</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Creation time.</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>Update time.</p>
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// <p>Remarks.</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

// Predefined struct for user
type DescribeIpWhiteListInternationalRequestParams struct {
	// <p>Page number.</p>
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// <p>Page length.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>Allowlist name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Ip address</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>IP Whitelist Configuration Status</p><p>Enumeration values:</p><ul><li>0: all</li><li>1: allowlisted</li><li>2: allowlisting canceled</li></ul><p>Default value: 0</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeIpWhiteListInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>Page number.</p>
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// <p>Page length.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>Allowlist name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Ip address</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>IP Whitelist Configuration Status</p><p>Enumeration values:</p><ul><li>0: all</li><li>1: allowlisted</li><li>2: allowlisting canceled</li></ul><p>Default value: 0</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeIpWhiteListInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpWhiteListInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "CaptchaAppid")
	delete(f, "Name")
	delete(f, "Ip")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpWhiteListInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpWhiteListInternationalResponseParams struct {
	// <p>Result data</p>
	Data *DescribeCaptchaIpWhiteListDataNew `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Captcha status code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha information</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIpWhiteListInternationalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpWhiteListInternationalResponseParams `json:"Response"`
}

func (r *DescribeIpWhiteListInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpWhiteListInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCaptchaInfoInternationalRequestParams struct {
	// <p>Captcha appid</p>
	CaptchaAppId *string `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// <p>Captcha name</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Verification level</p><p>Enumeration values:</p><ul><li>1: Experience-oriented</li><li>2: Balanced</li><li>3: Security-focused</li></ul><p>Default value: 1</p>
	VerifyRank *string `json:"VerifyRank,omitnil,omitempty" name:"VerifyRank"`

	// <p>Verification method</p><p>Enumeration values:</p><ul><li>1: Invisible verification. UserSetCapType input 1, DisableInvisibleSwitch must</li><li>2: Slide verification</li><li>8: Graphical verification</li><li>9: Voice verification</li></ul>
	UserSetCapType *string `json:"UserSetCapType,omitnil,omitempty" name:"UserSetCapType"`

	// <p>Interception mode</p><p>Enumeration values:</p><ul><li>notify: perception mode</li><li>block: interception mode</li></ul>
	DefendMode *string `json:"DefendMode,omitnil,omitempty" name:"DefendMode"`

	// <p>Whether to enable captcha encryption. 0: Off. 1: On</p>
	CheckAppidSwitch *int64 `json:"CheckAppidSwitch,omitnil,omitempty" name:"CheckAppidSwitch"`

	// <p>Whether to enable Non-repeating IV</p><p>Enumeration values:</p><ul><li>0: Off</li><li>1: On</li></ul><p>Input 1 is allowed only when CheckAppidSwitch is 1</p>
	CheckIvSwitch *int64 `json:"CheckIvSwitch,omitnil,omitempty" name:"CheckIvSwitch"`

	// <p>Verification mechanism: '0' One-Click Verification, '1' Always verify, '2' Invisible verification</p><p>Enumeration values:</p><ul><li>0: One-Click Verification</li><li>1: Always verify</li><li>2: Invisible verification. DisableInvisibleSwitch input 2, UserSetCapType must be 1</li></ul>
	DisableInvisibleSwitch *string `json:"DisableInvisibleSwitch,omitnil,omitempty" name:"DisableInvisibleSwitch"`

	// <p>Web domain name</p><p>Only valid when ChannelInfo is web</p>
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// <p>app BundleId</p><p>Valid only when ChannelInfo is ios</p>
	VerifyBundleId *string `json:"VerifyBundleId,omitnil,omitempty" name:"VerifyBundleId"`

	// <p>app package</p><p>Valid only when ChannelInfo is android</p>
	VerifyPackage *string `json:"VerifyPackage,omitnil,omitempty" name:"VerifyPackage"`

	// <p>Resource tag, key&amp;value format</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Checkbox display method. '0': minimalist mode, '1': full mode, '2': not set</p>
	CheckBoxStyle *string `json:"CheckBoxStyle,omitnil,omitempty" name:"CheckBoxStyle"`
}

type ModifyCaptchaInfoInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>Captcha appid</p>
	CaptchaAppId *string `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`

	// <p>Captcha name</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Verification level</p><p>Enumeration values:</p><ul><li>1: Experience-oriented</li><li>2: Balanced</li><li>3: Security-focused</li></ul><p>Default value: 1</p>
	VerifyRank *string `json:"VerifyRank,omitnil,omitempty" name:"VerifyRank"`

	// <p>Verification method</p><p>Enumeration values:</p><ul><li>1: Invisible verification. UserSetCapType input 1, DisableInvisibleSwitch must</li><li>2: Slide verification</li><li>8: Graphical verification</li><li>9: Voice verification</li></ul>
	UserSetCapType *string `json:"UserSetCapType,omitnil,omitempty" name:"UserSetCapType"`

	// <p>Interception mode</p><p>Enumeration values:</p><ul><li>notify: perception mode</li><li>block: interception mode</li></ul>
	DefendMode *string `json:"DefendMode,omitnil,omitempty" name:"DefendMode"`

	// <p>Whether to enable captcha encryption. 0: Off. 1: On</p>
	CheckAppidSwitch *int64 `json:"CheckAppidSwitch,omitnil,omitempty" name:"CheckAppidSwitch"`

	// <p>Whether to enable Non-repeating IV</p><p>Enumeration values:</p><ul><li>0: Off</li><li>1: On</li></ul><p>Input 1 is allowed only when CheckAppidSwitch is 1</p>
	CheckIvSwitch *int64 `json:"CheckIvSwitch,omitnil,omitempty" name:"CheckIvSwitch"`

	// <p>Verification mechanism: '0' One-Click Verification, '1' Always verify, '2' Invisible verification</p><p>Enumeration values:</p><ul><li>0: One-Click Verification</li><li>1: Always verify</li><li>2: Invisible verification. DisableInvisibleSwitch input 2, UserSetCapType must be 1</li></ul>
	DisableInvisibleSwitch *string `json:"DisableInvisibleSwitch,omitnil,omitempty" name:"DisableInvisibleSwitch"`

	// <p>Web domain name</p><p>Only valid when ChannelInfo is web</p>
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// <p>app BundleId</p><p>Valid only when ChannelInfo is ios</p>
	VerifyBundleId *string `json:"VerifyBundleId,omitnil,omitempty" name:"VerifyBundleId"`

	// <p>app package</p><p>Valid only when ChannelInfo is android</p>
	VerifyPackage *string `json:"VerifyPackage,omitnil,omitempty" name:"VerifyPackage"`

	// <p>Resource tag, key&amp;value format</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Checkbox display method. '0': minimalist mode, '1': full mode, '2': not set</p>
	CheckBoxStyle *string `json:"CheckBoxStyle,omitnil,omitempty" name:"CheckBoxStyle"`
}

func (r *ModifyCaptchaInfoInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCaptchaInfoInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "AppName")
	delete(f, "VerifyRank")
	delete(f, "UserSetCapType")
	delete(f, "DefendMode")
	delete(f, "CheckAppidSwitch")
	delete(f, "CheckIvSwitch")
	delete(f, "DisableInvisibleSwitch")
	delete(f, "VerifyDomain")
	delete(f, "VerifyBundleId")
	delete(f, "VerifyPackage")
	delete(f, "Tags")
	delete(f, "CheckBoxStyle")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCaptchaInfoInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCaptchaInfoInternationalResponseParams struct {
	// <p>Result data</p>
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Captcha status code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha information</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCaptchaInfoInternationalResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCaptchaInfoInternationalResponseParams `json:"Response"`
}

func (r *ModifyCaptchaInfoInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCaptchaInfoInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpWhiteListInternationalRequestParams struct {
	// <p>ip allowlist name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Record number</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>IP whitelist status</p><p>Enumeration values:</p><ul><li>0: enable</li><li>1: disable</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Remark information.</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifyIpWhiteListInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>ip allowlist name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Record number</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Captcha appid</p>
	CaptchaAppid *int64 `json:"CaptchaAppid,omitnil,omitempty" name:"CaptchaAppid"`

	// <p>IP whitelist status</p><p>Enumeration values:</p><ul><li>0: enable</li><li>1: disable</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Remark information.</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *ModifyIpWhiteListInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpWhiteListInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Id")
	delete(f, "CaptchaAppid")
	delete(f, "Status")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIpWhiteListInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpWhiteListInternationalResponseParams struct {
	// <p>Result data</p>
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Captcha status code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha information</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIpWhiteListInternationalResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIpWhiteListInternationalResponseParams `json:"Response"`
}

func (r *ModifyIpWhiteListInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpWhiteListInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderByInternational struct {
	// Sort by creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type RemoveCaptchaInfoInternationalRequestParams struct {
	// <p>Captcha AppId</p>
	CaptchaAppId *string `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`
}

type RemoveCaptchaInfoInternationalRequest struct {
	*tchttp.BaseRequest
	
	// <p>Captcha AppId</p>
	CaptchaAppId *string `json:"CaptchaAppId,omitnil,omitempty" name:"CaptchaAppId"`
}

func (r *RemoveCaptchaInfoInternationalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveCaptchaInfoInternationalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveCaptchaInfoInternationalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveCaptchaInfoInternationalResponseParams struct {
	// <p>Result data</p>
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Captcha status code</p>
	CaptchaCode *int64 `json:"CaptchaCode,omitnil,omitempty" name:"CaptchaCode"`

	// <p>Captcha information</p>
	CaptchaMsg *string `json:"CaptchaMsg,omitnil,omitempty" name:"CaptchaMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveCaptchaInfoInternationalResponse struct {
	*tchttp.BaseResponse
	Response *RemoveCaptchaInfoInternationalResponseParams `json:"Response"`
}

func (r *RemoveCaptchaInfoInternationalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveCaptchaInfoInternationalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}