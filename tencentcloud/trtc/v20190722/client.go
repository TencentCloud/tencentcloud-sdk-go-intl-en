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

package v20190722

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-07-22"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDescribeCallDetailRequest() (request *DescribeCallDetailRequest) {
    request = &DescribeCallDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCallDetail")
    return
}

func NewDescribeCallDetailResponse() (response *DescribeCallDetailResponse) {
    response = &DescribeCallDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the user list and user call quality data in a specified time period. It can query data of up to 6 users for the last 5 days, and the query time range cannot exceed 1 hour.
func (c *Client) DescribeCallDetail(request *DescribeCallDetailRequest) (response *DescribeCallDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCallDetailRequest()
    }
    response = NewDescribeCallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryScaleRequest() (request *DescribeHistoryScaleRequest) {
    request = &DescribeHistoryScaleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeHistoryScale")
    return
}

func NewDescribeHistoryScaleResponse() (response *DescribeHistoryScaleResponse) {
    response = &DescribeHistoryScaleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the number of historical rooms and users for the last 5 days. It can query once per minute.
func (c *Client) DescribeHistoryScale(request *DescribeHistoryScaleRequest) (response *DescribeHistoryScaleResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryScaleRequest()
    }
    response = NewDescribeHistoryScaleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealtimeNetworkRequest() (request *DescribeRealtimeNetworkRequest) {
    request = &DescribeRealtimeNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRealtimeNetwork")
    return
}

func NewDescribeRealtimeNetworkResponse() (response *DescribeRealtimeNetworkResponse) {
    response = &DescribeRealtimeNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query real-time network status for the last 24 hours according to `sdkappid`, including upstream and downstream packet losses. The query time range cannot exceed 1 hour.
func (c *Client) DescribeRealtimeNetwork(request *DescribeRealtimeNetworkRequest) (response *DescribeRealtimeNetworkResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeNetworkRequest()
    }
    response = NewDescribeRealtimeNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealtimeQualityRequest() (request *DescribeRealtimeQualityRequest) {
    request = &DescribeRealtimeQualityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRealtimeQuality")
    return
}

func NewDescribeRealtimeQualityResponse() (response *DescribeRealtimeQualityResponse) {
    response = &DescribeRealtimeQualityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query real-time quality data for the last 24 hours according to `sdkappid`, including the room entry success rate, instant playback rate of the first frame, audio lag rate, and video lag rate. The query time range cannot exceed 1 hour.
func (c *Client) DescribeRealtimeQuality(request *DescribeRealtimeQualityRequest) (response *DescribeRealtimeQualityResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeQualityRequest()
    }
    response = NewDescribeRealtimeQualityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealtimeScaleRequest() (request *DescribeRealtimeScaleRequest) {
    request = &DescribeRealtimeScaleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRealtimeScale")
    return
}

func NewDescribeRealtimeScaleResponse() (response *DescribeRealtimeScaleResponse) {
    response = &DescribeRealtimeScaleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the real-time scale for the last 24 hours according to `sdkappid`. The query time range cannot exceed 1 hour.
func (c *Client) DescribeRealtimeScale(request *DescribeRealtimeScaleRequest) (response *DescribeRealtimeScaleResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeScaleRequest()
    }
    response = NewDescribeRealtimeScaleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomInformationRequest() (request *DescribeRoomInformationRequest) {
    request = &DescribeRoomInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRoomInformation")
    return
}

func NewDescribeRoomInformationResponse() (response *DescribeRoomInformationResponse) {
    response = &DescribeRoomInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the room list for the last 5 days according to `sdkappid`. It returns 10 calls by default and up to 100 calls at a time.
func (c *Client) DescribeRoomInformation(request *DescribeRoomInformationRequest) (response *DescribeRoomInformationResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInformationRequest()
    }
    response = NewDescribeRoomInformationResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomRequest() (request *DismissRoomRequest) {
    request = &DismissRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DismissRoom")
    return
}

func NewDismissRoomResponse() (response *DismissRoomResponse) {
    response = &DismissRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to remove all users from a room and dismiss the room. It supports all platforms. For Android, iOS, Windows, and macOS, the TRTC SDK needs to be upgraded to v6.6 or above.
func (c *Client) DismissRoom(request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    if request == nil {
        request = NewDismissRoomRequest()
    }
    response = NewDismissRoomResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserRequest() (request *RemoveUserRequest) {
    request = &RemoveUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "RemoveUser")
    return
}

func NewRemoveUserResponse() (response *RemoveUserResponse) {
    response = &RemoveUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to remove a user from a room. It is applicable to scenarios where the anchor, room owner, or admin wants to kick out a user. It supports all platforms. For Android, iOS, Windows, and macOS, the TRTC SDK needs to be upgraded to v6.6 or above.
func (c *Client) RemoveUser(request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    if request == nil {
        request = NewRemoveUserRequest()
    }
    response = NewRemoveUserResponse()
    err = c.Send(request, response)
    return
}
