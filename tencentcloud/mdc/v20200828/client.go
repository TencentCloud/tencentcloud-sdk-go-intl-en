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

package v20200828

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-08-28"

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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateStreamLinkFlowRequest() (request *CreateStreamLinkFlowRequest) {
    request = &CreateStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "CreateStreamLinkFlow")
    
    
    return
}

func NewCreateStreamLinkFlowResponse() (response *CreateStreamLinkFlowResponse) {
    response = &CreateStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamLinkFlow
// This API is used to create a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkFlow(request *CreateStreamLinkFlowRequest) (response *CreateStreamLinkFlowResponse, err error) {
    return c.CreateStreamLinkFlowWithContext(context.Background(), request)
}

// CreateStreamLinkFlow
// This API is used to create a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkFlowWithContext(ctx context.Context, request *CreateStreamLinkFlowRequest) (response *CreateStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLinkOutputInfoRequest() (request *CreateStreamLinkOutputInfoRequest) {
    request = &CreateStreamLinkOutputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "CreateStreamLinkOutputInfo")
    
    
    return
}

func NewCreateStreamLinkOutputInfoResponse() (response *CreateStreamLinkOutputInfoResponse) {
    response = &CreateStreamLinkOutputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamLinkOutputInfo
// This API is used to create a StreamLink output.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) CreateStreamLinkOutputInfo(request *CreateStreamLinkOutputInfoRequest) (response *CreateStreamLinkOutputInfoResponse, err error) {
    return c.CreateStreamLinkOutputInfoWithContext(context.Background(), request)
}

// CreateStreamLinkOutputInfo
// This API is used to create a StreamLink output.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) CreateStreamLinkOutputInfoWithContext(ctx context.Context, request *CreateStreamLinkOutputInfoRequest) (response *CreateStreamLinkOutputInfoResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkOutputInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamLinkOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamLinkOutputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkFlowRequest() (request *DeleteStreamLinkFlowRequest) {
    request = &DeleteStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DeleteStreamLinkFlow")
    
    
    return
}

func NewDeleteStreamLinkFlowResponse() (response *DeleteStreamLinkFlowResponse) {
    response = &DeleteStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLinkFlow
// This API is used to delete a StreamLink flow.
//
// error code that may be returned:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkFlow(request *DeleteStreamLinkFlowRequest) (response *DeleteStreamLinkFlowResponse, err error) {
    return c.DeleteStreamLinkFlowWithContext(context.Background(), request)
}

// DeleteStreamLinkFlow
// This API is used to delete a StreamLink flow.
//
// error code that may be returned:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkFlowWithContext(ctx context.Context, request *DeleteStreamLinkFlowRequest) (response *DeleteStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkOutputRequest() (request *DeleteStreamLinkOutputRequest) {
    request = &DeleteStreamLinkOutputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DeleteStreamLinkOutput")
    
    
    return
}

func NewDeleteStreamLinkOutputResponse() (response *DeleteStreamLinkOutputResponse) {
    response = &DeleteStreamLinkOutputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLinkOutput
// This API is used to delete an output of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkOutput(request *DeleteStreamLinkOutputRequest) (response *DeleteStreamLinkOutputResponse, err error) {
    return c.DeleteStreamLinkOutputWithContext(context.Background(), request)
}

// DeleteStreamLinkOutput
// This API is used to delete an output of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkOutputWithContext(ctx context.Context, request *DeleteStreamLinkOutputRequest) (response *DeleteStreamLinkOutputResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkOutputRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamLinkOutput require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkOutputResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowRequest() (request *DescribeStreamLinkFlowRequest) {
    request = &DescribeStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlow")
    
    
    return
}

