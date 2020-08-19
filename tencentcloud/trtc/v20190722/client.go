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


func NewCreateTroubleInfoRequest() (request *CreateTroubleInfoRequest) {
    request = &CreateTroubleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "CreateTroubleInfo")
    return
}

func NewCreateTroubleInfoResponse() (response *CreateTroubleInfoResponse) {
    response = &CreateTroubleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create exception information.
func (c *Client) CreateTroubleInfo(request *CreateTroubleInfoRequest) (response *CreateTroubleInfoResponse, err error) {
    if request == nil {
        request = NewCreateTroubleInfoRequest()
    }
    response = NewCreateTroubleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalEventRequest() (request *DescribeAbnormalEventRequest) {
    request = &DescribeAbnormalEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeAbnormalEvent")
    return
}

func NewDescribeAbnormalEventResponse() (response *DescribeAbnormalEventResponse) {
    response = &DescribeAbnormalEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query users’ exceptional experience events according to `SDKAppID` and return the exceptional experience ID and possible causes. It queries data in last 24 hours, and the query period is up to 1 hour which can start and end on different days. For more information about exceptional experience ID mapping, please see here.
func (c *Client) DescribeAbnormalEvent(request *DescribeAbnormalEventRequest) (response *DescribeAbnormalEventResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalEventRequest()
    }
    response = NewDescribeAbnormalEventResponse()
    err = c.Send(request, response)
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

// This API is used to query the user list and user call quality data in a specified time period. It queries data of up to 6 users in the last 5 days. The query period is up to 1 hour, which must start and end on the same day.
func (c *Client) DescribeCallDetail(request *DescribeCallDetailRequest) (response *DescribeCallDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCallDetailRequest()
    }
    response = NewDescribeCallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetailEventRequest() (request *DescribeDetailEventRequest) {
    request = &DescribeDetailEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeDetailEvent")
    return
}

func NewDescribeDetailEventResponse() (response *DescribeDetailEventResponse) {
    response = &DescribeDetailEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query detailed events of a user such as room entry/exit and video enablement/disablement during a call. It can query data for the last 5 days.
func (c *Client) DescribeDetailEvent(request *DescribeDetailEventRequest) (response *DescribeDetailEventResponse, err error) {
    if request == nil {
        request = NewDescribeDetailEventRequest()
    }
    response = NewDescribeDetailEventResponse()
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

// This API is used to query the daily numbers of rooms and users under a specified `sdkqppid`. It can query data once per minute for the last 5 days. If a day has not ended, the numbers of rooms and users on the day cannot be queried.
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

func NewStartMCUMixTranscodeRequest() (request *StartMCUMixTranscodeRequest) {
    request = &StartMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StartMCUMixTranscode")
    return
}

func NewStartMCUMixTranscodeResponse() (response *StartMCUMixTranscodeResponse) {
    response = &StartMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable on-cloud stream mix and specify the layout position of each channel of video image in the mixed video image.
// 
// There may be multiple channels of audio/video streams in a TRTC room. You can call this API to request the Tencent Cloud server to combine multiple channels of video images into one channel, specify the position of each channel, and mix the multiple channels of audio so as to output one channel of audio/video stream for easier recording and live streaming.
// 
// You can use this API to perform the following operations:
// - Set the image and audio quality parameters of the final live stream, including video resolution, video bitrate, video frame rate, and audio quality.
// - Set the image layout, i.e., positions of all channels of images. You only need to set the layout once when enabling on-cloud stream mix, and the layout engine will automatically arrange the video images in the configured layout in subsequent operations.
// - Set the recording file name for future playback.
// - Set the CDN live stream ID for live streaming over CDN.
// 
// Currently, the following layout templates are supported:
// - Floating template: the entire screen will be covered by the video image of the first user who enters the room, and the video images of other users will be displayed as small images in horizontal rows from the bottom-left corner in room entry sequence. The screen can contain up to 4 lines with 4 small images each row, which float over the big image. Up to 1 big image and 15 small images are supported. If a user sends audio only, the user will still use an image spot.
// - Grid template: the screen is divided into user video images with the same dimensions. The more the users, the smaller the image dimensions. Up to 16 images are supported. If a user sends audio only, the user will still use an image spot.
// - Screen sharing template: it is suitable for video conferencing and online education. The shared screen (or camera of the anchor) is always displayed in the big image on the left of the screen, and the video images of other users are vertically displayed on the right in up to 2 columns with up to 8 small images in each column. Up to 1 big image and 15 small images are supported. If a user sends audio only, the user will still use an image spot.
// - Picture-in-picture template: it is suitable for mixing a pair of big/small images or a big image with the audio of other users. The small image floats over the big image, and the users in the big/small images and the display position of the small image can be specified.
func (c *Client) StartMCUMixTranscode(request *StartMCUMixTranscodeRequest) (response *StartMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeRequest()
    }
    response = NewStartMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewStopMCUMixTranscodeRequest() (request *StopMCUMixTranscodeRequest) {
    request = &StopMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StopMCUMixTranscode")
    return
}

func NewStopMCUMixTranscodeResponse() (response *StopMCUMixTranscodeResponse) {
    response = &StopMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to end On-Cloud MixTranscoding.
func (c *Client) StopMCUMixTranscode(request *StopMCUMixTranscodeRequest) (response *StopMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeRequest()
    }
    response = NewStopMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}
