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

package v20201229

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type DetailResults struct {
	// Result of the moderation. <br>`Normal`: normal content; `Porn`: pornographic content; `Abuse`: abusive content; **Ad**: advertising content; `Custom`: custom violating content
	Label *string `json:"Label,omitempty" name:"Label"`

	// Recommended follow-up action. <br>`Block`: block it automatically; `Review`: review the content again in human; **Pass**: pass
	// Note: This field may return `null`, indicating that no valid value can be found.
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// This field returns the matched keywords. This parameter can include multiple returned values, which means multiple keywords are matched. If no keyword is returned but there is a `Score`, it means that the result of `Label` is determined by a semantic model.
	// Note: This field may return `null`, indicating that no valid value can be found.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// This field indicates the convincing level of the `Label`, ranging from `0` (lowest) to `100` (highest). 
	// Note: This field may return `null`, indicating that no valid value can be found.
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// It indicates the library type corresponding with the keyword. Valid values: `1` (blocklist/allowlist library) and `2` (custom keyword library). If no custom keyword library is configured, the default value is 1.
	// Note: This field may return `null`, indicating that no valid value can be found.
	LibType *int64 `json:"LibType,omitempty" name:"LibType"`

	// This field is **only valid when `Label` is `Custom`. It returns the custom library ID to facilitate the library management and configuration.
	// Note: This field may return `null`, indicating that no valid value can be found.
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// This field is **only valid when `Label` is `Custom` (custom keyword)`. It returns the custom library name to facilitate the library management and configuration.
	// Note: This field may return `null`, indicating that no valid value can be found.
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// The field returns the second-level labels under the current label.
	// Note: This field may return `null`, indicating that no valid value can be found.
	SubLabel *string `json:"SubLabel,omitempty" name:"SubLabel"`
}

