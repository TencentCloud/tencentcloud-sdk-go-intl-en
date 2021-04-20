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

package v20180724

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AlarmEvent struct {

	// Event name
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Event display name
	Description *string `json:"Description,omitempty" name:"Description"`

	// Alarm policy type
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type AlarmHistory struct {

	// Alarm record ID
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// Monitor type
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// Policy type
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Alarm object
	AlarmObject *string `json:"AlarmObject,omitempty" name:"AlarmObject"`

	// Alarm content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Timestamp of the first occurrence
	FirstOccurTime *int64 `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// Timestamp of the last occurrence
	LastOccurTime *int64 `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// Alarm status. Valid values: ALARM (not resolved), OK (resolved), NO_CONF (expired), NO_DATA (insufficient data)
	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// VPC of alarm object for basic product alarm
	VPC *string `json:"VPC,omitempty" name:"VPC"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Instance group of alarm object
	InstanceGroup []*InstanceGroups `json:"InstanceGroup,omitempty" name:"InstanceGroup" list`

	// Recipient list
	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids" list`

	// Recipient group list
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups" list`

	// Alarm channel list. Valid values: SMS (SMS), EMAIL (email), CALL (phone), WECHAT (WeChat)
	NoticeWays []*string `json:"NoticeWays,omitempty" name:"NoticeWays" list`

	// Alarm policy ID, which can be used when you call APIs ([BindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40421?from_cn_redirect=1), [UnBindingAllPolicyObject](https://intl.cloud.tencent.com/document/product/248/40568?from_cn_redirect=1), [UnBindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40567?from_cn_redirect=1)) to bind/unbind instances or instance groups to/from an alarm policy
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Alarm type
	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`

	// Event ID
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Whether the policy exists. Valid values: 0 (no), 1 (yes)
	PolicyExists *int64 `json:"PolicyExists,omitempty" name:"PolicyExists"`

	// Metric information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MetricsInfo []*AlarmHistoryMetric `json:"MetricsInfo,omitempty" name:"MetricsInfo" list`
}

type AlarmHistoryMetric struct {

	// Namespace used to query data by Tencent Cloud service monitoring type
	QceNamespace *string `json:"QceNamespace,omitempty" name:"QceNamespace"`

	// Metric name
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Statistical period
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Value triggering alarm
	Value *string `json:"Value,omitempty" name:"Value"`

	// Metric display name
	Description *string `json:"Description,omitempty" name:"Description"`
}

type AlarmNotice struct {

	// Alarm notification template ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Alarm notification template name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Last modified time
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// Last modified by
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`

	// Alarm notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`

	// User notification list
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices" list`

	// Callback notification list
	// Note: this field may return null, indicating that no valid values can be obtained.
	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices" list`

	// Whether it is the system default notification template. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsPreset *int64 `json:"IsPreset,omitempty" name:"IsPreset"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`

	// List of IDs of the alarm policies bound to alarm notification template
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds" list`
}

type AlarmPolicy struct {

	// Alarm policy ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Alarm policy name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring)
	// Note: this field may return null, indicating that no valid values can be obtained.
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// Status. Valid values: 0 (disabled), 1 (enabled)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Number of instances bound to policy group
	// Note: this field may return null, indicating that no valid values can be obtained.
	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`

	// Project ID. Valid values: -1 (no project), 0 (default project)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Project name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Alarm policy type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Trigger condition template ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConditionTemplateId *string `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// Event trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`

	// Notification rule ID list
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds" list`

	// Notification rule list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Notices []*AlarmNotice `json:"Notices,omitempty" name:"Notices" list`

	// Triggered task list
	// Note: this field may return null, indicating that no valid values can be obtained.
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks" list`

	// Template policy group
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConditionsTemp *ConditionsTemp `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

	// `Uin` of the last modifying user
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// Update time
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// Region
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region []*string `json:"Region,omitempty" name:"Region" list`

	// Namespace display name
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" name:"NamespaceShowName"`

	// Whether it is the default policy. Valid values: 1 (yes), 0 (no)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// Whether the default policy can be set. Valid values: 1 (yes), 0 (no)
	// Note: this field may return null, indicating that no valid values can be obtained.
	CanSetDefault *int64 `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// Instance group ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Total number of instances in instance group
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// Instance group name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`

	// Trigger condition type. Valid values: STATIC (static threshold), DYNAMIC (dynamic)
	// Note: this field may return null, indicating that no valid values can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Policy ID for instance/instance group binding and unbinding APIs (BindingPolicyObject, UnBindingAllPolicyObject, UnBindingPolicyObject)
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// Tag
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagInstances []*TagInstance `json:"TagInstances,omitempty" name:"TagInstances" list`
}

type AlarmPolicyCondition struct {

	// Metric trigger condition operator. Valid values: 0 (OR), 1 (AND)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// Alarm trigger condition list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rules []*AlarmPolicyRule `json:"Rules,omitempty" name:"Rules" list`
}

type AlarmPolicyEventCondition struct {

	// Alarm trigger condition list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rules []*AlarmPolicyRule `json:"Rules,omitempty" name:"Rules" list`
}

type AlarmPolicyFilter struct {

	// Filter condition type. Valid values: DIMENSION (uses dimensions for filtering)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// JSON string generated by serializing the `AlarmPolicyDimension` two-dimensional array. The one-dimensional arrays are in OR relationship, and the elements in a one-dimensional array are in AND relationship
	// Note: this field may return null, indicating that no valid values can be obtained.
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type AlarmPolicyRule struct {

	// Metric name or event name. The supported metrics can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1) and the supported events via [DescribeAlarmEvents](https://intl.cloud.tencent.com/document/product/248/51284?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Statistical period in seconds. The valid values can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Operator
	// intelligent = intelligent detection without threshold
	// eq = equal to
	// ge = greater than or equal to
	// gt = greater than
	// le = less than or equal to
	// lt = less than
	// ne = not equal to
	// day_increase = day-on-day increase
	// day_decrease = day-on-day decrease
	// day_wave = day-on-day fluctuation
	// week_increase = week-on-week increase
	// week_decrease = week-on-week decrease
	// week_wave = week-on-week fluctuation
	// cycle_increase = cyclical increase
	// cycle_decrease = cyclical decrease
	// cycle_wave = cyclical fluctuation
	// re = regex match
	// The valid values can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Threshold. The valid value range can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Number of periods. `1`: continue for one period; `2`: continue for two periods; and so on. The valid values can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// Alarm interval in seconds. Valid values: 0 (do not repeat), 300 (alarm once every 5 minutes), 600 (alarm once every 10 minutes), 900 (alarm once every 15 minutes), 1800 (alarm once every 30 minutes), 3600 (alarm once every hour), 7200 (alarm once every 2 hours), 10800 (alarm once every 3 hours), 21600 (alarm once every 6 hours),  43200 (alarm once every 12 hours), 86400 (alarm once every day)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeFrequency *int64 `json:"NoticeFrequency,omitempty" name:"NoticeFrequency"`

	// Whether the alarm frequency increases exponentially. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsPowerNotice *int64 `json:"IsPowerNotice,omitempty" name:"IsPowerNotice"`

	// Filter condition for one single trigger rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	Filter *AlarmPolicyFilter `json:"Filter,omitempty" name:"Filter"`

	// Metric display name, which is used in the output parameter
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Unit, which is used in the output parameter
	// Note: this field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Trigger condition type. `STATIC`: static threshold; `dynamic`: dynamic threshold. If you do not specify this parameter when creating or editing a policy, `STATIC` is used by default.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

type AlarmPolicyTriggerTask struct {

	// Triggered task type. Valid value: AS (auto scaling)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Configuration information in JSON format, such as {"Key1":"Value1","Key2":"Value2"}
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskConfig *string `json:"TaskConfig,omitempty" name:"TaskConfig"`
}

type BindingPolicyObjectDimension struct {

	// Region name.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region ID.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Instance dimension information in the following format:
	// {"unInstanceId":"ins-00jvv9mo"}. The dimension information varies by Tencent Cloud services. For more information, please see:
	// [Dimension List](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1)
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// Event dimensions.
	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type BindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// Policy group ID. If `PolicyId` is used, this parameter will be ignored, and any value, e.g., 0, can be passed in.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Required. The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Dimensions of an object to be bound.
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// Alarm policy ID. If this parameter is used, `GroupId` will be ignored.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *BindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindingPolicyObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindingPolicyObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommonNamespace struct {

	// Namespace ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Namespace name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Namespace value
	Value *string `json:"Value,omitempty" name:"Value"`

	// Product name
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// Configuration information
	Config *string `json:"Config,omitempty" name:"Config"`

	// List of supported regions
	AvailableRegions []*string `json:"AvailableRegions,omitempty" name:"AvailableRegions" list`

	// Sort ID
	SortId *int64 `json:"SortId,omitempty" name:"SortId"`

	// Unique ID in Dashboard
	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

type ConditionsTemp struct {

	// Template name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Metric trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// Event trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Notification template name, which can contain up to 60 characters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`

	// User notifications (up to 5)
	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices" list`

	// Callback notifications (up to 3)
	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices" list`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm notification template ID
		NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// Value fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy name, which can contain up to 20 characters
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring)
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// Type of alarm policy, which can be obtained via [DescribeAllNamespaces](https://intl.cloud.tencent.com/document/product/248/48683?from_cn_redirect=1). An example value is `cvm_device`.
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Remarks with up to 100 letters, digits, underscores, and hyphens
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether to enable. Valid values: 0 (no), 1 (yes). Default value: 1. This parameter can be left empty
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Project ID. For products with different projects, a value other than `-1` must be passed in. `-1`: no project; `0`: default project. If no value is passed in, `-1` will be used. The supported project IDs can be viewed on the [**Account Center** > **Project Management**](https://console.cloud.tencent.com/project) page of the console.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID of trigger condition template. This parameter can be left empty.
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition. The supported metrics can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// Event trigger condition. The supported events can be queried via [DescribeAlarmEvents](https://intl.cloud.tencent.com/document/product/248/51284?from_cn_redirect=1).
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`

	// List of notification rule IDs, which can be obtained via [DescribeAlarmNotices](https://intl.cloud.tencent.com/document/product/248/51280?from_cn_redirect=1)
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds" list`

	// Triggered task list
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks" list`
}

func (r *CreateAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm policy ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// Alarm policy ID, which can be used when you call APIs ([BindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40421?from_cn_redirect=1), [UnBindingAllPolicyObject](https://intl.cloud.tencent.com/document/product/248/40568?from_cn_redirect=1), [UnBindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40567?from_cn_redirect=1)) to bind/unbind instances or instance groups to/from an alarm policy
		OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupCondition struct {

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to. This parameter is optional if a default comparative type is configured for the metric.
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// Comparative value. This parameter is optional if the metric has no requirement.
	CalcValue *float64 `json:"CalcValue,omitempty" name:"CalcValue"`

	// Data aggregation period in seconds. This parameter is optional if the metric has a default value.
	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// Number of consecutive periods after which an alarm will be triggered.
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// If a metric is created based on a template, the RuleId of the metric in the template must be passed in.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupEventCondition struct {

	// Alarm event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// If a metric is created based on a template, the RuleId of the metric in the template must be passed in.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// Policy group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Name of the view to which the policy group belongs. If the policy group is created based on a template, this parameter is optional.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// ID of the project to which the policy group belongs, which will be used for authentication.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID of a template-based policy group. This parameter is required only when the policy group is created based on a template.
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`

	// Whether the policy group is shielded. The value 0 indicates that the policy group is not shielded. The value 1 indicates that the policy group is shielded. The default value is 0.
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// Remarks of the policy group.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Insertion time in the format of Unix timestamp. If this parameter is not configured, the backend processing time is used.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// Alarm threshold rules in the policy group.
	Conditions []*CreatePolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// Event alarm rules in the policy group.
	EventConditions []*CreatePolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// Whether it is a backend call. If the value is 1, rules from the policy template will be used to fill in the `Conditions` and `EventConditions` fields.
	BackEndCall *int64 `json:"BackEndCall,omitempty" name:"BackEndCall"`

	// The 'AND' and 'OR' rules for alarm metrics. The value 0 indicates 'OR', which means that an alarm will be triggered when any rule is met. The value 1 indicates 'AND', which means that an alarm will be triggered only when all rules are met.
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

func (r *CreatePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the created policy group.
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataPoint struct {

	// Combination of instance object dimensions
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// The array of timestamps indicating at which points in time there is data. Missing timestamps have no data points (i.e., missed)
	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps" list`

	// The array of monitoring values, which is in one-to-one correspondence to Timestamps
	Values []*float64 `json:"Values,omitempty" name:"Values" list`
}

type DeleteAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm notification template ID list
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds" list`
}

func (r *DeleteAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmNoticesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmNoticesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID list
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds" list`
}

func (r *DeleteAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy group ID.
	GroupId []*int64 `json:"GroupId,omitempty" name:"GroupId" list`
}

func (r *DeletePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListAlarms struct {

	// Event type.
	// Note: This field may return null, indicating that no valid value was found.
	BusinessTypeDesc *string `json:"BusinessTypeDesc,omitempty" name:"BusinessTypeDesc"`

	// Event type.
	// Note: This field may return null, indicating that no valid value was found.
	AccidentTypeDesc *string `json:"AccidentTypeDesc,omitempty" name:"AccidentTypeDesc"`

	// ID of the event type. The value 1 indicates service issues. The value 2 indicates other subscriptions.
	// Note: This field may return null, indicating that no valid value was found.
	BusinessID *int64 `json:"BusinessID,omitempty" name:"BusinessID"`

	// Event status ID. The value 0 indicates that the event has been recovered. The value 1 indicates that the event has not been recovered.
	// Note: This field may return null, indicating that no valid value was found.
	EventStatus *int64 `json:"EventStatus,omitempty" name:"EventStatus"`

	// Affected object.
	// Note: This field may return null, indicating that no valid value was found.
	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`

	// Region where the event occurs.
	// Note: This field may return null, indicating that no valid value was found.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Time when the event occurs.
	// Note: This field may return null, indicating that no valid value was found.
	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid value was found.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAccidentEventListRequest struct {
	*tchttp.BaseRequest

	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Sorting rule by UpdateTime. Valid values: asc and desc.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// Sorting rule by OccurTime. Valid values: asc or desc. Sorting by UpdateTimeOrder takes priority.
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// Filter by event type. The value 1 indicates service issues. The value 2 indicates other subscriptions.
	AccidentType []*int64 `json:"AccidentType,omitempty" name:"AccidentType" list`

	// Filter by event. The value 1 indicates CVM storage issues. The value 2 indicates CVM network connection issues. The value 3 indicates that the CVM has an exception. The value 202 indicates that an ISP network jitter occurs.
	AccidentEvent []*int64 `json:"AccidentEvent,omitempty" name:"AccidentEvent" list`

	// Filter by event status. The value 0 indicates that the event has been recovered. The value 1 indicates that the event has not been recovered.
	AccidentStatus []*int64 `json:"AccidentStatus,omitempty" name:"AccidentStatus" list`

	// Filter by region where the event occurs. The value gz indicates Guangzhou. The value sh indicates Shanghai.
	AccidentRegion []*string `json:"AccidentRegion,omitempty" name:"AccidentRegion" list`

	// Filter by affected resource, such as ins-19a06bka.
	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`
}

func (r *DescribeAccidentEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccidentEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Platform event list.
	// Note: This field may return null, indicating that no valid value was found.
		Alarms []*DescribeAccidentEventListAlarms `json:"Alarms,omitempty" name:"Alarms" list`

		// Total number of platform events.
	// Note: This field may return null, indicating that no valid value was found.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccidentEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccidentEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmEventsRequest struct {
	*tchttp.BaseRequest

	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy type such as cvm_device, which is obtained through the `DescribeAllNamespaces` API
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeAlarmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm event list
		Events []*AlarmEvent `json:"Events,omitempty" name:"Events" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmHistoriesRequest struct {
	*tchttp.BaseRequest

	// Value fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Page number starting from 1. Default value: 1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–100. Default value: 20
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Sort by the first occurrence time in descending order by default. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitempty" name:"Order"`

	// Start time, which corresponds to `FirstOccurTime` (time when an alarm first occurred). The default value is the timestamp of a day earlier.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which corresponds to `FirstOccurTime` (time when an alarm first occurred). The default value is the current timestamp.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes" list`

	// Filter by alarm object. Fuzzy search with string is supported
	AlarmObject *string `json:"AlarmObject,omitempty" name:"AlarmObject"`

	// Filter by alarm status. Valid values: ALARM (not resolved), OK (resolved), NO_CONF (expired), NO_DATA (insufficient data). If this parameter is left empty, all will be queried by default
	AlarmStatus []*string `json:"AlarmStatus,omitempty" name:"AlarmStatus" list`

	// Filter by project ID. Valid values: -1 (no project), 0 (default project)
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// Filter by instance group ID
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitempty" name:"InstanceGroupIds" list`

	// Filter by policy type
	Namespaces []*MonitorTypeNamespace `json:"Namespaces,omitempty" name:"Namespaces" list`

	// Filter by metric name
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`

	// Fuzzy search by policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Fuzzy search by alarm content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Search by recipient
	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids" list`

	// Search by recipient group
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups" list`

	// Search by alarm policy ID list
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds" list`
}

func (r *DescribeAlarmHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmHistoriesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Alarm record list
		Histories []*AlarmHistory `json:"Histories,omitempty" name:"Histories" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmHistoriesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmMetricsRequest struct {
	*tchttp.BaseRequest

	// Value fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Monitor type filter. Valid values: MT_QCE (Tencent Cloud service monitoring)
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// Alarm policy type such as cvm_device, which is obtained through the `DescribeAllNamespaces` API
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeAlarmMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm metric list
		Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeCallbacksRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmNoticeCallbacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmNoticeCallbacksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeCallbacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm callback notification
	// Note: this field may return null, indicating that no valid values can be obtained.
		URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticeCallbacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmNoticeCallbacksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm notification template ID
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`
}

func (r *DescribeAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmNoticeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm notification template details
		Notice *AlarmNotice `json:"Notice,omitempty" name:"Notice"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmNoticeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Page number. Minimum value: 1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Sort by update time. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitempty" name:"Order"`

	// Root account `uid`, which is used to create preset notifications
	OwnerUid *int64 `json:"OwnerUid,omitempty" name:"OwnerUid"`

	// Alarm notification template name, which is used for fuzzy search
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter by recipient. The type of notified users should be selected for the alarm notification template. Valid values: USER (user), GROUP (user group). If this parameter is left empty, no filter by recipient will be performed
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// Recipient object list
	UserIds []*int64 `json:"UserIds,omitempty" name:"UserIds" list`

	// Recipient group list
	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of alarm notification templates
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Alarm notification template list
		Notices []*AlarmNotice `json:"Notices,omitempty" name:"Notices" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// Value fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Page number starting from 1. Default value: 1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–100. Default value: 20
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Fuzzy search by policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes" list`

	// Filter by namespace. For the values of different policy types, please see:
	// [Policy Type List](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1)
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces" list`

	// Alarm object list. The outer array corresponds to multiple instances.
	// Each inner array corresponds to one instance, where `object` corresponds to the dimension information of the instance. The format is as follows:
	// [
	// 	[{"name":"unInstanceId","value":"ins-qr888845g"}],
	// 	[{"name":"unInstanceId","value":"ins-qr8d555g"}]
	// 	...
	// ]
	// For the samples for different Tencent Cloud services, please see:
	// [Dimension List](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1)
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// Search by recipient `uid`, which should be queried by calling the CAM API. For more information, please see:
	// [ListUsers](https://intl.cloud.tencent.com/document/product/598/34587?from_cn_redirect=1)
	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids" list`

	// Search by recipient group `uid`, which should be queried by calling the CAM API. For more information, please see:
	// [ListGroups](https://intl.cloud.tencent.com/document/product/598/34589?from_cn_redirect=1)
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups" list`

	// Filter by default policy. Valid values: DEFAULT (display default policy), NOT_DEFAULT (display non-default policies). If this parameter is left empty, all policies will be displayed
	PolicyType []*string `json:"PolicyType,omitempty" name:"PolicyType" list`

	// Sort by field. For example, to sort by the last modification time, use Field: "UpdateTime".
	Field *string `json:"Field,omitempty" name:"Field"`

	// Sort order. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitempty" name:"Order"`

	// ID array of the policy project, which can be viewed on the following page:
	// [Project Management](https://console.cloud.tencent.com/project)
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// ID list of the notification template, which can be obtained by querying the notification template list.
	// [DescribeAlarmNotices](https://intl.cloud.tencent.com/document/product/248/51280?from_cn_redirect=1)
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds" list`

	// Filter by trigger condition. Valid values: STATIC (display policies with static threshold), DYNAMIC (display policies with dynamic threshold). If this parameter is left empty, all policies will be displayed
	RuleTypes []*string `json:"RuleTypes,omitempty" name:"RuleTypes" list`

	// Filter by alarm status. Valid values: [1]: enabled; [0]: disabled; [0, 1]: all
	Enable []*int64 `json:"Enable,omitempty" name:"Enable" list`

	// If `1` is passed in, alarm policies with no notification rules configured are queried. If it is left empty or other values are passed in, all alarm policies are queried.
	NotBindingNoticeRule *int64 `json:"NotBindingNoticeRule,omitempty" name:"NotBindingNoticeRule"`
}

func (r *DescribeAlarmPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of policies
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Policy array
		Policies []*AlarmPolicy `json:"Policies,omitempty" name:"Policies" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// Value fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy details
		Policy *AlarmPolicy `json:"Policy,omitempty" name:"Policy"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllNamespacesRequest struct {
	*tchttp.BaseRequest

	// Filter by use case. Valid values: ST_DASHBOARD (Dashboard type), ST_ALARM (alarm type)
	SceneType *string `json:"SceneType,omitempty" name:"SceneType"`

	// Value fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes" list`

	// Filter by namespace ID. If this parameter is left empty, all will be queried
	Ids []*string `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DescribeAllNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllNamespacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm policy type of Tencent Cloud service (disused)
		QceNamespaces *CommonNamespace `json:"QceNamespaces,omitempty" name:"QceNamespaces"`

		// Other alarm policy type (disused)
		CustomNamespaces *CommonNamespace `json:"CustomNamespaces,omitempty" name:"CustomNamespaces"`

		// Alarm policy type of Tencent Cloud service
		QceNamespacesNew []*CommonNamespace `json:"QceNamespacesNew,omitempty" name:"QceNamespacesNew" list`

		// Other alarm policy type (not supported currently)
		CustomNamespacesNew []*CommonNamespace `json:"CustomNamespacesNew,omitempty" name:"CustomNamespacesNew" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllNamespacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest

	// Service namespace. Tencent Cloud services have different namespaces. For more information on service namespaces, see the monitoring metric documentation of each service. For example, see [CVM Monitoring Metrics](https://intl.cloud.tencent.com/document/product/248/6843?from_cn_redirect=1) for the namespace of CVM
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric name. Tencent Cloud services have different metric names. For more information on metric names, see the monitoring metric documentation of each service. For example, see [CVM Monitoring Metrics](https://intl.cloud.tencent.com/document/product/248/6843?from_cn_redirect=1) for the metric names of CVM
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeBaseMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Listed of queried metric descriptions
		MetricSet []*MetricSet `json:"MetricSet,omitempty" name:"MetricSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListAlarms struct {

	// Alarm ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Project ID.
	// Note: This field may return null, indicating that no valid value was found.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Project name.
	// Note: This field may return null, indicating that no valid value was found.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Alarm status ID. Valid values: 0 (not resolved), 1 (resolved), 2/3/5 (insufficient data), 4 (expired)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Alarm status. Valid values: ALARM (not resolved), OK (resolved), NO_DATA (insufficient data), NO_CONF (expired)
	// Note: this field may return null, indicating that no valid values can be obtained.
	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// Policy group ID.
	// Note: This field may return null, indicating that no valid value was found.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Policy group name.
	// Note: This field may return null, indicating that no valid value was found.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Occurrence time.
	// Note: This field may return null, indicating that no valid value was found.
	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// Duration in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// End time.
	// Note: This field may return null, indicating that no valid value was found.
	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// Alarm content.
	// Note: This field may return null, indicating that no valid value was found.
	Content *string `json:"Content,omitempty" name:"Content"`

	// Alarm object.
	// Note: This field may return null, indicating that no valid value was found.
	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`

	// Alarm object ID.
	// Note: This field may return null, indicating that no valid value was found.
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// Policy type.
	// Note: This field may return null, indicating that no valid value was found.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// VPC, which is unique to CVM.
	// Note: This field may return null, indicating that no valid value was found.
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// Metric ID.
	// Note: This field may return null, indicating that no valid value was found.
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// Metric name.
	// Note: This field may return null, indicating that no valid value was found.
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Alarm type. The value 0 indicates metric alarms. The value 2 indicates product event alarms. The value 3 indicates platform event alarms.
	// Note: This field may return null, indicating that no valid value was found.
	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// Region.
	// Note: This field may return null, indicating that no valid value was found.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Dimensions of an alarm object.
	// Note: This field may return null, indicating that no valid value was found.
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// Notification method.
	// Note: This field may return null, indicating that no valid value was found.
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// Instance group information.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroup []*InstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup" list`
}

type DescribeBasicAlarmListRequest struct {
	*tchttp.BaseRequest

	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Sorting by occurrence time. Valid values: asc and desc.
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// Filter by project ID.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// Filter by policy type.
	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames" list`

	// Filter by alarm status.
	AlarmStatus []*int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus" list`

	// Filter by alarm object.
	ObjLike *string `json:"ObjLike,omitempty" name:"ObjLike"`

	// Filter by instance group ID.
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitempty" name:"InstanceGroupIds" list`

	// Filtering by metric names
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`
}

func (r *DescribeBasicAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBasicAlarmListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm list.
	// Note: This field may return null, indicating that no valid value was found.
		Alarms []*DescribeBasicAlarmListAlarms `json:"Alarms,omitempty" name:"Alarms" list`

		// Total number.
	// Note: This field may return null, indicating that no valid value was found.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBasicAlarmListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListDimension struct {

	// Region ID.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region abbreviation.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Combined JSON string of dimensions.
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// Combined JSON string of event dimensions.
	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type DescribeBindingPolicyObjectListInstance struct {

	// Unique ID of the object.
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// Dimension set of an object instance, which is a jsonObj string.
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// Whether the object is shielded. The value 0 indicates that the object is not shielded. The value 1 indicates that the object is shielded.
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// Region where the object resides.
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeBindingPolicyObjectListInstanceGroup struct {

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Alarm policy type name.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// Instance group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Number of instances.
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// Update time.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// Regions where the instances reside.
	// Note: This field may return null, indicating that no valid value was found.
	Regions []*string `json:"Regions,omitempty" name:"Regions" list`
}

type DescribeBindingPolicyObjectListRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Dimensions of filtering objects.
	Dimensions []*DescribeBindingPolicyObjectListDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

func (r *DescribeBindingPolicyObjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBindingPolicyObjectListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of bound object instances.
	// Note: This field may return null, indicating that no valid value was found.
		List []*DescribeBindingPolicyObjectListInstance `json:"List,omitempty" name:"List" list`

		// Total number of bound object instances.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Number of object instances that are not shielded.
		NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

		// Bound instance group information. This parameter is not configured if no instance group is bound.
	// Note: This field may return null, indicating that no valid value was found.
		InstanceGroup *DescribeBindingPolicyObjectListInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindingPolicyObjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBindingPolicyObjectListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorTypesRequest struct {
	*tchttp.BaseRequest

	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeMonitorTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring)
		MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorTypesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListCondition struct {

	// Policy view name.
	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`

	// Event alarm conditions.
	// Note: This field may return null, indicating that no valid value was found.
	EventMetrics []*DescribePolicyConditionListEventMetric `json:"EventMetrics,omitempty" name:"EventMetrics" list`

	// Whether to support multiple regions.
	IsSupportMultiRegion *bool `json:"IsSupportMultiRegion,omitempty" name:"IsSupportMultiRegion"`

	// Metric alarm conditions.
	// Note: This field may return null, indicating that no valid value was found.
	Metrics []*DescribePolicyConditionListMetric `json:"Metrics,omitempty" name:"Metrics" list`

	// Policy type name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Sorting ID.
	SortId *int64 `json:"SortId,omitempty" name:"SortId"`

	// Whether to support default policies.
	SupportDefault *bool `json:"SupportDefault,omitempty" name:"SupportDefault"`

	// List of regions that support this policy type.
	// Note: This field may return null, indicating that no valid value was found.
	SupportRegions []*string `json:"SupportRegions,omitempty" name:"SupportRegions" list`
}

type DescribePolicyConditionListConfigManual struct {

	// Check method.
	// Note: This field may return null, indicating that no valid value was found.
	CalcType *DescribePolicyConditionListConfigManualCalcType `json:"CalcType,omitempty" name:"CalcType"`

	// Threshold.
	// Note: This field may return null, indicating that no valid value was found.
	CalcValue *DescribePolicyConditionListConfigManualCalcValue `json:"CalcValue,omitempty" name:"CalcValue"`

	// Duration.
	// Note: This field may return null, indicating that no valid value was found.
	ContinueTime *DescribePolicyConditionListConfigManualContinueTime `json:"ContinueTime,omitempty" name:"ContinueTime"`

	// Data period.
	// Note: This field may return null, indicating that no valid value was found.
	Period *DescribePolicyConditionListConfigManualPeriod `json:"Period,omitempty" name:"Period"`

	// Number of periods.
	// Note: This field may return null, indicating that no valid value was found.
	PeriodNum *DescribePolicyConditionListConfigManualPeriodNum `json:"PeriodNum,omitempty" name:"PeriodNum"`

	// Statistics method.
	// Note: This field may return null, indicating that no valid value was found.
	StatType *DescribePolicyConditionListConfigManualStatType `json:"StatType,omitempty" name:"StatType"`
}

type DescribePolicyConditionListConfigManualCalcType struct {

	// Value of CalcType.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// Required or not.
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualCalcValue struct {

	// Default value.
	// Note: This field may return null, indicating that no valid value was found.
	Default *string `json:"Default,omitempty" name:"Default"`

	// Fixed value.
	// Note: This field may return null, indicating that no valid value was found.
	Fixed *string `json:"Fixed,omitempty" name:"Fixed"`

	// Maximum value.
	// Note: This field may return null, indicating that no valid value was found.
	Max *string `json:"Max,omitempty" name:"Max"`

	// Minimum value.
	// Note: This field may return null, indicating that no valid value was found.
	Min *string `json:"Min,omitempty" name:"Min"`

	// Required or not.
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualContinueTime struct {

	// Default duration in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// Custom durations in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// Required or not.
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriod struct {

	// Default period in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// Custom periods in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// Required or not.
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriodNum struct {

	// Number of default periods.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// Number of custom periods.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// Required or not.
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualStatType struct {

	// Data aggregation method in a period of 5 seconds.
	// Note: This field may return null, indicating that no valid value was found.
	P5 *string `json:"P5,omitempty" name:"P5"`

	// Data aggregation method in a period of 10 seconds.
	// Note: This field may return null, indicating that no valid value was found.
	P10 *string `json:"P10,omitempty" name:"P10"`

	// Data aggregation method in a period of 1 minute.
	// Note: This field may return null, indicating that no valid value was found.
	P60 *string `json:"P60,omitempty" name:"P60"`

	// Data aggregation method in a period of 5 minutes.
	// Note: This field may return null, indicating that no valid value was found.
	P300 *string `json:"P300,omitempty" name:"P300"`

	// Data aggregation method in a period of 10 minutes.
	// Note: This field may return null, indicating that no valid value was found.
	P600 *string `json:"P600,omitempty" name:"P600"`

	// Data aggregation method in a period of 30 minutes.
	// Note: This field may return null, indicating that no valid value was found.
	P1800 *string `json:"P1800,omitempty" name:"P1800"`

	// Data aggregation method in a period of 1 hour.
	// Note: This field may return null, indicating that no valid value was found.
	P3600 *string `json:"P3600,omitempty" name:"P3600"`

	// Data aggregation method in a period of 1 day.
	// Note: This field may return null, indicating that no valid value was found.
	P86400 *string `json:"P86400,omitempty" name:"P86400"`
}

type DescribePolicyConditionListEventMetric struct {

	// Event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Event name.
	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`

	// Whether to recover.
	NeedRecovered *bool `json:"NeedRecovered,omitempty" name:"NeedRecovered"`

	// Event type, which is a reserved field. Currently, it is fixed to 2.
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribePolicyConditionListMetric struct {

	// Metric configuration.
	// Note: This field may return null, indicating that no valid value was found.
	ConfigManual *DescribePolicyConditionListConfigManual `json:"ConfigManual,omitempty" name:"ConfigManual"`

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// Metric name.
	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`

	// Metric unit.
	MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`
}

type DescribePolicyConditionListRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribePolicyConditionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyConditionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of alarm policy conditions.
		Conditions []*DescribePolicyConditionListCondition `json:"Conditions,omitempty" name:"Conditions" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyConditionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyConditionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoCallback struct {

	// URL of the user callback API.
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// Status of the user callback API. The value 0 indicates that the API is not verified. The value 1 indicates that the API is verified. The value 2 indicates that a URL exists but the API fails to be verified.
	ValidFlag *int64 `json:"ValidFlag,omitempty" name:"ValidFlag"`

	// Verification code of the user callback API.
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type DescribePolicyGroupInfoCondition struct {

	// Metric name.
	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`

	// Data aggregation period in seconds.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// Threshold rule ID.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Metric unit.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. If the value is less than 0, no alarm will be triggered. If the value is 0, an alarm will be triggered only once. If the value is greater than 0, an alarm will be triggered at the interval of `triggerTime`.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to. The value 7 indicates day-on-day increase. The value 8 indicates day-on-day decrease. The value 9 indicates week-on-week increase. The value 10 indicates week-on-week decrease. The value 11 indicates periodical increase. The value 12 indicates periodical decrease.
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// Threshold.
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// Duration at which an alarm will be triggered in seconds.
	ContinueTime *int64 `json:"ContinueTime,omitempty" name:"ContinueTime"`

	// Alarm metric name.
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type DescribePolicyGroupInfoConditionTpl struct {

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Policy group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Policy type.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// Policy group remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// Update time.
	// Note: This field may return null, indicating that no valid value was found.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time.
	// Note: This field may return null, indicating that no valid value was found.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// Whether the 'AND' rule is used.
	// Note: This field may return null, indicating that no valid value was found.
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupInfoEventCondition struct {

	// Event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Event alarm rule ID.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Event name.
	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
}

type DescribePolicyGroupInfoReceiverInfo struct {

	// List of alarm recipient group IDs.
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList" list`

	// List of alarm recipient IDs.
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`

	// Start time of the alarm period. Value range: [0,86400). Convert the Unix timestamp to Beijing time and then remove the date. For example, 7200 indicates '10:0:0'.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the alarm period. The meaning is the same as that of StartTime.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Recipient type. Valid values: group and user.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// Alarm notification method. Valid values: "SMS", "SITE", "EMAIL", "CALL", and "WECHAT".
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// Uid of the alarm call recipient.
	// Note: This field may return null, indicating that no valid value was found.
	UidList []*int64 `json:"UidList,omitempty" name:"UidList" list`

	// Number of alarm call rounds.
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// Intervals of alarm call rounds in seconds.
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// Alarm call intervals for individuals in seconds.
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// Whether to send an alarm call delivery notice. The value 0 indicates that no notice needs to be sent. The value 1 indicates that a notice needs to be sent.
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// Alarm call notification time. Valid values: OCCUR (indicating that a notice is sent when the alarm is triggered) and RECOVER (indicating that a notice is sent when the alarm is recovered).
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor" list`

	// Notification method when an alarm is recovered. Valid value: SMS.
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify" list`

	// Alarm language.
	// Note: This field may return null, indicating that no valid value was found.
	ReceiveLanguage *string `json:"ReceiveLanguage,omitempty" name:"ReceiveLanguage"`
}

type DescribePolicyGroupInfoRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribePolicyGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy group name.
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// ID of the project to which the policy group belongs.
		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// Whether it is the default policy. The value 0 indicates that it is not the default policy. The value 1 indicates that it is the default policy.
		IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

		// Policy type.
		ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

		// Policy description
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// Policy type name.
		ShowName *string `json:"ShowName,omitempty" name:"ShowName"`

		// Uin that was last edited.
		LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

		// Last edited time.
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// Regions supported by this policy.
		Region []*string `json:"Region,omitempty" name:"Region" list`

		// List of policy type dimensions.
		DimensionGroup []*string `json:"DimensionGroup,omitempty" name:"DimensionGroup" list`

		// Threshold rule list.
	// Note: This field may return null, indicating that no valid value was found.
		ConditionsConfig []*DescribePolicyGroupInfoCondition `json:"ConditionsConfig,omitempty" name:"ConditionsConfig" list`

		// Product event rule list.
	// Note: This field may return null, indicating that no valid value was found.
		EventConfig []*DescribePolicyGroupInfoEventCondition `json:"EventConfig,omitempty" name:"EventConfig" list`

		// Recipient list.
	// Note: This field may return null, indicating that no valid value was found.
		ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos" list`

		// User callback information.
	// Note: This field may return null, indicating that no valid value was found.
		Callback *DescribePolicyGroupInfoCallback `json:"Callback,omitempty" name:"Callback"`

		// Template-based policy group.
	// Note: This field may return null, indicating that no valid value was found.
		ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

		// Whether the policy can be configured as the default policy.
		CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

		// Whether the 'AND' rule is used.
	// Note: This field may return null, indicating that no valid value was found.
		IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListGroup struct {

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Policy group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Whether it is enabled.
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`

	// Policy view name.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// Last modified time.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// Number of instances that are bound to the policy group.
	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`

	// Number of unshielded instances that are bound to the policy group.
	NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

	// Whether it is the default policy. The value 0 indicates that it is not the default policy. The value 1 indicates that it is the default policy.
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// Whether the policy can be configured as the default policy.
	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// Parent policy group ID.
	ParentGroupId *int64 `json:"ParentGroupId,omitempty" name:"ParentGroupId"`

	// Remarks of the policy group.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// ID of the project to which the policy group belongs.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Threshold rule list.
	// Note: This field may return null, indicating that no valid value was found.
	Conditions []*DescribePolicyGroupInfoCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// Product event rule list.
	// Note: This field may return null, indicating that no valid value was found.
	EventConditions []*DescribePolicyGroupInfoEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// Recipient list.
	// Note: This field may return null, indicating that no valid value was found.
	ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos" list`

	// Template-based policy group.
	// Note: This field may return null, indicating that no valid value was found.
	ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

	// Instance group that is bound to the policy group.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroup *DescribePolicyGroupListGroupInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

	// The 'AND' or 'OR' rule. The value 0 indicates the 'OR' rule (indicating that an alarm will be triggered if any rule meets the threshold condition). The value 1 indicates the 'AND' rule (indicating that an alarm will be triggered when all rules meet the threshold conditions).
	// Note: This field may return null, indicating that no valid value was found.
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupListGroupInstanceGroup struct {

	// Instance group name ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Policy type view name.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// Instance group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Number of instances.
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// Update time.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DescribePolicyGroupListRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Search by policy name.
	Like *string `json:"Like,omitempty" name:"Like"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Sort by update time. Valid values: asc and desc.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// Project ID list.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// List of alarm policy types.
	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames" list`

	// Whether to filter policy groups without recipients. The value 1 indicates that policy groups without recipients will be filtered. The value 0 indicates that policy groups without recipients will not be filtered.
	FilterUnuseReceiver *int64 `json:"FilterUnuseReceiver,omitempty" name:"FilterUnuseReceiver"`

	// Filter by recipient group.
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers" list`

	// Filter by recipient.
	ReceiverUserList []*string `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`

	// Dimension set field (json string), for example, [[{"name":"unInstanceId","value":"ins-6e4b2aaa"}]].
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// Template-based policy group IDs, which are separated by commas.
	ConditionTempGroupId *string `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`

	// Filter by recipient or recipient group. The value 'user' indicates by recipient. The value 'group' indicates by recipient group.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// Filter conditions. Whether the alarm policy has been enabled or disabled
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
}

func (r *DescribePolicyGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy group list.
	// Note: This field may return null, indicating that no valid value was found.
		GroupList []*DescribePolicyGroupListGroup `json:"GroupList,omitempty" name:"GroupList" list`

		// Total number of policy groups.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListDimensions struct {

	// Dimension name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Dimension value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeProductEventListEvents struct {

	// Event ID.
	// Note: This field may return null, indicating that no valid value was found.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Event name in Chinese.
	// Note: This field may return null, indicating that no valid value was found.
	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`

	// Event name in English.
	// Note: This field may return null, indicating that no valid value was found.
	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`

	// Event name abbreviation.
	// Note: This field may return null, indicating that no valid value was found.
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Product name in Chinese.
	// Note: This field may return null, indicating that no valid value was found.
	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`

	// Product name in English.
	// Note: This field may return null, indicating that no valid value was found.
	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`

	// Product name abbreviation.
	// Note: This field may return null, indicating that no valid value was found.
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Project ID.
	// Note: This field may return null, indicating that no valid value was found.
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region.
	// Note: This field may return null, indicating that no valid value was found.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Status.
	// Note: This field may return null, indicating that no valid value was found.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Whether to support alarms.
	// Note: This field may return null, indicating that no valid value was found.
	SupportAlarm *int64 `json:"SupportAlarm,omitempty" name:"SupportAlarm"`

	// Event type.
	// Note: This field may return null, indicating that no valid value was found.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Start time.
	// Note: This field may return null, indicating that no valid value was found.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid value was found.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Instance object information.
	// Note: This field may return null, indicating that no valid value was found.
	Dimensions []*DescribeProductEventListEventsDimensions `json:"Dimensions,omitempty" name:"Dimensions" list`

	// Additional information of the instance object.
	// Note: This field may return null, indicating that no valid value was found.
	AdditionMsg []*DescribeProductEventListEventsDimensions `json:"AdditionMsg,omitempty" name:"AdditionMsg" list`

	// Whether to configure alarms.
	// Note: This field may return null, indicating that no valid value was found.
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`

	// Policy information.
	// Note: This field may return null, indicating that no valid value was found.
	GroupInfo []*DescribeProductEventListEventsGroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo" list`

	// Display name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

type DescribeProductEventListEventsDimensions struct {

	// Dimension name in English.
	// Note: This field may return null, indicating that no valid value was found.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Dimension name in Chinese.
	// Note: This field may return null, indicating that no valid value was found.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Dimension value.
	// Note: This field may return null, indicating that no valid value was found.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeProductEventListEventsGroupInfo struct {

	// Policy ID.
	// Note: This field may return null, indicating that no valid value was found.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Policy name.
	// Note: This field may return null, indicating that no valid value was found.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type DescribeProductEventListOverView struct {

	// Number of events whose statuses have changed.
	// Note: This field may return null, indicating that no valid value was found.
	StatusChangeAmount *int64 `json:"StatusChangeAmount,omitempty" name:"StatusChangeAmount"`

	// Number of events whose alarm statuses are not configured.
	// Note: This field may return null, indicating that no valid value was found.
	UnConfigAlarmAmount *int64 `json:"UnConfigAlarmAmount,omitempty" name:"UnConfigAlarmAmount"`

	// Number of events with exceptions.
	// Note: This field may return null, indicating that no valid value was found.
	UnNormalEventAmount *int64 `json:"UnNormalEventAmount,omitempty" name:"UnNormalEventAmount"`

	// Number of events that have not been recovered.
	// Note: This field may return null, indicating that no valid value was found.
	UnRecoverAmount *int64 `json:"UnRecoverAmount,omitempty" name:"UnRecoverAmount"`
}

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest

	// API component name. It is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Filter by product type. For example, 'cvm' indicates Cloud Virtual Machine.
	ProductName []*string `json:"ProductName,omitempty" name:"ProductName" list`

	// Filter by product name. For example, "guest_reboot" indicates server restart.
	EventName []*string `json:"EventName,omitempty" name:"EventName" list`

	// Affected object, such as "ins-19708ino"
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId" list`

	// Filter by dimension, such as by public IP: 10.0.0.1.
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions" list`

	// Region filter parameter for service events, such as `gz`. For region abbreviations, please see [Region List](https://intl.cloud.tencent.com/document/product/248/50863?from_cn_redirect=1)
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList" list`

	// Filter by event type. Valid values: ["status_change","abnormal"], which indicate events whose statuses have changed and events with exceptions respectively.
	Type []*string `json:"Type,omitempty" name:"Type" list`

	// Filter by event status. Valid values: ["recover","alarm","-"], which indicate that an event has been recovered, has not been recovered, and has no status respectively.
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// Filter by project ID.
	Project []*string `json:"Project,omitempty" name:"Project" list`

	// Filter by alarm status configuration. The value 1 indicates that the alarm status has been configured. The value 0 indicates that the alarm status has not been configured.
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`

	// Sorting by update time. The value ASC indicates the ascending order. The value DESC indicates the descending order. The default value is DESC.
	TimeOrder *string `json:"TimeOrder,omitempty" name:"TimeOrder"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Page offset. The default value is 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of parameters that can be returned on each page. The default value is 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Event list
	// Note: This field may return null, indicating that no valid value was found.
		Events []*DescribeProductEventListEvents `json:"Events,omitempty" name:"Events" list`

		// Event statistics.
		OverView *DescribeProductEventListOverView `json:"OverView,omitempty" name:"OverView"`

		// Total number of events.
	// Note: This field may return null, indicating that no valid value was found.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDataRequest struct {
	*tchttp.BaseRequest

	// Module, whose value is fixed at `monitor`
	Module *string `json:"Module,omitempty" name:"Module"`

	// Namespace. Valid values: QCE/TKE
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric name list
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`

	// Dimension condition. The `=` and `in` operators are supported
	Conditions []*MidQueryCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// Statistical granularity in s. Default value: 300
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Start time, which is the current time by default, such as 2020-12-08T19:51:23+08:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is the current time by default, such as 2020-12-08T19:51:23+08:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// `groupBy` by the specified dimension
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`
}

func (r *DescribeStatisticDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStatisticDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistical period
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// Start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// End time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Monitoring data
		Data []*MetricData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStatisticDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {

	// Instance dimension name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Instance dimension value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DimensionsDesc struct {

	// Array of dimension names
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest

	// Namespace, such as QCE/CVM. For more information on the namespaces of each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1)
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric name, such as `CPUUsage`. For more information on the metrics of each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). The corresponding metric name is the `MetricName`
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Dimension combination of instance object in the format of `key-value` pair, such as [{"Name":"InstanceId","Value":"ins-j0hk02zo"}]. For more information on the dimensions of each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). The value in the dimension column is the `key` in the dimension combination, and the value corresponding to the `key` is the `value` in the combination
	Instances []*Instance `json:"Instances,omitempty" name:"Instances" list`

	// Monitoring statistical period in seconds, such as 60. Default value: 300. The statistical period varies by metric. For more information on the statistical periods supported by each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). The values in the statistical period column are the supported statistical periods
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Start time such as 2018-09-22T19:51:23+08:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is the current time by default, such as 2018-09-22T20:51:23+08:00. `EndTime` cannot be earlier than `StartTime`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistical period
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// Metric name
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// Array of data points
		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints" list`

		// Start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// End time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// Combination of instance dimensions
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type InstanceGroup struct {

	// Instance group ID.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Instance group name.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`
}

type InstanceGroups struct {

	// Instance group ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Instance group name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Metric struct {

	// Alarm policy type
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric name
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Metric display name
	Description *string `json:"Description,omitempty" name:"Description"`

	// Minimum value
	Min *float64 `json:"Min,omitempty" name:"Min"`

	// Maximum value
	Max *float64 `json:"Max,omitempty" name:"Max"`

	// Dimension list
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions" list`

	// Unit
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Metric configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	MetricConfig *MetricConfig `json:"MetricConfig,omitempty" name:"MetricConfig"`
}

type MetricConfig struct {

	// Allowed operator
	Operator []*string `json:"Operator,omitempty" name:"Operator" list`

	// Allowed data cycle in seconds
	Period []*int64 `json:"Period,omitempty" name:"Period" list`

	// Allowed number of continuous cycles
	ContinuePeriod []*int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod" list`
}

type MetricData struct {

	// Metric name
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Monitoring data point
	Points []*MetricDataPoint `json:"Points,omitempty" name:"Points" list`
}

type MetricDataPoint struct {

	// Combination of instance object dimensions
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// Data point list
	Values []*Point `json:"Values,omitempty" name:"Values" list`
}

type MetricDatum struct {

	// Metric name.
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Metric value.
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type MetricObjectMeaning struct {

	// Meaning of the metric in English
	En *string `json:"En,omitempty" name:"En"`

	// Meaning of the metric in Chinese
	Zh *string `json:"Zh,omitempty" name:"Zh"`
}

type MetricSet struct {

	// Namespace. Each Tencent Cloud product has a namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric Name
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Unit used by the metric
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Unit used by the metric
	UnitCname *string `json:"UnitCname,omitempty" name:"UnitCname"`

	// Statistical period in seconds supported by the metric, such as 60 and 300
	Period []*int64 `json:"Period,omitempty" name:"Period" list`

	// Metric method during the statistical period
	Periods []*PeriodsSt `json:"Periods,omitempty" name:"Periods" list`

	// Meaning of the statistical metric
	Meaning *MetricObjectMeaning `json:"Meaning,omitempty" name:"Meaning"`

	// Dimension description
	Dimensions []*DimensionsDesc `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type MidQueryCondition struct {

	// Dimension
	Key *string `json:"Key,omitempty" name:"Key"`

	// Operator. Valid values: eq (equal to), ne (not equal to), in
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Dimension value. If `Operator` is `eq` or `ne`, only the first element will be used
	Value []*string `json:"Value,omitempty" name:"Value" list`
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm notification rule name, which can contain up to 60 characters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`

	// Alarm notification template ID
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`

	// User notifications (up to 5)
	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices" list`

	// Callback notifications (up to 3)
	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices" list`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyConditionRequest struct {
	*tchttp.BaseRequest

	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// ID of trigger condition template. This parameter can be left empty.
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// Event trigger condition
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`
}

func (r *ModifyAlarmPolicyConditionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyConditionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyConditionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyConditionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyConditionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyInfoRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Field to be modified. Valid values: NAME (policy name), REMARK (policy remarks)
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value after modification
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyAlarmPolicyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyNoticeRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Alarm notification template ID list
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds" list`
}

func (r *ModifyAlarmPolicyNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyNoticeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyNoticeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Status. Valid values: 0 (disabled), 1 (enabled)
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyAlarmPolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyTasksRequest struct {
	*tchttp.BaseRequest

	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// List of tasks triggered by alarm policy. If this parameter is left empty, it indicates to unbind all tasks
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks" list`
}

func (r *ModifyAlarmPolicyTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversRequest struct {
	*tchttp.BaseRequest

	// ID of a policy group whose recipient needs to be modified.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Required. The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// New recipient information. If this parameter is not configured, all recipients will be deleted.
	ReceiverInfos []*ReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos" list`
}

func (r *ModifyAlarmReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmReceiversRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmReceiversResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupCondition struct {

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to.
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// Threshold.
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// Data period of the detected metric.
	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// Number of consecutive periods.
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// Alarm sending and convergence type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. If the value is less than 0, no alarm will be triggered. If the value is 0, an alarm will be triggered only once. If the value is greater than 0, an alarm will be triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// Rule ID. No filling means new addition while filling in ruleId means to modify existing rules.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyPolicyGroupEventCondition struct {

	// Event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Alarm sending and convergence type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. If the value is less than 0, no alarm will be triggered. If the value is 0, an alarm will be triggered only once. If the value is greater than 0, an alarm will be triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// Rule ID. No filling means new addition while filling in ruleId means to modify existing rules.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Alarm type.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// Policy group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The 'AND' and 'OR' rules for metric alarms. The value 1 indicates 'AND', which means that an alarm will be triggered only when all rules are met. The value 0 indicates 'OR', which means that an alarm will be triggered when any rule is met.
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// Metric alarm condition rules. No filling indicates that all existing metric alarm condition rules will be deleted.
	Conditions []*ModifyPolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// Event alarm conditions. No filling indicates that all existing event alarm conditions will be deleted.
	EventConditions []*ModifyPolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// Template-based policy group ID.
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`
}

func (r *ModifyPolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy group ID.
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorTypeNamespace struct {

	// Monitor type
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// Policy type value
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type PeriodsSt struct {

	// Period
	Period *string `json:"Period,omitempty" name:"Period"`

	// Statistical method
	StatType []*string `json:"StatType,omitempty" name:"StatType" list`
}

type Point struct {

	// Time point when this monitoring data point is generated
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Monitoring data point value
	// Note: this field may return null, indicating that no valid values can be obtained.
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type PutMonitorDataRequest struct {
	*tchttp.BaseRequest

	// A group of metrics and data.
	Metrics []*MetricDatum `json:"Metrics,omitempty" name:"Metrics" list`

	// IP address that is automatically specified when monitoring data is reported.
	AnnounceIp *string `json:"AnnounceIp,omitempty" name:"AnnounceIp"`

	// Timestamp that is automatically specified when monitoring data is reported.
	AnnounceTimestamp *uint64 `json:"AnnounceTimestamp,omitempty" name:"AnnounceTimestamp"`

	// IP address or product instance ID that is automatically specified when monitoring data is reported.
	AnnounceInstance *string `json:"AnnounceInstance,omitempty" name:"AnnounceInstance"`
}

func (r *PutMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutMonitorDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutMonitorDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReceiverInfo struct {

	// Start time of the alarm period. Value range: [0,86400). Convert the Unix timestamp to Beijing time and then remove the date. For example, 7200 indicates '10:0:0'.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the alarm period. The meaning is the same as that of StartTime.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Alarm notification method. Valid values: "SMS", "SITE", "EMAIL", "CALL", and "WECHAT".
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// Recipient type. Valid values: group and user.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// ReceiverId
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Alarm call notification time. Valid values: OCCUR (indicating that a notice is sent when the alarm is triggered) and RECOVER (indicating that a notice is sent when the alarm is recovered).
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor" list`

	// Uid of the alarm call recipient.
	UidList []*int64 `json:"UidList,omitempty" name:"UidList" list`

	// Number of alarm call rounds.
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// Alarm call intervals for individuals in seconds.
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// Intervals of alarm call rounds in seconds.
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// Notification method when an alarm is recovered. Valid value: SMS.
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify" list`

	// Whether to send an alarm call delivery notice. The value 0 indicates that no notice needs to be sent. The value 1 indicates that a notice needs to be sent.
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// Recipient group list. The list of recipient group IDs that is queried by a platform API.
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList" list`

	// Recipient list. The list of recipient IDs that is queried by a platform API.
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`

	// Language of received alarms. Enumerated values: zh-CN and en-US.
	ReceiveLanguage *string `json:"ReceiveLanguage,omitempty" name:"ReceiveLanguage"`
}

type SendCustomAlarmMsgRequest struct {
	*tchttp.BaseRequest

	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Message policy ID, which is configured on the custom message page of Cloud Monitor.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Custom message content that a user wants to send.
	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

func (r *SendCustomAlarmMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendCustomAlarmMsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendCustomAlarmMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCustomAlarmMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendCustomAlarmMsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetDefaultAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *SetDefaultAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDefaultAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetDefaultAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDefaultAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TagInstance struct {

	// Tag key
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Number of instances
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// Service type, for example, CVM
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Region ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Binding status. 2: bound; 1: binding
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BindingStatus *int64 `json:"BindingStatus,omitempty" name:"BindingStatus"`

	// Tag status. 2: existent; 1: nonexistent
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagStatus *int64 `json:"TagStatus,omitempty" name:"TagStatus"`
}

type URLNotice struct {

	// Callback URL, which can contain up to 256 characters
	// Note: this field may return null, indicating that no valid values can be obtained.
	URL *string `json:"URL,omitempty" name:"URL"`

	// Whether verification is passed. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsValid *int64 `json:"IsValid,omitempty" name:"IsValid"`

	// Verification code
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidationCode *string `json:"ValidationCode,omitempty" name:"ValidationCode"`
}

type UnBindingAllPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy group ID. If `PolicyId` is used, this parameter will be ignored, and any value, e.g., `0`, can be passed in.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Alarm policy ID. If this parameter is used, `GroupId` will be ignored.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *UnBindingAllPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingAllPolicyObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingAllPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingAllPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingAllPolicyObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy group ID. If `PolicyId` is used, this parameter will be ignored, and any value, e.g., `0`, can be passed in.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// List of unique IDs of the object instances to be deleted. `UniqueId` can be obtained from the output parameter `List` of the [DescribeBindingPolicyObjectList](https://intl.cloud.tencent.com/document/api/248/40570?from_cn_redirect=1) API
	UniqueId []*string `json:"UniqueId,omitempty" name:"UniqueId" list`

	// Instance group ID. The `UniqueId` parameter is invalid if object instances are deleted by instance group.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Alarm policy ID. If this parameter is used, `GroupId` will be ignored.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *UnBindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingPolicyObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingPolicyObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UserNotice struct {

	// Recipient type. Valid values: USER (user), GROUP (user group)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// Notification start time, which is expressed by the number of seconds since 00:00:00. Value range: 0–86399
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Notification end time, which is expressed by the number of seconds since 00:00:00. Value range: 0–86399
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Notification channel list. Valid values: EMAIL (email), SMS (SMS), CALL (phone), WECHAT (WeChat)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeWay []*string `json:"NoticeWay,omitempty" name:"NoticeWay" list`

	// User `uid` list
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserIds []*int64 `json:"UserIds,omitempty" name:"UserIds" list`

	// User group ID list
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Phone polling list
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneOrder []*int64 `json:"PhoneOrder,omitempty" name:"PhoneOrder" list`

	// Number of phone pollings. Value range: 1–5
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`

	// Call interval in seconds within one polling. Value range: 60–900
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`

	// Polling interval in seconds. Value range: 60–900
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`

	// Whether receipt notification is required. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NeedPhoneArriveNotice *int64 `json:"NeedPhoneArriveNotice,omitempty" name:"NeedPhoneArriveNotice"`
}
