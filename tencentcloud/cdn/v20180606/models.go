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

package v20180606

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CdnData struct {

	// Queries the specified metric:
	// flux: traffic (in bytes)
	// bandwidth: bandwidth (in bps)
	// request: number of requests
	// fluxHitRate: traffic hit rate (in %)
	// statusCode: status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx status codes will be returned (in entries)
	// 2XX: Returns the aggregate list of 2xx status codes and the data for status codes starting with 2 (in entries)
	// 3XX: Returns the aggregate list of 3xx status codes and the data for status codes starting with 3 (in entries)
	// 4XX: Returns the aggregate list of 4xx status codes and the data for status codes starting with 4 (in entries)
	// 5XX: Returns the aggregate list of 5xx status codes and the data for status codes starting with 5 (in entries)
	// Alternatively, you can specify a status code for querying.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Detailed data combination
	DetailData []*TimestampData `json:"DetailData,omitempty" name:"DetailData"`

	// Aggregate data combination
	SummarizedData *SummarizedData `json:"SummarizedData,omitempty" name:"SummarizedData"`
}

type CreateScdnFailedLogTaskRequest struct {
	*tchttp.BaseRequest

	// ID of the failed task to retry
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Region. Valid values: `mainland` and `overseas`.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateScdnFailedLogTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnFailedLogTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScdnFailedLogTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScdnFailedLogTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Creation result. 
	// 0: Creation succeeded
		Result *string `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScdnFailedLogTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnFailedLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest

	// Query start time, e.g., 2018-09-04 10:40:00. The returned result will be later than or equal to the specified time
	// The time will be rounded forward based on the granularity parameter `Interval`. For example, if the query start time is 2018-09-04 10:40:00 and the query time granularity is 1-hour, the time for the first returned entry will be 2018-09-04 10:00:00
	// The range between the start time and end time should be less than or equal to 90 days
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time, e.g. 2018-09-04 10:40:00. The returned result will be earlier than or equal to the specified time
	// The time will be rounded forward based on the granularity parameter `Interval`. For example, if the query end time is 2018-09-04 10:40:00 and the query time granularity is 1-hour, the time for the last returned entry will be 2018-09-04 10:00:00
	// The range between the start time and end time should be less than or equal to 90 days
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Time granularity, which can be:
	// `min`: 1-minute granularity. The query period cannot exceed 24 hours.
	// `5min`: 5-minute granularity. The query range cannot exceed 31 days.
	// `hour`: 1-hour granularity. The query period cannot exceed 31 days.
	// `day`: 1-day granularity. The query period cannot exceed 31 days.
	// 
	// Querying 1-minute granularity data is not supported if the `Area` field is `overseas`.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Domain name whose billing data is to be queried
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Project ID, which can be viewed [here](https://console.cloud.tencent.com/project)
	// If the `Domain` parameter is populated with specific domain name information, then the billing data of this domain name instead of the specified project will be returned
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Acceleration region whose billing data is to be queried:
	// mainland: in the mainland of China
	// overseas: outside the mainland of China
	// If this parameter is left empty, `mainland` will be used by default
	Area *string `json:"Area,omitempty" name:"Area"`

	// Country/region to be queried if `Area` is `overseas`
	// For district or country/region codes, please see [District Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E7.9C.81.E4.BB.BD.E6.98.A0.E5.B0.84)
	// If this parameter is left empty, all countries/regions will be queried
	District *int64 `json:"District,omitempty" name:"District"`

	// Billing statistics type
	// flux: bill-by-traffic
	// bandwidth: bill-by-bandwidth
	// Default value: `bandwidth`
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeBillingDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	delete(f, "Domain")
	delete(f, "Project")
	delete(f, "Area")
	delete(f, "District")
	delete(f, "Metric")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time granularity, which is specified by the parameter passed in during the query:
	// min: 1-minute
	// 5min: 5-minute
	// hour: 1-hour
	// day: 1-day
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// Data details
		Data []*ResourceBillingData `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDataRequest struct {
	*tchttp.BaseRequest

	// Queries start time, such as 2018-09-04 10:40:00; the returned result is later than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query end time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the first returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Queries end time, such as 2018-09-04 10:40:00; the returned result is earlier than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the last returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the metric to query, which can be:
	// `flux`: traffic (in bytes)
	// `fluxIn`: upstream traffic (in bytes), only used for the `ecdn` product
	// `fluxOut`: downstream traffic (in bytes), only used for the `ecdn` product
	// `bandwidth`: bandwidth (in bps)
	// `bandwidthIn`: upstream bandwidth (in bps), only used for the `ecdn` product
	// `bandwidthOut`: downstream bandwidth (in bps), only used for the `ecdn` product
	// `request`: number of requests
	// `hitRequest`: number of hit requests
	// `requestHitRate`: request hit rate (in % with two decimal digits)
	// `hitFlux`: hit traffic (in bytes)
	// `fluxHitRate`: traffic hit rate (in % with two decimal digits)
	// `statusCode`: status code. Number of 2xx, 3xx, 4xx, and 5xx status codes returned during the queried period.
	// `2xx`: lists the number of all status codes starting with **2** returned during the queried period based on the specified interval (if any)
	// `3xx`: lists the number of all status codes starting with **3** returned during the queried period based on the specified interval (if any)
	// `4xx`: lists the number of all status codes starting with **4** returned during the queried period based on the specified interval (if any)
	// `5xx`: lists the number of all status codes starting with **5** returned during the queried period based on the specified interval (if any)
	// Specifies the status code to query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Queries the information of specified domain names
	// Specifies a domain name to query
	// Specifies multiple domain names to query (30 at most at a time)
	// Queries all Specifies an account to query all domain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// Please note that if domain names are specified, this parameter will be ignored.
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Time granularity; valid values:
	// `min`: data with 1-minute granularity is returned when the queried period is no longer than 24 hours. This value is not supported if the service region you want to query is outside Mainland China;
	// `5min`: data with 5-minute granularity is returned when the queried period is no longer than 31 days;
	// `hour`: data with 1-hour granularity is returned when the queried period is no longer than 31 days;
	// `day`: data with 1-day granularity is returned when the queried period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) during a multi-domain-name query.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported)
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies an ISP when you query the CDN data within Mainland China. If this is left blank, all ISPs will be queried.
	// To view ISP codes, see [ISP Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// If you have specified an ISP, you cannot specify a province or an IP protocol for the same query.
	Isp *int64 `json:"Isp,omitempty" name:"Isp"`

	// Specifies a province when you query the CDN data within Mainland China. If this is left blank, all provinces will be queried.
	// Specifies a country/region when you query the CDN data outside Mainland China. If this is left blank, all countries/regions will be queried.
	// To view codes of provinces or countries/regions, see [Province Code Mappings](https://intl.cloud.tencent.com/document/product/228/6316?from_cn_redirect=1#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// If you have specified a province for your query on CDN data within mainland China, you cannot specify an ISP or an IP protocol for the same query.
	District *int64 `json:"District,omitempty" name:"District"`

	// Specifies the protocol to be queried; if you leave it blank, all protocols will be queried.
	// all: All protocols
	// http: specifies the HTTP metric to be queried
	// https: specifies the HTTPS metric to be queried
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Specifies the data source to be queried, which can be seen as the allowlist function.
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Specified IP protocol to be queried. If this parameter is left empty, all protocols will be queried
	// all: all protocols
	// ipv4: specifies to query IPv4 metrics
	// ipv6: specifies to query IPv6 metrics
	// If the IP protocol to be queried is specified, the district and ISP cannot be specified at the same time
	// Note: non-IPv6 allowlisted users cannot specify `ipv4` and `ipv6` for query
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Specifies a service region. If this value is left blank, CDN data within Mainland China will be queried.
	// `mainland`: specifies to query CDN data within Mainland China;
	// `overseas`: specifies to query CDN data outside Mainland China.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Specifies a region type for your query on CDN data outside Mainland China. If this parameter is left blank, data on the service region will be queried. This parameter is valid only when `Area` is `overseas`.
	// `server`: specifies to query data on the service region where Tencent Cloud CDN nodes are located;
	// `client`: specifies to query data on the client region where the request devices are located.
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`

	// Specifies the product to query, either `cdn` (default) or `ecdn`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeCdnDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	delete(f, "Detail")
	delete(f, "Isp")
	delete(f, "District")
	delete(f, "Protocol")
	delete(f, "DataSource")
	delete(f, "IpProtocol")
	delete(f, "Area")
	delete(f, "AreaType")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time granularity of the returned data. Specify one of the following during querying:
	// min: 1 minute
	// 5min: 5 minutes
	// hour: 1 hour
	// day: 1 day
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// Returned data details of the specified conditional query
		Data []*ResourceData `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataRequest struct {
	*tchttp.BaseRequest

	// Query start time, such as 2018-09-04 10:40:00; the returned result is later than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query end time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the first returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time, such as 2018-09-04 10:40:00; the returned result is earlier than or equal to the specified time.
	// According to the specified time granularity, forward rounding is applied; for example, if the query start time is 2018-09-04 10:40:00 and the query time granularity is 1 hour, the time for the last returned entry will be 2018-09-04 10:00:00.
	// The gap between the start time and end time should be less than or equal to 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the query metric, which can be:
	// flux: origin-pull traffic (in bytes)
	// bandwidth: origin-pull bandwidth (in bps)
	// request: number of origin-pull requests
	// failRequest: number of failed origin-pull requests
	// failRate: origin-pull failure rate (in %)
	// statusCode: origin-pull status code. The aggregate data for 2xx, 3xx, 4xx, and 5xx origin-pull status codes will be returned (in entries)
	// 2xx: Returns the aggregate list of 2xx origin-pull status codes and the data for origin-pull status codes starting with 2 (in entries)
	// 3xx: Returns the aggregate list of 3xx origin-pull status codes and the data for origin-pull status codes starting with 3 (in entries)
	// 4xx: Returns the aggregate list of 4xx origin-pull status codes and the data for origin-pull status codes starting with 4 (in entries)
	// 5xx: Returns the aggregate list of 5xx origin-pull status codes and the data for origin-pull status codes starting with 5 (in entries)
	// It is supported to specify a status code for query. The return will be empty if the status code has never been generated.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Specifies the list of domain names to be queried; up to 30 domain names can be queried at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Project ID, which can be viewed [here](https://console.cloud.tencent.com/project)
	// If the domain name is not specified, the specified project will be queried. Up to 30 acceleration domain names can be queried at a time
	// If the domain name information is specified, the domain name will prevail
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// Time granularity; valid values:
	// `min`: data with 1-minute granularity is returned when the queried period is no longer than 24 hours. This value is not supported if the service region you want to query is outside Mainland China;
	// `5min`: data with 5-minute granularity is returned when the queried period is no longer than 31 days;
	// `hour`: data with 1-hour granularity is returned when the queried period is no longer than 31 days;
	// `day`: data with 1-day granularity is returned when the queried period is longer than 31 days.
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The aggregate data for multiple domain names is returned by default (false) when multiple `Domains` are passed in.
	// You can set it to true to return the details for each Domain (the statusCode metric is currently not supported)
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Specifies a service region. If this value is left blank, CDN data within Mainland China will be queried.
	// `mainland`: specifies to query CDN data within Mainland China;
	// `overseas`: specifies to query CDN data outside Mainland China.
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeOriginDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	delete(f, "Detail")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time granularity of data statistics, which supports min (1 minute), 5min (5 minutes), hour (1 hour), and day (1 day).
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// Origin-pull data details of each resource.
		Data []*ResourceOriginData `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOriginDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest

	// List of directories. The protocol header such as "http://" or "https://" needs to be included.
	Paths []*string `json:"Paths,omitempty" name:"Paths"`

	// Purge type:
	// `flush`: purges updated resources
	// `delete`: purges all resources
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// Whether to encode Chinese characters before purge.
	UrlEncode *bool `json:"UrlEncode,omitempty" name:"UrlEncode"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Paths")
	delete(f, "FlushType")
	delete(f, "UrlEncode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurgePathCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Purge task ID. Directories submitted in one request share a task ID.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PushUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// List of URLs. The protocol header such as "http://" or "https://" needs to be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// Specifies the User-Agent header of an HTTP prefetch request when it is forwarded to the origin server
	// Default value: `TencentCdn`
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Destination region for the prefetch
	// `mainland`: prefetches resources to nodes within Mainland China
	// `overseas`: prefetches resources to nodes outside Mainland China
	// `global`: prefetches resources to global nodes
	// Default value: `mainland`. You can prefetch a URL to nodes in a region provided that CDN service has been enabled for the domain name in the URL in the region.
	Area *string `json:"Area,omitempty" name:"Area"`

	// If this parameter is `middle` or left empty, prefetch will be performed onto the intermediate node.
	// Note: resources prefetched outside the Chinese mainland will be cached to CDN nodes outside the Chinese mainland and the traffic generated will incur costs.
	Layer *string `json:"Layer,omitempty" name:"Layer"`

	// Whether to recursively resolve the M3U8 index file and prefetch the TS shards in it.
	// Notes:
	// 1. This feature requires that the M3U8 index file can be directly requested and obtained.
	// 2. In the M3U8 index file, currently only the TS shards at the first to the third level can be recursively resolved.
	// 3. Prefetching the TS shards obtained through recursive resolution consumes the daily prefetch quota. If the usage exceeds the quota, the feature will be disabled and TS shards will not be prefetched.
	ParseM3U8 *bool `json:"ParseM3U8,omitempty" name:"ParseM3U8"`

	// Specifies whether to disable Range GETs.
	// Notes:
	// This feature is in beta test.
	DisableRange *bool `json:"DisableRange,omitempty" name:"DisableRange"`
}

func (r *PushUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlsCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "UserAgent")
	delete(f, "Area")
	delete(f, "Layer")
	delete(f, "ParseM3U8")
	delete(f, "DisableRange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushUrlsCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PushUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the submitted task
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PushUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlsCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceBillingData struct {

	// Resource name, which is categorized as follows based on different query conditions:
	// Specific domain name: domain name details
	// multiDomains: aggregated details of multiple domain names
	// Project ID: displays the ID of the specified project to be queried
	// all: the details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Billing data details
	BillingData []*CdnData `json:"BillingData,omitempty" name:"BillingData"`
}

type ResourceData struct {

	// Resource name, which is classified as follows based on different query filters:
	// A single domain name: queries domain name details by a domain name. The details of the domain name will be displayed when the passed parameter `detail` is `true` (the `detail` parameter defaults to `false`).
	// Multiple domain names: queries domain name details by multiple domain names. The aggregated details of the domain names will be displayed.
	// Project ID: queries domain name details by a project ID. The aggregated details of the domain names of the project will be displayed.
	// `all`: account-level data, which is aggregated details of all domain names of an account.
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Data details of a resource
	CdnData []*CdnData `json:"CdnData,omitempty" name:"CdnData"`
}

type ResourceOriginData struct {

	// Resource name, which is classified as follows based on different query conditions:
	// A specific domain name: This indicates the details of this domain name
	// multiDomains: This indicates the aggregate details of multiple domain names
	// Project ID: This displays the ID of the specifically queried project
	// all: This indicates the details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Origin-pull data details
	OriginData []*CdnData `json:"OriginData,omitempty" name:"OriginData"`
}

type SummarizedData struct {

	// Aggregation method, which can be:
	// sum: aggregate summation
	// max: maximum value; in bandwidth mode, the peak bandwidth is calculated based on the aggregate data with 5-minute granularity.
	// avg: average value
	Name *string `json:"Name,omitempty" name:"Name"`

	// Aggregate data value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TimestampData struct {

	// Statistical point in time in forward rounding mode
	// Taking the 5-minute granularity as an example, 13:35:00 indicates that the statistical interval is between 13:35:00 and 13:39:59.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Data value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}