type Device struct {
	// This field indicates the IP address of the device used by the service subscriber. <br>
	// Note: Currently, only IPv4 addresses can be recorded.
	IP *string `json:"IP,omitempty" name:"IP"`

	// This field indicates the MAC address used by the service subscriber to facilitate device identification and management. Its format and value are consistent with those of the standard MAC address.
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// * In beta test. Available soon.*
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`

	// * In beta test. Available soon.*
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// This field represents the **IMEI** (International Mobile Equipment Identity) number of the device used by the service subscriber. IMEI can be used to identify each independent mobile communication device, such as a mobile phone, which is convenient for device identification and management. <br>Note: IMEI is formatted with **15 to 17 numbers only**.
	IMEI *string `json:"IMEI,omitempty" name:"IMEI"`

	// **Dedicated for iOS device**. This field indicates the **IDFA** (Identifier for Advertising) corresponding to the service subscriber. IDFA, a string of hexadecimal 32 characters including numbers and letters, is provided by Apple Inc. to identify the user.<br>
	// Note: Since the iOS14 update in 2021, Apple Inc. has allowed users to manually activate or deactivate IDFA, so the validity of the string identifier may be reduced.
	IDFA *string `json:"IDFA,omitempty" name:"IDFA"`

	// **Dedicated for iOS device**. This field indicates the **IDFV** (Identifier for Vendor) corresponding to the service subscriber. IDFV, a string of hexadecimal 32 characters including numbers and letters, is provided by Apple Inc. to identify the vendor. IDFV can also be used as a unique device identifier.
	IDFV *string `json:"IDFV,omitempty" name:"IDFV"`
}

type RiskDetails struct {
	// This field returns the risk categories after account information detection. Valid values: **RiskAccount** (the account is at risk), **RiskIP** (the IP address is at risk), and **RiskIMEI** (the mobile device identifier is at risk).
	Label *string `json:"Label,omitempty" name:"Label"`

	// This field returns the risk levels after account information detection. Valid values: **1** (suspected to be at risk) and **2** (malicious risk).
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

// Predefined struct for user
type TextModerationRequestParams struct {
	// This field indicates the text content of the object to be moderated. The text needs to be encoded in utf-8 format and encrypted with Base64. It can contain up to 10,000 characters, calculated by unicode encoding.
	Content *string `json:"Content,omitempty" name:"Content"`

	// This field indicates the specific policy number, which is used for the API call and can be configured in the CMS console. If it's not entered (left empty), the default moderation policy is adopted. If it's entered, the moderation policies are specified for business scenarios. <br>Note: Biztype contains 3 to 32 characters, including numbers, letters and underscores only. Different Biztypes are associated with different business scenarios and moderation policies. Ensure that you use the Biztype corresponding to the policy you want to apply.
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// This field indicates the data ID you assigned to the object to be moderated, which is convenient for you to identify and manage the file. <br>Value: this field can contain **up to 64 characters**, including uppercase and lowercase letters, numbers and four special symbols (_, -, @, #)
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// This field indicates the user information related with the object to be moderated, which can be used to identify violating users at risk.
	User *User `json:"User,omitempty" name:"User"`

	// This field indicates the device information related with the object to be moderated, which can be used to identify violating devices at risk.
	Device *Device `json:"Device,omitempty" name:"Device"`
}

type TextModerationRequest struct {
	*tchttp.BaseRequest
	
	// This field indicates the text content of the object to be moderated. The text needs to be encoded in utf-8 format and encrypted with Base64. It can contain up to 10,000 characters, calculated by unicode encoding.
	Content *string `json:"Content,omitempty" name:"Content"`

	// This field indicates the specific policy number, which is used for the API call and can be configured in the CMS console. If it's not entered (left empty), the default moderation policy is adopted. If it's entered, the moderation policies are specified for business scenarios. <br>Note: Biztype contains 3 to 32 characters, including numbers, letters and underscores only. Different Biztypes are associated with different business scenarios and moderation policies. Ensure that you use the Biztype corresponding to the policy you want to apply.
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// This field indicates the data ID you assigned to the object to be moderated, which is convenient for you to identify and manage the file. <br>Value: this field can contain **up to 64 characters**, including uppercase and lowercase letters, numbers and four special symbols (_, -, @, #)
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// This field indicates the user information related with the object to be moderated, which can be used to identify violating users at risk.
	User *User `json:"User,omitempty" name:"User"`

	// This field indicates the device information related with the object to be moderated, which can be used to identify violating devices at risk.
	Device *Device `json:"Device,omitempty" name:"Device"`
}

func (r *TextModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "User")
	delete(f, "Device")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextModerationResponseParams struct {
	// This field returns the BizType of the request parameters
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// This field returns the **negative label with the highest priority** in the moderation results (DetailResults), which indicates the moderation result recommended by the model. It is recommended that you handle different violations with the suggested values according to your business needs. <br>Returned values: **Normal**: normal content; **Porn**: pornographic content; **Abuse**: abusive content; **Ad**: advertising content; **Custom**: custom violating content, and others such as objectionable, insecure or inappropriate content.
	Label *string `json:"Label,omitempty" name:"Label"`

	// This field returns the follow-up moderation suggestions. The returned value indicates the recommended operation after obtaining the moderation result. It is recommended that you handle different violations with the suggested values according to your business needs. <br>Returned values: **Block**: block; **Review**: human moderation; **Pass**: pass
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// This field returns the keywords matched with the libraries in the moderated text under the current label to mark the specific violations (for example, *Friend me*). This parameter may have multiple returned values, indicating multiple keywords are matched. If the returned value is empty and the `Score` is not empty, it means that the negative label corresponding to the moderation result is a value returned from the semantic model judgment
	// Note: This field may return `null`, indicating that no valid value can be found.
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// This field returns the confidence level under the current label. Value range: 0 (**the lowest confidence level**) - 100 (**the highest confidence level**). The higher the value, the more likely the text is to belong to the category indicated by the current label. For example, *pornographic 99* indicates that the text is very likely to be pornographic, and *pornographic 0* indicates that the text is not pornographic
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// This field returns the moderation results based on the text libraries. For details, see `DetailResults` in the data structure
	// Note: This field may return `null`, indicating that no valid value can be found.
	DetailResults []*DetailResults `json:"DetailResults,omitempty" name:"DetailResults"`

	// This field returns the detection results of violating accounts at risk, mainly including violation categories and risk levels. For details, see `RiskDetails` in the data structure
	// Note: This field may return `null`, indicating that no valid value can be found.
	RiskDetails []*RiskDetails `json:"RiskDetails,omitempty" name:"RiskDetails"`

	// This field returns the extra information configured according to your needs. If it's not configured, the returned value is empty by default. <br>Note: the returned information varies based on different customers or Biztypes. If you need to configure this field, please submit a ticket or contact after-sales manager
	// Note: This field may return `null`, indicating that no valid value can be found.
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// This field returns the `DataId` in the request parameter corresponding to the moderated object
	// Note: This field may return `null`, indicating that no valid value can be found.
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// The field returns the second-level labels under the current label.
	// Note: This field may return `null`, indicating that no valid value can be found.
	SubLabel *string `json:"SubLabel,omitempty" name:"SubLabel"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextModerationResponse struct {
	*tchttp.BaseResponse
	Response *TextModerationResponseParams `json:"Response"`
}

func (r *TextModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// This field indicates the service subscriber ID. This ID can be used to optimize the moderation result judgment based on the account's violation records, which is helpful for auxiliary judgment when there is a risk of suspected violations.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// This field indicates the account nickname information of the service subscriber.
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// This field indicates the account type corresponding to the service subscriber ID.<br>
	// Use this field and the account ID (UserId) together to determine a unique account.
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// This field indicates the gender information of the service subscriber's account.<br>
	// Values: **0** (default value, indicating the gender is unknown), **1** (male) and **2** (female).
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// This field indicates the age information of the service subscriber's account.<br>
	// Values: Integers between **0** (default value, indicating that the age is unknown) and the number of (**custom maximum age**).
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// This field indicates the level information of the service subscriber's account.<br>
	// Values: **0** (default value, indicating that the level is unknown), **1** (lower level), **2** (medium level) and **3** (higher level). Currently, **custom levels are not supported**.
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// This field indicates the mobile phone number information of the service subscriber's account. The mobile phone numbers in various regions of the world can be recorded.<br>
	// Note: Please keep the format of mobile phone number uniform. For example, uniformly use the area code format (086/+86), etc.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// This field indicates the URL of the service subscriber's profile photos formatted with .png, .jpg, .jpeg, .bmp, .gif and .webp.
	// Note: Up to 5 MB is supported, and the minimum resolution is 256 x 256. When it takes more than 3 seconds to download, the "download timeout" is returned.
	HeadUrl *string `json:"HeadUrl,omitempty" name:"HeadUrl"`

	// This field indicates the profile information of service subscribers. It can contain up to 5,000 characters, including Chinese characters, letters and special symbols.
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}