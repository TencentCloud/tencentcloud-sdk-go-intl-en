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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type BindNewLVBDomainWithChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// The LVB domain name to bind
	LVBDomain *string `json:"LVBDomain,omitempty" name:"LVBDomain"`
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

type BindNewLVBDomainWithChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The LVB domain name bound successfully
		LVBDomain *string `json:"LVBDomain,omitempty" name:"LVBDomain"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Info []*CacheInfoInfo `json:"Info,omitempty" name:"Info"`
}

type CacheInfoInfo struct {

	// Timeout period (ms), which must be an integer multiple of 1000
	// .m3u8/.mpd: [1000, 60000]
	// .ts/.m4s/.mp4: [10000, 1800000]
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// File extension. Valid values: .m3u8, .ts, .mpd, .m4s, .mp4
	// Note: this field may return `null`, indicating that no valid value was found.
	Ext *string `json:"Ext,omitempty" name:"Ext"`
}

type ChannelInfo struct {

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Channel protocol.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Channel input and output.
	Points *PointInfo `json:"Points,omitempty" name:"Points"`

	// Cache configuration
	// Note: this field may return `null`, indicating that no valid value was found.
	CacheInfo *CacheInfo `json:"CacheInfo,omitempty" name:"CacheInfo"`
}

type CreateStreamPackageChannelEndpointRequest struct {
	*tchttp.BaseRequest

	// Channel ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Authentication information
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
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

type CreateStreamPackageChannelEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of the created channel endpoint
		Info *EndpointInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateStreamPackageChannelRequest struct {
	*tchttp.BaseRequest

	// Channel name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Channel protocol. Valid values: HLS, DASH
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Cache configuration
	CacheInfo *CacheInfo `json:"CacheInfo,omitempty" name:"CacheInfo"`
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

type CreateStreamPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel information
		Info *ChannelInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteStreamPackageChannelEndpointsRequest struct {
	*tchttp.BaseRequest

	// Channel ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// List of the URLs of the endpoints to delete
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
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

type DeleteStreamPackageChannelEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteStreamPackageChannelsRequest struct {
	*tchttp.BaseRequest

	// List of the IDs of the channels to delete
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
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

type DeleteStreamPackageChannelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of the information of successfully deleted channels
		SuccessInfos []*ChannelInfo `json:"SuccessInfos,omitempty" name:"SuccessInfos"`

		// List of the information of the channels that failed to be deleted
		FailInfos []*ChannelInfo `json:"FailInfos,omitempty" name:"FailInfos"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeStreamPackageChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID
	Id *string `json:"Id,omitempty" name:"Id"`
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

type DescribeStreamPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel information
		Info *ChannelInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeStreamPackageChannelsRequest struct {
	*tchttp.BaseRequest

	// Page number. Value range: [1, 1000]
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Value range: [1, 1000]
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
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

type DescribeStreamPackageChannelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of channel information
	// Note: this field may return `null`, indicating that no valid value was found.
		Infos []*ChannelInfo `json:"Infos,omitempty" name:"Infos"`

		// Page number
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number of entries
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type EndpointAuthInfo struct {

	// The security group allowlist in CIDR format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList"`

	// The security group blocklist in CIDR format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList"`

	// The authentication key. Its value is same as `X-TENCENT-PACKAGE` set in the HTTP request header.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
}

type EndpointInfo struct {

	// Endpoint name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Endpoint URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Endpoint authentication information.
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
}

type InputAuthInfo struct {

	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Password.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitempty" name:"Password"`
}

type InputInfo struct {

	// Channel input URL.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Channel input authentication information.
	AuthInfo *InputAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
}

type ModifyStreamPackageChannelEndpointRequest struct {
	*tchttp.BaseRequest

	// Channel ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel endpoint URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// New endpoint name
	Name *string `json:"Name,omitempty" name:"Name"`

	// New channel authentication information
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
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

type ModifyStreamPackageChannelEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyStreamPackageChannelInputAuthInfoRequest struct {
	*tchttp.BaseRequest

	// Channel ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel input URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// Authentication configuration. Valid values: `CLOSE`, `UPDATE`
	// `CLOSE`: disable authentication
	// `UPDATE`: update authentication information
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
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

type ModifyStreamPackageChannelInputAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel input authentication information
		AuthInfo *InputAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyStreamPackageChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// New channel name
	Name *string `json:"Name,omitempty" name:"Name"`

	// New channel protocol. Valid values: HLS, DASH
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Cache configuration
	CacheInfo *CacheInfo `json:"CacheInfo,omitempty" name:"CacheInfo"`
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

type ModifyStreamPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Inputs []*InputInfo `json:"Inputs,omitempty" name:"Inputs"`

	// Channel output list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Endpoints []*EndpointInfo `json:"Endpoints,omitempty" name:"Endpoints"`
}
