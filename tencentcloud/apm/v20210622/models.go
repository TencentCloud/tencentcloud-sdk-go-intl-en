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

type AgentOperationConfigView struct {
	// Whether allowlist configuration is enabled for the current API.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetentionValid *bool `json:"RetentionValid,omitnil,omitempty" name:"RetentionValid"`

	// Effective when RetentionValid is false. It indicates blocklist configuration in API settings. The APIs specified in the configuration do not support collection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IgnoreOperation *string `json:"IgnoreOperation,omitnil,omitempty" name:"IgnoreOperation"`

	// Effective when RetentionValid is true. It indicates allowlist configuration in API settings. Only the APIs specified in the configuration support collection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetentionOperation *string `json:"RetentionOperation,omitnil,omitempty" name:"RetentionOperation"`
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

type ApmAppConfig struct {
	// <p>Instance ID</p>
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// <p>Service name</p>
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// <p>URL convergence switch</p>
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// <p>URL convergence threshold</p>
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// <p>URL regular convergence</p>
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// <p>Exception filter regular</p>
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// <p>Error code filtering</p>
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// <p>Service component type</p>
	Components *string `json:"Components,omitnil,omitempty" name:"Components"`

	// <p>URL exclusion regular</p>
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// <p>Log source</p>
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// <p>Log region</p>
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// <p>Whether logging is enabled 0 Disabled 1 Enabled</p>
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// <p>Log topic ID</p>
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// <p>Interface Names to Filter</p>
	IgnoreOperationName *string `json:"IgnoreOperationName,omitnil,omitempty" name:"IgnoreOperationName"`

	// <p>CLS logset | ES cluster ID</p>
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// <p>Number of traces reported by the probe per second</p>
	TraceRateLimit *int64 `json:"TraceRateLimit,omitnil,omitempty" name:"TraceRateLimit"`

	// <p>Whether thread profiling is enabled</p>
	EnableSnapshot *bool `json:"EnableSnapshot,omitnil,omitempty" name:"EnableSnapshot"`

	// <p>Timeout threshold for thread profiling</p>
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// <p>Whether to enable agent</p>
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// <p>Component list</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// <p>Whether to enable link compression</p>
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// <p>Whether the application diagnosis switch is enabled</p>
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// <p>probe API related configuration</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitnil,omitempty" name:"AgentOperationConfigView"`

	// <p>Whether the application log configuration is enabled</p>
	EnableLogConfig *bool `json:"EnableLogConfig,omitnil,omitempty" name:"EnableLogConfig"`

	// <p>Application ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// <p>Whether the dashboard configuration is enabled: false (disabled, consistent with the business system)/true (enabled, hierarchical configuration)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitnil,omitempty" name:"EnableDashboardConfig"`

	// <p>Whether dashboard is associated: 0 Disabled 1 Enabled</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// <p>dashboard ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// <p>Whether the application-level configuration is enabled</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitnil,omitempty" name:"EnableSecurityConfig"`

	// <p>Whether the component vulnerability detection is enabled</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// <p>Whether SQL injection analysis is enabled</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// <p>Whether remote command execution analysis is enabled</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// <p>Whether Java Webshell detection and analysis is enabled</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// <p>CLS index type (0=full-text index, 1=key-value index)</p>
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// <p>Index key of traceId: This parameter is valid only when the CLS index type is key-value index.</p>
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// <p>Whether to enable file deletion detection (0 - disabled, 1 - enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file read detection (0 - disabled, 1 - enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file upload detection (0 - disabled, 1 - enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// <p>Whether to enable detection of arbitrary files (0 - disabled, 1 - enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// <p>Whether path traversal detection is enabled (0-disabled, 1-enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// <p>Whether to enable template engine injection detection (0-disable, 1-enable)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// <p>Whether script engine injection detection is enabled (0-disabled, 1-enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// <p>Whether expression injection detection is enabled (0-disabled, 1-enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// <p>Whether JNDI injection detection is enabled (0-disabled, 1-enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// <p>Whether JNI injection detection is enabled (0 - disabled, 1 - enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// <p>Whether to enable Webshell backdoor detection (0 - disabled, 1 - enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// <p>Whether deserialization detection is enabled (0-disabled, 1-enabled)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// <p>API name auto convergence switch (0-off, 1-on)</p>
	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitnil,omitempty" name:"UrlAutoConvergenceEnable"`

	// <p>URL long segment convergence threshold</p>
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// <p>URL digit segment convergence threshold</p>
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// <p>Fuse memory threshold of the probe</p>
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// <p>Probe fuse CPU threshold</p>
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// <p>Whether SQL parameter access is enabled</p>
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// <p>Slow SQL threshold</p>
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`

	// <p>Whether the masking rule is enabled</p>
	EnableDesensitizationRule *int64 `json:"EnableDesensitizationRule,omitnil,omitempty" name:"EnableDesensitizationRule"`

	// <p>Masking rule</p>
	DesensitizationRule *string `json:"DesensitizationRule,omitnil,omitempty" name:"DesensitizationRule"`

	// <p>Index key of spanId: This parameter is valid only when the CLS index type is key-value index.</p>
	LogSpanIdKey *string `json:"LogSpanIdKey,omitnil,omitempty" name:"LogSpanIdKey"`

	// <p>Automated performance analysis configuration</p>
	AutoProfilingConfig *AutoProfilingConfig `json:"AutoProfilingConfig,omitnil,omitempty" name:"AutoProfilingConfig"`

	// <p>Threshold configuration switch. true means use application level threshold; false means use business system level threshold.</p>
	EnableThresholdConfig *bool `json:"EnableThresholdConfig,omitnil,omitempty" name:"EnableThresholdConfig"`

	// <p>Error rate threshold (%) used to judge the application health status as "red".</p>
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// <p>Alert threshold for response time (ms), used to judge the application health status as "yellow".</p>
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// <p>Whether to use the built-in fuse threshold of the probe by default</p>
	UseDefaultFuseConfig *bool `json:"UseDefaultFuseConfig,omitnil,omitempty" name:"UseDefaultFuseConfig"`
}

type ApmApplicationConfigView struct {
	// <p>Business system ID</p>
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// <p>Application name</p>
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// <p>Interface Filtering</p>
	OperationNameFilter *string `json:"OperationNameFilter,omitnil,omitempty" name:"OperationNameFilter"`

	// <p>Error type filtering</p>
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// <p>HTTP status code filtering</p>
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// <p>Application diagnosis switch (abandoned)</p>
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// <p>URL convergence switch 0 Off 1 On</p>
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// <p>URL convergence threshold</p>
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// <p>URL regular convergence rules</p>
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// <p>URL exclusion rule regex</p>
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// <p>Whether logging is enabled 0 Disabled 1 Enabled</p>
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// <p>Log source</p>
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// <p>Logset</p>
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// <p>Log topic</p>
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// <p>Method stack snapshot switch. Enabled indicates true. false indicates disabled.</p>
	SnapshotEnable *bool `json:"SnapshotEnable,omitnil,omitempty" name:"SnapshotEnable"`

	// <p>Slow call monitoring trigger threshold</p>
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// <p>Master switch for probes</p>
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// <p>Component list toggle (abandoned)</p>
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// <p>Link compression switch (abandoned)</p>
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// <p>Link filtering configuration</p>
	AgentIgnoreOperation *string `json:"AgentIgnoreOperation,omitnil,omitempty" name:"AgentIgnoreOperation"`

	// <p>Enable the application security switch</p>
	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitnil,omitempty" name:"EnableSecurityConfig"`

	// <p>Whether SQL injection detection is enabled</p>
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// <p>Whether component vulnerability detection is enabled</p>
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// <p>Whether remote command execution detection is enabled</p>
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// <p>Whether memory leakage detection is enabled</p>
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// <p>Whether to enable detection of any file deletion</p>
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file read detection</p>
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file upload detection</p>
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// <p>Whether to enable detection of arbitrary files</p>
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// <p>Whether path traversal detection is enabled</p>
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// <p>Whether to enable template engine injection detection</p>
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// <p>Whether script engine injection detection is enabled</p>
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// <p>Whether expression injection detection is enabled</p>
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// <p>Whether JNDI injection detection is enabled</p>
	IsJndiInjectionAnalysis *int64 `json:"IsJndiInjectionAnalysis,omitnil,omitempty" name:"IsJndiInjectionAnalysis"`

	// <p>Whether JNI injection detection is enabled</p>
	IsJniInjectionAnalysis *int64 `json:"IsJniInjectionAnalysis,omitnil,omitempty" name:"IsJniInjectionAnalysis"`

	// <p>Whether Webshell backdoor detection is enabled</p>
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// <p>Whether deserialization detection is enabled</p>
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// <p>Whether the console switch is enabled</p>
	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitnil,omitempty" name:"EnableDashboardConfig"`

	// <p>Whether to associate with Dashboard</p>
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// <p>Dashboard topic</p>
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// <p>Fuse memory threshold of the probe</p>
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// <p>Probe fuse CPU threshold</p>
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// <p>Whether SQL parameter access is enabled</p>
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// <p>Slow SQL threshold</p>
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`

	// <p>Whether the masking rule is enabled</p>
	EnableDesensitizationRule *int64 `json:"EnableDesensitizationRule,omitnil,omitempty" name:"EnableDesensitizationRule"`

	// <p>Masking rule</p>
	DesensitizationRule *string `json:"DesensitizationRule,omitnil,omitempty" name:"DesensitizationRule"`

	// <p>Automated performance analysis task configuration</p>
	AutoProfilingConfig *AutoProfilingConfig `json:"AutoProfilingConfig,omitnil,omitempty" name:"AutoProfilingConfig"`

	// <p>Threshold configuration switch</p>
	EnableThresholdConfig *bool `json:"EnableThresholdConfig,omitnil,omitempty" name:"EnableThresholdConfig"`

	// <p>Error rate threshold</p><p>Unit: %</p>
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// <p>Alert threshold of response time</p><p>Unit: ms</p>
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// <p>Whether to use the default fuse threshold of the probe</p>
	UseDefaultFuseConfig *bool `json:"UseDefaultFuseConfig,omitnil,omitempty" name:"UseDefaultFuseConfig"`
}

type ApmAssociation struct {
	// <p>Instance ID of the associated product</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`

	// <p>Association status: 1 (enabled), 2 (not enabled), 3 (invalid)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>CKafka message topic</p>
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// <p>Ckafka consumption topic</p><p>Used for Kafka metric delivery</p>
	MetricTopic *string `json:"MetricTopic,omitnil,omitempty" name:"MetricTopic"`
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

	// Metric Chinese Name
	NameCN *string `json:"NameCN,omitnil,omitempty" name:"NameCN"`

	// Metric English name
	NameEN *string `json:"NameEN,omitnil,omitempty" name:"NameEN"`
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

	// Whether to enable detection of remote command execution (0 = disabled; 1 = enabled).
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// Whether to enable detection of Java webshell execution (0 = disabled; 1 = enabled).
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// CLS index type. (0 = full-text index; 1 = key-value index).
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// Index key of traceId. It is valid when the CLS index type is key-value index.
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// Whether to enable the detection of deleting arbitrary files. (0 - disabled; 1: enabled.)
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// Whether to enable the detection of reading arbitrary files. (0 - disabled; 1 - enabled.)
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// Whether to enable the detection of uploading arbitrary files. (0 - disabled; 1 - enabled.)
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// Whether to enable the detection of the inclusion of arbitrary files. (0: disabled, 1: enabled.)
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// Whether to enable traversal detection of the directory. (0 - disabled; 1 - enabled).
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// Whether to enable template engine injection detection. (0: disabled; 1: enabled.)
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// Whether to enable script engine injection detection. (0 - disabled; 1 - enabled.)
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// Whether to enable expression injection detection. (0 - disabled; 1 - enabled.)
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// Whether to enable JNDI injection detection. (0 - disabled; 1 - enabled.)
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// Whether to enable JNI injection detection. (0 - disabled, 1 - enabled).
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// Whether to enable Webshell backdoor detection. (0 - disabled; 1 - enabled).
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// Whether to enable deserialization detection. (0 - disabled; 1 - enabled).
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// Business system authentication token.
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// Convergence threshold for URL long segments.
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// Convergence threshold for URL numerical segments.
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// Index key of spanId: This parameter is valid only when the CLS index type is key-value index
	LogSpanIdKey *string `json:"LogSpanIdKey,omitnil,omitempty" name:"LogSpanIdKey"`
}

type ApmMetricRecord struct {
	// Field array, used for the query result of indicators.
	Fields []*ApmField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Tag array, used to distinguish the objects of groupby.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ApmPrometheusRules struct {
	// Metric match rule ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Metric match rule name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Applications where the rule takes effect. input an empty string for all applications.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Specifies the metric match rule status: 1 (enabled), 2 (disabled).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies the metric match rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`

	// Match type: 0 - precision match, 1 - prefix match, 2 - suffix match.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`
}

type ApmSampleConfig struct {
	// Instance ID.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// Service name.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Sampling name
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// API name.
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// Number of spans sampled
	SpanNum *int64 `json:"SpanNum,omitnil,omitempty" name:"SpanNum"`

	// Sampling configuration switch. 0: Off; 1: On
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Tag array
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Sampling rate.
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Specifies the matching method. 0: exact match (default); 1: prefix match; 2: suffix match.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Configuration ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type ApmServiceMetric struct {
	// Field array.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fields []*ApmField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Tag array
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Application information
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceDetail *ServiceDetail `json:"ServiceDetail,omitnil,omitempty" name:"ServiceDetail"`
}

type ApmTag struct {
	// Dimension key (column name, Tag key).
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Dimension value (tag value).
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ApmVulnerabilityServiceDetail struct {
	// Application instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceInstance *string `json:"ServiceInstance,omitnil,omitempty" name:"ServiceInstance"`

	// Path of the jar package with the vulnerability
	// Note: This field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Last occurrence time
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastOccurTime *int64 `json:"LastOccurTime,omitnil,omitempty" name:"LastOccurTime"`
}

type AutoProfilingConfig struct {
	// Auto CPU profiling task switch
	CpuProfilingEnable *bool `json:"CpuProfilingEnable,omitnil,omitempty" name:"CpuProfilingEnable"`

	// Auto memory profiling task switch
	MemoryProfilingEnable *bool `json:"MemoryProfilingEnable,omitnil,omitempty" name:"MemoryProfilingEnable"`

	// Auto CPU profiling task threshold
	CpuProfilingThreshold *int64 `json:"CpuProfilingThreshold,omitnil,omitempty" name:"CpuProfilingThreshold"`

	// Automated memory profiling task threshold
	MemoryProfilingThreshold *int64 `json:"MemoryProfilingThreshold,omitnil,omitempty" name:"MemoryProfilingThreshold"`

	// CPU auto profiling task duration
	CpuProfilingDuration *int64 `json:"CpuProfilingDuration,omitnil,omitempty" name:"CpuProfilingDuration"`

	// Memory auto profiling task duration
	MemoryProfilingDuration *int64 `json:"MemoryProfilingDuration,omitnil,omitempty" name:"MemoryProfilingDuration"`
}

type CVMMeta struct {
	// Region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type ComponentTopologyView struct {
	// Number of nodes at the service level.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Service *int64 `json:"Service,omitnil,omitempty" name:"Service"`

	// Node count of the database.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Database *int64 `json:"Database,omitnil,omitempty" name:"Database"`

	// Node count of the message queue.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MQ *int64 `json:"MQ,omitnil,omitempty" name:"MQ"`
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
type CreateApmPrometheusRuleRequestParams struct {
	// Metric match rule name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Applications where the rule takes effect. input an empty string for all applications.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Match type: 0 - precision match, 1 - prefix match, 2 - suffix match.
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// Specifies the rule for customer-defined metric names with cache hit.
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`

	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CreateApmPrometheusRuleRequest struct {
	*tchttp.BaseRequest
	
	// Metric match rule name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Applications where the rule takes effect. input an empty string for all applications.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Match type: 0 - precision match, 1 - prefix match, 2 - suffix match.
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// Specifies the rule for customer-defined metric names with cache hit.
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`

	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CreateApmPrometheusRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmPrometheusRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ServiceName")
	delete(f, "MetricMatchType")
	delete(f, "MetricNameRule")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApmPrometheusRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmPrometheusRuleResponseParams struct {
	// ID of the metric match rule
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApmPrometheusRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateApmPrometheusRuleResponseParams `json:"Response"`
}

func (r *CreateApmPrometheusRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmPrometheusRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmSampleConfigRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rate.
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// Sampling tags
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// API name.
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 0: exact match (default); 1: prefix match; 2: suffix match.
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type CreateApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rate.
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// Sampling tags
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// API name.
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 0: exact match (default); 1: prefix match; 2: suffix match.
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *CreateApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleRate")
	delete(f, "ServiceName")
	delete(f, "SampleName")
	delete(f, "Tags")
	delete(f, "OperationName")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmSampleConfigResponseParams struct {
	// Sampling configuration parameter
	ApmSampleConfig *ApmSampleConfig `json:"ApmSampleConfig,omitnil,omitempty" name:"ApmSampleConfig"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateApmSampleConfigResponseParams `json:"Response"`
}

func (r *CreateApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProfileTaskRequestParams struct {
	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// APM business system ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Application instance (online).
	ServiceInstance *string `json:"ServiceInstance,omitnil,omitempty" name:"ServiceInstance"`

	// Event type (cpu, alloc).
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// Specifies the task duration in milliseconds (ms). value range: 5 to 180 seconds.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Number of execution. value range: 1-100.
	AllTimes *int64 `json:"AllTimes,omitnil,omitempty" name:"AllTimes"`

	// Start timestamp. 0 indicates start from the current time (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Specifies the task execution interval in milliseconds (ms). value range: 10 to 600 seconds. cannot be less than 1.5 times the Duration.
	TaskInterval *int64 `json:"TaskInterval,omitnil,omitempty" name:"TaskInterval"`
}

type CreateProfileTaskRequest struct {
	*tchttp.BaseRequest
	
	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// APM business system ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Application instance (online).
	ServiceInstance *string `json:"ServiceInstance,omitnil,omitempty" name:"ServiceInstance"`

	// Event type (cpu, alloc).
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// Specifies the task duration in milliseconds (ms). value range: 5 to 180 seconds.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Number of execution. value range: 1-100.
	AllTimes *int64 `json:"AllTimes,omitnil,omitempty" name:"AllTimes"`

	// Start timestamp. 0 indicates start from the current time (unit: seconds).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Specifies the task execution interval in milliseconds (ms). value range: 10 to 600 seconds. cannot be less than 1.5 times the Duration.
	TaskInterval *int64 `json:"TaskInterval,omitnil,omitempty" name:"TaskInterval"`
}

func (r *CreateProfileTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProfileTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
	delete(f, "InstanceId")
	delete(f, "ServiceInstance")
	delete(f, "Event")
	delete(f, "Duration")
	delete(f, "AllTimes")
	delete(f, "StartTime")
	delete(f, "TaskInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProfileTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProfileTaskResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProfileTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProfileTaskResponseParams `json:"Response"`
}

func (r *CreateProfileTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProfileTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApmSampleConfigRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

type DeleteApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

func (r *DeleteApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApmSampleConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApmSampleConfigResponseParams `json:"Response"`
}

func (r *DeleteApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApmSampleConfigResponse) FromJsonString(s string) error {
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
type DescribeApmAllVulCountRequestParams struct {
	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeApmAllVulCountRequest struct {
	*tchttp.BaseRequest
	
	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeApmAllVulCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAllVulCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmAllVulCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAllVulCountResponseParams struct {
	// Vulnerability metrics as well as number of business systems	
	VulnerabilityList []*ApmMetricRecord `json:"VulnerabilityList,omitnil,omitempty" name:"VulnerabilityList"`

	// Total number of vulnerabilities
	VulnerabilityCount *int64 `json:"VulnerabilityCount,omitnil,omitempty" name:"VulnerabilityCount"`

	// Number of critical vulnerabilities
	ImportantVulnerabilityCount *int64 `json:"ImportantVulnerabilityCount,omitnil,omitempty" name:"ImportantVulnerabilityCount"`

	// High-risk vulnerability count
	CriticalVulnerabilityCount *int64 `json:"CriticalVulnerabilityCount,omitnil,omitempty" name:"CriticalVulnerabilityCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmAllVulCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmAllVulCountResponseParams `json:"Response"`
}

func (r *DescribeApmAllVulCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAllVulCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmApplicationConfigRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type DescribeApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Service name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

func (r *DescribeApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmApplicationConfigResponseParams struct {
	// Apm application configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApmAppConfig *ApmAppConfig `json:"ApmAppConfig,omitnil,omitempty" name:"ApmAppConfig"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmApplicationConfigResponseParams `json:"Response"`
}

func (r *DescribeApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAssociationRequestParams struct {
	// Associated product name. currently only supports Prometheus.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Business System Name
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeApmAssociationRequest struct {
	*tchttp.BaseRequest
	
	// Associated product name. currently only supports Prometheus.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Business System Name
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeApmAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAssociationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAssociationResponseParams struct {
	// Associated product instance ID.
	ApmAssociation *ApmAssociation `json:"ApmAssociation,omitnil,omitempty" name:"ApmAssociation"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmAssociationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmAssociationResponseParams `json:"Response"`
}

func (r *DescribeApmAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmInstancesRequestParams struct {
	// Tag list.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Filters by business system name, and fuzzy search is supported.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Filters by business system ID, and fuzzy search is supported.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

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

	// Filters by business system name, and fuzzy search is supported.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Filters by business system ID, and fuzzy search is supported.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

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
	delete(f, "InstanceId")
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
type DescribeApmPrometheusRuleRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeApmPrometheusRuleRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeApmPrometheusRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmPrometheusRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmPrometheusRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmPrometheusRuleResponseParams struct {
	// Specifies the metric match rule.
	ApmPrometheusRules []*ApmPrometheusRules `json:"ApmPrometheusRules,omitnil,omitempty" name:"ApmPrometheusRules"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmPrometheusRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmPrometheusRuleResponseParams `json:"Response"`
}

func (r *DescribeApmPrometheusRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmPrometheusRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmSQLInjectionDetailRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Query filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Aggregation dimension
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Metric list
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

type DescribeApmSQLInjectionDetailRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Query filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Aggregation dimension
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Metric list
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

func (r *DescribeApmSQLInjectionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmSQLInjectionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrderBy")
	delete(f, "Filters")
	delete(f, "GroupBy")
	delete(f, "Metrics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmSQLInjectionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmSQLInjectionDetailResponseParams struct {
	// SQL dimension information
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Link information
	Records []*ApmMetricRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmSQLInjectionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmSQLInjectionDetailResponseParams `json:"Response"`
}

func (r *DescribeApmSQLInjectionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmSQLInjectionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmSampleConfigRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

type DescribeApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

func (r *DescribeApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmSampleConfigResponseParams struct {
	// Sampling configuration list
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApmSampleConfigs []*ApmSampleConfig `json:"ApmSampleConfigs,omitnil,omitempty" name:"ApmSampleConfigs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmSampleConfigResponseParams `json:"Response"`
}

func (r *DescribeApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmServiceMetricRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Application ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// Start time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Whether to use the demo mode.
	Demo *bool `json:"Demo,omitnil,omitempty" name:"Demo"`

	// Application status filter criteria. valid values: health, warning, error. if multiple statuses are selected, separate them by commas, for example: "warning,error".
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// Tag list
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Page number.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Page size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeApmServiceMetricRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Application ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// Start time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Order.
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Whether to use the demo mode.
	Demo *bool `json:"Demo,omitnil,omitempty" name:"Demo"`

	// Application status filter criteria. valid values: health, warning, error. if multiple statuses are selected, separate them by commas, for example: "warning,error".
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// Tag list
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Page number.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Page size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeApmServiceMetricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmServiceMetricRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	delete(f, "ServiceID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrderBy")
	delete(f, "Demo")
	delete(f, "ServiceStatus")
	delete(f, "Tags")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmServiceMetricRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmServiceMetricResponseParams struct {
	// List of application metrics.
	ServiceMetricList []*ApmServiceMetric `json:"ServiceMetricList,omitnil,omitempty" name:"ServiceMetricList"`

	// Number of applications meeting the filtering conditions.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Warning of the abnormal number of applications.
	WarningErrorCount *int64 `json:"WarningErrorCount,omitnil,omitempty" name:"WarningErrorCount"`

	// Total number of applications.
	ApplicationCount *int64 `json:"ApplicationCount,omitnil,omitempty" name:"ApplicationCount"`

	// Page number.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Page size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Indicates the number of abnormal applications.
	ErrorCount *int64 `json:"ErrorCount,omitnil,omitempty" name:"ErrorCount"`

	// Warning of the number of applications.
	WarningCount *int64 `json:"WarningCount,omitnil,omitempty" name:"WarningCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmServiceMetricResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmServiceMetricResponseParams `json:"Response"`
}

func (r *DescribeApmServiceMetricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmServiceMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmVulnerabilityCountRequestParams struct {
	// APM business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// app name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data type of queried data. Attack detection is "attack_detect".
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeApmVulnerabilityCountRequest struct {
	*tchttp.BaseRequest
	
	// APM business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// app name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data type of queried data. Attack detection is "attack_detect".
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeApmVulnerabilityCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmVulnerabilityCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmVulnerabilityCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmVulnerabilityCountResponseParams struct {
	// Vulnerability metrics as well as number of business systems
	VulnerabilityList []*ApmMetricRecord `json:"VulnerabilityList,omitnil,omitempty" name:"VulnerabilityList"`

	// Total number of vulnerabilities
	VulnerabilityCount *int64 `json:"VulnerabilityCount,omitnil,omitempty" name:"VulnerabilityCount"`

	// Number of critical vulnerabilities
	ImportantVulnerabilityCount *int64 `json:"ImportantVulnerabilityCount,omitnil,omitempty" name:"ImportantVulnerabilityCount"`

	// High-risk vulnerability count
	CriticalVulnerabilityCount *int64 `json:"CriticalVulnerabilityCount,omitnil,omitempty" name:"CriticalVulnerabilityCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmVulnerabilityCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmVulnerabilityCountResponseParams `json:"Response"`
}

func (r *DescribeApmVulnerabilityCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmVulnerabilityCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmVulnerabilityDetailRequestParams struct {
	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// APM business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Conditional filtering, required service.name, instrumentation.name, version, vul.id
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeApmVulnerabilityDetailRequest struct {
	*tchttp.BaseRequest
	
	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// APM business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Conditional filtering, required service.name, instrumentation.name, version, vul.id
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeApmVulnerabilityDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmVulnerabilityDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmVulnerabilityDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmVulnerabilityDetailResponseParams struct {
	// Vulnerability details
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// List of business systems related to vulnerabilities
	ServiceInstanceList []*ApmVulnerabilityServiceDetail `json:"ServiceInstanceList,omitnil,omitempty" name:"ServiceInstanceList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmVulnerabilityDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmVulnerabilityDetailResponseParams `json:"Response"`
}

func (r *DescribeApmVulnerabilityDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmVulnerabilityDetailResponse) FromJsonString(s string) error {
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
	//  -If the time span is (48, +inf) hours, it is aggregated at an hourly granularity.
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
	//  -If the time span is (48, +inf) hours, it is aggregated at an hourly granularity.
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
	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Span query start timestamp (unit: seconds)</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Span query end timestamp (unit: s)</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Universal filter parameters supported filter key such as service.name</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Sort<br>Keys now supported:</p><ul><li>startTime</li><li>endTime</li><li>duration</li></ul><p>Values now supported:</p><ul><li>desc (sort in descending order)</li><li>asc (ascending order)</li></ul>
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Business service name. Console users please enter taw</p>
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// <p>Count of single-page projects. Defaults to 10000. Valid value range is 0–10000.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Page</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeGeneralOTSpanListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Span query start timestamp (unit: seconds)</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Span query end timestamp (unit: s)</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Universal filter parameters supported filter key such as service.name</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Sort<br>Keys now supported:</p><ul><li>startTime</li><li>endTime</li><li>duration</li></ul><p>Values now supported:</p><ul><li>desc (sort in descending order)</li><li>asc (ascending order)</li></ul>
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Business service name. Console users please enter taw</p>
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// <p>Count of single-page projects. Defaults to 10000. Valid value range is 0–10000.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Page</p>
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
	// <p>Total number</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>The Spans field contains all content of the link data. Since the data is compressed, perform the following three steps to restore the original text.</p><ol><li>Decode the text in the Spans field with Base64 to get the compressed byte[].</li><li>Decompress the compressed byte[] with gzip to get the byte array before compression.</li><li>Convert the byte array before compression to text using the UTF-8 character set.</li></ol>
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
	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Span query start timestamp (unit: seconds)</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Span query end timestamp (unit: s)</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Universal filter parameters supported filter key such as service.name</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Sort<br>Keys now supported:</p><ul><li>startTime</li><li>endTime</li><li>duration</li></ul><p>Values now supported:</p><ul><li>desc (sort in descending order)</li><li>asc (ascending order)</li></ul>
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Business service name. Console users please enter taw</p>
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// <p>Count of single-page projects. Defaults to 1000. Valid value range is 1–1000.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Page</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeGeneralSpanListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Span query start timestamp (unit: seconds)</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Span query end timestamp (unit: s)</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Universal filter parameters supported filter key such as service.name</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Sort<br>Keys now supported:</p><ul><li>startTime</li><li>endTime</li><li>duration</li></ul><p>Values now supported:</p><ul><li>desc (sort in descending order)</li><li>asc (ascending order)</li></ul>
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Business service name. Console users please enter taw</p>
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// <p>Count of single-page projects. Defaults to 1000. Valid value range is 1–1000.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Page</p>
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
	// <p>Total number</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Span pagination list</p>
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

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or filter criteria.
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

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

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or filter criteria.
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

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
	delete(f, "GroupBy")
	delete(f, "Filters")
	delete(f, "OrFilters")
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
type DescribeOPRAllVulCountRequestParams struct {
	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeOPRAllVulCountRequest struct {
	*tchttp.BaseRequest
	
	// unix second-level timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// unix second-level timestamp
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeOPRAllVulCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOPRAllVulCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOPRAllVulCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOPRAllVulCountResponseParams struct {
	// Vulnerability metrics as well as number of business systems	
	VulnerabilityList []*ApmMetricRecord `json:"VulnerabilityList,omitnil,omitempty" name:"VulnerabilityList"`

	// Total number of vulnerabilities
	VulnerabilityCount *int64 `json:"VulnerabilityCount,omitnil,omitempty" name:"VulnerabilityCount"`

	// Number of critical vulnerabilities
	ImportantVulnerabilityCount *int64 `json:"ImportantVulnerabilityCount,omitnil,omitempty" name:"ImportantVulnerabilityCount"`

	// High-risk vulnerability count
	CriticalVulnerabilityCount *int64 `json:"CriticalVulnerabilityCount,omitnil,omitempty" name:"CriticalVulnerabilityCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOPRAllVulCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOPRAllVulCountResponseParams `json:"Response"`
}

func (r *DescribeOPRAllVulCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOPRAllVulCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceOverviewRequestParams struct {
	// Metric list.
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// Business system id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Start time (unit: sec).
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

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

	// Aggregation dimension.
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// Start time (unit: sec).
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

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
	delete(f, "GroupBy")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
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

	// Start time (unit: sec).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

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

	// Start time (unit: sec).
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (unit: seconds).
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
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

// Predefined struct for user
type DescribeTopologyNewRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Querying start time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Upstream level.
	UpLevel *int64 `json:"UpLevel,omitnil,omitempty" name:"UpLevel"`

	// Application instance information.
	ServiceInstance *string `json:"ServiceInstance,omitnil,omitempty" name:"ServiceInstance"`

	// downstream hierarchy
	DownLevel *int64 `json:"DownLevel,omitnil,omitempty" name:"DownLevel"`

	// perspective
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// Filter.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Represents topic (for MQ topology)
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// View filtering list.
	Selectors *Selectors `json:"Selectors,omitnil,omitempty" name:"Selectors"`

	// View ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// TraceID
	TraceID *string `json:"TraceID,omitnil,omitempty" name:"TraceID"`

	// Specifies the top 5 slow response nodes query.
	IsSlowTopFive *bool `json:"IsSlowTopFive,omitnil,omitempty" name:"IsSlowTopFive"`

	// Whether the resource layer information is obtained.
	GetResource *bool `json:"GetResource,omitnil,omitempty" name:"GetResource"`

	// Filters by application tag.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Node type not displayed.
	Hidden *Selectors `json:"Hidden,omitnil,omitempty" name:"Hidden"`
}

type DescribeTopologyNewRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Querying start time.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application name
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Upstream level.
	UpLevel *int64 `json:"UpLevel,omitnil,omitempty" name:"UpLevel"`

	// Application instance information.
	ServiceInstance *string `json:"ServiceInstance,omitnil,omitempty" name:"ServiceInstance"`

	// downstream hierarchy
	DownLevel *int64 `json:"DownLevel,omitnil,omitempty" name:"DownLevel"`

	// perspective
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// Filter.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Represents topic (for MQ topology)
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// View filtering list.
	Selectors *Selectors `json:"Selectors,omitnil,omitempty" name:"Selectors"`

	// View ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// TraceID
	TraceID *string `json:"TraceID,omitnil,omitempty" name:"TraceID"`

	// Specifies the top 5 slow response nodes query.
	IsSlowTopFive *bool `json:"IsSlowTopFive,omitnil,omitempty" name:"IsSlowTopFive"`

	// Whether the resource layer information is obtained.
	GetResource *bool `json:"GetResource,omitnil,omitempty" name:"GetResource"`

	// Filters by application tag.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Node type not displayed.
	Hidden *Selectors `json:"Hidden,omitnil,omitempty" name:"Hidden"`
}

func (r *DescribeTopologyNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopologyNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ServiceName")
	delete(f, "UpLevel")
	delete(f, "ServiceInstance")
	delete(f, "DownLevel")
	delete(f, "View")
	delete(f, "Filters")
	delete(f, "Topic")
	delete(f, "Selectors")
	delete(f, "Id")
	delete(f, "TraceID")
	delete(f, "IsSlowTopFive")
	delete(f, "GetResource")
	delete(f, "Tags")
	delete(f, "Hidden")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopologyNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopologyNewResponseParams struct {
	// Node collection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nodes []*TopologyNode `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// Edge set.
	Edges []*TopologyEdgeNew `json:"Edges,omitnil,omitempty" name:"Edges"`

	// Whether the topology map is modified
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopologyModifyFlag *int64 `json:"TopologyModifyFlag,omitnil,omitempty" name:"TopologyModifyFlag"`

	// Number of nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Selectors *SelectorView `json:"Selectors,omitnil,omitempty" name:"Selectors"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopologyNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopologyNewResponseParams `json:"Response"`
}

func (r *DescribeTopologyNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopologyNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// <p>Filter method (=, !=, in)</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Filter dimension name</p><p>For details, see the actual interface field description</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Filter value. Use comma-separated multiple values for in filtering method.</p>
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

	// Metric data unit
	MetricUnit *string `json:"MetricUnit,omitnil,omitempty" name:"MetricUnit"`
}

// Predefined struct for user
type ModifyApmApplicationConfigRequestParams struct {
	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Application name</p>
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// <p>URL convergence switch, 0 Off | 1 On</p>
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// <p>URL convergence threshold</p>
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// <p>Exception filtering regex rules, comma-separated</p>
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// <p>URL convergence regex rules, comma-separated</p>
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// <p>Error code filtering, comma-separated</p>
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// <p>URL exclusion regex rule, comma-separated</p>
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// <p>Log switch 0 Disabled 1 Enabled</p>
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// <p>Log region</p>
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// <p>Log topic ID</p>
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// <p>CLS logset | ES cluster ID</p>
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// <p>Log source CLS | ES</p>
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// <p>Interfaces to Filter</p>
	IgnoreOperationName *string `json:"IgnoreOperationName,omitnil,omitempty" name:"IgnoreOperationName"`

	// <p>Whether thread profiling is enabled</p>
	EnableSnapshot *bool `json:"EnableSnapshot,omitnil,omitempty" name:"EnableSnapshot"`

	// <p>Timeout threshold for thread profiling</p>
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// <p>Whether to enable agent</p>
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// <p>Whether to enable link compression</p>
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// <p>Whether the switch for enabling application diagnosis is enabled</p>
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// <p>Component list</p>
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// <p>probe API related configuration</p>
	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitnil,omitempty" name:"AgentOperationConfigView"`

	// <p>Whether the application log configuration is enabled</p>
	EnableLogConfig *bool `json:"EnableLogConfig,omitnil,omitempty" name:"EnableLogConfig"`

	// <p>Whether the dashboard configuration is enabled: false (disabled, consistent with the business system)/true (enabled, application-level configuration)</p>
	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitnil,omitempty" name:"EnableDashboardConfig"`

	// <p>Whether to associate with dashboard: 0 off 1 on</p>
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// <p>dashboard ID</p>
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// <p>CLS index type (0=full-text index, 1=key-value index)</p>
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// <p>Index key of traceId: This parameter is valid only when the CLS index type is key-value index.</p>
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// <p>Whether application security configuration is enabled</p>
	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitnil,omitempty" name:"EnableSecurityConfig"`

	// <p>Whether SQL injection analysis is enabled</p>
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// <p>Whether component vulnerability detection is enabled</p>
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// <p>Whether remote command detection is enabled</p>
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// <p>Whether Java Webshell detection is enabled</p>
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// <p>Whether to enable detection of any file deletion (0 - turn off, 1 - turn on)</p>
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file read detection (0 - disabled, 1 - enabled)</p>
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file upload detection (0-disable, 1-enable)</p>
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// <p>Whether to enable detection of arbitrary files (0 - disabled, 1 - enabled)</p>
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// <p>Whether path traversal detection is enabled (0-disabled, 1-enabled)</p>
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// <p>Whether to enable template engine injection detection (0-disable, 1-enable)</p>
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// <p>Whether to enable script engine injection detection (0-disable, 1-enable)</p>
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// <p>Whether expression injection detection is enabled (0-disabled, 1-enabled)</p>
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// <p>Whether JNDI injection detection is enabled (0 - disabled, 1 - enabled)</p>
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// <p>Whether JNI injection detection is enabled (0-disabled, 1-enabled)</p>
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// <p>Whether to enable Webshell backdoor detection (0 - disabled, 1 - enabled)</p>
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// <p>Whether deserialization detection is enabled (0-disabled, 1-enabled)</p>
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// <p>API auto convergence switch, 0-off | 1-on</p>
	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitnil,omitempty" name:"UrlAutoConvergenceEnable"`

	// <p>URL long segment convergence threshold</p>
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// <p>URL digit segment convergence threshold</p>
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// <p>Fuse memory threshold of the probe</p>
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// <p>Probe fuse CPU threshold</p>
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// <p>Whether SQL parameter access is enabled</p>
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// <p>Slow SQL threshold</p>
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`

	// <p>Whether the masking rule is enabled</p>
	EnableDesensitizationRule *int64 `json:"EnableDesensitizationRule,omitnil,omitempty" name:"EnableDesensitizationRule"`

	// <p>Masking rule</p>
	DesensitizationRule *string `json:"DesensitizationRule,omitnil,omitempty" name:"DesensitizationRule"`

	// <p>Index key of spanId: This parameter is valid only when the CLS index type is key-value index.</p>
	LogSpanIdKey *string `json:"LogSpanIdKey,omitnil,omitempty" name:"LogSpanIdKey"`

	// <p>Automated performance analysis task configuration</p>
	AutoProfilingConfig *AutoProfilingConfig `json:"AutoProfilingConfig,omitnil,omitempty" name:"AutoProfilingConfig"`

	// <p>Threshold configuration switch. true means use application level threshold; false means use business system level threshold.</p>
	EnableThresholdConfig *bool `json:"EnableThresholdConfig,omitnil,omitempty" name:"EnableThresholdConfig"`

	// <p>Error rate threshold (%) used to judge the application health status as "red".</p>
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// <p>Alert threshold for response time (ms), used to judge the application health status as "yellow".</p>
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// <p>Whether to use the default fuse threshold of the probe</p>
	UseDefaultFuseConfig *bool `json:"UseDefaultFuseConfig,omitnil,omitempty" name:"UseDefaultFuseConfig"`
}

type ModifyApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Application name</p>
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// <p>URL convergence switch, 0 Off | 1 On</p>
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// <p>URL convergence threshold</p>
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// <p>Exception filtering regex rules, comma-separated</p>
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// <p>URL convergence regex rules, comma-separated</p>
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// <p>Error code filtering, comma-separated</p>
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// <p>URL exclusion regex rule, comma-separated</p>
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// <p>Log switch 0 Disabled 1 Enabled</p>
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// <p>Log region</p>
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// <p>Log topic ID</p>
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// <p>CLS logset | ES cluster ID</p>
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// <p>Log source CLS | ES</p>
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// <p>Interfaces to Filter</p>
	IgnoreOperationName *string `json:"IgnoreOperationName,omitnil,omitempty" name:"IgnoreOperationName"`

	// <p>Whether thread profiling is enabled</p>
	EnableSnapshot *bool `json:"EnableSnapshot,omitnil,omitempty" name:"EnableSnapshot"`

	// <p>Timeout threshold for thread profiling</p>
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// <p>Whether to enable agent</p>
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// <p>Whether to enable link compression</p>
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// <p>Whether the switch for enabling application diagnosis is enabled</p>
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// <p>Component list</p>
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// <p>probe API related configuration</p>
	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitnil,omitempty" name:"AgentOperationConfigView"`

	// <p>Whether the application log configuration is enabled</p>
	EnableLogConfig *bool `json:"EnableLogConfig,omitnil,omitempty" name:"EnableLogConfig"`

	// <p>Whether the dashboard configuration is enabled: false (disabled, consistent with the business system)/true (enabled, application-level configuration)</p>
	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitnil,omitempty" name:"EnableDashboardConfig"`

	// <p>Whether to associate with dashboard: 0 off 1 on</p>
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// <p>dashboard ID</p>
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// <p>CLS index type (0=full-text index, 1=key-value index)</p>
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// <p>Index key of traceId: This parameter is valid only when the CLS index type is key-value index.</p>
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// <p>Whether application security configuration is enabled</p>
	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitnil,omitempty" name:"EnableSecurityConfig"`

	// <p>Whether SQL injection analysis is enabled</p>
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// <p>Whether component vulnerability detection is enabled</p>
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// <p>Whether remote command detection is enabled</p>
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// <p>Whether Java Webshell detection is enabled</p>
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// <p>Whether to enable detection of any file deletion (0 - turn off, 1 - turn on)</p>
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file read detection (0 - disabled, 1 - enabled)</p>
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// <p>Whether to enable arbitrary file upload detection (0-disable, 1-enable)</p>
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// <p>Whether to enable detection of arbitrary files (0 - disabled, 1 - enabled)</p>
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// <p>Whether path traversal detection is enabled (0-disabled, 1-enabled)</p>
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// <p>Whether to enable template engine injection detection (0-disable, 1-enable)</p>
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// <p>Whether to enable script engine injection detection (0-disable, 1-enable)</p>
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// <p>Whether expression injection detection is enabled (0-disabled, 1-enabled)</p>
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// <p>Whether JNDI injection detection is enabled (0 - disabled, 1 - enabled)</p>
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// <p>Whether JNI injection detection is enabled (0-disabled, 1-enabled)</p>
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// <p>Whether to enable Webshell backdoor detection (0 - disabled, 1 - enabled)</p>
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// <p>Whether deserialization detection is enabled (0-disabled, 1-enabled)</p>
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// <p>API auto convergence switch, 0-off | 1-on</p>
	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitnil,omitempty" name:"UrlAutoConvergenceEnable"`

	// <p>URL long segment convergence threshold</p>
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// <p>URL digit segment convergence threshold</p>
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// <p>Fuse memory threshold of the probe</p>
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// <p>Probe fuse CPU threshold</p>
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// <p>Whether SQL parameter access is enabled</p>
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// <p>Slow SQL threshold</p>
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`

	// <p>Whether the masking rule is enabled</p>
	EnableDesensitizationRule *int64 `json:"EnableDesensitizationRule,omitnil,omitempty" name:"EnableDesensitizationRule"`

	// <p>Masking rule</p>
	DesensitizationRule *string `json:"DesensitizationRule,omitnil,omitempty" name:"DesensitizationRule"`

	// <p>Index key of spanId: This parameter is valid only when the CLS index type is key-value index.</p>
	LogSpanIdKey *string `json:"LogSpanIdKey,omitnil,omitempty" name:"LogSpanIdKey"`

	// <p>Automated performance analysis task configuration</p>
	AutoProfilingConfig *AutoProfilingConfig `json:"AutoProfilingConfig,omitnil,omitempty" name:"AutoProfilingConfig"`

	// <p>Threshold configuration switch. true means use application level threshold; false means use business system level threshold.</p>
	EnableThresholdConfig *bool `json:"EnableThresholdConfig,omitnil,omitempty" name:"EnableThresholdConfig"`

	// <p>Error rate threshold (%) used to judge the application health status as "red".</p>
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// <p>Alert threshold for response time (ms), used to judge the application health status as "yellow".</p>
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// <p>Whether to use the default fuse threshold of the probe</p>
	UseDefaultFuseConfig *bool `json:"UseDefaultFuseConfig,omitnil,omitempty" name:"UseDefaultFuseConfig"`
}

func (r *ModifyApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	delete(f, "UrlConvergenceSwitch")
	delete(f, "UrlConvergenceThreshold")
	delete(f, "ExceptionFilter")
	delete(f, "UrlConvergence")
	delete(f, "ErrorCodeFilter")
	delete(f, "UrlExclude")
	delete(f, "IsRelatedLog")
	delete(f, "LogRegion")
	delete(f, "LogTopicID")
	delete(f, "LogSet")
	delete(f, "LogSource")
	delete(f, "IgnoreOperationName")
	delete(f, "EnableSnapshot")
	delete(f, "SnapshotTimeout")
	delete(f, "AgentEnable")
	delete(f, "TraceSquash")
	delete(f, "EventEnable")
	delete(f, "InstrumentList")
	delete(f, "AgentOperationConfigView")
	delete(f, "EnableLogConfig")
	delete(f, "EnableDashboardConfig")
	delete(f, "IsRelatedDashboard")
	delete(f, "DashboardTopicID")
	delete(f, "LogIndexType")
	delete(f, "LogTraceIdKey")
	delete(f, "EnableSecurityConfig")
	delete(f, "IsSqlInjectionAnalysis")
	delete(f, "IsInstrumentationVulnerabilityScan")
	delete(f, "IsRemoteCommandExecutionAnalysis")
	delete(f, "IsMemoryHijackingAnalysis")
	delete(f, "IsDeleteAnyFileAnalysis")
	delete(f, "IsReadAnyFileAnalysis")
	delete(f, "IsUploadAnyFileAnalysis")
	delete(f, "IsIncludeAnyFileAnalysis")
	delete(f, "IsDirectoryTraversalAnalysis")
	delete(f, "IsTemplateEngineInjectionAnalysis")
	delete(f, "IsScriptEngineInjectionAnalysis")
	delete(f, "IsExpressionInjectionAnalysis")
	delete(f, "IsJNDIInjectionAnalysis")
	delete(f, "IsJNIInjectionAnalysis")
	delete(f, "IsWebshellBackdoorAnalysis")
	delete(f, "IsDeserializationAnalysis")
	delete(f, "UrlAutoConvergenceEnable")
	delete(f, "UrlLongSegmentThreshold")
	delete(f, "UrlNumberSegmentThreshold")
	delete(f, "DisableMemoryUsed")
	delete(f, "DisableCpuUsed")
	delete(f, "DbStatementParametersEnabled")
	delete(f, "SlowSQLThresholds")
	delete(f, "EnableDesensitizationRule")
	delete(f, "DesensitizationRule")
	delete(f, "LogSpanIdKey")
	delete(f, "AutoProfilingConfig")
	delete(f, "EnableThresholdConfig")
	delete(f, "ErrRateThreshold")
	delete(f, "ResponseDurationWarningThreshold")
	delete(f, "UseDefaultFuseConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmApplicationConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmApplicationConfigResponseParams `json:"Response"`
}

func (r *ModifyApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmAssociationRequestParams struct {
	// <p>Associated product name, currently only support Prometheus, CKafka</p>
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// <p>Status of the association relationship: // Association relationship status: 1 (enabled), 2 (disabled)</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>ID of the associated product instance</p>
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`

	// <p>CKafka message topic</p>
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// <p>Ckafka message topic</p>
	MetricTopic *string `json:"MetricTopic,omitnil,omitempty" name:"MetricTopic"`
}

type ModifyApmAssociationRequest struct {
	*tchttp.BaseRequest
	
	// <p>Associated product name, currently only support Prometheus, CKafka</p>
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// <p>Status of the association relationship: // Association relationship status: 1 (enabled), 2 (disabled)</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Business system ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>ID of the associated product instance</p>
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`

	// <p>CKafka message topic</p>
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// <p>Ckafka message topic</p>
	MetricTopic *string `json:"MetricTopic,omitnil,omitempty" name:"MetricTopic"`
}

func (r *ModifyApmAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmAssociationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "Status")
	delete(f, "InstanceId")
	delete(f, "PeerId")
	delete(f, "Topic")
	delete(f, "MetricTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmAssociationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmAssociationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmAssociationResponseParams `json:"Response"`
}

func (r *ModifyApmAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Whether to enable detection of the remote command attack.
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// Whether to enable detection of Java webshell.
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// CLS index type. (0 = full-text index; 1 = key-value index).
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// Index key of traceId. It is valid when the CLS index type is key-value index.
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// Whether to enable the detection of deleting arbitrary files. (0 - disabled; 1: enabled.)
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// Whether to enable the detection of reading arbitrary files. (0 - disabled; 1 - enabled.)
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// Whether to enable the detection of uploading arbitrary files. (0 - disabled; 1 - enabled.)
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// Whether to enable the detection of the inclusion of arbitrary files. (0: disabled, 1: enabled.)
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// Whether to enable traversal detection of the directory. (0 - disabled; 1 - enabled).
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// Whether to enable template engine injection detection. (0: disabled; 1: enabled.)
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// Whether to enable script engine injection detection. (0 - disabled; 1 - enabled.)
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// Whether to enable expression injection detection. (0 - disabled; 1 - enabled.)
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// Whether to enable JNDI injection detection. (0 - disabled; 1 - enabled.)
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// Whether to enable JNI injection detection. (0 - disabled, 1 - enabled).
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// Whether to enable Webshell backdoor detection. (0 - disabled; 1 - enabled).
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// Whether to enable deserialization detection. (0 - disabled; 1 - enabled).
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// Convergence threshold for URL long segments.
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// Convergence threshold for URL numerical segments.
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// Index key of spanId: This parameter is valid only when the CLS index type is key-value index
	LogSpanIdKey *string `json:"LogSpanIdKey,omitnil,omitempty" name:"LogSpanIdKey"`
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

	// Whether to enable detection of the remote command attack.
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// Whether to enable detection of Java webshell.
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// CLS index type. (0 = full-text index; 1 = key-value index).
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// Index key of traceId. It is valid when the CLS index type is key-value index.
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// Whether to enable the detection of deleting arbitrary files. (0 - disabled; 1: enabled.)
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// Whether to enable the detection of reading arbitrary files. (0 - disabled; 1 - enabled.)
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// Whether to enable the detection of uploading arbitrary files. (0 - disabled; 1 - enabled.)
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// Whether to enable the detection of the inclusion of arbitrary files. (0: disabled, 1: enabled.)
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// Whether to enable traversal detection of the directory. (0 - disabled; 1 - enabled).
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// Whether to enable template engine injection detection. (0: disabled; 1: enabled.)
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// Whether to enable script engine injection detection. (0 - disabled; 1 - enabled.)
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// Whether to enable expression injection detection. (0 - disabled; 1 - enabled.)
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// Whether to enable JNDI injection detection. (0 - disabled; 1 - enabled.)
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// Whether to enable JNI injection detection. (0 - disabled, 1 - enabled).
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// Whether to enable Webshell backdoor detection. (0 - disabled; 1 - enabled).
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// Whether to enable deserialization detection. (0 - disabled; 1 - enabled).
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// Convergence threshold for URL long segments.
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// Convergence threshold for URL numerical segments.
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// Index key of spanId: This parameter is valid only when the CLS index type is key-value index
	LogSpanIdKey *string `json:"LogSpanIdKey,omitnil,omitempty" name:"LogSpanIdKey"`
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
	delete(f, "IsRemoteCommandExecutionAnalysis")
	delete(f, "IsMemoryHijackingAnalysis")
	delete(f, "LogIndexType")
	delete(f, "LogTraceIdKey")
	delete(f, "IsDeleteAnyFileAnalysis")
	delete(f, "IsReadAnyFileAnalysis")
	delete(f, "IsUploadAnyFileAnalysis")
	delete(f, "IsIncludeAnyFileAnalysis")
	delete(f, "IsDirectoryTraversalAnalysis")
	delete(f, "IsTemplateEngineInjectionAnalysis")
	delete(f, "IsScriptEngineInjectionAnalysis")
	delete(f, "IsExpressionInjectionAnalysis")
	delete(f, "IsJNDIInjectionAnalysis")
	delete(f, "IsJNIInjectionAnalysis")
	delete(f, "IsWebshellBackdoorAnalysis")
	delete(f, "IsDeserializationAnalysis")
	delete(f, "UrlLongSegmentThreshold")
	delete(f, "UrlNumberSegmentThreshold")
	delete(f, "LogSpanIdKey")
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
type ModifyApmPrometheusRuleRequestParams struct {
	// Rule ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the rule name to modify.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Rule status: 1 (enabled), 2 (disabled), 3 (deleted).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Applications where the rule takes effect. input an empty string for all applications. if this is not modified, pass the old rule.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Match type: 0 - precision match, 1 - prefix match, 2 - suffix match (if not modified, the old rule must be passed).
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// Specifies the rule for customer-defined metric names with cache hit.
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`
}

type ModifyApmPrometheusRuleRequest struct {
	*tchttp.BaseRequest
	
	// Rule ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the rule name to modify.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Rule status: 1 (enabled), 2 (disabled), 3 (deleted).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Applications where the rule takes effect. input an empty string for all applications. if this is not modified, pass the old rule.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Match type: 0 - precision match, 1 - prefix match, 2 - suffix match (if not modified, the old rule must be passed).
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// Specifies the rule for customer-defined metric names with cache hit.
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`
}

func (r *ModifyApmPrometheusRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmPrometheusRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "ServiceName")
	delete(f, "MetricMatchType")
	delete(f, "MetricNameRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmPrometheusRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmPrometheusRuleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmPrometheusRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmPrometheusRuleResponseParams `json:"Response"`
}

func (r *ModifyApmPrometheusRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmPrometheusRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmSampleConfigRequestParams struct {
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// Sampling rate.
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Application name. specifies the application name. fill in the blank to take effect on all applications.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// API name.
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// Sampling tag
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Sampling switch. 0: Off; 1: On; 2: Delete
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Configuration ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 0: exact match (default); 1: prefix match; 2: suffix match.
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type ModifyApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// Business system ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sampling rule name.
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// Sampling rate.
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Application name. specifies the application name. fill in the blank to take effect on all applications.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// API name.
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// Sampling tag
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Sampling switch. 0: Off; 1: On; 2: Delete
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Configuration ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 0: exact match (default); 1: prefix match; 2: suffix match.
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *ModifyApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleName")
	delete(f, "SampleRate")
	delete(f, "ServiceName")
	delete(f, "OperationName")
	delete(f, "Tags")
	delete(f, "Status")
	delete(f, "Id")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmSampleConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmSampleConfigResponseParams `json:"Response"`
}

func (r *ModifyApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmServiceRequestParams struct {
	// Application ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// Application description
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// Tag list
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyApmServiceRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// Application description
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// Tag list
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyApmServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceID")
	delete(f, "ServiceDescription")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmServiceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmServiceResponseParams `json:"Response"`
}

func (r *ModifyApmServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmServiceResponse) FromJsonString(s string) error {
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

type Position struct {
	// Node horizontal coordinate
	// Note: This field may return null, indicating that no valid values can be obtained.
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// Node vertical coordinate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type QueryMetricItem struct {
	// Metric name.
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Year-Over-Year comparison is now supported for comparebyyesterday (compared to yesterday) and comparebylastweek (compared to last week). 
	Compares []*string `json:"Compares,omitnil,omitempty" name:"Compares"`

	// Year-On-Year, deprecated, not recommended for use.
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`
}

type Resource struct {
	// Resource type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// TKE resource layer information.
	TKEMeta []*TkeMeta `json:"TKEMeta,omitnil,omitempty" name:"TKEMeta"`

	// CVM resource information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CVMMeta []*CVMMeta `json:"CVMMeta,omitnil,omitempty" name:"CVMMeta"`
}

type SelectorView struct {
	// Component Count
	// Note: This field may return null, indicating that no valid values can be obtained.
	Component *ComponentTopologyView `json:"Component,omitnil,omitempty" name:"Component"`
}

type Selectors struct {
	// Selection status of the component.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Component []*string `json:"Component,omitnil,omitempty" name:"Component"`
}

type ServiceDetail struct {
	// <p>Application ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// <p>Business system ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// <p>User appid</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>main account uin</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateUIN *string `json:"CreateUIN,omitnil,omitempty" name:"CreateUIN"`

	// <p>Application name</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// <p>Application description</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// <p>Region</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Tag</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Business system name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Threshold configuration switch. true means use application level threshold; false means use business system level threshold.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableThresholdConfig *bool `json:"EnableThresholdConfig,omitnil,omitempty" name:"EnableThresholdConfig"`

	// <p>Error rate threshold (%) used to judge the application health status as "red".</p><p>Unit: %</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// <p>Alert threshold for response time (ms), used to judge application health status as "yellow".</p><p>Unit: ms</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`
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

type TkeMeta struct {
	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Cluster ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// pod name
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// workload
	Deployment *string `json:"Deployment,omitnil,omitempty" name:"Deployment"`

	// pod ip
	PodIP *string `json:"PodIP,omitnil,omitempty" name:"PodIP"`

	// node ip
	NodeIP *string `json:"NodeIP,omitnil,omitempty" name:"NodeIP"`
}

type TopologyEdgeNew struct {
	// Source node
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Edge ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Edge weight
	// Note: This field may return null, indicating that no valid values can be obtained.
	Weight *float64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Target node
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// response time
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Error rate
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrRate *float64 `json:"ErrRate,omitnil,omitempty" name:"ErrRate"`

	// throughput
	// Note: This field may return null, indicating that no valid values can be obtained.
	Qps *float64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// Edge type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Edge color
	// Note: This field may return null, indicating that no valid values can be obtained.
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// SQL call count
	// Note: This field may return null, indicating that no valid values can be obtained.
	SqlRequestCount *float64 `json:"SqlRequestCount,omitnil,omitempty" name:"SqlRequestCount"`

	// SQL call error count
	// Note: This field may return null, indicating that no valid values can be obtained.
	SqlErrorRequestCount *float64 `json:"SqlErrorRequestCount,omitnil,omitempty" name:"SqlErrorRequestCount"`

	// Source node type on the edge of the topology diagram. Application/MQ/DB.
	SourceComp *string `json:"SourceComp,omitnil,omitempty" name:"SourceComp"`

	// Target node type on the edge of the topology diagram. Application/MQ/DB.
	TargetComp *string `json:"TargetComp,omitnil,omitempty" name:"TargetComp"`
}

type TopologyNode struct {
	// Error rate
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrRate *float64 `json:"ErrRate,omitnil,omitempty" name:"ErrRate"`

	// Node type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Node name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Node weight
	// Note: This field may return null, indicating that no valid values can be obtained.
	Weight *float64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Node color
	// Note: This field may return null, indicating that no valid values can be obtained.
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// response time
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// throughput
	// Note: This field may return null, indicating that no valid values can be obtained.
	Qps *float64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// Node type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Node ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Node size
	// Note: This field may return null, indicating that no valid values can be obtained.
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// Indicate whether the node is a component
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsModule *bool `json:"IsModule,omitnil,omitempty" name:"IsModule"`

	// Node location information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Position *Position `json:"Position,omitnil,omitempty" name:"Position"`

	// Node tags
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether the node supports drill-down.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CanDrillDown *bool `json:"CanDrillDown,omitnil,omitempty" name:"CanDrillDown"`

	// Resource layer information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Resource *Resource `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Name of the topology node view.
	NodeView *string `json:"NodeView,omitnil,omitempty" name:"NodeView"`

	// Message consumption time of the MQ consumer, in ms.
	ConsumerDuration *float64 `json:"ConsumerDuration,omitnil,omitempty" name:"ConsumerDuration"`

	// Error rate of the MQ consumers, in %.
	ConsumerErrRate *float64 `json:"ConsumerErrRate,omitnil,omitempty" name:"ConsumerErrRate"`

	// Throughput of the message queue (MQ) consumer.
	ConsumerQps *float64 `json:"ConsumerQps,omitnil,omitempty" name:"ConsumerQps"`

	// Application ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}