func NewDescribeStreamLinkFlowResponse() (response *DescribeStreamLinkFlowResponse) {
    response = &DescribeStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlow
// This API is used to query the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
func (c *Client) DescribeStreamLinkFlow(request *DescribeStreamLinkFlowRequest) (response *DescribeStreamLinkFlowResponse, err error) {
    return c.DescribeStreamLinkFlowWithContext(context.Background(), request)
}

// DescribeStreamLinkFlow
// This API is used to query the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
func (c *Client) DescribeStreamLinkFlowWithContext(ctx context.Context, request *DescribeStreamLinkFlowRequest) (response *DescribeStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowLogsRequest() (request *DescribeStreamLinkFlowLogsRequest) {
    request = &DescribeStreamLinkFlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlowLogs")
    
    
    return
}

func NewDescribeStreamLinkFlowLogsResponse() (response *DescribeStreamLinkFlowLogsResponse) {
    response = &DescribeStreamLinkFlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlowLogs
// This API is used to query the logs of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_SORTTYPE = "InvalidParameter.SortType"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowLogs(request *DescribeStreamLinkFlowLogsRequest) (response *DescribeStreamLinkFlowLogsResponse, err error) {
    return c.DescribeStreamLinkFlowLogsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowLogs
// This API is used to query the logs of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_SORTTYPE = "InvalidParameter.SortType"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowLogsWithContext(ctx context.Context, request *DescribeStreamLinkFlowLogsRequest) (response *DescribeStreamLinkFlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowMediaStatisticsRequest() (request *DescribeStreamLinkFlowMediaStatisticsRequest) {
    request = &DescribeStreamLinkFlowMediaStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlowMediaStatistics")
    
    
    return
}

func NewDescribeStreamLinkFlowMediaStatisticsResponse() (response *DescribeStreamLinkFlowMediaStatisticsResponse) {
    response = &DescribeStreamLinkFlowMediaStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlowMediaStatistics
// This API is used to query the media quality of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowMediaStatistics(request *DescribeStreamLinkFlowMediaStatisticsRequest) (response *DescribeStreamLinkFlowMediaStatisticsResponse, err error) {
    return c.DescribeStreamLinkFlowMediaStatisticsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowMediaStatistics
// This API is used to query the media quality of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowMediaStatisticsWithContext(ctx context.Context, request *DescribeStreamLinkFlowMediaStatisticsRequest) (response *DescribeStreamLinkFlowMediaStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowMediaStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowMediaStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowMediaStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowRealtimeStatusRequest() (request *DescribeStreamLinkFlowRealtimeStatusRequest) {
    request = &DescribeStreamLinkFlowRealtimeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlowRealtimeStatus")
    
    
    return
}

func NewDescribeStreamLinkFlowRealtimeStatusResponse() (response *DescribeStreamLinkFlowRealtimeStatusResponse) {
    response = &DescribeStreamLinkFlowRealtimeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlowRealtimeStatus
// This API is used to query the current status of a flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLinkFlowRealtimeStatus(request *DescribeStreamLinkFlowRealtimeStatusRequest) (response *DescribeStreamLinkFlowRealtimeStatusResponse, err error) {
    return c.DescribeStreamLinkFlowRealtimeStatusWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowRealtimeStatus
// This API is used to query the current status of a flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLinkFlowRealtimeStatusWithContext(ctx context.Context, request *DescribeStreamLinkFlowRealtimeStatusRequest) (response *DescribeStreamLinkFlowRealtimeStatusResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowRealtimeStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowRealtimeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowRealtimeStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowSRTStatisticsRequest() (request *DescribeStreamLinkFlowSRTStatisticsRequest) {
    request = &DescribeStreamLinkFlowSRTStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlowSRTStatistics")
    
    
    return
}

func NewDescribeStreamLinkFlowSRTStatisticsResponse() (response *DescribeStreamLinkFlowSRTStatisticsResponse) {
    response = &DescribeStreamLinkFlowSRTStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlowSRTStatistics
// This API is used to query the SRT streaming performance of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowSRTStatistics(request *DescribeStreamLinkFlowSRTStatisticsRequest) (response *DescribeStreamLinkFlowSRTStatisticsResponse, err error) {
    return c.DescribeStreamLinkFlowSRTStatisticsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowSRTStatistics
// This API is used to query the SRT streaming performance of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowSRTStatisticsWithContext(ctx context.Context, request *DescribeStreamLinkFlowSRTStatisticsRequest) (response *DescribeStreamLinkFlowSRTStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowSRTStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowSRTStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowSRTStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowStatisticsRequest() (request *DescribeStreamLinkFlowStatisticsRequest) {
    request = &DescribeStreamLinkFlowStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlowStatistics")
    
    
    return
}

func NewDescribeStreamLinkFlowStatisticsResponse() (response *DescribeStreamLinkFlowStatisticsResponse) {
    response = &DescribeStreamLinkFlowStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlowStatistics
// This API is used to query the media quality of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowStatistics(request *DescribeStreamLinkFlowStatisticsRequest) (response *DescribeStreamLinkFlowStatisticsResponse, err error) {
    return c.DescribeStreamLinkFlowStatisticsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlowStatistics
// This API is used to query the media quality of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamLinkFlowStatisticsWithContext(ctx context.Context, request *DescribeStreamLinkFlowStatisticsRequest) (response *DescribeStreamLinkFlowStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlowStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowsRequest() (request *DescribeStreamLinkFlowsRequest) {
    request = &DescribeStreamLinkFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlows")
    
    
    return
}

func NewDescribeStreamLinkFlowsResponse() (response *DescribeStreamLinkFlowsResponse) {
    response = &DescribeStreamLinkFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlows
// This API is used to query the configuration information of multiple StreamLink flows in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkFlows(request *DescribeStreamLinkFlowsRequest) (response *DescribeStreamLinkFlowsResponse, err error) {
    return c.DescribeStreamLinkFlowsWithContext(context.Background(), request)
}

// DescribeStreamLinkFlows
// This API is used to query the configuration information of multiple StreamLink flows in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkFlowsWithContext(ctx context.Context, request *DescribeStreamLinkFlowsRequest) (response *DescribeStreamLinkFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkRegionsRequest() (request *DescribeStreamLinkRegionsRequest) {
    request = &DescribeStreamLinkRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkRegions")
    
    
    return
}

func NewDescribeStreamLinkRegionsResponse() (response *DescribeStreamLinkRegionsResponse) {
    response = &DescribeStreamLinkRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkRegions
// This API is used to query all StreamLink regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkRegions(request *DescribeStreamLinkRegionsRequest) (response *DescribeStreamLinkRegionsResponse, err error) {
    return c.DescribeStreamLinkRegionsWithContext(context.Background(), request)
}

// DescribeStreamLinkRegions
// This API is used to query all StreamLink regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkRegionsWithContext(ctx context.Context, request *DescribeStreamLinkRegionsRequest) (response *DescribeStreamLinkRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamLinkRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkFlowRequest() (request *ModifyStreamLinkFlowRequest) {
    request = &ModifyStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "ModifyStreamLinkFlow")
    
    
    return
}

