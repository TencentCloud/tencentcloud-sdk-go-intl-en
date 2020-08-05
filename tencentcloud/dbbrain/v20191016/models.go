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

package v20191016

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Event ID, which can be obtained through the `DescribeDBDiagHistory` API.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeDBDiagEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Diagnosis item.
		DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

		// Diagnosis type.
		DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

		// Event ID.
		EventId *int64 `json:"EventId,omitempty" name:"EventId"`

		// Event details.
		Explanation *string `json:"Explanation,omitempty" name:"Explanation"`

		// Summary.
		Outline *string `json:"Outline,omitempty" name:"Outline"`

		// Problem found.
		Problem *string `json:"Problem,omitempty" name:"Problem"`

		// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
		Severity *int64 `json:"Severity,omitempty" name:"Severity"`

		// Start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Suggestion.
		Suggestions *string `json:"Suggestions,omitempty" name:"Suggestions"`

		// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Metric *string `json:"Metric,omitempty" name:"Metric"`

		// End time.
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDiagEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as "2019-09-11 12:13:14".
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDBDiagHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Event description.
		Events []*DiagHistoryEventItem `json:"Events,omitempty" name:"Events" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDiagHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Query period in days. The end date is the current date and the query period is 7 days by default.
	RangeDays *int64 `json:"RangeDays,omitempty" name:"RangeDays"`
}

func (r *DescribeDBSpaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSpaceStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Disk usage growth in MB.
		Growth *int64 `json:"Growth,omitempty" name:"Growth"`

		// Available disk space in MB.
		Remain *int64 `json:"Remain,omitempty" name:"Remain"`

		// Total disk space in MB.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Estimated number of available days.
		AvailableDays *int64 `json:"AvailableDays,omitempty" name:"AvailableDays"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSpaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSpaceStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time range in seconds in histogram.
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// Number of slow logs in specified time range.
		TimeSeries []*TimeSlice `json:"TimeSeries,omitempty" name:"TimeSeries" list`

		// Instance CPU utilization monitoring data in specified time range.
		SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Sorting key. Valid values: QueryTime, ExecTimes, RowsSent, LockTime, RowsExamined.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTopSqlsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of top slow SQL statements
		Rows []*SlowLogTopSqlItem `json:"Rows,omitempty" name:"Rows" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTopSqlsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top tables. Default value: 20. Maximum value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize. Default value: PhysicalFileSize.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Start date. It can be as early as 6 days before the current date, and defaults to 6 days before the end date.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// End date. It can be as early as 6 days before the current date, and defaults to the current date.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time series list of the returned space statistics of top tables.
		TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitempty" name:"TopSpaceTableTimeSeries" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top tables. Default value: 20. Maximum value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top tables. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize. Default value: PhysicalFileSize.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeTopSpaceTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of the returned space statistics of top tables.
		TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitempty" name:"TopSpaceTables" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {

	// Diagnosis type.
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// Summary.
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// Diagnosis item.
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// Instance ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Reserved field
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Region
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`
}

type MonitorFloatMetric struct {

	// Metric name.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Metric value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Values []*float64 `json:"Values,omitempty" name:"Values" list`
}

type MonitorFloatMetricSeriesData struct {

	// Monitoring metric.
	Series []*MonitorFloatMetric `json:"Series,omitempty" name:"Series" list`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp" list`
}

type MonitorMetric struct {

	// Metric name.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Metric value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Values []*int64 `json:"Values,omitempty" name:"Values" list`
}

type MonitorMetricSeriesData struct {

	// Monitoring metric.
	Series []*MonitorMetric `json:"Series,omitempty" name:"Series" list`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp" list`
}

type SlowLogTopSqlItem struct {

	// Total SQL lock wait time
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// Maximum lock wait time
	LockTimeMax *float64 `json:"LockTimeMax,omitempty" name:"LockTimeMax"`

	// Minimum lock wait time
	LockTimeMin *float64 `json:"LockTimeMin,omitempty" name:"LockTimeMin"`

	// Total number of scanned rows
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// Maximum number of scanned rows
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitempty" name:"RowsExaminedMax"`

	// Minimum number of scanned rows
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitempty" name:"RowsExaminedMin"`

	// Total duration
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// Maximum execution time
	QueryTimeMax *float64 `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`

	// Minimum execution time
	QueryTimeMin *float64 `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`

	// Total number of returned rows
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// Maximum number of returned rows
	RowsSentMax *int64 `json:"RowsSentMax,omitempty" name:"RowsSentMax"`

	// Minimum number of returned rows
	RowsSentMin *int64 `json:"RowsSentMin,omitempty" name:"RowsSentMin"`

	// Number of executions
	ExecTimes *int64 `json:"ExecTimes,omitempty" name:"ExecTimes"`

	// SQL template
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// SQL with parameter (random)
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Schema
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// Ratio of total duration
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitempty" name:"QueryTimeRatio"`

	// Ratio of total SQL lock wait time
	LockTimeRatio *float64 `json:"LockTimeRatio,omitempty" name:"LockTimeRatio"`

	// Ratio of total number of scanned rows
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitempty" name:"RowsExaminedRatio"`

	// Ratio of total number of returned rows
	RowsSentRatio *float64 `json:"RowsSentRatio,omitempty" name:"RowsSentRatio"`
}

type TableSpaceData struct {

	// Table name.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitempty" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitempty" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitempty" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`

	// Fragmented rate (%).
	FragRatio *float64 `json:"FragRatio,omitempty" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// Size in MB of the physical file exclusive to a table.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitempty" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {

	// Table name.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Monitoring metric data in a unit of time interval.
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`
}

type TimeSlice struct {

	// Total number
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Statistics start time
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}
