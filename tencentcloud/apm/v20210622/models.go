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

package v20210622

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type APMKV struct {
	// Key value definition.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value definition.
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type APMKVItem struct {
	// Key value definition.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value definition.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ApmAgentInfo struct {
	// Agent download address.
	AgentDownloadURL *string `json:"AgentDownloadURL,omitnil,omitempty" name:"AgentDownloadURL"`

	// Collector reporting address.
	CollectorURL *string `json:"CollectorURL,omitnil,omitempty" name:"CollectorURL"`

	// Token information.
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// Public network reporting address.
	PublicCollectorURL *string `json:"PublicCollectorURL,omitnil,omitempty" name:"PublicCollectorURL"`

	// Self-Developed vpc report address.
	InnerCollectorURL *string `json:"InnerCollectorURL,omitnil,omitempty" name:"InnerCollectorURL"`

	// Private link reporting address.
	PrivateLinkCollectorURL *string `json:"PrivateLinkCollectorURL,omitnil,omitempty" name:"PrivateLinkCollectorURL"`
}

type ApmApplicationConfigView struct {
	// Business system id.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// Application name	.	
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// API filtering.
	OperationNameFilter *string `json:"OperationNameFilter,omitnil,omitempty" name:"OperationNameFilter"`

	// Error type filtering.
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// HTTP status code filtering.
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// Application diagnosis switch (deprecated).
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// URL convergence switch. 0: off; 1: on.
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// URL convergence threshold.	
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// URL convergence rule in the form of a regular expression.	
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// URL exclusion rule in the form of a regular expression.
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// Log feature switch. 0: off; 1: on.
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// Log source.	
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// Log set. 
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// Log topic.
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// Method stack snapshot switch: true to enable, false to disable.
	SnapshotEnable *bool `json:"SnapshotEnable,omitnil,omitempty" name:"SnapshotEnable"`

	// Slow call listening trigger threshold.
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// Probe master switch.
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// Component list switch (deprecated).
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// Link compression switch (deprecated).
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`
}

type ApmField struct {
	// Metric name.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Indicator numerical value.
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`

	// Units corresponding to the metric.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Year-Over-Year result array, recommended to use.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CompareVals []*APMKVItem `json:"CompareVals,omitnil,omitempty" name:"CompareVals"`

	// Indicator numerical value of the previous period in year-over-year comparison.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastPeriodValue []*APMKV `json:"LastPeriodValue,omitnil,omitempty" name:"LastPeriodValue"`

	// Year-On-Year metric value. deprecated, not recommended for use.
	CompareVal *string `json:"CompareVal,omitnil,omitempty" name:"CompareVal"`
}

type ApmInstanceDetail struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Business system name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Business system description information.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Status of the business system.
	// {Initializing; running; throttling}.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Region where the business system belongs.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Business system tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// AppID information.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Creator uin.
	CreateUin *string `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// Storage used (unit: mb).
	AmountOfUsedStorage *float64 `json:"AmountOfUsedStorage,omitnil,omitempty" name:"AmountOfUsedStorage"`

	// Quantity of server applications of the business system.
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// Average daily reported span count.
	CountOfReportSpanPerDay *int64 `json:"CountOfReportSpanPerDay,omitnil,omitempty" name:"CountOfReportSpanPerDay"`

	// Retention period of trace data (unit: days).
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// Business system report limit.
	SpanDailyCounters *int64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// Whether the business system billing is Activated (0 = not activated, 1 = activated).
	BillingInstance *int64 `json:"BillingInstance,omitnil,omitempty" name:"BillingInstance"`

	// Error warning line (unit: %).
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// Sampling rate (unit: %).
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Error sampling switch (0: off, 1: on).
	ErrorSample *int64 `json:"ErrorSample,omitnil,omitempty" name:"ErrorSample"`

	// Sampling slow call saving threshold (unit: ms).
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil,omitempty" name:"SlowRequestSavedThreshold"`

	// CLS log region.
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// Log source.
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// Log feature switch (0: off; 1: on).
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// Log topic id.
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// Quantity of client applications of the business system.
	ClientCount *int64 `json:"ClientCount,omitnil,omitempty" name:"ClientCount"`

	// The quantity of active applications in this business system in the last two days.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// CLS log set.
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// Retention period of metric data (unit: days).
	MetricDuration *int64 `json:"MetricDuration,omitnil,omitempty" name:"MetricDuration"`

	// List of custom display tags.
	CustomShowTags []*string `json:"CustomShowTags,omitnil,omitempty" name:"CustomShowTags"`

	// Business system billing mode (1: prepaid, 0: pay-as-you-go).
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Indicates whether the billing mode of the business system takes effect.
	PayModeEffective *bool `json:"PayModeEffective,omitnil,omitempty" name:"PayModeEffective"`

	// Response time warning line (unit: ms).
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// Whether it is free (0 = no; 1 = limited free; 2 = completely free), default 0.
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`

	// Indicates whether it is the default business system of tsf (0 = no, 1 = yes).
	DefaultTSF *int64 `json:"DefaultTSF,omitnil,omitempty" name:"DefaultTSF"`

	// Whether to associate the dashboard (0 = off, 1 = on).
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// Associated dashboard id.
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// Whether to enable component vulnerability detection (0 = no, 1 = yes).
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// Whether to enable sql injection analysis (0: off, 1: on).
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// Reasons for traffic throttling.
	// Official version quota;.
	// Trial version quota.
	// Trial version expiration;.
	// Account in arrears.
	// }.
	StopReason *int64 `json:"StopReason,omitnil,omitempty" name:"StopReason"`
}

type ApmMetricRecord struct {
	// Field array, used for the query result of indicators.
	Fields []*ApmField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Tag array, used to distinguish the objects of groupby.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ApmTag struct {
	// Dimension key (column name, Tag key).
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Dimension value (tag value).
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type CreateApmInstanceRequestParams struct {
	// Business system name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Business system description information.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Retention period of trace data (unit: days, the default storage duration is 3 days).
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// Business system tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The report quota value of the business system. the default value is 0, indicating no limit on the report quota. (obsolete).
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// Billing model of the business system (0: pay-as-you-go, 1: prepaid).
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Whether it is a free edition business system (0 = paid edition; 1 = tsf restricted free edition; 2 = free edition).
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`
}

type CreateApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Business system name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Business system description information.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Retention period of trace data (unit: days, the default storage duration is 3 days).
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// Business system tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The report quota value of the business system. the default value is 0, indicating no limit on the report quota. (obsolete).
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// Billing model of the business system (0: pay-as-you-go, 1: prepaid).
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Whether it is a free edition business system (0 = paid edition; 1 = tsf restricted free edition; 2 = free edition).
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`
}

func (r *CreateApmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "TraceDuration")
	delete(f, "Tags")
	delete(f, "SpanDailyCounters")
	delete(f, "PayMode")
	delete(f, "Free")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmInstanceResponseParams struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateApmInstanceResponseParams `json:"Response"`
}

func (r *CreateApmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAgentRequestParams struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Access method: currently supports access and reporting via skywalking, ot, and ebpf methods. if not specified, ot is used by default.
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// Reporting environment: currently supports pl (private network reporting), public (public network), and inner (self-developed vpc) environment reporting. if not specified, the default is public.
	NetworkMode *string `json:"NetworkMode,omitnil,omitempty" name:"NetworkMode"`

	// Language reporting is now supported for java, golang, php, python, dotnet, nodejs. if not specified, golang is used by default.
	LanguageEnvironment *string `json:"LanguageEnvironment,omitnil,omitempty" name:"LanguageEnvironment"`

	// Reporting method, deprecated.
	ReportMethod *string `json:"ReportMethod,omitnil,omitempty" name:"ReportMethod"`
}

type DescribeApmAgentRequest struct {
	*tchttp.BaseRequest
	
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Access method: currently supports access and reporting via skywalking, ot, and ebpf methods. if not specified, ot is used by default.
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// Reporting environment: currently supports pl (private network reporting), public (public network), and inner (self-developed vpc) environment reporting. if not specified, the default is public.
	NetworkMode *string `json:"NetworkMode,omitnil,omitempty" name:"NetworkMode"`

	// Language reporting is now supported for java, golang, php, python, dotnet, nodejs. if not specified, golang is used by default.
	LanguageEnvironment *string `json:"LanguageEnvironment,omitnil,omitempty" name:"LanguageEnvironment"`

	// Reporting method, deprecated.
	ReportMethod *string `json:"ReportMethod,omitnil,omitempty" name:"ReportMethod"`
}

func (r *DescribeApmAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentType")
	delete(f, "NetworkMode")
	delete(f, "LanguageEnvironment")
	delete(f, "ReportMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAgentResponseParams struct {
	// Agent information.
	ApmAgent *ApmAgentInfo `json:"ApmAgent,omitnil,omitempty" name:"ApmAgent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmAgentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmAgentResponseParams `json:"Response"`
}

