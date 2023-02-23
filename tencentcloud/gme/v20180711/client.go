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

package v20180711

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-11"

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


func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApp
// This API is used to create a GME application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// This API is used to create a GME application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoomMemberRequest() (request *DeleteRoomMemberRequest) {
    request = &DeleteRoomMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DeleteRoomMember")
    
    
    return
}

func NewDeleteRoomMemberResponse() (response *DeleteRoomMemberResponse) {
    response = &DeleteRoomMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoomMember
// This API is used to delete a room or remove members from the room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRoomMember(request *DeleteRoomMemberRequest) (response *DeleteRoomMemberResponse, err error) {
    return c.DeleteRoomMemberWithContext(context.Background(), request)
}

// DeleteRoomMember
// This API is used to delete a room or remove members from the room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRoomMemberWithContext(ctx context.Context, request *DeleteRoomMemberRequest) (response *DeleteRoomMemberResponse, err error) {
    if request == nil {
        request = NewDeleteRoomMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoomMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoomMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppStatisticsRequest() (request *DescribeAppStatisticsRequest) {
    request = &DescribeAppStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeAppStatistics")
    
    
    return
}

func NewDescribeAppStatisticsResponse() (response *DescribeAppStatisticsResponse) {
    response = &DescribeAppStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppStatistics
// This API is used to query the usage statistics of a GME application, including those of Voice Chat, Voice Message Service, Voice Analysis, etc. The maximum query period is the past 30 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_DATEOUTOFSIXTYDAYS = "InvalidParameter.DateOutOfSixtyDays"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatistics(request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    return c.DescribeAppStatisticsWithContext(context.Background(), request)
}

// DescribeAppStatistics
// This API is used to query the usage statistics of a GME application, including those of Voice Chat, Voice Message Service, Voice Analysis, etc. The maximum query period is the past 30 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_DATEOUTOFSIXTYDAYS = "InvalidParameter.DateOutOfSixtyDays"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatisticsWithContext(ctx context.Context, request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAppStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationDataRequest() (request *DescribeApplicationDataRequest) {
    request = &DescribeApplicationDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeApplicationData")
    
    
    return
}

func NewDescribeApplicationDataResponse() (response *DescribeApplicationDataResponse) {
    response = &DescribeApplicationDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationData
// This API is used to query data details for up to the past 90 days.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationData(request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    return c.DescribeApplicationDataWithContext(context.Background(), request)
}

// DescribeApplicationData
// This API is used to query data details for up to the past 90 days.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationDataWithContext(ctx context.Context, request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordInfoRequest() (request *DescribeRecordInfoRequest) {
    request = &DescribeRecordInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeRecordInfo")
    
    
    return
}

func NewDescribeRecordInfoResponse() (response *DescribeRecordInfoResponse) {
    response = &DescribeRecordInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordInfo
// This API is used to query a recording task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeRecordInfo(request *DescribeRecordInfoRequest) (response *DescribeRecordInfoResponse, err error) {
    return c.DescribeRecordInfoWithContext(context.Background(), request)
}

// DescribeRecordInfo
// This API is used to query a recording task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeRecordInfoWithContext(ctx context.Context, request *DescribeRecordInfoRequest) (response *DescribeRecordInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRecordInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeTaskInfo")
    
    
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskInfo
// This API is used to query the recording task in a room.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    return c.DescribeTaskInfoWithContext(context.Background(), request)
}

// DescribeTaskInfo
// This API is used to query the recording task in a room.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeTaskInfoWithContext(ctx context.Context, request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppStatusRequest() (request *ModifyAppStatusRequest) {
    request = &ModifyAppStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ModifyAppStatus")
    
    
    return
}

func NewModifyAppStatusResponse() (response *ModifyAppStatusResponse) {
    response = &ModifyAppStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAppStatus
// This API is used to change the status of an application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatus(request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    return c.ModifyAppStatusWithContext(context.Background(), request)
}

// ModifyAppStatus
// This API is used to change the status of an application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatusWithContext(ctx context.Context, request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    if request == nil {
        request = NewModifyAppStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAppStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordInfoRequest() (request *ModifyRecordInfoRequest) {
    request = &ModifyRecordInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ModifyRecordInfo")
    
    
    return
}

func NewModifyRecordInfoResponse() (response *ModifyRecordInfoResponse) {
    response = &ModifyRecordInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRecordInfo
// This API is used to modify recording configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyRecordInfo(request *ModifyRecordInfoRequest) (response *ModifyRecordInfoResponse, err error) {
    return c.ModifyRecordInfoWithContext(context.Background(), request)
}

// ModifyRecordInfo
// This API is used to modify recording configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyRecordInfoWithContext(ctx context.Context, request *ModifyRecordInfoRequest) (response *ModifyRecordInfoResponse, err error) {
    if request == nil {
        request = NewModifyRecordInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordInfoResponse()
    err = c.Send(request, response)
    return
}

func NewStartRecordRequest() (request *StartRecordRequest) {
    request = &StartRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "StartRecord")
    
    
    return
}

func NewStartRecordResponse() (response *StartRecordResponse) {
    response = &StartRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartRecord
// This API is used to start recording.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCEINUSE_TASKINUSE = "ResourceInUse.TaskInUse"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StartRecord(request *StartRecordRequest) (response *StartRecordResponse, err error) {
    return c.StartRecordWithContext(context.Background(), request)
}

// StartRecord
// This API is used to start recording.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCEINUSE_TASKINUSE = "ResourceInUse.TaskInUse"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StartRecordWithContext(ctx context.Context, request *StartRecordRequest) (response *StartRecordResponse, err error) {
    if request == nil {
        request = NewStartRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopRecordRequest() (request *StopRecordRequest) {
    request = &StopRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "StopRecord")
    
    
    return
}

func NewStopRecordResponse() (response *StopRecordResponse) {
    response = &StopRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopRecord
// This API is used to stop recording.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StopRecord(request *StopRecordRequest) (response *StopRecordResponse, err error) {
    return c.StopRecordWithContext(context.Background(), request)
}

// StopRecord
// This API is used to stop recording.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StopRecordWithContext(ctx context.Context, request *StopRecordRequest) (response *StopRecordResponse, err error) {
    if request == nil {
        request = NewStopRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRecordResponse()
    err = c.Send(request, response)
    return
}
