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
// This API is used to apply the collection configuration to the specified machine group.
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
// This API is used to apply the collection configuration to the specified machine group.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyConfigToMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyConfigToMachineGroupResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseKafkaConsumerResponse()
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
// This API is used to create a notification group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
// This API is used to create a notification group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmNoticeResponse()
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
// This API is used to create a collection rule configuration.
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
// This API is used to create a collection rule configuration.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigResponse()
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
// This API is used to create a shipping task.
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
// This API is used to create a shipping task.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerResponse()
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
func (c *Client) CreateCosRecharge(request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
    return c.CreateCosRechargeWithContext(context.Background(), request)
}

// CreateCosRecharge
// This API is used to create a COS import task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) CreateCosRechargeWithContext(ctx context.Context, request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
    if request == nil {
        request = NewCreateCosRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosRechargeResponse()
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
func (c *Client) CreateDataTransform(request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
    return c.CreateDataTransformWithContext(context.Background(), request)
}

// CreateDataTransform
// This API is used to create a data processing task.
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
func (c *Client) CreateDataTransformWithContext(ctx context.Context, request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
    if request == nil {
        request = NewCreateDataTransformRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataTransformResponse()
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
// This API is used to create a download task. To get the returned download address, call `DescribeExports` to view the task list. The `CosPath` parameter is also included for download address. For more information, visit https://intl.cloud.tencent.com/document/product/614/56449.?from_cn_redirect=1
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
// This API is used to create a download task. To get the returned download address, call `DescribeExports` to view the task list. The `CosPath` parameter is also included for download address. For more information, visit https://intl.cloud.tencent.com/document/product/614/56449.?from_cn_redirect=1
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExportResponse()
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
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
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
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndexWithContext(ctx context.Context, request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    if request == nil {
        request = NewCreateIndexRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMachineGroupResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateShipperResponse()
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
// This API is used to create a log topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TOPICCREATING = "FailedOperation.TopicCreating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
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
// This API is used to create a log topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TOPICCREATING = "FailedOperation.TopicCreating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmNoticeResponse()
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
// This API is used to delete a collection rule configuration.
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
// This API is used to delete a collection rule configuration.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigFromMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigFromMachineGroupResponse()
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
// This API is used to delete a shipping task.
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
// This API is used to delete a shipping task.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataTransformResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExportResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineGroupInfoResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShipperResponse()
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
// This API is used to delete a log topic.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
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
// This API is used to delete a log topic.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticesResponse()
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
func (c *Client) DescribeAlarmsWithContext(ctx context.Context, request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmsRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlertRecordHistoryWithContext(ctx context.Context, request *DescribeAlertRecordHistoryRequest) (response *DescribeAlertRecordHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeAlertRecordHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertRecordHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertRecordHistoryResponse()
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
// This API is used to get the machine group bound to a collection rule configuration.
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
// This API is used to get the machine group bound to a collection rule configuration.
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
// This API is used to get a collection rule configuration.
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
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    return c.DescribeConfigsWithContext(context.Background(), request)
}

// DescribeConfigs
// This API is used to get a collection rule configuration.
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
func (c *Client) DescribeConfigsWithContext(ctx context.Context, request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRecharges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRechargesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataTransformInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataTransformInfoResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexResponse()
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
// This API is used to search for content in the log context.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
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
// This API is used to search for content in the log context.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
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
// This API is used to get the machine status in the specified machine group.
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
// This API is used to get the machine status in the specified machine group.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachinesResponse()
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
// This API is used to get the list of topic partitions.
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
// This API is used to get the list of topic partitions.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePartitions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePartitionsResponse()
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
// This API is used to get the configuration of the task of shipping to COS.
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
// This API is used to get the configuration of the task of shipping to COS.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShippers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShippersResponse()
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
// This API is used to get the list of log topics and supports pagination.
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
//  OPERATIONDENIED_ANALYSISSWITCHCLOSE = "OperationDenied.AnalysisSwitchClose"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    return c.DescribeTopicsWithContext(context.Background(), request)
}

// DescribeTopics
// This API is used to get the list of log topics and supports pagination.
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
//  OPERATIONDENIED_ANALYSISSWITCHCLOSE = "OperationDenied.AnalysisSwitchClose"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicsWithContext(ctx context.Context, request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicsResponse()
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
// This API is used to get the records of alarm tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
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
// This API is used to get the records of alarm tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAlarmLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAlarmLogResponse()
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
// This API is used to merge a topic partition in read/write state. During merge, a topic partition ID can be specified, and CLS will automatically merge the partition adjacent to the right of the range.
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
// This API is used to merge a topic partition in read/write state. During merge, a topic partition ID can be specified, and CLS will automatically merge the partition adjacent to the right of the range.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmNoticeResponse()
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
// This API is used to modify a collection rule configuration.
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
// This API is used to modify a collection rule configuration.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigResponse()
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
// This API is used to modify a shipping task.
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
func (c *Client) ModifyConsumer(request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
    return c.ModifyConsumerWithContext(context.Background(), request)
}

// ModifyConsumer
// This API is used to modify a shipping task.
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
func (c *Client) ModifyConsumerWithContext(ctx context.Context, request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
    if request == nil {
        request = NewModifyConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerResponse()
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
func (c *Client) ModifyCosRecharge(request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
    return c.ModifyCosRechargeWithContext(context.Background(), request)
}

// ModifyCosRecharge
// This API is used to modify a COS import task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ModifyCosRechargeWithContext(ctx context.Context, request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
    if request == nil {
        request = NewModifyCosRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCosRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCosRechargeResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataTransformResponse()
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
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
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
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIndexResponse()
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
// This API is used to modify a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
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
// This API is used to modify a machine group.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMachineGroupResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyShipperResponse()
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
// This API is used to modify a log topic.
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    return c.ModifyTopicWithContext(context.Background(), request)
}

// ModifyTopic
// This API is used to modify a log topic.
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyTopicWithContext(ctx context.Context, request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    if request == nil {
        request = NewModifyTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("PreviewKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewPreviewKafkaRechargeResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryShipperTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryShipperTaskResponse()
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
// This API is used to search logs. It is subject to the default API rate limit, and the number of concurrent queries to the same log topic cannot exceed 15.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
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
// This API is used to search logs. It is subject to the default API rate limit, and the number of concurrent queries to the same log topic cannot exceed 15.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchLogResponse()
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
// This API is used to split a topic partition.
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
// This API is used to split a topic partition.
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
// ## Note
//
// To ensure log data reliability and help you use CLS more efficiently, we recommend you use the optimized API to upload logs. For more information about the API, see [Uploading Log via API](https://intl.cloud.tencent.com/document/product/614/16873?from_cn_redirect=1).
//
// 
//
// For the optimized API, we have developed an SDK (available in multiple languages) that provides features including async sending, resource control, automatic retry, graceful shutdown, and detection-based reporting. For details, see [Uploading Log via SDK](https://intl.cloud.tencent.com/document/product/614/67157?from_cn_redirect=1).
//
// 
//
// `UploadLog` allows you to synchronously upload log data. If you still want to continue to use this API instead of the optimized one, read this document.
//
// 
//
// ## Feature Description
//
// 
//
// This API is used to write logs to a specified log topic.
//
// 
//
// CLS provides the following two modes:
//
// 
//
// #### Load balancing mode
//
// 
//
// In this mode, logs will be automatically written to a target partition among all readable/writable partitions under the current log topic based on the load balancing principle. This mode is suitable for scenarios where sequential consumption is not needed.
//
// 
//
// #### Hash routing mode
//
// 
//
// In this mode, data will be written to a target partition that meets the range requirements based on the carried hash value (`X-CLS-HashKey`). For example, a log source can be bound to a topic partition through `HashKey`, strictly guaranteeing the sequence of the data written to and consumed in this partition.
//
// 
//
//                  
//
// 
//
// #### Input parameters (pb binary streams in `body`)
//
// 
//
// | Parameter       | Type    | Location | Required | Description                                                         |
//
// | ------------ | ------- | ---- | ---- | ------------------------------------------------------------ |
//
// | logGroupList | message | pb   | Yes   | The `logGroup` list, which describes the encapsulated log groups. We recommend you enter up to five `logGroup` values. |
//
// 
//
// `LogGroup` description:
//
// 
//
// | Parameter      | Required | Description                                                         |
//
// | ----------- | -------- | ------------------------------------------------------------ |
//
// | logs        | Yes       | Log array consisting of multiple `Log` values. The `Log` indicates a log, and a `LogGroup` can contain up to 10,000 `Log` values. |
//
// | contextFlow | No       | Unique `LogGroup` ID, which should be passed in if the context feature needs to be used. Format: "{Context ID}-{LogGroupID}". <br>Context ID: Uniquely identifies the context (a series of log files that are continuously scrolling or a series of logs that need to be sequenced), which is a 64-bit integer hex string. <br>LogGroupID: A 64-bit integer hex string that continuously increases, such as `102700A66102516A-59F59`.                        |
//
// | filename    | No       | Log filename                                                   |
//
// | source      | No       | Log source, which is generally the machine IP                           |
//
// | logTags     | No       | List of log tags                                               |
//
// 
//
// `Log` description:
//
// 
//
// | Parameter   | Required | Description                                                         |
//
// | -------- | -------- | ------------------------------------------------------------ |
//
// | time     | Yes       | Unix timestamp of log time in seconds or milliseconds (recommended)      |
//
// | contents | No       | Log content in key-value format. A log can contain multiple key-value pairs. |
//
// 
//
// `Content` description:
//
// 
//
// | Parameter | Required | Description                                                         |
//
// | ------ | -------- | ------------------------------------------------------------ |
//
// | key    | Yes       | Key of a field group in one log, which cannot start with `_`.                 |
//
// | value  | Yes       | Value of a field group. The `value` of one log cannot exceed 1 MB and the total `value` in `LogGroup` cannot exceed 5 MB. |
//
// 
//
// `LogTag` description:
//
// 
//
// | Parameter | Required | Description                             |
//
// | ------ | -------- | -------------------------------- |
//
// | key    | Yes       | Key of a custom tag                 |
//
// | value  | Yes       | Value corresponding to the custom tag key |
//
// 
//
// ## pb Compilation Example
//
// 
//
// This example shows you how to use the protoc compiler to compile a pb description file into a log upload API in C++.
//
// 
//
// > ?Currently, protoc supports compilation in multiple programming languages such as Java, C++, and Python. For more information, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// #### 1. Install Protocol Buffers
//
// 
//
// Download [Protocol Buffers](https://main.qcloudimg.com/raw/d7810aaf8b3073fbbc9d4049c21532aa/protobuf-2.6.1.tar.gz), decompress the package, and install the tool. The version used in the example is protobuf 2.6.1 running on CentOS 7.3. Run the following command to decompress the `protobuf-2.6.1.tar.gz` package to the `/usr/local` directory and go to the directory:
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
// Run the following commands to start compilation and installation and configure the environment variables:
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
// After the compilation succeeds, run the following command to check the version:
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
// #### 2. Create a pb description file
//
// 
//
// A pb description file is an agreed-on data interchange format for communication. To upload logs, compile the specified protocol format to an API in the target programming language and add the API to the project code. For more information, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// Create a pb message description file `cls.proto` based on the pb data format content specified by CLS.
//
// 
//
// > !The pb description file content cannot be modified, and the filename must end with `.proto`.
//
// 
//
// The content of `cls.proto` (pb description file) is as follows:
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
//         required string key   = 1; // Key of each field group
//
//         required string value = 2; // Value of each field group
//
//     }
//
//     required int64   time     = 1; // Unix timestamp
//
//     repeated Content contents = 2; // Multiple key-value pairs in one log
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
//     repeated Log    logs        = 1; // Log array consisting of multiple logs
//
//     optional string contextFlow = 2; // This parameter does not take effect currently
//
//     optional string filename    = 3; // Log filename
//
//     optional string source      = 4; // Log source, which is generally the machine IP
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
//     repeated LogGroup logGroupList = 1; // Log group list
//
// }
//
// ```
//
// 
//
// #### 3. Compile and generate the API
//
// 
//
// This example uses the proto compiler to generate a C++ file in the same directory as the `cls.proto` file. Run the following compilation command:
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
// > ?`--cpp_out=./` indicates that the file will be compiled in cpp format and output to the current directory. `./cls.proto` indicates the `cls.proto` description file in the current directory.
//
// 
//
// After the compilation succeeds, the code file in the corresponding programming language will be generated. This example generates the `cls.pb.h` header file and [cls.pb.cc](http://cls.pb.cc) code implementation file as shown below:
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
// #### 4. Call the API
//
// 
//
// Import the generated `cls.pb.h` header file into the code and call the API for data encapsulation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MISSINGCONTENT = "FailedOperation.MissingContent"
//  FAILEDOPERATION_READQPSLIMIT = "FailedOperation.ReadQpsLimit"
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
// ## Note
//
// To ensure log data reliability and help you use CLS more efficiently, we recommend you use the optimized API to upload logs. For more information about the API, see [Uploading Log via API](https://intl.cloud.tencent.com/document/product/614/16873?from_cn_redirect=1).
//
// 
//
// For the optimized API, we have developed an SDK (available in multiple languages) that provides features including async sending, resource control, automatic retry, graceful shutdown, and detection-based reporting. For details, see [Uploading Log via SDK](https://intl.cloud.tencent.com/document/product/614/67157?from_cn_redirect=1).
//
// 
//
// `UploadLog` allows you to synchronously upload log data. If you still want to continue to use this API instead of the optimized one, read this document.
//
// 
//
// ## Feature Description
//
// 
//
// This API is used to write logs to a specified log topic.
//
// 
//
// CLS provides the following two modes:
//
// 
//
// #### Load balancing mode
//
// 
//
// In this mode, logs will be automatically written to a target partition among all readable/writable partitions under the current log topic based on the load balancing principle. This mode is suitable for scenarios where sequential consumption is not needed.
//
// 
//
// #### Hash routing mode
//
// 
//
// In this mode, data will be written to a target partition that meets the range requirements based on the carried hash value (`X-CLS-HashKey`). For example, a log source can be bound to a topic partition through `HashKey`, strictly guaranteeing the sequence of the data written to and consumed in this partition.
//
// 
//
//                  
//
// 
//
// #### Input parameters (pb binary streams in `body`)
//
// 
//
// | Parameter       | Type    | Location | Required | Description                                                         |
//
// | ------------ | ------- | ---- | ---- | ------------------------------------------------------------ |
//
// | logGroupList | message | pb   | Yes   | The `logGroup` list, which describes the encapsulated log groups. We recommend you enter up to five `logGroup` values. |
//
// 
//
// `LogGroup` description:
//
// 
//
// | Parameter      | Required | Description                                                         |
//
// | ----------- | -------- | ------------------------------------------------------------ |
//
// | logs        | Yes       | Log array consisting of multiple `Log` values. The `Log` indicates a log, and a `LogGroup` can contain up to 10,000 `Log` values. |
//
// | contextFlow | No       | Unique `LogGroup` ID, which should be passed in if the context feature needs to be used. Format: "{Context ID}-{LogGroupID}". <br>Context ID: Uniquely identifies the context (a series of log files that are continuously scrolling or a series of logs that need to be sequenced), which is a 64-bit integer hex string. <br>LogGroupID: A 64-bit integer hex string that continuously increases, such as `102700A66102516A-59F59`.                        |
//
// | filename    | No       | Log filename                                                   |
//
// | source      | No       | Log source, which is generally the machine IP                           |
//
// | logTags     | No       | List of log tags                                               |
//
// 
//
// `Log` description:
//
// 
//
// | Parameter   | Required | Description                                                         |
//
// | -------- | -------- | ------------------------------------------------------------ |
//
// | time     | Yes       | Unix timestamp of log time in seconds or milliseconds (recommended)      |
//
// | contents | No       | Log content in key-value format. A log can contain multiple key-value pairs. |
//
// 
//
// `Content` description:
//
// 
//
// | Parameter | Required | Description                                                         |
//
// | ------ | -------- | ------------------------------------------------------------ |
//
// | key    | Yes       | Key of a field group in one log, which cannot start with `_`.                 |
//
// | value  | Yes       | Value of a field group. The `value` of one log cannot exceed 1 MB and the total `value` in `LogGroup` cannot exceed 5 MB. |
//
// 
//
// `LogTag` description:
//
// 
//
// | Parameter | Required | Description                             |
//
// | ------ | -------- | -------------------------------- |
//
// | key    | Yes       | Key of a custom tag                 |
//
// | value  | Yes       | Value corresponding to the custom tag key |
//
// 
//
// ## pb Compilation Example
//
// 
//
// This example shows you how to use the protoc compiler to compile a pb description file into a log upload API in C++.
//
// 
//
// > ?Currently, protoc supports compilation in multiple programming languages such as Java, C++, and Python. For more information, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// #### 1. Install Protocol Buffers
//
// 
//
// Download [Protocol Buffers](https://main.qcloudimg.com/raw/d7810aaf8b3073fbbc9d4049c21532aa/protobuf-2.6.1.tar.gz), decompress the package, and install the tool. The version used in the example is protobuf 2.6.1 running on CentOS 7.3. Run the following command to decompress the `protobuf-2.6.1.tar.gz` package to the `/usr/local` directory and go to the directory:
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
// Run the following commands to start compilation and installation and configure the environment variables:
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
// After the compilation succeeds, run the following command to check the version:
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
// #### 2. Create a pb description file
//
// 
//
// A pb description file is an agreed-on data interchange format for communication. To upload logs, compile the specified protocol format to an API in the target programming language and add the API to the project code. For more information, see [protoc](https://github.com/protocolbuffers/protobuf).
//
// 
//
// Create a pb message description file `cls.proto` based on the pb data format content specified by CLS.
//
// 
//
// > !The pb description file content cannot be modified, and the filename must end with `.proto`.
//
// 
//
// The content of `cls.proto` (pb description file) is as follows:
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
//         required string key   = 1; // Key of each field group
//
//         required string value = 2; // Value of each field group
//
//     }
//
//     required int64   time     = 1; // Unix timestamp
//
//     repeated Content contents = 2; // Multiple key-value pairs in one log
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
//     repeated Log    logs        = 1; // Log array consisting of multiple logs
//
//     optional string contextFlow = 2; // This parameter does not take effect currently
//
//     optional string filename    = 3; // Log filename
//
//     optional string source      = 4; // Log source, which is generally the machine IP
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
//     repeated LogGroup logGroupList = 1; // Log group list
//
// }
//
// ```
//
// 
//
// #### 3. Compile and generate the API
//
// 
//
// This example uses the proto compiler to generate a C++ file in the same directory as the `cls.proto` file. Run the following compilation command:
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
// > ?`--cpp_out=./` indicates that the file will be compiled in cpp format and output to the current directory. `./cls.proto` indicates the `cls.proto` description file in the current directory.
//
// 
//
// After the compilation succeeds, the code file in the corresponding programming language will be generated. This example generates the `cls.pb.h` header file and [cls.pb.cc](http://cls.pb.cc) code implementation file as shown below:
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
// #### 4. Call the API
//
// 
//
// Import the generated `cls.pb.h` header file into the code and call the API for data encapsulation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MISSINGCONTENT = "FailedOperation.MissingContent"
//  FAILEDOPERATION_READQPSLIMIT = "FailedOperation.ReadQpsLimit"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadLog require credential")
    }

    request.SetContext(ctx)
    request.SetBody(data)
    response = NewUploadLogResponse()
    err = c.SendOctetStream(request, response)
    return
}
