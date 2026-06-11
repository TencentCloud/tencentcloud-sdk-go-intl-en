// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20201016

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-10-16"

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


func NewAddMachineGroupInfoRequest() (request *AddMachineGroupInfoRequest) {
    request = &AddMachineGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "AddMachineGroupInfo")
    
    
    return
}

func NewAddMachineGroupInfoResponse() (response *AddMachineGroupInfoResponse) {
    response = &AddMachineGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddMachineGroupInfo
// This API is used to add machine group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
func (c *Client) AddMachineGroupInfo(request *AddMachineGroupInfoRequest) (response *AddMachineGroupInfoResponse, err error) {
    return c.AddMachineGroupInfoWithContext(context.Background(), request)
}

// AddMachineGroupInfo
// This API is used to add machine group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
func (c *Client) AddMachineGroupInfoWithContext(ctx context.Context, request *AddMachineGroupInfoRequest) (response *AddMachineGroupInfoResponse, err error) {
    if request == nil {
        request = NewAddMachineGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "AddMachineGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddMachineGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddMachineGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewApplyConfigToMachineGroupRequest() (request *ApplyConfigToMachineGroupRequest) {
    request = &ApplyConfigToMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ApplyConfigToMachineGroup")
    
    
    return
}

func NewApplyConfigToMachineGroupResponse() (response *ApplyConfigToMachineGroupResponse) {
    response = &ApplyConfigToMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyConfigToMachineGroup
// This API is used to apply the collection configuration to a specified machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ApplyConfigToMachineGroup(request *ApplyConfigToMachineGroupRequest) (response *ApplyConfigToMachineGroupResponse, err error) {
    return c.ApplyConfigToMachineGroupWithContext(context.Background(), request)
}

// ApplyConfigToMachineGroup
// This API is used to apply the collection configuration to a specified machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ApplyConfigToMachineGroupWithContext(ctx context.Context, request *ApplyConfigToMachineGroupRequest) (response *ApplyConfigToMachineGroupResponse, err error) {
    if request == nil {
        request = NewApplyConfigToMachineGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ApplyConfigToMachineGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyConfigToMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyConfigToMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCancelRebuildIndexTaskRequest() (request *CancelRebuildIndexTaskRequest) {
    request = &CancelRebuildIndexTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CancelRebuildIndexTask")
    
    
    return
}

func NewCancelRebuildIndexTaskResponse() (response *CancelRebuildIndexTaskResponse) {
    response = &CancelRebuildIndexTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelRebuildIndexTask
// This API is used to cancel creating a reindexing task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CancelRebuildIndexTask(request *CancelRebuildIndexTaskRequest) (response *CancelRebuildIndexTaskResponse, err error) {
    return c.CancelRebuildIndexTaskWithContext(context.Background(), request)
}

// CancelRebuildIndexTask
// This API is used to cancel creating a reindexing task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CancelRebuildIndexTaskWithContext(ctx context.Context, request *CancelRebuildIndexTaskRequest) (response *CancelRebuildIndexTaskResponse, err error) {
    if request == nil {
        request = NewCancelRebuildIndexTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CancelRebuildIndexTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelRebuildIndexTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelRebuildIndexTaskResponse()
    err = c.Send(request, response)
    return
}

func NewChatCompletionsRequest() (request *ChatCompletionsRequest) {
    request = &ChatCompletionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ChatCompletions")
    
    
    return
}

func NewChatCompletionsResponse() (response *ChatCompletionsResponse) {
    response = &ChatCompletionsResponse{} 
    return

}

// ChatCompletions
// Call the API to initiate a dialogue request.
//
// This API supports AI-powered logging features such as intelligently generating retrieval analysis statements.
//
// Note: When calling this API via SSE streaming, ensure the request domain name is set to cls.ai.tencentcloudapi.com (configurable as cls.ai.internal.tencentcloudapi.com in a private network environment).
//
// error code that may be returned:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ChatCompletions(request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    return c.ChatCompletionsWithContext(context.Background(), request)
}

// ChatCompletions
// Call the API to initiate a dialogue request.
//
// This API supports AI-powered logging features such as intelligently generating retrieval analysis statements.
//
// Note: When calling this API via SSE streaming, ensure the request domain name is set to cls.ai.tencentcloudapi.com (configurable as cls.ai.internal.tencentcloudapi.com in a private network environment).
//
// error code that may be returned:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ChatCompletionsWithContext(ctx context.Context, request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
    if request == nil {
        request = NewChatCompletionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ChatCompletions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatCompletions require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatCompletionsResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFunctionRequest() (request *CheckFunctionRequest) {
    request = &CheckFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CheckFunction")
    
    
    return
}

func NewCheckFunctionResponse() (response *CheckFunctionResponse) {
    response = &CheckFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckFunction
// This API is used to verify the syntax of data processing DSL functions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CROSSACCOUNTCONFLICT = "InvalidParameter.CrossAccountConflict"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CheckFunction(request *CheckFunctionRequest) (response *CheckFunctionResponse, err error) {
    return c.CheckFunctionWithContext(context.Background(), request)
}

// CheckFunction
// This API is used to verify the syntax of data processing DSL functions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CROSSACCOUNTCONFLICT = "InvalidParameter.CrossAccountConflict"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CheckFunctionWithContext(ctx context.Context, request *CheckFunctionRequest) (response *CheckFunctionResponse, err error) {
    if request == nil {
        request = NewCheckFunctionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CheckFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCheckRechargeKafkaServerRequest() (request *CheckRechargeKafkaServerRequest) {
    request = &CheckRechargeKafkaServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CheckRechargeKafkaServer")
    
    
    return
}

func NewCheckRechargeKafkaServerResponse() (response *CheckRechargeKafkaServerResponse) {
    response = &CheckRechargeKafkaServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckRechargeKafkaServer
// This API is used to check whether the Kafka service cluster is accessible.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckRechargeKafkaServer(request *CheckRechargeKafkaServerRequest) (response *CheckRechargeKafkaServerResponse, err error) {
    return c.CheckRechargeKafkaServerWithContext(context.Background(), request)
}

// CheckRechargeKafkaServer
// This API is used to check whether the Kafka service cluster is accessible.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckRechargeKafkaServerWithContext(ctx context.Context, request *CheckRechargeKafkaServerRequest) (response *CheckRechargeKafkaServerResponse, err error) {
    if request == nil {
        request = NewCheckRechargeKafkaServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CheckRechargeKafkaServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckRechargeKafkaServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckRechargeKafkaServerResponse()
    err = c.Send(request, response)
    return
}

func NewCloseKafkaConsumerRequest() (request *CloseKafkaConsumerRequest) {
    request = &CloseKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CloseKafkaConsumer")
    
    
    return
}

func NewCloseKafkaConsumerResponse() (response *CloseKafkaConsumerResponse) {
    response = &CloseKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseKafkaConsumer
// This API is used to disable Kafka consumption.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CloseKafkaConsumer(request *CloseKafkaConsumerRequest) (response *CloseKafkaConsumerResponse, err error) {
    return c.CloseKafkaConsumerWithContext(context.Background(), request)
}

// CloseKafkaConsumer
// This API is used to disable Kafka consumption.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CloseKafkaConsumerWithContext(ctx context.Context, request *CloseKafkaConsumerRequest) (response *CloseKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewCloseKafkaConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CloseKafkaConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCommitConsumerOffsetsRequest() (request *CommitConsumerOffsetsRequest) {
    request = &CommitConsumerOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CommitConsumerOffsets")
    
    
    return
}

func NewCommitConsumerOffsetsResponse() (response *CommitConsumerOffsetsResponse) {
    response = &CommitConsumerOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitConsumerOffsets
// This API is used to submit a consumption offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CommitConsumerOffsets(request *CommitConsumerOffsetsRequest) (response *CommitConsumerOffsetsResponse, err error) {
    return c.CommitConsumerOffsetsWithContext(context.Background(), request)
}

// CommitConsumerOffsets
// This API is used to submit a consumption offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CommitConsumerOffsetsWithContext(ctx context.Context, request *CommitConsumerOffsetsRequest) (response *CommitConsumerOffsetsResponse, err error) {
    if request == nil {
        request = NewCommitConsumerOffsetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CommitConsumerOffsets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitConsumerOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitConsumerOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmRequest() (request *CreateAlarmRequest) {
    request = &CreateAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateAlarm")
    
    
    return
}

func NewCreateAlarmResponse() (response *CreateAlarmResponse) {
    response = &CreateAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarm
// This API is used to create an alarm policy.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateAlarm(request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
    return c.CreateAlarmWithContext(context.Background(), request)
}

// CreateAlarm
// This API is used to create an alarm policy.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateAlarmWithContext(ctx context.Context, request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
    if request == nil {
        request = NewCreateAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmNoticeRequest() (request *CreateAlarmNoticeRequest) {
    request = &CreateAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateAlarmNotice")
    
    
    return
}

func NewCreateAlarmNoticeResponse() (response *CreateAlarmNoticeResponse) {
    response = &CreateAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmNotice
// This API is used to create a notification channel group with two configuration modes to choose from.
//
// 1. Easy mode provides the most basic notification channel function. The following parameters are required:
//
// - Type
//
// - NoticeReceivers
//
// - WebCallbacks
//
// 
//
// 2. Advanced mode: On the basis of easy mode, it supports setting rules, setting notification channels for different types of alarms, and escalating alarms. The following parameters are required:
//
// - NoticeRules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmNotice(request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    return c.CreateAlarmNoticeWithContext(context.Background(), request)
}

// CreateAlarmNotice
// This API is used to create a notification channel group with two configuration modes to choose from.
//
// 1. Easy mode provides the most basic notification channel function. The following parameters are required:
//
// - Type
//
// - NoticeReceivers
//
// - WebCallbacks
//
// 
//
// 2. Advanced mode: On the basis of easy mode, it supports setting rules, setting notification channels for different types of alarms, and escalating alarms. The following parameters are required:
//
// - NoticeRules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmNoticeWithContext(ctx context.Context, request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewCreateAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmShieldRequest() (request *CreateAlarmShieldRequest) {
    request = &CreateAlarmShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateAlarmShield")
    
    
    return
}

func NewCreateAlarmShieldResponse() (response *CreateAlarmShieldResponse) {
    response = &CreateAlarmShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmShield
// This API is used to create an alarm blocking rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmShield(request *CreateAlarmShieldRequest) (response *CreateAlarmShieldResponse, err error) {
    return c.CreateAlarmShieldWithContext(context.Background(), request)
}

// CreateAlarmShield
// This API is used to create an alarm blocking rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmShieldWithContext(ctx context.Context, request *CreateAlarmShieldRequest) (response *CreateAlarmShieldResponse, err error) {
    if request == nil {
        request = NewCreateAlarmShieldRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateAlarmShield")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmShield require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmShieldResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudProductLogCollectionRequest() (request *CreateCloudProductLogCollectionRequest) {
    request = &CreateCloudProductLogCollectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateCloudProductLogCollection")
    
    
    return
}

func NewCreateCloudProductLogCollectionResponse() (response *CreateCloudProductLogCollectionResponse) {
    response = &CreateCloudProductLogCollectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudProductLogCollection
// Cloud product integration uses internal APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDPRODUCTINVOCATIONERROR = "FailedOperation.CloudProductInvocationError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudProductLogCollection(request *CreateCloudProductLogCollectionRequest) (response *CreateCloudProductLogCollectionResponse, err error) {
    return c.CreateCloudProductLogCollectionWithContext(context.Background(), request)
}

// CreateCloudProductLogCollection
// Cloud product integration uses internal APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDPRODUCTINVOCATIONERROR = "FailedOperation.CloudProductInvocationError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudProductLogCollectionWithContext(ctx context.Context, request *CreateCloudProductLogCollectionRequest) (response *CreateCloudProductLogCollectionResponse, err error) {
    if request == nil {
        request = NewCreateCloudProductLogCollectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateCloudProductLogCollection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudProductLogCollection require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudProductLogCollectionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
    request = &CreateConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConfig")
    
    
    return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
    response = &CreateConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfig
// This API is used to create collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_CONFIG = "LimitExceeded.Config"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    return c.CreateConfigWithContext(context.Background(), request)
}

// CreateConfig
// This API is used to create collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_CONFIG = "LimitExceeded.Config"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateConfigWithContext(ctx context.Context, request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    if request == nil {
        request = NewCreateConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsoleRequest() (request *CreateConsoleRequest) {
    request = &CreateConsoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConsole")
    
    
    return
}

func NewCreateConsoleResponse() (response *CreateConsoleResponse) {
    response = &CreateConsoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsole
// This API is used to create the DataSight Console
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsole(request *CreateConsoleRequest) (response *CreateConsoleResponse, err error) {
    return c.CreateConsoleWithContext(context.Background(), request)
}

// CreateConsole
// This API is used to create the DataSight Console
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsoleWithContext(ctx context.Context, request *CreateConsoleRequest) (response *CreateConsoleResponse, err error) {
    if request == nil {
        request = NewCreateConsoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateConsole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerRequest() (request *CreateConsumerRequest) {
    request = &CreateConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConsumer")
    
    
    return
}

func NewCreateConsumerResponse() (response *CreateConsumerResponse) {
    response = &CreateConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsumer
// This API is used to create a CKafka delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsumer(request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    return c.CreateConsumerWithContext(context.Background(), request)
}

// CreateConsumer
// This API is used to create a CKafka delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsumerWithContext(ctx context.Context, request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    if request == nil {
        request = NewCreateConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
    request = &CreateConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConsumerGroup")
    
    
    return
}

func NewCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
    response = &CreateConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsumerGroup
// This API is used to check the heartbeat of a consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    return c.CreateConsumerGroupWithContext(context.Background(), request)
}

// CreateConsumerGroup
// This API is used to check the heartbeat of a consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsumerGroupWithContext(ctx context.Context, request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    if request == nil {
        request = NewCreateConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosRechargeRequest() (request *CreateCosRechargeRequest) {
    request = &CreateCosRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateCosRecharge")
    
    
    return
}

func NewCreateCosRechargeResponse() (response *CreateCosRechargeResponse) {
    response = &CreateCosRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosRecharge
// This API is used to create a COS import task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED_BILLINGCOSRECHARGEOUTOFLIMIT = "LimitExceeded.BillingCosRechargeOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateCosRecharge(request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
    return c.CreateCosRechargeWithContext(context.Background(), request)
}

// CreateCosRecharge
// This API is used to create a COS import task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED_BILLINGCOSRECHARGEOUTOFLIMIT = "LimitExceeded.BillingCosRechargeOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateCosRechargeWithContext(ctx context.Context, request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
    if request == nil {
        request = NewCreateCosRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateCosRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDashboardRequest() (request *CreateDashboardRequest) {
    request = &CreateDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDashboard")
    
    
    return
}

func NewCreateDashboardResponse() (response *CreateDashboardResponse) {
    response = &CreateDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDashboard
// This API is used to create a dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateDashboard(request *CreateDashboardRequest) (response *CreateDashboardResponse, err error) {
    return c.CreateDashboardWithContext(context.Background(), request)
}

// CreateDashboard
// This API is used to create a dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateDashboardWithContext(ctx context.Context, request *CreateDashboardRequest) (response *CreateDashboardResponse, err error) {
    if request == nil {
        request = NewCreateDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDashboardSubscribeRequest() (request *CreateDashboardSubscribeRequest) {
    request = &CreateDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDashboardSubscribe")
    
    
    return
}

func NewCreateDashboardSubscribeResponse() (response *CreateDashboardSubscribeResponse) {
    response = &CreateDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDashboardSubscribe
// This API is used to create a dashboard subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) CreateDashboardSubscribe(request *CreateDashboardSubscribeRequest) (response *CreateDashboardSubscribeResponse, err error) {
    return c.CreateDashboardSubscribeWithContext(context.Background(), request)
}

// CreateDashboardSubscribe
// This API is used to create a dashboard subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) CreateDashboardSubscribeWithContext(ctx context.Context, request *CreateDashboardSubscribeRequest) (response *CreateDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateDashboardSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateDashboardSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataTransformRequest() (request *CreateDataTransformRequest) {
    request = &CreateDataTransformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDataTransform")
    
    
    return
}

func NewCreateDataTransformResponse() (response *CreateDataTransformResponse) {
    response = &CreateDataTransformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataTransform
// This API is used to create a data processing task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CROSSACCOUNTCONFLICT = "InvalidParameter.CrossAccountConflict"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateDataTransform(request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
    return c.CreateDataTransformWithContext(context.Background(), request)
}

// CreateDataTransform
// This API is used to create a data processing task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CROSSACCOUNTCONFLICT = "InvalidParameter.CrossAccountConflict"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateDataTransformWithContext(ctx context.Context, request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
    if request == nil {
        request = NewCreateDataTransformRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateDataTransform")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataTransformResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeliverCloudFunctionRequest() (request *CreateDeliverCloudFunctionRequest) {
    request = &CreateDeliverCloudFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDeliverCloudFunction")
    
    
    return
}

func NewCreateDeliverCloudFunctionResponse() (response *CreateDeliverCloudFunctionResponse) {
    response = &CreateDeliverCloudFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeliverCloudFunction
// This API is used to create a delivery SCF task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeliverCloudFunction(request *CreateDeliverCloudFunctionRequest) (response *CreateDeliverCloudFunctionResponse, err error) {
    return c.CreateDeliverCloudFunctionWithContext(context.Background(), request)
}

// CreateDeliverCloudFunction
// This API is used to create a delivery SCF task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeliverCloudFunctionWithContext(ctx context.Context, request *CreateDeliverCloudFunctionRequest) (response *CreateDeliverCloudFunctionResponse, err error) {
    if request == nil {
        request = NewCreateDeliverCloudFunctionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateDeliverCloudFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeliverCloudFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeliverCloudFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDlcDeliverRequest() (request *CreateDlcDeliverRequest) {
    request = &CreateDlcDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDlcDeliver")
    
    
    return
}

func NewCreateDlcDeliverResponse() (response *CreateDlcDeliverResponse) {
    response = &CreateDlcDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDlcDeliver
// Create a DLC delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDlcDeliver(request *CreateDlcDeliverRequest) (response *CreateDlcDeliverResponse, err error) {
    return c.CreateDlcDeliverWithContext(context.Background(), request)
}

// CreateDlcDeliver
// Create a DLC delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDlcDeliverWithContext(ctx context.Context, request *CreateDlcDeliverRequest) (response *CreateDlcDeliverResponse, err error) {
    if request == nil {
        request = NewCreateDlcDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateDlcDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDlcDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDlcDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEsRechargeRequest() (request *CreateEsRechargeRequest) {
    request = &CreateEsRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateEsRecharge")
    
    
    return
}

func NewCreateEsRechargeResponse() (response *CreateEsRechargeResponse) {
    response = &CreateEsRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEsRecharge
// This API is used to create an es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEsRecharge(request *CreateEsRechargeRequest) (response *CreateEsRechargeResponse, err error) {
    return c.CreateEsRechargeWithContext(context.Background(), request)
}

// CreateEsRecharge
// This API is used to create an es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEsRechargeWithContext(ctx context.Context, request *CreateEsRechargeRequest) (response *CreateEsRechargeResponse, err error) {
    if request == nil {
        request = NewCreateEsRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateEsRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEsRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEsRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExportRequest() (request *CreateExportRequest) {
    request = &CreateExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateExport")
    
    
    return
}

func NewCreateExportResponse() (response *CreateExportResponse) {
    response = &CreateExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExport
// This API only creates download tasks. The download address returned by the task can be obtained by user invocation of [DescribeExports](https://www.tencentcloud.com/document/product/614/56449?from_cn_redirect=1) to view task list, which includes the CosPath parameter for the download address.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EXPORT = "LimitExceeded.Export"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateExport(request *CreateExportRequest) (response *CreateExportResponse, err error) {
    return c.CreateExportWithContext(context.Background(), request)
}

// CreateExport
// This API only creates download tasks. The download address returned by the task can be obtained by user invocation of [DescribeExports](https://www.tencentcloud.com/document/product/614/56449?from_cn_redirect=1) to view task list, which includes the CosPath parameter for the download address.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EXPORT = "LimitExceeded.Export"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateExportWithContext(ctx context.Context, request *CreateExportRequest) (response *CreateExportResponse, err error) {
    if request == nil {
        request = NewCreateExportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateExport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostMetricConfigRequest() (request *CreateHostMetricConfigRequest) {
    request = &CreateHostMetricConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateHostMetricConfig")
    
    
    return
}

func NewCreateHostMetricConfigResponse() (response *CreateHostMetricConfigResponse) {
    response = &CreateHostMetricConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHostMetricConfig
// This API is used to create host metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateHostMetricConfig(request *CreateHostMetricConfigRequest) (response *CreateHostMetricConfigResponse, err error) {
    return c.CreateHostMetricConfigWithContext(context.Background(), request)
}

// CreateHostMetricConfig
// This API is used to create host metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateHostMetricConfigWithContext(ctx context.Context, request *CreateHostMetricConfigRequest) (response *CreateHostMetricConfigResponse, err error) {
    if request == nil {
        request = NewCreateHostMetricConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateHostMetricConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHostMetricConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHostMetricConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIndexRequest() (request *CreateIndexRequest) {
    request = &CreateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateIndex")
    
    
    return
}

func NewCreateIndexResponse() (response *CreateIndexResponse) {
    response = &CreateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIndex
// This API is used to create an index.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXKEYOVERLIMIT = "LimitExceeded.IndexKeyOverLimit"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndex(request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    return c.CreateIndexWithContext(context.Background(), request)
}

// CreateIndex
// This API is used to create an index.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXKEYOVERLIMIT = "LimitExceeded.IndexKeyOverLimit"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndexWithContext(ctx context.Context, request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    if request == nil {
        request = NewCreateIndexRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateIndex")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKafkaRechargeRequest() (request *CreateKafkaRechargeRequest) {
    request = &CreateKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateKafkaRecharge")
    
    
    return
}

func NewCreateKafkaRechargeResponse() (response *CreateKafkaRechargeResponse) {
    response = &CreateKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKafkaRecharge
// This API is used to create a Kafka data subscription task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateKafkaRecharge(request *CreateKafkaRechargeRequest) (response *CreateKafkaRechargeResponse, err error) {
    return c.CreateKafkaRechargeWithContext(context.Background(), request)
}

// CreateKafkaRecharge
// This API is used to create a Kafka data subscription task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateKafkaRechargeWithContext(ctx context.Context, request *CreateKafkaRechargeRequest) (response *CreateKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewCreateKafkaRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateKafkaRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogsetRequest() (request *CreateLogsetRequest) {
    request = &CreateLogsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateLogset")
    
    
    return
}

func NewCreateLogsetResponse() (response *CreateLogsetResponse) {
    response = &CreateLogsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLogset
// This API is used to create a logset. The ID of the created logset is returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETCONFLICT = "FailedOperation.LogsetConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateLogset(request *CreateLogsetRequest) (response *CreateLogsetResponse, err error) {
    return c.CreateLogsetWithContext(context.Background(), request)
}

// CreateLogset
// This API is used to create a logset. The ID of the created logset is returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETCONFLICT = "FailedOperation.LogsetConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateLogsetWithContext(ctx context.Context, request *CreateLogsetRequest) (response *CreateLogsetResponse, err error) {
    if request == nil {
        request = NewCreateLogsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateLogset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogsetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMachineGroupRequest() (request *CreateMachineGroupRequest) {
    request = &CreateMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateMachineGroup")
    
    
    return
}

func NewCreateMachineGroupResponse() (response *CreateMachineGroupResponse) {
    response = &CreateMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMachineGroup
// This API is used to create a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MACHINEGROUP = "LimitExceeded.MachineGroup"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateMachineGroup(request *CreateMachineGroupRequest) (response *CreateMachineGroupResponse, err error) {
    return c.CreateMachineGroupWithContext(context.Background(), request)
}

// CreateMachineGroup
// This API is used to create a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MACHINEGROUP = "LimitExceeded.MachineGroup"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateMachineGroupWithContext(ctx context.Context, request *CreateMachineGroupRequest) (response *CreateMachineGroupResponse, err error) {
    if request == nil {
        request = NewCreateMachineGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateMachineGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMetricConfigRequest() (request *CreateMetricConfigRequest) {
    request = &CreateMetricConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateMetricConfig")
    
    
    return
}

func NewCreateMetricConfigResponse() (response *CreateMetricConfigResponse) {
    response = &CreateMetricConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMetricConfig
// This API is used to create metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMetricConfig(request *CreateMetricConfigRequest) (response *CreateMetricConfigResponse, err error) {
    return c.CreateMetricConfigWithContext(context.Background(), request)
}

// CreateMetricConfig
// This API is used to create metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMetricConfigWithContext(ctx context.Context, request *CreateMetricConfigRequest) (response *CreateMetricConfigResponse, err error) {
    if request == nil {
        request = NewCreateMetricConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateMetricConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMetricConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMetricConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMetricSubscribeRequest() (request *CreateMetricSubscribeRequest) {
    request = &CreateMetricSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateMetricSubscribe")
    
    
    return
}

func NewCreateMetricSubscribeResponse() (response *CreateMetricSubscribeResponse) {
    response = &CreateMetricSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMetricSubscribe
// This API is used to create metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMetricSubscribe(request *CreateMetricSubscribeRequest) (response *CreateMetricSubscribeResponse, err error) {
    return c.CreateMetricSubscribeWithContext(context.Background(), request)
}

// CreateMetricSubscribe
// This API is used to create metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMetricSubscribeWithContext(ctx context.Context, request *CreateMetricSubscribeRequest) (response *CreateMetricSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateMetricSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateMetricSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMetricSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMetricSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkApplicationRequest() (request *CreateNetworkApplicationRequest) {
    request = &CreateNetworkApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateNetworkApplication")
    
    
    return
}

func NewCreateNetworkApplicationResponse() (response *CreateNetworkApplicationResponse) {
    response = &CreateNetworkApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNetworkApplication
// This API is used to create a network application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNetworkApplication(request *CreateNetworkApplicationRequest) (response *CreateNetworkApplicationResponse, err error) {
    return c.CreateNetworkApplicationWithContext(context.Background(), request)
}

// CreateNetworkApplication
// This API is used to create a network application.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNetworkApplicationWithContext(ctx context.Context, request *CreateNetworkApplicationRequest) (response *CreateNetworkApplicationResponse, err error) {
    if request == nil {
        request = NewCreateNetworkApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateNetworkApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNoticeContentRequest() (request *CreateNoticeContentRequest) {
    request = &CreateNoticeContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateNoticeContent")
    
    
    return
}

func NewCreateNoticeContentResponse() (response *CreateNoticeContentResponse) {
    response = &CreateNoticeContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNoticeContent
// This API is used to create notification content.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateNoticeContent(request *CreateNoticeContentRequest) (response *CreateNoticeContentResponse, err error) {
    return c.CreateNoticeContentWithContext(context.Background(), request)
}

// CreateNoticeContent
// This API is used to create notification content.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateNoticeContentWithContext(ctx context.Context, request *CreateNoticeContentRequest) (response *CreateNoticeContentResponse, err error) {
    if request == nil {
        request = NewCreateNoticeContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateNoticeContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNoticeContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNoticeContentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRebuildIndexTaskRequest() (request *CreateRebuildIndexTaskRequest) {
    request = &CreateRebuildIndexTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateRebuildIndexTask")
    
    
    return
}

func NewCreateRebuildIndexTaskResponse() (response *CreateRebuildIndexTaskResponse) {
    response = &CreateRebuildIndexTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRebuildIndexTask
// This API is used to creates rebuild index tasks.
//
// Note:
//
// -A single log topic allows only one index reconstruction task at a time and can have up to 10 rebuild index task records. Delete task records that are no longer needed before creating an index task.
//
// -Logs within the same time range only allow rebuilding indexes once. Delete previous task records before rebuilding again.
//
// -Deleting a rebuild index task record restores the index data before rebuilding an index.
//
// -The log write traffic of the selected time range cannot exceed 5TB.
//
// -The index rebuilding time range is subject to the log time. When the log uploading time has a deviation exceeding one hour from the index rebuilding time range (for example, reporting a new log at 16:00 for 02:00 to CLS and rebuilding the log index for 00:00–12:00), the logs will not be rebuilt and cannot be retrieved subsequently. When reporting a new log to the reconstructed log time range, it will not be rebuilt and cannot be retrieved subsequently.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateRebuildIndexTask(request *CreateRebuildIndexTaskRequest) (response *CreateRebuildIndexTaskResponse, err error) {
    return c.CreateRebuildIndexTaskWithContext(context.Background(), request)
}

// CreateRebuildIndexTask
// This API is used to creates rebuild index tasks.
//
// Note:
//
// -A single log topic allows only one index reconstruction task at a time and can have up to 10 rebuild index task records. Delete task records that are no longer needed before creating an index task.
//
// -Logs within the same time range only allow rebuilding indexes once. Delete previous task records before rebuilding again.
//
// -Deleting a rebuild index task record restores the index data before rebuilding an index.
//
// -The log write traffic of the selected time range cannot exceed 5TB.
//
// -The index rebuilding time range is subject to the log time. When the log uploading time has a deviation exceeding one hour from the index rebuilding time range (for example, reporting a new log at 16:00 for 02:00 to CLS and rebuilding the log index for 00:00–12:00), the logs will not be rebuilt and cannot be retrieved subsequently. When reporting a new log to the reconstructed log time range, it will not be rebuilt and cannot be retrieved subsequently.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateRebuildIndexTaskWithContext(ctx context.Context, request *CreateRebuildIndexTaskRequest) (response *CreateRebuildIndexTaskResponse, err error) {
    if request == nil {
        request = NewCreateRebuildIndexTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateRebuildIndexTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRebuildIndexTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRebuildIndexTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordingRuleTaskRequest() (request *CreateRecordingRuleTaskRequest) {
    request = &CreateRecordingRuleTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateRecordingRuleTask")
    
    
    return
}

func NewCreateRecordingRuleTaskResponse() (response *CreateRecordingRuleTaskResponse) {
    response = &CreateRecordingRuleTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRecordingRuleTask
// Creating a Metric Pre-Aggregation Task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateRecordingRuleTask(request *CreateRecordingRuleTaskRequest) (response *CreateRecordingRuleTaskResponse, err error) {
    return c.CreateRecordingRuleTaskWithContext(context.Background(), request)
}

// CreateRecordingRuleTask
// Creating a Metric Pre-Aggregation Task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateRecordingRuleTaskWithContext(ctx context.Context, request *CreateRecordingRuleTaskRequest) (response *CreateRecordingRuleTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecordingRuleTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateRecordingRuleTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordingRuleTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordingRuleTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordingRuleYamlTaskRequest() (request *CreateRecordingRuleYamlTaskRequest) {
    request = &CreateRecordingRuleYamlTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateRecordingRuleYamlTask")
    
    
    return
}

func NewCreateRecordingRuleYamlTaskResponse() (response *CreateRecordingRuleYamlTaskResponse) {
    response = &CreateRecordingRuleYamlTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRecordingRuleYamlTask
// Creating a Metric Pre-Aggregation Task Through a YAML File
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateRecordingRuleYamlTask(request *CreateRecordingRuleYamlTaskRequest) (response *CreateRecordingRuleYamlTaskResponse, err error) {
    return c.CreateRecordingRuleYamlTaskWithContext(context.Background(), request)
}

// CreateRecordingRuleYamlTask
// Creating a Metric Pre-Aggregation Task Through a YAML File
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateRecordingRuleYamlTaskWithContext(ctx context.Context, request *CreateRecordingRuleYamlTaskRequest) (response *CreateRecordingRuleYamlTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecordingRuleYamlTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateRecordingRuleYamlTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordingRuleYamlTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordingRuleYamlTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScheduledSqlRequest() (request *CreateScheduledSqlRequest) {
    request = &CreateScheduledSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateScheduledSql")
    
    
    return
}

func NewCreateScheduledSqlResponse() (response *CreateScheduledSqlResponse) {
    response = &CreateScheduledSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScheduledSql
// This API is used to create a scheduled SQL analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateScheduledSql(request *CreateScheduledSqlRequest) (response *CreateScheduledSqlResponse, err error) {
    return c.CreateScheduledSqlWithContext(context.Background(), request)
}

// CreateScheduledSql
// This API is used to create a scheduled SQL analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateScheduledSqlWithContext(ctx context.Context, request *CreateScheduledSqlRequest) (response *CreateScheduledSqlResponse, err error) {
    if request == nil {
        request = NewCreateScheduledSqlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateScheduledSql")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScheduledSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScheduledSqlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSearchViewRequest() (request *CreateSearchViewRequest) {
    request = &CreateSearchViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateSearchView")
    
    
    return
}

func NewCreateSearchViewResponse() (response *CreateSearchViewResponse) {
    response = &CreateSearchViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSearchView
// Create a query view
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateSearchView(request *CreateSearchViewRequest) (response *CreateSearchViewResponse, err error) {
    return c.CreateSearchViewWithContext(context.Background(), request)
}

// CreateSearchView
// Create a query view
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateSearchViewWithContext(ctx context.Context, request *CreateSearchViewRequest) (response *CreateSearchViewResponse, err error) {
    if request == nil {
        request = NewCreateSearchViewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateSearchView")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSearchView require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSearchViewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateShipperRequest() (request *CreateShipperRequest) {
    request = &CreateShipperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateShipper")
    
    
    return
}

func NewCreateShipperResponse() (response *CreateShipperResponse) {
    response = &CreateShipperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateShipper
// This API is used to create a task to ship to COS. Note: To use this API, you need to check whether you have configured the role and permission for shipping to COS. If not, see **Viewing and Configuring Shipping Authorization** at https://intl.cloud.tencent.com/document/product/614/71623.?from_cn_redirect=1
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SHIPPERCONFLICT = "InvalidParameter.ShipperConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SHIPPER = "LimitExceeded.Shipper"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateShipper(request *CreateShipperRequest) (response *CreateShipperResponse, err error) {
    return c.CreateShipperWithContext(context.Background(), request)
}

// CreateShipper
// This API is used to create a task to ship to COS. Note: To use this API, you need to check whether you have configured the role and permission for shipping to COS. If not, see **Viewing and Configuring Shipping Authorization** at https://intl.cloud.tencent.com/document/product/614/71623.?from_cn_redirect=1
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SHIPPERCONFLICT = "InvalidParameter.ShipperConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SHIPPER = "LimitExceeded.Shipper"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateShipperWithContext(ctx context.Context, request *CreateShipperRequest) (response *CreateShipperResponse, err error) {
    if request == nil {
        request = NewCreateShipperRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateShipper")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateShipperResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSplunkDeliverRequest() (request *CreateSplunkDeliverRequest) {
    request = &CreateSplunkDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateSplunkDeliver")
    
    
    return
}

func NewCreateSplunkDeliverResponse() (response *CreateSplunkDeliverResponse) {
    response = &CreateSplunkDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSplunkDeliver
// Create a Splunk delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSplunkDeliver(request *CreateSplunkDeliverRequest) (response *CreateSplunkDeliverResponse, err error) {
    return c.CreateSplunkDeliverWithContext(context.Background(), request)
}

// CreateSplunkDeliver
// Create a Splunk delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSplunkDeliverWithContext(ctx context.Context, request *CreateSplunkDeliverRequest) (response *CreateSplunkDeliverResponse, err error) {
    if request == nil {
        request = NewCreateSplunkDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateSplunkDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSplunkDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSplunkDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateTopic")
    
    
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTopic
// This API is used to create logs or metric topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TOPICCREATING = "FailedOperation.TopicCreating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_BILLINGTOPICOUTOFLIMIT = "LimitExceeded.BillingTopicOutOfLimit"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TOPIC = "LimitExceeded.Topic"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// This API is used to create logs or metric topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TOPICCREATING = "FailedOperation.TopicCreating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_BILLINGTOPICOUTOFLIMIT = "LimitExceeded.BillingTopicOutOfLimit"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TOPIC = "LimitExceeded.Topic"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebCallbackRequest() (request *CreateWebCallbackRequest) {
    request = &CreateWebCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateWebCallback")
    
    
    return
}

func NewCreateWebCallbackResponse() (response *CreateWebCallbackResponse) {
    response = &CreateWebCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWebCallback
// This API is used to create alarm channel callback configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateWebCallback(request *CreateWebCallbackRequest) (response *CreateWebCallbackResponse, err error) {
    return c.CreateWebCallbackWithContext(context.Background(), request)
}

// CreateWebCallback
// This API is used to create alarm channel callback configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateWebCallbackWithContext(ctx context.Context, request *CreateWebCallbackRequest) (response *CreateWebCallbackResponse, err error) {
    if request == nil {
        request = NewCreateWebCallbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "CreateWebCallback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmRequest() (request *DeleteAlarmRequest) {
    request = &DeleteAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarm")
    
    
    return
}

func NewDeleteAlarmResponse() (response *DeleteAlarmResponse) {
    response = &DeleteAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarm
// This API is used to delete an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
func (c *Client) DeleteAlarm(request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
    return c.DeleteAlarmWithContext(context.Background(), request)
}

// DeleteAlarm
// This API is used to delete an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
func (c *Client) DeleteAlarmWithContext(ctx context.Context, request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmNoticeRequest() (request *DeleteAlarmNoticeRequest) {
    request = &DeleteAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarmNotice")
    
    
    return
}

func NewDeleteAlarmNoticeResponse() (response *DeleteAlarmNoticeResponse) {
    response = &DeleteAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmNotice
// This API is used to delete a notification group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteAlarmNotice(request *DeleteAlarmNoticeRequest) (response *DeleteAlarmNoticeResponse, err error) {
    return c.DeleteAlarmNoticeWithContext(context.Background(), request)
}

// DeleteAlarmNotice
// This API is used to delete a notification group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteAlarmNoticeWithContext(ctx context.Context, request *DeleteAlarmNoticeRequest) (response *DeleteAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmShieldRequest() (request *DeleteAlarmShieldRequest) {
    request = &DeleteAlarmShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarmShield")
    
    
    return
}

func NewDeleteAlarmShieldResponse() (response *DeleteAlarmShieldResponse) {
    response = &DeleteAlarmShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmShield
// This API is used to delete an alarm blocking rule. When the alarm blocking rule is active or invalid, it cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteAlarmShield(request *DeleteAlarmShieldRequest) (response *DeleteAlarmShieldResponse, err error) {
    return c.DeleteAlarmShieldWithContext(context.Background(), request)
}

// DeleteAlarmShield
// This API is used to delete an alarm blocking rule. When the alarm blocking rule is active or invalid, it cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteAlarmShieldWithContext(ctx context.Context, request *DeleteAlarmShieldRequest) (response *DeleteAlarmShieldResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmShieldRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteAlarmShield")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmShield require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmShieldResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudProductLogCollectionRequest() (request *DeleteCloudProductLogCollectionRequest) {
    request = &DeleteCloudProductLogCollectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteCloudProductLogCollection")
    
    
    return
}

func NewDeleteCloudProductLogCollectionResponse() (response *DeleteCloudProductLogCollectionResponse) {
    response = &DeleteCloudProductLogCollectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudProductLogCollection
// Cloud product integration uses internal APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDPRODUCTINVOCATIONERROR = "FailedOperation.CloudProductInvocationError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCloudProductLogCollection(request *DeleteCloudProductLogCollectionRequest) (response *DeleteCloudProductLogCollectionResponse, err error) {
    return c.DeleteCloudProductLogCollectionWithContext(context.Background(), request)
}

// DeleteCloudProductLogCollection
// Cloud product integration uses internal APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDPRODUCTINVOCATIONERROR = "FailedOperation.CloudProductInvocationError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCloudProductLogCollectionWithContext(ctx context.Context, request *DeleteCloudProductLogCollectionRequest) (response *DeleteCloudProductLogCollectionResponse, err error) {
    if request == nil {
        request = NewDeleteCloudProductLogCollectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteCloudProductLogCollection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudProductLogCollection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudProductLogCollectionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
    request = &DeleteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConfig")
    
    
    return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
    response = &DeleteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfig
// This API is used to delete collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    return c.DeleteConfigWithContext(context.Background(), request)
}

// DeleteConfig
// This API is used to delete collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DeleteConfigWithContext(ctx context.Context, request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigFromMachineGroupRequest() (request *DeleteConfigFromMachineGroupRequest) {
    request = &DeleteConfigFromMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigFromMachineGroup")
    
    
    return
}

func NewDeleteConfigFromMachineGroupResponse() (response *DeleteConfigFromMachineGroupResponse) {
    response = &DeleteConfigFromMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigFromMachineGroup
// This API is used to delete the collection configuration applied to a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteConfigFromMachineGroup(request *DeleteConfigFromMachineGroupRequest) (response *DeleteConfigFromMachineGroupResponse, err error) {
    return c.DeleteConfigFromMachineGroupWithContext(context.Background(), request)
}

// DeleteConfigFromMachineGroup
// This API is used to delete the collection configuration applied to a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteConfigFromMachineGroupWithContext(ctx context.Context, request *DeleteConfigFromMachineGroupRequest) (response *DeleteConfigFromMachineGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConfigFromMachineGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteConfigFromMachineGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigFromMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigFromMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsoleRequest() (request *DeleteConsoleRequest) {
    request = &DeleteConsoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConsole")
    
    
    return
}

func NewDeleteConsoleResponse() (response *DeleteConsoleResponse) {
    response = &DeleteConsoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsole
// This API is used to delete the DataSight Console
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConsole(request *DeleteConsoleRequest) (response *DeleteConsoleResponse, err error) {
    return c.DeleteConsoleWithContext(context.Background(), request)
}

// DeleteConsole
// This API is used to delete the DataSight Console
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConsoleWithContext(ctx context.Context, request *DeleteConsoleRequest) (response *DeleteConsoleResponse, err error) {
    if request == nil {
        request = NewDeleteConsoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteConsole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsumerRequest() (request *DeleteConsumerRequest) {
    request = &DeleteConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConsumer")
    
    
    return
}

func NewDeleteConsumerResponse() (response *DeleteConsumerResponse) {
    response = &DeleteConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsumer
// Deleting a CKafka Delivery Task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConsumer(request *DeleteConsumerRequest) (response *DeleteConsumerResponse, err error) {
    return c.DeleteConsumerWithContext(context.Background(), request)
}

// DeleteConsumer
// Deleting a CKafka Delivery Task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConsumerWithContext(ctx context.Context, request *DeleteConsumerRequest) (response *DeleteConsumerResponse, err error) {
    if request == nil {
        request = NewDeleteConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
    request = &DeleteConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConsumerGroup")
    
    
    return
}

func NewDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
    response = &DeleteConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsumerGroup
// Delete consumer groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    return c.DeleteConsumerGroupWithContext(context.Background(), request)
}

// DeleteConsumerGroup
// Delete consumer groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConsumerGroupWithContext(ctx context.Context, request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCosRechargeRequest() (request *DeleteCosRechargeRequest) {
    request = &DeleteCosRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteCosRecharge")
    
    
    return
}

func NewDeleteCosRechargeResponse() (response *DeleteCosRechargeResponse) {
    response = &DeleteCosRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCosRecharge
// This API is used to delete a cos Import Task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteCosRecharge(request *DeleteCosRechargeRequest) (response *DeleteCosRechargeResponse, err error) {
    return c.DeleteCosRechargeWithContext(context.Background(), request)
}

// DeleteCosRecharge
// This API is used to delete a cos Import Task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteCosRechargeWithContext(ctx context.Context, request *DeleteCosRechargeRequest) (response *DeleteCosRechargeResponse, err error) {
    if request == nil {
        request = NewDeleteCosRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteCosRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCosRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCosRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDashboardRequest() (request *DeleteDashboardRequest) {
    request = &DeleteDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteDashboard")
    
    
    return
}

func NewDeleteDashboardResponse() (response *DeleteDashboardResponse) {
    response = &DeleteDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDashboard
// This API is used to delete dashboards.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteDashboard(request *DeleteDashboardRequest) (response *DeleteDashboardResponse, err error) {
    return c.DeleteDashboardWithContext(context.Background(), request)
}

// DeleteDashboard
// This API is used to delete dashboards.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteDashboardWithContext(ctx context.Context, request *DeleteDashboardRequest) (response *DeleteDashboardResponse, err error) {
    if request == nil {
        request = NewDeleteDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDashboardSubscribeRequest() (request *DeleteDashboardSubscribeRequest) {
    request = &DeleteDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteDashboardSubscribe")
    
    
    return
}

func NewDeleteDashboardSubscribeResponse() (response *DeleteDashboardSubscribeResponse) {
    response = &DeleteDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDashboardSubscribe
// This API is used to delete dashboard subscriptions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) DeleteDashboardSubscribe(request *DeleteDashboardSubscribeRequest) (response *DeleteDashboardSubscribeResponse, err error) {
    return c.DeleteDashboardSubscribeWithContext(context.Background(), request)
}

// DeleteDashboardSubscribe
// This API is used to delete dashboard subscriptions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) DeleteDashboardSubscribeWithContext(ctx context.Context, request *DeleteDashboardSubscribeRequest) (response *DeleteDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewDeleteDashboardSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteDashboardSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataTransformRequest() (request *DeleteDataTransformRequest) {
    request = &DeleteDataTransformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteDataTransform")
    
    
    return
}

func NewDeleteDataTransformResponse() (response *DeleteDataTransformResponse) {
    response = &DeleteDataTransformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataTransform
// This API is used to delete a data processing task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteDataTransform(request *DeleteDataTransformRequest) (response *DeleteDataTransformResponse, err error) {
    return c.DeleteDataTransformWithContext(context.Background(), request)
}

// DeleteDataTransform
// This API is used to delete a data processing task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteDataTransformWithContext(ctx context.Context, request *DeleteDataTransformRequest) (response *DeleteDataTransformResponse, err error) {
    if request == nil {
        request = NewDeleteDataTransformRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteDataTransform")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataTransformResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDlcDeliverRequest() (request *DeleteDlcDeliverRequest) {
    request = &DeleteDlcDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteDlcDeliver")
    
    
    return
}

func NewDeleteDlcDeliverResponse() (response *DeleteDlcDeliverResponse) {
    response = &DeleteDlcDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDlcDeliver
// Delete a DLC delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDlcDeliver(request *DeleteDlcDeliverRequest) (response *DeleteDlcDeliverResponse, err error) {
    return c.DeleteDlcDeliverWithContext(context.Background(), request)
}

// DeleteDlcDeliver
// Delete a DLC delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDlcDeliverWithContext(ctx context.Context, request *DeleteDlcDeliverRequest) (response *DeleteDlcDeliverResponse, err error) {
    if request == nil {
        request = NewDeleteDlcDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteDlcDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDlcDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDlcDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEsRechargeRequest() (request *DeleteEsRechargeRequest) {
    request = &DeleteEsRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteEsRecharge")
    
    
    return
}

func NewDeleteEsRechargeResponse() (response *DeleteEsRechargeResponse) {
    response = &DeleteEsRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEsRecharge
// Delete es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEsRecharge(request *DeleteEsRechargeRequest) (response *DeleteEsRechargeResponse, err error) {
    return c.DeleteEsRechargeWithContext(context.Background(), request)
}

// DeleteEsRecharge
// Delete es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEsRechargeWithContext(ctx context.Context, request *DeleteEsRechargeRequest) (response *DeleteEsRechargeResponse, err error) {
    if request == nil {
        request = NewDeleteEsRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteEsRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEsRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEsRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteExportRequest() (request *DeleteExportRequest) {
    request = &DeleteExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteExport")
    
    
    return
}

func NewDeleteExportResponse() (response *DeleteExportResponse) {
    response = &DeleteExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteExport
// This API is used to delete a log download task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
func (c *Client) DeleteExport(request *DeleteExportRequest) (response *DeleteExportResponse, err error) {
    return c.DeleteExportWithContext(context.Background(), request)
}

// DeleteExport
// This API is used to delete a log download task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
func (c *Client) DeleteExportWithContext(ctx context.Context, request *DeleteExportRequest) (response *DeleteExportResponse, err error) {
    if request == nil {
        request = NewDeleteExportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteExport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExportResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHostMetricConfigRequest() (request *DeleteHostMetricConfigRequest) {
    request = &DeleteHostMetricConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteHostMetricConfig")
    
    
    return
}

func NewDeleteHostMetricConfigResponse() (response *DeleteHostMetricConfigResponse) {
    response = &DeleteHostMetricConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHostMetricConfig
// Delete host metric collection configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHostMetricConfig(request *DeleteHostMetricConfigRequest) (response *DeleteHostMetricConfigResponse, err error) {
    return c.DeleteHostMetricConfigWithContext(context.Background(), request)
}

// DeleteHostMetricConfig
// Delete host metric collection configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHostMetricConfigWithContext(ctx context.Context, request *DeleteHostMetricConfigRequest) (response *DeleteHostMetricConfigResponse, err error) {
    if request == nil {
        request = NewDeleteHostMetricConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteHostMetricConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHostMetricConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHostMetricConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIndexRequest() (request *DeleteIndexRequest) {
    request = &DeleteIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteIndex")
    
    
    return
}

func NewDeleteIndexResponse() (response *DeleteIndexResponse) {
    response = &DeleteIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIndex
// This API is used to delete the index configuration of a log topic. After deleting, you cannot retrieve or query the collected logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    return c.DeleteIndexWithContext(context.Background(), request)
}

// DeleteIndex
// This API is used to delete the index configuration of a log topic. After deleting, you cannot retrieve or query the collected logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteIndexWithContext(ctx context.Context, request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    if request == nil {
        request = NewDeleteIndexRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteIndex")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKafkaRechargeRequest() (request *DeleteKafkaRechargeRequest) {
    request = &DeleteKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteKafkaRecharge")
    
    
    return
}

func NewDeleteKafkaRechargeResponse() (response *DeleteKafkaRechargeResponse) {
    response = &DeleteKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKafkaRecharge
// This API is used to delete a Kafka data subscription task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteKafkaRecharge(request *DeleteKafkaRechargeRequest) (response *DeleteKafkaRechargeResponse, err error) {
    return c.DeleteKafkaRechargeWithContext(context.Background(), request)
}

// DeleteKafkaRecharge
// This API is used to delete a Kafka data subscription task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteKafkaRechargeWithContext(ctx context.Context, request *DeleteKafkaRechargeRequest) (response *DeleteKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewDeleteKafkaRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteKafkaRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogsetRequest() (request *DeleteLogsetRequest) {
    request = &DeleteLogsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteLogset")
    
    
    return
}

func NewDeleteLogsetResponse() (response *DeleteLogsetResponse) {
    response = &DeleteLogsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogset
// This API is used to delete a logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETNOTEMPTY = "FailedOperation.LogsetNotEmpty"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) DeleteLogset(request *DeleteLogsetRequest) (response *DeleteLogsetResponse, err error) {
    return c.DeleteLogsetWithContext(context.Background(), request)
}

// DeleteLogset
// This API is used to delete a logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETNOTEMPTY = "FailedOperation.LogsetNotEmpty"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) DeleteLogsetWithContext(ctx context.Context, request *DeleteLogsetRequest) (response *DeleteLogsetResponse, err error) {
    if request == nil {
        request = NewDeleteLogsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteLogset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogsetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineGroupRequest() (request *DeleteMachineGroupRequest) {
    request = &DeleteMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroup")
    
    
    return
}

func NewDeleteMachineGroupResponse() (response *DeleteMachineGroupResponse) {
    response = &DeleteMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMachineGroup
// This API is used to delete a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteMachineGroup(request *DeleteMachineGroupRequest) (response *DeleteMachineGroupResponse, err error) {
    return c.DeleteMachineGroupWithContext(context.Background(), request)
}

// DeleteMachineGroup
// This API is used to delete a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteMachineGroupWithContext(ctx context.Context, request *DeleteMachineGroupRequest) (response *DeleteMachineGroupResponse, err error) {
    if request == nil {
        request = NewDeleteMachineGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteMachineGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineGroupInfoRequest() (request *DeleteMachineGroupInfoRequest) {
    request = &DeleteMachineGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroupInfo")
    
    
    return
}

func NewDeleteMachineGroupInfoResponse() (response *DeleteMachineGroupInfoResponse) {
    response = &DeleteMachineGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMachineGroupInfo
// This API is used to delete machine group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteMachineGroupInfo(request *DeleteMachineGroupInfoRequest) (response *DeleteMachineGroupInfoResponse, err error) {
    return c.DeleteMachineGroupInfoWithContext(context.Background(), request)
}

// DeleteMachineGroupInfo
// This API is used to delete machine group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteMachineGroupInfoWithContext(ctx context.Context, request *DeleteMachineGroupInfoRequest) (response *DeleteMachineGroupInfoResponse, err error) {
    if request == nil {
        request = NewDeleteMachineGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteMachineGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMetricConfigRequest() (request *DeleteMetricConfigRequest) {
    request = &DeleteMetricConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteMetricConfig")
    
    
    return
}

func NewDeleteMetricConfigResponse() (response *DeleteMetricConfigResponse) {
    response = &DeleteMetricConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMetricConfig
// This API is used to delete metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMetricConfig(request *DeleteMetricConfigRequest) (response *DeleteMetricConfigResponse, err error) {
    return c.DeleteMetricConfigWithContext(context.Background(), request)
}

// DeleteMetricConfig
// This API is used to delete metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMetricConfigWithContext(ctx context.Context, request *DeleteMetricConfigRequest) (response *DeleteMetricConfigResponse, err error) {
    if request == nil {
        request = NewDeleteMetricConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteMetricConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMetricConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMetricConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMetricSubscribeRequest() (request *DeleteMetricSubscribeRequest) {
    request = &DeleteMetricSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteMetricSubscribe")
    
    
    return
}

func NewDeleteMetricSubscribeResponse() (response *DeleteMetricSubscribeResponse) {
    response = &DeleteMetricSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMetricSubscribe
// This API is used to delete metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMetricSubscribe(request *DeleteMetricSubscribeRequest) (response *DeleteMetricSubscribeResponse, err error) {
    return c.DeleteMetricSubscribeWithContext(context.Background(), request)
}

// DeleteMetricSubscribe
// This API is used to delete metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMetricSubscribeWithContext(ctx context.Context, request *DeleteMetricSubscribeRequest) (response *DeleteMetricSubscribeResponse, err error) {
    if request == nil {
        request = NewDeleteMetricSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteMetricSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMetricSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMetricSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkApplicationRequest() (request *DeleteNetworkApplicationRequest) {
    request = &DeleteNetworkApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteNetworkApplication")
    
    
    return
}

func NewDeleteNetworkApplicationResponse() (response *DeleteNetworkApplicationResponse) {
    response = &DeleteNetworkApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNetworkApplication
// Delete a web application
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNetworkApplication(request *DeleteNetworkApplicationRequest) (response *DeleteNetworkApplicationResponse, err error) {
    return c.DeleteNetworkApplicationWithContext(context.Background(), request)
}

// DeleteNetworkApplication
// Delete a web application
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNetworkApplicationWithContext(ctx context.Context, request *DeleteNetworkApplicationRequest) (response *DeleteNetworkApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteNetworkApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetworkApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNetworkApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNoticeContentRequest() (request *DeleteNoticeContentRequest) {
    request = &DeleteNoticeContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteNoticeContent")
    
    
    return
}

func NewDeleteNoticeContentResponse() (response *DeleteNoticeContentResponse) {
    response = &DeleteNoticeContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNoticeContent
// This API is used to delete notification content configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteNoticeContent(request *DeleteNoticeContentRequest) (response *DeleteNoticeContentResponse, err error) {
    return c.DeleteNoticeContentWithContext(context.Background(), request)
}

// DeleteNoticeContent
// This API is used to delete notification content configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteNoticeContentWithContext(ctx context.Context, request *DeleteNoticeContentRequest) (response *DeleteNoticeContentResponse, err error) {
    if request == nil {
        request = NewDeleteNoticeContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteNoticeContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNoticeContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNoticeContentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordingRuleTaskRequest() (request *DeleteRecordingRuleTaskRequest) {
    request = &DeleteRecordingRuleTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteRecordingRuleTask")
    
    
    return
}

func NewDeleteRecordingRuleTaskResponse() (response *DeleteRecordingRuleTaskResponse) {
    response = &DeleteRecordingRuleTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordingRuleTask
// This API is used to delete a pre-aggregation analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteRecordingRuleTask(request *DeleteRecordingRuleTaskRequest) (response *DeleteRecordingRuleTaskResponse, err error) {
    return c.DeleteRecordingRuleTaskWithContext(context.Background(), request)
}

// DeleteRecordingRuleTask
// This API is used to delete a pre-aggregation analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteRecordingRuleTaskWithContext(ctx context.Context, request *DeleteRecordingRuleTaskRequest) (response *DeleteRecordingRuleTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordingRuleTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteRecordingRuleTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordingRuleTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordingRuleTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordingRuleYamlTaskRequest() (request *DeleteRecordingRuleYamlTaskRequest) {
    request = &DeleteRecordingRuleYamlTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteRecordingRuleYamlTask")
    
    
    return
}

func NewDeleteRecordingRuleYamlTaskResponse() (response *DeleteRecordingRuleYamlTaskResponse) {
    response = &DeleteRecordingRuleYamlTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordingRuleYamlTask
// This API is used to delete the pre-aggregation task in yaml.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteRecordingRuleYamlTask(request *DeleteRecordingRuleYamlTaskRequest) (response *DeleteRecordingRuleYamlTaskResponse, err error) {
    return c.DeleteRecordingRuleYamlTaskWithContext(context.Background(), request)
}

// DeleteRecordingRuleYamlTask
// This API is used to delete the pre-aggregation task in yaml.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteRecordingRuleYamlTaskWithContext(ctx context.Context, request *DeleteRecordingRuleYamlTaskRequest) (response *DeleteRecordingRuleYamlTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordingRuleYamlTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteRecordingRuleYamlTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordingRuleYamlTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordingRuleYamlTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScheduledSqlRequest() (request *DeleteScheduledSqlRequest) {
    request = &DeleteScheduledSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteScheduledSql")
    
    
    return
}

func NewDeleteScheduledSqlResponse() (response *DeleteScheduledSqlResponse) {
    response = &DeleteScheduledSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScheduledSql
// This API is used to delete a scheduled SQL analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteScheduledSql(request *DeleteScheduledSqlRequest) (response *DeleteScheduledSqlResponse, err error) {
    return c.DeleteScheduledSqlWithContext(context.Background(), request)
}

// DeleteScheduledSql
// This API is used to delete a scheduled SQL analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteScheduledSqlWithContext(ctx context.Context, request *DeleteScheduledSqlRequest) (response *DeleteScheduledSqlResponse, err error) {
    if request == nil {
        request = NewDeleteScheduledSqlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteScheduledSql")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScheduledSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScheduledSqlResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSearchViewRequest() (request *DeleteSearchViewRequest) {
    request = &DeleteSearchViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteSearchView")
    
    
    return
}

func NewDeleteSearchViewResponse() (response *DeleteSearchViewResponse) {
    response = &DeleteSearchViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSearchView
// This API is used to delete a query view.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteSearchView(request *DeleteSearchViewRequest) (response *DeleteSearchViewResponse, err error) {
    return c.DeleteSearchViewWithContext(context.Background(), request)
}

// DeleteSearchView
// This API is used to delete a query view.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteSearchViewWithContext(ctx context.Context, request *DeleteSearchViewRequest) (response *DeleteSearchViewResponse, err error) {
    if request == nil {
        request = NewDeleteSearchViewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteSearchView")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSearchView require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSearchViewResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShipperRequest() (request *DeleteShipperRequest) {
    request = &DeleteShipperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteShipper")
    
    
    return
}

func NewDeleteShipperResponse() (response *DeleteShipperResponse) {
    response = &DeleteShipperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShipper
// This API is used to delete a COS shipping task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DeleteShipper(request *DeleteShipperRequest) (response *DeleteShipperResponse, err error) {
    return c.DeleteShipperWithContext(context.Background(), request)
}

// DeleteShipper
// This API is used to delete a COS shipping task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DeleteShipperWithContext(ctx context.Context, request *DeleteShipperRequest) (response *DeleteShipperResponse, err error) {
    if request == nil {
        request = NewDeleteShipperRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteShipper")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShipperResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSplunkDeliverRequest() (request *DeleteSplunkDeliverRequest) {
    request = &DeleteSplunkDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteSplunkDeliver")
    
    
    return
}

func NewDeleteSplunkDeliverResponse() (response *DeleteSplunkDeliverResponse) {
    response = &DeleteSplunkDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSplunkDeliver
// Delete a Splunk delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSplunkDeliver(request *DeleteSplunkDeliverRequest) (response *DeleteSplunkDeliverResponse, err error) {
    return c.DeleteSplunkDeliverWithContext(context.Background(), request)
}

// DeleteSplunkDeliver
// Delete a Splunk delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSplunkDeliverWithContext(ctx context.Context, request *DeleteSplunkDeliverRequest) (response *DeleteSplunkDeliverResponse, err error) {
    if request == nil {
        request = NewDeleteSplunkDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteSplunkDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSplunkDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSplunkDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopic
// This API is used to delete logs or metric topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASDATAFORMTASK = "OperationDenied.TopicHasDataFormTask"
//  OPERATIONDENIED_TOPICHASDELIVERFUNCTION = "OperationDenied.TopicHasDeliverFunction"
//  OPERATIONDENIED_TOPICHASEXTERNALDATASOURCECONFIG = "OperationDenied.TopicHasExternalDatasourceConfig"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// This API is used to delete logs or metric topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASDATAFORMTASK = "OperationDenied.TopicHasDataFormTask"
//  OPERATIONDENIED_TOPICHASDELIVERFUNCTION = "OperationDenied.TopicHasDeliverFunction"
//  OPERATIONDENIED_TOPICHASEXTERNALDATASOURCECONFIG = "OperationDenied.TopicHasExternalDatasourceConfig"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebCallbackRequest() (request *DeleteWebCallbackRequest) {
    request = &DeleteWebCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteWebCallback")
    
    
    return
}

func NewDeleteWebCallbackResponse() (response *DeleteWebCallbackResponse) {
    response = &DeleteWebCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWebCallback
// This API is used to delete alarm channel callback configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteWebCallback(request *DeleteWebCallbackRequest) (response *DeleteWebCallbackResponse, err error) {
    return c.DeleteWebCallbackWithContext(context.Background(), request)
}

// DeleteWebCallback
// This API is used to delete alarm channel callback configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteWebCallbackWithContext(ctx context.Context, request *DeleteWebCallbackRequest) (response *DeleteWebCallbackResponse, err error) {
    if request == nil {
        request = NewDeleteWebCallbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DeleteWebCallback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWebCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticesRequest() (request *DescribeAlarmNoticesRequest) {
    request = &DescribeAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmNotices")
    
    
    return
}

func NewDescribeAlarmNoticesResponse() (response *DescribeAlarmNoticesResponse) {
    response = &DescribeAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotices
// This API is used to get the notification group list.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlarmNotices(request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    return c.DescribeAlarmNoticesWithContext(context.Background(), request)
}

// DescribeAlarmNotices
// This API is used to get the notification group list.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlarmNoticesWithContext(ctx context.Context, request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeAlarmNotices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmShieldsRequest() (request *DescribeAlarmShieldsRequest) {
    request = &DescribeAlarmShieldsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmShields")
    
    
    return
}

func NewDescribeAlarmShieldsResponse() (response *DescribeAlarmShieldsResponse) {
    response = &DescribeAlarmShieldsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmShields
// This API is used to access alarm blocking configuration rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmShields(request *DescribeAlarmShieldsRequest) (response *DescribeAlarmShieldsResponse, err error) {
    return c.DescribeAlarmShieldsWithContext(context.Background(), request)
}

// DescribeAlarmShields
// This API is used to access alarm blocking configuration rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmShieldsWithContext(ctx context.Context, request *DescribeAlarmShieldsRequest) (response *DescribeAlarmShieldsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmShieldsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeAlarmShields")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmShields require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmShieldsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
    request = &DescribeAlarmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarms")
    
    
    return
}

func NewDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
    response = &DescribeAlarmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarms
// This API is used to get the alarm policy list.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    return c.DescribeAlarmsWithContext(context.Background(), request)
}

// DescribeAlarms
// This API is used to get the alarm policy list.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeAlarmsWithContext(ctx context.Context, request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeAlarms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertRecordHistoryRequest() (request *DescribeAlertRecordHistoryRequest) {
    request = &DescribeAlertRecordHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlertRecordHistory")
    
    
    return
}

func NewDescribeAlertRecordHistoryResponse() (response *DescribeAlertRecordHistoryResponse) {
    response = &DescribeAlertRecordHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertRecordHistory
// This API is used to get alarm records, such as today's uncleared alarms.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlertRecordHistory(request *DescribeAlertRecordHistoryRequest) (response *DescribeAlertRecordHistoryResponse, err error) {
    return c.DescribeAlertRecordHistoryWithContext(context.Background(), request)
}

// DescribeAlertRecordHistory
// This API is used to get alarm records, such as today's uncleared alarms.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlertRecordHistoryWithContext(ctx context.Context, request *DescribeAlertRecordHistoryRequest) (response *DescribeAlertRecordHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeAlertRecordHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeAlertRecordHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertRecordHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertRecordHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudProductLogTasksRequest() (request *DescribeCloudProductLogTasksRequest) {
    request = &DescribeCloudProductLogTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeCloudProductLogTasks")
    
    
    return
}

func NewDescribeCloudProductLogTasksResponse() (response *DescribeCloudProductLogTasksResponse) {
    response = &DescribeCloudProductLogTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudProductLogTasks
// Cloud product integration uses relevant APIs
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCloudProductLogTasks(request *DescribeCloudProductLogTasksRequest) (response *DescribeCloudProductLogTasksResponse, err error) {
    return c.DescribeCloudProductLogTasksWithContext(context.Background(), request)
}

// DescribeCloudProductLogTasks
// Cloud product integration uses relevant APIs
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCloudProductLogTasksWithContext(ctx context.Context, request *DescribeCloudProductLogTasksRequest) (response *DescribeCloudProductLogTasksResponse, err error) {
    if request == nil {
        request = NewDescribeCloudProductLogTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeCloudProductLogTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudProductLogTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudProductLogTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterBaseMetricConfigsRequest() (request *DescribeClusterBaseMetricConfigsRequest) {
    request = &DescribeClusterBaseMetricConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeClusterBaseMetricConfigs")
    
    
    return
}

func NewDescribeClusterBaseMetricConfigsResponse() (response *DescribeClusterBaseMetricConfigsResponse) {
    response = &DescribeClusterBaseMetricConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterBaseMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterBaseMetricConfigs(request *DescribeClusterBaseMetricConfigsRequest) (response *DescribeClusterBaseMetricConfigsResponse, err error) {
    return c.DescribeClusterBaseMetricConfigsWithContext(context.Background(), request)
}

// DescribeClusterBaseMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterBaseMetricConfigsWithContext(ctx context.Context, request *DescribeClusterBaseMetricConfigsRequest) (response *DescribeClusterBaseMetricConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterBaseMetricConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeClusterBaseMetricConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterBaseMetricConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterBaseMetricConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterMetricConfigsRequest() (request *DescribeClusterMetricConfigsRequest) {
    request = &DescribeClusterMetricConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeClusterMetricConfigs")
    
    
    return
}

func NewDescribeClusterMetricConfigsResponse() (response *DescribeClusterMetricConfigsResponse) {
    response = &DescribeClusterMetricConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterMetricConfigs(request *DescribeClusterMetricConfigsRequest) (response *DescribeClusterMetricConfigsResponse, err error) {
    return c.DescribeClusterMetricConfigsWithContext(context.Background(), request)
}

// DescribeClusterMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterMetricConfigsWithContext(ctx context.Context, request *DescribeClusterMetricConfigsRequest) (response *DescribeClusterMetricConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterMetricConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeClusterMetricConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterMetricConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterMetricConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigMachineGroupsRequest() (request *DescribeConfigMachineGroupsRequest) {
    request = &DescribeConfigMachineGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigMachineGroups")
    
    
    return
}

func NewDescribeConfigMachineGroupsResponse() (response *DescribeConfigMachineGroupsResponse) {
    response = &DescribeConfigMachineGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigMachineGroups
// This API is used to get the machine group bound to collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DescribeConfigMachineGroups(request *DescribeConfigMachineGroupsRequest) (response *DescribeConfigMachineGroupsResponse, err error) {
    return c.DescribeConfigMachineGroupsWithContext(context.Background(), request)
}

// DescribeConfigMachineGroups
// This API is used to get the machine group bound to collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DescribeConfigMachineGroupsWithContext(ctx context.Context, request *DescribeConfigMachineGroupsRequest) (response *DescribeConfigMachineGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigMachineGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConfigMachineGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigMachineGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigMachineGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
    request = &DescribeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigs")
    
    
    return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
    response = &DescribeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigs
// This API is used to get collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    return c.DescribeConfigsWithContext(context.Background(), request)
}

// DescribeConfigs
// This API is used to get collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigsWithContext(ctx context.Context, request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsolesRequest() (request *DescribeConsolesRequest) {
    request = &DescribeConsolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsoles")
    
    
    return
}

func NewDescribeConsolesResponse() (response *DescribeConsolesResponse) {
    response = &DescribeConsolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsoles
// Query the DataSight console instance list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeConsoles(request *DescribeConsolesRequest) (response *DescribeConsolesResponse, err error) {
    return c.DescribeConsolesWithContext(context.Background(), request)
}

// DescribeConsoles
// Query the DataSight console instance list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeConsolesWithContext(ctx context.Context, request *DescribeConsolesRequest) (response *DescribeConsolesResponse, err error) {
    if request == nil {
        request = NewDescribeConsolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConsoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerRequest() (request *DescribeConsumerRequest) {
    request = &DescribeConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumer")
    
    
    return
}

func NewDescribeConsumerResponse() (response *DescribeConsumerResponse) {
    response = &DescribeConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumer
// This API is used to query a shipping task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConsumer(request *DescribeConsumerRequest) (response *DescribeConsumerResponse, err error) {
    return c.DescribeConsumerWithContext(context.Background(), request)
}

// DescribeConsumer
// This API is used to query a shipping task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConsumerWithContext(ctx context.Context, request *DescribeConsumerRequest) (response *DescribeConsumerResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupsRequest() (request *DescribeConsumerGroupsRequest) {
    request = &DescribeConsumerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumerGroups")
    
    
    return
}

func NewDescribeConsumerGroupsResponse() (response *DescribeConsumerGroupsResponse) {
    response = &DescribeConsumerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerGroups
// This API is used to obtain the consumer group list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
    return c.DescribeConsumerGroupsWithContext(context.Background(), request)
}

// DescribeConsumerGroups
// This API is used to obtain the consumer group list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeConsumerGroupsWithContext(ctx context.Context, request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConsumerGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerOffsetsRequest() (request *DescribeConsumerOffsetsRequest) {
    request = &DescribeConsumerOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumerOffsets")
    
    
    return
}

func NewDescribeConsumerOffsetsResponse() (response *DescribeConsumerOffsetsResponse) {
    response = &DescribeConsumerOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerOffsets
// Obtaining the Consumer Group Point Information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeConsumerOffsets(request *DescribeConsumerOffsetsRequest) (response *DescribeConsumerOffsetsResponse, err error) {
    return c.DescribeConsumerOffsetsWithContext(context.Background(), request)
}

// DescribeConsumerOffsets
// Obtaining the Consumer Group Point Information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeConsumerOffsetsWithContext(ctx context.Context, request *DescribeConsumerOffsetsRequest) (response *DescribeConsumerOffsetsResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerOffsetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConsumerOffsets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerPreviewRequest() (request *DescribeConsumerPreviewRequest) {
    request = &DescribeConsumerPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumerPreview")
    
    
    return
}

func NewDescribeConsumerPreviewResponse() (response *DescribeConsumerPreviewResponse) {
    response = &DescribeConsumerPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerPreview
// This API is used to preview Kafka shipping data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConsumerPreview(request *DescribeConsumerPreviewRequest) (response *DescribeConsumerPreviewResponse, err error) {
    return c.DescribeConsumerPreviewWithContext(context.Background(), request)
}

// DescribeConsumerPreview
// This API is used to preview Kafka shipping data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConsumerPreviewWithContext(ctx context.Context, request *DescribeConsumerPreviewRequest) (response *DescribeConsumerPreviewResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConsumerPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumersRequest() (request *DescribeConsumersRequest) {
    request = &DescribeConsumersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumers")
    
    
    return
}

func NewDescribeConsumersResponse() (response *DescribeConsumersResponse) {
    response = &DescribeConsumersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumers
// This API is used to obtain the shipping rule information list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeConsumers(request *DescribeConsumersRequest) (response *DescribeConsumersResponse, err error) {
    return c.DescribeConsumersWithContext(context.Background(), request)
}

// DescribeConsumers
// This API is used to obtain the shipping rule information list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeConsumersWithContext(ctx context.Context, request *DescribeConsumersRequest) (response *DescribeConsumersResponse, err error) {
    if request == nil {
        request = NewDescribeConsumersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeConsumers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosRechargesRequest() (request *DescribeCosRechargesRequest) {
    request = &DescribeCosRechargesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeCosRecharges")
    
    
    return
}

func NewDescribeCosRechargesResponse() (response *DescribeCosRechargesResponse) {
    response = &DescribeCosRechargesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosRecharges
// This API is used to get COS import configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeCosRecharges(request *DescribeCosRechargesRequest) (response *DescribeCosRechargesResponse, err error) {
    return c.DescribeCosRechargesWithContext(context.Background(), request)
}

// DescribeCosRecharges
// This API is used to get COS import configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeCosRechargesWithContext(ctx context.Context, request *DescribeCosRechargesRequest) (response *DescribeCosRechargesResponse, err error) {
    if request == nil {
        request = NewDescribeCosRechargesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeCosRecharges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRecharges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRechargesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardSubscribesRequest() (request *DescribeDashboardSubscribesRequest) {
    request = &DescribeDashboardSubscribesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeDashboardSubscribes")
    
    
    return
}

func NewDescribeDashboardSubscribesResponse() (response *DescribeDashboardSubscribesResponse) {
    response = &DescribeDashboardSubscribesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDashboardSubscribes
// This API is used to obtain the dashboard subscription list, and supports paging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboardSubscribes(request *DescribeDashboardSubscribesRequest) (response *DescribeDashboardSubscribesResponse, err error) {
    return c.DescribeDashboardSubscribesWithContext(context.Background(), request)
}

// DescribeDashboardSubscribes
// This API is used to obtain the dashboard subscription list, and supports paging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboardSubscribesWithContext(ctx context.Context, request *DescribeDashboardSubscribesRequest) (response *DescribeDashboardSubscribesResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardSubscribesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeDashboardSubscribes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDashboardSubscribes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDashboardSubscribesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataTransformInfoRequest() (request *DescribeDataTransformInfoRequest) {
    request = &DescribeDataTransformInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformInfo")
    
    
    return
}

func NewDescribeDataTransformInfoResponse() (response *DescribeDataTransformInfoResponse) {
    response = &DescribeDataTransformInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataTransformInfo
// This API is used to get the basic information of data processing tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeDataTransformInfo(request *DescribeDataTransformInfoRequest) (response *DescribeDataTransformInfoResponse, err error) {
    return c.DescribeDataTransformInfoWithContext(context.Background(), request)
}

// DescribeDataTransformInfo
// This API is used to get the basic information of data processing tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeDataTransformInfoWithContext(ctx context.Context, request *DescribeDataTransformInfoRequest) (response *DescribeDataTransformInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataTransformInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeDataTransformInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataTransformInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataTransformInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDlcDeliversRequest() (request *DescribeDlcDeliversRequest) {
    request = &DescribeDlcDeliversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeDlcDelivers")
    
    
    return
}

func NewDescribeDlcDeliversResponse() (response *DescribeDlcDeliversResponse) {
    response = &DescribeDlcDeliversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDlcDelivers
// This API is used to search alarm channel callback configuration lists.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeDlcDelivers(request *DescribeDlcDeliversRequest) (response *DescribeDlcDeliversResponse, err error) {
    return c.DescribeDlcDeliversWithContext(context.Background(), request)
}

// DescribeDlcDelivers
// This API is used to search alarm channel callback configuration lists.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeDlcDeliversWithContext(ctx context.Context, request *DescribeDlcDeliversRequest) (response *DescribeDlcDeliversResponse, err error) {
    if request == nil {
        request = NewDescribeDlcDeliversRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeDlcDelivers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDlcDelivers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDlcDeliversResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEsRechargePreviewRequest() (request *DescribeEsRechargePreviewRequest) {
    request = &DescribeEsRechargePreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeEsRechargePreview")
    
    
    return
}

func NewDescribeEsRechargePreviewResponse() (response *DescribeEsRechargePreviewResponse) {
    response = &DescribeEsRechargePreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEsRechargePreview
// Import Preview
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEsRechargePreview(request *DescribeEsRechargePreviewRequest) (response *DescribeEsRechargePreviewResponse, err error) {
    return c.DescribeEsRechargePreviewWithContext(context.Background(), request)
}

// DescribeEsRechargePreview
// Import Preview
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEsRechargePreviewWithContext(ctx context.Context, request *DescribeEsRechargePreviewRequest) (response *DescribeEsRechargePreviewResponse, err error) {
    if request == nil {
        request = NewDescribeEsRechargePreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeEsRechargePreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEsRechargePreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEsRechargePreviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEsRechargesRequest() (request *DescribeEsRechargesRequest) {
    request = &DescribeEsRechargesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeEsRecharges")
    
    
    return
}

func NewDescribeEsRechargesResponse() (response *DescribeEsRechargesResponse) {
    response = &DescribeEsRechargesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEsRecharges
// Retrieve the es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEsRecharges(request *DescribeEsRechargesRequest) (response *DescribeEsRechargesResponse, err error) {
    return c.DescribeEsRechargesWithContext(context.Background(), request)
}

// DescribeEsRecharges
// Retrieve the es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEsRechargesWithContext(ctx context.Context, request *DescribeEsRechargesRequest) (response *DescribeEsRechargesResponse, err error) {
    if request == nil {
        request = NewDescribeEsRechargesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeEsRecharges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEsRecharges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEsRechargesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportsRequest() (request *DescribeExportsRequest) {
    request = &DescribeExportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeExports")
    
    
    return
}

func NewDescribeExportsResponse() (response *DescribeExportsResponse) {
    response = &DescribeExportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExports
// This API is used to get the list of log download tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExports(request *DescribeExportsRequest) (response *DescribeExportsResponse, err error) {
    return c.DescribeExportsWithContext(context.Background(), request)
}

// DescribeExports
// This API is used to get the list of log download tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExportsWithContext(ctx context.Context, request *DescribeExportsRequest) (response *DescribeExportsResponse, err error) {
    if request == nil {
        request = NewDescribeExportsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeExports")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostMetricConfigsRequest() (request *DescribeHostMetricConfigsRequest) {
    request = &DescribeHostMetricConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeHostMetricConfigs")
    
    
    return
}

func NewDescribeHostMetricConfigsResponse() (response *DescribeHostMetricConfigsResponse) {
    response = &DescribeHostMetricConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHostMetricConfigs(request *DescribeHostMetricConfigsRequest) (response *DescribeHostMetricConfigsResponse, err error) {
    return c.DescribeHostMetricConfigsWithContext(context.Background(), request)
}

// DescribeHostMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHostMetricConfigsWithContext(ctx context.Context, request *DescribeHostMetricConfigsRequest) (response *DescribeHostMetricConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeHostMetricConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeHostMetricConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostMetricConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostMetricConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexRequest() (request *DescribeIndexRequest) {
    request = &DescribeIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeIndex")
    
    
    return
}

func NewDescribeIndexResponse() (response *DescribeIndexResponse) {
    response = &DescribeIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIndex
// This API is used to get the index configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeIndex(request *DescribeIndexRequest) (response *DescribeIndexResponse, err error) {
    return c.DescribeIndexWithContext(context.Background(), request)
}

// DescribeIndex
// This API is used to get the index configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeIndexWithContext(ctx context.Context, request *DescribeIndexRequest) (response *DescribeIndexResponse, err error) {
    if request == nil {
        request = NewDescribeIndexRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeIndex")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaConsumerRequest() (request *DescribeKafkaConsumerRequest) {
    request = &DescribeKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumer")
    
    
    return
}

func NewDescribeKafkaConsumerResponse() (response *DescribeKafkaConsumerResponse) {
    response = &DescribeKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaConsumer
// This API is used to access Kafka protocol consumption information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumer(request *DescribeKafkaConsumerRequest) (response *DescribeKafkaConsumerResponse, err error) {
    return c.DescribeKafkaConsumerWithContext(context.Background(), request)
}

// DescribeKafkaConsumer
// This API is used to access Kafka protocol consumption information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerWithContext(ctx context.Context, request *DescribeKafkaConsumerRequest) (response *DescribeKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeKafkaConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaConsumerGroupDetailRequest() (request *DescribeKafkaConsumerGroupDetailRequest) {
    request = &DescribeKafkaConsumerGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumerGroupDetail")
    
    
    return
}

func NewDescribeKafkaConsumerGroupDetailResponse() (response *DescribeKafkaConsumerGroupDetailResponse) {
    response = &DescribeKafkaConsumerGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaConsumerGroupDetail
// Retrieve Kafka protocol consumption group details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerGroupDetail(request *DescribeKafkaConsumerGroupDetailRequest) (response *DescribeKafkaConsumerGroupDetailResponse, err error) {
    return c.DescribeKafkaConsumerGroupDetailWithContext(context.Background(), request)
}

// DescribeKafkaConsumerGroupDetail
// Retrieve Kafka protocol consumption group details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerGroupDetailWithContext(ctx context.Context, request *DescribeKafkaConsumerGroupDetailRequest) (response *DescribeKafkaConsumerGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaConsumerGroupDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeKafkaConsumerGroupDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaConsumerGroupDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaConsumerGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaConsumerGroupListRequest() (request *DescribeKafkaConsumerGroupListRequest) {
    request = &DescribeKafkaConsumerGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumerGroupList")
    
    
    return
}

func NewDescribeKafkaConsumerGroupListResponse() (response *DescribeKafkaConsumerGroupListResponse) {
    response = &DescribeKafkaConsumerGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaConsumerGroupList
// Retrieve the information list of Kafka protocol consumption groups
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerGroupList(request *DescribeKafkaConsumerGroupListRequest) (response *DescribeKafkaConsumerGroupListResponse, err error) {
    return c.DescribeKafkaConsumerGroupListWithContext(context.Background(), request)
}

// DescribeKafkaConsumerGroupList
// Retrieve the information list of Kafka protocol consumption groups
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerGroupListWithContext(ctx context.Context, request *DescribeKafkaConsumerGroupListRequest) (response *DescribeKafkaConsumerGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaConsumerGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeKafkaConsumerGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaConsumerGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaConsumerGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaConsumerPreviewRequest() (request *DescribeKafkaConsumerPreviewRequest) {
    request = &DescribeKafkaConsumerPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumerPreview")
    
    
    return
}

func NewDescribeKafkaConsumerPreviewResponse() (response *DescribeKafkaConsumerPreviewResponse) {
    response = &DescribeKafkaConsumerPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaConsumerPreview
// This API is used to preview the Kafka consumption.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerPreview(request *DescribeKafkaConsumerPreviewRequest) (response *DescribeKafkaConsumerPreviewResponse, err error) {
    return c.DescribeKafkaConsumerPreviewWithContext(context.Background(), request)
}

// DescribeKafkaConsumerPreview
// This API is used to preview the Kafka consumption.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerPreviewWithContext(ctx context.Context, request *DescribeKafkaConsumerPreviewRequest) (response *DescribeKafkaConsumerPreviewResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaConsumerPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeKafkaConsumerPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaConsumerPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaConsumerPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaConsumerTopicsRequest() (request *DescribeKafkaConsumerTopicsRequest) {
    request = &DescribeKafkaConsumerTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumerTopics")
    
    
    return
}

func NewDescribeKafkaConsumerTopicsResponse() (response *DescribeKafkaConsumerTopicsResponse) {
    response = &DescribeKafkaConsumerTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaConsumerTopics
// This API is used to obtain the topic information list of Kafka consumption.
//
// error code that may be returned:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerTopics(request *DescribeKafkaConsumerTopicsRequest) (response *DescribeKafkaConsumerTopicsResponse, err error) {
    return c.DescribeKafkaConsumerTopicsWithContext(context.Background(), request)
}

// DescribeKafkaConsumerTopics
// This API is used to obtain the topic information list of Kafka consumption.
//
// error code that may be returned:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerTopicsWithContext(ctx context.Context, request *DescribeKafkaConsumerTopicsRequest) (response *DescribeKafkaConsumerTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaConsumerTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeKafkaConsumerTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaConsumerTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaConsumerTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaRechargesRequest() (request *DescribeKafkaRechargesRequest) {
    request = &DescribeKafkaRechargesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaRecharges")
    
    
    return
}

func NewDescribeKafkaRechargesResponse() (response *DescribeKafkaRechargesResponse) {
    response = &DescribeKafkaRechargesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaRecharges
// This API is used to get the list of Kafka data subscription tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaRecharges(request *DescribeKafkaRechargesRequest) (response *DescribeKafkaRechargesResponse, err error) {
    return c.DescribeKafkaRechargesWithContext(context.Background(), request)
}

// DescribeKafkaRecharges
// This API is used to get the list of Kafka data subscription tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaRechargesWithContext(ctx context.Context, request *DescribeKafkaRechargesRequest) (response *DescribeKafkaRechargesResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaRechargesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeKafkaRecharges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaRecharges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaRechargesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogContextRequest() (request *DescribeLogContextRequest) {
    request = &DescribeLogContextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeLogContext")
    
    
    return
}

func NewDescribeLogContextResponse() (response *DescribeLogContextResponse) {
    response = &DescribeLogContextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogContext
// This API is used to search for content near the log context. For more details, see [Context Search](https://intl.cloud.tencent.com/document/product/614/53248?from_cn_redirect=1).The maximum value of API's return data packet is 49MB. It is recommended to enable gzip compression (HTTP Request Header Accept-Encoding: gzip).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogContext(request *DescribeLogContextRequest) (response *DescribeLogContextResponse, err error) {
    return c.DescribeLogContextWithContext(context.Background(), request)
}

// DescribeLogContext
// This API is used to search for content near the log context. For more details, see [Context Search](https://intl.cloud.tencent.com/document/product/614/53248?from_cn_redirect=1).The maximum value of API's return data packet is 49MB. It is recommended to enable gzip compression (HTTP Request Header Accept-Encoding: gzip).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogContextWithContext(ctx context.Context, request *DescribeLogContextRequest) (response *DescribeLogContextResponse, err error) {
    if request == nil {
        request = NewDescribeLogContextRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeLogContext")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogContext require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogContextResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogHistogramRequest() (request *DescribeLogHistogramRequest) {
    request = &DescribeLogHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeLogHistogram")
    
    
    return
}

func NewDescribeLogHistogramResponse() (response *DescribeLogHistogramResponse) {
    response = &DescribeLogHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogHistogram
// This API is used to get a log count histogram.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogHistogram(request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
    return c.DescribeLogHistogramWithContext(context.Background(), request)
}

// DescribeLogHistogram
// This API is used to get a log count histogram.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogHistogramWithContext(ctx context.Context, request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
    if request == nil {
        request = NewDescribeLogHistogramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeLogHistogram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogsetsRequest() (request *DescribeLogsetsRequest) {
    request = &DescribeLogsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeLogsets")
    
    
    return
}

func NewDescribeLogsetsResponse() (response *DescribeLogsetsResponse) {
    response = &DescribeLogsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogsets
// This API is used to get the list of logsets.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogsets(request *DescribeLogsetsRequest) (response *DescribeLogsetsResponse, err error) {
    return c.DescribeLogsetsWithContext(context.Background(), request)
}

// DescribeLogsets
// This API is used to get the list of logsets.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogsetsWithContext(ctx context.Context, request *DescribeLogsetsRequest) (response *DescribeLogsetsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeLogsets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineGroupConfigsRequest() (request *DescribeMachineGroupConfigsRequest) {
    request = &DescribeMachineGroupConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroupConfigs")
    
    
    return
}

func NewDescribeMachineGroupConfigsResponse() (response *DescribeMachineGroupConfigsResponse) {
    response = &DescribeMachineGroupConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMachineGroupConfigs
// This API is used to get the collection rule configuration bound to a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachineGroupConfigs(request *DescribeMachineGroupConfigsRequest) (response *DescribeMachineGroupConfigsResponse, err error) {
    return c.DescribeMachineGroupConfigsWithContext(context.Background(), request)
}

// DescribeMachineGroupConfigs
// This API is used to get the collection rule configuration bound to a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachineGroupConfigsWithContext(ctx context.Context, request *DescribeMachineGroupConfigsRequest) (response *DescribeMachineGroupConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeMachineGroupConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeMachineGroupConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineGroupConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineGroupConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineGroupsRequest() (request *DescribeMachineGroupsRequest) {
    request = &DescribeMachineGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroups")
    
    
    return
}

func NewDescribeMachineGroupsResponse() (response *DescribeMachineGroupsResponse) {
    response = &DescribeMachineGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMachineGroups
// This API is used to get the list of machine groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMachineGroups(request *DescribeMachineGroupsRequest) (response *DescribeMachineGroupsResponse, err error) {
    return c.DescribeMachineGroupsWithContext(context.Background(), request)
}

// DescribeMachineGroups
// This API is used to get the list of machine groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMachineGroupsWithContext(ctx context.Context, request *DescribeMachineGroupsRequest) (response *DescribeMachineGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeMachineGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeMachineGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
    request = &DescribeMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMachines")
    
    
    return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
    response = &DescribeMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMachines
// This API is used to get the status of a machine under the specified machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_AGENTVERSIONNOTEXIST = "ResourceNotFound.AgentVersionNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    return c.DescribeMachinesWithContext(context.Background(), request)
}

// DescribeMachines
// This API is used to get the status of a machine under the specified machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_AGENTVERSIONNOTEXIST = "ResourceNotFound.AgentVersionNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachinesWithContext(ctx context.Context, request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeMachines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricCorrectDimensionRequest() (request *DescribeMetricCorrectDimensionRequest) {
    request = &DescribeMetricCorrectDimensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMetricCorrectDimension")
    
    
    return
}

func NewDescribeMetricCorrectDimensionResponse() (response *DescribeMetricCorrectDimensionResponse) {
    response = &DescribeMetricCorrectDimensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricCorrectDimension
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricCorrectDimension(request *DescribeMetricCorrectDimensionRequest) (response *DescribeMetricCorrectDimensionResponse, err error) {
    return c.DescribeMetricCorrectDimensionWithContext(context.Background(), request)
}

// DescribeMetricCorrectDimension
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricCorrectDimensionWithContext(ctx context.Context, request *DescribeMetricCorrectDimensionRequest) (response *DescribeMetricCorrectDimensionResponse, err error) {
    if request == nil {
        request = NewDescribeMetricCorrectDimensionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeMetricCorrectDimension")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricCorrectDimension require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricCorrectDimensionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricSubscribePreviewRequest() (request *DescribeMetricSubscribePreviewRequest) {
    request = &DescribeMetricSubscribePreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMetricSubscribePreview")
    
    
    return
}

func NewDescribeMetricSubscribePreviewResponse() (response *DescribeMetricSubscribePreviewResponse) {
    response = &DescribeMetricSubscribePreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricSubscribePreview
// This API is used to create metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricSubscribePreview(request *DescribeMetricSubscribePreviewRequest) (response *DescribeMetricSubscribePreviewResponse, err error) {
    return c.DescribeMetricSubscribePreviewWithContext(context.Background(), request)
}

// DescribeMetricSubscribePreview
// This API is used to create metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricSubscribePreviewWithContext(ctx context.Context, request *DescribeMetricSubscribePreviewRequest) (response *DescribeMetricSubscribePreviewResponse, err error) {
    if request == nil {
        request = NewDescribeMetricSubscribePreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeMetricSubscribePreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricSubscribePreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricSubscribePreviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricSubscribesRequest() (request *DescribeMetricSubscribesRequest) {
    request = &DescribeMetricSubscribesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMetricSubscribes")
    
    
    return
}

func NewDescribeMetricSubscribesResponse() (response *DescribeMetricSubscribesResponse) {
    response = &DescribeMetricSubscribesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricSubscribes
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricSubscribes(request *DescribeMetricSubscribesRequest) (response *DescribeMetricSubscribesResponse, err error) {
    return c.DescribeMetricSubscribesWithContext(context.Background(), request)
}

// DescribeMetricSubscribes
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricSubscribesWithContext(ctx context.Context, request *DescribeMetricSubscribesRequest) (response *DescribeMetricSubscribesResponse, err error) {
    if request == nil {
        request = NewDescribeMetricSubscribesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeMetricSubscribes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricSubscribes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricSubscribesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkApplicationDetailRequest() (request *DescribeNetworkApplicationDetailRequest) {
    request = &DescribeNetworkApplicationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeNetworkApplicationDetail")
    
    
    return
}

func NewDescribeNetworkApplicationDetailResponse() (response *DescribeNetworkApplicationDetailResponse) {
    response = &DescribeNetworkApplicationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkApplicationDetail
// Retrieve web application details
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeNetworkApplicationDetail(request *DescribeNetworkApplicationDetailRequest) (response *DescribeNetworkApplicationDetailResponse, err error) {
    return c.DescribeNetworkApplicationDetailWithContext(context.Background(), request)
}

// DescribeNetworkApplicationDetail
// Retrieve web application details
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeNetworkApplicationDetailWithContext(ctx context.Context, request *DescribeNetworkApplicationDetailRequest) (response *DescribeNetworkApplicationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkApplicationDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeNetworkApplicationDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkApplicationDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkApplicationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkApplicationsRequest() (request *DescribeNetworkApplicationsRequest) {
    request = &DescribeNetworkApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeNetworkApplications")
    
    
    return
}

func NewDescribeNetworkApplicationsResponse() (response *DescribeNetworkApplicationsResponse) {
    response = &DescribeNetworkApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkApplications
// Retrieve the network application list
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeNetworkApplications(request *DescribeNetworkApplicationsRequest) (response *DescribeNetworkApplicationsResponse, err error) {
    return c.DescribeNetworkApplicationsWithContext(context.Background(), request)
}

// DescribeNetworkApplications
// Retrieve the network application list
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeNetworkApplicationsWithContext(ctx context.Context, request *DescribeNetworkApplicationsRequest) (response *DescribeNetworkApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkApplicationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeNetworkApplications")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNoticeContentsRequest() (request *DescribeNoticeContentsRequest) {
    request = &DescribeNoticeContentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeNoticeContents")
    
    
    return
}

func NewDescribeNoticeContentsResponse() (response *DescribeNoticeContentsResponse) {
    response = &DescribeNoticeContentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNoticeContents
// This API is used to obtain the notification content list.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeNoticeContents(request *DescribeNoticeContentsRequest) (response *DescribeNoticeContentsResponse, err error) {
    return c.DescribeNoticeContentsWithContext(context.Background(), request)
}

// DescribeNoticeContents
// This API is used to obtain the notification content list.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeNoticeContentsWithContext(ctx context.Context, request *DescribeNoticeContentsRequest) (response *DescribeNoticeContentsResponse, err error) {
    if request == nil {
        request = NewDescribeNoticeContentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeNoticeContents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNoticeContents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNoticeContentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePartitionsRequest() (request *DescribePartitionsRequest) {
    request = &DescribePartitionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribePartitions")
    
    
    return
}

func NewDescribePartitionsResponse() (response *DescribePartitionsResponse) {
    response = &DescribePartitionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePartitions
// This API is deprecated. If needed, please use the DescribeTopics API to get the number of partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribePartitions(request *DescribePartitionsRequest) (response *DescribePartitionsResponse, err error) {
    return c.DescribePartitionsWithContext(context.Background(), request)
}

// DescribePartitions
// This API is deprecated. If needed, please use the DescribeTopics API to get the number of partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribePartitionsWithContext(ctx context.Context, request *DescribePartitionsRequest) (response *DescribePartitionsResponse, err error) {
    if request == nil {
        request = NewDescribePartitionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribePartitions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePartitions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePartitionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRebuildIndexTasksRequest() (request *DescribeRebuildIndexTasksRequest) {
    request = &DescribeRebuildIndexTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeRebuildIndexTasks")
    
    
    return
}

func NewDescribeRebuildIndexTasksResponse() (response *DescribeRebuildIndexTasksResponse) {
    response = &DescribeRebuildIndexTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRebuildIndexTasks
// This API is used to obtain the list of rebuild index tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeRebuildIndexTasks(request *DescribeRebuildIndexTasksRequest) (response *DescribeRebuildIndexTasksResponse, err error) {
    return c.DescribeRebuildIndexTasksWithContext(context.Background(), request)
}

// DescribeRebuildIndexTasks
// This API is used to obtain the list of rebuild index tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeRebuildIndexTasksWithContext(ctx context.Context, request *DescribeRebuildIndexTasksRequest) (response *DescribeRebuildIndexTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRebuildIndexTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeRebuildIndexTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRebuildIndexTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRebuildIndexTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingRuleTaskRequest() (request *DescribeRecordingRuleTaskRequest) {
    request = &DescribeRecordingRuleTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeRecordingRuleTask")
    
    
    return
}

func NewDescribeRecordingRuleTaskResponse() (response *DescribeRecordingRuleTaskResponse) {
    response = &DescribeRecordingRuleTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordingRuleTask
// This API is used to retrieve the pre-aggregation task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeRecordingRuleTask(request *DescribeRecordingRuleTaskRequest) (response *DescribeRecordingRuleTaskResponse, err error) {
    return c.DescribeRecordingRuleTaskWithContext(context.Background(), request)
}

// DescribeRecordingRuleTask
// This API is used to retrieve the pre-aggregation task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeRecordingRuleTaskWithContext(ctx context.Context, request *DescribeRecordingRuleTaskRequest) (response *DescribeRecordingRuleTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingRuleTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeRecordingRuleTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordingRuleTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordingRuleTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingRuleYamlTaskRequest() (request *DescribeRecordingRuleYamlTaskRequest) {
    request = &DescribeRecordingRuleYamlTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeRecordingRuleYamlTask")
    
    
    return
}

func NewDescribeRecordingRuleYamlTaskResponse() (response *DescribeRecordingRuleYamlTaskResponse) {
    response = &DescribeRecordingRuleYamlTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordingRuleYamlTask
// This API is used to retrieve the pre-aggregation task list in yaml.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeRecordingRuleYamlTask(request *DescribeRecordingRuleYamlTaskRequest) (response *DescribeRecordingRuleYamlTaskResponse, err error) {
    return c.DescribeRecordingRuleYamlTaskWithContext(context.Background(), request)
}

// DescribeRecordingRuleYamlTask
// This API is used to retrieve the pre-aggregation task list in yaml.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeRecordingRuleYamlTaskWithContext(ctx context.Context, request *DescribeRecordingRuleYamlTaskRequest) (response *DescribeRecordingRuleYamlTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingRuleYamlTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeRecordingRuleYamlTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordingRuleYamlTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordingRuleYamlTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScheduledSqlInfoRequest() (request *DescribeScheduledSqlInfoRequest) {
    request = &DescribeScheduledSqlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeScheduledSqlInfo")
    
    
    return
}

func NewDescribeScheduledSqlInfoResponse() (response *DescribeScheduledSqlInfoResponse) {
    response = &DescribeScheduledSqlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScheduledSqlInfo
// This API is used to access the scheduled SQL analysis task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeScheduledSqlInfo(request *DescribeScheduledSqlInfoRequest) (response *DescribeScheduledSqlInfoResponse, err error) {
    return c.DescribeScheduledSqlInfoWithContext(context.Background(), request)
}

// DescribeScheduledSqlInfo
// This API is used to access the scheduled SQL analysis task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeScheduledSqlInfoWithContext(ctx context.Context, request *DescribeScheduledSqlInfoRequest) (response *DescribeScheduledSqlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeScheduledSqlInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeScheduledSqlInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScheduledSqlInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScheduledSqlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchViewsRequest() (request *DescribeSearchViewsRequest) {
    request = &DescribeSearchViewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeSearchViews")
    
    
    return
}

func NewDescribeSearchViewsResponse() (response *DescribeSearchViewsResponse) {
    response = &DescribeSearchViewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchViews
// Query view list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) DescribeSearchViews(request *DescribeSearchViewsRequest) (response *DescribeSearchViewsResponse, err error) {
    return c.DescribeSearchViewsWithContext(context.Background(), request)
}

// DescribeSearchViews
// Query view list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) DescribeSearchViewsWithContext(ctx context.Context, request *DescribeSearchViewsRequest) (response *DescribeSearchViewsResponse, err error) {
    if request == nil {
        request = NewDescribeSearchViewsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeSearchViews")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchViews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchViewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShipperTasksRequest() (request *DescribeShipperTasksRequest) {
    request = &DescribeShipperTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeShipperTasks")
    
    
    return
}

func NewDescribeShipperTasksResponse() (response *DescribeShipperTasksResponse) {
    response = &DescribeShipperTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShipperTasks
// This API is used to get the list of shipping tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DescribeShipperTasks(request *DescribeShipperTasksRequest) (response *DescribeShipperTasksResponse, err error) {
    return c.DescribeShipperTasksWithContext(context.Background(), request)
}

// DescribeShipperTasks
// This API is used to get the list of shipping tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DescribeShipperTasksWithContext(ctx context.Context, request *DescribeShipperTasksRequest) (response *DescribeShipperTasksResponse, err error) {
    if request == nil {
        request = NewDescribeShipperTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeShipperTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShipperTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShipperTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShippersRequest() (request *DescribeShippersRequest) {
    request = &DescribeShippersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeShippers")
    
    
    return
}

func NewDescribeShippersResponse() (response *DescribeShippersResponse) {
    response = &DescribeShippersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShippers
// This API is used to get the configuration of the task shipped to COS.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeShippers(request *DescribeShippersRequest) (response *DescribeShippersResponse, err error) {
    return c.DescribeShippersWithContext(context.Background(), request)
}

// DescribeShippers
// This API is used to get the configuration of the task shipped to COS.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeShippersWithContext(ctx context.Context, request *DescribeShippersRequest) (response *DescribeShippersResponse, err error) {
    if request == nil {
        request = NewDescribeShippersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeShippers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShippers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShippersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSplunkDeliversRequest() (request *DescribeSplunkDeliversRequest) {
    request = &DescribeSplunkDeliversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeSplunkDelivers")
    
    
    return
}

func NewDescribeSplunkDeliversResponse() (response *DescribeSplunkDeliversResponse) {
    response = &DescribeSplunkDeliversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSplunkDelivers
// Retrieve the Splunk delivery task list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) DescribeSplunkDelivers(request *DescribeSplunkDeliversRequest) (response *DescribeSplunkDeliversResponse, err error) {
    return c.DescribeSplunkDeliversWithContext(context.Background(), request)
}

// DescribeSplunkDelivers
// Retrieve the Splunk delivery task list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) DescribeSplunkDeliversWithContext(ctx context.Context, request *DescribeSplunkDeliversRequest) (response *DescribeSplunkDeliversResponse, err error) {
    if request == nil {
        request = NewDescribeSplunkDeliversRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeSplunkDelivers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSplunkDelivers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSplunkDeliversResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSplunkPreviewRequest() (request *DescribeSplunkPreviewRequest) {
    request = &DescribeSplunkPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeSplunkPreview")
    
    
    return
}

func NewDescribeSplunkPreviewResponse() (response *DescribeSplunkPreviewResponse) {
    response = &DescribeSplunkPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSplunkPreview
// splunk delivery task preview
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSplunkPreview(request *DescribeSplunkPreviewRequest) (response *DescribeSplunkPreviewResponse, err error) {
    return c.DescribeSplunkPreviewWithContext(context.Background(), request)
}

// DescribeSplunkPreview
// splunk delivery task preview
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSplunkPreviewWithContext(ctx context.Context, request *DescribeSplunkPreviewRequest) (response *DescribeSplunkPreviewResponse, err error) {
    if request == nil {
        request = NewDescribeSplunkPreviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeSplunkPreview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSplunkPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSplunkPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicBaseMetricConfigsRequest() (request *DescribeTopicBaseMetricConfigsRequest) {
    request = &DescribeTopicBaseMetricConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeTopicBaseMetricConfigs")
    
    
    return
}

func NewDescribeTopicBaseMetricConfigsResponse() (response *DescribeTopicBaseMetricConfigsResponse) {
    response = &DescribeTopicBaseMetricConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicBaseMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicBaseMetricConfigs(request *DescribeTopicBaseMetricConfigsRequest) (response *DescribeTopicBaseMetricConfigsResponse, err error) {
    return c.DescribeTopicBaseMetricConfigsWithContext(context.Background(), request)
}

// DescribeTopicBaseMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicBaseMetricConfigsWithContext(ctx context.Context, request *DescribeTopicBaseMetricConfigsRequest) (response *DescribeTopicBaseMetricConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicBaseMetricConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeTopicBaseMetricConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicBaseMetricConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicBaseMetricConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicMetricConfigsRequest() (request *DescribeTopicMetricConfigsRequest) {
    request = &DescribeTopicMetricConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeTopicMetricConfigs")
    
    
    return
}

func NewDescribeTopicMetricConfigsResponse() (response *DescribeTopicMetricConfigsResponse) {
    response = &DescribeTopicMetricConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicMetricConfigs(request *DescribeTopicMetricConfigsRequest) (response *DescribeTopicMetricConfigsResponse, err error) {
    return c.DescribeTopicMetricConfigsWithContext(context.Background(), request)
}

// DescribeTopicMetricConfigs
// This API is used to obtain metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicMetricConfigsWithContext(ctx context.Context, request *DescribeTopicMetricConfigsRequest) (response *DescribeTopicMetricConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicMetricConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeTopicMetricConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicMetricConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicMetricConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
    request = &DescribeTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeTopics")
    
    
    return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
    response = &DescribeTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopics
// This API is used to obtain logs or metric topic lists and supports pagination.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ANALYSISSWITCHCLOSE = "OperationDenied.AnalysisSwitchClose"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    return c.DescribeTopicsWithContext(context.Background(), request)
}

// DescribeTopics
// This API is used to obtain logs or metric topic lists and supports pagination.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ANALYSISSWITCHCLOSE = "OperationDenied.AnalysisSwitchClose"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicsWithContext(ctx context.Context, request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebCallbacksRequest() (request *DescribeWebCallbacksRequest) {
    request = &DescribeWebCallbacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeWebCallbacks")
    
    
    return
}

func NewDescribeWebCallbacksResponse() (response *DescribeWebCallbacksResponse) {
    response = &DescribeWebCallbacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebCallbacks
// This API is used to search alarm channel callback configuration lists.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeWebCallbacks(request *DescribeWebCallbacksRequest) (response *DescribeWebCallbacksResponse, err error) {
    return c.DescribeWebCallbacksWithContext(context.Background(), request)
}

// DescribeWebCallbacks
// This API is used to search alarm channel callback configuration lists.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeWebCallbacksWithContext(ctx context.Context, request *DescribeWebCallbacksRequest) (response *DescribeWebCallbacksResponse, err error) {
    if request == nil {
        request = NewDescribeWebCallbacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "DescribeWebCallbacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebCallbacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebCallbacksResponse()
    err = c.Send(request, response)
    return
}

func NewEstimateRebuildIndexTaskRequest() (request *EstimateRebuildIndexTaskRequest) {
    request = &EstimateRebuildIndexTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "EstimateRebuildIndexTask")
    
    
    return
}

func NewEstimateRebuildIndexTaskResponse() (response *EstimateRebuildIndexTaskResponse) {
    response = &EstimateRebuildIndexTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EstimateRebuildIndexTask
// This API is used to estimate rebuild index tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) EstimateRebuildIndexTask(request *EstimateRebuildIndexTaskRequest) (response *EstimateRebuildIndexTaskResponse, err error) {
    return c.EstimateRebuildIndexTaskWithContext(context.Background(), request)
}

// EstimateRebuildIndexTask
// This API is used to estimate rebuild index tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) EstimateRebuildIndexTaskWithContext(ctx context.Context, request *EstimateRebuildIndexTaskRequest) (response *EstimateRebuildIndexTaskResponse, err error) {
    if request == nil {
        request = NewEstimateRebuildIndexTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "EstimateRebuildIndexTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EstimateRebuildIndexTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewEstimateRebuildIndexTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetAlarmLogRequest() (request *GetAlarmLogRequest) {
    request = &GetAlarmLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "GetAlarmLog")
    
    
    return
}

func NewGetAlarmLogResponse() (response *GetAlarmLogResponse) {
    response = &GetAlarmLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAlarmLog
// This API is used to access alarm policy execution details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAlarmLog(request *GetAlarmLogRequest) (response *GetAlarmLogResponse, err error) {
    return c.GetAlarmLogWithContext(context.Background(), request)
}

// GetAlarmLog
// This API is used to access alarm policy execution details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAlarmLogWithContext(ctx context.Context, request *GetAlarmLogRequest) (response *GetAlarmLogResponse, err error) {
    if request == nil {
        request = NewGetAlarmLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "GetAlarmLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAlarmLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAlarmLogResponse()
    err = c.Send(request, response)
    return
}

func NewGetClsServiceRequest() (request *GetClsServiceRequest) {
    request = &GetClsServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "GetClsService")
    
    
    return
}

func NewGetClsServiceResponse() (response *GetClsServiceResponse) {
    response = &GetClsServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetClsService
// This API is used to check whether CLS is enabled.
//
// This API is used to fill in any region for Region, recommend using Guangzhou (ap-guangzhou).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetClsService(request *GetClsServiceRequest) (response *GetClsServiceResponse, err error) {
    return c.GetClsServiceWithContext(context.Background(), request)
}

// GetClsService
// This API is used to check whether CLS is enabled.
//
// This API is used to fill in any region for Region, recommend using Guangzhou (ap-guangzhou).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetClsServiceWithContext(ctx context.Context, request *GetClsServiceRequest) (response *GetClsServiceResponse, err error) {
    if request == nil {
        request = NewGetClsServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "GetClsService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetClsService require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetClsServiceResponse()
    err = c.Send(request, response)
    return
}

func NewGetMetricLabelValuesRequest() (request *GetMetricLabelValuesRequest) {
    request = &GetMetricLabelValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "GetMetricLabelValues")
    
    
    return
}

func NewGetMetricLabelValuesResponse() (response *GetMetricLabelValuesResponse) {
    response = &GetMetricLabelValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMetricLabelValues
// This API is used to obtain the list of time series label values.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) GetMetricLabelValues(request *GetMetricLabelValuesRequest) (response *GetMetricLabelValuesResponse, err error) {
    return c.GetMetricLabelValuesWithContext(context.Background(), request)
}

// GetMetricLabelValues
// This API is used to obtain the list of time series label values.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) GetMetricLabelValuesWithContext(ctx context.Context, request *GetMetricLabelValuesRequest) (response *GetMetricLabelValuesResponse, err error) {
    if request == nil {
        request = NewGetMetricLabelValuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "GetMetricLabelValues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMetricLabelValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMetricLabelValuesResponse()
    err = c.Send(request, response)
    return
}

func NewMergePartitionRequest() (request *MergePartitionRequest) {
    request = &MergePartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "MergePartition")
    
    
    return
}

func NewMergePartitionResponse() (response *MergePartitionResponse) {
    response = &MergePartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MergePartition
// This API is deprecated. If needed, please use the ModifyTopic API to change the number of partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MergePartition(request *MergePartitionRequest) (response *MergePartitionResponse, err error) {
    return c.MergePartitionWithContext(context.Background(), request)
}

// MergePartition
// This API is deprecated. If needed, please use the ModifyTopic API to change the number of partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MergePartitionWithContext(ctx context.Context, request *MergePartitionRequest) (response *MergePartitionResponse, err error) {
    if request == nil {
        request = NewMergePartitionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "MergePartition")
    
    if c.GetCredential() == nil {
        return nil, errors.New("MergePartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewMergePartitionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmRequest() (request *ModifyAlarmRequest) {
    request = &ModifyAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarm")
    
    
    return
}

func NewModifyAlarmResponse() (response *ModifyAlarmResponse) {
    response = &ModifyAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarm
// This API is used to modify an alarm policy. At least one valid configuration item needs to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDALARM = "FailedOperation.InvalidAlarm"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyAlarm(request *ModifyAlarmRequest) (response *ModifyAlarmResponse, err error) {
    return c.ModifyAlarmWithContext(context.Background(), request)
}

// ModifyAlarm
// This API is used to modify an alarm policy. At least one valid configuration item needs to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDALARM = "FailedOperation.InvalidAlarm"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyAlarmWithContext(ctx context.Context, request *ModifyAlarmRequest) (response *ModifyAlarmResponse, err error) {
    if request == nil {
        request = NewModifyAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmNoticeRequest() (request *ModifyAlarmNoticeRequest) {
    request = &ModifyAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarmNotice")
    
    
    return
}

func NewModifyAlarmNoticeResponse() (response *ModifyAlarmNoticeResponse) {
    response = &ModifyAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmNotice
// This API is used to modify a notification group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyAlarmNotice(request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    return c.ModifyAlarmNoticeWithContext(context.Background(), request)
}

// ModifyAlarmNotice
// This API is used to modify a notification group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyAlarmNoticeWithContext(ctx context.Context, request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmShieldRequest() (request *ModifyAlarmShieldRequest) {
    request = &ModifyAlarmShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarmShield")
    
    
    return
}

func NewModifyAlarmShieldResponse() (response *ModifyAlarmShieldResponse) {
    response = &ModifyAlarmShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmShield
// This API is used to modify alarm blocking rules. When the alarm blocking rule is invalid, it cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyAlarmShield(request *ModifyAlarmShieldRequest) (response *ModifyAlarmShieldResponse, err error) {
    return c.ModifyAlarmShieldWithContext(context.Background(), request)
}

// ModifyAlarmShield
// This API is used to modify alarm blocking rules. When the alarm blocking rule is invalid, it cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyAlarmShieldWithContext(ctx context.Context, request *ModifyAlarmShieldRequest) (response *ModifyAlarmShieldResponse, err error) {
    if request == nil {
        request = NewModifyAlarmShieldRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyAlarmShield")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmShield require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmShieldResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudProductLogCollectionRequest() (request *ModifyCloudProductLogCollectionRequest) {
    request = &ModifyCloudProductLogCollectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyCloudProductLogCollection")
    
    
    return
}

func NewModifyCloudProductLogCollectionResponse() (response *ModifyCloudProductLogCollectionResponse) {
    response = &ModifyCloudProductLogCollectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudProductLogCollection
// Cloud product integration uses internal APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDPRODUCTINVOCATIONERROR = "FailedOperation.CloudProductInvocationError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudProductLogCollection(request *ModifyCloudProductLogCollectionRequest) (response *ModifyCloudProductLogCollectionResponse, err error) {
    return c.ModifyCloudProductLogCollectionWithContext(context.Background(), request)
}

// ModifyCloudProductLogCollection
// Cloud product integration uses internal APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDPRODUCTINVOCATIONERROR = "FailedOperation.CloudProductInvocationError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudProductLogCollectionWithContext(ctx context.Context, request *ModifyCloudProductLogCollectionRequest) (response *ModifyCloudProductLogCollectionResponse, err error) {
    if request == nil {
        request = NewModifyCloudProductLogCollectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyCloudProductLogCollection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudProductLogCollection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudProductLogCollectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConfigRequest() (request *ModifyConfigRequest) {
    request = &ModifyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConfig")
    
    
    return
}

func NewModifyConfigResponse() (response *ModifyConfigResponse) {
    response = &ModifyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConfig
// This API is used to modify collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyConfig(request *ModifyConfigRequest) (response *ModifyConfigResponse, err error) {
    return c.ModifyConfigWithContext(context.Background(), request)
}

// ModifyConfig
// This API is used to modify collection rule configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyConfigWithContext(ctx context.Context, request *ModifyConfigRequest) (response *ModifyConfigResponse, err error) {
    if request == nil {
        request = NewModifyConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsoleRequest() (request *ModifyConsoleRequest) {
    request = &ModifyConsoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConsole")
    
    
    return
}

func NewModifyConsoleResponse() (response *ModifyConsoleResponse) {
    response = &ModifyConsoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsole
// This API is used to edit the DataSight Console
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConsole(request *ModifyConsoleRequest) (response *ModifyConsoleResponse, err error) {
    return c.ModifyConsoleWithContext(context.Background(), request)
}

// ModifyConsole
// This API is used to edit the DataSight Console
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConsoleWithContext(ctx context.Context, request *ModifyConsoleRequest) (response *ModifyConsoleResponse, err error) {
    if request == nil {
        request = NewModifyConsoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyConsole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerRequest() (request *ModifyConsumerRequest) {
    request = &ModifyConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConsumer")
    
    
    return
}

func NewModifyConsumerResponse() (response *ModifyConsumerResponse) {
    response = &ModifyConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumer
// This API is used to modify a CKafka delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) ModifyConsumer(request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
    return c.ModifyConsumerWithContext(context.Background(), request)
}

// ModifyConsumer
// This API is used to modify a CKafka delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) ModifyConsumerWithContext(ctx context.Context, request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
    if request == nil {
        request = NewModifyConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerGroupRequest() (request *ModifyConsumerGroupRequest) {
    request = &ModifyConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConsumerGroup")
    
    
    return
}

func NewModifyConsumerGroupResponse() (response *ModifyConsumerGroupResponse) {
    response = &ModifyConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumerGroup
// This API is used to update the consumer group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyConsumerGroup(request *ModifyConsumerGroupRequest) (response *ModifyConsumerGroupResponse, err error) {
    return c.ModifyConsumerGroupWithContext(context.Background(), request)
}

// ModifyConsumerGroup
// This API is used to update the consumer group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyConsumerGroupWithContext(ctx context.Context, request *ModifyConsumerGroupRequest) (response *ModifyConsumerGroupResponse, err error) {
    if request == nil {
        request = NewModifyConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCosRechargeRequest() (request *ModifyCosRechargeRequest) {
    request = &ModifyCosRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyCosRecharge")
    
    
    return
}

func NewModifyCosRechargeResponse() (response *ModifyCosRechargeResponse) {
    response = &ModifyCosRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCosRecharge
// This API is used to modify a COS import task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION_MODIFYBILLINGCOSRECHARGENOSUPPORT = "UnsupportedOperation.ModifyBillingCosRechargeNoSupport"
func (c *Client) ModifyCosRecharge(request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
    return c.ModifyCosRechargeWithContext(context.Background(), request)
}

// ModifyCosRecharge
// This API is used to modify a COS import task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION_MODIFYBILLINGCOSRECHARGENOSUPPORT = "UnsupportedOperation.ModifyBillingCosRechargeNoSupport"
func (c *Client) ModifyCosRechargeWithContext(ctx context.Context, request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
    if request == nil {
        request = NewModifyCosRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyCosRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCosRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCosRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDashboardRequest() (request *ModifyDashboardRequest) {
    request = &ModifyDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyDashboard")
    
    
    return
}

func NewModifyDashboardResponse() (response *ModifyDashboardResponse) {
    response = &ModifyDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDashboard
// This API is used to modify the dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyDashboard(request *ModifyDashboardRequest) (response *ModifyDashboardResponse, err error) {
    return c.ModifyDashboardWithContext(context.Background(), request)
}

// ModifyDashboard
// This API is used to modify the dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyDashboardWithContext(ctx context.Context, request *ModifyDashboardRequest) (response *ModifyDashboardResponse, err error) {
    if request == nil {
        request = NewModifyDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDashboardSubscribeRequest() (request *ModifyDashboardSubscribeRequest) {
    request = &ModifyDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyDashboardSubscribe")
    
    
    return
}

func NewModifyDashboardSubscribeResponse() (response *ModifyDashboardSubscribeResponse) {
    response = &ModifyDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDashboardSubscribe
// This API is used to modify dashboard subscriptions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) ModifyDashboardSubscribe(request *ModifyDashboardSubscribeRequest) (response *ModifyDashboardSubscribeResponse, err error) {
    return c.ModifyDashboardSubscribeWithContext(context.Background(), request)
}

// ModifyDashboardSubscribe
// This API is used to modify dashboard subscriptions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) ModifyDashboardSubscribeWithContext(ctx context.Context, request *ModifyDashboardSubscribeRequest) (response *ModifyDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewModifyDashboardSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyDashboardSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataTransformRequest() (request *ModifyDataTransformRequest) {
    request = &ModifyDataTransformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyDataTransform")
    
    
    return
}

func NewModifyDataTransformResponse() (response *ModifyDataTransformResponse) {
    response = &ModifyDataTransformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDataTransform
// This API is used to modify a data processing task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CROSSACCOUNTCONFLICT = "InvalidParameter.CrossAccountConflict"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyDataTransform(request *ModifyDataTransformRequest) (response *ModifyDataTransformResponse, err error) {
    return c.ModifyDataTransformWithContext(context.Background(), request)
}

// ModifyDataTransform
// This API is used to modify a data processing task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CROSSACCOUNTCONFLICT = "InvalidParameter.CrossAccountConflict"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyDataTransformWithContext(ctx context.Context, request *ModifyDataTransformRequest) (response *ModifyDataTransformResponse, err error) {
    if request == nil {
        request = NewModifyDataTransformRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyDataTransform")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataTransformResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDlcDeliverRequest() (request *ModifyDlcDeliverRequest) {
    request = &ModifyDlcDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyDlcDeliver")
    
    
    return
}

func NewModifyDlcDeliverResponse() (response *ModifyDlcDeliverResponse) {
    response = &ModifyDlcDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDlcDeliver
// Modify a DLC delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDlcDeliver(request *ModifyDlcDeliverRequest) (response *ModifyDlcDeliverResponse, err error) {
    return c.ModifyDlcDeliverWithContext(context.Background(), request)
}

// ModifyDlcDeliver
// Modify a DLC delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDlcDeliverWithContext(ctx context.Context, request *ModifyDlcDeliverRequest) (response *ModifyDlcDeliverResponse, err error) {
    if request == nil {
        request = NewModifyDlcDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyDlcDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDlcDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDlcDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEsRechargeRequest() (request *ModifyEsRechargeRequest) {
    request = &ModifyEsRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyEsRecharge")
    
    
    return
}

func NewModifyEsRechargeResponse() (response *ModifyEsRechargeResponse) {
    response = &ModifyEsRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEsRecharge
// Modify es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEsRecharge(request *ModifyEsRechargeRequest) (response *ModifyEsRechargeResponse, err error) {
    return c.ModifyEsRechargeWithContext(context.Background(), request)
}

// ModifyEsRecharge
// Modify es import configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEsRechargeWithContext(ctx context.Context, request *ModifyEsRechargeRequest) (response *ModifyEsRechargeResponse, err error) {
    if request == nil {
        request = NewModifyEsRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyEsRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEsRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEsRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostMetricConfigRequest() (request *ModifyHostMetricConfigRequest) {
    request = &ModifyHostMetricConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyHostMetricConfig")
    
    
    return
}

func NewModifyHostMetricConfigResponse() (response *ModifyHostMetricConfigResponse) {
    response = &ModifyHostMetricConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostMetricConfig
// Modify host metric collection configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHostMetricConfig(request *ModifyHostMetricConfigRequest) (response *ModifyHostMetricConfigResponse, err error) {
    return c.ModifyHostMetricConfigWithContext(context.Background(), request)
}

// ModifyHostMetricConfig
// Modify host metric collection configuration
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHostMetricConfigWithContext(ctx context.Context, request *ModifyHostMetricConfigRequest) (response *ModifyHostMetricConfigResponse, err error) {
    if request == nil {
        request = NewModifyHostMetricConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyHostMetricConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostMetricConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostMetricConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIndexRequest() (request *ModifyIndexRequest) {
    request = &ModifyIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyIndex")
    
    
    return
}

func NewModifyIndexResponse() (response *ModifyIndexResponse) {
    response = &ModifyIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIndex
// This API is used to modify the index configuration. It is subject to the default request frequency limit, and the number of concurrent requests to the same log topic cannot exceed 1, i.e., the index configuration of only one log topic can be modified at a time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  LIMITEXCEEDED_INDEXKEYOVERLIMIT = "LimitExceeded.IndexKeyOverLimit"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIndex(request *ModifyIndexRequest) (response *ModifyIndexResponse, err error) {
    return c.ModifyIndexWithContext(context.Background(), request)
}

// ModifyIndex
// This API is used to modify the index configuration. It is subject to the default request frequency limit, and the number of concurrent requests to the same log topic cannot exceed 1, i.e., the index configuration of only one log topic can be modified at a time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  LIMITEXCEEDED_INDEXKEYOVERLIMIT = "LimitExceeded.IndexKeyOverLimit"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIndexWithContext(ctx context.Context, request *ModifyIndexRequest) (response *ModifyIndexResponse, err error) {
    if request == nil {
        request = NewModifyIndexRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyIndex")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIndexResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKafkaConsumerRequest() (request *ModifyKafkaConsumerRequest) {
    request = &ModifyKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaConsumer")
    
    
    return
}

func NewModifyKafkaConsumerResponse() (response *ModifyKafkaConsumerResponse) {
    response = &ModifyKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKafkaConsumer
// This API is used to modify Kafka protocol consumption information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyKafkaConsumer(request *ModifyKafkaConsumerRequest) (response *ModifyKafkaConsumerResponse, err error) {
    return c.ModifyKafkaConsumerWithContext(context.Background(), request)
}

// ModifyKafkaConsumer
// This API is used to modify Kafka protocol consumption information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyKafkaConsumerWithContext(ctx context.Context, request *ModifyKafkaConsumerRequest) (response *ModifyKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewModifyKafkaConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyKafkaConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKafkaConsumerGroupOffsetRequest() (request *ModifyKafkaConsumerGroupOffsetRequest) {
    request = &ModifyKafkaConsumerGroupOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaConsumerGroupOffset")
    
    
    return
}

func NewModifyKafkaConsumerGroupOffsetResponse() (response *ModifyKafkaConsumerGroupOffsetResponse) {
    response = &ModifyKafkaConsumerGroupOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKafkaConsumerGroupOffset
// This API is used to modify Kafka protocol consumption group offsets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyKafkaConsumerGroupOffset(request *ModifyKafkaConsumerGroupOffsetRequest) (response *ModifyKafkaConsumerGroupOffsetResponse, err error) {
    return c.ModifyKafkaConsumerGroupOffsetWithContext(context.Background(), request)
}

// ModifyKafkaConsumerGroupOffset
// This API is used to modify Kafka protocol consumption group offsets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyKafkaConsumerGroupOffsetWithContext(ctx context.Context, request *ModifyKafkaConsumerGroupOffsetRequest) (response *ModifyKafkaConsumerGroupOffsetResponse, err error) {
    if request == nil {
        request = NewModifyKafkaConsumerGroupOffsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyKafkaConsumerGroupOffset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKafkaConsumerGroupOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKafkaConsumerGroupOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKafkaRechargeRequest() (request *ModifyKafkaRechargeRequest) {
    request = &ModifyKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaRecharge")
    
    
    return
}

func NewModifyKafkaRechargeResponse() (response *ModifyKafkaRechargeResponse) {
    response = &ModifyKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKafkaRecharge
// This API is used to modify a Kafka data subscription task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyKafkaRecharge(request *ModifyKafkaRechargeRequest) (response *ModifyKafkaRechargeResponse, err error) {
    return c.ModifyKafkaRechargeWithContext(context.Background(), request)
}

// ModifyKafkaRecharge
// This API is used to modify a Kafka data subscription task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyKafkaRechargeWithContext(ctx context.Context, request *ModifyKafkaRechargeRequest) (response *ModifyKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewModifyKafkaRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyKafkaRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogsetRequest() (request *ModifyLogsetRequest) {
    request = &ModifyLogsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyLogset")
    
    
    return
}

func NewModifyLogsetResponse() (response *ModifyLogsetResponse) {
    response = &ModifyLogsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLogset
// This API is used to modify a logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERIODMODIFYFORBIDDEN = "FailedOperation.PeriodModifyForbidden"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) ModifyLogset(request *ModifyLogsetRequest) (response *ModifyLogsetResponse, err error) {
    return c.ModifyLogsetWithContext(context.Background(), request)
}

// ModifyLogset
// This API is used to modify a logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERIODMODIFYFORBIDDEN = "FailedOperation.PeriodModifyForbidden"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) ModifyLogsetWithContext(ctx context.Context, request *ModifyLogsetRequest) (response *ModifyLogsetResponse, err error) {
    if request == nil {
        request = NewModifyLogsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyLogset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLogset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLogsetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMachineGroupRequest() (request *ModifyMachineGroupRequest) {
    request = &ModifyMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyMachineGroup")
    
    
    return
}

func NewModifyMachineGroupResponse() (response *ModifyMachineGroupResponse) {
    response = &ModifyMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMachineGroup
// Modify machine group.
//
// Note: Modifying the interface will directly overwrite historical data and change it to valid input parameters this time. Please be cautious when calling this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ModifyMachineGroup(request *ModifyMachineGroupRequest) (response *ModifyMachineGroupResponse, err error) {
    return c.ModifyMachineGroupWithContext(context.Background(), request)
}

// ModifyMachineGroup
// Modify machine group.
//
// Note: Modifying the interface will directly overwrite historical data and change it to valid input parameters this time. Please be cautious when calling this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ModifyMachineGroupWithContext(ctx context.Context, request *ModifyMachineGroupRequest) (response *ModifyMachineGroupResponse, err error) {
    if request == nil {
        request = NewModifyMachineGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyMachineGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMetricConfigRequest() (request *ModifyMetricConfigRequest) {
    request = &ModifyMetricConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyMetricConfig")
    
    
    return
}

func NewModifyMetricConfigResponse() (response *ModifyMetricConfigResponse) {
    response = &ModifyMetricConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMetricConfig
// This API is used to create metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMetricConfig(request *ModifyMetricConfigRequest) (response *ModifyMetricConfigResponse, err error) {
    return c.ModifyMetricConfigWithContext(context.Background(), request)
}

// ModifyMetricConfig
// This API is used to create metric collection configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMetricConfigWithContext(ctx context.Context, request *ModifyMetricConfigRequest) (response *ModifyMetricConfigResponse, err error) {
    if request == nil {
        request = NewModifyMetricConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyMetricConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMetricConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMetricConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMetricSubscribeRequest() (request *ModifyMetricSubscribeRequest) {
    request = &ModifyMetricSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyMetricSubscribe")
    
    
    return
}

func NewModifyMetricSubscribeResponse() (response *ModifyMetricSubscribeResponse) {
    response = &ModifyMetricSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMetricSubscribe
// This API is used to modify metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMetricSubscribe(request *ModifyMetricSubscribeRequest) (response *ModifyMetricSubscribeResponse, err error) {
    return c.ModifyMetricSubscribeWithContext(context.Background(), request)
}

// ModifyMetricSubscribe
// This API is used to modify metric subscription configurations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMetricSubscribeWithContext(ctx context.Context, request *ModifyMetricSubscribeRequest) (response *ModifyMetricSubscribeResponse, err error) {
    if request == nil {
        request = NewModifyMetricSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyMetricSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMetricSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMetricSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkApplicationRequest() (request *ModifyNetworkApplicationRequest) {
    request = &ModifyNetworkApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyNetworkApplication")
    
    
    return
}

func NewModifyNetworkApplicationResponse() (response *ModifyNetworkApplicationResponse) {
    response = &ModifyNetworkApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNetworkApplication
// Modify a web application
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNetworkApplication(request *ModifyNetworkApplicationRequest) (response *ModifyNetworkApplicationResponse, err error) {
    return c.ModifyNetworkApplicationWithContext(context.Background(), request)
}

// ModifyNetworkApplication
// Modify a web application
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNetworkApplicationWithContext(ctx context.Context, request *ModifyNetworkApplicationRequest) (response *ModifyNetworkApplicationResponse, err error) {
    if request == nil {
        request = NewModifyNetworkApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyNetworkApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNoticeContentRequest() (request *ModifyNoticeContentRequest) {
    request = &ModifyNoticeContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyNoticeContent")
    
    
    return
}

func NewModifyNoticeContentResponse() (response *ModifyNoticeContentResponse) {
    response = &ModifyNoticeContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNoticeContent
// This API is used to modify notification content configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyNoticeContent(request *ModifyNoticeContentRequest) (response *ModifyNoticeContentResponse, err error) {
    return c.ModifyNoticeContentWithContext(context.Background(), request)
}

// ModifyNoticeContent
// This API is used to modify notification content configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyNoticeContentWithContext(ctx context.Context, request *ModifyNoticeContentRequest) (response *ModifyNoticeContentResponse, err error) {
    if request == nil {
        request = NewModifyNoticeContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyNoticeContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNoticeContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNoticeContentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordingRuleTaskRequest() (request *ModifyRecordingRuleTaskRequest) {
    request = &ModifyRecordingRuleTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyRecordingRuleTask")
    
    
    return
}

func NewModifyRecordingRuleTaskResponse() (response *ModifyRecordingRuleTaskResponse) {
    response = &ModifyRecordingRuleTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRecordingRuleTask
// This API is used to modify a scheduled pre-aggregation task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyRecordingRuleTask(request *ModifyRecordingRuleTaskRequest) (response *ModifyRecordingRuleTaskResponse, err error) {
    return c.ModifyRecordingRuleTaskWithContext(context.Background(), request)
}

// ModifyRecordingRuleTask
// This API is used to modify a scheduled pre-aggregation task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyRecordingRuleTaskWithContext(ctx context.Context, request *ModifyRecordingRuleTaskRequest) (response *ModifyRecordingRuleTaskResponse, err error) {
    if request == nil {
        request = NewModifyRecordingRuleTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyRecordingRuleTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordingRuleTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordingRuleTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordingRuleYamlTaskRequest() (request *ModifyRecordingRuleYamlTaskRequest) {
    request = &ModifyRecordingRuleYamlTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyRecordingRuleYamlTask")
    
    
    return
}

func NewModifyRecordingRuleYamlTaskResponse() (response *ModifyRecordingRuleYamlTaskResponse) {
    response = &ModifyRecordingRuleYamlTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRecordingRuleYamlTask
// Modifying a Metric Pre-Aggregation Task Through a YAML File
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyRecordingRuleYamlTask(request *ModifyRecordingRuleYamlTaskRequest) (response *ModifyRecordingRuleYamlTaskResponse, err error) {
    return c.ModifyRecordingRuleYamlTaskWithContext(context.Background(), request)
}

// ModifyRecordingRuleYamlTask
// Modifying a Metric Pre-Aggregation Task Through a YAML File
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyRecordingRuleYamlTaskWithContext(ctx context.Context, request *ModifyRecordingRuleYamlTaskRequest) (response *ModifyRecordingRuleYamlTaskResponse, err error) {
    if request == nil {
        request = NewModifyRecordingRuleYamlTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyRecordingRuleYamlTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordingRuleYamlTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordingRuleYamlTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScheduledSqlRequest() (request *ModifyScheduledSqlRequest) {
    request = &ModifyScheduledSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyScheduledSql")
    
    
    return
}

func NewModifyScheduledSqlResponse() (response *ModifyScheduledSqlResponse) {
    response = &ModifyScheduledSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyScheduledSql
// This API is used to modify a scheduled SQL analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyScheduledSql(request *ModifyScheduledSqlRequest) (response *ModifyScheduledSqlResponse, err error) {
    return c.ModifyScheduledSqlWithContext(context.Background(), request)
}

// ModifyScheduledSql
// This API is used to modify a scheduled SQL analysis task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyScheduledSqlWithContext(ctx context.Context, request *ModifyScheduledSqlRequest) (response *ModifyScheduledSqlResponse, err error) {
    if request == nil {
        request = NewModifyScheduledSqlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyScheduledSql")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyScheduledSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScheduledSqlResponse()
    err = c.Send(request, response)
    return
}

func NewModifySearchViewRequest() (request *ModifySearchViewRequest) {
    request = &ModifySearchViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifySearchView")
    
    
    return
}

func NewModifySearchViewResponse() (response *ModifySearchViewResponse) {
    response = &ModifySearchViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySearchView
// This API is used to modify a query view.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifySearchView(request *ModifySearchViewRequest) (response *ModifySearchViewResponse, err error) {
    return c.ModifySearchViewWithContext(context.Background(), request)
}

// ModifySearchView
// This API is used to modify a query view.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifySearchViewWithContext(ctx context.Context, request *ModifySearchViewRequest) (response *ModifySearchViewResponse, err error) {
    if request == nil {
        request = NewModifySearchViewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifySearchView")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySearchView require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySearchViewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyShipperRequest() (request *ModifyShipperRequest) {
    request = &ModifyShipperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyShipper")
    
    
    return
}

func NewModifyShipperResponse() (response *ModifyShipperResponse) {
    response = &ModifyShipperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyShipper
// This API is used to modify an existing shipping rule. To use this API, you need to grant CLS the write permission of the specified bucket.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) ModifyShipper(request *ModifyShipperRequest) (response *ModifyShipperResponse, err error) {
    return c.ModifyShipperWithContext(context.Background(), request)
}

// ModifyShipper
// This API is used to modify an existing shipping rule. To use this API, you need to grant CLS the write permission of the specified bucket.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) ModifyShipperWithContext(ctx context.Context, request *ModifyShipperRequest) (response *ModifyShipperResponse, err error) {
    if request == nil {
        request = NewModifyShipperRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyShipper")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyShipperResponse()
    err = c.Send(request, response)
    return
}

func NewModifySplunkDeliverRequest() (request *ModifySplunkDeliverRequest) {
    request = &ModifySplunkDeliverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifySplunkDeliver")
    
    
    return
}

func NewModifySplunkDeliverResponse() (response *ModifySplunkDeliverResponse) {
    response = &ModifySplunkDeliverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySplunkDeliver
// Modify information related to splunk delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySplunkDeliver(request *ModifySplunkDeliverRequest) (response *ModifySplunkDeliverResponse, err error) {
    return c.ModifySplunkDeliverWithContext(context.Background(), request)
}

// ModifySplunkDeliver
// Modify information related to splunk delivery task
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySplunkDeliverWithContext(ctx context.Context, request *ModifySplunkDeliverRequest) (response *ModifySplunkDeliverResponse, err error) {
    if request == nil {
        request = NewModifySplunkDeliverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifySplunkDeliver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySplunkDeliver require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySplunkDeliverResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
    request = &ModifyTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyTopic")
    
    
    return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
    response = &ModifyTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTopic
// This API is used to modify logs or metric topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    return c.ModifyTopicWithContext(context.Background(), request)
}

// ModifyTopic
// This API is used to modify logs or metric topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyTopicWithContext(ctx context.Context, request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    if request == nil {
        request = NewModifyTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebCallbackRequest() (request *ModifyWebCallbackRequest) {
    request = &ModifyWebCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyWebCallback")
    
    
    return
}

func NewModifyWebCallbackResponse() (response *ModifyWebCallbackResponse) {
    response = &ModifyWebCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWebCallback
// This API is used to modify alarm channel callback configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyWebCallback(request *ModifyWebCallbackRequest) (response *ModifyWebCallbackResponse, err error) {
    return c.ModifyWebCallbackWithContext(context.Background(), request)
}

// ModifyWebCallback
// This API is used to modify alarm channel callback configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyWebCallbackWithContext(ctx context.Context, request *ModifyWebCallbackRequest) (response *ModifyWebCallbackResponse, err error) {
    if request == nil {
        request = NewModifyWebCallbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "ModifyWebCallback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClawServiceRequest() (request *OpenClawServiceRequest) {
    request = &OpenClawServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "OpenClawService")
    
    
    return
}

func NewOpenClawServiceResponse() (response *OpenClawServiceResponse) {
    response = &OpenClawServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClawService
// This API is used to create resources and indexes dependent on OpenClaw.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXKEYOVERLIMIT = "LimitExceeded.IndexKeyOverLimit"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OpenClawService(request *OpenClawServiceRequest) (response *OpenClawServiceResponse, err error) {
    return c.OpenClawServiceWithContext(context.Background(), request)
}

// OpenClawService
// This API is used to create resources and indexes dependent on OpenClaw.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXKEYOVERLIMIT = "LimitExceeded.IndexKeyOverLimit"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OpenClawServiceWithContext(ctx context.Context, request *OpenClawServiceRequest) (response *OpenClawServiceResponse, err error) {
    if request == nil {
        request = NewOpenClawServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "OpenClawService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClawService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClawServiceResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClsServiceRequest() (request *OpenClsServiceRequest) {
    request = &OpenClsServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "OpenClsService")
    
    
    return
}

func NewOpenClsServiceResponse() (response *OpenClsServiceResponse) {
    response = &OpenClsServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClsService
// Enable logging
//
// This API is used to enable CLS in all regions by filling any region for Region, recommend using Guangzhou (ap-guangzhou).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) OpenClsService(request *OpenClsServiceRequest) (response *OpenClsServiceResponse, err error) {
    return c.OpenClsServiceWithContext(context.Background(), request)
}

// OpenClsService
// Enable logging
//
// This API is used to enable CLS in all regions by filling any region for Region, recommend using Guangzhou (ap-guangzhou).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) OpenClsServiceWithContext(ctx context.Context, request *OpenClsServiceRequest) (response *OpenClsServiceResponse, err error) {
    if request == nil {
        request = NewOpenClsServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "OpenClsService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClsService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClsServiceResponse()
    err = c.Send(request, response)
    return
}

func NewOpenKafkaConsumerRequest() (request *OpenKafkaConsumerRequest) {
    request = &OpenKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "OpenKafkaConsumer")
    
    
    return
}

func NewOpenKafkaConsumerResponse() (response *OpenKafkaConsumerResponse) {
    response = &OpenKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenKafkaConsumer
// This API is used to enable the Kafka consumption feature.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXPORTCONFLICT = "InvalidParameter.ExportConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) OpenKafkaConsumer(request *OpenKafkaConsumerRequest) (response *OpenKafkaConsumerResponse, err error) {
    return c.OpenKafkaConsumerWithContext(context.Background(), request)
}

// OpenKafkaConsumer
// This API is used to enable the Kafka consumption feature.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXPORTCONFLICT = "InvalidParameter.ExportConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) OpenKafkaConsumerWithContext(ctx context.Context, request *OpenKafkaConsumerRequest) (response *OpenKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewOpenKafkaConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "OpenKafkaConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewKafkaRechargeRequest() (request *PreviewKafkaRechargeRequest) {
    request = &PreviewKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "PreviewKafkaRecharge")
    
    
    return
}

func NewPreviewKafkaRechargeResponse() (response *PreviewKafkaRechargeResponse) {
    response = &PreviewKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PreviewKafkaRecharge
// This API is used to preview the logs of Kafka data subscription tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) PreviewKafkaRecharge(request *PreviewKafkaRechargeRequest) (response *PreviewKafkaRechargeResponse, err error) {
    return c.PreviewKafkaRechargeWithContext(context.Background(), request)
}

// PreviewKafkaRecharge
// This API is used to preview the logs of Kafka data subscription tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) PreviewKafkaRechargeWithContext(ctx context.Context, request *PreviewKafkaRechargeRequest) (response *PreviewKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewPreviewKafkaRechargeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "PreviewKafkaRecharge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PreviewKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewPreviewKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMetricRequest() (request *QueryMetricRequest) {
    request = &QueryMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "QueryMetric")
    
    
    return
}

func NewQueryMetricResponse() (response *QueryMetricResponse) {
    response = &QueryMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryMetric
// Query the latest metric value at a specified time.
//
// If there is no metric data in the 5 minutes before that moment, there will be no query result.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) QueryMetric(request *QueryMetricRequest) (response *QueryMetricResponse, err error) {
    return c.QueryMetricWithContext(context.Background(), request)
}

// QueryMetric
// Query the latest metric value at a specified time.
//
// If there is no metric data in the 5 minutes before that moment, there will be no query result.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) QueryMetricWithContext(ctx context.Context, request *QueryMetricRequest) (response *QueryMetricResponse, err error) {
    if request == nil {
        request = NewQueryMetricRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "QueryMetric")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMetric require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMetricResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRangeMetricRequest() (request *QueryRangeMetricRequest) {
    request = &QueryRangeMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "QueryRangeMetric")
    
    
    return
}

func NewQueryRangeMetricResponse() (response *QueryRangeMetricResponse) {
    response = &QueryRangeMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryRangeMetric
// This API is used to query the trend of metrics within a specified time range.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryRangeMetric(request *QueryRangeMetricRequest) (response *QueryRangeMetricResponse, err error) {
    return c.QueryRangeMetricWithContext(context.Background(), request)
}

// QueryRangeMetric
// This API is used to query the trend of metrics within a specified time range.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryRangeMetricWithContext(ctx context.Context, request *QueryRangeMetricRequest) (response *QueryRangeMetricResponse, err error) {
    if request == nil {
        request = NewQueryRangeMetricRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "QueryRangeMetric")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryRangeMetric require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryRangeMetricResponse()
    err = c.Send(request, response)
    return
}

func NewRetryShipperTaskRequest() (request *RetryShipperTaskRequest) {
    request = &RetryShipperTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "RetryShipperTask")
    
    
    return
}

func NewRetryShipperTaskResponse() (response *RetryShipperTaskResponse) {
    response = &RetryShipperTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryShipperTask
// This API is used to retry a failed shipping task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHIPPERTASKNOTTORETRY = "FailedOperation.ShipperTaskNotToRetry"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
//  RESOURCENOTFOUND_SHIPPERTASKNOTEXIST = "ResourceNotFound.ShipperTaskNotExist"
func (c *Client) RetryShipperTask(request *RetryShipperTaskRequest) (response *RetryShipperTaskResponse, err error) {
    return c.RetryShipperTaskWithContext(context.Background(), request)
}

// RetryShipperTask
// This API is used to retry a failed shipping task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHIPPERTASKNOTTORETRY = "FailedOperation.ShipperTaskNotToRetry"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
//  RESOURCENOTFOUND_SHIPPERTASKNOTEXIST = "ResourceNotFound.ShipperTaskNotExist"
func (c *Client) RetryShipperTaskWithContext(ctx context.Context, request *RetryShipperTaskRequest) (response *RetryShipperTaskResponse, err error) {
    if request == nil {
        request = NewRetryShipperTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "RetryShipperTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryShipperTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryShipperTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSearchCosRechargeInfoRequest() (request *SearchCosRechargeInfoRequest) {
    request = &SearchCosRechargeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SearchCosRechargeInfo")
    
    
    return
}

func NewSearchCosRechargeInfoResponse() (response *SearchCosRechargeInfoResponse) {
    response = &SearchCosRechargeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchCosRechargeInfo
// This API is used to preview COS import information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOFILE = "FailedOperation.BucketNoFile"
//  FAILEDOPERATION_DECOMPRESSFILE = "FailedOperation.DecompressFile"
//  FAILEDOPERATION_DOWNLOADFILE = "FailedOperation.DownLoadFile"
//  FAILEDOPERATION_GETLISTFILE = "FailedOperation.GetListFile"
//  FAILEDOPERATION_PREVIEWFILE = "FailedOperation.PreviewFile"
//  FAILEDOPERATION_READFILE = "FailedOperation.ReadFile"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchCosRechargeInfo(request *SearchCosRechargeInfoRequest) (response *SearchCosRechargeInfoResponse, err error) {
    return c.SearchCosRechargeInfoWithContext(context.Background(), request)
}

// SearchCosRechargeInfo
// This API is used to preview COS import information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOFILE = "FailedOperation.BucketNoFile"
//  FAILEDOPERATION_DECOMPRESSFILE = "FailedOperation.DecompressFile"
//  FAILEDOPERATION_DOWNLOADFILE = "FailedOperation.DownLoadFile"
//  FAILEDOPERATION_GETLISTFILE = "FailedOperation.GetListFile"
//  FAILEDOPERATION_PREVIEWFILE = "FailedOperation.PreviewFile"
//  FAILEDOPERATION_READFILE = "FailedOperation.ReadFile"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchCosRechargeInfoWithContext(ctx context.Context, request *SearchCosRechargeInfoRequest) (response *SearchCosRechargeInfoResponse, err error) {
    if request == nil {
        request = NewSearchCosRechargeInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "SearchCosRechargeInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchCosRechargeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchCosRechargeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSearchDashboardSubscribeRequest() (request *SearchDashboardSubscribeRequest) {
    request = &SearchDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SearchDashboardSubscribe")
    
    
    return
}

func NewSearchDashboardSubscribeResponse() (response *SearchDashboardSubscribeResponse) {
    response = &SearchDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchDashboardSubscribe
// This API is used to preview the dashboard subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) SearchDashboardSubscribe(request *SearchDashboardSubscribeRequest) (response *SearchDashboardSubscribeResponse, err error) {
    return c.SearchDashboardSubscribeWithContext(context.Background(), request)
}

// SearchDashboardSubscribe
// This API is used to preview the dashboard subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) SearchDashboardSubscribeWithContext(ctx context.Context, request *SearchDashboardSubscribeRequest) (response *SearchDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewSearchDashboardSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "SearchDashboardSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
    request = &SearchLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SearchLog")
    
    
    return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
    response = &SearchLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchLog
// This API is used to retrieve and analyze logs. Please note the following matters when using this API.
//
// 1. Besides being subject to the default API request rate limit, for a single log topic, the number of concurrent queries cannot exceed 15.
//
// 2. The API's return data packet maximum is 49MB. It is recommended to enable gzip compression (HTTP Request Header Accept-Encoding: gzip).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESOURCES = "LimitExceeded.SearchResources"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  OPERATIONDENIED_OPERATIONNOTSUPPORTINSEARCHLOW = "OperationDenied.OperationNotSupportInSearchLow"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
    return c.SearchLogWithContext(context.Background(), request)
}

// SearchLog
// This API is used to retrieve and analyze logs. Please note the following matters when using this API.
//
// 1. Besides being subject to the default API request rate limit, for a single log topic, the number of concurrent queries cannot exceed 15.
//
// 2. The API's return data packet maximum is 49MB. It is recommended to enable gzip compression (HTTP Request Header Accept-Encoding: gzip).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESOURCES = "LimitExceeded.SearchResources"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  OPERATIONDENIED_OPERATIONNOTSUPPORTINSEARCHLOW = "OperationDenied.OperationNotSupportInSearchLow"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchLogWithContext(ctx context.Context, request *SearchLogRequest) (response *SearchLogResponse, err error) {
    if request == nil {
        request = NewSearchLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "SearchLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchLogResponse()
    err = c.Send(request, response)
    return
}

func NewSendConsumerHeartbeatRequest() (request *SendConsumerHeartbeatRequest) {
    request = &SendConsumerHeartbeatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SendConsumerHeartbeat")
    
    
    return
}

func NewSendConsumerHeartbeatResponse() (response *SendConsumerHeartbeatResponse) {
    response = &SendConsumerHeartbeatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendConsumerHeartbeat
// This API is used to check the heartbeat of a consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) SendConsumerHeartbeat(request *SendConsumerHeartbeatRequest) (response *SendConsumerHeartbeatResponse, err error) {
    return c.SendConsumerHeartbeatWithContext(context.Background(), request)
}

// SendConsumerHeartbeat
// This API is used to check the heartbeat of a consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) SendConsumerHeartbeatWithContext(ctx context.Context, request *SendConsumerHeartbeatRequest) (response *SendConsumerHeartbeatResponse, err error) {
    if request == nil {
        request = NewSendConsumerHeartbeatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "SendConsumerHeartbeat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendConsumerHeartbeat require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendConsumerHeartbeatResponse()
    err = c.Send(request, response)
    return
}

func NewSplitPartitionRequest() (request *SplitPartitionRequest) {
    request = &SplitPartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SplitPartition")
    
    
    return
}

func NewSplitPartitionResponse() (response *SplitPartitionResponse) {
    response = &SplitPartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SplitPartition
// This API is deprecated. If needed, please use the ModifyTopic API to change the number of partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SplitPartition(request *SplitPartitionRequest) (response *SplitPartitionResponse, err error) {
    return c.SplitPartitionWithContext(context.Background(), request)
}

// SplitPartition
// This API is deprecated. If needed, please use the ModifyTopic API to change the number of partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SplitPartitionWithContext(ctx context.Context, request *SplitPartitionRequest) (response *SplitPartitionResponse, err error) {
    if request == nil {
        request = NewSplitPartitionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "SplitPartition")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SplitPartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewSplitPartitionResponse()
    err = c.Send(request, response)
    return
}

func NewUploadLogRequest() (request *UploadLogRequest) {
    request = &UploadLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "UploadLog")
    request.SetContentType("application/octet-stream")
    
    return
}

func NewUploadLogResponse() (response *UploadLogResponse) {
    response = &UploadLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadLog
// ## Notification
//
// To ensure the reliability of your log data and use the log service more efficiently, we recommend that you use the CLS-optimized API to upload structured logs (https://www.tencentcloud.com/document/product/614/16873?from_cn_redirect=1).
//
// 
//
// Meanwhile, we have specially optimized and customized SDKs in multiple languages for this API for you to choose from. The SDK provides unified async sending, resource control, automatic retry, graceful shutdown, perception reporting, and other features to improve the log reporting function. For details, refer to [SDK Collection](https://www.tencentcloud.com/document/product/614/67157?from_cn_redirect=1).
//
// 
//
// Meanwhile, the log upload API also supports synchronous log data upload. If you select to continue using this API, refer to the following text.
//
// 
//
// ## Feature Description
//
// 
//
// This API is used to write logs to the designated log topic.
//
// 
//
// #### Input parameter (pb binary stream, located in body)
//
// 
//
// | Field name | Data type | Location | Must | Description |
//
// | ------------ | ------- | ---- | ---- | ------------------------------------------------------------ |
//
// | logGroupList | message | pb   | Yes | logGroup list, encapsulated content of the log group list. It is advisable not to exceed 5 logGroups. |
//
// 
//
// Group description:
//
// 
//
// | Field name | Required or optional | Description |
//
// | ----------- | -------- | ------------------------------------------------------------ |
//
// | logs        | is       | a log array, which means a set of multiple logs. One Log represents one log, and the number of logs in one LogGroup cannot exceed 10000 |
//
// | contextFlow | No | The unique ID of LogGroup, which must be imported when using context features. Format: "{context ID}-{LogGroupID}".<br>Context ID: A unique identifier for a context (a series of consecutively scrolled log files or a sequence of logs requiring order preservation), a 64-bit integer string in base 16.<br>LogGroupID: A consecutively incremental integer string in base 16. Example: "102700A66102516A-59F59".
//
// | filename    | No       | Log file name |
//
// | source      | No       | Log source, using machine IP as a label in general use       |
//
// | logTags     | No       | Log tag list                                               |
//
// 
//
// Log description:
//
// 
//
// | field name | Required or optional | Description |
//
// | -------- | -------- | ------------------------------------------------------------ |
//
// | time     | is       | log time (Unix timestamp), supports second, millisecond, microsecond, milliseconds is recommended |
//
// | contents | No | Key-value formatted log content, representing multiple key-value composites in a log |
//
// 
//
// Content description:
//
// 
//
// | Field name | Required or optional | Description |
//
// | ------ | -------- | ------------------------------------------------------------ |
//
// | key    | Yes       | The key value of a field group in a single-line log. It cannot start with `_` |
//
// | value  | Yes       | The value of a field group in a single-line log. The value of a single-line log must not exceed 1MB, and the sum of ALL values in a LogGroup cannot exceed 5MB. |
//
// 
//
// LogTag description:
//
// 
//
// | Field name | Required or optional | Description |
//
// | ------ | -------- | -------------------------------- |
//
// | key    | Yes      | Custom tag key                 |
//
// | value  | is       | value corresponding to the custom tag key |
//
// 
//
// ## PB Compilation Example
//
// 
//
// This example shows how to use the official protoc compiler to compile and generate a C++ language adjustable log upload API from a description file.
//
// 
//
// Currently protoc officially supports compilation for languages such as Java, C++, and Python. For details, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// #### 1. Protocol Buffer installation
//
// 
//
// Download [Protocol Buffer](https://main.qcloudimg.com/raw/d7810aaf8b3073fbbc9d4049c21532aa/protobuf-2.6.1.tar.gz), unzip and install. The example version is protobuf 2.6.1, and the environment is Centos 7.3 system. Decompress the `protobuf-2.6.1.tar.gz` compressed package to the `/usr/local` directory and enter the directory. Execute the command as follows:
//
// 
//
// ```
//
// tar -zxvf protobuf-2.6.1.tar.gz -C /usr/local/ && cd /usr/local/protobuf-2.6.1
//
// ```
//
// 
//
// Start compilation and installation, configure environment variables, execute the command as follows:
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# ./configure 
//
// [root@VM_0_8_centos protobuf-2.6.1]# make && make install
//
// [root@VM_0_8_centos protobuf-2.6.1]# export PATH=$PATH:/usr/local/protobuf-2.6.1/bin
//
// ```
//
// 
//
// After successful compilation, view the version using the following command:
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --version
//
// liprotoc 2.6.1
//
// ```
//
// 
//
// #### 2. Create PB description file
//
// 
//
// The PB description file is the data interchange format agreed by the communication parties. When uploading logs, compile the specified protocol format into the calling interface of the corresponding language version, then add to engineering code. For details, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// Create a local PB message description file cls.proto based on the PB data format specified by the log service.
//
// 
//
// !PB description file content immutable, and the file name must end with `.proto`.
//
// 
//
// The content of cls.proto (PB description file) is as follows:
//
// 
//
// ```
//
// package cls;
//
// 
//
// message Log
//
// {
//
//     message Content
//
//     {
//
// required string key = 1; // key for each group of fields
//
// required string value = 2; // The value of the group field
//
//     }
//
// required int64   time     = 1; // A timestamp in UNIX time format
//
// repeated Content contents = 2; // multiple kv combinations in a log
//
// }
//
// 
//
// message LogTag
//
// {
//
//     required string key       = 1;
//
//     required string value     = 2;
//
// }
//
// 
//
// message LogGroup
//
// {
//
// repeated Log    logs        = 1; // log array composed of multiple logs
//
// optional string contextFlow = 2; // Currently no utility
//
// optional string filename = 3; // log file name
//
// optional string source      = 4; // log source, general use machine IP
//
//     repeated LogTag logTags     = 5;
//
// }
//
// 
//
// message LogGroupList
//
// {
//
// repeated LogGroup logGroupList = 1; // log group list
//
// }
//
// ```
//
// 
//
// #### 3. Compile and generate
//
// 
//
// In this example, use the proto compiler to generate C++ language files under the same directory as the cls.proto file and execute the following compilation command:
//
// 
//
// ```
//
// protoc --cpp_out=./ ./cls.proto 
//
// ```
//
// 
//
// `--cpp_out=./` means compile to cpp format and output in the current directory. `./cls.proto` refers to the cls.proto description file located in the current directory.
//
// 
//
// After successful compilation, it will output the code file in the corresponding language. This routine generates the cls.pb.h header file and the cls.pb.cc code implementation file, as shown below:
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --cpp_out=./ ./cls.proto
//
// [root@VM_0_8_centos protobuf-2.6.1]# ls
//
// cls.pb.cc cls.pb.h cls.proto
//
// ```
//
// 
//
// #### 4. Call
//
// 
//
// Import the generated cls.pb.h header file into the code and call the interface to encapsulate the data format.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MISSINGCONTENT = "FailedOperation.MissingContent"
//  FAILEDOPERATION_READQPSLIMIT = "FailedOperation.ReadQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  FAILEDOPERATION_WRITEQPSLIMIT = "FailedOperation.WriteQpsLimit"
//  FAILEDOPERATION_WRITETRAFFICLIMIT = "FailedOperation.WriteTrafficLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  LIMITEXCEEDED_LOGSIZE = "LimitExceeded.LogSize"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadLog(request *UploadLogRequest, data []byte) (response *UploadLogResponse, err error) {
    return c.UploadLogWithContext(context.Background(), request, data)
}

// UploadLog
// ## Notification
//
// To ensure the reliability of your log data and use the log service more efficiently, we recommend that you use the CLS-optimized API to upload structured logs (https://www.tencentcloud.com/document/product/614/16873?from_cn_redirect=1).
//
// 
//
// Meanwhile, we have specially optimized and customized SDKs in multiple languages for this API for you to choose from. The SDK provides unified async sending, resource control, automatic retry, graceful shutdown, perception reporting, and other features to improve the log reporting function. For details, refer to [SDK Collection](https://www.tencentcloud.com/document/product/614/67157?from_cn_redirect=1).
//
// 
//
// Meanwhile, the log upload API also supports synchronous log data upload. If you select to continue using this API, refer to the following text.
//
// 
//
// ## Feature Description
//
// 
//
// This API is used to write logs to the designated log topic.
//
// 
//
// #### Input parameter (pb binary stream, located in body)
//
// 
//
// | Field name | Data type | Location | Must | Description |
//
// | ------------ | ------- | ---- | ---- | ------------------------------------------------------------ |
//
// | logGroupList | message | pb   | Yes | logGroup list, encapsulated content of the log group list. It is advisable not to exceed 5 logGroups. |
//
// 
//
// Group description:
//
// 
//
// | Field name | Required or optional | Description |
//
// | ----------- | -------- | ------------------------------------------------------------ |
//
// | logs        | is       | a log array, which means a set of multiple logs. One Log represents one log, and the number of logs in one LogGroup cannot exceed 10000 |
//
// | contextFlow | No | The unique ID of LogGroup, which must be imported when using context features. Format: "{context ID}-{LogGroupID}".<br>Context ID: A unique identifier for a context (a series of consecutively scrolled log files or a sequence of logs requiring order preservation), a 64-bit integer string in base 16.<br>LogGroupID: A consecutively incremental integer string in base 16. Example: "102700A66102516A-59F59".
//
// | filename    | No       | Log file name |
//
// | source      | No       | Log source, using machine IP as a label in general use       |
//
// | logTags     | No       | Log tag list                                               |
//
// 
//
// Log description:
//
// 
//
// | field name | Required or optional | Description |
//
// | -------- | -------- | ------------------------------------------------------------ |
//
// | time     | is       | log time (Unix timestamp), supports second, millisecond, microsecond, milliseconds is recommended |
//
// | contents | No | Key-value formatted log content, representing multiple key-value composites in a log |
//
// 
//
// Content description:
//
// 
//
// | Field name | Required or optional | Description |
//
// | ------ | -------- | ------------------------------------------------------------ |
//
// | key    | Yes       | The key value of a field group in a single-line log. It cannot start with `_` |
//
// | value  | Yes       | The value of a field group in a single-line log. The value of a single-line log must not exceed 1MB, and the sum of ALL values in a LogGroup cannot exceed 5MB. |
//
// 
//
// LogTag description:
//
// 
//
// | Field name | Required or optional | Description |
//
// | ------ | -------- | -------------------------------- |
//
// | key    | Yes      | Custom tag key                 |
//
// | value  | is       | value corresponding to the custom tag key |
//
// 
//
// ## PB Compilation Example
//
// 
//
// This example shows how to use the official protoc compiler to compile and generate a C++ language adjustable log upload API from a description file.
//
// 
//
// Currently protoc officially supports compilation for languages such as Java, C++, and Python. For details, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// #### 1. Protocol Buffer installation
//
// 
//
// Download [Protocol Buffer](https://main.qcloudimg.com/raw/d7810aaf8b3073fbbc9d4049c21532aa/protobuf-2.6.1.tar.gz), unzip and install. The example version is protobuf 2.6.1, and the environment is Centos 7.3 system. Decompress the `protobuf-2.6.1.tar.gz` compressed package to the `/usr/local` directory and enter the directory. Execute the command as follows:
//
// 
//
// ```
//
// tar -zxvf protobuf-2.6.1.tar.gz -C /usr/local/ && cd /usr/local/protobuf-2.6.1
//
// ```
//
// 
//
// Start compilation and installation, configure environment variables, execute the command as follows:
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# ./configure 
//
// [root@VM_0_8_centos protobuf-2.6.1]# make && make install
//
// [root@VM_0_8_centos protobuf-2.6.1]# export PATH=$PATH:/usr/local/protobuf-2.6.1/bin
//
// ```
//
// 
//
// After successful compilation, view the version using the following command:
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --version
//
// liprotoc 2.6.1
//
// ```
//
// 
//
// #### 2. Create PB description file
//
// 
//
// The PB description file is the data interchange format agreed by the communication parties. When uploading logs, compile the specified protocol format into the calling interface of the corresponding language version, then add to engineering code. For details, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// Create a local PB message description file cls.proto based on the PB data format specified by the log service.
//
// 
//
// !PB description file content immutable, and the file name must end with `.proto`.
//
// 
//
// The content of cls.proto (PB description file) is as follows:
//
// 
//
// ```
//
// package cls;
//
// 
//
// message Log
//
// {
//
//     message Content
//
//     {
//
// required string key = 1; // key for each group of fields
//
// required string value = 2; // The value of the group field
//
//     }
//
// required int64   time     = 1; // A timestamp in UNIX time format
//
// repeated Content contents = 2; // multiple kv combinations in a log
//
// }
//
// 
//
// message LogTag
//
// {
//
//     required string key       = 1;
//
//     required string value     = 2;
//
// }
//
// 
//
// message LogGroup
//
// {
//
// repeated Log    logs        = 1; // log array composed of multiple logs
//
// optional string contextFlow = 2; // Currently no utility
//
// optional string filename = 3; // log file name
//
// optional string source      = 4; // log source, general use machine IP
//
//     repeated LogTag logTags     = 5;
//
// }
//
// 
//
// message LogGroupList
//
// {
//
// repeated LogGroup logGroupList = 1; // log group list
//
// }
//
// ```
//
// 
//
// #### 3. Compile and generate
//
// 
//
// In this example, use the proto compiler to generate C++ language files under the same directory as the cls.proto file and execute the following compilation command:
//
// 
//
// ```
//
// protoc --cpp_out=./ ./cls.proto 
//
// ```
//
// 
//
// `--cpp_out=./` means compile to cpp format and output in the current directory. `./cls.proto` refers to the cls.proto description file located in the current directory.
//
// 
//
// After successful compilation, it will output the code file in the corresponding language. This routine generates the cls.pb.h header file and the cls.pb.cc code implementation file, as shown below:
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --cpp_out=./ ./cls.proto
//
// [root@VM_0_8_centos protobuf-2.6.1]# ls
//
// cls.pb.cc cls.pb.h cls.proto
//
// ```
//
// 
//
// #### 4. Call
//
// 
//
// Import the generated cls.pb.h header file into the code and call the interface to encapsulate the data format.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MISSINGCONTENT = "FailedOperation.MissingContent"
//  FAILEDOPERATION_READQPSLIMIT = "FailedOperation.ReadQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  FAILEDOPERATION_WRITEQPSLIMIT = "FailedOperation.WriteQpsLimit"
//  FAILEDOPERATION_WRITETRAFFICLIMIT = "FailedOperation.WriteTrafficLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  LIMITEXCEEDED_LOGSIZE = "LimitExceeded.LogSize"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadLogWithContext(ctx context.Context, request *UploadLogRequest, data []byte) (response *UploadLogResponse, err error) {
    if request == nil {
        request = NewUploadLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cls", APIVersion, "UploadLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadLog require credential")
    }

    request.SetContext(ctx)
    request.SetBody(data)
    response = NewUploadLogResponse()
    err = c.SendOctetStream(request, response)
    return
}
