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

type BindingPolicyObjectDimension struct {

	// Region name.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region ID.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Dimensions.
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// Event dimensions.
	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type BindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Required. It is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Instance group ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Dimensions of an object to be bound.
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
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

type CreatePolicyGroupCondition struct {

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to. This parameter may not be set if a default comparative type is set for a metric.
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// Comparative value. This parameter may not be set if a metric has no requirement.
	CalcValue *float64 `json:"CalcValue,omitempty" name:"CalcValue"`

	// Data statistics period in seconds. This parameter may not be set if a metric has a default value.
	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// Number of consecutive periods after which an alarm will be triggered.
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// If a metric is created based on a template, the RuleId of the metric in the template must be input.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupEventCondition struct {

	// Alarm event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm will be triggered only once. The value >0 indicates that an alarm will be triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// If a metric is created based on a template, the RuleId of the metric in the template must be input.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// Policy group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Name of the view to which the policy group belongs. If the policy group is created based on a template, this parameter may not be set.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// ID of the project to which the policy group belongs, which will be used for authentication.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID of a template-based policy group. This parameter is required only when the policy group is created based on a template.
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`

	// Whether the policy group is shielded. The value 0 indicates that the policy group is not shielded. The value 1 indicates that the policy group is shielded. The default value is 0.
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// Remarks of the policy group.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Insertion time in the format of Unix timestamp. If you do not set this parameter, the background processing time is used.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// Alarm threshold rule in the policy group.
	Conditions []*CreatePolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// Event alarm rules in the policy group.
	EventConditions []*CreatePolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// Whether to invoke at the background. Only when the value is 1, the rules in the background pull policy template are filled into the Conditions and EventConditions fields.
	BackEndCall *int64 `json:"BackEndCall,omitempty" name:"BackEndCall"`

	// The “AND” and “OR” rules for alarm metrics. The value 0 indicates “OR”, which means that an alarm will be reported when any rule is met. The value 1 indicates “AND”, which means that an alarm will be reported only when all rules are met.
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

	// API module name. The value for the current API is monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Start time, which is the timestamp one day earlier by default.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time, which is the current timestamp by default.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Number of parameters that can be returned on each page. Value range: 1 - 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Parameter offset on each page. The value starts from 0 and the default value is 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Sorting rule by UpdateTime. Valid values: asc and desc.
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// Sorting rule by OccurTime. Valid values: asc or desc. Sorting by UpdateTimeOrder takes a higher priority.
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// Filter by event type. The value 1 indicates service issues. The value 2 indicates other subscriptions.
	AccidentType []*int64 `json:"AccidentType,omitempty" name:"AccidentType" list`

	// Filter by event. The value 1 indicates CVM storage issues. The value 2 indicates CVM network connection issues. The value 3 indicates that the CVM runs exceptionally. The value 202 indicates that an ISP network jitter occurs.
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

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest

	// Service namespace. Different Tencent Cloud services have different namespaces. For more information on service namespaces, see the monitoring API documentation of each product. For example, you can see [CVM Monitoring APIs](https://cloud.tencent.com/document/api/248/30385) for the namespace of CVM.
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric name. Different Tencent Cloud services have different metric names. For more information on service metric names, see the monitoring API documentation of each product. For example, you can see the [CVM Monitoring APIs](https://cloud.tencent.com/document/api/248/30385) for the metric names of CVM.
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

	// Alarm status ID.
	// Note: This field may return null, indicating that no valid value was found.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Alarm status.
	// Note: This field may return null, indicating that no valid value was found.
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

	// API module name. The value for the current API is monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Start time, which is the timestamp one day ago by default.
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

	// 
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

		// Bound instance group information. You do not need to set this parameter if no instance group is bound.
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

	// Optional duration in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// Required or not.
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriod struct {

	// Default period in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// Optional period in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// Required or not.
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriodNum struct {

	// Number of default periods.
	// Note: This field may return null, indicating that no valid value was found.
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// Number of optional periods.
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

	// Data aggregation method in a period of 1 second.
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

	// Event type, which is a reserved field. At present, it is fixed to 2.
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

		// List of Alarm policy conditions.
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

	// Data statistics period in seconds.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Metric ID.
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// Threshold rule ID.
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Metric unit.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Alarm sending and converging type. The value 0 indicates that alarms are sent consecutively. The value 1 indicates that alarms are sent exponentially.
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// Alarm sending period in seconds. The value <0 indicates that no alarm will be triggered. The value 0 indicates that an alarm is triggered only once. The value >0 indicates that an alarm is triggered at the interval of triggerTime.
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// Comparative type. The value 1 indicates greater than. The value 2 indicates greater than or equal to. The value 3 indicates smaller than. The value 4 indicates smaller than or equal to. The value 5 indicates equal to. The value 6 indicates not equal to. The value 7 indicates day-on-day increase. The value 8 indicates day-on-day decrease. The value 9 indicates week-on-week increase. The value 10 indicates week-on-week decrease. The value 11 indicates periodical increase. The value 12 indicates periodical decrease.
	// Note: This field may return null, indicating that no valid value was found.
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// Threshold.
	// Note: This field may return null, indicating that no valid value was found.
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// Duration at which an alarm will be triggered in seconds.
	// Note: This field may return null, indicating that no valid value was found.
	ContinueTime *int64 `json:"ContinueTime,omitempty" name:"ContinueTime"`
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

	// Whether the “AND” rule is used.
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

	// List of alarm receiver group IDs.
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList" list`

	// List of alarm recipient IDs.
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`

	// Start time of the alarm period. Value range: [0,86400). Convert the Unix timestamp to Beijing time and then remove the date. For example, 7200 indicates “10:0:0”.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the alarm period. The meaning is the same as that of StartTime.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Recipient type. Valid values: group and user.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// Alarm notification type. Valid values: "SMS", "SITE", "EMAIL", "CALL", and "WECHAT".
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// Uid of the alarm call receiver.
	// Note: This field may return null, indicating that no valid value was found.
	UidList []*int64 `json:"UidList,omitempty" name:"UidList" list`

	// Number of alarm call rounds.
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// Round interval of alarm calls in seconds.
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// Person interval of alarm calls in seconds.
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// Whether to send an alarm call delivery notice. The value 0 indicates that no notice needs to be sent. The value 1 indicates that a notice needs to be sent.
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// Alarm call notification time. Valid values: OCCUR (indicating that a notice is sent when the alarm is reported) and RECOVER (indicating that a notice is sent when the alarm is cleared).
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor" list`

	// Notification method when an alarm is cleared. Valid value: SMS.
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

		// Last update time.
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// Regions that support this policy.
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

		// Whether the policy can be set as the default policy.
		CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

		// Whether the “AND” rule is used.
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

	// Uin that is last edited.
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// Last update time.
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time.
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// Number of instances that are bound to the policy group.
	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`

	// Number of unshielded instances that are bound to the policy group.
	NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

	// Whether it is the default policy. The value 0 indicates that it is not the default policy. The value 1 indicates that it is the default policy.
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// Whether the policy can be set as the default policy.
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

	// The “AND” or “OR” rule. The value 0 indicates the “OR” rule (indicating that an alarm will be reported if any rule reaches the threshold condition). The value 1 indicates the “AND” rule (indicating that an alarm will be reported when all rules reach the threshold conditions).
	// Note: This field may return null, indicating that no valid value was found.
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupListGroupInstanceGroup struct {

	// Instance group name ID.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// Policy type view name.
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// Uin that is last edited.
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

	// Filter by recipient or recipient group. The value “user” indicates by recipient. The value “group” indicates by recipient group.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
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

	// Chinese event name.
	// Note: This field may return null, indicating that no valid value was found.
	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`

	// English event name.
	// Note: This field may return null, indicating that no valid value was found.
	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`

	// Event name abbreviation.
	// Note: This field may return null, indicating that no valid value was found.
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Chinese product name.
	// Note: This field may return null, indicating that no valid value was found.
	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`

	// English product name.
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
}

type DescribeProductEventListEventsDimensions struct {

	// English dimension name.
	// Note: This field may return null, indicating that no valid value was found.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Chinese dimension name.
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

	// Number of exceptional events.
	// Note: This field may return null, indicating that no valid value was found.
	UnNormalEventAmount *int64 `json:"UnNormalEventAmount,omitempty" name:"UnNormalEventAmount"`

	// Number of events that have not been recovered.
	// Note: This field may return null, indicating that no valid value was found.
	UnRecoverAmount *int64 `json:"UnRecoverAmount,omitempty" name:"UnRecoverAmount"`
}

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest

	// API module name. It is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Filter by product type. For example, “cvm” indicates Cloud Virtual Machine.
	ProductName []*string `json:"ProductName,omitempty" name:"ProductName" list`

	// Filter by product name. For example, "guest_reboot" indicates server restart.
	EventName []*string `json:"EventName,omitempty" name:"EventName" list`

	// Affected object, such as ins-19708ino.
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId" list`

	// Filter by dimension, such as by public IP: 10.0.0.1.
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions" list`

	// Filter by region, such as by gz.
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList" list`

	// Filter by event type. Valid values: ["status_change","abnormal"], which indicate events whose statuses have changed and exceptional events respectively.
	Type []*string `json:"Type,omitempty" name:"Type" list`

	// Filter by event status. Valid values: ["recover","alarm","-"], which indicate that an event has been recovered, has not been recovered, and has no status respectively.
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// Filter by project ID.
	Project []*string `json:"Project,omitempty" name:"Project" list`

	// Filter by alarm status configuration. The value 1 indicates that the alarm status has been configured. The value 0 indicates that the alarm status has not been configured.
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`

	// Sorting by update time. The value ASC indicates the ascending order. The value DESC indicates the descending order. The default value is DESC.
	TimeOrder *string `json:"TimeOrder,omitempty" name:"TimeOrder"`

	// Start time, which is the timestamp one day ago by default.
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

	// Namespace. Each Tencent Cloud product has a namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric name. For detailed metric descriptions of each Tencent Cloud product, see the corresponding [Monitoring API](https://cloud.tencent.com/document/product/248/30384) document
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Combination of instance object dimensions
	Instances []*Instance `json:"Instances,omitempty" name:"Instances" list`

	// Monitoring statistical period. The default value is 300, and the unit is s
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Start time such as 2018-09-22T19:51:23+08:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time. Uses the current time by default and cannot be earlier than StartTime
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

type ModifyAlarmReceiversRequest struct {
	*tchttp.BaseRequest

	// ID of a policy group whose recipient needs to be modified.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Required. The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// New recipient information. If this parameter is not set, all recipients will be deleted.
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

type PeriodsSt struct {

	// Period
	Period *string `json:"Period,omitempty" name:"Period"`

	// Statistical method
	StatType []*string `json:"StatType,omitempty" name:"StatType" list`
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

	// Start time of the alarm period. Value range: [0,86400). Convert the Unix timestamp to Beijing time and then remove the date. For example, 7200 indicates “10:0:0”.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the alarm period. The meaning is the same as that of StartTime.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Alarm notification type. Valid values: "SMS", "SITE", "EMAIL", "CALL", and "WECHAT".
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// Recipient type. Valid values: group and user.
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Call alarm notification time. Valid values: OCCUR (indicating that a notice is sent when the alarm is reported) and RECOVER (indicating that a notice is sent when the alarm is cleared).
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor" list`

	// Uid of the alarm call receiver.
	UidList []*int64 `json:"UidList,omitempty" name:"UidList" list`

	// Number of alarm call rounds.
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// Person interval of alarm calls in seconds.
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// Round interval of alarm calls in seconds.
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// Notification method when an alarm is cleared. Valid value: SMS.
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

	// API module name. The value for the current API is monitor.
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

type UnBindingAllPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// The value is fixed to monitor.
	Module *string `json:"Module,omitempty" name:"Module"`

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
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

	// Policy group ID.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// List of unique IDs of object instances to be deleted.
	UniqueId []*string `json:"UniqueId,omitempty" name:"UniqueId" list`

	// Instance group ID. The UniqueId parameter is invalid if object instances are deleted by instance group.
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
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