func NewModifyStreamLinkFlowResponse() (response *ModifyStreamLinkFlowResponse) {
    response = &ModifyStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamLinkFlow
// This API is used to modify the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkFlow(request *ModifyStreamLinkFlowRequest) (response *ModifyStreamLinkFlowResponse, err error) {
    return c.ModifyStreamLinkFlowWithContext(context.Background(), request)
}

// ModifyStreamLinkFlow
// This API is used to modify the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkFlowWithContext(ctx context.Context, request *ModifyStreamLinkFlowRequest) (response *ModifyStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkInputRequest() (request *ModifyStreamLinkInputRequest) {
    request = &ModifyStreamLinkInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "ModifyStreamLinkInput")
    
    
    return
}

func NewModifyStreamLinkInputResponse() (response *ModifyStreamLinkInputResponse) {
    response = &ModifyStreamLinkInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamLinkInput
// This API is used to modify an input of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkInput(request *ModifyStreamLinkInputRequest) (response *ModifyStreamLinkInputResponse, err error) {
    return c.ModifyStreamLinkInputWithContext(context.Background(), request)
}

// ModifyStreamLinkInput
// This API is used to modify an input of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkInputWithContext(ctx context.Context, request *ModifyStreamLinkInputRequest) (response *ModifyStreamLinkInputResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkInputRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkInput require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkInputResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkOutputInfoRequest() (request *ModifyStreamLinkOutputInfoRequest) {
    request = &ModifyStreamLinkOutputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "ModifyStreamLinkOutputInfo")
    
    
    return
}

func NewModifyStreamLinkOutputInfoResponse() (response *ModifyStreamLinkOutputInfoResponse) {
    response = &ModifyStreamLinkOutputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamLinkOutputInfo
// This API is used to modify an output of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkOutputInfo(request *ModifyStreamLinkOutputInfoRequest) (response *ModifyStreamLinkOutputInfoResponse, err error) {
    return c.ModifyStreamLinkOutputInfoWithContext(context.Background(), request)
}

// ModifyStreamLinkOutputInfo
// This API is used to modify an output of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkOutputInfoWithContext(ctx context.Context, request *ModifyStreamLinkOutputInfoRequest) (response *ModifyStreamLinkOutputInfoResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkOutputInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamLinkOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamLinkOutputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewStartStreamLinkFlowRequest() (request *StartStreamLinkFlowRequest) {
    request = &StartStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "StartStreamLinkFlow")
    
    
    return
}

func NewStartStreamLinkFlowResponse() (response *StartStreamLinkFlowResponse) {
    response = &StartStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartStreamLinkFlow
// This API is used to start a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StartStreamLinkFlow(request *StartStreamLinkFlowRequest) (response *StartStreamLinkFlowResponse, err error) {
    return c.StartStreamLinkFlowWithContext(context.Background(), request)
}

// StartStreamLinkFlow
// This API is used to start a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StartStreamLinkFlowWithContext(ctx context.Context, request *StartStreamLinkFlowRequest) (response *StartStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStartStreamLinkFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewStopStreamLinkFlowRequest() (request *StopStreamLinkFlowRequest) {
    request = &StopStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "StopStreamLinkFlow")
    
    
    return
}

func NewStopStreamLinkFlowResponse() (response *StopStreamLinkFlowResponse) {
    response = &StopStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopStreamLinkFlow
// This API is used to stop a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StopStreamLinkFlow(request *StopStreamLinkFlowRequest) (response *StopStreamLinkFlowResponse, err error) {
    return c.StopStreamLinkFlowWithContext(context.Background(), request)
}

// StopStreamLinkFlow
// This API is used to stop a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StopStreamLinkFlowWithContext(ctx context.Context, request *StopStreamLinkFlowRequest) (response *StopStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStopStreamLinkFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopStreamLinkFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}
