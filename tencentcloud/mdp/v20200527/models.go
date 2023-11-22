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

package v20200527

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type BindNewLVBDomainWithChannelRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// The LVB domain name to bind
	LVBDomain *string `json:"LVBDomain,omitnil" name:"LVBDomain"`
}

type BindNewLVBDomainWithChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// The LVB domain name to bind
	LVBDomain *string `json:"LVBDomain,omitnil" name:"LVBDomain"`
}

func (r *BindNewLVBDomainWithChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNewLVBDomainWithChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "LVBDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindNewLVBDomainWithChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindNewLVBDomainWithChannelResponseParams struct {
	// The LVB domain name bound successfully
	LVBDomain *string `json:"LVBDomain,omitnil" name:"LVBDomain"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BindNewLVBDomainWithChannelResponse struct {
	*tchttp.BaseResponse
	Response *BindNewLVBDomainWithChannelResponseParams `json:"Response"`
}

func (r *BindNewLVBDomainWithChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindNewLVBDomainWithChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CacheInfo struct {
	// List of timeout parameter configuration
	// Note: this field may return `null`, indicating that no valid value was found.
	Info []*CacheInfoInfo `json:"Info,omitnil" name:"Info"`
}

type CacheInfoInfo struct {
	// Timeout period (ms), which must be an integer multiple of 1000
	// .m3u8/.mpd: [1000, 60000]
	// .ts/.m4s/.mp4: [10000, 1800000]
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`

	// File extension. Valid values: .m3u8, .ts, .mpd, .m4s, .mp4
	// Note: this field may return `null`, indicating that no valid value was found.
	Ext *string `json:"Ext,omitnil" name:"Ext"`
}

type ChannelInfo struct {
	// Channel ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Channel protocol.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Channel input and output.
	Points *PointInfo `json:"Points,omitnil" name:"Points"`

	// Cache configuration
	// Note: this field may return `null`, indicating that no valid value was found.
	CacheInfo *CacheInfo `json:"CacheInfo,omitnil" name:"CacheInfo"`
}

// Predefined struct for user
type CreateStreamPackageChannelEndpointRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Endpoint name, which must contain 1 to 32 characters and supports digits, letters, and underscores
	Name *string `json:"Name,omitnil" name:"Name"`

	// Authentication information
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`
}

type CreateStreamPackageChannelEndpointRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Endpoint name, which must contain 1 to 32 characters and supports digits, letters, and underscores
	Name *string `json:"Name,omitnil" name:"Name"`

	// Authentication information
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`
}

func (r *CreateStreamPackageChannelEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamPackageChannelEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "AuthInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamPackageChannelEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamPackageChannelEndpointResponseParams struct {
	// Information of the created channel endpoint
	Info *EndpointInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamPackageChannelEndpointResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamPackageChannelEndpointResponseParams `json:"Response"`
}

func (r *CreateStreamPackageChannelEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamPackageChannelEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamPackageChannelRequestParams struct {
	// Channel name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Channel protocol. Valid values: HLS, DASH
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Cache configuration
	CacheInfo *CacheInfo `json:"CacheInfo,omitnil" name:"CacheInfo"`
}

type CreateStreamPackageChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Channel protocol. Valid values: HLS, DASH
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Cache configuration
	CacheInfo *CacheInfo `json:"CacheInfo,omitnil" name:"CacheInfo"`
}

