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

package v20221121

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AIAgentAsset struct {
	// <p>ID identifier</p>
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>agent name</p>
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// <p>agent model name usage</p>
	AgentModel []*string `json:"AgentModel,omitnil,omitempty" name:"AgentModel"`

	// <p>Instance ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>metadata risk list. Has the following enumeration values: 1. AK_TMP 2. USER_DATA</p>
	MetadataRiskList []*string `json:"MetadataRiskList,omitnil,omitempty" name:"MetadataRiskList"`

	// <p>First detection time</p>
	IdentityTimeFirst *string `json:"IdentityTimeFirst,omitnil,omitempty" name:"IdentityTimeFirst"`

	// <p>Latest detected time</p>
	IdentityTimeLast *string `json:"IdentityTimeLast,omitnil,omitempty" name:"IdentityTimeLast"`

	// <p>Detect method. Has the following enumeration values: 1. FINGER Detect via asset fingerprinting 2. NETWORK Detect via network access mode</p>
	IdentityMethod *string `json:"IdentityMethod,omitnil,omitempty" name:"IdentityMethod"`

	// <p>Exposure status. Has the following enumeration values. 1. EXPOSED; 2. UNEXPOSED;</p><ol start="3"><li>UNKNOWN;</li></ol>
	ExposureStatus *string `json:"ExposureStatus,omitnil,omitempty" name:"ExposureStatus"`

	// <p>Corresponding path when metadata is at risk</p>
	MetadataRiskURL *string `json:"MetadataRiskURL,omitnil,omitempty" name:"MetadataRiskURL"`

	// <p>None</p>
	SkillState *SkillState `json:"SkillState,omitnil,omitempty" name:"SkillState"`

	// <p>Traffic sandbox plug-in status</p>
	TrafficPluginState *TrafficPluginState `json:"TrafficPluginState,omitnil,omitempty" name:"TrafficPluginState"`

	// <p>Sandbox rule status for traffic</p>
	TrafficRuleState []*TrafficRuleState `json:"TrafficRuleState,omitnil,omitempty" name:"TrafficRuleState"`

	// <p>Command sandbox plug-in status</p>
	CommandPluginState *CommandPluginState `json:"CommandPluginState,omitnil,omitempty" name:"CommandPluginState"`
}

type AKInfo struct {
	// ak id.
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// ak specific value. returns temporary key when temporary key is used.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Associated account.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type AccessCredentialOutput struct {
	// Credential key name (original), such as SecretId, SecretKey, Token
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Credential key-value (masked)
	// Supplementary description: Reserve the first 3 and last 4 digits, replace the middle with ***; replace all with *** if the length is less than 7.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AccessKeyAlarm struct {
	// <p>Alarm name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Alarm level<br>0-Unavailable 1-Notification 2-Low risk 3-Medium risk 4-High risk 5-Critical</p>
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>Alarm record ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>Alarm rule ID</p>
	AlarmRuleID *int64 `json:"AlarmRuleID,omitnil,omitempty" name:"AlarmRuleID"`

	// <p>Alarm type<br>0 Abnormal call<br>1 Leak monitoring</p>
	AlarmType *int64 `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// <p>Access key</p>
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// <p>Access Key ID</p>
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// <p>Access key remark</p>
	AccessKeyRemark *string `json:"AccessKeyRemark,omitnil,omitempty" name:"AccessKeyRemark"`

	// <p>Last alarm time</p>
	LastAlarmTime *string `json:"LastAlarmTime,omitnil,omitempty" name:"LastAlarmTime"`

	// <p>Alarm status<br>0-unprocessed 1-processed 2-ignored</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Aggregate date</p>
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// <p>Alarm tag</p>
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// <p>Uin of the main account</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>Nickname of the main account</p>
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// <p>Sub-account Uin</p>
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>Sub-account nickname</p>
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// <p>Account type<br>0 Root account AK 1 Sub-account AK 2 Temporary key</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>App ID</p>
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>Leakage evidence</p>
	LeakEvidence []*string `json:"LeakEvidence,omitnil,omitempty" name:"LeakEvidence"`

	// <p>Whether support editing trust account</p>
	IsSupportEditWhiteAccount *bool `json:"IsSupportEditWhiteAccount,omitnil,omitempty" name:"IsSupportEditWhiteAccount"`

	// <p>Alert evidence</p>
	Evidence *string `json:"Evidence,omitnil,omitempty" name:"Evidence"`

	// <p>Alarm rule flag</p>
	RuleKey *string `json:"RuleKey,omitnil,omitempty" name:"RuleKey"`

	// <p>Cloud vendor type 0:Tencent Cloud 1:Amazon Web Services 2:Microsoft Azure 3:Google Cloud 4:Alibaba Cloud 5:Huawei Cloud</p>
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// <p>Alarm AI analysis status<br>-1 Analysis failed<br>0 Not analyzed<br>1 Under analysis<br>2 Analysis successful, real alarm<br>3 Analysis successful, suspicious alarm</p>
	AIStatus *int64 `json:"AIStatus,omitnil,omitempty" name:"AIStatus"`

	// <p>First alarm timestamp (in seconds)</p>
	FirstAlarmTimestamp *int64 `json:"FirstAlarmTimestamp,omitnil,omitempty" name:"FirstAlarmTimestamp"`

	// <p>Last alarm timestamp (in seconds)</p>
	LastAlarmTimestamp *int64 `json:"LastAlarmTimestamp,omitnil,omitempty" name:"LastAlarmTimestamp"`

	// <p>AI analysis failure description. Empty string if not failed.</p>
	AIFailedReason *string `json:"AIFailedReason,omitnil,omitempty" name:"AIFailedReason"`
}

type AccessKeyAlarmCount struct {
	// Access key ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Access key.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Alarm count.
	AlarmCount *int64 `json:"AlarmCount,omitnil,omitempty" name:"AlarmCount"`

	// Security credentials status. valid values: 0 (disabled), 1 (enabled), 2 (deleted).
	AccessKeyStatus *int64 `json:"AccessKeyStatus,omitnil,omitempty" name:"AccessKeyStatus"`

	// AK creation time.
	AccessKeyCreateTime *string `json:"AccessKeyCreateTime,omitnil,omitempty" name:"AccessKeyCreateTime"`

	// AK last usage time. returns "-" if never used.
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`
}

type AccessKeyAlarmInfo struct {
	// Alarm type/risktype.
	// Alarm type:.
	// Abnormal calls.
	// Leakage detection.
	// 2 custom.
	// 
	// Risk type:.
	// Configuration risk.
	// Custom risk.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Alarm count/number of risks.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type AccessKeyAsset struct {
	// AK id.
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// AK name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Account associate APPID.
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Account associate Uin belonging to main account.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Nickname of the main account.
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// Sub-Account Uin belonging to.
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Sub-Account nickname.
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// Root account AK.
	// Sub-Account AK.
	// 2 temporary key.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Security advice enumeration.
	// Normal.
	// Process now.
	// 2 recommend reinforcement.
	Advice *int64 `json:"Advice,omitnil,omitempty" name:"Advice"`

	// Alarm information list.
	AccessKeyAlarmList []*AccessKeyAlarmInfo `json:"AccessKeyAlarmList,omitnil,omitempty" name:"AccessKeyAlarmList"`

	// Risk information list.
	AccessKeyRiskList []*AccessKeyAlarmInfo `json:"AccessKeyRiskList,omitnil,omitempty" name:"AccessKeyRiskList"`

	// Source IP quantity.
	IPCount *int64 `json:"IPCount,omitnil,omitempty" name:"IPCount"`

	// Creation time.
	// 
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last access Time
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// AK status. 
	// 0: disabled.
	// 1: enabled.
	// 2: deleted (deleted in cam, the security center still retains the previous log).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 0 means detected.
	// 1 indicates detecting.
	CheckStatus *int64 `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// Cloud vendor type 0: tencent cloud 1: amazon web services 2: microsoft azure 3: google cloud 4: alibaba cloud 5: huawei cloud.
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

type AccessKeyRisk struct {
	// Risk name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Risk level.
	// 0 - unavailable 1 - Note 2 - low risk 3 - medium risk 4 - high risk 5 - critical.
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// Risk record ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Risk rule ID.
	RiskRuleID *int64 `json:"RiskRuleID,omitnil,omitempty" name:"RiskRuleID"`

	// Risk type.
	// Configuration risk.
	RiskType *int64 `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// Access key.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// Access key remark.
	AccessKeyRemark *string `json:"AccessKeyRemark,omitnil,omitempty" name:"AccessKeyRemark"`

	// Detection time of risk.
	RiskTime *string `json:"RiskTime,omitnil,omitempty" name:"RiskTime"`

	// Risk status.
	// 0 - unprocessed 2 - ignored 3 - converged.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk Tag.
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Risk evidence.
	Evidence *string `json:"Evidence,omitnil,omitempty" name:"Evidence"`

	// Risk description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Account associate Uin belonging to main account.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Nickname of the main account.
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// Sub-Account Uin belonging to.
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Sub-Account nickname.
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// Account type.
	// 0 root account AK 1 sub-account AK.
	// 2 temporary key.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Detection status.
	// 0: detected.
	// 1 indicates detecting.
	CheckStatus *int64 `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// App ID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Query parameter corresponding to the risk.
	QueryParam *string `json:"QueryParam,omitnil,omitempty" name:"QueryParam"`

	// Cloud type 0 for tencent cloud 4 for alibaba cloud.
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// Related AK list, including AK name and remark.
	RelatedAK []*AKInfo `json:"RelatedAK,omitnil,omitempty" name:"RelatedAK"`
}

type AccessKeyUser struct {
	// Account ID.
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 0 root account 1 sub-account.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Access method.
	// 0 API
	// 1 console and API.
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// Security recommendation enumerate 0 normal 1 process immediately 2 recommend reinforcement.
	Advice *int64 `json:"Advice,omitnil,omitempty" name:"Advice"`

	// Alarm information list.
	AccessKeyAlarmList []*AccessKeyAlarmInfo `json:"AccessKeyAlarmList,omitnil,omitempty" name:"AccessKeyAlarmList"`

	// Risk information list.
	AccessKeyRiskList []*AccessKeyAlarmInfo `json:"AccessKeyRiskList,omitnil,omitempty" name:"AccessKeyRiskList"`

	// Account associate APPID.
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Nickname of the main account.
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// Sub-Account nickname.
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// Account Uin belonging to main account.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Account self uin, same as root account uin when it is the root account.
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Login IP.
	LoginIP *string `json:"LoginIP,omitnil,omitempty" name:"LoginIP"`

	// Login address.
	LoginLocation *string `json:"LoginLocation,omitnil,omitempty" name:"LoginLocation"`

	// Log-In time.
	LoginTime *string `json:"LoginTime,omitnil,omitempty" name:"LoginTime"`

	// ISP name
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// Whether operation protection is enabled.
	// 0 not enabled.
	// 1: enabled.
	ActionFlag *int64 `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`

	// Is login protection enabled?.
	// 0 not enabled.
	// 1: enabled.
	LoginFlag *int64 `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 0 means detected. 1 means detecting.
	CheckStatus *int64 `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// Cloud vendor type 0: tencent cloud 1: amazon web services 2: microsoft azure 3: google cloud 4: alibaba cloud 5: huawei cloud.
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

// Predefined struct for user
type AddNewBindRoleUserRequestParams struct {

}

type AddNewBindRoleUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *AddNewBindRoleUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNewBindRoleUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNewBindRoleUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNewBindRoleUserResponseParams struct {
	// `0`: successful. Other values: failed.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddNewBindRoleUserResponse struct {
	*tchttp.BaseResponse
	Response *AddNewBindRoleUserResponseParams `json:"Response"`
}

func (r *AddNewBindRoleUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNewBindRoleUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertExtraInfo struct {
	// Related attack events
	RelateEvent *RelatedEvent `json:"RelateEvent,omitnil,omitempty" name:"RelateEvent"`

	// Leaked content
	LeakContent *string `json:"LeakContent,omitnil,omitempty" name:"LeakContent"`

	// Leak API
	LeakAPI *string `json:"LeakAPI,omitnil,omitempty" name:"LeakAPI"`

	// secretID
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// hit rule
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Rule description
	RuleDesc *string `json:"RuleDesc,omitnil,omitempty" name:"RuleDesc"`

	// Protocol port
	ProtocolPort *string `json:"ProtocolPort,omitnil,omitempty" name:"ProtocolPort"`

	// Attack content
	AttackContent *string `json:"AttackContent,omitnil,omitempty" name:"AttackContent"`

	// Attack IP profiling
	AttackIPProfile *string `json:"AttackIPProfile,omitnil,omitempty" name:"AttackIPProfile"`

	// Attack IP tag
	AttackIPTags *string `json:"AttackIPTags,omitnil,omitempty" name:"AttackIPTags"`

	// Request method
	RequestMethod *string `json:"RequestMethod,omitnil,omitempty" name:"RequestMethod"`

	// HTTP log
	HttpLog *string `json:"HttpLog,omitnil,omitempty" name:"HttpLog"`

	// Attacked domain name
	AttackDomain *string `json:"AttackDomain,omitnil,omitempty" name:"AttackDomain"`

	// File path
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// user_agent
	UserAgent *string `json:"UserAgent,omitnil,omitempty" name:"UserAgent"`

	// Request headers
	RequestHeaders *string `json:"RequestHeaders,omitnil,omitempty" name:"RequestHeaders"`

	// Login username
	LoginUserName *string `json:"LoginUserName,omitnil,omitempty" name:"LoginUserName"`

	// Vulnerability name
	VulnerabilityName *string `json:"VulnerabilityName,omitnil,omitempty" name:"VulnerabilityName"`

	// Public vulnerability and exposure
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Service process
	ServiceProcess *string `json:"ServiceProcess,omitnil,omitempty" name:"ServiceProcess"`

	// Filename
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// File MD5
	FileMD5 *string `json:"FileMD5,omitnil,omitempty" name:"FileMD5"`

	// Last access time of the file
	FileLastAccessTime *string `json:"FileLastAccessTime,omitnil,omitempty" name:"FileLastAccessTime"`

	// file modification time
	FileModifyTime *string `json:"FileModifyTime,omitnil,omitempty" name:"FileModifyTime"`

	// Last access Time
	RecentAccessTime *string `json:"RecentAccessTime,omitnil,omitempty" name:"RecentAccessTime"`

	// Last modification time
	RecentModifyTime *string `json:"RecentModifyTime,omitnil,omitempty" name:"RecentModifyTime"`

	// Virus name
	VirusName *string `json:"VirusName,omitnil,omitempty" name:"VirusName"`

	// Virus file tag
	VirusFileTags *string `json:"VirusFileTags,omitnil,omitempty" name:"VirusFileTags"`

	// behavioral characteristics
	BehavioralCharacteristics *string `json:"BehavioralCharacteristics,omitnil,omitempty" name:"BehavioralCharacteristics"`

	// process name (PID)
	ProcessNamePID *string `json:"ProcessNamePID,omitnil,omitempty" name:"ProcessNamePID"`

	// Process path
	ProcessPath *string `json:"ProcessPath,omitnil,omitempty" name:"ProcessPath"`

	// Command line of the process
	ProcessCommandLine *string `json:"ProcessCommandLine,omitnil,omitempty" name:"ProcessCommandLine"`

	// Process permission
	ProcessPermissions *string `json:"ProcessPermissions,omitnil,omitempty" name:"ProcessPermissions"`

	// Execute commands
	ExecutedCommand *string `json:"ExecutedCommand,omitnil,omitempty" name:"ExecutedCommand"`

	// Affected Filename
	AffectedFileName *string `json:"AffectedFileName,omitnil,omitempty" name:"AffectedFileName"`

	// bait path
	DecoyPath *string `json:"DecoyPath,omitnil,omitempty" name:"DecoyPath"`

	// Malicious process file size
	MaliciousProcessFileSize *string `json:"MaliciousProcessFileSize,omitnil,omitempty" name:"MaliciousProcessFileSize"`

	// Malicious process file MD5
	MaliciousProcessFileMD5 *string `json:"MaliciousProcessFileMD5,omitnil,omitempty" name:"MaliciousProcessFileMD5"`

	// Malicious process name (PID)
	MaliciousProcessNamePID *string `json:"MaliciousProcessNamePID,omitnil,omitempty" name:"MaliciousProcessNamePID"`

	// Malicious process path
	MaliciousProcessPath *string `json:"MaliciousProcessPath,omitnil,omitempty" name:"MaliciousProcessPath"`

	// malicious process start time
	MaliciousProcessStartTime *string `json:"MaliciousProcessStartTime,omitnil,omitempty" name:"MaliciousProcessStartTime"`

	// command content
	CommandContent *string `json:"CommandContent,omitnil,omitempty" name:"CommandContent"`

	// Startup user
	StartupUser *string `json:"StartupUser,omitnil,omitempty" name:"StartupUser"`

	// User group
	UserGroup *string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// Add new permission
	NewPermissions *string `json:"NewPermissions,omitnil,omitempty" name:"NewPermissions"`

	// Parent process
	ParentProcess *string `json:"ParentProcess,omitnil,omitempty" name:"ParentProcess"`

	// Class name
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// class loader
	ClassLoader *string `json:"ClassLoader,omitnil,omitempty" name:"ClassLoader"`

	// File size
	ClassFileSize *string `json:"ClassFileSize,omitnil,omitempty" name:"ClassFileSize"`

	// Class file MD5
	ClassFileMD5 *string `json:"ClassFileMD5,omitnil,omitempty" name:"ClassFileMD5"`

	// Parent class name
	ParentClassName *string `json:"ParentClassName,omitnil,omitempty" name:"ParentClassName"`

	// inherit an API
	InheritedInterface *string `json:"InheritedInterface,omitnil,omitempty" name:"InheritedInterface"`

	// Annotation
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// payload content
	PayloadContent *string `json:"PayloadContent,omitnil,omitempty" name:"PayloadContent"`

	// Callback address profile
	CallbackAddressPortrait *string `json:"CallbackAddressPortrait,omitnil,omitempty" name:"CallbackAddressPortrait"`

	// Callback address tag
	CallbackAddressTag *string `json:"CallbackAddressTag,omitnil,omitempty" name:"CallbackAddressTag"`

	// Process MD5
	ProcessMD5 *string `json:"ProcessMD5,omitnil,omitempty" name:"ProcessMD5"`

	// File permission
	FilePermission *string `json:"FilePermission,omitnil,omitempty" name:"FilePermission"`

	// Information field from log analysis
	FromLogAnalysisData []*KeyValue `json:"FromLogAnalysisData,omitnil,omitempty" name:"FromLogAnalysisData"`

	// probe hit
	HitProbe *string `json:"HitProbe,omitnil,omitempty" name:"HitProbe"`

	// hit honeypot
	HitHoneyPot *string `json:"HitHoneyPot,omitnil,omitempty" name:"HitHoneyPot"`

	// command list
	CommandList *string `json:"CommandList,omitnil,omitempty" name:"CommandList"`

	// Attack event description
	AttackEventDesc *string `json:"AttackEventDesc,omitnil,omitempty" name:"AttackEventDesc"`

	// Process information
	ProcessInfo *string `json:"ProcessInfo,omitnil,omitempty" name:"ProcessInfo"`

	// Login username & password
	UserNameAndPwd *string `json:"UserNameAndPwd,omitnil,omitempty" name:"UserNameAndPwd"`

	// Host protection policy ID
	StrategyID *string `json:"StrategyID,omitnil,omitempty" name:"StrategyID"`

	// Host protection policy name
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// Host protection hit policy is a combination of policy ID and Policy Name
	HitStrategy *string `json:"HitStrategy,omitnil,omitempty" name:"HitStrategy"`

	// Process name
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// PID
	PID *string `json:"PID,omitnil,omitempty" name:"PID"`

	// Container Pod name
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Container Pod ID
	PodID *string `json:"PodID,omitnil,omitempty" name:"PodID"`

	// Http response
	Response *string `json:"Response,omitnil,omitempty" name:"Response"`

	// system call
	SystemCall *string `json:"SystemCall,omitnil,omitempty" name:"SystemCall"`

	// Operation type
	Verb *string `json:"Verb,omitnil,omitempty" name:"Verb"`

	// Log ID.
	LogID *string `json:"LogID,omitnil,omitempty" name:"LogID"`

	// Change content
	Different *string `json:"Different,omitnil,omitempty" name:"Different"`

	// Event type
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Event description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Destination address (container reverse shell)
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// Malicious request domain name (container malicious outbound connection)
	MaliciousRequestDomain *string `json:"MaliciousRequestDomain,omitnil,omitempty" name:"MaliciousRequestDomain"`

	// Rule Type (Container K8sAPI Exception Request)
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Requested Resource (Container K8sAPI Exception Request)
	RequestURI *string `json:"RequestURI,omitnil,omitempty" name:"RequestURI"`

	// Request Initiating User (Container K8sAPI Exception Request)
	RequestUser *string `json:"RequestUser,omitnil,omitempty" name:"RequestUser"`

	// Request Object (Container K8sAPI Exception Request)
	RequestObject *string `json:"RequestObject,omitnil,omitempty" name:"RequestObject"`

	// Response object (container K8sAPI exception request)
	ResponseObject *string `json:"ResponseObject,omitnil,omitempty" name:"ResponseObject"`

	// File type (Container file tamper)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Tag feature (malicious outbound connection of container)
	TIType *string `json:"TIType,omitnil,omitempty" name:"TIType"`

	// Source IP Address (Container K8sAPI Exception Request)
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`
}

