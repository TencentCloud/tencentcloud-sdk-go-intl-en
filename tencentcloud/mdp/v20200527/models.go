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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ChannelInfo struct {

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Channel protocol.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Channel input and output.
	Points *PointInfo `json:"Points,omitempty" name:"Points"`
}

type CreateMediaPackageChannelEndpointRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Authentication information.
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
}

func (r *CreateMediaPackageChannelEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaPackageChannelEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaPackageChannelEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The information of the created channel endpoint.
		Info *EndpointInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaPackageChannelEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaPackageChannelEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaPackageChannelRequest struct {
	*tchttp.BaseRequest

	// Channel name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Channel protocol. Valid values: HLS, DASH.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *CreateMediaPackageChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaPackageChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel information.
		Info *ChannelInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaPackageChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaPackageChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaPackageChannelEndpointsRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// The list of endpoint URLs.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *DeleteMediaPackageChannelEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaPackageChannelEndpointsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaPackageChannelEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaPackageChannelEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaPackageChannelEndpointsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaPackageChannelsRequest struct {
	*tchttp.BaseRequest

	// The ID list of channels to be deleted.
	Ids []*string `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteMediaPackageChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaPackageChannelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaPackageChannelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The information list of channels that have been deleted.
		SuccessInfos []*ChannelInfo `json:"SuccessInfos,omitempty" name:"SuccessInfos" list`

		// The information list of channels that failed to be deleted.
		FailInfos []*ChannelInfo `json:"FailInfos,omitempty" name:"FailInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaPackageChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaPackageChannelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaPackageChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeMediaPackageChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaPackageChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel information.
		Info *ChannelInfo `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaPackageChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaPackageChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaPackageChannelsRequest struct {
	*tchttp.BaseRequest

	// Page number. Value range: [1, 1000].
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// The size of each page. Value range: [1, 1000].
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeMediaPackageChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaPackageChannelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaPackageChannelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The list of channel outputs.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Infos []*ChannelInfo `json:"Infos,omitempty" name:"Infos" list`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// The size of each page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaPackageChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaPackageChannelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EndpointAuthInfo struct {

	// The security group whitelist in CIDR format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList" list`

	// The security group blacklist in CIDR format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList" list`

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

type ModifyMediaPackageChannelEndpointRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel endpoint URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// The channel name after modification.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The channel authentication after modification.
	AuthInfo *EndpointAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
}

func (r *ModifyMediaPackageChannelEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaPackageChannelEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaPackageChannelEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaPackageChannelEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaPackageChannelEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaPackageChannelInputAuthInfoRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Channel input URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Authentication configuration type. Valid values: CLOSE, UPDATE.
	// CLOSE: disable authentication.
	// UPDATE: update authentication.
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *ModifyMediaPackageChannelInputAuthInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaPackageChannelInputAuthInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaPackageChannelInputAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Channel input authentication information.
		AuthInfo *InputAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaPackageChannelInputAuthInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaPackageChannelInputAuthInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaPackageChannelRequest struct {
	*tchttp.BaseRequest

	// Channel ID.
	Id *string `json:"Id,omitempty" name:"Id"`

	// The channel name after modification.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The channel protocol after modification. Valid values: HLS, DASH.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *ModifyMediaPackageChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaPackageChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaPackageChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaPackageChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaPackageChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PointInfo struct {

	// Channel input list.
	Inputs []*InputInfo `json:"Inputs,omitempty" name:"Inputs" list`

	// Channel output list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Endpoints []*EndpointInfo `json:"Endpoints,omitempty" name:"Endpoints" list`
}