func (r *CreateStreamPackageChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamPackageChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Protocol")
	delete(f, "CacheInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamPackageChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamPackageChannelResponseParams struct {
	// Channel information
	Info *ChannelInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamPackageChannelResponseParams `json:"Response"`
}

func (r *CreateStreamPackageChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamPackageChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamPackageHarvestJobRequestParams struct {
	// Task ID, a globally unique identifier, ID can contain up to 128 characters. Supported characters are numbers, letters, underscores (_), and dashes (-).
	ID *string `json:"ID,omitnil" name:"ID"`

	// The associated channel name.
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// The associated endpoint name, StreamPackage only supports harvesting content from HLS and DASH endpoints that have startover turned on.
	EndpointName *string `json:"EndpointName,omitnil" name:"EndpointName"`

	// Time format, supports the following types: 
	// 1. Epoch seconds 
	// 2. ISO-8601.
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`

	// Task start time supports two formats for TimeFormat input: 
	// 1. Epoch seconds: The input box is a numeric input box, and only positive integers can be entered. 
	// 2. ISO-8601: The supported format is ISO time, for example: 2023-08-01T10:00:00+08:00.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Task end time supports two formats for TimeFormat input: 
	// 1. Epoch seconds: The input box is a numeric input box, and only positive integers can be entered. 
	// 2. ISO-8601: The supported format is ISO time, for example: 2023-08-01T10:00:00+08:00.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The COS bucket address where the recording file is stored in COS, format: https://{Bucket_name}-{AppId}.cos.{Region}.myqcloud.com/.
	Destination *string `json:"Destination,omitnil" name:"Destination"`

	// The path in the cos bucket and identifier for the parent manifest for the live-to-VOD asset, fotmat: {path 1}/xxx/{path n}/{name}.m3u8.
	Manifest *string `json:"Manifest,omitnil" name:"Manifest"`
}

type CreateStreamPackageHarvestJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID, a globally unique identifier, ID can contain up to 128 characters. Supported characters are numbers, letters, underscores (_), and dashes (-).
	ID *string `json:"ID,omitnil" name:"ID"`

	// The associated channel name.
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// The associated endpoint name, StreamPackage only supports harvesting content from HLS and DASH endpoints that have startover turned on.
	EndpointName *string `json:"EndpointName,omitnil" name:"EndpointName"`

	// Time format, supports the following types: 
	// 1. Epoch seconds 
	// 2. ISO-8601.
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`

	// Task start time supports two formats for TimeFormat input: 
	// 1. Epoch seconds: The input box is a numeric input box, and only positive integers can be entered. 
	// 2. ISO-8601: The supported format is ISO time, for example: 2023-08-01T10:00:00+08:00.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Task end time supports two formats for TimeFormat input: 
	// 1. Epoch seconds: The input box is a numeric input box, and only positive integers can be entered. 
	// 2. ISO-8601: The supported format is ISO time, for example: 2023-08-01T10:00:00+08:00.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The COS bucket address where the recording file is stored in COS, format: https://{Bucket_name}-{AppId}.cos.{Region}.myqcloud.com/.
	Destination *string `json:"Destination,omitnil" name:"Destination"`

	// The path in the cos bucket and identifier for the parent manifest for the live-to-VOD asset, fotmat: {path 1}/xxx/{path n}/{name}.m3u8.
	Manifest *string `json:"Manifest,omitnil" name:"Manifest"`
}

func (r *CreateStreamPackageHarvestJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamPackageHarvestJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "ChannelName")
	delete(f, "EndpointName")
	delete(f, "TimeFormat")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Destination")
	delete(f, "Manifest")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamPackageHarvestJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamPackageHarvestJobResponseParams struct {
	// HarvestJob information.
	Info *HarvestJobResp `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStreamPackageHarvestJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamPackageHarvestJobResponseParams `json:"Response"`
}

func (r *CreateStreamPackageHarvestJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamPackageHarvestJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageChannelEndpointsRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// List of the URLs of the endpoints to delete
	Urls []*string `json:"Urls,omitnil" name:"Urls"`
}

type DeleteStreamPackageChannelEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// List of the URLs of the endpoints to delete
	Urls []*string `json:"Urls,omitnil" name:"Urls"`
}

func (r *DeleteStreamPackageChannelEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageChannelEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Urls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamPackageChannelEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageChannelEndpointsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamPackageChannelEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamPackageChannelEndpointsResponseParams `json:"Response"`
}

func (r *DeleteStreamPackageChannelEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageChannelEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageChannelsRequestParams struct {
	// List of the IDs of the channels to delete
	Ids []*string `json:"Ids,omitnil" name:"Ids"`
}

type DeleteStreamPackageChannelsRequest struct {
	*tchttp.BaseRequest
	
	// List of the IDs of the channels to delete
	Ids []*string `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteStreamPackageChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamPackageChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageChannelsResponseParams struct {
	// List of the information of successfully deleted channels
	SuccessInfos []*ChannelInfo `json:"SuccessInfos,omitnil" name:"SuccessInfos"`

	// List of the information of the channels that failed to be deleted
	FailInfos []*ChannelInfo `json:"FailInfos,omitnil" name:"FailInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamPackageChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamPackageChannelsResponseParams `json:"Response"`
}

func (r *DeleteStreamPackageChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageHarvestJobRequestParams struct {
	// Task ID, a globally unique identifier.
	ID *string `json:"ID,omitnil" name:"ID"`
}

type DeleteStreamPackageHarvestJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID, a globally unique identifier.
	ID *string `json:"ID,omitnil" name:"ID"`
}

func (r *DeleteStreamPackageHarvestJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageHarvestJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamPackageHarvestJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageHarvestJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamPackageHarvestJobResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamPackageHarvestJobResponseParams `json:"Response"`
}

func (r *DeleteStreamPackageHarvestJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageHarvestJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageHarvestJobsRequestParams struct {
	// Task IDs, id is a globally unique identifier.
	IDs []*string `json:"IDs,omitnil" name:"IDs"`
}

type DeleteStreamPackageHarvestJobsRequest struct {
	*tchttp.BaseRequest
	
	// Task IDs, id is a globally unique identifier.
	IDs []*string `json:"IDs,omitnil" name:"IDs"`
}

func (r *DeleteStreamPackageHarvestJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageHarvestJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamPackageHarvestJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamPackageHarvestJobsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteStreamPackageHarvestJobsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamPackageHarvestJobsResponseParams `json:"Response"`
}

func (r *DeleteStreamPackageHarvestJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamPackageHarvestJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageChannelRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeStreamPackageChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeStreamPackageChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamPackageChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageChannelResponseParams struct {
	// Channel information
	Info *ChannelInfo `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamPackageChannelResponseParams `json:"Response"`
}

func (r *DescribeStreamPackageChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageChannelsRequestParams struct {
	// Page number. Value range: [1, 1000]
	PageNum *uint64 `json:"PageNum,omitnil" name:"PageNum"`

	// Number of entries per page. Value range: [1, 1000]
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeStreamPackageChannelsRequest struct {
	*tchttp.BaseRequest
	
	// Page number. Value range: [1, 1000]
	PageNum *uint64 `json:"PageNum,omitnil" name:"PageNum"`

	// Number of entries per page. Value range: [1, 1000]
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeStreamPackageChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamPackageChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageChannelsResponseParams struct {
	// List of channel information
	// Note: this field may return `null`, indicating that no valid value was found.
	Infos []*ChannelInfo `json:"Infos,omitnil" name:"Infos"`

	// Page number
	PageNum *uint64 `json:"PageNum,omitnil" name:"PageNum"`

	// Number of entries per page
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`

	// Total number of entries
	TotalNum *uint64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// Total number of pages
	TotalPage *uint64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamPackageChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamPackageChannelsResponseParams `json:"Response"`
}

func (r *DescribeStreamPackageChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageHarvestJobRequestParams struct {
	// Task ID, a globally unique identifier.
	ID *string `json:"ID,omitnil" name:"ID"`
}

type DescribeStreamPackageHarvestJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID, a globally unique identifier.
	ID *string `json:"ID,omitnil" name:"ID"`
}

func (r *DescribeStreamPackageHarvestJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageHarvestJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamPackageHarvestJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageHarvestJobResponseParams struct {
	// HarvestJob information.
	Info *HarvestJobResp `json:"Info,omitnil" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamPackageHarvestJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamPackageHarvestJobResponseParams `json:"Response"`
}

func (r *DescribeStreamPackageHarvestJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageHarvestJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageHarvestJobsRequestParams struct {
	// The bound channel name. If not passed, all channels will be queried by default.
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// Page number.
	PageNum *uint64 `json:"PageNum,omitnil" name:"PageNum"`

	// PageSize.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeStreamPackageHarvestJobsRequest struct {
	*tchttp.BaseRequest
	
	// The bound channel name. If not passed, all channels will be queried by default.
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// Page number.
	PageNum *uint64 `json:"PageNum,omitnil" name:"PageNum"`

	// PageSize.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeStreamPackageHarvestJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageHarvestJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelName")
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamPackageHarvestJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamPackageHarvestJobsResponseParams struct {
	// HarvestJob information list.
	Infos []*HarvestJobResp `json:"Infos,omitnil" name:"Infos"`

	// Page number.
	PageNum *uint64 `json:"PageNum,omitnil" name:"PageNum"`

	// PageSize
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`

	// TotalNum
	TotalNum *uint64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamPackageHarvestJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamPackageHarvestJobsResponseParams `json:"Response"`
}

func (r *DescribeStreamPackageHarvestJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPackageHarvestJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndpointAuthInfo struct {
	// The security group allowlist in CIDR format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WhiteIpList []*string `json:"WhiteIpList,omitnil" name:"WhiteIpList"`

	// The security group blocklist in CIDR format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BlackIpList []*string `json:"BlackIpList,omitnil" name:"BlackIpList"`

	// The authentication key. Its value is same as `X-TENCENT-PACKAGE` set in the HTTP request header.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthKey *string `json:"AuthKey,omitnil" name:"AuthKey"`
}

type EndpointInfo struct {
	// Endpoint name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Endpoint URL.
	Url *string `json:"Url,omitnil" name:"Url"`

	// Endpoint authentication information.
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`
}

type HarvestJobResp struct {
	// Job ID, a globally unique identifier.
	ID *string `json:"ID,omitnil" name:"ID"`

	// The associated channel name.
	ChannelName *string `json:"ChannelName,omitnil" name:"ChannelName"`

	// The associated endpoint name.
	EndpointName *string `json:"EndpointName,omitnil" name:"EndpointName"`

	// Time format, supports the following types: 1. Epoch seconds 2. ISO-8601
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`

	// Job start time.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Job end time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The path where the recording file is stored in COS.
	Destination *string `json:"Destination,omitnil" name:"Destination"`

	// The file name of the recording file stored in COS.
	Manifest *string `json:"Manifest,omitnil" name:"Manifest"`

	// The task status is divided into running: Running, execution completed: Completed, and execution failure: Failed.
	Status *string `json:"Status,omitnil" name:"Status"`

	// Job error message.
	ErrMessage *string `json:"ErrMessage,omitnil" name:"ErrMessage"`

	// Job creation time, timestamp in seconds.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// The associated ChannelID.
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// The region corresponding to the harvest job.
	Region *string `json:"Region,omitnil" name:"Region"`
}

type InputAuthInfo struct {
	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Password.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitnil" name:"Password"`
}

type InputInfo struct {
	// Channel input URL.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil" name:"Url"`

	// Channel input authentication information.
	AuthInfo *InputAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`
}

// Predefined struct for user
type ModifyStreamPackageChannelEndpointRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel endpoint URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// New endpoint name
	Name *string `json:"Name,omitnil" name:"Name"`

	// New channel authentication information
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`
}

type ModifyStreamPackageChannelEndpointRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel endpoint URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// New endpoint name
	Name *string `json:"Name,omitnil" name:"Name"`

	// New channel authentication information
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`
}

func (r *ModifyStreamPackageChannelEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamPackageChannelEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Url")
	delete(f, "Name")
	delete(f, "AuthInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamPackageChannelEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamPackageChannelEndpointResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamPackageChannelEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamPackageChannelEndpointResponseParams `json:"Response"`
}

func (r *ModifyStreamPackageChannelEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamPackageChannelEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamPackageChannelInputAuthInfoRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel input URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// Authentication configuration. Valid values: `CLOSE`, `UPDATE`
	// `CLOSE`: disable authentication
	// `UPDATE`: update authentication information
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`
}

type ModifyStreamPackageChannelInputAuthInfoRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Channel input URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// Authentication configuration. Valid values: `CLOSE`, `UPDATE`
	// `CLOSE`: disable authentication
	// `UPDATE`: update authentication information
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`
}

func (r *ModifyStreamPackageChannelInputAuthInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamPackageChannelInputAuthInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Url")
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamPackageChannelInputAuthInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamPackageChannelInputAuthInfoResponseParams struct {
	// Channel input authentication information
	AuthInfo *InputAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamPackageChannelInputAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamPackageChannelInputAuthInfoResponseParams `json:"Response"`
}

func (r *ModifyStreamPackageChannelInputAuthInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamPackageChannelInputAuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamPackageChannelRequestParams struct {
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// New channel name
	Name *string `json:"Name,omitnil" name:"Name"`

	// New channel protocol. Valid values: HLS, DASH
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Cache configuration
	CacheInfo *CacheInfo `json:"CacheInfo,omitnil" name:"CacheInfo"`
}

type ModifyStreamPackageChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// New channel name
	Name *string `json:"Name,omitnil" name:"Name"`

	// New channel protocol. Valid values: HLS, DASH
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Cache configuration
	CacheInfo *CacheInfo `json:"CacheInfo,omitnil" name:"CacheInfo"`
}

func (r *ModifyStreamPackageChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamPackageChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "Protocol")
	delete(f, "CacheInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamPackageChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamPackageChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyStreamPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamPackageChannelResponseParams `json:"Response"`
}

func (r *ModifyStreamPackageChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamPackageChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PointInfo struct {
	// Channel input list.
	Inputs []*InputInfo `json:"Inputs,omitnil" name:"Inputs"`

	// Channel output list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Endpoints []*EndpointInfo `json:"Endpoints,omitnil" name:"Endpoints"`
}

// Predefined struct for user
type UnbindCdnDomainWithChannelRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// CDN playback domain name
	CdnDomain *string `json:"CdnDomain,omitnil" name:"CdnDomain"`
}

type UnbindCdnDomainWithChannelRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// CDN playback domain name
	CdnDomain *string `json:"CdnDomain,omitnil" name:"CdnDomain"`
}

func (r *UnbindCdnDomainWithChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCdnDomainWithChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "CdnDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindCdnDomainWithChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindCdnDomainWithChannelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnbindCdnDomainWithChannelResponse struct {
	*tchttp.BaseResponse
	Response *UnbindCdnDomainWithChannelResponseParams `json:"Response"`
}

func (r *UnbindCdnDomainWithChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCdnDomainWithChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}