type AlertInfo struct {
	// alarm ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// alarm name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Alarm source
	// CFW: Cloud Firewall
	// WAF: Web application firewall
	// CWP: Host Security
	// CSIP: Cloud Security Center
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// alarm level
	// Prompt.
	// 2: Low risk
	// 3: Medium risk
	// 4: High risk
	// 5: Critical
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// attacker
	Attacker *RoleInfo `json:"Attacker,omitnil,omitempty" name:"Attacker"`

	// victim
	Victim *RoleInfo `json:"Victim,omitnil,omitempty" name:"Victim"`

	// Evidence data (such as attack content, base64 encoded)
	EvidenceData *string `json:"EvidenceData,omitnil,omitempty" name:"EvidenceData"`

	// evidence location (for example protocol port)
	EvidenceLocation *string `json:"EvidenceLocation,omitnil,omitempty" name:"EvidenceLocation"`

	// Evidence Path
	EvidencePath *string `json:"EvidencePath,omitnil,omitempty" name:"EvidencePath"`

	// Initial alarm time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Latest Alarm Time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Alarm count
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Emergency Mitigation Suggestions
	UrgentSuggestion *string `json:"UrgentSuggestion,omitnil,omitempty" name:"UrgentSuggestion"`

	// Radical Treatment Suggestion
	RemediationSuggestion *string `json:"RemediationSuggestion,omitnil,omitempty" name:"RemediationSuggestion"`

	// Processing status
	// 0: unprocessed, 1: ignored, 2: processed
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Alarm Handling Type
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// Major Category of Alarm
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Alarm Subcategory
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// Dropdown Field
	ExtraInfo *AlertExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// Aggregate Fields
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Alarm Date
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// appid
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Account name
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// account ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Behavior
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// risk detection
	RiskInvestigation *string `json:"RiskInvestigation,omitnil,omitempty" name:"RiskInvestigation"`

	// Risk handling
	RiskTreatment *string `json:"RiskTreatment,omitnil,omitempty" name:"RiskTreatment"`

	// log type
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Statement retrieval
	LogSearch *string `json:"LogSearch,omitnil,omitempty" name:"LogSearch"`
}

type AssetBaseInfoResponse struct {
	// vpc-id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc-name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Operating system.
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP address
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Asset type
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Number of accounts
	AccountNum *uint64 `json:"AccountNum,omitnil,omitempty" name:"AccountNum"`

	// Number of Ports
	PortNum *uint64 `json:"PortNum,omitnil,omitempty" name:"PortNum"`

	// Process quantity
	ProcessNum *uint64 `json:"ProcessNum,omitnil,omitempty" name:"ProcessNum"`

	// Number of Software Applications
	SoftApplicationNum *uint64 `json:"SoftApplicationNum,omitnil,omitempty" name:"SoftApplicationNum"`

	// Database Count
	DatabaseNum *uint64 `json:"DatabaseNum,omitnil,omitempty" name:"DatabaseNum"`

	// Number of Web Applications
	WebApplicationNum *uint64 `json:"WebApplicationNum,omitnil,omitempty" name:"WebApplicationNum"`

	// Number of services
	ServiceNum *uint64 `json:"ServiceNum,omitnil,omitempty" name:"ServiceNum"`

	// Web Framework Count
	WebFrameworkNum *uint64 `json:"WebFrameworkNum,omitnil,omitempty" name:"WebFrameworkNum"`

	// Website Count
	WebSiteNum *uint64 `json:"WebSiteNum,omitnil,omitempty" name:"WebSiteNum"`

	// Jar Package Count
	JarPackageNum *uint64 `json:"JarPackageNum,omitnil,omitempty" name:"JarPackageNum"`

	// Started Service Count
	StartServiceNum *uint64 `json:"StartServiceNum,omitnil,omitempty" name:"StartServiceNum"`

	// Number of Scheduled Tasks
	ScheduledTaskNum *uint64 `json:"ScheduledTaskNum,omitnil,omitempty" name:"ScheduledTaskNum"`

	// Number of Environment Variables
	EnvironmentVariableNum *uint64 `json:"EnvironmentVariableNum,omitnil,omitempty" name:"EnvironmentVariableNum"`

	// Number of Kernel Modules
	KernelModuleNum *uint64 `json:"KernelModuleNum,omitnil,omitempty" name:"KernelModuleNum"`

	// System Installation Package Count
	SystemInstallationPackageNum *uint64 `json:"SystemInstallationPackageNum,omitnil,omitempty" name:"SystemInstallationPackageNum"`

	// remaining protection duration
	SurplusProtectDay *uint64 `json:"SurplusProtectDay,omitnil,omitempty" name:"SurplusProtectDay"`

	// Whether client is installed. 1 for Installed, 0 for Not Installed.
	CWPStatus *uint64 `json:"CWPStatus,omitnil,omitempty" name:"CWPStatus"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Protection level
	ProtectLevel *string `json:"ProtectLevel,omitnil,omitempty" name:"ProtectLevel"`

	// protection duration
	ProtectedDay *uint64 `json:"ProtectedDay,omitnil,omitempty" name:"ProtectedDay"`
}

type AssetCluster struct {
	// Tenant ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Tenant uin.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Tenant Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Cluster ID.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Cluster name.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Cluster type.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Cluster Creation Time
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Cluster protection status, enumerate on the left, display on the right.
	// Protection status of the cluster. 
	// 0: not connected.
	// Unprotected. 
	// 2: partial protection. 
	// 3: under protection. 
	// 4: access exception. 
	// 5: accessing. 
	// Uninstalling. 
	// 7: uninstallation exception.
	ProtectStatus *int64 `json:"ProtectStatus,omitnil,omitempty" name:"ProtectStatus"`

	// Access information, being empty indicates no access exception info.
	ProtectInfo *string `json:"ProtectInfo,omitnil,omitempty" name:"ProtectInfo"`

	// VPC id.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// kubernetes version.
	KubernetesVersion *string `json:"KubernetesVersion,omitnil,omitempty" name:"KubernetesVersion"`

	// Runtime component.
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Runtime component version.
	ComponentVersion *string `json:"ComponentVersion,omitnil,omitempty" name:"ComponentVersion"`

	// Component status.
	ComponentStatus *string `json:"ComponentStatus,omitnil,omitempty" name:"ComponentStatus"`

	// Health Checkup Time
	CheckTime *string `json:"CheckTime,omitnil,omitempty" name:"CheckTime"`

	// Associated hosts.
	MachineCount *int64 `json:"MachineCount,omitnil,omitempty" name:"MachineCount"`

	// Associated Pod Count
	PodCount *int64 `json:"PodCount,omitnil,omitempty" name:"PodCount"`

	// Associated Service Count
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// Vulnerability risk.
	VulRisk *int64 `json:"VulRisk,omitnil,omitempty" name:"VulRisk"`

	// Configuration risk.
	CFGRisk *int64 `json:"CFGRisk,omitnil,omitempty" name:"CFGRisk"`

	// Health Checkup Count
	CheckCount *int64 `json:"CheckCount,omitnil,omitempty" name:"CheckCount"`

	// Whether it is core. 1: Core; 2: Non-core.
	IsCore *int64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Cloud asset type: 0: tencent cloud, 1: aws, 2: azure.
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

type AssetClusterPod struct {
	// Tenant ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Tenant UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Tenant name
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Pod ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Pod name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Pod Creation Time
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Host ID
	MachineId *string `json:"MachineId,omitnil,omitempty" name:"MachineId"`

	// host name
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// pod ip
	PodIp *string `json:"PodIp,omitnil,omitempty" name:"PodIp"`

	// Associated Service Count
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// Associated container number
	ContainerCount *int64 `json:"ContainerCount,omitnil,omitempty" name:"ContainerCount"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP address
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Whether it is core. 1: Core; 2: Non-core.
	IsCore *int64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type AssetInfoDetail struct {
	// User appid.
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// CVE id
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`

	// Scan Status. 0-Not Scanned by Default; 1-Scanning; 2-Scan Completed; 3-Scan Error.
	IsScan *int64 `json:"IsScan,omitnil,omitempty" name:"IsScan"`

	// Number of Affected Assets
	InfluenceAsset *int64 `json:"InfluenceAsset,omitnil,omitempty" name:"InfluenceAsset"`

	// Number of Unfixed Assets
	NotRepairAsset *int64 `json:"NotRepairAsset,omitnil,omitempty" name:"NotRepairAsset"`

	// Unprotected Asset Count
	NotProtectAsset *int64 `json:"NotProtectAsset,omitnil,omitempty" name:"NotProtectAsset"`

	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task Percentage
	TaskPercent *int64 `json:"TaskPercent,omitnil,omitempty" name:"TaskPercent"`

	// Task Time
	TaskTime *int64 `json:"TaskTime,omitnil,omitempty" name:"TaskTime"`

	// Scan time
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`
}

type AssetInstanceTypeMap struct {
	// Asset type.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Asset type.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Mapping of asset type and instance type.
	InstanceTypeList []*FilterDataObject `json:"InstanceTypeList,omitnil,omitempty" name:"InstanceTypeList"`
}

type AssetProcessItem struct {
	// Cloud account ID.
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Tenant ID.
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Account name.
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// Instance ID.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP address
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Process ID
	ProcessID *string `json:"ProcessID,omitnil,omitempty" name:"ProcessID"`

	// Process name
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// Command line
	CmdLine *string `json:"CmdLine,omitnil,omitempty" name:"CmdLine"`

	// Listening port list.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type AssetRiskItem struct {
	// <p>Tenant ID</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Cloud vendor</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Cloud vendor name</p>
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// <p>Cloud account name</p>
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// <p>Cloud Account ID</p>
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>First discovery time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Risk status</p>
	RiskStatus *int64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// <p>Risk name</p>
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// <p>Check type</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>Risk level</p>
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// <p>Risk rule ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>Disposal categorization</p>
	Classify *string `json:"Classify,omitnil,omitempty" name:"Classify"`

	// <p>Cybersecurity classified protection compliance</p>
	StandardTerms []*StandardTerm `json:"StandardTerms,omitnil,omitempty" name:"StandardTerms"`

	// <p>Asset type</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>Asset type icon</p>
	AssetTypeIconURL *string `json:"AssetTypeIconURL,omitnil,omitempty" name:"AssetTypeIconURL"`

	// <p>Asset type</p>
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`
}

type AssetTag struct {
	// Tag Key, can be letters, digits, and underscores.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag Value, can be letters, digits, and underscores.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type AssetViewCFGRisk struct {
	// The unique ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Configuration name
	CFGName *string `json:"CFGName,omitnil,omitempty" name:"CFGName"`

	// Check type
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level, low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// Source of the task
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// relevant standards
	CFGSTD *string `json:"CFGSTD,omitnil,omitempty" name:"CFGSTD"`

	// Configuration details.
	CFGDescribe *string `json:"CFGDescribe,omitnil,omitempty" name:"CFGDescribe"`

	// Fix suggestion
	CFGFix *string `json:"CFGFix,omitnil,omitempty" name:"CFGFix"`

	// URL of the help documentation
	CFGHelpURL *string `json:"CFGHelpURL,omitnil,omitempty" name:"CFGHelpURL"`

	// Data entry key
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// When the asset type is LBL, show this field to locate the specific LB.
	ClbId *string `json:"ClbId,omitnil,omitempty" name:"ClbId"`
}

type AssetViewPortRisk struct {
	// Port
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level, low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Network protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Suggested action. `0`: Keep as it is; `1`: Block access requests; `2`: Block the port
	Suggestion *uint64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Status, 0 unprocessed, 1 processed, 2 ignored, 3 defended by cloud protection
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Recognition Source. See Enumeration Return for details.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Service judgment, high-risk service, web service, other service
	ServiceJudge *string `json:"ServiceJudge,omitnil,omitempty" name:"ServiceJudge"`

	// Status, 0 unprocessed, 1 processed, 2 ignored, 3 defended by cloud protection, 4 no action is required
	XspmStatus *uint64 `json:"XspmStatus,omitnil,omitempty" name:"XspmStatus"`
}

type AssetViewVULRisk struct {
	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level: low - low risk, high - high risk, middle - medium risk, info - note, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Vulnerability type
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Vulnerability description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Vulnerability impact component.
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Technology reference.
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// Vulnerability impact version.
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// Risks.
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// CVE number
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Fixing solution
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// POC ID
	POCId *string `json:"POCId,omitnil,omitempty" name:"POCId"`

	// Scan Source
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// CWPP edition
	CWPVersion *int64 `json:"CWPVersion,omitnil,omitempty" name:"CWPVersion"`

	// Whether it can be fixed 
	IsSupportRepair *bool `json:"IsSupportRepair,omitnil,omitempty" name:"IsSupportRepair"`

	// Whether it can be detected
	IsSupportDetect *bool `json:"IsSupportDetect,omitnil,omitempty" name:"IsSupportDetect"`

	// Instance UUID
	InstanceUUID *string `json:"InstanceUUID,omitnil,omitempty" name:"InstanceUUID"`

	// Payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Emergency Vulnerability Type. 1-Emergency Vulnerability; 0-Non-emergency Vulnerability.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`
}

type AssetViewVULRiskData struct {
	// Impact assets.
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level: low - low risk, high - high risk, middle - medium risk, info - note, extreme - serious.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Asset type.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Component.
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Latest Recognition Time
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First Recognition Time
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status, 0 unprocessed, 1 tagged, 2 ignored, 3 processed, 4 under disposal, 5 detecting, 6 partially processed.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User appid.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Vulnerability type.
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Port.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Vulnerability impact component.
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Vulnerability impact version.
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// Risks.
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// pocid
	POCId *string `json:"POCId,omitnil,omitempty" name:"POCId"`

	// Scan Source
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Host version.
	CWPVersion *int64 `json:"CWPVersion,omitnil,omitempty" name:"CWPVersion"`

	// Instance UUID
	InstanceUUID *string `json:"InstanceUUID,omitnil,omitempty" name:"InstanceUUID"`

	// Payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Emergency Vulnerability Type. 1-Emergency Vulnerability; 0-Non-emergency Vulnerability.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// CVSS score
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// Frontend index id.
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// pcmgrId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`

	// Report ID
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Vulnerability Tag.
	VulTag []*string `json:"VulTag,omitnil,omitempty" name:"VulTag"`

	// Vulnerability disclosure time.
	DisclosureTime *string `json:"DisclosureTime,omitnil,omitempty" name:"DisclosureTime"`

	// Attack intensity.
	AttackHeat *uint64 `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// Whether the vulnerability is mandatory. 1 for yes, 0 for no.
	IsSuggest *int64 `json:"IsSuggest,omitnil,omitempty" name:"IsSuggest"`

	// Disposal task ID.
	HandleTaskId *string `json:"HandleTaskId,omitnil,omitempty" name:"HandleTaskId"`

	// Engine source.
	EngineSource *string `json:"EngineSource,omitnil,omitempty" name:"EngineSource"`

	// New vulnerability risk id (same as RiskId in the network-wide vulnerabilities table).
	VulRiskId *string `json:"VulRiskId,omitnil,omitempty" name:"VulRiskId"`

	// New version vulnerability id.
	TvdID *string `json:"TvdID,omitnil,omitempty" name:"TvdID"`

	// Is it possible to perform a one-click physical examination, 1 - yes, 0 - not allowed.
	IsOneClick *uint64 `json:"IsOneClick,omitnil,omitempty" name:"IsOneClick"`

	// Whether to perform a POC scan. valid values: 0 (not a POC), 1 (POC).
	IsPOC *uint64 `json:"IsPOC,omitnil,omitempty" name:"IsPOC"`
}

type AssetViewWeakPassRisk struct {
	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level: low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID, handle risk usage
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Weak password type
	PasswordType *string `json:"PasswordType,omitnil,omitempty" name:"PasswordType"`

	// Source of the task
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Vulnerability type
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Vulnerability URL
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// Fix suggestion
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// proof
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type AttributeOptionSet struct {
	// cvm instance type
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// CVM instance name.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BugInfoDetail struct {
	// Vulnerability ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// PocId Corresponding to Vulnerability
	PatchId *string `json:"PatchId,omitnil,omitempty" name:"PatchId"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// Vulnerability Severity: High, Middle, Low, Info.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// CVSS Score
	CVSSScore *string `json:"CVSSScore,omitnil,omitempty" name:"CVSSScore"`

	// CVE ID
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`

	// Vulnerability tag
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Vulnerability Type. 1: Web Application; 2: System Component Vulnerabilities; 3: Configuration Risk.
	VULCategory *uint64 `json:"VULCategory,omitnil,omitempty" name:"VULCategory"`

	// Vulnerability Affected System
	ImpactOs *string `json:"ImpactOs,omitnil,omitempty" name:"ImpactOs"`

	// Affected component
	ImpactCOMPENT *string `json:"ImpactCOMPENT,omitnil,omitempty" name:"ImpactCOMPENT"`

	// Vulnerability Affected Version
	ImpactVersion *string `json:"ImpactVersion,omitnil,omitempty" name:"ImpactVersion"`

	// Link
	Reference *string `json:"Reference,omitnil,omitempty" name:"Reference"`

	// Vulnerability description
	VULDescribe *string `json:"VULDescribe,omitnil,omitempty" name:"VULDescribe"`

	// Fixing suggestion
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Product Support Status, returned in real time.
	ProSupport *uint64 `json:"ProSupport,omitnil,omitempty" name:"ProSupport"`

	// Published or Not. 0 for No, 1 for Yes.
	IsPublish *uint64 `json:"IsPublish,omitnil,omitempty" name:"IsPublish"`

	// Release time.
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Vulnerability Subcategory
	SubCategory *string `json:"SubCategory,omitnil,omitempty" name:"SubCategory"`
}

type CFGViewCFGRisk struct {
	// Impact assets.
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// Risk level: low - low risk, high - high risk, middle - medium risk, info - note, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Latest Recognition Time
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First Recognition Time
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status. 0-Unprocessed; 1-Disposed; 2-Ignored.
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// Unique ID of Asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Asset Subtype
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Front-end Index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// User appid.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Configuration name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFGName *string `json:"CFGName,omitnil,omitempty" name:"CFGName"`

	// Check type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// -
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFGSTD *string `json:"CFGSTD,omitnil,omitempty" name:"CFGSTD"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFGDescribe *string `json:"CFGDescribe,omitnil,omitempty" name:"CFGDescribe"`

	// Fixing suggestion
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFGFix *string `json:"CFGFix,omitnil,omitempty" name:"CFGFix"`

	// Help documentation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFGHelpURL *string `json:"CFGHelpURL,omitnil,omitempty" name:"CFGHelpURL"`
}

type CICDToken struct {
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>appid</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>CI/CD name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Token for integration</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// <p>Scanning result storage duration</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Scanned file</p>
	FileCnt *uint64 `json:"FileCnt,omitnil,omitempty" name:"FileCnt"`

	// <p>Latest scan status</p>
	LastScanStatus *string `json:"LastScanStatus,omitnil,omitempty" name:"LastScanStatus"`

	// <p>Last scan time.</p>
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`
}

type CVMAssetVO struct {
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Protection status
	CWPStatus *uint64 `json:"CWPStatus,omitnil,omitempty" name:"CWPStatus"`

	// Asset creation time.
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC Name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// App ID information
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Nickname.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Availability zone
	AvailableArea *string `json:"AvailableArea,omitnil,omitempty" name:"AvailableArea"`

	// Core or Not
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Subnet Name
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// CWP Agent UUID.
	InstanceUuid *string `json:"InstanceUuid,omitnil,omitempty" name:"InstanceUuid"`

	// CVM host UUID.
	InstanceQUuid *string `json:"InstanceQUuid,omitnil,omitempty" name:"InstanceQUuid"`

	// OS Name
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// Partition
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// CPU Information
	CPUInfo *string `json:"CPUInfo,omitnil,omitempty" name:"CPUInfo"`

	// CPU Size
	CPUSize *uint64 `json:"CPUSize,omitnil,omitempty" name:"CPUSize"`

	// CPU Load
	CPULoad *string `json:"CPULoad,omitnil,omitempty" name:"CPULoad"`

	// Memory size.
	MemorySize *string `json:"MemorySize,omitnil,omitempty" name:"MemorySize"`

	// Memory Load
	MemoryLoad *string `json:"MemoryLoad,omitnil,omitempty" name:"MemoryLoad"`

	// Hard disk size.
	DiskSize *string `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Hard Disk Load
	DiskLoad *string `json:"DiskLoad,omitnil,omitempty" name:"DiskLoad"`

	// Number of Accounts
	AccountCount *string `json:"AccountCount,omitnil,omitempty" name:"AccountCount"`

	// Number of Processes
	ProcessCount *string `json:"ProcessCount,omitnil,omitempty" name:"ProcessCount"`

	// Software application.
	AppCount *string `json:"AppCount,omitnil,omitempty" name:"AppCount"`

	// Listening port
	PortCount *uint64 `json:"PortCount,omitnil,omitempty" name:"PortCount"`

	// Network attack.
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// Network access.
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Network Interception
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// Inbound peak bandwidth.
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// Outbound peak bandwidth.
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// Total inbound traffic.
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// Outbound cumulative traffic.
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Malicious outgoing request.
	NetWorkOut *uint64 `json:"NetWorkOut,omitnil,omitempty" name:"NetWorkOut"`

	// Port risk.
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// Vulnerability risk.
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// Configuration risk.
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Scan Task Count
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// memberId
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Full OS Name
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// Risk service exposure.
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// Simulated Attack Tool Status. 0 indicates not installed. 1 indicates installed. 2 indicates offline.
	BASAgentStatus *int64 `json:"BASAgentStatus,omitnil,omitempty" name:"BASAgentStatus"`

	// 1-New Asset; 0-Not a New Asset
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// 0: not installed; 1: install; 2: installing.
	CVMAgentStatus *int64 `json:"CVMAgentStatus,omitnil,omitempty" name:"CVMAgentStatus"`

	// 1: enable 0: not enabled.
	CVMStatus *int64 `json:"CVMStatus,omitnil,omitempty" name:"CVMStatus"`

	// 1: client installed 0: not installed 2: Agentless.
	DefenseModel *int64 `json:"DefenseModel,omitnil,omitempty" name:"DefenseModel"`

	// 1: installed 0: not installed.
	TatStatus *int64 `json:"TatStatus,omitnil,omitempty" name:"TatStatus"`

	// cpu trend chart.
	CpuTrend []*Element `json:"CpuTrend,omitnil,omitempty" name:"CpuTrend"`

	// Memory trend chart.
	MemoryTrend []*Element `json:"MemoryTrend,omitnil,omitempty" name:"MemoryTrend"`

	// 1: agent online 0: agent offline 2: host offline.
	AgentStatus *int64 `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// Number of shutdowns this month.
	CloseDefenseCount *int64 `json:"CloseDefenseCount,omitnil,omitempty" name:"CloseDefenseCount"`

	// Running state.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Security group data.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Physical memory occupied KB.
	AgentMemRss *int64 `json:"AgentMemRss,omitnil,omitempty" name:"AgentMemRss"`

	// CPU utilization percentage.
	AgentCpuPer *float64 `json:"AgentCpuPer,omitnil,omitempty" name:"AgentCpuPer"`

	// Actual appid belonging to cvm.
	RealAppid *int64 `json:"RealAppid,omitnil,omitempty" name:"RealAppid"`

	// Cloud asset type: 0: tencent cloud, 1: aws, 2: azure.
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// Host protection status enumeration.
	// 0: not installed.
	// Basic edition protection.
	// 2: inclusive edition protection.
	// 3: protection by pro edition.
	// 4: ultimate edition protection.
	// 5: offline.
	// 6: shutdown.
	ProtectStatus *int64 `json:"ProtectStatus,omitnil,omitempty" name:"ProtectStatus"`

	// Last offline time.
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`
}

type CallRecord struct {
	// Invocation record ID.
	CallID *string `json:"CallID,omitnil,omitempty" name:"CallID"`

	// Access key.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Access key remark.
	AccessKeyRemark *string `json:"AccessKeyRemark,omitnil,omitempty" name:"AccessKeyRemark"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// Source IP of the call.
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// Source IP of the call remark.
	SourceIPRemark *string `json:"SourceIPRemark,omitnil,omitempty" name:"SourceIPRemark"`

	// Source IP region of the call.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// IP type 0: within the account (unremarked) 1: outside the account (unremarked) 2: within the account (remarked) 3: outside the account (remarked).
	IPType *int64 `json:"IPType,omitnil,omitempty" name:"IPType"`

	// Call interface name.
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Call the product name.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Invocation type.
	// 0: console invocation.
	// 1:API
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Type of user: CAMUser/root/AssumedRole.
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// User/Role name.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Policy List
	PolicySet []*string `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// Number of calls.
	CallCount *int64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// Error code.
	// 0: Successful
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// First time call time.
	FirstCallTime *string `json:"FirstCallTime,omitnil,omitempty" name:"FirstCallTime"`

	// Call time.
	LastCallTime *string `json:"LastCallTime,omitnil,omitempty" name:"LastCallTime"`

	// IP associated asset ID. if an empty string, means not associated with.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Associated asset name of the IP.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Aggregate date.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// appid
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Display status.
	ShowStatus *bool `json:"ShowStatus,omitnil,omitempty" name:"ShowStatus"`

	// Carrier.
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// vpc information list outside the account.
	VpcInfo []*SourceIPVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// Request client list.
	ReqClient []*string `json:"ReqClient,omitnil,omitempty" name:"ReqClient"`
}

type CheckViewRiskItem struct {
	// <p>Check item rule ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>Risk name</p>
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// <p>Check type</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>Risk level</p>
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// <p>1 risk item exists</p>
	RiskDesc *string `json:"RiskDesc,omitnil,omitempty" name:"RiskDesc"`

	// <p>First discovery time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Risk update time</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Cloud vendor</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Risk status</p>
	RiskStatus *uint64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// <p>Number of affected assets</p>
	AssetCount *uint64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// <p>Number of risks</p>
	RiskCount *uint64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// <p>Asset type</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>Event type</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>Disposal categorization</p>
	Classify *string `json:"Classify,omitnil,omitempty" name:"Classify"`

	// <p>cspm standard clauses</p>
	StandardTerms []*StandardTerm `json:"StandardTerms,omitnil,omitempty" name:"StandardTerms"`

	// <p>Asset type icon</p>
	AssetTypeIconURL *string `json:"AssetTypeIconURL,omitnil,omitempty" name:"AssetTypeIconURL"`
}

type ClbListenerListInfo struct {
	// Listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// listener name
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// CLB Id
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB name
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Load balancing ip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port.
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Region.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC id
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`

	// CLB Type
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// Listener Domain Name
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// CLB domain name
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`
}

type CloudCountDesc struct {
	// 0 means Tencent Cloud
	// 1 indicates AWS
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// Account Quantity
	CloudCount *int64 `json:"CloudCount,omitnil,omitempty" name:"CloudCount"`

	// Description of The Cloud Account Type
	CloudDesc *string `json:"CloudDesc,omitnil,omitempty" name:"CloudDesc"`
}

type CommandPluginState struct {
	// <p>Plug-in installation status (upper layer aggregation)<br>Enumeration value:<br>NONE: Not installed<br>INSTALLING: Installing<br>INSTALLED: Installed<br>INSTALL_FAIL: Installation failure</p>
	InstallStatus *string `json:"InstallStatus,omitnil,omitempty" name:"InstallStatus"`
}

// Predefined struct for user
type CreateAccessKeyCheckTaskRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Risk list.
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`

	// Access key list.
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// Account uin list.
	SubUinList []*string `json:"SubUinList,omitnil,omitempty" name:"SubUinList"`

	// Risk rule id list.
	RiskRuleIDList []*int64 `json:"RiskRuleIDList,omitnil,omitempty" name:"RiskRuleIDList"`
}

type CreateAccessKeyCheckTaskRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Risk list.
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`

	// Access key list.
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// Account uin list.
	SubUinList []*string `json:"SubUinList,omitnil,omitempty" name:"SubUinList"`

	// Risk rule id list.
	RiskRuleIDList []*int64 `json:"RiskRuleIDList,omitnil,omitempty" name:"RiskRuleIDList"`
}

func (r *CreateAccessKeyCheckTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyCheckTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "RiskIDList")
	delete(f, "AccessKeyList")
	delete(f, "SubUinList")
	delete(f, "RiskRuleIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessKeyCheckTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeyCheckTaskResponseParams struct {
	// 0 indicates success. 1 indicates failure.
	Code *uint64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Error message
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessKeyCheckTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessKeyCheckTaskResponseParams `json:"Response"`
}

func (r *CreateAccessKeyCheckTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeySyncTaskRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CreateAccessKeySyncTaskRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CreateAccessKeySyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeySyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessKeySyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeySyncTaskResponseParams struct {
	// Initiate a sync task.
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 0: success; 1: failure.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Error message
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessKeySyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessKeySyncTaskResponseParams `json:"Response"`
}

func (r *CreateAccessKeySyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeySyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainAndIpRequestParams struct {
	// Public IP/domain name
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// Public IP/domain name
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDomainAndIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAndIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "MemberId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainAndIpResponseParams struct {
	// Number of created assets
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainAndIpResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainAndIpResponseParams `json:"Response"`
}

func (r *CreateDomainAndIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCAccessTokenRequestParams struct {
	// <p>CI/CD name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Result storage duration (30/60/90/120/150/180 days)</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type CreateIaCAccessTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>CI/CD name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Result storage duration (30/60/90/120/150/180 days)</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *CreateIaCAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIaCAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCAccessTokenResponseParams struct {
	// <p>Token integration</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIaCAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateIaCAccessTokenResponseParams `json:"Response"`
}

func (r *CreateIaCAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileExportJobRequestParams struct {
	// <p>Filter conditions.</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CreateIaCFileExportJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>Filter conditions.</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CreateIaCFileExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIaCFileExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileExportJobResponseParams struct {
	// <p>Task ID.</p>
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIaCFileExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateIaCFileExportJobResponseParams `json:"Response"`
}

func (r *CreateIaCFileExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileReScanTaskRequestParams struct {
	// <p>File ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CreateIaCFileReScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>File ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CreateIaCFileReScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileReScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIaCFileReScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileReScanTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIaCFileReScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateIaCFileReScanTaskResponseParams `json:"Response"`
}

func (r *CreateIaCFileReScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileReScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskCenterScanTaskRequestParams struct {
	// Task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Values: `0` (Scan all); `1` (Scan specific assets); `2` (Scan all expect the specified assets); `3` (Custom assets). When `ScanAssetType=1/2`, `Assets` is required. When `ScanAssetType=3`, `SelfDefiningAssets` is required. 
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// Project to scan: port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// Task type. `0`: Scheduled task, `1`: Scan immediately; `2`: Scanned at the specified time; `3`: Custom. When ScanPlanType=0,2,3, `ScanPlanContent` is required.
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// List of assets to scan
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// Details of a scheduled scan task
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// IP/Domain name/URL
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Request initiation source, vss means vulnerability scan service, the user of CSC fill in csip, default csip
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// Advanced settings
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// Scan task mode: `0` (Standard), `1` (Quick), `2` (Advanced). Default: `0`
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Asset tags
	Tags *AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Task completed callback webhook url
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

type CreateRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Values: `0` (Scan all); `1` (Scan specific assets); `2` (Scan all expect the specified assets); `3` (Custom assets). When `ScanAssetType=1/2`, `Assets` is required. When `ScanAssetType=3`, `SelfDefiningAssets` is required. 
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// Project to scan: port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// Task type. `0`: Scheduled task, `1`: Scan immediately; `2`: Scanned at the specified time; `3`: Custom. When ScanPlanType=0,2,3, `ScanPlanContent` is required.
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// List of assets to scan
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// Details of a scheduled scan task
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// IP/Domain name/URL
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Request initiation source, vss means vulnerability scan service, the user of CSC fill in csip, default csip
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// Advanced settings
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// Scan task mode: `0` (Standard), `1` (Quick), `2` (Advanced). Default: `0`
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Asset tags
	Tags *AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Task completed callback webhook url
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

func (r *CreateRiskCenterScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskCenterScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "ScanAssetType")
	delete(f, "ScanItem")
	delete(f, "ScanPlanType")
	delete(f, "MemberId")
	delete(f, "Assets")
	delete(f, "ScanPlanContent")
	delete(f, "SelfDefiningAssets")
	delete(f, "ScanFrom")
	delete(f, "TaskAdvanceCFG")
	delete(f, "TaskMode")
	delete(f, "Tags")
	delete(f, "FinishWebHook")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRiskCenterScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskCenterScanTaskResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// `0`: Task created successfully. `-1`: There are unauthorized assets. 
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// List of unauthorized assets
	UnAuthAsset []*string `json:"UnAuthAsset,omitnil,omitempty" name:"UnAuthAsset"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRiskCenterScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRiskCenterScanTaskResponseParams `json:"Response"`
}

func (r *CreateRiskCenterScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskCenterScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillScanRequestParams struct {
	// Base64 encoding of the ZIP file content
	// Input limit: File size limit 7MB (before encoding), only effective ZIP format.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// Filename for server log
	// Parameter format: such as my-skill.zip
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type CreateSkillScanRequest struct {
	*tchttp.BaseRequest
	
	// Base64 encoding of the ZIP file content
	// Input limit: File size limit 7MB (before encoding), only effective ZIP format.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// Filename for server log
	// Parameter format: such as my-skill.zip
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

func (r *CreateSkillScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileBase64")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSkillScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillScanResponseParams struct {
	// SHA256 Hash of the file, used for polling the DescribeSkillScanResult API
	// Parameter format: sha256:<64-bit hex>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// Engine version number actually bound to the current request. The caller should save and explicitly input it in the follow-up DescribeSkillScanResult.
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Task status, fixed as SCANNING, indicates the task is received.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Readable operation result description
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSkillScanResponse struct {
	*tchttp.BaseResponse
	Response *CreateSkillScanResponseParams `json:"Response"`
}

func (r *CreateSkillScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialEffectScope struct {
	// Whether to exclude the mode
	// Enumeration values:
	// 0: Inclusion mode (only takes effect on the Real Server in Instances). At this point, Instances is required.
	// 1: Exclusion mode (Machines in Instances do not take effect, remaining machines take effect). At this point, Instances is selectable (Empty list means all machines take effect).
	Exclude *int64 `json:"Exclude,omitnil,omitempty" name:"Exclude"`

	// Machine instance ID list. Required when Exclude is 0, means only these machines can access the credential; Option when Exclude is 1, means these machines cannot access the credential (Empty list means all machines take effect).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Instances []*string `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type CsipRiskCenterStatistics struct {
	// Total Number of Port Risks
	PortTotal *uint64 `json:"PortTotal,omitnil,omitempty" name:"PortTotal"`

	// High Port Risk Count
	PortHighLevel *uint64 `json:"PortHighLevel,omitnil,omitempty" name:"PortHighLevel"`

	// 	Total number of weak password risks.
	WeakPasswordTotal *uint64 `json:"WeakPasswordTotal,omitnil,omitempty" name:"WeakPasswordTotal"`

	// High Weak Password Risk Count
	WeakPasswordHighLevel *uint64 `json:"WeakPasswordHighLevel,omitnil,omitempty" name:"WeakPasswordHighLevel"`

	// Website Risk Count
	WebsiteTotal *uint64 `json:"WebsiteTotal,omitnil,omitempty" name:"WebsiteTotal"`

	// Number of High Risks on Websites
	WebsiteHighLevel *uint64 `json:"WebsiteHighLevel,omitnil,omitempty" name:"WebsiteHighLevel"`

	// Time of the Latest Scan
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Number of vulnerability risks.
	VULTotal *int64 `json:"VULTotal,omitnil,omitempty" name:"VULTotal"`

	// Number of High-Risk Vulnerability Risks
	VULHighLevel *int64 `json:"VULHighLevel,omitnil,omitempty" name:"VULHighLevel"`

	// Number of Configuration Item Risks
	CFGTotal *int64 `json:"CFGTotal,omitnil,omitempty" name:"CFGTotal"`

	// Number of High-Risk Configuration Item Risks
	CFGHighLevel *int64 `json:"CFGHighLevel,omitnil,omitempty" name:"CFGHighLevel"`

	// Mapping Service Risk Count
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServerTotal *int64 `json:"ServerTotal,omitnil,omitempty" name:"ServerTotal"`

	// High Mapping Service Risk Count
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServerHighLevel *int64 `json:"ServerHighLevel,omitnil,omitempty" name:"ServerHighLevel"`

	// Number of host baseline risks.
	HostBaseLineRiskTotal *int64 `json:"HostBaseLineRiskTotal,omitnil,omitempty" name:"HostBaseLineRiskTotal"`

	// Number of high-risk risks.
	HostBaseLineRiskHighLevel *int64 `json:"HostBaseLineRiskHighLevel,omitnil,omitempty" name:"HostBaseLineRiskHighLevel"`

	// Baseline risk count of the container.
	PodBaseLineRiskTotal *int64 `json:"PodBaseLineRiskTotal,omitnil,omitempty" name:"PodBaseLineRiskTotal"`

	// Number of high-risk baseline risks in the container.
	PodBaseLineRiskHighLevel *int64 `json:"PodBaseLineRiskHighLevel,omitnil,omitempty" name:"PodBaseLineRiskHighLevel"`
}

type DBAssetVO struct {
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// vpcid
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc Tag.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Domain
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Asset creation time.
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Configuration risk.
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Network attack.
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// Network access.
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Scan Task
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// User appid.
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Nickname Alias
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Port.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Private IP address
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Status.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Core or Not
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type DataSearchBug struct {
	// Query status code
	StateCode *string `json:"StateCode,omitnil,omitempty" name:"StateCode"`

	// Vulnerability details
	DataBug []*BugInfoDetail `json:"DataBug,omitnil,omitempty" name:"DataBug"`

	// Vulnerability impact assets details
	DataAsset []*AssetInfoDetail `json:"DataAsset,omitnil,omitempty" name:"DataAsset"`

	// True supports scanning. False does not support scanning.
	VSSScan *bool `json:"VSSScan,omitnil,omitempty" name:"VSSScan"`

	// 0-Not Supported; 1-Supported
	CWPScan *string `json:"CWPScan,omitnil,omitempty" name:"CWPScan"`

	// 1 indicates virtual patches supported, 0 or null indicates not supported.
	CFWPatch *string `json:"CFWPatch,omitnil,omitempty" name:"CFWPatch"`

	// 0-Not Supported; 1-Supported
	WafPatch *int64 `json:"WafPatch,omitnil,omitempty" name:"WafPatch"`

	// 0-Not Supported; 1-Supported
	CWPFix *int64 `json:"CWPFix,omitnil,omitempty" name:"CWPFix"`

	// Product Support Status
	DataSupport []*ProductSupport `json:"DataSupport,omitnil,omitempty" name:"DataSupport"`

	// cveId
	CveId *string `json:"CveId,omitnil,omitempty" name:"CveId"`
}

type DbAssetInfo struct {
	// Cloud Defense Status
	CFWStatus *uint64 `json:"CFWStatus,omitnil,omitempty" name:"CFWStatus"`

	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// VPC information
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Asset type
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// VPC IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// VPC information
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Cloud Defense Protection Edition
	CFWProtectLevel *uint64 `json:"CFWProtectLevel,omitnil,omitempty" name:"CFWProtectLevel"`

	// Tag Information
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`
}

// Predefined struct for user
type DeleteDomainAndIpRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// asset
	Content []*PublicIpDomainListKey `json:"Content,omitnil,omitempty" name:"Content"`

	// Whether to retain the path configuration. `1`: Retain; Others: Do not retain. It defaults to do not retain if not specified.
	RetainPath *int64 `json:"RetainPath,omitnil,omitempty" name:"RetainPath"`

	// Whether to ignore this asset in the future. `1`: Ignore; Others: Do not ignore. It defaults to ignore if not specified.
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil,omitempty" name:"IgnoreAsset"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Deletion mode. Values: `ALL` (delete all). If it's not specified, `Content` is required.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeleteDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// asset
	Content []*PublicIpDomainListKey `json:"Content,omitnil,omitempty" name:"Content"`

	// Whether to retain the path configuration. `1`: Retain; Others: Do not retain. It defaults to do not retain if not specified.
	RetainPath *int64 `json:"RetainPath,omitnil,omitempty" name:"RetainPath"`

	// Whether to ignore this asset in the future. `1`: Ignore; Others: Do not ignore. It defaults to ignore if not specified.
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil,omitempty" name:"IgnoreAsset"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Deletion mode. Values: `ALL` (delete all). If it's not specified, `Content` is required.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DeleteDomainAndIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAndIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Content")
	delete(f, "RetainPath")
	delete(f, "IgnoreAsset")
	delete(f, "Tags")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainAndIpResponseParams struct {
	// Number of deleted assets
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainAndIpResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainAndIpResponseParams `json:"Response"`
}

func (r *DeleteDomainAndIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCAccessTokenRequestParams struct {
	// <p>Delete ID list</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteIaCAccessTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>Delete ID list</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteIaCAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIaCAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCAccessTokenResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIaCAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIaCAccessTokenResponseParams `json:"Response"`
}

func (r *DeleteIaCAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCFileRequestParams struct {
	// <p>Delete ID list</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteIaCFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>Delete ID list</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteIaCFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIaCFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCFileResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIaCFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIaCFileResponseParams `json:"Response"`
}

func (r *DeleteIaCFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskRequestParams struct {
	// task id and target AppID list
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DeleteRiskScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// task id and target AppID list
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DeleteRiskScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRiskScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRiskScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRiskScanTaskResponseParams `json:"Response"`
}

func (r *DeleteRiskScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAgentAssetListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// filter
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAIAgentAssetListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// filter
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAIAgentAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAgentAssetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAgentAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAgentAssetListResponseParams struct {
	// asset list
	AssetList []*AIAgentAsset `json:"AssetList,omitnil,omitempty" name:"AssetList"`

	// Total number of assets
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIAgentAssetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIAgentAssetListResponseParams `json:"Response"`
}

func (r *DescribeAIAgentAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAgentAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAKAnalysisDetailRequestParams struct {
	// Alarm record ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAKAnalysisDetailRequest struct {
	*tchttp.BaseRequest
	
	// Alarm record ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAKAnalysisDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAKAnalysisDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAKAnalysisDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAKAnalysisDetailResponseParams struct {
	// Alarm AI analysis status -1 Analysis failed 0 Not analyzed 1 Under analysis 2 Analysis successful, real alarm 3 Analysis successful, suspicious alarm
	AIStatus *int64 `json:"AIStatus,omitnil,omitempty" name:"AIStatus"`

	// AI Analysis Task ID
	AITaskID *string `json:"AITaskID,omitnil,omitempty" name:"AITaskID"`

	// Alarm AI analysis result, base64 format, avoid data interception
	AIResult *string `json:"AIResult,omitnil,omitempty" name:"AIResult"`

	// Feedback and Suggestions
	Feedback *string `json:"Feedback,omitnil,omitempty" name:"Feedback"`

	// Feedback status  0 means no feedback, 1 means recognized, 2 means not recognized
	FeedbackResult *int64 `json:"FeedbackResult,omitnil,omitempty" name:"FeedbackResult"`

	// Reason for failure
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAKAnalysisDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAKAnalysisDetailResponseParams `json:"Response"`
}

func (r *DescribeAKAnalysisDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAKAnalysisDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalCallRecordRequestParams struct {
	// Alarm rule ID.
	AlarmRuleID *int64 `json:"AlarmRuleID,omitnil,omitempty" name:"AlarmRuleID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Source IP of the call.
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAbnormalCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// Alarm rule ID.
	AlarmRuleID *int64 `json:"AlarmRuleID,omitnil,omitempty" name:"AlarmRuleID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Source IP of the call.
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAbnormalCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmRuleID")
	delete(f, "MemberId")
	delete(f, "AccessKey")
	delete(f, "SourceIP")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalCallRecordResponseParams struct {
	// Invocation record list.
	Data []*CallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of records.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAbnormalCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalCallRecordResponseParams `json:"Response"`
}

func (r *DescribeAbnormalCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalCallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmDetailRequestParams struct {
	// Alarm record ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAccessKeyAlarmDetailRequest struct {
	*tchttp.BaseRequest
	
	// Alarm record ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAccessKeyAlarmDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyAlarmDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmDetailResponseParams struct {
	// Alarm information.
	AlarmInfo *AccessKeyAlarm `json:"AlarmInfo,omitnil,omitempty" name:"AlarmInfo"`

	// Number of CAM policies in the associated account.
	CamCount *int64 `json:"CamCount,omitnil,omitempty" name:"CamCount"`

	// Number of AK risks.
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// Alarm policy description.
	AlarmDesc *string `json:"AlarmDesc,omitnil,omitempty" name:"AlarmDesc"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyAlarmDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyAlarmDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyAlarmDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmRequestParams struct {
	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// ID of the source IP.
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// Account UIN
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type DescribeAccessKeyAlarmRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// ID of the source IP.
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// Account UIN
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

func (r *DescribeAccessKeyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	delete(f, "AccessKeyID")
	delete(f, "SourceIPID")
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmResponseParams struct {
	// Alarm list.
	Data []*AccessKeyAlarm `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyAlarmResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAssetRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAccessKeyAssetRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAccessKeyAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAssetResponseParams struct {
	// Access key asset list.
	Data []*AccessKeyAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// Total quantity.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyAssetResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskDetailRequestParams struct {
	// Risk record ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAccessKeyRiskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Risk record ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAccessKeyRiskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyRiskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskDetailResponseParams struct {
	// Risk list.
	RiskInfo *AccessKeyRisk `json:"RiskInfo,omitnil,omitempty" name:"RiskInfo"`

	// Total number of CAM policies.
	CamCount *int64 `json:"CamCount,omitnil,omitempty" name:"CamCount"`

	// Number of associated alarms for the account.
	AlarmCount *int64 `json:"AlarmCount,omitnil,omitempty" name:"AlarmCount"`

	// Access method 0 API 1 console and API.
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// Access key Alarm count list.
	AccessKeyAlarmCount []*AccessKeyAlarmCount `json:"AccessKeyAlarmCount,omitnil,omitempty" name:"AccessKeyAlarmCount"`

	// Whether operation protection is enabled. valid values: 0 (not enabled), 1 (enabled).
	ActionFlag *int64 `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`

	// Whether login protection is enabled. valid values: 0 (not enabled), 1 (enabled).
	LoginFlag *int64 `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyRiskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyRiskDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyRiskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskRequestParams struct {
	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// Account UIN
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type DescribeAccessKeyRiskRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// Account UIN
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

func (r *DescribeAccessKeyRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	delete(f, "AccessKeyID")
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskResponseParams struct {
	// Risk list.
	Data []*AccessKeyRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyRiskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyRiskResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserDetailRequestParams struct {
	// Account uin itself.
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAccessKeyUserDetailRequest struct {
	*tchttp.BaseRequest
	
	// Account uin itself.
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAccessKeyUserDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubUin")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyUserDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserDetailResponseParams struct {
	// Account detailed information.
	User *AccessKeyUser `json:"User,omitnil,omitempty" name:"User"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyUserDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyUserDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyUserDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAccessKeyUserListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAccessKeyUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserListResponseParams struct {
	// Account list.
	Data []*AccessKeyUser `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyUserListResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertListRequestParams struct {
	// Tag search filter criteria
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Member ID of the Called Group Account
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// 0: Default all 1: Asset ID 2: Domain name
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

type DescribeAlertListRequest struct {
	*tchttp.BaseRequest
	
	// Tag search filter criteria
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Member ID of the Called Group Account
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// 0: Default all 1: Asset ID 2: Domain name
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

func (r *DescribeAlertListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	delete(f, "OperatedMemberId")
	delete(f, "AssetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlertListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertListResponseParams struct {
	// All alarms list
	AlertList []*AlertInfo `json:"AlertList,omitnil,omitempty" name:"AlertList"`

	// Number of Major Categories of Alarm
	AlertTypeCount []*TagCount `json:"AlertTypeCount,omitnil,omitempty" name:"AlertTypeCount"`

	// Total number of alarms
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 0:succeed 1:timeout
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Return status
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlertListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlertListResponseParams `json:"Response"`
}

func (r *DescribeAlertListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessListRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Cloud service provider.
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

type DescribeAssetProcessListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Cloud service provider.
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

func (r *DescribeAssetProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "Provider")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessListResponseParams struct {
	// Process quantity.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Process list.
	AssetProcessList []*AssetProcessItem `json:"AssetProcessList,omitnil,omitempty" name:"AssetProcessList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetProcessListResponseParams `json:"Response"`
}

func (r *DescribeAssetProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetRiskListRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeAssetRiskListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeAssetRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetRiskListResponseParams struct {
	// Number of risks from asset perspective
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Risk list from asset perspective
	AssetRiskList []*AssetRiskItem `json:"AssetRiskList,omitnil,omitempty" name:"AssetRiskList"`

	// Standard name collection
	StandardNameList []*StandardItem `json:"StandardNameList,omitnil,omitempty" name:"StandardNameList"`

	// Asset type collection
	AssetTypeList []*AttributeOptionSet `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetRiskListResponseParams `json:"Response"`
}

func (r *DescribeAssetRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetViewVulRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tag
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeAssetViewVulRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tag
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeAssetViewVulRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetViewVulRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetViewVulRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetViewVulRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Vulnerability Risk List from Asset's Perspective
	Data []*AssetViewVULRiskData `json:"Data,omitnil,omitempty" name:"Data"`

	// Status list
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// Danger Level List
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// Source List
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// Vulnerability Type List
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// Asset Type List
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// tag enumeration.
	Tags []*FilterDataObject `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetViewVulRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetViewVulRiskListResponseParams `json:"Response"`
}

func (r *DescribeAssetViewVulRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetViewVulRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCFWAssetStatisticsRequestParams struct {

}

type DescribeCFWAssetStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCFWAssetStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCFWAssetStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCFWAssetStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCFWAssetStatisticsResponseParams struct {
	// Total number of network assets
	NetworkTotal *int64 `json:"NetworkTotal,omitnil,omitempty" name:"NetworkTotal"`

	// Asset CLB Quantity
	ClbTotal *int64 `json:"ClbTotal,omitnil,omitempty" name:"ClbTotal"`

	// Number of NATs
	NatTotal *int64 `json:"NatTotal,omitnil,omitempty" name:"NatTotal"`

	// Number of Public IP Addresses
	PublicAssetTotal *int64 `json:"PublicAssetTotal,omitnil,omitempty" name:"PublicAssetTotal"`

	// Number of hosts
	CVMAssetTotal *int64 `json:"CVMAssetTotal,omitnil,omitempty" name:"CVMAssetTotal"`

	// Configuration risk.
	CFGTotal *int64 `json:"CFGTotal,omitnil,omitempty" name:"CFGTotal"`

	// Port risk.
	PortTotal *int64 `json:"PortTotal,omitnil,omitempty" name:"PortTotal"`

	// Content risk.
	WebsiteTotal *int64 `json:"WebsiteTotal,omitnil,omitempty" name:"WebsiteTotal"`

	// Risk service exposure.
	ServerTotal *int64 `json:"ServerTotal,omitnil,omitempty" name:"ServerTotal"`

	// Weak password risk.
	WeakPasswordTotal *int64 `json:"WeakPasswordTotal,omitnil,omitempty" name:"WeakPasswordTotal"`

	// Vulnerability risk.
	VULTotal *int64 `json:"VULTotal,omitnil,omitempty" name:"VULTotal"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCFWAssetStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCFWAssetStatisticsResponseParams `json:"Response"`
}

func (r *DescribeCFWAssetStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCFWAssetStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSIPRiskStatisticsRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCSIPRiskStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCSIPRiskStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSIPRiskStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCSIPRiskStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSIPRiskStatisticsResponseParams struct {
	// Asset Overview Data
	Data *CsipRiskCenterStatistics `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCSIPRiskStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCSIPRiskStatisticsResponseParams `json:"Response"`
}

func (r *DescribeCSIPRiskStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSIPRiskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetInfoRequestParams struct {
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeCVMAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeCVMAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCVMAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetInfoResponseParams struct {
	// Data.
	Data *AssetBaseInfoResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCVMAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCVMAssetInfoResponseParams `json:"Response"`
}

func (r *DescribeCVMAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCVMAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCVMAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCVMAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetsResponseParams struct {
	// Total number.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Machine list
	Data []*CVMAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Region list
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Protection status
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// VPC Enumeration
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// Asset Type Enumeration
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// Operating System Enumeration
	SystemTypeList []*FilterDataObject `json:"SystemTypeList,omitnil,omitempty" name:"SystemTypeList"`

	// IP List
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil,omitempty" name:"IpTypeList"`

	// AppID List
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// Availability Zone List
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// OS List
	OsList []*FilterDataObject `json:"OsList,omitnil,omitempty" name:"OsList"`

	// Mapping of asset type and instance type.
	AssetMapInstanceTypeList []*AssetInstanceTypeMap `json:"AssetMapInstanceTypeList,omitnil,omitempty" name:"AssetMapInstanceTypeList"`

	// Public network private network enumeration.
	PublicPrivateAttr []*FilterDataObject `json:"PublicPrivateAttr,omitnil,omitempty" name:"PublicPrivateAttr"`

	// Host protection status.
	ProtectStatusList []*FilterDataObject `json:"ProtectStatusList,omitnil,omitempty" name:"ProtectStatusList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCVMAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCVMAssetsResponseParams `json:"Response"`
}

func (r *DescribeCVMAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallRecordRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// ID of the source IP for the call.
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// Access account uin.
	AccUin *string `json:"AccUin,omitnil,omitempty" name:"AccUin"`

	// Access key. Note: Temporary key is unsupported.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Access key ID.
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// ID of the source IP for the call.
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// Access account uin.
	AccUin *string `json:"AccUin,omitnil,omitempty" name:"AccUin"`

	// Access key. Note: Temporary key is unsupported.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "AccessKeyID")
	delete(f, "SourceIPID")
	delete(f, "AccUin")
	delete(f, "AccessKey")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallRecordResponseParams struct {
	// Invocation record list.
	Data []*CallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of records.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallRecordResponseParams `json:"Response"`
}

func (r *DescribeCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckViewRisksRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>Filter content</p>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Page size.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting type</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>Sorting field.</p>
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeCheckViewRisksRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>Filter content</p>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Page size.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting type</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>Sorting field.</p>
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeCheckViewRisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckViewRisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCheckViewRisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckViewRisksResponseParams struct {
	// <p>Number of risks from check perspective</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Risk list in check perspective</p>
	CheckViewRiskList []*CheckViewRiskItem `json:"CheckViewRiskList,omitnil,omitempty" name:"CheckViewRiskList"`

	// <p>Tag list of cspm standard from a checking perspective</p>
	StandardNameList []*StandardItem `json:"StandardNameList,omitnil,omitempty" name:"StandardNameList"`

	// <p>Asset type collection</p>
	AssetTypeList []*AttributeOptionSet `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// <p>Cloud vendor type collection</p>
	ProviderList []*AttributeOptionSet `json:"ProviderList,omitnil,omitempty" name:"ProviderList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCheckViewRisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCheckViewRisksResponseParams `json:"Response"`
}

func (r *DescribeCheckViewRisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckViewRisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeClusterAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeClusterAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAssetsResponseParams struct {
	// List
	Data []*AssetCluster `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Cluster Type Enumeration
	ClusterTypeList []*FilterDataObject `json:"ClusterTypeList,omitnil,omitempty" name:"ClusterTypeList"`

	// Cluster Status Enumeration
	ClusterStatusList []*FilterDataObject `json:"ClusterStatusList,omitnil,omitempty" name:"ClusterStatusList"`

	// Component Status Enumeration
	ComponentStatusList []*FilterDataObject `json:"ComponentStatusList,omitnil,omitempty" name:"ComponentStatusList"`

	// VPC Enumeration
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// Region Enumeration
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Tenant Enumeration
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// Cluster protection status enumeration.
	ProtectStatusList []*FilterDataObject `json:"ProtectStatusList,omitnil,omitempty" name:"ProtectStatusList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAssetsResponseParams `json:"Response"`
}

func (r *DescribeClusterAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPodAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeClusterPodAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeClusterPodAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPodAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPodAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPodAssetsResponseParams struct {
	// Data list
	Data []*AssetClusterPod `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of cluster pod status
	PodStatusList []*FilterDataObject `json:"PodStatusList,omitnil,omitempty" name:"PodStatusList"`

	// List of namespaces
	NamespaceList []*FilterDataObject `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`

	// List of regions
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of users (AppId)
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterPodAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPodAssetsResponseParams `json:"Response"`
}

func (r *DescribeClusterPodAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPodAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigCheckRulesRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeConfigCheckRulesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeConfigCheckRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigCheckRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigCheckRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigCheckRulesResponseParams struct {
	// Number of risk rules
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Risk rule list
	RuleList []*RiskRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Cloud vendor type options
	ProviderList []*AttributeOptionSet `json:"ProviderList,omitnil,omitempty" name:"ProviderList"`

	// Risk level type option
	RiskLevelList []*AttributeOptionSet `json:"RiskLevelList,omitnil,omitempty" name:"RiskLevelList"`

	// Disposal categorization options
	DispositionTypeList []*AttributeOptionSet `json:"DispositionTypeList,omitnil,omitempty" name:"DispositionTypeList"`

	// Check type options
	CheckTypeList []*AttributeOptionSet `json:"CheckTypeList,omitnil,omitempty" name:"CheckTypeList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigCheckRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigCheckRulesResponseParams `json:"Response"`
}

func (r *DescribeConfigCheckRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigCheckRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetInfoRequestParams struct {
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeDbAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeDbAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetInfoResponseParams struct {
	// DB Asset Details
	Data *DbAssetInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbAssetInfoResponseParams `json:"Response"`
}

func (r *DescribeDbAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset types. Values: MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

type DescribeDbAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset types. Values: MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

func (r *DescribeDbAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "AssetTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetsResponseParams struct {
	// Total number.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Total number of assets
	Data []*DBAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Region Enumeration
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Asset Type Enumeration
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// VPC Enumeration
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// Appid Enumeration
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// Public network private network enumeration
	PublicPrivateAttr []*FilterDataObject `json:"PublicPrivateAttr,omitnil,omitempty" name:"PublicPrivateAttr"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbAssetsResponseParams `json:"Response"`
}

func (r *DescribeDbAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeDomainAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeDomainAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAssetsResponseParams struct {
	// Total number.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Domain name list
	Data []*DomainAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Protection Status List
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// Asset Attribution List
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil,omitempty" name:"AssetLocationList"`

	// Asset Type List
	SourceTypeList []*FilterDataObject `json:"SourceTypeList,omitnil,omitempty" name:"SourceTypeList"`

	// Region list
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainAssetsResponseParams `json:"Response"`
}

func (r *DescribeDomainAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposeAssetCategoryRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeExposeAssetCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeExposeAssetCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposeAssetCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExposeAssetCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposeAssetCategoryResponseParams struct {
	// Cloud boundary analytics asset classification list.
	ExposeAssetTypeList []*ExposeAssetTypeItem `json:"ExposeAssetTypeList,omitnil,omitempty" name:"ExposeAssetTypeList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExposeAssetCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExposeAssetCategoryResponseParams `json:"Response"`
}

func (r *DescribeExposeAssetCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposeAssetCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposePathRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Asset ID.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset IP.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Asset domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Port or port range.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type DescribeExposePathRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Asset ID.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset IP.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Asset domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Port or port range.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

func (r *DescribeExposePathRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposePathRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "AssetId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExposePathRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposePathResponseParams struct {
	// Cloud boundary analysis path within node.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExposePathResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExposePathResponseParams `json:"Response"`
}

func (r *DescribeExposePathResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposuresRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeExposuresRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeExposuresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposuresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExposuresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposuresResponseParams struct {
	// Cloud boundary analytics number of assets.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Cloud boundary analytics asset list.
	ExposeList []*ExposesItem `json:"ExposeList,omitnil,omitempty" name:"ExposeList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExposuresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExposuresResponseParams `json:"Response"`
}

func (r *DescribeExposuresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposuresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeGatewayAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeGatewayAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayAssetsResponseParams struct {
	// List
	Data []*GateWayAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Region list
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Asset Type List
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// VPC List
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// AppID List
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewayAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayAssetsResponseParams `json:"Response"`
}

func (r *DescribeGatewayAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHighBaseLineRiskListRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Cloud account ID.
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// Cloud service provider.
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

type DescribeHighBaseLineRiskListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Cloud account ID.
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// Cloud service provider.
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

func (r *DescribeHighBaseLineRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHighBaseLineRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "CloudAccountID")
	delete(f, "Provider")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHighBaseLineRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHighBaseLineRiskListResponseParams struct {
	// Number of high-risk risks.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// High-Risk baseline risk list.
	HighBaseLineRiskList []*HighBaseLineRiskItem `json:"HighBaseLineRiskList,omitnil,omitempty" name:"HighBaseLineRiskList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHighBaseLineRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHighBaseLineRiskListResponseParams `json:"Response"`
}

func (r *DescribeHighBaseLineRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHighBaseLineRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileListRequestParams struct {
	// <p>Filter conditions.</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeIaCFileListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Filter conditions.</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeIaCFileListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCFileListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileListResponseParams struct {
	// <p>List</p>
	List []*IaCFile `json:"List,omitnil,omitempty" name:"List"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCFileListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCFileListResponseParams `json:"Response"`
}

func (r *DescribeIaCFileListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileOverviewRequestParams struct {
	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeIaCFileOverviewRequest struct {
	*tchttp.BaseRequest
	
	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeIaCFileOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCFileOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileOverviewResponseParams struct {
	// <p>Number of files.</p>
	TotalFile *uint64 `json:"TotalFile,omitnil,omitempty" name:"TotalFile"`

	// <p>Number of risk files (1: Dockerfile, 2: Terraform, 3: KubernetesYaml)</p>
	RiskFile []*KeyValueInt `json:"RiskFile,omitnil,omitempty" name:"RiskFile"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCFileOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCFileOverviewResponseParams `json:"Response"`
}

func (r *DescribeIaCFileOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileReportRequestParams struct {
	// <p>Asset ID</p>
	AssetId *uint64 `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeIaCFileReportRequest struct {
	*tchttp.BaseRequest
	
	// <p>Asset ID</p>
	AssetId *uint64 `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeIaCFileReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCFileReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileReportResponseParams struct {
	// <p>Detect file</p>
	File *string `json:"File,omitnil,omitempty" name:"File"`

	// <p>Detection status (0: pending scan, 1: detecting, 2: completed, 3: detection exception)</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Detection time</p>
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// <p>Risk list</p>
	Risks []*IaCFileRisk `json:"Risks,omitnil,omitempty" name:"Risks"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCFileReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCFileReportResponseParams `json:"Response"`
}

func (r *DescribeIaCFileReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCTokenListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>Filter conditions.</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeIaCTokenListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>Filter conditions.</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeIaCTokenListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCTokenListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCTokenListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCTokenListResponseParams struct {
	// <p>List</p>
	List []*CICDToken `json:"List,omitnil,omitempty" name:"List"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCTokenListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCTokenListResponseParams `json:"Response"`
}

func (r *DescribeIaCTokenListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCTokenListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialListRequestParams struct {
	// Filter criteria list: supported filter conditions as follows:
	// CredentialName - Credential name (fuzzy matching)
	// CredentialType - Credential type (exact match). Parameter values: access, sts.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeKeySandboxCredentialListRequest struct {
	*tchttp.BaseRequest
	
	// Filter criteria list: supported filter conditions as follows:
	// CredentialName - Credential name (fuzzy matching)
	// CredentialType - Credential type (exact match). Parameter values: access, sts.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeKeySandboxCredentialListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeySandboxCredentialListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialListResponseParams struct {
	// Credential data list
	Data []*KeySandboxCredential `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeySandboxCredentialListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeySandboxCredentialListResponseParams `json:"Response"`
}

func (r *DescribeKeySandboxCredentialListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialRequestParams struct {
	// Credential ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeKeySandboxCredentialRequest struct {
	*tchttp.BaseRequest
	
	// Credential ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeKeySandboxCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeySandboxCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialResponseParams struct {
	// Credential ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// Credential name
	CredentialName *string `json:"CredentialName,omitnil,omitempty" name:"CredentialName"`

	// Credential Type
	// Enumeration value:
	// access: standard key
	// sts: STS temporary key
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// Effective machine scope
	CredentialEffectScope *CredentialEffectScope `json:"CredentialEffectScope,omitnil,omitempty" name:"CredentialEffectScope"`

	// Normal key credential data (masked). Returned when CredentialType is access.
	// Supplementary explanation: Key is the original text, and Value is the masked value (reserve the first 3 and last 4 digits, with *** as substitution in the middle).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Access []*AccessCredentialOutput `json:"Access,omitnil,omitempty" name:"Access"`

	// STS credential data (masked). Returned when CredentialType is sts.
	// Supplementary description: System is the original text, SecretID and SecretKey are masked values (reserve the first 3 and last 4 characters, with *** as substitution in the middle).
	// Note: This field may return null, indicating that no valid values can be obtained.
	STS *STSCredentialOutput `json:"STS,omitnil,omitempty" name:"STS"`

	// Creation time.
	// Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	// Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeySandboxCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeySandboxCredentialResponseParams `json:"Response"`
}

func (r *DescribeKeySandboxCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeListenerListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeListenerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerListResponseParams struct {
	// Total number.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Listener list
	Data []*ClbListenerListInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenerListResponseParams `json:"Response"`
}

func (r *DescribeListenerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNICAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeNICAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeNICAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNICAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNICAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNICAssetsResponseParams struct {
	// List
	Data []*NICAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Region list
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Asset Type List
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// VPC List
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// AppID List
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNICAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNICAssetsResponseParams `json:"Response"`
}

func (r *DescribeNICAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNICAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationInfoRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeOrganizationInfoRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeOrganizationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationInfoResponseParams struct {
	// Total quantity.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Group User List
	Data []*OrganizationInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationInfoResponseParams `json:"Response"`
}

func (r *DescribeOrganizationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationUserInfoRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Does not support multi-cloud.
	NotSupportCloud *bool `json:"NotSupportCloud,omitnil,omitempty" name:"NotSupportCloud"`
}

type DescribeOrganizationUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Does not support multi-cloud.
	NotSupportCloud *bool `json:"NotSupportCloud,omitnil,omitempty" name:"NotSupportCloud"`
}

func (r *DescribeOrganizationUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "NotSupportCloud")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationUserInfoResponseParams struct {
	// Total quantity.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Group User List
	Data []*OrganizationUserInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Join method enumeration
	JoinTypeLst []*FilterDataObject `json:"JoinTypeLst,omitnil,omitempty" name:"JoinTypeLst"`

	// Cloud vendor enumeration
	CloudTypeLst []*FilterDataObject `json:"CloudTypeLst,omitnil,omitempty" name:"CloudTypeLst"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationUserInfoResponseParams `json:"Response"`
}

func (r *DescribeOrganizationUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtherCloudAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// -
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset type: MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

type DescribeOtherCloudAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// -
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset type: MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

func (r *DescribeOtherCloudAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtherCloudAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "AssetTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOtherCloudAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtherCloudAssetsResponseParams struct {
	// Total number.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Total number of assets
	Data []*DBAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Region Enumeration
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Asset Type Enumeration
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// VPC Enumeration
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// Appid Enumeration
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOtherCloudAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOtherCloudAssetsResponseParams `json:"Response"`
}

func (r *DescribeOtherCloudAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtherCloudAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribePublicIpAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribePublicIpAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicIpAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpAssetsResponseParams struct {
	// List
	Data []*IpAssetListVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Asset Attribution
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil,omitempty" name:"AssetLocationList"`

	// IP List Enumeration
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil,omitempty" name:"IpTypeList"`

	// Region List Enumeration
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Protection Enumeration
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// Asset Type Enumeration
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// AppId Enumeration
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicIpAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicIpAssetsResponseParams `json:"Response"`
}

func (r *DescribePublicIpAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryImageAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRepositoryImageAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRepositoryImageAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryImageAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoryImageAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryImageAssetsResponseParams struct {
	// Repository Image List
	Data []*RepositoryImageVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Region List
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRepositoryImageAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoryImageAssetsResponseParams `json:"Response"`
}

func (r *DescribeRepositoryImageAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryImageAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCallRecordRequestParams struct {
	// Risk record ID.
	RiskID *int64 `json:"RiskID,omitnil,omitempty" name:"RiskID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRiskCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// Risk record ID.
	RiskID *int64 `json:"RiskID,omitnil,omitempty" name:"RiskID"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRiskCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskID")
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCallRecordResponseParams struct {
	// Risk call record list.
	Data []*RiskCallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of records.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCallRecordResponseParams `json:"Response"`
}

func (r *DescribeRiskCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewCFGRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewCFGRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewCFGRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewCFGRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of configuration risks
	Data []*AssetViewCFGRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of configuration names
	CFGNameLists []*FilterDataObject `json:"CFGNameLists,omitnil,omitempty" name:"CFGNameLists"`

	// List of check types
	CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitnil,omitempty" name:"CheckTypeLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewCFGRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewCFGRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewPortRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewPortRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewPortRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewPortRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of configuration risks
	Data []*AssetViewPortRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of fix suggestions 
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitnil,omitempty" name:"SuggestionLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewPortRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewPortRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewVULRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewVULRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewVULRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewVULRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of vulnerabilities
	Data []*AssetViewVULRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// List of vulnerability types
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewVULRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewVULRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewWeakPasswordRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewWeakPasswordRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewWeakPasswordRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of risks
	Data []*AssetViewWeakPassRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of weak password types
	PasswordTypeLists []*FilterDataObject `json:"PasswordTypeLists,omitnil,omitempty" name:"PasswordTypeLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterCFGViewCFGRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRiskCenterCFGViewCFGRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterCFGViewCFGRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterCFGViewCFGRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterCFGViewCFGRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterCFGViewCFGRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Configuration Risk List from Asset's Perspective
	Data []*CFGViewCFGRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// Status list
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// Danger Level List
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// Configuration Name List
	CFGNameLists []*FilterDataObject `json:"CFGNameLists,omitnil,omitempty" name:"CFGNameLists"`

	// Check Type List
	CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitnil,omitempty" name:"CheckTypeLists"`

	// Asset Type List
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// Source List
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterCFGViewCFGRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterCFGViewCFGRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterCFGViewCFGRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterCFGViewCFGRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterPortViewPortRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterPortViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterPortViewPortRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterPortViewPortRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterPortViewPortRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterPortViewPortRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Port Risk List from Port's Perspective
	Data []*PortViewPortRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of suggestions
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitnil,omitempty" name:"SuggestionLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterPortViewPortRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterPortViewPortRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterPortViewPortRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterPortViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterServerRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterServerRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterServerRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterServerRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterServerRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterServerRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of services in risk
	Data []*ServerRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterServerRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterServerRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterServerRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterServerRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterVULViewVULRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterVULViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterVULViewVULRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterVULViewVULRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterVULViewVULRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterVULViewVULRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of vulnerabilities
	Data []*VULViewVULRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// List of vulnerability types
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterVULViewVULRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterVULViewVULRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterVULViewVULRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterVULViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterWebsiteRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterWebsiteRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterWebsiteRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterWebsiteRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterWebsiteRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterWebsiteRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Content Risk List
	Data []*WebsiteRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of risk types
	DetectEngineLists []*FilterDataObject `json:"DetectEngineLists,omitnil,omitempty" name:"DetectEngineLists"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterWebsiteRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterWebsiteRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterWebsiteRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterWebsiteRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDetailListRequestParams struct {
	// Risk rule ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRiskDetailListRequest struct {
	*tchttp.BaseRequest
	
	// Risk rule ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeRiskDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskRuleId")
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDetailListResponseParams struct {
	// Risk detail count from asset perspective
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Risk detail list from asset perspective
	AssetRiskDetailList []*RiskDetailItem `json:"AssetRiskDetailList,omitnil,omitempty" name:"AssetRiskDetailList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskDetailListResponseParams `json:"Response"`
}

func (r *DescribeRiskDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRuleDetailRequestParams struct {
	// <p>Risk rule ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`
}

type DescribeRiskRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>Risk rule ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`
}

func (r *DescribeRiskRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRuleDetailResponseParams struct {
	// <p>Risk rule ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>Cloud vendor</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Risk name</p>
	RiskName *string `json:"RiskName,omitnil,omitempty" name:"RiskName"`

	// <p>Risk damage</p>
	RiskInfluence *string `json:"RiskInfluence,omitnil,omitempty" name:"RiskInfluence"`

	// <p>Repair guidance</p>
	RiskFixAdvice *string `json:"RiskFixAdvice,omitnil,omitempty" name:"RiskFixAdvice"`

	// <p>Asset type</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeRiskRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRulesRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeRiskRulesRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeRiskRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRulesResponseParams struct {
	// Number of risk rules
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Risk rule list
	RiskRuleList []*RiskRuleItem `json:"RiskRuleList,omitnil,omitempty" name:"RiskRuleList"`

	// Instance type options
	InstanceTypeList []*AttributeOptionSet `json:"InstanceTypeList,omitnil,omitempty" name:"InstanceTypeList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskRulesResponseParams `json:"Response"`
}

func (r *DescribeRiskRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanReportListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeScanReportListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeScanReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanReportListResponseParams struct {
	// Total quantity.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Task Log List
	Data []*ScanTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// List of account UINs
	UINList []*string `json:"UINList,omitnil,omitempty" name:"UINList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanReportListResponseParams `json:"Response"`
}

func (r *DescribeScanReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStatisticRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Health check task id.
	TaskLogId *string `json:"TaskLogId,omitnil,omitempty" name:"TaskLogId"`
}

type DescribeScanStatisticRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Health check task id.
	TaskLogId *string `json:"TaskLogId,omitnil,omitempty" name:"TaskLogId"`
}

func (r *DescribeScanStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskLogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStatisticResponseParams struct {
	// Port service quantity.
	PortServiceCount *uint64 `json:"PortServiceCount,omitnil,omitempty" name:"PortServiceCount"`

	// Number of Web services.
	WebAppCount *uint64 `json:"WebAppCount,omitnil,omitempty" name:"WebAppCount"`

	// Weak Password Risk Count
	WeakPasswordCount *uint64 `json:"WeakPasswordCount,omitnil,omitempty" name:"WeakPasswordCount"`

	// Vulnerability risk quantity.
	VulCount *uint64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// High-Risk port service quantity.
	HighRiskPortServiceCount *uint64 `json:"HighRiskPortServiceCount,omitnil,omitempty" name:"HighRiskPortServiceCount"`

	// Number of Web services at risk.
	RiskWebAppCount *uint64 `json:"RiskWebAppCount,omitnil,omitempty" name:"RiskWebAppCount"`

	// Newly-Added port services in the last 7 days.
	PortServiceIncrement *uint64 `json:"PortServiceIncrement,omitnil,omitempty" name:"PortServiceIncrement"`

	// Newly-Added Web services in the last 7 days.
	WebAppIncrement *uint64 `json:"WebAppIncrement,omitnil,omitempty" name:"WebAppIncrement"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanStatisticResponseParams `json:"Response"`
}

func (r *DescribeScanStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Tags
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeScanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Tags
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeScanTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListResponseParams struct {
	// Total quantity.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Task Log List
	Data []*ScanTaskInfoList `json:"Data,omitnil,omitempty" name:"Data"`

	// Host Account ID List
	UINList []*string `json:"UINList,omitnil,omitempty" name:"UINList"`

	// Health Checkup Mode Filter List
	TaskModeList []*FilterDataObject `json:"TaskModeList,omitnil,omitempty" name:"TaskModeList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskListResponseParams `json:"Response"`
}

func (r *DescribeScanTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchBugInfoRequestParams struct {
	// 1 when return emergency vulnerability, 2 when return emergency vulnerability list, 3 when collocation input CVEId field display vulnerability data
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// CVE number of the vulnerability. It's required when `Id=3`.
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`
}

type DescribeSearchBugInfoRequest struct {
	*tchttp.BaseRequest
	
	// 1 when return emergency vulnerability, 2 when return emergency vulnerability list, 3 when collocation input CVEId field display vulnerability data
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// CVE number of the vulnerability. It's required when `Id=3`.
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`
}

func (r *DescribeSearchBugInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchBugInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "CVEId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchBugInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchBugInfoResponseParams struct {
	// Vulnerability information and asset information
	Data *DataSearchBug `json:"Data,omitnil,omitempty" name:"Data"`

	// Status code. Valid values: 0: successful; others: failed.
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Status message. Valid values: success: successful query; fail: failed query.
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSearchBugInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchBugInfoResponseParams `json:"Response"`
}

func (r *DescribeSearchBugInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchBugInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillScanPayInfoRequestParams struct {

}

type DescribeSkillScanPayInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSkillScanPayInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanPayInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillScanPayInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillScanPayInfoResponseParams struct {
	// <p>AppID of the associated tenant for the order</p>
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>Order status<br>Enumeration value:<br>0: Not purchased<br>1: Normal<br>2: Isolated<br>6: In trial<br>7: Expired<br>8: Trial expiration</p>
	OrderStatus *int64 `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// <p>Total quota</p>
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// <p>Consumed quota.</p>
	UsedCount *int64 `json:"UsedCount,omitnil,omitempty" name:"UsedCount"`

	// <p>Payment mode<br>Enumeration value:<br>0: Postpaid<br>1: Prepaid</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Auto-renewal flag<br>Enumeration value:<br>0: not set<br>1: auto-renewal<br>2: no auto-renewal</p>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>Resource ID</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>Purchase period</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>Duration unit</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Order start time</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>Order expiration time</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Open beta end time is fixed as 2026-06-30 23:59:59</p>
	BetaEndTime *string `json:"BetaEndTime,omitnil,omitempty" name:"BetaEndTime"`

	// <p>Server current time</p>
	TimeNow *string `json:"TimeNow,omitnil,omitempty" name:"TimeNow"`

	// <p>Tenant Uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>Tenant nickname</p>
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillScanPayInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillScanPayInfoResponseParams `json:"Response"`
}

func (r *DescribeSkillScanPayInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanPayInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillScanResultRequestParams struct {
	// SHA256 Hash of the ZIP file
	// Parameter format: sha256:<64-bit hex>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// Specify the engine version number
	// Value for reference: API response of CreateSkillScan
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Report signature address validity
	// Unit: hr
	// Default value: 8760 (1 year).
	// Supplementary explanation: The returned ReportURL takes effect.
	ReportURLExpireHours *int64 `json:"ReportURLExpireHours,omitnil,omitempty" name:"ReportURLExpireHours"`
}

type DescribeSkillScanResultRequest struct {
	*tchttp.BaseRequest
	
	// SHA256 Hash of the ZIP file
	// Parameter format: sha256:<64-bit hex>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// Specify the engine version number
	// Value for reference: API response of CreateSkillScan
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Report signature address validity
	// Unit: hr
	// Default value: 8760 (1 year).
	// Supplementary explanation: The returned ReportURL takes effect.
	ReportURLExpireHours *int64 `json:"ReportURLExpireHours,omitnil,omitempty" name:"ReportURLExpireHours"`
}

func (r *DescribeSkillScanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContentHash")
	delete(f, "EngineVersion")
	delete(f, "ReportURLExpireHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillScanResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillScanResultResponseParams struct {
	// Detection status
	// Enumeration value:
	// SUCCESS: Detection completed, results returned.
	// SCANNING: Detection in progress
	// NOT_FOUND: No detection record.
	// FAILED: Detection failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Detection result details. When Status=SUCCESS, most fields have values. When Status=SCANNING, it contains only ContentHash and CreatedAt. When Status=FAILED, it contains only ContentHash, FailedAt, and Message. When Status=NOT_FOUND, it contains only ContentHash.
	Data *SkillScanItem `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillScanResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillScanResultResponseParams `json:"Response"`
}

func (r *DescribeSkillScanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceIPAssetRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSourceIPAssetRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSourceIPAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIPAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceIPAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceIPAssetResponseParams struct {
	// Access key asset list.
	Data []*SourceIPAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// Total quantity.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSourceIPAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceIPAssetResponseParams `json:"Response"`
}

func (r *DescribeSourceIPAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIPAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubUserInfoRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSubUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSubUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubUserInfoResponseParams struct {
	// Total number.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Sub-user list
	Data []*SubUserInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Manufacturer list
	CloudTypeLst []*FilterDataObject `json:"CloudTypeLst,omitnil,omitempty" name:"CloudTypeLst"`

	// Enumerate appid belonging to main account
	OwnerAppIDLst []*FilterDataObject `json:"OwnerAppIDLst,omitnil,omitempty" name:"OwnerAppIDLst"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubUserInfoResponseParams `json:"Response"`
}

func (r *DescribeSubUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSubnetAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSubnetAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAssetsResponseParams struct {
	// Data list
	Data []*SubnetAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of regions
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of VPCs
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// List of AppIds
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// List of availability zones
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubnetAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetAssetsResponseParams `json:"Response"`
}

func (r *DescribeSubnetAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeTaskLogListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeTaskLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogListResponseParams struct {
	// Total quantity.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Report List
	Data []*TaskLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Pending Viewing Count
	NotViewNumber *int64 `json:"NotViewNumber,omitnil,omitempty" name:"NotViewNumber"`

	// Number of Report Templates
	ReportTemplateNumber *int64 `json:"ReportTemplateNumber,omitnil,omitempty" name:"ReportTemplateNumber"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogListResponseParams `json:"Response"`
}

func (r *DescribeTaskLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogURLRequestParams struct {
	// Type of the task. `0`: Preview; `1`: Download
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// List of task report IDs
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil,omitempty" name:"ReportItemKeyList"`

	// List of task IDs in the report
	ReportTaskIdList []*ReportTaskIdList `json:"ReportTaskIdList,omitnil,omitempty" name:"ReportTaskIdList"`
}

type DescribeTaskLogURLRequest struct {
	*tchttp.BaseRequest
	
	// Type of the task. `0`: Preview; `1`: Download
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// List of task report IDs
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil,omitempty" name:"ReportItemKeyList"`

	// List of task IDs in the report
	ReportTaskIdList []*ReportTaskIdList `json:"ReportTaskIdList,omitnil,omitempty" name:"ReportTaskIdList"`
}

func (r *DescribeTaskLogURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "MemberId")
	delete(f, "ReportItemKeyList")
	delete(f, "ReportTaskIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogURLResponseParams struct {
	// Temp download URL of the report
	Data []*TaskLogURL `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogURLResponseParams `json:"Response"`
}

func (r *DescribeTaskLogURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopAttackInfoRequestParams struct {
	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 1: Attack type 2: Attacker
	QueryType *int64 `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Member ID of the Called Group Account
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 0: Default all 1: Asset ID 2: Domain name
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

type DescribeTopAttackInfoRequest struct {
	*tchttp.BaseRequest
	
	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 1: Attack type 2: Attacker
	QueryType *int64 `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Member ID of the Called Group Account
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 0: Default all 1: Asset ID 2: Domain name
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

func (r *DescribeTopAttackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopAttackInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "MemberId")
	delete(f, "OperatedMemberId")
	delete(f, "AssetName")
	delete(f, "AssetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopAttackInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopAttackInfoResponseParams struct {
	// Top attack types/attackers by count
	TopAttackInfo []*TagCount `json:"TopAttackInfo,omitnil,omitempty" name:"TopAttackInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopAttackInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopAttackInfoResponseParams `json:"Response"`
}

func (r *DescribeTopAttackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopAttackInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUebaRuleRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter criteria.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeUebaRuleRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter criteria.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeUebaRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUebaRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUebaRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUebaRuleResponseParams struct {
	// Total number.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Policy List
	Data []*UebaRule `json:"Data,omitnil,omitempty" name:"Data"`

	// Alarm category enumeration for custom policy
	AlterType []*FilterDataObject `json:"AlterType,omitnil,omitempty" name:"AlterType"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUebaRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUebaRuleResponseParams `json:"Response"`
}

func (r *DescribeUebaRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUebaRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCallRecordRequestParams struct {
	// Account UIN
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeUserCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// Account UIN
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeUserCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubUin")
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCallRecordResponseParams struct {
	// Account call record list.
	Data []*UserCallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of records.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCallRecordResponseParams `json:"Response"`
}

func (r *DescribeUserCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Query condition.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVULListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Query condition.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVULListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVULListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULListResponseParams struct {
	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Vulnerability list
	Data []*VULBaseInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Vulnerability Type List
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// Risk level list.
	RiskLevels []*FilterDataObject `json:"RiskLevels,omitnil,omitempty" name:"RiskLevels"`

	// Tag.
	Tags []*FilterDataObject `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Product support.
	ProductSupport []*FilterDataObject `json:"ProductSupport,omitnil,omitempty" name:"ProductSupport"`

	// Product support.
	CheckStatus []*FilterDataObject `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// Attack intensity enumeration.
	AttackHeat []*FilterDataObject `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVULListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVULListResponseParams `json:"Response"`
}

func (r *DescribeVULListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskAdvanceCFGListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Filter conditions.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVULRiskAdvanceCFGListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Filter conditions.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVULRiskAdvanceCFGListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskAdvanceCFGListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVULRiskAdvanceCFGListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskAdvanceCFGListResponseParams struct {
	// Configuration item list
	Data []*VULRiskAdvanceCFGList `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Risk Level Filter List
	RiskLevelLists []*FilterDataObject `json:"RiskLevelLists,omitnil,omitempty" name:"RiskLevelLists"`

	// Vulnerability Type Filter List
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// Recognition Source Filter List
	CheckFromLists []*FilterDataObject `json:"CheckFromLists,omitnil,omitempty" name:"CheckFromLists"`

	// Vulnerability tag list.
	VulTagList []*FilterDataObject `json:"VulTagList,omitnil,omitempty" name:"VulTagList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVULRiskAdvanceCFGListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVULRiskAdvanceCFGListResponseParams `json:"Response"`
}

func (r *DescribeVULRiskAdvanceCFGListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskAdvanceCFGListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskDetailRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Risk id.
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// pcMgrId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`
}

type DescribeVULRiskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Risk id.
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// pcMgrId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`
}

func (r *DescribeVULRiskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "RiskId")
	delete(f, "PCMGRId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVULRiskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskDetailResponseParams struct {
	// Security product support.
	ServiceSupport []*ServiceSupport `json:"ServiceSupport,omitnil,omitempty" name:"ServiceSupport"`

	// Vulnerability trends.
	VulTrend []*VulTrend `json:"VulTrend,omitnil,omitempty" name:"VulTrend"`

	// Vulnerability supplementary information.
	VulData *VULRiskInfo `json:"VulData,omitnil,omitempty" name:"VulData"`

	// Assistant q&a id.
	QuestionId *string `json:"QuestionId,omitnil,omitempty" name:"QuestionId"`

	// Session ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVULRiskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVULRiskDetailResponseParams `json:"Response"`
}

func (r *DescribeVULRiskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAssetsRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVpcAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVpcAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAssetsResponseParams struct {
	// Data list
	Data []*Vpc `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of VPCs
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// List of regions
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of AppIds
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcAssetsResponseParams `json:"Response"`
}

func (r *DescribeVpcAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulRiskListRequestParams struct {
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Cloud account ID.
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// Cloud service provider.
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

type DescribeVulRiskListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Group account member id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting type
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field.
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Cloud account ID.
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// Cloud service provider.
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

func (r *DescribeVulRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "CloudAccountID")
	delete(f, "Provider")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulRiskListResponseParams struct {
	// Number of vulnerabilities
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Vulnerability list
	VulRiskList []*VulRiskItem `json:"VulRiskList,omitnil,omitempty" name:"VulRiskList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVulRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulRiskListResponseParams `json:"Response"`
}

func (r *DescribeVulRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulViewVulRiskListRequestParams struct {
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tag
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeVulViewVulRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Filtered Content
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tag
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeVulViewVulRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulViewVulRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulViewVulRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulViewVulRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Vulnerability Risk List from Vulnerability Asset's Perspective
	Data []*VULViewVULRiskData `json:"Data,omitnil,omitempty" name:"Data"`

	// Danger Level List
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// Source List
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// Vulnerability Type List
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// tag enumeration.
	Tags []*FilterDataObject `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVulViewVulRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulViewVulRiskListResponseParams `json:"Response"`
}

func (r *DescribeVulViewVulRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulViewVulRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAssetVO struct {
	// Asset ID
	AssetId []*string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset Name
	AssetName []*string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type
	AssetType []*string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Region.
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// WAF Status
	WAFStatus *uint64 `json:"WAFStatus,omitnil,omitempty" name:"WAFStatus"`

	// Asset Creation Time
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Account ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Account name
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Core or Not
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Whether it is on-cloud asset.
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// network attack
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// Network access
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Network Interception
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// Inbound peak bandwidth
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// Outbound peak bandwidth
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// Cumulative Inbound Traffic
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// Cumulative Outbound Traffic
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// port risk
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// Vulnerability risk
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// Configuration risk
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Scan Task
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Domain
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// Resolve IP
	SeverIp []*string `json:"SeverIp,omitnil,omitempty" name:"SeverIp"`

	// Number of Bot Attacks
	BotCount *uint64 `json:"BotCount,omitnil,omitempty" name:"BotCount"`

	// Weak password risk
	WeakPassword *uint64 `json:"WeakPassword,omitnil,omitempty" name:"WeakPassword"`

	// Content risk
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil,omitempty" name:"WebContentRisk"`

	// tag
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Associated instance type
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Member ID information
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// CC Attack
	CCAttack *int64 `json:"CCAttack,omitnil,omitempty" name:"CCAttack"`

	// Web Attack
	WebAttack *int64 `json:"WebAttack,omitnil,omitempty" name:"WebAttack"`

	// Number of Risk Service Exposures
	ServiceRisk *uint64 `json:"ServiceRisk,omitnil,omitempty" name:"ServiceRisk"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Random Layer-3 Domain of Assets to Be Recognized
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// TXT Record Content of Pending Confirmation Assets
	VerifyTXTRecord *string `json:"VerifyTXTRecord,omitnil,omitempty" name:"VerifyTXTRecord"`

	// Authentication Status of Assets Pending Recognition. 0: Pending Authentication; 1: Authentication Succeeded; 2: Authentication in Progress; 3: TXT Authentication Failed; 4: Manual Authentication Failed.
	VerifyStatus *int64 `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`

	// Bot Access Count
	BotAccessCount *int64 `json:"BotAccessCount,omitnil,omitempty" name:"BotAccessCount"`
}

type Element struct {
	// Statistics type.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Statistics Object
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExposeAssetTypeItem struct {
	// Cloud service provider.
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// Vendor name.
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// Asset type.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Asset type name.
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`
}

type ExposesItem struct {
	// <p>Cloud vendor</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Cloud account name</p>
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// <p>Cloud Account</p>
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// <p>Domain name</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>IP</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Port or port range</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Open</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Risk type</p>
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// <p>acl type</p>
	AclType *string `json:"AclType,omitnil,omitempty" name:"AclType"`

	// <p>acl list</p>
	AclList *string `json:"AclList,omitnil,omitempty" name:"AclList"`

	// <p>Asset ID</p>
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Asset type</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>Number of port services</p>
	PortServiceCount *uint64 `json:"PortServiceCount,omitnil,omitempty" name:"PortServiceCount"`

	// <p>Number of high-risk ports</p>
	HighRiskPortServiceCount *uint64 `json:"HighRiskPortServiceCount,omitnil,omitempty" name:"HighRiskPortServiceCount"`

	// <p>Number of web applications</p>
	WebAppCount *uint64 `json:"WebAppCount,omitnil,omitempty" name:"WebAppCount"`

	// <p>Number of web applications at risk</p>
	RiskWebAppCount *uint64 `json:"RiskWebAppCount,omitnil,omitempty" name:"RiskWebAppCount"`

	// <p>Number of weak passwords.</p>
	WeakPasswordCount *uint64 `json:"WeakPasswordCount,omitnil,omitempty" name:"WeakPasswordCount"`

	// <p>Vulnerability count</p>
	VulCount *uint64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// <p>First discovery time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Latest update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Instance Type Name</p>
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`

	// <p>Open status</p>
	DisplayStatus *string `json:"DisplayStatus,omitnil,omitempty" name:"DisplayStatus"`

	// <p>Port status</p>
	DisplayRiskType *string `json:"DisplayRiskType,omitnil,omitempty" name:"DisplayRiskType"`

	// <p>Scan task status</p>
	ScanTaskStatus *string `json:"ScanTaskStatus,omitnil,omitempty" name:"ScanTaskStatus"`

	// <p>uuid</p>
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// <p>Whether a security check has been performed</p>
	HasScan *string `json:"HasScan,omitnil,omitempty" name:"HasScan"`

	// <p>Tenant ID</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Tenant ID string</p>
	AppIdStr *string `json:"AppIdStr,omitnil,omitempty" name:"AppIdStr"`

	// <p>Record ID</p>
	ExposureID *uint64 `json:"ExposureID,omitnil,omitempty" name:"ExposureID"`

	// <p>Number of open ports</p>
	PortDetectCount *uint64 `json:"PortDetectCount,omitnil,omitempty" name:"PortDetectCount"`

	// <p>Port exposure result</p>
	PortDetectResult *string `json:"PortDetectResult,omitnil,omitempty" name:"PortDetectResult"`

	// <p>Tag.</p>
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// <p>Remark</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>Number of risks to be governed</p>
	ToGovernedRiskCount *uint64 `json:"ToGovernedRiskCount,omitnil,omitempty" name:"ToGovernedRiskCount"`

	// <p>Risk content to be governed</p>
	ToGovernedRiskContent *string `json:"ToGovernedRiskContent,omitnil,omitempty" name:"ToGovernedRiskContent"`

	// <p>Type icon of asset</p>
	AssetTypeIconURL *string `json:"AssetTypeIconURL,omitnil,omitempty" name:"AssetTypeIconURL"`

	// <p>Asset type 3D icon</p>
	AssetTypeIconSolidURL *string `json:"AssetTypeIconSolidURL,omitnil,omitempty" name:"AssetTypeIconSolidURL"`
}

type Filter struct {
	// Max number of returned results
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Query offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting order. Values: `asc` (ascending), `desc` (descending).
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Specify the field used for sorting
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Filtered columns and content
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Start time of the query period. 
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type FilterDataObject struct {
	// Filter value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Filter name
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type Filters struct {
	// Filter condition name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter condition value list
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Exact match: 1 - exact match; default - fuzzy match
	ExactMatch *string `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type GateWayAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Asset ID.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// VPC IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// VPC id.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC Name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Outbound peak bandwidth.
	OutboundPeakBandwidth *string `json:"OutboundPeakBandwidth,omitnil,omitempty" name:"OutboundPeakBandwidth"`

	// Inbound peak bandwidth.
	InboundPeakBandwidth *string `json:"InboundPeakBandwidth,omitnil,omitempty" name:"InboundPeakBandwidth"`

	// Cumulative Outbound Traffic
	OutboundCumulativeFlow *string `json:"OutboundCumulativeFlow,omitnil,omitempty" name:"OutboundCumulativeFlow"`

	// Cumulative Inbound Traffic
	InboundCumulativeFlow *string `json:"InboundCumulativeFlow,omitnil,omitempty" name:"InboundCumulativeFlow"`

	// Network attack.
	NetworkAttack *int64 `json:"NetworkAttack,omitnil,omitempty" name:"NetworkAttack"`

	// Expose ports.
	ExposedPort *int64 `json:"ExposedPort,omitnil,omitempty" name:"ExposedPort"`

	// Exposed vulnerability.
	ExposedVUL *int64 `json:"ExposedVUL,omitnil,omitempty" name:"ExposedVUL"`

	// Configuration risk.
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil,omitempty" name:"ConfigureRisk"`

	// Creation time.
	// 
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Number of tasks.
	ScanTask *int64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Nickname.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// IPv6 address
	AddressIPV6 *string `json:"AddressIPV6,omitnil,omitempty" name:"AddressIPV6"`

	// Core or Not
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Risk service exposure.
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Gateway Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// TSE's Actual Gateway Region
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// Weak password risk.
	WeakPasswordRisk *uint64 `json:"WeakPasswordRisk,omitnil,omitempty" name:"WeakPasswordRisk"`
}

type HighBaseLineRiskItem struct {
	// Cloud account ID.
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// Instance ID.
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`

	// Instance status
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Risk name.
	RiskName *string `json:"RiskName,omitnil,omitempty" name:"RiskName"`

	// Risk classification.
	RiskCategory *string `json:"RiskCategory,omitnil,omitempty" name:"RiskCategory"`

	// Risk level.
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Risk description.
	RiskDesc *string `json:"RiskDesc,omitnil,omitempty" name:"RiskDesc"`

	// Risk result.
	RiskResult *string `json:"RiskResult,omitnil,omitempty" name:"RiskResult"`

	// Fixing suggestion
	FixAdvice *string `json:"FixAdvice,omitnil,omitempty" name:"FixAdvice"`

	// Linux vulnerability.
	RiskCategoryName *string `json:"RiskCategoryName,omitnil,omitempty" name:"RiskCategoryName"`

	// Risk name.
	RiskLevelName *string `json:"RiskLevelName,omitnil,omitempty" name:"RiskLevelName"`

	// Instance status
	InstanceStatusName *string `json:"InstanceStatusName,omitnil,omitempty" name:"InstanceStatusName"`

	// First detection time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last discovery time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Tenant ID.
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`
}

type IaCFile struct {
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>appid</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>File ID</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>File name.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>CI/CD name</p>
	CICDName *string `json:"CICDName,omitnil,omitempty" name:"CICDName"`

	// <p>File path</p>
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// <p>File type (1: Dockerfile, 2: Terraform, 3: KubernetesYaml)</p>
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>Total number of risks</p>
	RiskTotalCnt *uint64 `json:"RiskTotalCnt,omitnil,omitempty" name:"RiskTotalCnt"`

	// <p>Risk level count (0: Low risk, 1: Medium risk, 2: High risk, 3: Critical)</p>
	RiskLevelCnt []*KeyValueInt `json:"RiskLevelCnt,omitnil,omitempty" name:"RiskLevelCnt"`

	// <p>Scan time</p>
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// <p>Detection status (0: pending scan, 1: detecting, 2: completed, 3: detection exception)</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Scan failure type (0: No failure, 1: Detection timeout, 2: File format parsing failed, 3: Detection failed)</p>
	FailType *int64 `json:"FailType,omitnil,omitempty" name:"FailType"`
}

type IaCFileRisk struct {
	// <p>Risk level (0: low-risk, 1: medium-risk, 2: high-risk, 3: critical)</p>
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>Row count of risk location</p>
	Line *uint64 `json:"Line,omitnil,omitempty" name:"Line"`

	// <p>rule name</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>Problem description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Repair recommendation</p>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`
}

type IpAssetListVO struct {
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset Name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Cloud Defense Status
	CFWStatus *uint64 `json:"CFWStatus,omitnil,omitempty" name:"CFWStatus"`

	// Asset creation time
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Public IP Type
	PublicIpType *uint64 `json:"PublicIpType,omitnil,omitempty" name:"PublicIpType"`

	// vpc
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC Name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Name
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Core
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// On-Cloud
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// network attack
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// Network access
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Network Interception
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// Inbound bandwidth
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// Outbound bandwidth
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// Inbound traffic
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// outbound traffic
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Port risk
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// Vulnerability risk
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// Configuration risk
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Scan Task
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// weak password
	WeakPassword *uint64 `json:"WeakPassword,omitnil,omitempty" name:"WeakPassword"`

	// Content risk
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil,omitempty" name:"WebContentRisk"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// EIP Primary Key
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// Member ID information
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// risk service exposure
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Asset Authentication Status. 0-Pending Authentication; 1-Authentication Succeeded; 2-Authentication in Progress; 3+-Authentication Failed.
	VerifyStatus *int64 `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`
}

type KeySandboxCredential struct {
	// Credential ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// Credential name
	CredentialName *string `json:"CredentialName,omitnil,omitempty" name:"CredentialName"`

	// Credential Type
	// Enumeration value:
	// access: normal Key (Key-Value pair)
	// sts: STS temporary key credential
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// Effective machine scope
	CredentialEffectScope *CredentialEffectScope `json:"CredentialEffectScope,omitnil,omitempty" name:"CredentialEffectScope"`

	// Creation time.
	// Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	// Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type KeyValue struct {
	// Field
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KeyValueInt struct {
	// <p>Key</p>
	Key *int64 `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Value.</p>
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyIaCTokenPeriodRequestParams struct {
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Scanning result storage cycle</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type ModifyIaCTokenPeriodRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Scanning result storage cycle</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *ModifyIaCTokenPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIaCTokenPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIaCTokenPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIaCTokenPeriodResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIaCTokenPeriodResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIaCTokenPeriodResponseParams `json:"Response"`
}

func (r *ModifyIaCTokenPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIaCTokenPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrganizationAccountStatusRequestParams struct {
	// Modify group account status. 1 Enable, 0 Disable.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyOrganizationAccountStatusRequest struct {
	*tchttp.BaseRequest
	
	// Modify group account status. 1 Enable, 0 Disable.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyOrganizationAccountStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrganizationAccountStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOrganizationAccountStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrganizationAccountStatusResponseParams struct {
	// If the returned value is 0, the modification was successful.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOrganizationAccountStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOrganizationAccountStatusResponseParams `json:"Response"`
}

func (r *ModifyOrganizationAccountStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrganizationAccountStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterRiskStatusRequestParams struct {
	// Data of risk assets
	RiskStatusKeys []*RiskCenterStatusKey `json:"RiskStatusKeys,omitnil,omitempty" name:"RiskStatusKeys"`

	// Specify how you want to change the risk status. `1`: Change to Handled, `2`: Change to Ignored; `3`: Remove from Handled; `4`: Remove from Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk type. `0`: Port risk; `1`: Vulnerability; `2`: Weak password; `3`: Website content risk; `4`: Configuration risk; `5`: Risk services
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyRiskCenterRiskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Data of risk assets
	RiskStatusKeys []*RiskCenterStatusKey `json:"RiskStatusKeys,omitnil,omitempty" name:"RiskStatusKeys"`

	// Specify how you want to change the risk status. `1`: Change to Handled, `2`: Change to Ignored; `3`: Remove from Handled; `4`: Remove from Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk type. `0`: Port risk; `1`: Vulnerability; `2`: Weak password; `3`: Website content risk; `4`: Configuration risk; `5`: Risk services
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyRiskCenterRiskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterRiskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskStatusKeys")
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRiskCenterRiskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterRiskStatusResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRiskCenterRiskStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRiskCenterRiskStatusResponseParams `json:"Response"`
}

func (r *ModifyRiskCenterRiskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterRiskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterScanTaskRequestParams struct {
	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 0: Full Scan; 1: Specified Asset Scan; 2: Excluded Asset Scan; 3: Manual Entry Scan. 1 and 2 require the Assets field; 3 requires SelfDefiningAssets.
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// Scan items. port/poc/weakpass/webcontent/configrisk
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0: Periodic Task; 1: Scan Now; 2: Scheduled Scan; 3: Custom. If 0, 2, 3, ScanPlanContent is required.
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// Task ID to be Modified
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Scanned Asset Information List
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// Scan Plan Details
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// IP/Domain/URL Array
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Advanced configuration.
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// Checkup Mode. 0: Standard Mode; 1: Quick Mode; 2: Advanced Mode. Standard Mode by default.
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Task complete callback webhook url.
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

type ModifyRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 0: Full Scan; 1: Specified Asset Scan; 2: Excluded Asset Scan; 3: Manual Entry Scan. 1 and 2 require the Assets field; 3 requires SelfDefiningAssets.
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// Scan items. port/poc/weakpass/webcontent/configrisk
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0: Periodic Task; 1: Scan Now; 2: Scheduled Scan; 3: Custom. If 0, 2, 3, ScanPlanContent is required.
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// Task ID to be Modified
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Scanned Asset Information List
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// Scan Plan Details
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// IP/Domain/URL Array
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Advanced configuration.
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// Checkup Mode. 0: Standard Mode; 1: Quick Mode; 2: Advanced Mode. Standard Mode by default.
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Task complete callback webhook url.
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

func (r *ModifyRiskCenterScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "ScanAssetType")
	delete(f, "ScanItem")
	delete(f, "ScanPlanType")
	delete(f, "TaskId")
	delete(f, "MemberId")
	delete(f, "Assets")
	delete(f, "ScanPlanContent")
	delete(f, "SelfDefiningAssets")
	delete(f, "TaskAdvanceCFG")
	delete(f, "TaskMode")
	delete(f, "FinishWebHook")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRiskCenterScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterScanTaskResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 0: Modification succeeded; others: Failed; -1: Unauthenticated assets exist.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unauthenticated Asset List
	UnAuthAsset []*string `json:"UnAuthAsset,omitnil,omitempty" name:"UnAuthAsset"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRiskCenterScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRiskCenterScanTaskResponseParams `json:"Response"`
}

func (r *ModifyRiskCenterScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUebaRuleSwitchRequestParams struct {
	// Policy ID
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Enabling status
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyUebaRuleSwitchRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Enabling status
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyUebaRuleSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUebaRuleSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "Status")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUebaRuleSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUebaRuleSwitchResponseParams struct {
	// 0: Success; 1: Failure
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Returned information.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUebaRuleSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUebaRuleSwitchResponseParams `json:"Response"`
}

func (r *ModifyUebaRuleSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUebaRuleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NICAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Asset ID.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// VPC IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Public IP address
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// VPC id.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC Name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Outbound peak bandwidth.
	OutboundPeakBandwidth *string `json:"OutboundPeakBandwidth,omitnil,omitempty" name:"OutboundPeakBandwidth"`

	// Inbound peak bandwidth.
	InboundPeakBandwidth *string `json:"InboundPeakBandwidth,omitnil,omitempty" name:"InboundPeakBandwidth"`

	// Cumulative Outbound Traffic
	OutboundCumulativeFlow *string `json:"OutboundCumulativeFlow,omitnil,omitempty" name:"OutboundCumulativeFlow"`

	// Cumulative Inbound Traffic
	InboundCumulativeFlow *string `json:"InboundCumulativeFlow,omitnil,omitempty" name:"InboundCumulativeFlow"`

	// Network attack.
	NetworkAttack *int64 `json:"NetworkAttack,omitnil,omitempty" name:"NetworkAttack"`

	// Expose ports.
	ExposedPort *int64 `json:"ExposedPort,omitnil,omitempty" name:"ExposedPort"`

	// Exposed vulnerability.
	ExposedVUL *int64 `json:"ExposedVUL,omitnil,omitempty" name:"ExposedVUL"`

	// Configuration risk.
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil,omitempty" name:"ConfigureRisk"`

	// Creation time.
	// 
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Number of tasks.
	ScanTask *int64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Nickname.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Core or Not
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type NewAlertKey struct {
	// User AppID to Be Changed
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Alarm category
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Alarm Subcategory
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// Alarm source
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Alarm name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Alarm Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Time
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Status.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type OrganizationInfo struct {
	// member account name
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Department Node Name, Account's Department
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Member/Admin/DelegatedAdmin/EntityAdmin, corresponding to Member/Administrator/Delegated Administrator/Entity Administrator
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Member Account ID
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Account Joining Method: Create/Invite.
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`

	// Group Name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// administrator account name
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// Administrator UIN
	AdminUin *string `json:"AdminUin,omitnil,omitempty" name:"AdminUin"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Number of departments
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// Number of members
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// Number of sub-accounts
	SubAccountCount *int64 `json:"SubAccountCount,omitnil,omitempty" name:"SubAccountCount"`

	// Number of abnormal sub-accounts
	AbnormalSubUserCount *int64 `json:"AbnormalSubUserCount,omitnil,omitempty" name:"AbnormalSubUserCount"`

	// Group Relationship Policy Permissions
	GroupPermission []*string `json:"GroupPermission,omitnil,omitempty" name:"GroupPermission"`

	// Membership Policy Permissions
	MemberPermission []*string `json:"MemberPermission,omitnil,omitempty" name:"MemberPermission"`

	// Group Payment Mode. 0: Self-payment; 1: Proxy Payment.
	GroupPayMode *int64 `json:"GroupPayMode,omitnil,omitempty" name:"GroupPayMode"`

	// Personal Payment Mode. 0: Self-payment; 1: Proxy payment.
	MemberPayMode *int64 `json:"MemberPayMode,omitnil,omitempty" name:"MemberPayMode"`

	// Not enabled if empty. Otherwise, different strings correspond to different versions. Common for General, regardless of version.
	CFWProtect *string `json:"CFWProtect,omitnil,omitempty" name:"CFWProtect"`

	// Not enabled if empty. Otherwise, different strings correspond to different versions. Common for General, regardless of version.
	WAFProtect *string `json:"WAFProtect,omitnil,omitempty" name:"WAFProtect"`

	// Not enabled if empty. Otherwise, different strings correspond to different versions. Common for General, regardless of version.
	CWPProtect *string `json:"CWPProtect,omitnil,omitempty" name:"CWPProtect"`

	// Array of Collections for All Departments
	Departments []*string `json:"Departments,omitnil,omitempty" name:"Departments"`

	// Member Creation Time
	MemberCreateTime *string `json:"MemberCreateTime,omitnil,omitempty" name:"MemberCreateTime"`

	// Advanced/Enterprise/Ultimate 
	CSIPProtect *string `json:"CSIPProtect,omitnil,omitempty" name:"CSIPProtect"`

	// 1 indicates the quota consumer.
	QuotaConsumer *int64 `json:"QuotaConsumer,omitnil,omitempty" name:"QuotaConsumer"`

	// Number of activations by admin/delegated admin
	EnableAdminCount *int64 `json:"EnableAdminCount,omitnil,omitempty" name:"EnableAdminCount"`

	// Account Multi-Cloud Information Statistics, in array format. Refer to the description of CloudCountDesc for details.
	CloudCountDesc []*CloudCountDesc `json:"CloudCountDesc,omitnil,omitempty" name:"CloudCountDesc"`

	// Total number of admins/delegated admins
	AdminCount *int64 `json:"AdminCount,omitnil,omitempty" name:"AdminCount"`
}

type OrganizationUserInfo struct {
	// Member Account UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// member account name
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Department Node Name, Account's Department
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Number of assets
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// Number of risks
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// Number of Attacks
	AttackCount *int64 `json:"AttackCount,omitnil,omitempty" name:"AttackCount"`

	// Member/Admin/; Member or Administrator
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Member Account ID
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Member Account AppID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Account Joining Method: Create/Invite.
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`

	// Not enabled if empty. Otherwise, different strings correspond to different versions. Common for General, regardless of version.
	CFWProtect *string `json:"CFWProtect,omitnil,omitempty" name:"CFWProtect"`

	// Not enabled if empty. Otherwise, different strings correspond to different versions. Common for General, regardless of version.
	WAFProtect *string `json:"WAFProtect,omitnil,omitempty" name:"WAFProtect"`

	// Not enabled if empty. Otherwise, different strings correspond to different versions. Common for General, regardless of version.
	CWPProtect *string `json:"CWPProtect,omitnil,omitempty" name:"CWPProtect"`

	// 1-Enabled; 0-Not Enabled.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Free       // Free Edition  Advanced   //Advanced Edition Enterprise //Enterprise Edition Ultimate   //Premium Edition
	CSIPProtect *string `json:"CSIPProtect,omitnil,omitempty" name:"CSIPProtect"`

	// 1 for quota consumer.
	QuotaConsumer *int64 `json:"QuotaConsumer,omitnil,omitempty" name:"QuotaConsumer"`

	// Account Type. 0 for Tencent Cloud account; 1 for AWS account.
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 0 for default value, 1 for 10 minutes, 2 for 1 hour, 3 for 24 hours.
	SyncFrequency *int64 `json:"SyncFrequency,omitnil,omitempty" name:"SyncFrequency"`

	// Whether the multi-cloud account is expired.
	IsExpired *bool `json:"IsExpired,omitnil,omitempty" name:"IsExpired"`

	// Multi-Cloud Account Permission List
	PermissionList []*string `json:"PermissionList,omitnil,omitempty" name:"PermissionList"`

	// 1
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// Tencent Cloud Group Account
	// Tencent Cloud access account
	// 2: non-Tencent Cloud
	TcMemberType *int64 `json:"TcMemberType,omitnil,omitempty" name:"TcMemberType"`

	// Number of sub-accounts.
	SubUserCount *int64 `json:"SubUserCount,omitnil,omitempty" name:"SubUserCount"`

	// Joining method details
	JoinTypeInfo *string `json:"JoinTypeInfo,omitnil,omitempty" name:"JoinTypeInfo"`
}

type PortRiskAdvanceCFGParamItem struct {
	// Port Collection, separated by commas.
	PortSets *string `json:"PortSets,omitnil,omitempty" name:"PortSets"`

	// Detection Item Type. 0-System-Defined; 1-User-Defined.
	CheckType *int64 `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// Detection item description
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Enable/Disable. 1-Enable; 0-Disable.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type PortViewPortRisk struct {
	// Unprocessed quantity.
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// Risk level, low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Network protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Suggested action. `0`: Keep as it is; `1`: Block access requests; `2`: Block the port
	Suggestion *uint64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Number of Affected Assets
	AffectAssetCount *string `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Source recognition
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Data entry key
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`
}

type ProductSupport struct {
	// True supports scanning. False does not support scanning.
	VSSScan *bool `json:"VSSScan,omitnil,omitempty" name:"VSSScan"`

	// 0-Not Supported; 1-Supported
	CWPScan *string `json:"CWPScan,omitnil,omitempty" name:"CWPScan"`

	// 1 indicates virtual patches supported, 0 or null indicates not supported.
	CFWPatch *string `json:"CFWPatch,omitnil,omitempty" name:"CFWPatch"`

	// 0-Not Supported; 1-Supported	
	WafPatch *int64 `json:"WafPatch,omitnil,omitempty" name:"WafPatch"`

	// 0-Not Supported; 1-Supported	
	CWPFix *int64 `json:"CWPFix,omitnil,omitempty" name:"CWPFix"`

	// cveid
	CveId *string `json:"CveId,omitnil,omitempty" name:"CveId"`
}

type PublicIpDomainListKey struct {
	// IP/Domain
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`
}

type RelatedEvent struct {
	// Event ID
	EventID *string `json:"EventID,omitnil,omitempty" name:"EventID"`

	// Event description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Number of Alarms Associated with Event
	RelatedCount *int64 `json:"RelatedCount,omitnil,omitempty" name:"RelatedCount"`
}

type ReportItemKey struct {
	// List of report IDs.
	TaskLogList []*string `json:"TaskLogList,omitnil,omitempty" name:"TaskLogList"`
}

type ReportTaskIdList struct {
	// List of task IDs
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type RepositoryImageVO struct {
	// User appid.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Nickname.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Mirror id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Image name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Image creation time.
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// Image Size with Unit
	InstanceSize *string `json:"InstanceSize,omitnil,omitempty" name:"InstanceSize"`

	// Build times.
	BuildCount *int64 `json:"BuildCount,omitnil,omitempty" name:"BuildCount"`

	// Image type.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Authorization status.
	AuthStatus *int64 `json:"AuthStatus,omitnil,omitempty" name:"AuthStatus"`

	// Mirror version.
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Repository address.
	RepositoryUrl *string `json:"RepositoryUrl,omitnil,omitempty" name:"RepositoryUrl"`

	// Repository name.
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Core or Not
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Vulnerability risk.
	VulRisk *int64 `json:"VulRisk,omitnil,omitempty" name:"VulRisk"`

	// Check task.
	CheckCount *int64 `json:"CheckCount,omitnil,omitempty" name:"CheckCount"`

	// Health Checkup Time
	CheckTime *string `json:"CheckTime,omitnil,omitempty" name:"CheckTime"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type RiskCallRecord struct {
	// API name.
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Chinese description of the API.
	EventDescCN *string `json:"EventDescCN,omitnil,omitempty" name:"EventDescCN"`

	// Interface description.
	EventDescEN *string `json:"EventDescEN,omitnil,omitempty" name:"EventDescEN"`

	// Product name
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Product Chinese Name
	ProductNameCN *string `json:"ProductNameCN,omitnil,omitempty" name:"ProductNameCN"`

	// Number of calls.
	CallCount *int64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`
}

type RiskCenterStatusKey struct {
	// Risk ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Public IP/domain name
	PublicIPDomain *string `json:"PublicIPDomain,omitnil,omitempty" name:"PublicIPDomain"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type RiskDetailItem struct {
	// <p>First discovery time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Risk status</p>
	RiskStatus *int64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// <p>Risk content</p>
	RiskContent *string `json:"RiskContent,omitnil,omitempty" name:"RiskContent"`

	// <p>Cloud service provider</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Vendor name</p>
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// <p>cloud account</p>
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// <p>Cloud account name</p>
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Risk ID</p>
	RiskId *uint64 `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// <p>Risk rule ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>Risk verification status</p>
	CheckStatus *string `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// <p>User AppID</p>
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>Asset type</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

type RiskRuleInfo struct {
	// Risk Check Item ID
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Cloud vendor name
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// risk name
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// Check type
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// Risk level
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Risk damage
	RiskInfluence *string `json:"RiskInfluence,omitnil,omitempty" name:"RiskInfluence"`

	// Risk remediation guide report link
	RiskFixAdvance *string `json:"RiskFixAdvance,omitnil,omitempty" name:"RiskFixAdvance"`

	// Boundary control
	DispositionType *string `json:"DispositionType,omitnil,omitempty" name:"DispositionType"`
}

type RiskRuleItem struct {
	// Risk Check Item ID
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// Cloud vendor name
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance Type Name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// risk name
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// Check type
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// Risk level
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// Risk damage
	RiskInfluence *string `json:"RiskInfluence,omitnil,omitempty" name:"RiskInfluence"`
}

type RoleInfo struct {
	// IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// HostIP
	HostIP *string `json:"HostIP,omitnil,omitempty" name:"HostIP"`

	// Original IP
	OriginIP *string `json:"OriginIP,omitnil,omitempty" name:"OriginIP"`

	// Port.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// asset ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// city
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Province
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// nation
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// Address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// latitude
	Latitude *string `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// longitude
	Longitude *string `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// Information.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// Domain
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Enterprise Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Account
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// Family Group
	Family *string `json:"Family,omitnil,omitempty" name:"Family"`

	// Virus name
	VirusName *string `json:"VirusName,omitnil,omitempty" name:"VirusName"`

	// MD5 Value
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// Malicious process filename
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 1-Host Assets; 2-Domain Assets; 3-Network Assets
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Information Fields of Source Log Analysis
	FromLogAnalysisData []*KeyValue `json:"FromLogAnalysisData,omitnil,omitempty" name:"FromLogAnalysisData"`

	// Container name
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// container ID
	ContainerID *string `json:"ContainerID,omitnil,omitempty" name:"ContainerID"`
}

type STSCredentialOutput struct {
	// Credential provider flag (original text), such as tencentCam, aws, aliyun
	System *string `json:"System,omitnil,omitempty" name:"System"`

	// SecretID (masked)
	// Supplementary description: Reserve the first 3 and last 4 digits, replace the middle with ***; replace all with *** if the length is less than 7.
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// SecretKey (masked)
	// Supplementary description: Reserve the first 3 and last 4 digits, replace the middle with ***; replace all with *** if the length is less than 7.
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
}

type ScanTaskInfo struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Task Status Code: 1-Awaiting Start, 2-Scanning, 3-Scan Error, 4-Scan Completed.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Task progress
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task Completion Time
	TaskTime *string `json:"TaskTime,omitnil,omitempty" name:"TaskTime"`

	// report ID
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// report name
	ReportName *string `json:"ReportName,omitnil,omitempty" name:"ReportName"`

	// Scanning Schedule. 0-Periodic Task; 1-Scan Now; 2-Scheduled Scan; 3-Custom.
	ScanPlan *int64 `json:"ScanPlan,omitnil,omitempty" name:"ScanPlan"`

	// Number of Associated Assets
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// APP ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Host Account ID
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type ScanTaskInfoList struct {
	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// start time of the task
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Cron Format
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// 0-Periodic Task; 1-Scan Now; 2-Scheduled Scan; 3-Custom.
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Creation time.
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Custom Specified Scan Asset Information
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Estimated Time
	PredictTime *int64 `json:"PredictTime,omitnil,omitempty" name:"PredictTime"`

	// Estimated Completion Time
	PredictEndTime *string `json:"PredictEndTime,omitnil,omitempty" name:"PredictEndTime"`

	// Report Count
	ReportNumber *int64 `json:"ReportNumber,omitnil,omitempty" name:"ReportNumber"`

	// Number of assets
	AssetNumber *int64 `json:"AssetNumber,omitnil,omitempty" name:"AssetNumber"`

	// Scan Status. 0-Initial Value; 1-Scanning; 2-Scan Completed; 3-Scan Error; 4-Scan Stopped.
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// Task progress
	Percent *float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// port/poc/weakpass/webcontent/configrisk
	ScanItem *string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0-Full Scan; 1-Specified Asset Scan; 2-Excluded Asset Scan; 3-Custom Specified Asset Scan.
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// VSS Subtask ID
	VSSTaskId *string `json:"VSSTaskId,omitnil,omitempty" name:"VSSTaskId"`

	// CSPM Subtask ID
	CSPMTaskId *string `json:"CSPMTaskId,omitnil,omitempty" name:"CSPMTaskId"`

	// Host Vulnerability Scan Subtask ID
	CWPPOCId *string `json:"CWPPOCId,omitnil,omitempty" name:"CWPPOCId"`

	// Host Baseline Subtask ID
	CWPBlId *string `json:"CWPBlId,omitnil,omitempty" name:"CWPBlId"`

	// VSS Subtask Progress
	VSSTaskProcess *int64 `json:"VSSTaskProcess,omitnil,omitempty" name:"VSSTaskProcess"`

	// CSPM Subtask Progress
	CSPMTaskProcess *uint64 `json:"CSPMTaskProcess,omitnil,omitempty" name:"CSPMTaskProcess"`

	// Host Vulnerability Scan Subtask Progress
	CWPPOCProcess *int64 `json:"CWPPOCProcess,omitnil,omitempty" name:"CWPPOCProcess"`

	// Host Baseline Subtask Progress
	CWPBlProcess *uint64 `json:"CWPBlProcess,omitnil,omitempty" name:"CWPBlProcess"`

	// Exception status code
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Exception information
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Number of Days for Periodic Task to Start
	StartDay *int64 `json:"StartDay,omitnil,omitempty" name:"StartDay"`

	// Scanning Frequency, in Days. 1-Daily; 7-Weekly; 30-Monthly; 0-Scan Once.
	Frequency *int64 `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// Completion Count
	CompleteNumber *int64 `json:"CompleteNumber,omitnil,omitempty" name:"CompleteNumber"`

	// Completed Asset Count
	CompleteAssetNumber *int64 `json:"CompleteAssetNumber,omitnil,omitempty" name:"CompleteAssetNumber"`

	// risk count
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// Asset
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// User Appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Host Account ID
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Checkup Mode. 0-Standard Mode; 1-Quick Mode; 2-Advanced Mode.
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Scan Source
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// Whether health checkup is limited or exempted. 0-No; 1-Yes.
	IsFree *int64 `json:"IsFree,omitnil,omitempty" name:"IsFree"`

	// Whether it can be deleted. 1-Yes; 0-No. For use with multi-account management.
	IsDelete *int64 `json:"IsDelete,omitnil,omitempty" name:"IsDelete"`

	// Task Source Type. 0: Default; 1: Assistant; 2: Health Checkup Items.
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type ServerRisk struct {
	// Service tag
	ServiceTag *string `json:"ServiceTag,omitnil,omitempty" name:"ServiceTag"`

	// Port.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Risk level, low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Network protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Risk details
	RiskDetails *string `json:"RiskDetails,omitnil,omitempty" name:"RiskDetails"`

	// Handling suggestion
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Status, 0 unprocessed, 1 processed, 2 ignored, 3 defended by cloud protection
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Service Snapshot
	ServiceSnapshot *string `json:"ServiceSnapshot,omitnil,omitempty" name:"ServiceSnapshot"`

	// Service Access URL
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// List Index Value
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Risk list
	RiskList []*ServerRiskSuggestion `json:"RiskList,omitnil,omitempty" name:"RiskList"`

	// Recommendation List
	SuggestionList []*ServerRiskSuggestion `json:"SuggestionList,omitnil,omitempty" name:"SuggestionList"`

	// HTTP Response Status Code
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// New risk level, high_risk high risk suspect Suspected Normal Does not have risks currently
	NewLevel *string `json:"NewLevel,omitnil,omitempty" name:"NewLevel"`

	// Status, 0 unprocessed, 1 processed, 2 ignored, 3 defended by cloud protection, 4 no action is required
	XspmStatus *uint64 `json:"XspmStatus,omitnil,omitempty" name:"XspmStatus"`
}

type ServerRiskSuggestion struct {
	// Title.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Details.
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`
}

type ServiceSupport struct {
	// Product name.
	// "cfw_waf_virtual", "cwp_detect", "cwp_defense", "cwp_fix"
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Total number of processed assets.
	SupportHandledCount *int64 `json:"SupportHandledCount,omitnil,omitempty" name:"SupportHandledCount"`

	// Total number of supported assets.
	SupportTotalCount *int64 `json:"SupportTotalCount,omitnil,omitempty" name:"SupportTotalCount"`

	// Whether the product is supported: 1 for supported; 0 for unsupported.
	IsSupport *bool `json:"IsSupport,omitnil,omitempty" name:"IsSupport"`
}

type SkillCapabilityTag struct {
	// Capacity tag identification, suitable for program judgment, filtering or aggregation usage
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// Capacity Tag Display Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type SkillRuleCatalogItem struct {
	// Fusion rule ID (9xxxx)
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Risk category name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type SkillScanEngineResult struct {
	// Sub-engine type
	// Enumeration value:
	// AI: AI engine
	// STATIC: Static analysis engine
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// The rule list hit by the engine
	RuleList []*SkillScanRuleHit `json:"RuleList,omitnil,omitempty" name:"RuleList"`
}

type SkillScanItem struct {
	// <p>Skill name</p>
	SkillName *string `json:"SkillName,omitnil,omitempty" name:"SkillName"`

	// <p>Skill description to help understand its primary purpose</p>
	SkillDescription *string `json:"SkillDescription,omitnil,omitempty" name:"SkillDescription"`

	// <p>SHA256 Hash of the ZIP file<br>Parameter format: sha256:&lt;64-bit hex&gt;</p>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// <p>The number of actual files after decompressing the original uploaded ZIP file, also within the billing limit. Each file is counted as one limit after a successful scan.</p>
	UploadFileCount *int64 `json:"UploadFileCount,omitnil,omitempty" name:"UploadFileCount"`

	// <p>Comprehensive risk level<br>Enumeration value:<br>malicious: Malicious<br>suspicious: Suspicious<br>benign: Trustworthy</p>
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>Risk master tag fusion rule ID (9xxxx) is generated by the server from the hit fusion risk tags. It is empty when benign and no rule hits occur. The display name can be obtained via RuleCatalog.</p>
	PrimaryRuleID *string `json:"PrimaryRuleID,omitnil,omitempty" name:"PrimaryRuleID"`

	// <p>Comprehensive handling suggestions for guiding the caller to prioritize actions such as decommissioning, isolation, repair, and recheck. The historical result may be empty. Returns copywriting in English when Language=en-US is passed.</p>
	Mitigation *string `json:"Mitigation,omitnil,omitempty" name:"Mitigation"`

	// <p>Comprehensive risk description provides an overview of risks found in this detection. Returns English copy when Language=en-US is passed.</p>
	RiskDescription *string `json:"RiskDescription,omitnil,omitempty" name:"RiskDescription"`

	// <p>Security score value ranges from 0 to 100. Supplementary explanation: the higher the score, the more secure.</p>
	SecurityScore *int64 `json:"SecurityScore,omitnil,omitempty" name:"SecurityScore"`

	// <p>Engine version number used in this scan</p>
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>Skill ability tag list describes the ability features or application scenarios of Skill. It is not equal to risk tag and does not participate in risk level judgment. When Language=en-US is passed, Name switches to English while ID remains unchanged.</p>
	CapabilityTags []*SkillCapabilityTag `json:"CapabilityTags,omitnil,omitempty" name:"CapabilityTags"`

	// <p>Complete set of fusion rule directory, including all fusion rule categories (9xxxx). The caller can show category tags accordingly without the need to maintain a mapping table locally. Returns English name when Language=en-US is passed.</p>
	RuleCatalog []*SkillRuleCatalogItem `json:"RuleCatalog,omitnil,omitempty" name:"RuleCatalog"`

	// <p>Scan result details, grouped by sub-engine. Each element contains ScanType (engine type) and RuleList (hit rule list). RuleID within the rules uses fusion code (9xxxx) and can be cross-referenced with RuleCatalog. Description returns in English when Language=en-US is passed.</p>
	ScanItems []*SkillScanEngineResult `json:"ScanItems,omitnil,omitempty" name:"ScanItems"`

	// <p>Comprehensive security audit report address (pre-signed URL). The valid period is controlled by the request parameter ReportURLExpireHours.</p>
	ReportURL *string `json:"ReportURL,omitnil,omitempty" name:"ReportURL"`

	// <p>Scan completion time. Only available when Status=SUCCESS<br>Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format)</p>
	ScannedAt *string `json:"ScannedAt,omitnil,omitempty" name:"ScannedAt"`

	// <p>Task creation time. Only available when Status=SCANNING<br>Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format)</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>Failure time. Only valid when Status=FAILED<br>Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format)</p>
	FailedAt *string `json:"FailedAt,omitnil,omitempty" name:"FailedAt"`

	// <p>Failure reason description. Only available when Status=FAILED.</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type SkillScanRuleHit struct {
	// Fusion rule number (9xxxx) can be cross-referenced with RuleCatalog.
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Current description of the specific detection for the matched rule, including file location, behavioral features, risks, etc.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type SkillState struct {
	// SKILL Installation Status
	// Enumeration value:
	// 0: Not installed
	// Installing
	// 2: Installed
	// 3: Installation failure
	// 4: Uninstalling
	// 5: Uninstallation failed.
	SkillInstallStatus *int64 `json:"SkillInstallStatus,omitnil,omitempty" name:"SkillInstallStatus"`

	// SKILL installation/uninstallation operation time
	// Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format).
	SkillInstallTime *string `json:"SkillInstallTime,omitnil,omitempty" name:"SkillInstallTime"`

	// SKILL installation/uninstallation result description information
	SkillInstallResult *string `json:"SkillInstallResult,omitnil,omitempty" name:"SkillInstallResult"`
}

type SourceIPAsset struct {
	// id of the source IP.
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Source IP.
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Account associate APPID.
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// IP region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// API call method.
	// -1: uncounted.
	// 0: console invocation.
	// 1:API
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// IP type.
	// 0: within the account (unremarked).
	// 1: external accounts (unremarked).
	// 2: within the account (remarked).
	// 3: external account (remarked).
	IPType *int64 `json:"IPType,omitnil,omitempty" name:"IPType"`

	// Alarm information list.
	AccessKeyAlarmList []*AccessKeyAlarmInfo `json:"AccessKeyAlarmList,omitnil,omitempty" name:"AccessKeyAlarmList"`

	// ak information list.
	AKInfo []*AKInfo `json:"AKInfo,omitnil,omitempty" name:"AKInfo"`

	// Number of API calls.
	ActionCount *int64 `json:"ActionCount,omitnil,omitempty" name:"ActionCount"`

	// Last access Time
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// IP associated instance ID. if an empty string, represents an asset not within the account.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Associated instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Account associate Uin.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Nickname.
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// Display status.
	ShowStatus *bool `json:"ShowStatus,omitnil,omitempty" name:"ShowStatus"`

	// ISP field.
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// vpc information outside the account.
	VpcInfo []*SourceIPVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// Cloud type.
	// 0 for tencent cloud.
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

type SourceIPVpcInfo struct {
	// Account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// App ID of the VPC.
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// vpc id
	VpcID *string `json:"VpcID,omitnil,omitempty" name:"VpcID"`

	// vpc name.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`
}

type StandardItem struct {
	// Specification ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Standard name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type StandardTerm struct {
	// Tag.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Clause
	Terms []*string `json:"Terms,omitnil,omitempty" name:"Terms"`
}

type StatisticalFilter struct {
	// 0: Not based on statistical testing
	// 1: Occurrence count higher than a fixed value
	// 2: Occurrence count exceeds 100 percent of the period average
	// 3: Occurrences higher than 50 percent of the user average
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`

	// Statistical value
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type StopRiskCenterTaskRequestParams struct {
	// List of task IDs
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type StopRiskCenterTaskRequest struct {
	*tchttp.BaseRequest
	
	// List of task IDs
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *StopRiskCenterTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRiskCenterTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRiskCenterTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRiskCenterTaskResponseParams struct {
	// `0`: Operation succeeded; Others: failed
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopRiskCenterTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopRiskCenterTaskResponseParams `json:"Response"`
}

func (r *StopRiskCenterTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRiskCenterTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubUserInfo struct {
	// <p>Primary key ID, with no business significance, only serves as a unique key.</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>Sub-account Appid</p>
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>Sub-account UIn</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>Sub-account name</p>
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// <p>Root Account Appid</p>
	OwnerAppID *string `json:"OwnerAppID,omitnil,omitempty" name:"OwnerAppID"`

	// <p>Root account Uin</p>
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// <p>Root account name</p>
	OwnerNickName *string `json:"OwnerNickName,omitnil,omitempty" name:"OwnerNickName"`

	// <p>Member ID information belonging to main account</p>
	OwnerMemberID *string `json:"OwnerMemberID,omitnil,omitempty" name:"OwnerMemberID"`

	// <p>Account type. 0 indicates a Tencent Cloud account, and 1 indicates an AWS account.</p>
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// <p>Number of accessible services</p>
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// <p>Number of accessible APIs</p>
	InterfaceCount *int64 `json:"InterfaceCount,omitnil,omitempty" name:"InterfaceCount"`

	// <p>Number of accessible resources</p>
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// <p>Number of access/behavior logs</p>
	LogCount *int64 `json:"LogCount,omitnil,omitempty" name:"LogCount"`

	// <p>Permission configuration risk</p>
	ConfigRiskCount *int64 `json:"ConfigRiskCount,omitnil,omitempty" name:"ConfigRiskCount"`

	// <p>Dangerous behavior alarm</p>
	ActionRiskCount *int64 `json:"ActionRiskCount,omitnil,omitempty" name:"ActionRiskCount"`

	// <p>Whether to access operation audit log</p>
	IsAccessCloudAudit *bool `json:"IsAccessCloudAudit,omitnil,omitempty" name:"IsAccessCloudAudit"`

	// <p>Security check for configuration risk required or not</p>
	IsAccessCheck *bool `json:"IsAccessCheck,omitnil,omitempty" name:"IsAccessCheck"`

	// <p>Whether configure user behavior management policy</p>
	IsAccessUeba *bool `json:"IsAccessUeba,omitnil,omitempty" name:"IsAccessUeba"`

	// <p>Creation time (Unix timestamp).</p>
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type SubnetAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// User name
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// CIDR block
	CIDR *string `json:"CIDR,omitnil,omitempty" name:"CIDR"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of CVMs
	CVM *int64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// Number of available IPs
	AvailableIp *int64 `json:"AvailableIp,omitnil,omitempty" name:"AvailableIp"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Configuration risks
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil,omitempty" name:"ConfigureRisk"`

	// Number of tasks.
	ScanTask *int64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Core or Not
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type Tag struct {
	// Tag name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagCount struct {
	// Product Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of logs.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type Tags struct {
	// Host tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Host tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskAdvanceCFG struct {
	// Port Risk Advanced Configuration
	PortRisk []*PortRiskAdvanceCFGParamItem `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// Advanced vulnerability scan configuration
	VulRisk []*TaskCenterVulRiskInputParam `json:"VulRisk,omitnil,omitempty" name:"VulRisk"`

	// Advanced weak password check configuration
	WeakPwdRisk []*TaskCenterWeakPwdRiskInputParam `json:"WeakPwdRisk,omitnil,omitempty" name:"WeakPwdRisk"`

	// Advanced configuration risk scan configuration
	CFGRisk []*TaskCenterCFGRiskInputParam `json:"CFGRisk,omitnil,omitempty" name:"CFGRisk"`
}

type TaskAssetObject struct {
	// Asset name.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Asset category.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// IP, domain name, asset ID, database ID, and more
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Unique ID of Multi-Cloud Assets
	Arn *string `json:"Arn,omitnil,omitempty" name:"Arn"`
}

type TaskCenterCFGRiskInputParam struct {
	// Check item ID
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// Whether to enable. `0`: no, `1`: yes.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type TaskCenterVulRiskInputParam struct {
	// Risk ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// Whether to enable. `0`: no, `1`: yes.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type TaskCenterWeakPwdRiskInputParam struct {
	// Check item ID
	CheckItemId *int64 `json:"CheckItemId,omitnil,omitempty" name:"CheckItemId"`

	// Whether to enable. `0`: no, `1`: yes.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type TaskIdListKey struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// APP ID
	TargetAppId *string `json:"TargetAppId,omitnil,omitempty" name:"TargetAppId"`
}

type TaskLogInfo struct {
	// report name
	TaskLogName *string `json:"TaskLogName,omitnil,omitempty" name:"TaskLogName"`

	// Report ID.
	TaskLogId *string `json:"TaskLogId,omitnil,omitempty" name:"TaskLogId"`

	// Associated Asset Count
	AssetsNumber *int64 `json:"AssetsNumber,omitnil,omitempty" name:"AssetsNumber"`

	// Security Risk Count
	RiskNumber *int64 `json:"RiskNumber,omitnil,omitempty" name:"RiskNumber"`

	// Report generation time
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Task Status Code. 0-Initial Value; 1-Scanning; 2-Scan Completed; 3-Scan Error; 4-Stopped; 5-Halted; 6-Task Has Been Restarted.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Associated Task Name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Scan start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task Center Scan Task ID
	TaskCenterTaskId *string `json:"TaskCenterTaskId,omitnil,omitempty" name:"TaskCenterTaskId"`

	// Tenant ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Host Account ID
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// User name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Report Type. 1: Security Checkup; 2: Daily Report; 3: Weekly Report; 4: Monthly Report.
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// Report Template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type TaskLogURL struct {
	// Temporary Link for Report Download
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// Task Report ID
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// Task Report Name
	TaskLogName *string `json:"TaskLogName,omitnil,omitempty" name:"TaskLogName"`

	// APP ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type TrafficPluginState struct {
	// Plugin installation status (upper layer aggregation)
	// Enumeration value:
	// NONE: Not installed
	// INSTALLING
	// INSTALLED: Installed
	// INSTALL_FAIL: Installation failure
	InstallStatus *string `json:"InstallStatus,omitnil,omitempty" name:"InstallStatus"`

	// Plugin installation sub-status. The value corresponds to InstallStatus: empty string when not installed (InstallStatus=UNINSTALL); SUCCESS when successfully installed (InstallStatus=INSTALLED); specific failure reason when installation failure (InstallStatus=INSTALL_FAIL).
	// Enumeration value:
	// NOT_SUPPORT: Unsupported environment
	// CONTAINER_NOT_FOUND: Container does not exist.
	// RESTART required
	// CA_FAILED: CA failed
	// EBPF_FAILED: eBPF failed
	// IPTABLE_FAILED: iptables failed.
	// REDIRECT_FAILED: Traffic redirection failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Status copywriting (internationalization description derived from Status based on request language)
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Recent activity time of the plug-in
	// Parameter format: YYYY-MM-DDTHH:mm:ssZ (ISO8601 format).
	ActivityTime *string `json:"ActivityTime,omitnil,omitempty" name:"ActivityTime"`
}

type TrafficRuleState struct {
	// <p>Sandbox plug-in module name</p>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>Sandbox rule status</p><p>Enumeration value:</p><ul><li>ON: Enable</li><li>OFF: Disable</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type UebaCustomRule struct {
	// Policy name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 1: Cloud account
	// 2: Custom user
	UserType *int64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// Occurrence time
	// 10 minutes
	// 2:1 hour
	// 3: One day
	// 4: A week
	// 5: One month
	TimeInterval *int64 `json:"TimeInterval,omitnil,omitempty" name:"TimeInterval"`

	// Event
	EventContent *UebaEventContent `json:"EventContent,omitnil,omitempty" name:"EventContent"`

	// Alarm name
	AlertName *string `json:"AlertName,omitnil,omitempty" name:"AlertName"`

	// Alarm type
	// Prompt.
	// 1: low
	// 2: Medium risk
	// 3: High risk
	// 4: Critical
	AlterLevel *int64 `json:"AlterLevel,omitnil,omitempty" name:"AlterLevel"`

	// Operator.
	Operator []*string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Operation object.
	OperateObject []*string `json:"OperateObject,omitnil,omitempty" name:"OperateObject"`

	// Operation method
	OperateMethod []*string `json:"OperateMethod,omitnil,omitempty" name:"OperateMethod"`

	// Log type
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Chinese name in logs
	LogTypeStr *string `json:"LogTypeStr,omitnil,omitempty" name:"LogTypeStr"`
}

type UebaEventContent struct {
	// Event type
	// 1: Statement retrieval
	// 2: Filter search
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Statement retrieval content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// retrieval condition
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Statistical condition
	StatisticalFilter *StatisticalFilter `json:"StatisticalFilter,omitnil,omitempty" name:"StatisticalFilter"`
}

type UebaRule struct {
	// Policy ID
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Policy type
	// System policy
	// custom policy
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Policy level
	// Prompt.
	// 1: low
	// 2: Medium risk
	// 3: High risk
	// 4: Critical
	RuleLevel *int64 `json:"RuleLevel,omitnil,omitempty" name:"RuleLevel"`

	// Policy content
	RuleContent *string `json:"RuleContent,omitnil,omitempty" name:"RuleContent"`

	// Policy switch
	RuleStatus *bool `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// Number of hits
	HitCount *uint64 `json:"HitCount,omitnil,omitempty" name:"HitCount"`

	// Associated account Appid.
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Multi-account, member ID
	MemberID *string `json:"MemberID,omitnil,omitempty" name:"MemberID"`

	// Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Nickname
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// Custom rule specific content
	CustomRuleDetail *UebaCustomRule `json:"CustomRuleDetail,omitnil,omitempty" name:"CustomRuleDetail"`

	// Cloud type
	// 0 Tencent Cloud
	// aws:1
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

// Predefined struct for user
type UpdateAccessKeyAlarmStatusRequestParams struct {
	// Status  0: unprocessed 1: fixed 2: ignored.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Alarm ID list.
	AlarmIDList []*int64 `json:"AlarmIDList,omitnil,omitempty" name:"AlarmIDList"`

	// Risk ID list.
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`
}

type UpdateAccessKeyAlarmStatusRequest struct {
	*tchttp.BaseRequest
	
	// Status  0: unprocessed 1: fixed 2: ignored.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Alarm ID list.
	AlarmIDList []*int64 `json:"AlarmIDList,omitnil,omitempty" name:"AlarmIDList"`

	// Risk ID list.
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`
}

func (r *UpdateAccessKeyAlarmStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyAlarmStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "MemberId")
	delete(f, "AlarmIDList")
	delete(f, "RiskIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccessKeyAlarmStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyAlarmStatusResponseParams struct {
	// 0: Success; 1: Failure
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Error message
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccessKeyAlarmStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccessKeyAlarmStatusResponseParams `json:"Response"`
}

func (r *UpdateAccessKeyAlarmStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyAlarmStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyRemarkRequestParams struct {
	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Source IP name.
	SourceIPList []*string `json:"SourceIPList,omitnil,omitempty" name:"SourceIPList"`

	// ak name.
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// ID of the source IP.
	SourceIPIDList []*uint64 `json:"SourceIPIDList,omitnil,omitempty" name:"SourceIPIDList"`

	// AK ID.
	AccessKeyIDList []*uint64 `json:"AccessKeyIDList,omitnil,omitempty" name:"AccessKeyIDList"`
}

type UpdateAccessKeyRemarkRequest struct {
	*tchttp.BaseRequest
	
	// Remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Source IP name.
	SourceIPList []*string `json:"SourceIPList,omitnil,omitempty" name:"SourceIPList"`

	// ak name.
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// ID of the source IP.
	SourceIPIDList []*uint64 `json:"SourceIPIDList,omitnil,omitempty" name:"SourceIPIDList"`

	// AK ID.
	AccessKeyIDList []*uint64 `json:"AccessKeyIDList,omitnil,omitempty" name:"AccessKeyIDList"`
}

func (r *UpdateAccessKeyRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Remark")
	delete(f, "MemberId")
	delete(f, "SourceIPList")
	delete(f, "AccessKeyList")
	delete(f, "SourceIPIDList")
	delete(f, "AccessKeyIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccessKeyRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyRemarkResponseParams struct {
	// 0: success; 1: failure.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Error message
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccessKeyRemarkResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccessKeyRemarkResponseParams `json:"Response"`
}

func (r *UpdateAccessKeyRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertStatusListRequestParams struct {
	// Alarm ID list
	ID []*NewAlertKey `json:"ID,omitnil,omitempty" name:"ID"`

	// Operation type. 
	// 1: Revoke disposal 
	// 2: Marked with processed 
	// 3: Marked as ignored 
	// 4: Cancel tag disposal
	// 5: Unmark ignore
	OperateType *int64 `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Member ID of the Called Group Account
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`
}

type UpdateAlertStatusListRequest struct {
	*tchttp.BaseRequest
	
	// Alarm ID list
	ID []*NewAlertKey `json:"ID,omitnil,omitempty" name:"ID"`

	// Operation type. 
	// 1: Revoke disposal 
	// 2: Marked with processed 
	// 3: Marked as ignored 
	// 4: Cancel tag disposal
	// 5: Unmark ignore
	OperateType *int64 `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// Group Account Member ID
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Member ID of the Called Group Account
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`
}

func (r *UpdateAlertStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "OperateType")
	delete(f, "MemberId")
	delete(f, "OperatedMemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAlertStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertStatusListResponseParams struct {
	// Result Information
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Result Code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAlertStatusListResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAlertStatusListResponseParams `json:"Response"`
}

func (r *UpdateAlertStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserCallRecord struct {
	// Source IP of the call.
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// Invocation type.
	// 0: console invocation.
	// 1:API
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Number of calls.
	CallCount *int64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// Error code.
	// 0: Successful
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// First time call time.
	FirstCallTime *string `json:"FirstCallTime,omitnil,omitempty" name:"FirstCallTime"`

	// Call time.
	LastCallTime *string `json:"LastCallTime,omitnil,omitempty" name:"LastCallTime"`

	// Source IP of the call remark.
	SourceIPRemark *string `json:"SourceIPRemark,omitnil,omitempty" name:"SourceIPRemark"`

	// Source IP region of the call.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// User/Role name.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Aggregate date.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// appid
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Carrier.
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`
}

type VULBaseInfo struct {
	// Risk level.
	// High - high risk, middle - medium risk, low - low risk, info - Note.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Component.
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Release date.
	PublishTime *string `json:"PublishTime,omitnil,omitempty" name:"PublishTime"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Number of Affected Assets
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// Risk ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// Vulnerability type.
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Vulnerability Payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Vulnerability impact component.
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Technology reference.
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// Vulnerability impact version.
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// Risks.
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// User Nickname
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User appid.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Fixing suggestion
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Emergency Vulnerability Type. 1-Emergency Vulnerability; 0-Non-emergency Vulnerability.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// CVSS score
	// Note: This field may return null, indicating that no valid values can be obtained.
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// Attack intensity.
	// 0/1/2/3 
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttackHeat *int64 `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// Detection status 0 unscanned 1 scan in progress 2 scan complete.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// 1/0 whether compulsory.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSuggest *int64 `json:"IsSuggest,omitnil,omitempty" name:"IsSuggest"`

	// Tag.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VulTag []*string `json:"VulTag,omitnil,omitempty" name:"VulTag"`

	// Support products: "cfw_waf_virtual", "cwp_detect", "cwp_defense", "cwp_fix" (comma-separated).
	// Note: This field may return null, indicating that no valid values can be obtained.
	SupportProduct *string `json:"SupportProduct,omitnil,omitempty" name:"SupportProduct"`

	// Vulnerability detection task id.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Primary key
	// Note: This field may return null, indicating that no valid values can be obtained.
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Vulnerability id old version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PcmgrID *string `json:"PcmgrID,omitnil,omitempty" name:"PcmgrID"`

	// Vulnerability id new version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TvdID *string `json:"TvdID,omitnil,omitempty" name:"TvdID"`
}

type VULRiskAdvanceCFGList struct {
	// Risk ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// Risk level, low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Source of the check task
	CheckFrom *string `json:"CheckFrom,omitnil,omitempty" name:"CheckFrom"`

	// Enable/Disable. 1-Enable; 0-Disable.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Risk type.
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Affected versions
	ImpactVersion *string `json:"ImpactVersion,omitnil,omitempty" name:"ImpactVersion"`

	// CVE
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Vulnerability tag
	VULTag []*string `json:"VULTag,omitnil,omitempty" name:"VULTag"`

	// Repair method
	FixMethod []*string `json:"FixMethod,omitnil,omitempty" name:"FixMethod"`

	// Disclosure time
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// Emergency Vulnerability Type. 1-Emergency Vulnerability; 0-Non-emergency Vulnerability.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// Vulnerability description
	VULDescribe *string `json:"VULDescribe,omitnil,omitempty" name:"VULDescribe"`

	// Affected component
	ImpactComponent *string `json:"ImpactComponent,omitnil,omitempty" name:"ImpactComponent"`

	// Vulnerability Payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Technology reference
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// CVSS Score
	CVSS *string `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// Attack intensity
	AttackHeat *string `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// Security Product Support Status
	ServiceSupport []*ServiceSupport `json:"ServiceSupport,omitnil,omitempty" name:"ServiceSupport"`

	// Latest detection time
	RecentScanTime *string `json:"RecentScanTime,omitnil,omitempty" name:"RecentScanTime"`
}

type VULRiskInfo struct {
	// Fixing suggestion
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Technology reference/reference link.
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// Vulnerability description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Affected component.
	ImpactComponent []*VulImpactComponentInfo `json:"ImpactComponent,omitnil,omitempty" name:"ImpactComponent"`
}

type VULViewVULRisk struct {
	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Affected assets
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// Risk level, low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Number of Affected Assets
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// Risk ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Scan Source. See API Return Enumeration Type for details.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Vulnerability type
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// CVE number
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Vulnerability Payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Affected component
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Technology reference
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// Vulnerability Affected Version
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// risks
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Fixing suggestion
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Emergency Vulnerability Type. 1-Emergency Vulnerability; 0-Non-emergency Vulnerability.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`
}

type VULViewVULRiskData struct {
	// Port.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Impact assets.
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// Risk level: low - low risk, high - high risk, middle - medium risk, info - note, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Component.
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Latest Recognition Time
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First Recognition Time
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Number of Affected Assets
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// Risk ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// Scan Source. See API Return Enumeration Type for details.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Front-end Index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Vulnerability type.
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Vulnerability Payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Vulnerability impact component.
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Vulnerability impact version.
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// Risks.
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User appid.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Emergency Vulnerability Type. 1-Emergency Vulnerability; 0-Non-emergency Vulnerability.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// CVSS score
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// PCMGRId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`

	// Vulnerability tag. during searching, emergency mandatory parameter VulTag=SuggestRepair/EMGCVul.
	VulTag []*string `json:"VulTag,omitnil,omitempty" name:"VulTag"`

	// Vulnerability disclosure time.
	DisclosureTime *string `json:"DisclosureTime,omitnil,omitempty" name:"DisclosureTime"`

	// Attack intensity.
	AttackHeat *uint64 `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// Whether a mandatory vulnerability: 1 - yes; 0 - no.
	IsSuggest *int64 `json:"IsSuggest,omitnil,omitempty" name:"IsSuggest"`

	// Disposal task id.
	HandleTaskId *string `json:"HandleTaskId,omitnil,omitempty" name:"HandleTaskId"`

	// Engine source.
	EngineSource *string `json:"EngineSource,omitnil,omitempty" name:"EngineSource"`

	// New vulnerability risk id.
	VulRiskId *string `json:"VulRiskId,omitnil,omitempty" name:"VulRiskId"`

	// New version vulnerability id.
	TvdID *string `json:"TvdID,omitnil,omitempty" name:"TvdID"`

	// Is it possible to perform a one-click physical examination. valid values: 1-yes, 0-not allowed.
	IsOneClick *uint64 `json:"IsOneClick,omitnil,omitempty" name:"IsOneClick"`
}

type Vpc struct {
	// Subnet (32-bit mask)
	Subnet *uint64 `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// Connected VPC (32-bit mask)
	ConnectedVpc *uint64 `json:"ConnectedVpc,omitnil,omitempty" name:"ConnectedVpc"`

	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CVM (only 32-bit)
	CVM *uint64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// Tag.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// DNS Domain
	DNS []*string `json:"DNS,omitnil,omitempty" name:"DNS"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// CIDR block
	CIDR *string `json:"CIDR,omitnil,omitempty" name:"CIDR"`

	// Asset creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// User name
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// New Asset or Not. 1: New
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Whether it is a core asset. 1-Yes, 2-No.
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`
}

type VulImpactComponentInfo struct {
	// Component name
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Version name.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type VulRiskItem struct {
	// Cloud account ID.
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// Instance ID.
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`

	// Instance status
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Creation time.
	// 
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Vulnerability name
	VulName *string `json:"VulName,omitnil,omitempty" name:"VulName"`

	// Vulnerability type.
	VulCategory *string `json:"VulCategory,omitnil,omitempty" name:"VulCategory"`

	// Vulnerability level
	VulLevel *string `json:"VulLevel,omitnil,omitempty" name:"VulLevel"`

	// CVE id.
	CveID *string `json:"CveID,omitnil,omitempty" name:"CveID"`

	// Vulnerability description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Container ID.
	ContainerID *string `json:"ContainerID,omitnil,omitempty" name:"ContainerID"`

	// Vulnerability risk remediation recommendation.
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Linux vulnerability.
	VulCategoryName *string `json:"VulCategoryName,omitnil,omitempty" name:"VulCategoryName"`

	// Vulnerability level name.
	VulLevelName *string `json:"VulLevelName,omitnil,omitempty" name:"VulLevelName"`

	// Instance status chinese information.
	InstanceStatusName *string `json:"InstanceStatusName,omitnil,omitempty" name:"InstanceStatusName"`

	// Tenant ID.
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`
}

type VulTrend struct {
	// Number of affected assets.
	AffectAssetCount *int64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// Number of users affected.
	AffectUserCount *int64 `json:"AffectUserCount,omitnil,omitempty" name:"AffectUserCount"`

	// Number of attacks.
	AttackCount *int64 `json:"AttackCount,omitnil,omitempty" name:"AttackCount"`

	// Time
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type WebsiteRisk struct {
	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level, low - low risk, high - high risk, middle - medium risk, info - notification, extreme - critical.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Last detected
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID, use to handle risk
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Nickname
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// URL of the risk
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// URL of the risk file
	URLPath *string `json:"URLPath,omitnil,omitempty" name:"URLPath"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Check type.
	DetectEngine *string `json:"DetectEngine,omitnil,omitempty" name:"DetectEngine"`

	// Result description.
	ResultDescribe *string `json:"ResultDescribe,omitnil,omitempty" name:"ResultDescribe"`

	// Source URL
	SourceURL *string `json:"SourceURL,omitnil,omitempty" name:"SourceURL"`

	// Source file URL
	SourceURLPath *string `json:"SourceURLPath,omitnil,omitempty" name:"SourceURLPath"`
}

type WhereFilter struct {
	// Filter item
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Central platform definition:.
	// 1 equal 2 larger than 3 less than 4 greater than or equal to 5 less than or equal to 6 not equal to 9 fuzzy matching 13 non-fuzzy matching 14 bitwise and.
	// Exact match fills 7. fuzzy matching fills 9. 
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`
}