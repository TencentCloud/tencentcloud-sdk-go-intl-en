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

type DataPoint struct {

	// Combination of instance object dimensions
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// The array of timestamps indicating at which points in time there is data. Missing timestamps have no data points (i.e., missed)
	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps" list`

	// The array of monitoring values, which is in one-to-one correspondence to Timestamps
	Values []*float64 `json:"Values,omitempty" name:"Values" list`
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest

	// Business namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Metric name
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

type PeriodsSt struct {

	// Period
	Period *string `json:"Period,omitempty" name:"Period"`

	// Statistical method
	StatType []*string `json:"StatType,omitempty" name:"StatType" list`
}