func (r *DescribeApmAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmInstancesRequestParams struct {
	// Tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Filter by business system name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Filter by business system id.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether to query the official demo business system (0 = non-demo business system, 1 = demo business system, default is 0).
	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitnil,omitempty" name:"DemoInstanceFlag"`

	// Whether to query all regional business systems (0 = do not query all regions, 1 = query all regions, default is 0).
	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitnil,omitempty" name:"AllRegionsFlag"`
}

type DescribeApmInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Filter by business system name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Filter by business system id.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether to query the official demo business system (0 = non-demo business system, 1 = demo business system, default is 0).
	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitnil,omitempty" name:"DemoInstanceFlag"`

	// Whether to query all regional business systems (0 = do not query all regions, 1 = query all regions, default is 0).
	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitnil,omitempty" name:"AllRegionsFlag"`
}

func (r *DescribeApmInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	delete(f, "InstanceName")
	delete(f, "InstanceIds")
	delete(f, "DemoInstanceFlag")
	delete(f, "AllRegionsFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmInstancesResponseParams struct {
	// APM business system list.
	Instances []*ApmInstanceDetail `json:"Instances,omitnil,omitempty" name:"Instances"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmInstancesResponseParams `json:"Response"`
}

func (r *DescribeApmInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralApmApplicationConfigRequestParams struct {
	// Application name.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeGeneralApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// Application name.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeGeneralApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralApmApplicationConfigResponseParams struct {
	// Application configuration item.
	ApmApplicationConfigView *ApmApplicationConfigView `json:"ApmApplicationConfigView,omitnil,omitempty" name:"ApmApplicationConfigView"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralApmApplicationConfigResponseParams `json:"Response"`
}

func (r *DescribeGeneralApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralMetricDataRequestParams struct {
	// Metric name to be queried, which cannot be customized. (for details, see <https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1>.).
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// View name. the input cannot be customized. [for details, see.](https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1).
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// The dimension information to be filtered; different views have corresponding metric dimensions. (for details, see <https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1>.).
	Filters []*GeneralFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Aggregated dimension; different views have corresponding metric dimensions. (for details, see <https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1>.).
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// The timestamp of the start time, supporting the query of metric data within 30 days. (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The timestamp of the end time, supporting the query of metric data within 30 days. (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to aggregate by a fixed time span: enter 1 for values of 1 and greater, and 0 if not filled in.
	// -If 0 is filled in, it calculates the metric data from the start time to the cutoff time.
	// - if 1 is filled in, the aggregation granularity will be selected according to the time span from the start time to the deadline:.
	//  -If the time span is (0,12) hours, it is aggregated by one-minute granularity.
	//  -If the time span is [12,48] hours, it is aggregated at a five-minute granularity.
	//  -If the time span is (48, +∞) hours, it is aggregated at an hourly granularity.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Sort query metrics.
	// Key: enter the tencentcloud api metric name. [for details, see](https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1) .
	// Value: specify the sorting method:.     
	// -Asc: sorts query metrics in ascending order.
	// - desc: sort query metrics in descending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Maximum number of queried metrics. currently, up to 50 data entries can be displayed. the value range for pagesize is 1-50. submit pagesize to show the limited number based on the value of pagesize.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeGeneralMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// Metric name to be queried, which cannot be customized. (for details, see <https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1>.).
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// View name. the input cannot be customized. [for details, see.](https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1).
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// The dimension information to be filtered; different views have corresponding metric dimensions. (for details, see <https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1>.).
	Filters []*GeneralFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Aggregated dimension; different views have corresponding metric dimensions. (for details, see <https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1>.).
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// The timestamp of the start time, supporting the query of metric data within 30 days. (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The timestamp of the end time, supporting the query of metric data within 30 days. (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to aggregate by a fixed time span: enter 1 for values of 1 and greater, and 0 if not filled in.
	// -If 0 is filled in, it calculates the metric data from the start time to the cutoff time.
	// - if 1 is filled in, the aggregation granularity will be selected according to the time span from the start time to the deadline:.
	//  -If the time span is (0,12) hours, it is aggregated by one-minute granularity.
	//  -If the time span is [12,48] hours, it is aggregated at a five-minute granularity.
	//  -If the time span is (48, +∞) hours, it is aggregated at an hourly granularity.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Sort query metrics.
	// Key: enter the tencentcloud api metric name. [for details, see](https://intl.cloud.tencent.com/document/product/248/101681?from_cn_redirect=1) .
	// Value: specify the sorting method:.     
	// -Asc: sorts query metrics in ascending order.
	// - desc: sort query metrics in descending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Maximum number of queried metrics. currently, up to 50 data entries can be displayed. the value range for pagesize is 1-50. submit pagesize to show the limited number based on the value of pagesize.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeGeneralMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Metrics")
	delete(f, "InstanceId")
	delete(f, "ViewName")
	delete(f, "Filters")
	delete(f, "GroupBy")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	delete(f, "OrderBy")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralMetricDataResponseParams struct {
	// Indicator result set.
	Records []*Line `json:"Records,omitnil,omitempty" name:"Records"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralMetricDataResponseParams `json:"Response"`
}

func (r *DescribeGeneralMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralOTSpanListRequestParams struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span query start timestamp (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span query end timestamp (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Universal filter parameters.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort
	// .
	// The currently supported keys are:.
	// 
	// -StartTime (start time).
	// -EndTime (end time).
	// -Duration (response time).
	// 
	// The currently supported values are:.
	// 
	// - desc: sort in descending order.
	// -Asc - ascending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The service name of the business itself. console users should fill in taw.
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Number of items per page, defaults to 10,000, valid value range is 0 – 10,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeGeneralOTSpanListRequest struct {
	*tchttp.BaseRequest
	
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span query start timestamp (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span query end timestamp (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Universal filter parameters.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort
	// .
	// The currently supported keys are:.
	// 
	// -StartTime (start time).
	// -EndTime (end time).
	// -Duration (response time).
	// 
	// The currently supported values are:.
	// 
	// - desc: sort in descending order.
	// -Asc - ascending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The service name of the business itself. console users should fill in taw.
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Number of items per page, defaults to 10,000, valid value range is 0 – 10,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeGeneralOTSpanListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralOTSpanListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "BusinessName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralOTSpanListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralOTSpanListResponseParams struct {
	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The trace structure containing the query results spans. the string after the opentelemetry standard trace structure is hashed. first, the trace is converted into a json string using ptrace.jsonmarshaler, then compressed with gzip, and finally converted into a base64 standard string.
	Spans *string `json:"Spans,omitnil,omitempty" name:"Spans"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralOTSpanListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralOTSpanListResponseParams `json:"Response"`
}

func (r *DescribeGeneralOTSpanListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralOTSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralSpanListRequestParams struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span query start timestamp (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span query end timestamp (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Universal filter parameters.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort
	// .
	// The currently supported keys are:.
	// 
	// -StartTime (start time).
	// -EndTime (end time).
	// -Duration (response time).
	// 
	// The currently supported values are:.
	// 
	// - desc: sort in descending order.
	// -Asc - ascending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The service name of the business itself. console users should fill in taw.
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Number of items per page, defaults to 10,000, valid values are 0 to 10,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeGeneralSpanListRequest struct {
	*tchttp.BaseRequest
	
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span query start timestamp (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span query end timestamp (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Universal filter parameters.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort
	// .
	// The currently supported keys are:.
	// 
	// -StartTime (start time).
	// -EndTime (end time).
	// -Duration (response time).
	// 
	// The currently supported values are:.
	// 
	// - desc: sort in descending order.
	// -Asc - ascending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The service name of the business itself. console users should fill in taw.
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Number of items per page, defaults to 10,000, valid values are 0 to 10,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeGeneralSpanListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralSpanListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "BusinessName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralSpanListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralSpanListResponseParams struct {
	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Span pagination list.
	Spans []*Span `json:"Spans,omitnil,omitempty" name:"Spans"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralSpanListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralSpanListResponseParams `json:"Response"`
}

func (r *DescribeGeneralSpanListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricRecordsRequestParams struct {
	// Metric list.
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time (unit: sec).
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or filter criteria.
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Sort
	// .
	// The currently supported keys are:.
	// 
	// -StartTime (start time).
	// -EndTime (end time).
	// -Duration (response time).
	// 
	// The currently supported values are:.
	// 
	// - desc: sort in descending order.
	// -Asc - ascending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The service name of the business itself. console users should fill in taw.
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Special handling of query results.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Size per page, defaults to 1,000, valid value range is 0 – 1,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Paging starting point.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page number.
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// Page length.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeMetricRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Metric list.
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time (unit: sec).
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or filter criteria.
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Sort
	// .
	// The currently supported keys are:.
	// 
	// -StartTime (start time).
	// -EndTime (end time).
	// -Duration (response time).
	// 
	// The currently supported values are:.
	// 
	// - desc: sort in descending order.
	// -Asc - ascending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// The service name of the business itself. console users should fill in taw.
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Special handling of query results.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Size per page, defaults to 1,000, valid value range is 0 – 1,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Paging starting point.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page number.
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// Page length.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeMetricRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Metrics")
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "OrFilters")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "BusinessName")
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMetricRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricRecordsResponseParams struct {
	// Indicator result set.
	Records []*ApmMetricRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// Number of metric query result sets.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMetricRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMetricRecordsResponseParams `json:"Response"`
}

func (r *DescribeMetricRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceOverviewRequestParams struct {
	// Metric list.
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Start time (unit: sec).
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting method
	// .
	// Value: fill in.
	// -Asc: sorts query metrics in ascending order.
	// - desc: sort query metrics in descending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Page size.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Paging starting point.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeServiceOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Metric list.
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Start time (unit: sec).
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting method
	// .
	// Value: fill in.
	// -Asc: sorts query metrics in ascending order.
	// - desc: sort query metrics in descending order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Page size.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Paging starting point.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeServiceOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Metrics")
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "GroupBy")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceOverviewResponseParams struct {
	// Indicator result set.
	Records []*ApmMetricRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceOverviewResponseParams `json:"Response"`
}

func (r *DescribeServiceOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesRequestParams struct {
	// Dimension name.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Start time (unit: sec).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Or filter criteria.
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// Usage type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest
	
	// Dimension name.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Start time (unit: sec).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Or filter criteria.
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// Usage type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeTagValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrFilters")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesResponseParams struct {
	// Dimension value list.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagValuesResponseParams `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Filtering method (=, !=, in).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter dimension name.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Filter value. uses commas to separate multiple values in in filtering method.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type GeneralFilter struct {
	// Filter dimension name.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Values after filtering.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Instrument struct {
	// Component name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Component switch.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type Line struct {
	// Metric name.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Metric chinese name.
	MetricNameCN *string `json:"MetricNameCN,omitnil,omitempty" name:"MetricNameCN"`

	// Time series.
	TimeSerial []*int64 `json:"TimeSerial,omitnil,omitempty" name:"TimeSerial"`

	// Data sequence.
	DataSerial []*float64 `json:"DataSerial,omitnil,omitempty" name:"DataSerial"`

	// Dimension list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type ModifyApmInstanceRequestParams struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Business system name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Business system description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Retention period of trace data (unit: days).
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// Billing switch.
	OpenBilling *bool `json:"OpenBilling,omitnil,omitempty" name:"OpenBilling"`

	// Business system report limit.
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// Error rate warning line. when the average error rate of the application exceeds this threshold, the system will give an abnormal note.
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// Sampling rate (unit: %).
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Error sampling switch (0: off, 1: on).
	ErrorSample *int64 `json:"ErrorSample,omitnil,omitempty" name:"ErrorSample"`

	// Sampling slow call saving threshold (unit: ms).
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil,omitempty" name:"SlowRequestSavedThreshold"`

	// Log feature switch (0: off; 1: on).
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// Log region, which takes effect after the log feature is enabled.
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// CLS log topic id, which takes effect after the log feature is enabled.
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// Logset, which takes effect only after the log feature is enabled.
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// Log source, which takes effect only after the log feature is enabled.
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// List of custom display tags.
	CustomShowTags []*string `json:"CustomShowTags,omitnil,omitempty" name:"CustomShowTags"`

	// Modify billing mode (1: prepaid, 0: pay-as-you-go).
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Response time warning line.
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// Whether it is free (0 = paid edition; 1 = tsf restricted free edition; 2 = free edition), default 0.
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`

	// Whether to associate the dashboard (0 = off, 1 = on).
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// Associated dashboard id, which takes effect after the associated dashboard is enabled.
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// SQL injection detection switch (0: off, 1: on).
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// Whether to enable component vulnerability detection (0 = no, 1 = yes).
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`
}

type ModifyApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Business system name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Business system description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Retention period of trace data (unit: days).
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// Billing switch.
	OpenBilling *bool `json:"OpenBilling,omitnil,omitempty" name:"OpenBilling"`

	// Business system report limit.
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// Error rate warning line. when the average error rate of the application exceeds this threshold, the system will give an abnormal note.
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// Sampling rate (unit: %).
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Error sampling switch (0: off, 1: on).
	ErrorSample *int64 `json:"ErrorSample,omitnil,omitempty" name:"ErrorSample"`

	// Sampling slow call saving threshold (unit: ms).
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil,omitempty" name:"SlowRequestSavedThreshold"`

	// Log feature switch (0: off; 1: on).
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// Log region, which takes effect after the log feature is enabled.
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// CLS log topic id, which takes effect after the log feature is enabled.
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// Logset, which takes effect only after the log feature is enabled.
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// Log source, which takes effect only after the log feature is enabled.
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// List of custom display tags.
	CustomShowTags []*string `json:"CustomShowTags,omitnil,omitempty" name:"CustomShowTags"`

	// Modify billing mode (1: prepaid, 0: pay-as-you-go).
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Response time warning line.
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// Whether it is free (0 = paid edition; 1 = tsf restricted free edition; 2 = free edition), default 0.
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`

	// Whether to associate the dashboard (0 = off, 1 = on).
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// Associated dashboard id, which takes effect after the associated dashboard is enabled.
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// SQL injection detection switch (0: off, 1: on).
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// Whether to enable component vulnerability detection (0 = no, 1 = yes).
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`
}

func (r *ModifyApmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Tags")
	delete(f, "Description")
	delete(f, "TraceDuration")
	delete(f, "OpenBilling")
	delete(f, "SpanDailyCounters")
	delete(f, "ErrRateThreshold")
	delete(f, "SampleRate")
	delete(f, "ErrorSample")
	delete(f, "SlowRequestSavedThreshold")
	delete(f, "IsRelatedLog")
	delete(f, "LogRegion")
	delete(f, "LogTopicID")
	delete(f, "LogSet")
	delete(f, "LogSource")
	delete(f, "CustomShowTags")
	delete(f, "PayMode")
	delete(f, "ResponseDurationWarningThreshold")
	delete(f, "Free")
	delete(f, "IsRelatedDashboard")
	delete(f, "DashboardTopicID")
	delete(f, "IsSqlInjectionAnalysis")
	delete(f, "IsInstrumentationVulnerabilityScan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmInstanceResponseParams `json:"Response"`
}

func (r *ModifyApmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGeneralApmApplicationConfigRequestParams struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Fields to be modified. the key and value respectively specify the field name and field value.
	// .
	// For specific fields, please refer to.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Name of the application list that requires configuration modification.	
	ServiceNames []*string `json:"ServiceNames,omitnil,omitempty" name:"ServiceNames"`
}

type ModifyGeneralApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Fields to be modified. the key and value respectively specify the field name and field value.
	// .
	// For specific fields, please refer to.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Name of the application list that requires configuration modification.	
	ServiceNames []*string `json:"ServiceNames,omitnil,omitempty" name:"ServiceNames"`
}

func (r *ModifyGeneralApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGeneralApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Tags")
	delete(f, "ServiceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGeneralApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGeneralApmApplicationConfigResponseParams struct {
	// Description of the returned value.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGeneralApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGeneralApmApplicationConfigResponseParams `json:"Response"`
}

func (r *ModifyGeneralApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGeneralApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderBy struct {
	// Sort field (starttime, endtime, duration are supported).
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// ASC: sequential sorting / desc: reverse sorting.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QueryMetricItem struct {
	// Metric name.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Year-Over-Year comparison is now supported for comparebyyesterday (compared to yesterday) and comparebylastweek (compared to last week). 
	Compares []*string `json:"Compares,omitnil,omitempty" name:"Compares"`

	// Year-On-Year, deprecated, not recommended for use.
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`
}

type Span struct {
	// Trace ID
	TraceID *string `json:"TraceID,omitnil,omitempty" name:"TraceID"`

	// Log.
	Logs []*SpanLog `json:"Logs,omitnil,omitempty" name:"Logs"`

	// Tag.
	Tags []*SpanTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Submit application service information.
	Process *SpanProcess `json:"Process,omitnil,omitempty" name:"Process"`

	// Generated timestamp (ms).
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Span name.
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// Association relationship.
	References []*SpanReference `json:"References,omitnil,omitempty" name:"References"`

	// Generated timestamp (ms).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Duration (ms).
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Span ID
	SpanID *string `json:"SpanID,omitnil,omitempty" name:"SpanID"`

	// Generated timestamp (ms).
	StartTimeMillis *int64 `json:"StartTimeMillis,omitnil,omitempty" name:"StartTimeMillis"`

	// Parent Span ID
	ParentSpanID *string `json:"ParentSpanID,omitnil,omitempty" name:"ParentSpanID"`
}

type SpanLog struct {
	// Log timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Tag.
	Fields []*SpanTag `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type SpanProcess struct {
	// Application service name.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Tags Tag array.
	Tags []*SpanTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SpanReference struct {
	// Type of association relationship.
	RefType *string `json:"RefType,omitnil,omitempty" name:"RefType"`

	// Span ID
	SpanID *string `json:"SpanID,omitnil,omitempty" name:"SpanID"`

	// Trace ID
	TraceID *string `json:"TraceID,omitnil,omitempty" name:"TraceID"`
}

type SpanTag struct {
	// Tag type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Tag key.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	// .
	// Note: this field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type TerminateApmInstanceRequestParams struct {
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TerminateApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TerminateApmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateApmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateApmInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateApmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateApmInstanceResponseParams `json:"Response"`
}

func (r *TerminateApmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}