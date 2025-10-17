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

package v20180724

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AlarmEvent struct {
	// Event name
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Event display name
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Alarm policy type
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type AlarmHierarchicalNotice struct {
	// Notification template ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// The list of alarm notification levels. The values `Remind` and `Serious` indicate that the notification template only sends alarms at the `Remind` and `Serious` levels.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Classification []*string `json:"Classification,omitnil,omitempty" name:"Classification"`
}

type AlarmHierarchicalValue struct {
	// Threshold for the `Remind` level
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remind *string `json:"Remind,omitnil,omitempty" name:"Remind"`

	// Threshold for the `Warn` level
	// Note: This field may return null, indicating that no valid values can be obtained.
	Warn *string `json:"Warn,omitnil,omitempty" name:"Warn"`

	// Threshold for the `Serious` level
	// Note: This field may return null, indicating that no valid values can be obtained.
	Serious *string `json:"Serious,omitnil,omitempty" name:"Serious"`
}

type AlarmHistory struct {
	// Alarm record ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// Monitor type
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Policy type
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Alarm object
	AlarmObject *string `json:"AlarmObject,omitnil,omitempty" name:"AlarmObject"`

	// Alarm content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Timestamp of the first occurrence
	FirstOccurTime *int64 `json:"FirstOccurTime,omitnil,omitempty" name:"FirstOccurTime"`

	// Timestamp of the last occurrence
	LastOccurTime *int64 `json:"LastOccurTime,omitnil,omitempty" name:"LastOccurTime"`

	// Alarm status. Valid values: ALARM (not resolved), OK (resolved), NO_CONF (expired), NO_DATA (insufficient data)
	AlarmStatus *string `json:"AlarmStatus,omitnil,omitempty" name:"AlarmStatus"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// VPC of alarm object for basic product alarm
	VPC *string `json:"VPC,omitnil,omitempty" name:"VPC"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Instance group of alarm object
	InstanceGroup []*InstanceGroups `json:"InstanceGroup,omitnil,omitempty" name:"InstanceGroup"`

	// Recipient list
	ReceiverUids []*int64 `json:"ReceiverUids,omitnil,omitempty" name:"ReceiverUids"`

	// Recipient group list
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// Alarm channel list. Valid values: SMS (SMS), EMAIL (email), CALL (phone), WECHAT (WeChat)
	NoticeWays []*string `json:"NoticeWays,omitnil,omitempty" name:"NoticeWays"`

	// Alarm policy ID, which can be used when you call APIs ([BindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40421?from_cn_redirect=1), [UnBindingAllPolicyObject](https://intl.cloud.tencent.com/document/product/248/40568?from_cn_redirect=1), [UnBindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40567?from_cn_redirect=1)) to bind/unbind instances or instance groups to/from an alarm policy
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// Alarm type
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// Event ID
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Whether the policy exists. Valid values: 0 (no), 1 (yes)
	PolicyExists *int64 `json:"PolicyExists,omitnil,omitempty" name:"PolicyExists"`

	// Metric information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MetricsInfo []*AlarmHistoryMetric `json:"MetricsInfo,omitnil,omitempty" name:"MetricsInfo"`

	// Dimension information of an instance that triggered alarms.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type AlarmHistoryMetric struct {
	// Namespace used to query data by Tencent Cloud service monitoring type
	QceNamespace *string `json:"QceNamespace,omitnil,omitempty" name:"QceNamespace"`

	// Metric name
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Statistical period
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Value triggering alarm
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Metric display name
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AlarmNotice struct {
	// Alarm notification template ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Alarm notification template name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Last modified time
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Last modified by
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdatedBy *string `json:"UpdatedBy,omitnil,omitempty" name:"UpdatedBy"`

	// Alarm notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeType *string `json:"NoticeType,omitnil,omitempty" name:"NoticeType"`

	// User notification list
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserNotices []*UserNotice `json:"UserNotices,omitnil,omitempty" name:"UserNotices"`

	// Callback notification list
	// Note: this field may return null, indicating that no valid values can be obtained.
	URLNotices []*URLNotice `json:"URLNotices,omitnil,omitempty" name:"URLNotices"`

	// Whether it is the system default notification template. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsPreset *int64 `json:"IsPreset,omitnil,omitempty" name:"IsPreset"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeLanguage *string `json:"NoticeLanguage,omitnil,omitempty" name:"NoticeLanguage"`

	// List of IDs of the alarm policies bound to alarm notification template
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// Backend AMP consumer ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AMPConsumerId *string `json:"AMPConsumerId,omitnil,omitempty" name:"AMPConsumerId"`

	// Channel to push alarm notifications to CLS.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CLSNotices []*CLSNotice `json:"CLSNotices,omitnil,omitempty" name:"CLSNotices"`

	// Tags bound to a notification template
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AlarmPolicy struct {
	// Alarm policy ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Alarm policy name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring)
	// Note: this field may return null, indicating that no valid values can be obtained.
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Status. Valid values: 0 (disabled), 1 (enabled)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Number of instances bound to policy group
	// Note: this field may return null, indicating that no valid values can be obtained.
	UseSum *int64 `json:"UseSum,omitnil,omitempty" name:"UseSum"`

	// Project ID. Valid values: -1 (no project), 0 (default project)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Alarm policy type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Trigger condition template ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConditionTemplateId *string `json:"ConditionTemplateId,omitnil,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Condition *AlarmPolicyCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Event trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitnil,omitempty" name:"EventCondition"`

	// Notification rule ID list
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Notification rule list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Notices []*AlarmNotice `json:"Notices,omitnil,omitempty" name:"Notices"`

	// Triggered task list
	// Note: this field may return null, indicating that no valid values can be obtained.
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitnil,omitempty" name:"TriggerTasks"`

	// Template policy group
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConditionsTemp *ConditionsTemp `json:"ConditionsTemp,omitnil,omitempty" name:"ConditionsTemp"`

	// `Uin` of the last modifying user
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastEditUin *string `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Update time
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Region
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// Namespace display name
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceShowName *string `json:"NamespaceShowName,omitnil,omitempty" name:"NamespaceShowName"`

	// Whether it is the default policy. Valid values: 1 (yes), 0 (no)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Whether the default policy can be set. Valid values: 1 (yes), 0 (no)
	// Note: this field may return null, indicating that no valid values can be obtained.
	CanSetDefault *int64 `json:"CanSetDefault,omitnil,omitempty" name:"CanSetDefault"`

	// Instance group ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Total number of instances in instance group
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceSum *int64 `json:"InstanceSum,omitnil,omitempty" name:"InstanceSum"`

	// Instance group name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceGroupName *string `json:"InstanceGroupName,omitnil,omitempty" name:"InstanceGroupName"`

	// Trigger condition type. Valid values: STATIC (static threshold), DYNAMIC (dynamic)
	// Note: this field may return null, indicating that no valid values can be obtained.
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Policy ID for instance/instance group binding and unbinding APIs (BindingPolicyObject, UnBindingAllPolicyObject, UnBindingPolicyObject)
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// Tag
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagInstances []*TagInstance `json:"TagInstances,omitnil,omitempty" name:"TagInstances"`

	// Information on the filter dimension associated with a policy.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FilterDimensionsParam *string `json:"FilterDimensionsParam,omitnil,omitempty" name:"FilterDimensionsParam"`

	// Whether it is a quick alarm policy.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsOneClick *int64 `json:"IsOneClick,omitnil,omitempty" name:"IsOneClick"`

	// Whether the quick alarm policy is enabled.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OneClickStatus *int64 `json:"OneClickStatus,omitnil,omitempty" name:"OneClickStatus"`

	// The number of advanced metrics.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AdvancedMetricNumber *int64 `json:"AdvancedMetricNumber,omitnil,omitempty" name:"AdvancedMetricNumber"`

	// Whether the policy is associated with all objects
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsBindAll *int64 `json:"IsBindAll,omitnil,omitempty" name:"IsBindAll"`

	// Policy tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AlarmPolicyCondition struct {
	// Judgment condition of an alarm trigger condition (`0`: Any; `1`: All; `2`: Composite). When the value is set to `2` (i.e., composite trigger conditions), this parameter should be used together with `ComplexExpression`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`

	// Alarm trigger condition list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rules []*AlarmPolicyRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// The judgment expression of composite alarm trigger conditions, which is valid when the value of `IsUnionRule` is `2`. This parameter is used to determine that an alarm condition is met only when the expression values are `True` for multiple trigger conditions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComplexExpression *string `json:"ComplexExpression,omitnil,omitempty" name:"ComplexExpression"`
}

type AlarmPolicyEventCondition struct {
	// Alarm trigger condition list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Rules []*AlarmPolicyRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type AlarmPolicyFilter struct {
	// Filter condition type. Valid values: DIMENSION (uses dimensions for filtering)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// JSON string generated by serializing the `AlarmPolicyDimension` two-dimensional array. The one-dimensional arrays are in OR relationship, and the elements in a one-dimensional array are in AND relationship
	// Note: this field may return null, indicating that no valid values can be obtained.
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type AlarmPolicyRule struct {
	// Metric name or event name. The supported metrics can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1) and the supported events via [DescribeAlarmEvents](https://intl.cloud.tencent.com/document/product/248/51284?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Statistical period in seconds. The valid values can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

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
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Threshold. The valid value range can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Number of periods. `1`: continue for one period; `2`: continue for two periods; and so on. The valid values can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ContinuePeriod *int64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`

	// Alarm interval in seconds. Valid values: 0 (do not repeat), 300 (alarm once every 5 minutes), 600 (alarm once every 10 minutes), 900 (alarm once every 15 minutes), 1800 (alarm once every 30 minutes), 3600 (alarm once every hour), 7200 (alarm once every 2 hours), 10800 (alarm once every 3 hours), 21600 (alarm once every 6 hours),  43200 (alarm once every 12 hours), 86400 (alarm once every day)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NoticeFrequency *int64 `json:"NoticeFrequency,omitnil,omitempty" name:"NoticeFrequency"`

	// Whether the alarm frequency increases exponentially. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsPowerNotice *int64 `json:"IsPowerNotice,omitnil,omitempty" name:"IsPowerNotice"`

	// Filter condition for one single trigger rule
	// Note: this field may return null, indicating that no valid values can be obtained.
	Filter *AlarmPolicyFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Metric display name, which is used in the output parameter
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Unit, which is used in the output parameter
	// Note: this field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Trigger condition type. `STATIC`: static threshold; `dynamic`: dynamic threshold. If you do not specify this parameter when creating or editing a policy, `STATIC` is used by default.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// Whether it is an advanced metric. 0: No; 1: Yes.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsAdvanced *int64 `json:"IsAdvanced,omitnil,omitempty" name:"IsAdvanced"`

	// Whether the advanced metric feature is enabled. 0: No; 1: Yes.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsOpen *int64 `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`

	// Integration center product ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Maximum value
	// Note: This field may return null, indicating that no valid values can be obtained.
	ValueMax *float64 `json:"ValueMax,omitnil,omitempty" name:"ValueMax"`

	// Minimum value
	// Note: This field may return null, indicating that no valid values can be obtained.
	ValueMin *float64 `json:"ValueMin,omitnil,omitempty" name:"ValueMin"`

	// The configuration of alarm level threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	HierarchicalValue *AlarmHierarchicalValue `json:"HierarchicalValue,omitnil,omitempty" name:"HierarchicalValue"`
}

type AlarmPolicyTriggerTask struct {
	// Triggered task type. Valid value: AS (auto scaling)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Configuration information in JSON format, such as {"Key1":"Value1","Key2":"Value2"}
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskConfig *string `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`
}

// Predefined struct for user
type BindPrometheusManagedGrafanaRequestParams struct {
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Grafana instance ID
	GrafanaId *string `json:"GrafanaId,omitnil,omitempty" name:"GrafanaId"`
}

type BindPrometheusManagedGrafanaRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Grafana instance ID
	GrafanaId *string `json:"GrafanaId,omitnil,omitempty" name:"GrafanaId"`
}

func (r *BindPrometheusManagedGrafanaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPrometheusManagedGrafanaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GrafanaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindPrometheusManagedGrafanaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindPrometheusManagedGrafanaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindPrometheusManagedGrafanaResponse struct {
	*tchttp.BaseResponse
	Response *BindPrometheusManagedGrafanaResponseParams `json:"Response"`
}

func (r *BindPrometheusManagedGrafanaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPrometheusManagedGrafanaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindingPolicyObjectDimension struct {
	// Region name.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region ID.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Instance dimension information in the following format:
	// {"unInstanceId":"ins-00jvv9mo"}. The dimension information varies by Tencent Cloud services. For more information, please see:
	// [Dimension List](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1)
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Event dimensions.
	EventDimensions *string `json:"EventDimensions,omitnil,omitempty" name:"EventDimensions"`
}

// Predefined struct for user
type BindingPolicyObjectRequestParams struct {
	// Required. The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID, such as `4739573`. This parameter will be disused soon. Another parameter `PolicyId` is recommended.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm policy ID, such as `policy-gh892hg0`. At least one of the two parameters, `PolicyId` and `GroupId`, must be specified; otherwise, an error will be reported. `PolicyId` is preferred over `GroupId` when both of them are specified.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Dimensions of an object to be bound.
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`

	// Whether the event alert has been configured
	EbEventFlag *int64 `json:"EbEventFlag,omitnil,omitempty" name:"EbEventFlag"`
}

type BindingPolicyObjectRequest struct {
	*tchttp.BaseRequest
	
	// Required. The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID, such as `4739573`. This parameter will be disused soon. Another parameter `PolicyId` is recommended.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm policy ID, such as `policy-gh892hg0`. At least one of the two parameters, `PolicyId` and `GroupId`, must be specified; otherwise, an error will be reported. `PolicyId` is preferred over `GroupId` when both of them are specified.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Dimensions of an object to be bound.
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`

	// Whether the event alert has been configured
	EbEventFlag *int64 `json:"EbEventFlag,omitnil,omitempty" name:"EbEventFlag"`
}

func (r *BindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindingPolicyObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "PolicyId")
	delete(f, "InstanceGroupId")
	delete(f, "Dimensions")
	delete(f, "EbSubject")
	delete(f, "EbEventFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindingPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindingPolicyObjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *BindingPolicyObjectResponseParams `json:"Response"`
}

func (r *BindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindingPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CLSNotice struct {
	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Logset ID.
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// Topic ID.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Status. Valid values: `0` (disabled), `1` (enabled). Default value: `1` (enabled). This parameter can be left empty.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

// Predefined struct for user
type CheckIsPrometheusNewUserRequestParams struct {

}

type CheckIsPrometheusNewUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CheckIsPrometheusNewUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIsPrometheusNewUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIsPrometheusNewUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIsPrometheusNewUserResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckIsPrometheusNewUserResponse struct {
	*tchttp.BaseResponse
	Response *CheckIsPrometheusNewUserResponseParams `json:"Response"`
}

func (r *CheckIsPrometheusNewUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIsPrometheusNewUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanGrafanaInstanceRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CleanGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CleanGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CleanGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanGrafanaInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CleanGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CleanGrafanaInstanceResponseParams `json:"Response"`
}

func (r *CleanGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonNamespace struct {
	// Namespace ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Namespace name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Namespace value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Product name
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Configuration information
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// List of supported regions
	AvailableRegions []*string `json:"AvailableRegions,omitnil,omitempty" name:"AvailableRegions"`

	// Sort ID
	SortId *int64 `json:"SortId,omitnil,omitempty" name:"SortId"`

	// Unique ID in Dashboard
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`
}

type CommonNamespaceNew struct {
	// Namespace ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Namespace name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Monitoring type
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Dimension information
	Dimensions []*DimensionNew `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type Condition struct {
	// Alarm notification frequency.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// Predefined repeated notification policy. `0`: One-time alarm; `1`: exponential alarm; `2`: consecutive alarm.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`

	// Detection method.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CalcType *string `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// Detection value.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CalcValue *string `json:"CalcValue,omitnil,omitempty" name:"CalcValue"`

	// Duration in seconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ContinueTime *string `json:"ContinueTime,omitnil,omitempty" name:"ContinueTime"`

	// Metric ID.
	MetricID *int64 `json:"MetricID,omitnil,omitempty" name:"MetricID"`

	// Displayed metric name.
	MetricDisplayName *string `json:"MetricDisplayName,omitnil,omitempty" name:"MetricDisplayName"`

	// Statistical period.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Rule ID.
	RuleID *int64 `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// Metric unit.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Whether it is an advanced metric. Valid values: `0` (no), `1` (yes).
	IsAdvanced *int64 `json:"IsAdvanced,omitnil,omitempty" name:"IsAdvanced"`

	// Whether the advance metric feature is enabled. Valid values: `0` (no), `1` (yes).
	IsOpen *int64 `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`

	// Product ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type ConditionsTemp struct {
	// Template name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Metric trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Condition *AlarmPolicyCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Event trigger condition
	// Note: this field may return null, indicating that no valid values can be obtained.
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitnil,omitempty" name:"EventCondition"`
}

// Predefined struct for user
type CreateAlarmNoticeRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Notification template name, which can contain up to 60 characters
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	NoticeType *string `json:"NoticeType,omitnil,omitempty" name:"NoticeType"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	NoticeLanguage *string `json:"NoticeLanguage,omitnil,omitempty" name:"NoticeLanguage"`

	// User notifications (up to 5)
	UserNotices []*UserNotice `json:"UserNotices,omitnil,omitempty" name:"UserNotices"`

	// Callback notifications (up to 3)
	URLNotices []*URLNotice `json:"URLNotices,omitnil,omitempty" name:"URLNotices"`

	// The operation of pushing alarm notifications to CLS. Up to one CLS log topic can be configured.
	CLSNotices []*CLSNotice `json:"CLSNotices,omitnil,omitempty" name:"CLSNotices"`

	// Tags bound to a template
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Notification template name, which can contain up to 60 characters
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	NoticeType *string `json:"NoticeType,omitnil,omitempty" name:"NoticeType"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	NoticeLanguage *string `json:"NoticeLanguage,omitnil,omitempty" name:"NoticeLanguage"`

	// User notifications (up to 5)
	UserNotices []*UserNotice `json:"UserNotices,omitnil,omitempty" name:"UserNotices"`

	// Callback notifications (up to 3)
	URLNotices []*URLNotice `json:"URLNotices,omitnil,omitempty" name:"URLNotices"`

	// The operation of pushing alarm notifications to CLS. Up to one CLS log topic can be configured.
	CLSNotices []*CLSNotice `json:"CLSNotices,omitnil,omitempty" name:"CLSNotices"`

	// Tags bound to a template
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Name")
	delete(f, "NoticeType")
	delete(f, "NoticeLanguage")
	delete(f, "UserNotices")
	delete(f, "URLNotices")
	delete(f, "CLSNotices")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmNoticeResponseParams struct {
	// Alarm notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmNoticeResponseParams `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmPolicyRequestParams struct {
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy name, which can contain up to 20 characters
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring)
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Type of alarm policy, which can be obtained via [DescribeAllNamespaces](https://intl.cloud.tencent.com/document/product/248/48683?from_cn_redirect=1). For the monitoring of Tencent Cloud services, the value of this parameter is `QceNamespacesNew.N.Id` of the output parameter of `DescribeAllNamespaces`, for example, `cvm_device`.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Remarks with up to 100 letters, digits, underscores, and hyphens
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Whether to enable. Valid values: 0 (no), 1 (yes). Default value: 1. This parameter can be left empty
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Project ID. For products with different projects, a value other than `-1` must be passed in. `-1`: no project; `0`: default project. If no value is passed in, `-1` will be used. The supported project IDs can be viewed on the [**Account Center** > **Project Management**](https://console.cloud.tencent.com/project) page of the console.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Trigger condition template ID. Pass in this parameter if the policy is associated with the trigger condition template; otherwise, pass in the `Condition` parameter. The trigger condition template ID can be obtained via [`DescribeConditionsTemplateList`](https://intl.cloud.tencent.com/document/api/248/70250?from_cn_redirect=1).
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitnil,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition. The supported metrics can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	Condition *AlarmPolicyCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Event trigger condition. The supported events can be queried via [DescribeAlarmEvents](https://intl.cloud.tencent.com/document/product/248/51284?from_cn_redirect=1).
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitnil,omitempty" name:"EventCondition"`

	// List of notification rule IDs, which can be obtained via [DescribeAlarmNotices](https://intl.cloud.tencent.com/document/product/248/51280?from_cn_redirect=1)
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Triggered task list
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitnil,omitempty" name:"TriggerTasks"`

	// Global filter.
	Filter *AlarmPolicyFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Aggregation dimension list, which is used to specify which dimension keys data is grouped by.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Tags bound to a template
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Log alarm information
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitnil,omitempty" name:"LogAlarmReqInfo"`

	// Notification rules for different alarm levels
	HierarchicalNotices []*AlarmHierarchicalNotice `json:"HierarchicalNotices,omitnil,omitempty" name:"HierarchicalNotices"`

	// A dedicated field for migration policies. 0: Implement authentication logic; 1: Skip authentication logic.
	MigrateFlag *int64 `json:"MigrateFlag,omitnil,omitempty" name:"MigrateFlag"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`
}

type CreateAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy name, which can contain up to 20 characters
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring)
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Type of alarm policy, which can be obtained via [DescribeAllNamespaces](https://intl.cloud.tencent.com/document/product/248/48683?from_cn_redirect=1). For the monitoring of Tencent Cloud services, the value of this parameter is `QceNamespacesNew.N.Id` of the output parameter of `DescribeAllNamespaces`, for example, `cvm_device`.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Remarks with up to 100 letters, digits, underscores, and hyphens
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Whether to enable. Valid values: 0 (no), 1 (yes). Default value: 1. This parameter can be left empty
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Project ID. For products with different projects, a value other than `-1` must be passed in. `-1`: no project; `0`: default project. If no value is passed in, `-1` will be used. The supported project IDs can be viewed on the [**Account Center** > **Project Management**](https://console.cloud.tencent.com/project) page of the console.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Trigger condition template ID. Pass in this parameter if the policy is associated with the trigger condition template; otherwise, pass in the `Condition` parameter. The trigger condition template ID can be obtained via [`DescribeConditionsTemplateList`](https://intl.cloud.tencent.com/document/api/248/70250?from_cn_redirect=1).
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitnil,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition. The supported metrics can be queried via [DescribeAlarmMetrics](https://intl.cloud.tencent.com/document/product/248/51283?from_cn_redirect=1).
	Condition *AlarmPolicyCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Event trigger condition. The supported events can be queried via [DescribeAlarmEvents](https://intl.cloud.tencent.com/document/product/248/51284?from_cn_redirect=1).
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitnil,omitempty" name:"EventCondition"`

	// List of notification rule IDs, which can be obtained via [DescribeAlarmNotices](https://intl.cloud.tencent.com/document/product/248/51280?from_cn_redirect=1)
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Triggered task list
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitnil,omitempty" name:"TriggerTasks"`

	// Global filter.
	Filter *AlarmPolicyFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Aggregation dimension list, which is used to specify which dimension keys data is grouped by.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Tags bound to a template
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Log alarm information
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitnil,omitempty" name:"LogAlarmReqInfo"`

	// Notification rules for different alarm levels
	HierarchicalNotices []*AlarmHierarchicalNotice `json:"HierarchicalNotices,omitnil,omitempty" name:"HierarchicalNotices"`

	// A dedicated field for migration policies. 0: Implement authentication logic; 1: Skip authentication logic.
	MigrateFlag *int64 `json:"MigrateFlag,omitnil,omitempty" name:"MigrateFlag"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`
}

func (r *CreateAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyName")
	delete(f, "MonitorType")
	delete(f, "Namespace")
	delete(f, "Remark")
	delete(f, "Enable")
	delete(f, "ProjectId")
	delete(f, "ConditionTemplateId")
	delete(f, "Condition")
	delete(f, "EventCondition")
	delete(f, "NoticeIds")
	delete(f, "TriggerTasks")
	delete(f, "Filter")
	delete(f, "GroupBy")
	delete(f, "Tags")
	delete(f, "LogAlarmReqInfo")
	delete(f, "HierarchicalNotices")
	delete(f, "MigrateFlag")
	delete(f, "EbSubject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmPolicyResponseParams struct {
	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Alarm policy ID, which can be used when you call APIs ([BindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40421?from_cn_redirect=1), [UnBindingAllPolicyObject](https://intl.cloud.tencent.com/document/product/248/40568?from_cn_redirect=1), [UnBindingPolicyObject](https://intl.cloud.tencent.com/document/product/248/40567?from_cn_redirect=1)) to bind/unbind instances or instance groups to/from an alarm policy
	OriginId *string `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmPolicyResponseParams `json:"Response"`
}

func (r *CreateAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertRuleRequestParams struct {
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Alerting rule expression. For more information, see <a href="https://www.tencentcloud.com/document/product/1116/43192?lang=en&pg=">Alerting Rule Description</a>
	Expr *string `json:"Expr,omitnil,omitempty" name:"Expr"`

	// List of alert notification template IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Rule status code. Valid values:
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Rule alert duration
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// List of tags
	Labels []*PrometheusRuleKV `json:"Labels,omitnil,omitempty" name:"Labels"`

	// List of annotations.
	// 
	// Alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively.
	Annotations []*PrometheusRuleKV `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// Alerting rule template category
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CreateAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Alerting rule expression. For more information, see <a href="https://www.tencentcloud.com/document/product/1116/43192?lang=en&pg=">Alerting Rule Description</a>
	Expr *string `json:"Expr,omitnil,omitempty" name:"Expr"`

	// List of alert notification template IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Rule status code. Valid values:
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Rule alert duration
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// List of tags
	Labels []*PrometheusRuleKV `json:"Labels,omitnil,omitempty" name:"Labels"`

	// List of annotations.
	// 
	// Alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively.
	Annotations []*PrometheusRuleKV `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// Alerting rule template category
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *CreateAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	delete(f, "Expr")
	delete(f, "Receivers")
	delete(f, "RuleState")
	delete(f, "Duration")
	delete(f, "Labels")
	delete(f, "Annotations")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertRuleResponseParams struct {
	// Rule ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlertRuleResponseParams `json:"Response"`
}

func (r *CreateAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExporterIntegrationRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Integrated configuration
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type CreateExporterIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Integrated configuration
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *CreateExporterIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExporterIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Content")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExporterIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExporterIntegrationResponseParams struct {
	// The list of successfully created integrations.
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExporterIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateExporterIntegrationResponseParams `json:"Response"`
}

func (r *CreateExporterIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExporterIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaInstanceRequestParams struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of subnet IDs
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Initial Grafana password
	GrafanaInitPassword *string `json:"GrafanaInitPassword,omitnil,omitempty" name:"GrafanaInitPassword"`

	// Whether to enable public network access
	EnableInternet *bool `json:"EnableInternet,omitnil,omitempty" name:"EnableInternet"`

	// Tag
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type CreateGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of subnet IDs
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Initial Grafana password
	GrafanaInitPassword *string `json:"GrafanaInitPassword,omitnil,omitempty" name:"GrafanaInitPassword"`

	// Whether to enable public network access
	EnableInternet *bool `json:"EnableInternet,omitnil,omitempty" name:"EnableInternet"`

	// Tag
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *CreateGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "VpcId")
	delete(f, "SubnetIds")
	delete(f, "GrafanaInitPassword")
	delete(f, "EnableInternet")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaInstanceResponseParams struct {
	// Instance name
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateGrafanaInstanceResponseParams `json:"Response"`
}

func (r *CreateGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaIntegrationRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration type, such as “tencent-cloud-prometheus”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Integration configuration
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CreateGrafanaIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration type, such as “tencent-cloud-prometheus”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Integration configuration
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *CreateGrafanaIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGrafanaIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaIntegrationResponseParams struct {
	// Integration ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGrafanaIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateGrafanaIntegrationResponseParams `json:"Response"`
}

func (r *CreateGrafanaIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaNotificationChannelRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Default value: `1`. This parameter has been deprecated. Please use `OrganizationIds` instead.
	OrgId *int64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// Array of notification channel IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Array of extra organization IDs. This parameter has been deprecated. Please use `OrganizationIds` instead.
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitnil,omitempty" name:"ExtraOrgIds"`

	// Array of all valid organization IDs. Default value: `1`.
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

type CreateGrafanaNotificationChannelRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Default value: `1`. This parameter has been deprecated. Please use `OrganizationIds` instead.
	OrgId *int64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// Array of notification channel IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Array of extra organization IDs. This parameter has been deprecated. Please use `OrganizationIds` instead.
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitnil,omitempty" name:"ExtraOrgIds"`

	// Array of all valid organization IDs. Default value: `1`.
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

func (r *CreateGrafanaNotificationChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaNotificationChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ChannelName")
	delete(f, "OrgId")
	delete(f, "Receivers")
	delete(f, "ExtraOrgIds")
	delete(f, "OrganizationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGrafanaNotificationChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaNotificationChannelResponseParams struct {
	// Channel ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGrafanaNotificationChannelResponse struct {
	*tchttp.BaseResponse
	Response *CreateGrafanaNotificationChannelResponseParams `json:"Response"`
}

func (r *CreateGrafanaNotificationChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaNotificationChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupCondition struct {
	// Metric ID.
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to. This parameter is optional if a default comparative type is configured for the metric.
	CalcType *int64 `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// Comparative value. This parameter is optional if the metric has no requirement.
	CalcValue *float64 `json:"CalcValue,omitnil,omitempty" name:"CalcValue"`

	// Data aggregation period in seconds. This parameter is optional if the metric has a default value.
	CalcPeriod *int64 `json:"CalcPeriod,omitnil,omitempty" name:"CalcPeriod"`

	// Number of consecutive periods after which an alarm will be triggered.
	ContinuePeriod *int64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`

	// If a metric is created based on a template, the `RuleId` of the metric in the template must be passed in.
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type CreatePolicyGroupEventCondition struct {
	// Alarm event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// If a metric is created based on a template, the `RuleId` of the metric in the template must be passed in.
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

// Predefined struct for user
type CreatePolicyGroupRequestParams struct {
	// Policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Name of the view to which the policy group belongs. If the policy group is created based on a template, this parameter is optional.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// ID of the project to which the policy group belongs, which will be used for authentication.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// ID of a template-based policy group. This parameter is required only when the policy group is created based on a template.
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitnil,omitempty" name:"ConditionTempGroupId"`

	// Whether the policy group is shielded. The value 0 indicates that the policy group is not shielded. The value 1 indicates that the policy group is shielded. The default value is 0.
	IsShielded *int64 `json:"IsShielded,omitnil,omitempty" name:"IsShielded"`

	// Remarks of the policy group.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Insertion time in the format of Unix timestamp. If this parameter is not configured, the backend processing time is used.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Alarm threshold rules in the policy group.
	Conditions []*CreatePolicyGroupCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Event alarm rules in the policy group.
	EventConditions []*CreatePolicyGroupEventCondition `json:"EventConditions,omitnil,omitempty" name:"EventConditions"`

	// Whether it is a backend call. Rules pulled from the policy template will be used to fill in the `Conditions` and `EventConditions` fields only when the value of this parameter is `1`.
	BackEndCall *int64 `json:"BackEndCall,omitnil,omitempty" name:"BackEndCall"`

	// The 'AND' and 'OR' rules for alarm metrics. The value 0 indicates 'OR', which means that an alarm will be triggered when any rule is met. The value 1 indicates 'AND', which means that an alarm will be triggered only when all rules are met.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`
}

type CreatePolicyGroupRequest struct {
	*tchttp.BaseRequest
	
	// Policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Name of the view to which the policy group belongs. If the policy group is created based on a template, this parameter is optional.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// ID of the project to which the policy group belongs, which will be used for authentication.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// ID of a template-based policy group. This parameter is required only when the policy group is created based on a template.
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitnil,omitempty" name:"ConditionTempGroupId"`

	// Whether the policy group is shielded. The value 0 indicates that the policy group is not shielded. The value 1 indicates that the policy group is shielded. The default value is 0.
	IsShielded *int64 `json:"IsShielded,omitnil,omitempty" name:"IsShielded"`

	// Remarks of the policy group.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Insertion time in the format of Unix timestamp. If this parameter is not configured, the backend processing time is used.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Alarm threshold rules in the policy group.
	Conditions []*CreatePolicyGroupCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Event alarm rules in the policy group.
	EventConditions []*CreatePolicyGroupEventCondition `json:"EventConditions,omitnil,omitempty" name:"EventConditions"`

	// Whether it is a backend call. Rules pulled from the policy template will be used to fill in the `Conditions` and `EventConditions` fields only when the value of this parameter is `1`.
	BackEndCall *int64 `json:"BackEndCall,omitnil,omitempty" name:"BackEndCall"`

	// The 'AND' and 'OR' rules for alarm metrics. The value 0 indicates 'OR', which means that an alarm will be triggered when any rule is met. The value 1 indicates 'AND', which means that an alarm will be triggered only when all rules are met.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`
}

func (r *CreatePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Module")
	delete(f, "ViewName")
	delete(f, "ProjectId")
	delete(f, "ConditionTempGroupId")
	delete(f, "IsShielded")
	delete(f, "Remark")
	delete(f, "InsertTime")
	delete(f, "Conditions")
	delete(f, "EventConditions")
	delete(f, "BackEndCall")
	delete(f, "IsUnionRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyGroupResponseParams struct {
	// ID of the created policy group.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyGroupResponseParams `json:"Response"`
}

func (r *CreatePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAgentRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreatePrometheusAgentRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreatePrometheusAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAgentResponseParams struct {
	// ID of a successfully created agent.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAgentResponseParams `json:"Response"`
}

func (r *CreatePrometheusAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertGroupRequestParams struct {
	// prometheus instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert group name. Not allowed to have the same name as other Alert groups.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Alert group status.
	// 2 - enable.
	// 3 - disabled.
	// Specifies to overwrite all alert rule statuses under the `Rules` field when not empty.
	GroupState *int64 `json:"GroupState,omitnil,omitempty" name:"GroupState"`

	// Alert notification template ID list of tencent cloud observability platform, such as Consumer-xxxx or notice-xxxx.
	AMPReceivers []*string `json:"AMPReceivers,omitnil,omitempty" name:"AMPReceivers"`

	// Specifies the custom Alert notification template.
	CustomReceiver *PrometheusAlertCustomReceiver `json:"CustomReceiver,omitnil,omitempty" name:"CustomReceiver"`

	// Alert notification cycle (convergence time). defaults to 1h if left empty.
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// Specifies the Alert rule list to be created.
	Rules []*PrometheusAlertGroupRuleSet `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreatePrometheusAlertGroupRequest struct {
	*tchttp.BaseRequest
	
	// prometheus instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert group name. Not allowed to have the same name as other Alert groups.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Alert group status.
	// 2 - enable.
	// 3 - disabled.
	// Specifies to overwrite all alert rule statuses under the `Rules` field when not empty.
	GroupState *int64 `json:"GroupState,omitnil,omitempty" name:"GroupState"`

	// Alert notification template ID list of tencent cloud observability platform, such as Consumer-xxxx or notice-xxxx.
	AMPReceivers []*string `json:"AMPReceivers,omitnil,omitempty" name:"AMPReceivers"`

	// Specifies the custom Alert notification template.
	CustomReceiver *PrometheusAlertCustomReceiver `json:"CustomReceiver,omitnil,omitempty" name:"CustomReceiver"`

	// Alert notification cycle (convergence time). defaults to 1h if left empty.
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// Specifies the Alert rule list to be created.
	Rules []*PrometheusAlertGroupRuleSet `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreatePrometheusAlertGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupName")
	delete(f, "GroupState")
	delete(f, "AMPReceivers")
	delete(f, "CustomReceiver")
	delete(f, "RepeatInterval")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAlertGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertGroupResponseParams struct {
	// Created Alert group ID, which matches the regular expression `alert-[a-z0-9]{8}`.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusAlertGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAlertGroupResponseParams `json:"Response"`
}

func (r *CreatePrometheusAlertGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert configuration
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

type CreatePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert configuration
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

func (r *CreatePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertPolicyResponseParams struct {
	// Alerting rule ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *CreatePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusClusterAgentRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent list
	Agents []*PrometheusClusterAgentBasic `json:"Agents,omitnil,omitempty" name:"Agents"`
}

type CreatePrometheusClusterAgentRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent list
	Agents []*PrometheusClusterAgentBasic `json:"Agents,omitnil,omitempty" name:"Agents"`
}

func (r *CreatePrometheusClusterAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusClusterAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Agents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusClusterAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusClusterAgentResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusClusterAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusClusterAgentResponseParams `json:"Response"`
}

func (r *CreatePrometheusClusterAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusClusterAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Configuration of service monitors
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// Configuration of pod monitors
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// Configuration of Prometheus raw job
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`
}

type CreatePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Configuration of service monitors
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// Configuration of pod monitors
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// Configuration of Prometheus raw job
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`
}

func (r *CreatePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusConfigResponseParams `json:"Response"`
}

func (r *CreatePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusGlobalNotificationRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert notification channel
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

type CreatePrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert notification channel
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

func (r *CreatePrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Notification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusGlobalNotificationResponseParams struct {
	// ID of the global alert notification channel
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *CreatePrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusMultiTenantInstancePostPayModeRequestParams struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Data retention period in days. Valid values: 15, 30, 45.
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance tag
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// The Grafana instance to be associated
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitnil,omitempty" name:"GrafanaInstanceId"`
}

type CreatePrometheusMultiTenantInstancePostPayModeRequest struct {
	*tchttp.BaseRequest
	
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Data retention period in days. Valid values: 15, 30, 45.
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance tag
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// The Grafana instance to be associated
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitnil,omitempty" name:"GrafanaInstanceId"`
}

func (r *CreatePrometheusMultiTenantInstancePostPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusMultiTenantInstancePostPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DataRetentionTime")
	delete(f, "Zone")
	delete(f, "TagSpecification")
	delete(f, "GrafanaInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusMultiTenantInstancePostPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusMultiTenantInstancePostPayModeResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusMultiTenantInstancePostPayModeResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusMultiTenantInstancePostPayModeResponseParams `json:"Response"`
}

func (r *CreatePrometheusMultiTenantInstancePostPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusMultiTenantInstancePostPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusRecordRuleYamlRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// YAML content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreatePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// YAML content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreatePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Content")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusRecordRuleYamlResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *CreatePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusScrapeJobRequestParams struct {
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID, such as “agent-abcd1234”. It can be obtained on the **Agent Management** page in the console.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Scrape task ID in the format of “job_name:xx”
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreatePrometheusScrapeJobRequest struct {
	*tchttp.BaseRequest
	
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID, such as “agent-abcd1234”. It can be obtained on the **Agent Management** page in the console.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Scrape task ID in the format of “job_name:xx”
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *CreatePrometheusScrapeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusScrapeJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusScrapeJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusScrapeJobResponseParams struct {
	// ID of a successfully created scrape task.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusScrapeJobResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusScrapeJobResponseParams `json:"Response"`
}

func (r *CreatePrometheusScrapeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusScrapeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTempRequestParams struct {
	// Template settings
	Template *PrometheusTemp `json:"Template,omitnil,omitempty" name:"Template"`
}

type CreatePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// Template settings
	Template *PrometheusTemp `json:"Template,omitnil,omitempty" name:"Template"`
}

func (r *CreatePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTempResponseParams struct {
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusTempResponseParams `json:"Response"`
}

func (r *CreatePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordingRuleRequestParams struct {
	// Recording rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Recording rule group content in YAML format
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

type CreateRecordingRuleRequest struct {
	*tchttp.BaseRequest
	
	// Recording rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Recording rule group content in YAML format
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

func (r *CreateRecordingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Group")
	delete(f, "InstanceId")
	delete(f, "RuleState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordingRuleResponseParams struct {
	// Rule ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRecordingRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordingRuleResponseParams `json:"Response"`
}

func (r *CreateRecordingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSSOAccountRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User account ID, such as “10000000”.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Permission
	Role []*GrafanaAccountRole `json:"Role,omitnil,omitempty" name:"Role"`

	// Remarks
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
}

type CreateSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User account ID, such as “10000000”.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Permission
	Role []*GrafanaAccountRole `json:"Role,omitnil,omitempty" name:"Role"`

	// Remarks
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
}

func (r *CreateSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	delete(f, "Role")
	delete(f, "Notes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSSOAccountResponseParams struct {
	// The added user UIN
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateSSOAccountResponseParams `json:"Response"`
}

func (r *CreateSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceDiscoveryRequestParams struct {
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <li>TKE: ID of the integrated TKE cluster</li>
	KubeClusterId *string `json:"KubeClusterId,omitnil,omitempty" name:"KubeClusterId"`

	// Kubernetes cluster type:
	// <li> 1 = TKE </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Scrape configuration type. Valid values:
	// <li> 1 = ServiceMonitor</li>
	// <li> 2 = PodMonitor</li>
	// <li> 3 = JobMonitor</li>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Scrape configuration information
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

type CreateServiceDiscoveryRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <li>TKE: ID of the integrated TKE cluster</li>
	KubeClusterId *string `json:"KubeClusterId,omitnil,omitempty" name:"KubeClusterId"`

	// Kubernetes cluster type:
	// <li> 1 = TKE </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Scrape configuration type. Valid values:
	// <li> 1 = ServiceMonitor</li>
	// <li> 2 = PodMonitor</li>
	// <li> 3 = JobMonitor</li>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Scrape configuration information
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

func (r *CreateServiceDiscoveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceDiscoveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeClusterId")
	delete(f, "KubeType")
	delete(f, "Type")
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceDiscoveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceDiscoveryResponseParams struct {
	// The scrape configuration information returned after successful creation
	ServiceDiscovery *ServiceDiscoveryItem `json:"ServiceDiscovery,omitnil,omitempty" name:"ServiceDiscovery"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceDiscoveryResponseParams `json:"Response"`
}

func (r *CreateServiceDiscoveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceDiscoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataPoint struct {
	// Combination of instance object dimensions
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// The array of timestamps indicating at which points in time there is data. Missing timestamps have no data points (i.e., missed)
	Timestamps []*float64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// The array of monitoring values, which is in one-to-one correspondence to Timestamps
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type DeleteAlarmNoticesRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm notification template ID list
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Binding between a notification template and a policy
	NoticeBindPolicys []*NoticeBindPolicys `json:"NoticeBindPolicys,omitnil,omitempty" name:"NoticeBindPolicys"`
}

type DeleteAlarmNoticesRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm notification template ID list
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Binding between a notification template and a policy
	NoticeBindPolicys []*NoticeBindPolicys `json:"NoticeBindPolicys,omitnil,omitempty" name:"NoticeBindPolicys"`
}

func (r *DeleteAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "NoticeIds")
	delete(f, "NoticeBindPolicys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmNoticesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmNoticesResponseParams `json:"Response"`
}

func (r *DeleteAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmPolicyRequestParams struct {
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID list
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`
}

type DeleteAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID list
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`
}

func (r *DeleteAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmPolicyResponseParams `json:"Response"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlertRulesRequestParams struct {
	// List of rule IDs
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteAlertRulesRequest struct {
	*tchttp.BaseRequest
	
	// List of rule IDs
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteAlertRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlertRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlertRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlertRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlertRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlertRulesResponseParams `json:"Response"`
}

func (r *DeleteAlertRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlertRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExporterIntegrationRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteExporterIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteExporterIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExporterIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	delete(f, "Kind")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExporterIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExporterIntegrationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExporterIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExporterIntegrationResponseParams `json:"Response"`
}

func (r *DeleteExporterIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExporterIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaInstanceRequestParams struct {
	// Array of instance names
	InstanceIDs []*string `json:"InstanceIDs,omitnil,omitempty" name:"InstanceIDs"`
}

type DeleteGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance names
	InstanceIDs []*string `json:"InstanceIDs,omitnil,omitempty" name:"InstanceIDs"`
}

func (r *DeleteGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGrafanaInstanceResponseParams `json:"Response"`
}

func (r *DeleteGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaIntegrationRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration ID, such as “integration-abcd1234”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`
}

type DeleteGrafanaIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration ID, such as “integration-abcd1234”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`
}

func (r *DeleteGrafanaIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGrafanaIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaIntegrationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGrafanaIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGrafanaIntegrationResponseParams `json:"Response"`
}

func (r *DeleteGrafanaIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaNotificationChannelRequestParams struct {
	// Array of channel IDs, such as “nchannel-abcd1234”.
	ChannelIDs []*string `json:"ChannelIDs,omitnil,omitempty" name:"ChannelIDs"`

	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteGrafanaNotificationChannelRequest struct {
	*tchttp.BaseRequest
	
	// Array of channel IDs, such as “nchannel-abcd1234”.
	ChannelIDs []*string `json:"ChannelIDs,omitnil,omitempty" name:"ChannelIDs"`

	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteGrafanaNotificationChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaNotificationChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelIDs")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGrafanaNotificationChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaNotificationChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGrafanaNotificationChannelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGrafanaNotificationChannelResponseParams `json:"Response"`
}

func (r *DeleteGrafanaNotificationChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaNotificationChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyGroupRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID.
	GroupId []*int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeletePolicyGroupRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID.
	GroupId []*int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeletePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePolicyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeletePolicyGroupResponseParams `json:"Response"`
}

func (r *DeletePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertGroupsRequestParams struct {
	// prometheus instance id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert group IDs that needs to be deleted, such as alert-xxxxx.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DeletePrometheusAlertGroupsRequest struct {
	*tchttp.BaseRequest
	
	// prometheus instance id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert group IDs that needs to be deleted, such as alert-xxxxx.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *DeletePrometheusAlertGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusAlertGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusAlertGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusAlertGroupsResponseParams `json:"Response"`
}

func (r *DeletePrometheusAlertGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of alerting rule IDs
	AlertIds []*string `json:"AlertIds,omitnil,omitempty" name:"AlertIds"`

	// Alerting rule name
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DeletePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of alerting rule IDs
	AlertIds []*string `json:"AlertIds,omitnil,omitempty" name:"AlertIds"`

	// Alerting rule name
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DeletePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertIds")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *DeletePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusClusterAgentRequestParams struct {
	// Agent list
	Agents []*PrometheusAgentInfo `json:"Agents,omitnil,omitempty" name:"Agents"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeletePrometheusClusterAgentRequest struct {
	*tchttp.BaseRequest
	
	// Agent list
	Agents []*PrometheusAgentInfo `json:"Agents,omitnil,omitempty" name:"Agents"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeletePrometheusClusterAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusClusterAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agents")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusClusterAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusClusterAgentResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusClusterAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusClusterAgentResponseParams `json:"Response"`
}

func (r *DeletePrometheusClusterAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusClusterAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of names of the service monitors to be deleted
	ServiceMonitors []*string `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// List of names of the pod monitors to be deleted
	PodMonitors []*string `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// List of names of the raw jobs to be deleted
	RawJobs []*string `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`
}

type DeletePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of names of the service monitors to be deleted
	ServiceMonitors []*string `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// List of names of the pod monitors to be deleted
	PodMonitors []*string `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// List of names of the raw jobs to be deleted
	RawJobs []*string `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`
}

func (r *DeletePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusConfigResponseParams `json:"Response"`
}

func (r *DeletePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusRecordRuleYamlRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of recording rules
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DeletePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of recording rules
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DeletePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusRecordRuleYamlResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *DeletePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusScrapeJobsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// List of task IDs
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`
}

type DeletePrometheusScrapeJobsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// List of task IDs
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`
}

func (r *DeletePrometheusScrapeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusScrapeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "JobIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusScrapeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusScrapeJobsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusScrapeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusScrapeJobsResponseParams `json:"Response"`
}

func (r *DeletePrometheusScrapeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusScrapeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempRequestParams struct {
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeletePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeletePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTempResponseParams `json:"Response"`
}

func (r *DeletePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempSyncRequestParams struct {
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// List of unsynced objects
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type DeletePrometheusTempSyncRequest struct {
	*tchttp.BaseRequest
	
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// List of unsynced objects
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *DeletePrometheusTempSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTempSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempSyncResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusTempSyncResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTempSyncResponseParams `json:"Response"`
}

func (r *DeletePrometheusTempSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordingRulesRequestParams struct {
	// List of rule IDs
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteRecordingRulesRequest struct {
	*tchttp.BaseRequest
	
	// List of rule IDs
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteRecordingRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordingRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordingRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordingRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordingRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordingRulesResponseParams `json:"Response"`
}

func (r *DeleteRecordingRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordingRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSSOAccountRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User account ID, such as “10000000”.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User account ID, such as “10000000”.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSSOAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSSOAccountResponseParams `json:"Response"`
}

func (r *DeleteSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListAlarms struct {
	// Event type.
	// Note: This field may return null, indicating that no valid value was found.
	BusinessTypeDesc *string `json:"BusinessTypeDesc,omitnil,omitempty" name:"BusinessTypeDesc"`

	// Event type.
	// Note: This field may return null, indicating that no valid value was found.
	AccidentTypeDesc *string `json:"AccidentTypeDesc,omitnil,omitempty" name:"AccidentTypeDesc"`

	// ID of the event type. The value 1 indicates service issues. The value 2 indicates other subscriptions.
	// Note: This field may return null, indicating that no valid value was found.
	BusinessID *int64 `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`

	// Event status ID. The value 0 indicates that the event has been recovered. The value 1 indicates that the event has not been recovered.
	// Note: This field may return null, indicating that no valid value was found.
	EventStatus *int64 `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// Affected object.
	// Note: This field may return null, indicating that no valid value was found.
	AffectResource *string `json:"AffectResource,omitnil,omitempty" name:"AffectResource"`

	// Region where the event occurs.
	// Note: This field may return null, indicating that no valid value was found.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Time when the event occurs.
	// Note: This field may return null, indicating that no valid value was found.
	OccurTime *string `json:"OccurTime,omitnil,omitempty" name:"OccurTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid value was found.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type DescribeAccidentEventListRequestParams struct {
	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting rule by UpdateTime. Valid values: asc and desc.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitnil,omitempty" name:"UpdateTimeOrder"`

	// Sorting rule by OccurTime. Valid values: asc or desc. Sorting by UpdateTimeOrder takes priority.
	OccurTimeOrder *string `json:"OccurTimeOrder,omitnil,omitempty" name:"OccurTimeOrder"`

	// Filter by event type. The value 1 indicates service issues. The value 2 indicates other subscriptions.
	AccidentType []*int64 `json:"AccidentType,omitnil,omitempty" name:"AccidentType"`

	// Filter by event. The value 1 indicates CVM storage issues. The value 2 indicates CVM network connection issues. The value 3 indicates that the CVM has an exception. The value 202 indicates that an ISP network jitter occurs.
	AccidentEvent []*int64 `json:"AccidentEvent,omitnil,omitempty" name:"AccidentEvent"`

	// Filter by event status. The value 0 indicates that the event has been recovered. The value 1 indicates that the event has not been recovered.
	AccidentStatus []*int64 `json:"AccidentStatus,omitnil,omitempty" name:"AccidentStatus"`

	// Filter by region where the event occurs. The value gz indicates Guangzhou. The value sh indicates Shanghai.
	AccidentRegion []*string `json:"AccidentRegion,omitnil,omitempty" name:"AccidentRegion"`

	// Filter by affected resource, such as ins-19a06bka.
	AffectResource *string `json:"AffectResource,omitnil,omitempty" name:"AffectResource"`
}

type DescribeAccidentEventListRequest struct {
	*tchttp.BaseRequest
	
	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting rule by UpdateTime. Valid values: asc and desc.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitnil,omitempty" name:"UpdateTimeOrder"`

	// Sorting rule by OccurTime. Valid values: asc or desc. Sorting by UpdateTimeOrder takes priority.
	OccurTimeOrder *string `json:"OccurTimeOrder,omitnil,omitempty" name:"OccurTimeOrder"`

	// Filter by event type. The value 1 indicates service issues. The value 2 indicates other subscriptions.
	AccidentType []*int64 `json:"AccidentType,omitnil,omitempty" name:"AccidentType"`

	// Filter by event. The value 1 indicates CVM storage issues. The value 2 indicates CVM network connection issues. The value 3 indicates that the CVM has an exception. The value 202 indicates that an ISP network jitter occurs.
	AccidentEvent []*int64 `json:"AccidentEvent,omitnil,omitempty" name:"AccidentEvent"`

	// Filter by event status. The value 0 indicates that the event has been recovered. The value 1 indicates that the event has not been recovered.
	AccidentStatus []*int64 `json:"AccidentStatus,omitnil,omitempty" name:"AccidentStatus"`

	// Filter by region where the event occurs. The value gz indicates Guangzhou. The value sh indicates Shanghai.
	AccidentRegion []*string `json:"AccidentRegion,omitnil,omitempty" name:"AccidentRegion"`

	// Filter by affected resource, such as ins-19a06bka.
	AffectResource *string `json:"AffectResource,omitnil,omitempty" name:"AffectResource"`
}

func (r *DescribeAccidentEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccidentEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "UpdateTimeOrder")
	delete(f, "OccurTimeOrder")
	delete(f, "AccidentType")
	delete(f, "AccidentEvent")
	delete(f, "AccidentStatus")
	delete(f, "AccidentRegion")
	delete(f, "AffectResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccidentEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccidentEventListResponseParams struct {
	// Platform event list.
	// Note: This field may return null, indicating that no valid value was found.
	Alarms []*DescribeAccidentEventListAlarms `json:"Alarms,omitnil,omitempty" name:"Alarms"`

	// Total number of platform events.
	// Note: This field may return null, indicating that no valid value was found.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccidentEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccidentEventListResponseParams `json:"Response"`
}

func (r *DescribeAccidentEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccidentEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmEventsRequestParams struct {
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy type such as cvm_device, which is obtained through the `DescribeAllNamespaces` API
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Monitoring type, such as `MT_QCE`, which is set to default.
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

type DescribeAlarmEventsRequest struct {
	*tchttp.BaseRequest
	
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy type such as cvm_device, which is obtained through the `DescribeAllNamespaces` API
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Monitoring type, such as `MT_QCE`, which is set to default.
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

func (r *DescribeAlarmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Namespace")
	delete(f, "MonitorType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmEventsResponseParams struct {
	// Alarm event list
	Events []*AlarmEvent `json:"Events,omitnil,omitempty" name:"Events"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmEventsResponseParams `json:"Response"`
}

func (r *DescribeAlarmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmHistoriesRequestParams struct {
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Page number starting from 1. Default value: 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–100. Default value: 20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Sort by the first occurrence time in descending order by default. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Start time, which is the timestamp one day ago by default and the time when the alarm `FirstOccurTime` first occurs. An alarm record can be searched only if its `FirstOccurTime` is later than the `StartTime`.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp and the time when the alarm `FirstOccurTime` first occurs. An alarm record can be searched only if its `FirstOccurTime` is earlier than the `EndTime`.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter by monitor type. Valid values: "MT_QCE" (Tencent Cloud service monitoring), "MT_TAW" (application performance monitoring), "MT_RUM" (frontend performance monitoring), "MT_PROBE" (cloud automated testing). If this parameter is left empty, all types will be queried by default.
	MonitorTypes []*string `json:"MonitorTypes,omitnil,omitempty" name:"MonitorTypes"`

	// Filter by alarm object. Fuzzy search with string is supported
	AlarmObject *string `json:"AlarmObject,omitnil,omitempty" name:"AlarmObject"`

	// Filter by alarm status. Valid values: ALARM (not resolved), OK (resolved), NO_CONF (expired), NO_DATA (insufficient data). If this parameter is left empty, all will be queried by default
	AlarmStatus []*string `json:"AlarmStatus,omitnil,omitempty" name:"AlarmStatus"`

	// Filter by project ID. Valid values: `-1` (no project), `0` (default project)
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Filter by instance group ID
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitnil,omitempty" name:"InstanceGroupIds"`

	// Filter by policy type. Monitoring type and policy type are first-level and second-level filters respectively and both need to be passed in. For example, `[{"MonitorType": "MT_QCE", "Namespace": "cvm_device"}]`
	Namespaces []*MonitorTypeNamespace `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// Filter by metric name
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// Fuzzy search by policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Fuzzy search by alarm content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Search by recipient
	ReceiverUids []*int64 `json:"ReceiverUids,omitnil,omitempty" name:"ReceiverUids"`

	// Search by recipient group
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// Search by alarm policy ID list
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`
}

type DescribeAlarmHistoriesRequest struct {
	*tchttp.BaseRequest
	
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Page number starting from 1. Default value: 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–100. Default value: 20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Sort by the first occurrence time in descending order by default. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Start time, which is the timestamp one day ago by default and the time when the alarm `FirstOccurTime` first occurs. An alarm record can be searched only if its `FirstOccurTime` is later than the `StartTime`.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp and the time when the alarm `FirstOccurTime` first occurs. An alarm record can be searched only if its `FirstOccurTime` is earlier than the `EndTime`.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter by monitor type. Valid values: "MT_QCE" (Tencent Cloud service monitoring), "MT_TAW" (application performance monitoring), "MT_RUM" (frontend performance monitoring), "MT_PROBE" (cloud automated testing). If this parameter is left empty, all types will be queried by default.
	MonitorTypes []*string `json:"MonitorTypes,omitnil,omitempty" name:"MonitorTypes"`

	// Filter by alarm object. Fuzzy search with string is supported
	AlarmObject *string `json:"AlarmObject,omitnil,omitempty" name:"AlarmObject"`

	// Filter by alarm status. Valid values: ALARM (not resolved), OK (resolved), NO_CONF (expired), NO_DATA (insufficient data). If this parameter is left empty, all will be queried by default
	AlarmStatus []*string `json:"AlarmStatus,omitnil,omitempty" name:"AlarmStatus"`

	// Filter by project ID. Valid values: `-1` (no project), `0` (default project)
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Filter by instance group ID
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitnil,omitempty" name:"InstanceGroupIds"`

	// Filter by policy type. Monitoring type and policy type are first-level and second-level filters respectively and both need to be passed in. For example, `[{"MonitorType": "MT_QCE", "Namespace": "cvm_device"}]`
	Namespaces []*MonitorTypeNamespace `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// Filter by metric name
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// Fuzzy search by policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Fuzzy search by alarm content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Search by recipient
	ReceiverUids []*int64 `json:"ReceiverUids,omitnil,omitempty" name:"ReceiverUids"`

	// Search by recipient group
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// Search by alarm policy ID list
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`
}

func (r *DescribeAlarmHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Order")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MonitorTypes")
	delete(f, "AlarmObject")
	delete(f, "AlarmStatus")
	delete(f, "ProjectIds")
	delete(f, "InstanceGroupIds")
	delete(f, "Namespaces")
	delete(f, "MetricNames")
	delete(f, "PolicyName")
	delete(f, "Content")
	delete(f, "ReceiverUids")
	delete(f, "ReceiverGroups")
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmHistoriesResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Alarm record list
	Histories []*AlarmHistory `json:"Histories,omitnil,omitempty" name:"Histories"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmHistoriesResponseParams `json:"Response"`
}

func (r *DescribeAlarmHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmMetricsRequestParams struct {
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Monitor type filter. Valid values: MT_QCE (Tencent Cloud service monitoring)
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Alarm policy type such as cvm_device, which is obtained through the `DescribeAllNamespaces` API
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeAlarmMetricsRequest struct {
	*tchttp.BaseRequest
	
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Monitor type filter. Valid values: MT_QCE (Tencent Cloud service monitoring)
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Alarm policy type such as cvm_device, which is obtained through the `DescribeAllNamespaces` API
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeAlarmMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "MonitorType")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmMetricsResponseParams struct {
	// Alarm metric list
	Metrics []*Metric `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmMetricsResponseParams `json:"Response"`
}

func (r *DescribeAlarmMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticeCallbacksRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`
}

type DescribeAlarmNoticeCallbacksRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`
}

func (r *DescribeAlarmNoticeCallbacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeCallbacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticeCallbacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticeCallbacksResponseParams struct {
	// Alarm callback notification
	// Note: this field may return null, indicating that no valid values can be obtained.
	URLNotices []*URLNotice `json:"URLNotices,omitnil,omitempty" name:"URLNotices"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNoticeCallbacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNoticeCallbacksResponseParams `json:"Response"`
}

func (r *DescribeAlarmNoticeCallbacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeCallbacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticeRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`
}

type DescribeAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`
}

func (r *DescribeAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "NoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticeResponseParams struct {
	// Alarm notification template details
	Notice *AlarmNotice `json:"Notice,omitnil,omitempty" name:"Notice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNoticeResponseParams `json:"Response"`
}

func (r *DescribeAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticesRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Page number. Minimum value: 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–200
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Sort by update time. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Root account `uid`, which is used to create preset notifications
	OwnerUid *int64 `json:"OwnerUid,omitnil,omitempty" name:"OwnerUid"`

	// Alarm notification template name, which is used for fuzzy search
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter by recipient. The type of notified users should be selected for the alarm notification template. Valid values: USER (user), GROUP (user group). If this parameter is left empty, no filter by recipient will be performed
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Recipient object list
	UserIds []*int64 `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// Recipient group list
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Filter by notification template ID. If an empty array is passed in or if this parameter is left empty, the filter operation will not be performed.
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Filter templates by tag
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Schedule list
	OnCallFormIDs []*string `json:"OnCallFormIDs,omitnil,omitempty" name:"OnCallFormIDs"`
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Page number. Minimum value: 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–200
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Sort by update time. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Root account `uid`, which is used to create preset notifications
	OwnerUid *int64 `json:"OwnerUid,omitnil,omitempty" name:"OwnerUid"`

	// Alarm notification template name, which is used for fuzzy search
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter by recipient. The type of notified users should be selected for the alarm notification template. Valid values: USER (user), GROUP (user group). If this parameter is left empty, no filter by recipient will be performed
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Recipient object list
	UserIds []*int64 `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// Recipient group list
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Filter by notification template ID. If an empty array is passed in or if this parameter is left empty, the filter operation will not be performed.
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Filter templates by tag
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Schedule list
	OnCallFormIDs []*string `json:"OnCallFormIDs,omitnil,omitempty" name:"OnCallFormIDs"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Order")
	delete(f, "OwnerUid")
	delete(f, "Name")
	delete(f, "ReceiverType")
	delete(f, "UserIds")
	delete(f, "GroupIds")
	delete(f, "NoticeIds")
	delete(f, "Tags")
	delete(f, "OnCallFormIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticesResponseParams struct {
	// Total number of alarm notification templates
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Alarm notification template list
	Notices []*AlarmNotice `json:"Notices,omitnil,omitempty" name:"Notices"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNoticesResponseParams `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmPoliciesRequestParams struct {
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Page number starting from 1. Default value: 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–100. Default value: 20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy search by policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default
	MonitorTypes []*string `json:"MonitorTypes,omitnil,omitempty" name:"MonitorTypes"`

	// Filter by namespace. For the values of different policy types, please see:
	// [Policy Type List](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1)
	Namespaces []*string `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// The alarm object list, which is a JSON string. The outer array corresponds to multiple instances, and the inner array is the dimension of an object. For example, “CVM - Basic Monitor” can be written as:
	// `[ {"Dimensions": {"unInstanceId": "ins-qr8d555g"}}, {"Dimensions": {"unInstanceId": "ins-qr8d555h"}} ]`
	// You can also refer to the “Example 2” below.
	// 
	// For more information on the parameter samples of different Tencent Cloud services, see [Product Policy Type and Dimension Information](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1).
	// 
	// Note: If `1` is passed in for `NeedCorrespondence`, the relationship between a policy and an instance needs to be returned. You can pass in up to 20 alarm object dimensions to avoid request timeout.
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Search by recipient. You can get the user list with the API [ListUsers](https://intl.cloud.tencent.com/document/product/598/34587?from_cn_redirect=1) in “Cloud Access Management” or query the sub-user information with the API [GetUser](https://intl.cloud.tencent.com/document/product/598/34590?from_cn_redirect=1). The `Uid` field in the returned result should be entered here.
	ReceiverUids []*int64 `json:"ReceiverUids,omitnil,omitempty" name:"ReceiverUids"`

	// Search by recipient group. You can get the user group list with the API [ListGroups](https://intl.cloud.tencent.com/document/product/598/34589?from_cn_redirect=1) in “Cloud Access Management” or query the user group list where a sub-user is in with the API [ListGroupsForUser](https://intl.cloud.tencent.com/document/product/598/34588?from_cn_redirect=1). The `GroupId` field in the returned result should be entered here.
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// Filter by default policy. Valid values: DEFAULT (display default policy), NOT_DEFAULT (display non-default policies). If this parameter is left empty, all policies will be displayed
	PolicyType []*string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// Sort by field. For example, to sort by the last modification time, use Field: "UpdateTime".
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// Sort order. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// ID array of the policy project, which can be viewed on the following page:
	// [Project Management](https://console.cloud.tencent.com/project)
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// List of the notification template IDs, which can be obtained by querying the notification template list.
	// It can be queried with the API [DescribeAlarmNotices](https://intl.cloud.tencent.com/document/product/248/51280?from_cn_redirect=1).
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Filter by trigger condition. Valid values: STATIC (display policies with static threshold), DYNAMIC (display policies with dynamic threshold). If this parameter is left empty, all policies will be displayed
	RuleTypes []*string `json:"RuleTypes,omitnil,omitempty" name:"RuleTypes"`

	// Filter by alarm status. Valid values: [1]: enabled; [0]: disabled; [0, 1]: all
	Enable []*int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// If `1` is passed in, alarm policies with no notification rules configured are queried. If it is left empty or other values are passed in, all alarm policies are queried.
	NotBindingNoticeRule *int64 `json:"NotBindingNoticeRule,omitnil,omitempty" name:"NotBindingNoticeRule"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Whether the relationship between a policy and the input parameter filter dimension is required. `1`: Yes. `0`: No. Default value: `0`.
	NeedCorrespondence *int64 `json:"NeedCorrespondence,omitnil,omitempty" name:"NeedCorrespondence"`

	// Filter alarm policy by triggered task (such as auto scaling task). Up to 10 tasks can be specified.
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitnil,omitempty" name:"TriggerTasks"`

	// Filter by quick alarm policy. If this parameter is left empty, all policies are displayed. `ONECLICK`: Display quick alarm policies; `NOT_ONECLICK`: Display non-quick alarm policies.
	OneClickPolicyType []*string `json:"OneClickPolicyType,omitnil,omitempty" name:"OneClickPolicyType"`

	// Whether the returned result needs to filter policies associated with all objects. Valid values: `1` (Yes), `0` (No).
	NotBindAll *int64 `json:"NotBindAll,omitnil,omitempty" name:"NotBindAll"`

	// Whether the returned result needs to filter policies associated with instance groups. Valid values: `1` (Yes), `0` (No).
	NotInstanceGroup *int64 `json:"NotInstanceGroup,omitnil,omitempty" name:"NotInstanceGroup"`

	// Filter policies by tag
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// ID of the TencentCloud Managed Service for Prometheus instance, which is used for customizing a metric policy.
	PromInsId *string `json:"PromInsId,omitnil,omitempty" name:"PromInsId"`

	// Search by schedule
	ReceiverOnCallFormIDs []*string `json:"ReceiverOnCallFormIDs,omitnil,omitempty" name:"ReceiverOnCallFormIDs"`
}

type DescribeAlarmPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Page number starting from 1. Default value: 1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of entries per page. Value range: 1–100. Default value: 20
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy search by policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default
	MonitorTypes []*string `json:"MonitorTypes,omitnil,omitempty" name:"MonitorTypes"`

	// Filter by namespace. For the values of different policy types, please see:
	// [Policy Type List](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1)
	Namespaces []*string `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// The alarm object list, which is a JSON string. The outer array corresponds to multiple instances, and the inner array is the dimension of an object. For example, “CVM - Basic Monitor” can be written as:
	// `[ {"Dimensions": {"unInstanceId": "ins-qr8d555g"}}, {"Dimensions": {"unInstanceId": "ins-qr8d555h"}} ]`
	// You can also refer to the “Example 2” below.
	// 
	// For more information on the parameter samples of different Tencent Cloud services, see [Product Policy Type and Dimension Information](https://intl.cloud.tencent.com/document/product/248/50397?from_cn_redirect=1).
	// 
	// Note: If `1` is passed in for `NeedCorrespondence`, the relationship between a policy and an instance needs to be returned. You can pass in up to 20 alarm object dimensions to avoid request timeout.
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Search by recipient. You can get the user list with the API [ListUsers](https://intl.cloud.tencent.com/document/product/598/34587?from_cn_redirect=1) in “Cloud Access Management” or query the sub-user information with the API [GetUser](https://intl.cloud.tencent.com/document/product/598/34590?from_cn_redirect=1). The `Uid` field in the returned result should be entered here.
	ReceiverUids []*int64 `json:"ReceiverUids,omitnil,omitempty" name:"ReceiverUids"`

	// Search by recipient group. You can get the user group list with the API [ListGroups](https://intl.cloud.tencent.com/document/product/598/34589?from_cn_redirect=1) in “Cloud Access Management” or query the user group list where a sub-user is in with the API [ListGroupsForUser](https://intl.cloud.tencent.com/document/product/598/34588?from_cn_redirect=1). The `GroupId` field in the returned result should be entered here.
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// Filter by default policy. Valid values: DEFAULT (display default policy), NOT_DEFAULT (display non-default policies). If this parameter is left empty, all policies will be displayed
	PolicyType []*string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// Sort by field. For example, to sort by the last modification time, use Field: "UpdateTime".
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// Sort order. Valid values: ASC (ascending), DESC (descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// ID array of the policy project, which can be viewed on the following page:
	// [Project Management](https://console.cloud.tencent.com/project)
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// List of the notification template IDs, which can be obtained by querying the notification template list.
	// It can be queried with the API [DescribeAlarmNotices](https://intl.cloud.tencent.com/document/product/248/51280?from_cn_redirect=1).
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Filter by trigger condition. Valid values: STATIC (display policies with static threshold), DYNAMIC (display policies with dynamic threshold). If this parameter is left empty, all policies will be displayed
	RuleTypes []*string `json:"RuleTypes,omitnil,omitempty" name:"RuleTypes"`

	// Filter by alarm status. Valid values: [1]: enabled; [0]: disabled; [0, 1]: all
	Enable []*int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// If `1` is passed in, alarm policies with no notification rules configured are queried. If it is left empty or other values are passed in, all alarm policies are queried.
	NotBindingNoticeRule *int64 `json:"NotBindingNoticeRule,omitnil,omitempty" name:"NotBindingNoticeRule"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Whether the relationship between a policy and the input parameter filter dimension is required. `1`: Yes. `0`: No. Default value: `0`.
	NeedCorrespondence *int64 `json:"NeedCorrespondence,omitnil,omitempty" name:"NeedCorrespondence"`

	// Filter alarm policy by triggered task (such as auto scaling task). Up to 10 tasks can be specified.
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitnil,omitempty" name:"TriggerTasks"`

	// Filter by quick alarm policy. If this parameter is left empty, all policies are displayed. `ONECLICK`: Display quick alarm policies; `NOT_ONECLICK`: Display non-quick alarm policies.
	OneClickPolicyType []*string `json:"OneClickPolicyType,omitnil,omitempty" name:"OneClickPolicyType"`

	// Whether the returned result needs to filter policies associated with all objects. Valid values: `1` (Yes), `0` (No).
	NotBindAll *int64 `json:"NotBindAll,omitnil,omitempty" name:"NotBindAll"`

	// Whether the returned result needs to filter policies associated with instance groups. Valid values: `1` (Yes), `0` (No).
	NotInstanceGroup *int64 `json:"NotInstanceGroup,omitnil,omitempty" name:"NotInstanceGroup"`

	// Filter policies by tag
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// ID of the TencentCloud Managed Service for Prometheus instance, which is used for customizing a metric policy.
	PromInsId *string `json:"PromInsId,omitnil,omitempty" name:"PromInsId"`

	// Search by schedule
	ReceiverOnCallFormIDs []*string `json:"ReceiverOnCallFormIDs,omitnil,omitempty" name:"ReceiverOnCallFormIDs"`
}

func (r *DescribeAlarmPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "PolicyName")
	delete(f, "MonitorTypes")
	delete(f, "Namespaces")
	delete(f, "Dimensions")
	delete(f, "ReceiverUids")
	delete(f, "ReceiverGroups")
	delete(f, "PolicyType")
	delete(f, "Field")
	delete(f, "Order")
	delete(f, "ProjectIds")
	delete(f, "NoticeIds")
	delete(f, "RuleTypes")
	delete(f, "Enable")
	delete(f, "NotBindingNoticeRule")
	delete(f, "InstanceGroupId")
	delete(f, "NeedCorrespondence")
	delete(f, "TriggerTasks")
	delete(f, "OneClickPolicyType")
	delete(f, "NotBindAll")
	delete(f, "NotInstanceGroup")
	delete(f, "Tags")
	delete(f, "PromInsId")
	delete(f, "ReceiverOnCallFormIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmPoliciesResponseParams struct {
	// Total number of policies
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Policy array
	Policies []*AlarmPolicy `json:"Policies,omitnil,omitempty" name:"Policies"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAlarmPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmPolicyRequestParams struct {
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmPolicyResponseParams struct {
	// Policy details
	Policy *AlarmPolicy `json:"Policy,omitnil,omitempty" name:"Policy"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmPolicyResponseParams `json:"Response"`
}

func (r *DescribeAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertRulesRequestParams struct {
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status code. Valid values:
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Alerting rule template category
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeAlertRulesRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status code. Valid values:
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Alerting rule template category
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeAlertRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleId")
	delete(f, "RuleState")
	delete(f, "RuleName")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlertRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertRulesResponseParams struct {
	// Number of alerting rules
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Alerting rule details
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlertRuleSet []*PrometheusRuleSet `json:"AlertRuleSet,omitnil,omitempty" name:"AlertRuleSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlertRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlertRulesResponseParams `json:"Response"`
}

func (r *DescribeAlertRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllNamespacesRequestParams struct {
	// Filter by use case. Currently, the only valid value is `ST_ALARM` (alarm type).
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default
	MonitorTypes []*string `json:"MonitorTypes,omitnil,omitempty" name:"MonitorTypes"`

	// Filter by namespace ID. If this parameter is left empty, all will be queried
	Ids []*string `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DescribeAllNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// Filter by use case. Currently, the only valid value is `ST_ALARM` (alarm type).
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Value fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default
	MonitorTypes []*string `json:"MonitorTypes,omitnil,omitempty" name:"MonitorTypes"`

	// Filter by namespace ID. If this parameter is left empty, all will be queried
	Ids []*string `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DescribeAllNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneType")
	delete(f, "Module")
	delete(f, "MonitorTypes")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllNamespacesResponseParams struct {
	// Alarm policy type of Tencent Cloud service (disused)
	QceNamespaces *CommonNamespace `json:"QceNamespaces,omitnil,omitempty" name:"QceNamespaces"`

	// Other alarm policy type (disused)
	CustomNamespaces *CommonNamespace `json:"CustomNamespaces,omitnil,omitempty" name:"CustomNamespaces"`

	// Alarm policy type of Tencent Cloud service
	QceNamespacesNew []*CommonNamespace `json:"QceNamespacesNew,omitnil,omitempty" name:"QceNamespacesNew"`

	// Other alarm policy type (not supported currently)
	CustomNamespacesNew []*CommonNamespace `json:"CustomNamespacesNew,omitnil,omitempty" name:"CustomNamespacesNew"`

	// General alarm policy type, including TAPM, RUM, and CAT.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CommonNamespaces []*CommonNamespaceNew `json:"CommonNamespaces,omitnil,omitempty" name:"CommonNamespaces"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllNamespacesResponseParams `json:"Response"`
}

func (r *DescribeAllNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseMetricsRequestParams struct {
	// Service namespace. Tencent Cloud services have different namespaces. For more information on service namespaces, see the monitoring metric documentation of each service. For example, see [CVM Monitoring Metrics](https://intl.cloud.tencent.com/document/product/248/6843?from_cn_redirect=1) for the namespace of CVM
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric name. Tencent Cloud services have different metric names. For more information on metric names, see the monitoring metric documentation of each service. For example, see [CVM Monitoring Metrics](https://intl.cloud.tencent.com/document/product/248/6843?from_cn_redirect=1) for the metric names of CVM
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Filter by dimension. This parameter is optional.
	Dimensions []*string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest
	
	// Service namespace. Tencent Cloud services have different namespaces. For more information on service namespaces, see the monitoring metric documentation of each service. For example, see [CVM Monitoring Metrics](https://intl.cloud.tencent.com/document/product/248/6843?from_cn_redirect=1) for the namespace of CVM
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric name. Tencent Cloud services have different metric names. For more information on metric names, see the monitoring metric documentation of each service. For example, see [CVM Monitoring Metrics](https://intl.cloud.tencent.com/document/product/248/6843?from_cn_redirect=1) for the metric names of CVM
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Filter by dimension. This parameter is optional.
	Dimensions []*string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

func (r *DescribeBaseMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	delete(f, "MetricName")
	delete(f, "Dimensions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaseMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseMetricsResponseParams struct {
	// Listed of queried metric descriptions
	MetricSet []*MetricSet `json:"MetricSet,omitnil,omitempty" name:"MetricSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaseMetricsResponseParams `json:"Response"`
}

func (r *DescribeBaseMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListAlarms struct {
	// Alarm ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Project ID.
	// Note: This field may return null, indicating that no valid value was found.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name.
	// Note: This field may return null, indicating that no valid value was found.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Alarm status ID. Valid values: 0 (not resolved), 1 (resolved), 2/3/5 (insufficient data), 4 (expired)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Alarm status. Valid values: ALARM (not resolved), OK (resolved), NO_DATA (insufficient data), NO_CONF (expired)
	// Note: this field may return null, indicating that no valid values can be obtained.
	AlarmStatus *string `json:"AlarmStatus,omitnil,omitempty" name:"AlarmStatus"`

	// Policy group ID.
	// Note: This field may return null, indicating that no valid value was found.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Policy group name.
	// Note: This field may return null, indicating that no valid value was found.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Occurrence time.
	// Note: This field may return null, indicating that no valid value was found.
	FirstOccurTime *string `json:"FirstOccurTime,omitnil,omitempty" name:"FirstOccurTime"`

	// Duration in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// End time.
	// Note: This field may return null, indicating that no valid value was found.
	LastOccurTime *string `json:"LastOccurTime,omitnil,omitempty" name:"LastOccurTime"`

	// Alarm content.
	// Note: This field may return null, indicating that no valid value was found.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Alarm object.
	// Note: This field may return null, indicating that no valid value was found.
	ObjName *string `json:"ObjName,omitnil,omitempty" name:"ObjName"`

	// Alarm object ID.
	// Note: This field may return null, indicating that no valid value was found.
	ObjId *string `json:"ObjId,omitnil,omitempty" name:"ObjId"`

	// Policy type.
	// Note: This field may return null, indicating that no valid value was found.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// VPC, which is unique to CVM.
	// Note: This field may return null, indicating that no valid value was found.
	Vpc *string `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// Metric ID.
	// Note: This field may return null, indicating that no valid value was found.
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// Metric name.
	// Note: This field may return null, indicating that no valid value was found.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Alarm type. The value 0 indicates metric alarms. The value 2 indicates product event alarms. The value 3 indicates platform event alarms.
	// Note: This field may return null, indicating that no valid value was found.
	AlarmType *int64 `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// Region.
	// Note: This field may return null, indicating that no valid value was found.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Dimensions of an alarm object.
	// Note: This field may return null, indicating that no valid value was found.
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Notification method.
	// Note: This field may return null, indicating that no valid value was found.
	NotifyWay []*string `json:"NotifyWay,omitnil,omitempty" name:"NotifyWay"`

	// Instance group information.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroup []*InstanceGroup `json:"InstanceGroup,omitnil,omitempty" name:"InstanceGroup"`
}

// Predefined struct for user
type DescribeBasicAlarmListRequestParams struct {
	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting by occurrence time. Valid values: asc and desc.
	OccurTimeOrder *string `json:"OccurTimeOrder,omitnil,omitempty" name:"OccurTimeOrder"`

	// Filter by project ID.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Filter by policy type.
	ViewNames []*string `json:"ViewNames,omitnil,omitempty" name:"ViewNames"`

	// Filter by alarm status.
	AlarmStatus []*int64 `json:"AlarmStatus,omitnil,omitempty" name:"AlarmStatus"`

	// Filter by alarm object.
	ObjLike *string `json:"ObjLike,omitnil,omitempty" name:"ObjLike"`

	// Filter by instance group ID.
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitnil,omitempty" name:"InstanceGroupIds"`

	// Filtering by metric names
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`
}

type DescribeBasicAlarmListRequest struct {
	*tchttp.BaseRequest
	
	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting by occurrence time. Valid values: asc and desc.
	OccurTimeOrder *string `json:"OccurTimeOrder,omitnil,omitempty" name:"OccurTimeOrder"`

	// Filter by project ID.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Filter by policy type.
	ViewNames []*string `json:"ViewNames,omitnil,omitempty" name:"ViewNames"`

	// Filter by alarm status.
	AlarmStatus []*int64 `json:"AlarmStatus,omitnil,omitempty" name:"AlarmStatus"`

	// Filter by alarm object.
	ObjLike *string `json:"ObjLike,omitnil,omitempty" name:"ObjLike"`

	// Filter by instance group ID.
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitnil,omitempty" name:"InstanceGroupIds"`

	// Filtering by metric names
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`
}

func (r *DescribeBasicAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicAlarmListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OccurTimeOrder")
	delete(f, "ProjectIds")
	delete(f, "ViewNames")
	delete(f, "AlarmStatus")
	delete(f, "ObjLike")
	delete(f, "InstanceGroupIds")
	delete(f, "MetricNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicAlarmListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicAlarmListResponseParams struct {
	// Alarm list.
	// Note: This field may return null, indicating that no valid value was found.
	Alarms []*DescribeBasicAlarmListAlarms `json:"Alarms,omitnil,omitempty" name:"Alarms"`

	// Total number.
	// Note: This field may return null, indicating that no valid value was found.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Warning *string `json:"Warning,omitnil,omitempty" name:"Warning"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBasicAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicAlarmListResponseParams `json:"Response"`
}

func (r *DescribeBasicAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicAlarmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListDimension struct {
	// Region ID.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region abbreviation.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Combined JSON string of dimensions.
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Combined JSON string of event dimensions.
	EventDimensions *string `json:"EventDimensions,omitnil,omitempty" name:"EventDimensions"`
}

type DescribeBindingPolicyObjectListInstance struct {
	// Unique ID of the object.
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`

	// Dimension set of an object instance, which is a jsonObj string.
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Whether the object is shielded. The value 0 indicates that the object is not shielded. The value 1 indicates that the object is shielded.
	IsShielded *int64 `json:"IsShielded,omitnil,omitempty" name:"IsShielded"`

	// Region where the object resides.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type DescribeBindingPolicyObjectListInstanceGroup struct {
	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Alarm policy type name.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Instance group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Number of instances.
	InstanceSum *int64 `json:"InstanceSum,omitnil,omitempty" name:"InstanceSum"`

	// Update time.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Regions where the instances reside.
	// Note: This field may return null, indicating that no valid value was found.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

// Predefined struct for user
type DescribeBindingPolicyObjectListRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID. If the ID is in the format of “policy-xxxx”, please enter it in the `PolicyId` field. Enter 0 in this field.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm policy ID in the format of “policy-xxxx”. If a value has been entered in this field, you can enter 0 in the `GroupId` field.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The number of alarm objects returned each time. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 and is set to 0 by default. For example, the parameter `Offset=0&Limit=20` returns the zeroth to 19th alarm objects, and `Offset=20&Limit=20` returns the 20th to 39th alarm objects, and so on.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Dimensions of filtering objects.
	Dimensions []*DescribeBindingPolicyObjectListDimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type DescribeBindingPolicyObjectListRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID. If the ID is in the format of “policy-xxxx”, please enter it in the `PolicyId` field. Enter 0 in this field.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm policy ID in the format of “policy-xxxx”. If a value has been entered in this field, you can enter 0 in the `GroupId` field.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The number of alarm objects returned each time. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 and is set to 0 by default. For example, the parameter `Offset=0&Limit=20` returns the zeroth to 19th alarm objects, and `Offset=20&Limit=20` returns the 20th to 39th alarm objects, and so on.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Dimensions of filtering objects.
	Dimensions []*DescribeBindingPolicyObjectListDimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

func (r *DescribeBindingPolicyObjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindingPolicyObjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "PolicyId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Dimensions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindingPolicyObjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindingPolicyObjectListResponseParams struct {
	// List of bound object instances.
	// Note: This field may return null, indicating that no valid value was found.
	List []*DescribeBindingPolicyObjectListInstance `json:"List,omitnil,omitempty" name:"List"`

	// Total number of bound object instances.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of object instances that are not shielded.
	NoShieldedSum *int64 `json:"NoShieldedSum,omitnil,omitempty" name:"NoShieldedSum"`

	// Bound instance group information. This parameter is not configured if no instance group is bound.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroup *DescribeBindingPolicyObjectListInstanceGroup `json:"InstanceGroup,omitnil,omitempty" name:"InstanceGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBindingPolicyObjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindingPolicyObjectListResponseParams `json:"Response"`
}

func (r *DescribeBindingPolicyObjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindingPolicyObjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAgentCreatingProgressRequestParams struct {

}

type DescribeClusterAgentCreatingProgressRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeClusterAgentCreatingProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAgentCreatingProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAgentCreatingProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAgentCreatingProgressResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterAgentCreatingProgressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAgentCreatingProgressResponseParams `json:"Response"`
}

func (r *DescribeClusterAgentCreatingProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAgentCreatingProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConditionsTemplateListRequestParams struct {
	// The value is fixed to `monitor`.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// View name, which can be obtained via [DescribeAllNamespaces](https://intl.cloud.tencent.com/document/product/248/48683?from_cn_redirect=1). For the monitoring of Tencent Cloud services, the value of this parameter is `QceNamespacesNew.N.Id` of the output parameter of `DescribeAllNamespaces`, for example, `cvm_device`.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Filter by trigger condition template name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Filter by trigger condition template ID.
	GroupID *string `json:"GroupID,omitnil,omitempty" name:"GroupID"`

	// Pagination parameter, which specifies the number of returned results per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset starting from 0. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting method by update time. `asc`: Ascending order; `desc`: Descending order.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitnil,omitempty" name:"UpdateTimeOrder"`

	// Sorting order based on the number of associated policies. Valid values: `asc` (ascending order), `desc` (descending order).
	PolicyCountOrder *string `json:"PolicyCountOrder,omitnil,omitempty" name:"PolicyCountOrder"`
}

type DescribeConditionsTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to `monitor`.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// View name, which can be obtained via [DescribeAllNamespaces](https://intl.cloud.tencent.com/document/product/248/48683?from_cn_redirect=1). For the monitoring of Tencent Cloud services, the value of this parameter is `QceNamespacesNew.N.Id` of the output parameter of `DescribeAllNamespaces`, for example, `cvm_device`.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Filter by trigger condition template name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Filter by trigger condition template ID.
	GroupID *string `json:"GroupID,omitnil,omitempty" name:"GroupID"`

	// Pagination parameter, which specifies the number of returned results per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset starting from 0. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting method by update time. `asc`: Ascending order; `desc`: Descending order.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitnil,omitempty" name:"UpdateTimeOrder"`

	// Sorting order based on the number of associated policies. Valid values: `asc` (ascending order), `desc` (descending order).
	PolicyCountOrder *string `json:"PolicyCountOrder,omitnil,omitempty" name:"PolicyCountOrder"`
}

func (r *DescribeConditionsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConditionsTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "ViewName")
	delete(f, "GroupName")
	delete(f, "GroupID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "UpdateTimeOrder")
	delete(f, "PolicyCountOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConditionsTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConditionsTemplateListResponseParams struct {
	// Total number of templates.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Template list.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TemplateGroupList []*TemplateGroup `json:"TemplateGroupList,omitnil,omitempty" name:"TemplateGroupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConditionsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConditionsTemplateListResponseParams `json:"Response"`
}

func (r *DescribeConditionsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConditionsTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDNSConfigRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDNSConfigRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDNSConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDNSConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDNSConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDNSConfigResponseParams struct {
	// Array of DNS servers
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDNSConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDNSConfigResponseParams `json:"Response"`
}

func (r *DescribeDNSConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDNSConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExporterIntegrationsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeExporterIntegrationsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeExporterIntegrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExporterIntegrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	delete(f, "Kind")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExporterIntegrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExporterIntegrationsResponseParams struct {
	// List of integration configurations
	IntegrationSet []*IntegrationConfiguration `json:"IntegrationSet,omitnil,omitempty" name:"IntegrationSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExporterIntegrationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExporterIntegrationsResponseParams `json:"Response"`
}

func (r *DescribeExporterIntegrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExporterIntegrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaChannelsRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items to be queried
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Alert channel ID, such as “nchannel-abcd1234”.
	ChannelIds []*string `json:"ChannelIds,omitnil,omitempty" name:"ChannelIds"`

	// Alert channel status
	ChannelState *int64 `json:"ChannelState,omitnil,omitempty" name:"ChannelState"`
}

type DescribeGrafanaChannelsRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items to be queried
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Alert channel ID, such as “nchannel-abcd1234”.
	ChannelIds []*string `json:"ChannelIds,omitnil,omitempty" name:"ChannelIds"`

	// Alert channel status
	ChannelState *int64 `json:"ChannelState,omitnil,omitempty" name:"ChannelState"`
}

func (r *DescribeGrafanaChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ChannelName")
	delete(f, "ChannelIds")
	delete(f, "ChannelState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaChannelsResponseParams struct {
	// Array of alert channels
	NotificationChannelSet []*GrafanaChannel `json:"NotificationChannelSet,omitnil,omitempty" name:"NotificationChannelSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGrafanaChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaChannelsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaConfigRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeGrafanaConfigRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeGrafanaConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaConfigResponseParams struct {
	// JSON-encoded string
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGrafanaConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaConfigResponseParams `json:"Response"`
}

func (r *DescribeGrafanaConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaEnvironmentsRequestParams struct {
	// ID of a TencentCloud Managed Service for Grafana instance, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeGrafanaEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// ID of a TencentCloud Managed Service for Grafana instance, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeGrafanaEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaEnvironmentsResponseParams struct {
	// Environment variable string
	Envs *string `json:"Envs,omitnil,omitempty" name:"Envs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGrafanaEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaInstancesRequestParams struct {
	// Offset for query
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items to be queried
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Array of TCMG instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// TCMG instance name, which can be fuzzily matched by prefix.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Query status
	InstanceStatus []*int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Array of tag filters
	TagFilters []*PrometheusTag `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeGrafanaInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Offset for query
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items to be queried
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Array of TCMG instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// TCMG instance name, which can be fuzzily matched by prefix.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Query status
	InstanceStatus []*int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Array of tag filters
	TagFilters []*PrometheusTag `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeGrafanaInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIds")
	delete(f, "InstanceName")
	delete(f, "InstanceStatus")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaInstancesResponseParams struct {
	// This parameter has been disused. Use `Instances` instead.
	InstanceSet []*GrafanaInstanceInfo `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// Number of eligible instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instances
	Instances []*GrafanaInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGrafanaInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaInstancesResponseParams `json:"Response"`
}

func (r *DescribeGrafanaInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaIntegrationsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration ID
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`
}

type DescribeGrafanaIntegrationsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration ID
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`
}

func (r *DescribeGrafanaIntegrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaIntegrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationId")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaIntegrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaIntegrationsResponseParams struct {
	// Array of integrations
	IntegrationSet []*GrafanaIntegrationConfig `json:"IntegrationSet,omitnil,omitempty" name:"IntegrationSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGrafanaIntegrationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaIntegrationsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaIntegrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaIntegrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaNotificationChannelsRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items to be queried
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Alert channel ID, such as “nchannel-abcd1234”.
	ChannelIDs []*string `json:"ChannelIDs,omitnil,omitempty" name:"ChannelIDs"`

	// Alert channel status
	ChannelState *int64 `json:"ChannelState,omitnil,omitempty" name:"ChannelState"`
}

type DescribeGrafanaNotificationChannelsRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items to be queried
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Alert channel ID, such as “nchannel-abcd1234”.
	ChannelIDs []*string `json:"ChannelIDs,omitnil,omitempty" name:"ChannelIDs"`

	// Alert channel status
	ChannelState *int64 `json:"ChannelState,omitnil,omitempty" name:"ChannelState"`
}

func (r *DescribeGrafanaNotificationChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaNotificationChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ChannelName")
	delete(f, "ChannelIDs")
	delete(f, "ChannelState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaNotificationChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaNotificationChannelsResponseParams struct {
	// Array of notification channels
	NotificationChannelSet []*GrafanaNotificationChannel `json:"NotificationChannelSet,omitnil,omitempty" name:"NotificationChannelSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGrafanaNotificationChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaNotificationChannelsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaNotificationChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaNotificationChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaWhiteListRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeGrafanaWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeGrafanaWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaWhiteListResponseParams struct {
	// Array
	WhiteList []*string `json:"WhiteList,omitnil,omitempty" name:"WhiteList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGrafanaWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaWhiteListResponseParams `json:"Response"`
}

func (r *DescribeGrafanaWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstalledPluginsRequestParams struct {
	// TCMG instance ID, such as “grafana-kleu3gt0”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by plugin ID such as “grafana-piechart-panel”. You can view the IDs of installed plugins through the `DescribeInstalledPlugins` API.
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`
}

type DescribeInstalledPluginsRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-kleu3gt0”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by plugin ID such as “grafana-piechart-panel”. You can view the IDs of installed plugins through the `DescribeInstalledPlugins` API.
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`
}

func (r *DescribeInstalledPluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstalledPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PluginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstalledPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstalledPluginsResponseParams struct {
	// List of plugins
	PluginSet []*GrafanaPlugin `json:"PluginSet,omitnil,omitempty" name:"PluginSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstalledPluginsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstalledPluginsResponseParams `json:"Response"`
}

func (r *DescribeInstalledPluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstalledPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorTypesRequestParams struct {
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`
}

type DescribeMonitorTypesRequest struct {
	*tchttp.BaseRequest
	
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`
}

func (r *DescribeMonitorTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorTypesResponseParams struct {
	// Monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring)
	MonitorTypes []*string `json:"MonitorTypes,omitnil,omitempty" name:"MonitorTypes"`

	// Monitoring type details
	MonitorTypeInfos []*MonitorTypeInfo `json:"MonitorTypeInfos,omitnil,omitempty" name:"MonitorTypeInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMonitorTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorTypesResponseParams `json:"Response"`
}

func (r *DescribeMonitorTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListCondition struct {
	// Policy view name.
	PolicyViewName *string `json:"PolicyViewName,omitnil,omitempty" name:"PolicyViewName"`

	// Event alarm conditions.
	// Note: This field may return null, indicating that no valid value was found.
	EventMetrics []*DescribePolicyConditionListEventMetric `json:"EventMetrics,omitnil,omitempty" name:"EventMetrics"`

	// Whether to support multiple regions.
	IsSupportMultiRegion *bool `json:"IsSupportMultiRegion,omitnil,omitempty" name:"IsSupportMultiRegion"`

	// Metric alarm conditions.
	// Note: This field may return null, indicating that no valid value was found.
	Metrics []*DescribePolicyConditionListMetric `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Policy type name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sorting ID.
	SortId *int64 `json:"SortId,omitnil,omitempty" name:"SortId"`

	// Whether to support default policies.
	SupportDefault *bool `json:"SupportDefault,omitnil,omitempty" name:"SupportDefault"`

	// List of regions that support this policy type.
	// Note: This field may return null, indicating that no valid value was found.
	SupportRegions []*string `json:"SupportRegions,omitnil,omitempty" name:"SupportRegions"`

	// Deprecated information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeprecatingInfo *DescribePolicyConditionListResponseDeprecatingInfo `json:"DeprecatingInfo,omitnil,omitempty" name:"DeprecatingInfo"`
}

type DescribePolicyConditionListConfigManual struct {
	// Check method.
	// Note: This field may return null, indicating that no valid value was found.
	CalcType *DescribePolicyConditionListConfigManualCalcType `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// Threshold.
	// Note: This field may return null, indicating that no valid value was found.
	CalcValue *DescribePolicyConditionListConfigManualCalcValue `json:"CalcValue,omitnil,omitempty" name:"CalcValue"`

	// Duration.
	// Note: This field may return null, indicating that no valid value was found.
	ContinueTime *DescribePolicyConditionListConfigManualContinueTime `json:"ContinueTime,omitnil,omitempty" name:"ContinueTime"`

	// Data period.
	// Note: This field may return null, indicating that no valid value was found.
	Period *DescribePolicyConditionListConfigManualPeriod `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of periods.
	// Note: This field may return null, indicating that no valid value was found.
	PeriodNum *DescribePolicyConditionListConfigManualPeriodNum `json:"PeriodNum,omitnil,omitempty" name:"PeriodNum"`

	// Statistics method.
	// Note: This field may return null, indicating that no valid value was found.
	StatType *DescribePolicyConditionListConfigManualStatType `json:"StatType,omitnil,omitempty" name:"StatType"`
}

type DescribePolicyConditionListConfigManualCalcType struct {
	// Value of CalcType.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Required or not.
	Need *bool `json:"Need,omitnil,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualCalcValue struct {
	// Default value.
	// Note: This field may return null, indicating that no valid value was found.
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Fixed value.
	// Note: This field may return null, indicating that no valid value was found.
	Fixed *string `json:"Fixed,omitnil,omitempty" name:"Fixed"`

	// Maximum value.
	// Note: This field may return null, indicating that no valid value was found.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value.
	// Note: This field may return null, indicating that no valid value was found.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Required or not.
	Need *bool `json:"Need,omitnil,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualContinueTime struct {
	// Default duration in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitnil,omitempty" name:"Default"`

	// Custom durations in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Required or not.
	Need *bool `json:"Need,omitnil,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriod struct {
	// Default period in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitnil,omitempty" name:"Default"`

	// Custom periods in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Required or not.
	Need *bool `json:"Need,omitnil,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriodNum struct {
	// Number of default periods.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitnil,omitempty" name:"Default"`

	// Number of custom periods.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Required or not.
	Need *bool `json:"Need,omitnil,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualStatType struct {
	// Data aggregation method in a period of 5 seconds.
	// Note: This field may return null, indicating that no valid value was found.
	P5 *string `json:"P5,omitnil,omitempty" name:"P5"`

	// Data aggregation method in a period of 10 seconds.
	// Note: This field may return null, indicating that no valid value was found.
	P10 *string `json:"P10,omitnil,omitempty" name:"P10"`

	// Data aggregation method in a period of 1 minute.
	// Note: This field may return null, indicating that no valid value was found.
	P60 *string `json:"P60,omitnil,omitempty" name:"P60"`

	// Data aggregation method in a period of 5 minutes.
	// Note: This field may return null, indicating that no valid value was found.
	P300 *string `json:"P300,omitnil,omitempty" name:"P300"`

	// Data aggregation method in a period of 10 minutes.
	// Note: This field may return null, indicating that no valid value was found.
	P600 *string `json:"P600,omitnil,omitempty" name:"P600"`

	// Data aggregation method in a period of 30 minutes.
	// Note: This field may return null, indicating that no valid value was found.
	P1800 *string `json:"P1800,omitnil,omitempty" name:"P1800"`

	// Data aggregation method in a period of 1 hour.
	// Note: This field may return null, indicating that no valid value was found.
	P3600 *string `json:"P3600,omitnil,omitempty" name:"P3600"`

	// Data aggregation method in a period of 1 day.
	// Note: This field may return null, indicating that no valid value was found.
	P86400 *string `json:"P86400,omitnil,omitempty" name:"P86400"`
}

type DescribePolicyConditionListEventMetric struct {
	// Event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Event name.
	EventShowName *string `json:"EventShowName,omitnil,omitempty" name:"EventShowName"`

	// Whether to recover.
	NeedRecovered *bool `json:"NeedRecovered,omitnil,omitempty" name:"NeedRecovered"`

	// Event type, which is a reserved field. Currently, it is fixed to 2.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribePolicyConditionListMetric struct {
	// Metric configuration.
	// Note: This field may return null, indicating that no valid value was found.
	ConfigManual *DescribePolicyConditionListConfigManual `json:"ConfigManual,omitnil,omitempty" name:"ConfigManual"`

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// Metric name.
	MetricShowName *string `json:"MetricShowName,omitnil,omitempty" name:"MetricShowName"`

	// Metric unit.
	MetricUnit *string `json:"MetricUnit,omitnil,omitempty" name:"MetricUnit"`
}

// Predefined struct for user
type DescribePolicyConditionListRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`
}

type DescribePolicyConditionListRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`
}

func (r *DescribePolicyConditionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConditionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyConditionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyConditionListResponseParams struct {
	// List of alarm policy conditions.
	Conditions []*DescribePolicyConditionListCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyConditionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyConditionListResponseParams `json:"Response"`
}

func (r *DescribePolicyConditionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConditionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListResponseDeprecatingInfo struct {
	// Whether to hide
	// Note: This field may return null, indicating that no valid values can be obtained.
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`

	// Names of new views
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewViewNames []*string `json:"NewViewNames,omitnil,omitempty" name:"NewViewNames"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type DescribePolicyGroupInfoCallback struct {
	// URL of the user callback API.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// Status of the user callback API. The value 0 indicates that the API is not verified. The value 1 indicates that the API is verified. The value 2 indicates that a URL exists but the API fails to be verified.
	ValidFlag *int64 `json:"ValidFlag,omitnil,omitempty" name:"ValidFlag"`

	// Verification code of the user callback API.
	VerifyCode *string `json:"VerifyCode,omitnil,omitempty" name:"VerifyCode"`
}

type DescribePolicyGroupInfoCondition struct {
	// Metric name.
	MetricShowName *string `json:"MetricShowName,omitnil,omitempty" name:"MetricShowName"`

	// Data aggregation period in seconds.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// Threshold rule ID.
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Metric unit.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. If the value is less than 0, no alarm will be triggered. If the value is 0, an alarm will be triggered only once. If the value is greater than 0, an alarm will be triggered at the interval of `triggerTime`.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to. The value 7 indicates day-on-day increase. The value 8 indicates day-on-day decrease. The value 9 indicates week-on-week increase. The value 10 indicates week-on-week decrease. The value 11 indicates periodical increase. The value 12 indicates periodical decrease.
	CalcType *int64 `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// Threshold.
	CalcValue *string `json:"CalcValue,omitnil,omitempty" name:"CalcValue"`

	// Duration at which an alarm will be triggered in seconds.
	ContinueTime *int64 `json:"ContinueTime,omitnil,omitempty" name:"ContinueTime"`

	// Alarm metric name.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`
}

type DescribePolicyGroupInfoConditionTpl struct {
	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Policy type.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Policy group remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Update time.
	// Note: This field may return null, indicating that no valid value was found.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time.
	// Note: This field may return null, indicating that no valid value was found.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Whether the 'AND' rule is used.
	// Note: This field may return null, indicating that no valid value was found.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupInfoEventCondition struct {
	// Event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Event alarm rule ID.
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Event name.
	EventShowName *string `json:"EventShowName,omitnil,omitempty" name:"EventShowName"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`
}

type DescribePolicyGroupInfoReceiverInfo struct {
	// List of alarm recipient group IDs.
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitnil,omitempty" name:"ReceiverGroupList"`

	// List of alarm recipient IDs.
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitnil,omitempty" name:"ReceiverUserList"`

	// Start time of the alarm period. Value range: [0,86400). Convert the Unix timestamp to Beijing time and then remove the date. For example, 7200 indicates “10:0:0”.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the alarm period. The meaning is the same as that of StartTime.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Recipient type. Valid values: group and user.
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Alarm notification method. Valid values: "SMS", "SITE", "EMAIL", "CALL", and "WECHAT".
	NotifyWay []*string `json:"NotifyWay,omitnil,omitempty" name:"NotifyWay"`

	// Uid of the alarm call recipient.
	// Note: This field may return null, indicating that no valid value was found.
	UidList []*int64 `json:"UidList,omitnil,omitempty" name:"UidList"`

	// Number of alarm call rounds.
	RoundNumber *int64 `json:"RoundNumber,omitnil,omitempty" name:"RoundNumber"`

	// Intervals of alarm call rounds in seconds.
	RoundInterval *int64 `json:"RoundInterval,omitnil,omitempty" name:"RoundInterval"`

	// Alarm call intervals for individuals in seconds.
	PersonInterval *int64 `json:"PersonInterval,omitnil,omitempty" name:"PersonInterval"`

	// Whether to send an alarm call delivery notice. The value 0 indicates that no notice needs to be sent. The value 1 indicates that a notice needs to be sent.
	NeedSendNotice *int64 `json:"NeedSendNotice,omitnil,omitempty" name:"NeedSendNotice"`

	// Alarm call notification time. Valid values: OCCUR (indicating that a notice is sent when the alarm is triggered) and RECOVER (indicating that a notice is sent when the alarm is recovered).
	SendFor []*string `json:"SendFor,omitnil,omitempty" name:"SendFor"`

	// Notification method when an alarm is recovered. Valid value: SMS.
	RecoverNotify []*string `json:"RecoverNotify,omitnil,omitempty" name:"RecoverNotify"`

	// Alarm language.
	// Note: This field may return null, indicating that no valid value was found.
	ReceiveLanguage *string `json:"ReceiveLanguage,omitnil,omitempty" name:"ReceiveLanguage"`
}

// Predefined struct for user
type DescribePolicyGroupInfoRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribePolicyGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DescribePolicyGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyGroupInfoResponseParams struct {
	// Policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// ID of the project to which the policy group belongs.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether it is the default policy. The value 0 indicates that it is not the default policy. The value 1 indicates that it is the default policy.
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Policy type.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Policy description
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Policy type name.
	ShowName *string `json:"ShowName,omitnil,omitempty" name:"ShowName"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Last edited time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Regions supported by this policy.
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// List of policy type dimensions.
	DimensionGroup []*string `json:"DimensionGroup,omitnil,omitempty" name:"DimensionGroup"`

	// Threshold rule list.
	// Note: This field may return null, indicating that no valid value was found.
	ConditionsConfig []*DescribePolicyGroupInfoCondition `json:"ConditionsConfig,omitnil,omitempty" name:"ConditionsConfig"`

	// Product event rule list.
	// Note: This field may return null, indicating that no valid value was found.
	EventConfig []*DescribePolicyGroupInfoEventCondition `json:"EventConfig,omitnil,omitempty" name:"EventConfig"`

	// Recipient list.
	// Note: This field may return null, indicating that no valid value was found.
	ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitnil,omitempty" name:"ReceiverInfos"`

	// User callback information.
	// Note: This field may return null, indicating that no valid value was found.
	Callback *DescribePolicyGroupInfoCallback `json:"Callback,omitnil,omitempty" name:"Callback"`

	// Template-based policy group.
	// Note: This field may return null, indicating that no valid value was found.
	ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitnil,omitempty" name:"ConditionsTemp"`

	// Whether the policy can be configured as the default policy.
	CanSetDefault *bool `json:"CanSetDefault,omitnil,omitempty" name:"CanSetDefault"`

	// Whether the 'AND' rule is used.
	// Note: This field may return null, indicating that no valid value was found.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyGroupInfoResponseParams `json:"Response"`
}

func (r *DescribePolicyGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListGroup struct {
	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Whether it is enabled.
	IsOpen *bool `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`

	// Policy view name.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Last modified time.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Number of instances that are bound to the policy group.
	UseSum *int64 `json:"UseSum,omitnil,omitempty" name:"UseSum"`

	// Number of unshielded instances that are bound to the policy group.
	NoShieldedSum *int64 `json:"NoShieldedSum,omitnil,omitempty" name:"NoShieldedSum"`

	// Whether it is the default policy. The value 0 indicates that it is not the default policy. The value 1 indicates that it is the default policy.
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Whether the policy can be configured as the default policy.
	CanSetDefault *bool `json:"CanSetDefault,omitnil,omitempty" name:"CanSetDefault"`

	// Parent policy group ID.
	ParentGroupId *int64 `json:"ParentGroupId,omitnil,omitempty" name:"ParentGroupId"`

	// Remarks of the policy group.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// ID of the project to which the policy group belongs.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Threshold rule list.
	// Note: This field may return null, indicating that no valid value was found.
	Conditions []*DescribePolicyGroupInfoCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Product event rule list.
	// Note: This field may return null, indicating that no valid value was found.
	EventConditions []*DescribePolicyGroupInfoEventCondition `json:"EventConditions,omitnil,omitempty" name:"EventConditions"`

	// Recipient list.
	// Note: This field may return null, indicating that no valid value was found.
	ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitnil,omitempty" name:"ReceiverInfos"`

	// Template-based policy group.
	// Note: This field may return null, indicating that no valid value was found.
	ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitnil,omitempty" name:"ConditionsTemp"`

	// Instance group that is bound to the policy group.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroup *DescribePolicyGroupListGroupInstanceGroup `json:"InstanceGroup,omitnil,omitempty" name:"InstanceGroup"`

	// The 'AND' or 'OR' rule. The value 0 indicates the 'OR' rule (indicating that an alarm will be triggered if any rule meets the threshold condition). The value 1 indicates the 'AND' rule (indicating that an alarm will be triggered when all rules meet the threshold conditions).
	// Note: This field may return null, indicating that no valid value was found.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupListGroupInstanceGroup struct {
	// Instance group name ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Policy type view name.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Uin that was last edited.
	LastEditUin *string `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Instance group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Number of instances.
	InstanceSum *int64 `json:"InstanceSum,omitnil,omitempty" name:"InstanceSum"`

	// Update time.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`
}

// Predefined struct for user
type DescribePolicyGroupListRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Search by policy name.
	Like *string `json:"Like,omitnil,omitempty" name:"Like"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Sort by update time. Valid values: asc and desc.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitnil,omitempty" name:"UpdateTimeOrder"`

	// Project ID list.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// List of alarm policy types.
	ViewNames []*string `json:"ViewNames,omitnil,omitempty" name:"ViewNames"`

	// Whether to filter policy groups without recipients. The value 1 indicates that policy groups without recipients will be filtered. The value 0 indicates that policy groups without recipients will not be filtered.
	FilterUnuseReceiver *int64 `json:"FilterUnuseReceiver,omitnil,omitempty" name:"FilterUnuseReceiver"`

	// Filter by recipient group.
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Filter by recipient.
	ReceiverUserList []*string `json:"ReceiverUserList,omitnil,omitempty" name:"ReceiverUserList"`

	// Dimension set field (json string), for example, [[{"name":"unInstanceId","value":"ins-6e4b2aaa"}]].
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Template-based policy group IDs, which are separated by commas.
	ConditionTempGroupId *string `json:"ConditionTempGroupId,omitnil,omitempty" name:"ConditionTempGroupId"`

	// Filter by recipient or recipient group. The value 'user' indicates by recipient. The value 'group' indicates by recipient group.
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Filter conditions. Whether the alarm policy has been enabled or disabled
	IsOpen *bool `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`
}

type DescribePolicyGroupListRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Search by policy name.
	Like *string `json:"Like,omitnil,omitempty" name:"Like"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Sort by update time. Valid values: asc and desc.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitnil,omitempty" name:"UpdateTimeOrder"`

	// Project ID list.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// List of alarm policy types.
	ViewNames []*string `json:"ViewNames,omitnil,omitempty" name:"ViewNames"`

	// Whether to filter policy groups without recipients. The value 1 indicates that policy groups without recipients will be filtered. The value 0 indicates that policy groups without recipients will not be filtered.
	FilterUnuseReceiver *int64 `json:"FilterUnuseReceiver,omitnil,omitempty" name:"FilterUnuseReceiver"`

	// Filter by recipient group.
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Filter by recipient.
	ReceiverUserList []*string `json:"ReceiverUserList,omitnil,omitempty" name:"ReceiverUserList"`

	// Dimension set field (json string), for example, [[{"name":"unInstanceId","value":"ins-6e4b2aaa"}]].
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Template-based policy group IDs, which are separated by commas.
	ConditionTempGroupId *string `json:"ConditionTempGroupId,omitnil,omitempty" name:"ConditionTempGroupId"`

	// Filter by recipient or recipient group. The value 'user' indicates by recipient. The value 'group' indicates by recipient group.
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Filter conditions. Whether the alarm policy has been enabled or disabled
	IsOpen *bool `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`
}

func (r *DescribePolicyGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Like")
	delete(f, "InstanceGroupId")
	delete(f, "UpdateTimeOrder")
	delete(f, "ProjectIds")
	delete(f, "ViewNames")
	delete(f, "FilterUnuseReceiver")
	delete(f, "Receivers")
	delete(f, "ReceiverUserList")
	delete(f, "Dimensions")
	delete(f, "ConditionTempGroupId")
	delete(f, "ReceiverType")
	delete(f, "IsOpen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyGroupListResponseParams struct {
	// Policy group list.
	// Note: This field may return null, indicating that no valid value was found.
	GroupList []*DescribePolicyGroupListGroup `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// Total number of policy groups.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Warning *string `json:"Warning,omitnil,omitempty" name:"Warning"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyGroupListResponseParams `json:"Response"`
}

func (r *DescribePolicyGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListDimensions struct {
	// Dimension name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Dimension value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DescribeProductEventListEvents struct {
	// Event ID.
	// Note: This field may return null, indicating that no valid value was found.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Event name in Chinese.
	// Note: This field may return null, indicating that no valid value was found.
	EventCName *string `json:"EventCName,omitnil,omitempty" name:"EventCName"`

	// Event name in English.
	// Note: This field may return null, indicating that no valid value was found.
	EventEName *string `json:"EventEName,omitnil,omitempty" name:"EventEName"`

	// Event name abbreviation.
	// Note: This field may return null, indicating that no valid value was found.
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Product name in Chinese.
	// Note: This field may return null, indicating that no valid value was found.
	ProductCName *string `json:"ProductCName,omitnil,omitempty" name:"ProductCName"`

	// Product name in English.
	// Note: This field may return null, indicating that no valid value was found.
	ProductEName *string `json:"ProductEName,omitnil,omitempty" name:"ProductEName"`

	// Product name abbreviation.
	// Note: This field may return null, indicating that no valid value was found.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Project ID.
	// Note: This field may return null, indicating that no valid value was found.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Region.
	// Note: This field may return null, indicating that no valid value was found.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Status.
	// Note: This field may return null, indicating that no valid value was found.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to support alarms.
	// Note: This field may return null, indicating that no valid value was found.
	SupportAlarm *int64 `json:"SupportAlarm,omitnil,omitempty" name:"SupportAlarm"`

	// Event type.
	// Note: This field may return null, indicating that no valid value was found.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Start time.
	// Note: This field may return null, indicating that no valid value was found.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid value was found.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Instance object information.
	// Note: This field may return null, indicating that no valid value was found.
	Dimensions []*DescribeProductEventListEventsDimensions `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Additional information of the instance object.
	// Note: This field may return null, indicating that no valid value was found.
	AdditionMsg []*DescribeProductEventListEventsDimensions `json:"AdditionMsg,omitnil,omitempty" name:"AdditionMsg"`

	// Whether to configure alarms.
	// Note: This field may return null, indicating that no valid value was found.
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitnil,omitempty" name:"IsAlarmConfig"`

	// Policy information.
	// Note: This field may return null, indicating that no valid value was found.
	GroupInfo []*DescribeProductEventListEventsGroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

	// Display name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`
}

type DescribeProductEventListEventsDimensions struct {
	// Dimension name in English.
	// Note: This field may return null, indicating that no valid value was found.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Dimension name in Chinese.
	// Note: This field may return null, indicating that no valid value was found.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Dimension value.
	// Note: This field may return null, indicating that no valid value was found.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DescribeProductEventListEventsGroupInfo struct {
	// Policy ID.
	// Note: This field may return null, indicating that no valid value was found.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Policy name.
	// Note: This field may return null, indicating that no valid value was found.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type DescribeProductEventListOverView struct {
	// Number of events whose statuses have changed.
	// Note: This field may return null, indicating that no valid value was found.
	StatusChangeAmount *int64 `json:"StatusChangeAmount,omitnil,omitempty" name:"StatusChangeAmount"`

	// Number of events whose alarm statuses are not configured.
	// Note: This field may return null, indicating that no valid value was found.
	UnConfigAlarmAmount *int64 `json:"UnConfigAlarmAmount,omitnil,omitempty" name:"UnConfigAlarmAmount"`

	// Number of events with exceptions.
	// Note: This field may return null, indicating that no valid value was found.
	UnNormalEventAmount *int64 `json:"UnNormalEventAmount,omitnil,omitempty" name:"UnNormalEventAmount"`

	// Number of events that have not been recovered.
	// Note: This field may return null, indicating that no valid value was found.
	UnRecoverAmount *int64 `json:"UnRecoverAmount,omitnil,omitempty" name:"UnRecoverAmount"`
}

// Predefined struct for user
type DescribeProductEventListRequestParams struct {
	// API component name. It is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Filter by product type. For example, "cvm" indicates Cloud Virtual Machine.
	ProductName []*string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Filter by event name. For example, "guest_reboot" indicates instance restart.
	EventName []*string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Affected object, such as "ins-19708ino".
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by dimension, such as by public IP: 10.0.0.1.
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Region filter parameter for service events.
	RegionList []*string `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Filter by event type. Valid values: ["status_change","abnormal"], which indicate events whose statuses have changed and events with exceptions respectively.
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter by event status. Valid values: ["recover","alarm","-"], which indicate that an event has been recovered, has not been recovered, and has no status respectively.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Filter by project ID.
	Project []*string `json:"Project,omitnil,omitempty" name:"Project"`

	// Filter by alarm status configuration. The value 1 indicates that the alarm status has been configured. The value 0 indicates that the alarm status has not been configured.
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitnil,omitempty" name:"IsAlarmConfig"`

	// Sorting by update time. The value ASC indicates the ascending order. The value DESC indicates the descending order. The default value is DESC.
	TimeOrder *string `json:"TimeOrder,omitnil,omitempty" name:"TimeOrder"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page offset. The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of parameters that can be returned on each page. The default value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest
	
	// API component name. It is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Filter by product type. For example, "cvm" indicates Cloud Virtual Machine.
	ProductName []*string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Filter by event name. For example, "guest_reboot" indicates instance restart.
	EventName []*string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Affected object, such as "ins-19708ino".
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by dimension, such as by public IP: 10.0.0.1.
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Region filter parameter for service events.
	RegionList []*string `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Filter by event type. Valid values: ["status_change","abnormal"], which indicate events whose statuses have changed and events with exceptions respectively.
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter by event status. Valid values: ["recover","alarm","-"], which indicate that an event has been recovered, has not been recovered, and has no status respectively.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Filter by project ID.
	Project []*string `json:"Project,omitnil,omitempty" name:"Project"`

	// Filter by alarm status configuration. The value 1 indicates that the alarm status has been configured. The value 0 indicates that the alarm status has not been configured.
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitnil,omitempty" name:"IsAlarmConfig"`

	// Sorting by update time. The value ASC indicates the ascending order. The value DESC indicates the descending order. The default value is DESC.
	TimeOrder *string `json:"TimeOrder,omitnil,omitempty" name:"TimeOrder"`

	// Start time, which is the timestamp one day prior by default.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page offset. The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of parameters that can be returned on each page. The default value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeProductEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "ProductName")
	delete(f, "EventName")
	delete(f, "InstanceId")
	delete(f, "Dimensions")
	delete(f, "RegionList")
	delete(f, "Type")
	delete(f, "Status")
	delete(f, "Project")
	delete(f, "IsAlarmConfig")
	delete(f, "TimeOrder")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductEventListResponseParams struct {
	// Event list
	// Note: This field may return null, indicating that no valid value was found.
	Events []*DescribeProductEventListEvents `json:"Events,omitnil,omitempty" name:"Events"`

	// Event statistics.
	OverView *DescribeProductEventListOverView `json:"OverView,omitnil,omitempty" name:"OverView"`

	// Total number of events.
	// Note: This field may return null, indicating that no valid value was found.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductEventListResponseParams `json:"Response"`
}

func (r *DescribeProductEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentInstancesRequestParams struct {
	// Cluster ID
	// It can be the ID of a TKE, EKS, or edge cluster.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribePrometheusAgentInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	// It can be the ID of a TKE, EKS, or edge cluster.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribePrometheusAgentInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAgentInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentInstancesResponseParams struct {
	// List of instances associated with the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Instances []*string `json:"Instances,omitnil,omitempty" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAgentInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAgentInstancesResponseParams `json:"Response"`
}

func (r *DescribePrometheusAgentInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of agent IDs
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusAgentsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of agent IDs
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "AgentIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentsResponseParams struct {
	// List of agents
	// Note: This field may return null, indicating that no valid values can be obtained.
	AgentSet []*PrometheusAgent `json:"AgentSet,omitnil,omitempty" name:"AgentSet"`

	// Total number of agents
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAgentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAgentsResponseParams `json:"Response"`
}

func (r *DescribePrometheusAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertGroupsRequestParams struct {
	// prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned results. defaults to 20. maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Alert group ID, such as alert-xxxx.
	// List the alert group with the given ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Specifies the alert group name.
	// List alert groups which name contains the given string.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type DescribePrometheusAlertGroupsRequest struct {
	*tchttp.BaseRequest
	
	// prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned results. defaults to 20. maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Alert group ID, such as alert-xxxx.
	// List the alert group with the given ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Specifies the alert group name.
	// List alert groups which name contains the given string.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

func (r *DescribePrometheusAlertGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAlertGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertGroupsResponseParams struct {
	// Alert group information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlertGroupSet []*PrometheusAlertGroupSet `json:"AlertGroupSet,omitnil,omitempty" name:"AlertGroupSet"`

	// Total count of Alert groups.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAlertGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAlertGroupsResponseParams `json:"Response"`
}

func (r *DescribePrometheusAlertGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter
	// Valid values: `ID`, `Name`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter
	// Valid values: `ID`, `Name`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertPolicyResponseParams struct {
	// Alert details
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlertRules []*PrometheusAlertPolicyItem `json:"AlertRules,omitnil,omitempty" name:"AlertRules"`

	// Total number
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *DescribePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusClusterAgentsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusClusterAgentsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusClusterAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusClusterAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusClusterAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusClusterAgentsResponseParams struct {
	// Information of the associated cluster
	Agents []*PrometheusAgentOverview `json:"Agents,omitnil,omitempty" name:"Agents"`

	// The total number of the associated clusters
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Whether the TMP instance is associated with the cluster for the first time. If so, you need to configure recording rules for it. This also applies if it has no default recording rule.
	IsFirstBind *bool `json:"IsFirstBind,omitnil,omitempty" name:"IsFirstBind"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusClusterAgentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusClusterAgentsResponseParams `json:"Response"`
}

func (r *DescribePrometheusClusterAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusClusterAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusConfigResponseParams struct {
	// Global configuration
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// ServiceMonitor configuration
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// PodMonitor configuration
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// Raw jobs
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// Probes
	Probes []*PrometheusConfigItem `json:"Probes,omitnil,omitempty" name:"Probes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusConfigResponseParams `json:"Response"`
}

func (r *DescribePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalConfigRequestParams struct {
	// Instance-level scrape configuration
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to disable statistics
	DisableStatistics *bool `json:"DisableStatistics,omitnil,omitempty" name:"DisableStatistics"`
}

type DescribePrometheusGlobalConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance-level scrape configuration
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to disable statistics
	DisableStatistics *bool `json:"DisableStatistics,omitnil,omitempty" name:"DisableStatistics"`
}

func (r *DescribePrometheusGlobalConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DisableStatistics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusGlobalConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalConfigResponseParams struct {
	// Configuration content
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// List of service monitors and the corresponding targets information
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// List of pod monitors and the corresponding targets information
	// Note: This field may return null, indicating that no valid values can be obtained.
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// List of raw jobs and the corresponding targets information
	// Note: This field may return null, indicating that no valid values can be obtained.
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// List of probes and the corresponding targets information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Probes []*PrometheusConfigItem `json:"Probes,omitnil,omitempty" name:"Probes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusGlobalConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusGlobalConfigResponseParams `json:"Response"`
}

func (r *DescribePrometheusGlobalConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalNotificationRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribePrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalNotificationResponseParams struct {
	// Global alert notification channel
	// Note: This field may return null, indicating that no valid values can be obtained.
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *DescribePrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceDetailRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribePrometheusInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceDetailResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance status. Valid values:
	// 
	// `1`: Creating
	// `2`: Running
	// `3`: Abnormal
	// `4`: Rebooting
	// `5`: Terminating
	// `6`: Service suspended
	// `8`: Suspending service for overdue payment
	// `9`: Service suspended for overdue payment
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Billing status
	// 
	// `1`: Normal
	// `2`: Expired
	// `3`: Terminated
	// `4`: Assigning
	// `5`: Failed to assign
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeStatus *int64 `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// Whether Grafana is enabled
	// `0`: Disabled
	// `1`: Enabled
	EnableGrafana *int64 `json:"EnableGrafana,omitnil,omitempty" name:"EnableGrafana"`

	// Grafana dashboard URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrafanaURL *string `json:"GrafanaURL,omitnil,omitempty" name:"GrafanaURL"`

	// Instance billing mode. Valid values:
	// 
	// `2`: Monthly subscription
	// `3`: Pay-as-you-go
	InstanceChargeType *int64 `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Specification name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Storage period
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`

	// Expiration time of the purchased instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Auto-renewal flag
	// 
	// `0`: Auto-renewal not enabled
	// `1`: Auto-renewal enabled
	// `2`: Auto-renewal prohibited
	// `-1`: Invalid
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceInitStatusRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribePrometheusInstanceInitStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusInstanceInitStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceInitStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceInitStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceInitStatusResponseParams struct {
	// Instance initialization status. Valid values:
	// `uninitialized` 
	// `initializing`
	// `running`
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Initialize task steps
	// Note: This field may return null, indicating that no valid values can be obtained.
	Steps []*TaskStepInfo `json:"Steps,omitnil,omitempty" name:"Steps"`

	// EKS cluster ID of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	EksClusterId *string `json:"EksClusterId,omitnil,omitempty" name:"EksClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceInitStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceInitStatusResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceInitStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceInitStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceUsageRequestParams struct {
	// Query by one or more instance IDs. Instance ID is in the format of `prom-xxxxxxxx`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Start time
	StartCalcDate *string `json:"StartCalcDate,omitnil,omitempty" name:"StartCalcDate"`

	// End time
	EndCalcDate *string `json:"EndCalcDate,omitnil,omitempty" name:"EndCalcDate"`
}

type DescribePrometheusInstanceUsageRequest struct {
	*tchttp.BaseRequest
	
	// Query by one or more instance IDs. Instance ID is in the format of `prom-xxxxxxxx`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Start time
	StartCalcDate *string `json:"StartCalcDate,omitnil,omitempty" name:"StartCalcDate"`

	// End time
	EndCalcDate *string `json:"EndCalcDate,omitnil,omitempty" name:"EndCalcDate"`
}

func (r *DescribePrometheusInstanceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "StartCalcDate")
	delete(f, "EndCalcDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceUsageResponseParams struct {
	// Usage list
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageSet []*PrometheusInstanceTenantUsage `json:"UsageSet,omitnil,omitempty" name:"UsageSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceUsageResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesOverviewRequestParams struct {
	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Instance filter. Valid values:
	// `ID`: Filter by instance ID 
	// `Name`: Filter by instance name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusInstancesOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Instance filter. Valid values:
	// `ID`: Filter by instance ID 
	// `Name`: Filter by instance name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusInstancesOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstancesOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesOverviewResponseParams struct {
	// List of instances
	Instances []*PrometheusInstancesOverview `json:"Instances,omitnil,omitempty" name:"Instances"`

	// Total number of instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstancesOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstancesOverviewResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstancesOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesRequestParams struct {
	// Queries by instance ID or IDs. Instance ID is in the format of `prom-xxxxxxxx`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Filter by instance status
	// <ul>
	// <li>1: Creating</li>
	// <li>2: Running</li>
	// <li>3: Abnormal</li>
	// <li>4: Rebooting</li>
	// <li>5: Terminating</li>
	// <li>6: Service suspended</li>
	// <li>8: Suspending service for overdue payment</li>
	// <li>9: Service suspended for overdue payment</li>
	// </ul>
	InstanceStatus []*int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Filter by instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Filter by AZ in the format of `ap-guangzhou-1`
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Filter by tag key-value pair. The `tag-key` should be replaced with a specific tag key.
	TagFilters []*PrometheusTag `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// Filter by instance IPv4 address
	IPv4Address []*string `json:"IPv4Address,omitnil,omitempty" name:"IPv4Address"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter by billing mode
	// <li>2: Monthly subscription</li>
	// <li>3: Pay-as-you-go</li>
	InstanceChargeType *int64 `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`
}

type DescribePrometheusInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Queries by instance ID or IDs. Instance ID is in the format of `prom-xxxxxxxx`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Filter by instance status
	// <ul>
	// <li>1: Creating</li>
	// <li>2: Running</li>
	// <li>3: Abnormal</li>
	// <li>4: Rebooting</li>
	// <li>5: Terminating</li>
	// <li>6: Service suspended</li>
	// <li>8: Suspending service for overdue payment</li>
	// <li>9: Service suspended for overdue payment</li>
	// </ul>
	InstanceStatus []*int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Filter by instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Filter by AZ in the format of `ap-guangzhou-1`
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Filter by tag key-value pair. The `tag-key` should be replaced with a specific tag key.
	TagFilters []*PrometheusTag `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// Filter by instance IPv4 address
	IPv4Address []*string `json:"IPv4Address,omitnil,omitempty" name:"IPv4Address"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter by billing mode
	// <li>2: Monthly subscription</li>
	// <li>3: Pay-as-you-go</li>
	InstanceChargeType *int64 `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`
}

func (r *DescribePrometheusInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceStatus")
	delete(f, "InstanceName")
	delete(f, "Zones")
	delete(f, "TagFilters")
	delete(f, "IPv4Address")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesResponseParams struct {
	// List of instance details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceSet []*PrometheusInstancesItem `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstancesResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRuleYamlRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter. Valid values:
	// `Name`: Name
	// `Values`: List of target names
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter. Valid values:
	// `Name`: Name
	// `Values`: List of target names
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRuleYamlResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *DescribePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRulesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusRecordRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusRecordRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusRecordRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRulesResponseParams struct {
	// Recording rule
	Records []*PrometheusRecordRuleYamlItem `json:"Records,omitnil,omitempty" name:"Records"`

	// Total number
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusRecordRulesResponseParams `json:"Response"`
}

func (r *DescribePrometheusRecordRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusScrapeJobsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of task IDs
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusScrapeJobsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of task IDs
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusScrapeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusScrapeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "Name")
	delete(f, "JobIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusScrapeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusScrapeJobsResponseParams struct {
	// List of tasks
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScrapeJobSet []*PrometheusScrapeJob `json:"ScrapeJobSet,omitnil,omitempty" name:"ScrapeJobSet"`

	// Total number of tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusScrapeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusScrapeJobsResponseParams `json:"Response"`
}

func (r *DescribePrometheusScrapeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusScrapeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTargetsTMPRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Filters.
	// You can filter by `RawJob`, `Job`, `ServiceMonitor`, `PodMonitor`, or `Health`.
	// `Health` contains three values: `up`, `down`, `unknown`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusTargetsTMPRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Filters.
	// You can filter by `RawJob`, `Job`, `ServiceMonitor`, `PodMonitor`, or `Health`.
	// `Health` contains three values: `up`, `down`, `unknown`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusTargetsTMPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTargetsTMPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTargetsTMPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTargetsTMPResponseParams struct {
	// Targets information of all jobs
	Jobs []*PrometheusJobTargets `json:"Jobs,omitnil,omitempty" name:"Jobs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTargetsTMPResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTargetsTMPResponseParams `json:"Response"`
}

func (r *DescribePrometheusTargetsTMPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTargetsTMPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempRequestParams struct {
	// Fuzzy filter. Valid values:
	// `Level`: Filter by template level
	// `Name`: Filter by name
	// `Describe`: Filter by description
	// `ID`: Filter by templateId
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// Fuzzy filter. Valid values:
	// `Level`: Filter by template level
	// `Name`: Filter by name
	// `Describe`: Filter by description
	// `ID`: Filter by templateId
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempResponseParams struct {
	// List of templates
	Templates []*PrometheusTemp `json:"Templates,omitnil,omitempty" name:"Templates"`

	// Total number
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTempResponseParams `json:"Response"`
}

func (r *DescribePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempSyncRequestParams struct {
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribePrometheusTempSyncRequest struct {
	*tchttp.BaseRequest
	
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribePrometheusTempSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTempSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempSyncResponseParams struct {
	// Details of the sync target
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTempSyncResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTempSyncResponseParams `json:"Response"`
}

func (r *DescribePrometheusTempSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusZonesRequestParams struct {
	// Region ID. You only need to specify the value of either `RegionId` or `RegionName`.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name. You only need to specify the value of either `RegionId` or `RegionName`.
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type DescribePrometheusZonesRequest struct {
	*tchttp.BaseRequest
	
	// Region ID. You only need to specify the value of either `RegionId` or `RegionName`.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name. You only need to specify the value of either `RegionId` or `RegionName`.
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

func (r *DescribePrometheusZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionId")
	delete(f, "RegionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusZonesResponseParams struct {
	// Region list
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneSet []*PrometheusZoneItem `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusZonesResponseParams `json:"Response"`
}

func (r *DescribePrometheusZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingRulesRequestParams struct {
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeRecordingRulesRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeRecordingRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleId")
	delete(f, "RuleState")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordingRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingRulesResponseParams struct {
	// Number of rule groups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Rule group details
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecordingRuleSet []*RecordingRuleSet `json:"RecordingRuleSet,omitnil,omitempty" name:"RecordingRuleSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordingRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordingRulesResponseParams `json:"Response"`
}

func (r *DescribeRecordingRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSOAccountRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by account ID such as “10000”
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by account ID such as “10000”
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSOAccountResponseParams struct {
	// List of authorized accounts
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountSet []*GrafanaAccountInfo `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSSOAccountResponseParams `json:"Response"`
}

func (r *DescribeSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceDiscoveryRequestParams struct {
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <li>TKE: ID of the integrated TKE cluster</li>
	KubeClusterId *string `json:"KubeClusterId,omitnil,omitempty" name:"KubeClusterId"`

	// Kubernetes cluster type:
	// <li> 1 = TKE </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`
}

type DescribeServiceDiscoveryRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <li>TKE: ID of the integrated TKE cluster</li>
	KubeClusterId *string `json:"KubeClusterId,omitnil,omitempty" name:"KubeClusterId"`

	// Kubernetes cluster type:
	// <li> 1 = TKE </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`
}

func (r *DescribeServiceDiscoveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceDiscoveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeClusterId")
	delete(f, "KubeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceDiscoveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceDiscoveryResponseParams struct {
	// List of returned scrape configurations
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceDiscoverySet []*ServiceDiscoveryItem `json:"ServiceDiscoverySet,omitnil,omitempty" name:"ServiceDiscoverySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceDiscoveryResponseParams `json:"Response"`
}

func (r *DescribeServiceDiscoveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceDiscoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStatisticDataRequestParams struct {
	// Module, whose value is fixed at `monitor`
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Namespace. Valid values: `QCE`, `TKE2`.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric name list
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// Dimension condition. The `=` and `in` operators are supported
	Conditions []*MidQueryCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Statistical period in seconds. Default value: 300. Optional values: 60, 300, 3,600, and 86,400.
	// Due to the storage period limit, the statistical period is subject to the time range of statistics:
	// 60s: The time range is less than 12 hours, and the timespan between `StartTime` and the current time cannot exceed 15 days.
	// 300s: The time range is less than three days, and the timespan between `StartTime` and the current time cannot exceed 31 days.
	// 3,600s: The time range is less than 30 days, and the timespan between `StartTime` and the current time cannot exceed 93 days.
	// 86,400s: The time range is less than 186 days, and the timespan between `StartTime` and the current time cannot exceed 186 days.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time, which is the current time by default, such as 2020-12-08T19:51:23+08:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current time by default, such as 2020-12-08T19:51:23+08:00
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// `groupBy` by the specified dimension
	GroupBys []*string `json:"GroupBys,omitnil,omitempty" name:"GroupBys"`
}

type DescribeStatisticDataRequest struct {
	*tchttp.BaseRequest
	
	// Module, whose value is fixed at `monitor`
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Namespace. Valid values: `QCE`, `TKE2`.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric name list
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// Dimension condition. The `=` and `in` operators are supported
	Conditions []*MidQueryCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Statistical period in seconds. Default value: 300. Optional values: 60, 300, 3,600, and 86,400.
	// Due to the storage period limit, the statistical period is subject to the time range of statistics:
	// 60s: The time range is less than 12 hours, and the timespan between `StartTime` and the current time cannot exceed 15 days.
	// 300s: The time range is less than three days, and the timespan between `StartTime` and the current time cannot exceed 31 days.
	// 3,600s: The time range is less than 30 days, and the timespan between `StartTime` and the current time cannot exceed 93 days.
	// 86,400s: The time range is less than 186 days, and the timespan between `StartTime` and the current time cannot exceed 186 days.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time, which is the current time by default, such as 2020-12-08T19:51:23+08:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current time by default, such as 2020-12-08T19:51:23+08:00
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// `groupBy` by the specified dimension
	GroupBys []*string `json:"GroupBys,omitnil,omitempty" name:"GroupBys"`
}

func (r *DescribeStatisticDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Namespace")
	delete(f, "MetricNames")
	delete(f, "Conditions")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "GroupBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStatisticDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStatisticDataResponseParams struct {
	// Statistical period
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Monitoring data
	Data []*MetricData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStatisticDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStatisticDataResponseParams `json:"Response"`
}

func (r *DescribeStatisticDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrometheusInstanceRequestParams struct {
	// Instance ID. The instance must be terminated first.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyPrometheusInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. The instance must be terminated first.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyPrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrometheusInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPrometheusInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrometheusInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyPrometheusInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPrometheusInstanceResponseParams `json:"Response"`
}

func (r *DestroyPrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrometheusInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {
	// Instance dimension name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Instance dimension value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DimensionNew struct {
	// Dimension key ID displayed on the backend
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Dimension key name displayed on the frontend
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether it is required
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// List of supported operators
	Operators []*Operator `json:"Operators,omitnil,omitempty" name:"Operators"`

	// Whether multiple items can be selected
	IsMultiple *bool `json:"IsMultiple,omitnil,omitempty" name:"IsMultiple"`

	// Whether it can be modified after creation
	IsMutable *bool `json:"IsMutable,omitnil,omitempty" name:"IsMutable"`

	// Whether it is displayed to users
	IsVisible *bool `json:"IsVisible,omitnil,omitempty" name:"IsVisible"`

	// Whether it can be used to filter policies
	CanFilterPolicy *bool `json:"CanFilterPolicy,omitnil,omitempty" name:"CanFilterPolicy"`

	// Whether it can be used to filter historical alarms
	CanFilterHistory *bool `json:"CanFilterHistory,omitnil,omitempty" name:"CanFilterHistory"`

	// Whether it can be used as an aggregate dimension
	CanGroupBy *bool `json:"CanGroupBy,omitnil,omitempty" name:"CanGroupBy"`

	// Whether it must be used as an aggregate dimension
	MustGroupBy *bool `json:"MustGroupBy,omitnil,omitempty" name:"MustGroupBy"`

	// The key to be replaced on the frontend
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShowValueReplace *string `json:"ShowValueReplace,omitnil,omitempty" name:"ShowValueReplace"`
}

type DimensionsDesc struct {
	// Array of dimension names
	Dimensions []*string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

// Predefined struct for user
type EnableGrafanaInternetRequestParams struct {
	// TCMG instance ID, such as “grafana-kleu3gt0”.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Whether to enable public network access (`true`: Yes; `false`: No)
	EnableInternet *bool `json:"EnableInternet,omitnil,omitempty" name:"EnableInternet"`
}

type EnableGrafanaInternetRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-kleu3gt0”.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Whether to enable public network access (`true`: Yes; `false`: No)
	EnableInternet *bool `json:"EnableInternet,omitnil,omitempty" name:"EnableInternet"`
}

func (r *EnableGrafanaInternetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaInternetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "EnableInternet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableGrafanaInternetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGrafanaInternetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableGrafanaInternetResponse struct {
	*tchttp.BaseResponse
	Response *EnableGrafanaInternetResponseParams `json:"Response"`
}

func (r *EnableGrafanaInternetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaInternetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGrafanaSSORequestParams struct {
	// Whether to enable SSO (`true`: Yes; `false`: No)
	EnableSSO *bool `json:"EnableSSO,omitnil,omitempty" name:"EnableSSO"`

	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type EnableGrafanaSSORequest struct {
	*tchttp.BaseRequest
	
	// Whether to enable SSO (`true`: Yes; `false`: No)
	EnableSSO *bool `json:"EnableSSO,omitnil,omitempty" name:"EnableSSO"`

	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *EnableGrafanaSSORequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaSSORequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnableSSO")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableGrafanaSSORequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGrafanaSSOResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableGrafanaSSOResponse struct {
	*tchttp.BaseResponse
	Response *EnableGrafanaSSOResponseParams `json:"Response"`
}

func (r *EnableGrafanaSSOResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaSSOResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSSOCamCheckRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to enable CAM authentication (`true`: Yes; `false`: No)
	EnableSSOCamCheck *bool `json:"EnableSSOCamCheck,omitnil,omitempty" name:"EnableSSOCamCheck"`
}

type EnableSSOCamCheckRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to enable CAM authentication (`true`: Yes; `false`: No)
	EnableSSOCamCheck *bool `json:"EnableSSOCamCheck,omitnil,omitempty" name:"EnableSSOCamCheck"`
}

func (r *EnableSSOCamCheckRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSSOCamCheckRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EnableSSOCamCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableSSOCamCheckRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSSOCamCheckResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableSSOCamCheckResponse struct {
	*tchttp.BaseResponse
	Response *EnableSSOCamCheckResponseParams `json:"Response"`
}

func (r *EnableSSOCamCheckResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSSOCamCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventCondition struct {
	// Alarm notification frequency.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// Predefined repeated notification policy. `0`: One-time alarm; `1`: exponential alarm; `2`: consecutive alarm
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AlarmNotifyType *string `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`

	// Event ID.
	EventID *string `json:"EventID,omitnil,omitempty" name:"EventID"`

	// Displayed event name.
	EventDisplayName *string `json:"EventDisplayName,omitnil,omitempty" name:"EventDisplayName"`

	// Rule ID.
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`
}

type Filter struct {
	// Filter method. Valid values: `=`, `!=`, `in`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter dimension name
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Filter value. For the `in` filter method, separate multiple values by comma.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Filter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value range
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetMonitorDataRequestParams struct {
	// Namespace, such as QCE/CVM. For more information on the namespaces of each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1)
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric name, such as `CPUUsage`. Only one monitoring metric can be pulled at a time. For more information on the metrics of each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). The corresponding metric name is `MetricName`.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// The dimension combination for instance objects, which is in the form of a set of key-value pairs. The dimension fields for instances of different Tencent Cloud services are completely different. For example, the field is [{"Name":"InstanceId","Value":"ins-j0hk02zo"}] for CVM instances, [{"Name":"instanceId","Value":"ckafka-l49k54dd"}] for CKafka instances, and [{"Name":"appid","Value":"1258344699"},{"Name":"bucket","Value":"rig-1258344699"}] for COS instances. For more information on the dimensions of various Tencent Cloud services, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). In each document, the dimension column displays a dimension combination’s key, which has a corresponding value. A single request can get the data of up to 10 instances.
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// Monitoring statistical period in seconds, such as 60. Default value: 300. The statistical period varies by metric. For more information on the statistical periods supported by each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). The values in the statistical period column are the supported statistical periods. A single request can get up to 1,440 data points.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time such as 2018-09-22T19:51:23+08:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current time by default, such as 2018-09-22T20:51:23+08:00. `EndTime` cannot be earlier than `StartTime`
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest
	
	// Namespace, such as QCE/CVM. For more information on the namespaces of each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1)
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric name, such as `CPUUsage`. Only one monitoring metric can be pulled at a time. For more information on the metrics of each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). The corresponding metric name is `MetricName`.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// The dimension combination for instance objects, which is in the form of a set of key-value pairs. The dimension fields for instances of different Tencent Cloud services are completely different. For example, the field is [{"Name":"InstanceId","Value":"ins-j0hk02zo"}] for CVM instances, [{"Name":"instanceId","Value":"ckafka-l49k54dd"}] for CKafka instances, and [{"Name":"appid","Value":"1258344699"},{"Name":"bucket","Value":"rig-1258344699"}] for COS instances. For more information on the dimensions of various Tencent Cloud services, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). In each document, the dimension column displays a dimension combination’s key, which has a corresponding value. A single request can get the data of up to 10 instances.
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// Monitoring statistical period in seconds, such as 60. Default value: 300. The statistical period varies by metric. For more information on the statistical periods supported by each Tencent Cloud service, please see [Tencent Cloud Service Metrics](https://intl.cloud.tencent.com/document/product/248/6140?from_cn_redirect=1). The values in the statistical period column are the supported statistical periods. A single request can get up to 1,440 data points.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time such as 2018-09-22T19:51:23+08:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, which is the current time by default, such as 2018-09-22T20:51:23+08:00. `EndTime` cannot be earlier than `StartTime`
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	delete(f, "MetricName")
	delete(f, "Instances")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMonitorDataResponseParams struct {
	// Statistical period
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Metric name
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Array of data points
	DataPoints []*DataPoint `json:"DataPoints,omitnil,omitempty" name:"DataPoints"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Returned message
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *GetMonitorDataResponseParams `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPrometheusAgentManagementCommandRequestParams struct {
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type GetPrometheusAgentManagementCommandRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *GetPrometheusAgentManagementCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPrometheusAgentManagementCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPrometheusAgentManagementCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPrometheusAgentManagementCommandResponseParams struct {
	// Agent management command
	Command *ManagementCommand `json:"Command,omitnil,omitempty" name:"Command"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPrometheusAgentManagementCommandResponse struct {
	*tchttp.BaseResponse
	Response *GetPrometheusAgentManagementCommandResponseParams `json:"Response"`
}

func (r *GetPrometheusAgentManagementCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPrometheusAgentManagementCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrafanaAccountInfo struct {
	// User account ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// User permission
	Role []*GrafanaAccountRole `json:"Role,omitnil,omitempty" name:"Role"`

	// Remarks
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// Creation time
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User’s root account UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type GrafanaAccountRole struct {
	// Organization
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// Permission
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type GrafanaChannel struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Array of alert channel template IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Update time
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// All valid organizations in an alert channel
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

type GrafanaInstanceInfo struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Array of subnet IDs
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Grafana private network address
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// Grafana public network address
	InternalUrl *string `json:"InternalUrl,omitnil,omitempty" name:"InternalUrl"`

	// Creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Status. Valid values: `1` (creating), `2` (running), `3` (abnormal), `4` (restarting), `5` (stopping), `6` (stopped), `7` (deleted).
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Instance tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// Instance AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Billing mode. Valid value: `1` (monthly subscription).
	InstanceChargeType *int64 `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// VPC name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Subnet name
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// The full URL used to access this instance
	RootUrl *string `json:"RootUrl,omitnil,omitempty" name:"RootUrl"`

	// Whether to enable SSO
	EnableSSO *bool `json:"EnableSSO,omitnil,omitempty" name:"EnableSSO"`

	// Version number
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Whether to enable CAM authentication during SSO
	EnableSSOCamCheck *bool `json:"EnableSSOCamCheck,omitnil,omitempty" name:"EnableSSOCamCheck"`
}

type GrafanaIntegrationConfig struct {
	// Integration ID
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`

	// Integration type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Integration content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Integration description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Grafana redirection address
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrafanaURL *string `json:"GrafanaURL,omitnil,omitempty" name:"GrafanaURL"`
}

type GrafanaNotificationChannel struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Array of notification channel template IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Update time
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Default valid organization. This parameter has been deprecated. Please use `OrganizationIds` instead.
	OrgId *string `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// Extra valid organization. This parameter has been deprecated. Please use `OrganizationIds` instead.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitnil,omitempty" name:"ExtraOrgIds"`

	// Valid organization. This parameter has been deprecated. Please use `OrganizationIds` instead.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgIds *string `json:"OrgIds,omitnil,omitempty" name:"OrgIds"`

	// All valid organizations in an alert channel
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrganizationIds *string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

type GrafanaPlugin struct {
	// Grafana plugin ID
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// Grafana plugin version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

// Predefined struct for user
type InstallPluginsRequestParams struct {
	// Plugin information
	Plugins []*GrafanaPlugin `json:"Plugins,omitnil,omitempty" name:"Plugins"`

	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InstallPluginsRequest struct {
	*tchttp.BaseRequest
	
	// Plugin information
	Plugins []*GrafanaPlugin `json:"Plugins,omitnil,omitempty" name:"Plugins"`

	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *InstallPluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Plugins")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallPluginsResponseParams struct {
	// ID of the installed plugin
	// Note: This field may return null, indicating that no valid values can be obtained.
	PluginIds []*string `json:"PluginIds,omitnil,omitempty" name:"PluginIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstallPluginsResponse struct {
	*tchttp.BaseResponse
	Response *InstallPluginsResponseParams `json:"Response"`
}

func (r *InstallPluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// Combination of instance dimensions
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type InstanceGroup struct {
	// Instance group ID.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Instance group name.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceGroupName *string `json:"InstanceGroupName,omitnil,omitempty" name:"InstanceGroupName"`
}

type InstanceGroups struct {
	// Instance group ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Instance group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type IntegrationConfiguration struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance type
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// Instance description
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// Dashboard URL
	GrafanaDashboardURL *string `json:"GrafanaDashboardURL,omitnil,omitempty" name:"GrafanaDashboardURL"`
}

type Label struct {
	// Label name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Label value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogAlarmReq struct {
	// APM instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search condition
	Filter []*LogFilterInfo `json:"Filter,omitnil,omitempty" name:"Filter"`

	// The switch to enable/disable alarm merging
	AlarmMerge *string `json:"AlarmMerge,omitnil,omitempty" name:"AlarmMerge"`

	// Alarm merging time
	AlarmMergeTime *string `json:"AlarmMergeTime,omitnil,omitempty" name:"AlarmMergeTime"`
}

type LogFilterInfo struct {
	// Field name
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Comparison operator
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Field value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ManagementCommand struct {
	// Agent installation command
	// Note: This field may return null, indicating that no valid values can be obtained.
	Install *string `json:"Install,omitnil,omitempty" name:"Install"`

	// Agent restart command
	// Note: This field may return null, indicating that no valid values can be obtained.
	Restart *string `json:"Restart,omitnil,omitempty" name:"Restart"`

	// Agent stop command
	// Note: This field may return null, indicating that no valid values can be obtained.
	Stop *string `json:"Stop,omitnil,omitempty" name:"Stop"`

	// Agent status detection command
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusCheck *string `json:"StatusCheck,omitnil,omitempty" name:"StatusCheck"`

	// Agent log detection command
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogCheck *string `json:"LogCheck,omitnil,omitempty" name:"LogCheck"`
}

type Metric struct {
	// Alarm policy type
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric name
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Metric display name
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Minimum value
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Maximum value
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// Dimension list
	Dimensions []*string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Unit
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Metric configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	MetricConfig *MetricConfig `json:"MetricConfig,omitnil,omitempty" name:"MetricConfig"`

	// Whether it is an advanced metric. 1: Yes; 0: No.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsAdvanced *int64 `json:"IsAdvanced,omitnil,omitempty" name:"IsAdvanced"`

	// Whether the advanced metric feature is enabled. 1: Yes; 0: No.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsOpen *int64 `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`

	// Integration center product ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Matching operator
	// Note: This field may return null, indicating that no valid values can be obtained.
	Operators []*Operator `json:"Operators,omitnil,omitempty" name:"Operators"`

	// Metric monitoring granularity
	// Note: This field may return null, indicating that no valid values can be obtained.
	Periods []*int64 `json:"Periods,omitnil,omitempty" name:"Periods"`


	IsLatenessMetric *int64 `json:"IsLatenessMetric,omitnil,omitempty" name:"IsLatenessMetric"`
}

type MetricConfig struct {
	// Allowed operator
	Operator []*string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Allowed data cycle in seconds
	Period []*int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Allowed number of continuous cycles
	ContinuePeriod []*int64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`
}

type MetricData struct {
	// Metric name
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Monitoring data point
	Points []*MetricDataPoint `json:"Points,omitnil,omitempty" name:"Points"`
}

type MetricDataPoint struct {
	// Combination of instance object dimensions
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Data point list
	Values []*Point `json:"Values,omitnil,omitempty" name:"Values"`
}

type MetricObjectMeaning struct {
	// Meaning of the metric in English
	En *string `json:"En,omitnil,omitempty" name:"En"`

	// Meaning of the metric in Chinese
	Zh *string `json:"Zh,omitnil,omitempty" name:"Zh"`
}

type MetricSet struct {
	// Namespace. Each Tencent Cloud product has a namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Metric Name
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Unit used by the metric
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Unit used by the metric
	UnitCname *string `json:"UnitCname,omitnil,omitempty" name:"UnitCname"`

	// Statistical period in seconds supported by the metric, such as 60 and 300
	Period []*int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Metric method during the statistical period
	Periods []*PeriodsSt `json:"Periods,omitnil,omitempty" name:"Periods"`

	// Meaning of the statistical metric
	Meaning *MetricObjectMeaning `json:"Meaning,omitnil,omitempty" name:"Meaning"`

	// Dimension description
	Dimensions []*DimensionsDesc `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Metric name (in Chinese).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MetricCName *string `json:"MetricCName,omitnil,omitempty" name:"MetricCName"`

	// Metric name (in English).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MetricEName *string `json:"MetricEName,omitnil,omitempty" name:"MetricEName"`
}

type MidQueryCondition struct {
	// Dimension
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Operator. Valid values: eq (equal to), ne (not equal to), in
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Dimension value. If `Operator` is `eq` or `ne`, only the first element will be used
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyAlarmNoticeRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm notification rule name, which can contain up to 60 characters
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	NoticeType *string `json:"NoticeType,omitnil,omitempty" name:"NoticeType"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	NoticeLanguage *string `json:"NoticeLanguage,omitnil,omitempty" name:"NoticeLanguage"`

	// Alarm notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// User notifications (up to 5)
	UserNotices []*UserNotice `json:"UserNotices,omitnil,omitempty" name:"UserNotices"`

	// Callback notifications (up to 3)
	URLNotices []*URLNotice `json:"URLNotices,omitnil,omitempty" name:"URLNotices"`

	// The operation of pushing alarm notifications to CLS. Up to one CLS log topic can be configured.
	CLSNotices []*CLSNotice `json:"CLSNotices,omitnil,omitempty" name:"CLSNotices"`

	// List of IDs of the alerting rules bound to an alarm notification template
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm notification rule name, which can contain up to 60 characters
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values: ALARM (for unresolved alarms), OK (for resolved alarms), ALL (for all alarms)
	NoticeType *string `json:"NoticeType,omitnil,omitempty" name:"NoticeType"`

	// Notification language. Valid values: zh-CN (Chinese), en-US (English)
	NoticeLanguage *string `json:"NoticeLanguage,omitnil,omitempty" name:"NoticeLanguage"`

	// Alarm notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// User notifications (up to 5)
	UserNotices []*UserNotice `json:"UserNotices,omitnil,omitempty" name:"UserNotices"`

	// Callback notifications (up to 3)
	URLNotices []*URLNotice `json:"URLNotices,omitnil,omitempty" name:"URLNotices"`

	// The operation of pushing alarm notifications to CLS. Up to one CLS log topic can be configured.
	CLSNotices []*CLSNotice `json:"CLSNotices,omitnil,omitempty" name:"CLSNotices"`

	// List of IDs of the alerting rules bound to an alarm notification template
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Name")
	delete(f, "NoticeType")
	delete(f, "NoticeLanguage")
	delete(f, "NoticeId")
	delete(f, "UserNotices")
	delete(f, "URLNotices")
	delete(f, "CLSNotices")
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmNoticeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmNoticeResponseParams `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyConditionRequestParams struct {
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// ID of trigger condition template. This parameter can be left empty.
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitnil,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition
	Condition *AlarmPolicyCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Event trigger condition
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitnil,omitempty" name:"EventCondition"`

	// Global filter.
	Filter *AlarmPolicyFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Aggregation dimension list, which is used to specify which dimension keys data is grouped by.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Log alarm creation request parameters
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitnil,omitempty" name:"LogAlarmReqInfo"`

	// Template ID, which is dedicated to TencentCloud Managed Service for Prometheus.
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Status (`0`: Disabled; `1`: Enabled)
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Name of the policy dedicated to TMP
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`
}

type ModifyAlarmPolicyConditionRequest struct {
	*tchttp.BaseRequest
	
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// ID of trigger condition template. This parameter can be left empty.
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitnil,omitempty" name:"ConditionTemplateId"`

	// Metric trigger condition
	Condition *AlarmPolicyCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Event trigger condition
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitnil,omitempty" name:"EventCondition"`

	// Global filter.
	Filter *AlarmPolicyFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Aggregation dimension list, which is used to specify which dimension keys data is grouped by.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Log alarm creation request parameters
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitnil,omitempty" name:"LogAlarmReqInfo"`

	// Template ID, which is dedicated to TencentCloud Managed Service for Prometheus.
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Status (`0`: Disabled; `1`: Enabled)
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Name of the policy dedicated to TMP
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`
}

func (r *ModifyAlarmPolicyConditionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyConditionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "ConditionTemplateId")
	delete(f, "Condition")
	delete(f, "EventCondition")
	delete(f, "Filter")
	delete(f, "GroupBy")
	delete(f, "LogAlarmReqInfo")
	delete(f, "NoticeIds")
	delete(f, "Enable")
	delete(f, "PolicyName")
	delete(f, "EbSubject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyConditionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyConditionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyConditionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyConditionResponseParams `json:"Response"`
}

func (r *ModifyAlarmPolicyConditionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyConditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyInfoRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Field to be modified. Valid values: NAME (policy name), REMARK (policy remarks)
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value after modification
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ModifyAlarmPolicyInfoRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Field to be modified. Valid values: NAME (policy name), REMARK (policy remarks)
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value after modification
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

func (r *ModifyAlarmPolicyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "Key")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyInfoResponseParams `json:"Response"`
}

func (r *ModifyAlarmPolicyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyNoticeRequestParams struct {
	// Module name, which is specified as `monitor`.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID. If both `PolicyIds` and this parameter are returned, only `PolicyIds` takes effect.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// List of alarm notification template IDs.
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Alarm policy ID array, which can be used to associate notification templates with multiple alarm policies. Max value: 30.
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// Notification rules for different alarm levels
	HierarchicalNotices []*AlarmHierarchicalNotice `json:"HierarchicalNotices,omitnil,omitempty" name:"HierarchicalNotices"`
}

type ModifyAlarmPolicyNoticeRequest struct {
	*tchttp.BaseRequest
	
	// Module name, which is specified as `monitor`.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID. If both `PolicyIds` and this parameter are returned, only `PolicyIds` takes effect.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// List of alarm notification template IDs.
	NoticeIds []*string `json:"NoticeIds,omitnil,omitempty" name:"NoticeIds"`

	// Alarm policy ID array, which can be used to associate notification templates with multiple alarm policies. Max value: 30.
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// Notification rules for different alarm levels
	HierarchicalNotices []*AlarmHierarchicalNotice `json:"HierarchicalNotices,omitnil,omitempty" name:"HierarchicalNotices"`
}

func (r *ModifyAlarmPolicyNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "NoticeIds")
	delete(f, "PolicyIds")
	delete(f, "HierarchicalNotices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyNoticeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyNoticeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyNoticeResponseParams `json:"Response"`
}

func (r *ModifyAlarmPolicyNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyStatusRequestParams struct {
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Status. Valid values: 0 (disabled), 1 (enabled)
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ModifyAlarmPolicyStatusRequest struct {
	*tchttp.BaseRequest
	
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Status. Valid values: 0 (disabled), 1 (enabled)
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *ModifyAlarmPolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyStatusResponseParams `json:"Response"`
}

func (r *ModifyAlarmPolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyTasksRequestParams struct {
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// List of tasks triggered by alarm policy. If this parameter is left empty, it indicates to unbind all tasks
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitnil,omitempty" name:"TriggerTasks"`
}

type ModifyAlarmPolicyTasksRequest struct {
	*tchttp.BaseRequest
	
	// Module name. Enter "monitor" here
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// List of tasks triggered by alarm policy. If this parameter is left empty, it indicates to unbind all tasks
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitnil,omitempty" name:"TriggerTasks"`
}

func (r *ModifyAlarmPolicyTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "TriggerTasks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyTasksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyTasksResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyTasksResponseParams `json:"Response"`
}

func (r *ModifyAlarmPolicyTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmReceiversRequestParams struct {
	// ID of a policy group whose recipient needs to be modified.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Required. The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// New recipient information. If this parameter is not configured, all recipients will be deleted.
	ReceiverInfos []*ReceiverInfo `json:"ReceiverInfos,omitnil,omitempty" name:"ReceiverInfos"`
}

type ModifyAlarmReceiversRequest struct {
	*tchttp.BaseRequest
	
	// ID of a policy group whose recipient needs to be modified.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Required. The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// New recipient information. If this parameter is not configured, all recipients will be deleted.
	ReceiverInfos []*ReceiverInfo `json:"ReceiverInfos,omitnil,omitempty" name:"ReceiverInfos"`
}

func (r *ModifyAlarmReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmReceiversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Module")
	delete(f, "ReceiverInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmReceiversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmReceiversResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmReceiversResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmReceiversResponseParams `json:"Response"`
}

func (r *ModifyAlarmReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGrafanaInstanceRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TCMG instance name, such as “test”.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TCMG instance name, such as “test”.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGrafanaInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGrafanaInstanceResponseParams `json:"Response"`
}

func (r *ModifyGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupCondition struct {
	// Metric ID.
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to.
	CalcType *int64 `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// Threshold.
	CalcValue *string `json:"CalcValue,omitnil,omitempty" name:"CalcValue"`

	// Data period of the detected metric.
	CalcPeriod *int64 `json:"CalcPeriod,omitnil,omitempty" name:"CalcPeriod"`

	// Number of consecutive periods.
	ContinuePeriod *int64 `json:"ContinuePeriod,omitnil,omitempty" name:"ContinuePeriod"`

	// Alarm sending and convergence type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. If the value is less than 0, no alarm will be triggered. If the value is 0, an alarm will be triggered only once. If the value is greater than 0, an alarm will be triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// Rule ID. No filling means new addition while filling in ruleId means to modify existing rules.
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type ModifyPolicyGroupEventCondition struct {
	// Event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Alarm sending and convergence type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitnil,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. If the value is less than 0, no alarm will be triggered. If the value is 0, an alarm will be triggered only once. If the value is greater than 0, an alarm will be triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitnil,omitempty" name:"AlarmNotifyPeriod"`

	// Rule ID. No filling means new addition while filling in ruleId means to modify existing rules.
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

// Predefined struct for user
type ModifyPolicyGroupRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm type.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The 'AND' and 'OR' rules for metric alarms. The value 1 indicates 'AND', which means that an alarm will be triggered only when all rules are met. The value 0 indicates 'OR', which means that an alarm will be triggered when any rule is met.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`

	// Metric alarm condition rules. No filling indicates that all existing metric alarm condition rules will be deleted.
	Conditions []*ModifyPolicyGroupCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Event alarm conditions. No filling indicates that all existing event alarm conditions will be deleted.
	EventConditions []*ModifyPolicyGroupEventCondition `json:"EventConditions,omitnil,omitempty" name:"EventConditions"`

	// Template-based policy group ID.
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitnil,omitempty" name:"ConditionTempGroupId"`
}

type ModifyPolicyGroupRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm type.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// The 'AND' and 'OR' rules for metric alarms. The value 1 indicates 'AND', which means that an alarm will be triggered only when all rules are met. The value 0 indicates 'OR', which means that an alarm will be triggered when any rule is met.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`

	// Metric alarm condition rules. No filling indicates that all existing metric alarm condition rules will be deleted.
	Conditions []*ModifyPolicyGroupCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Event alarm conditions. No filling indicates that all existing event alarm conditions will be deleted.
	EventConditions []*ModifyPolicyGroupEventCondition `json:"EventConditions,omitnil,omitempty" name:"EventConditions"`

	// Template-based policy group ID.
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitnil,omitempty" name:"ConditionTempGroupId"`
}

func (r *ModifyPolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "ViewName")
	delete(f, "GroupName")
	delete(f, "IsUnionRule")
	delete(f, "Conditions")
	delete(f, "EventConditions")
	delete(f, "ConditionTempGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPolicyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPolicyGroupResponseParams struct {
	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPolicyGroupResponseParams `json:"Response"`
}

func (r *ModifyPolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAgentExternalLabelsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// New external labels
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`
}

type ModifyPrometheusAgentExternalLabelsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// New external labels
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`
}

func (r *ModifyPrometheusAgentExternalLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAgentExternalLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterId")
	delete(f, "ExternalLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAgentExternalLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAgentExternalLabelsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusAgentExternalLabelsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAgentExternalLabelsResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAgentExternalLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAgentExternalLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert configuration
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

type ModifyPrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert configuration
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

func (r *ModifyPrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Configuration of service monitors
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// Configuration of pod monitors
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// Configuration of Prometheus raw jobs
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`
}

type ModifyPrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Configuration of service monitors
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// Configuration of pod monitors
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// Configuration of Prometheus raw jobs
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`
}

func (r *ModifyPrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusConfigResponseParams `json:"Response"`
}

func (r *ModifyPrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusGlobalNotificationRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert notification channel
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

type ModifyPrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert notification channel
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

func (r *ModifyPrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Notification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusGlobalNotificationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *ModifyPrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusInstanceAttributesRequestParams struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Storage period. Valid values: 15, 30, 45. This parameter is not applicable to monthly subscribed instances.
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`
}

type ModifyPrometheusInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Storage period. Valid values: 15, 30, 45. This parameter is not applicable to monthly subscribed instances.
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`
}

func (r *ModifyPrometheusInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusInstanceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "InstanceId")
	delete(f, "DataRetentionTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusInstanceAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusInstanceAttributesResponseParams `json:"Response"`
}

func (r *ModifyPrometheusInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusRecordRuleYamlRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Recording instance name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// New content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ModifyPrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Recording instance name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// New content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *ModifyPrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusRecordRuleYamlResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *ModifyPrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTempRequestParams struct {
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Modified content
	Template *PrometheusTempModify `json:"Template,omitnil,omitempty" name:"Template"`
}

type ModifyPrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// Template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Modified content
	Template *PrometheusTempModify `json:"Template,omitnil,omitempty" name:"Template"`
}

func (r *ModifyPrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTempResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusTempResponseParams `json:"Response"`
}

func (r *ModifyPrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTypeInfo struct {
	// Monitoring type ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Monitoring type
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sort order
	SortId *int64 `json:"SortId,omitnil,omitempty" name:"SortId"`
}

type MonitorTypeNamespace struct {
	// Monitor type
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Policy type value
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type NoticeBindPolicys struct {
	// Alert notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// List of IDs of the alerting rules bound to an alarm notification template
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`
}

type Operator struct {
	// Operator ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Operator name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type PeriodsSt struct {
	// Period
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Statistical method
	StatType []*string `json:"StatType,omitnil,omitempty" name:"StatType"`
}

type Point struct {
	// Time point when this monitoring data point is generated
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Monitoring data point value
	// Note: this field may return null, indicating that no valid values can be obtained.
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type PolicyGroup struct {
	// Whether the alarm policy can be set to default.
	CanSetDefault *bool `json:"CanSetDefault,omitnil,omitempty" name:"CanSetDefault"`

	// Alarm policy group ID.
	GroupID *int64 `json:"GroupID,omitnil,omitempty" name:"GroupID"`

	// Alarm policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Whether the alarm policy is set to default.
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Whether the alarm policy is enabled.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// UIN of the last modifier.
	LastEditUin *int64 `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Number of unshielded instances.
	NoShieldedInstanceCount *int64 `json:"NoShieldedInstanceCount,omitnil,omitempty" name:"NoShieldedInstanceCount"`

	// Parent policy group ID.
	ParentGroupID *int64 `json:"ParentGroupID,omitnil,omitempty" name:"ParentGroupID"`

	// Project ID.
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// Alarm recipient information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ReceiverInfos []*PolicyGroupReceiverInfo `json:"ReceiverInfos,omitnil,omitempty" name:"ReceiverInfos"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Modification time.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The total number of associated instances.
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitnil,omitempty" name:"TotalInstanceCount"`

	// View.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Whether the logical relationship between rules is AND.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`
}

type PolicyGroupReceiverInfo struct {
	// End time of a valid time period.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether it is required to send notifications.
	NeedSendNotice *int64 `json:"NeedSendNotice,omitnil,omitempty" name:"NeedSendNotice"`

	// Alarm receiving channel.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NotifyWay []*string `json:"NotifyWay,omitnil,omitempty" name:"NotifyWay"`

	// Alarm call intervals for individuals in seconds.
	PersonInterval *int64 `json:"PersonInterval,omitnil,omitempty" name:"PersonInterval"`

	// Message recipient group list.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitnil,omitempty" name:"ReceiverGroupList"`

	// Recipient type.
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Recipient list. The list of recipient IDs that is queried by a platform API.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitnil,omitempty" name:"ReceiverUserList"`

	// Alarm resolution notification method.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RecoverNotify []*string `json:"RecoverNotify,omitnil,omitempty" name:"RecoverNotify"`

	// Alarm call interval per round in seconds.
	RoundInterval *int64 `json:"RoundInterval,omitnil,omitempty" name:"RoundInterval"`

	// Number of alarm call rounds.
	RoundNumber *int64 `json:"RoundNumber,omitnil,omitempty" name:"RoundNumber"`

	// Alarm call notification time. Valid values: `OCCUR` (indicating that a notification is sent when the alarm is triggered) and `RECOVER` (indicating that a notification is sent when the alarm is resolved).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SendFor []*string `json:"SendFor,omitnil,omitempty" name:"SendFor"`

	// Start time of a valid time period.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// UID of the alarm call recipient.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UIDList []*int64 `json:"UIDList,omitnil,omitempty" name:"UIDList"`
}

type PrometheusAgent struct {
	// Agent name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ipv4 *string `json:"Ipv4,omitnil,omitempty" name:"Ipv4"`

	// Heartbeat time
	// Note: This field may return null, indicating that no valid values can be obtained.
	HeartbeatTime *string `json:"HeartbeatTime,omitnil,omitempty" name:"HeartbeatTime"`

	// Last error
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastError *string `json:"LastError,omitnil,omitempty" name:"LastError"`

	// Agent version
	// Note: This field may return null, indicating that no valid values can be obtained.
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// Agent status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type PrometheusAgentInfo struct {
	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Remarks
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`
}

type PrometheusAgentOverview struct {
	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Agent status. Valid values: 
	// `normal`
	// `abnormal`
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// External labels
	// External labels, which will be attached to all metrics in this cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`

	// Cluster region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// ID of the VPC where the cluster resides
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Recorded information of failed operations, such as association.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// Agent name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type PrometheusAlertAllowTimeRange struct {
	// Seconds from 00:00:00.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Start *string `json:"Start,omitnil,omitempty" name:"Start"`

	// Seconds from 00:00:00.
	// Note: This field may return null, indicating that no valid values can be obtained.
	End *string `json:"End,omitnil,omitempty" name:"End"`
}

type PrometheusAlertCustomReceiver struct {
	// Notification customization type.
	// Alertmanager - self-built alertmanager in vpc.
	// webhook - webhook address in the vpc.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// alertmanager/webhook url (ip in the same vpc as the prometheus instance).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Specifies the time range for allowing Alert sending.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AllowedTimeRanges []*PrometheusAlertAllowTimeRange `json:"AllowedTimeRanges,omitnil,omitempty" name:"AllowedTimeRanges"`

	// alertmanager intranet cluster ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// alertmanager resides in the private network cluster type (tke/eks/tdcc).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type PrometheusAlertGroupRuleSet struct {
	// Alert rule name. same name is not allowed in the same Alert group.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Prometheus alert label list.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Labels []*PrometheusRuleKV `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Prometheus alert annotation list.
	// 
	// Alarm object and Alarm message are special fields of Prometheus Rule Annotations. they need to pass through Annotations for transmission. the corresponding keys are summary and description respectively.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Annotations []*PrometheusRuleKV `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// Alert will be triggered after 'Expr' satisfied for 'Duration'.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Alert expression. see <a href="https://www.tencentcloud.comom/document/product/1416/56008?from_cn_redirect=1">alert rule description</a>.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Expr *string `json:"Expr,omitnil,omitempty" name:"Expr"`

	// Alert rule status.
	// 2 - enable.
	// 3 - disabled.
	// Enabled by default if left empty.
	// Note: This field may return null, indicating that no valid values can be obtained.
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

type PrometheusAlertGroupSet struct {
	// Alert group ID. must match the regular expression `alert-[a-z0-9]{8}`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alert group name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Alert template ID of tencent cloud observability platform (tcop). returns the converted notice ID of the Alert template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AMPReceivers []*string `json:"AMPReceivers,omitnil,omitempty" name:"AMPReceivers"`

	// Custom Alert template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CustomReceiver *PrometheusAlertCustomReceiver `json:"CustomReceiver,omitnil,omitempty" name:"CustomReceiver"`

	// Alert notification interval.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// If the Alert group is created via template, returns the template ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Alert rule details within the group.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*PrometheusAlertGroupRuleSet `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Group creation time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Group update time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type PrometheusAlertManagerConfig struct {
	// AlertManager URL
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Type of the cluster where AlertManager is deployed
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// ID of the cluster where AlertManager is deployed
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type PrometheusAlertPolicyItem struct {
	// Rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of rules
	Rules []*PrometheusAlertRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Alerting rule ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// If the alert comes from a template, `TemplateId` is the template ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Alert channel, which may be returned as null if used in a template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`

	// Last modification time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// If the alerting rule comes from the user cluster CRD resource definition, `ClusterId` is the cluster ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type PrometheusAlertRule struct {
	// Rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Prometheus statement
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Additional tags
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Alert sending template
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// Duration
	For *string `json:"For,omitnil,omitempty" name:"For"`

	// Rule description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// See `annotations` in the Prometheus rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	Annotations []*Label `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// Alerting rule status
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

type PrometheusClusterAgentBasic struct {
	// Cluster ID
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Cluster type
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether to enable public network CLB
	EnableExternal *bool `json:"EnableExternal,omitnil,omitempty" name:"EnableExternal"`

	// Pod configurations of components deployed in the cluster
	InClusterPodConfig *PrometheusClusterAgentPodConfig `json:"InClusterPodConfig,omitnil,omitempty" name:"InClusterPodConfig"`

	// External labels, which will be attached to all metrics collected by this cluster
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`

	// Whether to install the default collection configuration.
	NotInstallBasicScrape *bool `json:"NotInstallBasicScrape,omitnil,omitempty" name:"NotInstallBasicScrape"`

	// Whether to collect metrics (`true`: Drop all metrics; `false`: Collect default metrics)
	NotScrape *bool `json:"NotScrape,omitnil,omitempty" name:"NotScrape"`

	// Whether to enable the default recording rule
	OpenDefaultRecord *bool `json:"OpenDefaultRecord,omitnil,omitempty" name:"OpenDefaultRecord"`
}

type PrometheusClusterAgentPodConfig struct {
	// Whether to use HostNetWork
	HostNet *bool `json:"HostNet,omitnil,omitempty" name:"HostNet"`

	// A parameter used to specify the running nodes for a pod
	NodeSelector []*Label `json:"NodeSelector,omitnil,omitempty" name:"NodeSelector"`

	// Tolerable taints
	Tolerations []*Toleration `json:"Tolerations,omitnil,omitempty" name:"Tolerations"`
}

type PrometheusConfigItem struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Configuration content
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// If the configuration comes from a template, this parameter is the template ID, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Number of targets
	// Note: This field may return null, indicating that no valid values can be obtained.
	Targets *Targets `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type PrometheusInstanceGrantInfo struct {
	// Whether there is permission to operate on the billing information. Valid values: 1 (yes), 2 (no).
	HasChargeOperation *int64 `json:"HasChargeOperation,omitnil,omitempty" name:"HasChargeOperation"`

	// Whether there is permission to display the VPC information. Valid values: 1 (yes), 2 (no).
	HasVpcDisplay *int64 `json:"HasVpcDisplay,omitnil,omitempty" name:"HasVpcDisplay"`

	// Whether there is permission to change the Grafana status. Valid values: 1 (yes), 2 (no).
	HasGrafanaStatusChange *int64 `json:"HasGrafanaStatusChange,omitnil,omitempty" name:"HasGrafanaStatusChange"`

	// Whether there is permission to manage agents. Valid values: 1 (yes), 2 (no).
	HasAgentManage *int64 `json:"HasAgentManage,omitnil,omitempty" name:"HasAgentManage"`

	// Whether there is permission to manage TKE integrations. Valid values: 1 (yes), 2 (no).
	HasTkeManage *int64 `json:"HasTkeManage,omitnil,omitempty" name:"HasTkeManage"`

	// Whether there is permission to display the API information. Valid values: 1 (yes), 2 (no).
	HasApiOperation *int64 `json:"HasApiOperation,omitnil,omitempty" name:"HasApiOperation"`
}

type PrometheusInstanceTenantUsage struct {
	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Billing cycle
	// Note: This field may return null, indicating that no valid values can be obtained.
	CalcDate *string `json:"CalcDate,omitnil,omitempty" name:"CalcDate"`

	// Total usage
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *float64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Usage of basic (free) metrics
	// Note: This field may return null, indicating that no valid values can be obtained.
	Basic *float64 `json:"Basic,omitnil,omitempty" name:"Basic"`

	// Usage of paid metrics
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fee *float64 `json:"Fee,omitnil,omitempty" name:"Fee"`
}

type PrometheusInstancesItem struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance billing mode. Valid values:
	// <ul>
	// <li>2: Monthly subscription</li>
	// <li>3: Pay-as-you-go</li>
	// </ul>
	InstanceChargeType *int64 `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Storage period
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`

	// Instance status. Valid values:
	// <ul>
	// <li>1: Creating</li>
	// <li>2: Running</li>
	// <li>3: Abnormal</li>
	// <li>4: Rebooting</li>
	// <li>5: Terminating</li>
	// <li>6: Service suspended</li>
	// <li>8: Suspending service for overdue payment</li>
	// <li>9: Service suspended for overdue payment</li>
	// </ul>
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Grafana dashboard URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrafanaURL *string `json:"GrafanaURL,omitnil,omitempty" name:"GrafanaURL"`

	// Creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Whether Grafana is enabled
	// <li>0: Disabled</li>
	// <li>1: Enabled</li>
	EnableGrafana *int64 `json:"EnableGrafana,omitnil,omitempty" name:"EnableGrafana"`

	// Instance IPv4 address
	// Note: This field may return null, indicating that no valid values can be obtained.
	IPv4Address *string `json:"IPv4Address,omitnil,omitempty" name:"IPv4Address"`

	// List of tags associated with the instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// Expiration time of the purchased instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Billing status
	// <ul>
	// <li>1: Normal</li>
	// <li>2: Expired</li>
	// <li>3: Terminated</li>
	// <li>4: Assigning</li>
	// <li>5: Assignment failed</li>
	// </ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeStatus *int64 `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// Specification name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Auto-renewal flag
	// <ul>
	// <li>0: Auto-renewal not enabled</li>
	// <li>1: Auto-renewal enabled</li>
	// <li>2: Auto-renewal prohibited</li>
	// <li>-1: Invalid</ii>
	// </ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Expiring soon
	// <ul>
	// <li>0: No</li>
	// <li>1: Yes</li>
	// </ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsNearExpire *int64 `json:"IsNearExpire,omitnil,omitempty" name:"IsNearExpire"`

	// The token required for data writing
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthToken *string `json:"AuthToken,omitnil,omitempty" name:"AuthToken"`

	// Prometheus remote write address
	// Note: This field may return null, indicating that no valid values can be obtained.
	RemoteWrite *string `json:"RemoteWrite,omitnil,omitempty" name:"RemoteWrite"`

	// Prometheus HTTP API root address
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApiRootPath *string `json:"ApiRootPath,omitnil,omitempty" name:"ApiRootPath"`

	// Proxy address
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyAddress *string `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`

	// Grafana status
	// <ul>
	// <li>1: Creating</li>
	// <li>2: Running</li>
	// <li>3: Abnormal</li>
	// <li>4: Restarting</li>
	// <li>5: Terminating</li>
	// <li>6: Service suspended</li>
	// <li>7: Deleted</li>
	// </ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrafanaStatus *int64 `json:"GrafanaStatus,omitnil,omitempty" name:"GrafanaStatus"`

	// Grafana IP allowlist, where IPs are separated by comma.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrafanaIpWhiteList *string `json:"GrafanaIpWhiteList,omitnil,omitempty" name:"GrafanaIpWhiteList"`

	// Instance authorization information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Grant *PrometheusInstanceGrantInfo `json:"Grant,omitnil,omitempty" name:"Grant"`

	// ID of the bound Grafana instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitnil,omitempty" name:"GrafanaInstanceId"`

	// The alert rule limit
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlertRuleLimit *int64 `json:"AlertRuleLimit,omitnil,omitempty" name:"AlertRuleLimit"`

	// The recording rule limit
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecordingRuleLimit *int64 `json:"RecordingRuleLimit,omitnil,omitempty" name:"RecordingRuleLimit"`

	// Migration status. 0: Not migrating; 1: Migrating from source instance; 2: Migrating to target instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MigrationType *int64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`
}

type PrometheusInstancesOverview struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Running status. Valid values: `1` (creating); `2` (running); `3` (abnormal); `4` (restarting); `5` (terminating); `6` (stopped); `7` (deleted).
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Billing status. Valid values: `1` (normal); `2` (expired); `3` (terminated); `4` (assigning); `5` (failed to assign)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeStatus *int64 `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// Whether Grafana is enabled. Valid values: `0` (no); `1` (yes).
	EnableGrafana *int64 `json:"EnableGrafana,omitnil,omitempty" name:"EnableGrafana"`

	// Grafana dashboard URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrafanaURL *string `json:"GrafanaURL,omitnil,omitempty" name:"GrafanaURL"`

	// Instance payment type. Valid values: `1` (trial edition); `2` (prepaid)
	InstanceChargeType *int64 `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Specification name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Storage period
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`

	// Expiration time of the purchased instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Auto-renewal flag. Valid values: `0` (auto-renewal not enabled); `1` (auto-renewal enabled); `2` (auto-renewal prohibited); `-1` (invalid).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Total number of bound clusters
	BoundTotal *int64 `json:"BoundTotal,omitnil,omitempty" name:"BoundTotal"`

	// Total number of bound clusters in the normal status
	BoundNormal *int64 `json:"BoundNormal,omitnil,omitempty" name:"BoundNormal"`

	// Resource pack status (`0`: Unavailable; `1`: Available)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourcePackageStatus *int64 `json:"ResourcePackageStatus,omitnil,omitempty" name:"ResourcePackageStatus"`

	// Resource pack specification name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourcePackageSpecName *string `json:"ResourcePackageSpecName,omitnil,omitempty" name:"ResourcePackageSpecName"`
}

type PrometheusJobTargets struct {

}

type PrometheusNotificationItem struct {
	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Channel type. Default value: `amp`. Valid values:
	// `amp`
	// `webhook`
	// `alertmanager`
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// If `Type` is `webhook`, this field is required.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WebHook *string `json:"WebHook,omitnil,omitempty" name:"WebHook"`

	// If `Type` is `alertmanager`, this field is required.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlertManager *PrometheusAlertManagerConfig `json:"AlertManager,omitnil,omitempty" name:"AlertManager"`

	// Convergence time
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// Effect start time
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// Effect end time
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// Alert notification channel. Valid values: `SMS`, `EMAIL`, `CALL`, `WECHAT`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NotifyWay []*string `json:"NotifyWay,omitnil,omitempty" name:"NotifyWay"`

	// Alert recipient group (user group)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReceiverGroups []*string `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// Alert call sequence.
	// Note: If `NotifyWay` is `CALL`, this parameter will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneNotifyOrder []*uint64 `json:"PhoneNotifyOrder,omitnil,omitempty" name:"PhoneNotifyOrder"`

	// Number of alert calls.
	// Note: If `NotifyWay` is `CALL`, this parameter will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitnil,omitempty" name:"PhoneCircleTimes"`

	// Alert call interval within a cycle in seconds.
	// Note: If `NotifyWay` is `CALL`, this parameter will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitnil,omitempty" name:"PhoneInnerInterval"`

	// Alert call cycle interval in seconds.
	// Note: If `NotifyWay` is `CALL`, this parameter will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitnil,omitempty" name:"PhoneCircleInterval"`

	// Alert call receipt notification
	// Note: If `NotifyWay` is `CALL`, this parameter will be used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneArriveNotice *bool `json:"PhoneArriveNotice,omitnil,omitempty" name:"PhoneArriveNotice"`
}

type PrometheusRecordRuleYamlItem struct {
	// Instance name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Last update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// YAML content
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// If the recording rule comes from a template, `TemplateId` is the template ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// If the recording rule comes from the user cluster CRD resource definition, `ClusterId` is the cluster ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// id
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Number of rules
	// Note: This field may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type PrometheusRuleKV struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PrometheusRuleSet struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Rule status code
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Rule category
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// List of rule tags
	// Note: This field may return null, indicating that no valid values can be obtained.
	Labels []*PrometheusRuleKV `json:"Labels,omitnil,omitempty" name:"Labels"`

	// List of rule annotations
	// Note: This field may return null, indicating that no valid values can be obtained.
	Annotations []*PrometheusRuleKV `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// Rule expression
	// Note: This field may return null, indicating that no valid values can be obtained.
	Expr *string `json:"Expr,omitnil,omitempty" name:"Expr"`

	// Rule alert duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// List of alert recipient groups
	// Note: This field may return null, indicating that no valid values can be obtained.
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Rule status. Valid values:
	// <li>unknown: Unknown</li>
	// <li>pending: Loading</li>
	// <li>ok: Running</li>
	// <li>err: Error</li>
	Health *string `json:"Health,omitnil,omitempty" name:"Health"`

	// Rule creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Rule update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type PrometheusScrapeJob struct {
	// Task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Agent ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type PrometheusTag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PrometheusTemp struct {
	// Template name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template dimension. Valid values:
	// `instance`
	// `cluster`
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Template description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// This parameter is valid if `Level` is `instance`.
	// List of recording rules in the template
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitnil,omitempty" name:"RecordRules"`

	// This parameter is valid if `Level` is `cluster`.
	// List of ServiceMonitor rules in the template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// This parameter is valid if `Level` is `cluster`.
	// List of PodMonitor rules in the template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// This parameter is valid if `Level` is `cluster`.
	// List of RawJob rules in the template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// Template ID, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Last update time, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The current version, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Whether it is the default template provided by the system, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// This parameter is valid if `Level` is `instance`.
	// List of alert configurations in the template
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlertDetailRules []*PrometheusAlertPolicyItem `json:"AlertDetailRules,omitnil,omitempty" name:"AlertDetailRules"`

	// Number of associated instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetsTotal *int64 `json:"TargetsTotal,omitnil,omitempty" name:"TargetsTotal"`
}

type PrometheusTempModify struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// This parameter is valid if `Level` is `cluster`.
	// List of ServiceMonitor rules in the template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// This parameter is valid if `Level` is `cluster`.
	// List of PodMonitor rules in the template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// This parameter is valid if `Level` is `cluster`.
	// List of RawJob rules in the template.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// This parameter is valid if `Level` is `instance`.
	// List of recording rules in the template
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitnil,omitempty" name:"RecordRules"`

	// Modification content, which is valid only if template type is `Alert`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlertDetailRules []*PrometheusAlertPolicyItem `json:"AlertDetailRules,omitnil,omitempty" name:"AlertDetailRules"`
}

type PrometheusTemplateSyncTarget struct {
	// Target region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Target instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster ID, which is required only if the `Level` of the collection template is `cluster`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Last sync time, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SyncTime *string `json:"SyncTime,omitnil,omitempty" name:"SyncTime"`

	// The currently used template version, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Cluster type, which is required only if the `Level` of the collection template is `cluster`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Instance name, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Cluster name, which is used as an output parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type PrometheusZoneItem struct {
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// AZ status. Valid values: `0`(Unavailable), `1` (Available).
	ZoneState *int64 `json:"ZoneState,omitnil,omitempty" name:"ZoneState"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type ReceiverInfo struct {
	// Start time of the alarm period. Value range: [0,86400). Convert the Unix timestamp to Beijing time and then remove the date. For example, 7200 indicates “10:0:0”.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the alarm period. The meaning is the same as that of StartTime.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Alarm notification method. Valid values: "SMS", "SITE", "EMAIL", "CALL", and "WECHAT".
	NotifyWay []*string `json:"NotifyWay,omitnil,omitempty" name:"NotifyWay"`

	// Recipient type. Valid values: group and user.
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// ReceiverId
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Alarm call notification time. Valid values: OCCUR (indicating that a notice is sent when the alarm is triggered) and RECOVER (indicating that a notice is sent when the alarm is recovered).
	SendFor []*string `json:"SendFor,omitnil,omitempty" name:"SendFor"`

	// UID of the phone call alarm.
	UidList []*int64 `json:"UidList,omitnil,omitempty" name:"UidList"`

	// Number of alarm call rounds.
	RoundNumber *int64 `json:"RoundNumber,omitnil,omitempty" name:"RoundNumber"`

	// Alarm call intervals for individuals in seconds.
	PersonInterval *int64 `json:"PersonInterval,omitnil,omitempty" name:"PersonInterval"`

	// Intervals of alarm call rounds in seconds.
	RoundInterval *int64 `json:"RoundInterval,omitnil,omitempty" name:"RoundInterval"`

	// Notification method when an alarm is recovered. Valid value: SMS.
	RecoverNotify []*string `json:"RecoverNotify,omitnil,omitempty" name:"RecoverNotify"`

	// Whether to send an alarm call delivery notice. The value 0 indicates that no notice needs to be sent. The value 1 indicates that a notice needs to be sent.
	NeedSendNotice *int64 `json:"NeedSendNotice,omitnil,omitempty" name:"NeedSendNotice"`

	// Recipient group list. The list of recipient group IDs that is queried by API.
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitnil,omitempty" name:"ReceiverGroupList"`

	// Recipient list. The list of recipient IDs that is queried by API.
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitnil,omitempty" name:"ReceiverUserList"`

	// Language of received alarms. Enumerated values: zh-CN and en-US.
	ReceiveLanguage *string `json:"ReceiveLanguage,omitnil,omitempty" name:"ReceiveLanguage"`
}

type RecordingRuleSet struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status code
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Rule group
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Number of rules
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Rule creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Rule update time
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Rule name
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

// Predefined struct for user
type ResumeGrafanaInstanceRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ResumeGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ResumeGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeGrafanaInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeGrafanaInstanceResponseParams `json:"Response"`
}

func (r *ResumeGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunPrometheusInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Subnet ID. Initialization is performed with the subnet used by the instance by default and can also be performed with the subnet passed in by this parameter.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type RunPrometheusInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Subnet ID. Initialization is performed with the subnet used by the instance by default and can also be performed with the subnet passed in by this parameter.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

func (r *RunPrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunPrometheusInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunPrometheusInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunPrometheusInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunPrometheusInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RunPrometheusInstanceResponseParams `json:"Response"`
}

func (r *RunPrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunPrometheusInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendCustomAlarmMsgRequestParams struct {
	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Message policy ID, which is configured on the custom message page.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Custom message content that a user wants to send.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`
}

type SendCustomAlarmMsgRequest struct {
	*tchttp.BaseRequest
	
	// API component name. The value for the current API is monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Message policy ID, which is configured on the custom message page.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Custom message content that a user wants to send.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`
}

func (r *SendCustomAlarmMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCustomAlarmMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "Msg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendCustomAlarmMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendCustomAlarmMsgResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendCustomAlarmMsgResponse struct {
	*tchttp.BaseResponse
	Response *SendCustomAlarmMsgResponseParams `json:"Response"`
}

func (r *SendCustomAlarmMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCustomAlarmMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceDiscoveryItem struct {
	// Scrape configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Namespace of the scrape configuration
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Scrape configuration type: ServiceMonitor/PodMonitor
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Namespace selection method
	// Note: This field may return null, indicating that no valid values can be obtained.
	NamespaceSelector *string `json:"NamespaceSelector,omitnil,omitempty" name:"NamespaceSelector"`

	// Label selection method
	// Note: This field may return null, indicating that no valid values can be obtained.
	Selector *string `json:"Selector,omitnil,omitempty" name:"Selector"`

	// `Endpoints` information (PodMonitor does not have this parameter)
	Endpoints *string `json:"Endpoints,omitnil,omitempty" name:"Endpoints"`

	// Scrape configuration information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Yaml *string `json:"Yaml,omitnil,omitempty" name:"Yaml"`
}

// Predefined struct for user
type SetDefaultAlarmPolicyRequestParams struct {
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type SetDefaultAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Module name, which is fixed at "monitor"
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Alarm policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *SetDefaultAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDefaultAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDefaultAlarmPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetDefaultAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *SetDefaultAlarmPolicyResponseParams `json:"Response"`
}

func (r *SetDefaultAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPrometheusTempRequestParams struct {
	// Instance ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Sync target
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type SyncPrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Sync target
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *SyncPrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncPrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPrometheusTempResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncPrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *SyncPrometheusTempResponseParams `json:"Response"`
}

func (r *SyncPrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagInstance struct {
	// Tag key
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Number of instances
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceSum *int64 `json:"InstanceSum,omitnil,omitempty" name:"InstanceSum"`

	// Service type, for example, CVM
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Region ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Binding status. 2: bound; 1: binding
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BindingStatus *int64 `json:"BindingStatus,omitnil,omitempty" name:"BindingStatus"`

	// Tag status. 2: existent; 1: nonexistent
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagStatus *int64 `json:"TagStatus,omitnil,omitempty" name:"TagStatus"`
}

type Targets struct {
	// The total count
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of online targets
	// Note: This field may return null, indicating that no valid values can be obtained.
	Up *uint64 `json:"Up,omitnil,omitempty" name:"Up"`

	// Number of offline targets
	// Note: This field may return null, indicating that no valid values can be obtained.
	Down *uint64 `json:"Down,omitnil,omitempty" name:"Down"`

	// Number of unknown status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unknown *uint64 `json:"Unknown,omitnil,omitempty" name:"Unknown"`
}

type TaskStepInfo struct {
	// Step name
	Step *string `json:"Step,omitnil,omitempty" name:"Step"`

	// Lifecycle
	// `pending`
	// `running`
	// `success`
	// `failed`
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// Step start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// Step end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// If `LifeState` is `failed`, this field displays the error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailedMsg *string `json:"FailedMsg,omitnil,omitempty" name:"FailedMsg"`
}

type TemplateGroup struct {
	// Metric alarm rules.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Conditions []*Condition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// Event alarm rules.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EventConditions []*EventCondition `json:"EventConditions,omitnil,omitempty" name:"EventConditions"`

	// The associated alarm policy groups.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PolicyGroups []*PolicyGroup `json:"PolicyGroups,omitnil,omitempty" name:"PolicyGroups"`

	// Template-based policy group ID.
	GroupID *int64 `json:"GroupID,omitnil,omitempty" name:"GroupID"`

	// Template-based policy group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// UIN of the last modifier.
	LastEditUin *int64 `json:"LastEditUin,omitnil,omitempty" name:"LastEditUin"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Update time.
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// View.
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// Whether the logical relationship between rules is AND.
	IsUnionRule *int64 `json:"IsUnionRule,omitnil,omitempty" name:"IsUnionRule"`
}

// Predefined struct for user
type TerminatePrometheusInstancesRequestParams struct {
	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type TerminatePrometheusInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *TerminatePrometheusInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminatePrometheusInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminatePrometheusInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminatePrometheusInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminatePrometheusInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminatePrometheusInstancesResponseParams `json:"Response"`
}

func (r *TerminatePrometheusInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminatePrometheusInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Toleration struct {
	// Key of the taint to which the toleration is applied
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The key-value relationship
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// The taint effect to be matched
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`
}

type URLNotice struct {
	// Callback URL, which can contain up to 256 characters
	// Note: this field may return null, indicating that no valid values can be obtained.
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// Whether verification is passed. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsValid *int64 `json:"IsValid,omitnil,omitempty" name:"IsValid"`

	// Verification code
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidationCode *string `json:"ValidationCode,omitnil,omitempty" name:"ValidationCode"`

	// Start time of the notification in seconds, which is calculated from 00:00:00.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the notification in seconds, which is calculated from 00:00:00.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Notification cycle. The values 1-7 indicate Monday to Sunday.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Weekday []*int64 `json:"Weekday,omitnil,omitempty" name:"Weekday"`
}

// Predefined struct for user
type UnBindingAllPolicyObjectRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID. If `PolicyId` is used, this parameter will be ignored, and any value, e.g., `0`, can be passed in.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm policy ID. If this parameter is used, `GroupId` will be ignored.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`

	// Whether the event alert has been configured
	EbEventFlag *int64 `json:"EbEventFlag,omitnil,omitempty" name:"EbEventFlag"`
}

type UnBindingAllPolicyObjectRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID. If `PolicyId` is used, this parameter will be ignored, and any value, e.g., `0`, can be passed in.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alarm policy ID. If this parameter is used, `GroupId` will be ignored.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`

	// Whether the event alert has been configured
	EbEventFlag *int64 `json:"EbEventFlag,omitnil,omitempty" name:"EbEventFlag"`
}

func (r *UnBindingAllPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingAllPolicyObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "PolicyId")
	delete(f, "EbSubject")
	delete(f, "EbEventFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindingAllPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindingAllPolicyObjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnBindingAllPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *UnBindingAllPolicyObjectResponseParams `json:"Response"`
}

func (r *UnBindingAllPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingAllPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindingPolicyObjectRequestParams struct {
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID. If `PolicyId` is used, this parameter will be ignored, and any value, e.g., `0`, can be passed in.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// List of unique IDs of the object instances to be deleted. `UniqueId` can be obtained from the output parameter `List` of the [DescribeBindingPolicyObjectList](https://intl.cloud.tencent.com/document/api/248/40570?from_cn_redirect=1) API
	UniqueId []*string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`

	// Instance group ID. The `UniqueId` parameter is invalid if object instances are deleted by instance group.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Alarm policy ID. If this parameter is used, `GroupId` will be ignored.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`

	// Whether the event alert has been configured
	EbEventFlag *int64 `json:"EbEventFlag,omitnil,omitempty" name:"EbEventFlag"`
}

type UnBindingPolicyObjectRequest struct {
	*tchttp.BaseRequest
	
	// The value is fixed to monitor.
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// Policy group ID. If `PolicyId` is used, this parameter will be ignored, and any value, e.g., `0`, can be passed in.
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// List of unique IDs of the object instances to be deleted. `UniqueId` can be obtained from the output parameter `List` of the [DescribeBindingPolicyObjectList](https://intl.cloud.tencent.com/document/api/248/40570?from_cn_redirect=1) API
	UniqueId []*string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`

	// Instance group ID. The `UniqueId` parameter is invalid if object instances are deleted by instance group.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Alarm policy ID. If this parameter is used, `GroupId` will be ignored.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The alert configured for an event
	EbSubject *string `json:"EbSubject,omitnil,omitempty" name:"EbSubject"`

	// Whether the event alert has been configured
	EbEventFlag *int64 `json:"EbEventFlag,omitnil,omitempty" name:"EbEventFlag"`
}

func (r *UnBindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingPolicyObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "UniqueId")
	delete(f, "InstanceGroupId")
	delete(f, "PolicyId")
	delete(f, "EbSubject")
	delete(f, "EbEventFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindingPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindingPolicyObjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnBindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *UnBindingPolicyObjectResponseParams `json:"Response"`
}

func (r *UnBindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindPrometheusManagedGrafanaRequestParams struct {
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Grafana instance ID
	GrafanaId *string `json:"GrafanaId,omitnil,omitempty" name:"GrafanaId"`
}

type UnbindPrometheusManagedGrafanaRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Grafana instance ID
	GrafanaId *string `json:"GrafanaId,omitnil,omitempty" name:"GrafanaId"`
}

func (r *UnbindPrometheusManagedGrafanaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindPrometheusManagedGrafanaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GrafanaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindPrometheusManagedGrafanaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindPrometheusManagedGrafanaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindPrometheusManagedGrafanaResponse struct {
	*tchttp.BaseResponse
	Response *UnbindPrometheusManagedGrafanaResponseParams `json:"Response"`
}

func (r *UnbindPrometheusManagedGrafanaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindPrometheusManagedGrafanaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaDashboardRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus integration code, indicating to delete the corresponding dashboard. Valid values:
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitnil,omitempty" name:"IntegrationCodes"`
}

type UninstallGrafanaDashboardRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus integration code, indicating to delete the corresponding dashboard. Valid values:
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitnil,omitempty" name:"IntegrationCodes"`
}

func (r *UninstallGrafanaDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallGrafanaDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaDashboardResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UninstallGrafanaDashboardResponse struct {
	*tchttp.BaseResponse
	Response *UninstallGrafanaDashboardResponseParams `json:"Response"`
}

func (r *UninstallGrafanaDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaPluginsRequestParams struct {
	// Array of plugin IDs, such as "PluginIds": [ "grafana-clock-panel" ]. The plugin ID can be obtained through the `DescribePluginOverviews` API.
	PluginIds []*string `json:"PluginIds,omitnil,omitempty" name:"PluginIds"`

	// TCMG instance ID, such as “grafana-abcdefg”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type UninstallGrafanaPluginsRequest struct {
	*tchttp.BaseRequest
	
	// Array of plugin IDs, such as "PluginIds": [ "grafana-clock-panel" ]. The plugin ID can be obtained through the `DescribePluginOverviews` API.
	PluginIds []*string `json:"PluginIds,omitnil,omitempty" name:"PluginIds"`

	// TCMG instance ID, such as “grafana-abcdefg”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *UninstallGrafanaPluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallGrafanaPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaPluginsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UninstallGrafanaPluginsResponse struct {
	*tchttp.BaseResponse
	Response *UninstallGrafanaPluginsResponseParams `json:"Response"`
}

func (r *UninstallGrafanaPluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertRuleRequestParams struct {
	// Prometheus alerting rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Alerting rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Alerting rule expression
	Expr *string `json:"Expr,omitnil,omitempty" name:"Expr"`

	// Alerting rule duration
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// List of alerting rule recipient groups
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// List of alerting rule tags
	Labels []*PrometheusRuleKV `json:"Labels,omitnil,omitempty" name:"Labels"`

	// List of alerting rule annotations.
	// 
	// Alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively.
	Annotations []*PrometheusRuleKV `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// Alerting rule template category
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type UpdateAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus alerting rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`

	// Alerting rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Alerting rule expression
	Expr *string `json:"Expr,omitnil,omitempty" name:"Expr"`

	// Alerting rule duration
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// List of alerting rule recipient groups
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// List of alerting rule tags
	Labels []*PrometheusRuleKV `json:"Labels,omitnil,omitempty" name:"Labels"`

	// List of alerting rule annotations.
	// 
	// Alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively.
	Annotations []*PrometheusRuleKV `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// Alerting rule template category
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *UpdateAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "InstanceId")
	delete(f, "RuleState")
	delete(f, "RuleName")
	delete(f, "Expr")
	delete(f, "Duration")
	delete(f, "Receivers")
	delete(f, "Labels")
	delete(f, "Annotations")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAlertRuleResponseParams `json:"Response"`
}

func (r *UpdateAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertRuleStateRequestParams struct {
	// List of rule IDs
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule status code. Valid values:
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

type UpdateAlertRuleStateRequest struct {
	*tchttp.BaseRequest
	
	// List of rule IDs
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rule status code. Valid values:
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

func (r *UpdateAlertRuleStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	delete(f, "InstanceId")
	delete(f, "RuleState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAlertRuleStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertRuleStateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAlertRuleStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAlertRuleStateResponseParams `json:"Response"`
}

func (r *UpdateAlertRuleStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDNSConfigRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of DNS servers
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`
}

type UpdateDNSConfigRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of DNS servers
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`
}

func (r *UpdateDNSConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDNSConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NameServers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDNSConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDNSConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDNSConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDNSConfigResponseParams `json:"Response"`
}

func (r *UpdateDNSConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDNSConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateExporterIntegrationRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Configuration content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type UpdateExporterIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kubernetes cluster type. Valid values:
	// <li> 1 = TKE </li>
	// <li> 2 = EKS </li>
	// <li> 3 = MEKS </li>
	KubeType *int64 `json:"KubeType,omitnil,omitempty" name:"KubeType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Type
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Configuration content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *UpdateExporterIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateExporterIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	delete(f, "Kind")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateExporterIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateExporterIntegrationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateExporterIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateExporterIntegrationResponseParams `json:"Response"`
}

func (r *UpdateExporterIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateExporterIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaConfigRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// JSON-encoded string
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type UpdateGrafanaConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// JSON-encoded string
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *UpdateGrafanaConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateGrafanaConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaConfigResponseParams `json:"Response"`
}

func (r *UpdateGrafanaConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaEnvironmentsRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Environment variable string
	Envs *string `json:"Envs,omitnil,omitempty" name:"Envs"`
}

type UpdateGrafanaEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Environment variable string
	Envs *string `json:"Envs,omitnil,omitempty" name:"Envs"`
}

func (r *UpdateGrafanaEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Envs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaEnvironmentsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateGrafanaEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaEnvironmentsResponseParams `json:"Response"`
}

func (r *UpdateGrafanaEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaIntegrationRequestParams struct {
	// Integration ID, such as “integration-abcd1234”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`

	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration type, such as “tencent-cloud-prometheus”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Integration content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type UpdateGrafanaIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// Integration ID, such as “integration-abcd1234”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	IntegrationId *string `json:"IntegrationId,omitnil,omitempty" name:"IntegrationId"`

	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Integration type, such as “tencent-cloud-prometheus”. You can view it by going to the instance details page and clicking **Tencent Cloud Service Integration** > **Integration List**.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Integration content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *UpdateGrafanaIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntegrationId")
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaIntegrationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateGrafanaIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaIntegrationResponseParams `json:"Response"`
}

func (r *UpdateGrafanaIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaNotificationChannelRequestParams struct {
	// Channel ID, such as “nchannel-abcd1234”.
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Array of notification channel IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// This parameter has been deprecated. Please use `OrganizationIds` instead.
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitnil,omitempty" name:"ExtraOrgIds"`

	// Array of valid organization IDs
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

type UpdateGrafanaNotificationChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID, such as “nchannel-abcd1234”.
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert channel name, such as “test”.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Array of notification channel IDs
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// This parameter has been deprecated. Please use `OrganizationIds` instead.
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitnil,omitempty" name:"ExtraOrgIds"`

	// Array of valid organization IDs
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

func (r *UpdateGrafanaNotificationChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaNotificationChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "InstanceId")
	delete(f, "ChannelName")
	delete(f, "Receivers")
	delete(f, "ExtraOrgIds")
	delete(f, "OrganizationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaNotificationChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaNotificationChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateGrafanaNotificationChannelResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaNotificationChannelResponseParams `json:"Response"`
}

func (r *UpdateGrafanaNotificationChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaNotificationChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaWhiteListRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of public IPs (such as “127.0.0.1”) in the allowlist, which can be viewed through the `DescribeGrafanaWhiteList` API.
	Whitelist []*string `json:"Whitelist,omitnil,omitempty" name:"Whitelist"`
}

type UpdateGrafanaWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Array of public IPs (such as “127.0.0.1”) in the allowlist, which can be viewed through the `DescribeGrafanaWhiteList` API.
	Whitelist []*string `json:"Whitelist,omitnil,omitempty" name:"Whitelist"`
}

func (r *UpdateGrafanaWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Whitelist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaWhiteListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateGrafanaWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaWhiteListResponseParams `json:"Response"`
}

func (r *UpdateGrafanaWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAgentStatusRequestParams struct {
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of agent IDs such as “agent-abcd1234”, which can be obtained on the **Agent Management** page in the console.
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`

	// Status to update
	// <li> 1 = enabled </li>
	// <li> 2 = disabled </li>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdatePrometheusAgentStatusRequest struct {
	*tchttp.BaseRequest
	
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of agent IDs such as “agent-abcd1234”, which can be obtained on the **Agent Management** page in the console.
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`

	// Status to update
	// <li> 1 = enabled </li>
	// <li> 2 = disabled </li>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *UpdatePrometheusAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAgentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentIds")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrometheusAgentStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAgentStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePrometheusAgentStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePrometheusAgentStatusResponseParams `json:"Response"`
}

func (r *UpdatePrometheusAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAlertGroupRequestParams struct {
	// prometheus instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert group ID, such as alert-xxxx.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alert group name. not allowed to have the same name as other alert groups.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Alert group status.
	// 2 - enable.
	// 3 - disabled.
	// Specifies the Alert rule status under the `Rules` field to be overwritten when not empty.
	GroupState *int64 `json:"GroupState,omitnil,omitempty" name:"GroupState"`

	// Alert notification template ID list of tencent cloud observability platform, such as Consumer-xxxx or notice-xxxx.
	AMPReceivers []*string `json:"AMPReceivers,omitnil,omitempty" name:"AMPReceivers"`

	// Specifies the custom Alert notification template.
	CustomReceiver *PrometheusAlertCustomReceiver `json:"CustomReceiver,omitnil,omitempty" name:"CustomReceiver"`

	// Alert notification cycle (convergence time). defaults to 1h if left empty.
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// Specifies the Alert rule list to be created.
	Rules []*PrometheusAlertGroupRuleSet `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type UpdatePrometheusAlertGroupRequest struct {
	*tchttp.BaseRequest
	
	// prometheus instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alert group ID, such as alert-xxxx.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Alert group name. not allowed to have the same name as other alert groups.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Alert group status.
	// 2 - enable.
	// 3 - disabled.
	// Specifies the Alert rule status under the `Rules` field to be overwritten when not empty.
	GroupState *int64 `json:"GroupState,omitnil,omitempty" name:"GroupState"`

	// Alert notification template ID list of tencent cloud observability platform, such as Consumer-xxxx or notice-xxxx.
	AMPReceivers []*string `json:"AMPReceivers,omitnil,omitempty" name:"AMPReceivers"`

	// Specifies the custom Alert notification template.
	CustomReceiver *PrometheusAlertCustomReceiver `json:"CustomReceiver,omitnil,omitempty" name:"CustomReceiver"`

	// Alert notification cycle (convergence time). defaults to 1h if left empty.
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// Specifies the Alert rule list to be created.
	Rules []*PrometheusAlertGroupRuleSet `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *UpdatePrometheusAlertGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAlertGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "GroupState")
	delete(f, "AMPReceivers")
	delete(f, "CustomReceiver")
	delete(f, "RepeatInterval")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrometheusAlertGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAlertGroupResponseParams struct {
	// Updated Alert group ID. must match the regular expression `alert-[a-z0-9]{8}`.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePrometheusAlertGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePrometheusAlertGroupResponseParams `json:"Response"`
}

func (r *UpdatePrometheusAlertGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAlertGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAlertGroupStateRequestParams struct {
	// TMP instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alarm group ID list, such as alert-xxxx.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Describes the group status of the alarm.
	// 2 - enable.
	// 3 - disabled.
	GroupState *int64 `json:"GroupState,omitnil,omitempty" name:"GroupState"`
}

type UpdatePrometheusAlertGroupStateRequest struct {
	*tchttp.BaseRequest
	
	// TMP instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Alarm group ID list, such as alert-xxxx.
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Describes the group status of the alarm.
	// 2 - enable.
	// 3 - disabled.
	GroupState *int64 `json:"GroupState,omitnil,omitempty" name:"GroupState"`
}

func (r *UpdatePrometheusAlertGroupStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAlertGroupStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupIds")
	delete(f, "GroupState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrometheusAlertGroupStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAlertGroupStateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePrometheusAlertGroupStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePrometheusAlertGroupStateResponseParams `json:"Response"`
}

func (r *UpdatePrometheusAlertGroupStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAlertGroupStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusScrapeJobRequestParams struct {
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID such as “agent-abcd1234”, which can be obtained on the **Agent Management** page in the console
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Scrape task ID such as “job-abcd1234”. You can go to the **Agent Management** page and obtain the ID in the pop-up **Create Scrape Task** window.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Scrape task ID in the format of “job_name:xx”
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type UpdatePrometheusScrapeJobRequest struct {
	*tchttp.BaseRequest
	
	// TMP instance ID, such as “prom-abcd1234”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent ID such as “agent-abcd1234”, which can be obtained on the **Agent Management** page in the console
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Scrape task ID such as “job-abcd1234”. You can go to the **Agent Management** page and obtain the ID in the pop-up **Create Scrape Task** window.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Scrape task ID in the format of “job_name:xx”
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *UpdatePrometheusScrapeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusScrapeJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "JobId")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrometheusScrapeJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusScrapeJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePrometheusScrapeJobResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePrometheusScrapeJobResponseParams `json:"Response"`
}

func (r *UpdatePrometheusScrapeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusScrapeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordingRuleRequestParams struct {
	// Recording rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Recording rule group content, which is in YAML format and is Base64-encoded.
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus recording rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

type UpdateRecordingRuleRequest struct {
	*tchttp.BaseRequest
	
	// Recording rule name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Recording rule group content, which is in YAML format and is Base64-encoded.
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Prometheus instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus recording rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule status code. Valid values:
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// Default value: 2 (enabled).
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

func (r *UpdateRecordingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Group")
	delete(f, "InstanceId")
	delete(f, "RuleId")
	delete(f, "RuleState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordingRuleResponseParams struct {
	// Rule ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRecordingRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordingRuleResponseParams `json:"Response"`
}

func (r *UpdateRecordingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSSOAccountRequestParams struct {
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User account ID, such as “10000000”.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Permission
	Role []*GrafanaAccountRole `json:"Role,omitnil,omitempty" name:"Role"`

	// Remarks
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
}

type UpdateSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-abcdefgh”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User account ID, such as “10000000”.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Permission
	Role []*GrafanaAccountRole `json:"Role,omitnil,omitempty" name:"Role"`

	// Remarks
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
}

func (r *UpdateSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	delete(f, "Role")
	delete(f, "Notes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSSOAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSSOAccountResponseParams `json:"Response"`
}

func (r *UpdateSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGrafanaDashboardRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus integration code, indicating to upgrade to the corresponding dashboard. Valid values:
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitnil,omitempty" name:"IntegrationCodes"`
}

type UpgradeGrafanaDashboardRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Prometheus integration code, indicating to upgrade to the corresponding dashboard. Valid values:
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitnil,omitempty" name:"IntegrationCodes"`
}

func (r *UpgradeGrafanaDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeGrafanaDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGrafanaDashboardResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeGrafanaDashboardResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeGrafanaDashboardResponseParams `json:"Response"`
}

func (r *UpgradeGrafanaDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGrafanaInstanceRequestParams struct {
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Version alias, such as v7.4.2.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type UpgradeGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// TCMG instance ID, such as “grafana-12345678”.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Version alias, such as v7.4.2.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

func (r *UpgradeGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGrafanaInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeGrafanaInstanceResponseParams `json:"Response"`
}

func (r *UpgradeGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserNotice struct {
	// Recipient type. Valid values: USER (user), GROUP (user group)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Notification start time, which is expressed by the number of seconds since 00:00:00. Value range: 0-86399
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Notification end time, which is expressed by the number of seconds since 00:00:00. Value range: 0-86399
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Notification channel list. Valid values: `EMAIL` (email), `SMS` (SMS), `CALL` (phone), `WECHAT` (WeChat), `RTX` (WeCom)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NoticeWay []*string `json:"NoticeWay,omitnil,omitempty" name:"NoticeWay"`

	// User `uid` list
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserIds []*int64 `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// User group ID list
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupIds []*int64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// Phone polling list
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneOrder []*int64 `json:"PhoneOrder,omitnil,omitempty" name:"PhoneOrder"`

	// Number of phone pollings. Value range: 1-5
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitnil,omitempty" name:"PhoneCircleTimes"`

	// Call interval in seconds within one polling. Value range: 60-900
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitnil,omitempty" name:"PhoneInnerInterval"`

	// Polling interval in seconds. Value range: 60-900
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitnil,omitempty" name:"PhoneCircleInterval"`

	// Whether receipt notification is required. Valid values: 0 (no), 1 (yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	NeedPhoneArriveNotice *int64 `json:"NeedPhoneArriveNotice,omitnil,omitempty" name:"NeedPhoneArriveNotice"`

	// Dial type. `SYNC` (simultaneous dial), `CIRCLE` (polled dial). Default value: `CIRCLE`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PhoneCallType *string `json:"PhoneCallType,omitnil,omitempty" name:"PhoneCallType"`

	// Notification cycle. The values 1-7 indicate Monday to Sunday.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Weekday []*int64 `json:"Weekday,omitnil,omitempty" name:"Weekday"`

	// List of schedule IDs
	// Note: This field may return null, indicating that no valid values can be obtained.
	OnCallFormIDs []*string `json:"OnCallFormIDs,omitnil,omitempty" name:"OnCallFormIDs"`
}