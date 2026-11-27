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

package v20230110

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccountFactoryItem struct {
	// Unique ID of the Account Factory baseline item, which can contain only English letters, numbers, and symbols (@,._[]-:()+=), and must be 2 to 128 characters in length.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Baseline item name, which is unique, supports only the combination of English letters, digits, and symbols (@&_[]-), and is 1 to 25 characters in length.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Baseline item English name, which must be unique, can contain only English letters, numbers, spaces, and symbols (@&_[]-), and must be 1 to 64 characters in length.
	NameEn *string `json:"NameEn,omitnil,omitempty" name:"NameEn"`

	// Baseline item weight. A smaller value indicates a higher priority. Value range: greater than or equal to 0.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Whether the baseline item is required. Valid values: 1 (yes) and 0 (no).
	Required *int64 `json:"Required,omitnil,omitempty" name:"Required"`

	// Baseline item dependencies. The value range of N depends on the number of other baseline items that the current baseline item depends on.
	DependsOn []*DependsOnItem `json:"DependsOn,omitnil,omitempty" name:"DependsOn"`

	// Baseline description, which is 2 to 256 English characters in length. The default value is empty.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Baseline item English description, which is 2 to 1024 English characters in length. The default value is empty.
	DescriptionEn *string `json:"DescriptionEn,omitnil,omitempty" name:"DescriptionEn"`

	// Baseline category, which is 2 to 32 English characters in length and cannot be empty.
	Classify *string `json:"Classify,omitnil,omitempty" name:"Classify"`

	// Baseline English category, which is 2 to 64 English characters in length and cannot be empty.
	ClassifyEn *string `json:"ClassifyEn,omitnil,omitempty" name:"ClassifyEn"`
}

type BaselineConfigItem struct {
	// Specifies the unique identifier for account factory baseline item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Account factory baseline item configuration. different baseline item configuration parameters.
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`
}

type BaselineInfoItem struct {
	// Specifies the unique identifier for account factory baseline item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Account factory baseline item configuration. different baseline item configuration parameters.
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// Specifies the number of accounts for baseline applications.
	ApplyCount *int64 `json:"ApplyCount,omitnil,omitempty" name:"ApplyCount"`
}

type BaselineStepTaskInfo struct {
	// Specifies the unique Id of the task, which can only contain english letters and digits, and is a 16-character random string.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Specifies the unique identifier for the baseline feature item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Specifies the member account uin of the applied baseline item.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Baseline item application status. Running means the baseline item is in application. Success means the baseline item application is successful. Failed means the baseline item application failure. Pending means the baseline item is Pending application. Skipped means the baseline item is Skipped.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	ErrCode *string `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message
	ErrMessage *string `json:"ErrMessage,omitnil,omitempty" name:"ErrMessage"`

	// Baseline item deployment output.
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Creation time, represented in ISO8601 standard format as yyyy-MM-dd hh:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the last update time in ISO8601 standard representation with format yyyy-MM-dd hh:MM:ss.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type BatchApplyAccountBaselinesRequestParams struct {
	// Member account UIN, which is also the UIN of the account to which the baseline is applied.
	MemberUinList []*int64 `json:"MemberUinList,omitnil,omitempty" name:"MemberUinList"`

	// List of baseline item configuration information.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

type BatchApplyAccountBaselinesRequest struct {
	*tchttp.BaseRequest
	
	// Member account UIN, which is also the UIN of the account to which the baseline is applied.
	MemberUinList []*int64 `json:"MemberUinList,omitnil,omitempty" name:"MemberUinList"`

	// List of baseline item configuration information.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

func (r *BatchApplyAccountBaselinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchApplyAccountBaselinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUinList")
	delete(f, "BaselineConfigItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchApplyAccountBaselinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchApplyAccountBaselinesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchApplyAccountBaselinesResponse struct {
	*tchttp.BaseResponse
	Response *BatchApplyAccountBaselinesResponseParams `json:"Response"`
}

func (r *BatchApplyAccountBaselinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchApplyAccountBaselinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependsOnItem struct {
	// Dependency item type. Valid values: LandingZoneSetUp (Landing Zone dependency) and AccountFactorySetUp (Account Factory dependency).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Feature item unique ID, which can contain only English letters, numbers, and symbols (@,._[]-:()+=), and must be 2 to 128 characters in length.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`
}

// Predefined struct for user
type GetAccountFactoryBaselineRequestParams struct {

}

type GetAccountFactoryBaselineRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAccountFactoryBaselineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountFactoryBaselineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAccountFactoryBaselineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAccountFactoryBaselineResponseParams struct {
	// Specifies the uin of the main account to which the resource belongs.
	OwnerUin *int64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Specifies the baseline item name, which must be unique and can only contain a combination of english letters, digits, chinese characters, and symbols @, &_[]-, with a length of 1-25 chinese or english characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of baseline item configurations.
	BaselineConfigItems []*BaselineInfoItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`

	// Creation time, represented in ISO8601 standard format as yyyy-MM-dd hh:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the last update time in ISO8601 standard representation with format yyyy-MM-dd hh:MM:ss.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAccountFactoryBaselineResponse struct {
	*tchttp.BaseResponse
	Response *GetAccountFactoryBaselineResponseParams `json:"Response"`
}

func (r *GetAccountFactoryBaselineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountFactoryBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccountFactoryBaselineItemsRequestParams struct {
	// Maximum number of returned records. Value range: 0-200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Value range: greater than or equal to 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListAccountFactoryBaselineItemsRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of returned records. Value range: 0-200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Value range: greater than or equal to 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListAccountFactoryBaselineItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountFactoryBaselineItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAccountFactoryBaselineItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccountFactoryBaselineItemsResponseParams struct {
	// <p>Account Factory baseline list.<br><strong>1. TCC-AF_ACCOUNT_CONTACT Account contact</strong><br>This baseline item is used to create message recipients (contacts) under a member account.<br>Field    Type    Required    Default value    Valid values and description<br>AccountContacts    Object array    Yes    None    Contact list, up to 50 entries, cannot be empty<br>AccountContacts[].Name    String    Yes    None    Contact name, must be an existing name in the account; cannot be duplicated within the same list; no fixed format restrictions<br>AccountContacts[].CountryCode    String    Yes    None    Phone country code, such as 86; no fixed enumeration<br>AccountContacts[].PhoneNum    String    Yes    None    Phone number, must be verified before this baseline item can be applied; format is determined by the verification process, with no local format restrictions<br>AccountContacts[].Email    String    Yes    None    Email address, must be verified before this baseline item can be applied; format is determined by the verification process, with no local format restrictions<br>AccountContacts[].Remark    String    No    None    Remarks, no length limit specified<br><pre><code>{  &quot;AccountContacts&quot;: [    {      &quot;Name&quot;: &quot;ops-admin&quot;,      &quot;CountryCode&quot;: &quot;86&quot;,      &quot;PhoneNum&quot;: &quot;13800***000&quot;,      &quot;Email&quot;: &quot;ops@ex***le.com&quot;,      &quot;Remark&quot;: &quot;Ops contact&quot;    }  ]}</code></pre><strong>2. TCC-AF_ACCOUNT_NOTIFICATION Message recipient</strong><br>China site:<br>Field    Type    Required    Default value    Valid values and description<br>Subscriptions    Object array    Yes    None    Subscription list, recommended not to exceed 50 entries; if exceeded, process in batches, will not be rejected due to the number of entries<br>Subscriptions[].Name    String    Yes    None    Single sub-user name; enter an existing sub-user name, otherwise the subscription does not take effect<br>Subscriptions[].ProductId    Integer    Yes    None    Product ID, no fixed enumeration, query through DescribeProductIds<br>Subscriptions[].MsgType    Integer    Yes    None    Message type ID, no fixed enumeration, query through DescribeFirstSiteMsgType and DescribeSecondMsgType<br>Subscriptions[].Channel    Integer    Yes    None    Message delivery channel, bitmask, see the table below for values<br>Subscriptions[].Status    String    Yes    None    Only Open or Close; only Open creates a subscription, Close or other values do not take effect</p><p>Channel bitmask:<br>Channel    Value<br>Inbox message    1<br>Email    2<br>SMS    4<br>WeChat    8 (not available on the international site)<br>Voice    16<br>WeCom    32 (not available on the international site)</p><p>Channel combination    Value<br>Inbox message + Email    3<br>Inbox message + SMS    5<br>Email + SMS    6<br>Inbox message + Email + SMS    7</p><p>International site fields:<br>Field    Type    Required    Default value    Valid values and description<br>Subscriptions    Object array    Yes    None    Subscription list, recommended not to exceed 50 entries; if exceeded, process in batches, will not be rejected due to the number of entries<br>Subscriptions[].Name    String array    Yes    None    Sub-user name list; enter existing sub-user names, names will be deduplicated<br>Subscriptions[].MsgType    Integer    Yes    None    Message type ID, no fixed enumeration, query through DescribeFirstSiteMsgType and DescribeSecondMsgType<br>Subscriptions[].Channel    Integer    Yes    None    Message delivery channel, bitmask, values same as the table above</p><p>China site example:<br><pre><code>{  &quot;Subscriptions&quot;: [    {      &quot;Name&quot;: &quot;ops-admin&quot;,      &quot;ProductId&quot;: 1001,      &quot;MsgType&quot;: 2001,      &quot;Channel&quot;: 3,      &quot;Status&quot;: &quot;Open&quot;    }  ]}</code></pre>International site example:<br><pre><code>{  &quot;Subscriptions&quot;: [    {      &quot;Name&quot;: [&quot;ops-admin&quot;, &quot;audit&quot;],      &quot;MsgType&quot;: 2001,      &quot;Channel&quot;: 6    }  ]}</code></pre><strong>3. TCC-AF_CAM_USER_PASSWORD_POLICY CAM password policy</strong><br>Field    Type    Required    Default value    Valid values and description<br>MinimumLength    Integer    Yes    None    Minimum password length, no fixed enumeration or upper/lower limit specified; whether the value is valid is determined by the CAM password policy<br>MustContain    String    Yes    None    Characters that the password must contain, no fixed enumeration or valid character set specified; whether the value is valid is determined by the CAM password policy<br>ForcePasswordChange    Integer    Yes    None    Forced password change period, no fixed enumeration or unit specified; whether the value is valid is determined by the CAM password policy<br>ReusePasswordLimit    Integer    Yes    None    Password reuse limit, no fixed enumeration or upper/lower limit specified; whether the value is valid is determined by the CAM password policy<br>RetryPasswordLimit    Integer    Yes    None    Password retry limit, no fixed enumeration or upper/lower limit specified; whether the value is valid is determined by the CAM password policy<br>OnlyAdminCanResetPassword    Integer    Yes    None    Whether only administrators can reset passwords, no fixed enumeration specified; whether the value is valid is determined by the CAM password policy<br><pre><code>{  &quot;MinimumLength&quot;: 19,  &quot;MustContain&quot;: &quot;A&quot;,  &quot;ForcePasswordChange&quot;: 0,  &quot;ReusePasswordLimit&quot;: 2,  &quot;RetryPasswordLimit&quot;: 1,  &quot;OnlyAdminCanResetPassword&quot;: 0}</code></pre><strong>4. TCC-AF_CAM_USER_SECURITY CAM user security settings</strong><br>Field    Type    Required    Default value    Valid values and description<br>DefaultMFASettings    Object    Yes    None    Login and operation protection settings<br>DefaultMFASettings.LoginStrategy    Integer    Yes    None    1: enforce verification for all users; 2: depend on individual user settings<br>DefaultMFASettings.ActionStrategy    Integer    Yes    None    1: enforce verification for all users; 2: depend on individual user settings; this baseline item must be set to 2<br>DefaultMFASettings.LoginDetail    String array    Yes    None    Valid options: Stoken, Phone, Wechat, U2FToken, Mail; must include Stoken<br>LoginStatusSettings    Object    Yes    None    Session duration settings<br>LoginStatusSettings.SessionTimeout    Integer    No    None    Session timeout in seconds; no fixed upper/lower limit specified; whether the value is valid is determined by CAM<br>LoginStatusSettings.IdleSessionTimeout    Integer    No    None    Idle session timeout in seconds; no fixed upper/lower limit specified; whether the value is valid is determined by CAM<br><pre><code>{  &quot;DefaultMFASettings&quot;: {    &quot;LoginStrategy&quot;: 1,    &quot;ActionStrategy&quot;: 2,    &quot;LoginDetail&quot;: [&quot;Stoken&quot;, &quot;Phone&quot;]  },  &quot;LoginStatusSettings&quot;: {    &quot;SessionTimeout&quot;: 7200,    &quot;IdleSessionTimeout&quot;: 900  }}</code></pre><strong>5. TCC-AF_VPC_SUBNET VPC and subnet</strong><br>Field    Type    Required    Default value    Valid values and description<br>VpcName    String    Yes    None    VPC name, up to 60 bytes<br>CidrBlock    String    Yes    None    VPC CIDR block, only addresses within the three private network segments 10.0.0.0/12, 172.16.0.0/12, and 192.168.0.0/16 are supported<br>RegionName    String    Yes    None    Region, fill in according to "How to fill in regions, AZs, and dynamic values"; the region of this baseline item is specified by this field<br>Subnets    Object array    No    None    Subnet list; if not specified or empty, only the VPC is created, and no subnets are created<br>Subnets[].CidrBlock    String    Yes    None    Subnet CIDR block, must be within the range of CidrBlock, and subnet CIDR blocks within the same VPC cannot overlap<br>Subnets[].SubnetName    String    Yes    None    Subnet name, up to 60 bytes<br>Subnets[].Zone    String    Yes    None    AZ, in the format of ap-guangzhou-2; no fixed enumeration, query through DescribeZones<br><pre><code>{  &quot;VpcName&quot;: &quot;af-vpc&quot;,  &quot;CidrBlock&quot;: &quot;10.0.0.0/16&quot;,  &quot;RegionName&quot;: &quot;ap-guangzhou&quot;,  &quot;Subnets&quot;: [    {      &quot;CidrBlock&quot;: &quot;10.0.1.0/24&quot;,      &quot;SubnetName&quot;: &quot;af-subnet-1&quot;,      &quot;Zone&quot;: &quot;ap-guangzhou-2&quot;    }  ]}</code></pre><strong>6. TCC-AF_SECURITY_GROUP Security group</strong><br>Field    Type    Required    Default value    Valid values and description<br>GroupName    String    Yes    None    Security group name, up to 60 characters<br>GroupDescription    String    No    None    Security group description, up to 100 characters<br>Region    String    Yes    None    Region, fill in according to "How to fill in regions, AZs, and dynamic values"<br>SecurityGroupPolicySet    Object    Yes    None    Security group rule set<br>SecurityGroupPolicySet.Ingress    Object array    No    None    Inbound rule list<br>SecurityGroupPolicySet.Egress    Object array    No    None    Outbound rule list<br>Ingress[].Protocol, Egress[].Protocol    String    Yes    None    TCP, UDP, ICMP, ICMPv6, ALL<br>Ingress[].Port, Egress[].Port    String    Yes    None    all, a single port, or a port range (such as 80 or 80-90); must be set to all when Protocol is ALL<br>Ingress[].Action, Egress[].Action    String    Yes    None    ACCEPT: allow traffic; DROP: drop traffic<br>Ingress[].CidrBlock, Egress[].CidrBlock    String    No    None    IPv4 CIDR block or IP, such as 10.0.0.0/16 or 0.0.0.0/0; mutually exclusive with Ipv6CidrBlock<br>Ingress[].Ipv6CidrBlock, Egress[].Ipv6CidrBlock    String    No    None    Not supported yet, do not fill in<br>Ingress[].SecurityGroupId, Egress[].SecurityGroupId    String    No    None    Referenced security group ID, in the format of sg-ohuuioma; no fixed enumeration, query through the resource list<br>Ingress[].AddressTemplate.AddressId    String    No    None    IP address template ID, in the format of ipm-2uw6ujo6; when using an address template, this field and AddressGroupId must be filled in together<br>Ingress[].AddressTemplate.AddressGroupId    String    No    None    IP address group ID, in the format of ipmg-2uw6ujo6; when using an address template, this field and AddressId must be filled in together<br>Ingress[].PolicyDescription, Egress[].PolicyDescription    String    No    None    Rule description, no length limit specified; whether the value is valid is determined by the cloud product<br><pre><code>{  &quot;GroupName&quot;: &quot;af-sg&quot;,  &quot;GroupDescription&quot;: &quot;account factory baseline&quot;,  &quot;Region&quot;: &quot;ap-guangzhou&quot;,  &quot;SecurityGroupPolicySet&quot;: {    &quot;Ingress&quot;: [      {        &quot;Protocol&quot;: &quot;TCP&quot;,        &quot;Port&quot;: &quot;22&quot;,        &quot;CidrBlock&quot;: &quot;10.0.0.0/16&quot;,        &quot;Action&quot;: &quot;ACCEPT&quot;,        &quot;PolicyDescription&quot;: &quot;ssh&quot;      },      {        &quot;Protocol&quot;: &quot;TCP&quot;,        &quot;Port&quot;: &quot;443&quot;,        &quot;AddressTemplate&quot;: {          &quot;AddressId&quot;: &quot;ipm-2uw6ujo6&quot;,          &quot;AddressGroupId&quot;: &quot;ipmg-2uw6ujo6&quot;        },        &quot;Action&quot;: &quot;ACCEPT&quot;      }    ],    &quot;Egress&quot;: [      {        &quot;Protocol&quot;: &quot;ALL&quot;,        &quot;Port&quot;: &quot;all&quot;,        &quot;CidrBlock&quot;: &quot;0.0.0.0/0&quot;,        &quot;Action&quot;: &quot;ACCEPT&quot;,        &quot;PolicyDescription&quot;: &quot;allow all&quot;      }    ]  }}</code></pre><strong>7. TCC-AF_PRESET_TAG Preset tag</strong><br>Field    Type    Required    Default value    Valid values and description<br>TagValuePairs    Object array    Yes    None    Tag key-value pair list<br>TagValuePairs[].Key    String    Yes    None    Tag key, no fixed length or character restrictions specified; whether the value is valid is determined by the tag service<br>TagValuePairs[].Values    String array    Yes    None    List of tag values corresponding to the key; no fixed length or character restrictions specified; whether the value is valid is determined by the tag service<br>Constraint: the total number of entries in the array cannot exceed 30<br><pre><code>{  &quot;TagValuePairs&quot;: [    { &quot;Key&quot;: &quot;env&quot;, &quot;Values&quot;: [&quot;prod&quot;, &quot;dev&quot;] },    { &quot;Key&quot;: &quot;owner&quot;, &quot;Values&quot;: [&quot;ops&quot;] }  ]}</code></pre><strong>8. TCC-AF_SHARE_IMAGE Shared image</strong><br>Field    Type    Required    Default value    Valid values and description<br>Images    Object array    Yes    None    List of images to be shared, no upper limit specified<br>Images[].Region    String    Yes    None    Region where the image resides, fill in according to "How to fill in regions, AZs, and dynamic values"<br>Images[].ImageId    String    Yes    None    Image ID, in the format of img-gvbnzy6f; no fixed enumeration, query through DescribeImages or the console, the image must be in the available state<br><pre><code>{  &quot;Images&quot;: [    {      &quot;Region&quot;: &quot;ap-guangzhou&quot;,      &quot;ImageId&quot;: &quot;img-gvbnzy6f&quot;    }  ]}</code></pre></p>
	BaselineItems []*AccountFactoryItem `json:"BaselineItems,omitnil,omitempty" name:"BaselineItems"`

	// <p>Total number.</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAccountFactoryBaselineItemsResponse struct {
	*tchttp.BaseResponse
	Response *ListAccountFactoryBaselineItemsResponseParams `json:"Response"`
}

func (r *ListAccountFactoryBaselineItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountFactoryBaselineItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployStepTasksRequestParams struct {
	// Specifies the unique identifier for the feature item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Maximum number of returned records. value ranges from 0 to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. valid values are equal to or greater than 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListDeployStepTasksRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the unique identifier for the feature item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Maximum number of returned records. value ranges from 0 to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. valid values are equal to or greater than 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListDeployStepTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployStepTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Identifier")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDeployStepTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployStepTasksResponseParams struct {
	// Account factory baseline function application information list.
	BaselineDeployStepTaskList []*BaselineStepTaskInfo `json:"BaselineDeployStepTaskList,omitnil,omitempty" name:"BaselineDeployStepTaskList"`

	// Total quantity.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDeployStepTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListDeployStepTasksResponseParams `json:"Response"`
}

func (r *ListDeployStepTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployStepTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccountFactoryBaselineRequestParams struct {
	// Specifies the baseline name, which must be unique and can only contain a combination of english letters, digits, chinese characters, and symbols @, &_[]-, with a length of 1-25 chinese or english characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Baseline configuration. overwrite update. can be accessed through controlcenter:GetAccountFactoryBaseline to query existing baseline configuration. can be accessed through controlcenter:ListAccountFactoryBaselineItems to query supported baseline list.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

type UpdateAccountFactoryBaselineRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the baseline name, which must be unique and can only contain a combination of english letters, digits, chinese characters, and symbols @, &_[]-, with a length of 1-25 chinese or english characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Baseline configuration. overwrite update. can be accessed through controlcenter:GetAccountFactoryBaseline to query existing baseline configuration. can be accessed through controlcenter:ListAccountFactoryBaselineItems to query supported baseline list.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

func (r *UpdateAccountFactoryBaselineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccountFactoryBaselineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "BaselineConfigItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccountFactoryBaselineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccountFactoryBaselineResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccountFactoryBaselineResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccountFactoryBaselineResponseParams `json:"Response"`
}

func (r *UpdateAccountFactoryBaselineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccountFactoryBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}